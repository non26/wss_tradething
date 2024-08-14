package wssconfig

import viper "github.com/spf13/viper"

func ReadConfig() (c *AppConfig, err error) {
	viper.New().ConfigFileUsed()
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./wssconfig")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.Unmarshal(&c)

	return c, nil
}
