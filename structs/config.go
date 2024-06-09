package structs

type DbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SLLMode  string
	TimeZone string
}

type LogConfig struct {
	Name     string
	LogFile  string
	LogLevel string
}
