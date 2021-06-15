package wInterface

import . "github.com/du5/v5sdk_go/ws/wImpl"

// 请求数据
type WSParam interface {
	EventType() Event
	ToMap() *map[string]string
}
