package schema

import (
	"entgo.io/ent"
	// "entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	// "github.com/shopspring/decimal"
	"entgo.io/ent/schema/edge"
	"time"
)

// Cart holds the schema definition for the Cart entity.
type Cart struct {
	ent.Schema
}

// Fields of the Cart.
// visitor_id cart_status
func (Cart) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Nillable(),
	}
}

func (Cart) Annotations() []schema.Annotation {
	return []schema.Annotation{
			entproto.Message(),
	}
}

// Edges of the Cart.
func (Cart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cart_details", CartDetail.Type).
				Ref("carts").
				Annotations(entproto.Field(5)),
		edge.From("visitor", Visitor.Type).
				Ref("cart").
				Unique().
				Required(),
	}
}