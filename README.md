# golang-graphql-benchmark

benchmark of graphql framework in golang.

* [graphql-go/graphql](https://github.com/graphql-go/graphql)
* [playlyfe/go-graphql](https://github.com/playlyfe/go-graphql)
* [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)
* [samsarahq/thunder](https://github.com/samsarahq/thunder)

## Environment

```
OS: Ubuntu 14.04
Memory: 24G
CPU: Intel(R) Xeon(R) CPU X5670  @ 2.93GHz
```

## Benchmark Result

Please execute the following command.

```
$ git clone https://github.com/appleboy/golang-graphql-benchmark.git
$ cd golang-graphql-benchmark
$ go test -v -bench=. -benchmem
```

Result:

```
BenchmarkGoGraphQLMaster-12                10000            137017 ns/op           27712 B/op        500 allocs/op
BenchmarkPlaylyfeGraphQLMaster-12         200000             11353 ns/op            3174 B/op         61 allocs/op
BenchmarkGophersGraphQLMaster-12          100000             12459 ns/op            3877 B/op         38 allocs/op
BenchmarkThunderGraphQLMaster-12          200000              8149 ns/op            2192 B/op         48 allocs/op
```

set `benchtime` to `3s`

```
$ go test -v -bench=. -benchmem -benchtime=3s
```

Result: 

```
BenchmarkGoGraphQLMaster-12                30000            134769 ns/op           27711 B/op        500 allocs/op
BenchmarkPlaylyfeGraphQLMaster-12         500000             12034 ns/op            3174 B/op         61 allocs/op
BenchmarkGophersGraphQLMaster-12          300000             12525 ns/op            3877 B/op         38 allocs/op
BenchmarkThunderGraphQLMaster-12          500000              8194 ns/op            2192 B/op         48 allocs/op
```
