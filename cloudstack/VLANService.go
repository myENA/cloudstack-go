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

type VLANServiceIface interface {
	CreateVlanIpRange(p *CreateVlanIpRangeParams) (*CreateVlanIpRangeResponse, error)
	NewCreateVlanIpRangeParams() *CreateVlanIpRangeParams
	DedicateGuestVlanRange(p *DedicateGuestVlanRangeParams) (*DedicateGuestVlanRangeResponse, error)
	NewDedicateGuestVlanRangeParams(physicalnetworkid string, vlanrange string) *DedicateGuestVlanRangeParams
	DeleteVlanIpRange(p *DeleteVlanIpRangeParams) (*DeleteVlanIpRangeResponse, error)
	NewDeleteVlanIpRangeParams(id string) *DeleteVlanIpRangeParams
	ListDedicatedGuestVlanRanges(p *ListDedicatedGuestVlanRangesParams) (*ListDedicatedGuestVlanRangesResponse, error)
	NewListDedicatedGuestVlanRangesParams() *ListDedicatedGuestVlanRangesParams
	GetDedicatedGuestVlanRangeByID(id string, opts ...OptionFunc) (*DedicatedGuestVlanRange, int, error)
	ListVlanIpRanges(p *ListVlanIpRangesParams) (*ListVlanIpRangesResponse, error)
	NewListVlanIpRangesParams() *ListVlanIpRangesParams
	GetVlanIpRangeByID(id string, opts ...OptionFunc) (*VlanIpRange, int, error)
	ReleaseDedicatedGuestVlanRange(p *ReleaseDedicatedGuestVlanRangeParams) (*ReleaseDedicatedGuestVlanRangeResponse, error)
	NewReleaseDedicatedGuestVlanRangeParams(id string) *ReleaseDedicatedGuestVlanRangeParams
}

type CreateVlanIpRangeParams struct {
	P map[string]interface{}
}

func (P *CreateVlanIpRangeParams) toURLValues() url.Values {
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
	if v, found := P.P["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := P.P["endipv6"]; found {
		u.Set("endipv6", v.(string))
	}
	if v, found := P.P["forsystemvms"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forsystemvms", vv)
	}
	if v, found := P.P["forvirtualnetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvirtualnetwork", vv)
	}
	if v, found := P.P["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := P.P["ip6cidr"]; found {
		u.Set("ip6cidr", v.(string))
	}
	if v, found := P.P["ip6gateway"]; found {
		u.Set("ip6gateway", v.(string))
	}
	if v, found := P.P["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := P.P["startipv6"]; found {
		u.Set("startipv6", v.(string))
	}
	if v, found := P.P["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *CreateVlanIpRangeParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateVlanIpRangeParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateVlanIpRangeParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetEndip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endip"] = v
}

func (P *CreateVlanIpRangeParams) GetEndip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endip"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetEndipv6(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endipv6"] = v
}

func (P *CreateVlanIpRangeParams) GetEndipv6() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endipv6"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetForsystemvms(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forsystemvms"] = v
}

func (P *CreateVlanIpRangeParams) GetForsystemvms() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forsystemvms"].(bool)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetForvirtualnetwork(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forvirtualnetwork"] = v
}

func (P *CreateVlanIpRangeParams) GetForvirtualnetwork() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forvirtualnetwork"].(bool)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetGateway(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gateway"] = v
}

func (P *CreateVlanIpRangeParams) GetGateway() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gateway"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetIp6cidr(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ip6cidr"] = v
}

func (P *CreateVlanIpRangeParams) GetIp6cidr() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ip6cidr"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetIp6gateway(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ip6gateway"] = v
}

func (P *CreateVlanIpRangeParams) GetIp6gateway() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ip6gateway"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetNetmask(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["netmask"] = v
}

func (P *CreateVlanIpRangeParams) GetNetmask() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["netmask"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *CreateVlanIpRangeParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *CreateVlanIpRangeParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *CreateVlanIpRangeParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateVlanIpRangeParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetStartip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startip"] = v
}

func (P *CreateVlanIpRangeParams) GetStartip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startip"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetStartipv6(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startipv6"] = v
}

func (P *CreateVlanIpRangeParams) GetStartipv6() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startipv6"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetVlan(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vlan"] = v
}

func (P *CreateVlanIpRangeParams) GetVlan() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vlan"].(string)
	return value, ok
}

func (P *CreateVlanIpRangeParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateVlanIpRangeParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVlanIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *VLANService) NewCreateVlanIpRangeParams() *CreateVlanIpRangeParams {
	P := &CreateVlanIpRangeParams{}
	P.P = make(map[string]interface{})
	return P
}

// Creates a VLAN IP range.
func (s *VLANService) CreateVlanIpRange(p *CreateVlanIpRangeParams) (*CreateVlanIpRangeResponse, error) {
	resp, err := s.cs.newRequest("createVlanIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	//var r CreateVlanIpRangeResponse
	var r struct {
		Vlan CreateVlanIpRangeResponse `json:"vlan"`
	}
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r.Vlan, nil
}

type CreateVlanIpRangeResponse struct {
	Account           string `json:"account"`
	Description       string `json:"description"`
	Domain            string `json:"domain"`
	Domainid          string `json:"domainid"`
	Endip             string `json:"endip"`
	Endipv6           string `json:"endipv6"`
	Forsystemvms      bool   `json:"forsystemvms"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork"`
	Gateway           string `json:"gateway"`
	Id                string `json:"id"`
	Ip6cidr           string `json:"ip6cidr"`
	Ip6gateway        string `json:"ip6gateway"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Netmask           string `json:"netmask"`
	Networkid         string `json:"networkid"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Project           string `json:"project"`
	Projectid         string `json:"projectid"`
	Startip           string `json:"startip"`
	Startipv6         string `json:"startipv6"`
	Vlan              string `json:"vlan"`
	Zoneid            string `json:"zoneid"`
}

type DedicateGuestVlanRangeParams struct {
	P map[string]interface{}
}

func (P *DedicateGuestVlanRangeParams) toURLValues() url.Values {
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
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["vlanrange"]; found {
		u.Set("vlanrange", v.(string))
	}
	return u
}

func (P *DedicateGuestVlanRangeParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DedicateGuestVlanRangeParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DedicateGuestVlanRangeParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DedicateGuestVlanRangeParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *DedicateGuestVlanRangeParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *DedicateGuestVlanRangeParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *DedicateGuestVlanRangeParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *DedicateGuestVlanRangeParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *DedicateGuestVlanRangeParams) SetVlanrange(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vlanrange"] = v
}

func (P *DedicateGuestVlanRangeParams) GetVlanrange() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vlanrange"].(string)
	return value, ok
}

// You should always use this function to get a new DedicateGuestVlanRangeParams instance,
// as then you are sure you have configured all required params
func (s *VLANService) NewDedicateGuestVlanRangeParams(physicalnetworkid string, vlanrange string) *DedicateGuestVlanRangeParams {
	P := &DedicateGuestVlanRangeParams{}
	P.P = make(map[string]interface{})
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["vlanrange"] = vlanrange
	return P
}

// Dedicates a guest vlan range to an account
func (s *VLANService) DedicateGuestVlanRange(p *DedicateGuestVlanRangeParams) (*DedicateGuestVlanRangeResponse, error) {
	resp, err := s.cs.newRequest("dedicateGuestVlanRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r DedicateGuestVlanRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DedicateGuestVlanRangeResponse struct {
	Account           string `json:"account"`
	Domain            string `json:"domain"`
	Domainid          string `json:"domainid"`
	Guestvlanrange    string `json:"guestvlanrange"`
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid int64  `json:"physicalnetworkid"`
	Project           string `json:"project"`
	Projectid         string `json:"projectid"`
	Zoneid            int64  `json:"zoneid"`
}

type DeleteVlanIpRangeParams struct {
	P map[string]interface{}
}

func (P *DeleteVlanIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteVlanIpRangeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteVlanIpRangeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVlanIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *VLANService) NewDeleteVlanIpRangeParams(id string) *DeleteVlanIpRangeParams {
	P := &DeleteVlanIpRangeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Creates a VLAN IP range.
func (s *VLANService) DeleteVlanIpRange(p *DeleteVlanIpRangeParams) (*DeleteVlanIpRangeResponse, error) {
	resp, err := s.cs.newRequest("deleteVlanIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVlanIpRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteVlanIpRangeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteVlanIpRangeResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteVlanIpRangeResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListDedicatedGuestVlanRangesParams struct {
	P map[string]interface{}
}

func (P *ListDedicatedGuestVlanRangesParams) toURLValues() url.Values {
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
	if v, found := P.P["guestvlanrange"]; found {
		u.Set("guestvlanrange", v.(string))
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
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListDedicatedGuestVlanRangesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListDedicatedGuestVlanRangesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListDedicatedGuestVlanRangesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListDedicatedGuestVlanRangesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListDedicatedGuestVlanRangesParams) SetGuestvlanrange(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["guestvlanrange"] = v
}

func (P *ListDedicatedGuestVlanRangesParams) GetGuestvlanrange() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["guestvlanrange"].(string)
	return value, ok
}

func (P *ListDedicatedGuestVlanRangesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListDedicatedGuestVlanRangesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListDedicatedGuestVlanRangesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListDedicatedGuestVlanRangesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListDedicatedGuestVlanRangesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListDedicatedGuestVlanRangesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListDedicatedGuestVlanRangesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListDedicatedGuestVlanRangesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListDedicatedGuestVlanRangesParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListDedicatedGuestVlanRangesParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *ListDedicatedGuestVlanRangesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListDedicatedGuestVlanRangesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListDedicatedGuestVlanRangesParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListDedicatedGuestVlanRangesParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListDedicatedGuestVlanRangesParams instance,
// as then you are sure you have configured all required params
func (s *VLANService) NewListDedicatedGuestVlanRangesParams() *ListDedicatedGuestVlanRangesParams {
	P := &ListDedicatedGuestVlanRangesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VLANService) GetDedicatedGuestVlanRangeByID(id string, opts ...OptionFunc) (*DedicatedGuestVlanRange, int, error) {
	P := &ListDedicatedGuestVlanRangesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListDedicatedGuestVlanRanges(P)
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
		return l.DedicatedGuestVlanRanges[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for DedicatedGuestVlanRange UUID: %s!", id)
}

// Lists dedicated guest vlan ranges
func (s *VLANService) ListDedicatedGuestVlanRanges(p *ListDedicatedGuestVlanRangesParams) (*ListDedicatedGuestVlanRangesResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedGuestVlanRanges", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedGuestVlanRangesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedGuestVlanRangesResponse struct {
	Count                    int                        `json:"count"`
	DedicatedGuestVlanRanges []*DedicatedGuestVlanRange `json:"dedicatedguestvlanrange"`
}

type DedicatedGuestVlanRange struct {
	Account           string `json:"account"`
	Domain            string `json:"domain"`
	Domainid          string `json:"domainid"`
	Guestvlanrange    string `json:"guestvlanrange"`
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid int64  `json:"physicalnetworkid"`
	Project           string `json:"project"`
	Projectid         string `json:"projectid"`
	Zoneid            int64  `json:"zoneid"`
}

type ListVlanIpRangesParams struct {
	P map[string]interface{}
}

func (P *ListVlanIpRangesParams) toURLValues() url.Values {
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
	if v, found := P.P["forvirtualnetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvirtualnetwork", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
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
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListVlanIpRangesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListVlanIpRangesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListVlanIpRangesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetForvirtualnetwork(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forvirtualnetwork"] = v
}

func (P *ListVlanIpRangesParams) GetForvirtualnetwork() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forvirtualnetwork"].(bool)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListVlanIpRangesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListVlanIpRangesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListVlanIpRangesParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListVlanIpRangesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListVlanIpRangesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListVlanIpRangesParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListVlanIpRangesParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListVlanIpRangesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetVlan(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vlan"] = v
}

func (P *ListVlanIpRangesParams) GetVlan() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vlan"].(string)
	return value, ok
}

func (P *ListVlanIpRangesParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListVlanIpRangesParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVlanIpRangesParams instance,
// as then you are sure you have configured all required params
func (s *VLANService) NewListVlanIpRangesParams() *ListVlanIpRangesParams {
	P := &ListVlanIpRangesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VLANService) GetVlanIpRangeByID(id string, opts ...OptionFunc) (*VlanIpRange, int, error) {
	P := &ListVlanIpRangesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVlanIpRanges(P)
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
		return l.VlanIpRanges[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VlanIpRange UUID: %s!", id)
}

// Lists all VLAN IP ranges.
func (s *VLANService) ListVlanIpRanges(p *ListVlanIpRangesParams) (*ListVlanIpRangesResponse, error) {
	resp, err := s.cs.newRequest("listVlanIpRanges", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVlanIpRangesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVlanIpRangesResponse struct {
	Count        int            `json:"count"`
	VlanIpRanges []*VlanIpRange `json:"vlaniprange"`
}

type VlanIpRange struct {
	Account           string `json:"account"`
	Description       string `json:"description"`
	Domain            string `json:"domain"`
	Domainid          string `json:"domainid"`
	Endip             string `json:"endip"`
	Endipv6           string `json:"endipv6"`
	Forsystemvms      bool   `json:"forsystemvms"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork"`
	Gateway           string `json:"gateway"`
	Id                string `json:"id"`
	Ip6cidr           string `json:"ip6cidr"`
	Ip6gateway        string `json:"ip6gateway"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Netmask           string `json:"netmask"`
	Networkid         string `json:"networkid"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Project           string `json:"project"`
	Projectid         string `json:"projectid"`
	Startip           string `json:"startip"`
	Startipv6         string `json:"startipv6"`
	Vlan              string `json:"vlan"`
	Zoneid            string `json:"zoneid"`
}

type ReleaseDedicatedGuestVlanRangeParams struct {
	P map[string]interface{}
}

func (P *ReleaseDedicatedGuestVlanRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ReleaseDedicatedGuestVlanRangeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ReleaseDedicatedGuestVlanRangeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseDedicatedGuestVlanRangeParams instance,
// as then you are sure you have configured all required params
func (s *VLANService) NewReleaseDedicatedGuestVlanRangeParams(id string) *ReleaseDedicatedGuestVlanRangeParams {
	P := &ReleaseDedicatedGuestVlanRangeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Releases a dedicated guest vlan range to the system
func (s *VLANService) ReleaseDedicatedGuestVlanRange(p *ReleaseDedicatedGuestVlanRangeParams) (*ReleaseDedicatedGuestVlanRangeResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedGuestVlanRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedGuestVlanRangeResponse
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

type ReleaseDedicatedGuestVlanRangeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
