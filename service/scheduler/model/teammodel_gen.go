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
	teamFieldNames          = builder.RawFieldNames(&Team{})
	teamRows                = strings.Join(teamFieldNames, ",")
	teamRowsExpectAutoSet   = strings.Join(stringx.Remove(teamFieldNames, "`create_at`", "`create_time`", "`update_at`", "`update_time`"), ",")
	teamRowsWithPlaceHolder = strings.Join(stringx.Remove(teamFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), "=?,") + "=?"

	cacheTeamIdPrefix = "cache:team:id:"
)

type (
	teamModel interface {
		Insert(ctx context.Context, data *Team) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Team, error)
		Update(ctx context.Context, data *Team) error
		Delete(ctx context.Context, id string) error
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindList(current, pageSize int64, nickName, major, position, telephone, TeamTypeId, TenantId string) (*[]TwyTeamUser, error)
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
		FindOneByTeamTypeIdUserId(teamTypeId, userId, tenantId string) (*Team, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Team) (sql.Result, error)
		TransDelete(ctx context.Context, session sqlx.Session, id string) error
		Count(nickName, major, position, telephone, TeamTypeId, TenantId string) int64
	}

	defaultTeamModel struct {
		sqlc.CachedConn
		table string
	}

	Team struct {
		Id          string         `db:"id"`           // 部门人员ID
		CreatedAt   time.Time      `db:"created_at"`   // 创建时间
		UpdatedAt   sql.NullTime   `db:"updated_at"`   // 更新时间
		CreatedName string         `db:"created_name"` // 创建人
		UpdatedName sql.NullString `db:"updated_name"` // 更新人
		UserId      string         `db:"user_id"`      // 用户表ID
		TeamTypeId  string         `db:"team_type_id"` // 部门表ID
		TenantId    string         `db:"tenant_id"`    // 租户ID
	}
	TwyTeamUser struct {
		Id        string         `db:"tid"`       // 部门人员ID
		Uid       string         `db:"uid"`       // 用户ID
		Account   string         `db:"account"`   // 用户名
		NickName  string         `db:"nick_name"` // 昵称
		Major     sql.NullString `db:"major"`     // 专业
		Position  sql.NullString `db:"position"`  // 岗位
		Avatar    sql.NullString `db:"avatar"`    // 用户头像
		Email     string         `db:"email"`     // 用户邮箱
		Telephone string         `db:"telephone"` // 用户电话
		State     int64          `db:"state"`     // 用户电话
	}
)

func newTeamModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTeamModel {
	return &defaultTeamModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`team`",
	}
}

func (m *defaultTeamModel) Delete(ctx context.Context, id string) error {
	teamIdKey := fmt.Sprintf("%s%v", cacheTeamIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, teamIdKey)
	return err
}

func (m *defaultTeamModel) FindOne(ctx context.Context, id string) (*Team, error) {
	teamIdKey := fmt.Sprintf("%s%v", cacheTeamIdPrefix, id)
	var resp Team
	err := m.QueryRowCtx(ctx, &resp, teamIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", teamRows, m.table)
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

func (m *defaultTeamModel) Insert(ctx context.Context, data *Team) (sql.Result, error) {
	teamIdKey := fmt.Sprintf("%s%v", cacheTeamIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, teamRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.CreatedAt, data.UpdatedAt, data.CreatedName, data.UpdatedName, data.UserId, data.TeamTypeId, data.TenantId)
	}, teamIdKey)
	return ret, err
}

func (m *defaultTeamModel) Update(ctx context.Context, data *Team) error {
	teamIdKey := fmt.Sprintf("%s%v", cacheTeamIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, teamRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.CreatedName, data.UpdatedName, data.UserId, data.TeamTypeId, data.TenantId, data.Id)
	}, teamIdKey)
	return err
}

func (m *defaultTeamModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(teamRows).From(m.table)
}

func (m *defaultTeamModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultTeamModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultTeamModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultTeamModel) FindList(current, pageSize int64, nickName, major, position, telephone, TeamTypeId, TenantId string) (*[]TwyTeamUser, error) {
	var resp []TwyTeamUser
	var where string
	var limit string
	if len(nickName) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "sys_user.nick_name", "%"+nickName+"%")
	}
	if len(major) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "sys_user.major", "%"+major+"%")
	}
	if len(position) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "sys_user.position", "%"+position+"%")
	}
	if len(telephone) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "sys_user.telephone", "%"+telephone+"%")
	}
	if len(TeamTypeId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "team.team_type_id", TeamTypeId)
	}
	if len(TenantId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "team.tenant_id", TenantId)
	}

	if current != 0 && pageSize != 0 {
		limit += fmt.Sprintf(" limit %s,%s ", strconv.FormatInt((current-1)*pageSize, 10), strconv.FormatInt(pageSize, 10))
	}
	query := fmt.Sprintf("SELECT "+
		"team.id as tid,"+
		"sys_user.id as uid,"+
		"sys_user.account,"+
		"sys_user.nick_name,"+
		"sys_user.major,"+
		"sys_user.position,"+
		"sys_user.avatar,"+
		"sys_user.email,"+
		"sys_user.telephone,"+
		"sys_user.state  "+
		"FROM "+
		"sys_user "+
		"LEFT JOIN "+
		"team "+
		"ON "+
		"sys_user.id = team.user_id "+
		"WHERE "+
		"sys_user.deleted_at  is null "+
		"%s  ORDER BY sys_user.created_at DESC %s ", where, limit)
	err := m.CachedConn.QueryRowsNoCache(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return &resp, nil
	default:
		return nil, err
	}
}


func (m *defaultTeamModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultTeamModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTeamIdPrefix, primary)
}

func (m *defaultTeamModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", teamRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTeamModel) tableName() string {
	return m.table
}

func (m *defaultTeamModel) FindOneByTeamTypeIdUserId(teamTypeId, userId, tenantId string) (*Team, error) {
	var resp Team
	var where string
	if len(tenantId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "tenant_id", tenantId)
	}
	if len(teamTypeId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "team_type_id", teamTypeId)
	}
	if len(userId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "user_id", userId)
	}
	query := fmt.Sprintf("select %s from %s where 1=1 %s limit 1", teamRows, m.table, where)
	err := m.CachedConn.QueryRowNoCache(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlc.ErrNotFound
	default:
		return nil, err
	}
}


func (m *defaultTeamModel) TransInsert(ctx context.Context, session sqlx.Session, data *Team) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, teamRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, data.Id, data.CreatedAt, data.UpdatedAt, data.CreatedName, data.UpdatedName, data.UserId, data.TeamTypeId, data.TenantId)

	return ret, err
}

func (m *defaultTeamModel) TransDelete(ctx context.Context, session sqlx.Session, id string) error {
	twyTeamIdKey := fmt.Sprintf("%s%v", cacheTeamIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return session.ExecCtx(ctx, query, id)
	}, twyTeamIdKey)
	return err
}

func (m *defaultTeamModel) Count(nickName, major, position, telephone, TeamTypeId, TenantId string) int64 {
	var count int64
	var where string
	if len(nickName) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "sys_user.nick_name", "%"+nickName+"%")
	}
	if len(major) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "sys_user.major", "%"+major+"%")
	}
	if len(position) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "sys_user.position", "%"+position+"%")
	}
	if len(telephone) != 0 {
		where += fmt.Sprintf(" AND %s like '%s'", "sys_user.telephone", "%"+telephone+"%")
	}
	if len(TeamTypeId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "team.team_type_id", TeamTypeId)
	}
	if len(TenantId) != 0 {
		where += fmt.Sprintf(" AND %s = '%s'", "team.tenant_id", TenantId)
	}
	query := fmt.Sprintf("SELECT count(*) as count "+
		"FROM "+
		"sys_user "+
		"LEFT JOIN "+
		"team "+
		"ON "+
		"sys_user.id = team.user_id "+
		"WHERE "+
		"sys_user.deleted_at is null  %s", where)
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