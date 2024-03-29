package flags

import "flag"

type Option struct {
	DB   bool   // 初始化数据库
	Port int    // 更换端口
	Load string // 导入数据库文件

}

// 解析命令行参数
func Parse() (option *Option) {
	option = new(Option)
	flag.BoolVar(&option.DB,"db",false,"初始化数据库")
	flag.IntVar(&option.Port,"port",0,"程序运行端口")
	flag.StringVar(&option.Load,"load","","导入sql数据库")
	flag.Parse()
	return option
}
// 根据不同的参数运行不同的结果
func (option Option) Run() bool{
	if option.DB{
		DB()
		return true
	}
	if option.Port != 0{
		Port(option.Port)
		return false
	}
	if option.Load !=""{
		Load()
		return true
	}
	return false
}