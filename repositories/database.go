package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golobby/container/v3"
	"github.com/brianmorais/go-user-importation/configuration"

	_ "github.com/denisenkom/go-mssqldb"
)

func getWriteConnection() *sql.DB {
	var config configuration.Settings
	container.Resolve(&config)

	dbSettings := config.Databases.WriteDatabase
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", dbSettings.Server, dbSettings.User, dbSettings.Password, dbSettings.Port, dbSettings.DatabaseName)

	sqlObj, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatalln("Erro ao abrir conexão com banco de escrita:", err.Error())
	}

	return sqlObj
}

func getReadConnection() *sql.DB {
	var config configuration.Settings
	container.Resolve(&config)

	dbSettings := config.Databases.ReadDatabase
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", dbSettings.Server, dbSettings.User, dbSettings.Password, dbSettings.Port, dbSettings.DatabaseName)

	sqlObj, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatalln("Erro ao abrir conexão com banco de leitura:", err.Error())
	}

	return sqlObj
}
