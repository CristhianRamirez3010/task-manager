package impl

import (
	"fmt"
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useHistoryTokensModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
)

type UsePersonalDataRepoImpl struct {
	loginRepo *UseLoginRepoImpl
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
				usePersonalDataModel.DOCUMENT,
				usePersonalDataModel.PHONE,
				usePersonalDataModel.COUNTRY,
				usePersonalDataModel.DOCUMENTTYPE,
				usePersonalDataModel.LOGIN_ID,
				usePersonalDataModel.USER_REGISTER,
				usePersonalDataModel.DATE_REGISTER,
				usePersonalDataModel.USER_UPDATE,
				usePersonalDataModel.DATE_UPDATE,
			},
		},
	}
}

func (p *UsePersonalDataRepoImpl) LoadUseLoginRepoImpl(login *UseLoginRepoImpl) {
	p.loginRepo = login
}

func (p *UsePersonalDataRepoImpl) GetDataByLoginId(loginId *int64) (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto) {

	db, errDto := p.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	query := p.buildQuery([]*string{
		p.addSelect(),
		p.addWhere(fmt.Sprintf("%s=?", usePersonalDataModel.LOGIN_ID)),
	})
	rows, err := db.Query(*query, *loginId)
	if err != nil {
		return nil, utils.Logger("Error with the query", errDefault, http.StatusInternalServerError, err.Error())
	}
	var usePersonalData usePersonalDataModel.UsePersonalDataModel
	for rows.Next() {
		usePersonalData.ScanModel(rows)
	}
	return &usePersonalData, nil
}

func (p *UsePersonalDataRepoImpl) GetDataByToken(token string) (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto) {
	db, errDto := p.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	personRef := "person"
	loginRef := "login"
	tokenRef := "histTok"

	query := p.buildQuery([]*string{
		p.addSelectWithRef(personRef),
		p.addInnerJoin(useLoginModel.TABLE_NAME, loginRef, fmt.Sprintf(" %s.%s=%s.%s ", loginRef, useLoginModel.ID, personRef, usePersonalDataModel.LOGIN_ID)),
		p.addInnerJoin(useHistoryTokensModel.TABLE_NAME, tokenRef, fmt.Sprintf(" %s.%s=%s.%s ", tokenRef, useHistoryTokensModel.LOGIN_ID, loginRef, useLoginModel.ID)),
		p.addWhere(*utils.BuildStrFromArray([]string{
			fmt.Sprintf(" %s.%s=? ", tokenRef, useHistoryTokensModel.TOKEN),
		})),
	})
	rows, err := db.Query(*query, token)
	if err != nil {
		return nil, utils.Logger("Error with the query (GetDataByToken,PersonalDataRepo)", errDefault, http.StatusInternalServerError, err.Error())
	}
	var personalData usePersonalDataModel.UsePersonalDataModel
	if rows.Next() {
		errDto = personalData.ScanModel(rows)
		if errDto != nil {
			return nil, errDto
		}
	}

	return &personalData, nil
}

func (p *UsePersonalDataRepoImpl) New(personaldata *usePersonalDataModel.UsePersonalDataModel) *errorManagerDto.ErrorManagerDto {
	db, errDto := p.loadConnection()
	if errDto != nil {
		return errDto
	}
	defer db.Close()

	query := p.buildQuery([]*string{
		p.addInsert(),
	})
	_, err := db.Exec(*query,
		personaldata.Id,
		personaldata.Name,
		personaldata.Surname,
		personaldata.Document,
		personaldata.Phone,
		personaldata.Country,
		personaldata.DocumentType,
		personaldata.LoginId,
		personaldata.UserRegister,
		personaldata.DateRegister,
		personaldata.UserUpdate,
		personaldata.DateUpdate,
	)
	if err != nil {
		return utils.Logger("Error with the insert in PersonalDataRepo (New())", errDefault, http.StatusInternalServerError, err.Error())
	}
	return nil
}
