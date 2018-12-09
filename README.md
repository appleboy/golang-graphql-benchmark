# golang-graphql-benchmark

[![Build Status](https://cloud.drone.io/api/badges/appleboy/golang-graphql-benchmark/status.svg)](https://cloud.drone.io/appleboy/golang-graphql-benchmark)

benchmark of graphql framework in golang.

* [graphql-go/graphql](https://github.com/graphql-go/graphql) version: 2018-09-12T00:08:44Z
* [playlyfe/go-graphql](https://github.com/playlyfe/go-graphql) version: 2017-04-28T20:40:03Z
* [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go) version: 2018-06-09T14:05:35Z
* [samsarahq/thunder](https://github.com/samsarahq/thunder) version: 2018-08-21T22:33:29Z

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
BenchmarkGoGraphQLMaster-12                10000            146557 ns/op           28138 B/op        506 allocs/op
BenchmarkPlaylyfeGraphQLMaster-12         200000             11761 ns/op            3175 B/op         61 allocs/op
BenchmarkGophersGraphQLMaster-12          100000             12985 ns/op            3877 B/op         38 allocs/op
BenchmarkThunderGraphQLMaster-12          200000              8722 ns/op            2176 B/op         48 allocs/op
```

set `benchtime` to `3s`

```
$ go test -v -bench=Master -benchmem -benchtime=3s
```

Result: 

```
BenchmarkGoGraphQLMaster-12                30000            146054 ns/op           28136 B/op        506 allocs/op
BenchmarkPlaylyfeGraphQLMaster-12         500000             12062 ns/op            3175 B/op         61 allocs/op
BenchmarkGophersGraphQLMaster-12          300000             12955 ns/op            3877 B/op         38 allocs/op
BenchmarkThunderGraphQLMaster-12          500000              8756 ns/op            2176 B/op         48 allocs/op
```

Testing with http framwork using [Gin](https://github.com/gin-gonic/gin)

```
$ go test -v -bench=Route -benchmem
```

Result:

```
BenchmarkGinHttpRoute-12                  300000              3998 ns/op            1423 B/op         18 allocs/op
BenchmarkGinGoGraphQLRoute-12              20000             66478 ns/op           17599 B/op        226 allocs/op
BenchmarkGinGopherGraphQLRoute-12        1000000              1740 ns/op            1066 B/op          6 allocs/op
BenchmarkGinThunderGraphQLRoute-12        500000              2847 ns/op            1444 B/op         11 allocs/op
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
    Latency    16.22ms   16.63ms 236.12ms   86.17%
    Req/Sec     2.60k   233.08     4.02k    73.14%
  Latency Distribution
     50%   10.90ms
     75%   22.76ms
     90%   38.29ms
     99%   75.06ms
  933231 requests in 30.03s, 132.61MB read
Requests/sec:  31071.70
Transfer/sec:      4.42MB
```

## [gin + graph-gophers](golang/graph-gophers)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.39ms    3.61ms 206.07ms   88.34%
    Req/Sec     8.14k     1.18k   29.35k    91.89%
  Latency Distribution
     50%    3.61ms
     75%    4.93ms
     90%    7.57ms
     99%   18.61ms
  2918001 requests in 30.10s, 372.90MB read
Requests/sec:  96944.53
Transfer/sec:     12.39MB
```

## [gin + thunder](golang/thunder)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency "http://localhost:8080/graphql"
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.76ms    4.00ms 209.77ms   87.50%
    Req/Sec     5.76k   510.83    12.64k    91.89%
  Latency Distribution
     50%    5.48ms
     75%    7.21ms
     90%    9.20ms
     99%   14.06ms
  2066881 requests in 30.09s, 376.49MB read
Requests/sec:  68685.50
Transfer/sec:     12.51MB
```

## [gin + josn](golang/gin-json)

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/json.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.40ms    3.07ms  43.04ms   89.83%
    Req/Sec    11.31k   671.31    19.84k    72.50%
  Latency Distribution
     50%    2.58ms
     75%    3.78ms
     90%    6.45ms
     99%   16.51ms
  4057155 requests in 30.06s, 576.51MB read
Requests/sec: 134982.90
Transfer/sec:     19.18MB
```

## Sumary

|                   | Requests/sec |
| ----------------- | ------------ |
| graphql-go        | 31071.70     |
| **graph-gophers** | **96944.53** |
| thunder           | 68685.50     |

Without graphql (only gin render json output)

|                      | Requests/sec  |
| -------------------- | ------------- |
| json without graphql | **134982.90** |
