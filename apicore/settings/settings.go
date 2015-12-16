package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var environments = map[string]string{
	"production":    "src/github.com/mindaugasrmn/goproject/apicore/settings/prod.json",
	"preproduction": "src/github.com/mindaugasrmn/goproject/apicore/settings/pre.json",
	"tests":         "src/github.com/mindaugasrmn/goproject/apicore/settings/tests.json",
}

var confFile = "src/github.com/mindaugasrmn/goproject/apicore/settings/config.json"
var env = "preproduction"

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

type Config struct {
	Mongo_svr string
	Mongo_db  string
	Redis_srv string
	Redis_pwd string
	Env       string
}

var settings Settings = Settings{}
var config Config = Config{}


func Init() {
	var conf = ReadConfig()
	var env = conf.Env
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading cert file", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

func ReadConfig()  Config {
	file, err := ioutil.ReadFile(confFile)
	if err != nil {
		fmt.Println("Error while reading config.json file", err)
	}
    config = Config{}
	jsonErr := json.Unmarshal(file, &config)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
    return config
}

