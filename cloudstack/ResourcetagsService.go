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

type ResourcetagsServiceIface interface {
	CreateTags(p *CreateTagsParams) (*CreateTagsResponse, error)
	NewCreateTagsParams(resourceids []string, resourcetype string, tags map[string]string) *CreateTagsParams
	DeleteTags(p *DeleteTagsParams) (*DeleteTagsResponse, error)
	NewDeleteTagsParams(resourceids []string, resourcetype string) *DeleteTagsParams
	ListStorageTags(p *ListStorageTagsParams) (*ListStorageTagsResponse, error)
	NewListStorageTagsParams() *ListStorageTagsParams
	GetStorageTagID(keyword string, opts ...OptionFunc) (string, int, error)
	ListTags(p *ListTagsParams) (*ListTagsResponse, error)
	NewListTagsParams() *ListTagsParams
}

type CreateTagsParams struct {
	P map[string]interface{}
}

func (P *CreateTagsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["customer"]; found {
		u.Set("customer", v.(string))
	}
	if v, found := P.P["resourceids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("resourceids", vv)
	}
	if v, found := P.P["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
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

func (P *CreateTagsParams) SetCustomer(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customer"] = v
}

func (P *CreateTagsParams) GetCustomer() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customer"].(string)
	return value, ok
}

func (P *CreateTagsParams) SetResourceids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourceids"] = v
}

func (P *CreateTagsParams) GetResourceids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourceids"].([]string)
	return value, ok
}

func (P *CreateTagsParams) SetResourcetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcetype"] = v
}

func (P *CreateTagsParams) GetResourcetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcetype"].(string)
	return value, ok
}

func (P *CreateTagsParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *CreateTagsParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new CreateTagsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcetagsService) NewCreateTagsParams(resourceids []string, resourcetype string, tags map[string]string) *CreateTagsParams {
	P := &CreateTagsParams{}
	P.P = make(map[string]interface{})
	P.P["resourceids"] = resourceids
	P.P["resourcetype"] = resourcetype
	P.P["tags"] = tags
	return P
}

// Creates resource tag(s)
func (s *ResourcetagsService) CreateTags(p *CreateTagsParams) (*CreateTagsResponse, error) {
	resp, err := s.cs.newRequest("createTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTagsResponse
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

type CreateTagsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteTagsParams struct {
	P map[string]interface{}
}

func (P *DeleteTagsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["resourceids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("resourceids", vv)
	}
	if v, found := P.P["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			if m[k] != "" {
				u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
			}
		}
	}
	return u
}

func (P *DeleteTagsParams) SetResourceids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourceids"] = v
}

func (P *DeleteTagsParams) GetResourceids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourceids"].([]string)
	return value, ok
}

func (P *DeleteTagsParams) SetResourcetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcetype"] = v
}

func (P *DeleteTagsParams) GetResourcetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcetype"].(string)
	return value, ok
}

func (P *DeleteTagsParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *DeleteTagsParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new DeleteTagsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcetagsService) NewDeleteTagsParams(resourceids []string, resourcetype string) *DeleteTagsParams {
	P := &DeleteTagsParams{}
	P.P = make(map[string]interface{})
	P.P["resourceids"] = resourceids
	P.P["resourcetype"] = resourcetype
	return P
}

// Deleting resource tag(s)
func (s *ResourcetagsService) DeleteTags(p *DeleteTagsParams) (*DeleteTagsResponse, error) {
	resp, err := s.cs.newRequest("deleteTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTagsResponse
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

type DeleteTagsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListStorageTagsParams struct {
	P map[string]interface{}
}

func (P *ListStorageTagsParams) toURLValues() url.Values {
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
	return u
}

func (P *ListStorageTagsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListStorageTagsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListStorageTagsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListStorageTagsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListStorageTagsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListStorageTagsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListStorageTagsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcetagsService) NewListStorageTagsParams() *ListStorageTagsParams {
	P := &ListStorageTagsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ResourcetagsService) GetStorageTagID(keyword string, opts ...OptionFunc) (string, int, error) {
	P := &ListStorageTagsParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListStorageTags(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.StorageTags[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.StorageTags {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Lists storage tags
func (s *ResourcetagsService) ListStorageTags(p *ListStorageTagsParams) (*ListStorageTagsResponse, error) {
	resp, err := s.cs.newRequest("listStorageTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStorageTagsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStorageTagsResponse struct {
	Count       int           `json:"count"`
	StorageTags []*StorageTag `json:"storagetag"`
}

type StorageTag struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Poolid    int64  `json:"poolid"`
}

type ListTagsParams struct {
	P map[string]interface{}
}

func (P *ListTagsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["customer"]; found {
		u.Set("customer", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := P.P["key"]; found {
		u.Set("key", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["resourceid"]; found {
		u.Set("resourceid", v.(string))
	}
	if v, found := P.P["resourcetype"]; found {
		u.Set("resourcetype", v.(string))
	}
	if v, found := P.P["value"]; found {
		u.Set("value", v.(string))
	}
	return u
}

func (P *ListTagsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListTagsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListTagsParams) SetCustomer(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customer"] = v
}

func (P *ListTagsParams) GetCustomer() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customer"].(string)
	return value, ok
}

func (P *ListTagsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListTagsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListTagsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListTagsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListTagsParams) SetKey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["key"] = v
}

func (P *ListTagsParams) GetKey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["key"].(string)
	return value, ok
}

func (P *ListTagsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListTagsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListTagsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListTagsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListTagsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListTagsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListTagsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListTagsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListTagsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListTagsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListTagsParams) SetResourceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourceid"] = v
}

func (P *ListTagsParams) GetResourceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourceid"].(string)
	return value, ok
}

func (P *ListTagsParams) SetResourcetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcetype"] = v
}

func (P *ListTagsParams) GetResourcetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcetype"].(string)
	return value, ok
}

func (P *ListTagsParams) SetValue(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["value"] = v
}

func (P *ListTagsParams) GetValue() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["value"].(string)
	return value, ok
}

// You should always use this function to get a new ListTagsParams instance,
// as then you are sure you have configured all required params
func (s *ResourcetagsService) NewListTagsParams() *ListTagsParams {
	P := &ListTagsParams{}
	P.P = make(map[string]interface{})
	return P
}

// List resource tag(s)
func (s *ResourcetagsService) ListTags(p *ListTagsParams) (*ListTagsResponse, error) {
	resp, err := s.cs.newRequest("listTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTagsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTagsResponse struct {
	Count int    `json:"count"`
	Tags  []*Tag `json:"tag"`
}

type Tag struct {
	Account      string `json:"account"`
	Customer     string `json:"customer"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Key          string `json:"key"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Resourceid   string `json:"resourceid"`
	Resourcetype string `json:"resourcetype"`
	Value        string `json:"value"`
}
