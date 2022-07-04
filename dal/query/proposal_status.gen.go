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

func newProposalStatus(db *gorm.DB) proposalStatus {
	_proposalStatus := proposalStatus{}

	_proposalStatus.proposalStatusDo.UseDB(db)
	_proposalStatus.proposalStatusDo.UseModel(&model.ProposalStatus{})

	tableName := _proposalStatus.proposalStatusDo.TableName()
	_proposalStatus.ALL = field.NewField(tableName, "*")
	_proposalStatus.ID = field.NewInt32(tableName, "id")
	_proposalStatus.DaoAddress = field.NewString(tableName, "dao_address")
	_proposalStatus.ProposalType = field.NewString(tableName, "proposal_type")
	_proposalStatus.Status = field.NewInt32(tableName, "status")

	_proposalStatus.fillFieldMap()

	return _proposalStatus
}

type proposalStatus struct {
	proposalStatusDo proposalStatusDo

	ALL          field.Field
	ID           field.Int32
	DaoAddress   field.String
	ProposalType field.String
	Status       field.Int32

	fieldMap map[string]field.Expr
}

func (p proposalStatus) Table(newTableName string) *proposalStatus {
	p.proposalStatusDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p proposalStatus) As(alias string) *proposalStatus {
	p.proposalStatusDo.DO = *(p.proposalStatusDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *proposalStatus) updateTableName(table string) *proposalStatus {
	p.ALL = field.NewField(table, "*")
	p.ID = field.NewInt32(table, "id")
	p.DaoAddress = field.NewString(table, "dao_address")
	p.ProposalType = field.NewString(table, "proposal_type")
	p.Status = field.NewInt32(table, "status")

	p.fillFieldMap()

	return p
}

func (p *proposalStatus) WithContext(ctx context.Context) *proposalStatusDo {
	return p.proposalStatusDo.WithContext(ctx)
}

func (p proposalStatus) TableName() string { return p.proposalStatusDo.TableName() }

func (p proposalStatus) Alias() string { return p.proposalStatusDo.Alias() }

func (p *proposalStatus) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *proposalStatus) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["id"] = p.ID
	p.fieldMap["dao_address"] = p.DaoAddress
	p.fieldMap["proposal_type"] = p.ProposalType
	p.fieldMap["status"] = p.Status
}

func (p proposalStatus) clone(db *gorm.DB) proposalStatus {
	p.proposalStatusDo.ReplaceDB(db)
	return p
}

type proposalStatusDo struct{ gen.DO }

func (p proposalStatusDo) Debug() *proposalStatusDo {
	return p.withDO(p.DO.Debug())
}

func (p proposalStatusDo) WithContext(ctx context.Context) *proposalStatusDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p proposalStatusDo) Clauses(conds ...clause.Expression) *proposalStatusDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p proposalStatusDo) Returning(value interface{}, columns ...string) *proposalStatusDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p proposalStatusDo) Not(conds ...gen.Condition) *proposalStatusDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p proposalStatusDo) Or(conds ...gen.Condition) *proposalStatusDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p proposalStatusDo) Select(conds ...field.Expr) *proposalStatusDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p proposalStatusDo) Where(conds ...gen.Condition) *proposalStatusDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p proposalStatusDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *proposalStatusDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p proposalStatusDo) Order(conds ...field.Expr) *proposalStatusDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p proposalStatusDo) Distinct(cols ...field.Expr) *proposalStatusDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p proposalStatusDo) Omit(cols ...field.Expr) *proposalStatusDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p proposalStatusDo) Join(table schema.Tabler, on ...field.Expr) *proposalStatusDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p proposalStatusDo) LeftJoin(table schema.Tabler, on ...field.Expr) *proposalStatusDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p proposalStatusDo) RightJoin(table schema.Tabler, on ...field.Expr) *proposalStatusDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p proposalStatusDo) Group(cols ...field.Expr) *proposalStatusDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p proposalStatusDo) Having(conds ...gen.Condition) *proposalStatusDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p proposalStatusDo) Limit(limit int) *proposalStatusDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p proposalStatusDo) Offset(offset int) *proposalStatusDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p proposalStatusDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *proposalStatusDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p proposalStatusDo) Unscoped() *proposalStatusDo {
	return p.withDO(p.DO.Unscoped())
}

func (p proposalStatusDo) Create(values ...*model.ProposalStatus) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p proposalStatusDo) CreateInBatches(values []*model.ProposalStatus, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p proposalStatusDo) Save(values ...*model.ProposalStatus) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p proposalStatusDo) First() (*model.ProposalStatus, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalStatus), nil
	}
}

func (p proposalStatusDo) Take() (*model.ProposalStatus, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalStatus), nil
	}
}

func (p proposalStatusDo) Last() (*model.ProposalStatus, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalStatus), nil
	}
}

func (p proposalStatusDo) Find() ([]*model.ProposalStatus, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProposalStatus), err
}

func (p proposalStatusDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProposalStatus, err error) {
	buf := make([]*model.ProposalStatus, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p proposalStatusDo) FindInBatches(result *[]*model.ProposalStatus, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p proposalStatusDo) Attrs(attrs ...field.AssignExpr) *proposalStatusDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p proposalStatusDo) Assign(attrs ...field.AssignExpr) *proposalStatusDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p proposalStatusDo) Joins(field field.RelationField) *proposalStatusDo {
	return p.withDO(p.DO.Joins(field))
}

func (p proposalStatusDo) Preload(field field.RelationField) *proposalStatusDo {
	return p.withDO(p.DO.Preload(field))
}

func (p proposalStatusDo) FirstOrInit() (*model.ProposalStatus, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalStatus), nil
	}
}

func (p proposalStatusDo) FirstOrCreate() (*model.ProposalStatus, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalStatus), nil
	}
}

func (p proposalStatusDo) FindByPage(offset int, limit int) (result []*model.ProposalStatus, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p proposalStatusDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p *proposalStatusDo) withDO(do gen.Dao) *proposalStatusDo {
	p.DO = *do.(*gen.DO)
	return p
}