package schemaDataRepo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/store"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg/msgThing"
)

func (d *SchemaDataRepo) InsertEventData(ctx context.Context, productID string,
	deviceName string, event *msgThing.EventData) error {
	param, err := json.Marshal(event.Params)
	if err != nil {
		return errors.System.AddDetail("param json parse failure")
	}
	sql := fmt.Sprintf(
		"insert into %s using %s tags('%s','%s') (`ts`,`eventID`,`eventType`, `param`) values (?,?,?,?);",
		d.GetEventTableName(productID, deviceName), d.GetEventStableName(), productID, deviceName)
	if _, err := d.t.ExecContext(ctx, sql, event.TimeStamp, event.Identifier, event.Type, param); err != nil {
		return err
	}
	return nil
}

func (d *SchemaDataRepo) fmtSql(f msgThing.FilterOpt, sql sq.SelectBuilder) sq.SelectBuilder {
	if f.ProductID != "" {
		sql = sql.Where("`productID`=? ", f.ProductID)
	}
	if len(f.DeviceNames) != 0 {
		sql = sql.Where(fmt.Sprintf("`deviceName` in (%v)", store.ArrayToSql(f.DeviceNames)))
	}
	if f.DataID != "" {
		sql = sql.Where("`eventID`=? ", f.DataID)
	}
	if len(f.Types) != 0 {
		sql = sql.Where(fmt.Sprintf("`eventType` in (%v)", store.ArrayToSql(f.Types)))
	}
	return sql
}

func (d *SchemaDataRepo) GetEventDataByFilter(
	ctx context.Context,
	filter msgThing.FilterOpt) ([]*msgThing.EventData, error) {
	sql := sq.Select("*").From(d.GetEventStableName()).OrderBy("`ts` desc")
	sql = d.fmtSql(filter, sql)
	sql = filter.Page.FmtSql(sql)
	sqlStr, value, err := sql.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := d.t.QueryContext(ctx, sqlStr, value...)
	if err != nil {
		return nil, err
	}
	var datas []map[string]any
	store.Scan(rows, &datas)
	retEvents := make([]*msgThing.EventData, 0, len(datas))
	for _, v := range datas {
		retEvents = append(retEvents, ToEventData(v))
	}
	return retEvents, nil
}

func (d *SchemaDataRepo) GetEventCountByFilter(
	ctx context.Context,
	filter msgThing.FilterOpt) (int64, error) {
	sqSql := sq.Select("count(1)").From(d.GetEventStableName())
	sqSql = d.fmtSql(filter, sqSql)
	sqSql = filter.Page.FmtWhere(sqSql)
	sqlStr, value, err := sqSql.ToSql()
	if err != nil {
		return 0, err
	}
	row := d.t.QueryRowContext(ctx, sqlStr, value...)
	var total int64
	err = row.Scan(&total)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return total, nil
}
