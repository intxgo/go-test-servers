package config


type Socks5Config struct {
	Enabled bool `yaml:"enabled"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`
}

type TcpSocketConfig struct {
	Enabled bool `yaml:"enabled"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type SslSocketConfig struct {
	Enabled bool `yaml:"enabled"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Cert string `yaml:"certfile"`
	Key string `yaml:"keyfile"`
	Ca string `yaml:"cafile"`
}

type Config struct {
	TcpSocket TcpSocketConfig `yaml:"socket"`
	Socks5 Socks5Config `yaml:"socks5"`
	SslSocket SslSocketConfig `yaml:"ssl-socket"`
}