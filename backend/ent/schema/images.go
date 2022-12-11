package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	// "entgo.io/ent/schema/field"
)

// Images holds the schema definition for the Images entity.
// path name
type Images struct {
	ent.Schema
}

// Fields of the Images.
func (Images) Fields() []ent.Field {
	return nil
}

// Edges of the Images.
func (Images) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("items", Item.Type).
				Ref("images"),
}
}
