package setting

type Config struct {
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int64  `mapstructure:"MaxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int64  `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	LogLevel    string `mapstucture:"log_level"`
	FileLogName string `mapstucture:"file_log_name"`
	MaxBackups  int    `mapstucture:"max_backups"`
	MaxAge      int    `mapstucture:"max_age"`
	MaxSize     int    `mapstucture:"max_size"`
	Compress    bool   `mapstucture:"compress"`
}
