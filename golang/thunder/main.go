package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

type thunderServer struct{}

// registerQuery registers the root query type.
func (s *thunderServer) registerQuery(schema *schemabuilder.Schema) {
	obj := schema.Query()
	obj.FieldFunc("hello", func() string {
		return "world"
	})
}

func (s *thunderServer) schema() *graphql.Schema {
	builder := schemabuilder.NewSchema()
	s.registerQuery(builder)
	return builder.MustBuild()
}

func thunderGraphQLHandler() gin.HandlerFunc {
	server := &thunderServer{}
	thunderSchema := server.schema()
	introspection.AddIntrospectionToSchema(thunderSchema)
	r := graphql.HTTPHandler(thunderSchema)

	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	r := gin.New()
	r.POST("/graphql", thunderGraphQLHandler())

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with POST      : curl -g --request POST 'http://localhost:8080/graphql?query={hello}'")
	r.Run() // listen and serve on 0.0.0.0:8080
}
