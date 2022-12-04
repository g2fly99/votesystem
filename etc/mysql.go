package etc

type mySqlConfT struct {
	IpAddr   string `yaml:"addr"` // 127.0.0.1:13306
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Debug    bool   `yaml:"debug"`
}
