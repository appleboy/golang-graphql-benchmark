
# Tesing

how to test graphql using curl

```
$ curl -X POST -H "Content-Type: application/json" \
  --data '{"query":"{ hello }", "operationName":"", "variables": null}' \
  http://localhost:8080/graphql
```

## gin + graphql-go

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency "http://localhost:8080/graphql"
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   100.02ms  136.05ms   1.53s    86.88%
    Req/Sec   602.68    227.40     2.38k    71.36%
  Latency Distribution
     50%   37.90ms
     75%  141.16ms
     90%  276.93ms
     99%  623.23ms
  215793 requests in 30.10s, 30.66MB read
  Socket errors: connect 0, read 177, write 0, timeout 0
Requests/sec:   7169.26
Transfer/sec:      1.02MB
```

## gin + graph-gophers

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency "http://localhost:8080/graphql"
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    24.28ms   25.76ms 394.02ms   93.54%
    Req/Sec     1.60k   525.17    10.90k    77.93%
  Latency Distribution
     50%   19.27ms
     75%   26.00ms
     90%   40.26ms
     99%  138.60ms
  570824 requests in 30.11s, 72.95MB read
  Socket errors: connect 0, read 6, write 3, timeout 0
Requests/sec:  18959.51
Transfer/sec:      2.42MB
```

## gin + thunder

```
$ wrk -t12 -c400 -d30s --timeout 10s --script=golang/post.lua --latency "http://localhost:8080/graphql"
Running 30s test @ http://localhost:8080/graphql
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    68.56ms  102.66ms   1.15s    87.76%
    Req/Sec     1.05k   393.60     3.68k    74.09%
  Latency Distribution
     50%   22.35ms
     75%   91.03ms
     90%  194.78ms
     99%  482.67ms
  375764 requests in 30.10s, 68.45MB read
  Socket errors: connect 0, read 2, write 0, timeout 0
Requests/sec:  12482.84
Transfer/sec:      2.27MB
```
