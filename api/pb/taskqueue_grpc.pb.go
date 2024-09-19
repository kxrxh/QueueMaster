// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: api/taskqueue.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TaskQueue_SubmitTask_FullMethodName         = "/taskqueue.TaskQueue/SubmitTask"
	TaskQueue_GetTaskStatus_FullMethodName      = "/taskqueue.TaskQueue/GetTaskStatus"
	TaskQueue_StreamTasksResults_FullMethodName = "/taskqueue.TaskQueue/StreamTasksResults"
)

// TaskQueueClient is the client API for TaskQueue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskQueueClient interface {
	SubmitTask(ctx context.Context, in *SubmitTaskRequest, opts ...grpc.CallOption) (*SubmitTaskResponse, error)
	GetTaskStatus(ctx context.Context, in *GetTaskStatusRequest, opts ...grpc.CallOption) (*GetTaskStatusResponse, error)
	StreamTasksResults(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamTasksResultsRequest, StreamTaskResultResponse], error)
}

type taskQueueClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskQueueClient(cc grpc.ClientConnInterface) TaskQueueClient {
	return &taskQueueClient{cc}
}

func (c *taskQueueClient) SubmitTask(ctx context.Context, in *SubmitTaskRequest, opts ...grpc.CallOption) (*SubmitTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitTaskResponse)
	err := c.cc.Invoke(ctx, TaskQueue_SubmitTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskQueueClient) GetTaskStatus(ctx context.Context, in *GetTaskStatusRequest, opts ...grpc.CallOption) (*GetTaskStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTaskStatusResponse)
	err := c.cc.Invoke(ctx, TaskQueue_GetTaskStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskQueueClient) StreamTasksResults(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamTasksResultsRequest, StreamTaskResultResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TaskQueue_ServiceDesc.Streams[0], TaskQueue_StreamTasksResults_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamTasksResultsRequest, StreamTaskResultResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskQueue_StreamTasksResultsClient = grpc.BidiStreamingClient[StreamTasksResultsRequest, StreamTaskResultResponse]

// TaskQueueServer is the server API for TaskQueue service.
// All implementations must embed UnimplementedTaskQueueServer
// for forward compatibility.
type TaskQueueServer interface {
	SubmitTask(context.Context, *SubmitTaskRequest) (*SubmitTaskResponse, error)
	GetTaskStatus(context.Context, *GetTaskStatusRequest) (*GetTaskStatusResponse, error)
	StreamTasksResults(grpc.BidiStreamingServer[StreamTasksResultsRequest, StreamTaskResultResponse]) error
	mustEmbedUnimplementedTaskQueueServer()
}

// UnimplementedTaskQueueServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskQueueServer struct{}

func (UnimplementedTaskQueueServer) SubmitTask(context.Context, *SubmitTaskRequest) (*SubmitTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTask not implemented")
}
func (UnimplementedTaskQueueServer) GetTaskStatus(context.Context, *GetTaskStatusRequest) (*GetTaskStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskStatus not implemented")
}
func (UnimplementedTaskQueueServer) StreamTasksResults(grpc.BidiStreamingServer[StreamTasksResultsRequest, StreamTaskResultResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamTasksResults not implemented")
}
func (UnimplementedTaskQueueServer) mustEmbedUnimplementedTaskQueueServer() {}
func (UnimplementedTaskQueueServer) testEmbeddedByValue()                   {}

// UnsafeTaskQueueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskQueueServer will
// result in compilation errors.
type UnsafeTaskQueueServer interface {
	mustEmbedUnimplementedTaskQueueServer()
}

func RegisterTaskQueueServer(s grpc.ServiceRegistrar, srv TaskQueueServer) {
	// If the following call pancis, it indicates UnimplementedTaskQueueServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaskQueue_ServiceDesc, srv)
}

func _TaskQueue_SubmitTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskQueueServer).SubmitTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskQueue_SubmitTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskQueueServer).SubmitTask(ctx, req.(*SubmitTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskQueue_GetTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskQueueServer).GetTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskQueue_GetTaskStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskQueueServer).GetTaskStatus(ctx, req.(*GetTaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskQueue_StreamTasksResults_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskQueueServer).StreamTasksResults(&grpc.GenericServerStream[StreamTasksResultsRequest, StreamTaskResultResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskQueue_StreamTasksResultsServer = grpc.BidiStreamingServer[StreamTasksResultsRequest, StreamTaskResultResponse]

// TaskQueue_ServiceDesc is the grpc.ServiceDesc for TaskQueue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskQueue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "taskqueue.TaskQueue",
	HandlerType: (*TaskQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitTask",
			Handler:    _TaskQueue_SubmitTask_Handler,
		},
		{
			MethodName: "GetTaskStatus",
			Handler:    _TaskQueue_GetTaskStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTasksResults",
			Handler:       _TaskQueue_StreamTasksResults_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/taskqueue.proto",
}
