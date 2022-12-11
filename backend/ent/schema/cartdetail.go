package schema

import (
	"entgo.io/ent"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// CartDetail holds the schema definition for the CartDetail entity.
type CartDetail struct {
	ent.Schema
}

// Fields of the CartDetail.
// cart_id product_id category_id data quantity options price tax tax_rate discount_price discount_rate discount_name
func (CartDetail) Fields() []ent.Field {
	return nil
}

func (CartDetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
			entproto.Message(),
	}
}

// Edges of the CartDetail.
func (CartDetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carts", Cart.Type).
				Unique().
				Annotations(entproto.Field(5)),
	}
}