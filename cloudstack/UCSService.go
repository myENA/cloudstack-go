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

type UCSServiceIface interface {
	AddUcsManager(P *AddUcsManagerParams) (*AddUcsManagerResponse, error)
	NewAddUcsManagerParams(password string, url string, username string, zoneid string) *AddUcsManagerParams
	AssociateUcsProfileToBlade(P *AssociateUcsProfileToBladeParams) (*AssociateUcsProfileToBladeResponse, error)
	NewAssociateUcsProfileToBladeParams(bladeid string, profiledn string, ucsmanagerid string) *AssociateUcsProfileToBladeParams
	DeleteUcsManager(P *DeleteUcsManagerParams) (*DeleteUcsManagerResponse, error)
	NewDeleteUcsManagerParams(ucsmanagerid string) *DeleteUcsManagerParams
	ListUcsBlades(P *ListUcsBladesParams) (*ListUcsBladesResponse, error)
	NewListUcsBladesParams(ucsmanagerid string) *ListUcsBladesParams
	ListUcsManagers(P *ListUcsManagersParams) (*ListUcsManagersResponse, error)
	NewListUcsManagersParams() *ListUcsManagersParams
	GetUcsManagerID(keyword string, opts ...OptionFunc) (string, int, error)
	GetUcsManagerByName(name string, opts ...OptionFunc) (*UcsManager, int, error)
	GetUcsManagerByID(id string, opts ...OptionFunc) (*UcsManager, int, error)
	ListUcsProfiles(P *ListUcsProfilesParams) (*ListUcsProfilesResponse, error)
	NewListUcsProfilesParams(ucsmanagerid string) *ListUcsProfilesParams
}

type AddUcsManagerParams struct {
	P map[string]interface{}
}

func (P *AddUcsManagerParams) toURLValues() url.Values {
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

func (P *AddUcsManagerParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *AddUcsManagerParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *AddUcsManagerParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddUcsManagerParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddUcsManagerParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddUcsManagerParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddUcsManagerParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddUcsManagerParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

func (P *AddUcsManagerParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *AddUcsManagerParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddUcsManagerParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewAddUcsManagerParams(password string, url string, username string, zoneid string) *AddUcsManagerParams {
	P := &AddUcsManagerParams{}
	P.P = make(map[string]interface{})
	P.P["password"] = password
	P.P["url"] = url
	P.P["username"] = username
	P.P["zoneid"] = zoneid
	return P
}

// Adds a Ucs manager
func (s *UCSService) AddUcsManager(p *AddUcsManagerParams) (*AddUcsManagerResponse, error) {
	resp, err := s.cs.newRequest("addUcsManager", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddUcsManagerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddUcsManagerResponse struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Zoneid    string `json:"zoneid"`
}

type AssociateUcsProfileToBladeParams struct {
	P map[string]interface{}
}

func (P *AssociateUcsProfileToBladeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["bladeid"]; found {
		u.Set("bladeid", v.(string))
	}
	if v, found := P.P["profiledn"]; found {
		u.Set("profiledn", v.(string))
	}
	if v, found := P.P["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (P *AssociateUcsProfileToBladeParams) SetBladeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bladeid"] = v
}

func (P *AssociateUcsProfileToBladeParams) GetBladeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bladeid"].(string)
	return value, ok
}

func (P *AssociateUcsProfileToBladeParams) SetProfiledn(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["profiledn"] = v
}

func (P *AssociateUcsProfileToBladeParams) GetProfiledn() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["profiledn"].(string)
	return value, ok
}

func (P *AssociateUcsProfileToBladeParams) SetUcsmanagerid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ucsmanagerid"] = v
}

func (P *AssociateUcsProfileToBladeParams) GetUcsmanagerid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ucsmanagerid"].(string)
	return value, ok
}

// You should always use this function to get a new AssociateUcsProfileToBladeParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewAssociateUcsProfileToBladeParams(bladeid string, profiledn string, ucsmanagerid string) *AssociateUcsProfileToBladeParams {
	P := &AssociateUcsProfileToBladeParams{}
	P.P = make(map[string]interface{})
	P.P["bladeid"] = bladeid
	P.P["profiledn"] = profiledn
	P.P["ucsmanagerid"] = ucsmanagerid
	return P
}

// associate a profile to a blade
func (s *UCSService) AssociateUcsProfileToBlade(p *AssociateUcsProfileToBladeParams) (*AssociateUcsProfileToBladeResponse, error) {
	resp, err := s.cs.newRequest("associateUcsProfileToBlade", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssociateUcsProfileToBladeResponse
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

type AssociateUcsProfileToBladeResponse struct {
	Bladedn      string `json:"bladedn"`
	Hostid       string `json:"hostid"`
	Id           string `json:"id"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Profiledn    string `json:"profiledn"`
	Ucsmanagerid string `json:"ucsmanagerid"`
}

type DeleteUcsManagerParams struct {
	P map[string]interface{}
}

func (P *DeleteUcsManagerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (P *DeleteUcsManagerParams) SetUcsmanagerid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ucsmanagerid"] = v
}

func (P *DeleteUcsManagerParams) GetUcsmanagerid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ucsmanagerid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteUcsManagerParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewDeleteUcsManagerParams(ucsmanagerid string) *DeleteUcsManagerParams {
	P := &DeleteUcsManagerParams{}
	P.P = make(map[string]interface{})
	P.P["ucsmanagerid"] = ucsmanagerid
	return P
}

// Delete a Ucs manager
func (s *UCSService) DeleteUcsManager(p *DeleteUcsManagerParams) (*DeleteUcsManagerResponse, error) {
	resp, err := s.cs.newRequest("deleteUcsManager", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteUcsManagerResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteUcsManagerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteUcsManagerResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteUcsManagerResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListUcsBladesParams struct {
	P map[string]interface{}
}

func (P *ListUcsBladesParams) toURLValues() url.Values {
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
	if v, found := P.P["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (P *ListUcsBladesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListUcsBladesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListUcsBladesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListUcsBladesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListUcsBladesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListUcsBladesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListUcsBladesParams) SetUcsmanagerid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ucsmanagerid"] = v
}

func (P *ListUcsBladesParams) GetUcsmanagerid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ucsmanagerid"].(string)
	return value, ok
}

// You should always use this function to get a new ListUcsBladesParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewListUcsBladesParams(ucsmanagerid string) *ListUcsBladesParams {
	P := &ListUcsBladesParams{}
	P.P = make(map[string]interface{})
	P.P["ucsmanagerid"] = ucsmanagerid
	return P
}

// List ucs blades
func (s *UCSService) ListUcsBlades(p *ListUcsBladesParams) (*ListUcsBladesResponse, error) {
	resp, err := s.cs.newRequest("listUcsBlades", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUcsBladesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUcsBladesResponse struct {
	Count     int         `json:"count"`
	UcsBlades []*UcsBlade `json:"ucsblade"`
}

type UcsBlade struct {
	Bladedn      string `json:"bladedn"`
	Hostid       string `json:"hostid"`
	Id           string `json:"id"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Profiledn    string `json:"profiledn"`
	Ucsmanagerid string `json:"ucsmanagerid"`
}

type ListUcsManagersParams struct {
	P map[string]interface{}
}

func (P *ListUcsManagersParams) toURLValues() url.Values {
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
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListUcsManagersParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListUcsManagersParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListUcsManagersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListUcsManagersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListUcsManagersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListUcsManagersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListUcsManagersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListUcsManagersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListUcsManagersParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListUcsManagersParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListUcsManagersParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewListUcsManagersParams() *ListUcsManagersParams {
	P := &ListUcsManagersParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UCSService) GetUcsManagerID(keyword string, opts ...OptionFunc) (string, int, error) {
	P := &ListUcsManagersParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListUcsManagers(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.UcsManagers[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.UcsManagers {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UCSService) GetUcsManagerByName(name string, opts ...OptionFunc) (*UcsManager, int, error) {
	id, count, err := s.GetUcsManagerID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetUcsManagerByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UCSService) GetUcsManagerByID(id string, opts ...OptionFunc) (*UcsManager, int, error) {
	P := &ListUcsManagersParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListUcsManagers(P)
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
		return l.UcsManagers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for UcsManager UUID: %s!", id)
}

// List ucs manager
func (s *UCSService) ListUcsManagers(p *ListUcsManagersParams) (*ListUcsManagersResponse, error) {
	resp, err := s.cs.newRequest("listUcsManagers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUcsManagersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUcsManagersResponse struct {
	Count       int           `json:"count"`
	UcsManagers []*UcsManager `json:"ucsmanager"`
}

type UcsManager struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Zoneid    string `json:"zoneid"`
}

type ListUcsProfilesParams struct {
	P map[string]interface{}
}

func (P *ListUcsProfilesParams) toURLValues() url.Values {
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
	if v, found := P.P["ucsmanagerid"]; found {
		u.Set("ucsmanagerid", v.(string))
	}
	return u
}

func (P *ListUcsProfilesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListUcsProfilesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListUcsProfilesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListUcsProfilesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListUcsProfilesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListUcsProfilesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListUcsProfilesParams) SetUcsmanagerid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ucsmanagerid"] = v
}

func (P *ListUcsProfilesParams) GetUcsmanagerid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ucsmanagerid"].(string)
	return value, ok
}

// You should always use this function to get a new ListUcsProfilesParams instance,
// as then you are sure you have configured all required params
func (s *UCSService) NewListUcsProfilesParams(ucsmanagerid string) *ListUcsProfilesParams {
	P := &ListUcsProfilesParams{}
	P.P = make(map[string]interface{})
	P.P["ucsmanagerid"] = ucsmanagerid
	return P
}

// List profile in ucs manager
func (s *UCSService) ListUcsProfiles(p *ListUcsProfilesParams) (*ListUcsProfilesResponse, error) {
	resp, err := s.cs.newRequest("listUcsProfiles", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUcsProfilesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUcsProfilesResponse struct {
	Count       int           `json:"count"`
	UcsProfiles []*UcsProfile `json:"ucsprofile"`
}

type UcsProfile struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Ucsdn     string `json:"ucsdn"`
}
