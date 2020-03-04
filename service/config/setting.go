package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type App struct {
	JwtSecret string
	// PageSize  int
	// PrefixUrl string

	// RuntimeRootPath string

	// ImageSavePath  string
	// ImageMaxSize   int
	// ImageAllowExts []string

	// ExportSavePath string
	// QrCodeSavePath string
	// FontSavePath   string

	// LogSavePath string
	// LogSaveName string
	// LogFileExt  string
	// TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	// RunMode      string
	Host         string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	// cors
	AllowedOrigins []string
	AllowedHeaders []string
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	// TablePrefix string
	Encoding string
}

var DbReadSetting = &Database{}
var DbWriteSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *viper.Viper

// Setup initialize the configuration instance
func Setup() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfg = viper.New()
	cfg.AddConfigPath(path + "/conf")
	cfg.SetConfigName("config")
	cfg.SetConfigType("yaml")
	if err := cfg.ReadInConfig(); err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/config.yaml': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database.read", DbReadSetting)
	mapTo("database.write", DbWriteSetting)
	mapTo("redis", RedisSetting)

	// AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.UnmarshalKey(section, v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
