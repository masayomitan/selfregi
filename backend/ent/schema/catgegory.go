package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/dialect"
)

// Categories holds the schema definition for the Categories entity.
type Category struct {
	ent.Schema
}

// Fields of the Categories.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}),
		field.Bool("is_display").
			Default(false),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", Item.Type),
	}
}

func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}