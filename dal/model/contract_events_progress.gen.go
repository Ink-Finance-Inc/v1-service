// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameContractEventsProgress = "contract_events_progress"

// ContractEventsProgress mapped from table <contract_events_progress>
type ContractEventsProgress struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ContractAddress string    `gorm:"column:contract_address" json:"contract_address"`
	Events          string    `gorm:"column:events" json:"events"`
	StartHeight     int64     `gorm:"column:start_height" json:"start_height"`
	CurrentHeight   int64     `gorm:"column:current_height" json:"current_height"`
	Valid           int32     `gorm:"column:valid" json:"valid"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	Comments        string    `gorm:"column:comments" json:"comments"`
}

// TableName ContractEventsProgress's table name
func (*ContractEventsProgress) TableName() string {
	return TableNameContractEventsProgress
}
