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

type ISOServiceIface interface {
	AttachIso(p *AttachIsoParams) (*AttachIsoResponse, error)
	NewAttachIsoParams(id string, virtualmachineid string) *AttachIsoParams
	CopyIso(p *CopyIsoParams) (*CopyIsoResponse, error)
	NewCopyIsoParams(id string) *CopyIsoParams
	DeleteIso(p *DeleteIsoParams) (*DeleteIsoResponse, error)
	NewDeleteIsoParams(id string) *DeleteIsoParams
	DetachIso(p *DetachIsoParams) (*DetachIsoResponse, error)
	NewDetachIsoParams(virtualmachineid string) *DetachIsoParams
	ExtractIso(p *ExtractIsoParams) (*ExtractIsoResponse, error)
	NewExtractIsoParams(id string, mode string) *ExtractIsoParams
	ListIsoPermissions(p *ListIsoPermissionsParams) (*ListIsoPermissionsResponse, error)
	NewListIsoPermissionsParams(id string) *ListIsoPermissionsParams
	GetIsoPermissionByID(id string, opts ...OptionFunc) (*IsoPermission, int, error)
	ListIsos(p *ListIsosParams) (*ListIsosResponse, error)
	NewListIsosParams() *ListIsosParams
	GetIsoID(name string, isofilter string, zoneid string, opts ...OptionFunc) (string, int, error)
	GetIsoByName(name string, isofilter string, zoneid string, opts ...OptionFunc) (*Iso, int, error)
	GetIsoByID(id string, opts ...OptionFunc) (*Iso, int, error)
	RegisterIso(p *RegisterIsoParams) (*RegisterIsoResponse, error)
	NewRegisterIsoParams(displaytext string, name string, url string, zoneid string) *RegisterIsoParams
	UpdateIso(p *UpdateIsoParams) (*UpdateIsoResponse, error)
	NewUpdateIsoParams(id string) *UpdateIsoParams
	UpdateIsoPermissions(p *UpdateIsoPermissionsParams) (*UpdateIsoPermissionsResponse, error)
	NewUpdateIsoPermissionsParams(id string) *UpdateIsoPermissionsParams
}

type AttachIsoParams struct {
	P map[string]interface{}
}

func (P *AttachIsoParams) toURLValues() url.Values {
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
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *AttachIsoParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *AttachIsoParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *AttachIsoParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *AttachIsoParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *AttachIsoParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *AttachIsoParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new AttachIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewAttachIsoParams(id string, virtualmachineid string) *AttachIsoParams {
	P := &AttachIsoParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Attaches an ISO to a virtual machine.
func (s *ISOService) AttachIso(p *AttachIsoParams) (*AttachIsoResponse, error) {
	resp, err := s.cs.newRequest("attachIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AttachIsoResponse
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

type AttachIsoResponse struct {
	Account               string                           `json:"account"`
	Affinitygroup         []AttachIsoResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                           `json:"backupofferingid"`
	Backupofferingname    string                           `json:"backupofferingname"`
	Bootmode              string                           `json:"bootmode"`
	Boottype              string                           `json:"boottype"`
	Cpunumber             int                              `json:"cpunumber"`
	Cpuspeed              int                              `json:"cpuspeed"`
	Cpuused               string                           `json:"cpuused"`
	Created               string                           `json:"created"`
	Details               map[string]string                `json:"details"`
	Diskioread            int64                            `json:"diskioread"`
	Diskiowrite           int64                            `json:"diskiowrite"`
	Diskkbsread           int64                            `json:"diskkbsread"`
	Diskkbswrite          int64                            `json:"diskkbswrite"`
	Diskofferingid        string                           `json:"diskofferingid"`
	Diskofferingname      string                           `json:"diskofferingname"`
	Displayname           string                           `json:"displayname"`
	Displayvm             bool                             `json:"displayvm"`
	Domain                string                           `json:"domain"`
	Domainid              string                           `json:"domainid"`
	Forvirtualnetwork     bool                             `json:"forvirtualnetwork"`
	Group                 string                           `json:"group"`
	Groupid               string                           `json:"groupid"`
	Guestosid             string                           `json:"guestosid"`
	Haenable              bool                             `json:"haenable"`
	Hasannotations        bool                             `json:"hasannotations"`
	Hostid                string                           `json:"hostid"`
	Hostname              string                           `json:"hostname"`
	Hypervisor            string                           `json:"hypervisor"`
	Icon                  string                           `json:"icon"`
	Id                    string                           `json:"id"`
	Instancename          string                           `json:"instancename"`
	Isdynamicallyscalable bool                             `json:"isdynamicallyscalable"`
	Isodisplaytext        string                           `json:"isodisplaytext"`
	Isoid                 string                           `json:"isoid"`
	Isoname               string                           `json:"isoname"`
	JobID                 string                           `json:"jobid"`
	Jobstatus             int                              `json:"jobstatus"`
	Keypair               string                           `json:"keypair"`
	Lastupdated           string                           `json:"lastupdated"`
	Memory                int                              `json:"memory"`
	Memoryintfreekbs      int64                            `json:"memoryintfreekbs"`
	Memorykbs             int64                            `json:"memorykbs"`
	Memorytargetkbs       int64                            `json:"memorytargetkbs"`
	Name                  string                           `json:"name"`
	Networkkbsread        int64                            `json:"networkkbsread"`
	Networkkbswrite       int64                            `json:"networkkbswrite"`
	Nic                   []Nic                            `json:"nic"`
	Osdisplayname         string                           `json:"osdisplayname"`
	Ostypeid              string                           `json:"ostypeid"`
	Password              string                           `json:"password"`
	Passwordenabled       bool                             `json:"passwordenabled"`
	Pooltype              string                           `json:"pooltype"`
	Project               string                           `json:"project"`
	Projectid             string                           `json:"projectid"`
	Publicip              string                           `json:"publicip"`
	Publicipid            string                           `json:"publicipid"`
	Readonlydetails       string                           `json:"readonlydetails"`
	Receivedbytes         int64                            `json:"receivedbytes"`
	Rootdeviceid          int64                            `json:"rootdeviceid"`
	Rootdevicetype        string                           `json:"rootdevicetype"`
	Securitygroup         []AttachIsoResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                            `json:"sentbytes"`
	Serviceofferingid     string                           `json:"serviceofferingid"`
	Serviceofferingname   string                           `json:"serviceofferingname"`
	Servicestate          string                           `json:"servicestate"`
	State                 string                           `json:"state"`
	Tags                  []Tags                           `json:"tags"`
	Templatedisplaytext   string                           `json:"templatedisplaytext"`
	Templateid            string                           `json:"templateid"`
	Templatename          string                           `json:"templatename"`
	Userid                string                           `json:"userid"`
	Username              string                           `json:"username"`
	Vgpu                  string                           `json:"vgpu"`
	Zoneid                string                           `json:"zoneid"`
	Zonename              string                           `json:"zonename"`
}

type AttachIsoResponseSecuritygroup struct {
	Account             string                               `json:"account"`
	Description         string                               `json:"description"`
	Domain              string                               `json:"domain"`
	Domainid            string                               `json:"domainid"`
	Egressrule          []AttachIsoResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                               `json:"id"`
	Ingressrule         []AttachIsoResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                               `json:"name"`
	Project             string                               `json:"project"`
	Projectid           string                               `json:"projectid"`
	Tags                []Tags                               `json:"tags"`
	Virtualmachinecount int                                  `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                        `json:"virtualmachineids"`
}

type AttachIsoResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type AttachIsoResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *AttachIsoResponse) UnmarshalJSON(b []byte) error {
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

	type alias AttachIsoResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CopyIsoParams struct {
	P map[string]interface{}
}

func (P *CopyIsoParams) toURLValues() url.Values {
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

func (P *CopyIsoParams) SetDestzoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["destzoneid"] = v
}

func (P *CopyIsoParams) GetDestzoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["destzoneid"].(string)
	return value, ok
}

func (P *CopyIsoParams) SetDestzoneids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["destzoneids"] = v
}

func (P *CopyIsoParams) GetDestzoneids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["destzoneids"].([]string)
	return value, ok
}

func (P *CopyIsoParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *CopyIsoParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *CopyIsoParams) SetSourcezoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sourcezoneid"] = v
}

func (P *CopyIsoParams) GetSourcezoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sourcezoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CopyIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewCopyIsoParams(id string) *CopyIsoParams {
	P := &CopyIsoParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Copies an iso from one zone to another.
func (s *ISOService) CopyIso(p *CopyIsoParams) (*CopyIsoResponse, error) {
	resp, err := s.cs.newRequest("copyIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CopyIsoResponse
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

type CopyIsoResponse struct {
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

func (r *CopyIsoResponse) UnmarshalJSON(b []byte) error {
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

	type alias CopyIsoResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteIsoParams struct {
	P map[string]interface{}
}

func (P *DeleteIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *DeleteIsoParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteIsoParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *DeleteIsoParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *DeleteIsoParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewDeleteIsoParams(id string) *DeleteIsoParams {
	P := &DeleteIsoParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes an ISO file.
func (s *ISOService) DeleteIso(p *DeleteIsoParams) (*DeleteIsoResponse, error) {
	resp, err := s.cs.newRequest("deleteIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteIsoResponse
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

type DeleteIsoResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DetachIsoParams struct {
	P map[string]interface{}
}

func (P *DetachIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *DetachIsoParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *DetachIsoParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *DetachIsoParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *DetachIsoParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new DetachIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewDetachIsoParams(virtualmachineid string) *DetachIsoParams {
	P := &DetachIsoParams{}
	P.P = make(map[string]interface{})
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Detaches any ISO file (if any) currently attached to a virtual machine.
func (s *ISOService) DetachIso(p *DetachIsoParams) (*DetachIsoResponse, error) {
	resp, err := s.cs.newRequest("detachIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DetachIsoResponse
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

type DetachIsoResponse struct {
	Account               string                           `json:"account"`
	Affinitygroup         []DetachIsoResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                           `json:"backupofferingid"`
	Backupofferingname    string                           `json:"backupofferingname"`
	Bootmode              string                           `json:"bootmode"`
	Boottype              string                           `json:"boottype"`
	Cpunumber             int                              `json:"cpunumber"`
	Cpuspeed              int                              `json:"cpuspeed"`
	Cpuused               string                           `json:"cpuused"`
	Created               string                           `json:"created"`
	Details               map[string]string                `json:"details"`
	Diskioread            int64                            `json:"diskioread"`
	Diskiowrite           int64                            `json:"diskiowrite"`
	Diskkbsread           int64                            `json:"diskkbsread"`
	Diskkbswrite          int64                            `json:"diskkbswrite"`
	Diskofferingid        string                           `json:"diskofferingid"`
	Diskofferingname      string                           `json:"diskofferingname"`
	Displayname           string                           `json:"displayname"`
	Displayvm             bool                             `json:"displayvm"`
	Domain                string                           `json:"domain"`
	Domainid              string                           `json:"domainid"`
	Forvirtualnetwork     bool                             `json:"forvirtualnetwork"`
	Group                 string                           `json:"group"`
	Groupid               string                           `json:"groupid"`
	Guestosid             string                           `json:"guestosid"`
	Haenable              bool                             `json:"haenable"`
	Hasannotations        bool                             `json:"hasannotations"`
	Hostid                string                           `json:"hostid"`
	Hostname              string                           `json:"hostname"`
	Hypervisor            string                           `json:"hypervisor"`
	Icon                  string                           `json:"icon"`
	Id                    string                           `json:"id"`
	Instancename          string                           `json:"instancename"`
	Isdynamicallyscalable bool                             `json:"isdynamicallyscalable"`
	Isodisplaytext        string                           `json:"isodisplaytext"`
	Isoid                 string                           `json:"isoid"`
	Isoname               string                           `json:"isoname"`
	JobID                 string                           `json:"jobid"`
	Jobstatus             int                              `json:"jobstatus"`
	Keypair               string                           `json:"keypair"`
	Lastupdated           string                           `json:"lastupdated"`
	Memory                int                              `json:"memory"`
	Memoryintfreekbs      int64                            `json:"memoryintfreekbs"`
	Memorykbs             int64                            `json:"memorykbs"`
	Memorytargetkbs       int64                            `json:"memorytargetkbs"`
	Name                  string                           `json:"name"`
	Networkkbsread        int64                            `json:"networkkbsread"`
	Networkkbswrite       int64                            `json:"networkkbswrite"`
	Nic                   []Nic                            `json:"nic"`
	Osdisplayname         string                           `json:"osdisplayname"`
	Ostypeid              string                           `json:"ostypeid"`
	Password              string                           `json:"password"`
	Passwordenabled       bool                             `json:"passwordenabled"`
	Pooltype              string                           `json:"pooltype"`
	Project               string                           `json:"project"`
	Projectid             string                           `json:"projectid"`
	Publicip              string                           `json:"publicip"`
	Publicipid            string                           `json:"publicipid"`
	Readonlydetails       string                           `json:"readonlydetails"`
	Receivedbytes         int64                            `json:"receivedbytes"`
	Rootdeviceid          int64                            `json:"rootdeviceid"`
	Rootdevicetype        string                           `json:"rootdevicetype"`
	Securitygroup         []DetachIsoResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                            `json:"sentbytes"`
	Serviceofferingid     string                           `json:"serviceofferingid"`
	Serviceofferingname   string                           `json:"serviceofferingname"`
	Servicestate          string                           `json:"servicestate"`
	State                 string                           `json:"state"`
	Tags                  []Tags                           `json:"tags"`
	Templatedisplaytext   string                           `json:"templatedisplaytext"`
	Templateid            string                           `json:"templateid"`
	Templatename          string                           `json:"templatename"`
	Userid                string                           `json:"userid"`
	Username              string                           `json:"username"`
	Vgpu                  string                           `json:"vgpu"`
	Zoneid                string                           `json:"zoneid"`
	Zonename              string                           `json:"zonename"`
}

type DetachIsoResponseSecuritygroup struct {
	Account             string                               `json:"account"`
	Description         string                               `json:"description"`
	Domain              string                               `json:"domain"`
	Domainid            string                               `json:"domainid"`
	Egressrule          []DetachIsoResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                               `json:"id"`
	Ingressrule         []DetachIsoResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                               `json:"name"`
	Project             string                               `json:"project"`
	Projectid           string                               `json:"projectid"`
	Tags                []Tags                               `json:"tags"`
	Virtualmachinecount int                                  `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                        `json:"virtualmachineids"`
}

type DetachIsoResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type DetachIsoResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *DetachIsoResponse) UnmarshalJSON(b []byte) error {
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

	type alias DetachIsoResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ExtractIsoParams struct {
	P map[string]interface{}
}

func (P *ExtractIsoParams) toURLValues() url.Values {
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

func (P *ExtractIsoParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ExtractIsoParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ExtractIsoParams) SetMode(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["mode"] = v
}

func (P *ExtractIsoParams) GetMode() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["mode"].(string)
	return value, ok
}

func (P *ExtractIsoParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *ExtractIsoParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *ExtractIsoParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ExtractIsoParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ExtractIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewExtractIsoParams(id string, mode string) *ExtractIsoParams {
	P := &ExtractIsoParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["mode"] = mode
	return P
}

// Extracts an ISO
func (s *ISOService) ExtractIso(p *ExtractIsoParams) (*ExtractIsoResponse, error) {
	resp, err := s.cs.newRequest("extractIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExtractIsoResponse
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

type ExtractIsoResponse struct {
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

type ListIsoPermissionsParams struct {
	P map[string]interface{}
}

func (P *ListIsoPermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ListIsoPermissionsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListIsoPermissionsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ListIsoPermissionsParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewListIsoPermissionsParams(id string) *ListIsoPermissionsParams {
	P := &ListIsoPermissionsParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoPermissionByID(id string, opts ...OptionFunc) (*IsoPermission, int, error) {
	P := &ListIsoPermissionsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListIsoPermissions(P)
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
		return l.IsoPermissions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for IsoPermission UUID: %s!", id)
}

// List iso visibility and all accounts that have permissions to view this iso.
func (s *ISOService) ListIsoPermissions(p *ListIsoPermissionsParams) (*ListIsoPermissionsResponse, error) {
	resp, err := s.cs.newRequest("listIsoPermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListIsoPermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListIsoPermissionsResponse struct {
	Count          int              `json:"count"`
	IsoPermissions []*IsoPermission `json:"isopermission"`
}

type IsoPermission struct {
	Account    []string `json:"account"`
	Domainid   string   `json:"domainid"`
	Id         string   `json:"id"`
	Ispublic   bool     `json:"ispublic"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Projectids []string `json:"projectids"`
}

type ListIsosParams struct {
	P map[string]interface{}
}

func (P *ListIsosParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
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
	if v, found := P.P["isofilter"]; found {
		u.Set("isofilter", v.(string))
	}
	if v, found := P.P["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := P.P["isready"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isready", vv)
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
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListIsosParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListIsosParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListIsosParams) SetBootable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bootable"] = v
}

func (P *ListIsosParams) GetBootable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bootable"].(bool)
	return value, ok
}

func (P *ListIsosParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListIsosParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListIsosParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *ListIsosParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *ListIsosParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListIsosParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListIsosParams) SetIsofilter(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isofilter"] = v
}

func (P *ListIsosParams) GetIsofilter() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isofilter"].(string)
	return value, ok
}

func (P *ListIsosParams) SetIspublic(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ispublic"] = v
}

func (P *ListIsosParams) GetIspublic() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ispublic"].(bool)
	return value, ok
}

func (P *ListIsosParams) SetIsready(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isready"] = v
}

func (P *ListIsosParams) GetIsready() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isready"].(bool)
	return value, ok
}

func (P *ListIsosParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListIsosParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListIsosParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListIsosParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListIsosParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListIsosParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListIsosParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListIsosParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListIsosParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListIsosParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListIsosParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListIsosParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListIsosParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListIsosParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListIsosParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListIsosParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

func (P *ListIsosParams) SetShowremoved(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showremoved"] = v
}

func (P *ListIsosParams) GetShowremoved() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showremoved"].(bool)
	return value, ok
}

func (P *ListIsosParams) SetShowunique(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showunique"] = v
}

func (P *ListIsosParams) GetShowunique() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showunique"].(bool)
	return value, ok
}

func (P *ListIsosParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListIsosParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListIsosParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListIsosParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListIsosParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewListIsosParams() *ListIsosParams {
	P := &ListIsosParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoID(name string, isofilter string, zoneid string, opts ...OptionFunc) (string, int, error) {
	P := &ListIsosParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name
	P.P["isofilter"] = isofilter
	P.P["zoneid"] = zoneid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListIsos(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Isos[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Isos {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoByName(name string, isofilter string, zoneid string, opts ...OptionFunc) (*Iso, int, error) {
	id, count, err := s.GetIsoID(name, isofilter, zoneid, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetIsoByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ISOService) GetIsoByID(id string, opts ...OptionFunc) (*Iso, int, error) {
	P := &ListIsosParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListIsos(P)
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
		return l.Isos[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Iso UUID: %s!", id)
}

// Lists all available ISO files.
func (s *ISOService) ListIsos(p *ListIsosParams) (*ListIsosResponse, error) {
	resp, err := s.cs.newRequest("listIsos", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListIsosResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListIsosResponse struct {
	Count int    `json:"count"`
	Isos  []*Iso `json:"iso"`
}

type Iso struct {
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

func (r *Iso) UnmarshalJSON(b []byte) error {
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

	type alias Iso
	return json.Unmarshal(b, (*alias)(r))
}

type RegisterIsoParams struct {
	P map[string]interface{}
}

func (P *RegisterIsoParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["bootable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootable", vv)
	}
	if v, found := P.P["checksum"]; found {
		u.Set("checksum", v.(string))
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
	if v, found := P.P["imagestoreuuid"]; found {
		u.Set("imagestoreuuid", v.(string))
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
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *RegisterIsoParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *RegisterIsoParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *RegisterIsoParams) SetBootable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bootable"] = v
}

func (P *RegisterIsoParams) GetBootable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bootable"].(bool)
	return value, ok
}

func (P *RegisterIsoParams) SetChecksum(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["checksum"] = v
}

func (P *RegisterIsoParams) GetChecksum() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["checksum"].(string)
	return value, ok
}

func (P *RegisterIsoParams) SetDirectdownload(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["directdownload"] = v
}

func (P *RegisterIsoParams) GetDirectdownload() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["directdownload"].(bool)
	return value, ok
}

func (P *RegisterIsoParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *RegisterIsoParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *RegisterIsoParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *RegisterIsoParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *RegisterIsoParams) SetImagestoreuuid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["imagestoreuuid"] = v
}

func (P *RegisterIsoParams) GetImagestoreuuid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["imagestoreuuid"].(string)
	return value, ok
}

func (P *RegisterIsoParams) SetIsdynamicallyscalable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isdynamicallyscalable"] = v
}

func (P *RegisterIsoParams) GetIsdynamicallyscalable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isdynamicallyscalable"].(bool)
	return value, ok
}

func (P *RegisterIsoParams) SetIsextractable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isextractable"] = v
}

func (P *RegisterIsoParams) GetIsextractable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isextractable"].(bool)
	return value, ok
}

func (P *RegisterIsoParams) SetIsfeatured(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isfeatured"] = v
}

func (P *RegisterIsoParams) GetIsfeatured() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isfeatured"].(bool)
	return value, ok
}

func (P *RegisterIsoParams) SetIspublic(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ispublic"] = v
}

func (P *RegisterIsoParams) GetIspublic() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ispublic"].(bool)
	return value, ok
}

func (P *RegisterIsoParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *RegisterIsoParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *RegisterIsoParams) SetOstypeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ostypeid"] = v
}

func (P *RegisterIsoParams) GetOstypeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ostypeid"].(string)
	return value, ok
}

func (P *RegisterIsoParams) SetPasswordenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["passwordenabled"] = v
}

func (P *RegisterIsoParams) GetPasswordenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["passwordenabled"].(bool)
	return value, ok
}

func (P *RegisterIsoParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *RegisterIsoParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *RegisterIsoParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *RegisterIsoParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *RegisterIsoParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *RegisterIsoParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new RegisterIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewRegisterIsoParams(displaytext string, name string, url string, zoneid string) *RegisterIsoParams {
	P := &RegisterIsoParams{}
	P.P = make(map[string]interface{})
	P.P["displaytext"] = displaytext
	P.P["name"] = name
	P.P["url"] = url
	P.P["zoneid"] = zoneid
	return P
}

// Registers an existing ISO into the CloudStack Cloud.
func (s *ISOService) RegisterIso(p *RegisterIsoParams) (*RegisterIsoResponse, error) {
	resp, err := s.cs.newRequest("registerIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r RegisterIsoResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterIsoResponse struct {
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

func (r *RegisterIsoResponse) UnmarshalJSON(b []byte) error {
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

	type alias RegisterIsoResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateIsoParams struct {
	P map[string]interface{}
}

func (P *UpdateIsoParams) toURLValues() url.Values {
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
	return u
}

func (P *UpdateIsoParams) SetBootable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bootable"] = v
}

func (P *UpdateIsoParams) GetBootable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bootable"].(bool)
	return value, ok
}

func (P *UpdateIsoParams) SetCleanupdetails(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cleanupdetails"] = v
}

func (P *UpdateIsoParams) GetCleanupdetails() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cleanupdetails"].(bool)
	return value, ok
}

func (P *UpdateIsoParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *UpdateIsoParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *UpdateIsoParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *UpdateIsoParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *UpdateIsoParams) SetFormat(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["format"] = v
}

func (P *UpdateIsoParams) GetFormat() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["format"].(string)
	return value, ok
}

func (P *UpdateIsoParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateIsoParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateIsoParams) SetIsdynamicallyscalable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isdynamicallyscalable"] = v
}

func (P *UpdateIsoParams) GetIsdynamicallyscalable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isdynamicallyscalable"].(bool)
	return value, ok
}

func (P *UpdateIsoParams) SetIsrouting(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrouting"] = v
}

func (P *UpdateIsoParams) GetIsrouting() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrouting"].(bool)
	return value, ok
}

func (P *UpdateIsoParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateIsoParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateIsoParams) SetOstypeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ostypeid"] = v
}

func (P *UpdateIsoParams) GetOstypeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ostypeid"].(string)
	return value, ok
}

func (P *UpdateIsoParams) SetPasswordenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["passwordenabled"] = v
}

func (P *UpdateIsoParams) GetPasswordenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["passwordenabled"].(bool)
	return value, ok
}

func (P *UpdateIsoParams) SetRequireshvm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["requireshvm"] = v
}

func (P *UpdateIsoParams) GetRequireshvm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["requireshvm"].(bool)
	return value, ok
}

func (P *UpdateIsoParams) SetSortkey(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sortkey"] = v
}

func (P *UpdateIsoParams) GetSortkey() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sortkey"].(int)
	return value, ok
}

func (P *UpdateIsoParams) SetSshkeyenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sshkeyenabled"] = v
}

func (P *UpdateIsoParams) GetSshkeyenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sshkeyenabled"].(bool)
	return value, ok
}

// You should always use this function to get a new UpdateIsoParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewUpdateIsoParams(id string) *UpdateIsoParams {
	P := &UpdateIsoParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates an ISO file.
func (s *ISOService) UpdateIso(p *UpdateIsoParams) (*UpdateIsoResponse, error) {
	resp, err := s.cs.newRequest("updateIso", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateIsoResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateIsoResponse struct {
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

func (r *UpdateIsoResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateIsoResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateIsoPermissionsParams struct {
	P map[string]interface{}
}

func (P *UpdateIsoPermissionsParams) toURLValues() url.Values {
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

func (P *UpdateIsoPermissionsParams) SetAccounts(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accounts"] = v
}

func (P *UpdateIsoPermissionsParams) GetAccounts() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accounts"].([]string)
	return value, ok
}

func (P *UpdateIsoPermissionsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateIsoPermissionsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateIsoPermissionsParams) SetIsextractable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isextractable"] = v
}

func (P *UpdateIsoPermissionsParams) GetIsextractable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isextractable"].(bool)
	return value, ok
}

func (P *UpdateIsoPermissionsParams) SetIsfeatured(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isfeatured"] = v
}

func (P *UpdateIsoPermissionsParams) GetIsfeatured() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isfeatured"].(bool)
	return value, ok
}

func (P *UpdateIsoPermissionsParams) SetIspublic(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ispublic"] = v
}

func (P *UpdateIsoPermissionsParams) GetIspublic() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ispublic"].(bool)
	return value, ok
}

func (P *UpdateIsoPermissionsParams) SetOp(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["op"] = v
}

func (P *UpdateIsoPermissionsParams) GetOp() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["op"].(string)
	return value, ok
}

func (P *UpdateIsoPermissionsParams) SetProjectids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectids"] = v
}

func (P *UpdateIsoPermissionsParams) GetProjectids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectids"].([]string)
	return value, ok
}

// You should always use this function to get a new UpdateIsoPermissionsParams instance,
// as then you are sure you have configured all required params
func (s *ISOService) NewUpdateIsoPermissionsParams(id string) *UpdateIsoPermissionsParams {
	P := &UpdateIsoPermissionsParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates ISO permissions
func (s *ISOService) UpdateIsoPermissions(p *UpdateIsoPermissionsParams) (*UpdateIsoPermissionsResponse, error) {
	resp, err := s.cs.newRequest("updateIsoPermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateIsoPermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateIsoPermissionsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateIsoPermissionsResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateIsoPermissionsResponse
	return json.Unmarshal(b, (*alias)(r))
}
