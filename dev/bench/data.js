window.BENCHMARK_DATA = {
  "lastUpdate": 1709878857490,
  "repoUrl": "https://github.com/trym-b/warchaeology",
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
          "id": "33e32ca664dda8433f3f06097bbf56cbd803f825",
          "message": "commiting to hg-pages",
          "timestamp": "2024-03-08T07:19:59+01:00",
          "tree_id": "43739ff7945e39bd2d75cc1cc4c52d8e8ea60c13",
          "url": "https://github.com/trym-b/warchaeology/commit/33e32ca664dda8433f3f06097bbf56cbd803f825"
        },
        "date": 1709878857031,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFib10",
            "value": 316.7,
            "unit": "ns/op",
            "extra": "3841693 times\n4 procs"
          },
          {
            "name": "BenchmarkFib20",
            "value": 39564,
            "unit": "ns/op",
            "extra": "30318 times\n4 procs"
          },
          {
            "name": "BenchmarkFib20WithAuxMetric",
            "value": 39613,
            "unit": "ns/op\t         4.000 auxMetricUnits",
            "extra": "30519 times\n4 procs"
          },
          {
            "name": "BenchmarkFib20WithAuxMetric - ns/op",
            "value": 39613,
            "unit": "ns/op",
            "extra": "30519 times\n4 procs"
          },
          {
            "name": "BenchmarkFib20WithAuxMetric - auxMetricUnits",
            "value": 4,
            "unit": "auxMetricUnits",
            "extra": "30519 times\n4 procs"
          }
        ]
      }
    ]
  }
}