package utils

import "github.com/spf13/viper"

type ENVConfig struct {
	Enviroment   string `mapstructure:"ENVIROMENT"`
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBSource     string `mapstructure:"DB_SOURCE"`
	SVAddress    string `mapstructure:"SV_ADDRESS"`
	MigrationDir string `mapstructure:"MIGRATION_DIR"`
	SecretCode   string `mapstructure:"SECRET_CODE"`
}

func LoadConfig() (config ENVConfig, err error) {
	//load config directly from constant
	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	//check if viper can't find/read the config file
	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}
