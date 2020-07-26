package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/appleboy/golang-graphql-benchmark/golang/gqlgen"
)

func main() {
	h := handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &gqlgen.Resolver{}}))
	http.Handle("/graphql", h)

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with GET      : curl -g --request GET 'http://localhost:8080/graphql?query={hello}'")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
