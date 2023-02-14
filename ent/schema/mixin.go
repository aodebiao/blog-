package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type CommonMixin struct {
	mixin.Schema
}


func (CommonMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Optional().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			Optional().
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Default(nil).
			Optional().
			UpdateDefault(nil),
		field.Int("del_flag").
			Default(1).
			Optional(),
	}
}