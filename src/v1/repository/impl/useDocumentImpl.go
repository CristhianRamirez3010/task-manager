package impl

import (
	"log"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/connections"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models"
)

type UseDocumentImpl struct{}

func BuildUseDocumentImpl() UseDocumentImpl {
	return UseDocumentImpl{}
}

func (u *UseDocumentImpl) GetAll() []models.UseDocumentModel {
	var documents []models.UseDocumentModel

	db, err := connections.BuildMySQLConnection(
		"",
		1,
		1,
		1,
	).ConnectDBMysql()

	if err != nil {
		log.Fatal("error with connection", err)
	}
	defer db.Close()

	rows, err := db.Query("")
	if err != nil {
		log.Fatal("Problems with the query", err)
	}

	for rows.Next() {
		document := models.UseDocumentModel{}
		document.ScanModel(rows)
		documents = append(documents, document)
	}

	return documents
}
