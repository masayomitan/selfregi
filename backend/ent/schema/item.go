package schema

import (
	"entgo.io/ent"
	// "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
// category_id name detail normal_price special_price tax_rate stock is_soldout is_display
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type).
				Unique(),
		edge.To("images", Images.Type),
	}
}
