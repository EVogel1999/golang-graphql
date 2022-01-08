package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"golang-graphql/author"
	"net/http"
)

// Creates the root query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"author":  author.FindAuthorById,
		"authors": author.GetAuthors,
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})

func main() {
	handler := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil)
}
