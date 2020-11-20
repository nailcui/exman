package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var commands = make(map[string]string)

func init() {
	commands["grep"] = "grep -o \"1\\d\\d\"	只打印匹配部分"
	commands["zip"] = "zip -qr dist.zip dist	将一个文件夹压缩"
	commands["git"] = "git checkout -b dev_0.1 origin/dev_0.1	将远程分支检出一个本地分支"
}

func main() {
	r := gin.Default()
	r.GET("/man/:name", func(c *gin.Context) {
		name := c.Param("name")
		content := commands[name]
		if len(content) != 0 {
			c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("%s\n", content)))
			return
		}
		c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("not found: %s\n", name)))
	})
	err := r.Run(":10010")
	if err != nil {
		log.Panicln(err)
	}
}
