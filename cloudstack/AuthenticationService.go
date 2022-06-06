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

type AuthenticationServiceIface interface {
	Login(P *LoginParams) (*LoginResponse, error)
	NewLoginParams(password string, username string) *LoginParams
	Logout(P *LogoutParams) (*LogoutResponse, error)
	NewLogoutParams() *LogoutParams
}

type LoginParams struct {
	P map[string]interface{}
}

func (P *LoginParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := P.P["domainId"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("domainId", vv)
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *LoginParams) SetDomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domain"] = v
}

func (P *LoginParams) GetDomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domain"].(string)
	return value, ok
}

func (P *LoginParams) SetDomainId(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainId"] = v
}

func (P *LoginParams) GetDomainId() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainId"].(int64)
	return value, ok
}

func (P *LoginParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *LoginParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *LoginParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *LoginParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new LoginParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLoginParams(password string, username string) *LoginParams {
	P := &LoginParams{}
	P.P = make(map[string]interface{})
	P.P["password"] = password
	P.P["username"] = username
	return P
}

// Logs a user into the CloudStack. A successful login attempt will generate a JSESSIONID cookie value that can be passed in subsequent Query command calls until the "logout" command has been issued or the session has expired.
func (s *AuthenticationService) Login(p *LoginParams) (*LoginResponse, error) {
	resp, err := s.cs.newRequest("login", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LoginResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LoginResponse struct {
	Account        string `json:"account"`
	Domainid       string `json:"domainid"`
	Firstname      string `json:"firstname"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Lastname       string `json:"lastname"`
	Registered     string `json:"registered"`
	Sessionkey     string `json:"sessionkey"`
	Timeout        int    `json:"timeout"`
	Timezone       string `json:"timezone"`
	Timezoneoffset string `json:"timezoneoffset"`
	Type           string `json:"type"`
	Userid         string `json:"userid"`
	Username       string `json:"username"`
}

type LogoutParams struct {
	P map[string]interface{}
}

func (P *LogoutParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	return u
}

// You should always use this function to get a new LogoutParams instance,
// as then you are sure you have configured all required params
func (s *AuthenticationService) NewLogoutParams() *LogoutParams {
	P := &LogoutParams{}
	P.P = make(map[string]interface{})
	return P
}

// Logs out the user
func (s *AuthenticationService) Logout(p *LogoutParams) (*LogoutResponse, error) {
	resp, err := s.cs.newRequest("logout", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LogoutResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LogoutResponse struct {
	Description string `json:"description"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
}
