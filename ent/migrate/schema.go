// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticleColumns holds the columns for the "article" table.
	ArticleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "del_flag", Type: field.TypeInt, Nullable: true, Default: 1},
		{Name: "category_id", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(64)"}},
		{Name: "content", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"mysql": "longtext"}},
		{Name: "status", Type: field.TypeInt8, Nullable: true, SchemaType: map[string]string{"mysql": "tinyint(1)"}},
		{Name: "cover_url", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "varchar(255)"}},
	}
	// ArticleTable holds the schema information for the "article" table.
	ArticleTable = &schema.Table{
		Name:       "article",
		Columns:    ArticleColumns,
		PrimaryKey: []*schema.Column{ArticleColumns[0]},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"mysql": "varchar(32)"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "del_flag", Type: field.TypeInt, Nullable: true, Default: 1},
		{Name: "user_name", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(16)"}},
		{Name: "password", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(255)"}},
		{Name: "status", Type: field.TypeInt8, Nullable: true, SchemaType: map[string]string{"mysql": "tinyint(1)"}},
		{Name: "avatar_url", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(255)"}},
		{Name: "login_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"mysql": "varchar(32)"}},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticleTable,
		UserTable,
	}
)

func init() {
	ArticleTable.Annotation = &entsql.Annotation{
		Table: "article",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
