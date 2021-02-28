package autoload

type Database struct {
	Driver      string `yaml:"Driver"`
	DbName      string `yaml:"DbName"`
	Host        string `yaml:"Host"`
	UserName    string `yaml:"UserName"`
	Password    string `yaml:"Password"`
	Charset     string `yaml:"Charset"`
	Port        int    `yaml:"Port"`
	Prefix      string `yaml:"Prefix"`
	MaxIdleConn int    `yaml:"MaxIdleConn"`
	MaxOpenConn int    `yaml:"MaxOpenConn"`
}
