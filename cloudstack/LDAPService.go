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
)

type LDAPServiceIface interface {
	AddLdapConfiguration(P *AddLdapConfigurationParams) (*AddLdapConfigurationResponse, error)
	NewAddLdapConfigurationParams(hostname string, port int) *AddLdapConfigurationParams
	DeleteLdapConfiguration(P *DeleteLdapConfigurationParams) (*DeleteLdapConfigurationResponse, error)
	NewDeleteLdapConfigurationParams(hostname string) *DeleteLdapConfigurationParams
	ImportLdapUsers(P *ImportLdapUsersParams) (*ImportLdapUsersResponse, error)
	NewImportLdapUsersParams() *ImportLdapUsersParams
	LdapConfig(P *LdapConfigParams) (*LdapConfigResponse, error)
	NewLdapConfigParams() *LdapConfigParams
	LdapCreateAccount(P *LdapCreateAccountParams) (*LdapCreateAccountResponse, error)
	NewLdapCreateAccountParams(username string) *LdapCreateAccountParams
	LdapRemove(P *LdapRemoveParams) (*LdapRemoveResponse, error)
	NewLdapRemoveParams() *LdapRemoveParams
	LinkDomainToLdap(P *LinkDomainToLdapParams) (*LinkDomainToLdapResponse, error)
	NewLinkDomainToLdapParams(accounttype int, domainid string, lDAPType string) *LinkDomainToLdapParams
	ListLdapConfigurations(P *ListLdapConfigurationsParams) (*ListLdapConfigurationsResponse, error)
	NewListLdapConfigurationsParams() *ListLdapConfigurationsParams
	ListLdapUsers(P *ListLdapUsersParams) (*ListLdapUsersResponse, error)
	NewListLdapUsersParams() *ListLdapUsersParams
	SearchLdap(P *SearchLdapParams) (*SearchLdapResponse, error)
	NewSearchLdapParams(query string) *SearchLdapParams
}

type AddLdapConfigurationParams struct {
	P map[string]interface{}
}

func (P *AddLdapConfigurationParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := P.P["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	return u
}

func (P *AddLdapConfigurationParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *AddLdapConfigurationParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *AddLdapConfigurationParams) SetHostname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostname"] = v
}

func (P *AddLdapConfigurationParams) GetHostname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostname"].(string)
	return value, ok
}

func (P *AddLdapConfigurationParams) SetPort(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["port"] = v
}

func (P *AddLdapConfigurationParams) GetPort() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["port"].(int)
	return value, ok
}

// You should always use this function to get a new AddLdapConfigurationParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewAddLdapConfigurationParams(hostname string, port int) *AddLdapConfigurationParams {
	P := &AddLdapConfigurationParams{}
	P.P = make(map[string]interface{})
	P.P["hostname"] = hostname
	P.P["port"] = port
	return P
}

// Add a new Ldap Configuration
func (s *LDAPService) AddLdapConfiguration(p *AddLdapConfigurationParams) (*AddLdapConfigurationResponse, error) {
	resp, err := s.cs.newRequest("addLdapConfiguration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddLdapConfigurationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddLdapConfigurationResponse struct {
	Domainid  string `json:"domainid"`
	Hostname  string `json:"hostname"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Port      int    `json:"port"`
}

type DeleteLdapConfigurationParams struct {
	P map[string]interface{}
}

func (P *DeleteLdapConfigurationParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := P.P["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	return u
}

func (P *DeleteLdapConfigurationParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DeleteLdapConfigurationParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *DeleteLdapConfigurationParams) SetHostname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostname"] = v
}

func (P *DeleteLdapConfigurationParams) GetHostname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostname"].(string)
	return value, ok
}

func (P *DeleteLdapConfigurationParams) SetPort(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["port"] = v
}

func (P *DeleteLdapConfigurationParams) GetPort() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["port"].(int)
	return value, ok
}

// You should always use this function to get a new DeleteLdapConfigurationParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewDeleteLdapConfigurationParams(hostname string) *DeleteLdapConfigurationParams {
	P := &DeleteLdapConfigurationParams{}
	P.P = make(map[string]interface{})
	P.P["hostname"] = hostname
	return P
}

// Remove an Ldap Configuration
func (s *LDAPService) DeleteLdapConfiguration(p *DeleteLdapConfigurationParams) (*DeleteLdapConfigurationResponse, error) {
	resp, err := s.cs.newRequest("deleteLdapConfiguration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLdapConfigurationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteLdapConfigurationResponse struct {
	Domainid  string `json:"domainid"`
	Hostname  string `json:"hostname"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Port      int    `json:"port"`
}

type ImportLdapUsersParams struct {
	P map[string]interface{}
}

func (P *ImportLdapUsersParams) toURLValues() url.Values {
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
	if v, found := P.P["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["group"]; found {
		u.Set("group", v.(string))
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
	if v, found := P.P["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := P.P["timezone"]; found {
		u.Set("timezone", v.(string))
	}
	return u
}

func (P *ImportLdapUsersParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ImportLdapUsersParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ImportLdapUsersParams) SetAccountdetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountdetails"] = v
}

func (P *ImportLdapUsersParams) GetAccountdetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountdetails"].(map[string]string)
	return value, ok
}

func (P *ImportLdapUsersParams) SetAccounttype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accounttype"] = v
}

func (P *ImportLdapUsersParams) GetAccounttype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accounttype"].(int)
	return value, ok
}

func (P *ImportLdapUsersParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ImportLdapUsersParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ImportLdapUsersParams) SetGroup(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["group"] = v
}

func (P *ImportLdapUsersParams) GetGroup() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["group"].(string)
	return value, ok
}

func (P *ImportLdapUsersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ImportLdapUsersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ImportLdapUsersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ImportLdapUsersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ImportLdapUsersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ImportLdapUsersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ImportLdapUsersParams) SetRoleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roleid"] = v
}

func (P *ImportLdapUsersParams) GetRoleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roleid"].(string)
	return value, ok
}

func (P *ImportLdapUsersParams) SetTimezone(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["timezone"] = v
}

func (P *ImportLdapUsersParams) GetTimezone() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["timezone"].(string)
	return value, ok
}

// You should always use this function to get a new ImportLdapUsersParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewImportLdapUsersParams() *ImportLdapUsersParams {
	P := &ImportLdapUsersParams{}
	P.P = make(map[string]interface{})
	return P
}

// Import LDAP users
func (s *LDAPService) ImportLdapUsers(p *ImportLdapUsersParams) (*ImportLdapUsersResponse, error) {
	resp, err := s.cs.newRequest("importLdapUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ImportLdapUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ImportLdapUsersResponse struct {
	Conflictingusersource string `json:"conflictingusersource"`
	Domain                string `json:"domain"`
	Email                 string `json:"email"`
	Firstname             string `json:"firstname"`
	JobID                 string `json:"jobid"`
	Jobstatus             int    `json:"jobstatus"`
	Lastname              string `json:"lastname"`
	Principal             string `json:"principal"`
	Username              string `json:"username"`
}

type LdapConfigParams struct {
	P map[string]interface{}
}

func (P *LdapConfigParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["binddn"]; found {
		u.Set("binddn", v.(string))
	}
	if v, found := P.P["bindpass"]; found {
		u.Set("bindpass", v.(string))
	}
	if v, found := P.P["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := P.P["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := P.P["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	if v, found := P.P["queryfilter"]; found {
		u.Set("queryfilter", v.(string))
	}
	if v, found := P.P["searchbase"]; found {
		u.Set("searchbase", v.(string))
	}
	if v, found := P.P["ssl"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ssl", vv)
	}
	if v, found := P.P["truststore"]; found {
		u.Set("truststore", v.(string))
	}
	if v, found := P.P["truststorepass"]; found {
		u.Set("truststorepass", v.(string))
	}
	return u
}

func (P *LdapConfigParams) SetBinddn(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["binddn"] = v
}

func (P *LdapConfigParams) GetBinddn() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["binddn"].(string)
	return value, ok
}

func (P *LdapConfigParams) SetBindpass(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bindpass"] = v
}

func (P *LdapConfigParams) GetBindpass() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bindpass"].(string)
	return value, ok
}

func (P *LdapConfigParams) SetHostname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostname"] = v
}

func (P *LdapConfigParams) GetHostname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostname"].(string)
	return value, ok
}

func (P *LdapConfigParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *LdapConfigParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *LdapConfigParams) SetPort(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["port"] = v
}

func (P *LdapConfigParams) GetPort() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["port"].(int)
	return value, ok
}

func (P *LdapConfigParams) SetQueryfilter(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["queryfilter"] = v
}

func (P *LdapConfigParams) GetQueryfilter() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["queryfilter"].(string)
	return value, ok
}

func (P *LdapConfigParams) SetSearchbase(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["searchbase"] = v
}

func (P *LdapConfigParams) GetSearchbase() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["searchbase"].(string)
	return value, ok
}

func (P *LdapConfigParams) SetSsl(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ssl"] = v
}

func (P *LdapConfigParams) GetSsl() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ssl"].(bool)
	return value, ok
}

func (P *LdapConfigParams) SetTruststore(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["truststore"] = v
}

func (P *LdapConfigParams) GetTruststore() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["truststore"].(string)
	return value, ok
}

func (P *LdapConfigParams) SetTruststorepass(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["truststorepass"] = v
}

func (P *LdapConfigParams) GetTruststorepass() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["truststorepass"].(string)
	return value, ok
}

// You should always use this function to get a new LdapConfigParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewLdapConfigParams() *LdapConfigParams {
	P := &LdapConfigParams{}
	P.P = make(map[string]interface{})
	return P
}

// (Deprecated, use addLdapConfiguration) Configure the LDAP context for this site.
func (s *LDAPService) LdapConfig(p *LdapConfigParams) (*LdapConfigResponse, error) {
	resp, err := s.cs.newRequest("ldapConfig", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LdapConfigResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LdapConfigResponse struct {
	Binddn      string `json:"binddn"`
	Bindpass    string `json:"bindpass"`
	Hostname    string `json:"hostname"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Port        string `json:"port"`
	Queryfilter string `json:"queryfilter"`
	Searchbase  string `json:"searchbase"`
	Ssl         string `json:"ssl"`
}

type LdapCreateAccountParams struct {
	P map[string]interface{}
}

func (P *LdapCreateAccountParams) toURLValues() url.Values {
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
	if v, found := P.P["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
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

func (P *LdapCreateAccountParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *LdapCreateAccountParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *LdapCreateAccountParams) SetAccountdetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountdetails"] = v
}

func (P *LdapCreateAccountParams) GetAccountdetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountdetails"].(map[string]string)
	return value, ok
}

func (P *LdapCreateAccountParams) SetAccountid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountid"] = v
}

func (P *LdapCreateAccountParams) GetAccountid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountid"].(string)
	return value, ok
}

func (P *LdapCreateAccountParams) SetAccounttype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accounttype"] = v
}

func (P *LdapCreateAccountParams) GetAccounttype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accounttype"].(int)
	return value, ok
}

func (P *LdapCreateAccountParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *LdapCreateAccountParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *LdapCreateAccountParams) SetNetworkdomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdomain"] = v
}

func (P *LdapCreateAccountParams) GetNetworkdomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdomain"].(string)
	return value, ok
}

func (P *LdapCreateAccountParams) SetRoleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roleid"] = v
}

func (P *LdapCreateAccountParams) GetRoleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roleid"].(string)
	return value, ok
}

func (P *LdapCreateAccountParams) SetTimezone(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["timezone"] = v
}

func (P *LdapCreateAccountParams) GetTimezone() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["timezone"].(string)
	return value, ok
}

func (P *LdapCreateAccountParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *LdapCreateAccountParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

func (P *LdapCreateAccountParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *LdapCreateAccountParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new LdapCreateAccountParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewLdapCreateAccountParams(username string) *LdapCreateAccountParams {
	P := &LdapCreateAccountParams{}
	P.P = make(map[string]interface{})
	P.P["username"] = username
	return P
}

// Creates an account from an LDAP user
func (s *LDAPService) LdapCreateAccount(p *LdapCreateAccountParams) (*LdapCreateAccountResponse, error) {
	resp, err := s.cs.newRequest("ldapCreateAccount", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LdapCreateAccountResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LdapCreateAccountResponse struct {
	Accountdetails            map[string]string               `json:"accountdetails"`
	Accounttype               int                             `json:"accounttype"`
	Cpuavailable              string                          `json:"cpuavailable"`
	Cpulimit                  string                          `json:"cpulimit"`
	Cputotal                  int64                           `json:"cputotal"`
	Created                   string                          `json:"created"`
	Defaultzoneid             string                          `json:"defaultzoneid"`
	Domain                    string                          `json:"domain"`
	Domainid                  string                          `json:"domainid"`
	Domainpath                string                          `json:"domainpath"`
	Groups                    []string                        `json:"groups"`
	Icon                      string                          `json:"icon"`
	Id                        string                          `json:"id"`
	Ipavailable               string                          `json:"ipavailable"`
	Iplimit                   string                          `json:"iplimit"`
	Iptotal                   int64                           `json:"iptotal"`
	Iscleanuprequired         bool                            `json:"iscleanuprequired"`
	Isdefault                 bool                            `json:"isdefault"`
	JobID                     string                          `json:"jobid"`
	Jobstatus                 int                             `json:"jobstatus"`
	Memoryavailable           string                          `json:"memoryavailable"`
	Memorylimit               string                          `json:"memorylimit"`
	Memorytotal               int64                           `json:"memorytotal"`
	Name                      string                          `json:"name"`
	Networkavailable          string                          `json:"networkavailable"`
	Networkdomain             string                          `json:"networkdomain"`
	Networklimit              string                          `json:"networklimit"`
	Networktotal              int64                           `json:"networktotal"`
	Primarystorageavailable   string                          `json:"primarystorageavailable"`
	Primarystoragelimit       string                          `json:"primarystoragelimit"`
	Primarystoragetotal       int64                           `json:"primarystoragetotal"`
	Projectavailable          string                          `json:"projectavailable"`
	Projectlimit              string                          `json:"projectlimit"`
	Projecttotal              int64                           `json:"projecttotal"`
	Receivedbytes             int64                           `json:"receivedbytes"`
	Roleid                    string                          `json:"roleid"`
	Rolename                  string                          `json:"rolename"`
	Roletype                  string                          `json:"roletype"`
	Secondarystorageavailable string                          `json:"secondarystorageavailable"`
	Secondarystoragelimit     string                          `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64                         `json:"secondarystoragetotal"`
	Sentbytes                 int64                           `json:"sentbytes"`
	Snapshotavailable         string                          `json:"snapshotavailable"`
	Snapshotlimit             string                          `json:"snapshotlimit"`
	Snapshottotal             int64                           `json:"snapshottotal"`
	State                     string                          `json:"state"`
	Templateavailable         string                          `json:"templateavailable"`
	Templatelimit             string                          `json:"templatelimit"`
	Templatetotal             int64                           `json:"templatetotal"`
	User                      []LdapCreateAccountResponseUser `json:"user"`
	Vmavailable               string                          `json:"vmavailable"`
	Vmlimit                   string                          `json:"vmlimit"`
	Vmrunning                 int                             `json:"vmrunning"`
	Vmstopped                 int                             `json:"vmstopped"`
	Vmtotal                   int64                           `json:"vmtotal"`
	Volumeavailable           string                          `json:"volumeavailable"`
	Volumelimit               string                          `json:"volumelimit"`
	Volumetotal               int64                           `json:"volumetotal"`
	Vpcavailable              string                          `json:"vpcavailable"`
	Vpclimit                  string                          `json:"vpclimit"`
	Vpctotal                  int64                           `json:"vpctotal"`
}

type LdapCreateAccountResponseUser struct {
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

type LdapRemoveParams struct {
	P map[string]interface{}
}

func (P *LdapRemoveParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	return u
}

// You should always use this function to get a new LdapRemoveParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewLdapRemoveParams() *LdapRemoveParams {
	P := &LdapRemoveParams{}
	P.P = make(map[string]interface{})
	return P
}

// (Deprecated , use deleteLdapConfiguration) Remove the LDAP context for this site.
func (s *LDAPService) LdapRemove(p *LdapRemoveParams) (*LdapRemoveResponse, error) {
	resp, err := s.cs.newRequest("ldapRemove", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LdapRemoveResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LdapRemoveResponse struct {
	Binddn      string `json:"binddn"`
	Bindpass    string `json:"bindpass"`
	Hostname    string `json:"hostname"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Port        string `json:"port"`
	Queryfilter string `json:"queryfilter"`
	Searchbase  string `json:"searchbase"`
	Ssl         string `json:"ssl"`
}

type LinkDomainToLdapParams struct {
	P map[string]interface{}
}

func (P *LinkDomainToLdapParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["accounttype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("accounttype", vv)
	}
	if v, found := P.P["admin"]; found {
		u.Set("admin", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["ldapdomain"]; found {
		u.Set("ldapdomain", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (P *LinkDomainToLdapParams) SetAccounttype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accounttype"] = v
}

func (P *LinkDomainToLdapParams) GetAccounttype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accounttype"].(int)
	return value, ok
}

func (P *LinkDomainToLdapParams) SetAdmin(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["admin"] = v
}

func (P *LinkDomainToLdapParams) GetAdmin() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["admin"].(string)
	return value, ok
}

func (P *LinkDomainToLdapParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *LinkDomainToLdapParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *LinkDomainToLdapParams) SetLdapdomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ldapdomain"] = v
}

func (P *LinkDomainToLdapParams) GetLdapdomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ldapdomain"].(string)
	return value, ok
}

func (P *LinkDomainToLdapParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *LinkDomainToLdapParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *LinkDomainToLdapParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *LinkDomainToLdapParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new LinkDomainToLdapParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewLinkDomainToLdapParams(accounttype int, domainid string, lDAPType string) *LinkDomainToLdapParams {
	P := &LinkDomainToLdapParams{}
	P.P = make(map[string]interface{})
	P.P["accounttype"] = accounttype
	P.P["domainid"] = domainid
	P.P["type"] = lDAPType
	return P
}

// link an existing cloudstack domain to group or OU in ldap
func (s *LDAPService) LinkDomainToLdap(p *LinkDomainToLdapParams) (*LinkDomainToLdapResponse, error) {
	resp, err := s.cs.newRequest("linkDomainToLdap", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r LinkDomainToLdapResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type LinkDomainToLdapResponse struct {
	Accountid   string `json:"accountid"`
	Accounttype int    `json:"accounttype"`
	Domainid    string `json:"domainid"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Ldapdomain  string `json:"ldapdomain"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type ListLdapConfigurationsParams struct {
	P map[string]interface{}
}

func (P *ListLdapConfigurationsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["hostname"]; found {
		u.Set("hostname", v.(string))
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
	if v, found := P.P["port"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("port", vv)
	}
	return u
}

func (P *ListLdapConfigurationsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListLdapConfigurationsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListLdapConfigurationsParams) SetHostname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostname"] = v
}

func (P *ListLdapConfigurationsParams) GetHostname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostname"].(string)
	return value, ok
}

func (P *ListLdapConfigurationsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListLdapConfigurationsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListLdapConfigurationsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListLdapConfigurationsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListLdapConfigurationsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListLdapConfigurationsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListLdapConfigurationsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListLdapConfigurationsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListLdapConfigurationsParams) SetPort(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["port"] = v
}

func (P *ListLdapConfigurationsParams) GetPort() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["port"].(int)
	return value, ok
}

// You should always use this function to get a new ListLdapConfigurationsParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewListLdapConfigurationsParams() *ListLdapConfigurationsParams {
	P := &ListLdapConfigurationsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists all LDAP configurations
func (s *LDAPService) ListLdapConfigurations(p *ListLdapConfigurationsParams) (*ListLdapConfigurationsResponse, error) {
	resp, err := s.cs.newRequest("listLdapConfigurations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLdapConfigurationsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLdapConfigurationsResponse struct {
	Count              int                  `json:"count"`
	LdapConfigurations []*LdapConfiguration `json:"ldapconfiguration"`
}

type LdapConfiguration struct {
	Domainid  string `json:"domainid"`
	Hostname  string `json:"hostname"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Port      int    `json:"port"`
}

type ListLdapUsersParams struct {
	P map[string]interface{}
}

func (P *ListLdapUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["listtype"]; found {
		u.Set("listtype", v.(string))
	}
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := P.P["userfilter"]; found {
		u.Set("userfilter", v.(string))
	}
	return u
}

func (P *ListLdapUsersParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListLdapUsersParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListLdapUsersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListLdapUsersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListLdapUsersParams) SetListtype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listtype"] = v
}

func (P *ListLdapUsersParams) GetListtype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listtype"].(string)
	return value, ok
}

func (P *ListLdapUsersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListLdapUsersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListLdapUsersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListLdapUsersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListLdapUsersParams) SetUserfilter(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userfilter"] = v
}

func (P *ListLdapUsersParams) GetUserfilter() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userfilter"].(string)
	return value, ok
}

// You should always use this function to get a new ListLdapUsersParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewListLdapUsersParams() *ListLdapUsersParams {
	P := &ListLdapUsersParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists LDAP Users according to the specifications from the user request.
func (s *LDAPService) ListLdapUsers(p *ListLdapUsersParams) (*ListLdapUsersResponse, error) {
	resp, err := s.cs.newRequest("listLdapUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLdapUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLdapUsersResponse struct {
	Count     int         `json:"count"`
	LdapUsers []*LdapUser `json:"ldapuser"`
}

type LdapUser struct {
	Conflictingusersource string `json:"conflictingusersource"`
	Domain                string `json:"domain"`
	Email                 string `json:"email"`
	Firstname             string `json:"firstname"`
	JobID                 string `json:"jobid"`
	Jobstatus             int    `json:"jobstatus"`
	Lastname              string `json:"lastname"`
	Principal             string `json:"principal"`
	Username              string `json:"username"`
}

type SearchLdapParams struct {
	P map[string]interface{}
}

func (P *SearchLdapParams) toURLValues() url.Values {
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
	if v, found := P.P["query"]; found {
		u.Set("query", v.(string))
	}
	return u
}

func (P *SearchLdapParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *SearchLdapParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *SearchLdapParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *SearchLdapParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *SearchLdapParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *SearchLdapParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *SearchLdapParams) SetQuery(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["query"] = v
}

func (P *SearchLdapParams) GetQuery() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["query"].(string)
	return value, ok
}

// You should always use this function to get a new SearchLdapParams instance,
// as then you are sure you have configured all required params
func (s *LDAPService) NewSearchLdapParams(query string) *SearchLdapParams {
	P := &SearchLdapParams{}
	P.P = make(map[string]interface{})
	P.P["query"] = query
	return P
}

// Searches LDAP based on the username attribute
func (s *LDAPService) SearchLdap(p *SearchLdapParams) (*SearchLdapResponse, error) {
	resp, err := s.cs.newRequest("searchLdap", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r SearchLdapResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type SearchLdapResponse struct {
	Conflictingusersource string `json:"conflictingusersource"`
	Domain                string `json:"domain"`
	Email                 string `json:"email"`
	Firstname             string `json:"firstname"`
	JobID                 string `json:"jobid"`
	Jobstatus             int    `json:"jobstatus"`
	Lastname              string `json:"lastname"`
	Principal             string `json:"principal"`
	Username              string `json:"username"`
}
