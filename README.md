# rbtree_skiplist_benchmark


goos: linux
goarch: amd64
Benchmark_RbTreeNew-16          20000000               111 ns/op
Benchmark_SkipListNew-16        10000000               162 ns/op
Benchmark_MapNew-16             30000000                56.4 ns/op
Benchmark_RbTreeInsert-16        3000000               510 ns/op
Benchmark_SkipListInsert-16      5000000               439 ns/op
Benchmark_MapInsert-16           5000000               345 ns/op
Benchmark_RbTreeLoad-16         50000000                30.5 ns/op
Benchmark_SkipListLoad-16       20000000                64.6 ns/op
Benchmark_MapLoad-16            100000000               21.9 ns/op
Benchmark_RbTreeRange-16          200000              6201 ns/op
Benchmark_SkipListRange-16        300000              5513 ns/op
Benchmark_MapRange-16             100000             22849 ns/op
PASS
ok      _/home/cpp2go/test/rbtree_skiplist_benchmark-master       23.249s
