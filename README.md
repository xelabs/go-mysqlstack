[![Build Status](https://travis-ci.org/XeLabs/go-mysqlstack.png)](https://travis-ci.org/XeLabs/go-mysqlstack) [![codecov.io](https://codecov.io/gh/XeLabs/go-mysqlstack/graphs/badge.svg)](https://codecov.io/gh/XeLabs/go-mysqlstack/branch/master)

# What go-mysqlstack is

***go-mysqlstack*** is an Unmatched & GC-Optimized MySQL protocol library implementing

Protocol is based on [mysqlproto-go](https://github.com/pubnative/mysqlproto-go) and [go-sql-driver](https://github.com/go-sql-driver/mysql)

# How good GC-Optimized is

```
$go run benchmarks/main.go
go-mysqlstack:  records read 100        HEAP 10         time 517.087µs
mysqldriver-go: records read 100        HEAP 83         time 1.346706ms
go-sql-driver:  records read 100        HEAP 292        time 1.416546ms

go-mysqlstack:  records read 1000       HEAP 9          time 774.07µs
mysqldriver-go: records read 1000       HEAP 514        time 810.645µs
go-sql-driver:  records read 1000       HEAP 2510       time 1.151545ms

go-mysqlstack:  records read 2000       HEAP 6          time 990.627µs
mysqldriver-go: records read 2000       HEAP 1023       time 1.150353ms
go-sql-driver:  records read 2000       HEAP 5007       time 1.437157ms

go-mysqlstack:  records read 4000       HEAP 7          time 1.787648ms
mysqldriver-go: records read 4000       HEAP 2037       time 1.882848ms
go-sql-driver:  records read 4000       HEAP 10007      time 2.248269ms

go-mysqlstack:  records read 10000      HEAP 6          time 3.376725ms
mysqldriver-go: records read 10000      HEAP 5090       time 3.638042ms
go-sql-driver:  records read 10000      HEAP 25007      time 4.089699ms

go-mysqlstack:  records read 50000      HEAP 6          time 14.705967ms
mysqldriver-go: records read 50000      HEAP 45463      time 16.029054ms
go-sql-driver:  records read 50000      HEAP 145007     time 19.108608ms
```
