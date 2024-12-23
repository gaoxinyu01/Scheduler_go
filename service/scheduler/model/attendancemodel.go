package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AttendanceModel = (*customAttendanceModel)(nil)

type (
	// AttendanceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAttendanceModel.

	AttendanceModel interface {
		attendanceModel
	}

	customAttendanceModel struct {
		*defaultAttendanceModel
	}
)

// NewAttendanceModel returns a modelx for the database table.
func NewAttendanceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AttendanceModel {
	return &customAttendanceModel{
		defaultAttendanceModel: newAttendanceModel(conn, c, opts...),
	}
}
