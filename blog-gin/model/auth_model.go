/**
 *
 * @Description: 权限管理模型
 * @Version: 1.0.0
 * @Date: 2020-01-08 16:11
 */
package model

import gormadapter "github.com/casbin/gorm-adapter"

type CasbinRules struct {
}

func init() {
	gormadapter.NewAdapterByDB(GetConnect())
}
