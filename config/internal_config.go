package config

var (
	AppMode = Get("APP_MODE", "DEV")
	AppPort = Get("APP_PORT", "8080")

	// DATABASE
	DBHost     = GetRequired("DB_HOST")
	DBPort     = GetRequired("DB_PORT")
	DBUsername = GetRequired("DB_USER")
	DBPassword = GetRequired("DB_PASS")
	DBName     = GetRequired("DB_NAME")
	DBTimeZone = GetRequired("DB_TIMEZONE")

	// DATABASE CONNECTION POOL
	DBMaxIdleConn     = Get("DB_MAX_IDLE_CONNS", "10")
	DBMaxOpenConn     = Get("DB_MAX_OPEN_CONNS", "20")
	DBConnMaxIdleTime = Get("DB_CONN_MAX_IDLETIME_IN_MINUTES", "60")
	DBConnMaxLifetime = Get("DB_CONN_MAX_LIFETIME_IN_MINUTES", "10")
)
