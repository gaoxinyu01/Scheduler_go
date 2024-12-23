package logic

import (
	"Scheduler_go/common/workflow/engine"
	"Scheduler_go/service/workflow/model"
	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"reflect"
	"strings"
	"time"

	. "Scheduler_go/common/workflow/modelx"
	"github.com/zeromicro/go-zero/core/logx"
)

type ProcInstStartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcInstStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcInstStartLogic {
	return &ProcInstStartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 流程实例
func (l *ProcInstStartLogic) ProcInstStart(in *workflowclient.ProcInstStartReq) (resp *workflowclient.CommonResp, err error) {

	//实例初始化
	_, _, err = l.instanceInit(in.ProcId, in.BusinessId, in.VariablesJson)
	if err != nil {
		return nil, err
	}

	return &workflowclient.CommonResp{}, nil
}

// 1、流程实例初始化 2、保存实例变量 返回:流程实例ID、开始节点
func (l *ProcInstStartLogic) instanceInit(ProcessID int64, BusinessID string, VariableJson string) (int64, Node, error) {
	//获取流程定义(流程中所有node)
	nodes, err := l.GetProcCache(ProcessID)
	if err != nil {
		return 0, Node{}, err
	}

	//检查流程节点中的事件是否都已经注册
	err = l.VerifyEvents(ProcessID, nodes)
	if err != nil {
		return 0, Node{}, err
	}

	//获取流程开始节点ID
	type result struct {
		Node_ID string
	}
	var r result

	_, err = l.svcCtx.ProcExecutionModel.FindOneByProcIdAndNodeType(l.ctx, ProcessID, 0)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return 0, Node{}, fmt.Errorf("ProcExecution没有该ID：%v", ProcessID)
		}
		return 0, Node{}, err
	}
	if r.Node_ID == "" {
		return 0, Node{}, fmt.Errorf("无法获取流程ID为%d的开始节点", ProcessID)
	}

	//获得开始节点
	StartNode := nodes[r.Node_ID]

	//-----------------------------------开始处理数据-----------------------------------
	//1、在proc_inst表中生成一条记录
	//2、在proc_inst_variable表中记录流程实例的变量
	//3、返回proc_inst_id(流程实例ID)

	//获取流程定义信息
	res, err := l.svcCtx.ProcDefModel.FindOne(l.ctx, ProcessID)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return 0, Node{}, fmt.Errorf("ProcDef没有该ID：%v", ProcessID)
		}
		return 0, Node{}, err
	}

	// 开启事务添加
	err = l.svcCtx.ProcInstModel.TransCtx(l.ctx, func(ctx context.Context, sqlx sqlx.Session) error {
		//添加流程实例
		procInst, err := l.svcCtx.ProcInstModel.TransInsert(ctx, sqlx, &model.ProcInst{
			CreatedAt:     time.Now(),                                                      // 创建时间
			ProcId:        ProcessID,                                                       // 流程ID
			ProcName:      res.Name,                                                        // 流程名称
			ProcVersion:   res.Version,                                                     // 流程版本号
			BusinessId:    BusinessID,                                                      // 业务ID
			Starter:       res.CreateUserId.String,                                         // 流程发起人用户ID
			CurrentNodeId: StartNode.NodeID,                                                // 当前进行节点ID
			VariablesJson: sql.NullString{String: VariableJson, Valid: VariableJson != ""}, // 变量(Json)
			Status:        0,                                                               // 状态 0 未完成（审批中） 1 已完成 2 撤销
			TenantId:      res.TenantId,                                                    // 租户ID
			Data:          res.Data,                                                        //
			CreatedName:   res.CreatedName,                                                 // 创建人
		})
		if err != nil {
			return err
		}

		procInstId, _ := procInst.LastInsertId()
		//保存流程变量
		//获取变量数组
		var variables []Variable
		engine.Json2Struct(VariableJson, &variables)
		for _, v := range variables {

			procInstVariable, err := l.svcCtx.ProcInstVariableModel.FindOneByProcInstIdAndKey(procInstId, v.Key)
			if err != nil {
				return err
			}
			if procInstVariable.Id == 0 { //说明数据库中无此数据
				//插入
				_, err = l.svcCtx.ProcInstVariableModel.TransInsert(ctx, sqlx, &model.ProcInstVariable{
					CreatedAt:   time.Now(),      // 创建时间
					ProcInstId:  procInstId,      // 流程实例ID
					Key:         v.Key,           // 变量key
					Value:       v.Value,         // 变量value
					TenantId:    res.TenantId,    // 租户ID
					Data:        res.Data,        //
					CreatedName: res.CreatedName, // 创建人
				})
				if err != nil {
					return err
				}
			} else {
				//数据库中已有数据
				//更新
				procInstVariable.UpdatedName.String = res.CreatedName
				procInstVariable.UpdatedName.Valid = true
				procInstVariable.UpdatedAt.Time = time.Now()
				procInstVariable.UpdatedAt.Valid = true
				procInstVariable.Value = v.Value
				err = l.svcCtx.ProcInstVariableModel.TransUpdate(ctx, sqlx, procInstVariable)
			}

		}
		//获取流程起始人
		users, err := l.resolveNodeUser(procInstId, StartNode)

		//更新起始人到流程实例表
		resProcInst, err := l.svcCtx.ProcInstModel.FindOne(l.ctx, procInstId)
		if err != nil {
			if errors.Is(err, sqlc.ErrNotFound) {
				return fmt.Errorf("ProcInst没有该ID：%v", procInstId)
			}
			return err
		}

		resProcInst.UpdatedName.String = resProcInst.CreatedName
		resProcInst.UpdatedName.Valid = true
		resProcInst.UpdatedAt.Time = time.Now()
		resProcInst.UpdatedAt.Valid = true
		resProcInst.Starter = users[0]
		err = l.svcCtx.ProcInstModel.Update(l.ctx, resProcInst)
		return nil
	})

	return ProcessID, StartNode, nil
}

// map [NodeID]Node
type ProcNodes map[string]Node

// 定义流程cache其结构为 map [ProcID]ProcNodes
var ProcCache = make(map[int64]ProcNodes)

// 从缓存中获取流程节点定义
func (l *ProcInstStartLogic) GetProcCache(ProcessID int64) (ProcNodes, error) {
	if nodes, ok := ProcCache[ProcessID]; ok {
		return nodes, nil
	} else {
		process, err := l.GetProcessDefine(ProcessID)
		if err != nil {
			return nil, err
		}
		pn := make(ProcNodes)
		for _, n := range process.Nodes {
			pn[n.NodeID] = n
		}
		ProcCache[ProcessID] = pn
	}
	return ProcCache[ProcessID], nil
}

// 获取流程定义 by 流程ID
func (l *ProcInstStartLogic) GetProcessDefine(ProcessID int64) (Process, error) {
	type result struct {
		Resource string
	}
	var r result
	_, err := l.svcCtx.ProcDefModel.FindOne(l.ctx, ProcessID)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return Process{}, fmt.Errorf("ProcDef没有该ID：%v", ProcessID)
		}
		return Process{}, err
	}

	//如果数据库中没有找到ProcessID对应的流程,则直接报错
	if r.Resource == "" {
		return Process{}, errors.New("未找到对应流程定义")
	}

	return engine.ProcessParse(r.Resource)
}

type method struct {
	S interface{}    //method所在的struct，这是函数执行的第一个参数
	M reflect.Method //方法
}

// 事件池，所有的事件都会在流程引擎启动的时候注册到这里
var EventPool = make(map[string]method)

// 检查流程:
// 1、是否注册
// 2、参数是否正确
func (l *ProcInstStartLogic) VerifyEvents(ProcessID int64, Nodes ProcNodes) error {
	//获取流程定义
	process, err := l.GetProcessDefine(ProcessID)
	if err != nil {
		return err
	}

	//验证流程事件(目前只有撤销事件)
	for _, event := range process.RevokeEvents {
		if e, ok := EventPool[event]; !ok {
			return fmt.Errorf("事件%s尚未导入", event)
		} else {
			if err := verifyProcEventParameters(e.M); err != nil {
				return err
			}
		}
	}

	//各个节点中开始、结束事件 and 任务完成事件,先放入一个数组
	var nodeEvents []string
	for _, node := range Nodes {
		nodeEvents = append(nodeEvents, node.NodeStartEvents...)
		nodeEvents = append(nodeEvents, node.NodeEndEvents...)
		nodeEvents = append(nodeEvents, node.TaskFinishEvents...)
	}

	//各个节点中事件可能有重复的，需做去重
	nodeEventsSet := engine.MakeUnique(nodeEvents)

	//验证节点事件
	for _, event := range nodeEventsSet {
		if e, ok := EventPool[event]; !ok {
			return fmt.Errorf("事件%s尚未导入", event)
		} else {
			if err := verifyNodeEventParameters(e.M); err != nil {
				return err
			}
		}
	}

	return nil
}

// 验证流程事件(目前只有流程撤销事件)参数是否正确
// 流程撤销事件  func签名必须是func(struct *interface{}, ProcessInstanceID int,RevokeUserID string) error
func verifyProcEventParameters(m reflect.Method) error {
	//自定义函数必须是3个参数，参数0：*struct{} 1:int 2:String
	if m.Type.NumIn() != 3 || m.Type.NumOut() != 1 {
		fmt.Errorf("warning:事件方法 %s 入参、出参数量不匹配,此函数无法运行", m.Name)
	}

	if m.Type.In(1).Kind().String() != "int" {
		fmt.Errorf("warning:事件方法 %s 参数1不是int类型,此函数无法运行", m.Name)
	}

	if m.Type.In(2).Kind().String() != "string" {
		fmt.Errorf("warning:事件方法 %s 参数2不是string类型,此函数无法运行", m.Name)
	}

	if !TypeIsError(m.Type.Out(0)) {
		fmt.Errorf("warning:事件方法 %s 返回参数不是error类型,此函数无法运行", m.Name)
	}
	return nil
}

// 验证节点事件(1、节点开始  2、节点结束 3、任务结束)参数是否正确
// 1、节点开始、结束事件     func签名必须是func(struct *interface{}, ProcessInstanceID int, CurrentNode *Node, PrevNode Node) error
// 2、任务完成事件          func签名必须是func(struct *interface{}, TaskID int, CurrentNode *Node, PrevNode Node) error
func verifyNodeEventParameters(m reflect.Method) error {
	//自定义函数必须是4个参数，参数0：*struct{} 1:int 2:Node 3:Node
	if m.Type.NumIn() != 4 || m.Type.NumOut() != 1 {
		fmt.Errorf("warning:事件方法 %s 入参、出参数量不匹配,此函数无法运行", m.Name)
	}

	if m.Type.In(1).Kind().String() != "int" {
		fmt.Errorf("warning:事件方法 %s 参数1不是int类型,此函数无法运行", m.Name)
	}

	if m.Type.In(2).ConvertibleTo(reflect.TypeOf(&Node{})) != true {
		fmt.Errorf("warning:事件方法 %s 参数2不是*Node类型,此函数无法运行", m.Name)

	}

	if m.Type.In(3).ConvertibleTo(reflect.TypeOf(Node{})) != true {
		fmt.Errorf("warning:事件方法 %s 参数3不是Node类型,此函数无法运行", m.Name)
	}

	if !TypeIsError(m.Type.Out(0)) {
		fmt.Errorf("warning:事件方法 %s 返回参数不是error类型,此函数无法运行", m.Name)
	}
	return nil
}

func TypeIsError(Type reflect.Type) bool {
	//只要实现 Error() string方法的，就认为是实现了error接口
	//所以，要判断type中是否有一个方法名叫Error，无传入参数，输出string

	//如果都没有方法，自然不可能实现error
	if Type.NumMethod() >= 1 {
		for i := 0; i < Type.NumMethod(); i++ {
			if Type.Method(i).Name == "Error" {
				//是否无传入参数
				if Type.Method(i).Type.NumIn() != 0 {
					return false
				}

				//是否只有一个输出参数
				if Type.Method(i).Type.NumOut() != 1 {
					return false
				}

				//输出参数是否是string
				if Type.Method(i).Type.Out(0).Kind().String() != "string" {
					return false
				}

				return true
			}
		}
	}

	return false
}

// 解析节点用户
// 1、获得用户变量
// 2、用户去重
func (l *ProcInstStartLogic) resolveNodeUser(ProcessInstanceID int64, node Node) ([]string, error) {
	//匹配节点用户变量
	kv, err := l.ResolveVariables(ProcessInstanceID, node.UserIDs)
	if err != nil {
		return nil, err
	}

	//使用map去重，因为有可能某几个变量指向同一个用户，重复的用户会产生重复的任务
	var usersMap = make(map[string]string)
	for _, v := range kv {
		usersMap[v] = ""
	}

	//生成user数组
	var users []string
	for k, _ := range usersMap {
		users = append(users, k)
	}

	return users, nil
}

// 判断传入字符串是否是变量(是否以$开头)
func IsVariable(Key string) bool {
	if strings.HasPrefix(Key, "$") {
		return true
	}
	return false
}

// 去掉变量前缀"$"
func RemovePrefix(variable string) string {
	return strings.Replace(variable, "$", "", -1)
}

// 解析变量,获取并设置其value,返回map(注意，如果不是变量，则原样存储在map中)
func (l *ProcInstStartLogic) ResolveVariables(ProcessInstanceID int64, Variables []string) (map[string]string, error) {
	result := make(map[string]string)
	for _, v := range Variables {
		if IsVariable(v) {
			value, ok, err := l.SetVariable(ProcessInstanceID, v)
			if err != nil {
				return nil, err
			}
			if !ok {
				return nil, errors.New("无法匹配变量:" + v)
			}
			result[v] = value
		} else {
			result[v] = v
		}
	}
	return result, nil
}

func (l *ProcInstStartLogic) SetVariable(ProcessInstanceID int64, variable string) (string, bool, error) {
	Key := RemovePrefix(variable)

	//判断是否有匹配的值
	exists := false
	res, err := l.svcCtx.ProcInstVariableModel.FindOneByProcInstIdAndKey(ProcessInstanceID, Key)
	if err == nil {
		exists = true
		return res.Value, exists, nil
	} else {
		return "", false, err
	}
}
