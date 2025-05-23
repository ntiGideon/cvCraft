package models

import "github.com/ntiGideon/cvCraft/internal/validator"

type TemplateConfig struct {
	ColorScheme   ColorScheme   `json:"color_scheme"`
	FontSettings  FontSettings  `json:"font_settings"`
	SectionOrder  []string      `json:"section_order"`
	LayoutOptions LayoutOptions `json:"layout_options"`
}

type ColorScheme struct {
	Primary    string `json:"primary"`
	Secondary  string `json:"secondary"`
	Background string `json:"background"`
	Text       string `json:"text"`
	Accent     string `json:"accent"`
	Header     string `json:"header"`
}

type FontSettings struct {
	HeaderFont string `json:"header_font"`
	BodyFont   string `json:"body_font"`
	FontSize   string `json:"font_size"`
	LineHeight string `json:"line_height"`
}

type LayoutOptions struct {
	Columns        int    `json:"columns"` // 1 or 2 column layout
	SectionSpacing string `json:"section_spacing"`
	HeaderStyle    string `json:"header_style"` // underline, border-left, etc.
}

type CreateTemplateConfigDto struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Premium     bool   `form:"premium"`
	Config      *TemplateConfig

	validator.Validator `form:"-"`
}
