package router

import (
	"compress/flate"
	"net/http"
	"strings"

	"github.com/NYTimes/gziphandler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/thospol/go-graphql/endpoint"
	"github.com/thospol/go-graphql/gql"
)

// New new router
func New() *chi.Mux {
	// Create a new router
	r := chi.NewRouter()

	// Add some middleware to our router
	r.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		gziphandler.MustNewGzipLevelHandler(flate.BestSpeed),
		middleware.RequestID,
		middleware.RealIP,
		middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes,    // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,       // recover from panics without crashing server
	)

	// Create our root query for graphql
	rootQuery := gql.NewRootQuery()

	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		panic(err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	endpoint := endpoint.NewEndpoint(&sc)

	r.Route("/", func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Get("/statics/*", func(w http.ResponseWriter, r *http.Request) {
				routeCtx := chi.RouteContext(r.Context())
				prefix := strings.TrimSuffix(routeCtx.RoutePattern(), "/*")
				fs := http.StripPrefix(prefix, http.FileServer(http.Dir("./assets")))
				fs.ServeHTTP(w, r)
			})

			r.Post("/graphql", endpoint.Graphql())
		})
	})

	return r
}
