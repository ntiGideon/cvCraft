package dbClient

import (
	"context"
	"github.com/ntiGideon/cvCraft/ent"
	"github.com/ntiGideon/cvCraft/internal/models"
)

type TemplateRepository struct {
	Client *ent.Client
}

func (r *TemplateRepository) CreateTemplate(ctx context.Context, dto *models.CreateTemplateConfigDto) error {
	_, err := r.Client.Template.Create().
		SetName(dto.Name).
		SetDescription(dto.Description).
		SetPremium(dto.Premium).
		SetConfig(&dto.Config).
		Save(ctx)
	return err
}

func (r *TemplateRepository) ListTemplates(ctx context.Context) (ent.Templates, error) {
	templates, err := r.Client.Template.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return templates, nil
}
