# Birdie
Key-value хранилище, рассчитанное главным образом на операцию BulkStore, когда данные сохраняются большими пачками.

В проекте имеются бэнчмарки, для запуска можно использовать `make bench`:
```shell
cpu: Intel(R) Core(TM) i5-5350U CPU @ 1.80GHz
BenchmarkHash-4                              5703685             206.3 ns/op              31 B/op          2 allocs/op
BenchmarkStore-4                             1000000              1246 ns/op             283 B/op          7 allocs/op
BenchmarkLoad-4                              2292718             516.4 ns/op              31 B/op          2 allocs/op
BenchmarkBulkStore/count_1000-4                 3322            986949 ns/op          122849 B/op       1038 allocs/op
BenchmarkBulkStore/count_10000-4                 268          11222237 ns/op         1025112 B/op      10386 allocs/op
BenchmarkBulkStore/count_100000-4                 38          38469056 ns/op         1307768 B/op     101995 allocs/op
BenchmarkBulkStore_Parallel/count_1000-4        4150            343325 ns/op            8001 B/op       1000 allocs/op
BenchmarkBulkStore_Parallel/count_10000-4        342           3637982 ns/op           80002 B/op      10000 allocs/op
BenchmarkBulkStore_Parallel/count_100000-4        38          37536009 ns/op          800013 B/op     100000 allocs/op
```
Пачка 100К записей сохраняется в хранилище в среднем за 35ms, но выгоднее всего по скорости выполнения вставлять пакеты до 10К записей.

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
