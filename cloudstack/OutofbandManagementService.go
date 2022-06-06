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

type OutofbandManagementServiceIface interface {
	ChangeOutOfBandManagementPassword(P *ChangeOutOfBandManagementPasswordParams) (*ChangeOutOfBandManagementPasswordResponse, error)
	NewChangeOutOfBandManagementPasswordParams(hostid string) *ChangeOutOfBandManagementPasswordParams
	ConfigureOutOfBandManagement(P *ConfigureOutOfBandManagementParams) (*OutOfBandManagementResponse, error)
	NewConfigureOutOfBandManagementParams(address string, driver string, hostid string, password string, port string, username string) *ConfigureOutOfBandManagementParams
	IssueOutOfBandManagementPowerAction(P *IssueOutOfBandManagementPowerActionParams) (*IssueOutOfBandManagementPowerActionResponse, error)
	NewIssueOutOfBandManagementPowerActionParams(action string, hostid string) *IssueOutOfBandManagementPowerActionParams
}

type ChangeOutOfBandManagementPasswordParams struct {
	P map[string]interface{}
}

func (P *ChangeOutOfBandManagementPasswordParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	return u
}

func (P *ChangeOutOfBandManagementPasswordParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ChangeOutOfBandManagementPasswordParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *ChangeOutOfBandManagementPasswordParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *ChangeOutOfBandManagementPasswordParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

// You should always use this function to get a new ChangeOutOfBandManagementPasswordParams instance,
// as then you are sure you have configured all required params
func (s *OutofbandManagementService) NewChangeOutOfBandManagementPasswordParams(hostid string) *ChangeOutOfBandManagementPasswordParams {
	P := &ChangeOutOfBandManagementPasswordParams{}
	P.P = make(map[string]interface{})
	P.P["hostid"] = hostid
	return P
}

// Changes out-of-band management interface password on the host and updates the interface configuration in CloudStack if the operation succeeds, else reverts the old password
func (s *OutofbandManagementService) ChangeOutOfBandManagementPassword(p *ChangeOutOfBandManagementPasswordParams) (*ChangeOutOfBandManagementPasswordResponse, error) {
	resp, err := s.cs.newRequest("changeOutOfBandManagementPassword", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeOutOfBandManagementPasswordResponse
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

type ChangeOutOfBandManagementPasswordResponse struct {
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}

type ConfigureOutOfBandManagementParams struct {
	P map[string]interface{}
}

func (P *ConfigureOutOfBandManagementParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["address"]; found {
		u.Set("address", v.(string))
	}
	if v, found := P.P["driver"]; found {
		u.Set("driver", v.(string))
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["port"]; found {
		u.Set("port", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *ConfigureOutOfBandManagementParams) SetAddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["address"] = v
}

func (P *ConfigureOutOfBandManagementParams) GetAddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["address"].(string)
	return value, ok
}

func (P *ConfigureOutOfBandManagementParams) SetDriver(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["driver"] = v
}

func (P *ConfigureOutOfBandManagementParams) GetDriver() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["driver"].(string)
	return value, ok
}

func (P *ConfigureOutOfBandManagementParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ConfigureOutOfBandManagementParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *ConfigureOutOfBandManagementParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *ConfigureOutOfBandManagementParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *ConfigureOutOfBandManagementParams) SetPort(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["port"] = v
}

func (P *ConfigureOutOfBandManagementParams) GetPort() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["port"].(string)
	return value, ok
}

func (P *ConfigureOutOfBandManagementParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *ConfigureOutOfBandManagementParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new ConfigureOutOfBandManagementParams instance,
// as then you are sure you have configured all required params
func (s *OutofbandManagementService) NewConfigureOutOfBandManagementParams(address string, driver string, hostid string, password string, port string, username string) *ConfigureOutOfBandManagementParams {
	P := &ConfigureOutOfBandManagementParams{}
	P.P = make(map[string]interface{})
	P.P["address"] = address
	P.P["driver"] = driver
	P.P["hostid"] = hostid
	P.P["password"] = password
	P.P["port"] = port
	P.P["username"] = username
	return P
}

// Configures a host's out-of-band management interface
func (s *OutofbandManagementService) ConfigureOutOfBandManagement(p *ConfigureOutOfBandManagementParams) (*OutOfBandManagementResponse, error) {
	resp, err := s.cs.newRequest("configureOutOfBandManagement", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r OutOfBandManagementResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type OutOfBandManagementResponse struct {
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}

type IssueOutOfBandManagementPowerActionParams struct {
	P map[string]interface{}
}

func (P *IssueOutOfBandManagementPowerActionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["timeout"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("timeout", vv)
	}
	return u
}

func (P *IssueOutOfBandManagementPowerActionParams) SetAction(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["action"] = v
}

func (P *IssueOutOfBandManagementPowerActionParams) GetAction() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["action"].(string)
	return value, ok
}

func (P *IssueOutOfBandManagementPowerActionParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *IssueOutOfBandManagementPowerActionParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *IssueOutOfBandManagementPowerActionParams) SetTimeout(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["timeout"] = v
}

func (P *IssueOutOfBandManagementPowerActionParams) GetTimeout() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["timeout"].(int64)
	return value, ok
}

// You should always use this function to get a new IssueOutOfBandManagementPowerActionParams instance,
// as then you are sure you have configured all required params
func (s *OutofbandManagementService) NewIssueOutOfBandManagementPowerActionParams(action string, hostid string) *IssueOutOfBandManagementPowerActionParams {
	P := &IssueOutOfBandManagementPowerActionParams{}
	P.P = make(map[string]interface{})
	P.P["action"] = action
	P.P["hostid"] = hostid
	return P
}

// Initiates the specified power action to the host's out-of-band management interface
func (s *OutofbandManagementService) IssueOutOfBandManagementPowerAction(p *IssueOutOfBandManagementPowerActionParams) (*IssueOutOfBandManagementPowerActionResponse, error) {
	resp, err := s.cs.newRequest("issueOutOfBandManagementPowerAction", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r IssueOutOfBandManagementPowerActionResponse
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

type IssueOutOfBandManagementPowerActionResponse struct {
	Action      string `json:"action"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Driver      string `json:"driver"`
	Enabled     bool   `json:"enabled"`
	Hostid      string `json:"hostid"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Powerstate  string `json:"powerstate"`
	Status      bool   `json:"status"`
	Username    string `json:"username"`
}
