package storage

import (
	"database/sql"
	"fmt"

	"github.com/AJRDRGZ/go-db/pkg/invoice"
	"github.com/AJRDRGZ/go-db/pkg/invoiceheader"
	"github.com/AJRDRGZ/go-db/pkg/invoiceitem"
)

// PsqlInvoice used for work with postgres - invoice
type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// NewPsqlInvoice return a new pointer of PsqlInvoice
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Create implement the interface invoice.Storage
func (p *PsqlInvoice) Create(m *invoice.Model) error {
	// iniciar una transacción
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	// transaccion para el encabezado de la factura
	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		// si hay un error, hacer rollback para deshacer los cambios
		tx.Rollback()
		return fmt.Errorf("header: %w", err)
	}
	fmt.Printf("Factura creada con id: %d \n", m.Header.ID)

	// transaccion para el detalle de la factura
	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		// si hay un error, hacer rollback para deshacer los cambios
		tx.Rollback()
		return fmt.Errorf("items: %w", err)
	}
	fmt.Printf("items creados: %d \n", len(m.Items))

	// si todo salió bien, hacer commit para guardar los cambios
	return tx.Commit()
}
