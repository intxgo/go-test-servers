package config

type ServerType string

const (
	Socket ServerType = "socket"
	Socks5 ServerType = "socks5"
	Ssl ServerType = "ssl-socket"
)

type ServerConfig struct {
	Type ServerType `yaml:"type"`
	Enabled bool `yaml:"enabled"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`

	// Socks5 specific
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`

	// Ssl specific
	Cert string `yaml:"certfile"`
	Key string `yaml:"keyfile"`
	Ca string `yaml:"cafile"`
}

type Config struct {
	Servers []ServerConfig `yaml:"servers"`
}