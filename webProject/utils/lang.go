package utils

import "fmt"

func Trans(key string, params ...interface{}) string {
	return fmt.Sprintf(t(key), params)
}
