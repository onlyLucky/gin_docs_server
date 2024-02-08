package config

import "fmt"

type Mysql struct {
	Host         string `yaml:"host"`                                 // 服务器地址:端口
	Port         int    `yaml:"port"`                                 //端口
	Config       string `yaml:"config"`                               // 高级配置
	DB           string `yaml:"db"`                                   // 数据库名
	Username     string `yaml:"username"`                             // 数据库用户名
	Password     string `yaml:"password"`                             // 数据库密码
	LogLevel     string `yaml:"logLevel"`                             // 是否开启Gorm全局日志
}

// 用户名:密码啊@tcp(ip:端口)/数据库的名字?数据库配置
func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",m.Username,m.Password,m.Host,m.Port,m.DB,m.Config)
}

