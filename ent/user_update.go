// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aodebiao/ent/predicate"
	"aodebiao/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUserName sets the "user_name" field.
func (uu *UserUpdate) SetUserName(s string) *UserUpdate {
	uu.mutation.SetUserName(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetDelFlag sets the "del_flag" field.
func (uu *UserUpdate) SetDelFlag(i int8) *UserUpdate {
	uu.mutation.ResetDelFlag()
	uu.mutation.SetDelFlag(i)
	return uu
}

// AddDelFlag adds i to the "del_flag" field.
func (uu *UserUpdate) AddDelFlag(i int8) *UserUpdate {
	uu.mutation.AddDelFlag(i)
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(i int8) *UserUpdate {
	uu.mutation.ResetStatus()
	uu.mutation.SetStatus(i)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(i *int8) *UserUpdate {
	if i != nil {
		uu.SetStatus(*i)
	}
	return uu
}

// AddStatus adds i to the "status" field.
func (uu *UserUpdate) AddStatus(i int8) *UserUpdate {
	uu.mutation.AddStatus(i)
	return uu
}

// ClearStatus clears the value of the "status" field.
func (uu *UserUpdate) ClearStatus() *UserUpdate {
	uu.mutation.ClearStatus()
	return uu
}

// SetAvatarURL sets the "avatar_url" field.
func (uu *UserUpdate) SetAvatarURL(s string) *UserUpdate {
	uu.mutation.SetAvatarURL(s)
	return uu
}

// SetLoginName sets the "login_name" field.
func (uu *UserUpdate) SetLoginName(s string) *UserUpdate {
	uu.mutation.SetLoginName(s)
	return uu
}

// SetNillableLoginName sets the "login_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLoginName(s *string) *UserUpdate {
	if s != nil {
		uu.SetLoginName(*s)
	}
	return uu
}

// ClearLoginName clears the value of the "login_name" field.
func (uu *UserUpdate) ClearLoginName() *UserUpdate {
	uu.mutation.ClearLoginName()
	return uu
}

// SetCreateAt sets the "create_at" field.
func (uu *UserUpdate) SetCreateAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreateAt(t)
	return uu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreateAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreateAt(*t)
	}
	return uu
}

// ClearCreateAt clears the value of the "create_at" field.
func (uu *UserUpdate) ClearCreateAt() *UserUpdate {
	uu.mutation.ClearCreateAt()
	return uu
}

// SetUpdateAt sets the "update_at" field.
func (uu *UserUpdate) SetUpdateAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdateAt(t)
	return uu
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdateAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdateAt(*t)
	}
	return uu
}

// ClearUpdateAt clears the value of the "update_at" field.
func (uu *UserUpdate) ClearUpdateAt() *UserUpdate {
	uu.mutation.ClearUpdateAt()
	return uu
}

// SetDeleteAt sets the "delete_at" field.
func (uu *UserUpdate) SetDeleteAt(t time.Time) *UserUpdate {
	uu.mutation.SetDeleteAt(t)
	return uu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeleteAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeleteAt(*t)
	}
	return uu
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (uu *UserUpdate) ClearDeleteAt() *UserUpdate {
	uu.mutation.ClearDeleteAt()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.DelFlag(); ok {
		_spec.SetField(user.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := uu.mutation.AddedDelFlag(); ok {
		_spec.AddField(user.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := uu.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeInt8, value)
	}
	if uu.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeInt8)
	}
	if value, ok := uu.mutation.AvatarURL(); ok {
		_spec.SetField(user.FieldAvatarURL, field.TypeString, value)
	}
	if value, ok := uu.mutation.LoginName(); ok {
		_spec.SetField(user.FieldLoginName, field.TypeString, value)
	}
	if uu.mutation.LoginNameCleared() {
		_spec.ClearField(user.FieldLoginName, field.TypeString)
	}
	if value, ok := uu.mutation.CreateAt(); ok {
		_spec.SetField(user.FieldCreateAt, field.TypeTime, value)
	}
	if uu.mutation.CreateAtCleared() {
		_spec.ClearField(user.FieldCreateAt, field.TypeTime)
	}
	if value, ok := uu.mutation.UpdateAt(); ok {
		_spec.SetField(user.FieldUpdateAt, field.TypeTime, value)
	}
	if uu.mutation.UpdateAtCleared() {
		_spec.ClearField(user.FieldUpdateAt, field.TypeTime)
	}
	if value, ok := uu.mutation.DeleteAt(); ok {
		_spec.SetField(user.FieldDeleteAt, field.TypeTime, value)
	}
	if uu.mutation.DeleteAtCleared() {
		_spec.ClearField(user.FieldDeleteAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUserName sets the "user_name" field.
func (uuo *UserUpdateOne) SetUserName(s string) *UserUpdateOne {
	uuo.mutation.SetUserName(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetDelFlag sets the "del_flag" field.
func (uuo *UserUpdateOne) SetDelFlag(i int8) *UserUpdateOne {
	uuo.mutation.ResetDelFlag()
	uuo.mutation.SetDelFlag(i)
	return uuo
}

// AddDelFlag adds i to the "del_flag" field.
func (uuo *UserUpdateOne) AddDelFlag(i int8) *UserUpdateOne {
	uuo.mutation.AddDelFlag(i)
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(i int8) *UserUpdateOne {
	uuo.mutation.ResetStatus()
	uuo.mutation.SetStatus(i)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(i *int8) *UserUpdateOne {
	if i != nil {
		uuo.SetStatus(*i)
	}
	return uuo
}

// AddStatus adds i to the "status" field.
func (uuo *UserUpdateOne) AddStatus(i int8) *UserUpdateOne {
	uuo.mutation.AddStatus(i)
	return uuo
}

// ClearStatus clears the value of the "status" field.
func (uuo *UserUpdateOne) ClearStatus() *UserUpdateOne {
	uuo.mutation.ClearStatus()
	return uuo
}

// SetAvatarURL sets the "avatar_url" field.
func (uuo *UserUpdateOne) SetAvatarURL(s string) *UserUpdateOne {
	uuo.mutation.SetAvatarURL(s)
	return uuo
}

// SetLoginName sets the "login_name" field.
func (uuo *UserUpdateOne) SetLoginName(s string) *UserUpdateOne {
	uuo.mutation.SetLoginName(s)
	return uuo
}

// SetNillableLoginName sets the "login_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLoginName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLoginName(*s)
	}
	return uuo
}

// ClearLoginName clears the value of the "login_name" field.
func (uuo *UserUpdateOne) ClearLoginName() *UserUpdateOne {
	uuo.mutation.ClearLoginName()
	return uuo
}

// SetCreateAt sets the "create_at" field.
func (uuo *UserUpdateOne) SetCreateAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreateAt(t)
	return uuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreateAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreateAt(*t)
	}
	return uuo
}

// ClearCreateAt clears the value of the "create_at" field.
func (uuo *UserUpdateOne) ClearCreateAt() *UserUpdateOne {
	uuo.mutation.ClearCreateAt()
	return uuo
}

// SetUpdateAt sets the "update_at" field.
func (uuo *UserUpdateOne) SetUpdateAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdateAt(t)
	return uuo
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdateAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdateAt(*t)
	}
	return uuo
}

// ClearUpdateAt clears the value of the "update_at" field.
func (uuo *UserUpdateOne) ClearUpdateAt() *UserUpdateOne {
	uuo.mutation.ClearUpdateAt()
	return uuo
}

// SetDeleteAt sets the "delete_at" field.
func (uuo *UserUpdateOne) SetDeleteAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeleteAt(t)
	return uuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeleteAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeleteAt(*t)
	}
	return uuo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (uuo *UserUpdateOne) ClearDeleteAt() *UserUpdateOne {
	uuo.mutation.ClearDeleteAt()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UserName(); ok {
		_spec.SetField(user.FieldUserName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.DelFlag(); ok {
		_spec.SetField(user.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := uuo.mutation.AddedDelFlag(); ok {
		_spec.AddField(user.FieldDelFlag, field.TypeInt8, value)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := uuo.mutation.AddedStatus(); ok {
		_spec.AddField(user.FieldStatus, field.TypeInt8, value)
	}
	if uuo.mutation.StatusCleared() {
		_spec.ClearField(user.FieldStatus, field.TypeInt8)
	}
	if value, ok := uuo.mutation.AvatarURL(); ok {
		_spec.SetField(user.FieldAvatarURL, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LoginName(); ok {
		_spec.SetField(user.FieldLoginName, field.TypeString, value)
	}
	if uuo.mutation.LoginNameCleared() {
		_spec.ClearField(user.FieldLoginName, field.TypeString)
	}
	if value, ok := uuo.mutation.CreateAt(); ok {
		_spec.SetField(user.FieldCreateAt, field.TypeTime, value)
	}
	if uuo.mutation.CreateAtCleared() {
		_spec.ClearField(user.FieldCreateAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.UpdateAt(); ok {
		_spec.SetField(user.FieldUpdateAt, field.TypeTime, value)
	}
	if uuo.mutation.UpdateAtCleared() {
		_spec.ClearField(user.FieldUpdateAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.DeleteAt(); ok {
		_spec.SetField(user.FieldDeleteAt, field.TypeTime, value)
	}
	if uuo.mutation.DeleteAtCleared() {
		_spec.ClearField(user.FieldDeleteAt, field.TypeTime)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}