package graphql_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func runRequest(B *testing.B, r *gin.Engine, method, path string, body io.Reader) {
	// create fake request
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		panic(err)
	}
	w := httptest.NewRecorder()
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		r.ServeHTTP(w, req)
	}
}

var newSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "RootQueryType",
				Fields: graphql.Fields{
					"hello": &graphql.Field{
						Type: graphql.String,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return "world", nil
						},
					},
				},
			}),
	},
)

func goGraphQLHandler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.New(&handler.Config{
		Schema: &newSchema,
		Pretty: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Handler initializes the graphql middleware.
func gophersGraphQLHandler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	r := &relay.Handler{Schema: schema3}

	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

func BenchmarkGinHttpRoute(B *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.New()
	router.POST("/hello", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"data": gin.H{"hello": "world"},
			},
		)
	})
	runRequest(B, router, "POST", "/hello", strings.NewReader(`{"query":"{ hello }", "operationName":"", "variables": null}`))
}

func BenchmarkGinGoGraphQLRoute(B *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.New()
	router.GET("/graphql", goGraphQLHandler())
	runRequest(B, router, "GET", "/graphql?query={hello}", nil)
}

func BenchmarkGinGopherGraphQLRoute(B *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.New()
	router.POST("/graphql", gophersGraphQLHandler())
	runRequest(B, router, "POST", "/graphql", strings.NewReader(`{"query":"{ hello }", "operationName":"", "variables": null}`))
}
