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
		command := resource.Commands[name]
		if command.Examples == nil {
			if resource.Tags[name] != nil {
				returnExamples(c, resource.Tags[name])
				return
			}
			c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(fmt.Sprintf("not found: %s\n", name)))
		} else {
			returnExamples(c, command.Examples)
		}
	})
	err := r.Run(":8082")
	if err != nil {
		log.Panicln(err)
	}
}

func returnExamples(c *gin.Context, examples []resource.Example) {
	result := "\n"
	// 格式化结果如下:
	// # desc
	// $ cmd
	for _, example := range examples {
		result = result + "# " + example.Desc + "\n$ " + example.Cmd + "\n\n"
	}
	result = result + "\n"
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(fmt.Sprintf("%s", result)))
}
