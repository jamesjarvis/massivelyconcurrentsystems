window.BENCHMARK_DATA = {
  "lastUpdate": 1654110586924,
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
      }
    ]
  }
}