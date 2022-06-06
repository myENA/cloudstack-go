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

type NiciraNVPServiceIface interface {
	AddNiciraNvpDevice(P *AddNiciraNvpDeviceParams) (*AddNiciraNvpDeviceResponse, error)
	NewAddNiciraNvpDeviceParams(hostname string, password string, physicalnetworkid string, transportzoneuuid string, username string) *AddNiciraNvpDeviceParams
	DeleteNiciraNvpDevice(P *DeleteNiciraNvpDeviceParams) (*DeleteNiciraNvpDeviceResponse, error)
	NewDeleteNiciraNvpDeviceParams(nvpdeviceid string) *DeleteNiciraNvpDeviceParams
	ListNiciraNvpDevices(P *ListNiciraNvpDevicesParams) (*ListNiciraNvpDevicesResponse, error)
	NewListNiciraNvpDevicesParams() *ListNiciraNvpDevicesParams
}

type AddNiciraNvpDeviceParams struct {
	P map[string]interface{}
}

func (P *AddNiciraNvpDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := P.P["l2gatewayserviceuuid"]; found {
		u.Set("l2gatewayserviceuuid", v.(string))
	}
	if v, found := P.P["l3gatewayserviceuuid"]; found {
		u.Set("l3gatewayserviceuuid", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["transportzoneuuid"]; found {
		u.Set("transportzoneuuid", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *AddNiciraNvpDeviceParams) SetHostname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostname"] = v
}

func (P *AddNiciraNvpDeviceParams) GetHostname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostname"].(string)
	return value, ok
}

func (P *AddNiciraNvpDeviceParams) SetL2gatewayserviceuuid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["l2gatewayserviceuuid"] = v
}

func (P *AddNiciraNvpDeviceParams) GetL2gatewayserviceuuid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["l2gatewayserviceuuid"].(string)
	return value, ok
}

func (P *AddNiciraNvpDeviceParams) SetL3gatewayserviceuuid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["l3gatewayserviceuuid"] = v
}

func (P *AddNiciraNvpDeviceParams) GetL3gatewayserviceuuid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["l3gatewayserviceuuid"].(string)
	return value, ok
}

func (P *AddNiciraNvpDeviceParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddNiciraNvpDeviceParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddNiciraNvpDeviceParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddNiciraNvpDeviceParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddNiciraNvpDeviceParams) SetTransportzoneuuid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["transportzoneuuid"] = v
}

func (P *AddNiciraNvpDeviceParams) GetTransportzoneuuid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["transportzoneuuid"].(string)
	return value, ok
}

func (P *AddNiciraNvpDeviceParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddNiciraNvpDeviceParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddNiciraNvpDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NiciraNVPService) NewAddNiciraNvpDeviceParams(hostname string, password string, physicalnetworkid string, transportzoneuuid string, username string) *AddNiciraNvpDeviceParams {
	P := &AddNiciraNvpDeviceParams{}
	P.P = make(map[string]interface{})
	P.P["hostname"] = hostname
	P.P["password"] = password
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["transportzoneuuid"] = transportzoneuuid
	P.P["username"] = username
	return P
}

// Adds a Nicira NVP device
func (s *NiciraNVPService) AddNiciraNvpDevice(p *AddNiciraNvpDeviceParams) (*AddNiciraNvpDeviceResponse, error) {
	resp, err := s.cs.newRequest("addNiciraNvpDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNiciraNvpDeviceResponse
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

type AddNiciraNvpDeviceResponse struct {
	Hostname             string `json:"hostname"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	L2gatewayserviceuuid string `json:"l2gatewayserviceuuid"`
	L3gatewayserviceuuid string `json:"l3gatewayserviceuuid"`
	Niciradevicename     string `json:"niciradevicename"`
	Nvpdeviceid          string `json:"nvpdeviceid"`
	Physicalnetworkid    string `json:"physicalnetworkid"`
	Provider             string `json:"provider"`
	Transportzoneuuid    string `json:"transportzoneuuid"`
}

type DeleteNiciraNvpDeviceParams struct {
	P map[string]interface{}
}

func (P *DeleteNiciraNvpDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["nvpdeviceid"]; found {
		u.Set("nvpdeviceid", v.(string))
	}
	return u
}

func (P *DeleteNiciraNvpDeviceParams) SetNvpdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nvpdeviceid"] = v
}

func (P *DeleteNiciraNvpDeviceParams) GetNvpdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nvpdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNiciraNvpDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NiciraNVPService) NewDeleteNiciraNvpDeviceParams(nvpdeviceid string) *DeleteNiciraNvpDeviceParams {
	P := &DeleteNiciraNvpDeviceParams{}
	P.P = make(map[string]interface{})
	P.P["nvpdeviceid"] = nvpdeviceid
	return P
}

//  delete a nicira nvp device
func (s *NiciraNVPService) DeleteNiciraNvpDevice(p *DeleteNiciraNvpDeviceParams) (*DeleteNiciraNvpDeviceResponse, error) {
	resp, err := s.cs.newRequest("deleteNiciraNvpDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNiciraNvpDeviceResponse
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

type DeleteNiciraNvpDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListNiciraNvpDevicesParams struct {
	P map[string]interface{}
}

func (P *ListNiciraNvpDevicesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["nvpdeviceid"]; found {
		u.Set("nvpdeviceid", v.(string))
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

func (P *ListNiciraNvpDevicesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNiciraNvpDevicesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNiciraNvpDevicesParams) SetNvpdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nvpdeviceid"] = v
}

func (P *ListNiciraNvpDevicesParams) GetNvpdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nvpdeviceid"].(string)
	return value, ok
}

func (P *ListNiciraNvpDevicesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNiciraNvpDevicesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNiciraNvpDevicesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNiciraNvpDevicesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListNiciraNvpDevicesParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListNiciraNvpDevicesParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNiciraNvpDevicesParams instance,
// as then you are sure you have configured all required params
func (s *NiciraNVPService) NewListNiciraNvpDevicesParams() *ListNiciraNvpDevicesParams {
	P := &ListNiciraNvpDevicesParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists Nicira NVP devices
func (s *NiciraNVPService) ListNiciraNvpDevices(p *ListNiciraNvpDevicesParams) (*ListNiciraNvpDevicesResponse, error) {
	resp, err := s.cs.newRequest("listNiciraNvpDevices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNiciraNvpDevicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNiciraNvpDevicesResponse struct {
	Count            int                `json:"count"`
	NiciraNvpDevices []*NiciraNvpDevice `json:"niciranvpdevice"`
}

type NiciraNvpDevice struct {
	Hostname             string `json:"hostname"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	L2gatewayserviceuuid string `json:"l2gatewayserviceuuid"`
	L3gatewayserviceuuid string `json:"l3gatewayserviceuuid"`
	Niciradevicename     string `json:"niciradevicename"`
	Nvpdeviceid          string `json:"nvpdeviceid"`
	Physicalnetworkid    string `json:"physicalnetworkid"`
	Provider             string `json:"provider"`
	Transportzoneuuid    string `json:"transportzoneuuid"`
}
