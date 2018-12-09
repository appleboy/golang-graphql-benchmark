package graphqlbenchmark

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gqlH "github.com/99designs/gqlgen/handler"
	"github.com/appleboy/golang-graphql-benchmark/golang/gqlgen"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graphql-go/handler"
	thql "github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
)

func runRequest(B *testing.B, r *gin.Engine, method, path string) {
	// create fake request
	req, err := http.NewRequest(method, path, strings.NewReader(`{"query":"{ hello }", "operationName":"", "variables": null}`))
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

func goGraphQLHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema: &graphQLGoSchema,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func gophersGraphQLHandler() gin.HandlerFunc {
	r := &relay.Handler{Schema: gopherSchema}

	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

func thunderGraphQLHandler() gin.HandlerFunc {
	server := &thunderServer{}
	thunderSchema := server.schema()
	introspection.AddIntrospectionToSchema(thunderSchema)
	r := thql.HTTPHandler(thunderSchema)

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
	runRequest(B, router, "POST", "/hello")
}

func goGQLGenHandler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := gqlH.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &gqlgen.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func BenchmarkGinGQLGenRoute(B *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.New()
	router.POST("/graphql", goGQLGenHandler())
	runRequest(B, router, "POST", "/graphql")
}

func BenchmarkGinGoGraphQLRoute(B *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.New()
	router.POST("/graphql", goGraphQLHandler())
	runRequest(B, router, "POST", "/graphql")
}

func BenchmarkGinGopherGraphQLRoute(B *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.New()
	router.POST("/graphql", gophersGraphQLHandler())
	runRequest(B, router, "POST", "/graphql")
}

func BenchmarkGinThunderGraphQLRoute(B *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.New()
	router.POST("/graphql", thunderGraphQLHandler())
	runRequest(B, router, "POST", "/graphql")
}
