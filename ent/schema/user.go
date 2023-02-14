package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{

		field.String("id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Unique(),

		field.String("user_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(16)", // Override MySQL.
		}),

		field.String("password").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.Int8("del_flag").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Optional(),

		field.String("avatar_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}),

		field.String("login_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional(),

		field.Time("create_at").SchemaType(map[string]string{
			dialect.MySQL: "date", // Override MySQL.
		}).Optional(),

		field.Time("update_at").SchemaType(map[string]string{
			dialect.MySQL: "date", // Override MySQL.
		}).Optional(),

		field.Time("delete_at").SchemaType(map[string]string{
			dialect.MySQL: "date", // Override MySQL.
		}).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}
