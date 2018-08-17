i saw funny performance with msgpack+multiwriter so i wrote a thing

Output on August 17th, 2018 (go1.10.3):

```
[git:master]~/go/src/github.com/schmichael/codec-mw-bench$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/schmichael/codec-mw-bench
BenchmarkJSON_Blake2b_Multi-4              20000             64133 ns/op          28.35 MB/s       12594 B/op        212 allocs/op
BenchmarkJSON_Blake2b_NoMulti-4            20000             69513 ns/op          26.15 MB/s       12114 B/op        208 allocs/op
BenchmarkJSON_MD5_Multi-4                  20000             68048 ns/op          26.48 MB/s       12290 B/op        212 allocs/op
BenchmarkJSON_MD5_NoMulti-4                20000             70763 ns/op          25.47 MB/s       12114 B/op        208 allocs/op
BenchmarkMsgpack_Blake2b_Multi-4           50000             38915 ns/op          36.46 MB/s        8464 B/op        214 allocs/op
BenchmarkMsgpack_Blake2b_NoMulti-4        100000             13987 ns/op         101.45 MB/s        5584 B/op         10 allocs/op
BenchmarkMsgpack_MD5_Multi-4               50000             37498 ns/op          37.41 MB/s        8148 B/op        214 allocs/op
BenchmarkMsgpack_MD5_NoMulti-4            100000             14653 ns/op          95.75 MB/s        5584 B/op         10 allocs/op
PASS
ok      github.com/schmichael/codec-mw-bench    16.071s
```

Thanks to @apparentlymart for pointing me to why: https://github.com/ugorji/go/blob/00b869d2f4a5e27445c2d916fa106fc72c106d4c/codec/encode.go#L1036-L1051
