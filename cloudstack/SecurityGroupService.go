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

// Helper function for maintaining backwards compatibility
func convertAuthorizeSecurityGroupIngressResponse(b []byte) ([]byte, error) {
	var raw struct {
		Ingressrule []interface{} `json:"ingressrule"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	if len(raw.Ingressrule) != 1 {
		return b, nil
	}

	return json.Marshal(raw.Ingressrule[0])
}

// Helper function for maintaining backwards compatibility
func convertAuthorizeSecurityGroupEgressResponse(b []byte) ([]byte, error) {
	var raw struct {
		Egressrule []interface{} `json:"egressrule"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	if len(raw.Egressrule) != 1 {
		return b, nil
	}

	return json.Marshal(raw.Egressrule[0])
}

type SecurityGroupServiceIface interface {
	AuthorizeSecurityGroupEgress(P *AuthorizeSecurityGroupEgressParams) (*AuthorizeSecurityGroupEgressResponse, error)
	NewAuthorizeSecurityGroupEgressParams() *AuthorizeSecurityGroupEgressParams
	AuthorizeSecurityGroupIngress(P *AuthorizeSecurityGroupIngressParams) (*AuthorizeSecurityGroupIngressResponse, error)
	NewAuthorizeSecurityGroupIngressParams() *AuthorizeSecurityGroupIngressParams
	CreateSecurityGroup(P *CreateSecurityGroupParams) (*CreateSecurityGroupResponse, error)
	NewCreateSecurityGroupParams(name string) *CreateSecurityGroupParams
	DeleteSecurityGroup(P *DeleteSecurityGroupParams) (*DeleteSecurityGroupResponse, error)
	NewDeleteSecurityGroupParams() *DeleteSecurityGroupParams
	ListSecurityGroups(P *ListSecurityGroupsParams) (*ListSecurityGroupsResponse, error)
	NewListSecurityGroupsParams() *ListSecurityGroupsParams
	GetSecurityGroupID(keyword string, opts ...OptionFunc) (string, int, error)
	GetSecurityGroupByName(name string, opts ...OptionFunc) (*SecurityGroup, int, error)
	GetSecurityGroupByID(id string, opts ...OptionFunc) (*SecurityGroup, int, error)
	RevokeSecurityGroupEgress(P *RevokeSecurityGroupEgressParams) (*RevokeSecurityGroupEgressResponse, error)
	NewRevokeSecurityGroupEgressParams(id string) *RevokeSecurityGroupEgressParams
	RevokeSecurityGroupIngress(P *RevokeSecurityGroupIngressParams) (*RevokeSecurityGroupIngressResponse, error)
	NewRevokeSecurityGroupIngressParams(id string) *RevokeSecurityGroupIngressParams
}

type AuthorizeSecurityGroupEgressParams struct {
	P map[string]interface{}
}

func (P *AuthorizeSecurityGroupEgressParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := P.P["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := P.P["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["securitygroupid"]; found {
		u.Set("securitygroupid", v.(string))
	}
	if v, found := P.P["securitygroupname"]; found {
		u.Set("securitygroupname", v.(string))
	}
	if v, found := P.P["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := P.P["usersecuritygrouplist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].account", i), k)
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].group", i), m[k])
		}
	}
	return u
}

func (P *AuthorizeSecurityGroupEgressParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetCidrlist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetCidrlist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].([]string)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetEndport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endport"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetEndport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endport"].(int)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetIcmpcode(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmpcode"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetIcmpcode() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmpcode"].(int)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetIcmptype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmptype"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetIcmptype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmptype"].(int)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetSecuritygroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupid"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetSecuritygroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupid"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetSecuritygroupname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupname"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetSecuritygroupname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupname"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetStartport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startport"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetStartport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startport"].(int)
	return value, ok
}

func (P *AuthorizeSecurityGroupEgressParams) SetUsersecuritygrouplist(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["usersecuritygrouplist"] = v
}

func (P *AuthorizeSecurityGroupEgressParams) GetUsersecuritygrouplist() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["usersecuritygrouplist"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new AuthorizeSecurityGroupEgressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewAuthorizeSecurityGroupEgressParams() *AuthorizeSecurityGroupEgressParams {
	P := &AuthorizeSecurityGroupEgressParams{}
	P.P = make(map[string]interface{})
	return P
}

// Authorizes a particular egress rule for this security group
func (s *SecurityGroupService) AuthorizeSecurityGroupEgress(p *AuthorizeSecurityGroupEgressParams) (*AuthorizeSecurityGroupEgressResponse, error) {
	resp, err := s.cs.newRequest("authorizeSecurityGroupEgress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AuthorizeSecurityGroupEgressResponse
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

		b, err = convertAuthorizeSecurityGroupEgressResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AuthorizeSecurityGroupEgressResponse struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type AuthorizeSecurityGroupIngressParams struct {
	P map[string]interface{}
}

func (P *AuthorizeSecurityGroupIngressParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := P.P["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := P.P["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["securitygroupid"]; found {
		u.Set("securitygroupid", v.(string))
	}
	if v, found := P.P["securitygroupname"]; found {
		u.Set("securitygroupname", v.(string))
	}
	if v, found := P.P["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := P.P["usersecuritygrouplist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].account", i), k)
			u.Set(fmt.Sprintf("usersecuritygrouplist[%d].group", i), m[k])
		}
	}
	return u
}

func (P *AuthorizeSecurityGroupIngressParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetCidrlist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetCidrlist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].([]string)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetEndport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endport"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetEndport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endport"].(int)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetIcmpcode(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmpcode"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetIcmpcode() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmpcode"].(int)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetIcmptype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmptype"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetIcmptype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmptype"].(int)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetSecuritygroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupid"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetSecuritygroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupid"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetSecuritygroupname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupname"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetSecuritygroupname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupname"].(string)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetStartport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startport"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetStartport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startport"].(int)
	return value, ok
}

func (P *AuthorizeSecurityGroupIngressParams) SetUsersecuritygrouplist(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["usersecuritygrouplist"] = v
}

func (P *AuthorizeSecurityGroupIngressParams) GetUsersecuritygrouplist() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["usersecuritygrouplist"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new AuthorizeSecurityGroupIngressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewAuthorizeSecurityGroupIngressParams() *AuthorizeSecurityGroupIngressParams {
	P := &AuthorizeSecurityGroupIngressParams{}
	P.P = make(map[string]interface{})
	return P
}

// Authorizes a particular ingress rule for this security group
func (s *SecurityGroupService) AuthorizeSecurityGroupIngress(p *AuthorizeSecurityGroupIngressParams) (*AuthorizeSecurityGroupIngressResponse, error) {
	resp, err := s.cs.newRequest("authorizeSecurityGroupIngress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AuthorizeSecurityGroupIngressResponse
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

		b, err = convertAuthorizeSecurityGroupIngressResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AuthorizeSecurityGroupIngressResponse struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type CreateSecurityGroupParams struct {
	P map[string]interface{}
}

func (P *CreateSecurityGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
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

func (P *CreateSecurityGroupParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateSecurityGroupParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateSecurityGroupParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateSecurityGroupParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateSecurityGroupParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateSecurityGroupParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateSecurityGroupParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateSecurityGroupParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateSecurityGroupParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateSecurityGroupParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateSecurityGroupParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewCreateSecurityGroupParams(name string) *CreateSecurityGroupParams {
	P := &CreateSecurityGroupParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	return P
}

// Creates a security group
func (s *SecurityGroupService) CreateSecurityGroup(p *CreateSecurityGroupParams) (*CreateSecurityGroupResponse, error) {
	resp, err := s.cs.newRequest("createSecurityGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateSecurityGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateSecurityGroupResponse struct {
	Account             string                            `json:"account"`
	Description         string                            `json:"description"`
	Domain              string                            `json:"domain"`
	Domainid            string                            `json:"domainid"`
	Egressrule          []CreateSecurityGroupResponseRule `json:"egressrule"`
	Id                  string                            `json:"id"`
	Ingressrule         []CreateSecurityGroupResponseRule `json:"ingressrule"`
	JobID               string                            `json:"jobid"`
	Jobstatus           int                               `json:"jobstatus"`
	Name                string                            `json:"name"`
	Project             string                            `json:"project"`
	Projectid           string                            `json:"projectid"`
	Tags                []Tags                            `json:"tags"`
	Virtualmachinecount int                               `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                     `json:"virtualmachineids"`
}

type CreateSecurityGroupResponseRule struct {
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

type DeleteSecurityGroupParams struct {
	P map[string]interface{}
}

func (P *DeleteSecurityGroupParams) toURLValues() url.Values {
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
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *DeleteSecurityGroupParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DeleteSecurityGroupParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DeleteSecurityGroupParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DeleteSecurityGroupParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *DeleteSecurityGroupParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteSecurityGroupParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *DeleteSecurityGroupParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *DeleteSecurityGroupParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *DeleteSecurityGroupParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *DeleteSecurityGroupParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteSecurityGroupParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewDeleteSecurityGroupParams() *DeleteSecurityGroupParams {
	P := &DeleteSecurityGroupParams{}
	P.P = make(map[string]interface{})
	return P
}

// Deletes security group
func (s *SecurityGroupService) DeleteSecurityGroup(p *DeleteSecurityGroupParams) (*DeleteSecurityGroupResponse, error) {
	resp, err := s.cs.newRequest("deleteSecurityGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSecurityGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSecurityGroupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSecurityGroupResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteSecurityGroupResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListSecurityGroupsParams struct {
	P map[string]interface{}
}

func (P *ListSecurityGroupsParams) toURLValues() url.Values {
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["securitygroupname"]; found {
		u.Set("securitygroupname", v.(string))
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *ListSecurityGroupsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListSecurityGroupsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListSecurityGroupsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListSecurityGroupsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListSecurityGroupsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListSecurityGroupsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListSecurityGroupsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListSecurityGroupsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListSecurityGroupsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListSecurityGroupsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetSecuritygroupname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupname"] = v
}

func (P *ListSecurityGroupsParams) GetSecuritygroupname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupname"].(string)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListSecurityGroupsParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListSecurityGroupsParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *ListSecurityGroupsParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSecurityGroupsParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewListSecurityGroupsParams() *ListSecurityGroupsParams {
	P := &ListSecurityGroupsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SecurityGroupService) GetSecurityGroupID(keyword string, opts ...OptionFunc) (string, int, error) {
	P := &ListSecurityGroupsParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSecurityGroups(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.SecurityGroups[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SecurityGroups {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SecurityGroupService) GetSecurityGroupByName(name string, opts ...OptionFunc) (*SecurityGroup, int, error) {
	id, count, err := s.GetSecurityGroupID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSecurityGroupByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SecurityGroupService) GetSecurityGroupByID(id string, opts ...OptionFunc) (*SecurityGroup, int, error) {
	P := &ListSecurityGroupsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSecurityGroups(P)
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
		return l.SecurityGroups[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SecurityGroup UUID: %s!", id)
}

// Lists security groups
func (s *SecurityGroupService) ListSecurityGroups(p *ListSecurityGroupsParams) (*ListSecurityGroupsResponse, error) {
	resp, err := s.cs.newRequest("listSecurityGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSecurityGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSecurityGroupsResponse struct {
	Count          int              `json:"count"`
	SecurityGroups []*SecurityGroup `json:"securitygroup"`
}

type SecurityGroup struct {
	Account             string              `json:"account"`
	Description         string              `json:"description"`
	Domain              string              `json:"domain"`
	Domainid            string              `json:"domainid"`
	Egressrule          []SecurityGroupRule `json:"egressrule"`
	Id                  string              `json:"id"`
	Ingressrule         []SecurityGroupRule `json:"ingressrule"`
	JobID               string              `json:"jobid"`
	Jobstatus           int                 `json:"jobstatus"`
	Name                string              `json:"name"`
	Project             string              `json:"project"`
	Projectid           string              `json:"projectid"`
	Tags                []Tags              `json:"tags"`
	Virtualmachinecount int                 `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}       `json:"virtualmachineids"`
}

type SecurityGroupRule struct {
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

type RevokeSecurityGroupEgressParams struct {
	P map[string]interface{}
}

func (P *RevokeSecurityGroupEgressParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *RevokeSecurityGroupEgressParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RevokeSecurityGroupEgressParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RevokeSecurityGroupEgressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewRevokeSecurityGroupEgressParams(id string) *RevokeSecurityGroupEgressParams {
	P := &RevokeSecurityGroupEgressParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a particular egress rule from this security group
func (s *SecurityGroupService) RevokeSecurityGroupEgress(p *RevokeSecurityGroupEgressParams) (*RevokeSecurityGroupEgressResponse, error) {
	resp, err := s.cs.newRequest("revokeSecurityGroupEgress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevokeSecurityGroupEgressResponse
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

type RevokeSecurityGroupEgressResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RevokeSecurityGroupIngressParams struct {
	P map[string]interface{}
}

func (P *RevokeSecurityGroupIngressParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *RevokeSecurityGroupIngressParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RevokeSecurityGroupIngressParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RevokeSecurityGroupIngressParams instance,
// as then you are sure you have configured all required params
func (s *SecurityGroupService) NewRevokeSecurityGroupIngressParams(id string) *RevokeSecurityGroupIngressParams {
	P := &RevokeSecurityGroupIngressParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a particular ingress rule from this security group
func (s *SecurityGroupService) RevokeSecurityGroupIngress(p *RevokeSecurityGroupIngressParams) (*RevokeSecurityGroupIngressResponse, error) {
	resp, err := s.cs.newRequest("revokeSecurityGroupIngress", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevokeSecurityGroupIngressResponse
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

type RevokeSecurityGroupIngressResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
