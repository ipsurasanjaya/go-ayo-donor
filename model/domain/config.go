package domain

type Config struct {
	Host               string
	Port               string
	Schema             string
	DBName             string
	User               string
	Password           string
	MaxOpenConnections int
	MaxIdleConnections int
	SSLMode            string
	TimeZone           string
}
