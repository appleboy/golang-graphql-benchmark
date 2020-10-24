package main

import (
	"flag"
	fastgqlgn "golang-graphql-benchmark/fastgqlgn"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	prefork := flag.Bool("prefork", false, "enable prefork")
	flag.Parse()

	app := fiber.New(fiber.Config{
		Prefork: *prefork,
	})

	srv := handler.NewDefaultServer(fastgqlgn.NewExecutableSchema(fastgqlgn.Config{Resolvers: &fastgqlgn.Resolver{}}))
	gqlHandler := srv.Handler()

	app.All("/graphql", func(c *fiber.Ctx) error {
		gqlHandler(c.Context())
		return nil
	})

	log.Fatal(app.Listen(":8080"))
}
