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

type StratosphereSSPServiceIface interface {
	AddStratosphereSsp(p *AddStratosphereSspParams) (*AddStratosphereSspResponse, error)
	NewAddStratosphereSspParams(name string, url string, zoneid string) *AddStratosphereSspParams
	DeleteStratosphereSsp(p *DeleteStratosphereSspParams) (*DeleteStratosphereSspResponse, error)
	NewDeleteStratosphereSspParams(hostid string) *DeleteStratosphereSspParams
}

type AddStratosphereSspParams struct {
	P map[string]interface{}
}

func (P *AddStratosphereSspParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["tenantuuid"]; found {
		u.Set("tenantuuid", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *AddStratosphereSspParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *AddStratosphereSspParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *AddStratosphereSspParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddStratosphereSspParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddStratosphereSspParams) SetTenantuuid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tenantuuid"] = v
}

func (P *AddStratosphereSspParams) GetTenantuuid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tenantuuid"].(string)
	return value, ok
}

func (P *AddStratosphereSspParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddStratosphereSspParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddStratosphereSspParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddStratosphereSspParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

func (P *AddStratosphereSspParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *AddStratosphereSspParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddStratosphereSspParams instance,
// as then you are sure you have configured all required params
func (s *StratosphereSSPService) NewAddStratosphereSspParams(name string, url string, zoneid string) *AddStratosphereSspParams {
	P := &AddStratosphereSspParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	P.P["url"] = url
	P.P["zoneid"] = zoneid
	return P
}

// Adds stratosphere ssp server
func (s *StratosphereSSPService) AddStratosphereSsp(p *AddStratosphereSspParams) (*AddStratosphereSspResponse, error) {
	resp, err := s.cs.newRequest("addStratosphereSsp", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddStratosphereSspResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddStratosphereSspResponse struct {
	Hostid    string `json:"hostid"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Zoneid    string `json:"zoneid"`
}

type DeleteStratosphereSspParams struct {
	P map[string]interface{}
}

func (P *DeleteStratosphereSspParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (P *DeleteStratosphereSspParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *DeleteStratosphereSspParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteStratosphereSspParams instance,
// as then you are sure you have configured all required params
func (s *StratosphereSSPService) NewDeleteStratosphereSspParams(hostid string) *DeleteStratosphereSspParams {
	P := &DeleteStratosphereSspParams{}
	P.P = make(map[string]interface{})
	P.P["hostid"] = hostid
	return P
}

// Removes stratosphere ssp server
func (s *StratosphereSSPService) DeleteStratosphereSsp(p *DeleteStratosphereSspParams) (*DeleteStratosphereSspResponse, error) {
	resp, err := s.cs.newRequest("deleteStratosphereSsp", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteStratosphereSspResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteStratosphereSspResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteStratosphereSspResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteStratosphereSspResponse
	return json.Unmarshal(b, (*alias)(r))
}
