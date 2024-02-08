package schema

import "entgo.io/ent"

// Serie holds the schema definition for the Serie entity.
type Serie struct {
	ent.Schema
}

// Fields of the Serie.
func (Serie) Fields() []ent.Field {
  return []ent.Field{
    field.JSON("referential_id", []string{}).


    field.Float("rank").
        Optional(),
    field.Bool("active").
        Default(false),
    field.String("name").
        Unique(),
    field.Time("created_at").
        Default(time.Now),
    field.JSON("url", &url.URL{}).
        Optional(),
    field.JSON("strings", []string{}).
        Optional(),
    field.Enum("state").
        Values("on", "off").
        Optional(),
    field.UUID("uuid", uuid.UUID{}).
        Default(uuid.New),
  }
}

"id": 3856,
            "series": "The 06 Protocol (2022)",
            "year_began": 2022,
            "issue_count": 3,

// Edges of the Serie.
func (Serie) Edges() []ent.Edge {
	return nil
}
