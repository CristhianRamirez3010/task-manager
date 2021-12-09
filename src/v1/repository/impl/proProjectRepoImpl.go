package impl

import (
	"fmt"
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/proProjectModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usprPersonProjectModel"
)

type ProProjectRepoImpl struct {
	repositoryBase
}

func BuildProProjectRepoImpl() *ProProjectRepoImpl {
	return &ProProjectRepoImpl{
		repositoryBase: repositoryBase{
			Constants: constants.BuildConstants(),
			Table:     proProjectModel.TABLE_NAME,
			Fields: []string{
				proProjectModel.ID,
				proProjectModel.NAME,
				proProjectModel.DESCRIPTION,
				proProjectModel.USER_REGISTER,
				proProjectModel.DATE_REGISTER,
				proProjectModel.USER_UPDATE,
				proProjectModel.DATE_UPDATE,
			},
		},
	}
}

func (p *ProProjectRepoImpl) FindProjectByUser(user *usePersonalDataModel.UsePersonalDataModel) ([]*proProjectModel.ProProjectModel, *errorManagerDto.ErrorManagerDto) {
	db, errDto := p.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()
	projectRef := "proj"
	persProjRef := "uprpepr"
	personRef := "person"
	query := p.buildQuery([]*string{
		p.addSelectWithRef(projectRef),
		p.addInnerJoin(usprPersonProjectModel.TABLE_NAME, persProjRef, *utils.BuildStrFromArray([]string{
			fmt.Sprintf("%s.%s = %s.%s", projectRef, proProjectModel.ID, persProjRef, usprPersonProjectModel.PROJECT_ID),
		})),
		p.addInnerJoin(usePersonalDataModel.TABLE_NAME, personRef, *utils.BuildStrFromArray([]string{
			fmt.Sprintf("%s.%s=%s.%s", personRef, usePersonalDataModel.ID, persProjRef, usprPersonProjectModel.PEROSNALDATA_ID),
		})),
		p.addWhere(*utils.BuildStrFromArray([]string{
			fmt.Sprintf("%s.%s = ?", personRef, usePersonalDataModel.ID),
		})),
	})
	fmt.Println(*query, user.Id)
	rows, err := db.Query(*query, user.Id)
	if err != nil {
		return nil, utils.Logger("Error with the query (FindProjectByUser,ProjectRepo)", errDefault, http.StatusInternalServerError, err.Error())
	}

	var projects []*proProjectModel.ProProjectModel

	for rows.Next() {
		var project proProjectModel.ProProjectModel
		project.ScanModel(rows)
		projects = append(projects, &project)
	}

	return projects, nil
}
