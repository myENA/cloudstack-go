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

type NetworkACLServiceIface interface {
	CreateNetworkACL(P *CreateNetworkACLParams) (*CreateNetworkACLResponse, error)
	NewCreateNetworkACLParams(protocol string) *CreateNetworkACLParams
	CreateNetworkACLList(P *CreateNetworkACLListParams) (*CreateNetworkACLListResponse, error)
	NewCreateNetworkACLListParams(name string, vpcid string) *CreateNetworkACLListParams
	DeleteNetworkACL(P *DeleteNetworkACLParams) (*DeleteNetworkACLResponse, error)
	NewDeleteNetworkACLParams(id string) *DeleteNetworkACLParams
	DeleteNetworkACLList(P *DeleteNetworkACLListParams) (*DeleteNetworkACLListResponse, error)
	NewDeleteNetworkACLListParams(id string) *DeleteNetworkACLListParams
	ListNetworkACLLists(P *ListNetworkACLListsParams) (*ListNetworkACLListsResponse, error)
	NewListNetworkACLListsParams() *ListNetworkACLListsParams
	GetNetworkACLListID(name string, opts ...OptionFunc) (string, int, error)
	GetNetworkACLListByName(name string, opts ...OptionFunc) (*NetworkACLList, int, error)
	GetNetworkACLListByID(id string, opts ...OptionFunc) (*NetworkACLList, int, error)
	ListNetworkACLs(P *ListNetworkACLsParams) (*ListNetworkACLsResponse, error)
	NewListNetworkACLsParams() *ListNetworkACLsParams
	GetNetworkACLByID(id string, opts ...OptionFunc) (*NetworkACL, int, error)
	ReplaceNetworkACLList(P *ReplaceNetworkACLListParams) (*ReplaceNetworkACLListResponse, error)
	NewReplaceNetworkACLListParams(aclid string) *ReplaceNetworkACLListParams
	UpdateNetworkACLItem(P *UpdateNetworkACLItemParams) (*UpdateNetworkACLItemResponse, error)
	NewUpdateNetworkACLItemParams(id string) *UpdateNetworkACLItemParams
	UpdateNetworkACLList(P *UpdateNetworkACLListParams) (*UpdateNetworkACLListResponse, error)
	NewUpdateNetworkACLListParams(id string) *UpdateNetworkACLListParams
}

type CreateNetworkACLParams struct {
	P map[string]interface{}
}

func (P *CreateNetworkACLParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := P.P["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := P.P["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := P.P["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := P.P["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := P.P["number"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("number", vv)
	}
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["reason"]; found {
		u.Set("reason", v.(string))
	}
	if v, found := P.P["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := P.P["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (P *CreateNetworkACLParams) SetAclid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["aclid"] = v
}

func (P *CreateNetworkACLParams) GetAclid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["aclid"].(string)
	return value, ok
}

func (P *CreateNetworkACLParams) SetAction(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["action"] = v
}

func (P *CreateNetworkACLParams) GetAction() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["action"].(string)
	return value, ok
}

func (P *CreateNetworkACLParams) SetCidrlist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *CreateNetworkACLParams) GetCidrlist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].([]string)
	return value, ok
}

func (P *CreateNetworkACLParams) SetEndport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endport"] = v
}

func (P *CreateNetworkACLParams) GetEndport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endport"].(int)
	return value, ok
}

func (P *CreateNetworkACLParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateNetworkACLParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateNetworkACLParams) SetIcmpcode(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmpcode"] = v
}

func (P *CreateNetworkACLParams) GetIcmpcode() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmpcode"].(int)
	return value, ok
}

func (P *CreateNetworkACLParams) SetIcmptype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmptype"] = v
}

func (P *CreateNetworkACLParams) GetIcmptype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmptype"].(int)
	return value, ok
}

func (P *CreateNetworkACLParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *CreateNetworkACLParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *CreateNetworkACLParams) SetNumber(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["number"] = v
}

func (P *CreateNetworkACLParams) GetNumber() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["number"].(int)
	return value, ok
}

func (P *CreateNetworkACLParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *CreateNetworkACLParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *CreateNetworkACLParams) SetReason(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["reason"] = v
}

func (P *CreateNetworkACLParams) GetReason() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["reason"].(string)
	return value, ok
}

func (P *CreateNetworkACLParams) SetStartport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startport"] = v
}

func (P *CreateNetworkACLParams) GetStartport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startport"].(int)
	return value, ok
}

func (P *CreateNetworkACLParams) SetTraffictype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["traffictype"] = v
}

func (P *CreateNetworkACLParams) GetTraffictype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new CreateNetworkACLParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewCreateNetworkACLParams(protocol string) *CreateNetworkACLParams {
	P := &CreateNetworkACLParams{}
	P.P = make(map[string]interface{})
	P.P["protocol"] = protocol
	return P
}

// Creates a ACL rule in the given network (the network has to belong to VPC)
func (s *NetworkACLService) CreateNetworkACL(p *CreateNetworkACLParams) (*CreateNetworkACLResponse, error) {
	resp, err := s.cs.newRequest("createNetworkACL", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateNetworkACLResponse
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

type CreateNetworkACLResponse struct {
	Aclid       string `json:"aclid"`
	Aclname     string `json:"aclname"`
	Action      string `json:"action"`
	Cidrlist    string `json:"cidrlist"`
	Endport     string `json:"endport"`
	Fordisplay  bool   `json:"fordisplay"`
	Icmpcode    int    `json:"icmpcode"`
	Icmptype    int    `json:"icmptype"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Number      int    `json:"number"`
	Protocol    string `json:"protocol"`
	Reason      string `json:"reason"`
	Startport   string `json:"startport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Traffictype string `json:"traffictype"`
}

type CreateNetworkACLListParams struct {
	P map[string]interface{}
}

func (P *CreateNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (P *CreateNetworkACLListParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateNetworkACLListParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateNetworkACLListParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateNetworkACLListParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateNetworkACLListParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateNetworkACLListParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateNetworkACLListParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *CreateNetworkACLListParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewCreateNetworkACLListParams(name string, vpcid string) *CreateNetworkACLListParams {
	P := &CreateNetworkACLListParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	P.P["vpcid"] = vpcid
	return P
}

// Creates a network ACL for the given VPC
func (s *NetworkACLService) CreateNetworkACLList(p *CreateNetworkACLListParams) (*CreateNetworkACLListResponse, error) {
	resp, err := s.cs.newRequest("createNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateNetworkACLListResponse
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

type CreateNetworkACLListResponse struct {
	Description string `json:"description"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Vpcid       string `json:"vpcid"`
}

type DeleteNetworkACLParams struct {
	P map[string]interface{}
}

func (P *DeleteNetworkACLParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteNetworkACLParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteNetworkACLParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkACLParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewDeleteNetworkACLParams(id string) *DeleteNetworkACLParams {
	P := &DeleteNetworkACLParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a network ACL
func (s *NetworkACLService) DeleteNetworkACL(p *DeleteNetworkACLParams) (*DeleteNetworkACLResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkACL", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkACLResponse
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

type DeleteNetworkACLResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteNetworkACLListParams struct {
	P map[string]interface{}
}

func (P *DeleteNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteNetworkACLListParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteNetworkACLListParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewDeleteNetworkACLListParams(id string) *DeleteNetworkACLListParams {
	P := &DeleteNetworkACLListParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a network ACL
func (s *NetworkACLService) DeleteNetworkACLList(p *DeleteNetworkACLListParams) (*DeleteNetworkACLListResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkACLListResponse
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

type DeleteNetworkACLListResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListNetworkACLListsParams struct {
	P map[string]interface{}
}

func (P *ListNetworkACLListsParams) toURLValues() url.Values {
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
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
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
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
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
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (P *ListNetworkACLListsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListNetworkACLListsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListNetworkACLListsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListNetworkACLListsParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListNetworkACLListsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListNetworkACLListsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNetworkACLListsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListNetworkACLListsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListNetworkACLListsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListNetworkACLListsParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNetworkACLListsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNetworkACLListsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListNetworkACLListsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListNetworkACLListsParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *ListNetworkACLListsParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworkACLListsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewListNetworkACLListsParams() *ListNetworkACLListsParams {
	P := &ListNetworkACLListsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkACLService) GetNetworkACLListID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListNetworkACLListsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetworkACLLists(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.NetworkACLLists[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NetworkACLLists {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkACLService) GetNetworkACLListByName(name string, opts ...OptionFunc) (*NetworkACLList, int, error) {
	id, count, err := s.GetNetworkACLListID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetNetworkACLListByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkACLService) GetNetworkACLListByID(id string, opts ...OptionFunc) (*NetworkACLList, int, error) {
	P := &ListNetworkACLListsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListNetworkACLLists(P)
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
		return l.NetworkACLLists[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for NetworkACLList UUID: %s!", id)
}

// Lists all network ACLs
func (s *NetworkACLService) ListNetworkACLLists(p *ListNetworkACLListsParams) (*ListNetworkACLListsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkACLLists", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkACLListsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkACLListsResponse struct {
	Count           int               `json:"count"`
	NetworkACLLists []*NetworkACLList `json:"networkacllist"`
}

type NetworkACLList struct {
	Description string `json:"description"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Vpcid       string `json:"vpcid"`
}

type ListNetworkACLsParams struct {
	P map[string]interface{}
}

func (P *ListNetworkACLsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := P.P["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
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
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
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
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (P *ListNetworkACLsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListNetworkACLsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListNetworkACLsParams) SetAclid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["aclid"] = v
}

func (P *ListNetworkACLsParams) GetAclid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["aclid"].(string)
	return value, ok
}

func (P *ListNetworkACLsParams) SetAction(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["action"] = v
}

func (P *ListNetworkACLsParams) GetAction() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["action"].(string)
	return value, ok
}

func (P *ListNetworkACLsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListNetworkACLsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListNetworkACLsParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListNetworkACLsParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListNetworkACLsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListNetworkACLsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListNetworkACLsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListNetworkACLsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListNetworkACLsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNetworkACLsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNetworkACLsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListNetworkACLsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListNetworkACLsParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListNetworkACLsParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListNetworkACLsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNetworkACLsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNetworkACLsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNetworkACLsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListNetworkACLsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListNetworkACLsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListNetworkACLsParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *ListNetworkACLsParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *ListNetworkACLsParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListNetworkACLsParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListNetworkACLsParams) SetTraffictype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["traffictype"] = v
}

func (P *ListNetworkACLsParams) GetTraffictype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworkACLsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewListNetworkACLsParams() *ListNetworkACLsParams {
	P := &ListNetworkACLsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkACLService) GetNetworkACLByID(id string, opts ...OptionFunc) (*NetworkACL, int, error) {
	P := &ListNetworkACLsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListNetworkACLs(P)
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
		return l.NetworkACLs[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for NetworkACL UUID: %s!", id)
}

// Lists all network ACL items
func (s *NetworkACLService) ListNetworkACLs(p *ListNetworkACLsParams) (*ListNetworkACLsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkACLs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkACLsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkACLsResponse struct {
	Count       int           `json:"count"`
	NetworkACLs []*NetworkACL `json:"networkacl"`
}

type NetworkACL struct {
	Aclid       string `json:"aclid"`
	Aclname     string `json:"aclname"`
	Action      string `json:"action"`
	Cidrlist    string `json:"cidrlist"`
	Endport     string `json:"endport"`
	Fordisplay  bool   `json:"fordisplay"`
	Icmpcode    int    `json:"icmpcode"`
	Icmptype    int    `json:"icmptype"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Number      int    `json:"number"`
	Protocol    string `json:"protocol"`
	Reason      string `json:"reason"`
	Startport   string `json:"startport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Traffictype string `json:"traffictype"`
}

type ReplaceNetworkACLListParams struct {
	P map[string]interface{}
}

func (P *ReplaceNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["aclid"]; found {
		u.Set("aclid", v.(string))
	}
	if v, found := P.P["gatewayid"]; found {
		u.Set("gatewayid", v.(string))
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	return u
}

func (P *ReplaceNetworkACLListParams) SetAclid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["aclid"] = v
}

func (P *ReplaceNetworkACLListParams) GetAclid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["aclid"].(string)
	return value, ok
}

func (P *ReplaceNetworkACLListParams) SetGatewayid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gatewayid"] = v
}

func (P *ReplaceNetworkACLListParams) GetGatewayid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gatewayid"].(string)
	return value, ok
}

func (P *ReplaceNetworkACLListParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ReplaceNetworkACLListParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

// You should always use this function to get a new ReplaceNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewReplaceNetworkACLListParams(aclid string) *ReplaceNetworkACLListParams {
	P := &ReplaceNetworkACLListParams{}
	P.P = make(map[string]interface{})
	P.P["aclid"] = aclid
	return P
}

// Replaces ACL associated with a network or private gateway
func (s *NetworkACLService) ReplaceNetworkACLList(p *ReplaceNetworkACLListParams) (*ReplaceNetworkACLListResponse, error) {
	resp, err := s.cs.newRequest("replaceNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReplaceNetworkACLListResponse
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

type ReplaceNetworkACLListResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateNetworkACLItemParams struct {
	P map[string]interface{}
}

func (P *UpdateNetworkACLItemParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := P.P["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := P.P["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["number"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("number", vv)
	}
	if v, found := P.P["partialupgrade"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("partialupgrade", vv)
	}
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["reason"]; found {
		u.Set("reason", v.(string))
	}
	if v, found := P.P["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := P.P["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (P *UpdateNetworkACLItemParams) SetAction(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["action"] = v
}

func (P *UpdateNetworkACLItemParams) GetAction() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["action"].(string)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetCidrlist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *UpdateNetworkACLItemParams) GetCidrlist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].([]string)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateNetworkACLItemParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetEndport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endport"] = v
}

func (P *UpdateNetworkACLItemParams) GetEndport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endport"].(int)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateNetworkACLItemParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetIcmpcode(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmpcode"] = v
}

func (P *UpdateNetworkACLItemParams) GetIcmpcode() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmpcode"].(int)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetIcmptype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmptype"] = v
}

func (P *UpdateNetworkACLItemParams) GetIcmptype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmptype"].(int)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateNetworkACLItemParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetNumber(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["number"] = v
}

func (P *UpdateNetworkACLItemParams) GetNumber() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["number"].(int)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetPartialupgrade(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["partialupgrade"] = v
}

func (P *UpdateNetworkACLItemParams) GetPartialupgrade() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["partialupgrade"].(bool)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *UpdateNetworkACLItemParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetReason(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["reason"] = v
}

func (P *UpdateNetworkACLItemParams) GetReason() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["reason"].(string)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetStartport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startport"] = v
}

func (P *UpdateNetworkACLItemParams) GetStartport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startport"].(int)
	return value, ok
}

func (P *UpdateNetworkACLItemParams) SetTraffictype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["traffictype"] = v
}

func (P *UpdateNetworkACLItemParams) GetTraffictype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateNetworkACLItemParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewUpdateNetworkACLItemParams(id string) *UpdateNetworkACLItemParams {
	P := &UpdateNetworkACLItemParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates ACL item with specified ID
func (s *NetworkACLService) UpdateNetworkACLItem(p *UpdateNetworkACLItemParams) (*UpdateNetworkACLItemResponse, error) {
	resp, err := s.cs.newRequest("updateNetworkACLItem", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkACLItemResponse
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

type UpdateNetworkACLItemResponse struct {
	Aclid       string `json:"aclid"`
	Aclname     string `json:"aclname"`
	Action      string `json:"action"`
	Cidrlist    string `json:"cidrlist"`
	Endport     string `json:"endport"`
	Fordisplay  bool   `json:"fordisplay"`
	Icmpcode    int    `json:"icmpcode"`
	Icmptype    int    `json:"icmptype"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Number      int    `json:"number"`
	Protocol    string `json:"protocol"`
	Reason      string `json:"reason"`
	Startport   string `json:"startport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Traffictype string `json:"traffictype"`
}

type UpdateNetworkACLListParams struct {
	P map[string]interface{}
}

func (P *UpdateNetworkACLListParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (P *UpdateNetworkACLListParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateNetworkACLListParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateNetworkACLListParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *UpdateNetworkACLListParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *UpdateNetworkACLListParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateNetworkACLListParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateNetworkACLListParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateNetworkACLListParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateNetworkACLListParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateNetworkACLListParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateNetworkACLListParams instance,
// as then you are sure you have configured all required params
func (s *NetworkACLService) NewUpdateNetworkACLListParams(id string) *UpdateNetworkACLListParams {
	P := &UpdateNetworkACLListParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates network ACL list
func (s *NetworkACLService) UpdateNetworkACLList(p *UpdateNetworkACLListParams) (*UpdateNetworkACLListResponse, error) {
	resp, err := s.cs.newRequest("updateNetworkACLList", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkACLListResponse
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

type UpdateNetworkACLListResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
