package config

const (
	// ログ関連
	REQUEST_ID        string = "REQUEST_ID"
	REQUEST_ID_HEADER string = "X-Request-Id"
	USER_ID           string = "USER_ID"

	// DB関連
	DB_TYPE                 string = "postgres"
	DB_HOST                 string = "127.0.0.1"
	DB_PORT                 string = "5432"
	DB_USER                 string = "postgres"
	DB_PASSWORD             string = "password"
	DB_DATABASE_NAME        string = "testdb"
	DB_MAX_CONNECTION_NUM   int    = 10
	DB_MAX_IDLE_TIME_MINUTE int    = 5
)
