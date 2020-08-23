package main

import (
	"github.com/gin-gonic/gin"
)

/* 测试服务器 */

func main() {

	router := gin.Default()
	var num int
	router.GET("/", func(c *gin.Context) {

		num++
		c.String(200, "sid= %s\nid= %s\n", c.Query("sid"), c.Query("id"))
		// log.Printf("\n[%d 次请求]\nsid= %s\nid= %s\n", num, c.Query("sid"), c.Query("id"))
	})
	router.GET("/admin", func(c *gin.Context) {
		c.String(200,"账户")
	})
	_ = router.Run(":8080")

}
