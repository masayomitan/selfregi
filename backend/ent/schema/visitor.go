package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Visitor holds the schema definition for the Visitor entity.
type Visitor struct {
	ent.Schema
}

// Fields of the Visitor.
func (Visitor) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
		field.Int("sex"),
	}
}

// Edges of the Visitor.
func (Visitor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("managed_accounts", Account.Type),
		edge.To("carts", Cart.Type),
	}
}

func (Visitor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}