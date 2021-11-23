package impl

import (
	"fmt"
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
)

type UsePersonalDataRepoImpl struct {
	repositoryBase
}

func BuildUsePersonalDataRepoImpl() *UsePersonalDataRepoImpl {
	return &UsePersonalDataRepoImpl{
		repositoryBase: repositoryBase{
			Constants: constants.BuildConstants(),
			Table:     usePersonalDataModel.TABLE_NAME,
			Fields: []string{
				usePersonalDataModel.ID,
				usePersonalDataModel.NAME,
				usePersonalDataModel.SURNAME,
				usePersonalDataModel.IDENTIFICATION,
				usePersonalDataModel.PHONE,
				usePersonalDataModel.COUNTRY,
				usePersonalDataModel.LOGIN_ID,
				usePersonalDataModel.USER_REGISTER,
				usePersonalDataModel.DATE_REGISTER,
				usePersonalDataModel.USER_UPDATE,
				usePersonalDataModel.DATE_UPDATE,
			},
		},
	}
}

func (p *UsePersonalDataRepoImpl) GetDataByLoginId(loginId *int64) (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto) {

	db, errDto := p.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	db.Close()

	rows, err := db.Query(p.selectAll(fmt.Sprintf("%s=?", usePersonalDataModel.LOGIN_ID)), *loginId)
	if err != nil {
		return nil, utils.Logger("Error with the query", errDefault, http.StatusInternalServerError, err.Error())
	}
	var usePersonalData usePersonalDataModel.UsePersonalDataModel
	for rows.Next() {
		usePersonalData.ScanModel(rows)
	}
	return &usePersonalData, nil
}

func (p *UsePersonalDataRepoImpl) New(personaldata *usePersonalDataModel.UsePersonalDataModel) *errorManagerDto.ErrorManagerDto {
	db, errDto := p.loadConnection()
	if errDto != nil {
		return errDto
	}
	defer db.Close()

	_, err := db.Exec(p.insertAll(),
		personaldata.Id,
		personaldata.Name,
		personaldata.Surname,
		personaldata.Identification,
		personaldata.Phone,
		personaldata.Country,
		personaldata.LoginId,
		personaldata.UserRegister,
		personaldata.DateRegister,
		personaldata.UserUpdate,
		personaldata.DateUpdate,
	)
	fmt.Println(p.insertAll())
	if err != nil {
		return utils.Logger("Error with the insert in PersonalDataRepo (New())", errDefault, http.StatusInternalServerError, err.Error())
	}
	return nil
}
