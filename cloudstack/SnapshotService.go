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

type SnapshotServiceIface interface {
	CreateSnapshot(p *CreateSnapshotParams) (*CreateSnapshotResponse, error)
	NewCreateSnapshotParams(volumeid string) *CreateSnapshotParams
	CreateSnapshotPolicy(p *CreateSnapshotPolicyParams) (*CreateSnapshotPolicyResponse, error)
	NewCreateSnapshotPolicyParams(intervaltype string, maxsnaps int, schedule string, timezone string, volumeid string) *CreateSnapshotPolicyParams
	CreateVMSnapshot(p *CreateVMSnapshotParams) (*CreateVMSnapshotResponse, error)
	NewCreateVMSnapshotParams(virtualmachineid string) *CreateVMSnapshotParams
	DeleteSnapshot(p *DeleteSnapshotParams) (*DeleteSnapshotResponse, error)
	NewDeleteSnapshotParams(id string) *DeleteSnapshotParams
	DeleteSnapshotPolicies(p *DeleteSnapshotPoliciesParams) (*DeleteSnapshotPoliciesResponse, error)
	NewDeleteSnapshotPoliciesParams() *DeleteSnapshotPoliciesParams
	DeleteVMSnapshot(p *DeleteVMSnapshotParams) (*DeleteVMSnapshotResponse, error)
	NewDeleteVMSnapshotParams(vmsnapshotid string) *DeleteVMSnapshotParams
	ListSnapshotPolicies(p *ListSnapshotPoliciesParams) (*ListSnapshotPoliciesResponse, error)
	NewListSnapshotPoliciesParams() *ListSnapshotPoliciesParams
	GetSnapshotPolicyByID(id string, opts ...OptionFunc) (*SnapshotPolicy, int, error)
	ListSnapshots(p *ListSnapshotsParams) (*ListSnapshotsResponse, error)
	NewListSnapshotsParams() *ListSnapshotsParams
	GetSnapshotID(name string, opts ...OptionFunc) (string, int, error)
	GetSnapshotByName(name string, opts ...OptionFunc) (*Snapshot, int, error)
	GetSnapshotByID(id string, opts ...OptionFunc) (*Snapshot, int, error)
	ListVMSnapshot(p *ListVMSnapshotParams) (*ListVMSnapshotResponse, error)
	NewListVMSnapshotParams() *ListVMSnapshotParams
	GetVMSnapshotID(name string, opts ...OptionFunc) (string, int, error)
	RevertSnapshot(p *RevertSnapshotParams) (*RevertSnapshotResponse, error)
	NewRevertSnapshotParams(id string) *RevertSnapshotParams
	RevertToVMSnapshot(p *RevertToVMSnapshotParams) (*RevertToVMSnapshotResponse, error)
	NewRevertToVMSnapshotParams(vmsnapshotid string) *RevertToVMSnapshotParams
	UpdateSnapshotPolicy(p *UpdateSnapshotPolicyParams) (*UpdateSnapshotPolicyResponse, error)
	NewUpdateSnapshotPolicyParams() *UpdateSnapshotPolicyParams
}

type CreateSnapshotParams struct {
	P map[string]interface{}
}

func (P *CreateSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["asyncbackup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("asyncbackup", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["locationtype"]; found {
		u.Set("locationtype", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["policyid"]; found {
		u.Set("policyid", v.(string))
	}
	if v, found := P.P["quiescevm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("quiescevm", vv)
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (P *CreateSnapshotParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateSnapshotParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateSnapshotParams) SetAsyncbackup(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["asyncbackup"] = v
}

func (P *CreateSnapshotParams) GetAsyncbackup() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["asyncbackup"].(bool)
	return value, ok
}

func (P *CreateSnapshotParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateSnapshotParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateSnapshotParams) SetLocationtype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["locationtype"] = v
}

func (P *CreateSnapshotParams) GetLocationtype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["locationtype"].(string)
	return value, ok
}

func (P *CreateSnapshotParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateSnapshotParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateSnapshotParams) SetPolicyid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["policyid"] = v
}

func (P *CreateSnapshotParams) GetPolicyid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["policyid"].(string)
	return value, ok
}

func (P *CreateSnapshotParams) SetQuiescevm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["quiescevm"] = v
}

func (P *CreateSnapshotParams) GetQuiescevm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["quiescevm"].(bool)
	return value, ok
}

func (P *CreateSnapshotParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *CreateSnapshotParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *CreateSnapshotParams) SetVolumeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["volumeid"] = v
}

func (P *CreateSnapshotParams) GetVolumeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCreateSnapshotParams(volumeid string) *CreateSnapshotParams {
	P := &CreateSnapshotParams{}
	P.P = make(map[string]interface{})
	P.P["volumeid"] = volumeid
	return P
}

// Creates an instant snapshot of a volume.
func (s *SnapshotService) CreateSnapshot(p *CreateSnapshotParams) (*CreateSnapshotResponse, error) {
	resp, err := s.cs.newRequest("createSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSnapshotResponse
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

type CreateSnapshotResponse struct {
	Account        string `json:"account"`
	Created        string `json:"created"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	Intervaltype   string `json:"intervaltype"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Locationtype   string `json:"locationtype"`
	Name           string `json:"name"`
	Osdisplayname  string `json:"osdisplayname"`
	Ostypeid       string `json:"ostypeid"`
	Physicalsize   int64  `json:"physicalsize"`
	Project        string `json:"project"`
	Projectid      string `json:"projectid"`
	Revertable     bool   `json:"revertable"`
	Snapshottype   string `json:"snapshottype"`
	State          string `json:"state"`
	Tags           []Tags `json:"tags"`
	Virtualsize    int64  `json:"virtualsize"`
	Volumeid       string `json:"volumeid"`
	Volumename     string `json:"volumename"`
	Volumetype     string `json:"volumetype"`
	Zoneid         string `json:"zoneid"`
}

func (r *CreateSnapshotResponse) UnmarshalJSON(b []byte) error {
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

	type alias CreateSnapshotResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CreateSnapshotPolicyParams struct {
	P map[string]interface{}
}

func (P *CreateSnapshotPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["intervaltype"]; found {
		u.Set("intervaltype", v.(string))
	}
	if v, found := P.P["maxsnaps"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxsnaps", vv)
	}
	if v, found := P.P["schedule"]; found {
		u.Set("schedule", v.(string))
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := P.P["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (P *CreateSnapshotPolicyParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateSnapshotPolicyParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateSnapshotPolicyParams) SetIntervaltype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["intervaltype"] = v
}

func (P *CreateSnapshotPolicyParams) GetIntervaltype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["intervaltype"].(string)
	return value, ok
}

func (P *CreateSnapshotPolicyParams) SetMaxsnaps(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxsnaps"] = v
}

func (P *CreateSnapshotPolicyParams) GetMaxsnaps() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxsnaps"].(int)
	return value, ok
}

func (P *CreateSnapshotPolicyParams) SetSchedule(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["schedule"] = v
}

func (P *CreateSnapshotPolicyParams) GetSchedule() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["schedule"].(string)
	return value, ok
}

func (P *CreateSnapshotPolicyParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *CreateSnapshotPolicyParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *CreateSnapshotPolicyParams) SetTimezone(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["timezone"] = v
}

func (P *CreateSnapshotPolicyParams) GetTimezone() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["timezone"].(string)
	return value, ok
}

func (P *CreateSnapshotPolicyParams) SetVolumeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["volumeid"] = v
}

func (P *CreateSnapshotPolicyParams) GetVolumeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateSnapshotPolicyParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCreateSnapshotPolicyParams(intervaltype string, maxsnaps int, schedule string, timezone string, volumeid string) *CreateSnapshotPolicyParams {
	P := &CreateSnapshotPolicyParams{}
	P.P = make(map[string]interface{})
	P.P["intervaltype"] = intervaltype
	P.P["maxsnaps"] = maxsnaps
	P.P["schedule"] = schedule
	P.P["timezone"] = timezone
	P.P["volumeid"] = volumeid
	return P
}

// Creates a snapshot policy for the account.
func (s *SnapshotService) CreateSnapshotPolicy(p *CreateSnapshotPolicyParams) (*CreateSnapshotPolicyResponse, error) {
	resp, err := s.cs.newRequest("createSnapshotPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSnapshotPolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateSnapshotPolicyResponse struct {
	Fordisplay     bool   `json:"fordisplay"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	Intervaltype   int    `json:"intervaltype"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Maxsnaps       int    `json:"maxsnaps"`
	Schedule       string `json:"schedule"`
	Tags           []Tags `json:"tags"`
	Timezone       string `json:"timezone"`
	Volumeid       string `json:"volumeid"`
}

type CreateVMSnapshotParams struct {
	P map[string]interface{}
}

func (P *CreateVMSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["quiescevm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("quiescevm", vv)
	}
	if v, found := P.P["snapshotmemory"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("snapshotmemory", vv)
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *CreateVMSnapshotParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateVMSnapshotParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateVMSnapshotParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateVMSnapshotParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateVMSnapshotParams) SetQuiescevm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["quiescevm"] = v
}

func (P *CreateVMSnapshotParams) GetQuiescevm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["quiescevm"].(bool)
	return value, ok
}

func (P *CreateVMSnapshotParams) SetSnapshotmemory(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["snapshotmemory"] = v
}

func (P *CreateVMSnapshotParams) GetSnapshotmemory() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["snapshotmemory"].(bool)
	return value, ok
}

func (P *CreateVMSnapshotParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *CreateVMSnapshotParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewCreateVMSnapshotParams(virtualmachineid string) *CreateVMSnapshotParams {
	P := &CreateVMSnapshotParams{}
	P.P = make(map[string]interface{})
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Creates snapshot for a vm.
func (s *SnapshotService) CreateVMSnapshot(p *CreateVMSnapshotParams) (*CreateVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("createVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVMSnapshotResponse
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

type CreateVMSnapshotResponse struct {
	Account            string `json:"account"`
	Created            string `json:"created"`
	Current            bool   `json:"current"`
	Description        string `json:"description"`
	Displayname        string `json:"displayname"`
	Domain             string `json:"domain"`
	Domainid           string `json:"domainid"`
	Hasannotations     bool   `json:"hasannotations"`
	Hypervisor         string `json:"hypervisor"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Name               string `json:"name"`
	Parent             string `json:"parent"`
	ParentName         string `json:"parentName"`
	Project            string `json:"project"`
	Projectid          string `json:"projectid"`
	State              string `json:"state"`
	Tags               []Tags `json:"tags"`
	Type               string `json:"type"`
	Virtualmachineid   string `json:"virtualmachineid"`
	Virtualmachinename string `json:"virtualmachinename"`
	Zoneid             string `json:"zoneid"`
	Zonename           string `json:"zonename"`
}

type DeleteSnapshotParams struct {
	P map[string]interface{}
}

func (P *DeleteSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteSnapshotParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteSnapshotParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewDeleteSnapshotParams(id string) *DeleteSnapshotParams {
	P := &DeleteSnapshotParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a snapshot of a disk volume.
func (s *SnapshotService) DeleteSnapshot(p *DeleteSnapshotParams) (*DeleteSnapshotResponse, error) {
	resp, err := s.cs.newRequest("deleteSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSnapshotResponse
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

type DeleteSnapshotResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteSnapshotPoliciesParams struct {
	P map[string]interface{}
}

func (P *DeleteSnapshotPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	return u
}

func (P *DeleteSnapshotPoliciesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteSnapshotPoliciesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *DeleteSnapshotPoliciesParams) SetIds(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ids"] = v
}

func (P *DeleteSnapshotPoliciesParams) GetIds() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ids"].([]string)
	return value, ok
}

// You should always use this function to get a new DeleteSnapshotPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewDeleteSnapshotPoliciesParams() *DeleteSnapshotPoliciesParams {
	P := &DeleteSnapshotPoliciesParams{}
	P.P = make(map[string]interface{})
	return P
}

// Deletes snapshot policies for the account.
func (s *SnapshotService) DeleteSnapshotPolicies(p *DeleteSnapshotPoliciesParams) (*DeleteSnapshotPoliciesResponse, error) {
	resp, err := s.cs.newRequest("deleteSnapshotPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSnapshotPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSnapshotPoliciesResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSnapshotPoliciesResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteSnapshotPoliciesResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteVMSnapshotParams struct {
	P map[string]interface{}
}

func (P *DeleteVMSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["vmsnapshotid"]; found {
		u.Set("vmsnapshotid", v.(string))
	}
	return u
}

func (P *DeleteVMSnapshotParams) SetVmsnapshotid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmsnapshotid"] = v
}

func (P *DeleteVMSnapshotParams) GetVmsnapshotid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmsnapshotid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewDeleteVMSnapshotParams(vmsnapshotid string) *DeleteVMSnapshotParams {
	P := &DeleteVMSnapshotParams{}
	P.P = make(map[string]interface{})
	P.P["vmsnapshotid"] = vmsnapshotid
	return P
}

// Deletes a vmsnapshot.
func (s *SnapshotService) DeleteVMSnapshot(p *DeleteVMSnapshotParams) (*DeleteVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("deleteVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVMSnapshotResponse
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

type DeleteVMSnapshotResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListSnapshotPoliciesParams struct {
	P map[string]interface{}
}

func (P *ListSnapshotPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := P.P["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (P *ListSnapshotPoliciesParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListSnapshotPoliciesParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListSnapshotPoliciesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListSnapshotPoliciesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListSnapshotPoliciesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListSnapshotPoliciesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListSnapshotPoliciesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListSnapshotPoliciesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListSnapshotPoliciesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListSnapshotPoliciesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListSnapshotPoliciesParams) SetVolumeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["volumeid"] = v
}

func (P *ListSnapshotPoliciesParams) GetVolumeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSnapshotPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewListSnapshotPoliciesParams() *ListSnapshotPoliciesParams {
	P := &ListSnapshotPoliciesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotPolicyByID(id string, opts ...OptionFunc) (*SnapshotPolicy, int, error) {
	P := &ListSnapshotPoliciesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSnapshotPolicies(P)
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
		return l.SnapshotPolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SnapshotPolicy UUID: %s!", id)
}

// Lists snapshot policies.
func (s *SnapshotService) ListSnapshotPolicies(p *ListSnapshotPoliciesParams) (*ListSnapshotPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listSnapshotPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSnapshotPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSnapshotPoliciesResponse struct {
	Count            int               `json:"count"`
	SnapshotPolicies []*SnapshotPolicy `json:"snapshotpolicy"`
}

type SnapshotPolicy struct {
	Fordisplay     bool   `json:"fordisplay"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	Intervaltype   int    `json:"intervaltype"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Maxsnaps       int    `json:"maxsnaps"`
	Schedule       string `json:"schedule"`
	Tags           []Tags `json:"tags"`
	Timezone       string `json:"timezone"`
	Volumeid       string `json:"volumeid"`
}

type ListSnapshotsParams struct {
	P map[string]interface{}
}

func (P *ListSnapshotsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := P.P["intervaltype"]; found {
		u.Set("intervaltype", v.(string))
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
	if v, found := P.P["snapshottype"]; found {
		u.Set("snapshottype", v.(string))
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListSnapshotsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListSnapshotsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListSnapshotsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListSnapshotsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListSnapshotsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListSnapshotsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListSnapshotsParams) SetIds(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ids"] = v
}

func (P *ListSnapshotsParams) GetIds() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ids"].([]string)
	return value, ok
}

func (P *ListSnapshotsParams) SetIntervaltype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["intervaltype"] = v
}

func (P *ListSnapshotsParams) GetIntervaltype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["intervaltype"].(string)
	return value, ok
}

func (P *ListSnapshotsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListSnapshotsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListSnapshotsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListSnapshotsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListSnapshotsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListSnapshotsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListSnapshotsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListSnapshotsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListSnapshotsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListSnapshotsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListSnapshotsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListSnapshotsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListSnapshotsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListSnapshotsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListSnapshotsParams) SetSnapshottype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["snapshottype"] = v
}

func (P *ListSnapshotsParams) GetSnapshottype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["snapshottype"].(string)
	return value, ok
}

func (P *ListSnapshotsParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListSnapshotsParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListSnapshotsParams) SetVolumeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["volumeid"] = v
}

func (P *ListSnapshotsParams) GetVolumeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["volumeid"].(string)
	return value, ok
}

func (P *ListSnapshotsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListSnapshotsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSnapshotsParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewListSnapshotsParams() *ListSnapshotsParams {
	P := &ListSnapshotsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListSnapshotsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSnapshots(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Snapshots[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Snapshots {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotByName(name string, opts ...OptionFunc) (*Snapshot, int, error) {
	id, count, err := s.GetSnapshotID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSnapshotByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetSnapshotByID(id string, opts ...OptionFunc) (*Snapshot, int, error) {
	P := &ListSnapshotsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSnapshots(P)
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
		return l.Snapshots[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Snapshot UUID: %s!", id)
}

// Lists all available snapshots for the account.
func (s *SnapshotService) ListSnapshots(p *ListSnapshotsParams) (*ListSnapshotsResponse, error) {
	resp, err := s.cs.newRequest("listSnapshots", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSnapshotsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSnapshotsResponse struct {
	Count     int         `json:"count"`
	Snapshots []*Snapshot `json:"snapshot"`
}

type Snapshot struct {
	Account        string `json:"account"`
	Created        string `json:"created"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	Intervaltype   string `json:"intervaltype"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Locationtype   string `json:"locationtype"`
	Name           string `json:"name"`
	Osdisplayname  string `json:"osdisplayname"`
	Ostypeid       string `json:"ostypeid"`
	Physicalsize   int64  `json:"physicalsize"`
	Project        string `json:"project"`
	Projectid      string `json:"projectid"`
	Revertable     bool   `json:"revertable"`
	Snapshottype   string `json:"snapshottype"`
	State          string `json:"state"`
	Tags           []Tags `json:"tags"`
	Virtualsize    int64  `json:"virtualsize"`
	Volumeid       string `json:"volumeid"`
	Volumename     string `json:"volumename"`
	Volumetype     string `json:"volumetype"`
	Zoneid         string `json:"zoneid"`
}

func (r *Snapshot) UnmarshalJSON(b []byte) error {
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

	type alias Snapshot
	return json.Unmarshal(b, (*alias)(r))
}

type ListVMSnapshotParams struct {
	P map[string]interface{}
}

func (P *ListVMSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
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
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := P.P["vmsnapshotid"]; found {
		u.Set("vmsnapshotid", v.(string))
	}
	if v, found := P.P["vmsnapshotids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("vmsnapshotids", vv)
	}
	return u
}

func (P *ListVMSnapshotParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListVMSnapshotParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListVMSnapshotParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListVMSnapshotParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListVMSnapshotParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListVMSnapshotParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListVMSnapshotParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListVMSnapshotParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListVMSnapshotParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListVMSnapshotParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListVMSnapshotParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListVMSnapshotParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListVMSnapshotParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListVMSnapshotParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListVMSnapshotParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListVMSnapshotParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListVMSnapshotParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListVMSnapshotParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListVMSnapshotParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListVMSnapshotParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListVMSnapshotParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListVMSnapshotParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListVMSnapshotParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *ListVMSnapshotParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *ListVMSnapshotParams) SetVmsnapshotid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmsnapshotid"] = v
}

func (P *ListVMSnapshotParams) GetVmsnapshotid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmsnapshotid"].(string)
	return value, ok
}

func (P *ListVMSnapshotParams) SetVmsnapshotids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmsnapshotids"] = v
}

func (P *ListVMSnapshotParams) GetVmsnapshotids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmsnapshotids"].([]string)
	return value, ok
}

// You should always use this function to get a new ListVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewListVMSnapshotParams() *ListVMSnapshotParams {
	P := &ListVMSnapshotParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SnapshotService) GetVMSnapshotID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListVMSnapshotParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVMSnapshot(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VMSnapshot[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VMSnapshot {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// List virtual machine snapshot by conditions
func (s *SnapshotService) ListVMSnapshot(p *ListVMSnapshotParams) (*ListVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("listVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVMSnapshotResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVMSnapshotResponse struct {
	Count      int           `json:"count"`
	VMSnapshot []*VMSnapshot `json:"vmsnapshot"`
}

type VMSnapshot struct {
	Account            string `json:"account"`
	Created            string `json:"created"`
	Current            bool   `json:"current"`
	Description        string `json:"description"`
	Displayname        string `json:"displayname"`
	Domain             string `json:"domain"`
	Domainid           string `json:"domainid"`
	Hasannotations     bool   `json:"hasannotations"`
	Hypervisor         string `json:"hypervisor"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Name               string `json:"name"`
	Parent             string `json:"parent"`
	ParentName         string `json:"parentName"`
	Project            string `json:"project"`
	Projectid          string `json:"projectid"`
	State              string `json:"state"`
	Tags               []Tags `json:"tags"`
	Type               string `json:"type"`
	Virtualmachineid   string `json:"virtualmachineid"`
	Virtualmachinename string `json:"virtualmachinename"`
	Zoneid             string `json:"zoneid"`
	Zonename           string `json:"zonename"`
}

type RevertSnapshotParams struct {
	P map[string]interface{}
}

func (P *RevertSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *RevertSnapshotParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RevertSnapshotParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RevertSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewRevertSnapshotParams(id string) *RevertSnapshotParams {
	P := &RevertSnapshotParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// This is supposed to revert a volume snapshot. This command is only supported with KVM so far
func (s *SnapshotService) RevertSnapshot(p *RevertSnapshotParams) (*RevertSnapshotResponse, error) {
	resp, err := s.cs.newRequest("revertSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevertSnapshotResponse
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

type RevertSnapshotResponse struct {
	Account        string `json:"account"`
	Created        string `json:"created"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	Intervaltype   string `json:"intervaltype"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Locationtype   string `json:"locationtype"`
	Name           string `json:"name"`
	Osdisplayname  string `json:"osdisplayname"`
	Ostypeid       string `json:"ostypeid"`
	Physicalsize   int64  `json:"physicalsize"`
	Project        string `json:"project"`
	Projectid      string `json:"projectid"`
	Revertable     bool   `json:"revertable"`
	Snapshottype   string `json:"snapshottype"`
	State          string `json:"state"`
	Tags           []Tags `json:"tags"`
	Virtualsize    int64  `json:"virtualsize"`
	Volumeid       string `json:"volumeid"`
	Volumename     string `json:"volumename"`
	Volumetype     string `json:"volumetype"`
	Zoneid         string `json:"zoneid"`
}

func (r *RevertSnapshotResponse) UnmarshalJSON(b []byte) error {
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

	type alias RevertSnapshotResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RevertToVMSnapshotParams struct {
	P map[string]interface{}
}

func (P *RevertToVMSnapshotParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["vmsnapshotid"]; found {
		u.Set("vmsnapshotid", v.(string))
	}
	return u
}

func (P *RevertToVMSnapshotParams) SetVmsnapshotid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmsnapshotid"] = v
}

func (P *RevertToVMSnapshotParams) GetVmsnapshotid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmsnapshotid"].(string)
	return value, ok
}

// You should always use this function to get a new RevertToVMSnapshotParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewRevertToVMSnapshotParams(vmsnapshotid string) *RevertToVMSnapshotParams {
	P := &RevertToVMSnapshotParams{}
	P.P = make(map[string]interface{})
	P.P["vmsnapshotid"] = vmsnapshotid
	return P
}

// Revert VM from a vmsnapshot.
func (s *SnapshotService) RevertToVMSnapshot(p *RevertToVMSnapshotParams) (*RevertToVMSnapshotResponse, error) {
	resp, err := s.cs.newRequest("revertToVMSnapshot", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevertToVMSnapshotResponse
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

type RevertToVMSnapshotResponse struct {
	Account               string                                    `json:"account"`
	Affinitygroup         []RevertToVMSnapshotResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                    `json:"backupofferingid"`
	Backupofferingname    string                                    `json:"backupofferingname"`
	Bootmode              string                                    `json:"bootmode"`
	Boottype              string                                    `json:"boottype"`
	Cpunumber             int                                       `json:"cpunumber"`
	Cpuspeed              int                                       `json:"cpuspeed"`
	Cpuused               string                                    `json:"cpuused"`
	Created               string                                    `json:"created"`
	Details               map[string]string                         `json:"details"`
	Diskioread            int64                                     `json:"diskioread"`
	Diskiowrite           int64                                     `json:"diskiowrite"`
	Diskkbsread           int64                                     `json:"diskkbsread"`
	Diskkbswrite          int64                                     `json:"diskkbswrite"`
	Diskofferingid        string                                    `json:"diskofferingid"`
	Diskofferingname      string                                    `json:"diskofferingname"`
	Displayname           string                                    `json:"displayname"`
	Displayvm             bool                                      `json:"displayvm"`
	Domain                string                                    `json:"domain"`
	Domainid              string                                    `json:"domainid"`
	Forvirtualnetwork     bool                                      `json:"forvirtualnetwork"`
	Group                 string                                    `json:"group"`
	Groupid               string                                    `json:"groupid"`
	Guestosid             string                                    `json:"guestosid"`
	Haenable              bool                                      `json:"haenable"`
	Hasannotations        bool                                      `json:"hasannotations"`
	Hostid                string                                    `json:"hostid"`
	Hostname              string                                    `json:"hostname"`
	Hypervisor            string                                    `json:"hypervisor"`
	Icon                  string                                    `json:"icon"`
	Id                    string                                    `json:"id"`
	Instancename          string                                    `json:"instancename"`
	Isdynamicallyscalable bool                                      `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                    `json:"isodisplaytext"`
	Isoid                 string                                    `json:"isoid"`
	Isoname               string                                    `json:"isoname"`
	JobID                 string                                    `json:"jobid"`
	Jobstatus             int                                       `json:"jobstatus"`
	Keypair               string                                    `json:"keypair"`
	Lastupdated           string                                    `json:"lastupdated"`
	Memory                int                                       `json:"memory"`
	Memoryintfreekbs      int64                                     `json:"memoryintfreekbs"`
	Memorykbs             int64                                     `json:"memorykbs"`
	Memorytargetkbs       int64                                     `json:"memorytargetkbs"`
	Name                  string                                    `json:"name"`
	Networkkbsread        int64                                     `json:"networkkbsread"`
	Networkkbswrite       int64                                     `json:"networkkbswrite"`
	Nic                   []Nic                                     `json:"nic"`
	Osdisplayname         string                                    `json:"osdisplayname"`
	Ostypeid              string                                    `json:"ostypeid"`
	Password              string                                    `json:"password"`
	Passwordenabled       bool                                      `json:"passwordenabled"`
	Pooltype              string                                    `json:"pooltype"`
	Project               string                                    `json:"project"`
	Projectid             string                                    `json:"projectid"`
	Publicip              string                                    `json:"publicip"`
	Publicipid            string                                    `json:"publicipid"`
	Readonlydetails       string                                    `json:"readonlydetails"`
	Receivedbytes         int64                                     `json:"receivedbytes"`
	Rootdeviceid          int64                                     `json:"rootdeviceid"`
	Rootdevicetype        string                                    `json:"rootdevicetype"`
	Securitygroup         []RevertToVMSnapshotResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                     `json:"sentbytes"`
	Serviceofferingid     string                                    `json:"serviceofferingid"`
	Serviceofferingname   string                                    `json:"serviceofferingname"`
	Servicestate          string                                    `json:"servicestate"`
	State                 string                                    `json:"state"`
	Tags                  []Tags                                    `json:"tags"`
	Templatedisplaytext   string                                    `json:"templatedisplaytext"`
	Templateid            string                                    `json:"templateid"`
	Templatename          string                                    `json:"templatename"`
	Userid                string                                    `json:"userid"`
	Username              string                                    `json:"username"`
	Vgpu                  string                                    `json:"vgpu"`
	Zoneid                string                                    `json:"zoneid"`
	Zonename              string                                    `json:"zonename"`
}

type RevertToVMSnapshotResponseSecuritygroup struct {
	Account             string                                        `json:"account"`
	Description         string                                        `json:"description"`
	Domain              string                                        `json:"domain"`
	Domainid            string                                        `json:"domainid"`
	Egressrule          []RevertToVMSnapshotResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                        `json:"id"`
	Ingressrule         []RevertToVMSnapshotResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                        `json:"name"`
	Project             string                                        `json:"project"`
	Projectid           string                                        `json:"projectid"`
	Tags                []Tags                                        `json:"tags"`
	Virtualmachinecount int                                           `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                 `json:"virtualmachineids"`
}

type RevertToVMSnapshotResponseSecuritygroupRule struct {
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

type RevertToVMSnapshotResponseAffinitygroup struct {
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

func (r *RevertToVMSnapshotResponse) UnmarshalJSON(b []byte) error {
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

	type alias RevertToVMSnapshotResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateSnapshotPolicyParams struct {
	P map[string]interface{}
}

func (P *UpdateSnapshotPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *UpdateSnapshotPolicyParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateSnapshotPolicyParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateSnapshotPolicyParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateSnapshotPolicyParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateSnapshotPolicyParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateSnapshotPolicyParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateSnapshotPolicyParams instance,
// as then you are sure you have configured all required params
func (s *SnapshotService) NewUpdateSnapshotPolicyParams() *UpdateSnapshotPolicyParams {
	P := &UpdateSnapshotPolicyParams{}
	P.P = make(map[string]interface{})
	return P
}

// Updates the snapshot policy.
func (s *SnapshotService) UpdateSnapshotPolicy(p *UpdateSnapshotPolicyParams) (*UpdateSnapshotPolicyResponse, error) {
	resp, err := s.cs.newRequest("updateSnapshotPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateSnapshotPolicyResponse
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

type UpdateSnapshotPolicyResponse struct {
	Fordisplay     bool   `json:"fordisplay"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	Intervaltype   int    `json:"intervaltype"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Maxsnaps       int    `json:"maxsnaps"`
	Schedule       string `json:"schedule"`
	Tags           []Tags `json:"tags"`
	Timezone       string `json:"timezone"`
	Volumeid       string `json:"volumeid"`
}
