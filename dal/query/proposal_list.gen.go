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

func newProposalList(db *gorm.DB) proposalList {
	_proposalList := proposalList{}

	_proposalList.proposalListDo.UseDB(db)
	_proposalList.proposalListDo.UseModel(&model.ProposalList{})

	tableName := _proposalList.proposalListDo.TableName()
	_proposalList.ALL = field.NewField(tableName, "*")
	_proposalList.ID = field.NewInt32(tableName, "id")
	_proposalList.DaoAddress = field.NewString(tableName, "dao_address")
	_proposalList.ProposalID = field.NewString(tableName, "proposal_id")
	_proposalList.BlockHeight = field.NewInt64(tableName, "block_height")
	_proposalList.BlockTime = field.NewInt64(tableName, "block_time")
	_proposalList.Category = field.NewString(tableName, "category")
	_proposalList.SubCategory = field.NewString(tableName, "sub_category")
	_proposalList.Expiration = field.NewInt64(tableName, "expiration")

	_proposalList.fillFieldMap()

	return _proposalList
}

type proposalList struct {
	proposalListDo proposalListDo

	ALL         field.Field
	ID          field.Int32
	DaoAddress  field.String
	ProposalID  field.String
	BlockHeight field.Int64
	BlockTime   field.Int64
	Category    field.String
	SubCategory field.String
	Expiration  field.Int64

	fieldMap map[string]field.Expr
}

func (p proposalList) Table(newTableName string) *proposalList {
	p.proposalListDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p proposalList) As(alias string) *proposalList {
	p.proposalListDo.DO = *(p.proposalListDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *proposalList) updateTableName(table string) *proposalList {
	p.ALL = field.NewField(table, "*")
	p.ID = field.NewInt32(table, "id")
	p.DaoAddress = field.NewString(table, "dao_address")
	p.ProposalID = field.NewString(table, "proposal_id")
	p.BlockHeight = field.NewInt64(table, "block_height")
	p.BlockTime = field.NewInt64(table, "block_time")
	p.Category = field.NewString(table, "category")
	p.SubCategory = field.NewString(table, "sub_category")
	p.Expiration = field.NewInt64(table, "expiration")

	p.fillFieldMap()

	return p
}

func (p *proposalList) WithContext(ctx context.Context) *proposalListDo {
	return p.proposalListDo.WithContext(ctx)
}

func (p proposalList) TableName() string { return p.proposalListDo.TableName() }

func (p proposalList) Alias() string { return p.proposalListDo.Alias() }

func (p *proposalList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *proposalList) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 8)
	p.fieldMap["id"] = p.ID
	p.fieldMap["dao_address"] = p.DaoAddress
	p.fieldMap["proposal_id"] = p.ProposalID
	p.fieldMap["block_height"] = p.BlockHeight
	p.fieldMap["block_time"] = p.BlockTime
	p.fieldMap["category"] = p.Category
	p.fieldMap["sub_category"] = p.SubCategory
	p.fieldMap["expiration"] = p.Expiration
}

func (p proposalList) clone(db *gorm.DB) proposalList {
	p.proposalListDo.ReplaceDB(db)
	return p
}

type proposalListDo struct{ gen.DO }

func (p proposalListDo) Debug() *proposalListDo {
	return p.withDO(p.DO.Debug())
}

func (p proposalListDo) WithContext(ctx context.Context) *proposalListDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p proposalListDo) Clauses(conds ...clause.Expression) *proposalListDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p proposalListDo) Returning(value interface{}, columns ...string) *proposalListDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p proposalListDo) Not(conds ...gen.Condition) *proposalListDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p proposalListDo) Or(conds ...gen.Condition) *proposalListDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p proposalListDo) Select(conds ...field.Expr) *proposalListDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p proposalListDo) Where(conds ...gen.Condition) *proposalListDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p proposalListDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *proposalListDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p proposalListDo) Order(conds ...field.Expr) *proposalListDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p proposalListDo) Distinct(cols ...field.Expr) *proposalListDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p proposalListDo) Omit(cols ...field.Expr) *proposalListDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p proposalListDo) Join(table schema.Tabler, on ...field.Expr) *proposalListDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p proposalListDo) LeftJoin(table schema.Tabler, on ...field.Expr) *proposalListDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p proposalListDo) RightJoin(table schema.Tabler, on ...field.Expr) *proposalListDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p proposalListDo) Group(cols ...field.Expr) *proposalListDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p proposalListDo) Having(conds ...gen.Condition) *proposalListDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p proposalListDo) Limit(limit int) *proposalListDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p proposalListDo) Offset(offset int) *proposalListDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p proposalListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *proposalListDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p proposalListDo) Unscoped() *proposalListDo {
	return p.withDO(p.DO.Unscoped())
}

func (p proposalListDo) Create(values ...*model.ProposalList) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p proposalListDo) CreateInBatches(values []*model.ProposalList, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p proposalListDo) Save(values ...*model.ProposalList) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p proposalListDo) First() (*model.ProposalList, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalList), nil
	}
}

func (p proposalListDo) Take() (*model.ProposalList, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalList), nil
	}
}

func (p proposalListDo) Last() (*model.ProposalList, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalList), nil
	}
}

func (p proposalListDo) Find() ([]*model.ProposalList, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProposalList), err
}

func (p proposalListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProposalList, err error) {
	buf := make([]*model.ProposalList, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p proposalListDo) FindInBatches(result *[]*model.ProposalList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p proposalListDo) Attrs(attrs ...field.AssignExpr) *proposalListDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p proposalListDo) Assign(attrs ...field.AssignExpr) *proposalListDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p proposalListDo) Joins(field field.RelationField) *proposalListDo {
	return p.withDO(p.DO.Joins(field))
}

func (p proposalListDo) Preload(field field.RelationField) *proposalListDo {
	return p.withDO(p.DO.Preload(field))
}

func (p proposalListDo) FirstOrInit() (*model.ProposalList, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalList), nil
	}
}

func (p proposalListDo) FirstOrCreate() (*model.ProposalList, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalList), nil
	}
}

func (p proposalListDo) FindByPage(offset int, limit int) (result []*model.ProposalList, count int64, err error) {
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

func (p proposalListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p *proposalListDo) withDO(do gen.Dao) *proposalListDo {
	p.DO = *do.(*gen.DO)
	return p
}
