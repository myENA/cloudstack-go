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

type ImageStoreServiceIface interface {
	AddImageStore(P *AddImageStoreParams) (*AddImageStoreResponse, error)
	NewAddImageStoreParams(provider string) *AddImageStoreParams
	AddImageStoreS3(P *AddImageStoreS3Params) (*AddImageStoreS3Response, error)
	NewAddImageStoreS3Params(accesskey string, bucket string, endpoint string, secretkey string) *AddImageStoreS3Params
	CreateSecondaryStagingStore(P *CreateSecondaryStagingStoreParams) (*CreateSecondaryStagingStoreResponse, error)
	NewCreateSecondaryStagingStoreParams(url string) *CreateSecondaryStagingStoreParams
	DeleteImageStore(P *DeleteImageStoreParams) (*DeleteImageStoreResponse, error)
	NewDeleteImageStoreParams(id string) *DeleteImageStoreParams
	DeleteSecondaryStagingStore(P *DeleteSecondaryStagingStoreParams) (*DeleteSecondaryStagingStoreResponse, error)
	NewDeleteSecondaryStagingStoreParams(id string) *DeleteSecondaryStagingStoreParams
	ListImageStores(P *ListImageStoresParams) (*ListImageStoresResponse, error)
	NewListImageStoresParams() *ListImageStoresParams
	GetImageStoreID(name string, opts ...OptionFunc) (string, int, error)
	GetImageStoreByName(name string, opts ...OptionFunc) (*ImageStore, int, error)
	GetImageStoreByID(id string, opts ...OptionFunc) (*ImageStore, int, error)
	ListSecondaryStagingStores(P *ListSecondaryStagingStoresParams) (*ListSecondaryStagingStoresResponse, error)
	NewListSecondaryStagingStoresParams() *ListSecondaryStagingStoresParams
	GetSecondaryStagingStoreID(name string, opts ...OptionFunc) (string, int, error)
	GetSecondaryStagingStoreByName(name string, opts ...OptionFunc) (*SecondaryStagingStore, int, error)
	GetSecondaryStagingStoreByID(id string, opts ...OptionFunc) (*SecondaryStagingStore, int, error)
	UpdateCloudToUseObjectStore(P *UpdateCloudToUseObjectStoreParams) (*UpdateCloudToUseObjectStoreResponse, error)
	NewUpdateCloudToUseObjectStoreParams(provider string) *UpdateCloudToUseObjectStoreParams
}

type AddImageStoreParams struct {
	P map[string]interface{}
}

func (P *AddImageStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), m[k])
		}
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *AddImageStoreParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *AddImageStoreParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *AddImageStoreParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *AddImageStoreParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *AddImageStoreParams) SetProvider(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["provider"] = v
}

func (P *AddImageStoreParams) GetProvider() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["provider"].(string)
	return value, ok
}

func (P *AddImageStoreParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddImageStoreParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddImageStoreParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *AddImageStoreParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddImageStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewAddImageStoreParams(provider string) *AddImageStoreParams {
	P := &AddImageStoreParams{}
	P.P = make(map[string]interface{})
	P.P["provider"] = provider
	return P
}

// Adds backup image store.
func (s *ImageStoreService) AddImageStore(p *AddImageStoreParams) (*AddImageStoreResponse, error) {
	resp, err := s.cs.newRequest("addImageStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r AddImageStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddImageStoreResponse struct {
	Disksizetotal  int64  `json:"disksizetotal"`
	Disksizeused   int64  `json:"disksizeused"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Providername   string `json:"providername"`
	Readonly       bool   `json:"readonly"`
	Scope          string `json:"scope"`
	Url            string `json:"url"`
	Zoneid         string `json:"zoneid"`
	Zonename       string `json:"zonename"`
}

type AddImageStoreS3Params struct {
	P map[string]interface{}
}

func (P *AddImageStoreS3Params) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["accesskey"]; found {
		u.Set("accesskey", v.(string))
	}
	if v, found := P.P["bucket"]; found {
		u.Set("bucket", v.(string))
	}
	if v, found := P.P["connectiontimeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("connectiontimeout", vv)
	}
	if v, found := P.P["connectionttl"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("connectionttl", vv)
	}
	if v, found := P.P["endpoint"]; found {
		u.Set("endpoint", v.(string))
	}
	if v, found := P.P["maxerrorretry"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxerrorretry", vv)
	}
	if v, found := P.P["s3signer"]; found {
		u.Set("s3signer", v.(string))
	}
	if v, found := P.P["secretkey"]; found {
		u.Set("secretkey", v.(string))
	}
	if v, found := P.P["sockettimeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sockettimeout", vv)
	}
	if v, found := P.P["usehttps"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("usehttps", vv)
	}
	if v, found := P.P["usetcpkeepalive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("usetcpkeepalive", vv)
	}
	return u
}

func (P *AddImageStoreS3Params) SetAccesskey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accesskey"] = v
}

func (P *AddImageStoreS3Params) GetAccesskey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accesskey"].(string)
	return value, ok
}

func (P *AddImageStoreS3Params) SetBucket(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bucket"] = v
}

func (P *AddImageStoreS3Params) GetBucket() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bucket"].(string)
	return value, ok
}

func (P *AddImageStoreS3Params) SetConnectiontimeout(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["connectiontimeout"] = v
}

func (P *AddImageStoreS3Params) GetConnectiontimeout() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["connectiontimeout"].(int)
	return value, ok
}

func (P *AddImageStoreS3Params) SetConnectionttl(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["connectionttl"] = v
}

func (P *AddImageStoreS3Params) GetConnectionttl() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["connectionttl"].(int)
	return value, ok
}

func (P *AddImageStoreS3Params) SetEndpoint(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endpoint"] = v
}

func (P *AddImageStoreS3Params) GetEndpoint() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endpoint"].(string)
	return value, ok
}

func (P *AddImageStoreS3Params) SetMaxerrorretry(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxerrorretry"] = v
}

func (P *AddImageStoreS3Params) GetMaxerrorretry() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxerrorretry"].(int)
	return value, ok
}

func (P *AddImageStoreS3Params) SetS3signer(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["s3signer"] = v
}

func (P *AddImageStoreS3Params) GetS3signer() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["s3signer"].(string)
	return value, ok
}

func (P *AddImageStoreS3Params) SetSecretkey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["secretkey"] = v
}

func (P *AddImageStoreS3Params) GetSecretkey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["secretkey"].(string)
	return value, ok
}

func (P *AddImageStoreS3Params) SetSockettimeout(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sockettimeout"] = v
}

func (P *AddImageStoreS3Params) GetSockettimeout() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sockettimeout"].(int)
	return value, ok
}

func (P *AddImageStoreS3Params) SetUsehttps(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["usehttps"] = v
}

func (P *AddImageStoreS3Params) GetUsehttps() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["usehttps"].(bool)
	return value, ok
}

func (P *AddImageStoreS3Params) SetUsetcpkeepalive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["usetcpkeepalive"] = v
}

func (P *AddImageStoreS3Params) GetUsetcpkeepalive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["usetcpkeepalive"].(bool)
	return value, ok
}

// You should always use this function to get a new AddImageStoreS3Params instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewAddImageStoreS3Params(accesskey string, bucket string, endpoint string, secretkey string) *AddImageStoreS3Params {
	P := &AddImageStoreS3Params{}
	P.P = make(map[string]interface{})
	P.P["accesskey"] = accesskey
	P.P["bucket"] = bucket
	P.P["endpoint"] = endpoint
	P.P["secretkey"] = secretkey
	return P
}

// Adds S3 Image Store
func (s *ImageStoreService) AddImageStoreS3(p *AddImageStoreS3Params) (*AddImageStoreS3Response, error) {
	resp, err := s.cs.newRequest("addImageStoreS3", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddImageStoreS3Response
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddImageStoreS3Response struct {
	Disksizetotal  int64  `json:"disksizetotal"`
	Disksizeused   int64  `json:"disksizeused"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Providername   string `json:"providername"`
	Readonly       bool   `json:"readonly"`
	Scope          string `json:"scope"`
	Url            string `json:"url"`
	Zoneid         string `json:"zoneid"`
	Zonename       string `json:"zonename"`
}

type CreateSecondaryStagingStoreParams struct {
	P map[string]interface{}
}

func (P *CreateSecondaryStagingStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), m[k])
		}
	}
	if v, found := P.P["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := P.P["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *CreateSecondaryStagingStoreParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *CreateSecondaryStagingStoreParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *CreateSecondaryStagingStoreParams) SetProvider(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["provider"] = v
}

func (P *CreateSecondaryStagingStoreParams) GetProvider() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["provider"].(string)
	return value, ok
}

func (P *CreateSecondaryStagingStoreParams) SetScope(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["scope"] = v
}

func (P *CreateSecondaryStagingStoreParams) GetScope() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["scope"].(string)
	return value, ok
}

func (P *CreateSecondaryStagingStoreParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *CreateSecondaryStagingStoreParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *CreateSecondaryStagingStoreParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateSecondaryStagingStoreParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateSecondaryStagingStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewCreateSecondaryStagingStoreParams(url string) *CreateSecondaryStagingStoreParams {
	P := &CreateSecondaryStagingStoreParams{}
	P.P = make(map[string]interface{})
	P.P["url"] = url
	return P
}

// create secondary staging store.
func (s *ImageStoreService) CreateSecondaryStagingStore(p *CreateSecondaryStagingStoreParams) (*CreateSecondaryStagingStoreResponse, error) {
	resp, err := s.cs.newRequest("createSecondaryStagingStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateSecondaryStagingStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateSecondaryStagingStoreResponse struct {
	Disksizetotal  int64  `json:"disksizetotal"`
	Disksizeused   int64  `json:"disksizeused"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Providername   string `json:"providername"`
	Readonly       bool   `json:"readonly"`
	Scope          string `json:"scope"`
	Url            string `json:"url"`
	Zoneid         string `json:"zoneid"`
	Zonename       string `json:"zonename"`
}

type DeleteImageStoreParams struct {
	P map[string]interface{}
}

func (P *DeleteImageStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteImageStoreParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteImageStoreParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteImageStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewDeleteImageStoreParams(id string) *DeleteImageStoreParams {
	P := &DeleteImageStoreParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes an image store or Secondary Storage.
func (s *ImageStoreService) DeleteImageStore(p *DeleteImageStoreParams) (*DeleteImageStoreResponse, error) {
	resp, err := s.cs.newRequest("deleteImageStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteImageStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteImageStoreResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteImageStoreResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteImageStoreResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteSecondaryStagingStoreParams struct {
	P map[string]interface{}
}

func (P *DeleteSecondaryStagingStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteSecondaryStagingStoreParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteSecondaryStagingStoreParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteSecondaryStagingStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewDeleteSecondaryStagingStoreParams(id string) *DeleteSecondaryStagingStoreParams {
	P := &DeleteSecondaryStagingStoreParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a secondary staging store .
func (s *ImageStoreService) DeleteSecondaryStagingStore(p *DeleteSecondaryStagingStoreParams) (*DeleteSecondaryStagingStoreResponse, error) {
	resp, err := s.cs.newRequest("deleteSecondaryStagingStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSecondaryStagingStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSecondaryStagingStoreResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSecondaryStagingStoreResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteSecondaryStagingStoreResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListImageStoresParams struct {
	P map[string]interface{}
}

func (P *ListImageStoresParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
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
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := P.P["readonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("readonly", vv)
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListImageStoresParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListImageStoresParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListImageStoresParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListImageStoresParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListImageStoresParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListImageStoresParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListImageStoresParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListImageStoresParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListImageStoresParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListImageStoresParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListImageStoresParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *ListImageStoresParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *ListImageStoresParams) SetProvider(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["provider"] = v
}

func (P *ListImageStoresParams) GetProvider() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["provider"].(string)
	return value, ok
}

func (P *ListImageStoresParams) SetReadonly(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["readonly"] = v
}

func (P *ListImageStoresParams) GetReadonly() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["readonly"].(bool)
	return value, ok
}

func (P *ListImageStoresParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListImageStoresParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListImageStoresParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewListImageStoresParams() *ListImageStoresParams {
	P := &ListImageStoresParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetImageStoreID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListImageStoresParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListImageStores(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ImageStores[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ImageStores {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetImageStoreByName(name string, opts ...OptionFunc) (*ImageStore, int, error) {
	id, count, err := s.GetImageStoreID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetImageStoreByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetImageStoreByID(id string, opts ...OptionFunc) (*ImageStore, int, error) {
	P := &ListImageStoresParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListImageStores(P)
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
		return l.ImageStores[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ImageStore UUID: %s!", id)
}

// Lists image stores.
func (s *ImageStoreService) ListImageStores(p *ListImageStoresParams) (*ListImageStoresResponse, error) {
	resp, err := s.cs.newRequest("listImageStores", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListImageStoresResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListImageStoresResponse struct {
	Count       int           `json:"count"`
	ImageStores []*ImageStore `json:"imagestore"`
}

type ImageStore struct {
	Disksizetotal  int64  `json:"disksizetotal"`
	Disksizeused   int64  `json:"disksizeused"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Providername   string `json:"providername"`
	Readonly       bool   `json:"readonly"`
	Scope          string `json:"scope"`
	Url            string `json:"url"`
	Zoneid         string `json:"zoneid"`
	Zonename       string `json:"zonename"`
}

type ListSecondaryStagingStoresParams struct {
	P map[string]interface{}
}

func (P *ListSecondaryStagingStoresParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
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
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListSecondaryStagingStoresParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListSecondaryStagingStoresParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListSecondaryStagingStoresParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListSecondaryStagingStoresParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListSecondaryStagingStoresParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListSecondaryStagingStoresParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListSecondaryStagingStoresParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListSecondaryStagingStoresParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListSecondaryStagingStoresParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListSecondaryStagingStoresParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListSecondaryStagingStoresParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *ListSecondaryStagingStoresParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *ListSecondaryStagingStoresParams) SetProvider(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["provider"] = v
}

func (P *ListSecondaryStagingStoresParams) GetProvider() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["provider"].(string)
	return value, ok
}

func (P *ListSecondaryStagingStoresParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListSecondaryStagingStoresParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSecondaryStagingStoresParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewListSecondaryStagingStoresParams() *ListSecondaryStagingStoresParams {
	P := &ListSecondaryStagingStoresParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetSecondaryStagingStoreID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListSecondaryStagingStoresParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSecondaryStagingStores(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.SecondaryStagingStores[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SecondaryStagingStores {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetSecondaryStagingStoreByName(name string, opts ...OptionFunc) (*SecondaryStagingStore, int, error) {
	id, count, err := s.GetSecondaryStagingStoreID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSecondaryStagingStoreByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ImageStoreService) GetSecondaryStagingStoreByID(id string, opts ...OptionFunc) (*SecondaryStagingStore, int, error) {
	P := &ListSecondaryStagingStoresParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSecondaryStagingStores(P)
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
		return l.SecondaryStagingStores[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SecondaryStagingStore UUID: %s!", id)
}

// Lists secondary staging stores.
func (s *ImageStoreService) ListSecondaryStagingStores(p *ListSecondaryStagingStoresParams) (*ListSecondaryStagingStoresResponse, error) {
	resp, err := s.cs.newRequest("listSecondaryStagingStores", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSecondaryStagingStoresResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSecondaryStagingStoresResponse struct {
	Count                  int                      `json:"count"`
	SecondaryStagingStores []*SecondaryStagingStore `json:"secondarystagingstore"`
}

type SecondaryStagingStore struct {
	Disksizetotal  int64  `json:"disksizetotal"`
	Disksizeused   int64  `json:"disksizeused"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Providername   string `json:"providername"`
	Readonly       bool   `json:"readonly"`
	Scope          string `json:"scope"`
	Url            string `json:"url"`
	Zoneid         string `json:"zoneid"`
	Zonename       string `json:"zonename"`
}

type UpdateCloudToUseObjectStoreParams struct {
	P map[string]interface{}
}

func (P *UpdateCloudToUseObjectStoreParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), m[k])
		}
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	return u
}

func (P *UpdateCloudToUseObjectStoreParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *UpdateCloudToUseObjectStoreParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *UpdateCloudToUseObjectStoreParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateCloudToUseObjectStoreParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateCloudToUseObjectStoreParams) SetProvider(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["provider"] = v
}

func (P *UpdateCloudToUseObjectStoreParams) GetProvider() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["provider"].(string)
	return value, ok
}

func (P *UpdateCloudToUseObjectStoreParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *UpdateCloudToUseObjectStoreParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateCloudToUseObjectStoreParams instance,
// as then you are sure you have configured all required params
func (s *ImageStoreService) NewUpdateCloudToUseObjectStoreParams(provider string) *UpdateCloudToUseObjectStoreParams {
	P := &UpdateCloudToUseObjectStoreParams{}
	P.P = make(map[string]interface{})
	P.P["provider"] = provider
	return P
}

// Migrate current NFS secondary storages to use object store.
func (s *ImageStoreService) UpdateCloudToUseObjectStore(p *UpdateCloudToUseObjectStoreParams) (*UpdateCloudToUseObjectStoreResponse, error) {
	resp, err := s.cs.newRequest("updateCloudToUseObjectStore", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateCloudToUseObjectStoreResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateCloudToUseObjectStoreResponse struct {
	Disksizetotal  int64  `json:"disksizetotal"`
	Disksizeused   int64  `json:"disksizeused"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Providername   string `json:"providername"`
	Readonly       bool   `json:"readonly"`
	Scope          string `json:"scope"`
	Url            string `json:"url"`
	Zoneid         string `json:"zoneid"`
	Zonename       string `json:"zonename"`
}
