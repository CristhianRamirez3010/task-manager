package impl

import (
	"fmt"
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
)

type UseLoginRepoImpl struct {
	repositoryBase
}

func BuildUseLoginRepoImpl() UseLoginRepoImpl {
	return UseLoginRepoImpl{
		repositoryBase: repositoryBase{
			Constants: constants.BuildConstants(),
			Table:     useLoginModel.TABLE_NAME,
			Fields: []string{
				useLoginModel.ID,
				useLoginModel.EMAIL,
				useLoginModel.USER,
				useLoginModel.PASSOWRD,
				useLoginModel.USER_REGISTER,
				useLoginModel.DATE_REGISTER,
				useLoginModel.USER_UPDATE,
				useLoginModel.DATE_UPDATE,
			},
		},
	}
}

func (u UseLoginRepoImpl) FindByEmailAndUser(email string, user string) ([]*useLoginModel.UseLoginModel, *errorManagerDto.ErrorManagerDto) {
	var usersLogin []*useLoginModel.UseLoginModel
	db, errDto := u.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	rows, err := db.Query(u.selectAll(fmt.Sprintf("%s =$1 or %s = $2", useLoginModel.EMAIL, useLoginModel.USER)), email, user)
	if err != nil {
		return nil, utils.Logger("Error with the query in UseLoginRepo", errDefault, http.StatusInternalServerError, err.Error())
	}

	for rows.Next() {
		userLogin := useLoginModel.UseLoginModel{}
		userLogin.ScanModel(rows)
		usersLogin = append(usersLogin, &userLogin)
	}

	return usersLogin, nil
}
