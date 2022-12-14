// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"selfregi/ent/journals"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JournalsCreate is the builder for creating a Journals entity.
type JournalsCreate struct {
	config
	mutation *JournalsMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (jc *JournalsCreate) SetCreatedAt(t time.Time) *JournalsCreate {
	jc.mutation.SetCreatedAt(t)
	return jc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jc *JournalsCreate) SetNillableCreatedAt(t *time.Time) *JournalsCreate {
	if t != nil {
		jc.SetCreatedAt(*t)
	}
	return jc
}

// SetUpdatedAt sets the "updated_at" field.
func (jc *JournalsCreate) SetUpdatedAt(t time.Time) *JournalsCreate {
	jc.mutation.SetUpdatedAt(t)
	return jc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (jc *JournalsCreate) SetNillableUpdatedAt(t *time.Time) *JournalsCreate {
	if t != nil {
		jc.SetUpdatedAt(*t)
	}
	return jc
}

// SetDeletedAt sets the "deleted_at" field.
func (jc *JournalsCreate) SetDeletedAt(t time.Time) *JournalsCreate {
	jc.mutation.SetDeletedAt(t)
	return jc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (jc *JournalsCreate) SetNillableDeletedAt(t *time.Time) *JournalsCreate {
	if t != nil {
		jc.SetDeletedAt(*t)
	}
	return jc
}

// Mutation returns the JournalsMutation object of the builder.
func (jc *JournalsCreate) Mutation() *JournalsMutation {
	return jc.mutation
}

// Save creates the Journals in the database.
func (jc *JournalsCreate) Save(ctx context.Context) (*Journals, error) {
	var (
		err  error
		node *Journals
	)
	jc.defaults()
	if len(jc.hooks) == 0 {
		if err = jc.check(); err != nil {
			return nil, err
		}
		node, err = jc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JournalsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = jc.check(); err != nil {
				return nil, err
			}
			jc.mutation = mutation
			if node, err = jc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(jc.hooks) - 1; i >= 0; i-- {
			if jc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = jc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, jc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Journals)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from JournalsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (jc *JournalsCreate) SaveX(ctx context.Context) *Journals {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jc *JournalsCreate) Exec(ctx context.Context) error {
	_, err := jc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jc *JournalsCreate) ExecX(ctx context.Context) {
	if err := jc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jc *JournalsCreate) defaults() {
	if _, ok := jc.mutation.CreatedAt(); !ok {
		v := journals.DefaultCreatedAt()
		jc.mutation.SetCreatedAt(v)
	}
	if _, ok := jc.mutation.UpdatedAt(); !ok {
		v := journals.DefaultUpdatedAt()
		jc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jc *JournalsCreate) check() error {
	if _, ok := jc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Journals.created_at"`)}
	}
	if _, ok := jc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Journals.updated_at"`)}
	}
	return nil
}

func (jc *JournalsCreate) sqlSave(ctx context.Context) (*Journals, error) {
	_node, _spec := jc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (jc *JournalsCreate) createSpec() (*Journals, *sqlgraph.CreateSpec) {
	var (
		_node = &Journals{config: jc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: journals.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: journals.FieldID,
			},
		}
	)
	if value, ok := jc.mutation.CreatedAt(); ok {
		_spec.SetField(journals.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := jc.mutation.UpdatedAt(); ok {
		_spec.SetField(journals.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := jc.mutation.DeletedAt(); ok {
		_spec.SetField(journals.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// JournalsCreateBulk is the builder for creating many Journals entities in bulk.
type JournalsCreateBulk struct {
	config
	builders []*JournalsCreate
}

// Save creates the Journals entities in the database.
func (jcb *JournalsCreateBulk) Save(ctx context.Context) ([]*Journals, error) {
	specs := make([]*sqlgraph.CreateSpec, len(jcb.builders))
	nodes := make([]*Journals, len(jcb.builders))
	mutators := make([]Mutator, len(jcb.builders))
	for i := range jcb.builders {
		func(i int, root context.Context) {
			builder := jcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JournalsMutation)
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
					_, err = mutators[i+1].Mutate(root, jcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, jcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jcb *JournalsCreateBulk) SaveX(ctx context.Context) []*Journals {
	v, err := jcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jcb *JournalsCreateBulk) Exec(ctx context.Context) error {
	_, err := jcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcb *JournalsCreateBulk) ExecX(ctx context.Context) {
	if err := jcb.Exec(ctx); err != nil {
		panic(err)
	}
}
