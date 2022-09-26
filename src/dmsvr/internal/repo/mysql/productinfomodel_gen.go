// Code generated by goctl. DO NOT EDIT!

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	productInfoFieldNames          = builder.RawFieldNames(&ProductInfo{})
	productInfoRows                = strings.Join(productInfoFieldNames, ",")
	productInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(productInfoFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	productInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(productInfoFieldNames, "`productID`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
)

type (
	productInfoModel interface {
		Insert(ctx context.Context, data *ProductInfo) (sql.Result, error)
		FindOne(ctx context.Context, productID string) (*ProductInfo, error)
		FindOneByProductName(ctx context.Context, productName string) (*ProductInfo, error)
		Update(ctx context.Context, data *ProductInfo) error
		Delete(ctx context.Context, productID string) error
	}

	defaultProductInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ProductInfo struct {
		ProductID    string       `db:"productID"`    // 产品id
		ProductName  string       `db:"productName"`  // 产品名称
		ProductType  int64        `db:"productType"`  // 产品状态:1:开发中,2:审核中,3:已发布
		AuthMode     int64        `db:"authMode"`     // 认证方式:1:账密认证,2:秘钥认证
		DeviceType   int64        `db:"deviceType"`   // 设备类型:1:设备,2:网关,3:子设备
		CategoryID   int64        `db:"categoryID"`   // 产品品类
		NetType      int64        `db:"netType"`      // 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN
		DataProto    int64        `db:"dataProto"`    // 数据协议:1:自定义,2:数据模板
		AutoRegister int64        `db:"autoRegister"` // 动态注册:1:关闭,2:打开,3:打开并自动创建设备
		Secret       string       `db:"secret"`       // 动态注册产品秘钥
		Desc         string       `db:"desc"`         // 描述
		CreatedTime  time.Time    `db:"createdTime"`
		UpdatedTime  time.Time    `db:"updatedTime"`
		DeletedTime  sql.NullTime `db:"deletedTime"`
		DevStatus    string       `db:"devStatus"` // 产品状态
	}
)

func newProductInfoModel(conn sqlx.SqlConn) *defaultProductInfoModel {
	return &defaultProductInfoModel{
		conn:  conn,
		table: "`product_info`",
	}
}

func (m *defaultProductInfoModel) Delete(ctx context.Context, productID string) error {
	query := fmt.Sprintf("delete from %s where `productID` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, productID)
	return err
}

func (m *defaultProductInfoModel) FindOne(ctx context.Context, productID string) (*ProductInfo, error) {
	query := fmt.Sprintf("select %s from %s where `productID` = ? limit 1", productInfoRows, m.table)
	var resp ProductInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, productID)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductInfoModel) FindOneByProductName(ctx context.Context, productName string) (*ProductInfo, error) {
	var resp ProductInfo
	query := fmt.Sprintf("select %s from %s where `productName` = ? limit 1", productInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, productName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductInfoModel) Insert(ctx context.Context, data *ProductInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, productInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductID, data.ProductName, data.ProductType, data.AuthMode, data.DeviceType, data.CategoryID, data.NetType, data.DataProto, data.AutoRegister, data.Secret, data.Desc, data.CreatedTime, data.UpdatedTime, data.DeletedTime, data.DevStatus)
	return ret, err
}

func (m *defaultProductInfoModel) Update(ctx context.Context, newData *ProductInfo) error {
	query := fmt.Sprintf("update %s set %s where `productID` = ?", m.table, productInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.ProductName, newData.ProductType, newData.AuthMode, newData.DeviceType, newData.CategoryID, newData.NetType, newData.DataProto, newData.AutoRegister, newData.Secret, newData.Desc, newData.CreatedTime, newData.UpdatedTime, newData.DeletedTime, newData.DevStatus, newData.ProductID)
	return err
}

func (m *defaultProductInfoModel) tableName() string {
	return m.table
}
