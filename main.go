package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Username string
	Password string

	Header http.Header

	N     string
	Type  string
	Ac_id string
	Enc   string
}

func (c Config) String() string {
	s, _ := yaml.Marshal(&c)
	return string(s)
}

var verbose bool
var config Config

func main() {
	args := make(map[string]int)
	for k, v := range os.Args {
		args[v] = k
	}

	if _, ok := args["-v"]; ok {
		verbose = true
		fmt.Println("verbose mode on.\n")
	}
	var dir string
	if k, ok := args["-c"]; ok {
		if len(os.Args) < k+2 {
			fmt.Println("-c 参数错误")
			os.Exit(1)
		}
		dir = os.Args[k+1]
	} else {
		dir = "./config.yml"
	}

	data, err := ioutil.ReadFile(dir)
	if err != nil {
		fmt.Println("打开配置文件失败，请检查配置文件是否存在:", dir)
		os.Exit(1)
	}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		fmt.Println("加载配置文件失败，请检查配置文件内容")
		os.Exit(1)
	}

	if verbose {
		fmt.Println("config:")
		fmt.Println(config)
	}
	login()
}
