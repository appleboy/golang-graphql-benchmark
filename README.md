# golang-graphql-benchmark

[![Build Status](https://cloud.drone.io/api/badges/appleboy/golang-graphql-benchmark/status.svg)](https://cloud.drone.io/appleboy/golang-graphql-benchmark)

benchmark of graphql framework in [go1.14](https://golang.org/doc/go1.14) version.

* [graphql-go/graphql](https://github.com/graphql-go/graphql) version: v0.7.9
* [playlyfe/go-graphql](https://github.com/playlyfe/go-graphql) version: v0.0.0-20191219091308-23c3f22218ef
* [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go) version: v0.0.0-20200207002730-8334863f2c8b
* [samsarahq/thunder](https://github.com/samsarahq/thunder) version: v0.5.0
* [99designs/gqlgen](https://github.com/99designs/gqlgen) version: v0.11.3

## Benchmark Result (2020/03/15)

Please execute the following command.

```sh
git clone https://github.com/appleboy/golang-graphql-benchmark.git
cd golang-graphql-benchmark
go test -v -bench=Master -benchmem
```

Result:

```sh
BenchmarkGoGraphQLMaster-8                 15282             83858 ns/op           27490 B/op        458 allocs/op
BenchmarkPlaylyfeGraphQLMaster-8          161757              6852 ns/op            2854 B/op         56 allocs/op
BenchmarkGophersGraphQLMaster-8           191847              6256 ns/op            3673 B/op         38 allocs/op
BenchmarkThunderGraphQLMaster-8           426513              2734 ns/op            1336 B/op         30 allocs/op
```

Testing with http framwork using [Gin](https://github.com/gin-gonic/gin)

```sh
go test -v -bench=Route -benchmem
```

Result:

```sh
BenchmarkGinHttpRoute-8                   400288              2743 ns/op            1324 B/op         19 allocs/op
BenchmarkGinGQLGenRoute-8                 402793              3091 ns/op             944 B/op         13 allocs/op
BenchmarkGinGoGraphQLRoute-8               32685             36864 ns/op           17385 B/op        226 allocs/op
BenchmarkGinGopherGraphQLRoute-8         1339144               894 ns/op             973 B/op          6 allocs/op
BenchmarkGinThunderGraphQLRoute-8         660644              1725 ns/op            1179 B/op         11 allocs/op
```

## Benchmark from wrk benchmarking tool

[wrk](https://github.com/wg/wrk) - a HTTP benchmarking tool

### [gin + graphql-go](golang/graphql-go)

```sh
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    78.26ms  114.59ms   1.31s    86.37%
    Req/Sec     1.59k   537.11     4.15k    68.81%
  Latency Distribution
     50%    7.68ms
     75%  128.89ms
     90%  221.97ms
     99%  499.33ms
  571048 requests in 30.05s, 81.14MB read
  Socket errors: connect 0, read 233, write 0, timeout 0
Requests/sec:  19004.92
Transfer/sec:      2.70MB
```

### [gin + graph-gophers](golang/graph-gophers)

```sh
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.75ms   10.75ms 214.71ms   88.06%
    Req/Sec     3.71k   583.50     6.30k    72.11%
  Latency Distribution
     50%    7.21ms
     75%   12.15ms
     90%   23.80ms
     99%   53.21ms
  1331529 requests in 30.05s, 170.16MB read
  Socket errors: connect 0, read 236, write 0, timeout 0
Requests/sec:  44308.44
Transfer/sec:      5.66MB
```

### [gin + thunder](golang/thunder)

```sh
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    12.34ms   13.20ms 181.22ms   85.99%
    Req/Sec     3.43k   540.37     5.90k    72.64%
  Latency Distribution
     50%    8.09ms
     75%   17.52ms
     90%   30.13ms
     99%   58.72ms
  1232462 requests in 30.06s, 224.50MB read
  Socket errors: connect 0, read 232, write 0, timeout 0
Requests/sec:  40994.33
Transfer/sec:      7.47MB
```

### [gin + gqlgen](golang/gqlgen)

```sh
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.77ms    2.61ms  66.76ms   82.63%
    Req/Sec     4.18k   474.51     9.97k    83.78%
  Latency Distribution
     50%    7.48ms
     75%    8.47ms
     90%    9.99ms
     99%   18.09ms
  1501098 requests in 30.07s, 191.83MB read
  Socket errors: connect 0, read 239, write 0, timeout 0
Requests/sec:  49925.73
Transfer/sec:      6.38MB
```

### [gin + josn](golang/gin-json)

```sh
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/json.lua --latency http://localhost:8080/graphql
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.05ms    1.94ms  40.57ms   81.71%
    Req/Sec     4.64k   415.61     7.95k    85.89%
  Latency Distribution
     50%    6.89ms
     75%    7.66ms
     90%    8.79ms
     99%   14.50ms
  1663093 requests in 30.06s, 356.86MB read
  Socket errors: connect 0, read 236, write 0, timeout 0
  Non-2xx or 3xx responses: 1663093
Requests/sec:  55334.07
Transfer/sec:     11.87MB
```

## Summary

|                   | Requests/sec |
| ----------------- | ------------ |
| graphql-go        | 19004.92     |
| graph-gophers     | 44308.44     |
| thunder           | 40994.33     |
| gqlgen            | **49925.73** |

Without graphql (only gin render json output)

|                      | Requests/sec  |
| -------------------- | ------------- |
| json without graphql | **55334.07** |
