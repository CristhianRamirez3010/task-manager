package responseDto

import "github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"

type DesponseDto struct {
	Content interface{}
	Error   errorManagerDto.ErrorManagerDto
	Message string
}
