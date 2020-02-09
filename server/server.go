package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"github.com/simple-gql-api/gql"
)

type Server struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

func (s *Server) GraphQL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "No graphql query supplied in request body", 400)
			return
		}

		var rBody reqBody
		if err := json.NewDecoder(r.Body).Decode(*rBody); err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
			return
		}

		result := gql.ExecuteQuery(rBody.Query, *s.GqlSchema)
		render.JSON(w, r, result)
	}
}
