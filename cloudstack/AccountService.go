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

type AccountServiceIface interface {
	CreateAccount(p *CreateAccountParams) (*CreateAccountResponse, error)
	NewCreateAccountParams(email string, firstname string, lastname string, password string, username string) *CreateAccountParams
	DeleteAccount(p *DeleteAccountParams) (*DeleteAccountResponse, error)
	NewDeleteAccountParams(id string) *DeleteAccountParams
	DisableAccount(p *DisableAccountParams) (*DisableAccountResponse, error)
	NewDisableAccountParams(lock bool) *DisableAccountParams
	EnableAccount(p *EnableAccountParams) (*EnableAccountResponse, error)
	NewEnableAccountParams() *EnableAccountParams
	GetSolidFireAccountId(p *GetSolidFireAccountIdParams) (*GetSolidFireAccountIdResponse, error)
	NewGetSolidFireAccountIdParams(accountid string, storageid string) *GetSolidFireAccountIdParams
	ListAccounts(p *ListAccountsParams) (*ListAccountsResponse, error)
	NewListAccountsParams() *ListAccountsParams
	GetAccountID(name string, opts ...OptionFunc) (string, int, error)
	GetAccountByName(name string, opts ...OptionFunc) (*Account, int, error)
	GetAccountByID(id string, opts ...OptionFunc) (*Account, int, error)
	ListProjectAccounts(p *ListProjectAccountsParams) (*ListProjectAccountsResponse, error)
	NewListProjectAccountsParams(projectid string) *ListProjectAccountsParams
	GetProjectAccountID(keyword string, projectid string, opts ...OptionFunc) (string, int, error)
	LockAccount(p *LockAccountParams) (*LockAccountResponse, error)
	NewLockAccountParams(account string, domainid string) *LockAccountParams
	MarkDefaultZoneForAccount(p *MarkDefaultZoneForAccountParams) (*MarkDefaultZoneForAccountResponse, error)
	NewMarkDefaultZoneForAccountParams(account string, domainid string, zoneid string) *MarkDefaultZoneForAccountParams
	UpdateAccount(p *UpdateAccountParams) (*UpdateAccountResponse, error)
	NewUpdateAccountParams() *UpdateAccountParams
}

type CreateAccountParams struct {
	P map[string]interface{}
}

func (P *CreateAccountParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["accountdetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("accountdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("accountdetails[%d].value", i), m[k])
		}
	}
	if v, found := P.P["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := P.P["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := P.P["firstname"]; found {
		u.Set("firstname", v.(string))
	}
	if v, found := P.P["lastname"]; found {
		u.Set("lastname", v.(string))
	}
	if v, found := P.P["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := P.P["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *CreateAccountParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateAccountParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetAccountdetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountdetails"] = v
}

func (P *CreateAccountParams) GetAccountdetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountdetails"].(map[string]string)
	return value, ok
}

func (P *CreateAccountParams) SetAccountid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountid"] = v
}

func (P *CreateAccountParams) GetAccountid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountid"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetAccounttype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accounttype"] = v
}

func (P *CreateAccountParams) GetAccounttype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accounttype"].(int)
	return value, ok
}

func (P *CreateAccountParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateAccountParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetEmail(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["email"] = v
}

func (P *CreateAccountParams) GetEmail() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["email"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetFirstname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["firstname"] = v
}

func (P *CreateAccountParams) GetFirstname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["firstname"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetLastname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lastname"] = v
}

func (P *CreateAccountParams) GetLastname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lastname"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetNetworkdomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdomain"] = v
}

func (P *CreateAccountParams) GetNetworkdomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdomain"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *CreateAccountParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetRoleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roleid"] = v
}

func (P *CreateAccountParams) GetRoleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roleid"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetTimezone(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["timezone"] = v
}

func (P *CreateAccountParams) GetTimezone() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["timezone"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *CreateAccountParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

func (P *CreateAccountParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *CreateAccountParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new CreateAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewCreateAccountParams(email string, firstname string, lastname string, password string, username string) *CreateAccountParams {
	P := &CreateAccountParams{}
	P.P = make(map[string]interface{})
	P.P["email"] = email
	P.P["firstname"] = firstname
	P.P["lastname"] = lastname
	P.P["password"] = password
	P.P["username"] = username
	return P
}

// Creates an account
func (s *AccountService) CreateAccount(p *CreateAccountParams) (*CreateAccountResponse, error) {
	resp, err := s.cs.newRequest("createAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateAccountResponse struct {
	Accountdetails            map[string]string           `json:"accountdetails"`
	Accounttype               int                         `json:"accounttype"`
	Cpuavailable              string                      `json:"cpuavailable"`
	Cpulimit                  string                      `json:"cpulimit"`
	Cputotal                  int64                       `json:"cputotal"`
	Created                   string                      `json:"created"`
	Defaultzoneid             string                      `json:"defaultzoneid"`
	Domain                    string                      `json:"domain"`
	Domainid                  string                      `json:"domainid"`
	Domainpath                string                      `json:"domainpath"`
	Groups                    []string                    `json:"groups"`
	Icon                      string                      `json:"icon"`
	Id                        string                      `json:"id"`
	Ipavailable               string                      `json:"ipavailable"`
	Iplimit                   string                      `json:"iplimit"`
	Iptotal                   int64                       `json:"iptotal"`
	Iscleanuprequired         bool                        `json:"iscleanuprequired"`
	Isdefault                 bool                        `json:"isdefault"`
	JobID                     string                      `json:"jobid"`
	Jobstatus                 int                         `json:"jobstatus"`
	Memoryavailable           string                      `json:"memoryavailable"`
	Memorylimit               string                      `json:"memorylimit"`
	Memorytotal               int64                       `json:"memorytotal"`
	Name                      string                      `json:"name"`
	Networkavailable          string                      `json:"networkavailable"`
	Networkdomain             string                      `json:"networkdomain"`
	Networklimit              string                      `json:"networklimit"`
	Networktotal              int64                       `json:"networktotal"`
	Primarystorageavailable   string                      `json:"primarystorageavailable"`
	Primarystoragelimit       string                      `json:"primarystoragelimit"`
	Primarystoragetotal       int64                       `json:"primarystoragetotal"`
	Projectavailable          string                      `json:"projectavailable"`
	Projectlimit              string                      `json:"projectlimit"`
	Projecttotal              int64                       `json:"projecttotal"`
	Receivedbytes             int64                       `json:"receivedbytes"`
	Roleid                    string                      `json:"roleid"`
	Rolename                  string                      `json:"rolename"`
	Roletype                  string                      `json:"roletype"`
	Secondarystorageavailable string                      `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                      `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                     `json:"secondarystoragetotal"`
	Sentbytes                 int64                       `json:"sentbytes"`
	Snapshotavailable         string                      `json:"snapshotavailable"`
	Snapshotlimit             string                      `json:"snapshotlimit"`
	Snapshottotal             int64                       `json:"snapshottotal"`
	State                     string                      `json:"state"`
	Templateavailable         string                      `json:"templateavailable"`
	Templatelimit             string                      `json:"templatelimit"`
	Templatetotal             int64                       `json:"templatetotal"`
	User                      []CreateAccountResponseUser `json:"user"`
	Vmavailable               string                      `json:"vmavailable"`
	Vmlimit                   string                      `json:"vmlimit"`
	Vmrunning                 int                         `json:"vmrunning"`
	Vmstopped                 int                         `json:"vmstopped"`
	Vmtotal                   int64                       `json:"vmtotal"`
	Volumeavailable           string                      `json:"volumeavailable"`
	Volumelimit               string                      `json:"volumelimit"`
	Volumetotal               int64                       `json:"volumetotal"`
	Vpcavailable              string                      `json:"vpcavailable"`
	Vpclimit                  string                      `json:"vpclimit"`
	Vpctotal                  int64                       `json:"vpctotal"`
}

type CreateAccountResponseUser struct {
	Account             string `json:"account"`
	Accountid           string `json:"accountid"`
	Accounttype         int    `json:"accounttype"`
	Apikey              string `json:"apikey"`
	Created             string `json:"created"`
	Domain              string `json:"domain"`
	Domainid            string `json:"domainid"`
	Email               string `json:"email"`
	Firstname           string `json:"firstname"`
	Icon                string `json:"icon"`
	Id                  string `json:"id"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain"`
	Isdefault           bool   `json:"isdefault"`
	Lastname            string `json:"lastname"`
	Roleid              string `json:"roleid"`
	Rolename            string `json:"rolename"`
	Roletype            string `json:"roletype"`
	Secretkey           string `json:"secretkey"`
	State               string `json:"state"`
	Timezone            string `json:"timezone"`
	Username            string `json:"username"`
	Usersource          string `json:"usersource"`
}

type DeleteAccountParams struct {
	P map[string]interface{}
}

func (P *DeleteAccountParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteAccountParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteAccountParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewDeleteAccountParams(id string) *DeleteAccountParams {
	P := &DeleteAccountParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a account, and all users associated with this account
func (s *AccountService) DeleteAccount(p *DeleteAccountParams) (*DeleteAccountResponse, error) {
	resp, err := s.cs.newRequest("deleteAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAccountResponse
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

type DeleteAccountResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DisableAccountParams struct {
	P map[string]interface{}
}

func (P *DisableAccountParams) toURLValues() url.Values {
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
	if v, found := P.P["lock"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("lock", vv)
	}
	return u
}

func (P *DisableAccountParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DisableAccountParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DisableAccountParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DisableAccountParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *DisableAccountParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DisableAccountParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *DisableAccountParams) SetLock(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lock"] = v
}

func (P *DisableAccountParams) GetLock() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lock"].(bool)
	return value, ok
}

// You should always use this function to get a new DisableAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewDisableAccountParams(lock bool) *DisableAccountParams {
	P := &DisableAccountParams{}
	P.P = make(map[string]interface{})
	P.P["lock"] = lock
	return P
}

// Disables an account
func (s *AccountService) DisableAccount(p *DisableAccountParams) (*DisableAccountResponse, error) {
	resp, err := s.cs.newRequest("disableAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableAccountResponse
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

type DisableAccountResponse struct {
	Accountdetails            map[string]string            `json:"accountdetails"`
	Accounttype               int                          `json:"accounttype"`
	Cpuavailable              string                       `json:"cpuavailable"`
	Cpulimit                  string                       `json:"cpulimit"`
	Cputotal                  int64                        `json:"cputotal"`
	Created                   string                       `json:"created"`
	Defaultzoneid             string                       `json:"defaultzoneid"`
	Domain                    string                       `json:"domain"`
	Domainid                  string                       `json:"domainid"`
	Domainpath                string                       `json:"domainpath"`
	Groups                    []string                     `json:"groups"`
	Icon                      string                       `json:"icon"`
	Id                        string                       `json:"id"`
	Ipavailable               string                       `json:"ipavailable"`
	Iplimit                   string                       `json:"iplimit"`
	Iptotal                   int64                        `json:"iptotal"`
	Iscleanuprequired         bool                         `json:"iscleanuprequired"`
	Isdefault                 bool                         `json:"isdefault"`
	JobID                     string                       `json:"jobid"`
	Jobstatus                 int                          `json:"jobstatus"`
	Memoryavailable           string                       `json:"memoryavailable"`
	Memorylimit               string                       `json:"memorylimit"`
	Memorytotal               int64                        `json:"memorytotal"`
	Name                      string                       `json:"name"`
	Networkavailable          string                       `json:"networkavailable"`
	Networkdomain             string                       `json:"networkdomain"`
	Networklimit              string                       `json:"networklimit"`
	Networktotal              int64                        `json:"networktotal"`
	Primarystorageavailable   string                       `json:"primarystorageavailable"`
	Primarystoragelimit       string                       `json:"primarystoragelimit"`
	Primarystoragetotal       int64                        `json:"primarystoragetotal"`
	Projectavailable          string                       `json:"projectavailable"`
	Projectlimit              string                       `json:"projectlimit"`
	Projecttotal              int64                        `json:"projecttotal"`
	Receivedbytes             int64                        `json:"receivedbytes"`
	Roleid                    string                       `json:"roleid"`
	Rolename                  string                       `json:"rolename"`
	Roletype                  string                       `json:"roletype"`
	Secondarystorageavailable string                       `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                       `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                      `json:"secondarystoragetotal"`
	Sentbytes                 int64                        `json:"sentbytes"`
	Snapshotavailable         string                       `json:"snapshotavailable"`
	Snapshotlimit             string                       `json:"snapshotlimit"`
	Snapshottotal             int64                        `json:"snapshottotal"`
	State                     string                       `json:"state"`
	Templateavailable         string                       `json:"templateavailable"`
	Templatelimit             string                       `json:"templatelimit"`
	Templatetotal             int64                        `json:"templatetotal"`
	User                      []DisableAccountResponseUser `json:"user"`
	Vmavailable               string                       `json:"vmavailable"`
	Vmlimit                   string                       `json:"vmlimit"`
	Vmrunning                 int                          `json:"vmrunning"`
	Vmstopped                 int                          `json:"vmstopped"`
	Vmtotal                   int64                        `json:"vmtotal"`
	Volumeavailable           string                       `json:"volumeavailable"`
	Volumelimit               string                       `json:"volumelimit"`
	Volumetotal               int64                        `json:"volumetotal"`
	Vpcavailable              string                       `json:"vpcavailable"`
	Vpclimit                  string                       `json:"vpclimit"`
	Vpctotal                  int64                        `json:"vpctotal"`
}

type DisableAccountResponseUser struct {
	Account             string `json:"account"`
	Accountid           string `json:"accountid"`
	Accounttype         int    `json:"accounttype"`
	Apikey              string `json:"apikey"`
	Created             string `json:"created"`
	Domain              string `json:"domain"`
	Domainid            string `json:"domainid"`
	Email               string `json:"email"`
	Firstname           string `json:"firstname"`
	Icon                string `json:"icon"`
	Id                  string `json:"id"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain"`
	Isdefault           bool   `json:"isdefault"`
	Lastname            string `json:"lastname"`
	Roleid              string `json:"roleid"`
	Rolename            string `json:"rolename"`
	Roletype            string `json:"roletype"`
	Secretkey           string `json:"secretkey"`
	State               string `json:"state"`
	Timezone            string `json:"timezone"`
	Username            string `json:"username"`
	Usersource          string `json:"usersource"`
}

type EnableAccountParams struct {
	P map[string]interface{}
}

func (P *EnableAccountParams) toURLValues() url.Values {
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
	return u
}

func (P *EnableAccountParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *EnableAccountParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *EnableAccountParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *EnableAccountParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *EnableAccountParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *EnableAccountParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new EnableAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewEnableAccountParams() *EnableAccountParams {
	P := &EnableAccountParams{}
	P.P = make(map[string]interface{})
	return P
}

// Enables an account
func (s *AccountService) EnableAccount(p *EnableAccountParams) (*EnableAccountResponse, error) {
	resp, err := s.cs.newRequest("enableAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type EnableAccountResponse struct {
	Accountdetails            map[string]string           `json:"accountdetails"`
	Accounttype               int                         `json:"accounttype"`
	Cpuavailable              string                      `json:"cpuavailable"`
	Cpulimit                  string                      `json:"cpulimit"`
	Cputotal                  int64                       `json:"cputotal"`
	Created                   string                      `json:"created"`
	Defaultzoneid             string                      `json:"defaultzoneid"`
	Domain                    string                      `json:"domain"`
	Domainid                  string                      `json:"domainid"`
	Domainpath                string                      `json:"domainpath"`
	Groups                    []string                    `json:"groups"`
	Icon                      string                      `json:"icon"`
	Id                        string                      `json:"id"`
	Ipavailable               string                      `json:"ipavailable"`
	Iplimit                   string                      `json:"iplimit"`
	Iptotal                   int64                       `json:"iptotal"`
	Iscleanuprequired         bool                        `json:"iscleanuprequired"`
	Isdefault                 bool                        `json:"isdefault"`
	JobID                     string                      `json:"jobid"`
	Jobstatus                 int                         `json:"jobstatus"`
	Memoryavailable           string                      `json:"memoryavailable"`
	Memorylimit               string                      `json:"memorylimit"`
	Memorytotal               int64                       `json:"memorytotal"`
	Name                      string                      `json:"name"`
	Networkavailable          string                      `json:"networkavailable"`
	Networkdomain             string                      `json:"networkdomain"`
	Networklimit              string                      `json:"networklimit"`
	Networktotal              int64                       `json:"networktotal"`
	Primarystorageavailable   string                      `json:"primarystorageavailable"`
	Primarystoragelimit       string                      `json:"primarystoragelimit"`
	Primarystoragetotal       int64                       `json:"primarystoragetotal"`
	Projectavailable          string                      `json:"projectavailable"`
	Projectlimit              string                      `json:"projectlimit"`
	Projecttotal              int64                       `json:"projecttotal"`
	Receivedbytes             int64                       `json:"receivedbytes"`
	Roleid                    string                      `json:"roleid"`
	Rolename                  string                      `json:"rolename"`
	Roletype                  string                      `json:"roletype"`
	Secondarystorageavailable string                      `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                      `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                     `json:"secondarystoragetotal"`
	Sentbytes                 int64                       `json:"sentbytes"`
	Snapshotavailable         string                      `json:"snapshotavailable"`
	Snapshotlimit             string                      `json:"snapshotlimit"`
	Snapshottotal             int64                       `json:"snapshottotal"`
	State                     string                      `json:"state"`
	Templateavailable         string                      `json:"templateavailable"`
	Templatelimit             string                      `json:"templatelimit"`
	Templatetotal             int64                       `json:"templatetotal"`
	User                      []EnableAccountResponseUser `json:"user"`
	Vmavailable               string                      `json:"vmavailable"`
	Vmlimit                   string                      `json:"vmlimit"`
	Vmrunning                 int                         `json:"vmrunning"`
	Vmstopped                 int                         `json:"vmstopped"`
	Vmtotal                   int64                       `json:"vmtotal"`
	Volumeavailable           string                      `json:"volumeavailable"`
	Volumelimit               string                      `json:"volumelimit"`
	Volumetotal               int64                       `json:"volumetotal"`
	Vpcavailable              string                      `json:"vpcavailable"`
	Vpclimit                  string                      `json:"vpclimit"`
	Vpctotal                  int64                       `json:"vpctotal"`
}

type EnableAccountResponseUser struct {
	Account             string `json:"account"`
	Accountid           string `json:"accountid"`
	Accounttype         int    `json:"accounttype"`
	Apikey              string `json:"apikey"`
	Created             string `json:"created"`
	Domain              string `json:"domain"`
	Domainid            string `json:"domainid"`
	Email               string `json:"email"`
	Firstname           string `json:"firstname"`
	Icon                string `json:"icon"`
	Id                  string `json:"id"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain"`
	Isdefault           bool   `json:"isdefault"`
	Lastname            string `json:"lastname"`
	Roleid              string `json:"roleid"`
	Rolename            string `json:"rolename"`
	Roletype            string `json:"roletype"`
	Secretkey           string `json:"secretkey"`
	State               string `json:"state"`
	Timezone            string `json:"timezone"`
	Username            string `json:"username"`
	Usersource          string `json:"usersource"`
}

type GetSolidFireAccountIdParams struct {
	P map[string]interface{}
}

func (P *GetSolidFireAccountIdParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	return u
}

func (P *GetSolidFireAccountIdParams) SetAccountid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountid"] = v
}

func (P *GetSolidFireAccountIdParams) GetAccountid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountid"].(string)
	return value, ok
}

func (P *GetSolidFireAccountIdParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *GetSolidFireAccountIdParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

// You should always use this function to get a new GetSolidFireAccountIdParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewGetSolidFireAccountIdParams(accountid string, storageid string) *GetSolidFireAccountIdParams {
	P := &GetSolidFireAccountIdParams{}
	P.P = make(map[string]interface{})
	P.P["accountid"] = accountid
	P.P["storageid"] = storageid
	return P
}

// Get SolidFire Account ID
func (s *AccountService) GetSolidFireAccountId(p *GetSolidFireAccountIdParams) (*GetSolidFireAccountIdResponse, error) {
	resp, err := s.cs.newRequest("getSolidFireAccountId", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetSolidFireAccountIdResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetSolidFireAccountIdResponse struct {
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	SolidFireAccountId int64  `json:"solidFireAccountId"`
}

type ListAccountsParams struct {
	P map[string]interface{}
}

func (P *ListAccountsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["accounttype"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("accounttype", vv)
	}
	if v, found := P.P["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["iscleanuprequired"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("iscleanuprequired", vv)
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
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (P *ListAccountsParams) SetAccounttype(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accounttype"] = v
}

func (P *ListAccountsParams) GetAccounttype() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accounttype"].(int64)
	return value, ok
}

func (P *ListAccountsParams) SetDetails(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ListAccountsParams) GetDetails() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].([]string)
	return value, ok
}

func (P *ListAccountsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListAccountsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListAccountsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListAccountsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListAccountsParams) SetIscleanuprequired(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iscleanuprequired"] = v
}

func (P *ListAccountsParams) GetIscleanuprequired() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iscleanuprequired"].(bool)
	return value, ok
}

func (P *ListAccountsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListAccountsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListAccountsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListAccountsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListAccountsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListAccountsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListAccountsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListAccountsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListAccountsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListAccountsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListAccountsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListAccountsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListAccountsParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListAccountsParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

func (P *ListAccountsParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListAccountsParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

// You should always use this function to get a new ListAccountsParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewListAccountsParams() *ListAccountsParams {
	P := &ListAccountsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetAccountID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListAccountsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListAccounts(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Accounts[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Accounts {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetAccountByName(name string, opts ...OptionFunc) (*Account, int, error) {
	id, count, err := s.GetAccountID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetAccountByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetAccountByID(id string, opts ...OptionFunc) (*Account, int, error) {
	P := &ListAccountsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAccounts(P)
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
		return l.Accounts[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Account UUID: %s!", id)
}

// Lists accounts and provides detailed account information for listed accounts
func (s *AccountService) ListAccounts(p *ListAccountsParams) (*ListAccountsResponse, error) {
	resp, err := s.cs.newRequest("listAccounts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAccountsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAccountsResponse struct {
	Count    int        `json:"count"`
	Accounts []*Account `json:"account"`
}

type Account struct {
	Accountdetails            map[string]string `json:"accountdetails"`
	Accounttype               int               `json:"accounttype"`
	Cpuavailable              string            `json:"cpuavailable"`
	Cpulimit                  string            `json:"cpulimit"`
	Cputotal                  int64             `json:"cputotal"`
	Created                   string            `json:"created"`
	Defaultzoneid             string            `json:"defaultzoneid"`
	Domain                    string            `json:"domain"`
	Domainid                  string            `json:"domainid"`
	Domainpath                string            `json:"domainpath"`
	Groups                    []string          `json:"groups"`
	Icon                      string            `json:"icon"`
	Id                        string            `json:"id"`
	Ipavailable               string            `json:"ipavailable"`
	Iplimit                   string            `json:"iplimit"`
	Iptotal                   int64             `json:"iptotal"`
	Iscleanuprequired         bool              `json:"iscleanuprequired"`
	Isdefault                 bool              `json:"isdefault"`
	JobID                     string            `json:"jobid"`
	Jobstatus                 int               `json:"jobstatus"`
	Memoryavailable           string            `json:"memoryavailable"`
	Memorylimit               string            `json:"memorylimit"`
	Memorytotal               int64             `json:"memorytotal"`
	Name                      string            `json:"name"`
	Networkavailable          string            `json:"networkavailable"`
	Networkdomain             string            `json:"networkdomain"`
	Networklimit              string            `json:"networklimit"`
	Networktotal              int64             `json:"networktotal"`
	Primarystorageavailable   string            `json:"primarystorageavailable"`
	Primarystoragelimit       string            `json:"primarystoragelimit"`
	Primarystoragetotal       int64             `json:"primarystoragetotal"`
	Projectavailable          string            `json:"projectavailable"`
	Projectlimit              string            `json:"projectlimit"`
	Projecttotal              int64             `json:"projecttotal"`
	Receivedbytes             int64             `json:"receivedbytes"`
	Roleid                    string            `json:"roleid"`
	Rolename                  string            `json:"rolename"`
	Roletype                  string            `json:"roletype"`
	Secondarystorageavailable string            `json:"secondarystorageavailable"`
	Secondarystoragelimit     string            `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64           `json:"secondarystoragetotal"`
	Sentbytes                 int64             `json:"sentbytes"`
	Snapshotavailable         string            `json:"snapshotavailable"`
	Snapshotlimit             string            `json:"snapshotlimit"`
	Snapshottotal             int64             `json:"snapshottotal"`
	State                     string            `json:"state"`
	Templateavailable         string            `json:"templateavailable"`
	Templatelimit             string            `json:"templatelimit"`
	Templatetotal             int64             `json:"templatetotal"`
	User                      []AccountUser     `json:"user"`
	Vmavailable               string            `json:"vmavailable"`
	Vmlimit                   string            `json:"vmlimit"`
	Vmrunning                 int               `json:"vmrunning"`
	Vmstopped                 int               `json:"vmstopped"`
	Vmtotal                   int64             `json:"vmtotal"`
	Volumeavailable           string            `json:"volumeavailable"`
	Volumelimit               string            `json:"volumelimit"`
	Volumetotal               int64             `json:"volumetotal"`
	Vpcavailable              string            `json:"vpcavailable"`
	Vpclimit                  string            `json:"vpclimit"`
	Vpctotal                  int64             `json:"vpctotal"`
}

type AccountUser struct {
	Account             string `json:"account"`
	Accountid           string `json:"accountid"`
	Accounttype         int    `json:"accounttype"`
	Apikey              string `json:"apikey"`
	Created             string `json:"created"`
	Domain              string `json:"domain"`
	Domainid            string `json:"domainid"`
	Email               string `json:"email"`
	Firstname           string `json:"firstname"`
	Icon                string `json:"icon"`
	Id                  string `json:"id"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain"`
	Isdefault           bool   `json:"isdefault"`
	Lastname            string `json:"lastname"`
	Roleid              string `json:"roleid"`
	Rolename            string `json:"rolename"`
	Roletype            string `json:"roletype"`
	Secretkey           string `json:"secretkey"`
	State               string `json:"state"`
	Timezone            string `json:"timezone"`
	Username            string `json:"username"`
	Usersource          string `json:"usersource"`
}

type ListProjectAccountsParams struct {
	P map[string]interface{}
}

func (P *ListProjectAccountsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	if v, found := P.P["role"]; found {
		u.Set("role", v.(string))
	}
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (P *ListProjectAccountsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListProjectAccountsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListProjectAccountsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListProjectAccountsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListProjectAccountsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListProjectAccountsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListProjectAccountsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListProjectAccountsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListProjectAccountsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListProjectAccountsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListProjectAccountsParams) SetProjectroleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectroleid"] = v
}

func (P *ListProjectAccountsParams) GetProjectroleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectroleid"].(string)
	return value, ok
}

func (P *ListProjectAccountsParams) SetRole(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["role"] = v
}

func (P *ListProjectAccountsParams) GetRole() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["role"].(string)
	return value, ok
}

func (P *ListProjectAccountsParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *ListProjectAccountsParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

// You should always use this function to get a new ListProjectAccountsParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewListProjectAccountsParams(projectid string) *ListProjectAccountsParams {
	P := &ListProjectAccountsParams{}
	P.P = make(map[string]interface{})
	P.P["projectid"] = projectid
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AccountService) GetProjectAccountID(keyword string, projectid string, opts ...OptionFunc) (string, int, error) {
	P := &ListProjectAccountsParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword
	P.P["projectid"] = projectid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListProjectAccounts(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.ProjectAccounts[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ProjectAccounts {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Lists project's accounts
func (s *AccountService) ListProjectAccounts(p *ListProjectAccountsParams) (*ListProjectAccountsResponse, error) {
	resp, err := s.cs.newRequest("listProjectAccounts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectAccountsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListProjectAccountsResponse struct {
	Count           int               `json:"count"`
	ProjectAccounts []*ProjectAccount `json:"projectaccount"`
}

type ProjectAccount struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      string              `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type Tags struct {
	Account      string `json:"account"`
	Customer     string `json:"customer"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Key          string `json:"key"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Resourceid   string `json:"resourceid"`
	Resourcetype string `json:"resourcetype"`
	Value        string `json:"value"`
}

type LockAccountParams struct {
	P map[string]interface{}
}

func (P *LockAccountParams) toURLValues() url.Values {
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
	return u
}

func (P *LockAccountParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *LockAccountParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *LockAccountParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *LockAccountParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

// You should always use this function to get a new LockAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewLockAccountParams(account string, domainid string) *LockAccountParams {
	P := &LockAccountParams{}
	P.P = make(map[string]interface{})
	P.P["account"] = account
	P.P["domainid"] = domainid
	return P
}

// This deprecated function used to locks an account. Look for the API DisableAccount instead
func (s *AccountService) LockAccount(p *LockAccountParams) (*LockAccountResponse, error) {
	resp, err := s.cs.newRequest("lockAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LockAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LockAccountResponse struct {
	Accountdetails            map[string]string         `json:"accountdetails"`
	Accounttype               int                       `json:"accounttype"`
	Cpuavailable              string                    `json:"cpuavailable"`
	Cpulimit                  string                    `json:"cpulimit"`
	Cputotal                  int64                     `json:"cputotal"`
	Created                   string                    `json:"created"`
	Defaultzoneid             string                    `json:"defaultzoneid"`
	Domain                    string                    `json:"domain"`
	Domainid                  string                    `json:"domainid"`
	Domainpath                string                    `json:"domainpath"`
	Groups                    []string                  `json:"groups"`
	Icon                      string                    `json:"icon"`
	Id                        string                    `json:"id"`
	Ipavailable               string                    `json:"ipavailable"`
	Iplimit                   string                    `json:"iplimit"`
	Iptotal                   int64                     `json:"iptotal"`
	Iscleanuprequired         bool                      `json:"iscleanuprequired"`
	Isdefault                 bool                      `json:"isdefault"`
	JobID                     string                    `json:"jobid"`
	Jobstatus                 int                       `json:"jobstatus"`
	Memoryavailable           string                    `json:"memoryavailable"`
	Memorylimit               string                    `json:"memorylimit"`
	Memorytotal               int64                     `json:"memorytotal"`
	Name                      string                    `json:"name"`
	Networkavailable          string                    `json:"networkavailable"`
	Networkdomain             string                    `json:"networkdomain"`
	Networklimit              string                    `json:"networklimit"`
	Networktotal              int64                     `json:"networktotal"`
	Primarystorageavailable   string                    `json:"primarystorageavailable"`
	Primarystoragelimit       string                    `json:"primarystoragelimit"`
	Primarystoragetotal       int64                     `json:"primarystoragetotal"`
	Projectavailable          string                    `json:"projectavailable"`
	Projectlimit              string                    `json:"projectlimit"`
	Projecttotal              int64                     `json:"projecttotal"`
	Receivedbytes             int64                     `json:"receivedbytes"`
	Roleid                    string                    `json:"roleid"`
	Rolename                  string                    `json:"rolename"`
	Roletype                  string                    `json:"roletype"`
	Secondarystorageavailable string                    `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                    `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                   `json:"secondarystoragetotal"`
	Sentbytes                 int64                     `json:"sentbytes"`
	Snapshotavailable         string                    `json:"snapshotavailable"`
	Snapshotlimit             string                    `json:"snapshotlimit"`
	Snapshottotal             int64                     `json:"snapshottotal"`
	State                     string                    `json:"state"`
	Templateavailable         string                    `json:"templateavailable"`
	Templatelimit             string                    `json:"templatelimit"`
	Templatetotal             int64                     `json:"templatetotal"`
	User                      []LockAccountResponseUser `json:"user"`
	Vmavailable               string                    `json:"vmavailable"`
	Vmlimit                   string                    `json:"vmlimit"`
	Vmrunning                 int                       `json:"vmrunning"`
	Vmstopped                 int                       `json:"vmstopped"`
	Vmtotal                   int64                     `json:"vmtotal"`
	Volumeavailable           string                    `json:"volumeavailable"`
	Volumelimit               string                    `json:"volumelimit"`
	Volumetotal               int64                     `json:"volumetotal"`
	Vpcavailable              string                    `json:"vpcavailable"`
	Vpclimit                  string                    `json:"vpclimit"`
	Vpctotal                  int64                     `json:"vpctotal"`
}

type LockAccountResponseUser struct {
	Account             string `json:"account"`
	Accountid           string `json:"accountid"`
	Accounttype         int    `json:"accounttype"`
	Apikey              string `json:"apikey"`
	Created             string `json:"created"`
	Domain              string `json:"domain"`
	Domainid            string `json:"domainid"`
	Email               string `json:"email"`
	Firstname           string `json:"firstname"`
	Icon                string `json:"icon"`
	Id                  string `json:"id"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain"`
	Isdefault           bool   `json:"isdefault"`
	Lastname            string `json:"lastname"`
	Roleid              string `json:"roleid"`
	Rolename            string `json:"rolename"`
	Roletype            string `json:"roletype"`
	Secretkey           string `json:"secretkey"`
	State               string `json:"state"`
	Timezone            string `json:"timezone"`
	Username            string `json:"username"`
	Usersource          string `json:"usersource"`
}

type MarkDefaultZoneForAccountParams struct {
	P map[string]interface{}
}

func (P *MarkDefaultZoneForAccountParams) toURLValues() url.Values {
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
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *MarkDefaultZoneForAccountParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *MarkDefaultZoneForAccountParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *MarkDefaultZoneForAccountParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *MarkDefaultZoneForAccountParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *MarkDefaultZoneForAccountParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *MarkDefaultZoneForAccountParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new MarkDefaultZoneForAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewMarkDefaultZoneForAccountParams(account string, domainid string, zoneid string) *MarkDefaultZoneForAccountParams {
	P := &MarkDefaultZoneForAccountParams{}
	P.P = make(map[string]interface{})
	P.P["account"] = account
	P.P["domainid"] = domainid
	P.P["zoneid"] = zoneid
	return P
}

// Marks a default zone for this account
func (s *AccountService) MarkDefaultZoneForAccount(p *MarkDefaultZoneForAccountParams) (*MarkDefaultZoneForAccountResponse, error) {
	resp, err := s.cs.newRequest("markDefaultZoneForAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MarkDefaultZoneForAccountResponse
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

type MarkDefaultZoneForAccountResponse struct {
	Accountdetails            map[string]string                       `json:"accountdetails"`
	Accounttype               int                                     `json:"accounttype"`
	Cpuavailable              string                                  `json:"cpuavailable"`
	Cpulimit                  string                                  `json:"cpulimit"`
	Cputotal                  int64                                   `json:"cputotal"`
	Created                   string                                  `json:"created"`
	Defaultzoneid             string                                  `json:"defaultzoneid"`
	Domain                    string                                  `json:"domain"`
	Domainid                  string                                  `json:"domainid"`
	Domainpath                string                                  `json:"domainpath"`
	Groups                    []string                                `json:"groups"`
	Icon                      string                                  `json:"icon"`
	Id                        string                                  `json:"id"`
	Ipavailable               string                                  `json:"ipavailable"`
	Iplimit                   string                                  `json:"iplimit"`
	Iptotal                   int64                                   `json:"iptotal"`
	Iscleanuprequired         bool                                    `json:"iscleanuprequired"`
	Isdefault                 bool                                    `json:"isdefault"`
	JobID                     string                                  `json:"jobid"`
	Jobstatus                 int                                     `json:"jobstatus"`
	Memoryavailable           string                                  `json:"memoryavailable"`
	Memorylimit               string                                  `json:"memorylimit"`
	Memorytotal               int64                                   `json:"memorytotal"`
	Name                      string                                  `json:"name"`
	Networkavailable          string                                  `json:"networkavailable"`
	Networkdomain             string                                  `json:"networkdomain"`
	Networklimit              string                                  `json:"networklimit"`
	Networktotal              int64                                   `json:"networktotal"`
	Primarystorageavailable   string                                  `json:"primarystorageavailable"`
	Primarystoragelimit       string                                  `json:"primarystoragelimit"`
	Primarystoragetotal       int64                                   `json:"primarystoragetotal"`
	Projectavailable          string                                  `json:"projectavailable"`
	Projectlimit              string                                  `json:"projectlimit"`
	Projecttotal              int64                                   `json:"projecttotal"`
	Receivedbytes             int64                                   `json:"receivedbytes"`
	Roleid                    string                                  `json:"roleid"`
	Rolename                  string                                  `json:"rolename"`
	Roletype                  string                                  `json:"roletype"`
	Secondarystorageavailable string                                  `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                                  `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                                 `json:"secondarystoragetotal"`
	Sentbytes                 int64                                   `json:"sentbytes"`
	Snapshotavailable         string                                  `json:"snapshotavailable"`
	Snapshotlimit             string                                  `json:"snapshotlimit"`
	Snapshottotal             int64                                   `json:"snapshottotal"`
	State                     string                                  `json:"state"`
	Templateavailable         string                                  `json:"templateavailable"`
	Templatelimit             string                                  `json:"templatelimit"`
	Templatetotal             int64                                   `json:"templatetotal"`
	User                      []MarkDefaultZoneForAccountResponseUser `json:"user"`
	Vmavailable               string                                  `json:"vmavailable"`
	Vmlimit                   string                                  `json:"vmlimit"`
	Vmrunning                 int                                     `json:"vmrunning"`
	Vmstopped                 int                                     `json:"vmstopped"`
	Vmtotal                   int64                                   `json:"vmtotal"`
	Volumeavailable           string                                  `json:"volumeavailable"`
	Volumelimit               string                                  `json:"volumelimit"`
	Volumetotal               int64                                   `json:"volumetotal"`
	Vpcavailable              string                                  `json:"vpcavailable"`
	Vpclimit                  string                                  `json:"vpclimit"`
	Vpctotal                  int64                                   `json:"vpctotal"`
}

type MarkDefaultZoneForAccountResponseUser struct {
	Account             string `json:"account"`
	Accountid           string `json:"accountid"`
	Accounttype         int    `json:"accounttype"`
	Apikey              string `json:"apikey"`
	Created             string `json:"created"`
	Domain              string `json:"domain"`
	Domainid            string `json:"domainid"`
	Email               string `json:"email"`
	Firstname           string `json:"firstname"`
	Icon                string `json:"icon"`
	Id                  string `json:"id"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain"`
	Isdefault           bool   `json:"isdefault"`
	Lastname            string `json:"lastname"`
	Roleid              string `json:"roleid"`
	Rolename            string `json:"rolename"`
	Roletype            string `json:"roletype"`
	Secretkey           string `json:"secretkey"`
	State               string `json:"state"`
	Timezone            string `json:"timezone"`
	Username            string `json:"username"`
	Usersource          string `json:"usersource"`
}

type UpdateAccountParams struct {
	P map[string]interface{}
}

func (P *UpdateAccountParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["accountdetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("accountdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("accountdetails[%d].value", i), m[k])
		}
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := P.P["newname"]; found {
		u.Set("newname", v.(string))
	}
	if v, found := P.P["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	return u
}

func (P *UpdateAccountParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *UpdateAccountParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *UpdateAccountParams) SetAccountdetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountdetails"] = v
}

func (P *UpdateAccountParams) GetAccountdetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountdetails"].(map[string]string)
	return value, ok
}

func (P *UpdateAccountParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UpdateAccountParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UpdateAccountParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateAccountParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateAccountParams) SetNetworkdomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdomain"] = v
}

func (P *UpdateAccountParams) GetNetworkdomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdomain"].(string)
	return value, ok
}

func (P *UpdateAccountParams) SetNewname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["newname"] = v
}

func (P *UpdateAccountParams) GetNewname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["newname"].(string)
	return value, ok
}

func (P *UpdateAccountParams) SetRoleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roleid"] = v
}

func (P *UpdateAccountParams) GetRoleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roleid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateAccountParams instance,
// as then you are sure you have configured all required params
func (s *AccountService) NewUpdateAccountParams() *UpdateAccountParams {
	P := &UpdateAccountParams{}
	P.P = make(map[string]interface{})
	return P
}

// Updates account information for the authenticated user
func (s *AccountService) UpdateAccount(p *UpdateAccountParams) (*UpdateAccountResponse, error) {
	resp, err := s.cs.newRequest("updateAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateAccountResponse struct {
	Accountdetails            map[string]string           `json:"accountdetails"`
	Accounttype               int                         `json:"accounttype"`
	Cpuavailable              string                      `json:"cpuavailable"`
	Cpulimit                  string                      `json:"cpulimit"`
	Cputotal                  int64                       `json:"cputotal"`
	Created                   string                      `json:"created"`
	Defaultzoneid             string                      `json:"defaultzoneid"`
	Domain                    string                      `json:"domain"`
	Domainid                  string                      `json:"domainid"`
	Domainpath                string                      `json:"domainpath"`
	Groups                    []string                    `json:"groups"`
	Icon                      string                      `json:"icon"`
	Id                        string                      `json:"id"`
	Ipavailable               string                      `json:"ipavailable"`
	Iplimit                   string                      `json:"iplimit"`
	Iptotal                   int64                       `json:"iptotal"`
	Iscleanuprequired         bool                        `json:"iscleanuprequired"`
	Isdefault                 bool                        `json:"isdefault"`
	JobID                     string                      `json:"jobid"`
	Jobstatus                 int                         `json:"jobstatus"`
	Memoryavailable           string                      `json:"memoryavailable"`
	Memorylimit               string                      `json:"memorylimit"`
	Memorytotal               int64                       `json:"memorytotal"`
	Name                      string                      `json:"name"`
	Networkavailable          string                      `json:"networkavailable"`
	Networkdomain             string                      `json:"networkdomain"`
	Networklimit              string                      `json:"networklimit"`
	Networktotal              int64                       `json:"networktotal"`
	Primarystorageavailable   string                      `json:"primarystorageavailable"`
	Primarystoragelimit       string                      `json:"primarystoragelimit"`
	Primarystoragetotal       int64                       `json:"primarystoragetotal"`
	Projectavailable          string                      `json:"projectavailable"`
	Projectlimit              string                      `json:"projectlimit"`
	Projecttotal              int64                       `json:"projecttotal"`
	Receivedbytes             int64                       `json:"receivedbytes"`
	Roleid                    string                      `json:"roleid"`
	Rolename                  string                      `json:"rolename"`
	Roletype                  string                      `json:"roletype"`
	Secondarystorageavailable string                      `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                      `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                     `json:"secondarystoragetotal"`
	Sentbytes                 int64                       `json:"sentbytes"`
	Snapshotavailable         string                      `json:"snapshotavailable"`
	Snapshotlimit             string                      `json:"snapshotlimit"`
	Snapshottotal             int64                       `json:"snapshottotal"`
	State                     string                      `json:"state"`
	Templateavailable         string                      `json:"templateavailable"`
	Templatelimit             string                      `json:"templatelimit"`
	Templatetotal             int64                       `json:"templatetotal"`
	User                      []UpdateAccountResponseUser `json:"user"`
	Vmavailable               string                      `json:"vmavailable"`
	Vmlimit                   string                      `json:"vmlimit"`
	Vmrunning                 int                         `json:"vmrunning"`
	Vmstopped                 int                         `json:"vmstopped"`
	Vmtotal                   int64                       `json:"vmtotal"`
	Volumeavailable           string                      `json:"volumeavailable"`
	Volumelimit               string                      `json:"volumelimit"`
	Volumetotal               int64                       `json:"volumetotal"`
	Vpcavailable              string                      `json:"vpcavailable"`
	Vpclimit                  string                      `json:"vpclimit"`
	Vpctotal                  int64                       `json:"vpctotal"`
}

type UpdateAccountResponseUser struct {
	Account             string `json:"account"`
	Accountid           string `json:"accountid"`
	Accounttype         int    `json:"accounttype"`
	Apikey              string `json:"apikey"`
	Created             string `json:"created"`
	Domain              string `json:"domain"`
	Domainid            string `json:"domainid"`
	Email               string `json:"email"`
	Firstname           string `json:"firstname"`
	Icon                string `json:"icon"`
	Id                  string `json:"id"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain"`
	Isdefault           bool   `json:"isdefault"`
	Lastname            string `json:"lastname"`
	Roleid              string `json:"roleid"`
	Rolename            string `json:"rolename"`
	Roletype            string `json:"roletype"`
	Secretkey           string `json:"secretkey"`
	State               string `json:"state"`
	Timezone            string `json:"timezone"`
	Username            string `json:"username"`
	Usersource          string `json:"usersource"`
}
