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

type LoadBalancerServiceIface interface {
	AddNetscalerLoadBalancer(P *AddNetscalerLoadBalancerParams) (*AddNetscalerLoadBalancerResponse, error)
	NewAddNetscalerLoadBalancerParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddNetscalerLoadBalancerParams
	AssignCertToLoadBalancer(P *AssignCertToLoadBalancerParams) (*AssignCertToLoadBalancerResponse, error)
	NewAssignCertToLoadBalancerParams(certid string, lbruleid string) *AssignCertToLoadBalancerParams
	AssignToGlobalLoadBalancerRule(P *AssignToGlobalLoadBalancerRuleParams) (*AssignToGlobalLoadBalancerRuleResponse, error)
	NewAssignToGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *AssignToGlobalLoadBalancerRuleParams
	AssignToLoadBalancerRule(P *AssignToLoadBalancerRuleParams) (*AssignToLoadBalancerRuleResponse, error)
	NewAssignToLoadBalancerRuleParams(id string) *AssignToLoadBalancerRuleParams
	ConfigureNetscalerLoadBalancer(P *ConfigureNetscalerLoadBalancerParams) (*NetscalerLoadBalancerResponse, error)
	NewConfigureNetscalerLoadBalancerParams(lbdeviceid string) *ConfigureNetscalerLoadBalancerParams
	CreateGlobalLoadBalancerRule(P *CreateGlobalLoadBalancerRuleParams) (*CreateGlobalLoadBalancerRuleResponse, error)
	NewCreateGlobalLoadBalancerRuleParams(gslbdomainname string, gslbservicetype string, name string, regionid int) *CreateGlobalLoadBalancerRuleParams
	CreateLBHealthCheckPolicy(P *CreateLBHealthCheckPolicyParams) (*CreateLBHealthCheckPolicyResponse, error)
	NewCreateLBHealthCheckPolicyParams(lbruleid string) *CreateLBHealthCheckPolicyParams
	CreateLBStickinessPolicy(P *CreateLBStickinessPolicyParams) (*CreateLBStickinessPolicyResponse, error)
	NewCreateLBStickinessPolicyParams(lbruleid string, methodname string, name string) *CreateLBStickinessPolicyParams
	CreateLoadBalancer(P *CreateLoadBalancerParams) (*CreateLoadBalancerResponse, error)
	NewCreateLoadBalancerParams(algorithm string, instanceport int, name string, networkid string, scheme string, sourceipaddressnetworkid string, sourceport int) *CreateLoadBalancerParams
	CreateLoadBalancerRule(P *CreateLoadBalancerRuleParams) (*CreateLoadBalancerRuleResponse, error)
	NewCreateLoadBalancerRuleParams(algorithm string, name string, privateport int, publicport int) *CreateLoadBalancerRuleParams
	DeleteGlobalLoadBalancerRule(P *DeleteGlobalLoadBalancerRuleParams) (*DeleteGlobalLoadBalancerRuleResponse, error)
	NewDeleteGlobalLoadBalancerRuleParams(id string) *DeleteGlobalLoadBalancerRuleParams
	DeleteLBHealthCheckPolicy(P *DeleteLBHealthCheckPolicyParams) (*DeleteLBHealthCheckPolicyResponse, error)
	NewDeleteLBHealthCheckPolicyParams(id string) *DeleteLBHealthCheckPolicyParams
	DeleteLBStickinessPolicy(P *DeleteLBStickinessPolicyParams) (*DeleteLBStickinessPolicyResponse, error)
	NewDeleteLBStickinessPolicyParams(id string) *DeleteLBStickinessPolicyParams
	DeleteLoadBalancer(P *DeleteLoadBalancerParams) (*DeleteLoadBalancerResponse, error)
	NewDeleteLoadBalancerParams(id string) *DeleteLoadBalancerParams
	DeleteLoadBalancerRule(P *DeleteLoadBalancerRuleParams) (*DeleteLoadBalancerRuleResponse, error)
	NewDeleteLoadBalancerRuleParams(id string) *DeleteLoadBalancerRuleParams
	DeleteNetscalerLoadBalancer(P *DeleteNetscalerLoadBalancerParams) (*DeleteNetscalerLoadBalancerResponse, error)
	NewDeleteNetscalerLoadBalancerParams(lbdeviceid string) *DeleteNetscalerLoadBalancerParams
	DeleteSslCert(P *DeleteSslCertParams) (*DeleteSslCertResponse, error)
	NewDeleteSslCertParams(id string) *DeleteSslCertParams
	ListGlobalLoadBalancerRules(P *ListGlobalLoadBalancerRulesParams) (*ListGlobalLoadBalancerRulesResponse, error)
	NewListGlobalLoadBalancerRulesParams() *ListGlobalLoadBalancerRulesParams
	GetGlobalLoadBalancerRuleID(keyword string, opts ...OptionFunc) (string, int, error)
	GetGlobalLoadBalancerRuleByName(name string, opts ...OptionFunc) (*GlobalLoadBalancerRule, int, error)
	GetGlobalLoadBalancerRuleByID(id string, opts ...OptionFunc) (*GlobalLoadBalancerRule, int, error)
	ListLBHealthCheckPolicies(P *ListLBHealthCheckPoliciesParams) (*ListLBHealthCheckPoliciesResponse, error)
	NewListLBHealthCheckPoliciesParams() *ListLBHealthCheckPoliciesParams
	GetLBHealthCheckPolicyByID(id string, opts ...OptionFunc) (*LBHealthCheckPolicy, int, error)
	ListLBStickinessPolicies(P *ListLBStickinessPoliciesParams) (*ListLBStickinessPoliciesResponse, error)
	NewListLBStickinessPoliciesParams() *ListLBStickinessPoliciesParams
	GetLBStickinessPolicyByID(id string, opts ...OptionFunc) (*LBStickinessPolicy, int, error)
	ListLoadBalancerRuleInstances(P *ListLoadBalancerRuleInstancesParams) (*ListLoadBalancerRuleInstancesResponse, error)
	NewListLoadBalancerRuleInstancesParams(id string) *ListLoadBalancerRuleInstancesParams
	GetLoadBalancerRuleInstanceByID(id string, opts ...OptionFunc) (*VirtualMachine, int, error)
	ListLoadBalancerRules(P *ListLoadBalancerRulesParams) (*ListLoadBalancerRulesResponse, error)
	NewListLoadBalancerRulesParams() *ListLoadBalancerRulesParams
	GetLoadBalancerRuleID(name string, opts ...OptionFunc) (string, int, error)
	GetLoadBalancerRuleByName(name string, opts ...OptionFunc) (*LoadBalancerRule, int, error)
	GetLoadBalancerRuleByID(id string, opts ...OptionFunc) (*LoadBalancerRule, int, error)
	ListLoadBalancers(P *ListLoadBalancersParams) (*ListLoadBalancersResponse, error)
	NewListLoadBalancersParams() *ListLoadBalancersParams
	GetLoadBalancerID(name string, opts ...OptionFunc) (string, int, error)
	GetLoadBalancerByName(name string, opts ...OptionFunc) (*LoadBalancer, int, error)
	GetLoadBalancerByID(id string, opts ...OptionFunc) (*LoadBalancer, int, error)
	ListNetscalerLoadBalancers(P *ListNetscalerLoadBalancersParams) (*ListNetscalerLoadBalancersResponse, error)
	NewListNetscalerLoadBalancersParams() *ListNetscalerLoadBalancersParams
	ListSslCerts(P *ListSslCertsParams) (*ListSslCertsResponse, error)
	NewListSslCertsParams() *ListSslCertsParams
	RemoveCertFromLoadBalancer(P *RemoveCertFromLoadBalancerParams) (*RemoveCertFromLoadBalancerResponse, error)
	NewRemoveCertFromLoadBalancerParams(lbruleid string) *RemoveCertFromLoadBalancerParams
	RemoveFromGlobalLoadBalancerRule(P *RemoveFromGlobalLoadBalancerRuleParams) (*RemoveFromGlobalLoadBalancerRuleResponse, error)
	NewRemoveFromGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *RemoveFromGlobalLoadBalancerRuleParams
	RemoveFromLoadBalancerRule(P *RemoveFromLoadBalancerRuleParams) (*RemoveFromLoadBalancerRuleResponse, error)
	NewRemoveFromLoadBalancerRuleParams(id string) *RemoveFromLoadBalancerRuleParams
	UpdateGlobalLoadBalancerRule(P *UpdateGlobalLoadBalancerRuleParams) (*UpdateGlobalLoadBalancerRuleResponse, error)
	NewUpdateGlobalLoadBalancerRuleParams(id string) *UpdateGlobalLoadBalancerRuleParams
	UpdateLBHealthCheckPolicy(P *UpdateLBHealthCheckPolicyParams) (*UpdateLBHealthCheckPolicyResponse, error)
	NewUpdateLBHealthCheckPolicyParams(id string) *UpdateLBHealthCheckPolicyParams
	UpdateLBStickinessPolicy(P *UpdateLBStickinessPolicyParams) (*UpdateLBStickinessPolicyResponse, error)
	NewUpdateLBStickinessPolicyParams(id string) *UpdateLBStickinessPolicyParams
	UpdateLoadBalancer(P *UpdateLoadBalancerParams) (*UpdateLoadBalancerResponse, error)
	NewUpdateLoadBalancerParams(id string) *UpdateLoadBalancerParams
	UpdateLoadBalancerRule(P *UpdateLoadBalancerRuleParams) (*UpdateLoadBalancerRuleResponse, error)
	NewUpdateLoadBalancerRuleParams(id string) *UpdateLoadBalancerRuleParams
	UploadSslCert(P *UploadSslCertParams) (*UploadSslCertResponse, error)
	NewUploadSslCertParams(certificate string, name string, privatekey string) *UploadSslCertParams
}

type AddNetscalerLoadBalancerParams struct {
	P map[string]interface{}
}

func (P *AddNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["gslbprovider"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("gslbprovider", vv)
	}
	if v, found := P.P["gslbproviderprivateip"]; found {
		u.Set("gslbproviderprivateip", v.(string))
	}
	if v, found := P.P["gslbproviderpublicip"]; found {
		u.Set("gslbproviderpublicip", v.(string))
	}
	if v, found := P.P["isexclusivegslbprovider"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isexclusivegslbprovider", vv)
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

func (P *AddNetscalerLoadBalancerParams) SetGslbprovider(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gslbprovider"] = v
}

func (P *AddNetscalerLoadBalancerParams) GetGslbprovider() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gslbprovider"].(bool)
	return value, ok
}

func (P *AddNetscalerLoadBalancerParams) SetGslbproviderprivateip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gslbproviderprivateip"] = v
}

func (P *AddNetscalerLoadBalancerParams) GetGslbproviderprivateip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gslbproviderprivateip"].(string)
	return value, ok
}

func (P *AddNetscalerLoadBalancerParams) SetGslbproviderpublicip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gslbproviderpublicip"] = v
}

func (P *AddNetscalerLoadBalancerParams) GetGslbproviderpublicip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gslbproviderpublicip"].(string)
	return value, ok
}

func (P *AddNetscalerLoadBalancerParams) SetIsexclusivegslbprovider(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isexclusivegslbprovider"] = v
}

func (P *AddNetscalerLoadBalancerParams) GetIsexclusivegslbprovider() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isexclusivegslbprovider"].(bool)
	return value, ok
}

func (P *AddNetscalerLoadBalancerParams) SetNetworkdevicetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdevicetype"] = v
}

func (P *AddNetscalerLoadBalancerParams) GetNetworkdevicetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdevicetype"].(string)
	return value, ok
}

func (P *AddNetscalerLoadBalancerParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddNetscalerLoadBalancerParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddNetscalerLoadBalancerParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddNetscalerLoadBalancerParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddNetscalerLoadBalancerParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddNetscalerLoadBalancerParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddNetscalerLoadBalancerParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddNetscalerLoadBalancerParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAddNetscalerLoadBalancerParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddNetscalerLoadBalancerParams {
	P := &AddNetscalerLoadBalancerParams{}
	P.P = make(map[string]interface{})
	P.P["networkdevicetype"] = networkdevicetype
	P.P["password"] = password
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["url"] = url
	P.P["username"] = username
	return P
}

// Adds a netscaler load balancer device
func (s *LoadBalancerService) AddNetscalerLoadBalancer(p *AddNetscalerLoadBalancerParams) (*AddNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("addNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNetscalerLoadBalancerResponse
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

type AddNetscalerLoadBalancerResponse struct {
	Gslbprovider            bool     `json:"gslbprovider"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip"`
	Ipaddress               string   `json:"ipaddress"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider"`
	JobID                   string   `json:"jobid"`
	Jobstatus               int      `json:"jobstatus"`
	Lbdevicecapacity        int64    `json:"lbdevicecapacity"`
	Lbdevicededicated       bool     `json:"lbdevicededicated"`
	Lbdeviceid              string   `json:"lbdeviceid"`
	Lbdevicename            string   `json:"lbdevicename"`
	Lbdevicestate           string   `json:"lbdevicestate"`
	Physicalnetworkid       string   `json:"physicalnetworkid"`
	Podids                  []string `json:"podids"`
	Privateinterface        string   `json:"privateinterface"`
	Provider                string   `json:"provider"`
	Publicinterface         string   `json:"publicinterface"`
}

type AssignCertToLoadBalancerParams struct {
	P map[string]interface{}
}

func (P *AssignCertToLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["certid"]; found {
		u.Set("certid", v.(string))
	}
	if v, found := P.P["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	return u
}

func (P *AssignCertToLoadBalancerParams) SetCertid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["certid"] = v
}

func (P *AssignCertToLoadBalancerParams) GetCertid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["certid"].(string)
	return value, ok
}

func (P *AssignCertToLoadBalancerParams) SetLbruleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbruleid"] = v
}

func (P *AssignCertToLoadBalancerParams) GetLbruleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbruleid"].(string)
	return value, ok
}

// You should always use this function to get a new AssignCertToLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignCertToLoadBalancerParams(certid string, lbruleid string) *AssignCertToLoadBalancerParams {
	P := &AssignCertToLoadBalancerParams{}
	P.P = make(map[string]interface{})
	P.P["certid"] = certid
	P.P["lbruleid"] = lbruleid
	return P
}

// Assigns a certificate to a load balancer rule
func (s *LoadBalancerService) AssignCertToLoadBalancer(p *AssignCertToLoadBalancerParams) (*AssignCertToLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("assignCertToLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignCertToLoadBalancerResponse
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

type AssignCertToLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type AssignToGlobalLoadBalancerRuleParams struct {
	P map[string]interface{}
}

func (P *AssignToGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["gslblbruleweightsmap"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("gslblbruleweightsmap[%d].key", i), k)
			u.Set(fmt.Sprintf("gslblbruleweightsmap[%d].value", i), m[k])
		}
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["loadbalancerrulelist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("loadbalancerrulelist", vv)
	}
	return u
}

func (P *AssignToGlobalLoadBalancerRuleParams) SetGslblbruleweightsmap(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gslblbruleweightsmap"] = v
}

func (P *AssignToGlobalLoadBalancerRuleParams) GetGslblbruleweightsmap() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gslblbruleweightsmap"].(map[string]string)
	return value, ok
}

func (P *AssignToGlobalLoadBalancerRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *AssignToGlobalLoadBalancerRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *AssignToGlobalLoadBalancerRuleParams) SetLoadbalancerrulelist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["loadbalancerrulelist"] = v
}

func (P *AssignToGlobalLoadBalancerRuleParams) GetLoadbalancerrulelist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["loadbalancerrulelist"].([]string)
	return value, ok
}

// You should always use this function to get a new AssignToGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignToGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *AssignToGlobalLoadBalancerRuleParams {
	P := &AssignToGlobalLoadBalancerRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["loadbalancerrulelist"] = loadbalancerrulelist
	return P
}

// Assign load balancer rule or list of load balancer rules to a global load balancer rules.
func (s *LoadBalancerService) AssignToGlobalLoadBalancerRule(p *AssignToGlobalLoadBalancerRuleParams) (*AssignToGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("assignToGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignToGlobalLoadBalancerRuleResponse
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

type AssignToGlobalLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type AssignToLoadBalancerRuleParams struct {
	P map[string]interface{}
}

func (P *AssignToLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["virtualmachineids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("virtualmachineids", vv)
	}
	if v, found := P.P["vmidipmap"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("vmidipmap[%d].key", i), k)
			u.Set(fmt.Sprintf("vmidipmap[%d].value", i), m[k])
		}
	}
	return u
}

func (P *AssignToLoadBalancerRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *AssignToLoadBalancerRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *AssignToLoadBalancerRuleParams) SetVirtualmachineids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineids"] = v
}

func (P *AssignToLoadBalancerRuleParams) GetVirtualmachineids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineids"].([]string)
	return value, ok
}

func (P *AssignToLoadBalancerRuleParams) SetVmidipmap(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmidipmap"] = v
}

func (P *AssignToLoadBalancerRuleParams) GetVmidipmap() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmidipmap"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new AssignToLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignToLoadBalancerRuleParams(id string) *AssignToLoadBalancerRuleParams {
	P := &AssignToLoadBalancerRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Assigns virtual machine or a list of virtual machines to a load balancer rule.
func (s *LoadBalancerService) AssignToLoadBalancerRule(p *AssignToLoadBalancerRuleParams) (*AssignToLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("assignToLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignToLoadBalancerRuleResponse
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

type AssignToLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ConfigureNetscalerLoadBalancerParams struct {
	P map[string]interface{}
}

func (P *ConfigureNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["inline"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("inline", vv)
	}
	if v, found := P.P["lbdevicecapacity"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("lbdevicecapacity", vv)
	}
	if v, found := P.P["lbdevicededicated"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("lbdevicededicated", vv)
	}
	if v, found := P.P["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := P.P["podids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("podids", vv)
	}
	return u
}

func (P *ConfigureNetscalerLoadBalancerParams) SetInline(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["inline"] = v
}

func (P *ConfigureNetscalerLoadBalancerParams) GetInline() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["inline"].(bool)
	return value, ok
}

func (P *ConfigureNetscalerLoadBalancerParams) SetLbdevicecapacity(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbdevicecapacity"] = v
}

func (P *ConfigureNetscalerLoadBalancerParams) GetLbdevicecapacity() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbdevicecapacity"].(int64)
	return value, ok
}

func (P *ConfigureNetscalerLoadBalancerParams) SetLbdevicededicated(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbdevicededicated"] = v
}

func (P *ConfigureNetscalerLoadBalancerParams) GetLbdevicededicated() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbdevicededicated"].(bool)
	return value, ok
}

func (P *ConfigureNetscalerLoadBalancerParams) SetLbdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbdeviceid"] = v
}

func (P *ConfigureNetscalerLoadBalancerParams) GetLbdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbdeviceid"].(string)
	return value, ok
}

func (P *ConfigureNetscalerLoadBalancerParams) SetPodids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podids"] = v
}

func (P *ConfigureNetscalerLoadBalancerParams) GetPodids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podids"].([]string)
	return value, ok
}

// You should always use this function to get a new ConfigureNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewConfigureNetscalerLoadBalancerParams(lbdeviceid string) *ConfigureNetscalerLoadBalancerParams {
	P := &ConfigureNetscalerLoadBalancerParams{}
	P.P = make(map[string]interface{})
	P.P["lbdeviceid"] = lbdeviceid
	return P
}

// configures a netscaler load balancer device
func (s *LoadBalancerService) ConfigureNetscalerLoadBalancer(p *ConfigureNetscalerLoadBalancerParams) (*NetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("configureNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r NetscalerLoadBalancerResponse
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

type NetscalerLoadBalancerResponse struct {
	Gslbprovider            bool     `json:"gslbprovider"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip"`
	Ipaddress               string   `json:"ipaddress"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider"`
	JobID                   string   `json:"jobid"`
	Jobstatus               int      `json:"jobstatus"`
	Lbdevicecapacity        int64    `json:"lbdevicecapacity"`
	Lbdevicededicated       bool     `json:"lbdevicededicated"`
	Lbdeviceid              string   `json:"lbdeviceid"`
	Lbdevicename            string   `json:"lbdevicename"`
	Lbdevicestate           string   `json:"lbdevicestate"`
	Physicalnetworkid       string   `json:"physicalnetworkid"`
	Podids                  []string `json:"podids"`
	Privateinterface        string   `json:"privateinterface"`
	Provider                string   `json:"provider"`
	Publicinterface         string   `json:"publicinterface"`
}

type CreateGlobalLoadBalancerRuleParams struct {
	P map[string]interface{}
}

func (P *CreateGlobalLoadBalancerRuleParams) toURLValues() url.Values {
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
	if v, found := P.P["gslbdomainname"]; found {
		u.Set("gslbdomainname", v.(string))
	}
	if v, found := P.P["gslblbmethod"]; found {
		u.Set("gslblbmethod", v.(string))
	}
	if v, found := P.P["gslbservicetype"]; found {
		u.Set("gslbservicetype", v.(string))
	}
	if v, found := P.P["gslbstickysessionmethodname"]; found {
		u.Set("gslbstickysessionmethodname", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
	}
	return u
}

func (P *CreateGlobalLoadBalancerRuleParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateGlobalLoadBalancerRuleParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateGlobalLoadBalancerRuleParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateGlobalLoadBalancerRuleParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateGlobalLoadBalancerRuleParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateGlobalLoadBalancerRuleParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateGlobalLoadBalancerRuleParams) SetGslbdomainname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gslbdomainname"] = v
}

func (P *CreateGlobalLoadBalancerRuleParams) GetGslbdomainname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gslbdomainname"].(string)
	return value, ok
}

func (P *CreateGlobalLoadBalancerRuleParams) SetGslblbmethod(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gslblbmethod"] = v
}

func (P *CreateGlobalLoadBalancerRuleParams) GetGslblbmethod() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gslblbmethod"].(string)
	return value, ok
}

func (P *CreateGlobalLoadBalancerRuleParams) SetGslbservicetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gslbservicetype"] = v
}

func (P *CreateGlobalLoadBalancerRuleParams) GetGslbservicetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gslbservicetype"].(string)
	return value, ok
}

func (P *CreateGlobalLoadBalancerRuleParams) SetGslbstickysessionmethodname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gslbstickysessionmethodname"] = v
}

func (P *CreateGlobalLoadBalancerRuleParams) GetGslbstickysessionmethodname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gslbstickysessionmethodname"].(string)
	return value, ok
}

func (P *CreateGlobalLoadBalancerRuleParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateGlobalLoadBalancerRuleParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateGlobalLoadBalancerRuleParams) SetRegionid(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["regionid"] = v
}

func (P *CreateGlobalLoadBalancerRuleParams) GetRegionid() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["regionid"].(int)
	return value, ok
}

// You should always use this function to get a new CreateGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateGlobalLoadBalancerRuleParams(gslbdomainname string, gslbservicetype string, name string, regionid int) *CreateGlobalLoadBalancerRuleParams {
	P := &CreateGlobalLoadBalancerRuleParams{}
	P.P = make(map[string]interface{})
	P.P["gslbdomainname"] = gslbdomainname
	P.P["gslbservicetype"] = gslbservicetype
	P.P["name"] = name
	P.P["regionid"] = regionid
	return P
}

// Creates a global load balancer rule
func (s *LoadBalancerService) CreateGlobalLoadBalancerRule(p *CreateGlobalLoadBalancerRuleParams) (*CreateGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("createGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateGlobalLoadBalancerRuleResponse
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

type CreateGlobalLoadBalancerRuleResponse struct {
	Account                     string                                                 `json:"account"`
	Description                 string                                                 `json:"description"`
	Domain                      string                                                 `json:"domain"`
	Domainid                    string                                                 `json:"domainid"`
	Gslbdomainname              string                                                 `json:"gslbdomainname"`
	Gslblbmethod                string                                                 `json:"gslblbmethod"`
	Gslbservicetype             string                                                 `json:"gslbservicetype"`
	Gslbstickysessionmethodname string                                                 `json:"gslbstickysessionmethodname"`
	Id                          string                                                 `json:"id"`
	JobID                       string                                                 `json:"jobid"`
	Jobstatus                   int                                                    `json:"jobstatus"`
	Loadbalancerrule            []CreateGlobalLoadBalancerRuleResponseLoadbalancerrule `json:"loadbalancerrule"`
	Name                        string                                                 `json:"name"`
	Project                     string                                                 `json:"project"`
	Projectid                   string                                                 `json:"projectid"`
	Regionid                    int                                                    `json:"regionid"`
}

type CreateGlobalLoadBalancerRuleResponseLoadbalancerrule struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type CreateLBHealthCheckPolicyParams struct {
	P map[string]interface{}
}

func (P *CreateLBHealthCheckPolicyParams) toURLValues() url.Values {
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
	if v, found := P.P["healthythreshold"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("healthythreshold", vv)
	}
	if v, found := P.P["intervaltime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("intervaltime", vv)
	}
	if v, found := P.P["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := P.P["pingpath"]; found {
		u.Set("pingpath", v.(string))
	}
	if v, found := P.P["responsetimeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("responsetimeout", vv)
	}
	if v, found := P.P["unhealthythreshold"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("unhealthythreshold", vv)
	}
	return u
}

func (P *CreateLBHealthCheckPolicyParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateLBHealthCheckPolicyParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateLBHealthCheckPolicyParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateLBHealthCheckPolicyParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateLBHealthCheckPolicyParams) SetHealthythreshold(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["healthythreshold"] = v
}

func (P *CreateLBHealthCheckPolicyParams) GetHealthythreshold() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["healthythreshold"].(int)
	return value, ok
}

func (P *CreateLBHealthCheckPolicyParams) SetIntervaltime(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["intervaltime"] = v
}

func (P *CreateLBHealthCheckPolicyParams) GetIntervaltime() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["intervaltime"].(int)
	return value, ok
}

func (P *CreateLBHealthCheckPolicyParams) SetLbruleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbruleid"] = v
}

func (P *CreateLBHealthCheckPolicyParams) GetLbruleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbruleid"].(string)
	return value, ok
}

func (P *CreateLBHealthCheckPolicyParams) SetPingpath(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pingpath"] = v
}

func (P *CreateLBHealthCheckPolicyParams) GetPingpath() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pingpath"].(string)
	return value, ok
}

func (P *CreateLBHealthCheckPolicyParams) SetResponsetimeout(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["responsetimeout"] = v
}

func (P *CreateLBHealthCheckPolicyParams) GetResponsetimeout() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["responsetimeout"].(int)
	return value, ok
}

func (P *CreateLBHealthCheckPolicyParams) SetUnhealthythreshold(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["unhealthythreshold"] = v
}

func (P *CreateLBHealthCheckPolicyParams) GetUnhealthythreshold() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["unhealthythreshold"].(int)
	return value, ok
}

// You should always use this function to get a new CreateLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLBHealthCheckPolicyParams(lbruleid string) *CreateLBHealthCheckPolicyParams {
	P := &CreateLBHealthCheckPolicyParams{}
	P.P = make(map[string]interface{})
	P.P["lbruleid"] = lbruleid
	return P
}

// Creates a load balancer health check policy
func (s *LoadBalancerService) CreateLBHealthCheckPolicy(p *CreateLBHealthCheckPolicyParams) (*CreateLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("createLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLBHealthCheckPolicyResponse
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

type CreateLBHealthCheckPolicyResponse struct {
	Account           string                                               `json:"account"`
	Domain            string                                               `json:"domain"`
	Domainid          string                                               `json:"domainid"`
	Healthcheckpolicy []CreateLBHealthCheckPolicyResponseHealthcheckpolicy `json:"healthcheckpolicy"`
	JobID             string                                               `json:"jobid"`
	Jobstatus         int                                                  `json:"jobstatus"`
	Lbruleid          string                                               `json:"lbruleid"`
	Zoneid            string                                               `json:"zoneid"`
}

type CreateLBHealthCheckPolicyResponseHealthcheckpolicy struct {
	Description             string `json:"description"`
	Fordisplay              bool   `json:"fordisplay"`
	Healthcheckinterval     int    `json:"healthcheckinterval"`
	Healthcheckthresshold   int    `json:"healthcheckthresshold"`
	Id                      string `json:"id"`
	Pingpath                string `json:"pingpath"`
	Responsetime            int    `json:"responsetime"`
	State                   string `json:"state"`
	Unhealthcheckthresshold int    `json:"unhealthcheckthresshold"`
}

type CreateLBStickinessPolicyParams struct {
	P map[string]interface{}
}

func (P *CreateLBStickinessPolicyParams) toURLValues() url.Values {
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
	if v, found := P.P["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := P.P["methodname"]; found {
		u.Set("methodname", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["param"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("param[%d].key", i), k)
			u.Set(fmt.Sprintf("param[%d].value", i), m[k])
		}
	}
	return u
}

func (P *CreateLBStickinessPolicyParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateLBStickinessPolicyParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateLBStickinessPolicyParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateLBStickinessPolicyParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateLBStickinessPolicyParams) SetLbruleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbruleid"] = v
}

func (P *CreateLBStickinessPolicyParams) GetLbruleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbruleid"].(string)
	return value, ok
}

func (P *CreateLBStickinessPolicyParams) SetMethodname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["methodname"] = v
}

func (P *CreateLBStickinessPolicyParams) GetMethodname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["methodname"].(string)
	return value, ok
}

func (P *CreateLBStickinessPolicyParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateLBStickinessPolicyParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateLBStickinessPolicyParams) SetParam(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["param"] = v
}

func (P *CreateLBStickinessPolicyParams) GetParam() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["param"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new CreateLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLBStickinessPolicyParams(lbruleid string, methodname string, name string) *CreateLBStickinessPolicyParams {
	P := &CreateLBStickinessPolicyParams{}
	P.P = make(map[string]interface{})
	P.P["lbruleid"] = lbruleid
	P.P["methodname"] = methodname
	P.P["name"] = name
	return P
}

// Creates a load balancer stickiness policy
func (s *LoadBalancerService) CreateLBStickinessPolicy(p *CreateLBStickinessPolicyParams) (*CreateLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("createLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLBStickinessPolicyResponse
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

type CreateLBStickinessPolicyResponse struct {
	Account          string                                             `json:"account"`
	Description      string                                             `json:"description"`
	Domain           string                                             `json:"domain"`
	Domainid         string                                             `json:"domainid"`
	JobID            string                                             `json:"jobid"`
	Jobstatus        int                                                `json:"jobstatus"`
	Lbruleid         string                                             `json:"lbruleid"`
	Name             string                                             `json:"name"`
	State            string                                             `json:"state"`
	Stickinesspolicy []CreateLBStickinessPolicyResponseStickinesspolicy `json:"stickinesspolicy"`
	Zoneid           string                                             `json:"zoneid"`
}

type CreateLBStickinessPolicyResponseStickinesspolicy struct {
	Description string            `json:"description"`
	Fordisplay  bool              `json:"fordisplay"`
	Id          string            `json:"id"`
	Methodname  string            `json:"methodname"`
	Name        string            `json:"name"`
	Params      map[string]string `json:"params"`
	State       string            `json:"state"`
}

type CreateLoadBalancerParams struct {
	P map[string]interface{}
}

func (P *CreateLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["instanceport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("instanceport", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := P.P["scheme"]; found {
		u.Set("scheme", v.(string))
	}
	if v, found := P.P["sourceipaddress"]; found {
		u.Set("sourceipaddress", v.(string))
	}
	if v, found := P.P["sourceipaddressnetworkid"]; found {
		u.Set("sourceipaddressnetworkid", v.(string))
	}
	if v, found := P.P["sourceport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sourceport", vv)
	}
	return u
}

func (P *CreateLoadBalancerParams) SetAlgorithm(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["algorithm"] = v
}

func (P *CreateLoadBalancerParams) GetAlgorithm() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["algorithm"].(string)
	return value, ok
}

func (P *CreateLoadBalancerParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateLoadBalancerParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateLoadBalancerParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateLoadBalancerParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateLoadBalancerParams) SetInstanceport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["instanceport"] = v
}

func (P *CreateLoadBalancerParams) GetInstanceport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["instanceport"].(int)
	return value, ok
}

func (P *CreateLoadBalancerParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateLoadBalancerParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateLoadBalancerParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *CreateLoadBalancerParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *CreateLoadBalancerParams) SetScheme(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["scheme"] = v
}

func (P *CreateLoadBalancerParams) GetScheme() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["scheme"].(string)
	return value, ok
}

func (P *CreateLoadBalancerParams) SetSourceipaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sourceipaddress"] = v
}

func (P *CreateLoadBalancerParams) GetSourceipaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sourceipaddress"].(string)
	return value, ok
}

func (P *CreateLoadBalancerParams) SetSourceipaddressnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sourceipaddressnetworkid"] = v
}

func (P *CreateLoadBalancerParams) GetSourceipaddressnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sourceipaddressnetworkid"].(string)
	return value, ok
}

func (P *CreateLoadBalancerParams) SetSourceport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sourceport"] = v
}

func (P *CreateLoadBalancerParams) GetSourceport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sourceport"].(int)
	return value, ok
}

// You should always use this function to get a new CreateLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLoadBalancerParams(algorithm string, instanceport int, name string, networkid string, scheme string, sourceipaddressnetworkid string, sourceport int) *CreateLoadBalancerParams {
	P := &CreateLoadBalancerParams{}
	P.P = make(map[string]interface{})
	P.P["algorithm"] = algorithm
	P.P["instanceport"] = instanceport
	P.P["name"] = name
	P.P["networkid"] = networkid
	P.P["scheme"] = scheme
	P.P["sourceipaddressnetworkid"] = sourceipaddressnetworkid
	P.P["sourceport"] = sourceport
	return P
}

// Creates a load balancer
func (s *LoadBalancerService) CreateLoadBalancer(p *CreateLoadBalancerParams) (*CreateLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("createLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLoadBalancerResponse
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

type CreateLoadBalancerResponse struct {
	Account                  string                                           `json:"account"`
	Algorithm                string                                           `json:"algorithm"`
	Description              string                                           `json:"description"`
	Domain                   string                                           `json:"domain"`
	Domainid                 string                                           `json:"domainid"`
	Fordisplay               bool                                             `json:"fordisplay"`
	Id                       string                                           `json:"id"`
	JobID                    string                                           `json:"jobid"`
	Jobstatus                int                                              `json:"jobstatus"`
	Loadbalancerinstance     []CreateLoadBalancerResponseLoadbalancerinstance `json:"loadbalancerinstance"`
	Loadbalancerrule         []CreateLoadBalancerResponseLoadbalancerrule     `json:"loadbalancerrule"`
	Name                     string                                           `json:"name"`
	Networkid                string                                           `json:"networkid"`
	Project                  string                                           `json:"project"`
	Projectid                string                                           `json:"projectid"`
	Sourceipaddress          string                                           `json:"sourceipaddress"`
	Sourceipaddressnetworkid string                                           `json:"sourceipaddressnetworkid"`
	Tags                     []Tags                                           `json:"tags"`
}

type CreateLoadBalancerResponseLoadbalancerrule struct {
	Instanceport int    `json:"instanceport"`
	Sourceport   int    `json:"sourceport"`
	State        string `json:"state"`
}

type CreateLoadBalancerResponseLoadbalancerinstance struct {
	Id        string `json:"id"`
	Ipaddress string `json:"ipaddress"`
	Name      string `json:"name"`
	State     string `json:"state"`
}

type CreateLoadBalancerRuleParams struct {
	P map[string]interface{}
}

func (P *CreateLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := P.P["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := P.P["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := P.P["privateport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateport", vv)
	}
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := P.P["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	if v, found := P.P["publicport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicport", vv)
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *CreateLoadBalancerRuleParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateLoadBalancerRuleParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetAlgorithm(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["algorithm"] = v
}

func (P *CreateLoadBalancerRuleParams) GetAlgorithm() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["algorithm"].(string)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetCidrlist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *CreateLoadBalancerRuleParams) GetCidrlist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].([]string)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateLoadBalancerRuleParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateLoadBalancerRuleParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateLoadBalancerRuleParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateLoadBalancerRuleParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *CreateLoadBalancerRuleParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetOpenfirewall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["openfirewall"] = v
}

func (P *CreateLoadBalancerRuleParams) GetOpenfirewall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["openfirewall"].(bool)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetPrivateport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["privateport"] = v
}

func (P *CreateLoadBalancerRuleParams) GetPrivateport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["privateport"].(int)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *CreateLoadBalancerRuleParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetPublicipid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publicipid"] = v
}

func (P *CreateLoadBalancerRuleParams) GetPublicipid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publicipid"].(string)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetPublicport(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publicport"] = v
}

func (P *CreateLoadBalancerRuleParams) GetPublicport() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publicport"].(int)
	return value, ok
}

func (P *CreateLoadBalancerRuleParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateLoadBalancerRuleParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLoadBalancerRuleParams(algorithm string, name string, privateport int, publicport int) *CreateLoadBalancerRuleParams {
	P := &CreateLoadBalancerRuleParams{}
	P.P = make(map[string]interface{})
	P.P["algorithm"] = algorithm
	P.P["name"] = name
	P.P["privateport"] = privateport
	P.P["publicport"] = publicport
	return P
}

// Creates a load balancer rule
func (s *LoadBalancerService) CreateLoadBalancerRule(p *CreateLoadBalancerRuleParams) (*CreateLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("createLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLoadBalancerRuleResponse
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

type CreateLoadBalancerRuleResponse struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type DeleteGlobalLoadBalancerRuleParams struct {
	P map[string]interface{}
}

func (P *DeleteGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteGlobalLoadBalancerRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteGlobalLoadBalancerRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteGlobalLoadBalancerRuleParams(id string) *DeleteGlobalLoadBalancerRuleParams {
	P := &DeleteGlobalLoadBalancerRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a global load balancer rule.
func (s *LoadBalancerService) DeleteGlobalLoadBalancerRule(p *DeleteGlobalLoadBalancerRuleParams) (*DeleteGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteGlobalLoadBalancerRuleResponse
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

type DeleteGlobalLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLBHealthCheckPolicyParams struct {
	P map[string]interface{}
}

func (P *DeleteLBHealthCheckPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteLBHealthCheckPolicyParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteLBHealthCheckPolicyParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLBHealthCheckPolicyParams(id string) *DeleteLBHealthCheckPolicyParams {
	P := &DeleteLBHealthCheckPolicyParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a load balancer health check policy.
func (s *LoadBalancerService) DeleteLBHealthCheckPolicy(p *DeleteLBHealthCheckPolicyParams) (*DeleteLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLBHealthCheckPolicyResponse
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

type DeleteLBHealthCheckPolicyResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLBStickinessPolicyParams struct {
	P map[string]interface{}
}

func (P *DeleteLBStickinessPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteLBStickinessPolicyParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteLBStickinessPolicyParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLBStickinessPolicyParams(id string) *DeleteLBStickinessPolicyParams {
	P := &DeleteLBStickinessPolicyParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a load balancer stickiness policy.
func (s *LoadBalancerService) DeleteLBStickinessPolicy(p *DeleteLBStickinessPolicyParams) (*DeleteLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLBStickinessPolicyResponse
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

type DeleteLBStickinessPolicyResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLoadBalancerParams struct {
	P map[string]interface{}
}

func (P *DeleteLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteLoadBalancerParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteLoadBalancerParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLoadBalancerParams(id string) *DeleteLoadBalancerParams {
	P := &DeleteLoadBalancerParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a load balancer
func (s *LoadBalancerService) DeleteLoadBalancer(p *DeleteLoadBalancerParams) (*DeleteLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLoadBalancerResponse
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

type DeleteLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteLoadBalancerRuleParams struct {
	P map[string]interface{}
}

func (P *DeleteLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteLoadBalancerRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteLoadBalancerRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLoadBalancerRuleParams(id string) *DeleteLoadBalancerRuleParams {
	P := &DeleteLoadBalancerRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a load balancer rule.
func (s *LoadBalancerService) DeleteLoadBalancerRule(p *DeleteLoadBalancerRuleParams) (*DeleteLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLoadBalancerRuleResponse
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

type DeleteLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteNetscalerLoadBalancerParams struct {
	P map[string]interface{}
}

func (P *DeleteNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	return u
}

func (P *DeleteNetscalerLoadBalancerParams) SetLbdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbdeviceid"] = v
}

func (P *DeleteNetscalerLoadBalancerParams) GetLbdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbdeviceid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteNetscalerLoadBalancerParams(lbdeviceid string) *DeleteNetscalerLoadBalancerParams {
	P := &DeleteNetscalerLoadBalancerParams{}
	P.P = make(map[string]interface{})
	P.P["lbdeviceid"] = lbdeviceid
	return P
}

//  delete a netscaler load balancer device
func (s *LoadBalancerService) DeleteNetscalerLoadBalancer(p *DeleteNetscalerLoadBalancerParams) (*DeleteNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetscalerLoadBalancerResponse
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

type DeleteNetscalerLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteSslCertParams struct {
	P map[string]interface{}
}

func (P *DeleteSslCertParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteSslCertParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteSslCertParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteSslCertParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteSslCertParams(id string) *DeleteSslCertParams {
	P := &DeleteSslCertParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Delete a certificate to CloudStack
func (s *LoadBalancerService) DeleteSslCert(p *DeleteSslCertParams) (*DeleteSslCertResponse, error) {
	resp, err := s.cs.newRequest("deleteSslCert", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSslCertResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteSslCertResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteSslCertResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteSslCertResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListGlobalLoadBalancerRulesParams struct {
	P map[string]interface{}
}

func (P *ListGlobalLoadBalancerRulesParams) toURLValues() url.Values {
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
	if v, found := P.P["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
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

func (P *ListGlobalLoadBalancerRulesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListGlobalLoadBalancerRulesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListGlobalLoadBalancerRulesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListGlobalLoadBalancerRulesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListGlobalLoadBalancerRulesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListGlobalLoadBalancerRulesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListGlobalLoadBalancerRulesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListGlobalLoadBalancerRulesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListGlobalLoadBalancerRulesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListGlobalLoadBalancerRulesParams) SetRegionid(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["regionid"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetRegionid() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["regionid"].(int)
	return value, ok
}

func (P *ListGlobalLoadBalancerRulesParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListGlobalLoadBalancerRulesParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListGlobalLoadBalancerRulesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListGlobalLoadBalancerRulesParams() *ListGlobalLoadBalancerRulesParams {
	P := &ListGlobalLoadBalancerRulesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleID(keyword string, opts ...OptionFunc) (string, int, error) {
	P := &ListGlobalLoadBalancerRulesParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListGlobalLoadBalancerRules(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.GlobalLoadBalancerRules[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.GlobalLoadBalancerRules {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleByName(name string, opts ...OptionFunc) (*GlobalLoadBalancerRule, int, error) {
	id, count, err := s.GetGlobalLoadBalancerRuleID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetGlobalLoadBalancerRuleByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleByID(id string, opts ...OptionFunc) (*GlobalLoadBalancerRule, int, error) {
	P := &ListGlobalLoadBalancerRulesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListGlobalLoadBalancerRules(P)
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
		return l.GlobalLoadBalancerRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for GlobalLoadBalancerRule UUID: %s!", id)
}

// Lists load balancer rules.
func (s *LoadBalancerService) ListGlobalLoadBalancerRules(p *ListGlobalLoadBalancerRulesParams) (*ListGlobalLoadBalancerRulesResponse, error) {
	resp, err := s.cs.newRequest("listGlobalLoadBalancerRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListGlobalLoadBalancerRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListGlobalLoadBalancerRulesResponse struct {
	Count                   int                       `json:"count"`
	GlobalLoadBalancerRules []*GlobalLoadBalancerRule `json:"globalloadbalancerrule"`
}

type GlobalLoadBalancerRule struct {
	Account                     string                                   `json:"account"`
	Description                 string                                   `json:"description"`
	Domain                      string                                   `json:"domain"`
	Domainid                    string                                   `json:"domainid"`
	Gslbdomainname              string                                   `json:"gslbdomainname"`
	Gslblbmethod                string                                   `json:"gslblbmethod"`
	Gslbservicetype             string                                   `json:"gslbservicetype"`
	Gslbstickysessionmethodname string                                   `json:"gslbstickysessionmethodname"`
	Id                          string                                   `json:"id"`
	JobID                       string                                   `json:"jobid"`
	Jobstatus                   int                                      `json:"jobstatus"`
	Loadbalancerrule            []GlobalLoadBalancerRuleLoadbalancerrule `json:"loadbalancerrule"`
	Name                        string                                   `json:"name"`
	Project                     string                                   `json:"project"`
	Projectid                   string                                   `json:"projectid"`
	Regionid                    int                                      `json:"regionid"`
}

type GlobalLoadBalancerRuleLoadbalancerrule struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type ListLBHealthCheckPoliciesParams struct {
	P map[string]interface{}
}

func (P *ListLBHealthCheckPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (P *ListLBHealthCheckPoliciesParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListLBHealthCheckPoliciesParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListLBHealthCheckPoliciesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListLBHealthCheckPoliciesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListLBHealthCheckPoliciesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListLBHealthCheckPoliciesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListLBHealthCheckPoliciesParams) SetLbruleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbruleid"] = v
}

func (P *ListLBHealthCheckPoliciesParams) GetLbruleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbruleid"].(string)
	return value, ok
}

func (P *ListLBHealthCheckPoliciesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListLBHealthCheckPoliciesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListLBHealthCheckPoliciesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListLBHealthCheckPoliciesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListLBHealthCheckPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLBHealthCheckPoliciesParams() *ListLBHealthCheckPoliciesParams {
	P := &ListLBHealthCheckPoliciesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLBHealthCheckPolicyByID(id string, opts ...OptionFunc) (*LBHealthCheckPolicy, int, error) {
	P := &ListLBHealthCheckPoliciesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLBHealthCheckPolicies(P)
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
		return l.LBHealthCheckPolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LBHealthCheckPolicy UUID: %s!", id)
}

// Lists load balancer health check policies.
func (s *LoadBalancerService) ListLBHealthCheckPolicies(p *ListLBHealthCheckPoliciesParams) (*ListLBHealthCheckPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listLBHealthCheckPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLBHealthCheckPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLBHealthCheckPoliciesResponse struct {
	Count                 int                    `json:"count"`
	LBHealthCheckPolicies []*LBHealthCheckPolicy `json:"lbhealthcheckpolicy"`
}

type LBHealthCheckPolicy struct {
	Account           string                                 `json:"account"`
	Domain            string                                 `json:"domain"`
	Domainid          string                                 `json:"domainid"`
	Healthcheckpolicy []LBHealthCheckPolicyHealthcheckpolicy `json:"healthcheckpolicy"`
	JobID             string                                 `json:"jobid"`
	Jobstatus         int                                    `json:"jobstatus"`
	Lbruleid          string                                 `json:"lbruleid"`
	Zoneid            string                                 `json:"zoneid"`
}

type LBHealthCheckPolicyHealthcheckpolicy struct {
	Description             string `json:"description"`
	Fordisplay              bool   `json:"fordisplay"`
	Healthcheckinterval     int    `json:"healthcheckinterval"`
	Healthcheckthresshold   int    `json:"healthcheckthresshold"`
	Id                      string `json:"id"`
	Pingpath                string `json:"pingpath"`
	Responsetime            int    `json:"responsetime"`
	State                   string `json:"state"`
	Unhealthcheckthresshold int    `json:"unhealthcheckthresshold"`
}

type ListLBStickinessPoliciesParams struct {
	P map[string]interface{}
}

func (P *ListLBStickinessPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (P *ListLBStickinessPoliciesParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListLBStickinessPoliciesParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListLBStickinessPoliciesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListLBStickinessPoliciesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListLBStickinessPoliciesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListLBStickinessPoliciesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListLBStickinessPoliciesParams) SetLbruleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbruleid"] = v
}

func (P *ListLBStickinessPoliciesParams) GetLbruleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbruleid"].(string)
	return value, ok
}

func (P *ListLBStickinessPoliciesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListLBStickinessPoliciesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListLBStickinessPoliciesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListLBStickinessPoliciesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListLBStickinessPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLBStickinessPoliciesParams() *ListLBStickinessPoliciesParams {
	P := &ListLBStickinessPoliciesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLBStickinessPolicyByID(id string, opts ...OptionFunc) (*LBStickinessPolicy, int, error) {
	P := &ListLBStickinessPoliciesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLBStickinessPolicies(P)
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
		return l.LBStickinessPolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LBStickinessPolicy UUID: %s!", id)
}

// Lists load balancer stickiness policies.
func (s *LoadBalancerService) ListLBStickinessPolicies(p *ListLBStickinessPoliciesParams) (*ListLBStickinessPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listLBStickinessPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLBStickinessPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLBStickinessPoliciesResponse struct {
	Count                int                   `json:"count"`
	LBStickinessPolicies []*LBStickinessPolicy `json:"lbstickinesspolicy"`
}

type LBStickinessPolicy struct {
	Account          string                               `json:"account"`
	Description      string                               `json:"description"`
	Domain           string                               `json:"domain"`
	Domainid         string                               `json:"domainid"`
	JobID            string                               `json:"jobid"`
	Jobstatus        int                                  `json:"jobstatus"`
	Lbruleid         string                               `json:"lbruleid"`
	Name             string                               `json:"name"`
	State            string                               `json:"state"`
	Stickinesspolicy []LBStickinessPolicyStickinesspolicy `json:"stickinesspolicy"`
	Zoneid           string                               `json:"zoneid"`
}

type LBStickinessPolicyStickinesspolicy struct {
	Description string            `json:"description"`
	Fordisplay  bool              `json:"fordisplay"`
	Id          string            `json:"id"`
	Methodname  string            `json:"methodname"`
	Name        string            `json:"name"`
	Params      map[string]string `json:"params"`
	State       string            `json:"state"`
}

type ListLoadBalancerRuleInstancesParams struct {
	P map[string]interface{}
}

func (P *ListLoadBalancerRuleInstancesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["applied"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("applied", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["lbvmips"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("lbvmips", vv)
	}
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (P *ListLoadBalancerRuleInstancesParams) SetApplied(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["applied"] = v
}

func (P *ListLoadBalancerRuleInstancesParams) GetApplied() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["applied"].(bool)
	return value, ok
}

func (P *ListLoadBalancerRuleInstancesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListLoadBalancerRuleInstancesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListLoadBalancerRuleInstancesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListLoadBalancerRuleInstancesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListLoadBalancerRuleInstancesParams) SetLbvmips(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbvmips"] = v
}

func (P *ListLoadBalancerRuleInstancesParams) GetLbvmips() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbvmips"].(bool)
	return value, ok
}

func (P *ListLoadBalancerRuleInstancesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListLoadBalancerRuleInstancesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListLoadBalancerRuleInstancesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListLoadBalancerRuleInstancesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListLoadBalancerRuleInstancesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancerRuleInstancesParams(id string) *ListLoadBalancerRuleInstancesParams {
	P := &ListLoadBalancerRuleInstancesParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleInstanceByID(id string, opts ...OptionFunc) (*VirtualMachine, int, error) {
	P := &ListLoadBalancerRuleInstancesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLoadBalancerRuleInstances(P)
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
		return l.LoadBalancerRuleInstances[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancerRuleInstance UUID: %s!", id)
}

// List all virtual machine instances that are assigned to a load balancer rule.
func (s *LoadBalancerService) ListLoadBalancerRuleInstances(p *ListLoadBalancerRuleInstancesParams) (*ListLoadBalancerRuleInstancesResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancerRuleInstances", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancerRuleInstancesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLoadBalancerRuleInstancesResponse struct {
	Count                     int                         `json:"count"`
	LBRuleVMIDIPs             []*LoadBalancerRuleInstance `json:"lbrulevmidip"`
	LoadBalancerRuleInstances []*VirtualMachine           `json:"loadbalancerruleinstance"`
}

type LoadBalancerRuleInstance struct {
	JobID                    string          `json:"jobid"`
	Jobstatus                int             `json:"jobstatus"`
	Lbvmipaddresses          []string        `json:"lbvmipaddresses"`
	Loadbalancerruleinstance *VirtualMachine `json:"loadbalancerruleinstance"`
}

type ListLoadBalancerRulesParams struct {
	P map[string]interface{}
}

func (P *ListLoadBalancerRulesParams) toURLValues() url.Values {
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
	if v, found := P.P["publicipid"]; found {
		u.Set("publicipid", v.(string))
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
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListLoadBalancerRulesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListLoadBalancerRulesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListLoadBalancerRulesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListLoadBalancerRulesParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListLoadBalancerRulesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListLoadBalancerRulesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListLoadBalancerRulesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListLoadBalancerRulesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListLoadBalancerRulesParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListLoadBalancerRulesParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListLoadBalancerRulesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListLoadBalancerRulesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListLoadBalancerRulesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetPublicipid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publicipid"] = v
}

func (P *ListLoadBalancerRulesParams) GetPublicipid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publicipid"].(string)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListLoadBalancerRulesParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *ListLoadBalancerRulesParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *ListLoadBalancerRulesParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListLoadBalancerRulesParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListLoadBalancerRulesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancerRulesParams() *ListLoadBalancerRulesParams {
	P := &ListLoadBalancerRulesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListLoadBalancerRulesParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListLoadBalancerRules(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.LoadBalancerRules[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.LoadBalancerRules {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleByName(name string, opts ...OptionFunc) (*LoadBalancerRule, int, error) {
	id, count, err := s.GetLoadBalancerRuleID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetLoadBalancerRuleByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleByID(id string, opts ...OptionFunc) (*LoadBalancerRule, int, error) {
	P := &ListLoadBalancerRulesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLoadBalancerRules(P)
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
		return l.LoadBalancerRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancerRule UUID: %s!", id)
}

// Lists load balancer rules.
func (s *LoadBalancerService) ListLoadBalancerRules(p *ListLoadBalancerRulesParams) (*ListLoadBalancerRulesResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancerRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancerRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLoadBalancerRulesResponse struct {
	Count             int                 `json:"count"`
	LoadBalancerRules []*LoadBalancerRule `json:"loadbalancerrule"`
}

type LoadBalancerRule struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type ListLoadBalancersParams struct {
	P map[string]interface{}
}

func (P *ListLoadBalancersParams) toURLValues() url.Values {
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
	if v, found := P.P["scheme"]; found {
		u.Set("scheme", v.(string))
	}
	if v, found := P.P["sourceipaddress"]; found {
		u.Set("sourceipaddress", v.(string))
	}
	if v, found := P.P["sourceipaddressnetworkid"]; found {
		u.Set("sourceipaddressnetworkid", v.(string))
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

func (P *ListLoadBalancersParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListLoadBalancersParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListLoadBalancersParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListLoadBalancersParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListLoadBalancersParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListLoadBalancersParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListLoadBalancersParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListLoadBalancersParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListLoadBalancersParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListLoadBalancersParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListLoadBalancersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListLoadBalancersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListLoadBalancersParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListLoadBalancersParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListLoadBalancersParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListLoadBalancersParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListLoadBalancersParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListLoadBalancersParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListLoadBalancersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListLoadBalancersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListLoadBalancersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListLoadBalancersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListLoadBalancersParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListLoadBalancersParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListLoadBalancersParams) SetScheme(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["scheme"] = v
}

func (P *ListLoadBalancersParams) GetScheme() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["scheme"].(string)
	return value, ok
}

func (P *ListLoadBalancersParams) SetSourceipaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sourceipaddress"] = v
}

func (P *ListLoadBalancersParams) GetSourceipaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sourceipaddress"].(string)
	return value, ok
}

func (P *ListLoadBalancersParams) SetSourceipaddressnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sourceipaddressnetworkid"] = v
}

func (P *ListLoadBalancersParams) GetSourceipaddressnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sourceipaddressnetworkid"].(string)
	return value, ok
}

func (P *ListLoadBalancersParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListLoadBalancersParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new ListLoadBalancersParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancersParams() *ListLoadBalancersParams {
	P := &ListLoadBalancersParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListLoadBalancersParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListLoadBalancers(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.LoadBalancers[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.LoadBalancers {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerByName(name string, opts ...OptionFunc) (*LoadBalancer, int, error) {
	id, count, err := s.GetLoadBalancerID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetLoadBalancerByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerByID(id string, opts ...OptionFunc) (*LoadBalancer, int, error) {
	P := &ListLoadBalancersParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListLoadBalancers(P)
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
		return l.LoadBalancers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancer UUID: %s!", id)
}

// Lists load balancers
func (s *LoadBalancerService) ListLoadBalancers(p *ListLoadBalancersParams) (*ListLoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListLoadBalancersResponse struct {
	Count         int             `json:"count"`
	LoadBalancers []*LoadBalancer `json:"loadbalancer"`
}

type LoadBalancer struct {
	Account                  string                             `json:"account"`
	Algorithm                string                             `json:"algorithm"`
	Description              string                             `json:"description"`
	Domain                   string                             `json:"domain"`
	Domainid                 string                             `json:"domainid"`
	Fordisplay               bool                               `json:"fordisplay"`
	Id                       string                             `json:"id"`
	JobID                    string                             `json:"jobid"`
	Jobstatus                int                                `json:"jobstatus"`
	Loadbalancerinstance     []LoadBalancerLoadbalancerinstance `json:"loadbalancerinstance"`
	Loadbalancerrule         []LoadBalancerLoadbalancerrule     `json:"loadbalancerrule"`
	Name                     string                             `json:"name"`
	Networkid                string                             `json:"networkid"`
	Project                  string                             `json:"project"`
	Projectid                string                             `json:"projectid"`
	Sourceipaddress          string                             `json:"sourceipaddress"`
	Sourceipaddressnetworkid string                             `json:"sourceipaddressnetworkid"`
	Tags                     []Tags                             `json:"tags"`
}

type LoadBalancerLoadbalancerrule struct {
	Instanceport int    `json:"instanceport"`
	Sourceport   int    `json:"sourceport"`
	State        string `json:"state"`
}

type LoadBalancerLoadbalancerinstance struct {
	Id        string `json:"id"`
	Ipaddress string `json:"ipaddress"`
	Name      string `json:"name"`
	State     string `json:"state"`
}

type ListNetscalerLoadBalancersParams struct {
	P map[string]interface{}
}

func (P *ListNetscalerLoadBalancersParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
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

func (P *ListNetscalerLoadBalancersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNetscalerLoadBalancersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNetscalerLoadBalancersParams) SetLbdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbdeviceid"] = v
}

func (P *ListNetscalerLoadBalancersParams) GetLbdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbdeviceid"].(string)
	return value, ok
}

func (P *ListNetscalerLoadBalancersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNetscalerLoadBalancersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNetscalerLoadBalancersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNetscalerLoadBalancersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListNetscalerLoadBalancersParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListNetscalerLoadBalancersParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetscalerLoadBalancersParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListNetscalerLoadBalancersParams() *ListNetscalerLoadBalancersParams {
	P := &ListNetscalerLoadBalancersParams{}
	P.P = make(map[string]interface{})
	return P
}

// lists netscaler load balancer devices
func (s *LoadBalancerService) ListNetscalerLoadBalancers(p *ListNetscalerLoadBalancersParams) (*ListNetscalerLoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listNetscalerLoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetscalerLoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetscalerLoadBalancersResponse struct {
	Count                  int                      `json:"count"`
	NetscalerLoadBalancers []*NetscalerLoadBalancer `json:"netscalerloadbalancer"`
}

type NetscalerLoadBalancer struct {
	Gslbprovider            bool     `json:"gslbprovider"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip"`
	Ipaddress               string   `json:"ipaddress"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider"`
	JobID                   string   `json:"jobid"`
	Jobstatus               int      `json:"jobstatus"`
	Lbdevicecapacity        int64    `json:"lbdevicecapacity"`
	Lbdevicededicated       bool     `json:"lbdevicededicated"`
	Lbdeviceid              string   `json:"lbdeviceid"`
	Lbdevicename            string   `json:"lbdevicename"`
	Lbdevicestate           string   `json:"lbdevicestate"`
	Physicalnetworkid       string   `json:"physicalnetworkid"`
	Podids                  []string `json:"podids"`
	Privateinterface        string   `json:"privateinterface"`
	Provider                string   `json:"provider"`
	Publicinterface         string   `json:"publicinterface"`
}

type ListSslCertsParams struct {
	P map[string]interface{}
}

func (P *ListSslCertsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := P.P["certid"]; found {
		u.Set("certid", v.(string))
	}
	if v, found := P.P["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *ListSslCertsParams) SetAccountid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountid"] = v
}

func (P *ListSslCertsParams) GetAccountid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountid"].(string)
	return value, ok
}

func (P *ListSslCertsParams) SetCertid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["certid"] = v
}

func (P *ListSslCertsParams) GetCertid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["certid"].(string)
	return value, ok
}

func (P *ListSslCertsParams) SetLbruleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbruleid"] = v
}

func (P *ListSslCertsParams) GetLbruleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbruleid"].(string)
	return value, ok
}

func (P *ListSslCertsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListSslCertsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSslCertsParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListSslCertsParams() *ListSslCertsParams {
	P := &ListSslCertsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists SSL certificates
func (s *LoadBalancerService) ListSslCerts(p *ListSslCertsParams) (*ListSslCertsResponse, error) {
	resp, err := s.cs.newRequest("listSslCerts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSslCertsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSslCertsResponse struct {
	Count    int        `json:"count"`
	SslCerts []*SslCert `json:"sslcert"`
}

type SslCert struct {
	Account              string   `json:"account"`
	Certchain            string   `json:"certchain"`
	Certificate          string   `json:"certificate"`
	Domain               string   `json:"domain"`
	Domainid             string   `json:"domainid"`
	Fingerprint          string   `json:"fingerprint"`
	Id                   string   `json:"id"`
	JobID                string   `json:"jobid"`
	Jobstatus            int      `json:"jobstatus"`
	Loadbalancerrulelist []string `json:"loadbalancerrulelist"`
	Name                 string   `json:"name"`
	Project              string   `json:"project"`
	Projectid            string   `json:"projectid"`
}

type RemoveCertFromLoadBalancerParams struct {
	P map[string]interface{}
}

func (P *RemoveCertFromLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	return u
}

func (P *RemoveCertFromLoadBalancerParams) SetLbruleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbruleid"] = v
}

func (P *RemoveCertFromLoadBalancerParams) GetLbruleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbruleid"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveCertFromLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveCertFromLoadBalancerParams(lbruleid string) *RemoveCertFromLoadBalancerParams {
	P := &RemoveCertFromLoadBalancerParams{}
	P.P = make(map[string]interface{})
	P.P["lbruleid"] = lbruleid
	return P
}

// Removes a certificate from a load balancer rule
func (s *LoadBalancerService) RemoveCertFromLoadBalancer(p *RemoveCertFromLoadBalancerParams) (*RemoveCertFromLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("removeCertFromLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveCertFromLoadBalancerResponse
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

type RemoveCertFromLoadBalancerResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RemoveFromGlobalLoadBalancerRuleParams struct {
	P map[string]interface{}
}

func (P *RemoveFromGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["loadbalancerrulelist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("loadbalancerrulelist", vv)
	}
	return u
}

func (P *RemoveFromGlobalLoadBalancerRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RemoveFromGlobalLoadBalancerRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *RemoveFromGlobalLoadBalancerRuleParams) SetLoadbalancerrulelist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["loadbalancerrulelist"] = v
}

func (P *RemoveFromGlobalLoadBalancerRuleParams) GetLoadbalancerrulelist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["loadbalancerrulelist"].([]string)
	return value, ok
}

// You should always use this function to get a new RemoveFromGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveFromGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *RemoveFromGlobalLoadBalancerRuleParams {
	P := &RemoveFromGlobalLoadBalancerRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["loadbalancerrulelist"] = loadbalancerrulelist
	return P
}

// Removes a load balancer rule association with global load balancer rule
func (s *LoadBalancerService) RemoveFromGlobalLoadBalancerRule(p *RemoveFromGlobalLoadBalancerRuleParams) (*RemoveFromGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("removeFromGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveFromGlobalLoadBalancerRuleResponse
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

type RemoveFromGlobalLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RemoveFromLoadBalancerRuleParams struct {
	P map[string]interface{}
}

func (P *RemoveFromLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["virtualmachineids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("virtualmachineids", vv)
	}
	if v, found := P.P["vmidipmap"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("vmidipmap[%d].key", i), k)
			u.Set(fmt.Sprintf("vmidipmap[%d].value", i), m[k])
		}
	}
	return u
}

func (P *RemoveFromLoadBalancerRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RemoveFromLoadBalancerRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *RemoveFromLoadBalancerRuleParams) SetVirtualmachineids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineids"] = v
}

func (P *RemoveFromLoadBalancerRuleParams) GetVirtualmachineids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineids"].([]string)
	return value, ok
}

func (P *RemoveFromLoadBalancerRuleParams) SetVmidipmap(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmidipmap"] = v
}

func (P *RemoveFromLoadBalancerRuleParams) GetVmidipmap() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmidipmap"].(map[string]string)
	return value, ok
}

// You should always use this function to get a new RemoveFromLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveFromLoadBalancerRuleParams(id string) *RemoveFromLoadBalancerRuleParams {
	P := &RemoveFromLoadBalancerRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Removes a virtual machine or a list of virtual machines from a load balancer rule.
func (s *LoadBalancerService) RemoveFromLoadBalancerRule(p *RemoveFromLoadBalancerRuleParams) (*RemoveFromLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("removeFromLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveFromLoadBalancerRuleResponse
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

type RemoveFromLoadBalancerRuleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateGlobalLoadBalancerRuleParams struct {
	P map[string]interface{}
}

func (P *UpdateGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["gslblbmethod"]; found {
		u.Set("gslblbmethod", v.(string))
	}
	if v, found := P.P["gslbstickysessionmethodname"]; found {
		u.Set("gslbstickysessionmethodname", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *UpdateGlobalLoadBalancerRuleParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *UpdateGlobalLoadBalancerRuleParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *UpdateGlobalLoadBalancerRuleParams) SetGslblbmethod(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gslblbmethod"] = v
}

func (P *UpdateGlobalLoadBalancerRuleParams) GetGslblbmethod() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gslblbmethod"].(string)
	return value, ok
}

func (P *UpdateGlobalLoadBalancerRuleParams) SetGslbstickysessionmethodname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gslbstickysessionmethodname"] = v
}

func (P *UpdateGlobalLoadBalancerRuleParams) GetGslbstickysessionmethodname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gslbstickysessionmethodname"].(string)
	return value, ok
}

func (P *UpdateGlobalLoadBalancerRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateGlobalLoadBalancerRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateGlobalLoadBalancerRuleParams(id string) *UpdateGlobalLoadBalancerRuleParams {
	P := &UpdateGlobalLoadBalancerRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// update global load balancer rules.
func (s *LoadBalancerService) UpdateGlobalLoadBalancerRule(p *UpdateGlobalLoadBalancerRuleParams) (*UpdateGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("updateGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateGlobalLoadBalancerRuleResponse
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

type UpdateGlobalLoadBalancerRuleResponse struct {
	Account                     string                                                 `json:"account"`
	Description                 string                                                 `json:"description"`
	Domain                      string                                                 `json:"domain"`
	Domainid                    string                                                 `json:"domainid"`
	Gslbdomainname              string                                                 `json:"gslbdomainname"`
	Gslblbmethod                string                                                 `json:"gslblbmethod"`
	Gslbservicetype             string                                                 `json:"gslbservicetype"`
	Gslbstickysessionmethodname string                                                 `json:"gslbstickysessionmethodname"`
	Id                          string                                                 `json:"id"`
	JobID                       string                                                 `json:"jobid"`
	Jobstatus                   int                                                    `json:"jobstatus"`
	Loadbalancerrule            []UpdateGlobalLoadBalancerRuleResponseLoadbalancerrule `json:"loadbalancerrule"`
	Name                        string                                                 `json:"name"`
	Project                     string                                                 `json:"project"`
	Projectid                   string                                                 `json:"projectid"`
	Regionid                    int                                                    `json:"regionid"`
}

type UpdateGlobalLoadBalancerRuleResponseLoadbalancerrule struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type UpdateLBHealthCheckPolicyParams struct {
	P map[string]interface{}
}

func (P *UpdateLBHealthCheckPolicyParams) toURLValues() url.Values {
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

func (P *UpdateLBHealthCheckPolicyParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateLBHealthCheckPolicyParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateLBHealthCheckPolicyParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateLBHealthCheckPolicyParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateLBHealthCheckPolicyParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateLBHealthCheckPolicyParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLBHealthCheckPolicyParams(id string) *UpdateLBHealthCheckPolicyParams {
	P := &UpdateLBHealthCheckPolicyParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates load balancer health check policy
func (s *LoadBalancerService) UpdateLBHealthCheckPolicy(p *UpdateLBHealthCheckPolicyParams) (*UpdateLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("updateLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLBHealthCheckPolicyResponse
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

type UpdateLBHealthCheckPolicyResponse struct {
	Account           string                                               `json:"account"`
	Domain            string                                               `json:"domain"`
	Domainid          string                                               `json:"domainid"`
	Healthcheckpolicy []UpdateLBHealthCheckPolicyResponseHealthcheckpolicy `json:"healthcheckpolicy"`
	JobID             string                                               `json:"jobid"`
	Jobstatus         int                                                  `json:"jobstatus"`
	Lbruleid          string                                               `json:"lbruleid"`
	Zoneid            string                                               `json:"zoneid"`
}

type UpdateLBHealthCheckPolicyResponseHealthcheckpolicy struct {
	Description             string `json:"description"`
	Fordisplay              bool   `json:"fordisplay"`
	Healthcheckinterval     int    `json:"healthcheckinterval"`
	Healthcheckthresshold   int    `json:"healthcheckthresshold"`
	Id                      string `json:"id"`
	Pingpath                string `json:"pingpath"`
	Responsetime            int    `json:"responsetime"`
	State                   string `json:"state"`
	Unhealthcheckthresshold int    `json:"unhealthcheckthresshold"`
}

type UpdateLBStickinessPolicyParams struct {
	P map[string]interface{}
}

func (P *UpdateLBStickinessPolicyParams) toURLValues() url.Values {
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

func (P *UpdateLBStickinessPolicyParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateLBStickinessPolicyParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateLBStickinessPolicyParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateLBStickinessPolicyParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateLBStickinessPolicyParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateLBStickinessPolicyParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLBStickinessPolicyParams(id string) *UpdateLBStickinessPolicyParams {
	P := &UpdateLBStickinessPolicyParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates load balancer stickiness policy
func (s *LoadBalancerService) UpdateLBStickinessPolicy(p *UpdateLBStickinessPolicyParams) (*UpdateLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("updateLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLBStickinessPolicyResponse
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

type UpdateLBStickinessPolicyResponse struct {
	Account          string                                             `json:"account"`
	Description      string                                             `json:"description"`
	Domain           string                                             `json:"domain"`
	Domainid         string                                             `json:"domainid"`
	JobID            string                                             `json:"jobid"`
	Jobstatus        int                                                `json:"jobstatus"`
	Lbruleid         string                                             `json:"lbruleid"`
	Name             string                                             `json:"name"`
	State            string                                             `json:"state"`
	Stickinesspolicy []UpdateLBStickinessPolicyResponseStickinesspolicy `json:"stickinesspolicy"`
	Zoneid           string                                             `json:"zoneid"`
}

type UpdateLBStickinessPolicyResponseStickinesspolicy struct {
	Description string            `json:"description"`
	Fordisplay  bool              `json:"fordisplay"`
	Id          string            `json:"id"`
	Methodname  string            `json:"methodname"`
	Name        string            `json:"name"`
	Params      map[string]string `json:"params"`
	State       string            `json:"state"`
}

type UpdateLoadBalancerParams struct {
	P map[string]interface{}
}

func (P *UpdateLoadBalancerParams) toURLValues() url.Values {
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

func (P *UpdateLoadBalancerParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateLoadBalancerParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateLoadBalancerParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateLoadBalancerParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateLoadBalancerParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateLoadBalancerParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLoadBalancerParams(id string) *UpdateLoadBalancerParams {
	P := &UpdateLoadBalancerParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a load balancer
func (s *LoadBalancerService) UpdateLoadBalancer(p *UpdateLoadBalancerParams) (*UpdateLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("updateLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLoadBalancerResponse
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

type UpdateLoadBalancerResponse struct {
	Account                  string                                           `json:"account"`
	Algorithm                string                                           `json:"algorithm"`
	Description              string                                           `json:"description"`
	Domain                   string                                           `json:"domain"`
	Domainid                 string                                           `json:"domainid"`
	Fordisplay               bool                                             `json:"fordisplay"`
	Id                       string                                           `json:"id"`
	JobID                    string                                           `json:"jobid"`
	Jobstatus                int                                              `json:"jobstatus"`
	Loadbalancerinstance     []UpdateLoadBalancerResponseLoadbalancerinstance `json:"loadbalancerinstance"`
	Loadbalancerrule         []UpdateLoadBalancerResponseLoadbalancerrule     `json:"loadbalancerrule"`
	Name                     string                                           `json:"name"`
	Networkid                string                                           `json:"networkid"`
	Project                  string                                           `json:"project"`
	Projectid                string                                           `json:"projectid"`
	Sourceipaddress          string                                           `json:"sourceipaddress"`
	Sourceipaddressnetworkid string                                           `json:"sourceipaddressnetworkid"`
	Tags                     []Tags                                           `json:"tags"`
}

type UpdateLoadBalancerResponseLoadbalancerrule struct {
	Instanceport int    `json:"instanceport"`
	Sourceport   int    `json:"sourceport"`
	State        string `json:"state"`
}

type UpdateLoadBalancerResponseLoadbalancerinstance struct {
	Id        string `json:"id"`
	Ipaddress string `json:"ipaddress"`
	Name      string `json:"name"`
	State     string `json:"state"`
}

type UpdateLoadBalancerRuleParams struct {
	P map[string]interface{}
}

func (P *UpdateLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["algorithm"]; found {
		u.Set("algorithm", v.(string))
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
	if v, found := P.P["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	return u
}

func (P *UpdateLoadBalancerRuleParams) SetAlgorithm(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["algorithm"] = v
}

func (P *UpdateLoadBalancerRuleParams) GetAlgorithm() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["algorithm"].(string)
	return value, ok
}

func (P *UpdateLoadBalancerRuleParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateLoadBalancerRuleParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateLoadBalancerRuleParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *UpdateLoadBalancerRuleParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *UpdateLoadBalancerRuleParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateLoadBalancerRuleParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateLoadBalancerRuleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateLoadBalancerRuleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateLoadBalancerRuleParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateLoadBalancerRuleParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateLoadBalancerRuleParams) SetProtocol(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["protocol"] = v
}

func (P *UpdateLoadBalancerRuleParams) GetProtocol() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["protocol"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLoadBalancerRuleParams(id string) *UpdateLoadBalancerRuleParams {
	P := &UpdateLoadBalancerRuleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates load balancer
func (s *LoadBalancerService) UpdateLoadBalancerRule(p *UpdateLoadBalancerRuleParams) (*UpdateLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("updateLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLoadBalancerRuleResponse
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

type UpdateLoadBalancerRuleResponse struct {
	Account     string `json:"account"`
	Algorithm   string `json:"algorithm"`
	Cidrlist    string `json:"cidrlist"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Fordisplay  bool   `json:"fordisplay"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Networkid   string `json:"networkid"`
	Privateport string `json:"privateport"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	Protocol    string `json:"protocol"`
	Publicip    string `json:"publicip"`
	Publicipid  string `json:"publicipid"`
	Publicport  string `json:"publicport"`
	State       string `json:"state"`
	Tags        []Tags `json:"tags"`
	Zoneid      string `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type UploadSslCertParams struct {
	P map[string]interface{}
}

func (P *UploadSslCertParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["certchain"]; found {
		u.Set("certchain", v.(string))
	}
	if v, found := P.P["certificate"]; found {
		u.Set("certificate", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["enabledrevocationcheck"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabledrevocationcheck", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["privatekey"]; found {
		u.Set("privatekey", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *UploadSslCertParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *UploadSslCertParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *UploadSslCertParams) SetCertchain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["certchain"] = v
}

func (P *UploadSslCertParams) GetCertchain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["certchain"].(string)
	return value, ok
}

func (P *UploadSslCertParams) SetCertificate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["certificate"] = v
}

func (P *UploadSslCertParams) GetCertificate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["certificate"].(string)
	return value, ok
}

func (P *UploadSslCertParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UploadSslCertParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UploadSslCertParams) SetEnabledrevocationcheck(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enabledrevocationcheck"] = v
}

func (P *UploadSslCertParams) GetEnabledrevocationcheck() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enabledrevocationcheck"].(bool)
	return value, ok
}

func (P *UploadSslCertParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UploadSslCertParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UploadSslCertParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *UploadSslCertParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *UploadSslCertParams) SetPrivatekey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["privatekey"] = v
}

func (P *UploadSslCertParams) GetPrivatekey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["privatekey"].(string)
	return value, ok
}

func (P *UploadSslCertParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *UploadSslCertParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new UploadSslCertParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUploadSslCertParams(certificate string, name string, privatekey string) *UploadSslCertParams {
	P := &UploadSslCertParams{}
	P.P = make(map[string]interface{})
	P.P["certificate"] = certificate
	P.P["name"] = name
	P.P["privatekey"] = privatekey
	return P
}

// Upload a certificate to CloudStack
func (s *LoadBalancerService) UploadSslCert(p *UploadSslCertParams) (*UploadSslCertResponse, error) {
	resp, err := s.cs.newRequest("uploadSslCert", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadSslCertResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UploadSslCertResponse struct {
	Account              string   `json:"account"`
	Certchain            string   `json:"certchain"`
	Certificate          string   `json:"certificate"`
	Domain               string   `json:"domain"`
	Domainid             string   `json:"domainid"`
	Fingerprint          string   `json:"fingerprint"`
	Id                   string   `json:"id"`
	JobID                string   `json:"jobid"`
	Jobstatus            int      `json:"jobstatus"`
	Loadbalancerrulelist []string `json:"loadbalancerrulelist"`
	Name                 string   `json:"name"`
	Project              string   `json:"project"`
	Projectid            string   `json:"projectid"`
}
