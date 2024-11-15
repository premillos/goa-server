// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"com.goa/ent/oauthaccount"
	"com.goa/ent/user"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// OAuthAccount is the model entity for the OAuthAccount schema.
type OAuthAccount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider string `json:"provider,omitempty"`
	// ProviderUserID holds the value of the "provider_user_id" field.
	ProviderUserID string `json:"provider_user_id,omitempty"`
	// AccessToken holds the value of the "access_token" field.
	AccessToken string `json:"access_token,omitempty"`
	// RefreshToken holds the value of the "refresh_token" field.
	RefreshToken string `json:"refresh_token,omitempty"`
	// TokenExpiry holds the value of the "token_expiry" field.
	TokenExpiry time.Time `json:"token_expiry,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OAuthAccountQuery when eager-loading is set.
	Edges               OAuthAccountEdges `json:"edges"`
	user_oauth_accounts *uuid.UUID
	selectValues        sql.SelectValues
}

// OAuthAccountEdges holds the relations/edges for other nodes in the graph.
type OAuthAccountEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OAuthAccountEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OAuthAccount) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauthaccount.FieldID:
			values[i] = new(sql.NullInt64)
		case oauthaccount.FieldProvider, oauthaccount.FieldProviderUserID, oauthaccount.FieldAccessToken, oauthaccount.FieldRefreshToken:
			values[i] = new(sql.NullString)
		case oauthaccount.FieldTokenExpiry:
			values[i] = new(sql.NullTime)
		case oauthaccount.ForeignKeys[0]: // user_oauth_accounts
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OAuthAccount fields.
func (oa *OAuthAccount) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauthaccount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oa.ID = int(value.Int64)
		case oauthaccount.FieldProvider:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider", values[i])
			} else if value.Valid {
				oa.Provider = value.String
			}
		case oauthaccount.FieldProviderUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider_user_id", values[i])
			} else if value.Valid {
				oa.ProviderUserID = value.String
			}
		case oauthaccount.FieldAccessToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_token", values[i])
			} else if value.Valid {
				oa.AccessToken = value.String
			}
		case oauthaccount.FieldRefreshToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refresh_token", values[i])
			} else if value.Valid {
				oa.RefreshToken = value.String
			}
		case oauthaccount.FieldTokenExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field token_expiry", values[i])
			} else if value.Valid {
				oa.TokenExpiry = value.Time
			}
		case oauthaccount.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_oauth_accounts", values[i])
			} else if value.Valid {
				oa.user_oauth_accounts = new(uuid.UUID)
				*oa.user_oauth_accounts = *value.S.(*uuid.UUID)
			}
		default:
			oa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OAuthAccount.
// This includes values selected through modifiers, order, etc.
func (oa *OAuthAccount) Value(name string) (ent.Value, error) {
	return oa.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the OAuthAccount entity.
func (oa *OAuthAccount) QueryUser() *UserQuery {
	return NewOAuthAccountClient(oa.config).QueryUser(oa)
}

// Update returns a builder for updating this OAuthAccount.
// Note that you need to call OAuthAccount.Unwrap() before calling this method if this OAuthAccount
// was returned from a transaction, and the transaction was committed or rolled back.
func (oa *OAuthAccount) Update() *OAuthAccountUpdateOne {
	return NewOAuthAccountClient(oa.config).UpdateOne(oa)
}

// Unwrap unwraps the OAuthAccount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oa *OAuthAccount) Unwrap() *OAuthAccount {
	_tx, ok := oa.config.driver.(*txDriver)
	if !ok {
		panic("ent: OAuthAccount is not a transactional entity")
	}
	oa.config.driver = _tx.drv
	return oa
}

// String implements the fmt.Stringer.
func (oa *OAuthAccount) String() string {
	var builder strings.Builder
	builder.WriteString("OAuthAccount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oa.ID))
	builder.WriteString("provider=")
	builder.WriteString(oa.Provider)
	builder.WriteString(", ")
	builder.WriteString("provider_user_id=")
	builder.WriteString(oa.ProviderUserID)
	builder.WriteString(", ")
	builder.WriteString("access_token=")
	builder.WriteString(oa.AccessToken)
	builder.WriteString(", ")
	builder.WriteString("refresh_token=")
	builder.WriteString(oa.RefreshToken)
	builder.WriteString(", ")
	builder.WriteString("token_expiry=")
	builder.WriteString(oa.TokenExpiry.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// OAuthAccounts is a parsable slice of OAuthAccount.
type OAuthAccounts []*OAuthAccount
