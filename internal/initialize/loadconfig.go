package initialize

import (
	"fmt"

	"github.com/huyle/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func loadConfig() {
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
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		fmt.Println(fmt.Errorf("fail to unmarshal config: %w \v", err))
	}
}
