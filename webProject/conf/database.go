package conf

import (
	"webProject/model"
)

func InitDb(connectionString string) {
	model.InitConnection(connectionString)
}
