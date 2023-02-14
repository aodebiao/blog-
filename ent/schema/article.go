package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{

		field.String("id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}).Unique(),

		field.String("category_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(64)", // Override MySQL.
		}),

		field.Text("content").SchemaType(map[string]string{
			dialect.MySQL: "longtext", // Override MySQL.
		}).Optional(),

		field.Int8("del_flag").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
		}).Optional(),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint(1)", // Override MySQL.
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

		field.String("cover_url").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional(),
	}

}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}

func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "article"},
	}
}
