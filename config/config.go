package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"

	"GOGOGO/utils"
	"github.com/spf13/viper"
)

var jsonData map[string]interface{}

type config struct {
	Database dBConfig
	Redis    redisConfig
	Mongodb  mongoConfig
	Statsd   statsdConfig
	Server   serverConfig
	Chrome   chrome
	Rabbitmq rabbitmq
}

type rabbitmq struct {
	Host     string
	User     string
	Password string
}

type chrome struct {
	Headless bool
	Env      string
}

var SystemConfig config

func initJSON() {

	viper.SetConfigType("toml")

	if _, err := ioutil.ReadFile("./IS_DEV"); err == nil {
		viper.SetConfigName("config-dev") // 设置配置文件名 (不带后缀)
	} else if _, err = ioutil.ReadFile("./IS_TEST"); err == nil {
		viper.SetConfigName("config-test") // 设置配置文件名 (不带后缀)
	} else {
		viper.SetConfigName("config") // 设置配置文件名 (不带后缀)
	}
	viper.AddConfigPath(".")    // 第一个搜索路径
	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&SystemConfig) // 将配置信息绑定到结构体上
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
func initJSON2() {

	//原来的
	//bytes, err := ioutil.ReadFile("./config.json") // 原来的
	var bytes []byte
	var err error
	if _, er := ioutil.ReadFile("./IS_DEV"); er == nil { // todo 用IS_DEV文件是否存在判断是否开发模式
		bytes, err = ioutil.ReadFile("./config.dev.json")
	} else {
		bytes, err = ioutil.ReadFile("./config.json")
	}

	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")

	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}

type dBConfig struct {
	Dialect      string
	Database     string
	User         string
	Password     string
	Host         string
	Port         int
	Charset      string
	URL          string
	MaxIdleConns int
	MaxOpenConns int
}

// DBConfig 数据库相关配置
var DBConfig dBConfig

func initDB() {
	//utils.SetStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
	DBConfig = SystemConfig.Database
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
	DBConfig.URL = url
}

type redisConfig struct {
	Host      string
	Port      int
	Password  string
	URL       string
	MaxIdle   int
	MaxActive int
}

// RedisConfig redis相关配置
var RedisConfig redisConfig

func initRedis() {
	//utils.SetStructByJSON(&RedisConfig, jsonData["redis"].(map[string]interface{}))
	RedisConfig = SystemConfig.Redis
	url := fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port)
	RedisConfig.URL = url
}

type mongoConfig struct {
	URL      string
	Database string
}

// MongoConfig mongodb相关配置
var MongoConfig mongoConfig

func initMongo() {
	//utils.SetStructByJSON(&MongoConfig, jsonData["mongodb"].(map[string]interface{}))
	MongoConfig = SystemConfig.Mongodb
}

type serverConfig struct {
	APIPoweredBy       string
	SiteName           string
	Host               string
	ImgHost            string
	Env                string
	LogDir             string
	LogFile            string
	APIPrefix          string
	UploadImgDir       string
	ImgPath            string
	WxAppImgPath       string
	MaxMultipartMemory int
	Port               int
	StatsEnabled       bool
	CronEnabled        bool
	TokenSecret        string
	TokenMaxAge        int
	PassSalt           string
	LuosimaoVerifyURL  string
	LuosimaoAPIKey     string
	CrawlerName        string
	MailUser           string //域名邮箱账号
	MailPass           string //域名邮箱密码
	MailHost           string //smtp邮箱域名
	MailPort           int    //smtp邮箱端口
	MailFrom           string //邮件来源
	Github             string
	BaiduPushLink      string
	SessionID          string
	JavaServer         string
	PythonServer       string
	Ttf                string
	TtfPdf             string
}

// ServerConfig 服务器相关配置
var ServerConfig serverConfig

func initServer() {
	//utils.SetStructByJSON(&ServerConfig, jsonData["go"].(map[string]interface{}))
	ServerConfig = SystemConfig.Server
	sep := string(os.PathSeparator)
	execPath, _ := os.Getwd()
	length := utf8.RuneCountInString(execPath)
	lastChar := execPath[length-1:]
	if lastChar != sep {
		execPath = execPath + sep
	}
	if ServerConfig.UploadImgDir == "" {
		//pathArr := []string{"website", "static", "upload", "img"}
		pathArr := []string{"upload", "img"}
		uploadImgDir := execPath + strings.Join(pathArr, sep)
		ServerConfig.UploadImgDir = uploadImgDir
	}

	ymdStr := utils.GetTodayYMD("-")

	if ServerConfig.LogDir == "" {
		ServerConfig.LogDir = execPath
	} else {
		length := utf8.RuneCountInString(ServerConfig.LogDir)
		lastChar := ServerConfig.LogDir[length-1:]
		if lastChar != sep {
			ServerConfig.LogDir = ServerConfig.LogDir + sep
		}
	}
	ServerConfig.LogFile = ServerConfig.LogDir + ymdStr + ".log"
}

type statsdConfig struct {
	URL    string
	Prefix string
}

// StatsDConfig statsd相关配置
var StatsDConfig statsdConfig

func initStatsd() {
	//utils.SetStructByJSON(&StatsDConfig, jsonData["statsd"].(map[string]interface{}))
	StatsDConfig = SystemConfig.Statsd
}

func init() {
	initJSON()
	initDB()
	initRedis()
	initMongo()
	initServer()
	initStatsd()
}
