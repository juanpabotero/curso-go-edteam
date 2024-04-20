package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/AJRDRGZ/go-db/pkg/product"
	// importar el driver
	// _ significa que se usa el paquete pero no se llama directamente
	// se llama indirectamente a través de la función init() del paquete
	// previene usar metodos internos del paquete y cometer errores
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New: create the connection with db
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	// once permite ejecutar una sola vez una función para aplicar el patron singleton,
	// asi se llame multiples veces, solo se ejecuta una vez
	once.Do(func() {
		var err error
		// conectar a la base de datos
		db, err = sql.Open("postgres", "postgres://postgres:password@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}
		// comprobar la conexión
		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

func newMySQLDB() {
	// once permite ejecutar una sola vez una función para aplicar el patron singleton,
	// asi se llame multiples veces, solo se ejecuta una vez
	once.Do(func() {
		var err error
		// conectar a la base de datos
		// parseTime=true permite que los campos de tipo fecha sean manejados como time.Time
		db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}
		// comprobar la conexión
		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a mySQL")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}

// manejo de nulos
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	// si la fecha no es cero, entonces es válida
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	// estructuras intermedias para manejar nulos
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	// escanear los valores
	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}

// DAOProduct factory of product.Storage
func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQL:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("Driver not implemented")
	}
}
