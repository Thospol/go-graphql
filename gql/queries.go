package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/thospol/go-graphql/core/postgres"
	"github.com/thospol/go-graphql/repository"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

// NewRootQuery returns base query type. This is where we add all the base queries
func NewRootQuery() *Root {
	// Create a resolver holding our database. Resolver can be found in resolvers.go
	resolver := Resolver{
		database:       postgres.GetDatabase(),
		userRepository: repository.NewUserRepository(),
	}

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: graphql.NewList(User),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.UserResolver,
					},
				},
			},
		),
	}

	return &root
}
