package user_api

import (
	"fmt"
	"gin_docs_server/global"
	"gin_docs_server/models"
	"gin_docs_server/service/common/res"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (UserApi) UserRemoveView(c *gin.Context) {
	var cr models.IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var userList []models.UserModel
	global.DB.Find(&userList, cr.IDList)
	if len(userList) != len(cr.IDList) {
		// 传值的id 里面有不存在的人员
		res.FailWithMsg("数据一致性校验不通过", c)
		return
	}
	// 这里可以返回成功删除，删除失败的map
	for _, model := range userList {
		err = UserRemoveService(model)
		if err != nil {
			logrus.Errorf("删除用户%s失败 err: %s", model.UserName, err.Error())
		} else {
			logrus.Infof("删除用户 %s 成功", model.UserName)
		}
	}
	res.OKWithData(fmt.Sprintf("批量删除成功，共删除%d个用户", len(cr.IDList)), c)
	return
}

func UserRemoveService(user models.UserModel) (err error) {

	// 添加事务进行删除操作
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// ImageModel 连带删除
		var imageList []models.ImageModel
		tx.Find(&imageList, "userId = ?", user.ID)
		if len(imageList) > 0 {
			err = tx.Delete(&imageList).Error
			if err != nil {
				return err
			}
		}
		// loginModel 不用连带删除

		// UserCollDocModel 连带删除
		var userCollList []models.UserCollDocModel
		tx.Find(&userCollList, "userId = ?", user.ID)
		if len(userCollList) > 0 {
			err = tx.Delete(&userCollList).Error
			if err != nil {
				return err
			}
		}
		//  UserPwdDocModel 连带删除
		var userPwdList []models.UserPwdDocModel
		tx.Find(&userPwdList, "userId = ?", user.ID)
		if len(userPwdList) > 0 {
			err = tx.Delete(&userPwdList).Error
			if err != nil {
				return err
			}
		}
		// 删除用户
		err = tx.Delete(&user).Error
		// 提交事务
		return err
	})

	return err
}
