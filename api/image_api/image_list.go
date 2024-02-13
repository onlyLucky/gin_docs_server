package image_api

import (
	"fmt"
	"gin_docs_server/models"
	"gin_docs_server/service/common/list"
	"gin_docs_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type ImageListResponse struct {
	models.ImageModel
	WebPath string `json:"webPath"`
}

// @Tags 图片管理
// @Summary 上传列表
// @Description 上传列表
// @Param data query models.Pagination true "参数"
// @Param token header string true "token"
// @Router /api/uploadAvatarList [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[ImageListResponse]}
func (ImageApi) ImageListView(c *gin.Context) {
	var cr models.Pagination
	c.ShouldBindQuery(&cr)

	_list, count, _ := list.QueryList[models.ImageModel](models.ImageModel{}, list.Option{
		Pagination: cr,
		Likes:      []string{"fileName"},
	})
	fmt.Println(_list)
	var imageList = make([]ImageListResponse, 0)
	for _, model := range _list {
		imageList = append(imageList, ImageListResponse{
			ImageModel: model,
			WebPath:    model.WebPath(),
		})
	}
	res.OKWithList(imageList, count, c)
}
