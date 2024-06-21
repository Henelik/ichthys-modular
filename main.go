package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/henelik/ichthys-modular/pkg/config"
	"github.com/henelik/ichthys-modular/pkg/handler"
)

func main() {
	cfg := config.SetupConfig()

	// set up the logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(errors.Wrap(err, "failed to set up logger"))
	}

	zap.ReplaceGlobals(logger)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// set up the file server
	workDir, _ := os.Getwd()
	webDir := http.Dir(filepath.Join(workDir, cfg.WebDirectory))
	handler.SetupFileServer(r, "/web", webDir)

	zap.L().Info("running server at port 3000")

	http.ListenAndServe(":3000", r)
}
