// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Blog/ent/blog"
	"Blog/ent/predicate"
	"Blog/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogUpdate is the builder for updating Blog entities.
type BlogUpdate struct {
	config
	hooks    []Hook
	mutation *BlogMutation
}

// Where appends a list predicates to the BlogUpdate builder.
func (bu *BlogUpdate) Where(ps ...predicate.Blog) *BlogUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BlogUpdate) SetTitle(s string) *BlogUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetContent sets the "content" field.
func (bu *BlogUpdate) SetContent(s string) *BlogUpdate {
	bu.mutation.SetContent(s)
	return bu
}

// SetAuthor sets the "Author" field.
func (bu *BlogUpdate) SetAuthor(s string) *BlogUpdate {
	bu.mutation.SetAuthor(s)
	return bu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (bu *BlogUpdate) SetOwnerID(id int) *BlogUpdate {
	bu.mutation.SetOwnerID(id)
	return bu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (bu *BlogUpdate) SetNillableOwnerID(id *int) *BlogUpdate {
	if id != nil {
		bu = bu.SetOwnerID(*id)
	}
	return bu
}

// SetOwner sets the "owner" edge to the User entity.
func (bu *BlogUpdate) SetOwner(u *User) *BlogUpdate {
	return bu.SetOwnerID(u.ID)
}

// Mutation returns the BlogMutation object of the builder.
func (bu *BlogUpdate) Mutation() *BlogMutation {
	return bu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (bu *BlogUpdate) ClearOwner() *BlogUpdate {
	bu.mutation.ClearOwner()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, BlogMutation](ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlogUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlogUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlogUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BlogUpdate) check() error {
	if v, ok := bu.mutation.Title(); ok {
		if err := blog.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Blog.title": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Content(); ok {
		if err := blog.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Blog.content": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Author(); ok {
		if err := blog.AuthorValidator(v); err != nil {
			return &ValidationError{Name: "Author", err: fmt.Errorf(`ent: validator failed for field "Blog.Author": %w`, err)}
		}
	}
	return nil
}

func (bu *BlogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blog.Table,
			Columns: blog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blog.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(blog.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Content(); ok {
		_spec.SetField(blog.FieldContent, field.TypeString, value)
	}
	if value, ok := bu.mutation.Author(); ok {
		_spec.SetField(blog.FieldAuthor, field.TypeString, value)
	}
	if bu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   blog.OwnerTable,
			Columns: []string{blog.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   blog.OwnerTable,
			Columns: []string{blog.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BlogUpdateOne is the builder for updating a single Blog entity.
type BlogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlogMutation
}

// SetTitle sets the "title" field.
func (buo *BlogUpdateOne) SetTitle(s string) *BlogUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetContent sets the "content" field.
func (buo *BlogUpdateOne) SetContent(s string) *BlogUpdateOne {
	buo.mutation.SetContent(s)
	return buo
}

// SetAuthor sets the "Author" field.
func (buo *BlogUpdateOne) SetAuthor(s string) *BlogUpdateOne {
	buo.mutation.SetAuthor(s)
	return buo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (buo *BlogUpdateOne) SetOwnerID(id int) *BlogUpdateOne {
	buo.mutation.SetOwnerID(id)
	return buo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (buo *BlogUpdateOne) SetNillableOwnerID(id *int) *BlogUpdateOne {
	if id != nil {
		buo = buo.SetOwnerID(*id)
	}
	return buo
}

// SetOwner sets the "owner" edge to the User entity.
func (buo *BlogUpdateOne) SetOwner(u *User) *BlogUpdateOne {
	return buo.SetOwnerID(u.ID)
}

// Mutation returns the BlogMutation object of the builder.
func (buo *BlogUpdateOne) Mutation() *BlogMutation {
	return buo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (buo *BlogUpdateOne) ClearOwner() *BlogUpdateOne {
	buo.mutation.ClearOwner()
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlogUpdateOne) Select(field string, fields ...string) *BlogUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blog entity.
func (buo *BlogUpdateOne) Save(ctx context.Context) (*Blog, error) {
	return withHooks[*Blog, BlogMutation](ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlogUpdateOne) SaveX(ctx context.Context) *Blog {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlogUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlogUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BlogUpdateOne) check() error {
	if v, ok := buo.mutation.Title(); ok {
		if err := blog.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Blog.title": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Content(); ok {
		if err := blog.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Blog.content": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Author(); ok {
		if err := blog.AuthorValidator(v); err != nil {
			return &ValidationError{Name: "Author", err: fmt.Errorf(`ent: validator failed for field "Blog.Author": %w`, err)}
		}
	}
	return nil
}

func (buo *BlogUpdateOne) sqlSave(ctx context.Context) (_node *Blog, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blog.Table,
			Columns: blog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blog.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blog.FieldID)
		for _, f := range fields {
			if !blog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(blog.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Content(); ok {
		_spec.SetField(blog.FieldContent, field.TypeString, value)
	}
	if value, ok := buo.mutation.Author(); ok {
		_spec.SetField(blog.FieldAuthor, field.TypeString, value)
	}
	if buo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   blog.OwnerTable,
			Columns: []string{blog.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   blog.OwnerTable,
			Columns: []string{blog.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Blog{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
