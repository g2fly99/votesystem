package etc

type redisConfT struct {
	ipAddr   string `yaml:"addr"` // 127.0.0.1:13306
	password string `yaml:"password"`
}

func (this redisConfT) ServerAddr() string {
	return this.ipAddr
}

func (this redisConfT) Password() string {
	return this.password
}
