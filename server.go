package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
// 	"log"
)
/*
	测试服务器，附带session机制

*/
var store *sessions.CookieStore
func main() {

	newSession()
	router := gin.Default()
	// var num int

	router.GET("/", func(c *gin.Context) {
		// num++
		c.String(200, "sid= %s\nid= %s\n", c.Query("sid"), c.Query("id"))
		// log.Printf("\n[%d 次请求]\nsid= %s\nid= %s\n", num, c.Query("sid"), c.Query("id"))
	})
	router.GET("/admin", func(c *gin.Context) {
		c.String(200,"账户")
	})
	router.GET("/admin2", func(c *gin.Context) {
		c.String(200, "密码")
	})

	router.GET("/login", func(c *gin.Context) {
		user := c.Query("user")
		session, _ := store.Get(c.Request, "get_session")
		_, login := session.Values["name"].(string)
		if login {
			c.String(200, "已登录")
		}else {
			session.Options.MaxAge = 1 * 5
			session.Options.HttpOnly = true
			session.Values["name"] = user
			_ = sessions.Save(c.Request, c.Writer)

			c.String(200, "已获取cookie")

		}		
	})
	_ = router.Run(":8080")

}

func newSession() {

	store = sessions.NewCookieStore([]byte("test"))
}