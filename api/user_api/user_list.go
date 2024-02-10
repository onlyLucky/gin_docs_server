package user_api

import (
	"gin_docs_server/global"
	"gin_docs_server/models"
	"gin_docs_server/service/common/list"
	"gin_docs_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
	Key   string `json:"key" form:"key"`
}

func (UserApi) UserListView(c *gin.Context) {
	// var cr UserListRequest
	var cr models.Pagination
	c.ShouldBindQuery(&cr)

	_list, count, err := list.QueryList[models.UserModel](models.UserModel{}, list.Option{
		Pagination: cr,
		Likes:      []string{"nickName", "userName"},
		Preload:    []string{"RoleModel"},
		Debug:      true,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("获取用户列表失败", c)
		return
	}
	res.OKWithList(_list, count, c)

	/* var users []models.UserModel
	if cr.Limit <= 0 {
		cr.Limit = 10
	}

	if cr.Page <= 0 {
		cr.Page = 1
	}

	var count int //总数
	offset := (cr.Page - 1) * cr.Limit

	// query := global.DB.Where(fmt.Sprintf("nickName like '%%%s%%'", cr.Key)) // 这个写法很容易产生sql注入的 操作 key: ' or 1=1--用户
	query := global.DB.Where("nickName like ?",fmt.Sprintf("%%%s%%", cr.Key))
	global.DB.Model(models.UserModel{}).Where(query).Select("count(id)").Scan(&count)
	// global.DB.Debug().Limit(cr.Limit).Offset(offset).Find(&users)
	global.DB.Debug().Where(query).Limit(cr.Limit).Offset(offset).Find(&users)
	res.OKWithList(users,count, c) */

	return
}
