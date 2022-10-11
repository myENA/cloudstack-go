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

type SSHServiceIface interface {
	CreateSSHKeyPair(p *CreateSSHKeyPairParams) (*CreateSSHKeyPairResponse, error)
	NewCreateSSHKeyPairParams(name string) *CreateSSHKeyPairParams
	DeleteSSHKeyPair(p *DeleteSSHKeyPairParams) (*DeleteSSHKeyPairResponse, error)
	NewDeleteSSHKeyPairParams(name string) *DeleteSSHKeyPairParams
	ListSSHKeyPairs(p *ListSSHKeyPairsParams) (*ListSSHKeyPairsResponse, error)
	NewListSSHKeyPairsParams() *ListSSHKeyPairsParams
	GetSSHKeyPairID(name string, opts ...OptionFunc) (string, int, error)
	GetSSHKeyPairByName(name string, opts ...OptionFunc) (*SSHKeyPair, int, error)
	GetSSHKeyPairByID(id string, opts ...OptionFunc) (*SSHKeyPair, int, error)
	RegisterSSHKeyPair(p *RegisterSSHKeyPairParams) (*RegisterSSHKeyPairResponse, error)
	NewRegisterSSHKeyPairParams(name string, publickey string) *RegisterSSHKeyPairParams
	ResetSSHKeyForVirtualMachine(p *ResetSSHKeyForVirtualMachineParams) (*ResetSSHKeyForVirtualMachineResponse, error)
	NewResetSSHKeyForVirtualMachineParams(id string, keypair string) *ResetSSHKeyForVirtualMachineParams
}

type CreateSSHKeyPairParams struct {
	P map[string]interface{}
}

func (P *CreateSSHKeyPairParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *CreateSSHKeyPairParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateSSHKeyPairParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateSSHKeyPairParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateSSHKeyPairParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateSSHKeyPairParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateSSHKeyPairParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateSSHKeyPairParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateSSHKeyPairParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateSSHKeyPairParams instance,
// as then you are sure you have configured all required params
func (s *SSHService) NewCreateSSHKeyPairParams(name string) *CreateSSHKeyPairParams {
	P := &CreateSSHKeyPairParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	return P
}

// Create a new keypair and returns the private key
func (s *SSHService) CreateSSHKeyPair(p *CreateSSHKeyPairParams) (*CreateSSHKeyPairResponse, error) {
	resp, err := s.cs.newRequest("createSSHKeyPair", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateSSHKeyPairResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateSSHKeyPairResponse struct {
	Account        string `json:"account"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Fingerprint    string `json:"fingerprint"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Privatekey     string `json:"privatekey"`
}

type DeleteSSHKeyPairParams struct {
	P map[string]interface{}
}

func (P *DeleteSSHKeyPairParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *DeleteSSHKeyPairParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DeleteSSHKeyPairParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DeleteSSHKeyPairParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DeleteSSHKeyPairParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *DeleteSSHKeyPairParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *DeleteSSHKeyPairParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *DeleteSSHKeyPairParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *DeleteSSHKeyPairParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteSSHKeyPairParams instance,
// as then you are sure you have configured all required params
func (s *SSHService) NewDeleteSSHKeyPairParams(name string) *DeleteSSHKeyPairParams {
	P := &DeleteSSHKeyPairParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	return P
}

// Deletes a keypair by name
func (s *SSHService) DeleteSSHKeyPair(p *DeleteSSHKeyPairParams) (*DeleteSSHKeyPairResponse, error) {
	resp, err := s.cs.newRequest("deleteSSHKeyPair", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSSHKeyPairResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSSHKeyPairResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSSHKeyPairResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteSSHKeyPairResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListSSHKeyPairsParams struct {
	P map[string]interface{}
}

func (P *ListSSHKeyPairsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["fingerprint"]; found {
		u.Set("fingerprint", v.(string))
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *ListSSHKeyPairsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListSSHKeyPairsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListSSHKeyPairsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListSSHKeyPairsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListSSHKeyPairsParams) SetFingerprint(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fingerprint"] = v
}

func (P *ListSSHKeyPairsParams) GetFingerprint() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fingerprint"].(string)
	return value, ok
}

func (P *ListSSHKeyPairsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListSSHKeyPairsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListSSHKeyPairsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListSSHKeyPairsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListSSHKeyPairsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListSSHKeyPairsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListSSHKeyPairsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListSSHKeyPairsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListSSHKeyPairsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListSSHKeyPairsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListSSHKeyPairsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListSSHKeyPairsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListSSHKeyPairsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListSSHKeyPairsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListSSHKeyPairsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListSSHKeyPairsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSSHKeyPairsParams instance,
// as then you are sure you have configured all required params
func (s *SSHService) NewListSSHKeyPairsParams() *ListSSHKeyPairsParams {
	P := &ListSSHKeyPairsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SSHService) GetSSHKeyPairID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListSSHKeyPairsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSSHKeyPairs(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.SSHKeyPairs[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SSHKeyPairs {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SSHService) GetSSHKeyPairByName(name string, opts ...OptionFunc) (*SSHKeyPair, int, error) {
	id, count, err := s.GetSSHKeyPairID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSSHKeyPairByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SSHService) GetSSHKeyPairByID(id string, opts ...OptionFunc) (*SSHKeyPair, int, error) {
	P := &ListSSHKeyPairsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSSHKeyPairs(P)
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
		return l.SSHKeyPairs[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SSHKeyPair UUID: %s!", id)
}

// List registered keypairs
func (s *SSHService) ListSSHKeyPairs(p *ListSSHKeyPairsParams) (*ListSSHKeyPairsResponse, error) {
	resp, err := s.cs.newRequest("listSSHKeyPairs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSSHKeyPairsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSSHKeyPairsResponse struct {
	Count       int           `json:"count"`
	SSHKeyPairs []*SSHKeyPair `json:"sshkeypair"`
}

type SSHKeyPair struct {
	Account        string `json:"account"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Fingerprint    string `json:"fingerprint"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
}

type RegisterSSHKeyPairParams struct {
	P map[string]interface{}
}

func (P *RegisterSSHKeyPairParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["publickey"]; found {
		u.Set("publickey", v.(string))
	}
	return u
}

func (P *RegisterSSHKeyPairParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *RegisterSSHKeyPairParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *RegisterSSHKeyPairParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *RegisterSSHKeyPairParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *RegisterSSHKeyPairParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *RegisterSSHKeyPairParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *RegisterSSHKeyPairParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *RegisterSSHKeyPairParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *RegisterSSHKeyPairParams) SetPublickey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publickey"] = v
}

func (P *RegisterSSHKeyPairParams) GetPublickey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publickey"].(string)
	return value, ok
}

// You should always use this function to get a new RegisterSSHKeyPairParams instance,
// as then you are sure you have configured all required params
func (s *SSHService) NewRegisterSSHKeyPairParams(name string, publickey string) *RegisterSSHKeyPairParams {
	P := &RegisterSSHKeyPairParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	P.P["publickey"] = publickey
	return P
}

// Register a public key in a keypair under a certain name
func (s *SSHService) RegisterSSHKeyPair(p *RegisterSSHKeyPairParams) (*RegisterSSHKeyPairResponse, error) {
	resp, err := s.cs.newRequest("registerSSHKeyPair", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r RegisterSSHKeyPairResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterSSHKeyPairResponse struct {
	Account        string `json:"account"`
	Domain         string `json:"domain"`
	Domainid       string `json:"domainid"`
	Fingerprint    string `json:"fingerprint"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
}

type ResetSSHKeyForVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *ResetSSHKeyForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keypair"]; found {
		u.Set("keypair", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *ResetSSHKeyForVirtualMachineParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ResetSSHKeyForVirtualMachineParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ResetSSHKeyForVirtualMachineParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ResetSSHKeyForVirtualMachineParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ResetSSHKeyForVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ResetSSHKeyForVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ResetSSHKeyForVirtualMachineParams) SetKeypair(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keypair"] = v
}

func (P *ResetSSHKeyForVirtualMachineParams) GetKeypair() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keypair"].(string)
	return value, ok
}

func (P *ResetSSHKeyForVirtualMachineParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ResetSSHKeyForVirtualMachineParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ResetSSHKeyForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *SSHService) NewResetSSHKeyForVirtualMachineParams(id string, keypair string) *ResetSSHKeyForVirtualMachineParams {
	P := &ResetSSHKeyForVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["keypair"] = keypair
	return P
}

// Resets the SSH Key for virtual machine. The virtual machine must be in a "Stopped" state. [async]
func (s *SSHService) ResetSSHKeyForVirtualMachine(p *ResetSSHKeyForVirtualMachineParams) (*ResetSSHKeyForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("resetSSHKeyForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetSSHKeyForVirtualMachineResponse
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

type ResetSSHKeyForVirtualMachineResponse struct {
	Account               string                                              `json:"account"`
	Affinitygroup         []ResetSSHKeyForVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                              `json:"backupofferingid"`
	Backupofferingname    string                                              `json:"backupofferingname"`
	Bootmode              string                                              `json:"bootmode"`
	Boottype              string                                              `json:"boottype"`
	Cpunumber             int                                                 `json:"cpunumber"`
	Cpuspeed              int                                                 `json:"cpuspeed"`
	Cpuused               string                                              `json:"cpuused"`
	Created               string                                              `json:"created"`
	Details               map[string]string                                   `json:"details"`
	Diskioread            int64                                               `json:"diskioread"`
	Diskiowrite           int64                                               `json:"diskiowrite"`
	Diskkbsread           int64                                               `json:"diskkbsread"`
	Diskkbswrite          int64                                               `json:"diskkbswrite"`
	Diskofferingid        string                                              `json:"diskofferingid"`
	Diskofferingname      string                                              `json:"diskofferingname"`
	Displayname           string                                              `json:"displayname"`
	Displayvm             bool                                                `json:"displayvm"`
	Domain                string                                              `json:"domain"`
	Domainid              string                                              `json:"domainid"`
	Forvirtualnetwork     bool                                                `json:"forvirtualnetwork"`
	Group                 string                                              `json:"group"`
	Groupid               string                                              `json:"groupid"`
	Guestosid             string                                              `json:"guestosid"`
	Haenable              bool                                                `json:"haenable"`
	Hasannotations        bool                                                `json:"hasannotations"`
	Hostid                string                                              `json:"hostid"`
	Hostname              string                                              `json:"hostname"`
	Hypervisor            string                                              `json:"hypervisor"`
	Icon                  string                                              `json:"icon"`
	Id                    string                                              `json:"id"`
	Instancename          string                                              `json:"instancename"`
	Isdynamicallyscalable bool                                                `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                              `json:"isodisplaytext"`
	Isoid                 string                                              `json:"isoid"`
	Isoname               string                                              `json:"isoname"`
	JobID                 string                                              `json:"jobid"`
	Jobstatus             int                                                 `json:"jobstatus"`
	Keypair               string                                              `json:"keypair"`
	Lastupdated           string                                              `json:"lastupdated"`
	Memory                int                                                 `json:"memory"`
	Memoryintfreekbs      int64                                               `json:"memoryintfreekbs"`
	Memorykbs             int64                                               `json:"memorykbs"`
	Memorytargetkbs       int64                                               `json:"memorytargetkbs"`
	Name                  string                                              `json:"name"`
	Networkkbsread        int64                                               `json:"networkkbsread"`
	Networkkbswrite       int64                                               `json:"networkkbswrite"`
	Nic                   []Nic                                               `json:"nic"`
	Osdisplayname         string                                              `json:"osdisplayname"`
	Ostypeid              string                                              `json:"ostypeid"`
	Password              string                                              `json:"password"`
	Passwordenabled       bool                                                `json:"passwordenabled"`
	Pooltype              string                                              `json:"pooltype"`
	Project               string                                              `json:"project"`
	Projectid             string                                              `json:"projectid"`
	Publicip              string                                              `json:"publicip"`
	Publicipid            string                                              `json:"publicipid"`
	Readonlydetails       string                                              `json:"readonlydetails"`
	Receivedbytes         int64                                               `json:"receivedbytes"`
	Rootdeviceid          int64                                               `json:"rootdeviceid"`
	Rootdevicetype        string                                              `json:"rootdevicetype"`
	Securitygroup         []ResetSSHKeyForVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                               `json:"sentbytes"`
	Serviceofferingid     string                                              `json:"serviceofferingid"`
	Serviceofferingname   string                                              `json:"serviceofferingname"`
	Servicestate          string                                              `json:"servicestate"`
	State                 string                                              `json:"state"`
	Tags                  []Tags                                              `json:"tags"`
	Templatedisplaytext   string                                              `json:"templatedisplaytext"`
	Templateid            string                                              `json:"templateid"`
	Templatename          string                                              `json:"templatename"`
	Userid                string                                              `json:"userid"`
	Username              string                                              `json:"username"`
	Vgpu                  string                                              `json:"vgpu"`
	Zoneid                string                                              `json:"zoneid"`
	Zonename              string                                              `json:"zonename"`
}

type ResetSSHKeyForVirtualMachineResponseSecuritygroup struct {
	Account             string                                                  `json:"account"`
	Description         string                                                  `json:"description"`
	Domain              string                                                  `json:"domain"`
	Domainid            string                                                  `json:"domainid"`
	Egressrule          []ResetSSHKeyForVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                  `json:"id"`
	Ingressrule         []ResetSSHKeyForVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                  `json:"name"`
	Project             string                                                  `json:"project"`
	Projectid           string                                                  `json:"projectid"`
	Tags                []Tags                                                  `json:"tags"`
	Virtualmachinecount int                                                     `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                           `json:"virtualmachineids"`
}

type ResetSSHKeyForVirtualMachineResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type ResetSSHKeyForVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *ResetSSHKeyForVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias ResetSSHKeyForVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}
