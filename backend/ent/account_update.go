// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"selfregi/ent/account"
	"selfregi/ent/accountdetail"
	"selfregi/ent/predicate"
	"selfregi/ent/visitor"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AccountUpdate) SetDeletedAt(t time.Time) *AccountUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDeletedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (au *AccountUpdate) ClearDeletedAt() *AccountUpdate {
	au.mutation.ClearDeletedAt()
	return au
}

// SetVisitorID sets the "visitor_id" field.
func (au *AccountUpdate) SetVisitorID(i int) *AccountUpdate {
	au.mutation.ResetVisitorID()
	au.mutation.SetVisitorID(i)
	return au
}

// SetNillableVisitorID sets the "visitor_id" field if the given value is not nil.
func (au *AccountUpdate) SetNillableVisitorID(i *int) *AccountUpdate {
	if i != nil {
		au.SetVisitorID(*i)
	}
	return au
}

// AddVisitorID adds i to the "visitor_id" field.
func (au *AccountUpdate) AddVisitorID(i int) *AccountUpdate {
	au.mutation.AddVisitorID(i)
	return au
}

// ClearVisitorID clears the value of the "visitor_id" field.
func (au *AccountUpdate) ClearVisitorID() *AccountUpdate {
	au.mutation.ClearVisitorID()
	return au
}

// SetStatus sets the "status" field.
func (au *AccountUpdate) SetStatus(u uint) *AccountUpdate {
	au.mutation.ResetStatus()
	au.mutation.SetStatus(u)
	return au
}

// AddStatus adds u to the "status" field.
func (au *AccountUpdate) AddStatus(u int) *AccountUpdate {
	au.mutation.AddStatus(u)
	return au
}

// SetSubtotal sets the "subtotal" field.
func (au *AccountUpdate) SetSubtotal(i int) *AccountUpdate {
	au.mutation.ResetSubtotal()
	au.mutation.SetSubtotal(i)
	return au
}

// AddSubtotal adds i to the "subtotal" field.
func (au *AccountUpdate) AddSubtotal(i int) *AccountUpdate {
	au.mutation.AddSubtotal(i)
	return au
}

// SetTotal sets the "total" field.
func (au *AccountUpdate) SetTotal(i int) *AccountUpdate {
	au.mutation.ResetTotal()
	au.mutation.SetTotal(i)
	return au
}

// AddTotal adds i to the "total" field.
func (au *AccountUpdate) AddTotal(i int) *AccountUpdate {
	au.mutation.AddTotal(i)
	return au
}

// SetTax sets the "tax" field.
func (au *AccountUpdate) SetTax(i int) *AccountUpdate {
	au.mutation.ResetTax()
	au.mutation.SetTax(i)
	return au
}

// AddTax adds i to the "tax" field.
func (au *AccountUpdate) AddTax(i int) *AccountUpdate {
	au.mutation.AddTax(i)
	return au
}

// SetTaxRate sets the "tax_rate" field.
func (au *AccountUpdate) SetTaxRate(i int) *AccountUpdate {
	au.mutation.ResetTaxRate()
	au.mutation.SetTaxRate(i)
	return au
}

// AddTaxRate adds i to the "tax_rate" field.
func (au *AccountUpdate) AddTaxRate(i int) *AccountUpdate {
	au.mutation.AddTaxRate(i)
	return au
}

// SetDiscountClass sets the "discount_class" field.
func (au *AccountUpdate) SetDiscountClass(u uint) *AccountUpdate {
	au.mutation.ResetDiscountClass()
	au.mutation.SetDiscountClass(u)
	return au
}

// AddDiscountClass adds u to the "discount_class" field.
func (au *AccountUpdate) AddDiscountClass(u int) *AccountUpdate {
	au.mutation.AddDiscountClass(u)
	return au
}

// SetDiscountRate sets the "discount_rate" field.
func (au *AccountUpdate) SetDiscountRate(d decimal.Decimal) *AccountUpdate {
	au.mutation.ResetDiscountRate()
	au.mutation.SetDiscountRate(d)
	return au
}

// AddDiscountRate adds d to the "discount_rate" field.
func (au *AccountUpdate) AddDiscountRate(d decimal.Decimal) *AccountUpdate {
	au.mutation.AddDiscountRate(d)
	return au
}

// SetDiscountPrice sets the "discount_price" field.
func (au *AccountUpdate) SetDiscountPrice(i int) *AccountUpdate {
	au.mutation.ResetDiscountPrice()
	au.mutation.SetDiscountPrice(i)
	return au
}

// AddDiscountPrice adds i to the "discount_price" field.
func (au *AccountUpdate) AddDiscountPrice(i int) *AccountUpdate {
	au.mutation.AddDiscountPrice(i)
	return au
}

// SetPaidPrice sets the "paid_price" field.
func (au *AccountUpdate) SetPaidPrice(i int) *AccountUpdate {
	au.mutation.ResetPaidPrice()
	au.mutation.SetPaidPrice(i)
	return au
}

// AddPaidPrice adds i to the "paid_price" field.
func (au *AccountUpdate) AddPaidPrice(i int) *AccountUpdate {
	au.mutation.AddPaidPrice(i)
	return au
}

// SetChange sets the "change" field.
func (au *AccountUpdate) SetChange(i int) *AccountUpdate {
	au.mutation.ResetChange()
	au.mutation.SetChange(i)
	return au
}

// AddChange adds i to the "change" field.
func (au *AccountUpdate) AddChange(i int) *AccountUpdate {
	au.mutation.AddChange(i)
	return au
}

// AddManagedAccountDetailIDs adds the "managed_account_details" edge to the AccountDetail entity by IDs.
func (au *AccountUpdate) AddManagedAccountDetailIDs(ids ...int) *AccountUpdate {
	au.mutation.AddManagedAccountDetailIDs(ids...)
	return au
}

// AddManagedAccountDetails adds the "managed_account_details" edges to the AccountDetail entity.
func (au *AccountUpdate) AddManagedAccountDetails(a ...*AccountDetail) *AccountUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddManagedAccountDetailIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Visitor entity by ID.
func (au *AccountUpdate) SetOwnerID(id int) *AccountUpdate {
	au.mutation.SetOwnerID(id)
	return au
}

// SetOwner sets the "owner" edge to the Visitor entity.
func (au *AccountUpdate) SetOwner(v *Visitor) *AccountUpdate {
	return au.SetOwnerID(v.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearManagedAccountDetails clears all "managed_account_details" edges to the AccountDetail entity.
func (au *AccountUpdate) ClearManagedAccountDetails() *AccountUpdate {
	au.mutation.ClearManagedAccountDetails()
	return au
}

// RemoveManagedAccountDetailIDs removes the "managed_account_details" edge to AccountDetail entities by IDs.
func (au *AccountUpdate) RemoveManagedAccountDetailIDs(ids ...int) *AccountUpdate {
	au.mutation.RemoveManagedAccountDetailIDs(ids...)
	return au
}

// RemoveManagedAccountDetails removes "managed_account_details" edges to AccountDetail entities.
func (au *AccountUpdate) RemoveManagedAccountDetails(a ...*AccountDetail) *AccountUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveManagedAccountDetailIDs(ids...)
}

// ClearOwner clears the "owner" edge to the Visitor entity.
func (au *AccountUpdate) ClearOwner() *AccountUpdate {
	au.mutation.ClearOwner()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.VisitorID(); ok {
		if err := account.VisitorIDValidator(v); err != nil {
			return &ValidationError{Name: "visitor_id", err: fmt.Errorf(`ent: validator failed for field "Account.visitor_id": %w`, err)}
		}
	}
	if _, ok := au.mutation.OwnerID(); au.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Account.owner"`)
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(account.FieldDeletedAt, field.TypeTime, value)
	}
	if au.mutation.DeletedAtCleared() {
		_spec.ClearField(account.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := au.mutation.VisitorID(); ok {
		_spec.SetField(account.FieldVisitorID, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedVisitorID(); ok {
		_spec.AddField(account.FieldVisitorID, field.TypeInt, value)
	}
	if au.mutation.VisitorIDCleared() {
		_spec.ClearField(account.FieldVisitorID, field.TypeInt)
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeUint, value)
	}
	if value, ok := au.mutation.AddedStatus(); ok {
		_spec.AddField(account.FieldStatus, field.TypeUint, value)
	}
	if value, ok := au.mutation.Subtotal(); ok {
		_spec.SetField(account.FieldSubtotal, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedSubtotal(); ok {
		_spec.AddField(account.FieldSubtotal, field.TypeInt, value)
	}
	if value, ok := au.mutation.Total(); ok {
		_spec.SetField(account.FieldTotal, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedTotal(); ok {
		_spec.AddField(account.FieldTotal, field.TypeInt, value)
	}
	if value, ok := au.mutation.Tax(); ok {
		_spec.SetField(account.FieldTax, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedTax(); ok {
		_spec.AddField(account.FieldTax, field.TypeInt, value)
	}
	if value, ok := au.mutation.TaxRate(); ok {
		_spec.SetField(account.FieldTaxRate, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedTaxRate(); ok {
		_spec.AddField(account.FieldTaxRate, field.TypeInt, value)
	}
	if value, ok := au.mutation.DiscountClass(); ok {
		_spec.SetField(account.FieldDiscountClass, field.TypeUint, value)
	}
	if value, ok := au.mutation.AddedDiscountClass(); ok {
		_spec.AddField(account.FieldDiscountClass, field.TypeUint, value)
	}
	if value, ok := au.mutation.DiscountRate(); ok {
		_spec.SetField(account.FieldDiscountRate, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedDiscountRate(); ok {
		_spec.AddField(account.FieldDiscountRate, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.DiscountPrice(); ok {
		_spec.SetField(account.FieldDiscountPrice, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedDiscountPrice(); ok {
		_spec.AddField(account.FieldDiscountPrice, field.TypeInt, value)
	}
	if value, ok := au.mutation.PaidPrice(); ok {
		_spec.SetField(account.FieldPaidPrice, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedPaidPrice(); ok {
		_spec.AddField(account.FieldPaidPrice, field.TypeInt, value)
	}
	if value, ok := au.mutation.Change(); ok {
		_spec.SetField(account.FieldChange, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedChange(); ok {
		_spec.AddField(account.FieldChange, field.TypeInt, value)
	}
	if au.mutation.ManagedAccountDetailsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedManagedAccountDetailsIDs(); len(nodes) > 0 && !au.mutation.ManagedAccountDetailsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ManagedAccountDetailsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AccountUpdateOne) SetDeletedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDeletedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (auo *AccountUpdateOne) ClearDeletedAt() *AccountUpdateOne {
	auo.mutation.ClearDeletedAt()
	return auo
}

// SetVisitorID sets the "visitor_id" field.
func (auo *AccountUpdateOne) SetVisitorID(i int) *AccountUpdateOne {
	auo.mutation.ResetVisitorID()
	auo.mutation.SetVisitorID(i)
	return auo
}

// SetNillableVisitorID sets the "visitor_id" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableVisitorID(i *int) *AccountUpdateOne {
	if i != nil {
		auo.SetVisitorID(*i)
	}
	return auo
}

// AddVisitorID adds i to the "visitor_id" field.
func (auo *AccountUpdateOne) AddVisitorID(i int) *AccountUpdateOne {
	auo.mutation.AddVisitorID(i)
	return auo
}

// ClearVisitorID clears the value of the "visitor_id" field.
func (auo *AccountUpdateOne) ClearVisitorID() *AccountUpdateOne {
	auo.mutation.ClearVisitorID()
	return auo
}

// SetStatus sets the "status" field.
func (auo *AccountUpdateOne) SetStatus(u uint) *AccountUpdateOne {
	auo.mutation.ResetStatus()
	auo.mutation.SetStatus(u)
	return auo
}

// AddStatus adds u to the "status" field.
func (auo *AccountUpdateOne) AddStatus(u int) *AccountUpdateOne {
	auo.mutation.AddStatus(u)
	return auo
}

// SetSubtotal sets the "subtotal" field.
func (auo *AccountUpdateOne) SetSubtotal(i int) *AccountUpdateOne {
	auo.mutation.ResetSubtotal()
	auo.mutation.SetSubtotal(i)
	return auo
}

// AddSubtotal adds i to the "subtotal" field.
func (auo *AccountUpdateOne) AddSubtotal(i int) *AccountUpdateOne {
	auo.mutation.AddSubtotal(i)
	return auo
}

// SetTotal sets the "total" field.
func (auo *AccountUpdateOne) SetTotal(i int) *AccountUpdateOne {
	auo.mutation.ResetTotal()
	auo.mutation.SetTotal(i)
	return auo
}

// AddTotal adds i to the "total" field.
func (auo *AccountUpdateOne) AddTotal(i int) *AccountUpdateOne {
	auo.mutation.AddTotal(i)
	return auo
}

// SetTax sets the "tax" field.
func (auo *AccountUpdateOne) SetTax(i int) *AccountUpdateOne {
	auo.mutation.ResetTax()
	auo.mutation.SetTax(i)
	return auo
}

// AddTax adds i to the "tax" field.
func (auo *AccountUpdateOne) AddTax(i int) *AccountUpdateOne {
	auo.mutation.AddTax(i)
	return auo
}

// SetTaxRate sets the "tax_rate" field.
func (auo *AccountUpdateOne) SetTaxRate(i int) *AccountUpdateOne {
	auo.mutation.ResetTaxRate()
	auo.mutation.SetTaxRate(i)
	return auo
}

// AddTaxRate adds i to the "tax_rate" field.
func (auo *AccountUpdateOne) AddTaxRate(i int) *AccountUpdateOne {
	auo.mutation.AddTaxRate(i)
	return auo
}

// SetDiscountClass sets the "discount_class" field.
func (auo *AccountUpdateOne) SetDiscountClass(u uint) *AccountUpdateOne {
	auo.mutation.ResetDiscountClass()
	auo.mutation.SetDiscountClass(u)
	return auo
}

// AddDiscountClass adds u to the "discount_class" field.
func (auo *AccountUpdateOne) AddDiscountClass(u int) *AccountUpdateOne {
	auo.mutation.AddDiscountClass(u)
	return auo
}

// SetDiscountRate sets the "discount_rate" field.
func (auo *AccountUpdateOne) SetDiscountRate(d decimal.Decimal) *AccountUpdateOne {
	auo.mutation.ResetDiscountRate()
	auo.mutation.SetDiscountRate(d)
	return auo
}

// AddDiscountRate adds d to the "discount_rate" field.
func (auo *AccountUpdateOne) AddDiscountRate(d decimal.Decimal) *AccountUpdateOne {
	auo.mutation.AddDiscountRate(d)
	return auo
}

// SetDiscountPrice sets the "discount_price" field.
func (auo *AccountUpdateOne) SetDiscountPrice(i int) *AccountUpdateOne {
	auo.mutation.ResetDiscountPrice()
	auo.mutation.SetDiscountPrice(i)
	return auo
}

// AddDiscountPrice adds i to the "discount_price" field.
func (auo *AccountUpdateOne) AddDiscountPrice(i int) *AccountUpdateOne {
	auo.mutation.AddDiscountPrice(i)
	return auo
}

// SetPaidPrice sets the "paid_price" field.
func (auo *AccountUpdateOne) SetPaidPrice(i int) *AccountUpdateOne {
	auo.mutation.ResetPaidPrice()
	auo.mutation.SetPaidPrice(i)
	return auo
}

// AddPaidPrice adds i to the "paid_price" field.
func (auo *AccountUpdateOne) AddPaidPrice(i int) *AccountUpdateOne {
	auo.mutation.AddPaidPrice(i)
	return auo
}

// SetChange sets the "change" field.
func (auo *AccountUpdateOne) SetChange(i int) *AccountUpdateOne {
	auo.mutation.ResetChange()
	auo.mutation.SetChange(i)
	return auo
}

// AddChange adds i to the "change" field.
func (auo *AccountUpdateOne) AddChange(i int) *AccountUpdateOne {
	auo.mutation.AddChange(i)
	return auo
}

// AddManagedAccountDetailIDs adds the "managed_account_details" edge to the AccountDetail entity by IDs.
func (auo *AccountUpdateOne) AddManagedAccountDetailIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.AddManagedAccountDetailIDs(ids...)
	return auo
}

// AddManagedAccountDetails adds the "managed_account_details" edges to the AccountDetail entity.
func (auo *AccountUpdateOne) AddManagedAccountDetails(a ...*AccountDetail) *AccountUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddManagedAccountDetailIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Visitor entity by ID.
func (auo *AccountUpdateOne) SetOwnerID(id int) *AccountUpdateOne {
	auo.mutation.SetOwnerID(id)
	return auo
}

// SetOwner sets the "owner" edge to the Visitor entity.
func (auo *AccountUpdateOne) SetOwner(v *Visitor) *AccountUpdateOne {
	return auo.SetOwnerID(v.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearManagedAccountDetails clears all "managed_account_details" edges to the AccountDetail entity.
func (auo *AccountUpdateOne) ClearManagedAccountDetails() *AccountUpdateOne {
	auo.mutation.ClearManagedAccountDetails()
	return auo
}

// RemoveManagedAccountDetailIDs removes the "managed_account_details" edge to AccountDetail entities by IDs.
func (auo *AccountUpdateOne) RemoveManagedAccountDetailIDs(ids ...int) *AccountUpdateOne {
	auo.mutation.RemoveManagedAccountDetailIDs(ids...)
	return auo
}

// RemoveManagedAccountDetails removes "managed_account_details" edges to AccountDetail entities.
func (auo *AccountUpdateOne) RemoveManagedAccountDetails(a ...*AccountDetail) *AccountUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveManagedAccountDetailIDs(ids...)
}

// ClearOwner clears the "owner" edge to the Visitor entity.
func (auo *AccountUpdateOne) ClearOwner() *AccountUpdateOne {
	auo.mutation.ClearOwner()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	var (
		err  error
		node *Account
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.VisitorID(); ok {
		if err := account.VisitorIDValidator(v); err != nil {
			return &ValidationError{Name: "visitor_id", err: fmt.Errorf(`ent: validator failed for field "Account.visitor_id": %w`, err)}
		}
	}
	if _, ok := auo.mutation.OwnerID(); auo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Account.owner"`)
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: account.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(account.FieldDeletedAt, field.TypeTime, value)
	}
	if auo.mutation.DeletedAtCleared() {
		_spec.ClearField(account.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.VisitorID(); ok {
		_spec.SetField(account.FieldVisitorID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedVisitorID(); ok {
		_spec.AddField(account.FieldVisitorID, field.TypeInt, value)
	}
	if auo.mutation.VisitorIDCleared() {
		_spec.ClearField(account.FieldVisitorID, field.TypeInt)
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeUint, value)
	}
	if value, ok := auo.mutation.AddedStatus(); ok {
		_spec.AddField(account.FieldStatus, field.TypeUint, value)
	}
	if value, ok := auo.mutation.Subtotal(); ok {
		_spec.SetField(account.FieldSubtotal, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedSubtotal(); ok {
		_spec.AddField(account.FieldSubtotal, field.TypeInt, value)
	}
	if value, ok := auo.mutation.Total(); ok {
		_spec.SetField(account.FieldTotal, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedTotal(); ok {
		_spec.AddField(account.FieldTotal, field.TypeInt, value)
	}
	if value, ok := auo.mutation.Tax(); ok {
		_spec.SetField(account.FieldTax, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedTax(); ok {
		_spec.AddField(account.FieldTax, field.TypeInt, value)
	}
	if value, ok := auo.mutation.TaxRate(); ok {
		_spec.SetField(account.FieldTaxRate, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedTaxRate(); ok {
		_spec.AddField(account.FieldTaxRate, field.TypeInt, value)
	}
	if value, ok := auo.mutation.DiscountClass(); ok {
		_spec.SetField(account.FieldDiscountClass, field.TypeUint, value)
	}
	if value, ok := auo.mutation.AddedDiscountClass(); ok {
		_spec.AddField(account.FieldDiscountClass, field.TypeUint, value)
	}
	if value, ok := auo.mutation.DiscountRate(); ok {
		_spec.SetField(account.FieldDiscountRate, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedDiscountRate(); ok {
		_spec.AddField(account.FieldDiscountRate, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.DiscountPrice(); ok {
		_spec.SetField(account.FieldDiscountPrice, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedDiscountPrice(); ok {
		_spec.AddField(account.FieldDiscountPrice, field.TypeInt, value)
	}
	if value, ok := auo.mutation.PaidPrice(); ok {
		_spec.SetField(account.FieldPaidPrice, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedPaidPrice(); ok {
		_spec.AddField(account.FieldPaidPrice, field.TypeInt, value)
	}
	if value, ok := auo.mutation.Change(); ok {
		_spec.SetField(account.FieldChange, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedChange(); ok {
		_spec.AddField(account.FieldChange, field.TypeInt, value)
	}
	if auo.mutation.ManagedAccountDetailsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedManagedAccountDetailsIDs(); len(nodes) > 0 && !auo.mutation.ManagedAccountDetailsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ManagedAccountDetailsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
