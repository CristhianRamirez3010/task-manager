package models

import (
	"database/sql"
	"log"
)

type UseDocumentTypeModel struct {
	Id          int64
	Name        string
	Description string

	AuditorialBase
}

func (d *UseDocumentTypeModel) ScanModel(rows *sql.Rows) {

	err := rows.Scan(
		&d.Id,
		&d.Name,
		&d.Description,
		&d.UserRegister,
		&d.DateRegister,
		&d.UserUpdate,
		&d.UserUpdate,
	)
	if err != nil {
		log.Fatal("err", err.Error())
	}
}
