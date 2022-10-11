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
)

type ResourcemetadataServiceIface interface {
	AddResourceDetail(p *AddResourceDetailParams) (*AddResourceDetailResponse, error)
	NewAddResourceDetailParams(details map[string]string, resourceid string, resourcetype string) *AddResourceDetailParams
	GetVolumeSnapshotDetails(p *GetVolumeSnapshotDetailsParams) (*GetVolumeSnapshotDetailsResponse, error)
	NewGetVolumeSnapshotDetailsParams(snapshotid string) *GetVolumeSnapshotDetailsParams
	ListResourceDetails(p *ListResourceDetailsParams) (*ListResourceDetailsResponse, error)
	NewListResourceDetailsParams(resourcetype string) *ListResourceDetailsParams
	RemoveResourceDetail(p *RemoveResourceDetailParams) (*RemoveResourceDetailResponse, error)
	NewRemoveResourceDetailParams(resourceid string, resourcetype string) *RemoveResourceDetailParams
}

type AddResourceDetailParams struct {
	P map[string]interface{}
}

func (P *AddResourceDetailParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), m[k])
		}
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := P.P["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (P *AddResourceDetailParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *AddResourceDetailParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *AddResourceDetailParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *AddResourceDetailParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *AddResourceDetailParams) SetResourceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourceid"] = v
}

func (P *AddResourceDetailParams) GetResourceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourceid"].(string)
	return value, ok
}

func (P *AddResourceDetailParams) SetResourcetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcetype"] = v
}

func (P *AddResourceDetailParams) GetResourcetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new AddResourceDetailParams instance,
// as then you are sure you have configured all required params
func (s *ResourcemetadataService) NewAddResourceDetailParams(details map[string]string, resourceid string, resourcetype string) *AddResourceDetailParams {
	P := &AddResourceDetailParams{}
	P.P = make(map[string]interface{})
	P.P["details"] = details
	P.P["resourceid"] = resourceid
	P.P["resourcetype"] = resourcetype
	return P
}

// Adds detail for the Resource.
func (s *ResourcemetadataService) AddResourceDetail(p *AddResourceDetailParams) (*AddResourceDetailResponse, error) {
	resp, err := s.cs.newRequest("addResourceDetail", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddResourceDetailResponse
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

type AddResourceDetailResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type GetVolumeSnapshotDetailsParams struct {
	P map[string]interface{}
}

func (P *GetVolumeSnapshotDetailsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["snapshotid"]; found {
		u.Set("snapshotid", v.(string))
	}
	return u
}

func (P *GetVolumeSnapshotDetailsParams) SetSnapshotid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["snapshotid"] = v
}

func (P *GetVolumeSnapshotDetailsParams) GetSnapshotid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["snapshotid"].(string)
	return value, ok
}

// You should always use this function to get a new GetVolumeSnapshotDetailsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcemetadataService) NewGetVolumeSnapshotDetailsParams(snapshotid string) *GetVolumeSnapshotDetailsParams {
	P := &GetVolumeSnapshotDetailsParams{}
	P.P = make(map[string]interface{})
	P.P["snapshotid"] = snapshotid
	return P
}

// Get Volume Snapshot Details
func (s *ResourcemetadataService) GetVolumeSnapshotDetails(p *GetVolumeSnapshotDetailsParams) (*GetVolumeSnapshotDetailsResponse, error) {
	resp, err := s.cs.newRequest("getVolumeSnapshotDetails", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetVolumeSnapshotDetailsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetVolumeSnapshotDetailsResponse struct {
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	VolumeiScsiName string `json:"volumeiScsiName"`
}

type ListResourceDetailsParams struct {
	P map[string]interface{}
}

func (P *ListResourceDetailsParams) toURLValues() url.Values {
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
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := P.P["key"]; found {
		u.Set("key", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
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
	if v, found := P.P["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := P.P["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	if v, found := P.P["value"]; found {
		u.Set("value", v.(string))
	}
	return u
}

func (P *ListResourceDetailsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListResourceDetailsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListResourceDetailsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListResourceDetailsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListResourceDetailsParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListResourceDetailsParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListResourceDetailsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListResourceDetailsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListResourceDetailsParams) SetKey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["key"] = v
}

func (P *ListResourceDetailsParams) GetKey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["key"].(string)
	return value, ok
}

func (P *ListResourceDetailsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListResourceDetailsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListResourceDetailsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListResourceDetailsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListResourceDetailsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListResourceDetailsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListResourceDetailsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListResourceDetailsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListResourceDetailsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListResourceDetailsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListResourceDetailsParams) SetResourceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourceid"] = v
}

func (P *ListResourceDetailsParams) GetResourceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourceid"].(string)
	return value, ok
}

func (P *ListResourceDetailsParams) SetResourcetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcetype"] = v
}

func (P *ListResourceDetailsParams) GetResourcetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcetype"].(string)
	return value, ok
}

func (P *ListResourceDetailsParams) SetValue(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["value"] = v
}

func (P *ListResourceDetailsParams) GetValue() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["value"].(string)
	return value, ok
}

// You should always use this function to get a new ListResourceDetailsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcemetadataService) NewListResourceDetailsParams(resourcetype string) *ListResourceDetailsParams {
	P := &ListResourceDetailsParams{}
	P.P = make(map[string]interface{})
	P.P["resourcetype"] = resourcetype
	return P
}

// List resource detail(s)
func (s *ResourcemetadataService) ListResourceDetails(p *ListResourceDetailsParams) (*ListResourceDetailsResponse, error) {
	resp, err := s.cs.newRequest("listResourceDetails", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListResourceDetailsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListResourceDetailsResponse struct {
	Count           int               `json:"count"`
	ResourceDetails []*ResourceDetail `json:"resourcedetail"`
}

type ResourceDetail struct {
	Account      string `json:"account"`
	Customer     string `json:"customer"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Key          string `json:"key"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Resourceid   string `json:"resourceid"`
	Resourcetype string `json:"resourcetype"`
	Value        string `json:"value"`
}

type RemoveResourceDetailParams struct {
	P map[string]interface{}
}

func (P *RemoveResourceDetailParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["key"]; found {
		u.Set("key", v.(string))
	}
	if v, found := P.P["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := P.P["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	return u
}

func (P *RemoveResourceDetailParams) SetKey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["key"] = v
}

func (P *RemoveResourceDetailParams) GetKey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["key"].(string)
	return value, ok
}

func (P *RemoveResourceDetailParams) SetResourceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourceid"] = v
}

func (P *RemoveResourceDetailParams) GetResourceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourceid"].(string)
	return value, ok
}

func (P *RemoveResourceDetailParams) SetResourcetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcetype"] = v
}

func (P *RemoveResourceDetailParams) GetResourcetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcetype"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveResourceDetailParams instance,
// as then you are sure you have configured all required params
func (s *ResourcemetadataService) NewRemoveResourceDetailParams(resourceid string, resourcetype string) *RemoveResourceDetailParams {
	P := &RemoveResourceDetailParams{}
	P.P = make(map[string]interface{})
	P.P["resourceid"] = resourceid
	P.P["resourcetype"] = resourcetype
	return P
}

// Removes detail for the Resource.
func (s *ResourcemetadataService) RemoveResourceDetail(p *RemoveResourceDetailParams) (*RemoveResourceDetailResponse, error) {
	resp, err := s.cs.newRequest("removeResourceDetail", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveResourceDetailResponse
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

type RemoveResourceDetailResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
