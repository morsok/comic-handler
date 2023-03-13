package schema

import "entgo.io/ent"

// ReadingList holds the schema definition for the ReadingList entity.
type ReadingList struct {
	ent.Schema
}

// Fields of the ReadingList.
func (ReadingList) Fields() []ent.Field {
	return nil
}

// Edges of the ReadingList.
func (ReadingList) Edges() []ent.Edge {
	return nil
}
