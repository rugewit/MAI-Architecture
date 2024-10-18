package additional

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadViper(path string) error {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Could not read config", err)
		return err
	}
	return nil
}
