// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.2

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	schedulingFieldNames          = builder.RawFieldNames(&Scheduling{})
	schedulingRows                = strings.Join(schedulingFieldNames, ",")
	schedulingRowsExpectAutoSet   = strings.Join(stringx.Remove(schedulingFieldNames, "`create_at`", "`create_time`", "`update_at`", "`update_time`"), ",")
	schedulingRowsWithPlaceHolder = strings.Join(stringx.Remove(schedulingFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), "=?,") + "=?"

	cacheSchedulingIdPrefix = "cache:scheduling:id:"
)

type (
	schedulingModel interface {
		Insert(ctx context.Context, data *Scheduling) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Scheduling, error)
		Update(ctx context.Context, data *Scheduling) error
		Delete(ctx context.Context, id string) error
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindList(current, pageSize int64, time int64, name string, startTime int64, endTime int64, teamName string, userName string, tenantId string) (*[]SchedulingData, error)
		Count(time int64, name string, startTime int64, endTime int64, teamName string, userName string, tenantId string) int64
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
		TransInsert(ctx context.Context, session sqlx.Session, data *Scheduling) (sql.Result, error)
		FindOneByUserIdAndDayTime(userId string, startTime, endTime int64) int64
		FindListByUserId(userId, tenantId string, isFinished int64) (*[]Scheduling, error)
		TransDelete(ctx context.Context, session sqlx.Session, id string) error
	}

	defaultSchedulingModel struct {
		sqlc.CachedConn
		table string
	}

	Scheduling struct {
		Id           string         `db:"id"`             // 部门人员ID
		CreatedAt    time.Time      `db:"created_at"`     // 创建时间
		UpdatedAt    sql.NullTime   `db:"updated_at"`     // 更新时间
		DeletedAt    sql.NullTime   `db:"deleted_at"`     // 删除时间
		CreatedName  string         `db:"created_name"`   // 创建人
		UpdatedName  sql.NullString `db:"updated_name"`   // 更新人
		DeletedName  sql.NullString `db:"deleted_name"`   // 删除人
		Time         int64          `db:"time"`           // 排班日期
		Name         string         `db:"name"`           // 排班名称
		StartTime    int64          `db:"start_time"`     // 开始时间
		EndTime      int64          `db:"end_time"`       // 结束时间
		TeamName     string         `db:"team_name"`      // 执勤部门
		UserName     string         `db:"user_name"`      // 执勤人
		JobStartTime sql.NullInt64  `db:"job_start_time"` // 上班打卡时间
		JobEndTime   sql.NullInt64  `db:"job_end_time"`   // 下班打卡时间
		Colour       sql.NullString `db:"colour"`         // 颜色
		TeamId       string         `db:"team_id"`        // 部门ID
		UserId       string         `db:"user_id"`        // 用户ID
		TenantId     string         `db:"tenant_id"`      // 租户ID
		IsFinished   int64          `db:"is_finished"`    // 是否完成
	}

	SchedulingData struct {
		Id           sql.NullString `db:"id"`             // 部门人员ID
		TeamName     sql.NullString `db:"team_name"`      // 执勤部门
		CreatedAt    sql.NullTime   `db:"created_at"`     // 创建时间
		UpdatedAt    sql.NullTime   `db:"updated_at"`     // 更新时间
		CreatedName  sql.NullString `db:"created_name"`   // 创建人
		UpdatedName  sql.NullString `db:"updated_name"`   // 更新人
		Time         sql.NullInt64  `db:"time"`           // 排班日期
		Name         sql.NullString `db:"name"`           // 排班名称
		StartTime    sql.NullInt64  `db:"start_time"`     // 开始时间
		EndTime      sql.NullInt64  `db:"end_time"`       // 结束时间
		Colour       sql.NullString `db:"colour"`         // 颜色
		UserName     sql.NullString `db:"user_name"`      // 执勤人
		JobStartTime sql.NullInt64  `db:"job_start_time"` // 上班打卡时间
		JobEndTime   sql.NullInt64  `db:"job_end_time"`   // 下班打卡时间
		IsFinished   sql.NullInt64  `db:"is_finished"`    // 是否完成
	}

	SchedulingAddList struct {
		CreatedName     string           `json:"created_name"`
		Time            int64            `json:"time"`
		Name            string           `json:"name"`
		StartTime       int64            `json:"start_time"`
		EndTime         int64            `json:"end_time"`
		TeamId          string           `json:"team_id"`
		SchedulingUsers []SchedulingUser `json:"scheduling_users"`
		Colour          string           `json:"colour"` // 颜色
		TenantId        string           `json:"tenant_d"`
	}

	SchedulingUser struct {
		UserId   string `json:"user_id"`   // 用户ID
		NickName string `json:"nick_name"` // 昵称
	}
)

func newSchedulingModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSchedulingModel {
	return &defaultSchedulingModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`scheduling`",
	}
}

func (m *defaultSchedulingModel) Delete(ctx context.Context, id string) error {
	schedulingIdKey := fmt.Sprintf("%s%v", cacheSchedulingIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, schedulingIdKey)
	return err
}

func (m *defaultSchedulingModel) FindOne(ctx context.Context, id string) (*Scheduling, error) {
	schedulingIdKey := fmt.Sprintf("%s%v", cacheSchedulingIdPrefix, id)
	var resp Scheduling
	err := m.QueryRowCtx(ctx, &resp, schedulingIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", schedulingRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}

}

func (m *defaultSchedulingModel) Insert(ctx context.Context, data *Scheduling) (sql.Result, error) {
	schedulingIdKey := fmt.Sprintf("%s%v", cacheSchedulingIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, schedulingRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName, data.Time, data.Name, data.StartTime, data.EndTime, data.TeamName, data.UserName, data.JobStartTime, data.JobEndTime, data.Colour, data.TeamId, data.UserId, data.TenantId, data.IsFinished)
	}, schedulingIdKey)
	return ret, err
}

func (m *defaultSchedulingModel) Update(ctx context.Context, data *Scheduling) error {
	schedulingIdKey := fmt.Sprintf("%s%v", cacheSchedulingIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, schedulingRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName, data.Time, data.Name, data.StartTime, data.EndTime, data.TeamName, data.UserName, data.JobStartTime, data.JobEndTime, data.Colour, data.TeamId, data.UserId, data.TenantId, data.IsFinished, data.Id)
	}, schedulingIdKey)
	return err
}

func (m *defaultSchedulingModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(schedulingRows).From(m.table)
}

func (m *defaultSchedulingModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultSchedulingModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultSchedulingModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *defaultSchedulingModel) FindList(current, pageSize int64, time int64, name string, startTime int64, endTime int64, teamName string, userName string, tenantId string) (*[]SchedulingData, error) {
	var resp []SchedulingData
	var where string
	if len(tenantId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "team_type.tenant_id", tenantId)
	}
	if len(name) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "scheduling.name", "%"+name+"%")
	}
	if time != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "scheduling.time", strconv.FormatInt(time, 10))
	}
	if startTime != 0 {
		where += fmt.Sprintf(" AND %s >= '%s'", "scheduling.start_time", strconv.FormatInt(startTime, 10))
	}
	if endTime != 0 {
		where += fmt.Sprintf(" AND %s <= '%s'", "scheduling.end_time", strconv.FormatInt(endTime, 10))
	}
	if len(teamName) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "scheduling.team_name", "%"+teamName+"%")
	}
	if len(userName) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "scheduling.user_name", "%"+userName+"%")
	}
	query := fmt.Sprintf("SELECT "+
		"scheduling.id, "+
		"team_type.`name` as team_name, "+
		"scheduling.created_at, "+
		"scheduling.updated_at, "+
		"scheduling.created_name, "+
		"scheduling.updated_name, "+
		"scheduling.time, "+
		"scheduling.`name`, "+
		"scheduling.start_time, "+
		"scheduling.end_time, "+
		"scheduling.colour, "+
		"scheduling.user_name, "+
		"scheduling.job_start_time, "+
		"scheduling.job_end_time, "+
		"scheduling.is_finished "+
		"FROM "+
		"team_type "+
		"LEFT JOIN "+
		"team "+
		"ON "+
		"team_type.id = team.team_type_id "+
		"LEFT JOIN "+
		"scheduling "+
		"ON "+
		"team.user_id = scheduling.user_id "+
		"WHERE scheduling.deleted_at is null "+
		"%s  ORDER BY scheduling.created_at DESC limit ?,?", where)
	err := m.CachedConn.QueryRowsNoCache(&resp, query, (current-1)*pageSize, pageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlc.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSchedulingModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultSchedulingModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSchedulingIdPrefix, primary)
}

func (m *defaultSchedulingModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", schedulingRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSchedulingModel) tableName() string {
	return m.table
}

func (m *defaultSchedulingModel) TransInsert(ctx context.Context, session sqlx.Session, data *Scheduling) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, schedulingRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, data.Id, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName, data.Time, data.Name, data.StartTime, data.EndTime, data.TeamName, data.UserName, data.JobStartTime, data.JobEndTime, data.Colour, data.TeamId, data.UserId, data.TenantId, data.IsFinished)

	return ret, err
}

func (m *defaultSchedulingModel) FindOneByUserIdAndDayTime(userId string, startTime, endTime int64) int64 {
	var count int64
	var where string
	if len(userId) > 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "user_id", userId)
	}
	if startTime != 0 {
		where += fmt.Sprintf(" AND %s >= '%v'", "time", endTime)
	}
	if endTime != 0 {
		where += fmt.Sprintf(" AND %s <= '%v'", "time", endTime)
	}
	query := fmt.Sprintf("SELECT count(*) as count from %s where 1=1  %s", m.table, where)
	err := m.CachedConn.QueryRowNoCache(&count, query)
	switch err {
	case nil:
		return count
	case sqlc.ErrNotFound:
		return 0
	default:
		return 0
	}
}
func (m *defaultSchedulingModel) TransDelete(ctx context.Context, session sqlx.Session, id string) error {

	twySchedulingIdKey := fmt.Sprintf("%s%v", cacheSchedulingIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return session.ExecCtx(ctx, query, id)
	}, twySchedulingIdKey)
	return err
}

func (m *defaultSchedulingModel) FindListByUserId(userId, tenantId string, isFinished int64) (*[]Scheduling, error) {
	var resp []Scheduling
	var where string
	if len(tenantId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "tenant_id", tenantId)
	}
	if len(userId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "user_id", userId)
	}
	if isFinished != 99 {
		where += fmt.Sprintf(" AND %s = '%s'", "is_finished", strconv.FormatInt(isFinished, 10))
	}
	query := fmt.Sprintf("select %s from %s where  deleted_at is null %s  ORDER BY created_at DESC ", schedulingRows, m.table, where)
	err := m.CachedConn.QueryRowsNoCache(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlc.ErrNotFound
	default:
		return nil, err
	}

}

func (m *defaultSchedulingModel) Count(time int64, name string, startTime int64, endTime int64, teamName string, userName string, tenantId string) int64 {
	var count int64
	var where string
	if len(tenantId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "team_type.tenant_id", tenantId)
	}
	if len(name) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "scheduling.name", "%"+name+"%")
	}
	if time != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "scheduling.time", strconv.FormatInt(time, 10))
	}
	if startTime != 0 {
		where += fmt.Sprintf(" AND %s >= '%s'", "scheduling.start_time", strconv.FormatInt(startTime, 10))
	}
	if endTime != 0 {
		where += fmt.Sprintf(" AND %s <= '%s'", "scheduling.end_time", strconv.FormatInt(endTime, 10))
	}
	if len(teamName) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "scheduling.team_name", "%"+teamName+"%")
	}
	if len(userName) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "scheduling.user_name", "%"+userName+"%")
	}
	query := fmt.Sprintf("SELECT count(team_type.`name`) as count "+
		"FROM "+
		"team_type "+
		"LEFT JOIN "+
		"team "+
		"ON "+
		"team_type.id = team.team_type_id "+
		"LEFT JOIN "+
		"scheduling "+
		"ON "+
		"team.user_id = scheduling.user_id "+
		"WHERE scheduling.deleted_at is null "+
		" %s", where)
	err := m.CachedConn.QueryRowNoCache(&count, query)
	switch err {
	case nil:
		return count
	case sqlc.ErrNotFound:
		return 0
	default:
		return 0
	}
}