package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("ユーザー名"),
		field.String("email").
			NotEmpty().
			Unique().
			Comment("メールアドレス"),
		field.String("password_hash").
			NotEmpty().
			Sensitive().
			Comment("パスワード"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("作成日時"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新日時"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
