package main

import (
	"fmt"
	"net/http"
    "io/ioutil"
    "os"
    "time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var savedText string

func main() {
    r := gin.Default()

    // 允许跨域
    r.Use(cors.Default())

    // 提供静态文件支持
    r.Static("/copy", "./static")

    // 保存文本的 API
    r.POST("/save", func(c *gin.Context) {
        var requestBody struct {
            Text string `json:"text"`
        }
        if err := c.BindJSON(&requestBody); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"message": "请求格式错误"})
            return
        }
        savedText = requestBody.Text
        // 获取当前时间
            currentTime := time.Now().Format("2006-01-02 15:04:05")

            // 以追加模式打开文件
            file, err := os.OpenFile("saved_text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                return
            }
            defer file.Close()

            // 追加时间和文本内容
            _, err = file.WriteString(fmt.Sprintf("[%s] %s\n", currentTime, requestBody.Text))
            if err != nil {
                return
            }

        data, err := ioutil.ReadFile("saved_text.txt")
        if err != nil {
            return
        }
    
        // 在控制台输出文件内容
        fmt.Println(string(data))
        // 返回保存成功的消息
        c.JSON(http.StatusOK, gin.H{"message": "保存成功"})
    })

    // 查看保存的文本的 API
    r.GET("/view", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"text": savedText})
    })
    // 监听并在 8080 端口运行
    r.Run(":30052")
}