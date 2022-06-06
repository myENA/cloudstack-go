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

type BrocadeVCSServiceIface interface {
	AddBrocadeVcsDevice(P *AddBrocadeVcsDeviceParams) (*AddBrocadeVcsDeviceResponse, error)
	NewAddBrocadeVcsDeviceParams(hostname string, password string, physicalnetworkid string, username string) *AddBrocadeVcsDeviceParams
	DeleteBrocadeVcsDevice(P *DeleteBrocadeVcsDeviceParams) (*DeleteBrocadeVcsDeviceResponse, error)
	NewDeleteBrocadeVcsDeviceParams(vcsdeviceid string) *DeleteBrocadeVcsDeviceParams
	ListBrocadeVcsDeviceNetworks(P *ListBrocadeVcsDeviceNetworksParams) (*ListBrocadeVcsDeviceNetworksResponse, error)
	NewListBrocadeVcsDeviceNetworksParams(vcsdeviceid string) *ListBrocadeVcsDeviceNetworksParams
	GetBrocadeVcsDeviceNetworkID(keyword string, vcsdeviceid string, opts ...OptionFunc) (string, int, error)
	ListBrocadeVcsDevices(P *ListBrocadeVcsDevicesParams) (*ListBrocadeVcsDevicesResponse, error)
	NewListBrocadeVcsDevicesParams() *ListBrocadeVcsDevicesParams
}

type AddBrocadeVcsDeviceParams struct {
	P map[string]interface{}
}

func (P *AddBrocadeVcsDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *AddBrocadeVcsDeviceParams) SetHostname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostname"] = v
}

func (P *AddBrocadeVcsDeviceParams) GetHostname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostname"].(string)
	return value, ok
}

func (P *AddBrocadeVcsDeviceParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddBrocadeVcsDeviceParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddBrocadeVcsDeviceParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddBrocadeVcsDeviceParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddBrocadeVcsDeviceParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddBrocadeVcsDeviceParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddBrocadeVcsDeviceParams instance,
// as then you are sure you have configured all required params
func (s *BrocadeVCSService) NewAddBrocadeVcsDeviceParams(hostname string, password string, physicalnetworkid string, username string) *AddBrocadeVcsDeviceParams {
	P := &AddBrocadeVcsDeviceParams{}
	P.P = make(map[string]interface{})
	P.P["hostname"] = hostname
	P.P["password"] = password
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["username"] = username
	return P
}

// Adds a Brocade VCS Switch
func (s *BrocadeVCSService) AddBrocadeVcsDevice(p *AddBrocadeVcsDeviceParams) (*AddBrocadeVcsDeviceResponse, error) {
	resp, err := s.cs.newRequest("addBrocadeVcsDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBrocadeVcsDeviceResponse
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

type AddBrocadeVcsDeviceResponse struct {
	Brocadedevicename string `json:"brocadedevicename"`
	Hostname          string `json:"hostname"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Vcsdeviceid       string `json:"vcsdeviceid"`
}

type DeleteBrocadeVcsDeviceParams struct {
	P map[string]interface{}
}

func (P *DeleteBrocadeVcsDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["vcsdeviceid"]; found {
		u.Set("vcsdeviceid", v.(string))
	}
	return u
}

func (P *DeleteBrocadeVcsDeviceParams) SetVcsdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vcsdeviceid"] = v
}

func (P *DeleteBrocadeVcsDeviceParams) GetVcsdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vcsdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBrocadeVcsDeviceParams instance,
// as then you are sure you have configured all required params
func (s *BrocadeVCSService) NewDeleteBrocadeVcsDeviceParams(vcsdeviceid string) *DeleteBrocadeVcsDeviceParams {
	P := &DeleteBrocadeVcsDeviceParams{}
	P.P = make(map[string]interface{})
	P.P["vcsdeviceid"] = vcsdeviceid
	return P
}

//  delete a Brocade VCS Switch
func (s *BrocadeVCSService) DeleteBrocadeVcsDevice(p *DeleteBrocadeVcsDeviceParams) (*DeleteBrocadeVcsDeviceResponse, error) {
	resp, err := s.cs.newRequest("deleteBrocadeVcsDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBrocadeVcsDeviceResponse
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

type DeleteBrocadeVcsDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListBrocadeVcsDeviceNetworksParams struct {
	P map[string]interface{}
}

func (P *ListBrocadeVcsDeviceNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
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
	if v, found := P.P["vcsdeviceid"]; found {
		u.Set("vcsdeviceid", v.(string))
	}
	return u
}

func (P *ListBrocadeVcsDeviceNetworksParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListBrocadeVcsDeviceNetworksParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListBrocadeVcsDeviceNetworksParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListBrocadeVcsDeviceNetworksParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListBrocadeVcsDeviceNetworksParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListBrocadeVcsDeviceNetworksParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListBrocadeVcsDeviceNetworksParams) SetVcsdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vcsdeviceid"] = v
}

func (P *ListBrocadeVcsDeviceNetworksParams) GetVcsdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vcsdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBrocadeVcsDeviceNetworksParams instance,
// as then you are sure you have configured all required params
func (s *BrocadeVCSService) NewListBrocadeVcsDeviceNetworksParams(vcsdeviceid string) *ListBrocadeVcsDeviceNetworksParams {
	P := &ListBrocadeVcsDeviceNetworksParams{}
	P.P = make(map[string]interface{})
	P.P["vcsdeviceid"] = vcsdeviceid
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *BrocadeVCSService) GetBrocadeVcsDeviceNetworkID(keyword string, vcsdeviceid string, opts ...OptionFunc) (string, int, error) {
	P := &ListBrocadeVcsDeviceNetworksParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword
	P.P["vcsdeviceid"] = vcsdeviceid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListBrocadeVcsDeviceNetworks(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.BrocadeVcsDeviceNetworks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.BrocadeVcsDeviceNetworks {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// lists network that are using a brocade vcs switch
func (s *BrocadeVCSService) ListBrocadeVcsDeviceNetworks(p *ListBrocadeVcsDeviceNetworksParams) (*ListBrocadeVcsDeviceNetworksResponse, error) {
	resp, err := s.cs.newRequest("listBrocadeVcsDeviceNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBrocadeVcsDeviceNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBrocadeVcsDeviceNetworksResponse struct {
	Count                    int                        `json:"count"`
	BrocadeVcsDeviceNetworks []*BrocadeVcsDeviceNetwork `json:"brocadevcsdevicenetwork"`
}

type BrocadeVcsDeviceNetwork struct {
	Account                     string                           `json:"account"`
	Aclid                       string                           `json:"aclid"`
	Aclname                     string                           `json:"aclname"`
	Acltype                     string                           `json:"acltype"`
	Broadcastdomaintype         string                           `json:"broadcastdomaintype"`
	Broadcasturi                string                           `json:"broadcasturi"`
	Canusefordeploy             bool                             `json:"canusefordeploy"`
	Cidr                        string                           `json:"cidr"`
	Created                     string                           `json:"created"`
	Details                     map[string]string                `json:"details"`
	Displaynetwork              bool                             `json:"displaynetwork"`
	Displaytext                 string                           `json:"displaytext"`
	Dns1                        string                           `json:"dns1"`
	Dns2                        string                           `json:"dns2"`
	Domain                      string                           `json:"domain"`
	Domainid                    string                           `json:"domainid"`
	Externalid                  string                           `json:"externalid"`
	Gateway                     string                           `json:"gateway"`
	Hasannotations              bool                             `json:"hasannotations"`
	Icon                        string                           `json:"icon"`
	Id                          string                           `json:"id"`
	Ip6cidr                     string                           `json:"ip6cidr"`
	Ip6gateway                  string                           `json:"ip6gateway"`
	Isdefault                   bool                             `json:"isdefault"`
	Ispersistent                bool                             `json:"ispersistent"`
	Issystem                    bool                             `json:"issystem"`
	JobID                       string                           `json:"jobid"`
	Jobstatus                   int                              `json:"jobstatus"`
	Name                        string                           `json:"name"`
	Netmask                     string                           `json:"netmask"`
	Networkcidr                 string                           `json:"networkcidr"`
	Networkdomain               string                           `json:"networkdomain"`
	Networkofferingavailability string                           `json:"networkofferingavailability"`
	Networkofferingconservemode bool                             `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                           `json:"networkofferingdisplaytext"`
	Networkofferingid           string                           `json:"networkofferingid"`
	Networkofferingname         string                           `json:"networkofferingname"`
	Physicalnetworkid           string                           `json:"physicalnetworkid"`
	Project                     string                           `json:"project"`
	Projectid                   string                           `json:"projectid"`
	Receivedbytes               int64                            `json:"receivedbytes"`
	Redundantrouter             bool                             `json:"redundantrouter"`
	Related                     string                           `json:"related"`
	Reservediprange             string                           `json:"reservediprange"`
	Restartrequired             bool                             `json:"restartrequired"`
	Sentbytes                   int64                            `json:"sentbytes"`
	Service                     []BrocadeVcsDeviceNetworkService `json:"service"`
	Specifyipranges             bool                             `json:"specifyipranges"`
	State                       string                           `json:"state"`
	Strechedl2subnet            bool                             `json:"strechedl2subnet"`
	Subdomainaccess             bool                             `json:"subdomainaccess"`
	Tags                        []Tags                           `json:"tags"`
	Traffictype                 string                           `json:"traffictype"`
	Type                        string                           `json:"type"`
	Vlan                        string                           `json:"vlan"`
	Vpcid                       string                           `json:"vpcid"`
	Vpcname                     string                           `json:"vpcname"`
	Zoneid                      string                           `json:"zoneid"`
	Zonename                    string                           `json:"zonename"`
	Zonesnetworkspans           []interface{}                    `json:"zonesnetworkspans"`
}

type BrocadeVcsDeviceNetworkService struct {
	Capability []BrocadeVcsDeviceNetworkServiceCapability `json:"capability"`
	Name       string                                     `json:"name"`
	Provider   []BrocadeVcsDeviceNetworkServiceProvider   `json:"provider"`
}

type BrocadeVcsDeviceNetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type BrocadeVcsDeviceNetworkServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListBrocadeVcsDevicesParams struct {
	P map[string]interface{}
}

func (P *ListBrocadeVcsDevicesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
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
	if v, found := P.P["vcsdeviceid"]; found {
		u.Set("vcsdeviceid", v.(string))
	}
	return u
}

func (P *ListBrocadeVcsDevicesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListBrocadeVcsDevicesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListBrocadeVcsDevicesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListBrocadeVcsDevicesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListBrocadeVcsDevicesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListBrocadeVcsDevicesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListBrocadeVcsDevicesParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListBrocadeVcsDevicesParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *ListBrocadeVcsDevicesParams) SetVcsdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vcsdeviceid"] = v
}

func (P *ListBrocadeVcsDevicesParams) GetVcsdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vcsdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBrocadeVcsDevicesParams instance,
// as then you are sure you have configured all required params
func (s *BrocadeVCSService) NewListBrocadeVcsDevicesParams() *ListBrocadeVcsDevicesParams {
	P := &ListBrocadeVcsDevicesParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists Brocade VCS Switches
func (s *BrocadeVCSService) ListBrocadeVcsDevices(p *ListBrocadeVcsDevicesParams) (*ListBrocadeVcsDevicesResponse, error) {
	resp, err := s.cs.newRequest("listBrocadeVcsDevices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBrocadeVcsDevicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBrocadeVcsDevicesResponse struct {
	Count             int                 `json:"count"`
	BrocadeVcsDevices []*BrocadeVcsDevice `json:"brocadevcsdevice"`
}

type BrocadeVcsDevice struct {
	Brocadedevicename string `json:"brocadedevicename"`
	Hostname          string `json:"hostname"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Vcsdeviceid       string `json:"vcsdeviceid"`
}
