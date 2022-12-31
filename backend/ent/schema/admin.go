package schema

import (
	"entgo.io/ent"
	// "entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
		field.String("password").Nillable(),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return nil
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStamp{},
	}
}