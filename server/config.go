package server

type Config struct {
	ServerPort   uint16
	GCPProjectId string
}

func LoadConfig() Config {
	cfg := Config{
		ServerPort:   3000,
		GCPProjectId: "mailto-practice-project",
	}
	return cfg
}
