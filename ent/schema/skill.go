package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Skill holds the schema definition for the Skill entity.
type Skill struct {
	ent.Schema
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("category").NotEmpty(), // e.g., "Technical", "Language", "Soft"
		field.Enum("level").Values("beginner", "intermediate", "advanced", "expert").Default("intermediate"),
	}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("skills").Required(),
	}
}
