package conf

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DbConnect DbConnect
}

type DbConnect struct {
	Port         string
	Host         string
	User         string
	Password     string
	DatabaseName string
	Sslmode      string
	TimeZone     string
}

func getFileName(version_name string) string {

	config_file := []string{"../config.", version_name, ".json"}

	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(config_file, ""))
	return filePath
}

func GetConfig() Configuration {
	var version_name = flag.String("mode", "dev", "select web-app mode")

	configuration := Configuration{}

	flag.Parse()

	err := gonfig.GetConf(getFileName(*version_name), &configuration)

	if err != nil {
		fmt.Println(err)
		os.Exit(500)
	}
	return configuration

}
