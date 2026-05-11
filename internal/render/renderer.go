package render

import (
	"html/template"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(pattern string) (*TemplateRenderer, error) {
	funcs := template.FuncMap{
		"weatherIcon":        weatherIcon,
		"weatherDescription": weatherDescription,
		"temperatureC":       temperatureC,
	}

	templates, err := template.New("").Funcs(funcs).ParseGlob(pattern)
	if err != nil {
		return nil, err
	}

	return &TemplateRenderer{templates: templates}, nil
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
