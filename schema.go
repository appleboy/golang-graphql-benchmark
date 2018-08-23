package graphql_benchmark

import (
	ggql "github.com/graph-gophers/graphql-go"
	"github.com/graphql-go/graphql"
	pgql "github.com/playlyfe/go-graphql"
	thql "github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

var graphQLGoSchema, _ = graphql.NewSchema(
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

var playlyfeSchema = `
    type RootQueryType {
        hello: String
    }
  `
var playlyfeResolvers = map[string]interface{}{
	"RootQueryType/hello": func(params *pgql.ResolveParams) (interface{}, error) {
		return "world", nil
	},
}
var playlyfeExecutor, _ = pgql.NewExecutor(playlyfeSchema, "RootQueryType", "", playlyfeResolvers)

type gopherHelloWorldResolver struct{}

func (r *gopherHelloWorldResolver) Hello() string {
	return "world"
}

var gopherSchema = ggql.MustParseSchema(`
schema {
  query: Query
}
type Query {
  hello: String!
}
`, &gopherHelloWorldResolver{})

type thunderServer struct{}

// registerQuery registers the root query type.
func (s *thunderServer) registerQuery(schema *schemabuilder.Schema) {
	obj := schema.Query()
	obj.FieldFunc("hello", func() string {
		return "world"
	})
}

func (s *thunderServer) schema() *thql.Schema {
	builder := schemabuilder.NewSchema()
	s.registerQuery(builder)
	return builder.MustBuild()
}
