package impl

import (
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/utils"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/usePersonalDataModel"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/repository"
)

const (
	errDefault      = "Error with the handler"
	errTokenInvalid = "Token invalid"
)

type handlerBase struct {
	contextDto *contextDto.ContextDto
}

func (h *handlerBase) validateToken() (*usePersonalDataModel.UsePersonalDataModel, *errorManagerDto.ErrorManagerDto) {
	if h.contextDto.AccessToken == "" {
		return nil, utils.Logger(errTokenInvalid, errTokenInvalid, http.StatusPreconditionFailed, "")
	}
	personalDataRepo := repository.BuildIUsePersonalDataRepo()
	historyTokenRepo := repository.BuildIUseHistoryTokensRepo()

	perosnalDataModel, errDto := personalDataRepo.GetDataByToken(h.contextDto.AccessToken)
	if errDto != nil {
		return nil, errDto
	} else if perosnalDataModel.Id < 1 {
		return nil, utils.Logger(errTokenInvalid, errTokenInvalid, http.StatusPreconditionFailed, "")
	}

	historyTokenModel, errDto := historyTokenRepo.GetLastTokenByPersonId(perosnalDataModel.Id)
	if errDto != nil {
		return nil, errDto
	} else if historyTokenModel.Id < 1 {
		return nil, utils.Logger("token don't found in data base", errTokenInvalid, http.StatusPreconditionFailed, "")
	} else if historyTokenModel.Token != h.contextDto.AccessToken {
		return nil, utils.Logger(errTokenInvalid, errTokenInvalid, http.StatusPreconditionFailed, "")
	}

	return perosnalDataModel, nil
}
