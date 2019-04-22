package jpush

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"
)

// ScheduleTime schedule time
type ScheduleTime time.Time

// UnmarshalJSON unmarshal json
func (s *ScheduleTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("2006-01-02 15:04:05", string(data))
	if err != nil {
		return err
	}
	*s = ScheduleTime(t)
	return nil
}

// MarshalJSON marshal json
func (s *ScheduleTime) MarshalJSON() (data []byte, err error) {
	t := time.Time(*s).Format("2006-01-02 15:04:05")
	return []byte(t), nil
}

// SchedulePeriodical schedule periodical
type SchedulePeriodical struct {
	Start     *ScheduleTime `json:"start,omitempty"`
	End       *ScheduleTime `json:"end,omitempty"`
	Time      string        `json:"time,omitempty"`
	TimeUnit  string        `json:"time_unit,omitempty"`
	Frequency int           `json:"frequency,int,omitempty"`
	Point     []string      `json:"point,omitempty"`
}

// ScheduleTrigger schedule trigger
type ScheduleTrigger struct {
	Single     *ScheduleTime       `json:"single,omitempty"`
	Periodical *SchedulePeriodical `json:"periodical,omitempty"`
}

// ScheduleRequest schedule request body
type ScheduleRequest struct {
	Cid     string           `json:"cid,omitempty"`
	Push    *PushRequest     `json:"push,omitempty"`
	Name    string           `json:"name,omitempty"`
	Enabled bool             `json:"enabled,omitempty"`
	Trigger *ScheduleTrigger `json:"trigger,omitempty"`
}

// ScheduleResponse new schedule response
type ScheduleResponse struct {
	ScheduleID string           `json:"schedule_id"`
	Name       string           `json:"name"`
	Enabled    bool             `json:"enabled,omitempty"`
	Trigger    *ScheduleTrigger `json:"trigger,omitempty"`
	Push       *PushRequest     `json:"push,omitempty"`
}

// SchedulePageResponse schedule page response
type SchedulePageResponse struct {
	TotalCount int                `json:"total_count"`
	TotalPages int                `json:"total_pages"`
	Page       int                `json:"page"`
	Schedules  []ScheduleResponse `json:"schedules"`
}

// ScheduleMsgsResponse schedule message
type ScheduleMsgsResponse struct {
	Count  int           `json:"count"`
	MsgIDs []interface{} `json:"msgids,omitempty"`
}

// Schedule create schedule
// POST /v3/schedules
func (j *JPush) Schedule(req *ScheduleRequest) (*ScheduleResponse, error) {
	url := j.GetURL("schedule")
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	if err != nil {
		return nil, err
	}
	ret := new(ScheduleResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// SchedulePage get schedule list
// GET /v3/schedules?page=
func (j *JPush) SchedulePage(page int) (*SchedulePageResponse, error) {
	url := j.GetURL("schedule")
	params := make(map[string]string)
	params["page"] = strconv.Itoa(page)

	resp, err := j.request("GET", url, nil, params)
	if err != nil {
		return nil, err
	}
	ret := new(SchedulePageResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// ScheduleID get schedule by id
// GET /v3/schedules/{schedule_id}
func (j *JPush) ScheduleID(scheduleID string) (*ScheduleResponse, error) {
	url := j.GetURL("schedule") + scheduleID

	resp, err := j.request("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}
	ret := new(ScheduleResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// ScheduleIDMsgs get all msg ids by schedule id
// GET /v3/schedules/{schedule_id}/msg_ids
func (j *JPush) ScheduleIDMsgs(scheduleID string) (*ScheduleMsgsResponse, error) {
	url := j.GetURL("schedule") + scheduleID + "/msg_ids"

	resp, err := j.request("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}
	ret := new(ScheduleMsgsResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// SchedulePut modify schedule
// PUT /v3/schedules/{schedule_id}
func (j *JPush) SchedulePut(scheduleID string, req *ScheduleRequest) (*ScheduleResponse, error) {
	url := j.GetURL("schedule") + scheduleID
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("PUT", url, bytes.NewReader(buf), nil)
	if err != nil {
		return nil, err
	}
	ret := new(ScheduleResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// ScheduleDelete delete schedule
// DELETE /v3/schedules/{schedule_id}
func (j *JPush) ScheduleDelete(scheduleID string) (*DefaultResponse, error) {
	url := j.GetURL("schedule") + scheduleID

	resp, err := j.request("DELETE", url, nil, nil)
	if err != nil {
		return nil, err
	}
	ret := new(DefaultResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
