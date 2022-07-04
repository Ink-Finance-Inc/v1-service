// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameScheduleInfo = "schedule_info"

// ScheduleInfo mapped from table <schedule_info>
type ScheduleInfo struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ScheduleID    int32     `gorm:"column:schedule_id" json:"schedule_id"`
	Duration      int64     `gorm:"column:duration" json:"duration"`
	ScheduleTimes int64     `gorm:"column:schedule_times" json:"schedule_times"`
	StartTime     int64     `gorm:"column:start_time" json:"start_time"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
	DaoAddress    string    `gorm:"column:dao_address" json:"dao_address"`
}

// TableName ScheduleInfo's table name
func (*ScheduleInfo) TableName() string {
	return TableNameScheduleInfo
}