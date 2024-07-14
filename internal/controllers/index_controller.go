package controllers

import (
	"html/template"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (healthController *IndexController) Index(context *gin.Context) {
	indexTemplate.Execute(context.Writer, nil)
}

var indexTemplate = template.Must(template.ParseFiles(filepath.Join("web", "index.html")))

func NewIndexController() *IndexController {
	return &IndexController{}
}
