package schema

import "entgo.io/ent"

// StoryArc holds the schema definition for the StoryArc entity.
type StoryArc struct {
	ent.Schema
}

// Fields of the StoryArc.
func (StoryArc) Fields() []ent.Field {
	return nil
}

// Edges of the StoryArc.
func (StoryArc) Edges() []ent.Edge {
	return nil
}
