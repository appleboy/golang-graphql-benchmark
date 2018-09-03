
## gin + graphql-go

```
$ wrk -t12 -c400 -d30s --timeout 10s "http://localhost:8080/graphql?query={hello}"
Running 30s test @ http://localhost:8080/graphql?query={hello}
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    91.23ms  127.96ms   1.49s    87.74%
    Req/Sec   680.67    240.06     1.83k    70.99%
  244055 requests in 30.11s, 37.01MB read
  Socket errors: connect 0, read 6, write 0, timeout 0
Requests/sec:   8106.74
Transfer/sec:      1.23MB
```
