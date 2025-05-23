package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PersonalInfo holds the schema definition for the PersonalInfo entity.
type PersonalInfo struct {
	ent.Schema
}

// Fields of the PersonalInfo.
func (PersonalInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("profession").NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("phone").Optional(),
		field.String("address").Optional(),
		field.String("website").Optional(),
		field.String("linkedin").Optional(),
		field.String("github").Optional(),
		field.String("twitter").Optional(),
		field.String("bio").Optional(),
		field.String("photo_url").Optional(),
	}
}

// Edges of the PersonalInfo.
func (PersonalInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("personal_info").Unique().Required(),
	}
}
