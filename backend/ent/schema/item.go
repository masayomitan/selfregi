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
		field.Int("category_id"),
		field.Int("is_display"),
		field.Int("tax"),
		field.Int("tax_rate"),
		field.Int("price"),
		field.Int("temporary_stock"),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Images.Type),
		edge.From("category", Categories.Type).
			Ref("items").
			Field("category_id").
			Unique().
			Required(),
	}
}

func (Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}