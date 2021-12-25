package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
)

func Logger(errIn string, errOut string, status int, errExeption string) *errorManagerDto.ErrorManagerDto {
	errCode := time.Now().Unix()
	log.Printf("|ErrIn:%s: (%d)|", errIn, errCode)
	log.Printf("|ErrOut:%s: (%d)|", errOut, errCode)
	if errExeption != "" {
		log.Printf("|ErrExeption:%s: (%d)|", errExeption, errCode)
	}
	return &errorManagerDto.ErrorManagerDto{
		Message: fmt.Sprintf("%s (%d)", errOut, errCode),
		Status:  status,
	}
}

func BuildStrFromArray(strList ...string) string {
	strReturn := ""

	for _, str := range strList {
		strReturn = fmt.Sprintf("%s%s", strReturn, str)
	}

	return strReturn
}

func ResponseManager(response *responseDto.ResponseDto) (int, interface{}) {
	if response.Error == (errorManagerDto.ErrorManagerDto{}) {
		return http.StatusOK, response
	} else {
		return response.Error.Status, response
	}
}

func ResponseErrorManager(errDto *errorManagerDto.ErrorManagerDto) *responseDto.ResponseDto {
	response := new(responseDto.ResponseDto)
	response.Error = *errDto
	if response.Error.Status == 0 {
		response.Error.Status = http.StatusInternalServerError
	}
	return response

}

func ResponseMessageManager(message string) *responseDto.ResponseDto {
	response := new(responseDto.ResponseDto)
	response.Message = message
	return response
}

func ResponseMssgAndContentManager(message string, content interface{}) *responseDto.ResponseDto {
	response := new(responseDto.ResponseDto)
	response.Message = message
	response.Content = content
	return response
}
