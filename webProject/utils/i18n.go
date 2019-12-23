package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

// 语言包字典
type langList map[string]Dictinary

type Dictinary *map[interface{}]interface{}

var d = make(langList)

func LoadLang(paths map[string]string) error {
	for langName, path := range paths {
		m := make(map[interface{}]interface{})
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(data, &m)
		if err != nil {
			return err
		}

		d[langName] = &m
	}
	return nil
}

// 获得语言包的内容 用.符号连接层级
func t(key string) (langT string) {

	keys := strings.Split(key, ".")
	langT = key
	name := GetApp() + "." + GetLang()
	if lang := dictF(keys, d[name]); lang != "" {
		langT = lang
	}
	return
}

func dictF(index []string, in *map[interface{}]interface{}) string {

	dict := *in
	// 检查非法操作（传值进来为空字符）
	if len(index) == 0 {
		return ""
	}

	keyName := index[0]
	if len(index) == 1 {
		if langT, ok := dict[keyName].(string); ok {
			return langT
		}
		return ""
	}

	if len(index) > 1 {
		index = index[1:]
		if d, ok := dict[keyName].(map[interface{}]interface{}); ok {
			return dictF(index, &d)
		}
	}
	return ""
}
