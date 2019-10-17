package setting

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"time"
)

type App struct {
	JwtSecret string
	PageSize  int `yaml:"page_size"`
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table_prefix"`
}

type Redis struct {
	Host        string        `yaml:"host"`
	Password    string        `yaml:"password"`
	MaxIdle     int           `yaml:"max_idle"`
	MaxActive   int           `yaml:"max_active"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

var Config = struct {
	App      App
	Server   Server
	Database Database
	Redis    Redis
}{}

// Setup initialize the configuration instance
func Setup() {
	YamlLoadFromPath("config/config.yml", &Config)

	//AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	//ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	//ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	//RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

}

//YamlLoadFromPath load from local file
func YamlLoadFromPath(path string, t interface{}) error {

	b, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(b, t); err != nil {
		return err
	}

	return nil
}
