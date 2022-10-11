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

type ConfigurationServiceIface interface {
	ListCapabilities(p *ListCapabilitiesParams) (*ListCapabilitiesResponse, error)
	NewListCapabilitiesParams() *ListCapabilitiesParams
	ListConfigurations(p *ListConfigurationsParams) (*ListConfigurationsResponse, error)
	NewListConfigurationsParams() *ListConfigurationsParams
	ListDeploymentPlanners(p *ListDeploymentPlannersParams) (*ListDeploymentPlannersResponse, error)
	NewListDeploymentPlannersParams() *ListDeploymentPlannersParams
	UpdateConfiguration(p *UpdateConfigurationParams) (*UpdateConfigurationResponse, error)
	NewUpdateConfigurationParams(name string) *UpdateConfigurationParams
}

type ListCapabilitiesParams struct {
	P map[string]interface{}
}

func (P *ListCapabilitiesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	return u
}

// You should always use this function to get a new ListCapabilitiesParams instance,
// as then you are sure you have configured all required params
func (s *ConfigurationService) NewListCapabilitiesParams() *ListCapabilitiesParams {
	P := &ListCapabilitiesParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists capabilities
func (s *ConfigurationService) ListCapabilities(p *ListCapabilitiesParams) (*ListCapabilitiesResponse, error) {
	resp, err := s.cs.newRequest("listCapabilities", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCapabilitiesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCapabilitiesResponse struct {
	Capabilities *Capability `json:"capability"`
}

type Capability struct {
	Allowusercreateprojects                      bool   `json:"allowusercreateprojects"`
	Allowuserexpungerecovervm                    bool   `json:"allowuserexpungerecovervm"`
	Allowuserexpungerecovervolume                bool   `json:"allowuserexpungerecovervolume"`
	Allowuserviewalldomainaccounts               bool   `json:"allowuserviewalldomainaccounts"`
	Allowuserviewdestroyedvm                     bool   `json:"allowuserviewdestroyedvm"`
	Apilimitinterval                             int    `json:"apilimitinterval"`
	Apilimitmax                                  int    `json:"apilimitmax"`
	Cloudstackversion                            string `json:"cloudstackversion"`
	Customdiskofferingmaxsize                    int64  `json:"customdiskofferingmaxsize"`
	Customdiskofferingminsize                    int64  `json:"customdiskofferingminsize"`
	Defaultuipagesize                            int64  `json:"defaultuipagesize"`
	Dynamicrolesenabled                          bool   `json:"dynamicrolesenabled"`
	JobID                                        string `json:"jobid"`
	Jobstatus                                    int    `json:"jobstatus"`
	Kubernetesclusterexperimentalfeaturesenabled bool   `json:"kubernetesclusterexperimentalfeaturesenabled"`
	Kubernetesserviceenabled                     bool   `json:"kubernetesserviceenabled"`
	Kvmsnapshotenabled                           bool   `json:"kvmsnapshotenabled"`
	Projectinviterequired                        bool   `json:"projectinviterequired"`
	Regionsecondaryenabled                       bool   `json:"regionsecondaryenabled"`
	Securitygroupsenabled                        bool   `json:"securitygroupsenabled"`
	SupportELB                                   string `json:"supportELB"`
	Userpublictemplateenabled                    bool   `json:"userpublictemplateenabled"`
}

type ListConfigurationsParams struct {
	P map[string]interface{}
}

func (P *ListConfigurationsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := P.P["category"]; found {
		u.Set("category", v.(string))
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["imagestoreuuid"]; found {
		u.Set("imagestoreuuid", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
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
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListConfigurationsParams) SetAccountid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountid"] = v
}

func (P *ListConfigurationsParams) GetAccountid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountid"].(string)
	return value, ok
}

func (P *ListConfigurationsParams) SetCategory(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["category"] = v
}

func (P *ListConfigurationsParams) GetCategory() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["category"].(string)
	return value, ok
}

func (P *ListConfigurationsParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *ListConfigurationsParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *ListConfigurationsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListConfigurationsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListConfigurationsParams) SetImagestoreuuid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["imagestoreuuid"] = v
}

func (P *ListConfigurationsParams) GetImagestoreuuid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["imagestoreuuid"].(string)
	return value, ok
}

func (P *ListConfigurationsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListConfigurationsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListConfigurationsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListConfigurationsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListConfigurationsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListConfigurationsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListConfigurationsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListConfigurationsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListConfigurationsParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *ListConfigurationsParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

func (P *ListConfigurationsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListConfigurationsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListConfigurationsParams instance,
// as then you are sure you have configured all required params
func (s *ConfigurationService) NewListConfigurationsParams() *ListConfigurationsParams {
	P := &ListConfigurationsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists all configurations.
func (s *ConfigurationService) ListConfigurations(p *ListConfigurationsParams) (*ListConfigurationsResponse, error) {
	resp, err := s.cs.newRequest("listConfigurations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListConfigurationsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListConfigurationsResponse struct {
	Count          int              `json:"count"`
	Configurations []*Configuration `json:"configuration"`
}

type Configuration struct {
	Category    string `json:"category"`
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Isdynamic   bool   `json:"isdynamic"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Scope       string `json:"scope"`
	Value       string `json:"value"`
}

type ListDeploymentPlannersParams struct {
	P map[string]interface{}
}

func (P *ListDeploymentPlannersParams) toURLValues() url.Values {
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
	return u
}

func (P *ListDeploymentPlannersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListDeploymentPlannersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListDeploymentPlannersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListDeploymentPlannersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListDeploymentPlannersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListDeploymentPlannersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListDeploymentPlannersParams instance,
// as then you are sure you have configured all required params
func (s *ConfigurationService) NewListDeploymentPlannersParams() *ListDeploymentPlannersParams {
	P := &ListDeploymentPlannersParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists all DeploymentPlanners available.
func (s *ConfigurationService) ListDeploymentPlanners(p *ListDeploymentPlannersParams) (*ListDeploymentPlannersResponse, error) {
	resp, err := s.cs.newRequest("listDeploymentPlanners", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDeploymentPlannersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDeploymentPlannersResponse struct {
	Count              int                  `json:"count"`
	DeploymentPlanners []*DeploymentPlanner `json:"deploymentplanner"`
}

type DeploymentPlanner struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type UpdateConfigurationParams struct {
	P map[string]interface{}
}

func (P *UpdateConfigurationParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["imagestoreuuid"]; found {
		u.Set("imagestoreuuid", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := P.P["value"]; found {
		u.Set("value", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *UpdateConfigurationParams) SetAccountid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountid"] = v
}

func (P *UpdateConfigurationParams) GetAccountid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountid"].(string)
	return value, ok
}

func (P *UpdateConfigurationParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *UpdateConfigurationParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *UpdateConfigurationParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UpdateConfigurationParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UpdateConfigurationParams) SetImagestoreuuid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["imagestoreuuid"] = v
}

func (P *UpdateConfigurationParams) GetImagestoreuuid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["imagestoreuuid"].(string)
	return value, ok
}

func (P *UpdateConfigurationParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateConfigurationParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateConfigurationParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *UpdateConfigurationParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

func (P *UpdateConfigurationParams) SetValue(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["value"] = v
}

func (P *UpdateConfigurationParams) GetValue() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["value"].(string)
	return value, ok
}

func (P *UpdateConfigurationParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *UpdateConfigurationParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateConfigurationParams instance,
// as then you are sure you have configured all required params
func (s *ConfigurationService) NewUpdateConfigurationParams(name string) *UpdateConfigurationParams {
	P := &UpdateConfigurationParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	return P
}

// Updates a configuration.
func (s *ConfigurationService) UpdateConfiguration(p *UpdateConfigurationParams) (*UpdateConfigurationResponse, error) {
	resp, err := s.cs.newRequest("updateConfiguration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r UpdateConfigurationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateConfigurationResponse struct {
	Category    string `json:"category"`
	Description string `json:"description"`
	Id          int64  `json:"id"`
	Isdynamic   bool   `json:"isdynamic"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Scope       string `json:"scope"`
	Value       string `json:"value"`
}
