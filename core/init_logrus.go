package core

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogRequest struct {
	LogPath string //日志目录
	AppName string // app的名字
	NoDate bool // 是否不需要按时间分割
	NoErr bool // 是否不单独存放error日志
	NoGlobal bool //是否不替换全局logrus
}

type LogFormatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// 日志按照时间分级储存本地hook
type DateHook struct{
	file     *os.File
  fileDate string //判断日期切换目录
	logPath string
	appName string
}
// 哪些等级的日志才会生效
func (DateHook) Levels() []logrus.Level{
	return logrus.AllLevels
}
func (hook DateHook) Fire(entry *logrus.Entry) error{
	timer := entry.Time.Format("2006-01-02")
  line, _ := entry.String()
  if hook.fileDate == timer {
    hook.file.Write([]byte(line))
    return nil
  }
  // 时间不等
  hook.file.Close()
  os.MkdirAll(path.Join(hook.logPath,timer), os.ModePerm)
  filename := path.Join(hook.logPath,timer,hook.appName+".log")

  hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
  hook.fileDate = timer
  hook.file.Write([]byte(line))
  return nil
}

// 日志按照时间分级储存本地error hook
type ErrorHook struct{
	file     *os.File
  fileDate string //判断日期切换目录
	logPath string
	appName string
}
// 哪些等级的日志才会生效
func (ErrorHook) Levels() []logrus.Level{
	return []logrus.Level{logrus.ErrorLevel}
}
func (hook ErrorHook) Fire(entry *logrus.Entry) error{
	timer := entry.Time.Format("2006-01-02")
  line, _ := entry.String()
  if hook.fileDate == timer {
    hook.file.Write([]byte(line))
    return nil
  }
  // 时间不等
  hook.file.Close()
  os.MkdirAll(path.Join(hook.logPath,timer), os.ModePerm)
  filename := path.Join(hook.logPath,timer,"err.log")

  hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
  hook.fileDate = timer
  hook.file.Write([]byte(line))
  return nil
}

// 可以不传参
func InitLogger(requestList ...LogRequest) *logrus.Logger {
	var request LogRequest
	if len(requestList)>0{
		request = requestList[0]
	}
	if request.LogPath == ""{
		request.LogPath = "logs"
	}
	if request.AppName == ""{
		request.AppName = "gvd"
	}
	mLog := logrus.New()               //新建一个实例
	mLog.SetOutput(os.Stdout)          //设置输出类型
	mLog.SetReportCaller(true)         //开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	mLog.SetLevel(logrus.DebugLevel)   //设置最低的Level
	if !request.NoDate{
		mLog.AddHook(&DateHook{
			logPath: request.LogPath,
			appName: request.AppName,
		})
	}
	if !request.NoErr{
		mLog.AddHook(&ErrorHook{
			logPath: request.LogPath,
			appName: request.AppName,
		})
	}
	if !request.NoGlobal{
		InitDefaultLogger()
	}
	return mLog
}

func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)          //设置输出类型
	logrus.SetReportCaller(true)         //开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	logrus.SetLevel(logrus.DebugLevel)   //设置最低的Level
}
