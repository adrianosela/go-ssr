package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/adrianosela/go-ssr/internal/vite"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	viteBuildDir         = filepath.Join("web", "dist")
	viteManifestFilePath = filepath.Join(viteBuildDir, ".vite", "manifest.json")

	templatesFilePath = filepath.Join("templates", "*.html")
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	viteManifest, err := vite.LoadManifest(viteManifestFilePath)
	if err != nil {
		logger.Fatal("failed to load vite manifest", zap.String("path", viteManifestFilePath), zap.Error(err))
	}

	templates, err := template.New("root").Funcs(vite.GetAllHelpers(viteManifest)).ParseGlob(templatesFilePath)
	if err != nil {
		logger.Fatal("failed to initialize templates", zap.Error(err))
	}

	r := gin.Default()
	r.Static("/static/vite", viteBuildDir)
	r.GET("/", func(ctx *gin.Context) {
		buf := bytes.NewBuffer(nil)
		err := templates.Lookup("demo.html").Execute(buf, map[string]any{
			"Config": map[string]any{
				"page":      "home",
				"csrfToken": "abc123",
				"version":   "1234",
			},
		})
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to execute template: %v", err))
			return
		}

		ctx.Data(http.StatusOK, "text/html", buf.Bytes())
	})

	if err := r.Run(":8082"); err != nil {
		logger.Fatal("failed to run gin http server", zap.Error(err))
	}
}
