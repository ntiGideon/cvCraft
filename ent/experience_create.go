// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ntiGideon/cvCraft/ent/experience"
	"github.com/ntiGideon/cvCraft/ent/resume"
)

// ExperienceCreate is the builder for creating a Experience entity.
type ExperienceCreate struct {
	config
	mutation *ExperienceMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (ec *ExperienceCreate) SetTitle(s string) *ExperienceCreate {
	ec.mutation.SetTitle(s)
	return ec
}

// SetCompany sets the "company" field.
func (ec *ExperienceCreate) SetCompany(s string) *ExperienceCreate {
	ec.mutation.SetCompany(s)
	return ec
}

// SetLocation sets the "location" field.
func (ec *ExperienceCreate) SetLocation(s string) *ExperienceCreate {
	ec.mutation.SetLocation(s)
	return ec
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (ec *ExperienceCreate) SetNillableLocation(s *string) *ExperienceCreate {
	if s != nil {
		ec.SetLocation(*s)
	}
	return ec
}

// SetStartDate sets the "start_date" field.
func (ec *ExperienceCreate) SetStartDate(t time.Time) *ExperienceCreate {
	ec.mutation.SetStartDate(t)
	return ec
}

// SetEndDate sets the "end_date" field.
func (ec *ExperienceCreate) SetEndDate(t time.Time) *ExperienceCreate {
	ec.mutation.SetEndDate(t)
	return ec
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (ec *ExperienceCreate) SetNillableEndDate(t *time.Time) *ExperienceCreate {
	if t != nil {
		ec.SetEndDate(*t)
	}
	return ec
}

// SetDescription sets the "description" field.
func (ec *ExperienceCreate) SetDescription(s string) *ExperienceCreate {
	ec.mutation.SetDescription(s)
	return ec
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ec *ExperienceCreate) SetNillableDescription(s *string) *ExperienceCreate {
	if s != nil {
		ec.SetDescription(*s)
	}
	return ec
}

// SetSkills sets the "skills" field.
func (ec *ExperienceCreate) SetSkills(s string) *ExperienceCreate {
	ec.mutation.SetSkills(s)
	return ec
}

// SetNillableSkills sets the "skills" field if the given value is not nil.
func (ec *ExperienceCreate) SetNillableSkills(s *string) *ExperienceCreate {
	if s != nil {
		ec.SetSkills(*s)
	}
	return ec
}

// SetCurrent sets the "current" field.
func (ec *ExperienceCreate) SetCurrent(b bool) *ExperienceCreate {
	ec.mutation.SetCurrent(b)
	return ec
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (ec *ExperienceCreate) SetNillableCurrent(b *bool) *ExperienceCreate {
	if b != nil {
		ec.SetCurrent(*b)
	}
	return ec
}

// AddResumeIDs adds the "resume" edge to the Resume entity by IDs.
func (ec *ExperienceCreate) AddResumeIDs(ids ...int) *ExperienceCreate {
	ec.mutation.AddResumeIDs(ids...)
	return ec
}

// AddResume adds the "resume" edges to the Resume entity.
func (ec *ExperienceCreate) AddResume(r ...*Resume) *ExperienceCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ec.AddResumeIDs(ids...)
}

// Mutation returns the ExperienceMutation object of the builder.
func (ec *ExperienceCreate) Mutation() *ExperienceMutation {
	return ec.mutation
}

// Save creates the Experience in the database.
func (ec *ExperienceCreate) Save(ctx context.Context) (*Experience, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *ExperienceCreate) SaveX(ctx context.Context) *Experience {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *ExperienceCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *ExperienceCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *ExperienceCreate) defaults() {
	if _, ok := ec.mutation.Current(); !ok {
		v := experience.DefaultCurrent
		ec.mutation.SetCurrent(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *ExperienceCreate) check() error {
	if _, ok := ec.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Experience.title"`)}
	}
	if v, ok := ec.mutation.Title(); ok {
		if err := experience.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Experience.title": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Company(); !ok {
		return &ValidationError{Name: "company", err: errors.New(`ent: missing required field "Experience.company"`)}
	}
	if v, ok := ec.mutation.Company(); ok {
		if err := experience.CompanyValidator(v); err != nil {
			return &ValidationError{Name: "company", err: fmt.Errorf(`ent: validator failed for field "Experience.company": %w`, err)}
		}
	}
	if _, ok := ec.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`ent: missing required field "Experience.start_date"`)}
	}
	if _, ok := ec.mutation.Current(); !ok {
		return &ValidationError{Name: "current", err: errors.New(`ent: missing required field "Experience.current"`)}
	}
	if len(ec.mutation.ResumeIDs()) == 0 {
		return &ValidationError{Name: "resume", err: errors.New(`ent: missing required edge "Experience.resume"`)}
	}
	return nil
}

func (ec *ExperienceCreate) sqlSave(ctx context.Context) (*Experience, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *ExperienceCreate) createSpec() (*Experience, *sqlgraph.CreateSpec) {
	var (
		_node = &Experience{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(experience.Table, sqlgraph.NewFieldSpec(experience.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.Title(); ok {
		_spec.SetField(experience.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ec.mutation.Company(); ok {
		_spec.SetField(experience.FieldCompany, field.TypeString, value)
		_node.Company = value
	}
	if value, ok := ec.mutation.Location(); ok {
		_spec.SetField(experience.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := ec.mutation.StartDate(); ok {
		_spec.SetField(experience.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := ec.mutation.EndDate(); ok {
		_spec.SetField(experience.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if value, ok := ec.mutation.Description(); ok {
		_spec.SetField(experience.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ec.mutation.Skills(); ok {
		_spec.SetField(experience.FieldSkills, field.TypeString, value)
		_node.Skills = value
	}
	if value, ok := ec.mutation.Current(); ok {
		_spec.SetField(experience.FieldCurrent, field.TypeBool, value)
		_node.Current = value
	}
	if nodes := ec.mutation.ResumeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   experience.ResumeTable,
			Columns: experience.ResumePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resume.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ExperienceCreateBulk is the builder for creating many Experience entities in bulk.
type ExperienceCreateBulk struct {
	config
	err      error
	builders []*ExperienceCreate
}

// Save creates the Experience entities in the database.
func (ecb *ExperienceCreateBulk) Save(ctx context.Context) ([]*Experience, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Experience, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExperienceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *ExperienceCreateBulk) SaveX(ctx context.Context) []*Experience {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *ExperienceCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *ExperienceCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
