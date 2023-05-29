package settings

type config struct {
	Application struct {
		Host         string
		Port         string
		ApiVersion   string
		TimeOut      int
		Secret       string
		SecretExpire int
		PageSize     int
		PageMaxSize  int
	}
	Mysql struct {
		DbType   string
		Host     string
		DbName   string
		Port     string
		UserName string
		Password string
		CharSet  string
	}
	Gorm struct {
		LogLevel    string
		MaxOpenConn int
		MaxIdleConn int
	}
}
