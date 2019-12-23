package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	Path "path"
	"strings"
	"webProject/utils"
)

func InitLang() error {

	path, e := os.Getwd()
	fmt.Println(path)
	if e != nil {
		panic(e)
	}
	dirList, e := ioutil.ReadDir(path + "/lang/")
	if e != nil {
		panic(e)
	}

	paths := make(map[string]string)

	for _, dirInfo := range dirList {
		langDirList, e := ioutil.ReadDir(path + "/lang/" + dirInfo.Name())
		if e != nil {
			panic(e)
		}
		for _, fileInfo := range langDirList {
			paths[dirInfo.Name()+"."+strings.TrimSuffix(fileInfo.Name(), Path.Ext(fileInfo.Name()))] = path + "/lang/" + dirInfo.Name() + "/" + fileInfo.Name()
		}
	}

	utils.LoadLang(paths)
	return nil
}
