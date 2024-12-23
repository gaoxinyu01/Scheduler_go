// file: operation.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
// source: scheduler.proto

package schedulerclient

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Scheduler_AttendanceAdd_FullMethodName          = "/schedulerclient.Scheduler/AttendanceAdd"
	Scheduler_AttendancePatch_FullMethodName        = "/schedulerclient.Scheduler/AttendancePatch"
	Scheduler_AttendanceDelete_FullMethodName       = "/schedulerclient.Scheduler/AttendanceDelete"
	Scheduler_AttendanceUpdate_FullMethodName       = "/schedulerclient.Scheduler/AttendanceUpdate"
	Scheduler_AttendanceFindOne_FullMethodName      = "/schedulerclient.Scheduler/AttendanceFindOne"
	Scheduler_AttendanceList_FullMethodName         = "/schedulerclient.Scheduler/AttendanceList"
	Scheduler_AttendanceFindOneDay_FullMethodName   = "/schedulerclient.Scheduler/AttendanceFindOneDay"
	Scheduler_AttendanceByDays_FullMethodName       = "/schedulerclient.Scheduler/AttendanceByDays"
	Scheduler_TeamTypeAdd_FullMethodName            = "/schedulerclient.Scheduler/TeamTypeAdd"
	Scheduler_TeamTypeDelete_FullMethodName         = "/schedulerclient.Scheduler/TeamTypeDelete"
	Scheduler_TeamTypeUpdate_FullMethodName         = "/schedulerclient.Scheduler/TeamTypeUpdate"
	Scheduler_TeamTypeFindList_FullMethodName       = "/schedulerclient.Scheduler/TeamTypeFindList"
	Scheduler_TeamAdd_FullMethodName                = "/schedulerclient.Scheduler/TeamAdd"
	Scheduler_TeamDelete_FullMethodName             = "/schedulerclient.Scheduler/TeamDelete"
	Scheduler_TeamUpdate_FullMethodName             = "/schedulerclient.Scheduler/TeamUpdate"
	Scheduler_TeamFindList_FullMethodName           = "/schedulerclient.Scheduler/TeamFindList"
	Scheduler_SchedulingTypeAdd_FullMethodName      = "/schedulerclient.Scheduler/SchedulingTypeAdd"
	Scheduler_SchedulingTypeDelete_FullMethodName   = "/schedulerclient.Scheduler/SchedulingTypeDelete"
	Scheduler_SchedulingTypeUpdate_FullMethodName   = "/schedulerclient.Scheduler/SchedulingTypeUpdate"
	Scheduler_SchedulingTypeFindList_FullMethodName = "/schedulerclient.Scheduler/SchedulingTypeFindList"
	Scheduler_SchedulingAdd_FullMethodName          = "/schedulerclient.Scheduler/SchedulingAdd"
	Scheduler_SchedulingDelete_FullMethodName       = "/schedulerclient.Scheduler/SchedulingDelete"
	Scheduler_SchedulingUpdate_FullMethodName       = "/schedulerclient.Scheduler/SchedulingUpdate"
	Scheduler_SchedulingFindList_FullMethodName     = "/schedulerclient.Scheduler/SchedulingFindList"
)

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerClient interface {
	// 考勤
	AttendanceAdd(ctx context.Context, in *AttendanceAddReq, opts ...grpc.CallOption) (*CommonResp, error)
	// 签退
	AttendancePatch(ctx context.Context, in *AttendancePatchReq, opts ...grpc.CallOption) (*CommonResp, error)
	AttendanceDelete(ctx context.Context, in *AttendanceDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
	AttendanceUpdate(ctx context.Context, in *AttendanceUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
	AttendanceFindOne(ctx context.Context, in *AttendanceFindOneReq, opts ...grpc.CallOption) (*AttendanceFindOneResp, error)
	AttendanceList(ctx context.Context, in *AttendanceListReq, opts ...grpc.CallOption) (*AttendanceListResp, error)
	// 获取某天考勤
	AttendanceFindOneDay(ctx context.Context, in *AttendanceFindOneDayReq, opts ...grpc.CallOption) (*AttendanceFindOneDayResp, error)
	// 根据时间段获取每日考勤
	AttendanceByDays(ctx context.Context, in *AttendanceByDaysReq, opts ...grpc.CallOption) (*AttendanceByDaysResp, error)
	// 部门
	TeamTypeAdd(ctx context.Context, in *TeamTypeAddReq, opts ...grpc.CallOption) (*CommonResp, error)
	TeamTypeDelete(ctx context.Context, in *TeamTypeDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
	TeamTypeUpdate(ctx context.Context, in *TeamTypeUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
	TeamTypeFindList(ctx context.Context, in *TeamTypeFindListReq, opts ...grpc.CallOption) (*TeamTypeFindListResp, error)
	// 部门人员表
	TeamAdd(ctx context.Context, in *TeamAddReq, opts ...grpc.CallOption) (*CommonResp, error)
	TeamDelete(ctx context.Context, in *TeamDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
	TeamUpdate(ctx context.Context, in *TeamUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
	TeamFindList(ctx context.Context, in *TeamFindListReq, opts ...grpc.CallOption) (*TeamFindListResp, error)
	// 排班类型
	SchedulingTypeAdd(ctx context.Context, in *SchedulingTypeAddReq, opts ...grpc.CallOption) (*CommonResp, error)
	SchedulingTypeDelete(ctx context.Context, in *SchedulingTypeDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
	SchedulingTypeUpdate(ctx context.Context, in *SchedulingTypeUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
	SchedulingTypeFindList(ctx context.Context, in *SchedulingTypeFindListReq, opts ...grpc.CallOption) (*SchedulingTypeFindListResp, error)
	// 排班列表
	SchedulingAdd(ctx context.Context, in *SchedulingAddReq, opts ...grpc.CallOption) (*CommonResp, error)
	SchedulingDelete(ctx context.Context, in *SchedulingDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
	SchedulingUpdate(ctx context.Context, in *SchedulingUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
	SchedulingFindList(ctx context.Context, in *SchedulingFindListReq, opts ...grpc.CallOption) (*SchedulingFindListResp, error)
}

type schedulerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerClient(cc grpc.ClientConnInterface) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) AttendanceAdd(ctx context.Context, in *AttendanceAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_AttendanceAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AttendancePatch(ctx context.Context, in *AttendancePatchReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_AttendancePatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AttendanceDelete(ctx context.Context, in *AttendanceDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_AttendanceDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AttendanceUpdate(ctx context.Context, in *AttendanceUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_AttendanceUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AttendanceFindOne(ctx context.Context, in *AttendanceFindOneReq, opts ...grpc.CallOption) (*AttendanceFindOneResp, error) {
	out := new(AttendanceFindOneResp)
	err := c.cc.Invoke(ctx, Scheduler_AttendanceFindOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AttendanceList(ctx context.Context, in *AttendanceListReq, opts ...grpc.CallOption) (*AttendanceListResp, error) {
	out := new(AttendanceListResp)
	err := c.cc.Invoke(ctx, Scheduler_AttendanceList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AttendanceFindOneDay(ctx context.Context, in *AttendanceFindOneDayReq, opts ...grpc.CallOption) (*AttendanceFindOneDayResp, error) {
	out := new(AttendanceFindOneDayResp)
	err := c.cc.Invoke(ctx, Scheduler_AttendanceFindOneDay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AttendanceByDays(ctx context.Context, in *AttendanceByDaysReq, opts ...grpc.CallOption) (*AttendanceByDaysResp, error) {
	out := new(AttendanceByDaysResp)
	err := c.cc.Invoke(ctx, Scheduler_AttendanceByDays_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) TeamTypeAdd(ctx context.Context, in *TeamTypeAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_TeamTypeAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) TeamTypeDelete(ctx context.Context, in *TeamTypeDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_TeamTypeDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) TeamTypeUpdate(ctx context.Context, in *TeamTypeUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_TeamTypeUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) TeamTypeFindList(ctx context.Context, in *TeamTypeFindListReq, opts ...grpc.CallOption) (*TeamTypeFindListResp, error) {
	out := new(TeamTypeFindListResp)
	err := c.cc.Invoke(ctx, Scheduler_TeamTypeFindList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) TeamAdd(ctx context.Context, in *TeamAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_TeamAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) TeamDelete(ctx context.Context, in *TeamDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_TeamDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) TeamUpdate(ctx context.Context, in *TeamUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_TeamUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) TeamFindList(ctx context.Context, in *TeamFindListReq, opts ...grpc.CallOption) (*TeamFindListResp, error) {
	out := new(TeamFindListResp)
	err := c.cc.Invoke(ctx, Scheduler_TeamFindList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SchedulingTypeAdd(ctx context.Context, in *SchedulingTypeAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_SchedulingTypeAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SchedulingTypeDelete(ctx context.Context, in *SchedulingTypeDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_SchedulingTypeDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SchedulingTypeUpdate(ctx context.Context, in *SchedulingTypeUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_SchedulingTypeUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SchedulingTypeFindList(ctx context.Context, in *SchedulingTypeFindListReq, opts ...grpc.CallOption) (*SchedulingTypeFindListResp, error) {
	out := new(SchedulingTypeFindListResp)
	err := c.cc.Invoke(ctx, Scheduler_SchedulingTypeFindList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SchedulingAdd(ctx context.Context, in *SchedulingAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_SchedulingAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SchedulingDelete(ctx context.Context, in *SchedulingDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_SchedulingDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SchedulingUpdate(ctx context.Context, in *SchedulingUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Scheduler_SchedulingUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) SchedulingFindList(ctx context.Context, in *SchedulingFindListReq, opts ...grpc.CallOption) (*SchedulingFindListResp, error) {
	out := new(SchedulingFindListResp)
	err := c.cc.Invoke(ctx, Scheduler_SchedulingFindList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
// All implementations must embed UnimplementedSchedulerServer
// for forward compatibility
type SchedulerServer interface {
	// 考勤
	AttendanceAdd(context.Context, *AttendanceAddReq) (*CommonResp, error)
	// 签退
	AttendancePatch(context.Context, *AttendancePatchReq) (*CommonResp, error)
	AttendanceDelete(context.Context, *AttendanceDeleteReq) (*CommonResp, error)
	AttendanceUpdate(context.Context, *AttendanceUpdateReq) (*CommonResp, error)
	AttendanceFindOne(context.Context, *AttendanceFindOneReq) (*AttendanceFindOneResp, error)
	AttendanceList(context.Context, *AttendanceListReq) (*AttendanceListResp, error)
	// 获取某天考勤
	AttendanceFindOneDay(context.Context, *AttendanceFindOneDayReq) (*AttendanceFindOneDayResp, error)
	// 根据时间段获取每日考勤
	AttendanceByDays(context.Context, *AttendanceByDaysReq) (*AttendanceByDaysResp, error)
	// 部门
	TeamTypeAdd(context.Context, *TeamTypeAddReq) (*CommonResp, error)
	TeamTypeDelete(context.Context, *TeamTypeDeleteReq) (*CommonResp, error)
	TeamTypeUpdate(context.Context, *TeamTypeUpdateReq) (*CommonResp, error)
	TeamTypeFindList(context.Context, *TeamTypeFindListReq) (*TeamTypeFindListResp, error)
	// 部门人员表
	TeamAdd(context.Context, *TeamAddReq) (*CommonResp, error)
	TeamDelete(context.Context, *TeamDeleteReq) (*CommonResp, error)
	TeamUpdate(context.Context, *TeamUpdateReq) (*CommonResp, error)
	TeamFindList(context.Context, *TeamFindListReq) (*TeamFindListResp, error)
	// 排班类型
	SchedulingTypeAdd(context.Context, *SchedulingTypeAddReq) (*CommonResp, error)
	SchedulingTypeDelete(context.Context, *SchedulingTypeDeleteReq) (*CommonResp, error)
	SchedulingTypeUpdate(context.Context, *SchedulingTypeUpdateReq) (*CommonResp, error)
	SchedulingTypeFindList(context.Context, *SchedulingTypeFindListReq) (*SchedulingTypeFindListResp, error)
	// 排班列表
	SchedulingAdd(context.Context, *SchedulingAddReq) (*CommonResp, error)
	SchedulingDelete(context.Context, *SchedulingDeleteReq) (*CommonResp, error)
	SchedulingUpdate(context.Context, *SchedulingUpdateReq) (*CommonResp, error)
	SchedulingFindList(context.Context, *SchedulingFindListReq) (*SchedulingFindListResp, error)
	mustEmbedUnimplementedSchedulerServer()
}

// UnimplementedSchedulerServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (UnimplementedSchedulerServer) AttendanceAdd(context.Context, *AttendanceAddReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttendanceAdd not implemented")
}
func (UnimplementedSchedulerServer) AttendancePatch(context.Context, *AttendancePatchReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttendancePatch not implemented")
}
func (UnimplementedSchedulerServer) AttendanceDelete(context.Context, *AttendanceDeleteReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttendanceDelete not implemented")
}
func (UnimplementedSchedulerServer) AttendanceUpdate(context.Context, *AttendanceUpdateReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttendanceUpdate not implemented")
}
func (UnimplementedSchedulerServer) AttendanceFindOne(context.Context, *AttendanceFindOneReq) (*AttendanceFindOneResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttendanceFindOne not implemented")
}
func (UnimplementedSchedulerServer) AttendanceList(context.Context, *AttendanceListReq) (*AttendanceListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttendanceList not implemented")
}
func (UnimplementedSchedulerServer) AttendanceFindOneDay(context.Context, *AttendanceFindOneDayReq) (*AttendanceFindOneDayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttendanceFindOneDay not implemented")
}
func (UnimplementedSchedulerServer) AttendanceByDays(context.Context, *AttendanceByDaysReq) (*AttendanceByDaysResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttendanceByDays not implemented")
}
func (UnimplementedSchedulerServer) TeamTypeAdd(context.Context, *TeamTypeAddReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamTypeAdd not implemented")
}
func (UnimplementedSchedulerServer) TeamTypeDelete(context.Context, *TeamTypeDeleteReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamTypeDelete not implemented")
}
func (UnimplementedSchedulerServer) TeamTypeUpdate(context.Context, *TeamTypeUpdateReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamTypeUpdate not implemented")
}
func (UnimplementedSchedulerServer) TeamTypeFindList(context.Context, *TeamTypeFindListReq) (*TeamTypeFindListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamTypeFindList not implemented")
}
func (UnimplementedSchedulerServer) TeamAdd(context.Context, *TeamAddReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamAdd not implemented")
}
func (UnimplementedSchedulerServer) TeamDelete(context.Context, *TeamDeleteReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamDelete not implemented")
}
func (UnimplementedSchedulerServer) TeamUpdate(context.Context, *TeamUpdateReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamUpdate not implemented")
}
func (UnimplementedSchedulerServer) TeamFindList(context.Context, *TeamFindListReq) (*TeamFindListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamFindList not implemented")
}
func (UnimplementedSchedulerServer) SchedulingTypeAdd(context.Context, *SchedulingTypeAddReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulingTypeAdd not implemented")
}
func (UnimplementedSchedulerServer) SchedulingTypeDelete(context.Context, *SchedulingTypeDeleteReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulingTypeDelete not implemented")
}
func (UnimplementedSchedulerServer) SchedulingTypeUpdate(context.Context, *SchedulingTypeUpdateReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulingTypeUpdate not implemented")
}
func (UnimplementedSchedulerServer) SchedulingTypeFindList(context.Context, *SchedulingTypeFindListReq) (*SchedulingTypeFindListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulingTypeFindList not implemented")
}
func (UnimplementedSchedulerServer) SchedulingAdd(context.Context, *SchedulingAddReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulingAdd not implemented")
}
func (UnimplementedSchedulerServer) SchedulingDelete(context.Context, *SchedulingDeleteReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulingDelete not implemented")
}
func (UnimplementedSchedulerServer) SchedulingUpdate(context.Context, *SchedulingUpdateReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulingUpdate not implemented")
}
func (UnimplementedSchedulerServer) SchedulingFindList(context.Context, *SchedulingFindListReq) (*SchedulingFindListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SchedulingFindList not implemented")
}
func (UnimplementedSchedulerServer) mustEmbedUnimplementedSchedulerServer() {}

// UnsafeSchedulerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerServer will
// result in compilation errors.
type UnsafeSchedulerServer interface {
	mustEmbedUnimplementedSchedulerServer()
}

func RegisterSchedulerServer(s grpc.ServiceRegistrar, srv SchedulerServer) {
	s.RegisterService(&Scheduler_ServiceDesc, srv)
}

func _Scheduler_AttendanceAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendanceAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AttendanceAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_AttendanceAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AttendanceAdd(ctx, req.(*AttendanceAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AttendancePatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendancePatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AttendancePatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_AttendancePatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AttendancePatch(ctx, req.(*AttendancePatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AttendanceDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendanceDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AttendanceDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_AttendanceDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AttendanceDelete(ctx, req.(*AttendanceDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AttendanceUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendanceUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AttendanceUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_AttendanceUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AttendanceUpdate(ctx, req.(*AttendanceUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AttendanceFindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendanceFindOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AttendanceFindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_AttendanceFindOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AttendanceFindOne(ctx, req.(*AttendanceFindOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AttendanceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendanceListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AttendanceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_AttendanceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AttendanceList(ctx, req.(*AttendanceListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AttendanceFindOneDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendanceFindOneDayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AttendanceFindOneDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_AttendanceFindOneDay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AttendanceFindOneDay(ctx, req.(*AttendanceFindOneDayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AttendanceByDays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttendanceByDaysReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AttendanceByDays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_AttendanceByDays_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AttendanceByDays(ctx, req.(*AttendanceByDaysReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_TeamTypeAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamTypeAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).TeamTypeAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_TeamTypeAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).TeamTypeAdd(ctx, req.(*TeamTypeAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_TeamTypeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamTypeDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).TeamTypeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_TeamTypeDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).TeamTypeDelete(ctx, req.(*TeamTypeDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_TeamTypeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamTypeUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).TeamTypeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_TeamTypeUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).TeamTypeUpdate(ctx, req.(*TeamTypeUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_TeamTypeFindList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamTypeFindListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).TeamTypeFindList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_TeamTypeFindList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).TeamTypeFindList(ctx, req.(*TeamTypeFindListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_TeamAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).TeamAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_TeamAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).TeamAdd(ctx, req.(*TeamAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_TeamDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).TeamDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_TeamDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).TeamDelete(ctx, req.(*TeamDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_TeamUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).TeamUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_TeamUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).TeamUpdate(ctx, req.(*TeamUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_TeamFindList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamFindListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).TeamFindList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_TeamFindList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).TeamFindList(ctx, req.(*TeamFindListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SchedulingTypeAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulingTypeAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SchedulingTypeAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_SchedulingTypeAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SchedulingTypeAdd(ctx, req.(*SchedulingTypeAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SchedulingTypeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulingTypeDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SchedulingTypeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_SchedulingTypeDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SchedulingTypeDelete(ctx, req.(*SchedulingTypeDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SchedulingTypeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulingTypeUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SchedulingTypeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_SchedulingTypeUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SchedulingTypeUpdate(ctx, req.(*SchedulingTypeUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SchedulingTypeFindList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulingTypeFindListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SchedulingTypeFindList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_SchedulingTypeFindList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SchedulingTypeFindList(ctx, req.(*SchedulingTypeFindListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SchedulingAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulingAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SchedulingAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_SchedulingAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SchedulingAdd(ctx, req.(*SchedulingAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SchedulingDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulingDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SchedulingDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_SchedulingDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SchedulingDelete(ctx, req.(*SchedulingDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SchedulingUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulingUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SchedulingUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_SchedulingUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SchedulingUpdate(ctx, req.(*SchedulingUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_SchedulingFindList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulingFindListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).SchedulingFindList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Scheduler_SchedulingFindList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).SchedulingFindList(ctx, req.(*SchedulingFindListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Scheduler_ServiceDesc is the grpc.ServiceDesc for Scheduler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scheduler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedulerclient.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AttendanceAdd",
			Handler:    _Scheduler_AttendanceAdd_Handler,
		},
		{
			MethodName: "AttendancePatch",
			Handler:    _Scheduler_AttendancePatch_Handler,
		},
		{
			MethodName: "AttendanceDelete",
			Handler:    _Scheduler_AttendanceDelete_Handler,
		},
		{
			MethodName: "AttendanceUpdate",
			Handler:    _Scheduler_AttendanceUpdate_Handler,
		},
		{
			MethodName: "AttendanceFindOne",
			Handler:    _Scheduler_AttendanceFindOne_Handler,
		},
		{
			MethodName: "AttendanceList",
			Handler:    _Scheduler_AttendanceList_Handler,
		},
		{
			MethodName: "AttendanceFindOneDay",
			Handler:    _Scheduler_AttendanceFindOneDay_Handler,
		},
		{
			MethodName: "AttendanceByDays",
			Handler:    _Scheduler_AttendanceByDays_Handler,
		},
		{
			MethodName: "TeamTypeAdd",
			Handler:    _Scheduler_TeamTypeAdd_Handler,
		},
		{
			MethodName: "TeamTypeDelete",
			Handler:    _Scheduler_TeamTypeDelete_Handler,
		},
		{
			MethodName: "TeamTypeUpdate",
			Handler:    _Scheduler_TeamTypeUpdate_Handler,
		},
		{
			MethodName: "TeamTypeFindList",
			Handler:    _Scheduler_TeamTypeFindList_Handler,
		},
		{
			MethodName: "TeamAdd",
			Handler:    _Scheduler_TeamAdd_Handler,
		},
		{
			MethodName: "TeamDelete",
			Handler:    _Scheduler_TeamDelete_Handler,
		},
		{
			MethodName: "TeamUpdate",
			Handler:    _Scheduler_TeamUpdate_Handler,
		},
		{
			MethodName: "TeamFindList",
			Handler:    _Scheduler_TeamFindList_Handler,
		},
		{
			MethodName: "SchedulingTypeAdd",
			Handler:    _Scheduler_SchedulingTypeAdd_Handler,
		},
		{
			MethodName: "SchedulingTypeDelete",
			Handler:    _Scheduler_SchedulingTypeDelete_Handler,
		},
		{
			MethodName: "SchedulingTypeUpdate",
			Handler:    _Scheduler_SchedulingTypeUpdate_Handler,
		},
		{
			MethodName: "SchedulingTypeFindList",
			Handler:    _Scheduler_SchedulingTypeFindList_Handler,
		},
		{
			MethodName: "SchedulingAdd",
			Handler:    _Scheduler_SchedulingAdd_Handler,
		},
		{
			MethodName: "SchedulingDelete",
			Handler:    _Scheduler_SchedulingDelete_Handler,
		},
		{
			MethodName: "SchedulingUpdate",
			Handler:    _Scheduler_SchedulingUpdate_Handler,
		},
		{
			MethodName: "SchedulingFindList",
			Handler:    _Scheduler_SchedulingFindList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler.proto",
}