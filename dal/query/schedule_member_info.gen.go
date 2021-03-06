// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"inkfinance.xyz/dal/model"
)

func newScheduleMemberInfo(db *gorm.DB) scheduleMemberInfo {
	_scheduleMemberInfo := scheduleMemberInfo{}

	_scheduleMemberInfo.scheduleMemberInfoDo.UseDB(db)
	_scheduleMemberInfo.scheduleMemberInfoDo.UseModel(&model.ScheduleMemberInfo{})

	tableName := _scheduleMemberInfo.scheduleMemberInfoDo.TableName()
	_scheduleMemberInfo.ALL = field.NewField(tableName, "*")
	_scheduleMemberInfo.ID = field.NewInt32(tableName, "id")
	_scheduleMemberInfo.ScheduleID = field.NewInt32(tableName, "schedule_id")
	_scheduleMemberInfo.MemberAddr = field.NewString(tableName, "member_addr")
	_scheduleMemberInfo.TokenAddr = field.NewString(tableName, "token_addr")
	_scheduleMemberInfo.CreateTime = field.NewTime(tableName, "create_time")
	_scheduleMemberInfo.UpdateTime = field.NewTime(tableName, "update_time")
	_scheduleMemberInfo.Valid = field.NewInt32(tableName, "valid")
	_scheduleMemberInfo.ListTitle = field.NewString(tableName, "list_title")
	_scheduleMemberInfo.Description = field.NewString(tableName, "description")
	_scheduleMemberInfo.OncePay = field.NewString(tableName, "once_pay")
	_scheduleMemberInfo.DaoAddress = field.NewString(tableName, "dao_address")

	_scheduleMemberInfo.fillFieldMap()

	return _scheduleMemberInfo
}

type scheduleMemberInfo struct {
	scheduleMemberInfoDo scheduleMemberInfoDo

	ALL         field.Field
	ID          field.Int32
	ScheduleID  field.Int32
	MemberAddr  field.String
	TokenAddr   field.String
	CreateTime  field.Time
	UpdateTime  field.Time
	Valid       field.Int32
	ListTitle   field.String
	Description field.String
	OncePay     field.String
	DaoAddress  field.String

	fieldMap map[string]field.Expr
}

func (s scheduleMemberInfo) Table(newTableName string) *scheduleMemberInfo {
	s.scheduleMemberInfoDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s scheduleMemberInfo) As(alias string) *scheduleMemberInfo {
	s.scheduleMemberInfoDo.DO = *(s.scheduleMemberInfoDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *scheduleMemberInfo) updateTableName(table string) *scheduleMemberInfo {
	s.ALL = field.NewField(table, "*")
	s.ID = field.NewInt32(table, "id")
	s.ScheduleID = field.NewInt32(table, "schedule_id")
	s.MemberAddr = field.NewString(table, "member_addr")
	s.TokenAddr = field.NewString(table, "token_addr")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.Valid = field.NewInt32(table, "valid")
	s.ListTitle = field.NewString(table, "list_title")
	s.Description = field.NewString(table, "description")
	s.OncePay = field.NewString(table, "once_pay")
	s.DaoAddress = field.NewString(table, "dao_address")

	s.fillFieldMap()

	return s
}

func (s *scheduleMemberInfo) WithContext(ctx context.Context) *scheduleMemberInfoDo {
	return s.scheduleMemberInfoDo.WithContext(ctx)
}

func (s scheduleMemberInfo) TableName() string { return s.scheduleMemberInfoDo.TableName() }

func (s scheduleMemberInfo) Alias() string { return s.scheduleMemberInfoDo.Alias() }

func (s *scheduleMemberInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *scheduleMemberInfo) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["id"] = s.ID
	s.fieldMap["schedule_id"] = s.ScheduleID
	s.fieldMap["member_addr"] = s.MemberAddr
	s.fieldMap["token_addr"] = s.TokenAddr
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["valid"] = s.Valid
	s.fieldMap["list_title"] = s.ListTitle
	s.fieldMap["description"] = s.Description
	s.fieldMap["once_pay"] = s.OncePay
	s.fieldMap["dao_address"] = s.DaoAddress
}

func (s scheduleMemberInfo) clone(db *gorm.DB) scheduleMemberInfo {
	s.scheduleMemberInfoDo.ReplaceDB(db)
	return s
}

type scheduleMemberInfoDo struct{ gen.DO }

func (s scheduleMemberInfoDo) Debug() *scheduleMemberInfoDo {
	return s.withDO(s.DO.Debug())
}

func (s scheduleMemberInfoDo) WithContext(ctx context.Context) *scheduleMemberInfoDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s scheduleMemberInfoDo) Clauses(conds ...clause.Expression) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s scheduleMemberInfoDo) Returning(value interface{}, columns ...string) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s scheduleMemberInfoDo) Not(conds ...gen.Condition) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s scheduleMemberInfoDo) Or(conds ...gen.Condition) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s scheduleMemberInfoDo) Select(conds ...field.Expr) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s scheduleMemberInfoDo) Where(conds ...gen.Condition) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s scheduleMemberInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *scheduleMemberInfoDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s scheduleMemberInfoDo) Order(conds ...field.Expr) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s scheduleMemberInfoDo) Distinct(cols ...field.Expr) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s scheduleMemberInfoDo) Omit(cols ...field.Expr) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s scheduleMemberInfoDo) Join(table schema.Tabler, on ...field.Expr) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s scheduleMemberInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *scheduleMemberInfoDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s scheduleMemberInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *scheduleMemberInfoDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s scheduleMemberInfoDo) Group(cols ...field.Expr) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s scheduleMemberInfoDo) Having(conds ...gen.Condition) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s scheduleMemberInfoDo) Limit(limit int) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s scheduleMemberInfoDo) Offset(offset int) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s scheduleMemberInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s scheduleMemberInfoDo) Unscoped() *scheduleMemberInfoDo {
	return s.withDO(s.DO.Unscoped())
}

func (s scheduleMemberInfoDo) Create(values ...*model.ScheduleMemberInfo) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s scheduleMemberInfoDo) CreateInBatches(values []*model.ScheduleMemberInfo, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s scheduleMemberInfoDo) Save(values ...*model.ScheduleMemberInfo) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s scheduleMemberInfoDo) First() (*model.ScheduleMemberInfo, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScheduleMemberInfo), nil
	}
}

func (s scheduleMemberInfoDo) Take() (*model.ScheduleMemberInfo, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScheduleMemberInfo), nil
	}
}

func (s scheduleMemberInfoDo) Last() (*model.ScheduleMemberInfo, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScheduleMemberInfo), nil
	}
}

func (s scheduleMemberInfoDo) Find() ([]*model.ScheduleMemberInfo, error) {
	result, err := s.DO.Find()
	return result.([]*model.ScheduleMemberInfo), err
}

func (s scheduleMemberInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ScheduleMemberInfo, err error) {
	buf := make([]*model.ScheduleMemberInfo, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s scheduleMemberInfoDo) FindInBatches(result *[]*model.ScheduleMemberInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s scheduleMemberInfoDo) Attrs(attrs ...field.AssignExpr) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s scheduleMemberInfoDo) Assign(attrs ...field.AssignExpr) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s scheduleMemberInfoDo) Joins(field field.RelationField) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Joins(field))
}

func (s scheduleMemberInfoDo) Preload(field field.RelationField) *scheduleMemberInfoDo {
	return s.withDO(s.DO.Preload(field))
}

func (s scheduleMemberInfoDo) FirstOrInit() (*model.ScheduleMemberInfo, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScheduleMemberInfo), nil
	}
}

func (s scheduleMemberInfoDo) FirstOrCreate() (*model.ScheduleMemberInfo, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ScheduleMemberInfo), nil
	}
}

func (s scheduleMemberInfoDo) FindByPage(offset int, limit int) (result []*model.ScheduleMemberInfo, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s scheduleMemberInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s *scheduleMemberInfoDo) withDO(do gen.Dao) *scheduleMemberInfoDo {
	s.DO = *do.(*gen.DO)
	return s
}
