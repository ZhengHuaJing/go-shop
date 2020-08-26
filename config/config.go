package config

import "time"

type App struct {
	PageSize int
}

type MD5 struct {
	Salt      string
	JwtSecret string
}

type Time struct {
	TimeFormat string
	DateFormat string
}

type Log struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
}

type File struct {
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string
}

type Casbin struct {
	CasbinConfPath  string
	ApiJsonFilePath string

	RoleAdmin      string
	RoleCommonUser string

	DefaultAdminUserName string
	DefaultAdminPassword string
}

type Captcha struct {
	ImgHeight int
	ImgWidth  int
	KeyLong   int
}

type Server struct {
	RunMode      string
	Host         string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Mysql struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	TestName    string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Email struct {
	HostName string
	Password string
	Host     string
	Port     int
}

type Alipay struct {
	AppID                string
	PrivateKey           string
	AppPublicCertPath    string
	AliPayPublicCertPath string
	AliPayRootCertPath   string
}

type Config struct {
	App     App
	MD5     MD5
	Time    Time
	Log     Log
	File    File
	Casbin  Casbin
	Server  Server
	Mysql   Mysql
	Redis   Redis
	Captcha Captcha
	Email   Email
	Alipay  Alipay
}
