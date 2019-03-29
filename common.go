package jpush

import "fmt"

// VERSION api client version
const (
	VERSION = "1.0"
)

// ZONES define api url
var ZONES = make(map[string]map[string]string)

func init() {
	ZONES["default"] = map[string]string{
		"push":     "https://api.jpush.cn/v3/",
		"report":   "https://report.jpush.cn/v3/",
		"device":   "https://device.jpush.cn/v3/devices/",
		"alias":    "https://device.jpush.cn/v3/aliases/",
		"tag":      "https://device.jpush.cn/v3/tags/",
		"schedule": "https://api.jpush.cn/v3/schedules/",
		"admin":    "https://admin.jpush.cn/v1/",
	}
	ZONES["bj"] = map[string]string{
		"push":     "https://bjapi.push.jiguang.cn/v3/",
		"report":   "https://bjapi.push.jiguang.cn/v3/report/",
		"device":   "https://bjapi.push.jiguang.cn/v3/device/",
		"alias":    "https://bjapi.push.jiguang.cn/v3/device/aliases/",
		"tag":      "https://bjapi.push.jiguang.cn/v3/device/tags/",
		"schedule": "https://bjapi.push.jiguang.cn/v3/push/schedules/",
		"admin":    "https://admin.jpush.cn/v1/",
	}
}

// GetURL get the api url address
func (j *JPush) GetURL(key string) string {
	if urls, ok := ZONES[j.Zone]; ok {
		if url, ok := urls[key]; ok {
			return url
		}
	}
	return ZONES["default"]["push"]
}

// ErrorMessage error message response
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse error response from api
type ErrorResponse struct {
	Error ErrorMessage `json:"error"`
}

func (e ErrorMessage) Error() string {
	return fmt.Sprintf("JPush Error %d: %s", e.Code, e.Message)
}

// DefaultResponse default null response
type DefaultResponse struct {
}
