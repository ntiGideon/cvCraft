package main

import (
	"bytes"
	"github.com/ntiGideon/cvCraft/internal/models"
)

func GenerateCSS(config *models.TemplateConfig) string {
	var css bytes.Buffer

	css.WriteString(`
        :root {
            --primary-color: ` + config.ColorScheme.Primary + `;
            --secondary-color: ` + config.ColorScheme.Secondary + `;
            --background-color: ` + config.ColorScheme.Background + `;
            --text-color: ` + config.ColorScheme.Text + `;
            --accent-color: ` + config.ColorScheme.Accent + `;
            --header-color: ` + config.ColorScheme.Header + `;
            
            --header-font: ` + config.FontSettings.HeaderFont + `;
            --body-font: ` + config.FontSettings.BodyFont + `;
            --base-font-size: ` + config.FontSettings.FontSize + `;
            --line-height: ` + config.FontSettings.LineHeight + `;
        }
        
        body {
            font-family: var(--body-font);
            font-size: var(--base-font-size);
            line-height: var(--line-height);
            color: var(--text-color);
            background-color: var(--background-color);
        }
        
        h1, h2, h3 {
            font-family: var(--header-font);
            color: var(--header-color);
        }
        
        a {
            color: var(--primary-color);
        }
    `)

	return css.String()
}
