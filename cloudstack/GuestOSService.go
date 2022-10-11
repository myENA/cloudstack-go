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

type GuestOSServiceIface interface {
	AddGuestOs(p *AddGuestOsParams) (*AddGuestOsResponse, error)
	NewAddGuestOsParams(details map[string]string, oscategoryid string, osdisplayname string) *AddGuestOsParams
	AddGuestOsMapping(p *AddGuestOsMappingParams) (*AddGuestOsMappingResponse, error)
	NewAddGuestOsMappingParams(hypervisor string, hypervisorversion string, osnameforhypervisor string) *AddGuestOsMappingParams
	ListGuestOsMapping(p *ListGuestOsMappingParams) (*ListGuestOsMappingResponse, error)
	NewListGuestOsMappingParams() *ListGuestOsMappingParams
	GetGuestOsMappingByID(id string, opts ...OptionFunc) (*GuestOsMapping, int, error)
	ListOsCategories(p *ListOsCategoriesParams) (*ListOsCategoriesResponse, error)
	NewListOsCategoriesParams() *ListOsCategoriesParams
	GetOsCategoryID(name string, opts ...OptionFunc) (string, int, error)
	GetOsCategoryByName(name string, opts ...OptionFunc) (*OsCategory, int, error)
	GetOsCategoryByID(id string, opts ...OptionFunc) (*OsCategory, int, error)
	ListOsTypes(p *ListOsTypesParams) (*ListOsTypesResponse, error)
	NewListOsTypesParams() *ListOsTypesParams
	GetOsTypeByID(id string, opts ...OptionFunc) (*OsType, int, error)
	RemoveGuestOs(p *RemoveGuestOsParams) (*RemoveGuestOsResponse, error)
	NewRemoveGuestOsParams(id string) *RemoveGuestOsParams
	RemoveGuestOsMapping(p *RemoveGuestOsMappingParams) (*RemoveGuestOsMappingResponse, error)
	NewRemoveGuestOsMappingParams(id string) *RemoveGuestOsMappingParams
	UpdateGuestOs(p *UpdateGuestOsParams) (*UpdateGuestOsResponse, error)
	NewUpdateGuestOsParams(details map[string]string, id string, osdisplayname string) *UpdateGuestOsParams
	UpdateGuestOsMapping(p *UpdateGuestOsMappingParams) (*UpdateGuestOsMappingResponse, error)
	NewUpdateGuestOsMappingParams(id string, osnameforhypervisor string) *UpdateGuestOsMappingParams
}

type AddGuestOsParams struct {
	P map[string]interface{}
}

func (P *AddGuestOsParams) toURLValues() url.Values {
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
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["oscategoryid"]; found {
		u.Set("oscategoryid", v.(string))
	}
	if v, found := P.P["osdisplayname"]; found {
		u.Set("osdisplayname", v.(string))
	}
	return u
}

func (P *AddGuestOsParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *AddGuestOsParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *AddGuestOsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *AddGuestOsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *AddGuestOsParams) SetOscategoryid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["oscategoryid"] = v
}

func (P *AddGuestOsParams) GetOscategoryid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["oscategoryid"].(string)
	return value, ok
}

func (P *AddGuestOsParams) SetOsdisplayname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["osdisplayname"] = v
}

func (P *AddGuestOsParams) GetOsdisplayname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["osdisplayname"].(string)
	return value, ok
}

// You should always use this function to get a new AddGuestOsParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewAddGuestOsParams(details map[string]string, oscategoryid string, osdisplayname string) *AddGuestOsParams {
	P := &AddGuestOsParams{}
	P.P = make(map[string]interface{})
	P.P["details"] = details
	P.P["oscategoryid"] = oscategoryid
	P.P["osdisplayname"] = osdisplayname
	return P
}

// Add a new guest OS type
func (s *GuestOSService) AddGuestOs(p *AddGuestOsParams) (*AddGuestOsResponse, error) {
	resp, err := s.cs.newRequest("addGuestOs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddGuestOsResponse
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

type AddGuestOsResponse struct {
	Description   string `json:"description"`
	Id            string `json:"id"`
	Isuserdefined bool   `json:"isuserdefined"`
	JobID         string `json:"jobid"`
	Jobstatus     int    `json:"jobstatus"`
	Oscategoryid  string `json:"oscategoryid"`
}

type AddGuestOsMappingParams struct {
	P map[string]interface{}
}

func (P *AddGuestOsMappingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["hypervisorversion"]; found {
		u.Set("hypervisorversion", v.(string))
	}
	if v, found := P.P["osdisplayname"]; found {
		u.Set("osdisplayname", v.(string))
	}
	if v, found := P.P["osnameforhypervisor"]; found {
		u.Set("osnameforhypervisor", v.(string))
	}
	if v, found := P.P["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	return u
}

func (P *AddGuestOsMappingParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *AddGuestOsMappingParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *AddGuestOsMappingParams) SetHypervisorversion(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisorversion"] = v
}

func (P *AddGuestOsMappingParams) GetHypervisorversion() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisorversion"].(string)
	return value, ok
}

func (P *AddGuestOsMappingParams) SetOsdisplayname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["osdisplayname"] = v
}

func (P *AddGuestOsMappingParams) GetOsdisplayname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["osdisplayname"].(string)
	return value, ok
}

func (P *AddGuestOsMappingParams) SetOsnameforhypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["osnameforhypervisor"] = v
}

func (P *AddGuestOsMappingParams) GetOsnameforhypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["osnameforhypervisor"].(string)
	return value, ok
}

func (P *AddGuestOsMappingParams) SetOstypeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ostypeid"] = v
}

func (P *AddGuestOsMappingParams) GetOstypeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ostypeid"].(string)
	return value, ok
}

// You should always use this function to get a new AddGuestOsMappingParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewAddGuestOsMappingParams(hypervisor string, hypervisorversion string, osnameforhypervisor string) *AddGuestOsMappingParams {
	P := &AddGuestOsMappingParams{}
	P.P = make(map[string]interface{})
	P.P["hypervisor"] = hypervisor
	P.P["hypervisorversion"] = hypervisorversion
	P.P["osnameforhypervisor"] = osnameforhypervisor
	return P
}

// Adds a guest OS name to hypervisor OS name mapping
func (s *GuestOSService) AddGuestOsMapping(p *AddGuestOsMappingParams) (*AddGuestOsMappingResponse, error) {
	resp, err := s.cs.newRequest("addGuestOsMapping", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddGuestOsMappingResponse
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

type AddGuestOsMappingResponse struct {
	Hypervisor          string `json:"hypervisor"`
	Hypervisorversion   string `json:"hypervisorversion"`
	Id                  string `json:"id"`
	Isuserdefined       string `json:"isuserdefined"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Osdisplayname       string `json:"osdisplayname"`
	Osnameforhypervisor string `json:"osnameforhypervisor"`
	Ostypeid            string `json:"ostypeid"`
}

func (r *AddGuestOsMappingResponse) UnmarshalJSON(b []byte) error {
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

	type alias AddGuestOsMappingResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListGuestOsMappingParams struct {
	P map[string]interface{}
}

func (P *ListGuestOsMappingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["hypervisorversion"]; found {
		u.Set("hypervisorversion", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (P *ListGuestOsMappingParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *ListGuestOsMappingParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *ListGuestOsMappingParams) SetHypervisorversion(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisorversion"] = v
}

func (P *ListGuestOsMappingParams) GetHypervisorversion() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisorversion"].(string)
	return value, ok
}

func (P *ListGuestOsMappingParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListGuestOsMappingParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListGuestOsMappingParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListGuestOsMappingParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListGuestOsMappingParams) SetOstypeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ostypeid"] = v
}

func (P *ListGuestOsMappingParams) GetOstypeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ostypeid"].(string)
	return value, ok
}

func (P *ListGuestOsMappingParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListGuestOsMappingParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListGuestOsMappingParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListGuestOsMappingParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListGuestOsMappingParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewListGuestOsMappingParams() *ListGuestOsMappingParams {
	P := &ListGuestOsMappingParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetGuestOsMappingByID(id string, opts ...OptionFunc) (*GuestOsMapping, int, error) {
	P := &ListGuestOsMappingParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListGuestOsMapping(P)
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
		return l.GuestOsMapping[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for GuestOsMapping UUID: %s!", id)
}

// Lists all available OS mappings for given hypervisor
func (s *GuestOSService) ListGuestOsMapping(p *ListGuestOsMappingParams) (*ListGuestOsMappingResponse, error) {
	resp, err := s.cs.newRequest("listGuestOsMapping", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListGuestOsMappingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListGuestOsMappingResponse struct {
	Count          int               `json:"count"`
	GuestOsMapping []*GuestOsMapping `json:"guestosmapping"`
}

type GuestOsMapping struct {
	Hypervisor          string `json:"hypervisor"`
	Hypervisorversion   string `json:"hypervisorversion"`
	Id                  string `json:"id"`
	Isuserdefined       string `json:"isuserdefined"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Osdisplayname       string `json:"osdisplayname"`
	Osnameforhypervisor string `json:"osnameforhypervisor"`
	Ostypeid            string `json:"ostypeid"`
}

func (r *GuestOsMapping) UnmarshalJSON(b []byte) error {
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

	type alias GuestOsMapping
	return json.Unmarshal(b, (*alias)(r))
}

type ListOsCategoriesParams struct {
	P map[string]interface{}
}

func (P *ListOsCategoriesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
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
	return u
}

func (P *ListOsCategoriesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListOsCategoriesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListOsCategoriesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListOsCategoriesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListOsCategoriesParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListOsCategoriesParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListOsCategoriesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListOsCategoriesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListOsCategoriesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListOsCategoriesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListOsCategoriesParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewListOsCategoriesParams() *ListOsCategoriesParams {
	P := &ListOsCategoriesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsCategoryID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListOsCategoriesParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListOsCategories(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.OsCategories[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.OsCategories {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsCategoryByName(name string, opts ...OptionFunc) (*OsCategory, int, error) {
	id, count, err := s.GetOsCategoryID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetOsCategoryByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsCategoryByID(id string, opts ...OptionFunc) (*OsCategory, int, error) {
	P := &ListOsCategoriesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListOsCategories(P)
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
		return l.OsCategories[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for OsCategory UUID: %s!", id)
}

// Lists all supported OS categories for this cloud.
func (s *GuestOSService) ListOsCategories(p *ListOsCategoriesParams) (*ListOsCategoriesResponse, error) {
	resp, err := s.cs.newRequest("listOsCategories", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListOsCategoriesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListOsCategoriesResponse struct {
	Count        int           `json:"count"`
	OsCategories []*OsCategory `json:"oscategory"`
}

type OsCategory struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type ListOsTypesParams struct {
	P map[string]interface{}
}

func (P *ListOsTypesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["oscategoryid"]; found {
		u.Set("oscategoryid", v.(string))
	}
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (P *ListOsTypesParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *ListOsTypesParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *ListOsTypesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListOsTypesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListOsTypesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListOsTypesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListOsTypesParams) SetOscategoryid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["oscategoryid"] = v
}

func (P *ListOsTypesParams) GetOscategoryid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["oscategoryid"].(string)
	return value, ok
}

func (P *ListOsTypesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListOsTypesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListOsTypesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListOsTypesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListOsTypesParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewListOsTypesParams() *ListOsTypesParams {
	P := &ListOsTypesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *GuestOSService) GetOsTypeByID(id string, opts ...OptionFunc) (*OsType, int, error) {
	P := &ListOsTypesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListOsTypes(P)
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
		return l.OsTypes[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for OsType UUID: %s!", id)
}

// Lists all supported OS types for this cloud.
func (s *GuestOSService) ListOsTypes(p *ListOsTypesParams) (*ListOsTypesResponse, error) {
	resp, err := s.cs.newRequest("listOsTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListOsTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListOsTypesResponse struct {
	Count   int       `json:"count"`
	OsTypes []*OsType `json:"ostype"`
}

type OsType struct {
	Description   string `json:"description"`
	Id            string `json:"id"`
	Isuserdefined bool   `json:"isuserdefined"`
	JobID         string `json:"jobid"`
	Jobstatus     int    `json:"jobstatus"`
	Oscategoryid  string `json:"oscategoryid"`
}

type RemoveGuestOsParams struct {
	P map[string]interface{}
}

func (P *RemoveGuestOsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *RemoveGuestOsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RemoveGuestOsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveGuestOsParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewRemoveGuestOsParams(id string) *RemoveGuestOsParams {
	P := &RemoveGuestOsParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Removes a Guest OS from listing.
func (s *GuestOSService) RemoveGuestOs(p *RemoveGuestOsParams) (*RemoveGuestOsResponse, error) {
	resp, err := s.cs.newRequest("removeGuestOs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveGuestOsResponse
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

type RemoveGuestOsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RemoveGuestOsMappingParams struct {
	P map[string]interface{}
}

func (P *RemoveGuestOsMappingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *RemoveGuestOsMappingParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RemoveGuestOsMappingParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveGuestOsMappingParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewRemoveGuestOsMappingParams(id string) *RemoveGuestOsMappingParams {
	P := &RemoveGuestOsMappingParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Removes a Guest OS Mapping.
func (s *GuestOSService) RemoveGuestOsMapping(p *RemoveGuestOsMappingParams) (*RemoveGuestOsMappingResponse, error) {
	resp, err := s.cs.newRequest("removeGuestOsMapping", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveGuestOsMappingResponse
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

type RemoveGuestOsMappingResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateGuestOsParams struct {
	P map[string]interface{}
}

func (P *UpdateGuestOsParams) toURLValues() url.Values {
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
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["osdisplayname"]; found {
		u.Set("osdisplayname", v.(string))
	}
	return u
}

func (P *UpdateGuestOsParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *UpdateGuestOsParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *UpdateGuestOsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateGuestOsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateGuestOsParams) SetOsdisplayname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["osdisplayname"] = v
}

func (P *UpdateGuestOsParams) GetOsdisplayname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["osdisplayname"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateGuestOsParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewUpdateGuestOsParams(details map[string]string, id string, osdisplayname string) *UpdateGuestOsParams {
	P := &UpdateGuestOsParams{}
	P.P = make(map[string]interface{})
	P.P["details"] = details
	P.P["id"] = id
	P.P["osdisplayname"] = osdisplayname
	return P
}

// Updates the information about Guest OS
func (s *GuestOSService) UpdateGuestOs(p *UpdateGuestOsParams) (*UpdateGuestOsResponse, error) {
	resp, err := s.cs.newRequest("updateGuestOs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateGuestOsResponse
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

type UpdateGuestOsResponse struct {
	Description   string `json:"description"`
	Id            string `json:"id"`
	Isuserdefined bool   `json:"isuserdefined"`
	JobID         string `json:"jobid"`
	Jobstatus     int    `json:"jobstatus"`
	Oscategoryid  string `json:"oscategoryid"`
}

type UpdateGuestOsMappingParams struct {
	P map[string]interface{}
}

func (P *UpdateGuestOsMappingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["osnameforhypervisor"]; found {
		u.Set("osnameforhypervisor", v.(string))
	}
	return u
}

func (P *UpdateGuestOsMappingParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateGuestOsMappingParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateGuestOsMappingParams) SetOsnameforhypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["osnameforhypervisor"] = v
}

func (P *UpdateGuestOsMappingParams) GetOsnameforhypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["osnameforhypervisor"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateGuestOsMappingParams instance,
// as then you are sure you have configured all required params
func (s *GuestOSService) NewUpdateGuestOsMappingParams(id string, osnameforhypervisor string) *UpdateGuestOsMappingParams {
	P := &UpdateGuestOsMappingParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["osnameforhypervisor"] = osnameforhypervisor
	return P
}

// Updates the information about Guest OS to Hypervisor specific name mapping
func (s *GuestOSService) UpdateGuestOsMapping(p *UpdateGuestOsMappingParams) (*UpdateGuestOsMappingResponse, error) {
	resp, err := s.cs.newRequest("updateGuestOsMapping", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateGuestOsMappingResponse
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

type UpdateGuestOsMappingResponse struct {
	Hypervisor          string `json:"hypervisor"`
	Hypervisorversion   string `json:"hypervisorversion"`
	Id                  string `json:"id"`
	Isuserdefined       string `json:"isuserdefined"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Osdisplayname       string `json:"osdisplayname"`
	Osnameforhypervisor string `json:"osnameforhypervisor"`
	Ostypeid            string `json:"ostypeid"`
}

func (r *UpdateGuestOsMappingResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateGuestOsMappingResponse
	return json.Unmarshal(b, (*alias)(r))
}
