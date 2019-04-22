package jpush

import (
	"bytes"
	"encoding/json"
)

// AdminAppRequest admin app request
type AdminAppRequest struct {
	AppName        string `json:"app_name"`
	AndroidPackage string `json:"android_package"`
	GroupName      string `json:"group_name"`
}

// AdminCertificateRequest admin certificate
type AdminCertificateRequest struct {
	DevCertificatePassword string `json:"devCertificatePassword,omitempty"`
	ProCertificatePassword string `json:"proCertificatePassword,omitempty"`
	DevCertificateFile     []byte `json:"devCertificateFile,omitempty"`
	ProCertificateFile     []byte `json:"proCertificateFile,omitempty"`
}

// AdminAppResponse new app response
type AdminAppResponse struct {
	AppKey         string `json:"app_key"`
	AndroidPackage string `json:"android_package"`
	IsNewCreated   bool   `json:"is_new_created"`
}

// AdminSuccessResponse success response
type AdminSuccessResponse struct {
	Success string `json:"success"`
}

// AdminApp admin app point
// POST /v1/admin/app
func (j *JPush) AdminApp(req **AdminAppRequest) (*AdminAppResponse, error) {
	url := j.GetURL("admin") + "app"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	if err != nil {
		return nil, err
	}
	ret := new(AdminAppResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// AdminAppDelete delete app
// POST /v1/app/{appkey}/delete
func (j *JPush) AdminAppDelete(appkey string) (*AdminSuccessResponse, error) {
	url := j.GetURL("admin") + "app/" + appkey + "/delete"

	resp, err := j.request("POST", url, nil, nil)
	if err != nil {
		return nil, err
	}
	ret := new(AdminSuccessResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// AdminAppCert admin certificate
// POST /v1/app/{appKey}/certificate
func (j *JPush) AdminAppCert(appkey string, req *AdminCertificateRequest) (*AdminSuccessResponse, error) {
	url := j.GetURL("admin") + "app/" + appkey + "/certificate"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := j.request("POST", url, bytes.NewReader(buf), nil)
	if err != nil {
		return nil, err
	}
	ret := new(AdminSuccessResponse)
	err = json.Unmarshal(resp, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
