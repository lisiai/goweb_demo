package settings

var (
	AppC   AppConfig
	MysqlC MysqlConfig
	RedisC RedisConfig
	LogC   LogConfig
)

type AppConfig struct {
	ContextPath string `mapstructure:"context_path"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Mode        string `mapstructure:"mode"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Dbname   int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"fileName"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Path       string `mapstructure:"path"`
}
