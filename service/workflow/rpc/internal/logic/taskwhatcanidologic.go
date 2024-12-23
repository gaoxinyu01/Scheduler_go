package logic

import (
	"Scheduler_go/common/workflow/engine"
	"Scheduler_go/common/workflow/modelx"
	"Scheduler_go/service/workflow/model"
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskWhatCanIDoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskWhatCanIDoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskWhatCanIDoLogic {
	return &TaskWhatCanIDoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 当前任务可以执行哪些操作
func (l *TaskWhatCanIDoLogic) TaskWhatCanIDo(in *workflowclient.TaskWhatCanIDoReq) (*workflowclient.TaskWhatCanIDoResp, error) {
	var act modelx.TaskAction
	act = modelx.TaskAction{
		CanPass:                     true,
		CanReject:                   true,
		CanFreeRejectToUpstreamNode: true,
		CanDirectlyToWhoRejectedMe:  true,
		CanRevoke:                   false} //初始化

	//获得task信息
	taskInfo, err := l.svcCtx.ProcTaskModel.FindOne(l.ctx, in.Taskid)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("ProcTask没有该ID：%v", in.Taskid)
		}
		return nil, err
	}

	//如果任务已经完成，则什么都做不了
	if taskInfo.IsFinished == 1 {
		return nil, nil
	}
	node, err := l.GetInstanceNode(taskInfo.ProcInstId, taskInfo.NodeId)

	if err != nil {
		return nil, nil
	}

	//起始节点不能做驳回动作 & 可驳回
	if node.NodeType == modelx.RootNode {
		act.CanReject = false
		act.CanFreeRejectToUpstreamNode = false
		act.CanRevoke = true
	}

	//会签节点不能使用DirectlyToWhoRejectedMe功能
	if taskInfo.IsCosigned == 1 {
		act.CanDirectlyToWhoRejectedMe = false
	}

	//此任务的上一节点并未做驳回,无法使用DirectlyToWhoRejectedMe功能
	err, PrevNodeIsReject := l.taskPrevNodeIsReject(taskInfo)
	if err != nil {
		return nil, err
	}
	if PrevNodeIsReject == false {
		act.CanDirectlyToWhoRejectedMe = false
	}

	return &workflowclient.TaskWhatCanIDoResp{}, nil
}
func (l *TaskWhatCanIDoLogic) GetInstanceNode(ProcessInstanceID int64, NodeID string) (modelx.Node, error) {
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
func (l *TaskWhatCanIDoLogic) GetProcCache(ProcessID int64) (ProcNodes, error) {
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
func (l *TaskWhatCanIDoLogic) GetProcessDefine(ProcessID int64) (modelx.Process, error) {
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

// 任务的上一个节点是不是做了驳回
func (l *TaskWhatCanIDoLogic) taskPrevNodeIsReject(TaskInfo *model.ProcTask) (error, bool) {
	//获得实际执行过程中上一个节点的BatchCode
	whereBuilder := l.svcCtx.ProcTaskModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别\"
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"batch_code ": TaskInfo.BatchCode,
	})
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"status ": 2,
	})

	all, err := l.svcCtx.ProcTaskModel.FindList(l.ctx, whereBuilder, 1, 9999999)
	if err != nil {
		return err, false
	}

	var list []*workflowclient.ProcTaskListData
	for _, item := range all {
		list = append(list, &workflowclient.ProcTaskListData{
			Id:                 item.Id,                                  //任务ID
			ProcId:             item.ProcId,                              //流程ID
			ProcInstId:         item.ProcInstId,                          //流程实例ID
			BusinessId:         item.BusinessId,                          //业务ID
			Starter:            item.Starter,                             //流程发起人用户ID
			NodeId:             item.NodeId,                              //节点ID
			NodeName:           item.NodeName,                            //节点名称
			PrevNodeId:         item.PrevNodeId,                          //上个处理节点ID
			IsCosigned:         item.IsCosigned,                          //任意一人通过即可 1:会签
			BatchCode:          item.BatchCode,                           //批次码.节点会被驳回，一个节点可能产生多批task,用此码做分别\"
			UserId:             item.UserId,                              //分配用户ID
			Status:             item.Status,                              //任务状态:0:初始 1:通过 2:驳回
			IsFinished:         item.IsFinished,                          //0:任务未完成 1:处理完成
			Comment:            item.Comment.String,                      //任务备注
			ProcInstCreateTime: item.ProcInstCreateTime.Time.UnixMilli(), //流程实例创建时间
			FinishedTime:       item.FinishedTime.UnixMilli(),            //处理任务时间
			Data:               item.Data.String,                         //
			CreatedAt:          item.CreatedAt.UnixMilli(),               //创建时间
			UpdatedAt:          item.UpdatedAt.Time.UnixMilli(),          //更新时间
			CreatedName:        item.CreatedName,                         //创建人
			UpdatedName:        item.UpdatedName.String,                  //更新人
		})
	}

	//没有找到,说明上一个节点中没有做驳回
	if len(list) < 0 {
		return nil, false
	} else {
		return nil, true
	}
}
