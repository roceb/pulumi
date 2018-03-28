// Code generated by protoc-gen-go.
// source: provider.proto
// DO NOT EDIT!

package pulumirpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConfigureRequest struct {
	Variables map[string]string `protobuf:"bytes,1,rep,name=variables" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ConfigureRequest) Reset()                    { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()               {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ConfigureRequest) GetVariables() map[string]string {
	if m != nil {
		return m.Variables
	}
	return nil
}

type InvokeRequest struct {
	Tok  string                   `protobuf:"bytes,1,opt,name=tok" json:"tok,omitempty"`
	Args *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=args" json:"args,omitempty"`
}

func (m *InvokeRequest) Reset()                    { *m = InvokeRequest{} }
func (m *InvokeRequest) String() string            { return proto.CompactTextString(m) }
func (*InvokeRequest) ProtoMessage()               {}
func (*InvokeRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *InvokeRequest) GetTok() string {
	if m != nil {
		return m.Tok
	}
	return ""
}

func (m *InvokeRequest) GetArgs() *google_protobuf1.Struct {
	if m != nil {
		return m.Args
	}
	return nil
}

type InvokeResponse struct {
	Return   *google_protobuf1.Struct `protobuf:"bytes,1,opt,name=return" json:"return,omitempty"`
	Failures []*CheckFailure          `protobuf:"bytes,2,rep,name=failures" json:"failures,omitempty"`
}

func (m *InvokeResponse) Reset()                    { *m = InvokeResponse{} }
func (m *InvokeResponse) String() string            { return proto.CompactTextString(m) }
func (*InvokeResponse) ProtoMessage()               {}
func (*InvokeResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *InvokeResponse) GetReturn() *google_protobuf1.Struct {
	if m != nil {
		return m.Return
	}
	return nil
}

func (m *InvokeResponse) GetFailures() []*CheckFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

type CheckRequest struct {
	Urn  string                   `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	Olds *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=olds" json:"olds,omitempty"`
	News *google_protobuf1.Struct `protobuf:"bytes,3,opt,name=news" json:"news,omitempty"`
}

func (m *CheckRequest) Reset()                    { *m = CheckRequest{} }
func (m *CheckRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()               {}
func (*CheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *CheckRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *CheckRequest) GetOlds() *google_protobuf1.Struct {
	if m != nil {
		return m.Olds
	}
	return nil
}

func (m *CheckRequest) GetNews() *google_protobuf1.Struct {
	if m != nil {
		return m.News
	}
	return nil
}

type CheckResponse struct {
	Inputs   *google_protobuf1.Struct `protobuf:"bytes,1,opt,name=inputs" json:"inputs,omitempty"`
	Failures []*CheckFailure          `protobuf:"bytes,2,rep,name=failures" json:"failures,omitempty"`
}

func (m *CheckResponse) Reset()                    { *m = CheckResponse{} }
func (m *CheckResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()               {}
func (*CheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *CheckResponse) GetInputs() *google_protobuf1.Struct {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *CheckResponse) GetFailures() []*CheckFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

type CheckFailure struct {
	Property string `protobuf:"bytes,1,opt,name=property" json:"property,omitempty"`
	Reason   string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *CheckFailure) Reset()                    { *m = CheckFailure{} }
func (m *CheckFailure) String() string            { return proto.CompactTextString(m) }
func (*CheckFailure) ProtoMessage()               {}
func (*CheckFailure) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *CheckFailure) GetProperty() string {
	if m != nil {
		return m.Property
	}
	return ""
}

func (m *CheckFailure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type DiffRequest struct {
	Id   string                   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Urn  string                   `protobuf:"bytes,2,opt,name=urn" json:"urn,omitempty"`
	Olds *google_protobuf1.Struct `protobuf:"bytes,3,opt,name=olds" json:"olds,omitempty"`
	News *google_protobuf1.Struct `protobuf:"bytes,4,opt,name=news" json:"news,omitempty"`
}

func (m *DiffRequest) Reset()                    { *m = DiffRequest{} }
func (m *DiffRequest) String() string            { return proto.CompactTextString(m) }
func (*DiffRequest) ProtoMessage()               {}
func (*DiffRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *DiffRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DiffRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *DiffRequest) GetOlds() *google_protobuf1.Struct {
	if m != nil {
		return m.Olds
	}
	return nil
}

func (m *DiffRequest) GetNews() *google_protobuf1.Struct {
	if m != nil {
		return m.News
	}
	return nil
}

type DiffResponse struct {
	Replaces            []string `protobuf:"bytes,1,rep,name=replaces" json:"replaces,omitempty"`
	Stables             []string `protobuf:"bytes,2,rep,name=stables" json:"stables,omitempty"`
	DeleteBeforeReplace bool     `protobuf:"varint,3,opt,name=deleteBeforeReplace" json:"deleteBeforeReplace,omitempty"`
}

func (m *DiffResponse) Reset()                    { *m = DiffResponse{} }
func (m *DiffResponse) String() string            { return proto.CompactTextString(m) }
func (*DiffResponse) ProtoMessage()               {}
func (*DiffResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *DiffResponse) GetReplaces() []string {
	if m != nil {
		return m.Replaces
	}
	return nil
}

func (m *DiffResponse) GetStables() []string {
	if m != nil {
		return m.Stables
	}
	return nil
}

func (m *DiffResponse) GetDeleteBeforeReplace() bool {
	if m != nil {
		return m.DeleteBeforeReplace
	}
	return false
}

type CreateRequest struct {
	Urn        string                   `protobuf:"bytes,1,opt,name=urn" json:"urn,omitempty"`
	Properties *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *CreateRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *CreateRequest) GetProperties() *google_protobuf1.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type CreateResponse struct {
	Id         string                   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Properties *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateResponse) GetProperties() *google_protobuf1.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type UpdateRequest struct {
	Id   string                   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Urn  string                   `protobuf:"bytes,2,opt,name=urn" json:"urn,omitempty"`
	Olds *google_protobuf1.Struct `protobuf:"bytes,3,opt,name=olds" json:"olds,omitempty"`
	News *google_protobuf1.Struct `protobuf:"bytes,4,opt,name=news" json:"news,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *UpdateRequest) GetOlds() *google_protobuf1.Struct {
	if m != nil {
		return m.Olds
	}
	return nil
}

func (m *UpdateRequest) GetNews() *google_protobuf1.Struct {
	if m != nil {
		return m.News
	}
	return nil
}

type UpdateResponse struct {
	Properties *google_protobuf1.Struct `protobuf:"bytes,1,opt,name=properties" json:"properties,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *UpdateResponse) GetProperties() *google_protobuf1.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type DeleteRequest struct {
	Id         string                   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Urn        string                   `protobuf:"bytes,2,opt,name=urn" json:"urn,omitempty"`
	Properties *google_protobuf1.Struct `protobuf:"bytes,3,opt,name=properties" json:"properties,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteRequest) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *DeleteRequest) GetProperties() *google_protobuf1.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigureRequest)(nil), "pulumirpc.ConfigureRequest")
	proto.RegisterType((*InvokeRequest)(nil), "pulumirpc.InvokeRequest")
	proto.RegisterType((*InvokeResponse)(nil), "pulumirpc.InvokeResponse")
	proto.RegisterType((*CheckRequest)(nil), "pulumirpc.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "pulumirpc.CheckResponse")
	proto.RegisterType((*CheckFailure)(nil), "pulumirpc.CheckFailure")
	proto.RegisterType((*DiffRequest)(nil), "pulumirpc.DiffRequest")
	proto.RegisterType((*DiffResponse)(nil), "pulumirpc.DiffResponse")
	proto.RegisterType((*CreateRequest)(nil), "pulumirpc.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "pulumirpc.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "pulumirpc.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "pulumirpc.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "pulumirpc.DeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceProvider service

type ResourceProviderClient interface {
	// Configure configures the resource provider with "globals" that control its behavior.
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// Invoke dynamically executes a built-in function in the provider.
	Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error)
	// Check validates that the given property bag is valid for a resource of the given type and returns the inputs
	// that should be passed to successive calls to Diff, Create, or Update for this resource. As a rule, the provider
	// inputs returned by a call to Check should preserve the original representation of the properties as present in
	// the program inputs. Though this rule is not required for correctness, violations thereof can negatively impact
	// the end-user experience, as the provider inputs are using for detecting and rendering diffs.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// Diff checks what impacts a hypothetical update will have on the resource's properties.
	Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error)
	// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
	// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Update updates an existing resource with new values.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PluginInfo, error)
}

type resourceProviderClient struct {
	cc *grpc.ClientConn
}

func NewResourceProviderClient(cc *grpc.ClientConn) ResourceProviderClient {
	return &resourceProviderClient{cc}
}

func (c *resourceProviderClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceProvider/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceProvider/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceProvider/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error) {
	out := new(DiffResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceProvider/Diff", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceProvider/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceProvider/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceProvider/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) GetPluginInfo(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := grpc.Invoke(ctx, "/pulumirpc.ResourceProvider/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceProvider service

type ResourceProviderServer interface {
	// Configure configures the resource provider with "globals" that control its behavior.
	Configure(context.Context, *ConfigureRequest) (*google_protobuf.Empty, error)
	// Invoke dynamically executes a built-in function in the provider.
	Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error)
	// Check validates that the given property bag is valid for a resource of the given type and returns the inputs
	// that should be passed to successive calls to Diff, Create, or Update for this resource. As a rule, the provider
	// inputs returned by a call to Check should preserve the original representation of the properties as present in
	// the program inputs. Though this rule is not required for correctness, violations thereof can negatively impact
	// the end-user experience, as the provider inputs are using for detecting and rendering diffs.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// Diff checks what impacts a hypothetical update will have on the resource's properties.
	Diff(context.Context, *DiffRequest) (*DiffResponse, error)
	// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
	// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Update updates an existing resource with new values.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
	Delete(context.Context, *DeleteRequest) (*google_protobuf.Empty, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(context.Context, *google_protobuf.Empty) (*PluginInfo, error)
}

func RegisterResourceProviderServer(s *grpc.Server, srv ResourceProviderServer) {
	s.RegisterService(&_ResourceProvider_serviceDesc, srv)
}

func _ResourceProvider_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Invoke(ctx, req.(*InvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Diff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Diff(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).GetPluginInfo(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.ResourceProvider",
	HandlerType: (*ResourceProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _ResourceProvider_Configure_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _ResourceProvider_Invoke_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _ResourceProvider_Check_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _ResourceProvider_Diff_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ResourceProvider_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResourceProvider_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResourceProvider_Delete_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _ResourceProvider_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}

func init() { proto.RegisterFile("provider.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xae, 0xe3, 0x34, 0x6f, 0x32, 0x6d, 0xa2, 0x68, 0x5f, 0x68, 0x8d, 0xcb, 0xa1, 0xf2, 0xa9,
	0x02, 0xc9, 0x45, 0xed, 0x81, 0x0f, 0x55, 0x2a, 0xea, 0x07, 0xd0, 0x0b, 0xaa, 0x8c, 0x40, 0x82,
	0x9b, 0x9b, 0x8c, 0x83, 0x89, 0xeb, 0x35, 0xeb, 0x5d, 0xa3, 0xf2, 0x0f, 0x10, 0x37, 0x8e, 0xfc,
	0x5a, 0xe4, 0xdd, 0xb5, 0xeb, 0x6d, 0xda, 0x26, 0xf4, 0xc2, 0xcd, 0xe3, 0xe7, 0x99, 0x8f, 0x67,
	0x76, 0x66, 0x17, 0x06, 0x19, 0xa3, 0x45, 0x3c, 0x46, 0xe6, 0x67, 0x8c, 0x72, 0x4a, 0x7a, 0x99,
	0x48, 0xc4, 0x79, 0xcc, 0xb2, 0x91, 0xbb, 0x9a, 0x25, 0x62, 0x12, 0xa7, 0x0a, 0x70, 0x37, 0x26,
	0x94, 0x4e, 0x12, 0xdc, 0x96, 0xd6, 0x99, 0x88, 0xb6, 0xf1, 0x3c, 0xe3, 0x17, 0x1a, 0x7c, 0x78,
	0x15, 0xcc, 0x39, 0x13, 0x23, 0xae, 0x50, 0xef, 0xb7, 0x05, 0xc3, 0x43, 0x9a, 0x46, 0xf1, 0x44,
	0x30, 0x0c, 0xf0, 0xab, 0xc0, 0x9c, 0x93, 0x37, 0xd0, 0x2b, 0x42, 0x16, 0x87, 0x67, 0x09, 0xe6,
	0x8e, 0xb5, 0x69, 0x6f, 0xad, 0xec, 0x3c, 0xf2, 0xeb, 0xe4, 0xfe, 0x55, 0xbe, 0xff, 0xa1, 0x22,
	0x1f, 0xa7, 0x9c, 0x5d, 0x04, 0x97, 0xce, 0xee, 0x1e, 0x0c, 0x4c, 0x90, 0x0c, 0xc1, 0x9e, 0xe2,
	0x85, 0x63, 0x6d, 0x5a, 0x5b, 0xbd, 0xa0, 0xfc, 0x24, 0xf7, 0x60, 0xb9, 0x08, 0x13, 0x81, 0x4e,
	0x4b, 0xfe, 0x53, 0xc6, 0x8b, 0xd6, 0x33, 0xcb, 0x7b, 0x0b, 0xfd, 0x93, 0xb4, 0xa0, 0xd3, 0xba,
	0xb0, 0x21, 0xd8, 0x9c, 0x4e, 0x2b, 0x67, 0x4e, 0xa7, 0xe4, 0x31, 0xb4, 0x43, 0x36, 0xc9, 0xa5,
	0xef, 0xca, 0xce, 0xba, 0xaf, 0xc4, 0xfa, 0x95, 0x58, 0xff, 0x9d, 0x14, 0x1b, 0x48, 0x92, 0x57,
	0xc0, 0xa0, 0x8a, 0x97, 0x67, 0x34, 0xcd, 0x91, 0x6c, 0x43, 0x87, 0x21, 0x17, 0x2c, 0x95, 0x31,
	0x6f, 0x09, 0xa0, 0x69, 0x64, 0x17, 0xba, 0x51, 0x18, 0x27, 0x82, 0x61, 0x99, 0xd3, 0x96, 0x2e,
	0x8d, 0xce, 0x7c, 0xc6, 0xd1, 0xf4, 0x95, 0xc2, 0x83, 0x9a, 0xe8, 0x7d, 0x87, 0x55, 0x89, 0x34,
	0x64, 0x54, 0x29, 0x7b, 0x41, 0xf9, 0x59, 0xca, 0xa0, 0xc9, 0x78, 0xbe, 0x8c, 0x92, 0x54, 0x92,
	0x53, 0xfc, 0x96, 0x3b, 0xf6, 0x1c, 0x72, 0x49, 0xf2, 0x04, 0xf4, 0x75, 0xee, 0x4b, 0xc9, 0x71,
	0x9a, 0x09, 0x9e, 0xcf, 0x95, 0xac, 0x68, 0x77, 0x93, 0x7c, 0xa0, 0x25, 0x6b, 0x84, 0xb8, 0xd0,
	0xcd, 0x18, 0xcd, 0x90, 0xf1, 0xea, 0xec, 0x6b, 0x9b, 0xac, 0x95, 0x87, 0x10, 0xe6, 0x34, 0xd5,
	0x13, 0xa0, 0x2d, 0xef, 0x87, 0x05, 0x2b, 0x47, 0x71, 0x14, 0x55, 0x6d, 0x1b, 0x40, 0x2b, 0x1e,
	0x6b, 0xef, 0x56, 0x3c, 0xae, 0xda, 0xd8, 0x9a, 0x6d, 0xa3, 0xfd, 0x37, 0x6d, 0x6c, 0x2f, 0xd2,
	0xc6, 0x02, 0x56, 0x55, 0x29, 0xba, 0x8b, 0x2e, 0x74, 0x19, 0x66, 0x49, 0x38, 0xd2, 0x1b, 0xd2,
	0x0b, 0x6a, 0x9b, 0x38, 0xf0, 0x5f, 0xce, 0xd5, 0xf2, 0xb4, 0x24, 0x54, 0x99, 0xe4, 0x09, 0xfc,
	0x3f, 0xc6, 0x04, 0x39, 0x1e, 0x60, 0x44, 0xcb, 0xfd, 0x91, 0x1e, 0xb2, 0xdc, 0x6e, 0x70, 0x1d,
	0xe4, 0x7d, 0x82, 0xfe, 0x21, 0xc3, 0x90, 0xe3, 0xcd, 0xb3, 0xf3, 0x14, 0x40, 0xb7, 0x32, 0xc6,
	0xb9, 0x13, 0xd4, 0xa0, 0x7a, 0x1f, 0x61, 0x50, 0xc5, 0xd6, 0xaa, 0xae, 0x76, 0xf8, 0xce, 0xa1,
	0x7f, 0x5a, 0xd0, 0x7f, 0x9f, 0x8d, 0x1b, 0x75, 0xff, 0xcb, 0xc3, 0x3b, 0x81, 0x41, 0x55, 0x8c,
	0x16, 0x6a, 0x0a, 0xb3, 0x16, 0x17, 0xf6, 0x05, 0xfa, 0x47, 0xf2, 0x98, 0x16, 0xd7, 0x65, 0xe6,
	0xb2, 0x17, 0xce, 0xb5, 0xf3, 0xab, 0x0d, 0xc3, 0x00, 0x73, 0x2a, 0xd8, 0x08, 0x4f, 0xf5, 0x53,
	0x40, 0x0e, 0xa0, 0x57, 0xdf, 0xbf, 0x64, 0xe3, 0x96, 0x5b, 0xd9, 0x5d, 0x9b, 0xc9, 0x71, 0x5c,
	0x3e, 0x0b, 0xde, 0x12, 0xd9, 0x87, 0x8e, 0xba, 0x07, 0x89, 0xd3, 0x08, 0x60, 0x5c, 0xb5, 0xee,
	0x83, 0x6b, 0x10, 0xd5, 0x3c, 0x6f, 0x89, 0xec, 0xc1, 0xb2, 0xdc, 0x6e, 0x32, 0x73, 0x13, 0x54,
	0xee, 0xce, 0x2c, 0x50, 0x7b, 0x3f, 0x87, 0x76, 0xb9, 0x4b, 0x64, 0xad, 0xc1, 0x69, 0xec, 0xb9,
	0xbb, 0x3e, 0xf3, 0xbf, 0x76, 0xdd, 0x87, 0x8e, 0x1a, 0x59, 0xa3, 0x72, 0x63, 0x43, 0x8c, 0xca,
	0xcd, 0xf9, 0x56, 0x01, 0xd4, 0x28, 0x18, 0x01, 0x8c, 0x51, 0x35, 0x02, 0x98, 0x73, 0x23, 0xa5,
	0x77, 0xd4, 0x00, 0x18, 0x01, 0x8c, 0x99, 0xb8, 0xa5, 0xf3, 0x2f, 0xa1, 0xff, 0x1a, 0xf9, 0xa9,
	0x7c, 0xbc, 0x4f, 0xd2, 0x88, 0x92, 0x1b, 0xa8, 0xee, 0xfd, 0x46, 0xf0, 0x4b, 0xba, 0xb7, 0x74,
	0xd6, 0x91, 0xc4, 0xdd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xf1, 0x98, 0x10, 0x1d, 0x08,
	0x00, 0x00,
}
