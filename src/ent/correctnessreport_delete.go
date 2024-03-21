// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/KenshiTech/unchained/ent/correctnessreport"
	"github.com/KenshiTech/unchained/ent/predicate"
)

// CorrectnessReportDelete is the builder for deleting a CorrectnessReport entity.
type CorrectnessReportDelete struct {
	config
	hooks    []Hook
	mutation *CorrectnessReportMutation
}

// Where appends a list predicates to the CorrectnessReportDelete builder.
func (crd *CorrectnessReportDelete) Where(ps ...predicate.CorrectnessReport) *CorrectnessReportDelete {
	crd.mutation.Where(ps...)
	return crd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crd *CorrectnessReportDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, crd.sqlExec, crd.mutation, crd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crd *CorrectnessReportDelete) ExecX(ctx context.Context) int {
	n, err := crd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crd *CorrectnessReportDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(correctnessreport.Table, sqlgraph.NewFieldSpec(correctnessreport.FieldID, field.TypeInt))
	if ps := crd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crd.mutation.done = true
	return affected, err
}

// CorrectnessReportDeleteOne is the builder for deleting a single CorrectnessReport entity.
type CorrectnessReportDeleteOne struct {
	crd *CorrectnessReportDelete
}

// Where appends a list predicates to the CorrectnessReportDelete builder.
func (crdo *CorrectnessReportDeleteOne) Where(ps ...predicate.CorrectnessReport) *CorrectnessReportDeleteOne {
	crdo.crd.mutation.Where(ps...)
	return crdo
}

// Exec executes the deletion query.
func (crdo *CorrectnessReportDeleteOne) Exec(ctx context.Context) error {
	n, err := crdo.crd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{correctnessreport.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crdo *CorrectnessReportDeleteOne) ExecX(ctx context.Context) {
	if err := crdo.Exec(ctx); err != nil {
		panic(err)
	}
}
