package image_api

import (
	"fmt"
	"gin_docs_server/global"
	"gin_docs_server/models"
	"gin_docs_server/service/common/res"
	"os"

	"github.com/gin-gonic/gin"
)

// @Tags 用户管理
// @Summary 上传头像列表删除
// @Description 上传头像列表删除
// @Param data body models.IDListRequest true "参数"
// @Param token header string true "token"
// @Router /api/user [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImageApi) ImageRemoveView(c *gin.Context) {
	var cr models.IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var imageList []models.ImageModel
	global.DB.Find(&imageList, cr.IDList)
	if len(cr.IDList) != len(imageList) {
		res.FailWithMsg("数据一致性校验不通过", c)
		return
	}
	for _, model := range imageList {
		imageRemove(model)
	}
	res.OKWithData(fmt.Sprintf("批量删除成功，共删除%d个图片列表", len(cr.IDList)), c)
	return
}

// 删除图片发现多个hash,只删除记录
func imageRemove(image models.ImageModel) {
	var count int64
	global.DB.Model(models.ImageModel{}).Where("hash = ?", image.Hash).Count(&count)
	if count == 1 {
		err := os.Remove(image.Path)
		if err != nil {
			global.Log.Errorf("删除文件%s 错误 %s", image.Hash, err.Error())
		} else {
			global.Log.Infof("删除文件%s 成功", image.Hash)
		}
	}
	global.DB.Delete(&image)
}
