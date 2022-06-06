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

type OvsElementServiceIface interface {
	ConfigureOvsElement(P *ConfigureOvsElementParams) (*OvsElementResponse, error)
	NewConfigureOvsElementParams(enabled bool, id string) *ConfigureOvsElementParams
	ListOvsElements(P *ListOvsElementsParams) (*ListOvsElementsResponse, error)
	NewListOvsElementsParams() *ListOvsElementsParams
	GetOvsElementByID(id string, opts ...OptionFunc) (*OvsElement, int, error)
}

type ConfigureOvsElementParams struct {
	P map[string]interface{}
}

func (P *ConfigureOvsElementParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ConfigureOvsElementParams) SetEnabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enabled"] = v
}

func (P *ConfigureOvsElementParams) GetEnabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enabled"].(bool)
	return value, ok
}

func (P *ConfigureOvsElementParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ConfigureOvsElementParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ConfigureOvsElementParams instance,
// as then you are sure you have configured all required params
func (s *OvsElementService) NewConfigureOvsElementParams(enabled bool, id string) *ConfigureOvsElementParams {
	P := &ConfigureOvsElementParams{}
	P.P = make(map[string]interface{})
	P.P["enabled"] = enabled
	P.P["id"] = id
	return P
}

// Configures an ovs element.
func (s *OvsElementService) ConfigureOvsElement(p *ConfigureOvsElementParams) (*OvsElementResponse, error) {
	resp, err := s.cs.newRequest("configureOvsElement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r OvsElementResponse
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

type OvsElementResponse struct {
	Account   string `json:"account"`
	Domain    string `json:"domain"`
	Domainid  string `json:"domainid"`
	Enabled   bool   `json:"enabled"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Nspid     string `json:"nspid"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
}

type ListOvsElementsParams struct {
	P map[string]interface{}
}

func (P *ListOvsElementsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["nspid"]; found {
		u.Set("nspid", v.(string))
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

func (P *ListOvsElementsParams) SetEnabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enabled"] = v
}

func (P *ListOvsElementsParams) GetEnabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enabled"].(bool)
	return value, ok
}

func (P *ListOvsElementsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListOvsElementsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListOvsElementsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListOvsElementsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListOvsElementsParams) SetNspid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nspid"] = v
}

func (P *ListOvsElementsParams) GetNspid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nspid"].(string)
	return value, ok
}

func (P *ListOvsElementsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListOvsElementsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListOvsElementsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListOvsElementsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListOvsElementsParams instance,
// as then you are sure you have configured all required params
func (s *OvsElementService) NewListOvsElementsParams() *ListOvsElementsParams {
	P := &ListOvsElementsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *OvsElementService) GetOvsElementByID(id string, opts ...OptionFunc) (*OvsElement, int, error) {
	P := &ListOvsElementsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListOvsElements(P)
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
		return l.OvsElements[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for OvsElement UUID: %s!", id)
}

// Lists all available ovs elements.
func (s *OvsElementService) ListOvsElements(p *ListOvsElementsParams) (*ListOvsElementsResponse, error) {
	resp, err := s.cs.newRequest("listOvsElements", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListOvsElementsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListOvsElementsResponse struct {
	Count       int           `json:"count"`
	OvsElements []*OvsElement `json:"ovselement"`
}

type OvsElement struct {
	Account   string `json:"account"`
	Domain    string `json:"domain"`
	Domainid  string `json:"domainid"`
	Enabled   bool   `json:"enabled"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Nspid     string `json:"nspid"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
}
