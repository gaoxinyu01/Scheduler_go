// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type ProcDefDelRequest struct {
	Id int64 `path:"id"` // 流程模板ID
}

type ProcDefInfoRequest struct {
	Id int64 `form:"id"` // 流程模板ID
}

type ProcDefListRequest struct {
	Current      int64  `form:"current,default=1,optional"`    // 页码
	PageSize     int64  `form:"page_size,default=10,optional"` // 页数
	Name         string `form:"name,optional"`                 // 流程名称
	Version      int64  `form:"version,default=99,optional"`   // 版本号
	ProcType     int64  `form:"proc_type,default=99,optional"` // 流程类型
	Resource     string `form:"resource,optional"`             // 流程定义模板
	CreateUserId string `form:"create_user_id,optional"`       // 创建者ID
	Source       string `form:"source,optional"`               // 来源
	Data         string `form:"data,optional"`                 //
}

type ProcDefSaveRequest struct {
	Resource     string `json:"resource"`       // 流程定义模板
	CreateUserId string `json:"create_user_id"` // 创建者ID
	Data         string `json:"data,optional"`  //
}

type ProcDefUpRequest struct {
	Id           int64  `json:"id"`                      // 流程模板ID
	Name         string `json:"name,optional"`           // 流程名称
	Version      int64  `json:"version,optional"`        // 版本号
	ProcType     int64  `json:"proc_type,optional"`      // 流程类型
	Resource     string `json:"resource,optional"`       // 流程定义模板
	CreateUserId string `json:"create_user_id,optional"` // 创建者ID
	Source       string `json:"source,optional"`         // 来源
	Data         string `json:"data,optional"`           //
}

type ProcInstDelRequest struct {
	Id int64 `path:"id"` // 流程实例ID
}

type ProcInstInfoRequest struct {
	Id int64 `form:"id"` // 流程实例ID
}

type ProcInstListRequest struct {
	Current       int64  `form:"current,default=1,optional"`       // 页码
	PageSize      int64  `form:"page_size,default=10,optional"`    // 页数
	ProcId        int64  `form:"proc_id,default=99,optional"`      // 流程ID
	ProcName      string `form:"proc_name,optional"`               // 流程名称
	ProcVersion   int64  `form:"proc_version,default=99,optional"` // 流程版本号
	BusinessId    string `form:"business_id,optional"`             // 业务ID
	Starter       string `form:"starter,optional"`                 // 流程发起人用户ID
	CurrentNodeId string `form:"current_node_id,optional"`         // 当前进行节点ID
	VariablesJson string `form:"variables_json,optional"`          // 变量(Json)
	Status        int64  `form:"status,default=99,optional"`       // 状态 0 未完成（审批中） 1 已完成 2 撤销
	Data          string `form:"data,optional"`                    //
}

type ProcInstRevokeRequest struct {
	Id           int64  `json:"id"`                                    // 流程实例ID
	RevokeUserID string `json:"revoke_user_id,optional"`               // 撤销发起用户ID
	Force        int64  `json:"force,default=0,options=0|1|,optional"` // 是否强制撤销
}

type ProcInstStartRequest struct {
	ProcId        int64  `json:"proc_id"`        // 流程ID
	BusinessId    string `json:"business_id"`    // 业务ID
	Comment       string `json:"comment"`        // 评论意见
	VariablesJson string `json:"variables_json"` // 变量(Json)
	Data          string `json:"data,optional"`  //预留
}

type ProcInstTaskHistoryRequest struct {
	Current  int64  `form:"current,default=1,optional"`    // 页码
	PageSize int64  `form:"page_size,default=10,optional"` // 页数
	InstId   int64  `form:"instid,default=99,optional"`    // 流程ID
	Data     string `form:"data,optional"`                 //
}

type ProcInstUpRequest struct {
	Id            int64  `json:"id"`                       // 流程实例ID
	ProcId        int64  `json:"proc_id,optional"`         // 流程ID
	ProcName      string `json:"proc_name,optional"`       // 流程名称
	ProcVersion   int64  `json:"proc_version,optional"`    // 流程版本号
	BusinessId    string `json:"business_id,optional"`     // 业务ID
	Starter       string `json:"starter,optional"`         // 流程发起人用户ID
	CurrentNodeId string `json:"current_node_id,optional"` // 当前进行节点ID
	VariablesJson string `json:"variables_json,optional"`  // 变量(Json)
	Status        int64  `json:"status,optional"`          // 状态 0 未完成（审批中） 1 已完成 2 撤销
	Data          string `json:"data,optional"`            //
}

type Response struct {
	Code int64       `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据
}

type TaskFinishedListRequest struct {
	Current         int64  `form:"current,default=1,optional"`            // 页码
	PageSize        int64  `form:"page_size,default=10,optional"`         // 页数
	UserId          string `form:"user_id"`                               // 任务ID
	ProcName        string `form:"proc_name,optional"`                    // 指定流程名称 ，非必填
	IgnoreStartByMe int64  `form:"ignore_start_by_me,default=0,optional"` //忽略由我开启流程,而生成处理人是我自己的任务  0 不忽略1 忽略
	SortByAsc       int64  `form:"sort_by_asc,default=0"`                 //是否按照任务生成时间升序排列,0 降序1 升序
	Data            string `form:"data,optional"`                         //
}

type TaskFreeRejectToUpstreamNodeRequest struct {
	TaskId         int64  `json:"task_id"`           // 任务ID
	Comment        string `json:"comment"`           // 评论意见
	VariablesJson  string `json:"variables_json"`    // 变量(Json)
	RejectToNodeID string `json:"reject_to_node_id"` //驳回到哪个节点
	Data           string `json:"data,optional"`     //
}

type TaskInfoRequest struct {
	TaskId int64 `form:"task_id"` // 任务ID
}

type TaskPassDirectlyToWhoRejectedMeRequest struct {
	TaskId        int64  `json:"task_id"`        // 任务ID
	Comment       string `json:"comment"`        // 评论意见
	VariablesJson string `json:"variables_json"` // 变量(Json)
	Data          string `json:"data,optional"`  //
}

type TaskPassRequest struct {
	TaskId                  int64  `json:"task_id"`                                        // 任务ID
	Comment                 string `json:"comment"`                                        // 评论意见
	VariablesJson           string `json:"variables_json"`                                 // 变量(Json)
	Data                    string `json:"data,optional"`                                  //
	DirectlyToWhoRejectedMe int64  `json:"directly_to_who_rejected_me,default=0,optional"` //任务通过(pass)时直接返回到上一个驳回我的节点 0不返回  1 返回
}

type TaskRejectRequest struct {
	TaskId        int64  `json:"task_id"`        // 任务ID
	Comment       string `json:"comment"`        // 评论意见
	VariablesJson string `json:"variables_json"` // 变量(Json)
	Data          string `json:"data,optional"`  //
}

type TaskToDoListRequest struct {
	Current   int64  `form:"current,default=1,optional"`    // 页码
	PageSize  int64  `form:"page_size,default=10,optional"` // 页数
	UserId    string `form:"user_id"`                       // 任务ID
	ProcId    int64  `form:"proc_id,optional"`              // 指定流程名称 ，非必填
	SortByAsc int64  `form:"sort_by_asc"`                   //是否按照任务生成时间升序排列,0 降序1 升序
	Data      string `form:"data,optional"`                 //
}

type TaskTransferRequest struct {
	TaskId int64    `json:"task_id"`       // 任务ID
	Users  []string `json:"users"`         //用户
	Data   string   `json:"data,optional"` //
}

type TaskUpstreamNodeListRequest struct {
	TaskId int64 `form:"task_id"` // 任务ID
}

type TaskWhatCanIDoRequest struct {
	TaskId int64 `form:"task_id"` // 任务ID
}