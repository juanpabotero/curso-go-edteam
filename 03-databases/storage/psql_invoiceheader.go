package storage

import (
	"database/sql"
	"fmt"

	"github.com/AJRDRGZ/go-db/pkg/invoiceheader"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id) 
	)`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers (client) VALUES ($1) RETURNING id, created_at`
)

// PsqlInvoiceHeader used for work with postgres - invoiceHeader
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader return a new pointer of PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Migrate implement the interface invoiceHeader.Storage
func (p *PsqlInvoiceHeader) Migrate() error {
	// preparar la query
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
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

	fmt.Println("migración de invoiceHeader ejecutada correctamente")
	return nil
}

// CreateTx implement the interface invoiceHeader.Storage
func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	// preparar la query, se hace sobre la transacción para poder hacer un rollback si es necesario
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	// cerrar el statement para liberar los recursos
	defer stmt.Close()

	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}
