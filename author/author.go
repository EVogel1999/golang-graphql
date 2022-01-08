package author

import (
	"errors"
	"github.com/graphql-go/graphql"
	"golang-graphql/database"
)

// authorType GraphQL author type
var authorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Author",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"pseudonym": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// FindAuthorById GraphQL function that will resolve to a specific author
var FindAuthorById = &graphql.Field{
	Type:        authorType,
	Description: "Gets a single author",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(string)
		if ok {
			author, err := database.GetAuthor(id)
			if err != nil {
				return database.Author{}, err
			}
			return author, nil
		}
		return database.Author{}, errors.New("could not parse id from query")
	},
}

// GetAuthors GraphQL function that will resolve to the entire list of authors
var GetAuthors = &graphql.Field{
	Type:        graphql.NewList(authorType),
	Description: "Gets the list of authors",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return database.GetAuthors(), nil
	},
}
