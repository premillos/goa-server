// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"com.goa/ent/oauthaccount"
	"com.goa/ent/user"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OAuthAccountCreate is the builder for creating a OAuthAccount entity.
type OAuthAccountCreate struct {
	config
	mutation *OAuthAccountMutation
	hooks    []Hook
}

// SetProvider sets the "provider" field.
func (oac *OAuthAccountCreate) SetProvider(s string) *OAuthAccountCreate {
	oac.mutation.SetProvider(s)
	return oac
}

// SetProviderUserID sets the "provider_user_id" field.
func (oac *OAuthAccountCreate) SetProviderUserID(s string) *OAuthAccountCreate {
	oac.mutation.SetProviderUserID(s)
	return oac
}

// SetAccessToken sets the "access_token" field.
func (oac *OAuthAccountCreate) SetAccessToken(s string) *OAuthAccountCreate {
	oac.mutation.SetAccessToken(s)
	return oac
}

// SetNillableAccessToken sets the "access_token" field if the given value is not nil.
func (oac *OAuthAccountCreate) SetNillableAccessToken(s *string) *OAuthAccountCreate {
	if s != nil {
		oac.SetAccessToken(*s)
	}
	return oac
}

// SetRefreshToken sets the "refresh_token" field.
func (oac *OAuthAccountCreate) SetRefreshToken(s string) *OAuthAccountCreate {
	oac.mutation.SetRefreshToken(s)
	return oac
}

// SetNillableRefreshToken sets the "refresh_token" field if the given value is not nil.
func (oac *OAuthAccountCreate) SetNillableRefreshToken(s *string) *OAuthAccountCreate {
	if s != nil {
		oac.SetRefreshToken(*s)
	}
	return oac
}

// SetTokenExpiry sets the "token_expiry" field.
func (oac *OAuthAccountCreate) SetTokenExpiry(t time.Time) *OAuthAccountCreate {
	oac.mutation.SetTokenExpiry(t)
	return oac
}

// SetNillableTokenExpiry sets the "token_expiry" field if the given value is not nil.
func (oac *OAuthAccountCreate) SetNillableTokenExpiry(t *time.Time) *OAuthAccountCreate {
	if t != nil {
		oac.SetTokenExpiry(*t)
	}
	return oac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (oac *OAuthAccountCreate) SetUserID(id uuid.UUID) *OAuthAccountCreate {
	oac.mutation.SetUserID(id)
	return oac
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (oac *OAuthAccountCreate) SetNillableUserID(id *uuid.UUID) *OAuthAccountCreate {
	if id != nil {
		oac = oac.SetUserID(*id)
	}
	return oac
}

// SetUser sets the "user" edge to the User entity.
func (oac *OAuthAccountCreate) SetUser(u *User) *OAuthAccountCreate {
	return oac.SetUserID(u.ID)
}

// Mutation returns the OAuthAccountMutation object of the builder.
func (oac *OAuthAccountCreate) Mutation() *OAuthAccountMutation {
	return oac.mutation
}

// Save creates the OAuthAccount in the database.
func (oac *OAuthAccountCreate) Save(ctx context.Context) (*OAuthAccount, error) {
	return withHooks(ctx, oac.sqlSave, oac.mutation, oac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oac *OAuthAccountCreate) SaveX(ctx context.Context) *OAuthAccount {
	v, err := oac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oac *OAuthAccountCreate) Exec(ctx context.Context) error {
	_, err := oac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oac *OAuthAccountCreate) ExecX(ctx context.Context) {
	if err := oac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oac *OAuthAccountCreate) check() error {
	if _, ok := oac.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required field "OAuthAccount.provider"`)}
	}
	if _, ok := oac.mutation.ProviderUserID(); !ok {
		return &ValidationError{Name: "provider_user_id", err: errors.New(`ent: missing required field "OAuthAccount.provider_user_id"`)}
	}
	return nil
}

func (oac *OAuthAccountCreate) sqlSave(ctx context.Context) (*OAuthAccount, error) {
	if err := oac.check(); err != nil {
		return nil, err
	}
	_node, _spec := oac.createSpec()
	if err := sqlgraph.CreateNode(ctx, oac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oac.mutation.id = &_node.ID
	oac.mutation.done = true
	return _node, nil
}

func (oac *OAuthAccountCreate) createSpec() (*OAuthAccount, *sqlgraph.CreateSpec) {
	var (
		_node = &OAuthAccount{config: oac.config}
		_spec = sqlgraph.NewCreateSpec(oauthaccount.Table, sqlgraph.NewFieldSpec(oauthaccount.FieldID, field.TypeInt))
	)
	if value, ok := oac.mutation.Provider(); ok {
		_spec.SetField(oauthaccount.FieldProvider, field.TypeString, value)
		_node.Provider = value
	}
	if value, ok := oac.mutation.ProviderUserID(); ok {
		_spec.SetField(oauthaccount.FieldProviderUserID, field.TypeString, value)
		_node.ProviderUserID = value
	}
	if value, ok := oac.mutation.AccessToken(); ok {
		_spec.SetField(oauthaccount.FieldAccessToken, field.TypeString, value)
		_node.AccessToken = value
	}
	if value, ok := oac.mutation.RefreshToken(); ok {
		_spec.SetField(oauthaccount.FieldRefreshToken, field.TypeString, value)
		_node.RefreshToken = value
	}
	if value, ok := oac.mutation.TokenExpiry(); ok {
		_spec.SetField(oauthaccount.FieldTokenExpiry, field.TypeTime, value)
		_node.TokenExpiry = value
	}
	if nodes := oac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oauthaccount.UserTable,
			Columns: []string{oauthaccount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_oauth_accounts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OAuthAccountCreateBulk is the builder for creating many OAuthAccount entities in bulk.
type OAuthAccountCreateBulk struct {
	config
	err      error
	builders []*OAuthAccountCreate
}

// Save creates the OAuthAccount entities in the database.
func (oacb *OAuthAccountCreateBulk) Save(ctx context.Context) ([]*OAuthAccount, error) {
	if oacb.err != nil {
		return nil, oacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oacb.builders))
	nodes := make([]*OAuthAccount, len(oacb.builders))
	mutators := make([]Mutator, len(oacb.builders))
	for i := range oacb.builders {
		func(i int, root context.Context) {
			builder := oacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OAuthAccountMutation)
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
					_, err = mutators[i+1].Mutate(root, oacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oacb *OAuthAccountCreateBulk) SaveX(ctx context.Context) []*OAuthAccount {
	v, err := oacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oacb *OAuthAccountCreateBulk) Exec(ctx context.Context) error {
	_, err := oacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oacb *OAuthAccountCreateBulk) ExecX(ctx context.Context) {
	if err := oacb.Exec(ctx); err != nil {
		panic(err)
	}
}