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
	"net/url"
	"strconv"
)

type LimitServiceIface interface {
	GetApiLimit(p *GetApiLimitParams) (*GetApiLimitResponse, error)
	NewGetApiLimitParams() *GetApiLimitParams
	ListResourceLimits(p *ListResourceLimitsParams) (*ListResourceLimitsResponse, error)
	NewListResourceLimitsParams() *ListResourceLimitsParams
	ResetApiLimit(p *ResetApiLimitParams) (*ResetApiLimitResponse, error)
	NewResetApiLimitParams() *ResetApiLimitParams
	UpdateResourceCount(p *UpdateResourceCountParams) (*UpdateResourceCountResponse, error)
	NewUpdateResourceCountParams(domainid string) *UpdateResourceCountParams
	UpdateResourceLimit(p *UpdateResourceLimitParams) (*UpdateResourceLimitResponse, error)
	NewUpdateResourceLimitParams(resourcetype int) *UpdateResourceLimitParams
}

type GetApiLimitParams struct {
	P map[string]interface{}
}

func (P *GetApiLimitParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	return u
}

// You should always use this function to get a new GetApiLimitParams instance,
// as then you are sure you have configured all required params
func (s *LimitService) NewGetApiLimitParams() *GetApiLimitParams {
	P := &GetApiLimitParams{}
	P.P = make(map[string]interface{})
	return P
}

// Get API limit count for the caller
func (s *LimitService) GetApiLimit(p *GetApiLimitParams) (*GetApiLimitResponse, error) {
	resp, err := s.cs.newRequest("getApiLimit", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetApiLimitResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetApiLimitResponse struct {
	Account     string `json:"account"`
	Accountid   string `json:"accountid"`
	ApiAllowed  int    `json:"apiAllowed"`
	ApiIssued   int    `json:"apiIssued"`
	ExpireAfter int64  `json:"expireAfter"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
}

type ListResourceLimitsParams struct {
	P map[string]interface{}
}

func (P *ListResourceLimitsParams) toURLValues() url.Values {
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
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("id", vv)
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
	if v, found := P.P["resourcetype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("resourcetype", vv)
	}
	if v, found := P.P["resourcetypename"]; found {
		u.Set("resourcetypename", v.(string))
	}
	return u
}

func (P *ListResourceLimitsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListResourceLimitsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListResourceLimitsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListResourceLimitsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListResourceLimitsParams) SetId(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListResourceLimitsParams) GetId() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(int64)
	return value, ok
}

func (P *ListResourceLimitsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListResourceLimitsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListResourceLimitsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListResourceLimitsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListResourceLimitsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListResourceLimitsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListResourceLimitsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListResourceLimitsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListResourceLimitsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListResourceLimitsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListResourceLimitsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListResourceLimitsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListResourceLimitsParams) SetResourcetype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcetype"] = v
}

func (P *ListResourceLimitsParams) GetResourcetype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcetype"].(int)
	return value, ok
}

func (P *ListResourceLimitsParams) SetResourcetypename(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcetypename"] = v
}

func (P *ListResourceLimitsParams) GetResourcetypename() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcetypename"].(string)
	return value, ok
}

// You should always use this function to get a new ListResourceLimitsParams instance,
// as then you are sure you have configured all required params
func (s *LimitService) NewListResourceLimitsParams() *ListResourceLimitsParams {
	P := &ListResourceLimitsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists resource limits.
func (s *LimitService) ListResourceLimits(p *ListResourceLimitsParams) (*ListResourceLimitsResponse, error) {
	resp, err := s.cs.newRequest("listResourceLimits", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListResourceLimitsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListResourceLimitsResponse struct {
	Count          int              `json:"count"`
	ResourceLimits []*ResourceLimit `json:"resourcelimit"`
}

type ResourceLimit struct {
	Account          string `json:"account"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Max              int64  `json:"max"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Resourcetype     string `json:"resourcetype"`
	Resourcetypename string `json:"resourcetypename"`
}

type ResetApiLimitParams struct {
	P map[string]interface{}
}

func (P *ResetApiLimitParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	return u
}

func (P *ResetApiLimitParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ResetApiLimitParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

// You should always use this function to get a new ResetApiLimitParams instance,
// as then you are sure you have configured all required params
func (s *LimitService) NewResetApiLimitParams() *ResetApiLimitParams {
	P := &ResetApiLimitParams{}
	P.P = make(map[string]interface{})
	return P
}

// Reset api count
func (s *LimitService) ResetApiLimit(p *ResetApiLimitParams) (*ResetApiLimitResponse, error) {
	resp, err := s.cs.newRequest("resetApiLimit", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetApiLimitResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ResetApiLimitResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ResetApiLimitResponse) UnmarshalJSON(b []byte) error {
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

	type alias ResetApiLimitResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateResourceCountParams struct {
	P map[string]interface{}
}

func (P *UpdateResourceCountParams) toURLValues() url.Values {
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["resourcetype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("resourcetype", vv)
	}
	return u
}

func (P *UpdateResourceCountParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *UpdateResourceCountParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *UpdateResourceCountParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UpdateResourceCountParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UpdateResourceCountParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *UpdateResourceCountParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *UpdateResourceCountParams) SetResourcetype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcetype"] = v
}

func (P *UpdateResourceCountParams) GetResourcetype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcetype"].(int)
	return value, ok
}

// You should always use this function to get a new UpdateResourceCountParams instance,
// as then you are sure you have configured all required params
func (s *LimitService) NewUpdateResourceCountParams(domainid string) *UpdateResourceCountParams {
	P := &UpdateResourceCountParams{}
	P.P = make(map[string]interface{})
	P.P["domainid"] = domainid
	return P
}

// Recalculate and update resource count for an account or domain.
func (s *LimitService) UpdateResourceCount(p *UpdateResourceCountParams) (*UpdateResourceCountResponse, error) {
	resp, err := s.cs.newRequest("updateResourceCount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateResourceCountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateResourceCountResponse struct {
	Account          string `json:"account"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Resourcecount    int64  `json:"resourcecount"`
	Resourcetype     string `json:"resourcetype"`
	Resourcetypename string `json:"resourcetypename"`
}

type UpdateResourceLimitParams struct {
	P map[string]interface{}
}

func (P *UpdateResourceLimitParams) toURLValues() url.Values {
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
	if v, found := P.P["max"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("max", vv)
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["resourcetype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("resourcetype", vv)
	}
	return u
}

func (P *UpdateResourceLimitParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *UpdateResourceLimitParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *UpdateResourceLimitParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UpdateResourceLimitParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UpdateResourceLimitParams) SetMax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["max"] = v
}

func (P *UpdateResourceLimitParams) GetMax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["max"].(int64)
	return value, ok
}

func (P *UpdateResourceLimitParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *UpdateResourceLimitParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *UpdateResourceLimitParams) SetResourcetype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcetype"] = v
}

func (P *UpdateResourceLimitParams) GetResourcetype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcetype"].(int)
	return value, ok
}

// You should always use this function to get a new UpdateResourceLimitParams instance,
// as then you are sure you have configured all required params
func (s *LimitService) NewUpdateResourceLimitParams(resourcetype int) *UpdateResourceLimitParams {
	P := &UpdateResourceLimitParams{}
	P.P = make(map[string]interface{})
	P.P["resourcetype"] = resourcetype
	return P
}

// Updates resource limits for an account or domain.
func (s *LimitService) UpdateResourceLimit(p *UpdateResourceLimitParams) (*UpdateResourceLimitResponse, error) {
	resp, err := s.cs.newRequest("updateResourceLimit", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateResourceLimitResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateResourceLimitResponse struct {
	Account          string `json:"account"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Max              int64  `json:"max"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Resourcetype     string `json:"resourcetype"`
	Resourcetypename string `json:"resourcetypename"`
}
