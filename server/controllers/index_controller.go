package controllers

import (
	"html/template"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type IIndexController interface {
	Index(context *gin.Context)
}

type IndexController struct {
	indexTemplate *template.Template
}

func (healthController *IndexController) Index(context *gin.Context) {
	healthController.indexTemplate.Execute(context.Writer, nil)
}

func NewIndexController() IIndexController {
	indexTemplate := template.Must(template.ParseFiles(filepath.Join("web", "index.html")))
	return &IndexController{indexTemplate}
}
