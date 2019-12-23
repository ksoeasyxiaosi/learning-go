package utils

import "os"

const API = "api"

var app = API

var lang = os.Getenv("LANG")

var corsList = []string{"*"}

func SetAppToApi() {
	app = API
}

func GetApp() string {
	return app
}

func GetLang() string {
	if lang == "" {
		lang = "zh-cn"
	}
	return lang
}

func GetCorsList() []string {
	return corsList
}
