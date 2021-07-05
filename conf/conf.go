package conf

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-redis/redis/v7"
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

type Redis struct {
	RedisClient *redis.Client
}

var RedisClient *Redis

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

func RedisInit() {
	dsn := os.Getenv("REDIS_DSN")

	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}
	RedisClient = &Redis{RedisClient: client}
}

func GetRedis() *Redis {

	return RedisClient
}
