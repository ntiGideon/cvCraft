package service

import (
	"github.com/fogleman/gg"
	"github.com/google/uuid"
	"github.com/ntiGideon/cvCraft/internal/models"
	"path/filepath"
)

type ThumbnailService struct {
	UploadPath string `json:"upload_path"`
}

func (ts *ThumbnailService) GenerateFromConfig(config *models.TemplateConfig) (string, error) {
	// Create a canvas
	const width = 400
	const height = 300
	dc := gg.NewContext(width, height)

	// Draw background
	dc.SetHexColor(config.ColorScheme.Background)
	dc.Clear()

	// Draw template preview elements
	dc.SetHexColor(config.ColorScheme.Primary)
	dc.DrawRectangle(50, 50, 300, 50)
	dc.Fill()

	dc.SetHexColor(config.ColorScheme.Text)
	if err := dc.LoadFontFace("static/fonts/"+config.FontSettings.HeaderFont+".ttf", 24); err == nil {
		dc.DrawStringAnchored(config.Name, width/2, 80, 0.5, 0.5)
	}

	// Save to file
	filename := "template_" + uuid.New().String() + ".png"
	path := filepath.Join(ts.UploadPath, filename)

	err := dc.SavePNG(path)
	if err != nil {
		return "", err
	}

	return "/" + filename, nil
}
