package handler

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// SetupFileServer sets up a static file serverer
func SetupFileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		zap.L().Info(
			"serving file",
			zap.String("r.URL", r.URL.String()),
			zap.String("path", path),
			zap.String("pathPrefix", pathPrefix),
		)
		fs.ServeHTTP(w, r)
	})
}
