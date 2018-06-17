# golang-graphql-benchmark

benchmark of graphql framework in golang.

* [graphql-go/graphql](github.com/graphql-go/graphql)
* [playlyfe/go-graphql](github.com/playlyfe/go-graphql)
* [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)

## Benchmark Result

Please execute the following command.

```
$ go test -v -bench .
```

Result:

```
BenchmarkGoGraphQLMaster-4                 20000            129572 ns/op
BenchmarkPlaylyfeGraphQLMaster-4          200000              9141 ns/op
BenchmarkGophersGraphQLMaster-4           200000              5263 ns/op
```
