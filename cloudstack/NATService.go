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

type NATServiceIface interface {
	CreateIpForwardingRule(p *CreateIpForwardingRuleParams) (*CreateIpForwardingRuleResponse, error)
	NewCreateIpForwardingRuleParams(ipaddressid string, protocol string, startport int) *CreateIpForwardingRuleParams
	DeleteIpForwardingRule(p *DeleteIpForwardingRuleParams) (*DeleteIpForwardingRuleResponse, error)
	NewDeleteIpForwardingRuleParams(id string) *DeleteIpForwardingRuleParams
	DisableStaticNat(p *DisableStaticNatParams) (*DisableStaticNatResponse, error)
	NewDisableStaticNatParams(ipaddressid string) *DisableStaticNatParams
	EnableStaticNat(p *EnableStaticNatParams) (*EnableStaticNatResponse, error)
	NewEnableStaticNatParams(ipaddressid string, virtualmachineid string) *EnableStaticNatParams
	ListIpForwardingRules(p *ListIpForwardingRulesParams) (*ListIpForwardingRulesResponse, error)
	NewListIpForwardingRulesParams() *ListIpForwardingRulesParams
	GetIpForwardingRuleByID(id string, opts ...OptionFunc) (*IpForwardingRule, int, error)
}

type CreateIpForwardingRuleParams struct {
	P map[string]interface{}
}

func (P *CreateIpForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := P.P["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := P.P["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := P.P["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	return u
}

func (P *CreateIpForwardingRuleParams) SetCidrlist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *CreateIpForwardingRuleParams) GetCidrlist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].([]string)
	return value, ok
}

func (P *CreateIpForwardingRuleParams) SetEndport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endport"] = v
}

func (P *CreateIpForwardingRuleParams) GetEndport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endport"].(int)
	return value, ok
}

func (P *CreateIpForwardingRuleParams) SetIpaddressid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddressid"] = v
}

func (P *CreateIpForwardingRuleParams) GetIpaddressid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddressid"].(string)
	return value, ok
}

func (P *CreateIpForwardingRuleParams) SetOpenfirewall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["openfirewall"] = v
}

func (P *CreateIpForwardingRuleParams) GetOpenfirewall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["openfirewall"].(bool)
	return value, ok
}

func (P *CreateIpForwardingRuleParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *CreateIpForwardingRuleParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *CreateIpForwardingRuleParams) SetStartport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startport"] = v
}

func (P *CreateIpForwardingRuleParams) GetStartport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startport"].(int)
	return value, ok
}

// You should always use this function to get a new CreateIpForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *NATService) NewCreateIpForwardingRuleParams(ipaddressid string, protocol string, startport int) *CreateIpForwardingRuleParams {
	P := &CreateIpForwardingRuleParams{}
	P.P = make(map[string]interface{})
	P.P["ipaddressid"] = ipaddressid
	P.P["protocol"] = protocol
	P.P["startport"] = startport
	return P
}

// Creates an IP forwarding rule
func (s *NATService) CreateIpForwardingRule(p *CreateIpForwardingRuleParams) (*CreateIpForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("createIpForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateIpForwardingRuleResponse
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

type CreateIpForwardingRuleResponse struct {
	Cidrlist                  string `json:"cidrlist"`
	Fordisplay                bool   `json:"fordisplay"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Ipaddressid               string `json:"ipaddressid"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Privateendport            string `json:"privateendport"`
	Privateport               string `json:"privateport"`
	Protocol                  string `json:"protocol"`
	Publicendport             string `json:"publicendport"`
	Publicport                string `json:"publicport"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vmguestip                 string `json:"vmguestip"`
}

type DeleteIpForwardingRuleParams struct {
	P map[string]interface{}
}

func (P *DeleteIpForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteIpForwardingRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteIpForwardingRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteIpForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *NATService) NewDeleteIpForwardingRuleParams(id string) *DeleteIpForwardingRuleParams {
	P := &DeleteIpForwardingRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes an IP forwarding rule
func (s *NATService) DeleteIpForwardingRule(p *DeleteIpForwardingRuleParams) (*DeleteIpForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteIpForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteIpForwardingRuleResponse
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

type DeleteIpForwardingRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DisableStaticNatParams struct {
	P map[string]interface{}
}

func (P *DisableStaticNatParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	return u
}

func (P *DisableStaticNatParams) SetIpaddressid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddressid"] = v
}

func (P *DisableStaticNatParams) GetIpaddressid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddressid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableStaticNatParams instance,
// as then you are sure you have configured all required params
func (s *NATService) NewDisableStaticNatParams(ipaddressid string) *DisableStaticNatParams {
	P := &DisableStaticNatParams{}
	P.P = make(map[string]interface{})
	P.P["ipaddressid"] = ipaddressid
	return P
}

// Disables static rule for given IP address
func (s *NATService) DisableStaticNat(p *DisableStaticNatParams) (*DisableStaticNatResponse, error) {
	resp, err := s.cs.newRequest("disableStaticNat", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableStaticNatResponse
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

type DisableStaticNatResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type EnableStaticNatParams struct {
	P map[string]interface{}
}

func (P *EnableStaticNatParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := P.P["vmguestip"]; found {
		u.Set("vmguestip", v.(string))
	}
	return u
}

func (P *EnableStaticNatParams) SetIpaddressid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddressid"] = v
}

func (P *EnableStaticNatParams) GetIpaddressid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddressid"].(string)
	return value, ok
}

func (P *EnableStaticNatParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *EnableStaticNatParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *EnableStaticNatParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *EnableStaticNatParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *EnableStaticNatParams) SetVmguestip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmguestip"] = v
}

func (P *EnableStaticNatParams) GetVmguestip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmguestip"].(string)
	return value, ok
}

// You should always use this function to get a new EnableStaticNatParams instance,
// as then you are sure you have configured all required params
func (s *NATService) NewEnableStaticNatParams(ipaddressid string, virtualmachineid string) *EnableStaticNatParams {
	P := &EnableStaticNatParams{}
	P.P = make(map[string]interface{})
	P.P["ipaddressid"] = ipaddressid
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Enables static NAT for given IP address
func (s *NATService) EnableStaticNat(p *EnableStaticNatParams) (*EnableStaticNatResponse, error) {
	resp, err := s.cs.newRequest("enableStaticNat", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableStaticNatResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type EnableStaticNatResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *EnableStaticNatResponse) UnmarshalJSON(b []byte) error {
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

	type alias EnableStaticNatResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListIpForwardingRulesParams struct {
	P map[string]interface{}
}

func (P *ListIpForwardingRulesParams) toURLValues() url.Values {
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
	if v, found := P.P["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
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
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *ListIpForwardingRulesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListIpForwardingRulesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListIpForwardingRulesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListIpForwardingRulesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListIpForwardingRulesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListIpForwardingRulesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListIpForwardingRulesParams) SetIpaddressid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddressid"] = v
}

func (P *ListIpForwardingRulesParams) GetIpaddressid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddressid"].(string)
	return value, ok
}

func (P *ListIpForwardingRulesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListIpForwardingRulesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListIpForwardingRulesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListIpForwardingRulesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListIpForwardingRulesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListIpForwardingRulesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListIpForwardingRulesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListIpForwardingRulesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListIpForwardingRulesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListIpForwardingRulesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListIpForwardingRulesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListIpForwardingRulesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListIpForwardingRulesParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *ListIpForwardingRulesParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new ListIpForwardingRulesParams instance,
// as then you are sure you have configured all required params
func (s *NATService) NewListIpForwardingRulesParams() *ListIpForwardingRulesParams {
	P := &ListIpForwardingRulesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NATService) GetIpForwardingRuleByID(id string, opts ...OptionFunc) (*IpForwardingRule, int, error) {
	P := &ListIpForwardingRulesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListIpForwardingRules(P)
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
		return l.IpForwardingRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for IpForwardingRule UUID: %s!", id)
}

// List the IP forwarding rules
func (s *NATService) ListIpForwardingRules(p *ListIpForwardingRulesParams) (*ListIpForwardingRulesResponse, error) {
	resp, err := s.cs.newRequest("listIpForwardingRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListIpForwardingRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListIpForwardingRulesResponse struct {
	Count             int                 `json:"count"`
	IpForwardingRules []*IpForwardingRule `json:"ipforwardingrule"`
}

type IpForwardingRule struct {
	Cidrlist                  string `json:"cidrlist"`
	Fordisplay                bool   `json:"fordisplay"`
	Id                        string `json:"id"`
	Ipaddress                 string `json:"ipaddress"`
	Ipaddressid               string `json:"ipaddressid"`
	JobID                     string `json:"jobid"`
	Jobstatus                 int    `json:"jobstatus"`
	Networkid                 string `json:"networkid"`
	Privateendport            string `json:"privateendport"`
	Privateport               string `json:"privateport"`
	Protocol                  string `json:"protocol"`
	Publicendport             string `json:"publicendport"`
	Publicport                string `json:"publicport"`
	State                     string `json:"state"`
	Tags                      []Tags `json:"tags"`
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname"`
	Virtualmachineid          string `json:"virtualmachineid"`
	Virtualmachinename        string `json:"virtualmachinename"`
	Vmguestip                 string `json:"vmguestip"`
}
