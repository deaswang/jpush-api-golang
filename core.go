package jpush

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

//JPush jpush core struct
type JPush struct {
	appKey       string
	masterSecret string
	auth         string
	Zone         string
	client       *http.Client
	Quota        int // 当前 AppKey 一个时间窗口内可调用次数
	Remaining    int // 当前时间窗口剩余的可用次数
	Reset        int // 距离时间窗口重置剩余的秒数
}

// NewJPush new jpush object
func NewJPush(key, secret string) *JPush {
	jpush := &JPush{appKey: key, masterSecret: secret}
	jpush.auth = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(key+":"+secret)))
	jpush.Zone = "default"
	jpush.client = &http.Client{}
	return jpush
}

// SetAuthorization set Authorization
func (j *JPush) SetAuthorization(key, secret string) {
	j.appKey = key
	j.masterSecret = secret
	j.auth = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(key+":"+secret)))
}

// SetZone set jpush zone
func (j *JPush) SetZone(zone string) {
	if _, ok := ZONES[j.Zone]; ok {
		j.Zone = zone
	}
}

// request request api func
func (j *JPush) request(method, url string, body io.Reader, params map[string]string) ([]byte, error) {
	httpReq, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	q := httpReq.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	httpReq.URL.RawQuery = q.Encode()

	httpReq.Header.Set("Authorization", j.auth)
	httpReq.Header.Set("User-Agent", "jpush-api-golang")
	httpReq.Header.Set("Content-Type", "application/json;charset:utf-8")
	httpReq.Header.Set("connection", "keep-alive")
	resp, err := j.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	limit, err := strconv.Atoi(resp.Header.Get("X-Rate-Limit-Quota"))
	if err == nil {
		j.Quota = limit
	}
	limit, err = strconv.Atoi(resp.Header.Get("X-Rate-Limit-Remaining"))
	if err == nil {
		j.Remaining = limit
	}
	limit, err = strconv.Atoi(resp.Header.Get("X-Rate-Limit-Reset"))
	if err == nil {
		j.Reset = limit
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		var jErr ErrorResponse
		err = json.Unmarshal(buf, &jErr)
		if err != nil {
			return nil, err
		}
		return nil, jErr.Error
	}
	return buf, nil
}

//GroupPush grouppush core struct
type GroupPush struct {
	JPush
}

// NewGroupPush new grouppush object
func NewGroupPush(key, secret string) *GroupPush {
	jpush := &GroupPush{}
	jpush.appKey = key
	jpush.masterSecret = secret
	jpush.auth = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("group-"+key+":"+secret)))
	jpush.Zone = "default"
	jpush.client = &http.Client{}
	return jpush
}

// SetAuthorization set grouppush authorization
func (j *GroupPush) SetAuthorization(key, secret string) {
	j.appKey = key
	j.masterSecret = secret
	j.auth = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("group-"+key+":"+secret)))
}
