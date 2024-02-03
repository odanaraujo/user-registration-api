package env

import (
	"github.com/spf13/viper"
)

var Env *config

/*
GO_ENV: Determina o ambiente, pode ser production, development, stage.
GO_PORT: Determina a porta que vamos usar para receber requisições, vamos usar a porta :8080,
DATABASE_URL: Aqui fica a url de conexão com o banco de dados.
*/
type config struct {
	GoEnv       string `mapstructure:"GO_ENV"`
	GoPort      string `mapstructure:"GO_PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func LoadingConfig(path string) (*config, error) {
	viper.SetConfigFile("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&Env); err != nil {
		return nil, err
	}

	return Env, nil
}
