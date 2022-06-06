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

type RegionServiceIface interface {
	AddRegion(P *AddRegionParams) (*AddRegionResponse, error)
	NewAddRegionParams(endpoint string, id int, name string) *AddRegionParams
	ListRegions(P *ListRegionsParams) (*ListRegionsResponse, error)
	NewListRegionsParams() *ListRegionsParams
	RemoveRegion(P *RemoveRegionParams) (*RemoveRegionResponse, error)
	NewRemoveRegionParams(id int) *RemoveRegionParams
	UpdateRegion(P *UpdateRegionParams) (*UpdateRegionResponse, error)
	NewUpdateRegionParams(id int) *UpdateRegionParams
}

type AddRegionParams struct {
	P map[string]interface{}
}

func (P *AddRegionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["endpoint"]; found {
		u.Set("endpoint", v.(string))
	}
	if v, found := P.P["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (P *AddRegionParams) SetEndpoint(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endpoint"] = v
}

func (P *AddRegionParams) GetEndpoint() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endpoint"].(string)
	return value, ok
}

func (P *AddRegionParams) SetId(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *AddRegionParams) GetId() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(int)
	return value, ok
}

func (P *AddRegionParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *AddRegionParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

// You should always use this function to get a new AddRegionParams instance,
// as then you are sure you have configured all required params
func (s *RegionService) NewAddRegionParams(endpoint string, id int, name string) *AddRegionParams {
	P := &AddRegionParams{}
	P.P = make(map[string]interface{})
	P.P["endpoint"] = endpoint
	P.P["id"] = id
	P.P["name"] = name
	return P
}

// Adds a Region
func (s *RegionService) AddRegion(p *AddRegionParams) (*AddRegionResponse, error) {
	resp, err := s.cs.newRequest("addRegion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddRegionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddRegionResponse struct {
	Endpoint                 string `json:"endpoint"`
	Gslbserviceenabled       bool   `json:"gslbserviceenabled"`
	Id                       int    `json:"id"`
	JobID                    string `json:"jobid"`
	Jobstatus                int    `json:"jobstatus"`
	Name                     string `json:"name"`
	Portableipserviceenabled bool   `json:"portableipserviceenabled"`
}

type ListRegionsParams struct {
	P map[string]interface{}
}

func (P *ListRegionsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
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

func (P *ListRegionsParams) SetId(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListRegionsParams) GetId() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(int)
	return value, ok
}

func (P *ListRegionsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListRegionsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListRegionsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListRegionsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListRegionsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListRegionsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListRegionsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListRegionsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListRegionsParams instance,
// as then you are sure you have configured all required params
func (s *RegionService) NewListRegionsParams() *ListRegionsParams {
	P := &ListRegionsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists Regions
func (s *RegionService) ListRegions(p *ListRegionsParams) (*ListRegionsResponse, error) {
	resp, err := s.cs.newRequest("listRegions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRegionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListRegionsResponse struct {
	Count   int       `json:"count"`
	Regions []*Region `json:"region"`
}

type Region struct {
	Endpoint                 string `json:"endpoint"`
	Gslbserviceenabled       bool   `json:"gslbserviceenabled"`
	Id                       int    `json:"id"`
	JobID                    string `json:"jobid"`
	Jobstatus                int    `json:"jobstatus"`
	Name                     string `json:"name"`
	Portableipserviceenabled bool   `json:"portableipserviceenabled"`
}

type RemoveRegionParams struct {
	P map[string]interface{}
}

func (P *RemoveRegionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
	}
	return u
}

func (P *RemoveRegionParams) SetId(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RemoveRegionParams) GetId() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(int)
	return value, ok
}

// You should always use this function to get a new RemoveRegionParams instance,
// as then you are sure you have configured all required params
func (s *RegionService) NewRemoveRegionParams(id int) *RemoveRegionParams {
	P := &RemoveRegionParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Removes specified region
func (s *RegionService) RemoveRegion(p *RemoveRegionParams) (*RemoveRegionResponse, error) {
	resp, err := s.cs.newRequest("removeRegion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveRegionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveRegionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *RemoveRegionResponse) UnmarshalJSON(b []byte) error {
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

	type alias RemoveRegionResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateRegionParams struct {
	P map[string]interface{}
}

func (P *UpdateRegionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["endpoint"]; found {
		u.Set("endpoint", v.(string))
	}
	if v, found := P.P["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (P *UpdateRegionParams) SetEndpoint(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endpoint"] = v
}

func (P *UpdateRegionParams) GetEndpoint() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endpoint"].(string)
	return value, ok
}

func (P *UpdateRegionParams) SetId(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateRegionParams) GetId() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(int)
	return value, ok
}

func (P *UpdateRegionParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateRegionParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateRegionParams instance,
// as then you are sure you have configured all required params
func (s *RegionService) NewUpdateRegionParams(id int) *UpdateRegionParams {
	P := &UpdateRegionParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a region
func (s *RegionService) UpdateRegion(p *UpdateRegionParams) (*UpdateRegionResponse, error) {
	resp, err := s.cs.newRequest("updateRegion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateRegionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateRegionResponse struct {
	Endpoint                 string `json:"endpoint"`
	Gslbserviceenabled       bool   `json:"gslbserviceenabled"`
	Id                       int    `json:"id"`
	JobID                    string `json:"jobid"`
	Jobstatus                int    `json:"jobstatus"`
	Name                     string `json:"name"`
	Portableipserviceenabled bool   `json:"portableipserviceenabled"`
}
