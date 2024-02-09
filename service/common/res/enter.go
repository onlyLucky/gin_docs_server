package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	List []T `json:"list"`
	Count int `json:"count"`
}

const (
	SUCCESS = 0
	FAIL = -1
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

func Fail(code int,data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg: msg,
	})
}

func FailWithMsg(msg string, c *gin.Context){
	Fail(FAIL,map[string]any{},msg,c)
}

func FailWithData(data any, c *gin.Context){
	Fail(FAIL,data,"系统错误",c)
}

func FailWithError(err error, c *gin.Context){
	
}