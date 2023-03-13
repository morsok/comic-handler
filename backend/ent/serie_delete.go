// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/morsok/comic-handler/ent/predicate"
	"github.com/morsok/comic-handler/ent/serie"
)

// SerieDelete is the builder for deleting a Serie entity.
type SerieDelete struct {
	config
	hooks    []Hook
	mutation *SerieMutation
}

// Where appends a list predicates to the SerieDelete builder.
func (sd *SerieDelete) Where(ps ...predicate.Serie) *SerieDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SerieDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, SerieMutation](ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SerieDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SerieDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: serie.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: serie.FieldID,
			},
		},
	}
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SerieDeleteOne is the builder for deleting a single Serie entity.
type SerieDeleteOne struct {
	sd *SerieDelete
}

// Where appends a list predicates to the SerieDelete builder.
func (sdo *SerieDeleteOne) Where(ps ...predicate.Serie) *SerieDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SerieDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{serie.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SerieDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
