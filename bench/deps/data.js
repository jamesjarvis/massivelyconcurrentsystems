window.BENCHMARK_DATA = {
  "lastUpdate": 1654110941552,
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
          "id": "9f322e2dbb84d81cbe94f874d77120bf9aa9144c",
          "message": "update to hopefully split up benchmark results into two directories on gh-pages",
          "timestamp": "2022-06-01T19:58:44+01:00",
          "tree_id": "9cfbceed2f6ca18fc53395db5ee806b53272e16e",
          "url": "https://github.com/jamesjarvis/massivelyconcurrentsystems/commit/9f322e2dbb84d81cbe94f874d77120bf9aa9144c"
        },
        "date": 1654110586506,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWaitManyDepsWithDeps/chans",
            "value": 21036936,
            "unit": "ns/op\t      21 B/op\t       0 allocs/op",
            "extra": "583 times\n2 procs"
          },
          {
            "name": "BenchmarkWaitManyDepsWithDeps/waitgroups",
            "value": 19913643,
            "unit": "ns/op\t      12 B/op\t       0 allocs/op",
            "extra": "584 times\n2 procs"
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
          "id": "05b80c2fd6b17849d4b5bb1882a3da4ab32959b2",
          "message": "add MIT License",
          "timestamp": "2022-06-01T20:14:32+01:00",
          "tree_id": "c1879f6fb600eac1ca7502875cd9f073a0a75115",
          "url": "https://github.com/jamesjarvis/massivelyconcurrentsystems/commit/05b80c2fd6b17849d4b5bb1882a3da4ab32959b2"
        },
        "date": 1654110940637,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWaitManyDepsWithDeps/chans",
            "value": 20439103,
            "unit": "ns/op\t       9 B/op\t       0 allocs/op",
            "extra": "631 times\n2 procs"
          },
          {
            "name": "BenchmarkWaitManyDepsWithDeps/waitgroups",
            "value": 19122520,
            "unit": "ns/op\t       9 B/op\t       0 allocs/op",
            "extra": "630 times\n2 procs"
          }
        ]
      }
    ]
  }
}