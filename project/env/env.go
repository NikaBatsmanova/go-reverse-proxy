package env

import (
	"go-reverse-proxy/project/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func Init() {
	var t types.T
	p, _ := filepath.Abs("./project/configs/config.yaml")
	conf, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(conf, &t)
	if err != nil {
		panic(err.Error())
	}
	os.Setenv("backend", t.Back)
	os.Setenv("port", strconv.Itoa(t.Port))
	os.Setenv("cache", strconv.Itoa(t.CacheSize))
}
