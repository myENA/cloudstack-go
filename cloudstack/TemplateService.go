//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type TemplateServiceIface interface {
	CopyTemplate(p *CopyTemplateParams) (*CopyTemplateResponse, error)
	NewCopyTemplateParams(id string) *CopyTemplateParams
	CreateTemplate(p *CreateTemplateParams) (*CreateTemplateResponse, error)
	NewCreateTemplateParams(displaytext string, name string, ostypeid string) *CreateTemplateParams
	DeleteTemplate(p *DeleteTemplateParams) (*DeleteTemplateResponse, error)
	NewDeleteTemplateParams(id string) *DeleteTemplateParams
	ExtractTemplate(p *ExtractTemplateParams) (*ExtractTemplateResponse, error)
	NewExtractTemplateParams(id string, mode string) *ExtractTemplateParams
	GetUploadParamsForTemplate(p *GetUploadParamsForTemplateParams) (*GetUploadParamsForTemplateResponse, error)
	NewGetUploadParamsForTemplateParams(displaytext string, format string, hypervisor string, name string, zoneid string) *GetUploadParamsForTemplateParams
	ListTemplatePermissions(p *ListTemplatePermissionsParams) (*ListTemplatePermissionsResponse, error)
	NewListTemplatePermissionsParams(id string) *ListTemplatePermissionsParams
	GetTemplatePermissionByID(id string, opts ...OptionFunc) (*TemplatePermission, int, error)
	ListTemplates(p *ListTemplatesParams) (*ListTemplatesResponse, error)
	NewListTemplatesParams(templatefilter string) *ListTemplatesParams
	GetTemplateID(name string, templatefilter string, zoneid string, opts ...OptionFunc) (string, int, error)
	GetTemplateByName(name string, templatefilter string, zoneid string, opts ...OptionFunc) (*Template, int, error)
	GetTemplateByID(id string, templatefilter string, opts ...OptionFunc) (*Template, int, error)
	PrepareTemplate(p *PrepareTemplateParams) (*PrepareTemplateResponse, error)
	NewPrepareTemplateParams(templateid string, zoneid string) *PrepareTemplateParams
	RegisterTemplate(p *RegisterTemplateParams) (*RegisterTemplateResponse, error)
	NewRegisterTemplateParams(displaytext string, format string, hypervisor string, name string, url string) *RegisterTemplateParams
	UpdateTemplate(p *UpdateTemplateParams) (*UpdateTemplateResponse, error)
	NewUpdateTemplateParams(id string) *UpdateTemplateParams
	UpdateTemplatePermissions(p *UpdateTemplatePermissionsParams) (*UpdateTemplatePermissionsResponse, error)
	NewUpdateTemplatePermissionsParams(id string) *UpdateTemplatePermissionsParams
	UpgradeRouterTemplate(p *UpgradeRouterTemplateParams) (*UpgradeRouterTemplateResponse, error)
	NewUpgradeRouterTemplateParams() *UpgradeRouterTemplateParams
}

type CopyTemplateParams struct {
	P map[string]interface{}
}

func (P *CopyTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["destzoneid"]; found {
		u.Set("destzoneid", v.(string))
	}
	if v, found := P.P["destzoneids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("destzoneids", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["sourcezoneid"]; found {
		u.Set("sourcezoneid", v.(string))
	}
	return u
}

func (P *CopyTemplateParams) SetDestzoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["destzoneid"] = v
}

func (P *CopyTemplateParams) GetDestzoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["destzoneid"].(string)
	return value, ok
}

func (P *CopyTemplateParams) SetDestzoneids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["destzoneids"] = v
}

func (P *CopyTemplateParams) GetDestzoneids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["destzoneids"].([]string)
	return value, ok
}

func (P *CopyTemplateParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *CopyTemplateParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *CopyTemplateParams) SetSourcezoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sourcezoneid"] = v
}

func (P *CopyTemplateParams) GetSourcezoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sourcezoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CopyTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewCopyTemplateParams(id string) *CopyTemplateParams {
	P := &CopyTemplateParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Copies a template from one zone to another.
func (s *TemplateService) CopyTemplate(p *CopyTemplateParams) (*CopyTemplateResponse, error) {
	resp, err := s.cs.newRequest("copyTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CopyTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CopyTemplateResponse struct {
	Account               string              `json:"account"`
	Accountid             string              `json:"accountid"`
	Bits                  int                 `json:"bits"`
	Bootable              bool                `json:"bootable"`
	Checksum              string              `json:"checksum"`
	Childtemplates        []interface{}       `json:"childtemplates"`
	Created               string              `json:"created"`
	CrossZones            bool                `json:"crossZones"`
	Deployasis            bool                `json:"deployasis"`
	Deployasisdetails     map[string]string   `json:"deployasisdetails"`
	Details               map[string]string   `json:"details"`
	Directdownload        bool                `json:"directdownload"`
	Displaytext           string              `json:"displaytext"`
	Domain                string              `json:"domain"`
	Domainid              string              `json:"domainid"`
	Downloaddetails       []map[string]string `json:"downloaddetails"`
	Format                string              `json:"format"`
	Hasannotations        bool                `json:"hasannotations"`
	Hostid                string              `json:"hostid"`
	Hostname              string              `json:"hostname"`
	Hypervisor            string              `json:"hypervisor"`
	Icon                  string              `json:"icon"`
	Id                    string              `json:"id"`
	Isdynamicallyscalable bool                `json:"isdynamicallyscalable"`
	Isextractable         bool                `json:"isextractable"`
	Isfeatured            bool                `json:"isfeatured"`
	Ispublic              bool                `json:"ispublic"`
	Isready               bool                `json:"isready"`
	JobID                 string              `json:"jobid"`
	Jobstatus             int                 `json:"jobstatus"`
	Name                  string              `json:"name"`
	Ostypeid              string              `json:"ostypeid"`
	Ostypename            string              `json:"ostypename"`
	Parenttemplateid      string              `json:"parenttemplateid"`
	Passwordenabled       bool                `json:"passwordenabled"`
	Physicalsize          int64               `json:"physicalsize"`
	Project               string              `json:"project"`
	Projectid             string              `json:"projectid"`
	Removed               string              `json:"removed"`
	Requireshvm           bool                `json:"requireshvm"`
	Size                  int64               `json:"size"`
	Sourcetemplateid      string              `json:"sourcetemplateid"`
	Sshkeyenabled         bool                `json:"sshkeyenabled"`
	Status                string              `json:"status"`
	Tags                  []Tags              `json:"tags"`
	Templatetag           string              `json:"templatetag"`
	Templatetype          string              `json:"templatetype"`
	Url                   string              `json:"url"`
	Zoneid                string              `json:"zoneid"`
	Zonename              string              `json:"zonename"`
}

func (r *CopyTemplateResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias CopyTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CreateTemplateParams struct {
	P map[string]interface{}
}

func (P *CreateTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["bits"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("bits", vv)
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := P.P["isfeatured"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isfeatured", vv)
	}
	if v, found := P.P["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := P.P["passwordenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passwordenabled", vv)
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["requireshvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("requireshvm", vv)
	}
	if v, found := P.P["snapshotid"]; found {
		u.Set("snapshotid", v.(string))
	}
	if v, found := P.P["sshkeyenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sshkeyenabled", vv)
	}
	if v, found := P.P["templatetag"]; found {
		u.Set("templatetag", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := P.P["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (P *CreateTemplateParams) SetBits(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bits"] = v
}

func (P *CreateTemplateParams) GetBits() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bits"].(int)
	return value, ok
}

func (P *CreateTemplateParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *CreateTemplateParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *CreateTemplateParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *CreateTemplateParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *CreateTemplateParams) SetIsdynamicallyscalable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isdynamicallyscalable"] = v
}

func (P *CreateTemplateParams) GetIsdynamicallyscalable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isdynamicallyscalable"].(bool)
	return value, ok
}

func (P *CreateTemplateParams) SetIsfeatured(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isfeatured"] = v
}

func (P *CreateTemplateParams) GetIsfeatured() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isfeatured"].(bool)
	return value, ok
}

func (P *CreateTemplateParams) SetIspublic(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ispublic"] = v
}

func (P *CreateTemplateParams) GetIspublic() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ispublic"].(bool)
	return value, ok
}

func (P *CreateTemplateParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateTemplateParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateTemplateParams) SetOstypeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ostypeid"] = v
}

func (P *CreateTemplateParams) GetOstypeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ostypeid"].(string)
	return value, ok
}

func (P *CreateTemplateParams) SetPasswordenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["passwordenabled"] = v
}

func (P *CreateTemplateParams) GetPasswordenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["passwordenabled"].(bool)
	return value, ok
}

func (P *CreateTemplateParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateTemplateParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *CreateTemplateParams) SetRequireshvm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["requireshvm"] = v
}

func (P *CreateTemplateParams) GetRequireshvm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["requireshvm"].(bool)
	return value, ok
}

func (P *CreateTemplateParams) SetSnapshotid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["snapshotid"] = v
}

func (P *CreateTemplateParams) GetSnapshotid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["snapshotid"].(string)
	return value, ok
}

func (P *CreateTemplateParams) SetSshkeyenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sshkeyenabled"] = v
}

func (P *CreateTemplateParams) GetSshkeyenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sshkeyenabled"].(bool)
	return value, ok
}

func (P *CreateTemplateParams) SetTemplatetag(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templatetag"] = v
}

func (P *CreateTemplateParams) GetTemplatetag() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templatetag"].(string)
	return value, ok
}

func (P *CreateTemplateParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *CreateTemplateParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *CreateTemplateParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *CreateTemplateParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *CreateTemplateParams) SetVolumeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["volumeid"] = v
}

func (P *CreateTemplateParams) GetVolumeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewCreateTemplateParams(displaytext string, name string, ostypeid string) *CreateTemplateParams {
	P := &CreateTemplateParams{}
	P.P = make(map[string]interface{})
	P.P["displaytext"] = displaytext
	P.P["name"] = name
	P.P["ostypeid"] = ostypeid
	return P
}

// Creates a template of a virtual machine. The virtual machine must be in a STOPPED state. A template created from this command is automatically designated as a private template visible to the account that created it.
func (s *TemplateService) CreateTemplate(p *CreateTemplateParams) (*CreateTemplateResponse, error) {
	resp, err := s.cs.newRequest("createTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateTemplateResponse struct {
	Account               string              `json:"account"`
	Accountid             string              `json:"accountid"`
	Bits                  int                 `json:"bits"`
	Bootable              bool                `json:"bootable"`
	Checksum              string              `json:"checksum"`
	Childtemplates        []interface{}       `json:"childtemplates"`
	Created               string              `json:"created"`
	CrossZones            bool                `json:"crossZones"`
	Deployasis            bool                `json:"deployasis"`
	Deployasisdetails     map[string]string   `json:"deployasisdetails"`
	Details               map[string]string   `json:"details"`
	Directdownload        bool                `json:"directdownload"`
	Displaytext           string              `json:"displaytext"`
	Domain                string              `json:"domain"`
	Domainid              string              `json:"domainid"`
	Downloaddetails       []map[string]string `json:"downloaddetails"`
	Format                string              `json:"format"`
	Hasannotations        bool                `json:"hasannotations"`
	Hostid                string              `json:"hostid"`
	Hostname              string              `json:"hostname"`
	Hypervisor            string              `json:"hypervisor"`
	Icon                  string              `json:"icon"`
	Id                    string              `json:"id"`
	Isdynamicallyscalable bool                `json:"isdynamicallyscalable"`
	Isextractable         bool                `json:"isextractable"`
	Isfeatured            bool                `json:"isfeatured"`
	Ispublic              bool                `json:"ispublic"`
	Isready               bool                `json:"isready"`
	JobID                 string              `json:"jobid"`
	Jobstatus             int                 `json:"jobstatus"`
	Name                  string              `json:"name"`
	Ostypeid              string              `json:"ostypeid"`
	Ostypename            string              `json:"ostypename"`
	Parenttemplateid      string              `json:"parenttemplateid"`
	Passwordenabled       bool                `json:"passwordenabled"`
	Physicalsize          int64               `json:"physicalsize"`
	Project               string              `json:"project"`
	Projectid             string              `json:"projectid"`
	Removed               string              `json:"removed"`
	Requireshvm           bool                `json:"requireshvm"`
	Size                  int64               `json:"size"`
	Sourcetemplateid      string              `json:"sourcetemplateid"`
	Sshkeyenabled         bool                `json:"sshkeyenabled"`
	Status                string              `json:"status"`
	Tags                  []Tags              `json:"tags"`
	Templatetag           string              `json:"templatetag"`
	Templatetype          string              `json:"templatetype"`
	Url                   string              `json:"url"`
	Zoneid                string              `json:"zoneid"`
	Zonename              string              `json:"zonename"`
}

func (r *CreateTemplateResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias CreateTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteTemplateParams struct {
	P map[string]interface{}
}

func (P *DeleteTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *DeleteTemplateParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *DeleteTemplateParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *DeleteTemplateParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteTemplateParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *DeleteTemplateParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *DeleteTemplateParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewDeleteTemplateParams(id string) *DeleteTemplateParams {
	P := &DeleteTemplateParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a template from the system. All virtual machines using the deleted template will not be affected.
func (s *TemplateService) DeleteTemplate(p *DeleteTemplateParams) (*DeleteTemplateResponse, error) {
	resp, err := s.cs.newRequest("deleteTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteTemplateResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ExtractTemplateParams struct {
	P map[string]interface{}
}

func (P *ExtractTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["mode"]; found {
		u.Set("mode", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ExtractTemplateParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ExtractTemplateParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ExtractTemplateParams) SetMode(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["mode"] = v
}

func (P *ExtractTemplateParams) GetMode() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["mode"].(string)
	return value, ok
}

func (P *ExtractTemplateParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *ExtractTemplateParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *ExtractTemplateParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ExtractTemplateParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ExtractTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewExtractTemplateParams(id string, mode string) *ExtractTemplateParams {
	P := &ExtractTemplateParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["mode"] = mode
	return P
}

// Extracts a template
func (s *TemplateService) ExtractTemplate(p *ExtractTemplateParams) (*ExtractTemplateResponse, error) {
	resp, err := s.cs.newRequest("extractTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExtractTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type ExtractTemplateResponse struct {
	Accountid        string `json:"accountid"`
	Created          string `json:"created"`
	ExtractId        string `json:"extractId"`
	ExtractMode      string `json:"extractMode"`
	Id               string `json:"id"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Name             string `json:"name"`
	Resultstring     string `json:"resultstring"`
	State            string `json:"state"`
	Status           string `json:"status"`
	Storagetype      string `json:"storagetype"`
	Uploadpercentage int    `json:"uploadpercentage"`
	Url              string `json:"url"`
	Zoneid           string `json:"zoneid"`
	Zonename         string `json:"zonename"`
}

type GetUploadParamsForTemplateParams struct {
	P map[string]interface{}
}

func (P *GetUploadParamsForTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["bits"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("bits", vv)
	}
	if v, found := P.P["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := P.P["deployasis"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("deployasis", vv)
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := P.P["isextractable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isextractable", vv)
	}
	if v, found := P.P["isfeatured"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isfeatured", vv)
	}
	if v, found := P.P["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := P.P["isrouting"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrouting", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := P.P["passwordenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passwordenabled", vv)
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["requireshvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("requireshvm", vv)
	}
	if v, found := P.P["sshkeyenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sshkeyenabled", vv)
	}
	if v, found := P.P["templatetag"]; found {
		u.Set("templatetag", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *GetUploadParamsForTemplateParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *GetUploadParamsForTemplateParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetBits(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bits"] = v
}

func (P *GetUploadParamsForTemplateParams) GetBits() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bits"].(int)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetChecksum(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["checksum"] = v
}

func (P *GetUploadParamsForTemplateParams) GetChecksum() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["checksum"].(string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetDeployasis(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["deployasis"] = v
}

func (P *GetUploadParamsForTemplateParams) GetDeployasis() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["deployasis"].(bool)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *GetUploadParamsForTemplateParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *GetUploadParamsForTemplateParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *GetUploadParamsForTemplateParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetFormat(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["format"] = v
}

func (P *GetUploadParamsForTemplateParams) GetFormat() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["format"].(string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *GetUploadParamsForTemplateParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetIsdynamicallyscalable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isdynamicallyscalable"] = v
}

func (P *GetUploadParamsForTemplateParams) GetIsdynamicallyscalable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isdynamicallyscalable"].(bool)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetIsextractable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isextractable"] = v
}

func (P *GetUploadParamsForTemplateParams) GetIsextractable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isextractable"].(bool)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetIsfeatured(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isfeatured"] = v
}

func (P *GetUploadParamsForTemplateParams) GetIsfeatured() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isfeatured"].(bool)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetIspublic(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ispublic"] = v
}

func (P *GetUploadParamsForTemplateParams) GetIspublic() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ispublic"].(bool)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetIsrouting(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrouting"] = v
}

func (P *GetUploadParamsForTemplateParams) GetIsrouting() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrouting"].(bool)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *GetUploadParamsForTemplateParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetOstypeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ostypeid"] = v
}

func (P *GetUploadParamsForTemplateParams) GetOstypeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ostypeid"].(string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetPasswordenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["passwordenabled"] = v
}

func (P *GetUploadParamsForTemplateParams) GetPasswordenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["passwordenabled"].(bool)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *GetUploadParamsForTemplateParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetRequireshvm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["requireshvm"] = v
}

func (P *GetUploadParamsForTemplateParams) GetRequireshvm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["requireshvm"].(bool)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetSshkeyenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sshkeyenabled"] = v
}

func (P *GetUploadParamsForTemplateParams) GetSshkeyenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sshkeyenabled"].(bool)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetTemplatetag(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templatetag"] = v
}

func (P *GetUploadParamsForTemplateParams) GetTemplatetag() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templatetag"].(string)
	return value, ok
}

func (P *GetUploadParamsForTemplateParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *GetUploadParamsForTemplateParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new GetUploadParamsForTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewGetUploadParamsForTemplateParams(displaytext string, format string, hypervisor string, name string, zoneid string) *GetUploadParamsForTemplateParams {
	P := &GetUploadParamsForTemplateParams{}
	P.P = make(map[string]interface{})
	P.P["displaytext"] = displaytext
	P.P["format"] = format
	P.P["hypervisor"] = hypervisor
	P.P["name"] = name
	P.P["zoneid"] = zoneid
	return P
}

// upload an existing template into the CloudStack cloud.
func (s *TemplateService) GetUploadParamsForTemplate(p *GetUploadParamsForTemplateParams) (*GetUploadParamsForTemplateResponse, error) {
	resp, err := s.cs.newRequest("getUploadParamsForTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response GetUploadParamsForTemplateResponse `json:"getuploadparams"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type GetUploadParamsForTemplateResponse struct {
	Expires   string `json:"expires"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Metadata  string `json:"metadata"`
	PostURL   string `json:"postURL"`
	Signature string `json:"signature"`
}

type ListTemplatePermissionsParams struct {
	P map[string]interface{}
}

func (P *ListTemplatePermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ListTemplatePermissionsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListTemplatePermissionsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ListTemplatePermissionsParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewListTemplatePermissionsParams(id string) *ListTemplatePermissionsParams {
	P := &ListTemplatePermissionsParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplatePermissionByID(id string, opts ...OptionFunc) (*TemplatePermission, int, error) {
	P := &ListTemplatePermissionsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListTemplatePermissions(P)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.TemplatePermissions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for TemplatePermission UUID: %s!", id)
}

// List template visibility and all accounts that have permissions to view this template.
func (s *TemplateService) ListTemplatePermissions(p *ListTemplatePermissionsParams) (*ListTemplatePermissionsResponse, error) {
	resp, err := s.cs.newRequest("listTemplatePermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTemplatePermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTemplatePermissionsResponse struct {
	Count               int                   `json:"count"`
	TemplatePermissions []*TemplatePermission `json:"templatepermission"`
}

type TemplatePermission struct {
	Account    []string `json:"account"`
	Domainid   string   `json:"domainid"`
	Id         string   `json:"id"`
	Ispublic   bool     `json:"ispublic"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Projectids []string `json:"projectids"`
}

type ListTemplatesParams struct {
	P map[string]interface{}
}

func (P *ListTemplatesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := P.P["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := P.P["parenttemplateid"]; found {
		u.Set("parenttemplateid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := P.P["showremoved"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showremoved", vv)
	}
	if v, found := P.P["showunique"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showunique", vv)
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["templatefilter"]; found {
		u.Set("templatefilter", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListTemplatesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListTemplatesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListTemplatesParams) SetDetails(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ListTemplatesParams) GetDetails() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].([]string)
	return value, ok
}

func (P *ListTemplatesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListTemplatesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListTemplatesParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *ListTemplatesParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *ListTemplatesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListTemplatesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListTemplatesParams) SetIds(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ids"] = v
}

func (P *ListTemplatesParams) GetIds() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ids"].([]string)
	return value, ok
}

func (P *ListTemplatesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListTemplatesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListTemplatesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListTemplatesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListTemplatesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListTemplatesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListTemplatesParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListTemplatesParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListTemplatesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListTemplatesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListTemplatesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListTemplatesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListTemplatesParams) SetParenttemplateid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["parenttemplateid"] = v
}

func (P *ListTemplatesParams) GetParenttemplateid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["parenttemplateid"].(string)
	return value, ok
}

func (P *ListTemplatesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListTemplatesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListTemplatesParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListTemplatesParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

func (P *ListTemplatesParams) SetShowremoved(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showremoved"] = v
}

func (P *ListTemplatesParams) GetShowremoved() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showremoved"].(bool)
	return value, ok
}

func (P *ListTemplatesParams) SetShowunique(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showunique"] = v
}

func (P *ListTemplatesParams) GetShowunique() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showunique"].(bool)
	return value, ok
}

func (P *ListTemplatesParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListTemplatesParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListTemplatesParams) SetTemplatefilter(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templatefilter"] = v
}

func (P *ListTemplatesParams) GetTemplatefilter() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templatefilter"].(string)
	return value, ok
}

func (P *ListTemplatesParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListTemplatesParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTemplatesParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewListTemplatesParams(templatefilter string) *ListTemplatesParams {
	P := &ListTemplatesParams{}
	P.P = make(map[string]interface{})
	P.P["templatefilter"] = templatefilter
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplateID(name string, templatefilter string, zoneid string, opts ...OptionFunc) (string, int, error) {
	P := &ListTemplatesParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name
	P.P["templatefilter"] = templatefilter
	P.P["zoneid"] = zoneid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListTemplates(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Templates[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Templates {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplateByName(name string, templatefilter string, zoneid string, opts ...OptionFunc) (*Template, int, error) {
	id, count, err := s.GetTemplateID(name, templatefilter, zoneid, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetTemplateByID(id, templatefilter, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *TemplateService) GetTemplateByID(id string, templatefilter string, opts ...OptionFunc) (*Template, int, error) {
	P := &ListTemplatesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id
	P.P["templatefilter"] = templatefilter

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListTemplates(P)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Templates[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Template UUID: %s!", id)
}

// List all public, private, and privileged templates.
func (s *TemplateService) ListTemplates(p *ListTemplatesParams) (*ListTemplatesResponse, error) {
	resp, err := s.cs.newRequest("listTemplates", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTemplatesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTemplatesResponse struct {
	Count     int         `json:"count"`
	Templates []*Template `json:"template"`
}

type Template struct {
	Account               string              `json:"account"`
	Accountid             string              `json:"accountid"`
	Bits                  int                 `json:"bits"`
	Bootable              bool                `json:"bootable"`
	Checksum              string              `json:"checksum"`
	Childtemplates        []interface{}       `json:"childtemplates"`
	Created               string              `json:"created"`
	CrossZones            bool                `json:"crossZones"`
	Deployasis            bool                `json:"deployasis"`
	Deployasisdetails     map[string]string   `json:"deployasisdetails"`
	Details               map[string]string   `json:"details"`
	Directdownload        bool                `json:"directdownload"`
	Displaytext           string              `json:"displaytext"`
	Domain                string              `json:"domain"`
	Domainid              string              `json:"domainid"`
	Downloaddetails       []map[string]string `json:"downloaddetails"`
	Format                string              `json:"format"`
	Hasannotations        bool                `json:"hasannotations"`
	Hostid                string              `json:"hostid"`
	Hostname              string              `json:"hostname"`
	Hypervisor            string              `json:"hypervisor"`
	Icon                  string              `json:"icon"`
	Id                    string              `json:"id"`
	Isdynamicallyscalable bool                `json:"isdynamicallyscalable"`
	Isextractable         bool                `json:"isextractable"`
	Isfeatured            bool                `json:"isfeatured"`
	Ispublic              bool                `json:"ispublic"`
	Isready               bool                `json:"isready"`
	JobID                 string              `json:"jobid"`
	Jobstatus             int                 `json:"jobstatus"`
	Name                  string              `json:"name"`
	Ostypeid              string              `json:"ostypeid"`
	Ostypename            string              `json:"ostypename"`
	Parenttemplateid      string              `json:"parenttemplateid"`
	Passwordenabled       bool                `json:"passwordenabled"`
	Physicalsize          int64               `json:"physicalsize"`
	Project               string              `json:"project"`
	Projectid             string              `json:"projectid"`
	Removed               string              `json:"removed"`
	Requireshvm           bool                `json:"requireshvm"`
	Size                  int64               `json:"size"`
	Sourcetemplateid      string              `json:"sourcetemplateid"`
	Sshkeyenabled         bool                `json:"sshkeyenabled"`
	Status                string              `json:"status"`
	Tags                  []Tags              `json:"tags"`
	Templatetag           string              `json:"templatetag"`
	Templatetype          string              `json:"templatetype"`
	Url                   string              `json:"url"`
	Zoneid                string              `json:"zoneid"`
	Zonename              string              `json:"zonename"`
}

func (r *Template) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias Template
	return json.Unmarshal(b, (*alias)(r))
}

type PrepareTemplateParams struct {
	P map[string]interface{}
}

func (P *PrepareTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := P.P["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *PrepareTemplateParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *PrepareTemplateParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

func (P *PrepareTemplateParams) SetTemplateid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templateid"] = v
}

func (P *PrepareTemplateParams) GetTemplateid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templateid"].(string)
	return value, ok
}

func (P *PrepareTemplateParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *PrepareTemplateParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new PrepareTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewPrepareTemplateParams(templateid string, zoneid string) *PrepareTemplateParams {
	P := &PrepareTemplateParams{}
	P.P = make(map[string]interface{})
	P.P["templateid"] = templateid
	P.P["zoneid"] = zoneid
	return P
}

// load template into primary storage
func (s *TemplateService) PrepareTemplate(p *PrepareTemplateParams) (*PrepareTemplateResponse, error) {
	resp, err := s.cs.newRequest("prepareTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PrepareTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type PrepareTemplateResponse struct {
	Account               string              `json:"account"`
	Accountid             string              `json:"accountid"`
	Bits                  int                 `json:"bits"`
	Bootable              bool                `json:"bootable"`
	Checksum              string              `json:"checksum"`
	Childtemplates        []interface{}       `json:"childtemplates"`
	Created               string              `json:"created"`
	CrossZones            bool                `json:"crossZones"`
	Deployasis            bool                `json:"deployasis"`
	Deployasisdetails     map[string]string   `json:"deployasisdetails"`
	Details               map[string]string   `json:"details"`
	Directdownload        bool                `json:"directdownload"`
	Displaytext           string              `json:"displaytext"`
	Domain                string              `json:"domain"`
	Domainid              string              `json:"domainid"`
	Downloaddetails       []map[string]string `json:"downloaddetails"`
	Format                string              `json:"format"`
	Hasannotations        bool                `json:"hasannotations"`
	Hostid                string              `json:"hostid"`
	Hostname              string              `json:"hostname"`
	Hypervisor            string              `json:"hypervisor"`
	Icon                  string              `json:"icon"`
	Id                    string              `json:"id"`
	Isdynamicallyscalable bool                `json:"isdynamicallyscalable"`
	Isextractable         bool                `json:"isextractable"`
	Isfeatured            bool                `json:"isfeatured"`
	Ispublic              bool                `json:"ispublic"`
	Isready               bool                `json:"isready"`
	JobID                 string              `json:"jobid"`
	Jobstatus             int                 `json:"jobstatus"`
	Name                  string              `json:"name"`
	Ostypeid              string              `json:"ostypeid"`
	Ostypename            string              `json:"ostypename"`
	Parenttemplateid      string              `json:"parenttemplateid"`
	Passwordenabled       bool                `json:"passwordenabled"`
	Physicalsize          int64               `json:"physicalsize"`
	Project               string              `json:"project"`
	Projectid             string              `json:"projectid"`
	Removed               string              `json:"removed"`
	Requireshvm           bool                `json:"requireshvm"`
	Size                  int64               `json:"size"`
	Sourcetemplateid      string              `json:"sourcetemplateid"`
	Sshkeyenabled         bool                `json:"sshkeyenabled"`
	Status                string              `json:"status"`
	Tags                  []Tags              `json:"tags"`
	Templatetag           string              `json:"templatetag"`
	Templatetype          string              `json:"templatetype"`
	Url                   string              `json:"url"`
	Zoneid                string              `json:"zoneid"`
	Zonename              string              `json:"zonename"`
}

func (r *PrepareTemplateResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias PrepareTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RegisterTemplateParams struct {
	P map[string]interface{}
}

func (P *RegisterTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["bits"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("bits", vv)
	}
	if v, found := P.P["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := P.P["deployasis"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("deployasis", vv)
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := P.P["directdownload"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("directdownload", vv)
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := P.P["isextractable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isextractable", vv)
	}
	if v, found := P.P["isfeatured"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isfeatured", vv)
	}
	if v, found := P.P["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := P.P["isrouting"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrouting", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := P.P["passwordenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passwordenabled", vv)
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["requireshvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("requireshvm", vv)
	}
	if v, found := P.P["sshkeyenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sshkeyenabled", vv)
	}
	if v, found := P.P["templatetag"]; found {
		u.Set("templatetag", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	if v, found := P.P["zoneids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneids", vv)
	}
	return u
}

func (P *RegisterTemplateParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *RegisterTemplateParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetBits(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bits"] = v
}

func (P *RegisterTemplateParams) GetBits() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bits"].(int)
	return value, ok
}

func (P *RegisterTemplateParams) SetChecksum(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["checksum"] = v
}

func (P *RegisterTemplateParams) GetChecksum() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["checksum"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetDeployasis(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["deployasis"] = v
}

func (P *RegisterTemplateParams) GetDeployasis() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["deployasis"].(bool)
	return value, ok
}

func (P *RegisterTemplateParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *RegisterTemplateParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *RegisterTemplateParams) SetDirectdownload(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["directdownload"] = v
}

func (P *RegisterTemplateParams) GetDirectdownload() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["directdownload"].(bool)
	return value, ok
}

func (P *RegisterTemplateParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *RegisterTemplateParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *RegisterTemplateParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetFormat(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["format"] = v
}

func (P *RegisterTemplateParams) GetFormat() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["format"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *RegisterTemplateParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetIsdynamicallyscalable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isdynamicallyscalable"] = v
}

func (P *RegisterTemplateParams) GetIsdynamicallyscalable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isdynamicallyscalable"].(bool)
	return value, ok
}

func (P *RegisterTemplateParams) SetIsextractable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isextractable"] = v
}

func (P *RegisterTemplateParams) GetIsextractable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isextractable"].(bool)
	return value, ok
}

func (P *RegisterTemplateParams) SetIsfeatured(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isfeatured"] = v
}

func (P *RegisterTemplateParams) GetIsfeatured() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isfeatured"].(bool)
	return value, ok
}

func (P *RegisterTemplateParams) SetIspublic(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ispublic"] = v
}

func (P *RegisterTemplateParams) GetIspublic() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ispublic"].(bool)
	return value, ok
}

func (P *RegisterTemplateParams) SetIsrouting(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrouting"] = v
}

func (P *RegisterTemplateParams) GetIsrouting() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrouting"].(bool)
	return value, ok
}

func (P *RegisterTemplateParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *RegisterTemplateParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetOstypeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ostypeid"] = v
}

func (P *RegisterTemplateParams) GetOstypeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ostypeid"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetPasswordenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["passwordenabled"] = v
}

func (P *RegisterTemplateParams) GetPasswordenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["passwordenabled"].(bool)
	return value, ok
}

func (P *RegisterTemplateParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *RegisterTemplateParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetRequireshvm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["requireshvm"] = v
}

func (P *RegisterTemplateParams) GetRequireshvm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["requireshvm"].(bool)
	return value, ok
}

func (P *RegisterTemplateParams) SetSshkeyenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sshkeyenabled"] = v
}

func (P *RegisterTemplateParams) GetSshkeyenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sshkeyenabled"].(bool)
	return value, ok
}

func (P *RegisterTemplateParams) SetTemplatetag(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templatetag"] = v
}

func (P *RegisterTemplateParams) GetTemplatetag() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templatetag"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *RegisterTemplateParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *RegisterTemplateParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

func (P *RegisterTemplateParams) SetZoneids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneids"] = v
}

func (P *RegisterTemplateParams) GetZoneids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneids"].([]string)
	return value, ok
}

// You should always use this function to get a new RegisterTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewRegisterTemplateParams(displaytext string, format string, hypervisor string, name string, url string) *RegisterTemplateParams {
	P := &RegisterTemplateParams{}
	P.P = make(map[string]interface{})
	P.P["displaytext"] = displaytext
	P.P["format"] = format
	P.P["hypervisor"] = hypervisor
	P.P["name"] = name
	P.P["url"] = url
	return P
}

// Registers an existing template into the CloudStack cloud.
func (s *TemplateService) RegisterTemplate(p *RegisterTemplateParams) (*RegisterTemplateResponse, error) {
	resp, err := s.cs.newRequest("registerTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RegisterTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterTemplateResponse struct {
	Count            int                 `json:"count"`
	RegisterTemplate []*RegisterTemplate `json:"template"`
}

type RegisterTemplate struct {
	Account               string              `json:"account"`
	Accountid             string              `json:"accountid"`
	Bits                  int                 `json:"bits"`
	Bootable              bool                `json:"bootable"`
	Checksum              string              `json:"checksum"`
	Childtemplates        []interface{}       `json:"childtemplates"`
	Created               string              `json:"created"`
	CrossZones            bool                `json:"crossZones"`
	Deployasis            bool                `json:"deployasis"`
	Deployasisdetails     map[string]string   `json:"deployasisdetails"`
	Details               map[string]string   `json:"details"`
	Directdownload        bool                `json:"directdownload"`
	Displaytext           string              `json:"displaytext"`
	Domain                string              `json:"domain"`
	Domainid              string              `json:"domainid"`
	Downloaddetails       []map[string]string `json:"downloaddetails"`
	Format                string              `json:"format"`
	Hasannotations        bool                `json:"hasannotations"`
	Hostid                string              `json:"hostid"`
	Hostname              string              `json:"hostname"`
	Hypervisor            string              `json:"hypervisor"`
	Icon                  string              `json:"icon"`
	Id                    string              `json:"id"`
	Isdynamicallyscalable bool                `json:"isdynamicallyscalable"`
	Isextractable         bool                `json:"isextractable"`
	Isfeatured            bool                `json:"isfeatured"`
	Ispublic              bool                `json:"ispublic"`
	Isready               bool                `json:"isready"`
	JobID                 string              `json:"jobid"`
	Jobstatus             int                 `json:"jobstatus"`
	Name                  string              `json:"name"`
	Ostypeid              string              `json:"ostypeid"`
	Ostypename            string              `json:"ostypename"`
	Parenttemplateid      string              `json:"parenttemplateid"`
	Passwordenabled       bool                `json:"passwordenabled"`
	Physicalsize          int64               `json:"physicalsize"`
	Project               string              `json:"project"`
	Projectid             string              `json:"projectid"`
	Removed               string              `json:"removed"`
	Requireshvm           bool                `json:"requireshvm"`
	Size                  int64               `json:"size"`
	Sourcetemplateid      string              `json:"sourcetemplateid"`
	Sshkeyenabled         bool                `json:"sshkeyenabled"`
	Status                string              `json:"status"`
	Tags                  []Tags              `json:"tags"`
	Templatetag           string              `json:"templatetag"`
	Templatetype          string              `json:"templatetype"`
	Url                   string              `json:"url"`
	Zoneid                string              `json:"zoneid"`
	Zonename              string              `json:"zonename"`
}

func (r *RegisterTemplate) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias RegisterTemplate
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateTemplateParams struct {
	P map[string]interface{}
}

func (P *UpdateTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
	}
	if v, found := P.P["cleanupdetails"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanupdetails", vv)
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := P.P["isrouting"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrouting", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := P.P["passwordenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passwordenabled", vv)
	}
	if v, found := P.P["requireshvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("requireshvm", vv)
	}
	if v, found := P.P["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	if v, found := P.P["sshkeyenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sshkeyenabled", vv)
	}
	if v, found := P.P["templatetype"]; found {
		u.Set("templatetype", v.(string))
	}
	return u
}

func (P *UpdateTemplateParams) SetBootable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bootable"] = v
}

func (P *UpdateTemplateParams) GetBootable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bootable"].(bool)
	return value, ok
}

func (P *UpdateTemplateParams) SetCleanupdetails(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cleanupdetails"] = v
}

func (P *UpdateTemplateParams) GetCleanupdetails() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cleanupdetails"].(bool)
	return value, ok
}

func (P *UpdateTemplateParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *UpdateTemplateParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *UpdateTemplateParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *UpdateTemplateParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *UpdateTemplateParams) SetFormat(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["format"] = v
}

func (P *UpdateTemplateParams) GetFormat() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["format"].(string)
	return value, ok
}

func (P *UpdateTemplateParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateTemplateParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateTemplateParams) SetIsdynamicallyscalable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isdynamicallyscalable"] = v
}

func (P *UpdateTemplateParams) GetIsdynamicallyscalable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isdynamicallyscalable"].(bool)
	return value, ok
}

func (P *UpdateTemplateParams) SetIsrouting(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrouting"] = v
}

func (P *UpdateTemplateParams) GetIsrouting() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrouting"].(bool)
	return value, ok
}

func (P *UpdateTemplateParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateTemplateParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateTemplateParams) SetOstypeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ostypeid"] = v
}

func (P *UpdateTemplateParams) GetOstypeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ostypeid"].(string)
	return value, ok
}

func (P *UpdateTemplateParams) SetPasswordenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["passwordenabled"] = v
}

func (P *UpdateTemplateParams) GetPasswordenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["passwordenabled"].(bool)
	return value, ok
}

func (P *UpdateTemplateParams) SetRequireshvm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["requireshvm"] = v
}

func (P *UpdateTemplateParams) GetRequireshvm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["requireshvm"].(bool)
	return value, ok
}

func (P *UpdateTemplateParams) SetSortkey(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sortkey"] = v
}

func (P *UpdateTemplateParams) GetSortkey() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sortkey"].(int)
	return value, ok
}

func (P *UpdateTemplateParams) SetSshkeyenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sshkeyenabled"] = v
}

func (P *UpdateTemplateParams) GetSshkeyenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sshkeyenabled"].(bool)
	return value, ok
}

func (P *UpdateTemplateParams) SetTemplatetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templatetype"] = v
}

func (P *UpdateTemplateParams) GetTemplatetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templatetype"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewUpdateTemplateParams(id string) *UpdateTemplateParams {
	P := &UpdateTemplateParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates attributes of a template.
func (s *TemplateService) UpdateTemplate(p *UpdateTemplateParams) (*UpdateTemplateResponse, error) {
	resp, err := s.cs.newRequest("updateTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateTemplateResponse struct {
	Account               string              `json:"account"`
	Accountid             string              `json:"accountid"`
	Bits                  int                 `json:"bits"`
	Bootable              bool                `json:"bootable"`
	Checksum              string              `json:"checksum"`
	Childtemplates        []interface{}       `json:"childtemplates"`
	Created               string              `json:"created"`
	CrossZones            bool                `json:"crossZones"`
	Deployasis            bool                `json:"deployasis"`
	Deployasisdetails     map[string]string   `json:"deployasisdetails"`
	Details               map[string]string   `json:"details"`
	Directdownload        bool                `json:"directdownload"`
	Displaytext           string              `json:"displaytext"`
	Domain                string              `json:"domain"`
	Domainid              string              `json:"domainid"`
	Downloaddetails       []map[string]string `json:"downloaddetails"`
	Format                string              `json:"format"`
	Hasannotations        bool                `json:"hasannotations"`
	Hostid                string              `json:"hostid"`
	Hostname              string              `json:"hostname"`
	Hypervisor            string              `json:"hypervisor"`
	Icon                  string              `json:"icon"`
	Id                    string              `json:"id"`
	Isdynamicallyscalable bool                `json:"isdynamicallyscalable"`
	Isextractable         bool                `json:"isextractable"`
	Isfeatured            bool                `json:"isfeatured"`
	Ispublic              bool                `json:"ispublic"`
	Isready               bool                `json:"isready"`
	JobID                 string              `json:"jobid"`
	Jobstatus             int                 `json:"jobstatus"`
	Name                  string              `json:"name"`
	Ostypeid              string              `json:"ostypeid"`
	Ostypename            string              `json:"ostypename"`
	Parenttemplateid      string              `json:"parenttemplateid"`
	Passwordenabled       bool                `json:"passwordenabled"`
	Physicalsize          int64               `json:"physicalsize"`
	Project               string              `json:"project"`
	Projectid             string              `json:"projectid"`
	Removed               string              `json:"removed"`
	Requireshvm           bool                `json:"requireshvm"`
	Size                  int64               `json:"size"`
	Sourcetemplateid      string              `json:"sourcetemplateid"`
	Sshkeyenabled         bool                `json:"sshkeyenabled"`
	Status                string              `json:"status"`
	Tags                  []Tags              `json:"tags"`
	Templatetag           string              `json:"templatetag"`
	Templatetype          string              `json:"templatetype"`
	Url                   string              `json:"url"`
	Zoneid                string              `json:"zoneid"`
	Zonename              string              `json:"zonename"`
}

func (r *UpdateTemplateResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias UpdateTemplateResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateTemplatePermissionsParams struct {
	P map[string]interface{}
}

func (P *UpdateTemplatePermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["accounts"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("accounts", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["isextractable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isextractable", vv)
	}
	if v, found := P.P["isfeatured"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isfeatured", vv)
	}
	if v, found := P.P["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := P.P["op"]; found {
		u.Set("op", v.(string))
	}
	if v, found := P.P["projectids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("projectids", vv)
	}
	return u
}

func (P *UpdateTemplatePermissionsParams) SetAccounts(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accounts"] = v
}

func (P *UpdateTemplatePermissionsParams) GetAccounts() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accounts"].([]string)
	return value, ok
}

func (P *UpdateTemplatePermissionsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateTemplatePermissionsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateTemplatePermissionsParams) SetIsextractable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isextractable"] = v
}

func (P *UpdateTemplatePermissionsParams) GetIsextractable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isextractable"].(bool)
	return value, ok
}

func (P *UpdateTemplatePermissionsParams) SetIsfeatured(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isfeatured"] = v
}

func (P *UpdateTemplatePermissionsParams) GetIsfeatured() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isfeatured"].(bool)
	return value, ok
}

func (P *UpdateTemplatePermissionsParams) SetIspublic(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ispublic"] = v
}

func (P *UpdateTemplatePermissionsParams) GetIspublic() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ispublic"].(bool)
	return value, ok
}

func (P *UpdateTemplatePermissionsParams) SetOp(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["op"] = v
}

func (P *UpdateTemplatePermissionsParams) GetOp() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["op"].(string)
	return value, ok
}

func (P *UpdateTemplatePermissionsParams) SetProjectids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectids"] = v
}

func (P *UpdateTemplatePermissionsParams) GetProjectids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectids"].([]string)
	return value, ok
}

// You should always use this function to get a new UpdateTemplatePermissionsParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewUpdateTemplatePermissionsParams(id string) *UpdateTemplatePermissionsParams {
	P := &UpdateTemplatePermissionsParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a template visibility permissions. A public template is visible to all accounts within the same domain. A private template is visible only to the owner of the template. A priviledged template is a private template with account permissions added. Only accounts specified under the template permissions are visible to them.
func (s *TemplateService) UpdateTemplatePermissions(p *UpdateTemplatePermissionsParams) (*UpdateTemplatePermissionsResponse, error) {
	resp, err := s.cs.newRequest("updateTemplatePermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateTemplatePermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateTemplatePermissionsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateTemplatePermissionsResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias UpdateTemplatePermissionsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpgradeRouterTemplateParams struct {
	P map[string]interface{}
}

func (P *UpgradeRouterTemplateParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *UpgradeRouterTemplateParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *UpgradeRouterTemplateParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *UpgradeRouterTemplateParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *UpgradeRouterTemplateParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *UpgradeRouterTemplateParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UpgradeRouterTemplateParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UpgradeRouterTemplateParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpgradeRouterTemplateParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpgradeRouterTemplateParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *UpgradeRouterTemplateParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *UpgradeRouterTemplateParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *UpgradeRouterTemplateParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UpgradeRouterTemplateParams instance,
// as then you are sure you have configured all required params
func (s *TemplateService) NewUpgradeRouterTemplateParams() *UpgradeRouterTemplateParams {
	P := &UpgradeRouterTemplateParams{}
	P.P = make(map[string]interface{})
	return P
}

// Upgrades router to use newer template
func (s *TemplateService) UpgradeRouterTemplate(p *UpgradeRouterTemplateParams) (*UpgradeRouterTemplateResponse, error) {
	resp, err := s.cs.newRequest("upgradeRouterTemplate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpgradeRouterTemplateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpgradeRouterTemplateResponse struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
}
