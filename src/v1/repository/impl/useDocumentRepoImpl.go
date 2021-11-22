package impl

import (
	"fmt"
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/userDocumentTypeModel"
)

const (
	errDefault = "Error with the repository"
)

type UseDocumentRepoImpl struct {
	RepositoryBase
}

func (u UseDocumentRepoImpl) selectAll(where string) string {
	if where != "" {
		where = fmt.Sprintf(" where %s", where)
	}
	return fmt.Sprintf("select %s,%s,%s,%s,%s,%s,%s from %s %s;",
		userDocumentTypeModel.ID,
		userDocumentTypeModel.NAME,
		userDocumentTypeModel.DESCRIPTION,
		userDocumentTypeModel.USER_REGISTER,
		userDocumentTypeModel.DATE_REGISTER,
		userDocumentTypeModel.USER_UPDATE,
		userDocumentTypeModel.DATE_UPDATE,
		userDocumentTypeModel.TABLE_NAME,
		where)
}

func BuildUseDocumentImpl() UseDocumentRepoImpl {
	return UseDocumentRepoImpl{
		RepositoryBase: RepositoryBase{Constants: constants.BuildConstants()},
	}
}

func (u UseDocumentRepoImpl) GetAll() ([]*userDocumentTypeModel.UseDocumentTypeModel, *errorManagerDto.ErrorManagerDto) {
	var documents []*userDocumentTypeModel.UseDocumentTypeModel
	db, errDto := u.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	rows, err := db.Query(u.selectAll(""))
	if err != nil {
		return nil, utils.Logger("Error with the query", errDefault, http.StatusInternalServerError, err.Error())
	}

	for rows.Next() {
		document := userDocumentTypeModel.UseDocumentTypeModel{}
		document.ScanModel(rows)
		documents = append(documents, &document)
	}

	return documents, nil
}
