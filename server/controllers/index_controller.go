package controllers

import (
	"html/template"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	indexTemplate *template.Template
}

func NewIndexController() *IndexController {
	indexTemplate := template.Must(template.ParseFiles(filepath.Join("web", "index.html")))
	return &IndexController{indexTemplate}
}

func (healthController *IndexController) Get(context *gin.Context) {
	healthController.indexTemplate.Execute(context.Writer, nil)
}
