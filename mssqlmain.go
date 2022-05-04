// Демо MSSQL-модуль формирования DSN и создания подключения к СУБД.

package main

import (
	"fmt"
	"log"
	"mssqldsn"

	"sqlinsertrs"
	"time"
)

const (
	insertIntegrTableSql = "INSERT DirectumRX.dbo.dBase SELECT Id, Discriminator, Name, BusinessUnit, ItemNumber, ItemName FROM DirectumRX.dbo.dBaseTest WHERE BusinessUnit = 65;"
)

func main() {
	start := time.Now()
	// структура DSN
	dd := mssqldsn.DataDsn{
		Debug:    true,
		User:     "user",
		Password: "password",
		Port:     1433,
		Server:   "rx-db-directum",
		Database: "DirectumRX",
	}
	// Вызов метода интерфейса, для формирования DSN и создания подключения к СУБД
	var d mssqldsn.ConDsner = dd
	db := d.SqlConDsn()
	defer db.Close()

	cs := make(chan string) // Channel of function mssql-request. Канал функции mssql-запроса
	go sqlinsertrs.SqlInserTrs(insertIntegrTableSql, db, cs)
	log.Println("\nРезультат sql-запроса: ", <-cs) // Получение данных запроса из канала горутины
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs время выполнения запроса\n", secs)
}
