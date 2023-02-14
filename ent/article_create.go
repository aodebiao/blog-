// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aodebiao/ent/article"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArticleCreate is the builder for creating a Article entity.
type ArticleCreate struct {
	config
	mutation *ArticleMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *ArticleCreate) SetCreatedAt(t time.Time) *ArticleCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableCreatedAt(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *ArticleCreate) SetUpdatedAt(t time.Time) *ArticleCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableUpdatedAt(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *ArticleCreate) SetDeletedAt(t time.Time) *ArticleCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableDeletedAt(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetDelFlag sets the "del_flag" field.
func (ac *ArticleCreate) SetDelFlag(i int) *ArticleCreate {
	ac.mutation.SetDelFlag(i)
	return ac
}

// SetNillableDelFlag sets the "del_flag" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableDelFlag(i *int) *ArticleCreate {
	if i != nil {
		ac.SetDelFlag(*i)
	}
	return ac
}

// SetCategoryID sets the "category_id" field.
func (ac *ArticleCreate) SetCategoryID(s string) *ArticleCreate {
	ac.mutation.SetCategoryID(s)
	return ac
}

// SetContent sets the "content" field.
func (ac *ArticleCreate) SetContent(s string) *ArticleCreate {
	ac.mutation.SetContent(s)
	return ac
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableContent(s *string) *ArticleCreate {
	if s != nil {
		ac.SetContent(*s)
	}
	return ac
}

// SetStatus sets the "status" field.
func (ac *ArticleCreate) SetStatus(i int8) *ArticleCreate {
	ac.mutation.SetStatus(i)
	return ac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableStatus(i *int8) *ArticleCreate {
	if i != nil {
		ac.SetStatus(*i)
	}
	return ac
}

// SetCoverURL sets the "cover_url" field.
func (ac *ArticleCreate) SetCoverURL(s string) *ArticleCreate {
	ac.mutation.SetCoverURL(s)
	return ac
}

// SetNillableCoverURL sets the "cover_url" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableCoverURL(s *string) *ArticleCreate {
	if s != nil {
		ac.SetCoverURL(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ArticleCreate) SetID(s string) *ArticleCreate {
	ac.mutation.SetID(s)
	return ac
}

// Mutation returns the ArticleMutation object of the builder.
func (ac *ArticleCreate) Mutation() *ArticleMutation {
	return ac.mutation
}

// Save creates the Article in the database.
func (ac *ArticleCreate) Save(ctx context.Context) (*Article, error) {
	ac.defaults()
	return withHooks[*Article, ArticleMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArticleCreate) SaveX(ctx context.Context) *Article {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArticleCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArticleCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArticleCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := article.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := article.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.DelFlag(); !ok {
		v := article.DefaultDelFlag
		ac.mutation.SetDelFlag(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArticleCreate) check() error {
	if _, ok := ac.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`ent: missing required field "Article.category_id"`)}
	}
	return nil
}

func (ac *ArticleCreate) sqlSave(ctx context.Context) (*Article, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Article.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArticleCreate) createSpec() (*Article, *sqlgraph.CreateSpec) {
	var (
		_node = &Article{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(article.Table, sqlgraph.NewFieldSpec(article.FieldID, field.TypeString))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(article.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(article.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(article.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ac.mutation.DelFlag(); ok {
		_spec.SetField(article.FieldDelFlag, field.TypeInt, value)
		_node.DelFlag = value
	}
	if value, ok := ac.mutation.CategoryID(); ok {
		_spec.SetField(article.FieldCategoryID, field.TypeString, value)
		_node.CategoryID = value
	}
	if value, ok := ac.mutation.Content(); ok {
		_spec.SetField(article.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(article.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := ac.mutation.CoverURL(); ok {
		_spec.SetField(article.FieldCoverURL, field.TypeString, value)
		_node.CoverURL = value
	}
	return _node, _spec
}

// ArticleCreateBulk is the builder for creating many Article entities in bulk.
type ArticleCreateBulk struct {
	config
	builders []*ArticleCreate
}

// Save creates the Article entities in the database.
func (acb *ArticleCreateBulk) Save(ctx context.Context) ([]*Article, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Article, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArticleMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArticleCreateBulk) SaveX(ctx context.Context) []*Article {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArticleCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArticleCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}