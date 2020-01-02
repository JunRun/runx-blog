/**
 *
 * @Description: 读取配置文件工具
 * @Version: 1.0.0
 * @Date: 2020-01-02 10:51
 */
package util

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"strings"
)

type Conf struct {
	Data map[interface{}]interface{}
}

var (
	Config *Conf
)

func init() {
	file, err := ioutil.ReadFile("./conf/config.yaml")
	if err != nil {
		fmt.Println("配置文件读取失败!")
		panic(err)
	}
	Config = &Conf{}
	err = yaml.Unmarshal(file, Config.Data)
	if err != nil {
		fmt.Println("解析配置文件失败")
		panic(err)
	}
}

func (c *Conf) Get(key string) interface{} {
	path := strings.Split(key, ".")
	data := c.Data
	for key, value := range path {
		v, ok := data[value]
		if !ok {
			break
		}
		if (key + 1) == len(path) {
			return v
		}
		//判断下一级是否是一个map 否 返回数据，是 继续向下搜索
		if data, ok = v.(map[interface{}]interface{}); !ok {
			return v
		}
	}
	return nil
}

// 从配置文件中获取string类型的值
func (c *Conf) GetString(param string) string {
	value := c.Get(param)
	switch value := value.(type) {
	case string:
		return value
	case bool, float64, int:
		return fmt.Sprint(value)
	default:
		return ""
	}
}

// 从配置文件中获取int类型的值
func (c *Conf) GetInt(param string) int {
	value := c.Get(param)
	switch value := value.(type) {
	case string:
		i, _ := strconv.Atoi(value)
		return i
	case int:
		return value
	case bool:
		if value {
			return 1
		}
		return 0
	case float64:
		return int(value)
	default:
		return 0
	}
}
