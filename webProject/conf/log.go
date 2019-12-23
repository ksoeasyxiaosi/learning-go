package conf

import (
	"webProject/utils"
)

func InitLog(ErrorLevel string) {

	utils.SetErrorLevel(ErrorLevel)
}
