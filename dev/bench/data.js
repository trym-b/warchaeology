window.BENCHMARK_DATA = {
  "lastUpdate": 1709812351513,
  "repoUrl": "https://github.com/nlnwa/warchaeology",
  "entries": {
    "Benchmark results": [
      {
        "commit": {
          "author": {
            "email": "trym.bremnes@gmail.com",
            "name": "Trym bremnes",
            "username": "trym-b"
          },
          "committer": {
            "email": "trym.bremnes@gmail.com",
            "name": "Trym bremnes",
            "username": "trym-b"
          },
          "distinct": true,
          "id": "d6f3df865f61bd83acc99d7d887bea42aa36db6b",
          "message": "fix -bench=.",
          "timestamp": "2024-03-07T12:38:59+01:00",
          "tree_id": "c5e4266db55751665e5dd6cc5e5180f6e30a99cb",
          "url": "https://github.com/nlnwa/warchaeology/commit/d6f3df865f61bd83acc99d7d887bea42aa36db6b"
        },
        "date": 1709812350965,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFib10",
            "value": 310.6,
            "unit": "ns/op",
            "extra": "3870499 times\n4 procs"
          },
          {
            "name": "BenchmarkFib20",
            "value": 39767,
            "unit": "ns/op",
            "extra": "30352 times\n4 procs"
          },
          {
            "name": "BenchmarkFib20WithAuxMetric",
            "value": 39662,
            "unit": "ns/op\t         4.000 auxMetricUnits",
            "extra": "30310 times\n4 procs"
          },
          {
            "name": "BenchmarkFib20WithAuxMetric - ns/op",
            "value": 39662,
            "unit": "ns/op",
            "extra": "30310 times\n4 procs"
          },
          {
            "name": "BenchmarkFib20WithAuxMetric - auxMetricUnits",
            "value": 4,
            "unit": "auxMetricUnits",
            "extra": "30310 times\n4 procs"
          }
        ]
      }
    ]
  }
}