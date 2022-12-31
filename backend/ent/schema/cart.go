package schema

import (
	"entgo.io/ent"
	// "entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	// "github.com/shopspring/decimal"
	"entgo.io/ent/schema/edge"
)

// Cart holds the schema definition for the Cart entity.
type Cart struct {
	ent.Schema
}

// Fields of the Cart.
func (Cart) Fields() []ent.Field {
	return []ent.Field{
		field.Int("visitor_id").
			Positive().
			Optional(),
		field.Uint("status").
			Nillable(),
	}
}

// Edges of the Cart.
func (Cart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cart_details", CartDetail.Type),
		edge.From("owner", Visitor.Type).
				Ref("carts").
				Unique().
				Required(),
	}
}

func (Cart) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}