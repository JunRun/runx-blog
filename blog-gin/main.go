/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-20 11:32
 */
package main

import (
	"fmt"

	_ "github.com/JunRun/blog-gin/conf"
	_ "github.com/JunRun/blog-gin/model"
	"github.com/JunRun/blog-gin/router"
	_ "github.com/JunRun/blog-gin/utils"
)

func main() {
	fmt.Println("Start----")
	_ = router.InitRouter()

}
