package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ntiGideon/cvCraft/internal/models"
	"time"
)

// Template holds the schema definition for the Template entity.
type Template struct {
	ent.Schema
}

// Fields of the Template.
func (Template) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.JSON("sections", []string{}).Default([]string{
			"personal_info",
			"summary",
			"experience",
			"education",
			"skills",
			"projects",
			"certifications",
		}),
		field.JSON("config", &models.TemplateConfig{}).
			Default(&models.TemplateConfig{
				ColorScheme: models.ColorScheme{
					Primary:    "#4a6cf7",
					Secondary:  "#6c757d",
					Background: "#ffffff",
					Text:       "#212529",
					Accent:     "#fd7e14",
					Header:     "#343a40",
				},
				FontSettings: models.FontSettings{
					HeaderFont: "Roboto",
					BodyFont:   "Open Sans",
					FontSize:   "14px",
					LineHeight: "1.6",
				},
				SectionOrder: []string{
					"personal_info",
					"summary",
					"experience",
					"education",
					"skills",
				},
				LayoutOptions: models.LayoutOptions{
					Columns:        1,
					SectionSpacing: "1.5rem",
					HeaderStyle:    "underline",
				},
			}),
		field.String("thumbnail_url").Optional(),
		field.String("preview_url").Optional(),
		field.String("style").Default("classic"), // classic, modern, creative, etc.
		field.Bool("premium").Default(false),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Template.
func (Template) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("templates").Unique(), // For user-created templates
		edge.To("resumes", Resume.Type),
	}
}
