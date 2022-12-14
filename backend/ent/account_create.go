// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"selfregi/ent/account"
	"selfregi/ent/accountdetail"
	"selfregi/ent/visitor"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AccountCreate) SetCreatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AccountCreate) SetUpdatedAt(t time.Time) *AccountCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdatedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AccountCreate) SetDeletedAt(t time.Time) *AccountCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableDeletedAt(t *time.Time) *AccountCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetVisitorID sets the "visitor_id" field.
func (ac *AccountCreate) SetVisitorID(i int) *AccountCreate {
	ac.mutation.SetVisitorID(i)
	return ac
}

// SetNillableVisitorID sets the "visitor_id" field if the given value is not nil.
func (ac *AccountCreate) SetNillableVisitorID(i *int) *AccountCreate {
	if i != nil {
		ac.SetVisitorID(*i)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *AccountCreate) SetStatus(u uint) *AccountCreate {
	ac.mutation.SetStatus(u)
	return ac
}

// SetSubtotal sets the "subtotal" field.
func (ac *AccountCreate) SetSubtotal(i int) *AccountCreate {
	ac.mutation.SetSubtotal(i)
	return ac
}

// SetTotal sets the "total" field.
func (ac *AccountCreate) SetTotal(i int) *AccountCreate {
	ac.mutation.SetTotal(i)
	return ac
}

// SetTax sets the "tax" field.
func (ac *AccountCreate) SetTax(i int) *AccountCreate {
	ac.mutation.SetTax(i)
	return ac
}

// SetTaxRate sets the "tax_rate" field.
func (ac *AccountCreate) SetTaxRate(i int) *AccountCreate {
	ac.mutation.SetTaxRate(i)
	return ac
}

// SetDiscountClass sets the "discount_class" field.
func (ac *AccountCreate) SetDiscountClass(u uint) *AccountCreate {
	ac.mutation.SetDiscountClass(u)
	return ac
}

// SetDiscountRate sets the "discount_rate" field.
func (ac *AccountCreate) SetDiscountRate(d decimal.Decimal) *AccountCreate {
	ac.mutation.SetDiscountRate(d)
	return ac
}

// SetDiscountPrice sets the "discount_price" field.
func (ac *AccountCreate) SetDiscountPrice(i int) *AccountCreate {
	ac.mutation.SetDiscountPrice(i)
	return ac
}

// SetPaidPrice sets the "paid_price" field.
func (ac *AccountCreate) SetPaidPrice(i int) *AccountCreate {
	ac.mutation.SetPaidPrice(i)
	return ac
}

// SetChange sets the "change" field.
func (ac *AccountCreate) SetChange(i int) *AccountCreate {
	ac.mutation.SetChange(i)
	return ac
}

// AddManagedAccountDetailIDs adds the "managed_account_details" edge to the AccountDetail entity by IDs.
func (ac *AccountCreate) AddManagedAccountDetailIDs(ids ...int) *AccountCreate {
	ac.mutation.AddManagedAccountDetailIDs(ids...)
	return ac
}

// AddManagedAccountDetails adds the "managed_account_details" edges to the AccountDetail entity.
func (ac *AccountCreate) AddManagedAccountDetails(a ...*AccountDetail) *AccountCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddManagedAccountDetailIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Visitor entity by ID.
func (ac *AccountCreate) SetOwnerID(id int) *AccountCreate {
	ac.mutation.SetOwnerID(id)
	return ac
}

// SetOwner sets the "owner" edge to the Visitor entity.
func (ac *AccountCreate) SetOwner(v *Visitor) *AccountCreate {
	return ac.SetOwnerID(v.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Account)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AccountMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := account.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := account.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Account.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Account.updated_at"`)}
	}
	if v, ok := ac.mutation.VisitorID(); ok {
		if err := account.VisitorIDValidator(v); err != nil {
			return &ValidationError{Name: "visitor_id", err: fmt.Errorf(`ent: validator failed for field "Account.visitor_id": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Account.status"`)}
	}
	if _, ok := ac.mutation.Subtotal(); !ok {
		return &ValidationError{Name: "subtotal", err: errors.New(`ent: missing required field "Account.subtotal"`)}
	}
	if _, ok := ac.mutation.Total(); !ok {
		return &ValidationError{Name: "total", err: errors.New(`ent: missing required field "Account.total"`)}
	}
	if _, ok := ac.mutation.Tax(); !ok {
		return &ValidationError{Name: "tax", err: errors.New(`ent: missing required field "Account.tax"`)}
	}
	if _, ok := ac.mutation.TaxRate(); !ok {
		return &ValidationError{Name: "tax_rate", err: errors.New(`ent: missing required field "Account.tax_rate"`)}
	}
	if _, ok := ac.mutation.DiscountClass(); !ok {
		return &ValidationError{Name: "discount_class", err: errors.New(`ent: missing required field "Account.discount_class"`)}
	}
	if _, ok := ac.mutation.DiscountRate(); !ok {
		return &ValidationError{Name: "discount_rate", err: errors.New(`ent: missing required field "Account.discount_rate"`)}
	}
	if _, ok := ac.mutation.DiscountPrice(); !ok {
		return &ValidationError{Name: "discount_price", err: errors.New(`ent: missing required field "Account.discount_price"`)}
	}
	if _, ok := ac.mutation.PaidPrice(); !ok {
		return &ValidationError{Name: "paid_price", err: errors.New(`ent: missing required field "Account.paid_price"`)}
	}
	if _, ok := ac.mutation.Change(); !ok {
		return &ValidationError{Name: "change", err: errors.New(`ent: missing required field "Account.change"`)}
	}
	if _, ok := ac.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Account.owner"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: account.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(account.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := ac.mutation.VisitorID(); ok {
		_spec.SetField(account.FieldVisitorID, field.TypeInt, value)
		_node.VisitorID = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeUint, value)
		_node.Status = &value
	}
	if value, ok := ac.mutation.Subtotal(); ok {
		_spec.SetField(account.FieldSubtotal, field.TypeInt, value)
		_node.Subtotal = &value
	}
	if value, ok := ac.mutation.Total(); ok {
		_spec.SetField(account.FieldTotal, field.TypeInt, value)
		_node.Total = &value
	}
	if value, ok := ac.mutation.Tax(); ok {
		_spec.SetField(account.FieldTax, field.TypeInt, value)
		_node.Tax = &value
	}
	if value, ok := ac.mutation.TaxRate(); ok {
		_spec.SetField(account.FieldTaxRate, field.TypeInt, value)
		_node.TaxRate = &value
	}
	if value, ok := ac.mutation.DiscountClass(); ok {
		_spec.SetField(account.FieldDiscountClass, field.TypeUint, value)
		_node.DiscountClass = &value
	}
	if value, ok := ac.mutation.DiscountRate(); ok {
		_spec.SetField(account.FieldDiscountRate, field.TypeFloat64, value)
		_node.DiscountRate = &value
	}
	if value, ok := ac.mutation.DiscountPrice(); ok {
		_spec.SetField(account.FieldDiscountPrice, field.TypeInt, value)
		_node.DiscountPrice = &value
	}
	if value, ok := ac.mutation.PaidPrice(); ok {
		_spec.SetField(account.FieldPaidPrice, field.TypeInt, value)
		_node.PaidPrice = &value
	}
	if value, ok := ac.mutation.Change(); ok {
		_spec.SetField(account.FieldChange, field.TypeInt, value)
		_node.Change = &value
	}
	if nodes := ac.mutation.ManagedAccountDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.ManagedAccountDetailsTable,
			Columns: []string{account.ManagedAccountDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: accountdetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.OwnerTable,
			Columns: []string{account.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: visitor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.visitor_managed_accounts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	builders []*AccountCreate
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
