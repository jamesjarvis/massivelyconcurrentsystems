window.BENCHMARK_DATA = {
  "lastUpdate": 1654108987366,
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
          "id": "52acc478a0c1d578ef86e30d1c546c981777ebc3",
          "message": "add storebench.yaml to use github pages for benchmark output",
          "timestamp": "2022-06-01T19:41:46+01:00",
          "tree_id": "590ec629551899769ad629c960bbeb1e90271be5",
          "url": "https://github.com/jamesjarvis/massivelyconcurrentsystems/commit/52acc478a0c1d578ef86e30d1c546c981777ebc3"
        },
        "date": 1654108986616,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkWaitManyDepsWithDeps/chans",
            "value": 23929561,
            "unit": "ns/op\t      25 B/op\t       0 allocs/op",
            "extra": "482 times\n2 procs"
          },
          {
            "name": "BenchmarkWaitManyDepsWithDeps/waitgroups",
            "value": 22601829,
            "unit": "ns/op\t      37 B/op\t       0 allocs/op",
            "extra": "531 times\n2 procs"
          }
        ]
      }
    ]
  }
}