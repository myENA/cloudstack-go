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

type AddressServiceIface interface {
	AssociateIpAddress(P *AssociateIpAddressParams) (*AssociateIpAddressResponse, error)
	NewAssociateIpAddressParams() *AssociateIpAddressParams
	DisassociateIpAddress(P *DisassociateIpAddressParams) (*DisassociateIpAddressResponse, error)
	NewDisassociateIpAddressParams(id string) *DisassociateIpAddressParams
	ListPublicIpAddresses(P *ListPublicIpAddressesParams) (*ListPublicIpAddressesResponse, error)
	NewListPublicIpAddressesParams() *ListPublicIpAddressesParams
	GetPublicIpAddressByID(id string, opts ...OptionFunc) (*PublicIpAddress, int, error)
	UpdateIpAddress(P *UpdateIpAddressParams) (*UpdateIpAddressResponse, error)
	NewUpdateIpAddressParams(id string) *UpdateIpAddressParams
}

type AssociateIpAddressParams struct {
	P map[string]interface{}
}

func (P *AssociateIpAddressParams) toURLValues() url.Values {
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
	if v, found := P.P["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := P.P["isportable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isportable", vv)
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
	}
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *AssociateIpAddressParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *AssociateIpAddressParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *AssociateIpAddressParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *AssociateIpAddressParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *AssociateIpAddressParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *AssociateIpAddressParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *AssociateIpAddressParams) SetIpaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddress"] = v
}

func (P *AssociateIpAddressParams) GetIpaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddress"].(string)
	return value, ok
}

func (P *AssociateIpAddressParams) SetIsportable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isportable"] = v
}

func (P *AssociateIpAddressParams) GetIsportable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isportable"].(bool)
	return value, ok
}

func (P *AssociateIpAddressParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *AssociateIpAddressParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *AssociateIpAddressParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *AssociateIpAddressParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *AssociateIpAddressParams) SetRegionid(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["regionid"] = v
}

func (P *AssociateIpAddressParams) GetRegionid() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["regionid"].(int)
	return value, ok
}

func (P *AssociateIpAddressParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *AssociateIpAddressParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

func (P *AssociateIpAddressParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *AssociateIpAddressParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AssociateIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewAssociateIpAddressParams() *AssociateIpAddressParams {
	P := &AssociateIpAddressParams{}
	P.P = make(map[string]interface{})
	return P
}

// Acquires and associates a public IP to an account.
func (s *AddressService) AssociateIpAddress(p *AssociateIpAddressParams) (*AssociateIpAddressResponse, error) {
	resp, err := s.cs.newRequest("associateIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssociateIpAddressResponse
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

type AssociateIpAddressResponse struct {
	Account                   string `json:"account"`
	Allocated                 string `json:"allocated"`
	Associatednetworkid       string `json:"associatednetworkid"`
	Associatednetworkname     string `json:"associatednetworkname"`
	Domain                    string `json:"domain"`
	Domainid                  string `json:"domainid"`
	Fordisplay                bool   `json:"fordisplay"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork"`
	Hasannotations            bool   `json:"hasannotations"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Isportable                bool   `json:"isportable"`
	Issourcenat               bool   `json:"issourcenat"`
	Isstaticnat               bool   `json:"isstaticnat"`
	Issystem                  bool   `json:"issystem"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Networkname               string `json:"networkname"`
	Physicalnetworkid         string `json:"physicalnetworkid"`
	Project                   string `json:"project"`
	Projectid                 string `json:"projectid"`
	Purpose                   string `json:"purpose"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vlanid                    string `json:"vlanid"`
	Vlanname                  string `json:"vlanname"`
	Vmipaddress               string `json:"vmipaddress"`
	Vpcid                     string `json:"vpcid"`
	Vpcname                   string `json:"vpcname"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type DisassociateIpAddressParams struct {
	P map[string]interface{}
}

func (P *DisassociateIpAddressParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DisassociateIpAddressParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DisassociateIpAddressParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DisassociateIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewDisassociateIpAddressParams(id string) *DisassociateIpAddressParams {
	P := &DisassociateIpAddressParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Disassociates an IP address from the account.
func (s *AddressService) DisassociateIpAddress(p *DisassociateIpAddressParams) (*DisassociateIpAddressResponse, error) {
	resp, err := s.cs.newRequest("disassociateIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisassociateIpAddressResponse
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

type DisassociateIpAddressResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListPublicIpAddressesParams struct {
	P map[string]interface{}
}

func (P *ListPublicIpAddressesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["allocatedonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("allocatedonly", vv)
	}
	if v, found := P.P["associatednetworkid"]; found {
		u.Set("associatednetworkid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["forloadbalancing"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forloadbalancing", vv)
	}
	if v, found := P.P["forvirtualnetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvirtualnetwork", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := P.P["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := P.P["issourcenat"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issourcenat", vv)
	}
	if v, found := P.P["isstaticnat"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isstaticnat", vv)
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
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
	if v, found := P.P["vlanid"]; found {
		u.Set("vlanid", v.(string))
	}
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListPublicIpAddressesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListPublicIpAddressesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetAllocatedonly(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocatedonly"] = v
}

func (P *ListPublicIpAddressesParams) GetAllocatedonly() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocatedonly"].(bool)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetAssociatednetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["associatednetworkid"] = v
}

func (P *ListPublicIpAddressesParams) GetAssociatednetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["associatednetworkid"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListPublicIpAddressesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListPublicIpAddressesParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetForloadbalancing(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forloadbalancing"] = v
}

func (P *ListPublicIpAddressesParams) GetForloadbalancing() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forloadbalancing"].(bool)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetForvirtualnetwork(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forvirtualnetwork"] = v
}

func (P *ListPublicIpAddressesParams) GetForvirtualnetwork() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forvirtualnetwork"].(bool)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListPublicIpAddressesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetIpaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddress"] = v
}

func (P *ListPublicIpAddressesParams) GetIpaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddress"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListPublicIpAddressesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetIssourcenat(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["issourcenat"] = v
}

func (P *ListPublicIpAddressesParams) GetIssourcenat() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["issourcenat"].(bool)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetIsstaticnat(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isstaticnat"] = v
}

func (P *ListPublicIpAddressesParams) GetIsstaticnat() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isstaticnat"].(bool)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListPublicIpAddressesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListPublicIpAddressesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListPublicIpAddressesParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListPublicIpAddressesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListPublicIpAddressesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListPublicIpAddressesParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListPublicIpAddressesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListPublicIpAddressesParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListPublicIpAddressesParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetVlanid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vlanid"] = v
}

func (P *ListPublicIpAddressesParams) GetVlanid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vlanid"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *ListPublicIpAddressesParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

func (P *ListPublicIpAddressesParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListPublicIpAddressesParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListPublicIpAddressesParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewListPublicIpAddressesParams() *ListPublicIpAddressesParams {
	P := &ListPublicIpAddressesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AddressService) GetPublicIpAddressByID(id string, opts ...OptionFunc) (*PublicIpAddress, int, error) {
	P := &ListPublicIpAddressesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListPublicIpAddresses(P)
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
		return l.PublicIpAddresses[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for PublicIpAddress UUID: %s!", id)
}

// Lists all public ip addresses
func (s *AddressService) ListPublicIpAddresses(p *ListPublicIpAddressesParams) (*ListPublicIpAddressesResponse, error) {
	resp, err := s.cs.newRequest("listPublicIpAddresses", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPublicIpAddressesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPublicIpAddressesResponse struct {
	Count             int                `json:"count"`
	PublicIpAddresses []*PublicIpAddress `json:"publicipaddress"`
}

type PublicIpAddress struct {
	Account                   string `json:"account"`
	Allocated                 string `json:"allocated"`
	Associatednetworkid       string `json:"associatednetworkid"`
	Associatednetworkname     string `json:"associatednetworkname"`
	Domain                    string `json:"domain"`
	Domainid                  string `json:"domainid"`
	Fordisplay                bool   `json:"fordisplay"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork"`
	Hasannotations            bool   `json:"hasannotations"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Isportable                bool   `json:"isportable"`
	Issourcenat               bool   `json:"issourcenat"`
	Isstaticnat               bool   `json:"isstaticnat"`
	Issystem                  bool   `json:"issystem"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Networkname               string `json:"networkname"`
	Physicalnetworkid         string `json:"physicalnetworkid"`
	Project                   string `json:"project"`
	Projectid                 string `json:"projectid"`
	Purpose                   string `json:"purpose"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vlanid                    string `json:"vlanid"`
	Vlanname                  string `json:"vlanname"`
	Vmipaddress               string `json:"vmipaddress"`
	Vpcid                     string `json:"vpcid"`
	Vpcname                   string `json:"vpcname"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}

type UpdateIpAddressParams struct {
	P map[string]interface{}
}

func (P *UpdateIpAddressParams) toURLValues() url.Values {
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

func (P *UpdateIpAddressParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateIpAddressParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateIpAddressParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateIpAddressParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateIpAddressParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateIpAddressParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateIpAddressParams instance,
// as then you are sure you have configured all required params
func (s *AddressService) NewUpdateIpAddressParams(id string) *UpdateIpAddressParams {
	P := &UpdateIpAddressParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates an IP address
func (s *AddressService) UpdateIpAddress(p *UpdateIpAddressParams) (*UpdateIpAddressResponse, error) {
	resp, err := s.cs.newRequest("updateIpAddress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateIpAddressResponse
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

type UpdateIpAddressResponse struct {
	Account                   string `json:"account"`
	Allocated                 string `json:"allocated"`
	Associatednetworkid       string `json:"associatednetworkid"`
	Associatednetworkname     string `json:"associatednetworkname"`
	Domain                    string `json:"domain"`
	Domainid                  string `json:"domainid"`
	Fordisplay                bool   `json:"fordisplay"`
	Forvirtualnetwork         bool   `json:"forvirtualnetwork"`
	Hasannotations            bool   `json:"hasannotations"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Isportable                bool   `json:"isportable"`
	Issourcenat               bool   `json:"issourcenat"`
	Isstaticnat               bool   `json:"isstaticnat"`
	Issystem                  bool   `json:"issystem"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Networkname               string `json:"networkname"`
	Physicalnetworkid         string `json:"physicalnetworkid"`
	Project                   string `json:"project"`
	Projectid                 string `json:"projectid"`
	Purpose                   string `json:"purpose"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vlanid                    string `json:"vlanid"`
	Vlanname                  string `json:"vlanname"`
	Vmipaddress               string `json:"vmipaddress"`
	Vpcid                     string `json:"vpcid"`
	Vpcname                   string `json:"vpcname"`
	Zoneid                    string `json:"zoneid"`
	Zonename                  string `json:"zonename"`
}
