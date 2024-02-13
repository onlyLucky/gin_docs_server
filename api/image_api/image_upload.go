package image_api

import (
	"fmt"
	"gin_docs_server/global"
	"gin_docs_server/models"
	"gin_docs_server/service/common/res"
	"gin_docs_server/utils/hash"
	"gin_docs_server/utils/jwts"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var ImageWhiteList = []string{
	"jpg",
	"png",
	"jpeg",
	"gif",
	"svg",
	"webp",
}

// @Tags 图片管理
// @Summary 上传头像
// @Description 上传头像
// @Param image formData file  true "文件上传"
// @Param token header string true "token"
// @Router /api/uploadAvatar [post]
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} res.Response{}
func (ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		res.FailWithMsg("图片参数错误", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)

	savePath := path.Join("uploads", claims.NickName, fileHeader.Filename)

	// 白名单判断
	if !InImageWhiteList(fileHeader.Filename, ImageWhiteList) {
		res.FailWithMsg("文件非法", c)
		return
	}
	// 文件大小判断 2MB
	if fileHeader.Size > int64(10*1024*1024) {
		res.FailWithMsg("文件过大", c)
		return
	}
	// 重复文件判断
	file, _ := fileHeader.Open()
	fileHash := hash.FileMd5(file)
	fmt.Println(fileHash)

	var imageModel models.ImageModel
	err = global.DB.Take(&imageModel, "hash = ?", fileHash).Error

	if err != nil {
		// 没有查询到文件hash

		// 判断数据库有没有重名的文件
		var count int64
		global.DB.Model(models.ImageModel{}).Where("path = ?", savePath).Count(&count)
		if count > 0 {
			// 重名存在的话，更改文件名称
			fileHeader.Filename = ReplaceFileName(fileHeader.Filename)
			savePath = path.Join("uploads", claims.NickName, fileHeader.Filename)
		}
		err = c.SaveUploadedFile(fileHeader, savePath)
		if err != nil {
			global.Log.Errorf("%s 文件保存错误 %s", savePath, err)
			res.FailWithMsg("图片上传失败", c)
			return
		}
	} else {
		// 查询到，修复path
		savePath = imageModel.Path
	}

	// 上传图片进行入库
	imageModel = models.ImageModel{
		UserID:   claims.UserID,
		FileName: fileHeader.Filename,
		Size:     fileHeader.Size,
		Path:     savePath,
		Hash:     fileHash,
	}
	err = global.DB.Create(&imageModel).Error
	if err != nil {
		global.Log.Errorln(err)
		res.FailWithMsg("图片上传失败", c)
		return
	}

	res.OK(imageModel.WebPath(), "图片上传成功", c)
	return
}

// 判断一个图片是否在白名单中
func InImageWhiteList(fileName string, whiteList []string) bool {
	// 截取文件后缀
	_list := strings.Split(fileName, ".") //
	if len(_list) < 2 {
		return false
	}
	suffix := strings.ToLower(_list[len(_list)-1])
	for _, s := range whiteList {
		if suffix == s {
			return true
		}
	}
	return false
}

func ReplaceFileName(oldFileName string) string {
	_list := strings.Split(oldFileName, ".")
	lastIndex := len(_list) - 2
	var newList []string
	for i, s := range _list {
		if i == lastIndex {
			newList = append(newList, fmt.Sprintf("%s_%d", s, time.Now().Unix()))
			continue
		}
		newList = append(newList, s)
	}
	return strings.Join(newList, ".")
}
