window.BENCHMARK_DATA = {
  "lastUpdate": 1656372496531,
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
      },
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
          "id": "b03e635d8668211e0b23a86e084b1e5aa00c6a17",
          "message": "Update README with badges and links and stuff",
          "timestamp": "2022-06-01T20:31:31+01:00",
          "tree_id": "017cd2f7aeaedbc226a984520d19e79dae5ddae3",
          "url": "https://github.com/jamesjarvis/massivelyconcurrentsystems/commit/b03e635d8668211e0b23a86e084b1e5aa00c6a17"
        },
        "date": 1654112247669,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBatchDispatcherSingleItem/default",
            "value": 465592,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25249 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_buffer",
            "value": 454491,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "26218 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/large_buffer",
            "value": 462881,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "25185 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_interval",
            "value": 16962355,
            "unit": "ns/op\t       7 B/op\t       0 allocs/op",
            "extra": "744 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_interval",
            "value": 3444120,
            "unit": "ns/op\t       1 B/op\t       0 allocs/op",
            "extra": "3454 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_batch_size",
            "value": 1152,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "10752702 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_batch_size",
            "value": 466852,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27255 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_number_of_workers",
            "value": 561756,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "21339 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/educated_guess_number_of_workers",
            "value": 562949,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "21298 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_number_of_workers",
            "value": 82642,
            "unit": "ns/op\t      13 B/op\t       0 allocs/op",
            "extra": "134071 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/guesstimate",
            "value": 23837566,
            "unit": "ns/op\t      20 B/op\t       0 allocs/op",
            "extra": "518 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/no_work",
            "value": 16.24,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "733414477 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/lots_of_work",
            "value": 1573871,
            "unit": "ns/op\t   48000 B/op\t    1000 allocs/op",
            "extra": "7683 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/more_work_than_buffer",
            "value": 1158663,
            "unit": "ns/op\t    4800 B/op\t     100 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/more_workers_than_work",
            "value": 1388898,
            "unit": "ns/op\t    4839 B/op\t     100 allocs/op",
            "extra": "8578 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/tiny_batch_size",
            "value": 37538,
            "unit": "ns/op\t    4800 B/op\t     100 allocs/op",
            "extra": "316791 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/huge_batch_size",
            "value": 1147701,
            "unit": "ns/op\t    4803 B/op\t     100 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/teeny_interval",
            "value": 35391444,
            "unit": "ns/op\t    4805 B/op\t     100 allocs/op",
            "extra": "375 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/slow_worker",
            "value": 6445124,
            "unit": "ns/op\t    4802 B/op\t     100 allocs/op",
            "extra": "1892 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/slow_worker_and_huge_batch_size",
            "value": 6430071,
            "unit": "ns/op\t    4801 B/op\t     100 allocs/op",
            "extra": "1886 times\n2 procs"
          }
        ]
      },
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
          "id": "2391354a9f60bfb42c19e10267a5c7d2e6c567ec",
          "message": "add some more badges because badges are cool",
          "timestamp": "2022-06-01T20:37:16+01:00",
          "tree_id": "54fec6aa67549a22ae0a4b174c4a93ff1dcc7bba",
          "url": "https://github.com/jamesjarvis/massivelyconcurrentsystems/commit/2391354a9f60bfb42c19e10267a5c7d2e6c567ec"
        },
        "date": 1654112647892,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBatchDispatcher/no_work",
            "value": 15.3,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "781878108 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/lots_of_work",
            "value": 1544224,
            "unit": "ns/op\t   48000 B/op\t    1000 allocs/op",
            "extra": "7400 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/more_work_than_buffer",
            "value": 1119064,
            "unit": "ns/op\t    4800 B/op\t     100 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/more_workers_than_work",
            "value": 1335189,
            "unit": "ns/op\t    4839 B/op\t     100 allocs/op",
            "extra": "8965 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/tiny_batch_size",
            "value": 33352,
            "unit": "ns/op\t    4800 B/op\t     100 allocs/op",
            "extra": "350844 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/huge_batch_size",
            "value": 1118914,
            "unit": "ns/op\t    4803 B/op\t     100 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/teeny_interval",
            "value": 37530332,
            "unit": "ns/op\t    4803 B/op\t     100 allocs/op",
            "extra": "1134 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/slow_worker",
            "value": 6432232,
            "unit": "ns/op\t    4802 B/op\t     100 allocs/op",
            "extra": "1898 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/slow_worker_and_huge_batch_size",
            "value": 6479476,
            "unit": "ns/op\t    4805 B/op\t     100 allocs/op",
            "extra": "1875 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/default",
            "value": 521816,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "23475 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_buffer",
            "value": 525793,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "22812 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/large_buffer",
            "value": 519005,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "23078 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_interval",
            "value": 15714907,
            "unit": "ns/op\t       6 B/op\t       0 allocs/op",
            "extra": "709 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_interval",
            "value": 4272265,
            "unit": "ns/op\t       1 B/op\t       0 allocs/op",
            "extra": "3139 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_batch_size",
            "value": 1270,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "9771079 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_batch_size",
            "value": 526270,
            "unit": "ns/op\t       1 B/op\t       0 allocs/op",
            "extra": "22855 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_number_of_workers",
            "value": 543877,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "22066 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/educated_guess_number_of_workers",
            "value": 543920,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "22076 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_number_of_workers",
            "value": 58424,
            "unit": "ns/op\t       6 B/op\t       0 allocs/op",
            "extra": "259633 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/guesstimate",
            "value": 24550306,
            "unit": "ns/op\t      17 B/op\t       0 allocs/op",
            "extra": "428 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "git@jamesjarvis.io",
            "name": "James Jarvis",
            "username": "jamesjarvis"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "382f373ca046940c2a8b284a8025d5709a810013",
          "message": "Merge pull request #3 from jamesjarvis/simplify-worker\n\nI think simplify the worker?",
          "timestamp": "2022-06-28T00:22:40+01:00",
          "tree_id": "cbbfe7b3440e4d575171d6a755811897870984e4",
          "url": "https://github.com/jamesjarvis/massivelyconcurrentsystems/commit/382f373ca046940c2a8b284a8025d5709a810013"
        },
        "date": 1656372496227,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkBatchDispatcherSingleItem/default",
            "value": 416853,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28561 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_buffer",
            "value": 407912,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "29198 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/large_buffer",
            "value": 412479,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28820 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_interval",
            "value": 1194,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "10551171 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_interval",
            "value": 3783270,
            "unit": "ns/op\t       1 B/op\t       0 allocs/op",
            "extra": "3057 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_batch_size",
            "value": 1164,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "10134744 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_batch_size",
            "value": 413902,
            "unit": "ns/op\t       1 B/op\t       0 allocs/op",
            "extra": "28911 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/tiny_number_of_workers",
            "value": 579944,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "20412 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/educated_guess_number_of_workers",
            "value": 586418,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "19827 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/huge_number_of_workers",
            "value": 241349,
            "unit": "ns/op\t      23 B/op\t       0 allocs/op",
            "extra": "63034 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcherSingleItem/guesstimate",
            "value": 1330,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "8981943 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/no_work",
            "value": 15.61,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "767198203 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/lots_of_work",
            "value": 1444710,
            "unit": "ns/op\t   48000 B/op\t    1000 allocs/op",
            "extra": "8350 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/more_work_than_buffer",
            "value": 1020599,
            "unit": "ns/op\t    4800 B/op\t     100 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/more_workers_than_work",
            "value": 1136706,
            "unit": "ns/op\t    4838 B/op\t     100 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/tiny_batch_size",
            "value": 47691,
            "unit": "ns/op\t    4800 B/op\t     100 allocs/op",
            "extra": "248778 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/huge_batch_size",
            "value": 1046668,
            "unit": "ns/op\t    4803 B/op\t     100 allocs/op",
            "extra": "10000 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/teeny_interval",
            "value": 57260,
            "unit": "ns/op\t    4800 B/op\t     100 allocs/op",
            "extra": "213645 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/slow_worker",
            "value": 6711947,
            "unit": "ns/op\t    4802 B/op\t     100 allocs/op",
            "extra": "1818 times\n2 procs"
          },
          {
            "name": "BenchmarkBatchDispatcher/slow_worker_and_huge_batch_size",
            "value": 6721900,
            "unit": "ns/op\t    4802 B/op\t     100 allocs/op",
            "extra": "1790 times\n2 procs"
          }
        ]
      }
    ]
  }
}