package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Certification holds the schema definition for the Certification entity.
type Certification struct {
	ent.Schema
}

// Fields of the Certification.
func (Certification) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("organization").NotEmpty(),
		field.Time("issue_date"),
		field.Time("expiration_date").Optional(),
		field.String("credential_id").Optional(),
		field.String("url").Optional(),
	}
}

// Edges of the Certification.
func (Certification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("certifications").Required(),
	}
}
