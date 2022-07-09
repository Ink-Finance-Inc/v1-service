// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameProposalDecision = "proposal_decision"

// ProposalDecision mapped from table <proposal_decision>
type ProposalDecision struct {
	ID                int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	DaoAddress        string    `gorm:"column:dao_address" json:"dao_address"`
	ProposalID        string    `gorm:"column:proposal_id" json:"proposal_id"`
	Subcategory       string    `gorm:"column:subcategory" json:"subcategory"`
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`
	Agree             int32     `gorm:"column:agree" json:"agree"`
	NewProposalHeight int64     `gorm:"column:new_proposal_height" json:"new_proposal_height"`
	AuditPeriod       int32     `gorm:"column:audit_period" json:"audit_period"`
	AgreeTime         int64     `gorm:"column:agree_time" json:"agree_time"`
	Category          string    `gorm:"column:category" json:"category"`
	Expiration        int64     `gorm:"column:expiration" json:"expiration"`
	ProposalTitle     string    `gorm:"column:proposal_title" json:"proposal_title"`
}

// TableName ProposalDecision's table name
func (*ProposalDecision) TableName() string {
	return TableNameProposalDecision
}
