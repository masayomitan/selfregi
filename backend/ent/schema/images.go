package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/dialect"
)

// Images holds the schema definition for the Images entity.
type Images struct {
	ent.Schema
}

// Fields of the Images.
func (Images) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}),
		field.String("path").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}),
	}
}

// Edges of the Images.
func (Images) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("items", Item.Type).
			Ref("images"),
	}
}

func (Images) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}
