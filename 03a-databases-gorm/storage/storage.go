package storage

import (
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
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
		db, err = gorm.Open("postgres", "postgres://edteam:edteam@localhost:7530/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
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
		db, err = gorm.Open("mysql", "edteam:edteam@tcp(localhost:7531)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("conectado a mySQL")
	})
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}
