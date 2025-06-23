package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("ошибка подключения к базе данных:", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("ошибка при попытке подключения к базе данных:", err)
	}

	// создание таблицы из кода файла SQL
	if err := createTasksTableFromFile(conn); err != nil {
		log.Fatal("ошибка при создании таблицы:", err)
	}

	DB = conn
}

func createTasksTableFromFile(conn *pgx.Conn) error {
	data, err := os.ReadFile("./pkg/database/schema.sql")
	if err != nil {
		return fmt.Errorf("ошибка чтения файла SQL: %v", err)
	}

	_, err = conn.Exec(context.Background(), string(data))
	if err != nil {
		return fmt.Errorf("ошибка выполнения SQL запроса: %v", err)
	}

	return nil
}
