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

## Benchmark from wrk benchmarking tool

[wrk](https://github.com/wg/wrk) - a HTTP benchmarking tool

### 

## [gin + graphql-go](golang/graphql-go)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency "http://localhost:8080/graphql"
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    17.92ms   20.76ms 348.48ms   86.34%
    Req/Sec     2.61k   354.24     4.45k    76.25%
  Latency Distribution
     50%   11.21ms
     75%   25.90ms
     90%   45.16ms
     99%   91.97ms
  927003 requests in 30.01s, 131.72MB read
Requests/sec:  30888.47
Transfer/sec:      4.39MB
```

## [gin + graph-gophers](golang/graph-gophers)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency "http://localhost:8080/graphql"
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.15ms    7.77ms 219.62ms   88.70%
    Req/Sec     7.56k     1.35k   44.65k    88.01%
  Latency Distribution
     50%    4.14ms
     75%    7.56ms
     90%   15.01ms
     99%   34.88ms
  2688654 requests in 30.10s, 343.59MB read
Requests/sec:  89330.39
Transfer/sec:     11.42MB
```

## [gin + thunder](golang/thunder)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency "http://localhost:8080/graphql"
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.09ms    3.91ms 207.04ms   79.71%
    Req/Sec     5.55k   475.25     9.01k    95.70%
  Latency Distribution
     50%    5.82ms
     75%    7.76ms
     90%   10.08ms
     99%   15.46ms
  1966319 requests in 30.01s, 358.17MB read
Requests/sec:  65515.21
Transfer/sec:     11.93MB
```

## Sumary

|               | Requests/sec |
|---------------|--------------|
| graphql-go    | 30888.47     |
| graph-gophers | **89330.39** |
| thunder       | 65515.21     |
