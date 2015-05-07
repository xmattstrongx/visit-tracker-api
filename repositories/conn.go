package repositories

import (
	"RestApiProject/configuration"
	"database/sql"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

const driverName = "postgres"

func initDB() (*runner.DB, error) {
	config := configuration.LoadConfiguration("config.json")
	runner, err := connectWIthTimeout(config.PsqlConfig)
	if err != nil {
		return nil, newBadConnectionError(err)
	}
	return runner, nil
}

func connectWIthTimeout(config configuration.PsqlConfiguration) (*runner.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", config.Username, config.Password, config.Hostname, config.Port, config.DatabaseName, config.SslMode)
	log.Print(connectionString)
	log.Printf("Matt connstring: %s", connectionString)

	// Before returning a connection we will try to ping the database to verify connection
	var dbRunner *runner.DB
	connected := make(chan error)
	go func() {
		db, err := sql.Open(driverName, connectionString)
		if err != nil {
			connected <- err
		}
		runner.MustPing(db)
		dbRunner = runner.NewDB(db, driverName)
		connected <- nil
	}()

	select {
	case result := <-connected:
		return dbRunner, result
	case <-time.After(time.Duration(10) * time.Second):
		return nil, newConnectionTimeoutError()
	}
}
