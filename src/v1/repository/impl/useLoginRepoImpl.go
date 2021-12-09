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

func BuildUseLoginRepoImpl() *UseLoginRepoImpl {
	return &UseLoginRepoImpl{
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

func (u *UseLoginRepoImpl) FindByEmailAndUser(email *string, user *string) ([]*useLoginModel.UseLoginModel, *errorManagerDto.ErrorManagerDto) {
	var usersLogin []*useLoginModel.UseLoginModel
	db, errDto := u.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	query := u.buildQuery([]*string{
		u.addSelect(),
		u.addWhere(fmt.Sprintf("%s=? or %s=?", useLoginModel.EMAIL, useLoginModel.USER)),
	})
	rows, err := db.Query(*query, email, user)
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

func (u *UseLoginRepoImpl) New(useLogin *useLoginModel.UseLoginModel) *errorManagerDto.ErrorManagerDto {
	db, errDto := u.loadConnection()
	if errDto != nil {
		return errDto
	}
	defer db.Close()

	query := u.buildQuery([]*string{
		u.addInsert(),
	})
	_, err := db.Exec(*query,
		useLogin.Id,
		useLogin.Email,
		useLogin.User,
		useLogin.Password,
		useLogin.UserRegister,
		useLogin.DateRegister,
		useLogin.UserUpdate,
		useLogin.DateUpdate,
	)
	if err != nil {
		return utils.Logger("Error with the insert in LoginRepo(new())", errDefault, http.StatusInternalServerError, err.Error())
	}
	return nil
}
