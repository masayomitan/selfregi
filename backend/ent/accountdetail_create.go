// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"selfregi/ent/account"
	"selfregi/ent/accountdetail"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// AccountDetailCreate is the builder for creating a AccountDetail entity.
type AccountDetailCreate struct {
	config
	mutation *AccountDetailMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (adc *AccountDetailCreate) SetCreatedAt(t time.Time) *AccountDetailCreate {
	adc.mutation.SetCreatedAt(t)
	return adc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (adc *AccountDetailCreate) SetNillableCreatedAt(t *time.Time) *AccountDetailCreate {
	if t != nil {
		adc.SetCreatedAt(*t)
	}
	return adc
}

// SetUpdatedAt sets the "updated_at" field.
func (adc *AccountDetailCreate) SetUpdatedAt(t time.Time) *AccountDetailCreate {
	adc.mutation.SetUpdatedAt(t)
	return adc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (adc *AccountDetailCreate) SetNillableUpdatedAt(t *time.Time) *AccountDetailCreate {
	if t != nil {
		adc.SetUpdatedAt(*t)
	}
	return adc
}

// SetDeletedAt sets the "deleted_at" field.
func (adc *AccountDetailCreate) SetDeletedAt(t time.Time) *AccountDetailCreate {
	adc.mutation.SetDeletedAt(t)
	return adc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (adc *AccountDetailCreate) SetNillableDeletedAt(t *time.Time) *AccountDetailCreate {
	if t != nil {
		adc.SetDeletedAt(*t)
	}
	return adc
}

// SetAccountID sets the "account_id" field.
func (adc *AccountDetailCreate) SetAccountID(i int) *AccountDetailCreate {
	adc.mutation.SetAccountID(i)
	return adc
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (adc *AccountDetailCreate) SetNillableAccountID(i *int) *AccountDetailCreate {
	if i != nil {
		adc.SetAccountID(*i)
	}
	return adc
}

// SetVisitorID sets the "visitor_id" field.
func (adc *AccountDetailCreate) SetVisitorID(i int) *AccountDetailCreate {
	adc.mutation.SetVisitorID(i)
	return adc
}

// SetNillableVisitorID sets the "visitor_id" field if the given value is not nil.
func (adc *AccountDetailCreate) SetNillableVisitorID(i *int) *AccountDetailCreate {
	if i != nil {
		adc.SetVisitorID(*i)
	}
	return adc
}

// SetProductID sets the "product_id" field.
func (adc *AccountDetailCreate) SetProductID(i int) *AccountDetailCreate {
	adc.mutation.SetProductID(i)
	return adc
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (adc *AccountDetailCreate) SetNillableProductID(i *int) *AccountDetailCreate {
	if i != nil {
		adc.SetProductID(*i)
	}
	return adc
}

// SetCategoryID sets the "category_id" field.
func (adc *AccountDetailCreate) SetCategoryID(i int) *AccountDetailCreate {
	adc.mutation.SetCategoryID(i)
	return adc
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (adc *AccountDetailCreate) SetNillableCategoryID(i *int) *AccountDetailCreate {
	if i != nil {
		adc.SetCategoryID(*i)
	}
	return adc
}

// SetData sets the "data" field.
func (adc *AccountDetailCreate) SetData(s []string) *AccountDetailCreate {
	adc.mutation.SetData(s)
	return adc
}

// SetOptions sets the "options" field.
func (adc *AccountDetailCreate) SetOptions(s []string) *AccountDetailCreate {
	adc.mutation.SetOptions(s)
	return adc
}

// SetQuantity sets the "quantity" field.
func (adc *AccountDetailCreate) SetQuantity(i int) *AccountDetailCreate {
	adc.mutation.SetQuantity(i)
	return adc
}

// SetPrice sets the "price" field.
func (adc *AccountDetailCreate) SetPrice(i int) *AccountDetailCreate {
	adc.mutation.SetPrice(i)
	return adc
}

// SetTax sets the "tax" field.
func (adc *AccountDetailCreate) SetTax(i int) *AccountDetailCreate {
	adc.mutation.SetTax(i)
	return adc
}

// SetTaxRate sets the "tax_rate" field.
func (adc *AccountDetailCreate) SetTaxRate(i int) *AccountDetailCreate {
	adc.mutation.SetTaxRate(i)
	return adc
}

// SetDiscountID sets the "discount_id" field.
func (adc *AccountDetailCreate) SetDiscountID(i int) *AccountDetailCreate {
	adc.mutation.SetDiscountID(i)
	return adc
}

// SetNillableDiscountID sets the "discount_id" field if the given value is not nil.
func (adc *AccountDetailCreate) SetNillableDiscountID(i *int) *AccountDetailCreate {
	if i != nil {
		adc.SetDiscountID(*i)
	}
	return adc
}

// SetDiscountName sets the "discount_name" field.
func (adc *AccountDetailCreate) SetDiscountName(s string) *AccountDetailCreate {
	adc.mutation.SetDiscountName(s)
	return adc
}

// SetDiscountClass sets the "discount_class" field.
func (adc *AccountDetailCreate) SetDiscountClass(u uint) *AccountDetailCreate {
	adc.mutation.SetDiscountClass(u)
	return adc
}

// SetDiscountRate sets the "discount_rate" field.
func (adc *AccountDetailCreate) SetDiscountRate(d decimal.Decimal) *AccountDetailCreate {
	adc.mutation.SetDiscountRate(d)
	return adc
}

// SetDiscountPrice sets the "discount_price" field.
func (adc *AccountDetailCreate) SetDiscountPrice(i int) *AccountDetailCreate {
	adc.mutation.SetDiscountPrice(i)
	return adc
}

// SetManagedAccountID sets the "managed_account" edge to the Account entity by ID.
func (adc *AccountDetailCreate) SetManagedAccountID(id int) *AccountDetailCreate {
	adc.mutation.SetManagedAccountID(id)
	return adc
}

// SetNillableManagedAccountID sets the "managed_account" edge to the Account entity by ID if the given value is not nil.
func (adc *AccountDetailCreate) SetNillableManagedAccountID(id *int) *AccountDetailCreate {
	if id != nil {
		adc = adc.SetManagedAccountID(*id)
	}
	return adc
}

// SetManagedAccount sets the "managed_account" edge to the Account entity.
func (adc *AccountDetailCreate) SetManagedAccount(a *Account) *AccountDetailCreate {
	return adc.SetManagedAccountID(a.ID)
}

// Mutation returns the AccountDetailMutation object of the builder.
func (adc *AccountDetailCreate) Mutation() *AccountDetailMutation {
	return adc.mutation
}

// Save creates the AccountDetail in the database.
func (adc *AccountDetailCreate) Save(ctx context.Context) (*AccountDetail, error) {
	var (
		err  error
		node *AccountDetail
	)
	adc.defaults()
	if len(adc.hooks) == 0 {
		if err = adc.check(); err != nil {
			return nil, err
		}
		node, err = adc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = adc.check(); err != nil {
				return nil, err
			}
			adc.mutation = mutation
			if node, err = adc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(adc.hooks) - 1; i >= 0; i-- {
			if adc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = adc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, adc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AccountDetail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountDetailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (adc *AccountDetailCreate) SaveX(ctx context.Context) *AccountDetail {
	v, err := adc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (adc *AccountDetailCreate) Exec(ctx context.Context) error {
	_, err := adc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adc *AccountDetailCreate) ExecX(ctx context.Context) {
	if err := adc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (adc *AccountDetailCreate) defaults() {
	if _, ok := adc.mutation.CreatedAt(); !ok {
		v := accountdetail.DefaultCreatedAt()
		adc.mutation.SetCreatedAt(v)
	}
	if _, ok := adc.mutation.UpdatedAt(); !ok {
		v := accountdetail.DefaultUpdatedAt()
		adc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (adc *AccountDetailCreate) check() error {
	if _, ok := adc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AccountDetail.created_at"`)}
	}
	if _, ok := adc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AccountDetail.updated_at"`)}
	}
	if v, ok := adc.mutation.AccountID(); ok {
		if err := accountdetail.AccountIDValidator(v); err != nil {
			return &ValidationError{Name: "account_id", err: fmt.Errorf(`ent: validator failed for field "AccountDetail.account_id": %w`, err)}
		}
	}
	if v, ok := adc.mutation.VisitorID(); ok {
		if err := accountdetail.VisitorIDValidator(v); err != nil {
			return &ValidationError{Name: "visitor_id", err: fmt.Errorf(`ent: validator failed for field "AccountDetail.visitor_id": %w`, err)}
		}
	}
	if v, ok := adc.mutation.ProductID(); ok {
		if err := accountdetail.ProductIDValidator(v); err != nil {
			return &ValidationError{Name: "product_id", err: fmt.Errorf(`ent: validator failed for field "AccountDetail.product_id": %w`, err)}
		}
	}
	if v, ok := adc.mutation.CategoryID(); ok {
		if err := accountdetail.CategoryIDValidator(v); err != nil {
			return &ValidationError{Name: "category_id", err: fmt.Errorf(`ent: validator failed for field "AccountDetail.category_id": %w`, err)}
		}
	}
	if _, ok := adc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "AccountDetail.quantity"`)}
	}
	if v, ok := adc.mutation.Quantity(); ok {
		if err := accountdetail.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "AccountDetail.quantity": %w`, err)}
		}
	}
	if _, ok := adc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "AccountDetail.price"`)}
	}
	if _, ok := adc.mutation.Tax(); !ok {
		return &ValidationError{Name: "tax", err: errors.New(`ent: missing required field "AccountDetail.tax"`)}
	}
	if _, ok := adc.mutation.TaxRate(); !ok {
		return &ValidationError{Name: "tax_rate", err: errors.New(`ent: missing required field "AccountDetail.tax_rate"`)}
	}
	if v, ok := adc.mutation.DiscountID(); ok {
		if err := accountdetail.DiscountIDValidator(v); err != nil {
			return &ValidationError{Name: "discount_id", err: fmt.Errorf(`ent: validator failed for field "AccountDetail.discount_id": %w`, err)}
		}
	}
	if _, ok := adc.mutation.DiscountName(); !ok {
		return &ValidationError{Name: "discount_name", err: errors.New(`ent: missing required field "AccountDetail.discount_name"`)}
	}
	if _, ok := adc.mutation.DiscountClass(); !ok {
		return &ValidationError{Name: "discount_class", err: errors.New(`ent: missing required field "AccountDetail.discount_class"`)}
	}
	if _, ok := adc.mutation.DiscountRate(); !ok {
		return &ValidationError{Name: "discount_rate", err: errors.New(`ent: missing required field "AccountDetail.discount_rate"`)}
	}
	if _, ok := adc.mutation.DiscountPrice(); !ok {
		return &ValidationError{Name: "discount_price", err: errors.New(`ent: missing required field "AccountDetail.discount_price"`)}
	}
	return nil
}

func (adc *AccountDetailCreate) sqlSave(ctx context.Context) (*AccountDetail, error) {
	_node, _spec := adc.createSpec()
	if err := sqlgraph.CreateNode(ctx, adc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (adc *AccountDetailCreate) createSpec() (*AccountDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &AccountDetail{config: adc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: accountdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: accountdetail.FieldID,
			},
		}
	)
	if value, ok := adc.mutation.CreatedAt(); ok {
		_spec.SetField(accountdetail.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := adc.mutation.UpdatedAt(); ok {
		_spec.SetField(accountdetail.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := adc.mutation.DeletedAt(); ok {
		_spec.SetField(accountdetail.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := adc.mutation.AccountID(); ok {
		_spec.SetField(accountdetail.FieldAccountID, field.TypeInt, value)
		_node.AccountID = value
	}
	if value, ok := adc.mutation.VisitorID(); ok {
		_spec.SetField(accountdetail.FieldVisitorID, field.TypeInt, value)
		_node.VisitorID = value
	}
	if value, ok := adc.mutation.ProductID(); ok {
		_spec.SetField(accountdetail.FieldProductID, field.TypeInt, value)
		_node.ProductID = value
	}
	if value, ok := adc.mutation.CategoryID(); ok {
		_spec.SetField(accountdetail.FieldCategoryID, field.TypeInt, value)
		_node.CategoryID = value
	}
	if value, ok := adc.mutation.Data(); ok {
		_spec.SetField(accountdetail.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	if value, ok := adc.mutation.Options(); ok {
		_spec.SetField(accountdetail.FieldOptions, field.TypeJSON, value)
		_node.Options = value
	}
	if value, ok := adc.mutation.Quantity(); ok {
		_spec.SetField(accountdetail.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := adc.mutation.Price(); ok {
		_spec.SetField(accountdetail.FieldPrice, field.TypeInt, value)
		_node.Price = &value
	}
	if value, ok := adc.mutation.Tax(); ok {
		_spec.SetField(accountdetail.FieldTax, field.TypeInt, value)
		_node.Tax = &value
	}
	if value, ok := adc.mutation.TaxRate(); ok {
		_spec.SetField(accountdetail.FieldTaxRate, field.TypeInt, value)
		_node.TaxRate = &value
	}
	if value, ok := adc.mutation.DiscountID(); ok {
		_spec.SetField(accountdetail.FieldDiscountID, field.TypeInt, value)
		_node.DiscountID = &value
	}
	if value, ok := adc.mutation.DiscountName(); ok {
		_spec.SetField(accountdetail.FieldDiscountName, field.TypeString, value)
		_node.DiscountName = &value
	}
	if value, ok := adc.mutation.DiscountClass(); ok {
		_spec.SetField(accountdetail.FieldDiscountClass, field.TypeUint, value)
		_node.DiscountClass = &value
	}
	if value, ok := adc.mutation.DiscountRate(); ok {
		_spec.SetField(accountdetail.FieldDiscountRate, field.TypeFloat64, value)
		_node.DiscountRate = &value
	}
	if value, ok := adc.mutation.DiscountPrice(); ok {
		_spec.SetField(accountdetail.FieldDiscountPrice, field.TypeInt, value)
		_node.DiscountPrice = &value
	}
	if nodes := adc.mutation.ManagedAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accountdetail.ManagedAccountTable,
			Columns: []string{accountdetail.ManagedAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_managed_account_details = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccountDetailCreateBulk is the builder for creating many AccountDetail entities in bulk.
type AccountDetailCreateBulk struct {
	config
	builders []*AccountDetailCreate
}

// Save creates the AccountDetail entities in the database.
func (adcb *AccountDetailCreateBulk) Save(ctx context.Context) ([]*AccountDetail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(adcb.builders))
	nodes := make([]*AccountDetail, len(adcb.builders))
	mutators := make([]Mutator, len(adcb.builders))
	for i := range adcb.builders {
		func(i int, root context.Context) {
			builder := adcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountDetailMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, adcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, adcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, adcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (adcb *AccountDetailCreateBulk) SaveX(ctx context.Context) []*AccountDetail {
	v, err := adcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (adcb *AccountDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := adcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adcb *AccountDetailCreateBulk) ExecX(ctx context.Context) {
	if err := adcb.Exec(ctx); err != nil {
		panic(err)
	}
}