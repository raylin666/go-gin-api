package autoload

type Redis struct {
	Addr     string `yaml:"Addr"`
	Port     int    `yaml:"Port"`
	Password string `yaml:"Password"`
	Db       int    `yaml:"Db"`
}
