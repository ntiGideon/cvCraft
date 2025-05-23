package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Resume holds the schema definition for the Resume entity.
type Resume struct {
	ent.Schema
}

// Fields of the Resume.
func (Resume) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Bool("shared").Default(false),
		field.String("share_code").Optional().Unique(),
		field.Enum("status").Values("draft", "published", "archived").Default("draft"),
	}
}

// Edges of the Resume.
func (Resume) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("resumes").Unique().Required(),
		edge.From("template", Template.Type).Ref("resumes").Unique(),
		edge.To("personal_info", PersonalInfo.Type).Unique(),
		edge.To("experiences", Experience.Type),
		edge.To("educations", Education.Type),
		edge.To("skills", Skill.Type),
		edge.To("projects", Project.Type),
		edge.To("certifications", Certification.Type),
	}
}
