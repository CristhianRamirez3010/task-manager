package impl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useDocumentTypeModel"
)

type UseDocumentRepoImpl struct {
	repositoryBase
}

func BuildUseDocumentImpl() *UseDocumentRepoImpl {
	return &UseDocumentRepoImpl{
		repositoryBase: repositoryBase{
			Constants: constants.BuildConstants(),
			Table:     useDocumentTypeModel.TABLE_NAME,
			Fields: []string{
				useDocumentTypeModel.ID,
				useDocumentTypeModel.NAME,
				useDocumentTypeModel.DESCRIPTION,
				useDocumentTypeModel.USER_REGISTER,
				useDocumentTypeModel.DATE_REGISTER,
				useDocumentTypeModel.USER_UPDATE,
				useDocumentTypeModel.DATE_UPDATE,
			},
		},
	}
}

func (u *UseDocumentRepoImpl) GetAll() ([]*useDocumentTypeModel.UseDocumentTypeModel, *errorManagerDto.ErrorManagerDto) {
	var documents []*useDocumentTypeModel.UseDocumentTypeModel
	db, errDto := u.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	query := u.buildQuery([]*string{
		u.addSelect(),
	})
	rows, err := db.Query(*query)
	if err != nil {
		return nil, utils.Logger("Error with the query", errDefault, http.StatusInternalServerError, err.Error())
	}

	for rows.Next() {
		document := useDocumentTypeModel.UseDocumentTypeModel{}
		document.ScanModel(rows)
		documents = append(documents, &document)
	}

	return documents, nil
}
