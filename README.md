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
$ go test -v -bench=Master -benchmem
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
$ go test -v -bench=Master -benchmem -benchtime=3s
```

Result: 

```
BenchmarkGoGraphQLMaster-12                30000            134769 ns/op           27711 B/op        500 allocs/op
BenchmarkPlaylyfeGraphQLMaster-12         500000             12034 ns/op            3174 B/op         61 allocs/op
BenchmarkGophersGraphQLMaster-12          300000             12525 ns/op            3877 B/op         38 allocs/op
BenchmarkThunderGraphQLMaster-12          500000              8194 ns/op            2192 B/op         48 allocs/op
```

Testing with http framwork using [Gin](https://github.com/gin-gonic/gin)

```
$ go test -v -bench=Route -benchmem
```

Result:

```
BenchmarkGinHttpRoute-12                  300000              4163 ns/op            1423 B/op         18 allocs/op
BenchmarkGinGoGraphQLRoute-12              10000            155427 ns/op           29145 B/op        515 allocs/op
BenchmarkGinGopherGraphQLRoute-12        1000000              1718 ns/op            1066 B/op          6 allocs/op
BenchmarkGinThunderGraphQLRoute-12        500000              2869 ns/op            1444 B/op         11 allocs/op
```
