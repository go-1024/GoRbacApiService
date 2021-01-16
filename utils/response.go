package utils

import (
	"github.com/gin-gonic/gin"
)
// 成功返回封装 参数 data interface{} 类型为可接受任意类型
func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "请求成功",
	})
}
// 错误返回封装
func SuccessErr(c *gin.Context, errCode int, msg interface{}){
	c.JSON(500, gin.H{
		"code": errCode,
		"msg":  msg,
	})
}