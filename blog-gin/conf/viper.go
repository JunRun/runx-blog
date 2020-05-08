/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/5/7 8:30 下午
 */
package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.AddConfigPath("./conf")
	Config.SetConfigName("config")
	Config.SetConfigType("yml")
	err := Config.ReadInConfig()
	if err != nil {
		fmt.Print("读取配置文件出错 ", err)
	}

}
