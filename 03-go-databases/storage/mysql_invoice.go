package storage

import (
	"database/sql"
	"fmt"

	"github.com/AJRDRGZ/go-db/pkg/invoice"
	"github.com/AJRDRGZ/go-db/pkg/invoiceheader"
	"github.com/AJRDRGZ/go-db/pkg/invoiceitem"
)

// MySQLInvoice used for work with MySQL - invoice
type MySQLInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// NewMySQLInvoice return a new pointer of MySQLInvoice
func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *MySQLInvoice {
	return &MySQLInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implement the interface invoice.Storage
func (p *MySQLInvoice) Create(m *invoice.Model) error {
	// iniciar una transacción
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	// transaccion para el encabezado de la factura
	// si hay un error, hacer rollback para deshacer los cambios
	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("header: %w", err)
	}
	fmt.Printf("Factura creada con id: %d \n", m.Header.ID)

	// transaccion para el detalle de la factura
	// si hay un error, hacer rollback para deshacer los cambios
	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("items: %w", err)
	}
	fmt.Printf("items creados: %d \n", len(m.Items))

	// si todo salió bien, hacer commit para guardar los cambios
	return tx.Commit()
}
