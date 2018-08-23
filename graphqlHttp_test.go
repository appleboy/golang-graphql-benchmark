package graphql_benchmark

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graphql-go/handler"
	thql "github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
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
		// contentType := w.Header().Get("Content-Type")
		// log.Println()
		// if contentType == "application/json" {
		// 	log.Println("test")
		// }
		// actualResponse := w.Body.String()
		// log.Println(actualResponse)
		// log.Println(contentType)
		// if actualResponse == `{"data":{"hello":"world"}}` {
		// 	log.Println("done")
		// }
	}
}

func goGraphQLHandler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	h := handler.New(&handler.Config{
		Schema: &graphQLGoSchema,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Handler initializes the graphql middleware.
func gophersGraphQLHandler() gin.HandlerFunc {
	// Creates a GraphQL-go HTTP handler with the defined schema
	r := &relay.Handler{Schema: gopherSchema}

	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

// Handler initializes the graphql middleware.
func thunderGraphQLHandler() gin.HandlerFunc {
	server := &thunderServer{}

	thunderSchema := server.schema()
	introspection.AddIntrospectionToSchema(thunderSchema)
	// Creates a GraphQL-go HTTP handler with the defined schema
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

func BenchmarkGinThunderGraphQLRoute(B *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	router := gin.New()
	router.POST("/graphql", thunderGraphQLHandler())
	runRequest(B, router, "POST", "/graphql", strings.NewReader(`{"query":"{ hello }", "operationName":"hello", "variables": null}`))
}
