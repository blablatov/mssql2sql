// Демо MSSQL-модуль формирования DSN и создания подключения к СУБД.

package mssqldsn

import (
	"database/sql"
	"fmt"
	"log"

	mssql "github.com/denisenkom/go-mssqldb"
)

type ConDsner interface {
	SqlConDsn() *sql.DB
}

// Структура DSN
type DataDsn struct {
	Debug    bool
	User     string
	Password string
	Port     int
	Server   string
	Database string
}

// Функция DSN
func (dd DataDsn) SqlConDsn() *sql.DB {
	connString := fmt.Sprintf("server=%s; user id=%s; password=%s; port=%d; database=%s", dd.Server, dd.User, dd.Password, dd.Port, dd.Database)
	if dd.Debug {
		fmt.Printf("connString:%s\n", connString)
	}

	// Create a new connector object by calling NewConnector. Создание объекта для подключения, через вызов NewConnector
	connector, err := mssql.NewConnector(connString)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Use SessionInitSql to set any options that cannot be set with the dsn string
	// SessionInitSql используется  для установки любых параметров, которые нельзя установить с помощью строки dsn.
	// With ANSI_NULLS set to ON, compare NULL data with = NULL or <> NULL will return 0 rows
	// Если ANSI_NULLS установлено ON, любое сравнение с NULL вернет 0 строк.
	connector.SessionInitSQL = "SET ANSI_NULLS ON"

	// Pass connector to sql.OpenDB to get a sql.DB object. Получение объекта sql.DB
	db := sql.OpenDB(connector)
	return db
}
