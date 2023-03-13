package schema

import "entgo.io/ent"

// Serie holds the schema definition for the Serie entity.
type Serie struct {
	ent.Schema
}

// Fields of the Serie.
func (Serie) Fields() []ent.Field {
	return nil
}

// Edges of the Serie.
func (Serie) Edges() []ent.Edge {
	return nil
}
