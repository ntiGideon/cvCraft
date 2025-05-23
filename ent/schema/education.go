package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Education holds the schema definition for the Education entity.
type Education struct {
	ent.Schema
}

// Fields of the Education.
func (Education) Fields() []ent.Field {
	return []ent.Field{
		field.String("institution").NotEmpty(),
		field.String("degree").NotEmpty(),
		field.String("field").NotEmpty(),
		field.Time("start_date"),
		field.Time("end_date").Optional(),
		field.String("description").Optional(),
		field.String("grade").Optional(),
	}
}

// Edges of the Education.
func (Education) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("educations").Required(),
	}
}
