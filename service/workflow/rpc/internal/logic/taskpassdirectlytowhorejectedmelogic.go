package logic

import (
	"Scheduler_go/common/datax"
	"Scheduler_go/common/workflow/engine"
	"Scheduler_go/common/workflow/modelx"
	"Scheduler_go/service/workflow/model"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"regexp"
	"strings"
	"time"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskPassDirectlyToWhoRejectedMeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskPassDirectlyToWhoRejectedMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskPassDirectlyToWhoRejectedMeLogic {
	return &TaskPassDirectlyToWhoRejectedMeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 任务通过后流程直接返回到上一个驳回我的节点
func (l *TaskPassDirectlyToWhoRejectedMeLogic) TaskPassDirectlyToWhoRejectedMe(in *workflowclient.TaskPassDirectlyToWhoRejectedMeReq) (*workflowclient.CommonResp, error) {
	taskInfo, err := l.svcCtx.ProcTaskModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcTask没有该ID：%v", in.TaskId)
		}
		return nil, err
	}

	//判断节点是否已处理
	if taskInfo.IsFinished == 1 {
		return nil, fmt.Errorf("节点ID%d已处理，无需操作", in.TaskId)
	}
	//获取task所在的node
	CurrentNode, err := l.GetInstanceNode(taskInfo.ProcId, taskInfo.NodeId)

	//起始节点不能做驳回
	if CurrentNode.NodeType == modelx.RootNode {
		return nil, errors.New("起始节点无法驳回!")
	}

	//reject to 节点
	RejectToNode, err := l.GetInstanceNode(taskInfo.ProcId, datax.ToString(in.TaskId))
	if err != nil {
		return nil, err
	}
	//判断节点是否已处理
	if taskInfo.IsFinished == 1 {
		return nil, fmt.Errorf("节点ID%d已处理，无需操作", taskInfo.Id)
	}
	err = l.svcCtx.ProcTaskModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		_, err = l.svcCtx.ProcTaskModel.TransInsert(l.ctx, session, &model.ProcTask{
			CreatedAt:          time.Now(),                  // 创建时间
			ProcId:             taskInfo.ProcId,             // 流程ID
			ProcInstId:         taskInfo.ProcInstId,         // 流程实例ID
			BusinessId:         taskInfo.BusinessId,         // 业务ID
			Starter:            taskInfo.Starter,            // 流程发起人用户ID
			NodeId:             taskInfo.NodeId,             // 节点ID
			NodeName:           taskInfo.NodeName,           // 节点名称
			PrevNodeId:         taskInfo.PrevNodeId,         // 上个处理节点ID
			IsCosigned:         taskInfo.IsCosigned,         // 任意一人通过即可 1:会签
			BatchCode:          taskInfo.BatchCode,          // 批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别\"
			UserId:             taskInfo.UserId,             // 分配用户ID
			Status:             2,                           // 任务状态:0:初始 1:通过 2:驳回
			IsFinished:         taskInfo.IsFinished,         // 0:任务未完成 1:处理完成
			Comment:            taskInfo.Comment,            // 任务备注
			ProcInstCreateTime: taskInfo.ProcInstCreateTime, // 流程实例创建时间
			FinishedTime:       taskInfo.FinishedTime,       // 处理任务时间
			TenantId:           in.TenantId,                 // 租户ID
			Data:               taskInfo.Data,               //
			CreatedName:        taskInfo.CreatedName,        // 创建人
		})
		if err != nil {
			return err
		}
		//1、非会签节点，一人通过即通过，所以要把其他人的任务finish掉
		//2、不论是否会签，都是一人驳回即驳回，所以需要把同一批次task的isfinish设置为1,让其他人不用再处理
		if (taskInfo.IsCosigned == 0 && taskInfo.Status == 1) || taskInfo.Status == 2 {
			resBatchCode, err := l.svcCtx.ProcTaskModel.FindOneByProcBatchCode(l.ctx, taskInfo.BatchCode)
			if err != nil {
				if errors.Is(err, sqlc.ErrNotFound) {
					return fmt.Errorf("ProcTask没有该ID：%v", taskInfo.BatchCode)
				}
				return err
			}

			resBatchCode.UpdatedName.String = resBatchCode.CreatedName
			resBatchCode.UpdatedName.Valid = true
			resBatchCode.UpdatedAt.Time = time.Now()
			resBatchCode.UpdatedAt.Valid = true
			resBatchCode.IsFinished = 1
			resBatchCode.FinishedTime = time.Now()

			err = l.svcCtx.ProcTaskModel.Update(l.ctx, resBatchCode)

		}

		//设置实例变量
		//获取变量数组
		var variables []modelx.Variable
		engine.Json2Struct(in.VariablesJson, &variables)
		for _, v := range variables {

			procInstVariable, err := l.svcCtx.ProcInstVariableModel.FindOneByProcInstIdAndKey(taskInfo.ProcInstId, v.Key)
			if err != nil {
				return err
			}
			if procInstVariable.Id == 0 { //说明数据库中无此数据
				//插入
				_, err = l.svcCtx.ProcInstVariableModel.TransInsert(ctx, session, &model.ProcInstVariable{
					CreatedAt:   time.Now(),           // 创建时间
					ProcInstId:  taskInfo.ProcInstId,  // 流程实例ID
					Key:         v.Key,                // 变量key
					Value:       v.Value,              // 变量value
					TenantId:    taskInfo.TenantId,    // 租户ID
					Data:        taskInfo.Data,        //
					CreatedName: taskInfo.CreatedName, // 创建人
				})
				if err != nil {
					return err
				}
			} else {
				//数据库中已有数据
				//更新
				procInstVariable.UpdatedName.String = taskInfo.CreatedName
				procInstVariable.UpdatedName.Valid = true
				procInstVariable.UpdatedAt.Time = time.Now()
				procInstVariable.UpdatedAt.Valid = true
				procInstVariable.Value = v.Value
				err = l.svcCtx.ProcInstVariableModel.TransUpdate(ctx, session, procInstVariable)
			}

		}

		return nil
	})

	err = l.ProcessNode(taskInfo.ProcInstId, &RejectToNode, CurrentNode)
	if err != nil {
		l.taskRevoke(taskInfo.Id)
	}

	return &workflowclient.CommonResp{}, nil
}

// 获取流程实例中某个Node 返回 Node
func (l *TaskPassDirectlyToWhoRejectedMeLogic) GetInstanceNode(ProcessInstanceID int64, NodeID string) (modelx.Node, error) {
	//获取task所在的node
	resProcInst, err := l.svcCtx.ProcInstModel.FindOne(l.ctx, ProcessInstanceID)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return modelx.Node{}, fmt.Errorf("ProcInst没有该ID：%v", ProcessInstanceID)
		}
		return modelx.Node{}, err
	}

	if err != nil {
		return modelx.Node{}, err
	}

	//从Cache中获得流程节点列表
	Nodes, err := l.GetProcCache(resProcInst.ProcId)
	if err != nil {
		return modelx.Node{}, err
	}
	//获得节点
	node, ok := Nodes[NodeID]
	if !ok {
		return modelx.Node{}, fmt.Errorf("ID为%d的流程实例中不存在ID为%s的节点", ProcessInstanceID, NodeID)
	}

	return node, nil
}

// 从缓存中获取流程节点定义
func (l *TaskPassDirectlyToWhoRejectedMeLogic) GetProcCache(ProcessID int64) (ProcNodes, error) {
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
func (l *TaskPassDirectlyToWhoRejectedMeLogic) GetProcessDefine(ProcessID int64) (modelx.Process, error) {
	type result struct {
		Resource string
	}
	var r result
	_, err := l.svcCtx.ProcDefModel.FindOne(l.ctx, ProcessID)

	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return modelx.Process{}, fmt.Errorf("ProcDef没有该ID：%v", ProcessID)
		}
		return modelx.Process{}, err
	}

	//如果数据库中没有找到ProcessID对应的流程,则直接报错
	if r.Resource == "" {
		return modelx.Process{}, errors.New("未找到对应流程定义")
	}

	return engine.ProcessParse(r.Resource)
}

// 任务提交之后需要处理后继工作：任务结束事件、节点结束事件等.
// 如果这些事件出错，则之前已提交的任务就成为了死任务，整个流程就被挂起.
// 所以，要在出错后对之前的任务做初始化恢复.
func (l *TaskPassDirectlyToWhoRejectedMeLogic) taskRevoke(TaskID int64) error {

	res, err := l.svcCtx.ProcTaskModel.FindOne(l.ctx, TaskID)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return fmt.Errorf("ProcTask没有该ID：%v", TaskID)
		}
		return err
	}
	res.IsFinished = 0
	res.Status = 0
	res.FinishedTime = time.Time{}
	res.UpdatedName.String = res.CreatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.ProcTaskModel.Update(l.ctx, res)

	return nil
}

// 处理节点,如：生成task、进行条件判断、处理结束节点等
func (l *TaskPassDirectlyToWhoRejectedMeLogic) ProcessNode(ProcessInstanceID int64, CurrentNode *modelx.Node, PrevNode modelx.Node) error {
	//这里处理开始事件
	err := RunNodeEvents(CurrentNode.NodeStartEvents, ProcessInstanceID, CurrentNode, PrevNode)
	if err != nil {
		return err
	}

	//开始节点也需要处理，因为开始节点可能因为驳回而重新回到开始节点，此时的开始节点=普通任务节点
	if CurrentNode.NodeType == modelx.RootNode {
		_, err := l.TaskNodeHandle(ProcessInstanceID, CurrentNode, PrevNode)
		if err != nil {
			return err
		}
	}

	if CurrentNode.NodeType == modelx.GateWayNode {
		err := l.GateWayNodeHandle(ProcessInstanceID, CurrentNode, PrevNode)
		if err != nil {
			return err
		}
	}

	if CurrentNode.NodeType == modelx.TaskNode {
		_, err := l.TaskNodeHandle(ProcessInstanceID, CurrentNode, PrevNode)
		if err != nil {
			return err
		}
	}

	if CurrentNode.NodeType == modelx.EndNode {
		err := l.EndNodeHandle(ProcessInstanceID, 1)
		if err != nil {
			return err
		}
	}

	return nil
}

// 任务节点处理 返回生成的taskid数组
func (l *TaskPassDirectlyToWhoRejectedMeLogic) TaskNodeHandle(ProcessInstanceID int64, CurrentNode *modelx.Node, PrevNode modelx.Node) ([]int64, error) {
	//获取节点用户
	users, err := l.resolveNodeUser(ProcessInstanceID, *CurrentNode)
	if err != nil {
		return nil, err
	}

	//如果没有处理人，则任务无法分配
	if len(users) < 1 {
		return nil, errors.New("未指定处理人，无法处理节点:" + CurrentNode.NodeName)
	}

	//开始节点只能有一个用户发起,不管多少用户，只要第一个
	//思考：如果开始节点有多个处理人，则可能进入会签状态，可能造成流程都无法正常开始
	//所以，开始节点只能有一个处理人，且默认就是非会签节点
	if CurrentNode.NodeType == modelx.RootNode {
		users = users[0:1]
	}
	res, err := l.svcCtx.ProcInstModel.FindOne(l.ctx, ProcessInstanceID)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcInst没有该ID：%v", ProcessInstanceID)
		}
		return nil, err
	}
	//生成Task
	var taskIDs []int64
	for _, user := range users {
		ProcTask, err := l.svcCtx.ProcTaskModel.Insert(l.ctx, &model.ProcTask{
			CreatedAt:   time.Now(),         // 创建时间
			ProcId:      res.ProcId,         // 流程ID
			ProcInstId:  ProcessInstanceID,  // 流程实例ID
			BusinessId:  res.BusinessId,     // 业务ID
			Starter:     res.Starter,        // 流程发起人用户ID
			NodeId:      CurrentNode.NodeID, // 节点ID
			PrevNodeId:  PrevNode.NodeID,    // 上个处理节点ID
			IsCosigned:  0,                  // 任意一人通过即可 1:会签
			UserId:      user,               // 分配用户ID
			Status:      1,                  // 任务状态:0:初始 1:通过 2:驳回
			IsFinished:  1,                  // 0:任务未完成 1:处理完成
			TenantId:    res.TenantId,       // 租户ID
			Data:        res.Data,           //
			CreatedName: res.CreatedName,    // 创建人
		})
		if err != nil {
			return nil, err
		}
		taskID, _ := ProcTask.LastInsertId()
		taskIDs = append(taskIDs, taskID)
	}

	if err != nil {
		return nil, err
	}
	return taskIDs, nil
}

// GateWay节点处理
func (l *TaskPassDirectlyToWhoRejectedMeLogic) GateWayNodeHandle(ProcessInstanceID int64, CurrentNode *modelx.Node, PrevTaskNode modelx.Node) error {
	//--------------------首先，混合节点需要确认所有的上级节点都处理完，才能做下一步--------------------
	var totalFinished int                          //所有已完成的上级节点
	totalPrevNodes := len(CurrentNode.PrevNodeIDs) //所有上级节点
	for _, nodeID := range CurrentNode.PrevNodeIDs {
		finished, err := l.InstanceNodeIsFinish(ProcessInstanceID, nodeID)
		if err != nil {
			return err
		}
		if finished != 0 {
			totalFinished++
		}
	}

	//如果是并行网关模式，还有尚未完成的上级节点，则退出
	if CurrentNode.GWConfig.WaitForAllPrevNode == 1 && totalPrevNodes != totalFinished {
		return nil
	}

	//如果是包含网关模式,连一个已完成的上级节点都没有，则退出
	if CurrentNode.GWConfig.WaitForAllPrevNode == 0 && totalFinished < 1 {
		return nil
	}

	//----------------------------计算条件----------------------------
	var conditionNodeIDs []string //condition指定的下级Node
	//一个GW节点可以有多个condition,所以要遍历
	for _, c := range CurrentNode.GWConfig.Conditions {
		//正则表达式，匹配以$开头的字母、数字、下划线
		reg := regexp.MustCompile(`[$]\w+`)
		//获取表达式中所有的变量
		variables := reg.FindAllString(c.Expression, -1)

		//替换表达式中的变量为值
		expression := c.Expression
		//获取变量对应的value
		kv, err := l.ResolveVariables(ProcessInstanceID, variables)
		if err != nil {
			return err
		}
		for k, v := range kv {
			expression = strings.Replace(expression, k, v, -1)
		}

		//首先通过正则表达式判断是否有SQL注入风险
		pattern := regexp.MustCompile("delete|truncate|insert|drop|create|select|update|set|from|grant|call|execute")
		match := pattern.FindString(strings.ToLower(expression))
		if match != "" {
			return errors.New("表达式中包含危险词,可能造成SQL注入!")
		}
		// 获取请求方法
		method := l.ctx.Value(expression)
		if method != nil {
			return err
		} else {
			conditionNodeIDs = append(conditionNodeIDs, c.NodeID)
		}
	}

	//-------将conditionNodeIDs和InevitableNodes中的值一起放入nextNodeIDs，这是真正要处理的节点ID-------
	//去重(节点ID如果重复，意味着一个节点要做N次处理，这是灾难)
	nextNodeIDs := MakeUnique(conditionNodeIDs, CurrentNode.GWConfig.InevitableNodes)

	//这里处理节点结束事件
	err := RunNodeEvents(CurrentNode.NodeEndEvents, ProcessInstanceID, CurrentNode, PrevTaskNode)
	if err != nil {
		return err
	}

	//------------------------------对下级节点进行处理------------------------------
	for _, nodeID := range nextNodeIDs {
		NextNode, err := l.GetInstanceNode(ProcessInstanceID, nodeID)
		if err != nil {
			return err
		}
		/*
			思考一个问题，ProcessNod函数的形参PrevNode应该传什么？
			如果传当前处理的GW节点本身，则要思考以下情况：
			节点定义是task1-gw1-gw2-task2，如果在gw1处理的最后，ProcessNode的PrevNode传gw1本身，那么task2就永远找不到task1了
			所以，在处理gw节点时,ProcessNod函数的形参PrevNode不能传gw本身，而是要传gw的上一节点，因为：
			1、只有任务节点才能开启一个gw
			2、直接把任务节点作为PrevTaskNode传入，就算下一个节点还是gw，重复此行为，之后的task节点还是可以获得上一个task节点
		*/
		err = l.ProcessNode(ProcessInstanceID, &NextNode, PrevTaskNode)
		if err != nil {
			return err
		}
	}

	return nil

}

// 判断特定实例中某一个节点是否已经完成
// 注意，finish只是代表节点是不是已经处理，不管处理的方式是驳回还是通过
// 一个流程实例中，由于驳回等原因，x节点可能出现多次。这里使用统计所有x节点的任务是否都finish来判断x节点是否finish
func (l *TaskPassDirectlyToWhoRejectedMeLogic) InstanceNodeIsFinish(ProcessInstanceID int64, NodeID string) (int64, error) {

	res, err := l.svcCtx.ProcTaskModel.FindOne(l.ctx, ProcessInstanceID)
	if err != nil {

		return res.IsFinished, nil
	} else {
		return res.IsFinished, err
	}
}

// 结束节点处理 结束节点只做收尾工作，将数据库中此流程实例产生的数据归档
// Status 流程实例状态 1:已完成 2:撤销
func (l *TaskPassDirectlyToWhoRejectedMeLogic) EndNodeHandle(ProcessInstanceID int64, Status int64) error {

	resProcTask, err := l.svcCtx.ProcTaskModel.FindOne(l.ctx, ProcessInstanceID)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return fmt.Errorf("ProcTask没有该ID：%v", ProcessInstanceID)
		}
		return err
	}
	resProcTask.IsFinished = 1
	resProcTask.FinishedTime = time.Now()
	resProcTask.UpdatedName.String = resProcTask.CreatedName
	resProcTask.UpdatedName.Valid = true
	resProcTask.UpdatedAt.Time = time.Now()
	resProcTask.UpdatedAt.Valid = true

	err = l.svcCtx.ProcTaskModel.Update(l.ctx, resProcTask)

	_, err = l.svcCtx.HistProcInstVariableModel.Insert(l.ctx, &model.HistProcInstVariable{
		CreatedAt:   time.Now(),              // 创建时间
		ProcInstId:  resProcTask.ProcInstId,  // 流程实例ID
		TenantId:    resProcTask.TenantId,    // 租户ID
		Data:        resProcTask.Data,        //
		CreatedName: resProcTask.CreatedName, // 创建人
	})
	if err != nil {
		return err
	}

	return nil
}

// 解析节点用户
// 1、获得用户变量
// 2、用户去重
func (l *TaskPassDirectlyToWhoRejectedMeLogic) resolveNodeUser(ProcessInstanceID int64, node modelx.Node) ([]string, error) {
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

// 解析变量,获取并设置其value,返回map(注意，如果不是变量，则原样存储在map中)
func (l *TaskPassDirectlyToWhoRejectedMeLogic) ResolveVariables(ProcessInstanceID int64, Variables []string) (map[string]string, error) {
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

// 从proc_inst_variable表中查找变量，若有则返回变量值,若无则返回false
func (l *TaskPassDirectlyToWhoRejectedMeLogic) SetVariable(ProcessInstanceID int64, variable string) (string, bool, error) {
	Key := RemovePrefix(variable)
	type result struct {
		Value string
	}
	var r result
	_, err := l.svcCtx.ProcInstVariableModel.FindOneByProcInstIdAndKey(ProcessInstanceID, Key)
	if err != nil {

		//判断是否有匹配的值
		exists := false
		if r.Value != "" {
			exists = true
		}

		return r.Value, exists, nil
	} else {
		return "", false, err
	}
}
