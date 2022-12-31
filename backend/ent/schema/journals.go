package schema

import (
	"entgo.io/ent"
	// "entgo.io/ent/schema/edge"
	// "entgo.io/ent/schema/field"
)

// Journals holds the schema definition for the Journals entity.
// class data
type Journals struct {
	ent.Schema
}

// Fields of the Journals.
func (Journals) Fields() []ent.Field {
	return nil
}

// Edges of the Journals.
func (Journals) Edges() []ent.Edge {
	return nil
}

func (Journals) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}