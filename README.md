# golang-graphql-benchmark

benchmark of graphql framework in golang.

* [graphql-go/graphql](github.com/graphql-go/graphql)
* [playlyfe/go-graphql](github.com/playlyfe/go-graphql)

## Benchmark Result

Please execute the following command.

```
$ go test -v -bench .
```

Result:

```
BenchmarkGoGraphQLMaster-4                 20000             81387 ns/op
BenchmarkPlaylyfeGraphQLMaster-4          200000              7791 ns/op
```
