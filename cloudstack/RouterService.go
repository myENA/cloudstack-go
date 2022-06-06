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

type RouterServiceIface interface {
	ChangeServiceForRouter(P *ChangeServiceForRouterParams) (*ChangeServiceForRouterResponse, error)
	NewChangeServiceForRouterParams(id string, serviceofferingid string) *ChangeServiceForRouterParams
	ConfigureVirtualRouterElement(P *ConfigureVirtualRouterElementParams) (*VirtualRouterElementResponse, error)
	NewConfigureVirtualRouterElementParams(enabled bool, id string) *ConfigureVirtualRouterElementParams
	CreateVirtualRouterElement(P *CreateVirtualRouterElementParams) (*CreateVirtualRouterElementResponse, error)
	NewCreateVirtualRouterElementParams(nspid string) *CreateVirtualRouterElementParams
	DestroyRouter(P *DestroyRouterParams) (*DestroyRouterResponse, error)
	NewDestroyRouterParams(id string) *DestroyRouterParams
	ListRouters(P *ListRoutersParams) (*ListRoutersResponse, error)
	NewListRoutersParams() *ListRoutersParams
	GetRouterID(name string, opts ...OptionFunc) (string, int, error)
	GetRouterByName(name string, opts ...OptionFunc) (*Router, int, error)
	GetRouterByID(id string, opts ...OptionFunc) (*Router, int, error)
	ListVirtualRouterElements(P *ListVirtualRouterElementsParams) (*ListVirtualRouterElementsResponse, error)
	NewListVirtualRouterElementsParams() *ListVirtualRouterElementsParams
	GetVirtualRouterElementByID(id string, opts ...OptionFunc) (*VirtualRouterElement, int, error)
	RebootRouter(P *RebootRouterParams) (*RebootRouterResponse, error)
	NewRebootRouterParams(id string) *RebootRouterParams
	StartRouter(P *StartRouterParams) (*StartRouterResponse, error)
	NewStartRouterParams(id string) *StartRouterParams
	StopRouter(P *StopRouterParams) (*StopRouterResponse, error)
	NewStopRouterParams(id string) *StopRouterParams
}

type ChangeServiceForRouterParams struct {
	P map[string]interface{}
}

func (P *ChangeServiceForRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (P *ChangeServiceForRouterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ChangeServiceForRouterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ChangeServiceForRouterParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *ChangeServiceForRouterParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

// You should always use this function to get a new ChangeServiceForRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewChangeServiceForRouterParams(id string, serviceofferingid string) *ChangeServiceForRouterParams {
	P := &ChangeServiceForRouterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["serviceofferingid"] = serviceofferingid
	return P
}

// Upgrades domain router to a new service offering
func (s *RouterService) ChangeServiceForRouter(p *ChangeServiceForRouterParams) (*ChangeServiceForRouterResponse, error) {
	resp, err := s.cs.newRequest("changeServiceForRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForRouterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ChangeServiceForRouterResponse struct {
	Account             string                                             `json:"account"`
	Created             string                                             `json:"created"`
	Dns1                string                                             `json:"dns1"`
	Dns2                string                                             `json:"dns2"`
	Domain              string                                             `json:"domain"`
	Domainid            string                                             `json:"domainid"`
	Gateway             string                                             `json:"gateway"`
	Guestipaddress      string                                             `json:"guestipaddress"`
	Guestmacaddress     string                                             `json:"guestmacaddress"`
	Guestnetmask        string                                             `json:"guestnetmask"`
	Guestnetworkid      string                                             `json:"guestnetworkid"`
	Guestnetworkname    string                                             `json:"guestnetworkname"`
	Hasannotations      bool                                               `json:"hasannotations"`
	Healthcheckresults  []ChangeServiceForRouterResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                               `json:"healthchecksfailed"`
	Hostid              string                                             `json:"hostid"`
	Hostname            string                                             `json:"hostname"`
	Hypervisor          string                                             `json:"hypervisor"`
	Id                  string                                             `json:"id"`
	Ip6dns1             string                                             `json:"ip6dns1"`
	Ip6dns2             string                                             `json:"ip6dns2"`
	Isredundantrouter   bool                                               `json:"isredundantrouter"`
	JobID               string                                             `json:"jobid"`
	Jobstatus           int                                                `json:"jobstatus"`
	Linklocalip         string                                             `json:"linklocalip"`
	Linklocalmacaddress string                                             `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                             `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                             `json:"linklocalnetworkid"`
	Name                string                                             `json:"name"`
	Networkdomain       string                                             `json:"networkdomain"`
	Nic                 []Nic                                              `json:"nic"`
	Podid               string                                             `json:"podid"`
	Podname             string                                             `json:"podname"`
	Project             string                                             `json:"project"`
	Projectid           string                                             `json:"projectid"`
	Publicip            string                                             `json:"publicip"`
	Publicmacaddress    string                                             `json:"publicmacaddress"`
	Publicnetmask       string                                             `json:"publicnetmask"`
	Publicnetworkid     string                                             `json:"publicnetworkid"`
	Redundantstate      string                                             `json:"redundantstate"`
	Requiresupgrade     bool                                               `json:"requiresupgrade"`
	Role                string                                             `json:"role"`
	Scriptsversion      string                                             `json:"scriptsversion"`
	Serviceofferingid   string                                             `json:"serviceofferingid"`
	Serviceofferingname string                                             `json:"serviceofferingname"`
	State               string                                             `json:"state"`
	Templateid          string                                             `json:"templateid"`
	Templatename        string                                             `json:"templatename"`
	Version             string                                             `json:"version"`
	Vpcid               string                                             `json:"vpcid"`
	Vpcname             string                                             `json:"vpcname"`
	Zoneid              string                                             `json:"zoneid"`
	Zonename            string                                             `json:"zonename"`
}

type ChangeServiceForRouterResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type ConfigureVirtualRouterElementParams struct {
	P map[string]interface{}
}

func (P *ConfigureVirtualRouterElementParams) toURLValues() url.Values {
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

func (P *ConfigureVirtualRouterElementParams) SetEnabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enabled"] = v
}

func (P *ConfigureVirtualRouterElementParams) GetEnabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enabled"].(bool)
	return value, ok
}

func (P *ConfigureVirtualRouterElementParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ConfigureVirtualRouterElementParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ConfigureVirtualRouterElementParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewConfigureVirtualRouterElementParams(enabled bool, id string) *ConfigureVirtualRouterElementParams {
	P := &ConfigureVirtualRouterElementParams{}
	P.P = make(map[string]interface{})
	P.P["enabled"] = enabled
	P.P["id"] = id
	return P
}

// Configures a virtual router element.
func (s *RouterService) ConfigureVirtualRouterElement(p *ConfigureVirtualRouterElementParams) (*VirtualRouterElementResponse, error) {
	resp, err := s.cs.newRequest("configureVirtualRouterElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r VirtualRouterElementResponse
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

type VirtualRouterElementResponse struct {
	Account   string `json:"account"`
	Domain    string `json:"domain"`
	Domainid  string `json:"domainid"`
	Enabled   bool   `json:"enabled"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Nspid     string `json:"nspid"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
}

type CreateVirtualRouterElementParams struct {
	P map[string]interface{}
}

func (P *CreateVirtualRouterElementParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["nspid"]; found {
		u.Set("nspid", v.(string))
	}
	if v, found := P.P["providertype"]; found {
		u.Set("providertype", v.(string))
	}
	return u
}

func (P *CreateVirtualRouterElementParams) SetNspid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nspid"] = v
}

func (P *CreateVirtualRouterElementParams) GetNspid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nspid"].(string)
	return value, ok
}

func (P *CreateVirtualRouterElementParams) SetProvidertype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["providertype"] = v
}

func (P *CreateVirtualRouterElementParams) GetProvidertype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["providertype"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVirtualRouterElementParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewCreateVirtualRouterElementParams(nspid string) *CreateVirtualRouterElementParams {
	P := &CreateVirtualRouterElementParams{}
	P.P = make(map[string]interface{})
	P.P["nspid"] = nspid
	return P
}

// Create a virtual router element.
func (s *RouterService) CreateVirtualRouterElement(p *CreateVirtualRouterElementParams) (*CreateVirtualRouterElementResponse, error) {
	resp, err := s.cs.newRequest("createVirtualRouterElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVirtualRouterElementResponse
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

type CreateVirtualRouterElementResponse struct {
	Account   string `json:"account"`
	Domain    string `json:"domain"`
	Domainid  string `json:"domainid"`
	Enabled   bool   `json:"enabled"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Nspid     string `json:"nspid"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
}

type DestroyRouterParams struct {
	P map[string]interface{}
}

func (P *DestroyRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DestroyRouterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DestroyRouterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DestroyRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewDestroyRouterParams(id string) *DestroyRouterParams {
	P := &DestroyRouterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Destroys a router.
func (s *RouterService) DestroyRouter(p *DestroyRouterParams) (*DestroyRouterResponse, error) {
	resp, err := s.cs.newRequest("destroyRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroyRouterResponse
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

type DestroyRouterResponse struct {
	Account             string                                    `json:"account"`
	Created             string                                    `json:"created"`
	Dns1                string                                    `json:"dns1"`
	Dns2                string                                    `json:"dns2"`
	Domain              string                                    `json:"domain"`
	Domainid            string                                    `json:"domainid"`
	Gateway             string                                    `json:"gateway"`
	Guestipaddress      string                                    `json:"guestipaddress"`
	Guestmacaddress     string                                    `json:"guestmacaddress"`
	Guestnetmask        string                                    `json:"guestnetmask"`
	Guestnetworkid      string                                    `json:"guestnetworkid"`
	Guestnetworkname    string                                    `json:"guestnetworkname"`
	Hasannotations      bool                                      `json:"hasannotations"`
	Healthcheckresults  []DestroyRouterResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                      `json:"healthchecksfailed"`
	Hostid              string                                    `json:"hostid"`
	Hostname            string                                    `json:"hostname"`
	Hypervisor          string                                    `json:"hypervisor"`
	Id                  string                                    `json:"id"`
	Ip6dns1             string                                    `json:"ip6dns1"`
	Ip6dns2             string                                    `json:"ip6dns2"`
	Isredundantrouter   bool                                      `json:"isredundantrouter"`
	JobID               string                                    `json:"jobid"`
	Jobstatus           int                                       `json:"jobstatus"`
	Linklocalip         string                                    `json:"linklocalip"`
	Linklocalmacaddress string                                    `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                    `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                    `json:"linklocalnetworkid"`
	Name                string                                    `json:"name"`
	Networkdomain       string                                    `json:"networkdomain"`
	Nic                 []Nic                                     `json:"nic"`
	Podid               string                                    `json:"podid"`
	Podname             string                                    `json:"podname"`
	Project             string                                    `json:"project"`
	Projectid           string                                    `json:"projectid"`
	Publicip            string                                    `json:"publicip"`
	Publicmacaddress    string                                    `json:"publicmacaddress"`
	Publicnetmask       string                                    `json:"publicnetmask"`
	Publicnetworkid     string                                    `json:"publicnetworkid"`
	Redundantstate      string                                    `json:"redundantstate"`
	Requiresupgrade     bool                                      `json:"requiresupgrade"`
	Role                string                                    `json:"role"`
	Scriptsversion      string                                    `json:"scriptsversion"`
	Serviceofferingid   string                                    `json:"serviceofferingid"`
	Serviceofferingname string                                    `json:"serviceofferingname"`
	State               string                                    `json:"state"`
	Templateid          string                                    `json:"templateid"`
	Templatename        string                                    `json:"templatename"`
	Version             string                                    `json:"version"`
	Vpcid               string                                    `json:"vpcid"`
	Vpcname             string                                    `json:"vpcname"`
	Zoneid              string                                    `json:"zoneid"`
	Zonename            string                                    `json:"zonename"`
}

type DestroyRouterResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type ListRoutersParams struct {
	P map[string]interface{}
}

func (P *ListRoutersParams) toURLValues() url.Values {
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
	if v, found := P.P["fetchhealthcheckresults"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fetchhealthcheckresults", vv)
	}
	if v, found := P.P["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := P.P["healthchecksfailed"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("healthchecksfailed", vv)
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
	if v, found := P.P["version"]; found {
		u.Set("version", v.(string))
	}
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListRoutersParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListRoutersParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *ListRoutersParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListRoutersParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetFetchhealthcheckresults(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fetchhealthcheckresults"] = v
}

func (P *ListRoutersParams) GetFetchhealthcheckresults() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fetchhealthcheckresults"].(bool)
	return value, ok
}

func (P *ListRoutersParams) SetForvpc(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forvpc"] = v
}

func (P *ListRoutersParams) GetForvpc() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forvpc"].(bool)
	return value, ok
}

func (P *ListRoutersParams) SetHealthchecksfailed(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["healthchecksfailed"] = v
}

func (P *ListRoutersParams) GetHealthchecksfailed() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["healthchecksfailed"].(bool)
	return value, ok
}

func (P *ListRoutersParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ListRoutersParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListRoutersParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListRoutersParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListRoutersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListRoutersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListRoutersParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListRoutersParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListRoutersParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListRoutersParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListRoutersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListRoutersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListRoutersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListRoutersParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListRoutersParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListRoutersParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListRoutersParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetVersion(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["version"] = v
}

func (P *ListRoutersParams) GetVersion() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["version"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *ListRoutersParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

func (P *ListRoutersParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListRoutersParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListRoutersParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewListRoutersParams() *ListRoutersParams {
	P := &ListRoutersParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetRouterID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListRoutersParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListRouters(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Routers[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Routers {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetRouterByName(name string, opts ...OptionFunc) (*Router, int, error) {
	id, count, err := s.GetRouterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetRouterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetRouterByID(id string, opts ...OptionFunc) (*Router, int, error) {
	P := &ListRoutersParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListRouters(P)
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
		return l.Routers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Router UUID: %s!", id)
}

// List routers.
func (s *RouterService) ListRouters(p *ListRoutersParams) (*ListRoutersResponse, error) {
	resp, err := s.cs.newRequest("listRouters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRoutersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListRoutersResponse struct {
	Count   int       `json:"count"`
	Routers []*Router `json:"router"`
}

type Router struct {
	Account             string                     `json:"account"`
	Created             string                     `json:"created"`
	Dns1                string                     `json:"dns1"`
	Dns2                string                     `json:"dns2"`
	Domain              string                     `json:"domain"`
	Domainid            string                     `json:"domainid"`
	Gateway             string                     `json:"gateway"`
	Guestipaddress      string                     `json:"guestipaddress"`
	Guestmacaddress     string                     `json:"guestmacaddress"`
	Guestnetmask        string                     `json:"guestnetmask"`
	Guestnetworkid      string                     `json:"guestnetworkid"`
	Guestnetworkname    string                     `json:"guestnetworkname"`
	Hasannotations      bool                       `json:"hasannotations"`
	Healthcheckresults  []RouterHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                       `json:"healthchecksfailed"`
	Hostid              string                     `json:"hostid"`
	Hostname            string                     `json:"hostname"`
	Hypervisor          string                     `json:"hypervisor"`
	Id                  string                     `json:"id"`
	Ip6dns1             string                     `json:"ip6dns1"`
	Ip6dns2             string                     `json:"ip6dns2"`
	Isredundantrouter   bool                       `json:"isredundantrouter"`
	JobID               string                     `json:"jobid"`
	Jobstatus           int                        `json:"jobstatus"`
	Linklocalip         string                     `json:"linklocalip"`
	Linklocalmacaddress string                     `json:"linklocalmacaddress"`
	Linklocalnetmask    string                     `json:"linklocalnetmask"`
	Linklocalnetworkid  string                     `json:"linklocalnetworkid"`
	Name                string                     `json:"name"`
	Networkdomain       string                     `json:"networkdomain"`
	Nic                 []Nic                      `json:"nic"`
	Podid               string                     `json:"podid"`
	Podname             string                     `json:"podname"`
	Project             string                     `json:"project"`
	Projectid           string                     `json:"projectid"`
	Publicip            string                     `json:"publicip"`
	Publicmacaddress    string                     `json:"publicmacaddress"`
	Publicnetmask       string                     `json:"publicnetmask"`
	Publicnetworkid     string                     `json:"publicnetworkid"`
	Redundantstate      string                     `json:"redundantstate"`
	Requiresupgrade     bool                       `json:"requiresupgrade"`
	Role                string                     `json:"role"`
	Scriptsversion      string                     `json:"scriptsversion"`
	Serviceofferingid   string                     `json:"serviceofferingid"`
	Serviceofferingname string                     `json:"serviceofferingname"`
	State               string                     `json:"state"`
	Templateid          string                     `json:"templateid"`
	Templatename        string                     `json:"templatename"`
	Version             string                     `json:"version"`
	Vpcid               string                     `json:"vpcid"`
	Vpcname             string                     `json:"vpcname"`
	Zoneid              string                     `json:"zoneid"`
	Zonename            string                     `json:"zonename"`
}

type RouterHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type ListVirtualRouterElementsParams struct {
	P map[string]interface{}
}

func (P *ListVirtualRouterElementsParams) toURLValues() url.Values {
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

func (P *ListVirtualRouterElementsParams) SetEnabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enabled"] = v
}

func (P *ListVirtualRouterElementsParams) GetEnabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enabled"].(bool)
	return value, ok
}

func (P *ListVirtualRouterElementsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListVirtualRouterElementsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListVirtualRouterElementsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListVirtualRouterElementsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListVirtualRouterElementsParams) SetNspid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nspid"] = v
}

func (P *ListVirtualRouterElementsParams) GetNspid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nspid"].(string)
	return value, ok
}

func (P *ListVirtualRouterElementsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListVirtualRouterElementsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListVirtualRouterElementsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListVirtualRouterElementsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListVirtualRouterElementsParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewListVirtualRouterElementsParams() *ListVirtualRouterElementsParams {
	P := &ListVirtualRouterElementsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RouterService) GetVirtualRouterElementByID(id string, opts ...OptionFunc) (*VirtualRouterElement, int, error) {
	P := &ListVirtualRouterElementsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVirtualRouterElements(P)
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
		return l.VirtualRouterElements[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VirtualRouterElement UUID: %s!", id)
}

// Lists all available virtual router elements.
func (s *RouterService) ListVirtualRouterElements(p *ListVirtualRouterElementsParams) (*ListVirtualRouterElementsResponse, error) {
	resp, err := s.cs.newRequest("listVirtualRouterElements", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVirtualRouterElementsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVirtualRouterElementsResponse struct {
	Count                 int                     `json:"count"`
	VirtualRouterElements []*VirtualRouterElement `json:"virtualrouterelement"`
}

type VirtualRouterElement struct {
	Account   string `json:"account"`
	Domain    string `json:"domain"`
	Domainid  string `json:"domainid"`
	Enabled   bool   `json:"enabled"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Nspid     string `json:"nspid"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
}

type RebootRouterParams struct {
	P map[string]interface{}
}

func (P *RebootRouterParams) toURLValues() url.Values {
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

func (P *RebootRouterParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *RebootRouterParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *RebootRouterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RebootRouterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RebootRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewRebootRouterParams(id string) *RebootRouterParams {
	P := &RebootRouterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Starts a router.
func (s *RouterService) RebootRouter(p *RebootRouterParams) (*RebootRouterResponse, error) {
	resp, err := s.cs.newRequest("rebootRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootRouterResponse
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

type RebootRouterResponse struct {
	Account             string                                   `json:"account"`
	Created             string                                   `json:"created"`
	Dns1                string                                   `json:"dns1"`
	Dns2                string                                   `json:"dns2"`
	Domain              string                                   `json:"domain"`
	Domainid            string                                   `json:"domainid"`
	Gateway             string                                   `json:"gateway"`
	Guestipaddress      string                                   `json:"guestipaddress"`
	Guestmacaddress     string                                   `json:"guestmacaddress"`
	Guestnetmask        string                                   `json:"guestnetmask"`
	Guestnetworkid      string                                   `json:"guestnetworkid"`
	Guestnetworkname    string                                   `json:"guestnetworkname"`
	Hasannotations      bool                                     `json:"hasannotations"`
	Healthcheckresults  []RebootRouterResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                     `json:"healthchecksfailed"`
	Hostid              string                                   `json:"hostid"`
	Hostname            string                                   `json:"hostname"`
	Hypervisor          string                                   `json:"hypervisor"`
	Id                  string                                   `json:"id"`
	Ip6dns1             string                                   `json:"ip6dns1"`
	Ip6dns2             string                                   `json:"ip6dns2"`
	Isredundantrouter   bool                                     `json:"isredundantrouter"`
	JobID               string                                   `json:"jobid"`
	Jobstatus           int                                      `json:"jobstatus"`
	Linklocalip         string                                   `json:"linklocalip"`
	Linklocalmacaddress string                                   `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                   `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                   `json:"linklocalnetworkid"`
	Name                string                                   `json:"name"`
	Networkdomain       string                                   `json:"networkdomain"`
	Nic                 []Nic                                    `json:"nic"`
	Podid               string                                   `json:"podid"`
	Podname             string                                   `json:"podname"`
	Project             string                                   `json:"project"`
	Projectid           string                                   `json:"projectid"`
	Publicip            string                                   `json:"publicip"`
	Publicmacaddress    string                                   `json:"publicmacaddress"`
	Publicnetmask       string                                   `json:"publicnetmask"`
	Publicnetworkid     string                                   `json:"publicnetworkid"`
	Redundantstate      string                                   `json:"redundantstate"`
	Requiresupgrade     bool                                     `json:"requiresupgrade"`
	Role                string                                   `json:"role"`
	Scriptsversion      string                                   `json:"scriptsversion"`
	Serviceofferingid   string                                   `json:"serviceofferingid"`
	Serviceofferingname string                                   `json:"serviceofferingname"`
	State               string                                   `json:"state"`
	Templateid          string                                   `json:"templateid"`
	Templatename        string                                   `json:"templatename"`
	Version             string                                   `json:"version"`
	Vpcid               string                                   `json:"vpcid"`
	Vpcname             string                                   `json:"vpcname"`
	Zoneid              string                                   `json:"zoneid"`
	Zonename            string                                   `json:"zonename"`
}

type RebootRouterResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type StartRouterParams struct {
	P map[string]interface{}
}

func (P *StartRouterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *StartRouterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *StartRouterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new StartRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewStartRouterParams(id string) *StartRouterParams {
	P := &StartRouterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Starts a router.
func (s *RouterService) StartRouter(p *StartRouterParams) (*StartRouterResponse, error) {
	resp, err := s.cs.newRequest("startRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartRouterResponse
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

type StartRouterResponse struct {
	Account             string                                  `json:"account"`
	Created             string                                  `json:"created"`
	Dns1                string                                  `json:"dns1"`
	Dns2                string                                  `json:"dns2"`
	Domain              string                                  `json:"domain"`
	Domainid            string                                  `json:"domainid"`
	Gateway             string                                  `json:"gateway"`
	Guestipaddress      string                                  `json:"guestipaddress"`
	Guestmacaddress     string                                  `json:"guestmacaddress"`
	Guestnetmask        string                                  `json:"guestnetmask"`
	Guestnetworkid      string                                  `json:"guestnetworkid"`
	Guestnetworkname    string                                  `json:"guestnetworkname"`
	Hasannotations      bool                                    `json:"hasannotations"`
	Healthcheckresults  []StartRouterResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                    `json:"healthchecksfailed"`
	Hostid              string                                  `json:"hostid"`
	Hostname            string                                  `json:"hostname"`
	Hypervisor          string                                  `json:"hypervisor"`
	Id                  string                                  `json:"id"`
	Ip6dns1             string                                  `json:"ip6dns1"`
	Ip6dns2             string                                  `json:"ip6dns2"`
	Isredundantrouter   bool                                    `json:"isredundantrouter"`
	JobID               string                                  `json:"jobid"`
	Jobstatus           int                                     `json:"jobstatus"`
	Linklocalip         string                                  `json:"linklocalip"`
	Linklocalmacaddress string                                  `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                  `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                  `json:"linklocalnetworkid"`
	Name                string                                  `json:"name"`
	Networkdomain       string                                  `json:"networkdomain"`
	Nic                 []Nic                                   `json:"nic"`
	Podid               string                                  `json:"podid"`
	Podname             string                                  `json:"podname"`
	Project             string                                  `json:"project"`
	Projectid           string                                  `json:"projectid"`
	Publicip            string                                  `json:"publicip"`
	Publicmacaddress    string                                  `json:"publicmacaddress"`
	Publicnetmask       string                                  `json:"publicnetmask"`
	Publicnetworkid     string                                  `json:"publicnetworkid"`
	Redundantstate      string                                  `json:"redundantstate"`
	Requiresupgrade     bool                                    `json:"requiresupgrade"`
	Role                string                                  `json:"role"`
	Scriptsversion      string                                  `json:"scriptsversion"`
	Serviceofferingid   string                                  `json:"serviceofferingid"`
	Serviceofferingname string                                  `json:"serviceofferingname"`
	State               string                                  `json:"state"`
	Templateid          string                                  `json:"templateid"`
	Templatename        string                                  `json:"templatename"`
	Version             string                                  `json:"version"`
	Vpcid               string                                  `json:"vpcid"`
	Vpcname             string                                  `json:"vpcname"`
	Zoneid              string                                  `json:"zoneid"`
	Zonename            string                                  `json:"zonename"`
}

type StartRouterResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}

type StopRouterParams struct {
	P map[string]interface{}
}

func (P *StopRouterParams) toURLValues() url.Values {
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

func (P *StopRouterParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *StopRouterParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *StopRouterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *StopRouterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new StopRouterParams instance,
// as then you are sure you have configured all required params
func (s *RouterService) NewStopRouterParams(id string) *StopRouterParams {
	P := &StopRouterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Stops a router.
func (s *RouterService) StopRouter(p *StopRouterParams) (*StopRouterResponse, error) {
	resp, err := s.cs.newRequest("stopRouter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopRouterResponse
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

type StopRouterResponse struct {
	Account             string                                 `json:"account"`
	Created             string                                 `json:"created"`
	Dns1                string                                 `json:"dns1"`
	Dns2                string                                 `json:"dns2"`
	Domain              string                                 `json:"domain"`
	Domainid            string                                 `json:"domainid"`
	Gateway             string                                 `json:"gateway"`
	Guestipaddress      string                                 `json:"guestipaddress"`
	Guestmacaddress     string                                 `json:"guestmacaddress"`
	Guestnetmask        string                                 `json:"guestnetmask"`
	Guestnetworkid      string                                 `json:"guestnetworkid"`
	Guestnetworkname    string                                 `json:"guestnetworkname"`
	Hasannotations      bool                                   `json:"hasannotations"`
	Healthcheckresults  []StopRouterResponseHealthcheckresults `json:"healthcheckresults"`
	Healthchecksfailed  bool                                   `json:"healthchecksfailed"`
	Hostid              string                                 `json:"hostid"`
	Hostname            string                                 `json:"hostname"`
	Hypervisor          string                                 `json:"hypervisor"`
	Id                  string                                 `json:"id"`
	Ip6dns1             string                                 `json:"ip6dns1"`
	Ip6dns2             string                                 `json:"ip6dns2"`
	Isredundantrouter   bool                                   `json:"isredundantrouter"`
	JobID               string                                 `json:"jobid"`
	Jobstatus           int                                    `json:"jobstatus"`
	Linklocalip         string                                 `json:"linklocalip"`
	Linklocalmacaddress string                                 `json:"linklocalmacaddress"`
	Linklocalnetmask    string                                 `json:"linklocalnetmask"`
	Linklocalnetworkid  string                                 `json:"linklocalnetworkid"`
	Name                string                                 `json:"name"`
	Networkdomain       string                                 `json:"networkdomain"`
	Nic                 []Nic                                  `json:"nic"`
	Podid               string                                 `json:"podid"`
	Podname             string                                 `json:"podname"`
	Project             string                                 `json:"project"`
	Projectid           string                                 `json:"projectid"`
	Publicip            string                                 `json:"publicip"`
	Publicmacaddress    string                                 `json:"publicmacaddress"`
	Publicnetmask       string                                 `json:"publicnetmask"`
	Publicnetworkid     string                                 `json:"publicnetworkid"`
	Redundantstate      string                                 `json:"redundantstate"`
	Requiresupgrade     bool                                   `json:"requiresupgrade"`
	Role                string                                 `json:"role"`
	Scriptsversion      string                                 `json:"scriptsversion"`
	Serviceofferingid   string                                 `json:"serviceofferingid"`
	Serviceofferingname string                                 `json:"serviceofferingname"`
	State               string                                 `json:"state"`
	Templateid          string                                 `json:"templateid"`
	Templatename        string                                 `json:"templatename"`
	Version             string                                 `json:"version"`
	Vpcid               string                                 `json:"vpcid"`
	Vpcname             string                                 `json:"vpcname"`
	Zoneid              string                                 `json:"zoneid"`
	Zonename            string                                 `json:"zonename"`
}

type StopRouterResponseHealthcheckresults struct {
	Checkname   string `json:"checkname"`
	Checktype   string `json:"checktype"`
	Details     string `json:"details"`
	Lastupdated string `json:"lastupdated"`
	Success     bool   `json:"success"`
}
