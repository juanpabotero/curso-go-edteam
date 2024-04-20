package storage

import (
	"database/sql"
	"fmt"

	"github.com/AJRDRGZ/go-db/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id) 
	)`
	psqlCreateProduct  = `INSERT INTO products(name, observations, price, created_at) VALUES($1, $2, $3, $4) RETURNING id`
	psqlGetAllProduct  = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
	psqlUpdateProduct  = `UPDATE products SET name = $1, observations = $2,
	price = $3, updated_at = $4 WHERE id = $5`
	psqlDeleteProduct = `DELETE FROM products WHERE id = $1`
)

// psqlProduct used for work with postgres - product
type psqlProduct struct {
	db *sql.DB
}

// newPsqlProduct return a new pointer of PsqlProduct
func newPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db}
}

// Migrate implement the interface product.Storage
func (p *psqlProduct) Migrate() error {
	// preparar la query
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	// cerrar el statement para liberar los recursos
	defer stmt.Close()
	// ejecutar la query
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de producto ejecutada correctamente")
	return nil
}

// Create implement the interface product.Storage
func (p *psqlProduct) Create(m *product.Model) error {
	// preparar la query
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	// cerrar el statement para liberar los recursos
	defer stmt.Close()
	// ejecutar la query
	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println("se creo el producto correctamente")
	return nil
}

// GetAll implement the interface product.Storage
func (p *psqlProduct) GetAll() (product.Models, error) {
	// preparar la query
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	// cerrar el statement para liberar los recursos
	defer stmt.Close()
	// ejecutar la query
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	// cerrar las filas para liberar los recursos
	defer rows.Close()
	// crear un slice de productos
	ms := make(product.Models, 0)
	// recorrer las filas
	for rows.Next() {
		// escanear las filas
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// GetByID implement the interface product.Storage
func (p *psqlProduct) GetByID(id uint) (*product.Model, error) {
	// preparar la query
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	// cerrar el statement para liberar los recursos
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Update implement the interface product.Storage
func (p *psqlProduct) Update(m *product.Model) error {
	// preparar la query
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	// cerrar el statement para liberar los recursos
	defer stmt.Close()
	// ejecutar la query
	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con id: %d", m.ID)
	}

	fmt.Println("se actualizó el producto correctamente")
	return nil
}

// Delete implement the interface product.Storage
func (p *psqlProduct) Delete(id uint) error {
	// preparar la query
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	// cerrar el statement para liberar los recursos
	defer stmt.Close()
	// ejecutar la query
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con id: %d", id)
	}

	fmt.Println("se eliminó el producto correctamente")
	return nil
}
