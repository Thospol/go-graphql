package gql

import "github.com/graphql-go/graphql"

// Model describes a graphql object containing a Model
var Model = graphql.NewObject(graphql.ObjectConfig{
	Name: "Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"createdAt": &graphql.Field{
			Type: graphql.DateTime,
		},
		"updatedAt": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deletedAt": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

// User describes a graphql object containing a User
var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"model": &graphql.Field{
				Type: Model,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"initialsName": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"profilePicture": &graphql.Field{
				Type: graphql.String,
			},
			"numberOfUnreadNotifications": &graphql.Field{
				Type: graphql.Int,
			},
			"deactivatedAt": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
