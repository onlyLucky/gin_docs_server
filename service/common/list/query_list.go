package list

import (
	"fmt"
	"gin_docs_server/global"
	"gin_docs_server/models"

	"gorm.io/gorm"
)

// 分页统一处理

type Option struct {
	models.Pagination
	Likes   []string // 模糊匹配的字段
	Debug   bool     //是否debug
	Where   *gorm.DB // 精细化条件查询
	Preload []string // 预加载
}

func QueryList[T any](model T, option Option) (list []T, count int, err error) {
	query := global.DB
	if option.Debug {
		query = query.Debug()
	}

	if option.Sort == "" {
		// 按时间排序
		option.Sort = "createdAt DESC"
	}

	if option.Limit <= 0 {
		option.Limit = 10
	}
	if option.Page <= 0 {
		option.Page = 1
	}

	if option.Where != nil {
		query.Where(option.Where)
	}

	if option.Key != "" {
		likeQuery := global.DB.Where("")
		for index, column := range option.Likes {
			if index == 0 {
				likeQuery.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			} else {
				likeQuery.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			}
		}
		if len(option.Likes) > 0 {
			query = query.Where(likeQuery)
		}
	}

	// 获取总数
	count = int(query.Find(&list).RowsAffected)

	// 预加载数据
	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}

	//分页处理
	offset := (option.Page - 1) * option.Limit
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return

	// 这里如果返回的参数已经在函数里面声明了，函数内部有赋值，底部要加一个return
}
