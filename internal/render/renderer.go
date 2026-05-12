package render

import (
	"html/template"
	"net/http"
	"weather-viewer/templates"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() (*TemplateRenderer, error) {
	funcs := template.FuncMap{
		"weatherIcon":        weatherIcon,
		"weatherDescription": weatherDescription,
		"temperatureC":       temperatureC,
	}

	tmpl, err := template.New("").
		Funcs(funcs).
		ParseFS(templates.FS, "*.html")
	if err != nil {
		return nil, err
	}

	return &TemplateRenderer{templates: tmpl}, nil
}

func (r *TemplateRenderer) Render(w http.ResponseWriter, name string, data any) {
	r.RenderStatus(w, http.StatusOK, name, data)
}

func (r *TemplateRenderer) RenderStatus(w http.ResponseWriter, status int, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if err := r.templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "template render error", http.StatusInternalServerError)
	}
}
