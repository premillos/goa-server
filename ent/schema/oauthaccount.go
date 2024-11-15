package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OAuthAccount holds the schema definition for the OAuthAccount entity.
type OAuthAccount struct {
	ent.Schema
}

// Fields of the OAuthAccount.
func (OAuthAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("provider"),                 // OAuth 服务提供商名称（如 "google"、"facebook"）
		field.String("provider_user_id"),         // OAuth 服务用户的唯一 ID（如 Google ID、Facebook ID）
		field.String("access_token").Optional(),  // 访问令牌
		field.String("refresh_token").Optional(), // 刷新令牌
		field.Time("token_expiry").Optional(),    // 令牌过期时间
	}
}

// Edges of the OAuthAccount.
func (OAuthAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("oauth_accounts").
			Unique(),
	}
}
