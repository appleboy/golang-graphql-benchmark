package graphql_benchmark

import (
	"context"
	"testing"

	"github.com/graphql-go/graphql"
	thunder "github.com/samsarahq/thunder/graphql"
)

func BenchmarkGoGraphQLMaster(b *testing.B) {
	// Disable SpecifiedRules
	// graphql.SpecifiedRules = []graphql.ValidationRuleFn{}
	for i := 0; i < b.N; i++ {
		graphql.Do(graphql.Params{
			Schema:        graphQLGoSchema,
			RequestString: "{hello}",
		})
	}
}

func BenchmarkPlaylyfeGraphQLMaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		context := map[string]interface{}{}
		variables := map[string]interface{}{}
		playlyfeExecutor.Execute(context, "{hello}", variables, "")
	}
}

func BenchmarkGophersGraphQLMaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		variables := map[string]interface{}{}
		gopherSchema.Exec(ctx, "{hello}", "", variables)
	}
}

func BenchmarkThunderGraphQLMaster(b *testing.B) {
	noArguments := func(json interface{}) (interface{}, error) {
		return nil, nil
	}
	var query = &thunder.Object{
		Name:   "Query",
		Fields: make(map[string]*thunder.Field),
	}

	query.Fields["hello"] = &thunder.Field{
		Resolve: func(ctx context.Context, source, args interface{}, selectionSet *thunder.SelectionSet) (interface{}, error) {
			return "world", nil
		},
		Type:           &thunder.Scalar{Type: "string"},
		ParseArguments: noArguments,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q := thunder.MustParse(`{hello}`, map[string]interface{}{})
		ctx := context.Background()
		e := thunder.Executor{}
		e.Execute(ctx, query, nil, q)
	}
}
