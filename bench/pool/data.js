window.BENCHMARK_DATA = {
  "lastUpdate": 1654111429679,
  "repoUrl": "https://github.com/jamesjarvis/massivelyconcurrentsystems",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "git@jamesjarvis.io",
            "name": "James Jarvis",
            "username": "jamesjarvis"
          },
          "committer": {
            "email": "git@jamesjarvis.io",
            "name": "James Jarvis",
            "username": "jamesjarvis"
          },
          "distinct": true,
          "id": "074caa6db4770acdebde2446ae5a03b814561639",
          "message": "seems I have to specify the go-version to get benchmark output",
          "timestamp": "2022-06-01T20:18:11+01:00",
          "tree_id": "0775769bea9e862b261eb08d904c3f154b11faa6",
          "url": "https://github.com/jamesjarvis/massivelyconcurrentsystems/commit/074caa6db4770acdebde2446ae5a03b814561639"
        },
        "date": 1654111429390,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBatchDispatcher/no_work",
            "value": 15.82,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "758325356 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/lots_of_work",
            "value": 1351356,
            "unit": "ns/op\t   48000 B/op\t    1000 allocs/op",
            "extra": "8772 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/more_work_than_buffer",
            "value": 1088528,
            "unit": "ns/op\t    4800 B/op\t     100 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/more_workers_than_work",
            "value": 1287822,
            "unit": "ns/op\t    4837 B/op\t     100 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/tiny_batch_size",
            "value": 27061,
            "unit": "ns/op\t    4800 B/op\t     100 allocs/op",
            "extra": "453526 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/huge_batch_size",
            "value": 1086451,
            "unit": "ns/op\t    4803 B/op\t     100 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/teeny_interval",
            "value": 27528751,
            "unit": "ns/op\t    4809 B/op\t     100 allocs/op",
            "extra": "421 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/slow_worker",
            "value": 6204767,
            "unit": "ns/op\t    4805 B/op\t     100 allocs/op",
            "extra": "1930 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/slow_worker_and_huge_batch_size",
            "value": 6211649,
            "unit": "ns/op\t    4802 B/op\t     100 allocs/op",
            "extra": "1936 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/default",
            "value": 393146,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "26655 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_buffer",
            "value": 426539,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27043 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/large_buffer",
            "value": 435951,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "30489 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_interval",
            "value": 12868532,
            "unit": "ns/op\t       4 B/op\t       0 allocs/op",
            "extra": "946 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_interval",
            "value": 3414647,
            "unit": "ns/op\t       1 B/op\t       0 allocs/op",
            "extra": "3446 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_batch_size",
            "value": 1090,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "10830558 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_batch_size",
            "value": 419594,
            "unit": "ns/op\t       1 B/op\t       0 allocs/op",
            "extra": "28806 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_number_of_workers",
            "value": 534872,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "22444 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/educated_guess_number_of_workers",
            "value": 535177,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "22419 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_number_of_workers",
            "value": 65874,
            "unit": "ns/op\t      11 B/op\t       0 allocs/op",
            "extra": "152631 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/guesstimate",
            "value": 18226268,
            "unit": "ns/op\t      18 B/op\t       0 allocs/op",
            "extra": "578 times\n2 procs"
          }
        ]
      }
    ]
  }
}