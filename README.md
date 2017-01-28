[![Build Status](https://travis-ci.org/XeLabs/go-mysqlstack.png)](https://travis-ci.org/XeLabs/go-mysqlstack) [![codecov.io](https://codecov.io/gh/XeLabs/go-mysqlstack/graphs/badge.svg)](https://codecov.io/gh/XeLabs/go-mysqlstack/branch/master)

# What go-mysqlstack is

***go-mysqlstack*** is an Unmatched & GC-Optimized MySQL protocol library implementing, use as little as possible, go-mysqlstack never wastes one malloc.

Protocol is based on [mysqlproto-go](https://github.com/pubnative/mysqlproto-go) and [go-sql-driver](https://github.com/go-sql-driver/mysql)

# How good GC-Optimized is

```
$go run benchmarks/main.go
go-mysqlstack:  records read 100        HEAP 17         Mallocs 20      AllocBytes 2728         time 452.727µs
mysqldriver-go: records read 100        HEAP 83         Mallocs 136     AllocBytes 5464         time 1.101469ms
go-sql-driver:  records read 100        HEAP 292        Mallocs 346     AllocBytes 13632        time 1.011363ms

go-mysqlstack:  records read 1000       HEAP 25         Mallocs 28      AllocBytes 15016        time 729.425µs
mysqldriver-go: records read 1000       HEAP 514        Mallocs 1016    AllocBytes 22288        time 897.252µs
go-sql-driver:  records read 1000       HEAP 2510       Mallocs 3013    AllocBytes 58048        time 1.040523ms

go-mysqlstack:  records read 2000       HEAP 31         Mallocs 34      AllocBytes 27048        time 1.085919ms
mysqldriver-go: records read 2000       HEAP 1023       Mallocs 2025    AllocBytes 44112        time 1.121449ms
go-sql-driver:  records read 2000       HEAP 5007       Mallocs 6010    AllocBytes 112256       time 1.302811ms

go-mysqlstack:  records read 4000       HEAP 48         Mallocs 51      AllocBytes 53160        time 1.631916ms
mysqldriver-go: records read 4000       HEAP 2037       Mallocs 4039    AllocBytes 84432        time 2.061695ms
go-sql-driver:  records read 4000       HEAP 10007      Mallocs 12010   AllocBytes 224256       time 1.965674ms

go-mysqlstack:  records read 10000      HEAP 100        Mallocs 103     AllocBytes 133048       time 3.409044ms
mysqldriver-go: records read 10000      HEAP 5090       Mallocs 10092   AllocBytes 213840       time 3.831417ms
go-sql-driver:  records read 10000      HEAP 25007      Mallocs 30010   AllocBytes 560272       time 4.016614ms

go-mysqlstack:  records read 50000      HEAP 475        Mallocs 478     AllocBytes 709048       time 15.189725ms
mysqldriver-go: records read 50000      HEAP 45463      Mallocs 50465   AllocBytes 1426768      time 16.87461ms
go-sql-driver:  records read 50000      HEAP 145007     Mallocs 150010  AllocBytes 3120272      time 17.252933ms

Legend:
 * HEAP:       total number of allocated objects
 * Mallocs:    number of mallocs
 * AllocBytes: bytes allocated (even if freed)
```
