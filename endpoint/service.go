package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/thospol/go-graphql/gql"
)

// Endpoint new enpoint interface
type Endpoint interface {
	Graphql() http.HandlerFunc
}

// NewEndpoint new endpoint
func NewEndpoint(sc *graphql.Schema) Endpoint {
	return &endpoint{
		GqlSchema: sc,
	}
}

// Server will hold connection to the db as well as handlers
type endpoint struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

// Graphql returns an http.HandlerFunc for our /graphql endpoint
func (ep *endpoint) Graphql() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check to ensure query was provided in the request body
		if r.Body == nil {
			http.Error(w, "Must provide graphql query in request body", 400)
			return
		}

		var rBody reqBody

		// Decode the request body into rBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		// Execute graphql query
		result := gql.ExecuteQuery(rBody.Query, *ep.GqlSchema)

		// render.JSON comes from the chi/render package and handles
		// marshalling to json, automatically escaping HTML and setting
		// the Content-Type as application/json.
		render.JSON(w, r, result)
	}
}
