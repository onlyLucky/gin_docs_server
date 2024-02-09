package res

import (
	"gin_docs_server/utils/valid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Code int

type Response struct {
	Code Code    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	List []T `json:"list"`
	Count int `json:"count"`
}

const (
	SUCCESS Code = 0
	ErrCode Code = -1 // 系统报错
	ValidCode Code = 1 // 参数检测报错
)

func OK(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: SUCCESS,
		Data: data,
		Msg: msg,
	})
}

func OKWithMsg(msg string, c *gin.Context){
	OK(map[string]any{},msg,c)
}

func OKWithData(data any, c *gin.Context){
	OK(data,"success",c)
}

func OKWithList[T any](list []T,count int, c *gin.Context){
	if len(list) == 0 {
		list = []T{}
	}
	OK(ListResponse[T]{
		List: list,
		Count: count,
	},"success",c)
}

func Fail(code Code,data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg: msg,
	})
}

func FailWithMsg(msg string, c *gin.Context){
	Fail(ErrCode,map[string]any{},msg,c)
}

func FailWithData(data any, c *gin.Context){
	Fail(ErrCode,data,"系统错误",c)
}

func FailWithError(err error, c *gin.Context){
	errorMsg := valid.Error(err)
	FailWithMsg(errorMsg, c)
}