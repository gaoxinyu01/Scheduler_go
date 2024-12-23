package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"Scheduler_go/service/workflow/rpc/internal/svc"
	"Scheduler_go/service/workflow/rpc/workflowclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskFinishedListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskFinishedListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskFinishedListLogic {
	return &TaskFinishedListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取已办任务
func (l *TaskFinishedListLogic) TaskFinishedList(in *workflowclient.TaskFinishedListReq) (*workflowclient.TaskFinishedListResp, error) {
	whereBuilder := l.svcCtx.ProcTaskModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 分配用户ID
	if len(in.UserId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"user_id ": "%" + in.UserId + "%",
		})
	}
	//
	if len(in.Data) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"data ": "%" + in.Data + "%",
		})
	}

	all, err := l.svcCtx.ProcTaskModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.ProcTaskModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	countBuilder = countBuilder.Where(squirrel.Eq{
		"tenant_id": in.TenantId,
	})

	// 分配用户ID
	if len(in.UserId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"user_id ": "%" + in.UserId + "%",
		})
	}
	//
	if len(in.Data) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"data ": "%" + in.Data + "%",
		})
	}
	count, err := l.svcCtx.ProcTaskModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
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

	return &workflowclient.TaskFinishedListResp{
		Total: count,
		List:  list,
	}, nil

}
