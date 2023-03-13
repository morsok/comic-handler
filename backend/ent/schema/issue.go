package schema

import "entgo.io/ent"

// Issue holds the schema definition for the Issue entity.
type Issue struct {
	ent.Schema
}

// Fields of the Issue.
func (Issue) Fields() []ent.Field {
	return nil
}

// Edges of the Issue.
func (Issue) Edges() []ent.Edge {
	return nil
}
