// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"selfregi/ent/categories"
	"selfregi/ent/item"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoriesCreate is the builder for creating a Categories entity.
type CategoriesCreate struct {
	config
	mutation *CategoriesMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CategoriesCreate) SetCreatedAt(t time.Time) *CategoriesCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CategoriesCreate) SetNillableCreatedAt(t *time.Time) *CategoriesCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CategoriesCreate) SetUpdatedAt(t time.Time) *CategoriesCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CategoriesCreate) SetNillableUpdatedAt(t *time.Time) *CategoriesCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CategoriesCreate) SetDeletedAt(t time.Time) *CategoriesCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CategoriesCreate) SetNillableDeletedAt(t *time.Time) *CategoriesCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CategoriesCreate) SetName(s string) *CategoriesCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetIsDisplay sets the "is_display" field.
func (cc *CategoriesCreate) SetIsDisplay(i int) *CategoriesCreate {
	cc.mutation.SetIsDisplay(i)
	return cc
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (cc *CategoriesCreate) AddItemIDs(ids ...int) *CategoriesCreate {
	cc.mutation.AddItemIDs(ids...)
	return cc
}

// AddItems adds the "items" edges to the Item entity.
func (cc *CategoriesCreate) AddItems(i ...*Item) *CategoriesCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cc.AddItemIDs(ids...)
}

// Mutation returns the CategoriesMutation object of the builder.
func (cc *CategoriesCreate) Mutation() *CategoriesMutation {
	return cc.mutation
}

// Save creates the Categories in the database.
func (cc *CategoriesCreate) Save(ctx context.Context) (*Categories, error) {
	var (
		err  error
		node *Categories
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoriesMutation)
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
		nv, ok := v.(*Categories)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CategoriesMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoriesCreate) SaveX(ctx context.Context) *Categories {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoriesCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoriesCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CategoriesCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := categories.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := categories.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoriesCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Categories.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Categories.updated_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Categories.name"`)}
	}
	if _, ok := cc.mutation.IsDisplay(); !ok {
		return &ValidationError{Name: "is_display", err: errors.New(`ent: missing required field "Categories.is_display"`)}
	}
	return nil
}

func (cc *CategoriesCreate) sqlSave(ctx context.Context) (*Categories, error) {
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

func (cc *CategoriesCreate) createSpec() (*Categories, *sqlgraph.CreateSpec) {
	var (
		_node = &Categories{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: categories.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: categories.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(categories.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(categories.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(categories.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(categories.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.IsDisplay(); ok {
		_spec.SetField(categories.FieldIsDisplay, field.TypeInt, value)
		_node.IsDisplay = value
	}
	if nodes := cc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   categories.ItemsTable,
			Columns: []string{categories.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CategoriesCreateBulk is the builder for creating many Categories entities in bulk.
type CategoriesCreateBulk struct {
	config
	builders []*CategoriesCreate
}

// Save creates the Categories entities in the database.
func (ccb *CategoriesCreateBulk) Save(ctx context.Context) ([]*Categories, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Categories, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoriesMutation)
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
func (ccb *CategoriesCreateBulk) SaveX(ctx context.Context) []*Categories {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoriesCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoriesCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}