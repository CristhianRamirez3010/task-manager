package responseDto

import "github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"

type ResponseDto struct {
	Content interface{}
	Error   errorManagerDto.ErrorManagerDto
	Message string
}
