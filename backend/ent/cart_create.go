// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"selfregi/ent/cart"
	"selfregi/ent/cartdetail"
	"selfregi/ent/visitor"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CartCreate is the builder for creating a Cart entity.
type CartCreate struct {
	config
	mutation *CartMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CartCreate) SetCreatedAt(t time.Time) *CartCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CartCreate) SetNillableCreatedAt(t *time.Time) *CartCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CartCreate) SetUpdatedAt(t time.Time) *CartCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CartCreate) SetNillableUpdatedAt(t *time.Time) *CartCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CartCreate) SetDeletedAt(t time.Time) *CartCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CartCreate) SetNillableDeletedAt(t *time.Time) *CartCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetVisitorID sets the "visitor_id" field.
func (cc *CartCreate) SetVisitorID(i int) *CartCreate {
	cc.mutation.SetVisitorID(i)
	return cc
}

// SetNillableVisitorID sets the "visitor_id" field if the given value is not nil.
func (cc *CartCreate) SetNillableVisitorID(i *int) *CartCreate {
	if i != nil {
		cc.SetVisitorID(*i)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CartCreate) SetStatus(u uint) *CartCreate {
	cc.mutation.SetStatus(u)
	return cc
}

// AddCartDetailIDs adds the "cart_details" edge to the CartDetail entity by IDs.
func (cc *CartCreate) AddCartDetailIDs(ids ...int) *CartCreate {
	cc.mutation.AddCartDetailIDs(ids...)
	return cc
}

// AddCartDetails adds the "cart_details" edges to the CartDetail entity.
func (cc *CartCreate) AddCartDetails(c ...*CartDetail) *CartCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCartDetailIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Visitor entity by ID.
func (cc *CartCreate) SetOwnerID(id int) *CartCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetOwner sets the "owner" edge to the Visitor entity.
func (cc *CartCreate) SetOwner(v *Visitor) *CartCreate {
	return cc.SetOwnerID(v.ID)
}

// Mutation returns the CartMutation object of the builder.
func (cc *CartCreate) Mutation() *CartMutation {
	return cc.mutation
}

// Save creates the Cart in the database.
func (cc *CartCreate) Save(ctx context.Context) (*Cart, error) {
	var (
		err  error
		node *Cart
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Cart)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CartMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CartCreate) SaveX(ctx context.Context) *Cart {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CartCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CartCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CartCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := cart.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := cart.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CartCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Cart.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Cart.updated_at"`)}
	}
	if v, ok := cc.mutation.VisitorID(); ok {
		if err := cart.VisitorIDValidator(v); err != nil {
			return &ValidationError{Name: "visitor_id", err: fmt.Errorf(`ent: validator failed for field "Cart.visitor_id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Cart.status"`)}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Cart.owner"`)}
	}
	return nil
}

func (cc *CartCreate) sqlSave(ctx context.Context) (*Cart, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CartCreate) createSpec() (*Cart, *sqlgraph.CreateSpec) {
	var (
		_node = &Cart{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cart.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cart.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(cart.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(cart.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(cart.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := cc.mutation.VisitorID(); ok {
		_spec.SetField(cart.FieldVisitorID, field.TypeInt, value)
		_node.VisitorID = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(cart.FieldStatus, field.TypeUint, value)
		_node.Status = &value
	}
	if nodes := cc.mutation.CartDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.CartDetailsTable,
			Columns: []string{cart.CartDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cartdetail.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.OwnerTable,
			Columns: []string{cart.OwnerColumn},
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
		_node.visitor_carts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CartCreateBulk is the builder for creating many Cart entities in bulk.
type CartCreateBulk struct {
	config
	builders []*CartCreate
}

// Save creates the Cart entities in the database.
func (ccb *CartCreateBulk) Save(ctx context.Context) ([]*Cart, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cart, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CartMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CartCreateBulk) SaveX(ctx context.Context) []*Cart {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CartCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CartCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
