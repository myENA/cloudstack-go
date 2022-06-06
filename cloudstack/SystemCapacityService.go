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

type SystemCapacityServiceIface interface {
	ListCapacity(P *ListCapacityParams) (*ListCapacityResponse, error)
	NewListCapacityParams() *ListCapacityParams
}

type ListCapacityParams struct {
	P map[string]interface{}
}

func (P *ListCapacityParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["fetchlatest"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fetchlatest", vv)
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
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["sortby"]; found {
		u.Set("sortby", v.(string))
	}
	if v, found := P.P["type"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("type", vv)
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListCapacityParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *ListCapacityParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *ListCapacityParams) SetFetchlatest(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fetchlatest"] = v
}

func (P *ListCapacityParams) GetFetchlatest() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fetchlatest"].(bool)
	return value, ok
}

func (P *ListCapacityParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListCapacityParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListCapacityParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListCapacityParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListCapacityParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListCapacityParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListCapacityParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListCapacityParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListCapacityParams) SetSortby(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sortby"] = v
}

func (P *ListCapacityParams) GetSortby() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sortby"].(string)
	return value, ok
}

func (P *ListCapacityParams) SetType(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ListCapacityParams) GetType() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(int)
	return value, ok
}

func (P *ListCapacityParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListCapacityParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListCapacityParams instance,
// as then you are sure you have configured all required params
func (s *SystemCapacityService) NewListCapacityParams() *ListCapacityParams {
	P := &ListCapacityParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists all the system wide capacities.
func (s *SystemCapacityService) ListCapacity(p *ListCapacityParams) (*ListCapacityResponse, error) {
	resp, err := s.cs.newRequest("listCapacity", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCapacityResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCapacityResponse struct {
	Count    int         `json:"count"`
	Capacity []*Capacity `json:"capacity"`
}

type Capacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}
