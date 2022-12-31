// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"selfregi/ent/images"
	"selfregi/ent/item"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImagesCreate is the builder for creating a Images entity.
type ImagesCreate struct {
	config
	mutation *ImagesMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *ImagesCreate) SetCreatedAt(t time.Time) *ImagesCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *ImagesCreate) SetNillableCreatedAt(t *time.Time) *ImagesCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *ImagesCreate) SetUpdatedAt(t time.Time) *ImagesCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *ImagesCreate) SetNillableUpdatedAt(t *time.Time) *ImagesCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetDeletedAt sets the "deleted_at" field.
func (ic *ImagesCreate) SetDeletedAt(t time.Time) *ImagesCreate {
	ic.mutation.SetDeletedAt(t)
	return ic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ic *ImagesCreate) SetNillableDeletedAt(t *time.Time) *ImagesCreate {
	if t != nil {
		ic.SetDeletedAt(*t)
	}
	return ic
}

// SetName sets the "name" field.
func (ic *ImagesCreate) SetName(s string) *ImagesCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetPath sets the "path" field.
func (ic *ImagesCreate) SetPath(s string) *ImagesCreate {
	ic.mutation.SetPath(s)
	return ic
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (ic *ImagesCreate) AddItemIDs(ids ...int) *ImagesCreate {
	ic.mutation.AddItemIDs(ids...)
	return ic
}

// AddItems adds the "items" edges to the Item entity.
func (ic *ImagesCreate) AddItems(i ...*Item) *ImagesCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ic.AddItemIDs(ids...)
}

// Mutation returns the ImagesMutation object of the builder.
func (ic *ImagesCreate) Mutation() *ImagesMutation {
	return ic.mutation
}

// Save creates the Images in the database.
func (ic *ImagesCreate) Save(ctx context.Context) (*Images, error) {
	var (
		err  error
		node *Images
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImagesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Images)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ImagesMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ImagesCreate) SaveX(ctx context.Context) *Images {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ImagesCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ImagesCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ImagesCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := images.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := images.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ImagesCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Images.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Images.updated_at"`)}
	}
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Images.name"`)}
	}
	if _, ok := ic.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Images.path"`)}
	}
	return nil
}

func (ic *ImagesCreate) sqlSave(ctx context.Context) (*Images, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ic *ImagesCreate) createSpec() (*Images, *sqlgraph.CreateSpec) {
	var (
		_node = &Images{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: images.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: images.FieldID,
			},
		}
	)
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(images.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(images.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.DeletedAt(); ok {
		_spec.SetField(images.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(images.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Path(); ok {
		_spec.SetField(images.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if nodes := ic.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   images.ItemsTable,
			Columns: images.ItemsPrimaryKey,
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

// ImagesCreateBulk is the builder for creating many Images entities in bulk.
type ImagesCreateBulk struct {
	config
	builders []*ImagesCreate
}

// Save creates the Images entities in the database.
func (icb *ImagesCreateBulk) Save(ctx context.Context) ([]*Images, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Images, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImagesMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ImagesCreateBulk) SaveX(ctx context.Context) []*Images {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ImagesCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ImagesCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
