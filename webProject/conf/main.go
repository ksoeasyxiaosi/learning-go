package conf

import (
	"github.com/joho/godotenv"
	"os"
)

func Init() {
	// 加载env全局配置
	godotenv.Load()
	if _, ok := os.LookupEnv("MYSQL"); !ok {
		panic("env文件未设置")
	}

	// 初始化语言包
	InitLang()

	// 初始化数据库连接
	InitDb(os.Getenv("MYSQL"))

	// 初始化cache
	InitCache(os.Getenv("CACHE_DB"))

	// 初始化日志级别
	InitLog(os.Getenv("LOG_LEVEL"))

}
