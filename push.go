package jpush

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Platform define platform entry
type Platform struct {
	isAll     bool
	Platforms []string
}

// SetAll set isAll value
func (p *Platform) SetAll(all bool) {
	p.isAll = all
}

// UnmarshalJSON unmarshal json
func (p *Platform) UnmarshalJSON(data []byte) error {
	if string(data) == "all" {
		p.isAll = true
		return nil
	}
	p.isAll = false
	return json.Unmarshal(data, &p.Platforms)
}

// MarshalJSON marshal json
func (p *Platform) MarshalJSON() (data []byte, err error) {
	if p.isAll {
		return []byte(`"all"`), nil
	}
	return json.Marshal(p.Platforms)
}

// Audience define Audience
type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationID []string `json:"registration_id,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	ABTest         []string `json:"abtest,omitempty"`
}

// PushAudience define audience entry
type PushAudience struct {
	isAll bool
	Aud   *Audience
}

// SetAll set isAll
func (p *PushAudience) SetAll(all bool) {
	p.isAll = all
}

// UnmarshalJSON unmarshal json
func (p *PushAudience) UnmarshalJSON(data []byte) error {
	if string(data) == "all" {
		p.isAll = true
		return nil
	}
	p.isAll = false
	return json.Unmarshal(data, p.Aud)
}

// MarshalJSON marshal json
func (p *PushAudience) MarshalJSON() (data []byte, err error) {
	if p.isAll {
		return []byte(`"all"`), nil
	}
	return json.Marshal(p.Aud)
}

// PushNotification define notification
type PushNotification struct {
	Alert    string                `json:"alert,omitempty"`
	Android  *NotificationAndroid  `json:"android,omitempty"`
	IOS      *NotificationIOS      `json:"ios,omitempty"`
	WinPhone *NotificationWinPhone `json:"winphone,omitempty"`
}

// NotificationAndroid define android notification
type NotificationAndroid struct {
	Alert      string                 `json:"alert"`
	Title      string                 `json:"title,omitempty"`
	BuilderID  int                    `json:"builder_id,int,omitempty"`
	Priority   int                    `json:"priority,omitempty"`
	Category   string                 `json:"category,omitempty"`
	Style      int                    `json:"style,int,omitempty"`
	AlertType  int                    `json:"alert_type,int,omitempty"`
	BigText    string                 `json:"big_text,omitempty"`
	Inbox      map[string]interface{} `json:"inbox,omitempty"`
	BigPicPath string                 `json:"big_pic_path,omitempty"`
	Extras     map[string]interface{} `json:"extras,omitempty"`
	LargeIcon  string                 `json:"large_icon,omitempty"`
	Intent     map[string]interface{} `json:"intent,omitempty"`
}

// NotificationIOS define ios notification
type NotificationIOS struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            int                    `json:"badge,int,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
}

// NotificationWinPhone define winphone notification
type NotificationWinPhone struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

// PushMessage define push message
type PushMessage struct {
	MsgContent  string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

// SmsMessage define sms message
type SmsMessage struct {
	DelayTime int                    `json:"delay_time,int"`
	TempID    float64                `json:"temp_id,float"`
	TempPara  map[string]interface{} `json:"temp_para,omitempty"`
}

// PushOptions define options
type PushOptions struct {
	SendNo          int    `json:"sendno,int,omitempty"`
	TimeToLive      int    `json:"time_to_live,int,omitempty"`
	OverrideMsgID   int64  `json:"override_msg_id,int64,omitempty"`
	ApnsProduction  bool   `json:"apns_production"`
	ApnsCollapseID  string `json:"apns_collapse_id,omitempty"`
	BigPushDuration int    `json:"big_push_duration,int,omitempty"`
}

// PushRequest define push request body
type PushRequest struct {
	Cid          string            `json:"cid,omitempty"`
	Platform     *Platform         `json:"platform"`
	Audience     *PushAudience     `json:"audience"`
	Notification *PushNotification `json:"notification,omitempty"`
	Message      *PushMessage      `json:"message,omitempty"`
	SmsMessage   *SmsMessage       `json:"sms_message,omitempty"`
	Options      *PushOptions      `json:"options,omitempty"`
}

// PushResponse define push repsone
type PushResponse struct {
	MsgID  string `json:"msg_id"`
	Sendno string `json:"sendno"`
}

// PushCIDResponse get cid response
type PushCIDResponse struct {
	Cids []string `json:"cidlist"`
}

// Push push notification or message to devices
// POST /v3/push
func (j *JPush) Push(req *PushRequest) (*PushResponse, error) {
	url := j.GetURL("push") + "push"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	if err != nil {
		return nil, err
	}
	ret := new(PushResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// PushGetCid get push by cid
// GET /v3/push/cid[?count=n[&type=xx]]
func (j *JPush) PushGetCid(count int, cidtype string) (*PushCIDResponse, error) {
	url := j.GetURL("push") + "push/cid"
	params := make(map[string]string)
	params["count"] = strconv.Itoa(count)
	params["type"] = cidtype

	resp, err := j.request("GET", url, nil, params)
	if err != nil {
		return nil, err
	}
	ret := new(PushCIDResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// GroupPush group push
// POST /v3/grouppush
func (j *GroupPush) GroupPush(req *PushRequest) (*PushResponse, error) {
	url := j.GetURL("push") + "grouppush"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	if err != nil {
		return nil, err
	}
	ret := new(PushResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// PushValidate push validate, not real push
// POST /v3/push/validate
func (j *JPush) PushValidate(req *PushRequest) (*PushResponse, error) {
	url := j.GetURL("push") + "push/validate"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	if err != nil {
		return nil, err
	}
	ret := new(PushResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
