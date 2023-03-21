package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "escalona_db"
)

type Evento struct {
	id        int
	nome      string
	data_hora string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error: Could not connect to the database", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not ping the database", err)
	}

	fmt.Println("Successfully connected!")

	// Inserir evento na tabela de escalonamento
	evento := Evento{
		nome:      "Evento de teste",
		data_hora: "2023-03-25 08:00:00",
	}

	sqlStatement := `INSERT INTO escalonamento (nome, data_hora) VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, evento.nome, evento.data_hora)
	if err != nil {
		log.Fatal("Error: Could not insert event", err)
	}

	fmt.Println("Successfully inserted event!")
}
