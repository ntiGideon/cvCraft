package models

import (
	"context"
)

type TemplateService interface {
	CreateTemplate(ctx context.Context, dto *CreateTemplateConfigDto) error
	//ListTemplates(ctx context.Context) (ent.Templates, error)
}
