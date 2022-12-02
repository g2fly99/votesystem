package etc

type mySqlConfT struct {
	ipAddr   string `yaml:"addr"` // 127.0.0.1:13306
	dbName   string `yaml:"dbName"`
	userName string `yaml:"user"`
	password string `yaml:"password"`
	debug    bool   `yaml:"debug"`
}

func (this mySqlConfT) UserName() string {
	return this.userName
}

func (this mySqlConfT) Password() string {
	return this.password
}

func (this mySqlConfT) ServerAddr() string {
	return this.ipAddr
}

func (this mySqlConfT) DbName() string {
	return this.dbName
}

func (this mySqlConfT) Debug() bool {
	return this.debug
}
