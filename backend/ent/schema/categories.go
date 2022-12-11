package schema

import (
	"entgo.io/ent"
	// "entgo.io/contrib/entproto"
	// "entgo.io/ent/schema"
	// "entgo.io/ent/schema/edge"
)

// Categories holds the schema definition for the Categories entity.
// name 
type Categories struct {
	ent.Schema
}

// Fields of the Categories.
func (Categories) Fields() []ent.Field {
	return nil
}

// Edges of the Categories.
func (Categories) Edges() []ent.Edge {
	return nil
}
