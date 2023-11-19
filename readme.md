# Running

```shell
docker-compose build

docker-compose up {dotnet|fastify}
```

# Checking stats

You can check in real time how much memory and cpu is consumed by the containers

```shell
docker stats $( docker ps --format '{{ .Names }}' )
```

# Benchmarking

```shell
autocannon -c 100 -d 40 -p 10 localhost:5105
```

# Results

## Dotnet

autocannon -c 100 -d 40 -p 10 localhost:5105
Running 40s test @ http://localhost:5105
100 connections with 10 pipelining factor

```
┌─────────┬──────┬───────┬───────┬───────┬──────────┬──────────┬─────────┐
│ Stat    │ 2.5% │ 50%   │ 97.5% │ 99%   │ Avg      │ Stdev    │ Max     │
├─────────┼──────┼───────┼───────┼───────┼──────────┼──────────┼─────────┤
│ Latency │ 7 ms │ 16 ms │ 34 ms │ 42 ms │ 17.95 ms │ 23.05 ms │ 1024 ms │
└─────────┴──────┴───────┴───────┴───────┴──────────┴──────────┴─────────┘
┌───────────┬─────────┬─────────┬─────────┬─────────┬──────────┬──────────┬─────────┐
│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg      │ Stdev    │ Min     │
├───────────┼─────────┼─────────┼─────────┼─────────┼──────────┼──────────┼─────────┤
│ Req/Sec   │ 6035    │ 6035    │ 55807   │ 70399   │ 54287.45 │ 10089.65 │ 6035    │
├───────────┼─────────┼─────────┼─────────┼─────────┼──────────┼──────────┼─────────┤
│ Bytes/Sec │ 1.06 MB │ 1.06 MB │ 9.82 MB │ 12.4 MB │ 9.55 MB  │ 1.78 MB  │ 1.06 MB │
└───────────┴─────────┴─────────┴─────────┴─────────┴──────────┴──────────┴─────────┘
```

## Fastify

```
┌─────────┬───────┬───────┬────────┬────────┬──────────┬──────────┬─────────┐
│ Stat    │ 2.5%  │ 50%   │ 97.5%  │ 99%    │ Avg      │ Stdev    │ Max     │
├─────────┼───────┼───────┼────────┼────────┼──────────┼──────────┼─────────┤
│ Latency │ 38 ms │ 51 ms │ 101 ms │ 111 ms │ 57.06 ms │ 56.81 ms │ 2508 ms │
└─────────┴───────┴───────┴────────┴────────┴──────────┴──────────┴─────────┘
┌───────────┬─────────┬─────────┬─────────┬─────────┬──────────┬─────────┬─────────┐
│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg      │ Stdev   │ Min     │
├───────────┼─────────┼─────────┼─────────┼─────────┼──────────┼─────────┼─────────┤
│ Req/Sec   │ 8187    │ 8187    │ 18031   │ 20895   │ 17366.45 │ 2617.06 │ 8185    │
├───────────┼─────────┼─────────┼─────────┼─────────┼──────────┼─────────┼─────────┤
│ Bytes/Sec │ 1.54 MB │ 1.54 MB │ 3.39 MB │ 3.93 MB │ 3.26 MB  │ 492 kB  │ 1.54 MB │
└───────────┴─────────┴─────────┴─────────┴─────────┴──────────┴─────────┴─────────┘
```

## Express

```
┌─────────┬────────┬────────┬────────┬────────┬───────────┬───────────┬─────────┐
│ Stat    │ 2.5%   │ 50%    │ 97.5%  │ 99%    │ Avg       │ Stdev     │ Max     │
├─────────┼────────┼────────┼────────┼────────┼───────────┼───────────┼─────────┤
│ Latency │ 128 ms │ 196 ms │ 270 ms │ 284 ms │ 209.11 ms │ 250.57 ms │ 8475 ms │
└─────────┴────────┴────────┴────────┴────────┴───────────┴───────────┴─────────┘
┌───────────┬────────┬────────┬────────┬─────────┬─────────┬────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%   │ Avg     │ Stdev  │ Min    │
├───────────┼────────┼────────┼────────┼─────────┼─────────┼────────┼────────┤
│ Req/Sec   │ 2891   │ 2891   │ 4759   │ 5523    │ 4760.53 │ 520.46 │ 2891   │
├───────────┼────────┼────────┼────────┼─────────┼─────────┼────────┼────────┤
│ Bytes/Sec │ 729 kB │ 729 kB │ 1.2 MB │ 1.39 MB │ 1.2 MB  │ 131 kB │ 729 kB │
└───────────┴────────┴────────┴────────┴─────────┴─────────┴────────┴────────┘
```
