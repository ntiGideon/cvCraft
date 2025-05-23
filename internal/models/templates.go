package models

import (
	"context"
	"github.com/ntiGideon/cvCraft/ent"
)

type TemplatesModelInterface interface{}

type TemplatesModel struct {
	Db *ent.Client
}

func (m *TemplatesModel) CreateTemplate(ctx context.Context, dto *CreateTemplateConfigDto) error {
	_, err := m.Db.Template.Create().
		SetName(dto.Name).
		SetDescription(dto.Description).
		SetPremium(dto.Premium).
		SetConfig(dto.Config).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
