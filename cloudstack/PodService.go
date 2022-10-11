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

type PodServiceIface interface {
	CreatePod(p *CreatePodParams) (*CreatePodResponse, error)
	NewCreatePodParams(gateway string, name string, netmask string, startip string, zoneid string) *CreatePodParams
	DedicatePod(p *DedicatePodParams) (*DedicatePodResponse, error)
	NewDedicatePodParams(domainid string, podid string) *DedicatePodParams
	DeletePod(p *DeletePodParams) (*DeletePodResponse, error)
	NewDeletePodParams(id string) *DeletePodParams
	ListDedicatedPods(p *ListDedicatedPodsParams) (*ListDedicatedPodsResponse, error)
	NewListDedicatedPodsParams() *ListDedicatedPodsParams
	ListPods(p *ListPodsParams) (*ListPodsResponse, error)
	NewListPodsParams() *ListPodsParams
	GetPodID(name string, opts ...OptionFunc) (string, int, error)
	GetPodByName(name string, opts ...OptionFunc) (*Pod, int, error)
	GetPodByID(id string, opts ...OptionFunc) (*Pod, int, error)
	ReleaseDedicatedPod(p *ReleaseDedicatedPodParams) (*ReleaseDedicatedPodResponse, error)
	NewReleaseDedicatedPodParams(podid string) *ReleaseDedicatedPodParams
	UpdatePod(p *UpdatePodParams) (*UpdatePodResponse, error)
	NewUpdatePodParams(id string) *UpdatePodParams
}

type CreatePodParams struct {
	P map[string]interface{}
}

func (P *CreatePodParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := P.P["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := P.P["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *CreatePodParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *CreatePodParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *CreatePodParams) SetEndip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endip"] = v
}

func (P *CreatePodParams) GetEndip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endip"].(string)
	return value, ok
}

func (P *CreatePodParams) SetGateway(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gateway"] = v
}

func (P *CreatePodParams) GetGateway() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gateway"].(string)
	return value, ok
}

func (P *CreatePodParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreatePodParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreatePodParams) SetNetmask(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["netmask"] = v
}

func (P *CreatePodParams) GetNetmask() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["netmask"].(string)
	return value, ok
}

func (P *CreatePodParams) SetStartip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startip"] = v
}

func (P *CreatePodParams) GetStartip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startip"].(string)
	return value, ok
}

func (P *CreatePodParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreatePodParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreatePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewCreatePodParams(gateway string, name string, netmask string, startip string, zoneid string) *CreatePodParams {
	P := &CreatePodParams{}
	P.P = make(map[string]interface{})
	P.P["gateway"] = gateway
	P.P["name"] = name
	P.P["netmask"] = netmask
	P.P["startip"] = startip
	P.P["zoneid"] = zoneid
	return P
}

// Creates a new Pod.
func (s *PodService) CreatePod(p *CreatePodParams) (*CreatePodResponse, error) {
	resp, err := s.cs.newRequest("createPod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePodResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreatePodResponse struct {
	Allocationstate string                      `json:"allocationstate"`
	Capacity        []CreatePodResponseCapacity `json:"capacity"`
	Endip           []string                    `json:"endip"`
	Forsystemvms    []string                    `json:"forsystemvms"`
	Gateway         string                      `json:"gateway"`
	Hasannotations  bool                        `json:"hasannotations"`
	Id              string                      `json:"id"`
	Ipranges        []CreatePodResponseIpranges `json:"ipranges"`
	JobID           string                      `json:"jobid"`
	Jobstatus       int                         `json:"jobstatus"`
	Name            string                      `json:"name"`
	Netmask         string                      `json:"netmask"`
	Startip         []string                    `json:"startip"`
	Vlanid          []string                    `json:"vlanid"`
	Zoneid          string                      `json:"zoneid"`
	Zonename        string                      `json:"zonename"`
}

type CreatePodResponseIpranges struct {
	Endip        string `json:"endip"`
	Forsystemvms string `json:"forsystemvms"`
	Startip      string `json:"startip"`
	Vlanid       string `json:"vlanid"`
}

type CreatePodResponseCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}

type DedicatePodParams struct {
	P map[string]interface{}
}

func (P *DedicatePodParams) toURLValues() url.Values {
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
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (P *DedicatePodParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DedicatePodParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DedicatePodParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DedicatePodParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *DedicatePodParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *DedicatePodParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicatePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewDedicatePodParams(domainid string, podid string) *DedicatePodParams {
	P := &DedicatePodParams{}
	P.P = make(map[string]interface{})
	P.P["domainid"] = domainid
	P.P["podid"] = podid
	return P
}

// Dedicates a Pod.
func (s *PodService) DedicatePod(p *DedicatePodParams) (*DedicatePodResponse, error) {
	resp, err := s.cs.newRequest("dedicatePod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicatePodResponse
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

type DedicatePodResponse struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Podid           string `json:"podid"`
	Podname         string `json:"podname"`
}

type DeletePodParams struct {
	P map[string]interface{}
}

func (P *DeletePodParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeletePodParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeletePodParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeletePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewDeletePodParams(id string) *DeletePodParams {
	P := &DeletePodParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a Pod.
func (s *PodService) DeletePod(p *DeletePodParams) (*DeletePodResponse, error) {
	resp, err := s.cs.newRequest("deletePod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePodResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeletePodResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeletePodResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeletePodResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListDedicatedPodsParams struct {
	P map[string]interface{}
}

func (P *ListDedicatedPodsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
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
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (P *ListDedicatedPodsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListDedicatedPodsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListDedicatedPodsParams) SetAffinitygroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["affinitygroupid"] = v
}

func (P *ListDedicatedPodsParams) GetAffinitygroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["affinitygroupid"].(string)
	return value, ok
}

func (P *ListDedicatedPodsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListDedicatedPodsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListDedicatedPodsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListDedicatedPodsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListDedicatedPodsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListDedicatedPodsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListDedicatedPodsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListDedicatedPodsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListDedicatedPodsParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListDedicatedPodsParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

// You should always use this function to get a new ListDedicatedPodsParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewListDedicatedPodsParams() *ListDedicatedPodsParams {
	P := &ListDedicatedPodsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists dedicated pods.
func (s *PodService) ListDedicatedPods(p *ListDedicatedPodsParams) (*ListDedicatedPodsResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedPods", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedPodsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedPodsResponse struct {
	Count         int             `json:"count"`
	DedicatedPods []*DedicatedPod `json:"dedicatedpod"`
}

type DedicatedPod struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Podid           string `json:"podid"`
	Podname         string `json:"podname"`
}

type ListPodsParams struct {
	P map[string]interface{}
}

func (P *ListPodsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
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
	if v, found := P.P["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListPodsParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *ListPodsParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *ListPodsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListPodsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListPodsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListPodsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListPodsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListPodsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListPodsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListPodsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListPodsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListPodsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListPodsParams) SetShowcapacities(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showcapacities"] = v
}

func (P *ListPodsParams) GetShowcapacities() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showcapacities"].(bool)
	return value, ok
}

func (P *ListPodsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListPodsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListPodsParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewListPodsParams() *ListPodsParams {
	P := &ListPodsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *PodService) GetPodID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListPodsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListPods(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Pods[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Pods {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *PodService) GetPodByName(name string, opts ...OptionFunc) (*Pod, int, error) {
	id, count, err := s.GetPodID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetPodByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *PodService) GetPodByID(id string, opts ...OptionFunc) (*Pod, int, error) {
	P := &ListPodsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListPods(P)
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
		return l.Pods[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Pod UUID: %s!", id)
}

// Lists all Pods.
func (s *PodService) ListPods(p *ListPodsParams) (*ListPodsResponse, error) {
	resp, err := s.cs.newRequest("listPods", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPodsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPodsResponse struct {
	Count int    `json:"count"`
	Pods  []*Pod `json:"pod"`
}

type Pod struct {
	Allocationstate string        `json:"allocationstate"`
	Capacity        []PodCapacity `json:"capacity"`
	Endip           []string      `json:"endip"`
	Forsystemvms    []string      `json:"forsystemvms"`
	Gateway         string        `json:"gateway"`
	Hasannotations  bool          `json:"hasannotations"`
	Id              string        `json:"id"`
	Ipranges        []PodIpranges `json:"ipranges"`
	JobID           string        `json:"jobid"`
	Jobstatus       int           `json:"jobstatus"`
	Name            string        `json:"name"`
	Netmask         string        `json:"netmask"`
	Startip         []string      `json:"startip"`
	Vlanid          []string      `json:"vlanid"`
	Zoneid          string        `json:"zoneid"`
	Zonename        string        `json:"zonename"`
}

type PodIpranges struct {
	Endip        string `json:"endip"`
	Forsystemvms string `json:"forsystemvms"`
	Startip      string `json:"startip"`
	Vlanid       string `json:"vlanid"`
}

type PodCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}

type ReleaseDedicatedPodParams struct {
	P map[string]interface{}
}

func (P *ReleaseDedicatedPodParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (P *ReleaseDedicatedPodParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ReleaseDedicatedPodParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseDedicatedPodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewReleaseDedicatedPodParams(podid string) *ReleaseDedicatedPodParams {
	P := &ReleaseDedicatedPodParams{}
	P.P = make(map[string]interface{})
	P.P["podid"] = podid
	return P
}

// Release the dedication for the pod
func (s *PodService) ReleaseDedicatedPod(p *ReleaseDedicatedPodParams) (*ReleaseDedicatedPodResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedPod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedPodResponse
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

type ReleaseDedicatedPodResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdatePodParams struct {
	P map[string]interface{}
}

func (P *UpdatePodParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := P.P["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := P.P["startip"]; found {
		u.Set("startip", v.(string))
	}
	return u
}

func (P *UpdatePodParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *UpdatePodParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *UpdatePodParams) SetEndip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endip"] = v
}

func (P *UpdatePodParams) GetEndip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endip"].(string)
	return value, ok
}

func (P *UpdatePodParams) SetGateway(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gateway"] = v
}

func (P *UpdatePodParams) GetGateway() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gateway"].(string)
	return value, ok
}

func (P *UpdatePodParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdatePodParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdatePodParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdatePodParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdatePodParams) SetNetmask(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["netmask"] = v
}

func (P *UpdatePodParams) GetNetmask() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["netmask"].(string)
	return value, ok
}

func (P *UpdatePodParams) SetStartip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startip"] = v
}

func (P *UpdatePodParams) GetStartip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startip"].(string)
	return value, ok
}

// You should always use this function to get a new UpdatePodParams instance,
// as then you are sure you have configured all required params
func (s *PodService) NewUpdatePodParams(id string) *UpdatePodParams {
	P := &UpdatePodParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a Pod.
func (s *PodService) UpdatePod(p *UpdatePodParams) (*UpdatePodResponse, error) {
	resp, err := s.cs.newRequest("updatePod", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdatePodResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdatePodResponse struct {
	Allocationstate string                      `json:"allocationstate"`
	Capacity        []UpdatePodResponseCapacity `json:"capacity"`
	Endip           []string                    `json:"endip"`
	Forsystemvms    []string                    `json:"forsystemvms"`
	Gateway         string                      `json:"gateway"`
	Hasannotations  bool                        `json:"hasannotations"`
	Id              string                      `json:"id"`
	Ipranges        []UpdatePodResponseIpranges `json:"ipranges"`
	JobID           string                      `json:"jobid"`
	Jobstatus       int                         `json:"jobstatus"`
	Name            string                      `json:"name"`
	Netmask         string                      `json:"netmask"`
	Startip         []string                    `json:"startip"`
	Vlanid          []string                    `json:"vlanid"`
	Zoneid          string                      `json:"zoneid"`
	Zonename        string                      `json:"zonename"`
}

type UpdatePodResponseIpranges struct {
	Endip        string `json:"endip"`
	Forsystemvms string `json:"forsystemvms"`
	Startip      string `json:"startip"`
	Vlanid       string `json:"vlanid"`
}

type UpdatePodResponseCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}
