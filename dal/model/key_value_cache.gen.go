// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameKeyValueCache = "key_value_cache"

// KeyValueCache mapped from table <key_value_cache>
type KeyValueCache struct {
	ID      int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	KeyIs   string `gorm:"column:key_is" json:"key_is"`
	ValueIs string `gorm:"column:value_is" json:"value_is"`
}

// TableName KeyValueCache's table name
func (*KeyValueCache) TableName() string {
	return TableNameKeyValueCache
}