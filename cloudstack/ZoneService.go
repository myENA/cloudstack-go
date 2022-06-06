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

type ZoneServiceIface interface {
	CreateZone(P *CreateZoneParams) (*CreateZoneResponse, error)
	NewCreateZoneParams(dns1 string, internaldns1 string, name string, networktype string) *CreateZoneParams
	DedicateZone(P *DedicateZoneParams) (*DedicateZoneResponse, error)
	NewDedicateZoneParams(domainid string, zoneid string) *DedicateZoneParams
	DeleteZone(P *DeleteZoneParams) (*DeleteZoneResponse, error)
	NewDeleteZoneParams(id string) *DeleteZoneParams
	DisableOutOfBandManagementForZone(P *DisableOutOfBandManagementForZoneParams) (*DisableOutOfBandManagementForZoneResponse, error)
	NewDisableOutOfBandManagementForZoneParams(zoneid string) *DisableOutOfBandManagementForZoneParams
	EnableOutOfBandManagementForZone(P *EnableOutOfBandManagementForZoneParams) (*EnableOutOfBandManagementForZoneResponse, error)
	NewEnableOutOfBandManagementForZoneParams(zoneid string) *EnableOutOfBandManagementForZoneParams
	DisableHAForZone(P *DisableHAForZoneParams) (*DisableHAForZoneResponse, error)
	NewDisableHAForZoneParams(zoneid string) *DisableHAForZoneParams
	EnableHAForZone(P *EnableHAForZoneParams) (*EnableHAForZoneResponse, error)
	NewEnableHAForZoneParams(zoneid string) *EnableHAForZoneParams
	ListDedicatedZones(P *ListDedicatedZonesParams) (*ListDedicatedZonesResponse, error)
	NewListDedicatedZonesParams() *ListDedicatedZonesParams
	ListZones(P *ListZonesParams) (*ListZonesResponse, error)
	NewListZonesParams() *ListZonesParams
	GetZoneID(name string, opts ...OptionFunc) (string, int, error)
	GetZoneByName(name string, opts ...OptionFunc) (*Zone, int, error)
	GetZoneByID(id string, opts ...OptionFunc) (*Zone, int, error)
	ReleaseDedicatedZone(P *ReleaseDedicatedZoneParams) (*ReleaseDedicatedZoneResponse, error)
	NewReleaseDedicatedZoneParams(zoneid string) *ReleaseDedicatedZoneParams
	UpdateZone(P *UpdateZoneParams) (*UpdateZoneResponse, error)
	NewUpdateZoneParams(id string) *UpdateZoneParams
}

type CreateZoneParams struct {
	P map[string]interface{}
}

func (P *CreateZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["dns1"]; found {
		u.Set("dns1", v.(string))
	}
	if v, found := P.P["dns2"]; found {
		u.Set("dns2", v.(string))
	}
	if v, found := P.P["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["guestcidraddress"]; found {
		u.Set("guestcidraddress", v.(string))
	}
	if v, found := P.P["internaldns1"]; found {
		u.Set("internaldns1", v.(string))
	}
	if v, found := P.P["internaldns2"]; found {
		u.Set("internaldns2", v.(string))
	}
	if v, found := P.P["ip6dns1"]; found {
		u.Set("ip6dns1", v.(string))
	}
	if v, found := P.P["ip6dns2"]; found {
		u.Set("ip6dns2", v.(string))
	}
	if v, found := P.P["localstorageenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("localstorageenabled", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networktype"]; found {
		u.Set("networktype", v.(string))
	}
	if v, found := P.P["securitygroupenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("securitygroupenabled", vv)
	}
	return u
}

func (P *CreateZoneParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *CreateZoneParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetDns1(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dns1"] = v
}

func (P *CreateZoneParams) GetDns1() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dns1"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetDns2(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dns2"] = v
}

func (P *CreateZoneParams) GetDns2() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dns2"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetDomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domain"] = v
}

func (P *CreateZoneParams) GetDomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domain"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateZoneParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetGuestcidraddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["guestcidraddress"] = v
}

func (P *CreateZoneParams) GetGuestcidraddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["guestcidraddress"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetInternaldns1(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["internaldns1"] = v
}

func (P *CreateZoneParams) GetInternaldns1() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["internaldns1"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetInternaldns2(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["internaldns2"] = v
}

func (P *CreateZoneParams) GetInternaldns2() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["internaldns2"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetIp6dns1(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ip6dns1"] = v
}

func (P *CreateZoneParams) GetIp6dns1() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ip6dns1"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetIp6dns2(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ip6dns2"] = v
}

func (P *CreateZoneParams) GetIp6dns2() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ip6dns2"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetLocalstorageenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["localstorageenabled"] = v
}

func (P *CreateZoneParams) GetLocalstorageenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["localstorageenabled"].(bool)
	return value, ok
}

func (P *CreateZoneParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateZoneParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetNetworktype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networktype"] = v
}

func (P *CreateZoneParams) GetNetworktype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networktype"].(string)
	return value, ok
}

func (P *CreateZoneParams) SetSecuritygroupenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupenabled"] = v
}

func (P *CreateZoneParams) GetSecuritygroupenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupenabled"].(bool)
	return value, ok
}

// You should always use this function to get a new CreateZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewCreateZoneParams(dns1 string, internaldns1 string, name string, networktype string) *CreateZoneParams {
	P := &CreateZoneParams{}
	P.P = make(map[string]interface{})
	P.P["dns1"] = dns1
	P.P["internaldns1"] = internaldns1
	P.P["name"] = name
	P.P["networktype"] = networktype
	return P
}

// Creates a Zone.
func (s *ZoneService) CreateZone(p *CreateZoneParams) (*CreateZoneResponse, error) {
	resp, err := s.cs.newRequest("createZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateZoneResponse struct {
	Allocationstate       string                       `json:"allocationstate"`
	Capacity              []CreateZoneResponseCapacity `json:"capacity"`
	Description           string                       `json:"description"`
	Dhcpprovider          string                       `json:"dhcpprovider"`
	Displaytext           string                       `json:"displaytext"`
	Dns1                  string                       `json:"dns1"`
	Dns2                  string                       `json:"dns2"`
	Domain                string                       `json:"domain"`
	Domainid              string                       `json:"domainid"`
	Domainname            string                       `json:"domainname"`
	Guestcidraddress      string                       `json:"guestcidraddress"`
	Hasannotations        bool                         `json:"hasannotations"`
	Icon                  string                       `json:"icon"`
	Id                    string                       `json:"id"`
	Internaldns1          string                       `json:"internaldns1"`
	Internaldns2          string                       `json:"internaldns2"`
	Ip6dns1               string                       `json:"ip6dns1"`
	Ip6dns2               string                       `json:"ip6dns2"`
	JobID                 string                       `json:"jobid"`
	Jobstatus             int                          `json:"jobstatus"`
	Localstorageenabled   bool                         `json:"localstorageenabled"`
	Name                  string                       `json:"name"`
	Networktype           string                       `json:"networktype"`
	Resourcedetails       map[string]string            `json:"resourcedetails"`
	Securitygroupsenabled bool                         `json:"securitygroupsenabled"`
	Tags                  []Tags                       `json:"tags"`
	Zonetoken             string                       `json:"zonetoken"`
}

type CreateZoneResponseCapacity struct {
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

type DedicateZoneParams struct {
	P map[string]interface{}
}

func (P *DedicateZoneParams) toURLValues() url.Values {
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
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *DedicateZoneParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DedicateZoneParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DedicateZoneParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DedicateZoneParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *DedicateZoneParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *DedicateZoneParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicateZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDedicateZoneParams(domainid string, zoneid string) *DedicateZoneParams {
	P := &DedicateZoneParams{}
	P.P = make(map[string]interface{})
	P.P["domainid"] = domainid
	P.P["zoneid"] = zoneid
	return P
}

// Dedicates a zones.
func (s *ZoneService) DedicateZone(p *DedicateZoneParams) (*DedicateZoneResponse, error) {
	resp, err := s.cs.newRequest("dedicateZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateZoneResponse
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

type DedicateZoneResponse struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Zoneid          string `json:"zoneid"`
	Zonename        string `json:"zonename"`
}

type DeleteZoneParams struct {
	P map[string]interface{}
}

func (P *DeleteZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteZoneParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteZoneParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDeleteZoneParams(id string) *DeleteZoneParams {
	P := &DeleteZoneParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a Zone.
func (s *ZoneService) DeleteZone(p *DeleteZoneParams) (*DeleteZoneResponse, error) {
	resp, err := s.cs.newRequest("deleteZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteZoneResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteZoneResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteZoneResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DisableOutOfBandManagementForZoneParams struct {
	P map[string]interface{}
}

func (P *DisableOutOfBandManagementForZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *DisableOutOfBandManagementForZoneParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *DisableOutOfBandManagementForZoneParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableOutOfBandManagementForZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDisableOutOfBandManagementForZoneParams(zoneid string) *DisableOutOfBandManagementForZoneParams {
	P := &DisableOutOfBandManagementForZoneParams{}
	P.P = make(map[string]interface{})
	P.P["zoneid"] = zoneid
	return P
}

// Disables out-of-band management for a zone
func (s *ZoneService) DisableOutOfBandManagementForZone(p *DisableOutOfBandManagementForZoneParams) (*DisableOutOfBandManagementForZoneResponse, error) {
	resp, err := s.cs.newRequest("disableOutOfBandManagementForZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableOutOfBandManagementForZoneResponse
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

type DisableOutOfBandManagementForZoneResponse struct {
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}

type EnableOutOfBandManagementForZoneParams struct {
	P map[string]interface{}
}

func (P *EnableOutOfBandManagementForZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *EnableOutOfBandManagementForZoneParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *EnableOutOfBandManagementForZoneParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableOutOfBandManagementForZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewEnableOutOfBandManagementForZoneParams(zoneid string) *EnableOutOfBandManagementForZoneParams {
	P := &EnableOutOfBandManagementForZoneParams{}
	P.P = make(map[string]interface{})
	P.P["zoneid"] = zoneid
	return P
}

// Enables out-of-band management for a zone
func (s *ZoneService) EnableOutOfBandManagementForZone(p *EnableOutOfBandManagementForZoneParams) (*EnableOutOfBandManagementForZoneResponse, error) {
	resp, err := s.cs.newRequest("enableOutOfBandManagementForZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableOutOfBandManagementForZoneResponse
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

type EnableOutOfBandManagementForZoneResponse struct {
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}

type DisableHAForZoneParams struct {
	P map[string]interface{}
}

func (P *DisableHAForZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *DisableHAForZoneParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *DisableHAForZoneParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableHAForZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewDisableHAForZoneParams(zoneid string) *DisableHAForZoneParams {
	P := &DisableHAForZoneParams{}
	P.P = make(map[string]interface{})
	P.P["zoneid"] = zoneid
	return P
}

// Disables HA for a zone
func (s *ZoneService) DisableHAForZone(p *DisableHAForZoneParams) (*DisableHAForZoneResponse, error) {
	resp, err := s.cs.newRequest("disableHAForZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableHAForZoneResponse
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

type DisableHAForZoneResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type EnableHAForZoneParams struct {
	P map[string]interface{}
}

func (P *EnableHAForZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *EnableHAForZoneParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *EnableHAForZoneParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableHAForZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewEnableHAForZoneParams(zoneid string) *EnableHAForZoneParams {
	P := &EnableHAForZoneParams{}
	P.P = make(map[string]interface{})
	P.P["zoneid"] = zoneid
	return P
}

// Enables HA for a zone
func (s *ZoneService) EnableHAForZone(p *EnableHAForZoneParams) (*EnableHAForZoneResponse, error) {
	resp, err := s.cs.newRequest("enableHAForZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableHAForZoneResponse
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

type EnableHAForZoneResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListDedicatedZonesParams struct {
	P map[string]interface{}
}

func (P *ListDedicatedZonesParams) toURLValues() url.Values {
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
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListDedicatedZonesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListDedicatedZonesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListDedicatedZonesParams) SetAffinitygroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["affinitygroupid"] = v
}

func (P *ListDedicatedZonesParams) GetAffinitygroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["affinitygroupid"].(string)
	return value, ok
}

func (P *ListDedicatedZonesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListDedicatedZonesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListDedicatedZonesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListDedicatedZonesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListDedicatedZonesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListDedicatedZonesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListDedicatedZonesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListDedicatedZonesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListDedicatedZonesParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListDedicatedZonesParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListDedicatedZonesParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewListDedicatedZonesParams() *ListDedicatedZonesParams {
	P := &ListDedicatedZonesParams{}
	P.P = make(map[string]interface{})
	return P
}

// List dedicated zones.
func (s *ZoneService) ListDedicatedZones(p *ListDedicatedZonesParams) (*ListDedicatedZonesResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedZones", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedZonesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedZonesResponse struct {
	Count          int              `json:"count"`
	DedicatedZones []*DedicatedZone `json:"dedicatedzone"`
}

type DedicatedZone struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Zoneid          string `json:"zoneid"`
	Zonename        string `json:"zonename"`
}

type ListZonesParams struct {
	P map[string]interface{}
}

func (P *ListZonesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["available"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("available", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
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
	if v, found := P.P["networktype"]; found {
		u.Set("networktype", v.(string))
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
	if v, found := P.P["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (P *ListZonesParams) SetAvailable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["available"] = v
}

func (P *ListZonesParams) GetAvailable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["available"].(bool)
	return value, ok
}

func (P *ListZonesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListZonesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListZonesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListZonesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListZonesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListZonesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListZonesParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListZonesParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListZonesParams) SetNetworktype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networktype"] = v
}

func (P *ListZonesParams) GetNetworktype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networktype"].(string)
	return value, ok
}

func (P *ListZonesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListZonesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListZonesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListZonesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListZonesParams) SetShowcapacities(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showcapacities"] = v
}

func (P *ListZonesParams) GetShowcapacities() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showcapacities"].(bool)
	return value, ok
}

func (P *ListZonesParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListZonesParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

func (P *ListZonesParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListZonesParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListZonesParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewListZonesParams() *ListZonesParams {
	P := &ListZonesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZoneID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListZonesParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListZones(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Zones[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Zones {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZoneByName(name string, opts ...OptionFunc) (*Zone, int, error) {
	id, count, err := s.GetZoneID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetZoneByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ZoneService) GetZoneByID(id string, opts ...OptionFunc) (*Zone, int, error) {
	P := &ListZonesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListZones(P)
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
		return l.Zones[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Zone UUID: %s!", id)
}

// Lists zones
func (s *ZoneService) ListZones(p *ListZonesParams) (*ListZonesResponse, error) {
	resp, err := s.cs.newRequest("listZones", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListZonesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListZonesResponse struct {
	Count int     `json:"count"`
	Zones []*Zone `json:"zone"`
}

type Zone struct {
	Allocationstate       string            `json:"allocationstate"`
	Capacity              []ZoneCapacity    `json:"capacity"`
	Description           string            `json:"description"`
	Dhcpprovider          string            `json:"dhcpprovider"`
	Displaytext           string            `json:"displaytext"`
	Dns1                  string            `json:"dns1"`
	Dns2                  string            `json:"dns2"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Domainname            string            `json:"domainname"`
	Guestcidraddress      string            `json:"guestcidraddress"`
	Hasannotations        bool              `json:"hasannotations"`
	Icon                  string            `json:"icon"`
	Id                    string            `json:"id"`
	Internaldns1          string            `json:"internaldns1"`
	Internaldns2          string            `json:"internaldns2"`
	Ip6dns1               string            `json:"ip6dns1"`
	Ip6dns2               string            `json:"ip6dns2"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Localstorageenabled   bool              `json:"localstorageenabled"`
	Name                  string            `json:"name"`
	Networktype           string            `json:"networktype"`
	Resourcedetails       map[string]string `json:"resourcedetails"`
	Securitygroupsenabled bool              `json:"securitygroupsenabled"`
	Tags                  []Tags            `json:"tags"`
	Zonetoken             string            `json:"zonetoken"`
}

type ZoneCapacity struct {
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

type ReleaseDedicatedZoneParams struct {
	P map[string]interface{}
}

func (P *ReleaseDedicatedZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ReleaseDedicatedZoneParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ReleaseDedicatedZoneParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseDedicatedZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewReleaseDedicatedZoneParams(zoneid string) *ReleaseDedicatedZoneParams {
	P := &ReleaseDedicatedZoneParams{}
	P.P = make(map[string]interface{})
	P.P["zoneid"] = zoneid
	return P
}

// Release dedication of zone
func (s *ZoneService) ReleaseDedicatedZone(p *ReleaseDedicatedZoneParams) (*ReleaseDedicatedZoneResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedZoneResponse
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

type ReleaseDedicatedZoneResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateZoneParams struct {
	P map[string]interface{}
}

func (P *UpdateZoneParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), m[k])
		}
	}
	if v, found := P.P["dhcpprovider"]; found {
		u.Set("dhcpprovider", v.(string))
	}
	if v, found := P.P["dns1"]; found {
		u.Set("dns1", v.(string))
	}
	if v, found := P.P["dns2"]; found {
		u.Set("dns2", v.(string))
	}
	if v, found := P.P["dnssearchorder"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("dnssearchorder", vv)
	}
	if v, found := P.P["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := P.P["guestcidraddress"]; found {
		u.Set("guestcidraddress", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["internaldns1"]; found {
		u.Set("internaldns1", v.(string))
	}
	if v, found := P.P["internaldns2"]; found {
		u.Set("internaldns2", v.(string))
	}
	if v, found := P.P["ip6dns1"]; found {
		u.Set("ip6dns1", v.(string))
	}
	if v, found := P.P["ip6dns2"]; found {
		u.Set("ip6dns2", v.(string))
	}
	if v, found := P.P["ispublic"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispublic", vv)
	}
	if v, found := P.P["localstorageenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("localstorageenabled", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	return u
}

func (P *UpdateZoneParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *UpdateZoneParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *UpdateZoneParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *UpdateZoneParams) SetDhcpprovider(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dhcpprovider"] = v
}

func (P *UpdateZoneParams) GetDhcpprovider() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dhcpprovider"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetDns1(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dns1"] = v
}

func (P *UpdateZoneParams) GetDns1() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dns1"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetDns2(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dns2"] = v
}

func (P *UpdateZoneParams) GetDns2() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dns2"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetDnssearchorder(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dnssearchorder"] = v
}

func (P *UpdateZoneParams) GetDnssearchorder() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dnssearchorder"].([]string)
	return value, ok
}

func (P *UpdateZoneParams) SetDomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domain"] = v
}

func (P *UpdateZoneParams) GetDomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domain"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetGuestcidraddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["guestcidraddress"] = v
}

func (P *UpdateZoneParams) GetGuestcidraddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["guestcidraddress"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateZoneParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetInternaldns1(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["internaldns1"] = v
}

func (P *UpdateZoneParams) GetInternaldns1() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["internaldns1"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetInternaldns2(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["internaldns2"] = v
}

func (P *UpdateZoneParams) GetInternaldns2() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["internaldns2"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetIp6dns1(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ip6dns1"] = v
}

func (P *UpdateZoneParams) GetIp6dns1() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ip6dns1"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetIp6dns2(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ip6dns2"] = v
}

func (P *UpdateZoneParams) GetIp6dns2() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ip6dns2"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetIspublic(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ispublic"] = v
}

func (P *UpdateZoneParams) GetIspublic() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ispublic"].(bool)
	return value, ok
}

func (P *UpdateZoneParams) SetLocalstorageenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["localstorageenabled"] = v
}

func (P *UpdateZoneParams) GetLocalstorageenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["localstorageenabled"].(bool)
	return value, ok
}

func (P *UpdateZoneParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateZoneParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateZoneParams) SetSortkey(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sortkey"] = v
}

func (P *UpdateZoneParams) GetSortkey() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sortkey"].(int)
	return value, ok
}

// You should always use this function to get a new UpdateZoneParams instance,
// as then you are sure you have configured all required params
func (s *ZoneService) NewUpdateZoneParams(id string) *UpdateZoneParams {
	P := &UpdateZoneParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a Zone.
func (s *ZoneService) UpdateZone(p *UpdateZoneParams) (*UpdateZoneResponse, error) {
	resp, err := s.cs.newRequest("updateZone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateZoneResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateZoneResponse struct {
	Allocationstate       string                       `json:"allocationstate"`
	Capacity              []UpdateZoneResponseCapacity `json:"capacity"`
	Description           string                       `json:"description"`
	Dhcpprovider          string                       `json:"dhcpprovider"`
	Displaytext           string                       `json:"displaytext"`
	Dns1                  string                       `json:"dns1"`
	Dns2                  string                       `json:"dns2"`
	Domain                string                       `json:"domain"`
	Domainid              string                       `json:"domainid"`
	Domainname            string                       `json:"domainname"`
	Guestcidraddress      string                       `json:"guestcidraddress"`
	Hasannotations        bool                         `json:"hasannotations"`
	Icon                  string                       `json:"icon"`
	Id                    string                       `json:"id"`
	Internaldns1          string                       `json:"internaldns1"`
	Internaldns2          string                       `json:"internaldns2"`
	Ip6dns1               string                       `json:"ip6dns1"`
	Ip6dns2               string                       `json:"ip6dns2"`
	JobID                 string                       `json:"jobid"`
	Jobstatus             int                          `json:"jobstatus"`
	Localstorageenabled   bool                         `json:"localstorageenabled"`
	Name                  string                       `json:"name"`
	Networktype           string                       `json:"networktype"`
	Resourcedetails       map[string]string            `json:"resourcedetails"`
	Securitygroupsenabled bool                         `json:"securitygroupsenabled"`
	Tags                  []Tags                       `json:"tags"`
	Zonetoken             string                       `json:"zonetoken"`
}

type UpdateZoneResponseCapacity struct {
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
