// Демо MSSQL-модуль формирования DSN и создания подключения к СУБД.

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/url"
	"sqlinsertrs"
	"strconv"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	user          = flag.String("user", "drxadmin", "the database user")
	password      = flag.String("password", "drxpassword", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "mssql.directum.server", "the database server")
	database      = flag.String("database", "DirectumRX", "the database name")
)

// Функция формирования строки подключения, dsn.
func makeConnURL() *url.URL {
	flag.Parse()
	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	var userInfo *url.Userinfo
	if *user != "" {
		userInfo = url.UserPassword(*user, *password)
	}
	return &url.URL{
		Scheme: "sqlserver",
		Host:   *server + ":" + strconv.Itoa(*port),
		User:   userInfo,
	}
}

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", *server, *user, *password, *port, *database)
	if *debug {
		fmt.Printf("connString:%s\n", connString)
	}

	// Create a new connector object by calling NewConnector. Создание объекта для подключения, через вызов NewConnector
	connector, err := mssql.NewConnector(connString)
	if err != nil {
		log.Println(err)
		return
	}

	// Use SessionInitSql to set any options that cannot be set with the dsn string
	// SessionInitSql используется  для установки любых параметров, которые нельзя установить с помощью строки dsn.
	// With ANSI_NULLS set to ON, compare NULL data with = NULL or <> NULL will return 0 rows
	// Если ANSI_NULLS установлено ON, любое сравнение с NULL вернет 0 строк.
	connector.SessionInitSQL = "SET ANSI_NULLS ON"

	// Pass connector to sql.OpenDB to get a sql.DB object. Получение объекта sql.DB
	db := sql.OpenDB(connector)
	defer db.Close()

	cs := make(chan string) // канал функции sqlinsert-запроса
	start := time.Now()
	go sqlinsertrs.SqlInserTrs(db, cs)
	log.Println("\n\nРезультат sql-запроса: ", <-cs) // Получение данных запроса из канала горутины
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs время выполнения запроса\n", secs)
}
