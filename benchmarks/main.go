// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"github.com/XeLabs/go-mysqlstack/driver"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pubnative/mysqldriver-go"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

const (
	user     = "benchyou"
	passwd   = "benchyou"
	address  = "192.168.0.2:3306"
	database = "sbtest"
)

func main() {
	debug.SetGCPercent(-1) // disable GC

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, passwd, address, database)
	db := mysqldriver.NewDB(dsn, 10)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	stack, err := driver.NewConn(user, passwd, address, database)
	if err != nil {
		panic(err)
	}

	suites := []int{100, 1000, 2000, 4000, 10000, 50000}
	for _, suite := range suites {
		preFillRecords(suite)
		objectsInHEAP(func() string { return readAllMysqlstack(stack) })
		objectsInHEAP(func() string { return readAllMysqldriver(db) })
		objectsInHEAP(func() string { return readAllGoSqlDriver(sqlDB) })
		fmt.Println("")
	}

	fmt.Println("Legend:\n" +
		" * HEAP:       total number of allocated objects\n" +
		" * Mallocs:    number of mallocs\n" +
		" * AllocBytes: bytes allocated (even if freed)\n")
}

func readAllMysqlstack(db driver.Conn) string {
	rows, err := db.Query("SELECT name FROM mysqldriver_benchmarks")
	if err != nil {
		panic(err)
	}

	count := 0
	for rows.Next() {
		name := rows.Datas()
		count++
		_ = name
	}

	return "go-mysqlstack:\trecords read " + strconv.Itoa(count)
}

func readAllMysqldriver(db *mysqldriver.DB) string {
	conn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	defer db.PutConn(conn)

	rows, err := conn.Query("SELECT name FROM mysqldriver_benchmarks")
	if err != nil {
		panic(err)
	}

	count := 0
	for rows.Next() {
		name := rows.String()
		count++
		_ = name
	}

	return "mysqldriver-go:\trecords read " + strconv.Itoa(count)
}

func readAllGoSqlDriver(db *sql.DB) string {
	rows, err := db.Query("SELECT name FROM mysqldriver_benchmarks")
	if err != nil {
		panic(err)
	}

	count := 0
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		count++
		_ = name
	}

	return "go-sql-driver:\trecords read " + strconv.Itoa(count)
}

func objectsInHEAP(fn func() string) {
	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)
	objects := memStats.HeapObjects
	totalAlloc := memStats.TotalAlloc
	mallocs := memStats.Mallocs
	now := time.Now()
	prefix := fn()
	took := time.Since(now)
	runtime.ReadMemStats(memStats)
	diff := memStats.HeapObjects - objects
	adiff := memStats.TotalAlloc - totalAlloc
	mdiff := memStats.Mallocs - mallocs
	fmt.Println(prefix, "\tHEAP", diff, "\tMallocs", mdiff, "\tAllocBytes", adiff, "\ttime", took)
}

func preFillRecords(num int) {
	conn, err := driver.NewConn(user, passwd, address, database)
	if err != nil {
		panic(err)
	}
	if _, err := conn.Exec(`DROP TABLE IF EXISTS mysqldriver_benchmarks`); err != nil {
		panic(err)
	}
	if _, err := conn.Exec(`CREATE TABLE mysqldriver_benchmarks (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255),
		age int,
		PRIMARY KEY (id)
	)`); err != nil {
		panic(err)
	}

	for i := 0; i < num; i++ {
		_, err := conn.Exec(`INSERT INTO mysqldriver_benchmarks(name) VALUES("name` + strconv.Itoa(i) + `")`)
		if err != nil {
			panic(err)
		}
	}

	// warm up
	rows, err := conn.Query("SELECT name FROM mysqldriver_benchmarks")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
	}

	conn.Close()
}
