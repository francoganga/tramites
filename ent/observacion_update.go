// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/francoganga/go_reno2/ent/observacion"
	"github.com/francoganga/go_reno2/ent/predicate"
)

// ObservacionUpdate is the builder for updating Observacion entities.
type ObservacionUpdate struct {
	config
	hooks    []Hook
	mutation *ObservacionMutation
}

// Where appends a list predicates to the ObservacionUpdate builder.
func (ou *ObservacionUpdate) Where(ps ...predicate.Observacion) *ObservacionUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// Mutation returns the ObservacionMutation object of the builder.
func (ou *ObservacionUpdate) Mutation() *ObservacionMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *ObservacionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObservacionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *ObservacionUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *ObservacionUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *ObservacionUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *ObservacionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   observacion.Table,
			Columns: observacion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: observacion.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{observacion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ObservacionUpdateOne is the builder for updating a single Observacion entity.
type ObservacionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ObservacionMutation
}

// Mutation returns the ObservacionMutation object of the builder.
func (ouo *ObservacionUpdateOne) Mutation() *ObservacionMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *ObservacionUpdateOne) Select(field string, fields ...string) *ObservacionUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Observacion entity.
func (ouo *ObservacionUpdateOne) Save(ctx context.Context) (*Observacion, error) {
	var (
		err  error
		node *Observacion
	)
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ObservacionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ouo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Observacion)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ObservacionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *ObservacionUpdateOne) SaveX(ctx context.Context) *Observacion {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *ObservacionUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *ObservacionUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *ObservacionUpdateOne) sqlSave(ctx context.Context) (_node *Observacion, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   observacion.Table,
			Columns: observacion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: observacion.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Observacion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, observacion.FieldID)
		for _, f := range fields {
			if !observacion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != observacion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &Observacion{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{observacion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
