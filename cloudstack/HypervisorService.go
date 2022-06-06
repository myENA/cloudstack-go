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

type HypervisorServiceIface interface {
	ListHypervisorCapabilities(P *ListHypervisorCapabilitiesParams) (*ListHypervisorCapabilitiesResponse, error)
	NewListHypervisorCapabilitiesParams() *ListHypervisorCapabilitiesParams
	GetHypervisorCapabilityByID(id string, opts ...OptionFunc) (*HypervisorCapability, int, error)
	ListHypervisors(P *ListHypervisorsParams) (*ListHypervisorsResponse, error)
	NewListHypervisorsParams() *ListHypervisorsParams
	UpdateHypervisorCapabilities(P *UpdateHypervisorCapabilitiesParams) (*UpdateHypervisorCapabilitiesResponse, error)
	NewUpdateHypervisorCapabilitiesParams() *UpdateHypervisorCapabilitiesParams
}

type ListHypervisorCapabilitiesParams struct {
	P map[string]interface{}
}

func (P *ListHypervisorCapabilitiesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
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
	return u
}

func (P *ListHypervisorCapabilitiesParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *ListHypervisorCapabilitiesParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *ListHypervisorCapabilitiesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListHypervisorCapabilitiesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListHypervisorCapabilitiesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListHypervisorCapabilitiesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListHypervisorCapabilitiesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListHypervisorCapabilitiesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListHypervisorCapabilitiesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListHypervisorCapabilitiesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListHypervisorCapabilitiesParams instance,
// as then you are sure you have configured all required params
func (s *HypervisorService) NewListHypervisorCapabilitiesParams() *ListHypervisorCapabilitiesParams {
	P := &ListHypervisorCapabilitiesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HypervisorService) GetHypervisorCapabilityByID(id string, opts ...OptionFunc) (*HypervisorCapability, int, error) {
	P := &ListHypervisorCapabilitiesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListHypervisorCapabilities(P)
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
		return l.HypervisorCapabilities[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for HypervisorCapability UUID: %s!", id)
}

// Lists all hypervisor capabilities.
func (s *HypervisorService) ListHypervisorCapabilities(p *ListHypervisorCapabilitiesParams) (*ListHypervisorCapabilitiesResponse, error) {
	resp, err := s.cs.newRequest("listHypervisorCapabilities", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHypervisorCapabilitiesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListHypervisorCapabilitiesResponse struct {
	Count                  int                     `json:"count"`
	HypervisorCapabilities []*HypervisorCapability `json:"hypervisorcapability"`
}

type HypervisorCapability struct {
	Hypervisor           string `json:"hypervisor"`
	Hypervisorversion    string `json:"hypervisorversion"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Maxdatavolumeslimit  int    `json:"maxdatavolumeslimit"`
	Maxguestslimit       int64  `json:"maxguestslimit"`
	Maxhostspercluster   int    `json:"maxhostspercluster"`
	Securitygroupenabled bool   `json:"securitygroupenabled"`
	Storagemotionenabled bool   `json:"storagemotionenabled"`
}

type ListHypervisorsParams struct {
	P map[string]interface{}
}

func (P *ListHypervisorsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListHypervisorsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListHypervisorsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListHypervisorsParams instance,
// as then you are sure you have configured all required params
func (s *HypervisorService) NewListHypervisorsParams() *ListHypervisorsParams {
	P := &ListHypervisorsParams{}
	P.P = make(map[string]interface{})
	return P
}

// List hypervisors
func (s *HypervisorService) ListHypervisors(p *ListHypervisorsParams) (*ListHypervisorsResponse, error) {
	resp, err := s.cs.newRequest("listHypervisors", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHypervisorsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListHypervisorsResponse struct {
	Count       int           `json:"count"`
	Hypervisors []*Hypervisor `json:"hypervisor"`
}

type Hypervisor struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type UpdateHypervisorCapabilitiesParams struct {
	P map[string]interface{}
}

func (P *UpdateHypervisorCapabilitiesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["maxguestslimit"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxguestslimit", vv)
	}
	if v, found := P.P["securitygroupenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("securitygroupenabled", vv)
	}
	return u
}

func (P *UpdateHypervisorCapabilitiesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateHypervisorCapabilitiesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateHypervisorCapabilitiesParams) SetMaxguestslimit(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxguestslimit"] = v
}

func (P *UpdateHypervisorCapabilitiesParams) GetMaxguestslimit() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxguestslimit"].(int64)
	return value, ok
}

func (P *UpdateHypervisorCapabilitiesParams) SetSecuritygroupenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupenabled"] = v
}

func (P *UpdateHypervisorCapabilitiesParams) GetSecuritygroupenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupenabled"].(bool)
	return value, ok
}

// You should always use this function to get a new UpdateHypervisorCapabilitiesParams instance,
// as then you are sure you have configured all required params
func (s *HypervisorService) NewUpdateHypervisorCapabilitiesParams() *UpdateHypervisorCapabilitiesParams {
	P := &UpdateHypervisorCapabilitiesParams{}
	P.P = make(map[string]interface{})
	return P
}

// Updates a hypervisor capabilities.
func (s *HypervisorService) UpdateHypervisorCapabilities(p *UpdateHypervisorCapabilitiesParams) (*UpdateHypervisorCapabilitiesResponse, error) {
	resp, err := s.cs.newRequest("updateHypervisorCapabilities", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateHypervisorCapabilitiesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateHypervisorCapabilitiesResponse struct {
	Hypervisor           string `json:"hypervisor"`
	Hypervisorversion    string `json:"hypervisorversion"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Maxdatavolumeslimit  int    `json:"maxdatavolumeslimit"`
	Maxguestslimit       int64  `json:"maxguestslimit"`
	Maxhostspercluster   int    `json:"maxhostspercluster"`
	Securitygroupenabled bool   `json:"securitygroupenabled"`
	Storagemotionenabled bool   `json:"storagemotionenabled"`
}
