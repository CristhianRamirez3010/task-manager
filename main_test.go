package main

import (
	"os"
	"testing"
)

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/CristhianRamirez3010/task-manager-go/src/config/contextDto"
// 	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
// 	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
// 	"github.com/CristhianRamirez3010/task-manager-go/src/v1/dto"
// 	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler"
// 	"github.com/CristhianRamirez3010/task-manager-go/src/v1/models/useLoginModel"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MyMockedObject struct {
// 	mock.Mock
// }

// func (u *MyMockedObject) GetDocuments() *responseDto.ResponseDto {
// 	return &responseDto.ResponseDto{Content: 2}
// }

// func (u *MyMockedObject) ValidateLogin(userLogin *useLoginModel.UseLoginModel) *responseDto.ResponseDto {
// 	return nil
// }

// func (u *MyMockedObject) CreateNewUser(userData *dto.UserdataDto) *responseDto.ResponseDto {
// 	return nil
// }

// func (u *MyMockedObject) createHashToke(loginModel *useLoginModel.UseLoginModel) (string, *errorManagerDto.ErrorManagerDto) {
// 	return "", nil
// }

// func (u *MyMockedObject) saveUserToken(token string, user *useLoginModel.UseLoginModel) *errorManagerDto.ErrorManagerDto {

// 	return nil
// }

// // func TestRun(t *testing.T) {

// // 	c, err := run()
// // 	a := assert.New(t)
// // 	a.Equal(err, nil, "error")
// // 	a.NotEqual(c, nil, "error")

// // }

// func TestRunTest(t *testing.T) {
// 	mock := new(MyMockedObject)
// 	a := assert.New(t)

// 	sert := new(TestModel)
// 	te := func(con *contextDto.ContextDto) handler.IUserHandler {
// 		return mock
// 	}
// 	buildIUserHandler = te
// 	e, err := sert.RunTest()

// 	fmt.Println(e, "-", err, "-", err != nil)

// 	a.Nil(err, "err 1")

// 	a.Equal(e, &responseDto.ResponseDto{Content: 2}, "errr 2")

// }

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
