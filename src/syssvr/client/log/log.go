// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package client

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckTokenReq     = sys.CheckTokenReq
	CheckTokenResp    = sys.CheckTokenResp
	ConfigResp        = sys.ConfigResp
	DateRange         = sys.DateRange
	JwtToken          = sys.JwtToken
	LoginLogCreateReq = sys.LoginLogCreateReq
	LoginLogIndexData = sys.LoginLogIndexData
	LoginLogIndexReq  = sys.LoginLogIndexReq
	LoginLogIndexResp = sys.LoginLogIndexResp
	LoginReq          = sys.LoginReq
	LoginResp         = sys.LoginResp
	Map               = sys.Map
	MenuCreateReq     = sys.MenuCreateReq
	MenuData          = sys.MenuData
	MenuDeleteReq     = sys.MenuDeleteReq
	MenuIndexReq      = sys.MenuIndexReq
	MenuIndexResp     = sys.MenuIndexResp
	MenuUpdateReq     = sys.MenuUpdateReq
	OperLogCreateReq  = sys.OperLogCreateReq
	OperLogIndexData  = sys.OperLogIndexData
	OperLogIndexReq   = sys.OperLogIndexReq
	OperLogIndexResp  = sys.OperLogIndexResp
	PageInfo          = sys.PageInfo
	Response          = sys.Response
	RoleCreateReq     = sys.RoleCreateReq
	RoleDeleteReq     = sys.RoleDeleteReq
	RoleIndexData     = sys.RoleIndexData
	RoleIndexReq      = sys.RoleIndexReq
	RoleIndexResp     = sys.RoleIndexResp
	RoleMenuUpdateReq = sys.RoleMenuUpdateReq
	RoleUpdateReq     = sys.RoleUpdateReq
	UserCreateReq     = sys.UserCreateReq
	UserCreateResp    = sys.UserCreateResp
	UserDeleteReq     = sys.UserDeleteReq
	UserIndexReq      = sys.UserIndexReq
	UserIndexResp     = sys.UserIndexResp
	UserInfo          = sys.UserInfo
	UserReadReq       = sys.UserReadReq
	UserReadResp      = sys.UserReadResp
	UserUpdateReq     = sys.UserUpdateReq

	Log interface {
		LoginLogIndex(ctx context.Context, in *LoginLogIndexReq, opts ...grpc.CallOption) (*LoginLogIndexResp, error)
		OperLogIndex(ctx context.Context, in *OperLogIndexReq, opts ...grpc.CallOption) (*OperLogIndexResp, error)
		LoginLogCreate(ctx context.Context, in *LoginLogCreateReq, opts ...grpc.CallOption) (*Response, error)
		OperLogCreate(ctx context.Context, in *OperLogCreateReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultLog struct {
		cli zrpc.Client
	}

	directLog struct {
		svcCtx *svc.ServiceContext
		svr    sys.LogServer
	}
)

func NewLog(cli zrpc.Client) Log {
	return &defaultLog{
		cli: cli,
	}
}

func NewDirectLog(svcCtx *svc.ServiceContext, svr sys.LogServer) Log {
	return &directLog{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultLog) LoginLogIndex(ctx context.Context, in *LoginLogIndexReq, opts ...grpc.CallOption) (*LoginLogIndexResp, error) {
	client := sys.NewLogClient(m.cli.Conn())
	return client.LoginLogIndex(ctx, in, opts...)
}

func (d *directLog) LoginLogIndex(ctx context.Context, in *LoginLogIndexReq, opts ...grpc.CallOption) (*LoginLogIndexResp, error) {
	return d.svr.LoginLogIndex(ctx, in)
}

func (m *defaultLog) OperLogIndex(ctx context.Context, in *OperLogIndexReq, opts ...grpc.CallOption) (*OperLogIndexResp, error) {
	client := sys.NewLogClient(m.cli.Conn())
	return client.OperLogIndex(ctx, in, opts...)
}

func (d *directLog) OperLogIndex(ctx context.Context, in *OperLogIndexReq, opts ...grpc.CallOption) (*OperLogIndexResp, error) {
	return d.svr.OperLogIndex(ctx, in)
}

func (m *defaultLog) LoginLogCreate(ctx context.Context, in *LoginLogCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewLogClient(m.cli.Conn())
	return client.LoginLogCreate(ctx, in, opts...)
}

func (d *directLog) LoginLogCreate(ctx context.Context, in *LoginLogCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.LoginLogCreate(ctx, in)
}

func (m *defaultLog) OperLogCreate(ctx context.Context, in *OperLogCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewLogClient(m.cli.Conn())
	return client.OperLogCreate(ctx, in, opts...)
}

func (d *directLog) OperLogCreate(ctx context.Context, in *OperLogCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.OperLogCreate(ctx, in)
}
