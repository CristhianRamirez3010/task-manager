package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
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

func BuildStrFromArray(strList []string) *string {
	strReturn := ""

	for _, str := range strList {
		strReturn = fmt.Sprintf("%s%s", strReturn, str)
	}

	return &strReturn
}
