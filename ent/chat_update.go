// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/IshinShibata/chatApp/ent/chat"
	"github.com/IshinShibata/chatApp/ent/predicate"
	"github.com/IshinShibata/chatApp/ent/user"
)

// ChatUpdate is the builder for updating Chat entities.
type ChatUpdate struct {
	config
	hooks    []Hook
	mutation *ChatMutation
}

// Where appends a list predicates to the ChatUpdate builder.
func (cu *ChatUpdate) Where(ps ...predicate.Chat) *ChatUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetTitle sets the "title" field.
func (cu *ChatUpdate) SetTitle(s string) *ChatUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cu *ChatUpdate) SetNillableTitle(s *string) *ChatUpdate {
	if s != nil {
		cu.SetTitle(*s)
	}
	return cu
}

// SetDescription sets the "description" field.
func (cu *ChatUpdate) SetDescription(s string) *ChatUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetImage sets the "image" field.
func (cu *ChatUpdate) SetImage(s string) *ChatUpdate {
	cu.mutation.SetImage(s)
	return cu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (cu *ChatUpdate) SetNillableImage(s *string) *ChatUpdate {
	if s != nil {
		cu.SetImage(*s)
	}
	return cu
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (cu *ChatUpdate) SetAuthorID(id int) *ChatUpdate {
	cu.mutation.SetAuthorID(id)
	return cu
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (cu *ChatUpdate) SetNillableAuthorID(id *int) *ChatUpdate {
	if id != nil {
		cu = cu.SetAuthorID(*id)
	}
	return cu
}

// SetAuthor sets the "author" edge to the User entity.
func (cu *ChatUpdate) SetAuthor(u *User) *ChatUpdate {
	return cu.SetAuthorID(u.ID)
}

// Mutation returns the ChatMutation object of the builder.
func (cu *ChatUpdate) Mutation() *ChatMutation {
	return cu.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (cu *ChatUpdate) ClearAuthor() *ChatUpdate {
	cu.mutation.ClearAuthor()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChatUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChatUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChatUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChatUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ChatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(chat.Table, chat.Columns, sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.SetField(chat.FieldTitle, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(chat.FieldDescription, field.TypeString, value)
	}
	if value, ok := cu.mutation.Image(); ok {
		_spec.SetField(chat.FieldImage, field.TypeString, value)
	}
	if cu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chat.AuthorTable,
			Columns: []string{chat.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chat.AuthorTable,
			Columns: []string{chat.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChatUpdateOne is the builder for updating a single Chat entity.
type ChatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatMutation
}

// SetTitle sets the "title" field.
func (cuo *ChatUpdateOne) SetTitle(s string) *ChatUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableTitle(s *string) *ChatUpdateOne {
	if s != nil {
		cuo.SetTitle(*s)
	}
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ChatUpdateOne) SetDescription(s string) *ChatUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetImage sets the "image" field.
func (cuo *ChatUpdateOne) SetImage(s string) *ChatUpdateOne {
	cuo.mutation.SetImage(s)
	return cuo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableImage(s *string) *ChatUpdateOne {
	if s != nil {
		cuo.SetImage(*s)
	}
	return cuo
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (cuo *ChatUpdateOne) SetAuthorID(id int) *ChatUpdateOne {
	cuo.mutation.SetAuthorID(id)
	return cuo
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableAuthorID(id *int) *ChatUpdateOne {
	if id != nil {
		cuo = cuo.SetAuthorID(*id)
	}
	return cuo
}

// SetAuthor sets the "author" edge to the User entity.
func (cuo *ChatUpdateOne) SetAuthor(u *User) *ChatUpdateOne {
	return cuo.SetAuthorID(u.ID)
}

// Mutation returns the ChatMutation object of the builder.
func (cuo *ChatUpdateOne) Mutation() *ChatMutation {
	return cuo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (cuo *ChatUpdateOne) ClearAuthor() *ChatUpdateOne {
	cuo.mutation.ClearAuthor()
	return cuo
}

// Where appends a list predicates to the ChatUpdate builder.
func (cuo *ChatUpdateOne) Where(ps ...predicate.Chat) *ChatUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChatUpdateOne) Select(field string, fields ...string) *ChatUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chat entity.
func (cuo *ChatUpdateOne) Save(ctx context.Context) (*Chat, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChatUpdateOne) SaveX(ctx context.Context) *Chat {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChatUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChatUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ChatUpdateOne) sqlSave(ctx context.Context) (_node *Chat, err error) {
	_spec := sqlgraph.NewUpdateSpec(chat.Table, chat.Columns, sqlgraph.NewFieldSpec(chat.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chat.FieldID)
		for _, f := range fields {
			if !chat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.SetField(chat.FieldTitle, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(chat.FieldDescription, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Image(); ok {
		_spec.SetField(chat.FieldImage, field.TypeString, value)
	}
	if cuo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chat.AuthorTable,
			Columns: []string{chat.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chat.AuthorTable,
			Columns: []string{chat.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Chat{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}