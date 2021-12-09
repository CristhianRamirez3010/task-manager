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

type UseHistoryTokensRepoImpl struct {
	repositoryBase
}

func BuildUseHistoryTokensRepoImpl() *UseHistoryTokensRepoImpl {
	return &UseHistoryTokensRepoImpl{
		repositoryBase: repositoryBase{
			Constants: constants.BuildConstants(),
			Table:     useHistoryTokensModel.TABLE_NAME,
			Fields: []string{
				useHistoryTokensModel.ID,
				useHistoryTokensModel.TOKEN,
				useHistoryTokensModel.FINISH,
				useHistoryTokensModel.LOGIN_ID,
				useHistoryTokensModel.USER_REGISTER,
				useHistoryTokensModel.DATE_REGISTER,
				useHistoryTokensModel.USER_UPDATE,
				useHistoryTokensModel.DATE_UPDATE,
			},
		},
	}
}

func (h *UseHistoryTokensRepoImpl) NewToken(hToken *useHistoryTokensModel.UseHistoryToekensModel) *errorManagerDto.ErrorManagerDto {
	db, errDto := h.loadConnection()
	if errDto != nil {
		return errDto
	}
	defer db.Close()

	query := h.buildQuery([]*string{
		h.addInsert(),
	})
	_, err := db.Exec(*query,
		0,
		hToken.Token,
		hToken.Finish,
		hToken.LoginId,
		hToken.UserRegister,
		hToken.DateRegister,
		hToken.UserUpdate,
		hToken.DateUpdate,
	)
	if err != nil {
		return utils.Logger("Error with the insert (NewToken,HistoryTokeRepo)", errDefault, http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (h *UseHistoryTokensRepoImpl) GetLastTokenByPersonId(personId int64) (*useHistoryTokensModel.UseHistoryToekensModel, *errorManagerDto.ErrorManagerDto) {
	db, errDto := h.loadConnection()
	if errDto != nil {
		return nil, errDto
	}
	defer db.Close()

	tokenRef := "token"
	personRef := "person"
	loginRef := "login"
	query := h.buildQuery([]*string{
		h.addSelectWithRef(tokenRef),
		h.addInnerJoin(useLoginModel.TABLE_NAME, loginRef, *utils.BuildStrFromArray([]string{
			fmt.Sprintf(" %s.%s = %s.%s ", loginRef, useLoginModel.ID, tokenRef, useHistoryTokensModel.LOGIN_ID),
		})),
		h.addInnerJoin(usePersonalDataModel.TABLE_NAME, personRef, *utils.BuildStrFromArray([]string{
			fmt.Sprintf(" %s.%s = %s.%s ", personRef, usePersonalDataModel.LOGIN_ID, loginRef, useHistoryTokensModel.ID),
		})),
		h.addWhere(*utils.BuildStrFromArray([]string{
			fmt.Sprintf(" %s.%s=? ", personRef, usePersonalDataModel.ID),
			fmt.Sprintf(" order by %s.%s  desc ", tokenRef, useHistoryTokensModel.ID),
		})),
	})
	rows, err := db.Query(*query, personId)
	if err != nil {
		return nil, utils.Logger("Error with teh query (GetLastTokenByPersonId,historyTokeRepo)", errDefault, http.StatusInternalServerError, err.Error())
	}

	var tokenModel useHistoryTokensModel.UseHistoryToekensModel
	if rows.Next() {
		errDto = tokenModel.ScanModel(rows)
		if errDto != nil {
			return nil, errDto
		}
	}
	return &tokenModel, nil
}
