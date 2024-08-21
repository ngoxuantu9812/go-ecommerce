package setting

type Config struct {
	Mysql MySQLSetting `mapstructure:"mysql"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    string `mapstructure:"MaxIdleConns"`
	MaxOpenConns    string `mapstructure:"maxOpenConns"`
	ConnMaxLifetime string `mapstructure:"connMaxLifetime"`
}
