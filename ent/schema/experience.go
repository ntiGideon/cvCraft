package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Experience holds the schema definition for the Experience entity.
type Experience struct {
	ent.Schema
}

// Fields of the Experience.
func (Experience) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("company").NotEmpty(),
		field.String("location").Optional(),
		field.Time("start_date"),
		field.Time("end_date").Optional(),
		field.String("description").Optional(),
		field.String("skills").Optional(), // Comma-separated or JSON
		field.Bool("current").Default(false),
	}
}

// Edges of the Experience.
func (Experience) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("experiences").Required(),
	}
}
