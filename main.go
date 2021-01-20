package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nailcui/exman/resource"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/man/:name", func(c *gin.Context) {
		name := c.Param("name")
		commands := resource.Commands[name]
		if commands == nil {
			c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(fmt.Sprintf("not found: %s\n", name)))
		} else {
			result := "\n"
			// 格式化结果如下:
			// # desc
			// $ cmd
			for _, command := range commands {
				result = result + "# " + command.Desc + "\n$ " + command.Cmd + "\n\n"
			}
			result = result + "\n"
			c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(fmt.Sprintf("%s", result)))
		}
	})
	err := r.Run(":10010")
	if err != nil {
		log.Panicln(err)
	}
}
