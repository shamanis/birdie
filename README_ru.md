![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)
![Golang version](https://img.shields.io/badge/Golang-1.19-green.svg)
![Go Test](https://github.com/shamanis/birdie/actions/workflows/go.yml/badge.svg?branch=master)
![Docker Build and Push](https://github.com/shamanis/birdie/actions/workflows/docker-image.yml/badge.svg)

# Birdie
Key-value хранилище, рассчитанное главным образом на операцию BulkStore, когда данные сохраняются большими пачками.

В проекте имеются бэнчмарки, для запуска можно использовать `make bench`:
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
Пачка 100К записей сохраняется в хранилище в среднем за 30ms.

Тут следует учитывать результаты параллельных бенчмарков, т.к. они больше похоже на условия реальной работы множества параллельных клиентов.

## Environment Variables
* `LISTEN_PORT` - порт, на котором будет запущен сервер. По-умолчанию `:50051`
* `LOG_LEVEL` - уровень логирования. По-молчанию `info`.

## Deploy
* Docker - смотри пример [docker-compose.yml](deployments/docker-compose.yml)
* Systemd service - смотри пример [birdie.service](init/birdie.service)

## Examples
* Пример реализации клиента [examples/client/main.go](examples/client/main.go)

## .proto
Для реализации своего клиента используйте всегда актуальный [birdie.proto](api/birdie.proto) конекретной версии сервера: 
