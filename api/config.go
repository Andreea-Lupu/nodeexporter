package api

type HTTPConfig struct {
	Host		string
	Port		string
}

type LogConfig struct {
	Level	string
	Output	string
}

type Config struct {
	HTTP	HTTPConfig
	Port	string
	Log		*LogConfig
}

func NewConfig() *Config {
	return &Config{
		HTTP:	HTTPConfig{Host: "127.0.0.1", Port: "5051"},
		Port:	"8000",
		Log:	&LogConfig{Level: "debug"},
	}
}
