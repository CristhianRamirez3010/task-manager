package models

import (
	"database/sql"
	"log"
)

type UseDocumentModel struct {
	Id                   int64
	Name                 string
	Type                 string
	UseDocumentTypeModel UseDocumentTypeModel

	AuditorialBase
}

func (d *UseDocumentModel) ScanModel(rows *sql.Rows) {

	err := rows.Scan(
		&d.Id,
		&d.Name,
		&d.Type,
		&d.UserRegister,
		&d.DateRegister,
		&d.UserUpdate,
		&d.UserUpdate,
	)
	if err != nil {
		log.Fatal("err", err.Error())
	}
}
