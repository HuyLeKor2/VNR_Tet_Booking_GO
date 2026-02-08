package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DBName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
	Security struct {
		JWT struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config files
	viper.SetConfigName("local")     // name of config file (without extension)
	viper.SetConfigType("yaml")      //

	//read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to read config: %w \n", err))
	}
	//read server config
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Port::", viper.GetString("security.jwt.key"))

	//config struct
	var config config
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(fmt.Errorf("fail to unmarshal config: %w \v", err))
	}

	fmt.Println("Config Port::", config.Server.Port)
}
