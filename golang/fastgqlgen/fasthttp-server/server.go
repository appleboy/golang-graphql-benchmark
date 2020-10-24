package main

import (
	"log"

	fastgqlgn "golang-graphql-benchmark/fastgqlgn"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/valyala/fasthttp"
)

func main() {
	srv := handler.NewDefaultServer(fastgqlgn.NewExecutableSchema(fastgqlgn.Config{Resolvers: &fastgqlgn.Resolver{}}))
	gqlHandler := srv.Handler()

	h := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/graphql":
			gqlHandler(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	log.Fatal(fasthttp.ListenAndServe(":8080", h))
}
