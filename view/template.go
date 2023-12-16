package view

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTmpl *template.Template
}

func NewTemplate(fs fs.FS, pattern ...string) *Template {
	tmpl, err := template.ParseFS(fs, pattern...)
	if err != nil {
		panic(err)
	}
	return &Template{htmlTmpl: tmpl}
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTmpl.Execute(w, data)
	if err != nil {
		log.Println("execution error:", err)
		http.Error(w, "error executing the template", http.StatusInternalServerError)
		return
	}
}
