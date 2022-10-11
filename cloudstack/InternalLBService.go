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

type InternalLBServiceIface interface {
	ConfigureInternalLoadBalancerElement(p *ConfigureInternalLoadBalancerElementParams) (*InternalLoadBalancerElementResponse, error)
	NewConfigureInternalLoadBalancerElementParams(enabled bool, id string) *ConfigureInternalLoadBalancerElementParams
	CreateInternalLoadBalancerElement(p *CreateInternalLoadBalancerElementParams) (*CreateInternalLoadBalancerElementResponse, error)
	NewCreateInternalLoadBalancerElementParams(nspid string) *CreateInternalLoadBalancerElementParams
	ListInternalLoadBalancerElements(p *ListInternalLoadBalancerElementsParams) (*ListInternalLoadBalancerElementsResponse, error)
	NewListInternalLoadBalancerElementsParams() *ListInternalLoadBalancerElementsParams
	GetInternalLoadBalancerElementByID(id string, opts ...OptionFunc) (*InternalLoadBalancerElement, int, error)
	ListInternalLoadBalancerVMs(p *ListInternalLoadBalancerVMsParams) (*ListInternalLoadBalancerVMsResponse, error)
	NewListInternalLoadBalancerVMsParams() *ListInternalLoadBalancerVMsParams
	GetInternalLoadBalancerVMID(name string, opts ...OptionFunc) (string, int, error)
	GetInternalLoadBalancerVMByName(name string, opts ...OptionFunc) (*InternalLoadBalancerVM, int, error)
	GetInternalLoadBalancerVMByID(id string, opts ...OptionFunc) (*InternalLoadBalancerVM, int, error)
	StartInternalLoadBalancerVM(p *StartInternalLoadBalancerVMParams) (*StartInternalLoadBalancerVMResponse, error)
	NewStartInternalLoadBalancerVMParams(id string) *StartInternalLoadBalancerVMParams
	StopInternalLoadBalancerVM(p *StopInternalLoadBalancerVMParams) (*StopInternalLoadBalancerVMResponse, error)
	NewStopInternalLoadBalancerVMParams(id string) *StopInternalLoadBalancerVMParams
}

type ConfigureInternalLoadBalancerElementParams struct {
	P map[string]interface{}
}

func (P *ConfigureInternalLoadBalancerElementParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ConfigureInternalLoadBalancerElementParams) SetEnabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enabled"] = v
}

func (P *ConfigureInternalLoadBalancerElementParams) GetEnabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enabled"].(bool)
	return value, ok
}

func (P *ConfigureInternalLoadBalancerElementParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ConfigureInternalLoadBalancerElementParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ConfigureInternalLoadBalancerElementParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewConfigureInternalLoadBalancerElementParams(enabled bool, id string) *ConfigureInternalLoadBalancerElementParams {
	P := &ConfigureInternalLoadBalancerElementParams{}
	P.P = make(map[string]interface{})
	P.P["enabled"] = enabled
	P.P["id"] = id
	return P
}

// Configures an Internal Load Balancer element.
func (s *InternalLBService) ConfigureInternalLoadBalancerElement(p *ConfigureInternalLoadBalancerElementParams) (*InternalLoadBalancerElementResponse, error) {
	resp, err := s.cs.newRequest("configureInternalLoadBalancerElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r InternalLoadBalancerElementResponse
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

type InternalLoadBalancerElementResponse struct {
	Enabled   bool   `json:"enabled"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Nspid     string `json:"nspid"`
}

type CreateInternalLoadBalancerElementParams struct {
	P map[string]interface{}
}

func (P *CreateInternalLoadBalancerElementParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["nspid"]; found {
		u.Set("nspid", v.(string))
	}
	return u
}

func (P *CreateInternalLoadBalancerElementParams) SetNspid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nspid"] = v
}

func (P *CreateInternalLoadBalancerElementParams) GetNspid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nspid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateInternalLoadBalancerElementParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewCreateInternalLoadBalancerElementParams(nspid string) *CreateInternalLoadBalancerElementParams {
	P := &CreateInternalLoadBalancerElementParams{}
	P.P = make(map[string]interface{})
	P.P["nspid"] = nspid
	return P
}

// Create an Internal Load Balancer element.
func (s *InternalLBService) CreateInternalLoadBalancerElement(p *CreateInternalLoadBalancerElementParams) (*CreateInternalLoadBalancerElementResponse, error) {
	resp, err := s.cs.newRequest("createInternalLoadBalancerElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateInternalLoadBalancerElementResponse
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

type CreateInternalLoadBalancerElementResponse struct {
	Enabled   bool   `json:"enabled"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Nspid     string `json:"nspid"`
}

type ListInternalLoadBalancerElementsParams struct {
	P map[string]interface{}
}

func (P *ListInternalLoadBalancerElementsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["nspid"]; found {
		u.Set("nspid", v.(string))
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

func (P *ListInternalLoadBalancerElementsParams) SetEnabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enabled"] = v
}

func (P *ListInternalLoadBalancerElementsParams) GetEnabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enabled"].(bool)
	return value, ok
}

func (P *ListInternalLoadBalancerElementsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListInternalLoadBalancerElementsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerElementsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListInternalLoadBalancerElementsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerElementsParams) SetNspid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nspid"] = v
}

func (P *ListInternalLoadBalancerElementsParams) GetNspid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nspid"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerElementsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListInternalLoadBalancerElementsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListInternalLoadBalancerElementsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListInternalLoadBalancerElementsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListInternalLoadBalancerElementsParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewListInternalLoadBalancerElementsParams() *ListInternalLoadBalancerElementsParams {
	P := &ListInternalLoadBalancerElementsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InternalLBService) GetInternalLoadBalancerElementByID(id string, opts ...OptionFunc) (*InternalLoadBalancerElement, int, error) {
	P := &ListInternalLoadBalancerElementsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListInternalLoadBalancerElements(P)
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
		return l.InternalLoadBalancerElements[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for InternalLoadBalancerElement UUID: %s!", id)
}

// Lists all available Internal Load Balancer elements.
func (s *InternalLBService) ListInternalLoadBalancerElements(p *ListInternalLoadBalancerElementsParams) (*ListInternalLoadBalancerElementsResponse, error) {
	resp, err := s.cs.newRequest("listInternalLoadBalancerElements", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListInternalLoadBalancerElementsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListInternalLoadBalancerElementsResponse struct {
	Count                        int                            `json:"count"`
	InternalLoadBalancerElements []*InternalLoadBalancerElement `json:"internalloadbalancerelement"`
}

type InternalLoadBalancerElement struct {
	Enabled   bool   `json:"enabled"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Nspid     string `json:"nspid"`
}

type ListInternalLoadBalancerVMsParams struct {
	P map[string]interface{}
}

func (P *ListInternalLoadBalancerVMsParams) toURLValues() url.Values {
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
	if v, found := P.P["fetchhealthcheckresults"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fetchhealthcheckresults", vv)
	}
	if v, found := P.P["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
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
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListInternalLoadBalancerVMsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetFetchhealthcheckresults(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fetchhealthcheckresults"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetFetchhealthcheckresults() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fetchhealthcheckresults"].(bool)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetForvpc(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forvpc"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetForvpc() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forvpc"].(bool)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

func (P *ListInternalLoadBalancerVMsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListInternalLoadBalancerVMsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListInternalLoadBalancerVMsParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewListInternalLoadBalancerVMsParams() *ListInternalLoadBalancerVMsParams {
	P := &ListInternalLoadBalancerVMsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InternalLBService) GetInternalLoadBalancerVMID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListInternalLoadBalancerVMsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListInternalLoadBalancerVMs(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.InternalLoadBalancerVMs[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.InternalLoadBalancerVMs {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InternalLBService) GetInternalLoadBalancerVMByName(name string, opts ...OptionFunc) (*InternalLoadBalancerVM, int, error) {
	id, count, err := s.GetInternalLoadBalancerVMID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetInternalLoadBalancerVMByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *InternalLBService) GetInternalLoadBalancerVMByID(id string, opts ...OptionFunc) (*InternalLoadBalancerVM, int, error) {
	P := &ListInternalLoadBalancerVMsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListInternalLoadBalancerVMs(P)
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
		return l.InternalLoadBalancerVMs[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for InternalLoadBalancerVM UUID: %s!", id)
}

// List internal LB VMs.
func (s *InternalLBService) ListInternalLoadBalancerVMs(p *ListInternalLoadBalancerVMsParams) (*ListInternalLoadBalancerVMsResponse, error) {
	resp, err := s.cs.newRequest("listInternalLoadBalancerVMs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListInternalLoadBalancerVMsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListInternalLoadBalancerVMsResponse struct {
	Count                   int                       `json:"count"`
	InternalLoadBalancerVMs []*InternalLoadBalancerVM `json:"internalloadbalancervm"`
}

type InternalLoadBalancerVM struct {
	Account             string                                     `json:"account"`
	Created             string                                     `json:"created"`
	Dns1                string                                     `json:"dns1"`
	Dns2                string                                     `json:"dns2"`
	Domain              string                                     `json:"domain"`
	Domainid            string                                     `json:"domainid"`
	Gateway             string                                     `json:"gateway"`
	Guestipaddress      string                                     `json:"guestipaddress"`
	Guestmacaddress     string                                     `json:"guestmacaddress"`
	Guestnetmask        string                                     `json:"guestnetmask"`
	Guestnetworkid      string                                     `json:"guestnetworkid"`
	Guestnetworkname    string                                     `json:"guestnetworkname"`
	Hasannotations      bool                                       `json:"hasannotations"`
	Healthcheckresults  []InternalLoadBalancerVMHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                       `json:"healthchecksfailed"`
	Hostid              string                                     `json:"hostid"`
	Hostname            string                                     `json:"hostname"`
	Hypervisor          string                                     `json:"hypervisor"`
	Id                  string                                     `json:"id"`
	Ip6dns1             string                                     `json:"ip6dns1"`
	Ip6dns2             string                                     `json:"ip6dns2"`
	Isredundantrouter   bool                                       `json:"isredundantrouter"`
	JobID               string                                     `json:"jobid"`
	Jobstatus           int                                        `json:"jobstatus"`
	Linklocalip         string                                     `json:"linklocalip"`
	Linklocalmacaddress string                                     `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                     `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                     `json:"linklocalnetworkid"`
	Name                string                                     `json:"name"`
	Networkdomain       string                                     `json:"networkdomain"`
	Nic                 []Nic                                      `json:"nic"`
	Podid               string                                     `json:"podid"`
	Podname             string                                     `json:"podname"`
	Project             string                                     `json:"project"`
	Projectid           string                                     `json:"projectid"`
	Publicip            string                                     `json:"publicip"`
	Publicmacaddress    string                                     `json:"publicmacaddress"`
	Publicnetmask       string                                     `json:"publicnetmask"`
	Publicnetworkid     string                                     `json:"publicnetworkid"`
	Redundantstate      string                                     `json:"redundantstate"`
	Requiresupgrade     bool                                       `json:"requiresupgrade"`
	Role                string                                     `json:"role"`
	Scriptsversion      string                                     `json:"scriptsversion"`
	Serviceofferingid   string                                     `json:"serviceofferingid"`
	Serviceofferingname string                                     `json:"serviceofferingname"`
	State               string                                     `json:"state"`
	Templateid          string                                     `json:"templateid"`
	Templatename        string                                     `json:"templatename"`
	Version             string                                     `json:"version"`
	Vpcid               string                                     `json:"vpcid"`
	Vpcname             string                                     `json:"vpcname"`
	Zoneid              string                                     `json:"zoneid"`
	Zonename            string                                     `json:"zonename"`
}

type InternalLoadBalancerVMHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type StartInternalLoadBalancerVMParams struct {
	P map[string]interface{}
}

func (P *StartInternalLoadBalancerVMParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *StartInternalLoadBalancerVMParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *StartInternalLoadBalancerVMParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new StartInternalLoadBalancerVMParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewStartInternalLoadBalancerVMParams(id string) *StartInternalLoadBalancerVMParams {
	P := &StartInternalLoadBalancerVMParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Starts an existing internal lb vm.
func (s *InternalLBService) StartInternalLoadBalancerVM(p *StartInternalLoadBalancerVMParams) (*StartInternalLoadBalancerVMResponse, error) {
	resp, err := s.cs.newRequest("startInternalLoadBalancerVM", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartInternalLoadBalancerVMResponse
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

type StartInternalLoadBalancerVMResponse struct {
	Account             string                                                  `json:"account"`
	Created             string                                                  `json:"created"`
	Dns1                string                                                  `json:"dns1"`
	Dns2                string                                                  `json:"dns2"`
	Domain              string                                                  `json:"domain"`
	Domainid            string                                                  `json:"domainid"`
	Gateway             string                                                  `json:"gateway"`
	Guestipaddress      string                                                  `json:"guestipaddress"`
	Guestmacaddress     string                                                  `json:"guestmacaddress"`
	Guestnetmask        string                                                  `json:"guestnetmask"`
	Guestnetworkid      string                                                  `json:"guestnetworkid"`
	Guestnetworkname    string                                                  `json:"guestnetworkname"`
	Hasannotations      bool                                                    `json:"hasannotations"`
	Healthcheckresults  []StartInternalLoadBalancerVMResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                                    `json:"healthchecksfailed"`
	Hostid              string                                                  `json:"hostid"`
	Hostname            string                                                  `json:"hostname"`
	Hypervisor          string                                                  `json:"hypervisor"`
	Id                  string                                                  `json:"id"`
	Ip6dns1             string                                                  `json:"ip6dns1"`
	Ip6dns2             string                                                  `json:"ip6dns2"`
	Isredundantrouter   bool                                                    `json:"isredundantrouter"`
	JobID               string                                                  `json:"jobid"`
	Jobstatus           int                                                     `json:"jobstatus"`
	Linklocalip         string                                                  `json:"linklocalip"`
	Linklocalmacaddress string                                                  `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                                  `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                                  `json:"linklocalnetworkid"`
	Name                string                                                  `json:"name"`
	Networkdomain       string                                                  `json:"networkdomain"`
	Nic                 []Nic                                                   `json:"nic"`
	Podid               string                                                  `json:"podid"`
	Podname             string                                                  `json:"podname"`
	Project             string                                                  `json:"project"`
	Projectid           string                                                  `json:"projectid"`
	Publicip            string                                                  `json:"publicip"`
	Publicmacaddress    string                                                  `json:"publicmacaddress"`
	Publicnetmask       string                                                  `json:"publicnetmask"`
	Publicnetworkid     string                                                  `json:"publicnetworkid"`
	Redundantstate      string                                                  `json:"redundantstate"`
	Requiresupgrade     bool                                                    `json:"requiresupgrade"`
	Role                string                                                  `json:"role"`
	Scriptsversion      string                                                  `json:"scriptsversion"`
	Serviceofferingid   string                                                  `json:"serviceofferingid"`
	Serviceofferingname string                                                  `json:"serviceofferingname"`
	State               string                                                  `json:"state"`
	Templateid          string                                                  `json:"templateid"`
	Templatename        string                                                  `json:"templatename"`
	Version             string                                                  `json:"version"`
	Vpcid               string                                                  `json:"vpcid"`
	Vpcname             string                                                  `json:"vpcname"`
	Zoneid              string                                                  `json:"zoneid"`
	Zonename            string                                                  `json:"zonename"`
}

type StartInternalLoadBalancerVMResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type StopInternalLoadBalancerVMParams struct {
	P map[string]interface{}
}

func (P *StopInternalLoadBalancerVMParams) toURLValues() url.Values {
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
	return u
}

func (P *StopInternalLoadBalancerVMParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *StopInternalLoadBalancerVMParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *StopInternalLoadBalancerVMParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *StopInternalLoadBalancerVMParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new StopInternalLoadBalancerVMParams instance,
// as then you are sure you have configured all required params
func (s *InternalLBService) NewStopInternalLoadBalancerVMParams(id string) *StopInternalLoadBalancerVMParams {
	P := &StopInternalLoadBalancerVMParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Stops an Internal LB vm.
func (s *InternalLBService) StopInternalLoadBalancerVM(p *StopInternalLoadBalancerVMParams) (*StopInternalLoadBalancerVMResponse, error) {
	resp, err := s.cs.newRequest("stopInternalLoadBalancerVM", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopInternalLoadBalancerVMResponse
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

type StopInternalLoadBalancerVMResponse struct {
	Account             string                                                 `json:"account"`
	Created             string                                                 `json:"created"`
	Dns1                string                                                 `json:"dns1"`
	Dns2                string                                                 `json:"dns2"`
	Domain              string                                                 `json:"domain"`
	Domainid            string                                                 `json:"domainid"`
	Gateway             string                                                 `json:"gateway"`
	Guestipaddress      string                                                 `json:"guestipaddress"`
	Guestmacaddress     string                                                 `json:"guestmacaddress"`
	Guestnetmask        string                                                 `json:"guestnetmask"`
	Guestnetworkid      string                                                 `json:"guestnetworkid"`
	Guestnetworkname    string                                                 `json:"guestnetworkname"`
	Hasannotations      bool                                                   `json:"hasannotations"`
	Healthcheckresults  []StopInternalLoadBalancerVMResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                                   `json:"healthchecksfailed"`
	Hostid              string                                                 `json:"hostid"`
	Hostname            string                                                 `json:"hostname"`
	Hypervisor          string                                                 `json:"hypervisor"`
	Id                  string                                                 `json:"id"`
	Ip6dns1             string                                                 `json:"ip6dns1"`
	Ip6dns2             string                                                 `json:"ip6dns2"`
	Isredundantrouter   bool                                                   `json:"isredundantrouter"`
	JobID               string                                                 `json:"jobid"`
	Jobstatus           int                                                    `json:"jobstatus"`
	Linklocalip         string                                                 `json:"linklocalip"`
	Linklocalmacaddress string                                                 `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                                 `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                                 `json:"linklocalnetworkid"`
	Name                string                                                 `json:"name"`
	Networkdomain       string                                                 `json:"networkdomain"`
	Nic                 []Nic                                                  `json:"nic"`
	Podid               string                                                 `json:"podid"`
	Podname             string                                                 `json:"podname"`
	Project             string                                                 `json:"project"`
	Projectid           string                                                 `json:"projectid"`
	Publicip            string                                                 `json:"publicip"`
	Publicmacaddress    string                                                 `json:"publicmacaddress"`
	Publicnetmask       string                                                 `json:"publicnetmask"`
	Publicnetworkid     string                                                 `json:"publicnetworkid"`
	Redundantstate      string                                                 `json:"redundantstate"`
	Requiresupgrade     bool                                                   `json:"requiresupgrade"`
	Role                string                                                 `json:"role"`
	Scriptsversion      string                                                 `json:"scriptsversion"`
	Serviceofferingid   string                                                 `json:"serviceofferingid"`
	Serviceofferingname string                                                 `json:"serviceofferingname"`
	State               string                                                 `json:"state"`
	Templateid          string                                                 `json:"templateid"`
	Templatename        string                                                 `json:"templatename"`
	Version             string                                                 `json:"version"`
	Vpcid               string                                                 `json:"vpcid"`
	Vpcname             string                                                 `json:"vpcname"`
	Zoneid              string                                                 `json:"zoneid"`
	Zonename            string                                                 `json:"zonename"`
}

type StopInternalLoadBalancerVMResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}
