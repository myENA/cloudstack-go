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
	"net/url"
	"strconv"
)

type BigSwitchBCFServiceIface interface {
	AddBigSwitchBcfDevice(p *AddBigSwitchBcfDeviceParams) (*AddBigSwitchBcfDeviceResponse, error)
	NewAddBigSwitchBcfDeviceParams(hostname string, nat bool, password string, physicalnetworkid string, username string) *AddBigSwitchBcfDeviceParams
	DeleteBigSwitchBcfDevice(p *DeleteBigSwitchBcfDeviceParams) (*DeleteBigSwitchBcfDeviceResponse, error)
	NewDeleteBigSwitchBcfDeviceParams(bcfdeviceid string) *DeleteBigSwitchBcfDeviceParams
	ListBigSwitchBcfDevices(p *ListBigSwitchBcfDevicesParams) (*ListBigSwitchBcfDevicesResponse, error)
	NewListBigSwitchBcfDevicesParams() *ListBigSwitchBcfDevicesParams
}

type AddBigSwitchBcfDeviceParams struct {
	P map[string]interface{}
}

func (P *AddBigSwitchBcfDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := P.P["nat"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("nat", vv)
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

func (P *AddBigSwitchBcfDeviceParams) SetHostname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostname"] = v
}

func (P *AddBigSwitchBcfDeviceParams) GetHostname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostname"].(string)
	return value, ok
}

func (P *AddBigSwitchBcfDeviceParams) SetNat(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nat"] = v
}

func (P *AddBigSwitchBcfDeviceParams) GetNat() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nat"].(bool)
	return value, ok
}

func (P *AddBigSwitchBcfDeviceParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddBigSwitchBcfDeviceParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddBigSwitchBcfDeviceParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddBigSwitchBcfDeviceParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddBigSwitchBcfDeviceParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddBigSwitchBcfDeviceParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddBigSwitchBcfDeviceParams instance,
// as then you are sure you have configured all required params
func (s *BigSwitchBCFService) NewAddBigSwitchBcfDeviceParams(hostname string, nat bool, password string, physicalnetworkid string, username string) *AddBigSwitchBcfDeviceParams {
	P := &AddBigSwitchBcfDeviceParams{}
	P.P = make(map[string]interface{})
	P.P["hostname"] = hostname
	P.P["nat"] = nat
	P.P["password"] = password
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["username"] = username
	return P
}

// Adds a BigSwitch BCF Controller device
func (s *BigSwitchBCFService) AddBigSwitchBcfDevice(p *AddBigSwitchBcfDeviceParams) (*AddBigSwitchBcfDeviceResponse, error) {
	resp, err := s.cs.newRequest("addBigSwitchBcfDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBigSwitchBcfDeviceResponse
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

type AddBigSwitchBcfDeviceResponse struct {
	Bcfdeviceid         string `json:"bcfdeviceid"`
	Bigswitchdevicename string `json:"bigswitchdevicename"`
	Hostname            string `json:"hostname"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Nat                 bool   `json:"nat"`
	Password            string `json:"password"`
	Physicalnetworkid   string `json:"physicalnetworkid"`
	Provider            string `json:"provider"`
	Username            string `json:"username"`
}

type DeleteBigSwitchBcfDeviceParams struct {
	P map[string]interface{}
}

func (P *DeleteBigSwitchBcfDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["bcfdeviceid"]; found {
		u.Set("bcfdeviceid", v.(string))
	}
	return u
}

func (P *DeleteBigSwitchBcfDeviceParams) SetBcfdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bcfdeviceid"] = v
}

func (P *DeleteBigSwitchBcfDeviceParams) GetBcfdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bcfdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBigSwitchBcfDeviceParams instance,
// as then you are sure you have configured all required params
func (s *BigSwitchBCFService) NewDeleteBigSwitchBcfDeviceParams(bcfdeviceid string) *DeleteBigSwitchBcfDeviceParams {
	P := &DeleteBigSwitchBcfDeviceParams{}
	P.P = make(map[string]interface{})
	P.P["bcfdeviceid"] = bcfdeviceid
	return P
}

// delete a BigSwitch BCF Controller device
func (s *BigSwitchBCFService) DeleteBigSwitchBcfDevice(p *DeleteBigSwitchBcfDeviceParams) (*DeleteBigSwitchBcfDeviceResponse, error) {
	resp, err := s.cs.newRequest("deleteBigSwitchBcfDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBigSwitchBcfDeviceResponse
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

type DeleteBigSwitchBcfDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListBigSwitchBcfDevicesParams struct {
	P map[string]interface{}
}

func (P *ListBigSwitchBcfDevicesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["bcfdeviceid"]; found {
		u.Set("bcfdeviceid", v.(string))
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
	return u
}

func (P *ListBigSwitchBcfDevicesParams) SetBcfdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bcfdeviceid"] = v
}

func (P *ListBigSwitchBcfDevicesParams) GetBcfdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bcfdeviceid"].(string)
	return value, ok
}

func (P *ListBigSwitchBcfDevicesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListBigSwitchBcfDevicesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListBigSwitchBcfDevicesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListBigSwitchBcfDevicesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListBigSwitchBcfDevicesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListBigSwitchBcfDevicesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListBigSwitchBcfDevicesParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListBigSwitchBcfDevicesParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBigSwitchBcfDevicesParams instance,
// as then you are sure you have configured all required params
func (s *BigSwitchBCFService) NewListBigSwitchBcfDevicesParams() *ListBigSwitchBcfDevicesParams {
	P := &ListBigSwitchBcfDevicesParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists BigSwitch BCF Controller devices
func (s *BigSwitchBCFService) ListBigSwitchBcfDevices(p *ListBigSwitchBcfDevicesParams) (*ListBigSwitchBcfDevicesResponse, error) {
	resp, err := s.cs.newRequest("listBigSwitchBcfDevices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBigSwitchBcfDevicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBigSwitchBcfDevicesResponse struct {
	Count               int                   `json:"count"`
	BigSwitchBcfDevices []*BigSwitchBcfDevice `json:"bigswitchbcfdevice"`
}

type BigSwitchBcfDevice struct {
	Bcfdeviceid         string `json:"bcfdeviceid"`
	Bigswitchdevicename string `json:"bigswitchdevicename"`
	Hostname            string `json:"hostname"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Nat                 bool   `json:"nat"`
	Password            string `json:"password"`
	Physicalnetworkid   string `json:"physicalnetworkid"`
	Provider            string `json:"provider"`
	Username            string `json:"username"`
}
