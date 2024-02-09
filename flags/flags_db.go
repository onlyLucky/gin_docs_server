package flags

import (
	"gin_docs_server/global"
	"gin_docs_server/models"

	"github.com/sirupsen/logrus"
)

// 初始化数据库表结构
func DB(){
	// 创建表时添加后缀
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.RoleModel{},
		&models.DocModel{},
		&models.UserCollDocModel{}, // 自定义连接表要放在两个连接表的后面。
		&models.RoleDocModel{},
		&models.ImageModel{},
		&models.UserPwdDocModel{},
		&models.LoginModel{},
		&models.DocDataModel{},
	)

	if err != nil {
		logrus.Fatalf("数据库迁移失败 err:%s", err.Error())
	}
	logrus.Infof("数据库迁移成功！！！")
}