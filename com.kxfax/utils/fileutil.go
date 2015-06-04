package utils

import (
	"io/ioutil"
	"os"
)

func Read(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(path+"文件未找到，请检查")
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	//    fmt.Println(string(fd))
	return string(fd)
}
