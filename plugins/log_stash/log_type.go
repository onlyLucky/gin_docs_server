package log_stash

import "encoding/json"

type LogType int

const (
	LoginType   LogType = 1
	ActionType  LogType = 2
	RunTimeType LogType = 3
)

// 转字符串
func (t LogType) String() string {
	switch t {
	case LoginType:
		return "loginType"
	case ActionType:
		return "actionType"
	case RunTimeType:
		return "runTimeType"
	}
	return ""
}

// 自定义类型转换为json
func (t LogType) MarshalJson() ([]byte, error) {
	return json.Marshal(t.String())
}
