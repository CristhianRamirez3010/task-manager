package utils

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
)

func GetName(i interface{}, fieldName string, tagName string) string {
	field, ok := reflect.TypeOf(i).Elem().FieldByName(fieldName)
	if !ok {
		panic("field not found")
	}
	return string(field.Tag.Get(tagName))
}

func GetNameByDb(i interface{}, fieldName string) string {
	return GetName(i, fieldName, "db")
}

func Logger(errIn string, errOut string, status int, errExeption string) *errorManagerDto.ErrorManagerDto {
	errCode := time.Now().Unix()
	log.Printf("|ErrIn:%s: (%d)|", errIn, errCode)
	if errExeption != "" {
		log.Printf("|ErrExeption:%s: (%d)|", errExeption, errCode)
	}
	return &errorManagerDto.ErrorManagerDto{
		Message: fmt.Sprintf("%s (%d)", errOut, errCode),
		Status:  status,
	}
}
