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
func convertFirewallServiceResponse(b []byte) ([]byte, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	if _, ok := raw["firewallrule"]; ok {
		return convertFirewallServiceListResponse(b)
	}

	for _, k := range []string{"endport", "startport"} {
		if sVal, ok := raw[k].(string); ok {
			iVal, err := strconv.Atoi(sVal)
			if err != nil {
				return nil, err
			}
			raw[k] = iVal
		}
	}

	return json.Marshal(raw)
}

// Helper function for maintaining backwards compatibility
func convertFirewallServiceListResponse(b []byte) ([]byte, error) {
	var rawList struct {
		Count         int                      `json:"count"`
		FirewallRules []map[string]interface{} `json:"firewallrule"`
	}

	if err := json.Unmarshal(b, &rawList); err != nil {
		return nil, err
	}

	for _, r := range rawList.FirewallRules {
		for _, k := range []string{"endport", "startport"} {
			if sVal, ok := r[k].(string); ok {
				iVal, err := strconv.Atoi(sVal)
				if err != nil {
					return nil, err
				}
				r[k] = iVal
			}
		}
	}

	return json.Marshal(rawList)
}

type FirewallServiceIface interface {
	AddPaloAltoFirewall(p *AddPaloAltoFirewallParams) (*AddPaloAltoFirewallResponse, error)
	NewAddPaloAltoFirewallParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddPaloAltoFirewallParams
	ConfigurePaloAltoFirewall(p *ConfigurePaloAltoFirewallParams) (*PaloAltoFirewallResponse, error)
	NewConfigurePaloAltoFirewallParams(fwdeviceid string) *ConfigurePaloAltoFirewallParams
	CreateEgressFirewallRule(p *CreateEgressFirewallRuleParams) (*CreateEgressFirewallRuleResponse, error)
	NewCreateEgressFirewallRuleParams(networkid string, protocol string) *CreateEgressFirewallRuleParams
	CreateFirewallRule(p *CreateFirewallRuleParams) (*CreateFirewallRuleResponse, error)
	NewCreateFirewallRuleParams(ipaddressid string, protocol string) *CreateFirewallRuleParams
	CreatePortForwardingRule(p *CreatePortForwardingRuleParams) (*CreatePortForwardingRuleResponse, error)
	NewCreatePortForwardingRuleParams(ipaddressid string, privateport int, protocol string, publicport int, virtualmachineid string) *CreatePortForwardingRuleParams
	DeleteEgressFirewallRule(p *DeleteEgressFirewallRuleParams) (*DeleteEgressFirewallRuleResponse, error)
	NewDeleteEgressFirewallRuleParams(id string) *DeleteEgressFirewallRuleParams
	DeleteFirewallRule(p *DeleteFirewallRuleParams) (*DeleteFirewallRuleResponse, error)
	NewDeleteFirewallRuleParams(id string) *DeleteFirewallRuleParams
	DeletePaloAltoFirewall(p *DeletePaloAltoFirewallParams) (*DeletePaloAltoFirewallResponse, error)
	NewDeletePaloAltoFirewallParams(fwdeviceid string) *DeletePaloAltoFirewallParams
	DeletePortForwardingRule(p *DeletePortForwardingRuleParams) (*DeletePortForwardingRuleResponse, error)
	NewDeletePortForwardingRuleParams(id string) *DeletePortForwardingRuleParams
	ListEgressFirewallRules(p *ListEgressFirewallRulesParams) (*ListEgressFirewallRulesResponse, error)
	NewListEgressFirewallRulesParams() *ListEgressFirewallRulesParams
	GetEgressFirewallRuleByID(id string, opts ...OptionFunc) (*EgressFirewallRule, int, error)
	ListFirewallRules(p *ListFirewallRulesParams) (*ListFirewallRulesResponse, error)
	NewListFirewallRulesParams() *ListFirewallRulesParams
	GetFirewallRuleByID(id string, opts ...OptionFunc) (*FirewallRule, int, error)
	ListPaloAltoFirewalls(p *ListPaloAltoFirewallsParams) (*ListPaloAltoFirewallsResponse, error)
	NewListPaloAltoFirewallsParams() *ListPaloAltoFirewallsParams
	ListPortForwardingRules(p *ListPortForwardingRulesParams) (*ListPortForwardingRulesResponse, error)
	NewListPortForwardingRulesParams() *ListPortForwardingRulesParams
	GetPortForwardingRuleByID(id string, opts ...OptionFunc) (*PortForwardingRule, int, error)
	UpdateEgressFirewallRule(p *UpdateEgressFirewallRuleParams) (*UpdateEgressFirewallRuleResponse, error)
	NewUpdateEgressFirewallRuleParams(id string) *UpdateEgressFirewallRuleParams
	UpdateFirewallRule(p *UpdateFirewallRuleParams) (*UpdateFirewallRuleResponse, error)
	NewUpdateFirewallRuleParams(id string) *UpdateFirewallRuleParams
	UpdatePortForwardingRule(p *UpdatePortForwardingRuleParams) (*UpdatePortForwardingRuleResponse, error)
	NewUpdatePortForwardingRuleParams(id string) *UpdatePortForwardingRuleParams
}

type AddPaloAltoFirewallParams struct {
	P map[string]interface{}
}

func (P *AddPaloAltoFirewallParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["networkdevicetype"]; found {
		u.Set("networkdevicetype", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *AddPaloAltoFirewallParams) SetNetworkdevicetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdevicetype"] = v
}

func (P *AddPaloAltoFirewallParams) GetNetworkdevicetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdevicetype"].(string)
	return value, ok
}

func (P *AddPaloAltoFirewallParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddPaloAltoFirewallParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddPaloAltoFirewallParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddPaloAltoFirewallParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddPaloAltoFirewallParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddPaloAltoFirewallParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddPaloAltoFirewallParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddPaloAltoFirewallParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddPaloAltoFirewallParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewAddPaloAltoFirewallParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddPaloAltoFirewallParams {
	P := &AddPaloAltoFirewallParams{}
	P.P = make(map[string]interface{})
	P.P["networkdevicetype"] = networkdevicetype
	P.P["password"] = password
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["url"] = url
	P.P["username"] = username
	return P
}

// Adds a Palo Alto firewall device
func (s *FirewallService) AddPaloAltoFirewall(p *AddPaloAltoFirewallParams) (*AddPaloAltoFirewallResponse, error) {
	resp, err := s.cs.newRequest("addPaloAltoFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddPaloAltoFirewallResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AddPaloAltoFirewallResponse struct {
	Fwdevicecapacity  int64  `json:"fwdevicecapacity"`
	Fwdeviceid        string `json:"fwdeviceid"`
	Fwdevicename      string `json:"fwdevicename"`
	Fwdevicestate     string `json:"fwdevicestate"`
	Ipaddress         string `json:"ipaddress"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Numretries        string `json:"numretries"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Privateinterface  string `json:"privateinterface"`
	Privatezone       string `json:"privatezone"`
	Provider          string `json:"provider"`
	Publicinterface   string `json:"publicinterface"`
	Publiczone        string `json:"publiczone"`
	Timeout           string `json:"timeout"`
	Usageinterface    string `json:"usageinterface"`
	Username          string `json:"username"`
	Zoneid            string `json:"zoneid"`
}

type ConfigurePaloAltoFirewallParams struct {
	P map[string]interface{}
}

func (P *ConfigurePaloAltoFirewallParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fwdevicecapacity"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("fwdevicecapacity", vv)
	}
	if v, found := P.P["fwdeviceid"]; found {
		u.Set("fwdeviceid", v.(string))
	}
	return u
}

func (P *ConfigurePaloAltoFirewallParams) SetFwdevicecapacity(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fwdevicecapacity"] = v
}

func (P *ConfigurePaloAltoFirewallParams) GetFwdevicecapacity() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fwdevicecapacity"].(int64)
	return value, ok
}

func (P *ConfigurePaloAltoFirewallParams) SetFwdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fwdeviceid"] = v
}

func (P *ConfigurePaloAltoFirewallParams) GetFwdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fwdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new ConfigurePaloAltoFirewallParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewConfigurePaloAltoFirewallParams(fwdeviceid string) *ConfigurePaloAltoFirewallParams {
	P := &ConfigurePaloAltoFirewallParams{}
	P.P = make(map[string]interface{})
	P.P["fwdeviceid"] = fwdeviceid
	return P
}

// Configures a Palo Alto firewall device
func (s *FirewallService) ConfigurePaloAltoFirewall(p *ConfigurePaloAltoFirewallParams) (*PaloAltoFirewallResponse, error) {
	resp, err := s.cs.newRequest("configurePaloAltoFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PaloAltoFirewallResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type PaloAltoFirewallResponse struct {
	Fwdevicecapacity  int64  `json:"fwdevicecapacity"`
	Fwdeviceid        string `json:"fwdeviceid"`
	Fwdevicename      string `json:"fwdevicename"`
	Fwdevicestate     string `json:"fwdevicestate"`
	Ipaddress         string `json:"ipaddress"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Numretries        string `json:"numretries"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Privateinterface  string `json:"privateinterface"`
	Privatezone       string `json:"privatezone"`
	Provider          string `json:"provider"`
	Publicinterface   string `json:"publicinterface"`
	Publiczone        string `json:"publiczone"`
	Timeout           string `json:"timeout"`
	Usageinterface    string `json:"usageinterface"`
	Username          string `json:"username"`
	Zoneid            string `json:"zoneid"`
}

type CreateEgressFirewallRuleParams struct {
	P map[string]interface{}
}

func (P *CreateEgressFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := P.P["destcidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("destcidrlist", vv)
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
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (P *CreateEgressFirewallRuleParams) SetCidrlist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *CreateEgressFirewallRuleParams) GetCidrlist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].([]string)
	return value, ok
}

func (P *CreateEgressFirewallRuleParams) SetDestcidrlist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["destcidrlist"] = v
}

func (P *CreateEgressFirewallRuleParams) GetDestcidrlist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["destcidrlist"].([]string)
	return value, ok
}

func (P *CreateEgressFirewallRuleParams) SetEndport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endport"] = v
}

func (P *CreateEgressFirewallRuleParams) GetEndport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endport"].(int)
	return value, ok
}

func (P *CreateEgressFirewallRuleParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateEgressFirewallRuleParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateEgressFirewallRuleParams) SetIcmpcode(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmpcode"] = v
}

func (P *CreateEgressFirewallRuleParams) GetIcmpcode() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmpcode"].(int)
	return value, ok
}

func (P *CreateEgressFirewallRuleParams) SetIcmptype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmptype"] = v
}

func (P *CreateEgressFirewallRuleParams) GetIcmptype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmptype"].(int)
	return value, ok
}

func (P *CreateEgressFirewallRuleParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *CreateEgressFirewallRuleParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *CreateEgressFirewallRuleParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *CreateEgressFirewallRuleParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *CreateEgressFirewallRuleParams) SetStartport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startport"] = v
}

func (P *CreateEgressFirewallRuleParams) GetStartport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startport"].(int)
	return value, ok
}

func (P *CreateEgressFirewallRuleParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *CreateEgressFirewallRuleParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new CreateEgressFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreateEgressFirewallRuleParams(networkid string, protocol string) *CreateEgressFirewallRuleParams {
	P := &CreateEgressFirewallRuleParams{}
	P.P = make(map[string]interface{})
	P.P["networkid"] = networkid
	P.P["protocol"] = protocol
	return P
}

// Creates a egress firewall rule for a given network
func (s *FirewallService) CreateEgressFirewallRule(p *CreateEgressFirewallRuleParams) (*CreateEgressFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("createEgressFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateEgressFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateEgressFirewallRuleResponse struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
}

type CreateFirewallRuleParams struct {
	P map[string]interface{}
}

func (P *CreateFirewallRuleParams) toURLValues() url.Values {
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
	if v, found := P.P["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (P *CreateFirewallRuleParams) SetCidrlist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *CreateFirewallRuleParams) GetCidrlist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].([]string)
	return value, ok
}

func (P *CreateFirewallRuleParams) SetEndport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endport"] = v
}

func (P *CreateFirewallRuleParams) GetEndport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endport"].(int)
	return value, ok
}

func (P *CreateFirewallRuleParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateFirewallRuleParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateFirewallRuleParams) SetIcmpcode(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmpcode"] = v
}

func (P *CreateFirewallRuleParams) GetIcmpcode() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmpcode"].(int)
	return value, ok
}

func (P *CreateFirewallRuleParams) SetIcmptype(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["icmptype"] = v
}

func (P *CreateFirewallRuleParams) GetIcmptype() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["icmptype"].(int)
	return value, ok
}

func (P *CreateFirewallRuleParams) SetIpaddressid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddressid"] = v
}

func (P *CreateFirewallRuleParams) GetIpaddressid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddressid"].(string)
	return value, ok
}

func (P *CreateFirewallRuleParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *CreateFirewallRuleParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *CreateFirewallRuleParams) SetStartport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startport"] = v
}

func (P *CreateFirewallRuleParams) GetStartport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startport"].(int)
	return value, ok
}

func (P *CreateFirewallRuleParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *CreateFirewallRuleParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new CreateFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreateFirewallRuleParams(ipaddressid string, protocol string) *CreateFirewallRuleParams {
	P := &CreateFirewallRuleParams{}
	P.P = make(map[string]interface{})
	P.P["ipaddressid"] = ipaddressid
	P.P["protocol"] = protocol
	return P
}

// Creates a firewall rule for a given IP address
func (s *FirewallService) CreateFirewallRule(p *CreateFirewallRuleParams) (*CreateFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("createFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateFirewallRuleResponse struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
}

type CreatePortForwardingRuleParams struct {
	P map[string]interface{}
}

func (P *CreatePortForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := P.P["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := P.P["privateendport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateendport", vv)
	}
	if v, found := P.P["privateport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateport", vv)
	}
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["publicendport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicendport", vv)
	}
	if v, found := P.P["publicport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicport", vv)
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := P.P["vmguestip"]; found {
		u.Set("vmguestip", v.(string))
	}
	return u
}

func (P *CreatePortForwardingRuleParams) SetCidrlist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *CreatePortForwardingRuleParams) GetCidrlist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].([]string)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreatePortForwardingRuleParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetIpaddressid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddressid"] = v
}

func (P *CreatePortForwardingRuleParams) GetIpaddressid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddressid"].(string)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *CreatePortForwardingRuleParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetOpenfirewall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["openfirewall"] = v
}

func (P *CreatePortForwardingRuleParams) GetOpenfirewall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["openfirewall"].(bool)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetPrivateendport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["privateendport"] = v
}

func (P *CreatePortForwardingRuleParams) GetPrivateendport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["privateendport"].(int)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetPrivateport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["privateport"] = v
}

func (P *CreatePortForwardingRuleParams) GetPrivateport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["privateport"].(int)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *CreatePortForwardingRuleParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetPublicendport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publicendport"] = v
}

func (P *CreatePortForwardingRuleParams) GetPublicendport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publicendport"].(int)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetPublicport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publicport"] = v
}

func (P *CreatePortForwardingRuleParams) GetPublicport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publicport"].(int)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *CreatePortForwardingRuleParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *CreatePortForwardingRuleParams) SetVmguestip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmguestip"] = v
}

func (P *CreatePortForwardingRuleParams) GetVmguestip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmguestip"].(string)
	return value, ok
}

// You should always use this function to get a new CreatePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreatePortForwardingRuleParams(ipaddressid string, privateport int, protocol string, publicport int, virtualmachineid string) *CreatePortForwardingRuleParams {
	P := &CreatePortForwardingRuleParams{}
	P.P = make(map[string]interface{})
	P.P["ipaddressid"] = ipaddressid
	P.P["privateport"] = privateport
	P.P["protocol"] = protocol
	P.P["publicport"] = publicport
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Creates a port forwarding rule
func (s *FirewallService) CreatePortForwardingRule(p *CreatePortForwardingRuleParams) (*CreatePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("createPortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePortForwardingRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreatePortForwardingRuleResponse struct {
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

type DeleteEgressFirewallRuleParams struct {
	P map[string]interface{}
}

func (P *DeleteEgressFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteEgressFirewallRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteEgressFirewallRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteEgressFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeleteEgressFirewallRuleParams(id string) *DeleteEgressFirewallRuleParams {
	P := &DeleteEgressFirewallRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes an egress firewall rule
func (s *FirewallService) DeleteEgressFirewallRule(p *DeleteEgressFirewallRuleParams) (*DeleteEgressFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteEgressFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteEgressFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteEgressFirewallRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteFirewallRuleParams struct {
	P map[string]interface{}
}

func (P *DeleteFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteFirewallRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteFirewallRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeleteFirewallRuleParams(id string) *DeleteFirewallRuleParams {
	P := &DeleteFirewallRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a firewall rule
func (s *FirewallService) DeleteFirewallRule(p *DeleteFirewallRuleParams) (*DeleteFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteFirewallRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeletePaloAltoFirewallParams struct {
	P map[string]interface{}
}

func (P *DeletePaloAltoFirewallParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fwdeviceid"]; found {
		u.Set("fwdeviceid", v.(string))
	}
	return u
}

func (P *DeletePaloAltoFirewallParams) SetFwdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fwdeviceid"] = v
}

func (P *DeletePaloAltoFirewallParams) GetFwdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fwdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new DeletePaloAltoFirewallParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeletePaloAltoFirewallParams(fwdeviceid string) *DeletePaloAltoFirewallParams {
	P := &DeletePaloAltoFirewallParams{}
	P.P = make(map[string]interface{})
	P.P["fwdeviceid"] = fwdeviceid
	return P
}

// delete a Palo Alto firewall device
func (s *FirewallService) DeletePaloAltoFirewall(p *DeletePaloAltoFirewallParams) (*DeletePaloAltoFirewallResponse, error) {
	resp, err := s.cs.newRequest("deletePaloAltoFirewall", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePaloAltoFirewallResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeletePaloAltoFirewallResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeletePortForwardingRuleParams struct {
	P map[string]interface{}
}

func (P *DeletePortForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeletePortForwardingRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeletePortForwardingRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeletePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeletePortForwardingRuleParams(id string) *DeletePortForwardingRuleParams {
	P := &DeletePortForwardingRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a port forwarding rule
func (s *FirewallService) DeletePortForwardingRule(p *DeletePortForwardingRuleParams) (*DeletePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("deletePortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePortForwardingRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeletePortForwardingRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListEgressFirewallRulesParams struct {
	P map[string]interface{}
}

func (P *ListEgressFirewallRulesParams) toURLValues() url.Values {
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
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (P *ListEgressFirewallRulesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListEgressFirewallRulesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListEgressFirewallRulesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListEgressFirewallRulesParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListEgressFirewallRulesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetIpaddressid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddressid"] = v
}

func (P *ListEgressFirewallRulesParams) GetIpaddressid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddressid"].(string)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListEgressFirewallRulesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListEgressFirewallRulesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListEgressFirewallRulesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListEgressFirewallRulesParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListEgressFirewallRulesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListEgressFirewallRulesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListEgressFirewallRulesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListEgressFirewallRulesParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListEgressFirewallRulesParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListEgressFirewallRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListEgressFirewallRulesParams() *ListEgressFirewallRulesParams {
	P := &ListEgressFirewallRulesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *FirewallService) GetEgressFirewallRuleByID(id string, opts ...OptionFunc) (*EgressFirewallRule, int, error) {
	P := &ListEgressFirewallRulesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListEgressFirewallRules(P)
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
		return l.EgressFirewallRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for EgressFirewallRule UUID: %s!", id)
}

// Lists all egress firewall rules for network ID.
func (s *FirewallService) ListEgressFirewallRules(p *ListEgressFirewallRulesParams) (*ListEgressFirewallRulesResponse, error) {
	resp, err := s.cs.newRequest("listEgressFirewallRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListEgressFirewallRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListEgressFirewallRulesResponse struct {
	Count               int                   `json:"count"`
	EgressFirewallRules []*EgressFirewallRule `json:"firewallrule"`
}

type EgressFirewallRule struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
}

type ListFirewallRulesParams struct {
	P map[string]interface{}
}

func (P *ListFirewallRulesParams) toURLValues() url.Values {
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
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (P *ListFirewallRulesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListFirewallRulesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListFirewallRulesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListFirewallRulesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListFirewallRulesParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListFirewallRulesParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListFirewallRulesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListFirewallRulesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListFirewallRulesParams) SetIpaddressid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddressid"] = v
}

func (P *ListFirewallRulesParams) GetIpaddressid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddressid"].(string)
	return value, ok
}

func (P *ListFirewallRulesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListFirewallRulesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListFirewallRulesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListFirewallRulesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListFirewallRulesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListFirewallRulesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListFirewallRulesParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListFirewallRulesParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListFirewallRulesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListFirewallRulesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListFirewallRulesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListFirewallRulesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListFirewallRulesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListFirewallRulesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListFirewallRulesParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListFirewallRulesParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListFirewallRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListFirewallRulesParams() *ListFirewallRulesParams {
	P := &ListFirewallRulesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *FirewallService) GetFirewallRuleByID(id string, opts ...OptionFunc) (*FirewallRule, int, error) {
	P := &ListFirewallRulesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListFirewallRules(P)
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
		return l.FirewallRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for FirewallRule UUID: %s!", id)
}

// Lists all firewall rules for an IP address.
func (s *FirewallService) ListFirewallRules(p *ListFirewallRulesParams) (*ListFirewallRulesResponse, error) {
	resp, err := s.cs.newRequest("listFirewallRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListFirewallRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListFirewallRulesResponse struct {
	Count         int             `json:"count"`
	FirewallRules []*FirewallRule `json:"firewallrule"`
}

type FirewallRule struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
}

type ListPaloAltoFirewallsParams struct {
	P map[string]interface{}
}

func (P *ListPaloAltoFirewallsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fwdeviceid"]; found {
		u.Set("fwdeviceid", v.(string))
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
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (P *ListPaloAltoFirewallsParams) SetFwdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fwdeviceid"] = v
}

func (P *ListPaloAltoFirewallsParams) GetFwdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fwdeviceid"].(string)
	return value, ok
}

func (P *ListPaloAltoFirewallsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListPaloAltoFirewallsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListPaloAltoFirewallsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListPaloAltoFirewallsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListPaloAltoFirewallsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListPaloAltoFirewallsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListPaloAltoFirewallsParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListPaloAltoFirewallsParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListPaloAltoFirewallsParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListPaloAltoFirewallsParams() *ListPaloAltoFirewallsParams {
	P := &ListPaloAltoFirewallsParams{}
	P.P = make(map[string]interface{})
	return P
}

// lists Palo Alto firewall devices in a physical network
func (s *FirewallService) ListPaloAltoFirewalls(p *ListPaloAltoFirewallsParams) (*ListPaloAltoFirewallsResponse, error) {
	resp, err := s.cs.newRequest("listPaloAltoFirewalls", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListPaloAltoFirewallsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPaloAltoFirewallsResponse struct {
	Count             int                 `json:"count"`
	PaloAltoFirewalls []*PaloAltoFirewall `json:"paloaltofirewall"`
}

type PaloAltoFirewall struct {
	Fwdevicecapacity  int64  `json:"fwdevicecapacity"`
	Fwdeviceid        string `json:"fwdeviceid"`
	Fwdevicename      string `json:"fwdevicename"`
	Fwdevicestate     string `json:"fwdevicestate"`
	Ipaddress         string `json:"ipaddress"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Numretries        string `json:"numretries"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Privateinterface  string `json:"privateinterface"`
	Privatezone       string `json:"privatezone"`
	Provider          string `json:"provider"`
	Publicinterface   string `json:"publicinterface"`
	Publiczone        string `json:"publiczone"`
	Timeout           string `json:"timeout"`
	Usageinterface    string `json:"usageinterface"`
	Username          string `json:"username"`
	Zoneid            string `json:"zoneid"`
}

type ListPortForwardingRulesParams struct {
	P map[string]interface{}
}

func (P *ListPortForwardingRulesParams) toURLValues() url.Values {
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
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	return u
}

func (P *ListPortForwardingRulesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListPortForwardingRulesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListPortForwardingRulesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListPortForwardingRulesParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListPortForwardingRulesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetIpaddressid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddressid"] = v
}

func (P *ListPortForwardingRulesParams) GetIpaddressid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddressid"].(string)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListPortForwardingRulesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListPortForwardingRulesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListPortForwardingRulesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListPortForwardingRulesParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListPortForwardingRulesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListPortForwardingRulesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListPortForwardingRulesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListPortForwardingRulesParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListPortForwardingRulesParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListPortForwardingRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListPortForwardingRulesParams() *ListPortForwardingRulesParams {
	P := &ListPortForwardingRulesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *FirewallService) GetPortForwardingRuleByID(id string, opts ...OptionFunc) (*PortForwardingRule, int, error) {
	P := &ListPortForwardingRulesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListPortForwardingRules(P)
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
		return l.PortForwardingRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for PortForwardingRule UUID: %s!", id)
}

// Lists all port forwarding rules for an IP address.
func (s *FirewallService) ListPortForwardingRules(p *ListPortForwardingRulesParams) (*ListPortForwardingRulesResponse, error) {
	resp, err := s.cs.newRequest("listPortForwardingRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListPortForwardingRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPortForwardingRulesResponse struct {
	Count               int                   `json:"count"`
	PortForwardingRules []*PortForwardingRule `json:"portforwardingrule"`
}

type PortForwardingRule struct {
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

type UpdateEgressFirewallRuleParams struct {
	P map[string]interface{}
}

func (P *UpdateEgressFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *UpdateEgressFirewallRuleParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateEgressFirewallRuleParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateEgressFirewallRuleParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateEgressFirewallRuleParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateEgressFirewallRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateEgressFirewallRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateEgressFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewUpdateEgressFirewallRuleParams(id string) *UpdateEgressFirewallRuleParams {
	P := &UpdateEgressFirewallRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates egress firewall rule
func (s *FirewallService) UpdateEgressFirewallRule(p *UpdateEgressFirewallRuleParams) (*UpdateEgressFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("updateEgressFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateEgressFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateEgressFirewallRuleResponse struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
}

type UpdateFirewallRuleParams struct {
	P map[string]interface{}
}

func (P *UpdateFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *UpdateFirewallRuleParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateFirewallRuleParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateFirewallRuleParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateFirewallRuleParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateFirewallRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateFirewallRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewUpdateFirewallRuleParams(id string) *UpdateFirewallRuleParams {
	P := &UpdateFirewallRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates firewall rule
func (s *FirewallService) UpdateFirewallRule(p *UpdateFirewallRuleParams) (*UpdateFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("updateFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateFirewallRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdateFirewallRuleResponse struct {
	Cidrlist     string `json:"cidrlist"`
	Destcidrlist string `json:"destcidrlist"`
	Endport      int    `json:"endport"`
	Fordisplay   bool   `json:"fordisplay"`
	Icmpcode     int    `json:"icmpcode"`
	Icmptype     int    `json:"icmptype"`
	Id           string `json:"id"`
	Ipaddress    string `json:"ipaddress"`
	Ipaddressid  string `json:"ipaddressid"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Networkid    string `json:"networkid"`
	Protocol     string `json:"protocol"`
	Startport    int    `json:"startport"`
	State        string `json:"state"`
	Tags         []Tags `json:"tags"`
}

type UpdatePortForwardingRuleParams struct {
	P map[string]interface{}
}

func (P *UpdatePortForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["privateendport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateendport", vv)
	}
	if v, found := P.P["privateport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateport", vv)
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := P.P["vmguestip"]; found {
		u.Set("vmguestip", v.(string))
	}
	return u
}

func (P *UpdatePortForwardingRuleParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdatePortForwardingRuleParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdatePortForwardingRuleParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdatePortForwardingRuleParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdatePortForwardingRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdatePortForwardingRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdatePortForwardingRuleParams) SetPrivateendport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["privateendport"] = v
}

func (P *UpdatePortForwardingRuleParams) GetPrivateendport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["privateendport"].(int)
	return value, ok
}

func (P *UpdatePortForwardingRuleParams) SetPrivateport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["privateport"] = v
}

func (P *UpdatePortForwardingRuleParams) GetPrivateport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["privateport"].(int)
	return value, ok
}

func (P *UpdatePortForwardingRuleParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *UpdatePortForwardingRuleParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *UpdatePortForwardingRuleParams) SetVmguestip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmguestip"] = v
}

func (P *UpdatePortForwardingRuleParams) GetVmguestip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmguestip"].(string)
	return value, ok
}

// You should always use this function to get a new UpdatePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewUpdatePortForwardingRuleParams(id string) *UpdatePortForwardingRuleParams {
	P := &UpdatePortForwardingRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a port forwarding rule. Only the private port and the virtual machine can be updated.
func (s *FirewallService) UpdatePortForwardingRule(p *UpdatePortForwardingRuleParams) (*UpdatePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("updatePortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdatePortForwardingRuleResponse
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

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type UpdatePortForwardingRuleResponse struct {
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
