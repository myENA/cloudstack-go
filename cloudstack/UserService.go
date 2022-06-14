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

type UserServiceIface interface {
	CreateUser(P *CreateUserParams) (*CreateUserResponse, error)
	NewCreateUserParams(account string, email string, firstname string, lastname string, password string, username string) *CreateUserParams
	DeleteUser(P *DeleteUserParams) (*DeleteUserResponse, error)
	NewDeleteUserParams(id string) *DeleteUserParams
	DisableUser(P *DisableUserParams) (*DisableUserResponse, error)
	NewDisableUserParams(id string) *DisableUserParams
	EnableUser(P *EnableUserParams) (*EnableUserResponse, error)
	NewEnableUserParams(id string) *EnableUserParams
	GetUser(P *GetUserParams) (*GetUserResponse, error)
	NewGetUserParams(userapikey string) *GetUserParams
	GetUserKeys(P *GetUserKeysParams) (*GetUserKeysResponse, error)
	NewGetUserKeysParams(id string) *GetUserKeysParams
	GetVirtualMachineUserData(P *GetVirtualMachineUserDataParams) (*GetVirtualMachineUserDataResponse, error)
	NewGetVirtualMachineUserDataParams(virtualmachineid string) *GetVirtualMachineUserDataParams
	ListUsers(P *ListUsersParams) (*ListUsersResponse, error)
	NewListUsersParams() *ListUsersParams
	GetUserByID(id string, opts ...OptionFunc) (*User, int, error)
	GetUserByUsernameAndDomain(username string, domainid string, opts ...OptionFunc) (*User, int, error)
	LockUser(P *LockUserParams) (*LockUserResponse, error)
	NewLockUserParams(id string) *LockUserParams
	RegisterUserKeys(P *RegisterUserKeysParams) (*RegisterUserKeysResponse, error)
	NewRegisterUserKeysParams(id string) *RegisterUserKeysParams
	UpdateUser(P *UpdateUserParams) (*UpdateUserResponse, error)
	NewUpdateUserParams(id string) *UpdateUserParams
}

type CreateUserParams struct {
	P map[string]interface{}
}

func (P *CreateUserParams) toURLValues() url.Values {
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
	if v, found := P.P["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := P.P["firstname"]; found {
		u.Set("firstname", v.(string))
	}
	if v, found := P.P["lastname"]; found {
		u.Set("lastname", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
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

func (P *CreateUserParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateUserParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateUserParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateUserParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateUserParams) SetEmail(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["email"] = v
}

func (P *CreateUserParams) GetEmail() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["email"].(string)
	return value, ok
}

func (P *CreateUserParams) SetFirstname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["firstname"] = v
}

func (P *CreateUserParams) GetFirstname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["firstname"].(string)
	return value, ok
}

func (P *CreateUserParams) SetLastname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lastname"] = v
}

func (P *CreateUserParams) GetLastname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lastname"].(string)
	return value, ok
}

func (P *CreateUserParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *CreateUserParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *CreateUserParams) SetTimezone(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["timezone"] = v
}

func (P *CreateUserParams) GetTimezone() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["timezone"].(string)
	return value, ok
}

func (P *CreateUserParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *CreateUserParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

func (P *CreateUserParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *CreateUserParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new CreateUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewCreateUserParams(account string, email string, firstname string, lastname string, password string, username string) *CreateUserParams {
	P := &CreateUserParams{}
	P.P = make(map[string]interface{})
	P.P["account"] = account
	P.P["email"] = email
	P.P["firstname"] = firstname
	P.P["lastname"] = lastname
	P.P["password"] = password
	P.P["username"] = username
	return P
}

// Creates a user for an account that already exists
func (s *UserService) CreateUser(p *CreateUserParams) (*CreateUserResponse, error) {
	resp, err := s.cs.newRequest("createUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateUserResponse struct {
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
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
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

type DeleteUserParams struct {
	P map[string]interface{}
}

func (P *DeleteUserParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteUserParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteUserParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewDeleteUserParams(id string) *DeleteUserParams {
	P := &DeleteUserParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a user for an account
func (s *UserService) DeleteUser(p *DeleteUserParams) (*DeleteUserResponse, error) {
	resp, err := s.cs.newRequest("deleteUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteUserResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteUserResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteUserResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DisableUserParams struct {
	P map[string]interface{}
}

func (P *DisableUserParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DisableUserParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DisableUserParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DisableUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewDisableUserParams(id string) *DisableUserParams {
	P := &DisableUserParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Disables a user account
func (s *UserService) DisableUser(p *DisableUserParams) (*DisableUserResponse, error) {
	resp, err := s.cs.newRequest("disableUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableUserResponse
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

type DisableUserResponse struct {
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
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
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

type EnableUserParams struct {
	P map[string]interface{}
}

func (P *EnableUserParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *EnableUserParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *EnableUserParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new EnableUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewEnableUserParams(id string) *EnableUserParams {
	P := &EnableUserParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Enables a user account
func (s *UserService) EnableUser(p *EnableUserParams) (*EnableUserResponse, error) {
	resp, err := s.cs.newRequest("enableUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r EnableUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type EnableUserResponse struct {
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
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
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

type GetUserParams struct {
	P map[string]interface{}
}

func (P *GetUserParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["userapikey"]; found {
		u.Set("userapikey", v.(string))
	}
	return u
}

func (P *GetUserParams) SetUserapikey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userapikey"] = v
}

func (P *GetUserParams) GetUserapikey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userapikey"].(string)
	return value, ok
}

// You should always use this function to get a new GetUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewGetUserParams(userapikey string) *GetUserParams {
	P := &GetUserParams{}
	P.P = make(map[string]interface{})
	P.P["userapikey"] = userapikey
	return P
}

// Find user account by API key
func (s *UserService) GetUser(p *GetUserParams) (*GetUserResponse, error) {
	resp, err := s.cs.newRequest("getUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetUserResponse struct {
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
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
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

type GetUserKeysParams struct {
	P map[string]interface{}
}

func (P *GetUserKeysParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *GetUserKeysParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *GetUserKeysParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new NewGetUserKeysParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewGetUserKeysParams(id string) *GetUserKeysParams {
	P := &GetUserKeysParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Find user keys
func (s *UserService) GetUserKeys(p *GetUserKeysParams) (*GetUserKeysResponse, error) {
	resp, err := s.cs.newRequest("getUserKeys", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UserKeysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r.UserKeys, nil
}

type UserKeysResponse struct {
	UserKeys GetUserKeysResponse `json:"userkeys"`
}

type GetUserKeysResponse struct {
	Apikey    string `json:"apikey"`
	Secretkey string `json:"secretkey"`
}

type GetVirtualMachineUserDataParams struct {
	P map[string]interface{}
}

func (P *GetVirtualMachineUserDataParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *GetVirtualMachineUserDataParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *GetVirtualMachineUserDataParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new GetVirtualMachineUserDataParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewGetVirtualMachineUserDataParams(virtualmachineid string) *GetVirtualMachineUserDataParams {
	P := &GetVirtualMachineUserDataParams{}
	P.P = make(map[string]interface{})
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Returns user data associated with the VM
func (s *UserService) GetVirtualMachineUserData(p *GetVirtualMachineUserDataParams) (*GetVirtualMachineUserDataResponse, error) {
	resp, err := s.cs.newRequest("getVirtualMachineUserData", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r GetVirtualMachineUserDataResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetVirtualMachineUserDataResponse struct {
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Userdata         string `json:"userdata"`
	Virtualmachineid string `json:"virtualmachineid"`
}

type ListUsersParams struct {
	P map[string]interface{}
}

func (P *ListUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["accounttype"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("accounttype", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
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
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *ListUsersParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListUsersParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListUsersParams) SetAccounttype(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accounttype"] = v
}

func (P *ListUsersParams) GetAccounttype() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accounttype"].(int64)
	return value, ok
}

func (P *ListUsersParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListUsersParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListUsersParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListUsersParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListUsersParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListUsersParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListUsersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListUsersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListUsersParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListUsersParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListUsersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListUsersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListUsersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListUsersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListUsersParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListUsersParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

func (P *ListUsersParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListUsersParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListUsersParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *ListUsersParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new ListUsersParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewListUsersParams() *ListUsersParams {
	P := &ListUsersParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UserService) GetUserByID(id string, opts ...OptionFunc) (*User, int, error) {
	P := &ListUsersParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListUsers(P)
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
		return l.Users[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for User UUID: %s!", id)
}

func (s *UserService) GetUserByUsernameAndDomain(username string, domainid string, opts ...OptionFunc) (*User, int, error) {
	P := &ListUsersParams{}
	P.P = make(map[string]interface{})

	P.P["username"] = username
	P.P["domainid"] = domainid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListUsers(P)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter username value=%s due to incorrect long value format, "+
				"or entity does not exist", username)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", username, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", username, l)
	}

	if l.Count == 1 {
		return l.Users[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Username: %s!", username)
}

// Lists user accounts
func (s *UserService) ListUsers(p *ListUsersParams) (*ListUsersResponse, error) {
	resp, err := s.cs.newRequest("listUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUsersResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"user"`
}

type User struct {
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
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
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

type LockUserParams struct {
	P map[string]interface{}
}

func (P *LockUserParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *LockUserParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *LockUserParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new LockUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewLockUserParams(id string) *LockUserParams {
	P := &LockUserParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Locks a user account
func (s *UserService) LockUser(p *LockUserParams) (*LockUserResponse, error) {
	resp, err := s.cs.newRequest("lockUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r LockUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LockUserResponse struct {
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
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
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

type RegisterUserKeysParams struct {
	P map[string]interface{}
}

func (P *RegisterUserKeysParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *RegisterUserKeysParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RegisterUserKeysParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RegisterUserKeysParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewRegisterUserKeysParams(id string) *RegisterUserKeysParams {
	P := &RegisterUserKeysParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// This command allows a user to register for the developer API, returning a secret key and an API key. This request is made through the integration API port, so it is a privileged command and must be made on behalf of a user. It is up to the implementer just how the username and password are entered, and then how that translates to an integration API request. Both secret key and API key should be returned to the user
func (s *UserService) RegisterUserKeys(p *RegisterUserKeysParams) (*RegisterUserKeysResponse, error) {
	resp, err := s.cs.newRequest("registerUserKeys", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r RegisterUserKeysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RegisterUserKeysResponse struct {
	Apikey    string `json:"apikey"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Secretkey string `json:"secretkey"`
}

type UpdateUserParams struct {
	P map[string]interface{}
}

func (P *UpdateUserParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["currentpassword"]; found {
		u.Set("currentpassword", v.(string))
	}
	if v, found := P.P["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := P.P["firstname"]; found {
		u.Set("firstname", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["lastname"]; found {
		u.Set("lastname", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	if v, found := P.P["userapikey"]; found {
		u.Set("userapikey", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := P.P["usersecretkey"]; found {
		u.Set("usersecretkey", v.(string))
	}
	return u
}

func (P *UpdateUserParams) SetCurrentpassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["currentpassword"] = v
}

func (P *UpdateUserParams) GetCurrentpassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["currentpassword"].(string)
	return value, ok
}

func (P *UpdateUserParams) SetEmail(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["email"] = v
}

func (P *UpdateUserParams) GetEmail() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["email"].(string)
	return value, ok
}

func (P *UpdateUserParams) SetFirstname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["firstname"] = v
}

func (P *UpdateUserParams) GetFirstname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["firstname"].(string)
	return value, ok
}

func (P *UpdateUserParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateUserParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateUserParams) SetLastname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lastname"] = v
}

func (P *UpdateUserParams) GetLastname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lastname"].(string)
	return value, ok
}

func (P *UpdateUserParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *UpdateUserParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *UpdateUserParams) SetTimezone(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["timezone"] = v
}

func (P *UpdateUserParams) GetTimezone() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["timezone"].(string)
	return value, ok
}

func (P *UpdateUserParams) SetUserapikey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userapikey"] = v
}

func (P *UpdateUserParams) GetUserapikey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userapikey"].(string)
	return value, ok
}

func (P *UpdateUserParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *UpdateUserParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

func (P *UpdateUserParams) SetUsersecretkey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["usersecretkey"] = v
}

func (P *UpdateUserParams) GetUsersecretkey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["usersecretkey"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateUserParams instance,
// as then you are sure you have configured all required params
func (s *UserService) NewUpdateUserParams(id string) *UpdateUserParams {
	P := &UpdateUserParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a user account
func (s *UserService) UpdateUser(p *UpdateUserParams) (*UpdateUserResponse, error) {
	resp, err := s.cs.newRequest("updateUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateUserResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateUserResponse struct {
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
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
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
