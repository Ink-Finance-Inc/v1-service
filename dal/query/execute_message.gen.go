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

func newExecuteMessage(db *gorm.DB) executeMessage {
	_executeMessage := executeMessage{}

	_executeMessage.executeMessageDo.UseDB(db)
	_executeMessage.executeMessageDo.UseModel(&model.ExecuteMessage{})

	tableName := _executeMessage.executeMessageDo.TableName()
	_executeMessage.ALL = field.NewField(tableName, "*")
	_executeMessage.ID = field.NewInt32(tableName, "id")
	_executeMessage.TopicID = field.NewString(tableName, "topic_id")
	_executeMessage.Address = field.NewString(tableName, "address")
	_executeMessage.ExecProposalID = field.NewString(tableName, "exec_proposal_id")
	_executeMessage.NowProposalID = field.NewString(tableName, "now_proposal_id")
	_executeMessage.Message = field.NewString(tableName, "message")

	_executeMessage.fillFieldMap()

	return _executeMessage
}

type executeMessage struct {
	executeMessageDo executeMessageDo

	ALL            field.Field
	ID             field.Int32
	TopicID        field.String
	Address        field.String
	ExecProposalID field.String
	NowProposalID  field.String
	Message        field.String

	fieldMap map[string]field.Expr
}

func (e executeMessage) Table(newTableName string) *executeMessage {
	e.executeMessageDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e executeMessage) As(alias string) *executeMessage {
	e.executeMessageDo.DO = *(e.executeMessageDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *executeMessage) updateTableName(table string) *executeMessage {
	e.ALL = field.NewField(table, "*")
	e.ID = field.NewInt32(table, "id")
	e.TopicID = field.NewString(table, "topic_id")
	e.Address = field.NewString(table, "address")
	e.ExecProposalID = field.NewString(table, "exec_proposal_id")
	e.NowProposalID = field.NewString(table, "now_proposal_id")
	e.Message = field.NewString(table, "message")

	e.fillFieldMap()

	return e
}

func (e *executeMessage) WithContext(ctx context.Context) *executeMessageDo {
	return e.executeMessageDo.WithContext(ctx)
}

func (e executeMessage) TableName() string { return e.executeMessageDo.TableName() }

func (e executeMessage) Alias() string { return e.executeMessageDo.Alias() }

func (e *executeMessage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *executeMessage) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 6)
	e.fieldMap["id"] = e.ID
	e.fieldMap["topic_id"] = e.TopicID
	e.fieldMap["address"] = e.Address
	e.fieldMap["exec_proposal_id"] = e.ExecProposalID
	e.fieldMap["now_proposal_id"] = e.NowProposalID
	e.fieldMap["message"] = e.Message
}

func (e executeMessage) clone(db *gorm.DB) executeMessage {
	e.executeMessageDo.ReplaceDB(db)
	return e
}

type executeMessageDo struct{ gen.DO }

func (e executeMessageDo) Debug() *executeMessageDo {
	return e.withDO(e.DO.Debug())
}

func (e executeMessageDo) WithContext(ctx context.Context) *executeMessageDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e executeMessageDo) Clauses(conds ...clause.Expression) *executeMessageDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e executeMessageDo) Returning(value interface{}, columns ...string) *executeMessageDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e executeMessageDo) Not(conds ...gen.Condition) *executeMessageDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e executeMessageDo) Or(conds ...gen.Condition) *executeMessageDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e executeMessageDo) Select(conds ...field.Expr) *executeMessageDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e executeMessageDo) Where(conds ...gen.Condition) *executeMessageDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e executeMessageDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *executeMessageDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e executeMessageDo) Order(conds ...field.Expr) *executeMessageDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e executeMessageDo) Distinct(cols ...field.Expr) *executeMessageDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e executeMessageDo) Omit(cols ...field.Expr) *executeMessageDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e executeMessageDo) Join(table schema.Tabler, on ...field.Expr) *executeMessageDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e executeMessageDo) LeftJoin(table schema.Tabler, on ...field.Expr) *executeMessageDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e executeMessageDo) RightJoin(table schema.Tabler, on ...field.Expr) *executeMessageDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e executeMessageDo) Group(cols ...field.Expr) *executeMessageDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e executeMessageDo) Having(conds ...gen.Condition) *executeMessageDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e executeMessageDo) Limit(limit int) *executeMessageDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e executeMessageDo) Offset(offset int) *executeMessageDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e executeMessageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *executeMessageDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e executeMessageDo) Unscoped() *executeMessageDo {
	return e.withDO(e.DO.Unscoped())
}

func (e executeMessageDo) Create(values ...*model.ExecuteMessage) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e executeMessageDo) CreateInBatches(values []*model.ExecuteMessage, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e executeMessageDo) Save(values ...*model.ExecuteMessage) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e executeMessageDo) First() (*model.ExecuteMessage, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteMessage), nil
	}
}

func (e executeMessageDo) Take() (*model.ExecuteMessage, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteMessage), nil
	}
}

func (e executeMessageDo) Last() (*model.ExecuteMessage, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteMessage), nil
	}
}

func (e executeMessageDo) Find() ([]*model.ExecuteMessage, error) {
	result, err := e.DO.Find()
	return result.([]*model.ExecuteMessage), err
}

func (e executeMessageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ExecuteMessage, err error) {
	buf := make([]*model.ExecuteMessage, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e executeMessageDo) FindInBatches(result *[]*model.ExecuteMessage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e executeMessageDo) Attrs(attrs ...field.AssignExpr) *executeMessageDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e executeMessageDo) Assign(attrs ...field.AssignExpr) *executeMessageDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e executeMessageDo) Joins(field field.RelationField) *executeMessageDo {
	return e.withDO(e.DO.Joins(field))
}

func (e executeMessageDo) Preload(field field.RelationField) *executeMessageDo {
	return e.withDO(e.DO.Preload(field))
}

func (e executeMessageDo) FirstOrInit() (*model.ExecuteMessage, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteMessage), nil
	}
}

func (e executeMessageDo) FirstOrCreate() (*model.ExecuteMessage, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ExecuteMessage), nil
	}
}

func (e executeMessageDo) FindByPage(offset int, limit int) (result []*model.ExecuteMessage, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e executeMessageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e *executeMessageDo) withDO(do gen.Dao) *executeMessageDo {
	e.DO = *do.(*gen.DO)
	return e
}
