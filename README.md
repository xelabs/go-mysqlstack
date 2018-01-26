[![Build Status](https://travis-ci.org/XeLabs/go-mysqlstack.png)](https://travis-ci.org/XeLabs/go-mysqlstack) [![Go Report Card](https://goreportcard.com/badge/github.com/XeLabs/go-mysqlstack)](https://goreportcard.com/report/github.com/XeLabs/go-mysqlstack) [![codecov.io](https://codecov.io/gh/XeLabs/go-mysqlstack/graphs/badge.svg)](https://codecov.io/gh/XeLabs/go-mysqlstack/branch/master)

# What go-mysqlstack is

***go-mysqlstack*** is an MySQL protocol library implementing.

Protocol is based on [mysqlproto-go](https://github.com/pubnative/mysqlproto-go) and [go-sql-driver](https://github.com/go-sql-driver/mysql)

## Running the tests

```
$ mkdir src
$ export GOPATH=`pwd`
$ go get -u github.com/XeLabs/go-mysqlstack/driver
$ cd github.com/XeLabs/go-mysqlstack/
$ make test
```

## Examples

1. examples/mysqld.go mocks a MySQL server by running:

```
$ go run example/mysqld.go
  2018/01/26 16:02:02.304376 mysqld.go:52:     [INFO]    mysqld.server.start.address[:4407]
```

2. examples/client.go mocks a client and query from the mock MySQL server:

```
$ go run example/client.go
  2018/01/26 16:06:10.779340 client.go:32:    [INFO]    results:[[[10 nice name]]]
```

## Production ready
