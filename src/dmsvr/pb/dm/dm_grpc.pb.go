// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/dm.proto

package dm

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

// DmClient is the client API for Dm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DmClient interface {
	//新增设备
	DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error)
	//更新设备
	DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error)
	//删除设备
	DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
	//获取设备信息列表
	DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error)
	//获取设备信息详情
	DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error)
	//新增设备
	ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error)
	//更新设备
	ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error)
	//删除设备
	ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
	//获取设备信息列表
	ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error)
	//获取设备信息详情
	ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error)
	//更新产品物模型
	ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error)
	//获取产品物模型
	ProductSchemaRead(ctx context.Context, in *ProductSchemaReadReq, opts ...grpc.CallOption) (*ProductSchema, error)
	//管理产品的固件
	ManageFirmware(ctx context.Context, in *ManageFirmwareReq, opts ...grpc.CallOption) (*Response, error)
	//获取产品固件信息
	GetFirmwareInfo(ctx context.Context, in *GetFirmwareInfoReq, opts ...grpc.CallOption) (*GetFirmwareInfoResp, error)
	//设备登录认证
	LoginAuth(ctx context.Context, in *LoginAuthReq, opts ...grpc.CallOption) (*Response, error)
	//设备操作认证
	AccessAuth(ctx context.Context, in *AccessAuthReq, opts ...grpc.CallOption) (*Response, error)
	//鉴定是否是root账号
	RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Response, error)
	//同步调用设备行为
	SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error)
	//同步调用设备属性
	SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error)
	//获取设备sdk调试日志
	DataSdkLogIndex(ctx context.Context, in *DataSdkLogIndexReq, opts ...grpc.CallOption) (*DataSdkLogIndexResp, error)
	//获取设备调试信息记录登入登出,操作
	DataHubLogIndex(ctx context.Context, in *DataHubLogIndexReq, opts ...grpc.CallOption) (*DataHubLogIndexResp, error)
	//获取设备数据信息
	DataSchemaLatestIndex(ctx context.Context, in *DataSchemaLatestIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error)
	//获取设备数据信息
	DataSchemaLogIndex(ctx context.Context, in *DataSchemaLogIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error)
}

type dmClient struct {
	cc grpc.ClientConnInterface
}

func NewDmClient(cc grpc.ClientConnInterface) DmClient {
	return &dmClient{cc}
}

func (c *dmClient) DeviceInfoCreate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/DeviceInfoCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) DeviceInfoUpdate(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/DeviceInfoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) DeviceInfoDelete(ctx context.Context, in *DeviceInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/DeviceInfoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) DeviceInfoIndex(ctx context.Context, in *DeviceInfoIndexReq, opts ...grpc.CallOption) (*DeviceInfoIndexResp, error) {
	out := new(DeviceInfoIndexResp)
	err := c.cc.Invoke(ctx, "/dm.Dm/DeviceInfoIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) DeviceInfoRead(ctx context.Context, in *DeviceInfoReadReq, opts ...grpc.CallOption) (*DeviceInfo, error) {
	out := new(DeviceInfo)
	err := c.cc.Invoke(ctx, "/dm.Dm/DeviceInfoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) ProductInfoCreate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/ProductInfoCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) ProductInfoUpdate(ctx context.Context, in *ProductInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/ProductInfoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) ProductInfoDelete(ctx context.Context, in *ProductInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/ProductInfoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) ProductInfoIndex(ctx context.Context, in *ProductInfoIndexReq, opts ...grpc.CallOption) (*ProductInfoIndexResp, error) {
	out := new(ProductInfoIndexResp)
	err := c.cc.Invoke(ctx, "/dm.Dm/ProductInfoIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) ProductInfoRead(ctx context.Context, in *ProductInfoReadReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	out := new(ProductInfo)
	err := c.cc.Invoke(ctx, "/dm.Dm/ProductInfoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) ProductSchemaUpdate(ctx context.Context, in *ProductSchemaUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/ProductSchemaUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) ProductSchemaRead(ctx context.Context, in *ProductSchemaReadReq, opts ...grpc.CallOption) (*ProductSchema, error) {
	out := new(ProductSchema)
	err := c.cc.Invoke(ctx, "/dm.Dm/ProductSchemaRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) ManageFirmware(ctx context.Context, in *ManageFirmwareReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/manageFirmware", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) GetFirmwareInfo(ctx context.Context, in *GetFirmwareInfoReq, opts ...grpc.CallOption) (*GetFirmwareInfoResp, error) {
	out := new(GetFirmwareInfoResp)
	err := c.cc.Invoke(ctx, "/dm.Dm/GetFirmwareInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) LoginAuth(ctx context.Context, in *LoginAuthReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/loginAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) AccessAuth(ctx context.Context, in *AccessAuthReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/accessAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) RootCheck(ctx context.Context, in *RootCheckReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dm.Dm/rootCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error) {
	out := new(SendActionResp)
	err := c.cc.Invoke(ctx, "/dm.Dm/sendAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error) {
	out := new(SendPropertyResp)
	err := c.cc.Invoke(ctx, "/dm.Dm/sendProperty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) DataSdkLogIndex(ctx context.Context, in *DataSdkLogIndexReq, opts ...grpc.CallOption) (*DataSdkLogIndexResp, error) {
	out := new(DataSdkLogIndexResp)
	err := c.cc.Invoke(ctx, "/dm.Dm/dataSdkLogIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) DataHubLogIndex(ctx context.Context, in *DataHubLogIndexReq, opts ...grpc.CallOption) (*DataHubLogIndexResp, error) {
	out := new(DataHubLogIndexResp)
	err := c.cc.Invoke(ctx, "/dm.Dm/dataHubLogIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) DataSchemaLatestIndex(ctx context.Context, in *DataSchemaLatestIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error) {
	out := new(DataSchemaIndexResp)
	err := c.cc.Invoke(ctx, "/dm.Dm/dataSchemaLatestIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dmClient) DataSchemaLogIndex(ctx context.Context, in *DataSchemaLogIndexReq, opts ...grpc.CallOption) (*DataSchemaIndexResp, error) {
	out := new(DataSchemaIndexResp)
	err := c.cc.Invoke(ctx, "/dm.Dm/dataSchemaLogIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DmServer is the server API for Dm service.
// All implementations must embed UnimplementedDmServer
// for forward compatibility
type DmServer interface {
	//新增设备
	DeviceInfoCreate(context.Context, *DeviceInfo) (*Response, error)
	//更新设备
	DeviceInfoUpdate(context.Context, *DeviceInfo) (*Response, error)
	//删除设备
	DeviceInfoDelete(context.Context, *DeviceInfoDeleteReq) (*Response, error)
	//获取设备信息列表
	DeviceInfoIndex(context.Context, *DeviceInfoIndexReq) (*DeviceInfoIndexResp, error)
	//获取设备信息详情
	DeviceInfoRead(context.Context, *DeviceInfoReadReq) (*DeviceInfo, error)
	//新增设备
	ProductInfoCreate(context.Context, *ProductInfo) (*Response, error)
	//更新设备
	ProductInfoUpdate(context.Context, *ProductInfo) (*Response, error)
	//删除设备
	ProductInfoDelete(context.Context, *ProductInfoDeleteReq) (*Response, error)
	//获取设备信息列表
	ProductInfoIndex(context.Context, *ProductInfoIndexReq) (*ProductInfoIndexResp, error)
	//获取设备信息详情
	ProductInfoRead(context.Context, *ProductInfoReadReq) (*ProductInfo, error)
	//更新产品物模型
	ProductSchemaUpdate(context.Context, *ProductSchemaUpdateReq) (*Response, error)
	//获取产品物模型
	ProductSchemaRead(context.Context, *ProductSchemaReadReq) (*ProductSchema, error)
	//管理产品的固件
	ManageFirmware(context.Context, *ManageFirmwareReq) (*Response, error)
	//获取产品固件信息
	GetFirmwareInfo(context.Context, *GetFirmwareInfoReq) (*GetFirmwareInfoResp, error)
	//设备登录认证
	LoginAuth(context.Context, *LoginAuthReq) (*Response, error)
	//设备操作认证
	AccessAuth(context.Context, *AccessAuthReq) (*Response, error)
	//鉴定是否是root账号
	RootCheck(context.Context, *RootCheckReq) (*Response, error)
	//同步调用设备行为
	SendAction(context.Context, *SendActionReq) (*SendActionResp, error)
	//同步调用设备属性
	SendProperty(context.Context, *SendPropertyReq) (*SendPropertyResp, error)
	//获取设备sdk调试日志
	DataSdkLogIndex(context.Context, *DataSdkLogIndexReq) (*DataSdkLogIndexResp, error)
	//获取设备调试信息记录登入登出,操作
	DataHubLogIndex(context.Context, *DataHubLogIndexReq) (*DataHubLogIndexResp, error)
	//获取设备数据信息
	DataSchemaLatestIndex(context.Context, *DataSchemaLatestIndexReq) (*DataSchemaIndexResp, error)
	//获取设备数据信息
	DataSchemaLogIndex(context.Context, *DataSchemaLogIndexReq) (*DataSchemaIndexResp, error)
	mustEmbedUnimplementedDmServer()
}

// UnimplementedDmServer must be embedded to have forward compatible implementations.
type UnimplementedDmServer struct {
}

func (UnimplementedDmServer) DeviceInfoCreate(context.Context, *DeviceInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceInfoCreate not implemented")
}
func (UnimplementedDmServer) DeviceInfoUpdate(context.Context, *DeviceInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceInfoUpdate not implemented")
}
func (UnimplementedDmServer) DeviceInfoDelete(context.Context, *DeviceInfoDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceInfoDelete not implemented")
}
func (UnimplementedDmServer) DeviceInfoIndex(context.Context, *DeviceInfoIndexReq) (*DeviceInfoIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceInfoIndex not implemented")
}
func (UnimplementedDmServer) DeviceInfoRead(context.Context, *DeviceInfoReadReq) (*DeviceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeviceInfoRead not implemented")
}
func (UnimplementedDmServer) ProductInfoCreate(context.Context, *ProductInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfoCreate not implemented")
}
func (UnimplementedDmServer) ProductInfoUpdate(context.Context, *ProductInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfoUpdate not implemented")
}
func (UnimplementedDmServer) ProductInfoDelete(context.Context, *ProductInfoDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfoDelete not implemented")
}
func (UnimplementedDmServer) ProductInfoIndex(context.Context, *ProductInfoIndexReq) (*ProductInfoIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfoIndex not implemented")
}
func (UnimplementedDmServer) ProductInfoRead(context.Context, *ProductInfoReadReq) (*ProductInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductInfoRead not implemented")
}
func (UnimplementedDmServer) ProductSchemaUpdate(context.Context, *ProductSchemaUpdateReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductSchemaUpdate not implemented")
}
func (UnimplementedDmServer) ProductSchemaRead(context.Context, *ProductSchemaReadReq) (*ProductSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductSchemaRead not implemented")
}
func (UnimplementedDmServer) ManageFirmware(context.Context, *ManageFirmwareReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageFirmware not implemented")
}
func (UnimplementedDmServer) GetFirmwareInfo(context.Context, *GetFirmwareInfoReq) (*GetFirmwareInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirmwareInfo not implemented")
}
func (UnimplementedDmServer) LoginAuth(context.Context, *LoginAuthReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAuth not implemented")
}
func (UnimplementedDmServer) AccessAuth(context.Context, *AccessAuthReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessAuth not implemented")
}
func (UnimplementedDmServer) RootCheck(context.Context, *RootCheckReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RootCheck not implemented")
}
func (UnimplementedDmServer) SendAction(context.Context, *SendActionReq) (*SendActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAction not implemented")
}
func (UnimplementedDmServer) SendProperty(context.Context, *SendPropertyReq) (*SendPropertyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendProperty not implemented")
}
func (UnimplementedDmServer) DataSdkLogIndex(context.Context, *DataSdkLogIndexReq) (*DataSdkLogIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataSdkLogIndex not implemented")
}
func (UnimplementedDmServer) DataHubLogIndex(context.Context, *DataHubLogIndexReq) (*DataHubLogIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataHubLogIndex not implemented")
}
func (UnimplementedDmServer) DataSchemaLatestIndex(context.Context, *DataSchemaLatestIndexReq) (*DataSchemaIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataSchemaLatestIndex not implemented")
}
func (UnimplementedDmServer) DataSchemaLogIndex(context.Context, *DataSchemaLogIndexReq) (*DataSchemaIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataSchemaLogIndex not implemented")
}
func (UnimplementedDmServer) mustEmbedUnimplementedDmServer() {}

// UnsafeDmServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DmServer will
// result in compilation errors.
type UnsafeDmServer interface {
	mustEmbedUnimplementedDmServer()
}

func RegisterDmServer(s grpc.ServiceRegistrar, srv DmServer) {
	s.RegisterService(&Dm_ServiceDesc, srv)
}

func _Dm_DeviceInfoCreate_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DeviceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).DeviceInfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/DeviceInfoCreate",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).DeviceInfoCreate(ctx, req.(*DeviceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_DeviceInfoUpdate_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DeviceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).DeviceInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/DeviceInfoUpdate",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).DeviceInfoUpdate(ctx, req.(*DeviceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_DeviceInfoDelete_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DeviceInfoDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).DeviceInfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/DeviceInfoDelete",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).DeviceInfoDelete(ctx, req.(*DeviceInfoDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_DeviceInfoIndex_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DeviceInfoIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).DeviceInfoIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/DeviceInfoIndex",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).DeviceInfoIndex(ctx, req.(*DeviceInfoIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_DeviceInfoRead_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DeviceInfoReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).DeviceInfoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/DeviceInfoRead",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).DeviceInfoRead(ctx, req.(*DeviceInfoReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_ProductInfoCreate_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).ProductInfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/ProductInfoCreate",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).ProductInfoCreate(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_ProductInfoUpdate_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ProductInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).ProductInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/ProductInfoUpdate",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).ProductInfoUpdate(ctx, req.(*ProductInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_ProductInfoDelete_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ProductInfoDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).ProductInfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/ProductInfoDelete",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).ProductInfoDelete(ctx, req.(*ProductInfoDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_ProductInfoIndex_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ProductInfoIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).ProductInfoIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/ProductInfoIndex",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).ProductInfoIndex(ctx, req.(*ProductInfoIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_ProductInfoRead_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ProductInfoReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).ProductInfoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/ProductInfoRead",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).ProductInfoRead(ctx, req.(*ProductInfoReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_ProductSchemaUpdate_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ProductSchemaUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).ProductSchemaUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/ProductSchemaUpdate",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).ProductSchemaUpdate(ctx, req.(*ProductSchemaUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_ProductSchemaRead_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ProductSchemaReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).ProductSchemaRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/ProductSchemaRead",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).ProductSchemaRead(ctx, req.(*ProductSchemaReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_ManageFirmware_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(ManageFirmwareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).ManageFirmware(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/manageFirmware",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).ManageFirmware(ctx, req.(*ManageFirmwareReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_GetFirmwareInfo_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetFirmwareInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).GetFirmwareInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/GetFirmwareInfo",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).GetFirmwareInfo(ctx, req.(*GetFirmwareInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_LoginAuth_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(LoginAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).LoginAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/loginAuth",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).LoginAuth(ctx, req.(*LoginAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_AccessAuth_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(AccessAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).AccessAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/accessAuth",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).AccessAuth(ctx, req.(*AccessAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_RootCheck_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(RootCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).RootCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/rootCheck",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).RootCheck(ctx, req.(*RootCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_SendAction_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(SendActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).SendAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/sendAction",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).SendAction(ctx, req.(*SendActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_SendProperty_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(SendPropertyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).SendProperty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/sendProperty",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).SendProperty(ctx, req.(*SendPropertyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_DataSdkLogIndex_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DataSdkLogIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).DataSdkLogIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/dataSdkLogIndex",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).DataSdkLogIndex(ctx, req.(*DataSdkLogIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_DataHubLogIndex_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DataHubLogIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).DataHubLogIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/dataHubLogIndex",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).DataHubLogIndex(ctx, req.(*DataHubLogIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_DataSchemaLatestIndex_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DataSchemaLatestIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).DataSchemaLatestIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/dataSchemaLatestIndex",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).DataSchemaLatestIndex(ctx, req.(*DataSchemaLatestIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dm_DataSchemaLogIndex_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(DataSchemaLogIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DmServer).DataSchemaLogIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dm.Dm/dataSchemaLogIndex",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(DmServer).DataSchemaLogIndex(ctx, req.(*DataSchemaLogIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Dm_ServiceDesc is the grpc.ServiceDesc for Dm service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dm_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dm.Dm",
	HandlerType: (*DmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeviceInfoCreate",
			Handler:    _Dm_DeviceInfoCreate_Handler,
		},
		{
			MethodName: "DeviceInfoUpdate",
			Handler:    _Dm_DeviceInfoUpdate_Handler,
		},
		{
			MethodName: "DeviceInfoDelete",
			Handler:    _Dm_DeviceInfoDelete_Handler,
		},
		{
			MethodName: "DeviceInfoIndex",
			Handler:    _Dm_DeviceInfoIndex_Handler,
		},
		{
			MethodName: "DeviceInfoRead",
			Handler:    _Dm_DeviceInfoRead_Handler,
		},
		{
			MethodName: "ProductInfoCreate",
			Handler:    _Dm_ProductInfoCreate_Handler,
		},
		{
			MethodName: "ProductInfoUpdate",
			Handler:    _Dm_ProductInfoUpdate_Handler,
		},
		{
			MethodName: "ProductInfoDelete",
			Handler:    _Dm_ProductInfoDelete_Handler,
		},
		{
			MethodName: "ProductInfoIndex",
			Handler:    _Dm_ProductInfoIndex_Handler,
		},
		{
			MethodName: "ProductInfoRead",
			Handler:    _Dm_ProductInfoRead_Handler,
		},
		{
			MethodName: "ProductSchemaUpdate",
			Handler:    _Dm_ProductSchemaUpdate_Handler,
		},
		{
			MethodName: "ProductSchemaRead",
			Handler:    _Dm_ProductSchemaRead_Handler,
		},
		{
			MethodName: "manageFirmware",
			Handler:    _Dm_ManageFirmware_Handler,
		},
		{
			MethodName: "GetFirmwareInfo",
			Handler:    _Dm_GetFirmwareInfo_Handler,
		},
		{
			MethodName: "loginAuth",
			Handler:    _Dm_LoginAuth_Handler,
		},
		{
			MethodName: "accessAuth",
			Handler:    _Dm_AccessAuth_Handler,
		},
		{
			MethodName: "rootCheck",
			Handler:    _Dm_RootCheck_Handler,
		},
		{
			MethodName: "sendAction",
			Handler:    _Dm_SendAction_Handler,
		},
		{
			MethodName: "sendProperty",
			Handler:    _Dm_SendProperty_Handler,
		},
		{
			MethodName: "dataSdkLogIndex",
			Handler:    _Dm_DataSdkLogIndex_Handler,
		},
		{
			MethodName: "dataHubLogIndex",
			Handler:    _Dm_DataHubLogIndex_Handler,
		},
		{
			MethodName: "dataSchemaLatestIndex",
			Handler:    _Dm_DataSchemaLatestIndex_Handler,
		},
		{
			MethodName: "dataSchemaLogIndex",
			Handler:    _Dm_DataSchemaLogIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dm.proto",
}
