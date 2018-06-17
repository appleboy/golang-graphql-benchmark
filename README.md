# golang-graphql-benchmark

benchmark of graphql framework in golang.

* [graphql-go/graphql](https://github.com/graphql-go/graphql)
* [playlyfe/go-graphql](https://github.com/playlyfe/go-graphql)
* [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)

## Benchmark Result

Please execute the following command.

```
$ git clone https://github.com/appleboy/golang-graphql-benchmark.git
$ cd golang-graphql-benchmark
$ go test -v -bench=. -benchmem
```

Result:

```
BenchmarkGoGraphQLMaster-4                 20000             84131 ns/op           27254 B/op        489 allocs/op
BenchmarkPlaylyfeGraphQLMaster-4          200000              7531 ns/op            2919 B/op         59 allocs/op
BenchmarkGophersGraphQLMaster-4           200000              5041 ns/op            3909 B/op         39 allocs/op
```
