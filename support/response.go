package support

import (
	"dpj-admin-api/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  int         `json:"status" `  // 业务状态码
	Message string      `json:"message" ` // 提示信息
	Data    interface{} `json:"data" `    // 任何数据
}

func NewResponse() *Response {
	return &Response{
		Status:  200,
		Message: "success",
		Data:    nil,
	}
}

// Wrapper 封装了gin.Context
type Wrapper struct {
	ctx *gin.Context
}

func WithContext(ctx *gin.Context) *Wrapper {
	return &Wrapper{ctx: ctx}
}

// Success 输出成功信息;data = json字符串
func (wrapper *Wrapper) Success(data interface{}) {
	resp := NewResponse()
	resp.Data = data
	wrapper.ctx.Header("Server", config.Get("app.name"))
	wrapper.ctx.Header("Server-Version", config.Get("app.version"))
	wrapper.ctx.JSON(http.StatusOK, resp)
}

// Error 输出错误信息
func (wrapper *Wrapper) Error(statusCode int, errMessage string) {
	resp := NewResponse()
	resp.Status = statusCode
	resp.Message = errMessage
	wrapper.ctx.Header("Server", config.Get("app.name"))
	wrapper.ctx.Header("Server-Version", config.Get("app.version"))
	wrapper.ctx.JSON(statusCode, resp)
}
