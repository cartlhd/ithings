package devicemanagelogic

import (
	"context"
	"fmt"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/repo/mysql"
	"github.com/spf13/cast"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceInfoIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeviceInfoIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceInfoIndexLogic {
	return &DeviceInfoIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取设备信息列表
func (l *DeviceInfoIndexLogic) DeviceInfoIndex(in *dm.DeviceInfoIndexReq) (*dm.DeviceInfoIndexResp, error) {
	l.Infof("%s req=%v", utils.FuncName(), utils.Fmt(in))
	var (
		info     []*dm.DeviceInfo
		size     int64
		page     int64 = 1
		pageSize int64 = 20
		err      error
	)
	if !(in.Page == nil || in.Page.Page == 0 || in.Page.Size == 0) {
		page = in.Page.Page
		pageSize = in.Page.Size
	}
	position := "POINT(0 0)"
	if in.Position != nil {
		position = fmt.Sprintf("POINT(%s)",
			cast.ToString(in.Position.Longitude)+" "+cast.ToString(in.Position.Latitude))
	}
	size, err = l.svcCtx.DeviceInfo.CountByFilter(
		l.ctx, mysql.DeviceFilter{
			ProductID:  in.ProductID,
			DeviceName: in.DeviceName,
			Tags:       in.Tags,
			Range:      in.Range,
			Position:   position,
		})
	if err != nil {
		return nil, err
	}
	di, err := l.svcCtx.DeviceInfo.FindByFilter(
		l.ctx, mysql.DeviceFilter{
			ProductID:  in.ProductID,
			DeviceName: in.DeviceName,
			Tags:       in.Tags,
			Range:      in.Range,
			Position:   position,
		}, def.PageInfo{Size: pageSize, Page: page})
	if err != nil {
		return nil, err
	}
	info = make([]*dm.DeviceInfo, 0, len(di))
	for _, v := range di {
		info = append(info, ToDeviceInfo(v))
	}
	return &dm.DeviceInfoIndexResp{List: info, Total: size}, nil
}
