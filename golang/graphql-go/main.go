package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Schema
var fields = graphql.Fields{
	"hello": &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	},
}
var rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
var schemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
var schema, _ = graphql.NewSchema(schemaConfig)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

// Handler initializes the graphql middleware.
func Handler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   false,
		GraphiQL: false,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	r := gin.New()
	r.POST("/graphql", Handler())

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={hello}'")
	r.Run() // listen and serve on 0.0.0.0:8080
}
