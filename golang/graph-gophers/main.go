package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type gopherHelloWorldResolver struct{}

func (r *gopherHelloWorldResolver) Hello() string {
	return "world"
}

var gopherSchema = graphql.MustParseSchema(`
schema {
  query: Query
}
type Query {
  hello: String!
}
`, &gopherHelloWorldResolver{})

func gophersGraphQLHandler() gin.HandlerFunc {
	r := &relay.Handler{Schema: gopherSchema}

	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	r := gin.New()
	r.POST("/graphql", gophersGraphQLHandler())

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with POST      : curl -g --request POST 'http://localhost:8080/graphql?query={hello}'")
	r.Run() // listen and serve on 0.0.0.0:8080
}
