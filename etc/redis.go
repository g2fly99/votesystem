package etc

type redisConfT struct {
	IpAddr string `yaml:"addr"` // 127.0.0.1:13306
	Passwd string `yaml:"password"`
}

func (this redisConfT) ServerAddr() string {
	return this.IpAddr
}

func (this redisConfT) Password() string {
	return this.Passwd
}
