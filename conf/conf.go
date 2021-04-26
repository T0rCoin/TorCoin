package conf

import (
	"fmt"
	"github.com/unknwon/goconfig"
	"os"
)

const(
	FilePath string = "E:\\Project\\TorCoin\\conf\\conf.ini"
	SingularPoint string = ""
)

func Init() map[string]string {

	xMap := make(map[string]string)

	xMap["ConfFilePath"] = FilePath

	return xMap

}

func Load() *goconfig.ConfigFile {

	cfg, err := goconfig.LoadConfigFile(FilePath)

	if err != nil{

		fmt.Printf("Error:CANT LOAD CONFIG FILE,%v\n",err)

		os.Exit(1)

	}

	return cfg
}

func Get(section string, key string) string{

	cfg, err := goconfig.LoadConfigFile(FilePath)

	if err != nil{

		fmt.Printf("Error:CANT LOAD CONFIG FILE,%v\n",err)

		os.Exit(1)

	}

	value, err := cfg.GetValue(section,key)

	if err != nil {

		fmt.Printf("Error:Load Config Error%v\n",err)

		os.Exit(1)

	}

	return value

}