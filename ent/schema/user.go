package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// 一个手机号和一个邮箱对应着一个用户  如果通过oauth登录 必须要获取邮箱地址  所以邮箱和手机号必须唯一

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("show_id").Optional().Default(""),
		field.String("name").
			Optional().
			Default(""),
		field.String("email").
			Optional().
			Unique(),
		field.String("password").
			Optional().
			Default(""),
		field.String("salt").NotEmpty(),
		field.String("phone").
			Optional().
			Unique(),
		field.String("avatar").
			Optional().
			Default(""),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("oauth_accounts", OAuthAccount.Type),
	}
}
