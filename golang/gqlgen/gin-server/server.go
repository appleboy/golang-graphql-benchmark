package main

import (
	"fmt"

	"github.com/99designs/gqlgen/handler"
	"github.com/appleboy/golang-graphql-benchmark/golang/gqlgen"
	"github.com/gin-gonic/gin"
)

// Handler initializes the graphql middleware.
func Handler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &gqlgen.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.New()
	r.POST("/graphql", Handler())

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with POST      : curl -g --request POST 'http://localhost:8080/graphql?query={hello}'")
	r.Run()
}
