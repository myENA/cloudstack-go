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

type DomainServiceIface interface {
	CreateDomain(P *CreateDomainParams) (*CreateDomainResponse, error)
	NewCreateDomainParams(name string) *CreateDomainParams
	DeleteDomain(P *DeleteDomainParams) (*DeleteDomainResponse, error)
	NewDeleteDomainParams(id string) *DeleteDomainParams
	ListDomainChildren(P *ListDomainChildrenParams) (*ListDomainChildrenResponse, error)
	NewListDomainChildrenParams() *ListDomainChildrenParams
	GetDomainChildrenID(name string, opts ...OptionFunc) (string, int, error)
	GetDomainChildrenByName(name string, opts ...OptionFunc) (*DomainChildren, int, error)
	GetDomainChildrenByID(id string, opts ...OptionFunc) (*DomainChildren, int, error)
	ListDomains(P *ListDomainsParams) (*ListDomainsResponse, error)
	NewListDomainsParams() *ListDomainsParams
	GetDomainID(name string, opts ...OptionFunc) (string, int, error)
	GetDomainByName(name string, opts ...OptionFunc) (*Domain, int, error)
	GetDomainByID(id string, opts ...OptionFunc) (*Domain, int, error)
	UpdateDomain(P *UpdateDomainParams) (*UpdateDomainResponse, error)
	NewUpdateDomainParams(id string) *UpdateDomainParams
}

type CreateDomainParams struct {
	P map[string]interface{}
}

func (P *CreateDomainParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := P.P["parentdomainid"]; found {
		u.Set("parentdomainid", v.(string))
	}
	return u
}

func (P *CreateDomainParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateDomainParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateDomainParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateDomainParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateDomainParams) SetNetworkdomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdomain"] = v
}

func (P *CreateDomainParams) GetNetworkdomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdomain"].(string)
	return value, ok
}

func (P *CreateDomainParams) SetParentdomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["parentdomainid"] = v
}

func (P *CreateDomainParams) GetParentdomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["parentdomainid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateDomainParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewCreateDomainParams(name string) *CreateDomainParams {
	P := &CreateDomainParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	return P
}

// Creates a domain
func (s *DomainService) CreateDomain(p *CreateDomainParams) (*CreateDomainResponse, error) {
	resp, err := s.cs.newRequest("createDomain", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateDomainResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateDomainResponse struct {
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Domaindetails             map[string]string `json:"domaindetails"`
	Hasannotations            bool              `json:"hasannotations"`
	Haschild                  bool              `json:"haschild"`
	Icon                      string            `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Level                     int               `json:"level"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Parentdomainid            string            `json:"parentdomainid"`
	Parentdomainname          string            `json:"parentdomainname"`
	Path                      string            `json:"path"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}

type DeleteDomainParams struct {
	P map[string]interface{}
}

func (P *DeleteDomainParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["cleanup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanup", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteDomainParams) SetCleanup(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cleanup"] = v
}

func (P *DeleteDomainParams) GetCleanup() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cleanup"].(bool)
	return value, ok
}

func (P *DeleteDomainParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteDomainParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteDomainParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewDeleteDomainParams(id string) *DeleteDomainParams {
	P := &DeleteDomainParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a specified domain
func (s *DomainService) DeleteDomain(p *DeleteDomainParams) (*DeleteDomainResponse, error) {
	resp, err := s.cs.newRequest("deleteDomain", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteDomainResponse
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

type DeleteDomainResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListDomainChildrenParams struct {
	P map[string]interface{}
}

func (P *ListDomainChildrenParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
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
	if v, found := P.P["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	return u
}

func (P *ListDomainChildrenParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListDomainChildrenParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListDomainChildrenParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListDomainChildrenParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListDomainChildrenParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListDomainChildrenParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListDomainChildrenParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListDomainChildrenParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListDomainChildrenParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListDomainChildrenParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListDomainChildrenParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListDomainChildrenParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListDomainChildrenParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListDomainChildrenParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListDomainChildrenParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListDomainChildrenParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

// You should always use this function to get a new ListDomainChildrenParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewListDomainChildrenParams() *ListDomainChildrenParams {
	P := &ListDomainChildrenParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainChildrenID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListDomainChildrenParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListDomainChildren(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.DomainChildren[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.DomainChildren {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainChildrenByName(name string, opts ...OptionFunc) (*DomainChildren, int, error) {
	id, count, err := s.GetDomainChildrenID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetDomainChildrenByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainChildrenByID(id string, opts ...OptionFunc) (*DomainChildren, int, error) {
	P := &ListDomainChildrenParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListDomainChildren(P)
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
		return l.DomainChildren[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for DomainChildren UUID: %s!", id)
}

// Lists all children domains belonging to a specified domain
func (s *DomainService) ListDomainChildren(p *ListDomainChildrenParams) (*ListDomainChildrenResponse, error) {
	resp, err := s.cs.newRequest("listDomainChildren", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDomainChildrenResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDomainChildrenResponse struct {
	Count          int               `json:"count"`
	DomainChildren []*DomainChildren `json:"domainchildren"`
}

type DomainChildren struct {
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Domaindetails             map[string]string `json:"domaindetails"`
	Hasannotations            bool              `json:"hasannotations"`
	Haschild                  bool              `json:"haschild"`
	Icon                      string            `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Level                     int               `json:"level"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Parentdomainid            string            `json:"parentdomainid"`
	Parentdomainname          string            `json:"parentdomainname"`
	Path                      string            `json:"path"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}

type ListDomainsParams struct {
	P map[string]interface{}
}

func (P *ListDomainsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["level"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("level", vv)
	}
	if v, found := P.P["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
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
	if v, found := P.P["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	return u
}

func (P *ListDomainsParams) SetDetails(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ListDomainsParams) GetDetails() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].([]string)
	return value, ok
}

func (P *ListDomainsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListDomainsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListDomainsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListDomainsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListDomainsParams) SetLevel(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["level"] = v
}

func (P *ListDomainsParams) GetLevel() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["level"].(int)
	return value, ok
}

func (P *ListDomainsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListDomainsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListDomainsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListDomainsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListDomainsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListDomainsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListDomainsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListDomainsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListDomainsParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListDomainsParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

// You should always use this function to get a new ListDomainsParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewListDomainsParams() *ListDomainsParams {
	P := &ListDomainsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListDomainsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListDomains(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Domains[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Domains {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainByName(name string, opts ...OptionFunc) (*Domain, int, error) {
	id, count, err := s.GetDomainID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetDomainByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DomainService) GetDomainByID(id string, opts ...OptionFunc) (*Domain, int, error) {
	P := &ListDomainsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListDomains(P)
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
		return l.Domains[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Domain UUID: %s!", id)
}

// Lists domains and provides detailed information for listed domains
func (s *DomainService) ListDomains(p *ListDomainsParams) (*ListDomainsResponse, error) {
	resp, err := s.cs.newRequest("listDomains", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDomainsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDomainsResponse struct {
	Count   int       `json:"count"`
	Domains []*Domain `json:"domain"`
}

type Domain struct {
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Domaindetails             map[string]string `json:"domaindetails"`
	Hasannotations            bool              `json:"hasannotations"`
	Haschild                  bool              `json:"haschild"`
	Icon                      string            `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Level                     int               `json:"level"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Parentdomainid            string            `json:"parentdomainid"`
	Parentdomainname          string            `json:"parentdomainname"`
	Path                      string            `json:"path"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}

type UpdateDomainParams struct {
	P map[string]interface{}
}

func (P *UpdateDomainParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	return u
}

func (P *UpdateDomainParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateDomainParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateDomainParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateDomainParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateDomainParams) SetNetworkdomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdomain"] = v
}

func (P *UpdateDomainParams) GetNetworkdomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdomain"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateDomainParams instance,
// as then you are sure you have configured all required params
func (s *DomainService) NewUpdateDomainParams(id string) *UpdateDomainParams {
	P := &UpdateDomainParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a domain with a new name
func (s *DomainService) UpdateDomain(p *UpdateDomainParams) (*UpdateDomainResponse, error) {
	resp, err := s.cs.newRequest("updateDomain", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r UpdateDomainResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateDomainResponse struct {
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Domaindetails             map[string]string `json:"domaindetails"`
	Hasannotations            bool              `json:"hasannotations"`
	Haschild                  bool              `json:"haschild"`
	Icon                      string            `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Level                     int               `json:"level"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Parentdomainid            string            `json:"parentdomainid"`
	Parentdomainname          string            `json:"parentdomainname"`
	Path                      string            `json:"path"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}
