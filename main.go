package main

import (
	"fmt"
	"gclass/oss"
	"gclass/utils"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"time"
)

func main() {
	crontab := cron.New()
	task := func() {
		// 取该路径下的第一个文件
		filePath := utils.ListAll("/mnt/backup/database/shop"+time.Now().Format("20060102"), 0)
		oss.Up(filePath, false)
	}
	// 每天1点执行task任务
	err := crontab.AddFunc("0 0 1 * * ?", task)
	if err != nil {
		fmt.Println(err)
	}
	crontab.Start()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/up", func(c *gin.Context) {
		// 本地文件路径,如:/www/backup/db/20210326105022.sql
		filePath := c.DefaultQuery("filePath", "")
		// 上传成功后,是否删除本地文件,默认不删除
		isDel := c.DefaultQuery("isDel", "-1")
		var isDelete = false
		if isDel != "-1" {
			isDelete = true
		}
		if filePath != "" {
			go oss.Up(filePath, isDelete)
		}
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	r.Run(":9996")
}
