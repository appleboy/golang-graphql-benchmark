# golang-graphql-benchmark

[![Build Status](https://cloud.drone.io/api/badges/appleboy/golang-graphql-benchmark/status.svg)](https://cloud.drone.io/appleboy/golang-graphql-benchmark)

benchmark of graphql framework in golang.

* [graphql-go/graphql](https://github.com/graphql-go/graphql) version: 2018-12-03T01:16:34Z
* [playlyfe/go-graphql](https://github.com/playlyfe/go-graphql) version: 2017-04-28T20:40:03Z
* [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go) version: 2018-11-28T22:09:52Z
* [samsarahq/thunder](https://github.com/samsarahq/thunder) version: 2018-12-05T21:49:39Z
* [99designs/gqlgen](https://github.com/99designs/gqlgen) version: 2018-12-02T22:03:39Z

## Environment

```
OS: Ubuntu 14.04
Memory: 24G
CPU: Intel(R) Xeon(R) CPU X5670  @ 2.93GHz
```

## Benchmark Result (2018/12/10)

Please execute the following command.

```
$ git clone https://github.com/appleboy/golang-graphql-benchmark.git
$ cd golang-graphql-benchmark
$ go test -v -bench=Master -benchmem
```

Result:

```
BenchmarkGoGraphQLMaster-12                10000            126120 ns/op           27216 B/op        452 allocs/op
BenchmarkPlaylyfeGraphQLMaster-12         200000             10094 ns/op            2919 B/op         57 allocs/op
BenchmarkGophersGraphQLMaster-12          100000             11242 ns/op            3762 B/op         39 allocs/op
BenchmarkThunderGraphQLMaster-12          200000              6985 ns/op            1968 B/op         42 allocs/op
```

set `benchtime` to `3s`

```
$ go test -v -bench=Master -benchmem -benchtime=3s
```

Result: 

```
BenchmarkGoGraphQLMaster-12                30000            124961 ns/op           27215 B/op        452 allocs/op
BenchmarkPlaylyfeGraphQLMaster-12         500000             10013 ns/op            2918 B/op         57 allocs/op
BenchmarkGophersGraphQLMaster-12          500000             11166 ns/op            3762 B/op         39 allocs/op
BenchmarkThunderGraphQLMaster-12          500000              6726 ns/op            1968 B/op         42 allocs/op
```

Testing with http framwork using [Gin](https://github.com/gin-gonic/gin)

```
$ go test -v -bench=Route -benchmem
```

Result:

```
BenchmarkGinHttpRoute-12                  300000              3719 ns/op            1279 B/op         18 allocs/op
BenchmarkGinGQLGenRoute-12                500000              3816 ns/op            1650 B/op         12 allocs/op
BenchmarkGinGoGraphQLRoute-12              30000             57065 ns/op           17195 B/op        221 allocs/op
BenchmarkGinGopherGraphQLRoute-12        1000000              1680 ns/op             969 B/op          6 allocs/op
BenchmarkGinThunderGraphQLRoute-12        500000              2854 ns/op            1204 B/op         11 allocs/op
```

## Benchmark from wrk benchmarking tool

[wrk](https://github.com/wg/wrk) - a HTTP benchmarking tool

### 

## [gin + graphql-go](golang/graphql-go)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    18.18ms   22.03ms 290.57ms   85.91%
    Req/Sec     2.81k   403.63     9.79k    78.24%
  Latency Distribution
     50%   10.80ms
     75%   26.46ms
     90%   48.24ms
     99%   96.89ms
  1006733 requests in 30.10s, 143.05MB read
Requests/sec:  33448.20
Transfer/sec:      4.75MB
```

## [gin + graph-gophers](golang/graph-gophers)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.35ms    9.92ms 228.60ms   91.21%
    Req/Sec     6.11k     2.12k   12.27k    72.83%
  Latency Distribution
     50%    5.06ms
     75%    9.00ms
     90%   16.11ms
     99%   45.78ms
  2179756 requests in 30.06s, 278.56MB read
Requests/sec:  72506.58
Transfer/sec:      9.27MB
```

## [gin + thunder](golang/thunder)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.10ms   16.33ms 258.10ms   87.78%
    Req/Sec     6.04k     1.04k   11.57k    81.34%
  Latency Distribution
     50%    3.40ms
     75%   15.15ms
     90%   31.05ms
     99%   74.53ms
  2148076 requests in 30.02s, 391.28MB read
Requests/sec:  71551.64
Transfer/sec:     13.03MB
```

## [gin + gqlgen](golang/gqlgen)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.61ms    8.24ms 257.19ms   92.70%
    Req/Sec     8.39k     1.31k   18.66k    82.80%
  Latency Distribution
     50%    4.42ms
     75%    8.93ms
     90%   12.21ms
     99%   29.50ms
  2991677 requests in 30.02s, 382.31MB read
Requests/sec:  99645.21
Transfer/sec:     12.73MB
```

## [gin + josn](golang/gin-json)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/json.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.19ms    4.81ms 204.15ms   88.81%
    Req/Sec    10.47k     1.93k   99.89k    89.61%
  Latency Distribution
     50%    3.17ms
     75%    5.07ms
     90%    9.69ms
     99%   23.51ms
  3752158 requests in 30.10s, 533.17MB read
Requests/sec: 124663.94
Transfer/sec:     17.71MB
```

## Sumary

|                   | Requests/sec |
| ----------------- | ------------ |
| graphql-go        | 33448.20     |
| graph-gophers     | 72506.58     |
| thunder           | 71551.64     |
| gqlgen            | **99645.21** |

Without graphql (only gin render json output)

|                      | Requests/sec  |
| -------------------- | ------------- |
| json without graphql | **124663.94** |
