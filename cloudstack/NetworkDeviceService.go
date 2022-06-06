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

type NetworkDeviceServiceIface interface {
	AddNetworkDevice(P *AddNetworkDeviceParams) (*AddNetworkDeviceResponse, error)
	NewAddNetworkDeviceParams() *AddNetworkDeviceParams
	DeleteNetworkDevice(P *DeleteNetworkDeviceParams) (*DeleteNetworkDeviceResponse, error)
	NewDeleteNetworkDeviceParams(id string) *DeleteNetworkDeviceParams
	ListNetworkDevice(P *ListNetworkDeviceParams) (*ListNetworkDeviceResponse, error)
	NewListNetworkDeviceParams() *ListNetworkDeviceParams
}

type AddNetworkDeviceParams struct {
	P map[string]interface{}
}

func (P *AddNetworkDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["networkdeviceparameterlist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("networkdeviceparameterlist[%d].key", i), k)
			u.Set(fmt.Sprintf("networkdeviceparameterlist[%d].value", i), m[k])
		}
	}
	if v, found := P.P["networkdevicetype"]; found {
		u.Set("networkdevicetype", v.(string))
	}
	return u
}

func (P *AddNetworkDeviceParams) SetNetworkdeviceparameterlist(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdeviceparameterlist"] = v
}

func (P *AddNetworkDeviceParams) GetNetworkdeviceparameterlist() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdeviceparameterlist"].(map[string]string)
	return value, ok
}

func (P *AddNetworkDeviceParams) SetNetworkdevicetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdevicetype"] = v
}

func (P *AddNetworkDeviceParams) GetNetworkdevicetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdevicetype"].(string)
	return value, ok
}

// You should always use this function to get a new AddNetworkDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NetworkDeviceService) NewAddNetworkDeviceParams() *AddNetworkDeviceParams {
	P := &AddNetworkDeviceParams{}
	P.P = make(map[string]interface{})
	return P
}

// Adds a network device of one of the following types: ExternalDhcp, ExternalFirewall, ExternalLoadBalancer, PxeServer
func (s *NetworkDeviceService) AddNetworkDevice(p *AddNetworkDeviceParams) (*AddNetworkDeviceResponse, error) {
	resp, err := s.cs.newRequest("addNetworkDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNetworkDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddNetworkDeviceResponse struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
}

type DeleteNetworkDeviceParams struct {
	P map[string]interface{}
}

func (P *DeleteNetworkDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteNetworkDeviceParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteNetworkDeviceParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NetworkDeviceService) NewDeleteNetworkDeviceParams(id string) *DeleteNetworkDeviceParams {
	P := &DeleteNetworkDeviceParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes network device.
func (s *NetworkDeviceService) DeleteNetworkDevice(p *DeleteNetworkDeviceParams) (*DeleteNetworkDeviceResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteNetworkDeviceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteNetworkDeviceResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteNetworkDeviceResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListNetworkDeviceParams struct {
	P map[string]interface{}
}

func (P *ListNetworkDeviceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["networkdeviceparameterlist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("networkdeviceparameterlist[%d].key", i), k)
			u.Set(fmt.Sprintf("networkdeviceparameterlist[%d].value", i), m[k])
		}
	}
	if v, found := P.P["networkdevicetype"]; found {
		u.Set("networkdevicetype", v.(string))
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

func (P *ListNetworkDeviceParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNetworkDeviceParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNetworkDeviceParams) SetNetworkdeviceparameterlist(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdeviceparameterlist"] = v
}

func (P *ListNetworkDeviceParams) GetNetworkdeviceparameterlist() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdeviceparameterlist"].(map[string]string)
	return value, ok
}

func (P *ListNetworkDeviceParams) SetNetworkdevicetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdevicetype"] = v
}

func (P *ListNetworkDeviceParams) GetNetworkdevicetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdevicetype"].(string)
	return value, ok
}

func (P *ListNetworkDeviceParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNetworkDeviceParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNetworkDeviceParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNetworkDeviceParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListNetworkDeviceParams instance,
// as then you are sure you have configured all required params
func (s *NetworkDeviceService) NewListNetworkDeviceParams() *ListNetworkDeviceParams {
	P := &ListNetworkDeviceParams{}
	P.P = make(map[string]interface{})
	return P
}

// List network devices
func (s *NetworkDeviceService) ListNetworkDevice(p *ListNetworkDeviceParams) (*ListNetworkDeviceResponse, error) {
	resp, err := s.cs.newRequest("listNetworkDevice", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkDeviceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkDeviceResponse struct {
	Count         int              `json:"count"`
	NetworkDevice []*NetworkDevice `json:"networkdevice"`
}

type NetworkDevice struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
}
