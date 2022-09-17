![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)
![Golang version](https://img.shields.io/badge/Golang-1.19-green.svg)
![Go Test](https://github.com/shamanis/birdie/actions/workflows/go.yml/badge.svg?branch=master)
![Docker Build and Push](https://github.com/shamanis/birdie/actions/workflows/docker-image.yml/badge.svg)

# Birdie
Key-value storage designed mainly for the BulkStore operation, when data is stored in large batches.

The project has benchmarks, you can use `make bench` to run it:
```shell
cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
BenchmarkHash-4                           3344268             357.7 ns/op              39 B/op          2 allocs/op
BenchmarkStore-4                          1000000              1425 ns/op             253 B/op          7 allocs/op
BenchmarkLoad-4                           1779650             777.4 ns/op              39 B/op          2 allocs/op
BenchmarkBulkStore/count_1000-4              2365            849593 ns/op          120534 B/op       1029 allocs/op
BenchmarkBulkStore/count_10000-4              241           5123830 ns/op          240012 B/op      10384 allocs/op
BenchmarkBulkStore/count_100000-4              22         235653234 ns/op        23109100 B/op     102137 allocs/op
BenchmarkBulkStore_Parallel/count_1000-4     6307            480595 ns/op           16000 B/op       1000 allocs/op
BenchmarkBulkStore_Parallel/count_10000-4     600           4034845 ns/op          160002 B/op      10000 allocs/op
BenchmarkBulkStore_Parallel/count_100000-4     60          26912932 ns/op         1600023 B/op     100000 allocs/op
```
A pack of 100K records is stored in storage in an average of 30ms.

Here you should pay attention to the results of parallel benchmarks, because they imitate the real work of clients in conditions of parallel connections.

## Environment Variables
* `LISTEN_NETWROK` - the network type of listener. May be `tcp` or `unix`. Default `tcp`.
* `LISTEN_ADDR` - the socket address on which the server will run. Default `0.0.0.0:50051`
* `LOG_LEVEL` - logging level. Default `info`.

## Deploy
* Docker - see [docker-compose.yml](deployments/docker-compose.yml)
* Systemd service - see [birdie.service](init/birdie.service)

## Examples
* Client implementation example [examples/client/main.go](examples/client/main.go)

## .proto
For implement your client use actual [birdie.proto](api/birdie.proto) from server version. 
