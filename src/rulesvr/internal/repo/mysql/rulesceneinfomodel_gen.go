// Code generated by goctl. DO NOT EDIT.

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
	ruleSceneInfoFieldNames          = builder.RawFieldNames(&RuleSceneInfo{})
	ruleSceneInfoRows                = strings.Join(ruleSceneInfoFieldNames, ",")
	ruleSceneInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(ruleSceneInfoFieldNames, "`id`", "`createdTime`", "`deletedTime`", "`updatedTime`"), ",")
	ruleSceneInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(ruleSceneInfoFieldNames, "`id`", "`createdTime`", "`deletedTime`", "`updatedTime`"), "=?,") + "=?"
)

type (
	ruleSceneInfoModel interface {
		Insert(ctx context.Context, data *RuleSceneInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RuleSceneInfo, error)
		FindOneByName(ctx context.Context, name string) (*RuleSceneInfo, error)
		Update(ctx context.Context, data *RuleSceneInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRuleSceneInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RuleSceneInfo struct {
		Id          int64          `db:"id"`          // id
		Name        string         `db:"name"`        // 场景名称
		TriggerType string         `db:"triggerType"` // 触发器类型 device: 设备触发 timer: 定时触发 manual:手动触发
		Trigger     sql.NullString `db:"trigger"`     // 触发器内容-根据触发器类型改变
		When        sql.NullString `db:"when"`        // 触发条件
		Then        sql.NullString `db:"then"`        // 满足条件时执行的动作
		Desc        string         `db:"desc"`        // 描述
		State       int64          `db:"state"`       // 状态（1启用 2禁用）
		CreatedTime time.Time      `db:"createdTime"` // 创建时间
		UpdatedTime time.Time      `db:"updatedTime"` // 更新时间
		DeletedTime sql.NullTime   `db:"deletedTime"` // 删除时间，默认为空，表示未删除，非空表示已删除
	}
)

func newRuleSceneInfoModel(conn sqlx.SqlConn) *defaultRuleSceneInfoModel {
	return &defaultRuleSceneInfoModel{
		conn:  conn,
		table: "`rule_scene_info`",
	}
}

func (m *defaultRuleSceneInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRuleSceneInfoModel) FindOne(ctx context.Context, id int64) (*RuleSceneInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ruleSceneInfoRows, m.table)
	var resp RuleSceneInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRuleSceneInfoModel) FindOneByName(ctx context.Context, name string) (*RuleSceneInfo, error) {
	var resp RuleSceneInfo
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", ruleSceneInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRuleSceneInfoModel) Insert(ctx context.Context, data *RuleSceneInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, ruleSceneInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.TriggerType, data.Trigger, data.When, data.Then, data.Desc, data.State)
	return ret, err
}

func (m *defaultRuleSceneInfoModel) Update(ctx context.Context, newData *RuleSceneInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ruleSceneInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Name, newData.TriggerType, newData.Trigger, newData.When, newData.Then, newData.Desc, newData.State, newData.Id)
	return err
}

func (m *defaultRuleSceneInfoModel) tableName() string {
	return m.table
}
