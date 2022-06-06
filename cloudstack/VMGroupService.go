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

type VMGroupServiceIface interface {
	CreateInstanceGroup(P *CreateInstanceGroupParams) (*CreateInstanceGroupResponse, error)
	NewCreateInstanceGroupParams(name string) *CreateInstanceGroupParams
	DeleteInstanceGroup(P *DeleteInstanceGroupParams) (*DeleteInstanceGroupResponse, error)
	NewDeleteInstanceGroupParams(id string) *DeleteInstanceGroupParams
	ListInstanceGroups(P *ListInstanceGroupsParams) (*ListInstanceGroupsResponse, error)
	NewListInstanceGroupsParams() *ListInstanceGroupsParams
	GetInstanceGroupID(name string, opts ...OptionFunc) (string, int, error)
	GetInstanceGroupByName(name string, opts ...OptionFunc) (*InstanceGroup, int, error)
	GetInstanceGroupByID(id string, opts ...OptionFunc) (*InstanceGroup, int, error)
	UpdateInstanceGroup(P *UpdateInstanceGroupParams) (*UpdateInstanceGroupResponse, error)
	NewUpdateInstanceGroupParams(id string) *UpdateInstanceGroupParams
}

type CreateInstanceGroupParams struct {
	P map[string]interface{}
}

func (P *CreateInstanceGroupParams) toURLValues() url.Values {
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
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *CreateInstanceGroupParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateInstanceGroupParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateInstanceGroupParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateInstanceGroupParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateInstanceGroupParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateInstanceGroupParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateInstanceGroupParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateInstanceGroupParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateInstanceGroupParams instance,
// as then you are sure you have configured all required params
func (s *VMGroupService) NewCreateInstanceGroupParams(name string) *CreateInstanceGroupParams {
	P := &CreateInstanceGroupParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	return P
}

// Creates a vm group
func (s *VMGroupService) CreateInstanceGroup(p *CreateInstanceGroupParams) (*CreateInstanceGroupResponse, error) {
	resp, err := s.cs.newRequest("createInstanceGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateInstanceGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateInstanceGroupResponse struct {
	Account        string `json:"account"`
	Created        string `json:"created"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Project        string `json:"project"`
	Projectid      string `json:"projectid"`
}

type DeleteInstanceGroupParams struct {
	P map[string]interface{}
}

func (P *DeleteInstanceGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteInstanceGroupParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteInstanceGroupParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteInstanceGroupParams instance,
// as then you are sure you have configured all required params
func (s *VMGroupService) NewDeleteInstanceGroupParams(id string) *DeleteInstanceGroupParams {
	P := &DeleteInstanceGroupParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a vm group
func (s *VMGroupService) DeleteInstanceGroup(p *DeleteInstanceGroupParams) (*DeleteInstanceGroupResponse, error) {
	resp, err := s.cs.newRequest("deleteInstanceGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteInstanceGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteInstanceGroupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteInstanceGroupResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteInstanceGroupResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListInstanceGroupsParams struct {
	P map[string]interface{}
}

func (P *ListInstanceGroupsParams) toURLValues() url.Values {
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
	return u
}

func (P *ListInstanceGroupsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListInstanceGroupsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListInstanceGroupsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListInstanceGroupsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListInstanceGroupsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListInstanceGroupsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListInstanceGroupsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListInstanceGroupsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListInstanceGroupsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListInstanceGroupsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListInstanceGroupsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListInstanceGroupsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListInstanceGroupsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListInstanceGroupsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListInstanceGroupsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListInstanceGroupsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListInstanceGroupsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListInstanceGroupsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListInstanceGroupsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListInstanceGroupsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ListInstanceGroupsParams instance,
// as then you are sure you have configured all required params
func (s *VMGroupService) NewListInstanceGroupsParams() *ListInstanceGroupsParams {
	P := &ListInstanceGroupsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VMGroupService) GetInstanceGroupID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListInstanceGroupsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListInstanceGroups(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.InstanceGroups[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.InstanceGroups {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VMGroupService) GetInstanceGroupByName(name string, opts ...OptionFunc) (*InstanceGroup, int, error) {
	id, count, err := s.GetInstanceGroupID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetInstanceGroupByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VMGroupService) GetInstanceGroupByID(id string, opts ...OptionFunc) (*InstanceGroup, int, error) {
	P := &ListInstanceGroupsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListInstanceGroups(P)
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
		return l.InstanceGroups[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for InstanceGroup UUID: %s!", id)
}

// Lists vm groups
func (s *VMGroupService) ListInstanceGroups(p *ListInstanceGroupsParams) (*ListInstanceGroupsResponse, error) {
	resp, err := s.cs.newRequest("listInstanceGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListInstanceGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListInstanceGroupsResponse struct {
	Count          int              `json:"count"`
	InstanceGroups []*InstanceGroup `json:"instancegroup"`
}

type InstanceGroup struct {
	Account        string `json:"account"`
	Created        string `json:"created"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Project        string `json:"project"`
	Projectid      string `json:"projectid"`
}

type UpdateInstanceGroupParams struct {
	P map[string]interface{}
}

func (P *UpdateInstanceGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (P *UpdateInstanceGroupParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateInstanceGroupParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateInstanceGroupParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateInstanceGroupParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateInstanceGroupParams instance,
// as then you are sure you have configured all required params
func (s *VMGroupService) NewUpdateInstanceGroupParams(id string) *UpdateInstanceGroupParams {
	P := &UpdateInstanceGroupParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a vm group
func (s *VMGroupService) UpdateInstanceGroup(p *UpdateInstanceGroupParams) (*UpdateInstanceGroupResponse, error) {
	resp, err := s.cs.newRequest("updateInstanceGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateInstanceGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateInstanceGroupResponse struct {
	Account        string `json:"account"`
	Created        string `json:"created"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Project        string `json:"project"`
	Projectid      string `json:"projectid"`
}
