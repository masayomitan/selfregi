package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/dialect"
)

// Categories holds the schema definition for the Categories entity.
// name 
type Categories struct {
	ent.Schema
}

// Fields of the Categories.
func (Categories) Fields() []ent.Field {
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

// Edges of the Categories.
func (Categories) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", Item.Type),
	}
}
