package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	// "entgo.io/ent/schema/field"
)

// Visitor holds the schema definition for the Visitor entity.
// name sex
type Visitor struct {
	ent.Schema
}

// Fields of the Visitor.
func (Visitor) Fields() []ent.Field {
	return []ent.Field{
		
	}
}

// Edges of the Visitor.
func (Visitor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
				Unique(),
		edge.To("cart", Cart.Type).
				Unique(),
	}
}
