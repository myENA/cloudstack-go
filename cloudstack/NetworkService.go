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

type NetworkServiceIface interface {
	AddNetworkServiceProvider(P *AddNetworkServiceProviderParams) (*AddNetworkServiceProviderResponse, error)
	NewAddNetworkServiceProviderParams(name string, physicalnetworkid string) *AddNetworkServiceProviderParams
	AddOpenDaylightController(P *AddOpenDaylightControllerParams) (*AddOpenDaylightControllerResponse, error)
	NewAddOpenDaylightControllerParams(password string, physicalnetworkid string, url string, username string) *AddOpenDaylightControllerParams
	CreateNetwork(P *CreateNetworkParams) (*CreateNetworkResponse, error)
	NewCreateNetworkParams(displaytext string, name string, networkofferingid string, zoneid string) *CreateNetworkParams
	CreatePhysicalNetwork(P *CreatePhysicalNetworkParams) (*CreatePhysicalNetworkResponse, error)
	NewCreatePhysicalNetworkParams(name string, zoneid string) *CreatePhysicalNetworkParams
	CreateServiceInstance(P *CreateServiceInstanceParams) (*CreateServiceInstanceResponse, error)
	NewCreateServiceInstanceParams(leftnetworkid string, name string, rightnetworkid string, serviceofferingid string, templateid string, zoneid string) *CreateServiceInstanceParams
	CreateStorageNetworkIpRange(P *CreateStorageNetworkIpRangeParams) (*CreateStorageNetworkIpRangeResponse, error)
	NewCreateStorageNetworkIpRangeParams(gateway string, netmask string, podid string, startip string) *CreateStorageNetworkIpRangeParams
	DedicatePublicIpRange(P *DedicatePublicIpRangeParams) (*DedicatePublicIpRangeResponse, error)
	NewDedicatePublicIpRangeParams(domainid string, id string) *DedicatePublicIpRangeParams
	DeleteNetwork(P *DeleteNetworkParams) (*DeleteNetworkResponse, error)
	NewDeleteNetworkParams(id string) *DeleteNetworkParams
	DeleteNetworkServiceProvider(P *DeleteNetworkServiceProviderParams) (*DeleteNetworkServiceProviderResponse, error)
	NewDeleteNetworkServiceProviderParams(id string) *DeleteNetworkServiceProviderParams
	DeleteOpenDaylightController(P *DeleteOpenDaylightControllerParams) (*DeleteOpenDaylightControllerResponse, error)
	NewDeleteOpenDaylightControllerParams(id string) *DeleteOpenDaylightControllerParams
	DeletePhysicalNetwork(P *DeletePhysicalNetworkParams) (*DeletePhysicalNetworkResponse, error)
	NewDeletePhysicalNetworkParams(id string) *DeletePhysicalNetworkParams
	DeleteStorageNetworkIpRange(P *DeleteStorageNetworkIpRangeParams) (*DeleteStorageNetworkIpRangeResponse, error)
	NewDeleteStorageNetworkIpRangeParams(id string) *DeleteStorageNetworkIpRangeParams
	ListNetscalerLoadBalancerNetworks(P *ListNetscalerLoadBalancerNetworksParams) (*ListNetscalerLoadBalancerNetworksResponse, error)
	NewListNetscalerLoadBalancerNetworksParams(lbdeviceid string) *ListNetscalerLoadBalancerNetworksParams
	GetNetscalerLoadBalancerNetworkID(keyword string, lbdeviceid string, opts ...OptionFunc) (string, int, error)
	ListNetworkIsolationMethods(P *ListNetworkIsolationMethodsParams) (*ListNetworkIsolationMethodsResponse, error)
	NewListNetworkIsolationMethodsParams() *ListNetworkIsolationMethodsParams
	ListNetworkServiceProviders(P *ListNetworkServiceProvidersParams) (*ListNetworkServiceProvidersResponse, error)
	NewListNetworkServiceProvidersParams() *ListNetworkServiceProvidersParams
	GetNetworkServiceProviderID(name string, opts ...OptionFunc) (string, int, error)
	ListNetworks(P *ListNetworksParams) (*ListNetworksResponse, error)
	NewListNetworksParams() *ListNetworksParams
	GetNetworkID(keyword string, opts ...OptionFunc) (string, int, error)
	GetNetworkByName(name string, opts ...OptionFunc) (*Network, int, error)
	GetNetworkByID(id string, opts ...OptionFunc) (*Network, int, error)
	ListNiciraNvpDeviceNetworks(P *ListNiciraNvpDeviceNetworksParams) (*ListNiciraNvpDeviceNetworksResponse, error)
	NewListNiciraNvpDeviceNetworksParams(nvpdeviceid string) *ListNiciraNvpDeviceNetworksParams
	GetNiciraNvpDeviceNetworkID(keyword string, nvpdeviceid string, opts ...OptionFunc) (string, int, error)
	ListOpenDaylightControllers(P *ListOpenDaylightControllersParams) (*ListOpenDaylightControllersResponse, error)
	NewListOpenDaylightControllersParams() *ListOpenDaylightControllersParams
	GetOpenDaylightControllerByID(id string, opts ...OptionFunc) (*OpenDaylightController, int, error)
	ListPaloAltoFirewallNetworks(P *ListPaloAltoFirewallNetworksParams) (*ListPaloAltoFirewallNetworksResponse, error)
	NewListPaloAltoFirewallNetworksParams(lbdeviceid string) *ListPaloAltoFirewallNetworksParams
	GetPaloAltoFirewallNetworkID(keyword string, lbdeviceid string, opts ...OptionFunc) (string, int, error)
	ListPhysicalNetworks(P *ListPhysicalNetworksParams) (*ListPhysicalNetworksResponse, error)
	NewListPhysicalNetworksParams() *ListPhysicalNetworksParams
	GetPhysicalNetworkID(name string, opts ...OptionFunc) (string, int, error)
	GetPhysicalNetworkByName(name string, opts ...OptionFunc) (*PhysicalNetwork, int, error)
	GetPhysicalNetworkByID(id string, opts ...OptionFunc) (*PhysicalNetwork, int, error)
	ListStorageNetworkIpRange(P *ListStorageNetworkIpRangeParams) (*ListStorageNetworkIpRangeResponse, error)
	NewListStorageNetworkIpRangeParams() *ListStorageNetworkIpRangeParams
	GetStorageNetworkIpRangeByID(id string, opts ...OptionFunc) (*StorageNetworkIpRange, int, error)
	ListSupportedNetworkServices(P *ListSupportedNetworkServicesParams) (*ListSupportedNetworkServicesResponse, error)
	NewListSupportedNetworkServicesParams() *ListSupportedNetworkServicesParams
	ReleasePublicIpRange(P *ReleasePublicIpRangeParams) (*ReleasePublicIpRangeResponse, error)
	NewReleasePublicIpRangeParams(id string) *ReleasePublicIpRangeParams
	RestartNetwork(P *RestartNetworkParams) (*RestartNetworkResponse, error)
	NewRestartNetworkParams(id string) *RestartNetworkParams
	UpdateNetwork(P *UpdateNetworkParams) (*UpdateNetworkResponse, error)
	NewUpdateNetworkParams(id string) *UpdateNetworkParams
	UpdateNetworkServiceProvider(P *UpdateNetworkServiceProviderParams) (*UpdateNetworkServiceProviderResponse, error)
	NewUpdateNetworkServiceProviderParams(id string) *UpdateNetworkServiceProviderParams
	UpdatePhysicalNetwork(P *UpdatePhysicalNetworkParams) (*UpdatePhysicalNetworkResponse, error)
	NewUpdatePhysicalNetworkParams(id string) *UpdatePhysicalNetworkParams
	UpdateStorageNetworkIpRange(P *UpdateStorageNetworkIpRangeParams) (*UpdateStorageNetworkIpRangeResponse, error)
	NewUpdateStorageNetworkIpRangeParams(id string) *UpdateStorageNetworkIpRangeParams
}

type AddNetworkServiceProviderParams struct {
	P map[string]interface{}
}

func (P *AddNetworkServiceProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["destinationphysicalnetworkid"]; found {
		u.Set("destinationphysicalnetworkid", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["servicelist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("servicelist", vv)
	}
	return u
}

func (P *AddNetworkServiceProviderParams) SetDestinationphysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["destinationphysicalnetworkid"] = v
}

func (P *AddNetworkServiceProviderParams) GetDestinationphysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["destinationphysicalnetworkid"].(string)
	return value, ok
}

func (P *AddNetworkServiceProviderParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *AddNetworkServiceProviderParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *AddNetworkServiceProviderParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddNetworkServiceProviderParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddNetworkServiceProviderParams) SetServicelist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["servicelist"] = v
}

func (P *AddNetworkServiceProviderParams) GetServicelist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["servicelist"].([]string)
	return value, ok
}

// You should always use this function to get a new AddNetworkServiceProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewAddNetworkServiceProviderParams(name string, physicalnetworkid string) *AddNetworkServiceProviderParams {
	P := &AddNetworkServiceProviderParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	P.P["physicalnetworkid"] = physicalnetworkid
	return P
}

// Adds a network serviceProvider to a physical network
func (s *NetworkService) AddNetworkServiceProvider(p *AddNetworkServiceProviderParams) (*AddNetworkServiceProviderResponse, error) {
	resp, err := s.cs.newRequest("addNetworkServiceProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNetworkServiceProviderResponse
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

type AddNetworkServiceProviderResponse struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	JobID                        string   `json:"jobid"`
	Jobstatus                    int      `json:"jobstatus"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type AddOpenDaylightControllerParams struct {
	P map[string]interface{}
}

func (P *AddOpenDaylightControllerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
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

func (P *AddOpenDaylightControllerParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddOpenDaylightControllerParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddOpenDaylightControllerParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddOpenDaylightControllerParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddOpenDaylightControllerParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddOpenDaylightControllerParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddOpenDaylightControllerParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddOpenDaylightControllerParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddOpenDaylightControllerParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewAddOpenDaylightControllerParams(password string, physicalnetworkid string, url string, username string) *AddOpenDaylightControllerParams {
	P := &AddOpenDaylightControllerParams{}
	P.P = make(map[string]interface{})
	P.P["password"] = password
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["url"] = url
	P.P["username"] = username
	return P
}

// Adds an OpenDyalight controler
func (s *NetworkService) AddOpenDaylightController(p *AddOpenDaylightControllerParams) (*AddOpenDaylightControllerResponse, error) {
	resp, err := s.cs.newRequest("addOpenDaylightController", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddOpenDaylightControllerResponse
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

type AddOpenDaylightControllerResponse struct {
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Name              string `json:"name"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Url               string `json:"url"`
	Username          string `json:"username"`
}

type CreateNetworkParams struct {
	P map[string]interface{}
}

func (P *CreateNetworkParams) toURLValues() url.Values {
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
	if v, found := P.P["acltype"]; found {
		u.Set("acltype", v.(string))
	}
	if v, found := P.P["bypassvlanoverlapcheck"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bypassvlanoverlapcheck", vv)
	}
	if v, found := P.P["displaynetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displaynetwork", vv)
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := P.P["endipv6"]; found {
		u.Set("endipv6", v.(string))
	}
	if v, found := P.P["externalid"]; found {
		u.Set("externalid", v.(string))
	}
	if v, found := P.P["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := P.P["hideipaddressusage"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("hideipaddressusage", vv)
	}
	if v, found := P.P["ip6cidr"]; found {
		u.Set("ip6cidr", v.(string))
	}
	if v, found := P.P["ip6gateway"]; found {
		u.Set("ip6gateway", v.(string))
	}
	if v, found := P.P["isolatedpvlan"]; found {
		u.Set("isolatedpvlan", v.(string))
	}
	if v, found := P.P["isolatedpvlantype"]; found {
		u.Set("isolatedpvlantype", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := P.P["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := P.P["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["routerip"]; found {
		u.Set("routerip", v.(string))
	}
	if v, found := P.P["routeripv6"]; found {
		u.Set("routeripv6", v.(string))
	}
	if v, found := P.P["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := P.P["startipv6"]; found {
		u.Set("startipv6", v.(string))
	}
	if v, found := P.P["subdomainaccess"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("subdomainaccess", vv)
	}
	if v, found := P.P["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *CreateNetworkParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateNetworkParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetAclid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["aclid"] = v
}

func (P *CreateNetworkParams) GetAclid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["aclid"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetAcltype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["acltype"] = v
}

func (P *CreateNetworkParams) GetAcltype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["acltype"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetBypassvlanoverlapcheck(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bypassvlanoverlapcheck"] = v
}

func (P *CreateNetworkParams) GetBypassvlanoverlapcheck() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bypassvlanoverlapcheck"].(bool)
	return value, ok
}

func (P *CreateNetworkParams) SetDisplaynetwork(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaynetwork"] = v
}

func (P *CreateNetworkParams) GetDisplaynetwork() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaynetwork"].(bool)
	return value, ok
}

func (P *CreateNetworkParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *CreateNetworkParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateNetworkParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetEndip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endip"] = v
}

func (P *CreateNetworkParams) GetEndip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endip"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetEndipv6(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endipv6"] = v
}

func (P *CreateNetworkParams) GetEndipv6() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endipv6"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetExternalid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["externalid"] = v
}

func (P *CreateNetworkParams) GetExternalid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["externalid"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetGateway(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gateway"] = v
}

func (P *CreateNetworkParams) GetGateway() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gateway"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetHideipaddressusage(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hideipaddressusage"] = v
}

func (P *CreateNetworkParams) GetHideipaddressusage() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hideipaddressusage"].(bool)
	return value, ok
}

func (P *CreateNetworkParams) SetIp6cidr(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ip6cidr"] = v
}

func (P *CreateNetworkParams) GetIp6cidr() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ip6cidr"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetIp6gateway(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ip6gateway"] = v
}

func (P *CreateNetworkParams) GetIp6gateway() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ip6gateway"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetIsolatedpvlan(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isolatedpvlan"] = v
}

func (P *CreateNetworkParams) GetIsolatedpvlan() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isolatedpvlan"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetIsolatedpvlantype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isolatedpvlantype"] = v
}

func (P *CreateNetworkParams) GetIsolatedpvlantype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isolatedpvlantype"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateNetworkParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetNetmask(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["netmask"] = v
}

func (P *CreateNetworkParams) GetNetmask() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["netmask"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetNetworkdomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdomain"] = v
}

func (P *CreateNetworkParams) GetNetworkdomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdomain"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetNetworkofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkofferingid"] = v
}

func (P *CreateNetworkParams) GetNetworkofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkofferingid"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *CreateNetworkParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateNetworkParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetRouterip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["routerip"] = v
}

func (P *CreateNetworkParams) GetRouterip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["routerip"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetRouteripv6(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["routeripv6"] = v
}

func (P *CreateNetworkParams) GetRouteripv6() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["routeripv6"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetStartip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startip"] = v
}

func (P *CreateNetworkParams) GetStartip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startip"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetStartipv6(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startipv6"] = v
}

func (P *CreateNetworkParams) GetStartipv6() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startipv6"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetSubdomainaccess(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["subdomainaccess"] = v
}

func (P *CreateNetworkParams) GetSubdomainaccess() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["subdomainaccess"].(bool)
	return value, ok
}

func (P *CreateNetworkParams) SetVlan(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vlan"] = v
}

func (P *CreateNetworkParams) GetVlan() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vlan"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *CreateNetworkParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

func (P *CreateNetworkParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateNetworkParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateNetworkParams(displaytext string, name string, networkofferingid string, zoneid string) *CreateNetworkParams {
	P := &CreateNetworkParams{}
	P.P = make(map[string]interface{})
	P.P["displaytext"] = displaytext
	P.P["name"] = name
	P.P["networkofferingid"] = networkofferingid
	P.P["zoneid"] = zoneid
	return P
}

// Creates a network
func (s *NetworkService) CreateNetwork(p *CreateNetworkParams) (*CreateNetworkResponse, error) {
	resp, err := s.cs.newRequest("createNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateNetworkResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateNetworkResponse struct {
	Account                     string                         `json:"account"`
	Aclid                       string                         `json:"aclid"`
	Aclname                     string                         `json:"aclname"`
	Acltype                     string                         `json:"acltype"`
	Broadcastdomaintype         string                         `json:"broadcastdomaintype"`
	Broadcasturi                string                         `json:"broadcasturi"`
	Canusefordeploy             bool                           `json:"canusefordeploy"`
	Cidr                        string                         `json:"cidr"`
	Created                     string                         `json:"created"`
	Details                     map[string]string              `json:"details"`
	Displaynetwork              bool                           `json:"displaynetwork"`
	Displaytext                 string                         `json:"displaytext"`
	Dns1                        string                         `json:"dns1"`
	Dns2                        string                         `json:"dns2"`
	Domain                      string                         `json:"domain"`
	Domainid                    string                         `json:"domainid"`
	Externalid                  string                         `json:"externalid"`
	Gateway                     string                         `json:"gateway"`
	Hasannotations              bool                           `json:"hasannotations"`
	Icon                        string                         `json:"icon"`
	Id                          string                         `json:"id"`
	Ip6cidr                     string                         `json:"ip6cidr"`
	Ip6gateway                  string                         `json:"ip6gateway"`
	Isdefault                   bool                           `json:"isdefault"`
	Ispersistent                bool                           `json:"ispersistent"`
	Issystem                    bool                           `json:"issystem"`
	JobID                       string                         `json:"jobid"`
	Jobstatus                   int                            `json:"jobstatus"`
	Name                        string                         `json:"name"`
	Netmask                     string                         `json:"netmask"`
	Networkcidr                 string                         `json:"networkcidr"`
	Networkdomain               string                         `json:"networkdomain"`
	Networkofferingavailability string                         `json:"networkofferingavailability"`
	Networkofferingconservemode bool                           `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                         `json:"networkofferingdisplaytext"`
	Networkofferingid           string                         `json:"networkofferingid"`
	Networkofferingname         string                         `json:"networkofferingname"`
	Physicalnetworkid           string                         `json:"physicalnetworkid"`
	Project                     string                         `json:"project"`
	Projectid                   string                         `json:"projectid"`
	Receivedbytes               int64                          `json:"receivedbytes"`
	Redundantrouter             bool                           `json:"redundantrouter"`
	Related                     string                         `json:"related"`
	Reservediprange             string                         `json:"reservediprange"`
	Restartrequired             bool                           `json:"restartrequired"`
	Sentbytes                   int64                          `json:"sentbytes"`
	Service                     []CreateNetworkResponseService `json:"service"`
	Specifyipranges             bool                           `json:"specifyipranges"`
	State                       string                         `json:"state"`
	Strechedl2subnet            bool                           `json:"strechedl2subnet"`
	Subdomainaccess             bool                           `json:"subdomainaccess"`
	Tags                        []Tags                         `json:"tags"`
	Traffictype                 string                         `json:"traffictype"`
	Type                        string                         `json:"type"`
	Vlan                        string                         `json:"vlan"`
	Vpcid                       string                         `json:"vpcid"`
	Vpcname                     string                         `json:"vpcname"`
	Zoneid                      string                         `json:"zoneid"`
	Zonename                    string                         `json:"zonename"`
	Zonesnetworkspans           []interface{}                  `json:"zonesnetworkspans"`
}

type CreateNetworkResponseService struct {
	Capability []CreateNetworkResponseServiceCapability `json:"capability"`
	Name       string                                   `json:"name"`
	Provider   []CreateNetworkResponseServiceProvider   `json:"provider"`
}

type CreateNetworkResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type CreateNetworkResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type CreatePhysicalNetworkParams struct {
	P map[string]interface{}
}

func (P *CreatePhysicalNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["broadcastdomainrange"]; found {
		u.Set("broadcastdomainrange", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["isolationmethods"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("isolationmethods", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networkspeed"]; found {
		u.Set("networkspeed", v.(string))
	}
	if v, found := P.P["tags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("tags", vv)
	}
	if v, found := P.P["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *CreatePhysicalNetworkParams) SetBroadcastdomainrange(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["broadcastdomainrange"] = v
}

func (P *CreatePhysicalNetworkParams) GetBroadcastdomainrange() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["broadcastdomainrange"].(string)
	return value, ok
}

func (P *CreatePhysicalNetworkParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreatePhysicalNetworkParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreatePhysicalNetworkParams) SetIsolationmethods(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isolationmethods"] = v
}

func (P *CreatePhysicalNetworkParams) GetIsolationmethods() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isolationmethods"].([]string)
	return value, ok
}

func (P *CreatePhysicalNetworkParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreatePhysicalNetworkParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreatePhysicalNetworkParams) SetNetworkspeed(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkspeed"] = v
}

func (P *CreatePhysicalNetworkParams) GetNetworkspeed() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkspeed"].(string)
	return value, ok
}

func (P *CreatePhysicalNetworkParams) SetTags(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *CreatePhysicalNetworkParams) GetTags() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].([]string)
	return value, ok
}

func (P *CreatePhysicalNetworkParams) SetVlan(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vlan"] = v
}

func (P *CreatePhysicalNetworkParams) GetVlan() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vlan"].(string)
	return value, ok
}

func (P *CreatePhysicalNetworkParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreatePhysicalNetworkParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreatePhysicalNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreatePhysicalNetworkParams(name string, zoneid string) *CreatePhysicalNetworkParams {
	P := &CreatePhysicalNetworkParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	P.P["zoneid"] = zoneid
	return P
}

// Creates a physical network
func (s *NetworkService) CreatePhysicalNetwork(p *CreatePhysicalNetworkParams) (*CreatePhysicalNetworkResponse, error) {
	resp, err := s.cs.newRequest("createPhysicalNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePhysicalNetworkResponse
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

type CreatePhysicalNetworkResponse struct {
	Broadcastdomainrange string `json:"broadcastdomainrange"`
	Domainid             string `json:"domainid"`
	Id                   string `json:"id"`
	Isolationmethods     string `json:"isolationmethods"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Name                 string `json:"name"`
	Networkspeed         string `json:"networkspeed"`
	State                string `json:"state"`
	Tags                 string `json:"tags"`
	Vlan                 string `json:"vlan"`
	Zoneid               string `json:"zoneid"`
	Zonename             string `json:"zonename"`
}

type CreateServiceInstanceParams struct {
	P map[string]interface{}
}

func (P *CreateServiceInstanceParams) toURLValues() url.Values {
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
	if v, found := P.P["leftnetworkid"]; found {
		u.Set("leftnetworkid", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["rightnetworkid"]; found {
		u.Set("rightnetworkid", v.(string))
	}
	if v, found := P.P["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := P.P["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *CreateServiceInstanceParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateServiceInstanceParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateServiceInstanceParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateServiceInstanceParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateServiceInstanceParams) SetLeftnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["leftnetworkid"] = v
}

func (P *CreateServiceInstanceParams) GetLeftnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["leftnetworkid"].(string)
	return value, ok
}

func (P *CreateServiceInstanceParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateServiceInstanceParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateServiceInstanceParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateServiceInstanceParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *CreateServiceInstanceParams) SetRightnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["rightnetworkid"] = v
}

func (P *CreateServiceInstanceParams) GetRightnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["rightnetworkid"].(string)
	return value, ok
}

func (P *CreateServiceInstanceParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *CreateServiceInstanceParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

func (P *CreateServiceInstanceParams) SetTemplateid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templateid"] = v
}

func (P *CreateServiceInstanceParams) GetTemplateid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templateid"].(string)
	return value, ok
}

func (P *CreateServiceInstanceParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateServiceInstanceParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateServiceInstanceParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateServiceInstanceParams(leftnetworkid string, name string, rightnetworkid string, serviceofferingid string, templateid string, zoneid string) *CreateServiceInstanceParams {
	P := &CreateServiceInstanceParams{}
	P.P = make(map[string]interface{})
	P.P["leftnetworkid"] = leftnetworkid
	P.P["name"] = name
	P.P["rightnetworkid"] = rightnetworkid
	P.P["serviceofferingid"] = serviceofferingid
	P.P["templateid"] = templateid
	P.P["zoneid"] = zoneid
	return P
}

// Creates a system virtual-machine that implements network services
func (s *NetworkService) CreateServiceInstance(p *CreateServiceInstanceParams) (*CreateServiceInstanceResponse, error) {
	resp, err := s.cs.newRequest("createServiceInstance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateServiceInstanceResponse
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

type CreateServiceInstanceResponse struct {
	Account     string `json:"account"`
	Displayname string `json:"displayname"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
}

type CreateStorageNetworkIpRangeParams struct {
	P map[string]interface{}
}

func (P *CreateStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := P.P["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := P.P["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := P.P["vlan"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("vlan", vv)
	}
	return u
}

func (P *CreateStorageNetworkIpRangeParams) SetEndip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endip"] = v
}

func (P *CreateStorageNetworkIpRangeParams) GetEndip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endip"].(string)
	return value, ok
}

func (P *CreateStorageNetworkIpRangeParams) SetGateway(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gateway"] = v
}

func (P *CreateStorageNetworkIpRangeParams) GetGateway() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gateway"].(string)
	return value, ok
}

func (P *CreateStorageNetworkIpRangeParams) SetNetmask(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["netmask"] = v
}

func (P *CreateStorageNetworkIpRangeParams) GetNetmask() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["netmask"].(string)
	return value, ok
}

func (P *CreateStorageNetworkIpRangeParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *CreateStorageNetworkIpRangeParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *CreateStorageNetworkIpRangeParams) SetStartip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startip"] = v
}

func (P *CreateStorageNetworkIpRangeParams) GetStartip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startip"].(string)
	return value, ok
}

func (P *CreateStorageNetworkIpRangeParams) SetVlan(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vlan"] = v
}

func (P *CreateStorageNetworkIpRangeParams) GetVlan() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vlan"].(int)
	return value, ok
}

// You should always use this function to get a new CreateStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewCreateStorageNetworkIpRangeParams(gateway string, netmask string, podid string, startip string) *CreateStorageNetworkIpRangeParams {
	P := &CreateStorageNetworkIpRangeParams{}
	P.P = make(map[string]interface{})
	P.P["gateway"] = gateway
	P.P["netmask"] = netmask
	P.P["podid"] = podid
	P.P["startip"] = startip
	return P
}

// Creates a Storage network IP range.
func (s *NetworkService) CreateStorageNetworkIpRange(p *CreateStorageNetworkIpRangeParams) (*CreateStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("createStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateStorageNetworkIpRangeResponse
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

type CreateStorageNetworkIpRangeResponse struct {
	Endip     string `json:"endip"`
	Gateway   string `json:"gateway"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Netmask   string `json:"netmask"`
	Networkid string `json:"networkid"`
	Podid     string `json:"podid"`
	Startip   string `json:"startip"`
	Vlan      int    `json:"vlan"`
	Zoneid    string `json:"zoneid"`
}

type DedicatePublicIpRangeParams struct {
	P map[string]interface{}
}

func (P *DedicatePublicIpRangeParams) toURLValues() url.Values {
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *DedicatePublicIpRangeParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DedicatePublicIpRangeParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DedicatePublicIpRangeParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DedicatePublicIpRangeParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *DedicatePublicIpRangeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DedicatePublicIpRangeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *DedicatePublicIpRangeParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *DedicatePublicIpRangeParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicatePublicIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDedicatePublicIpRangeParams(domainid string, id string) *DedicatePublicIpRangeParams {
	P := &DedicatePublicIpRangeParams{}
	P.P = make(map[string]interface{})
	P.P["domainid"] = domainid
	P.P["id"] = id
	return P
}

// Dedicates a Public IP range to an account
func (s *NetworkService) DedicatePublicIpRange(p *DedicatePublicIpRangeParams) (*DedicatePublicIpRangeResponse, error) {
	resp, err := s.cs.newRequest("dedicatePublicIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicatePublicIpRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DedicatePublicIpRangeResponse struct {
	Account           string `json:"account"`
	Description       string `json:"description"`
	Domain            string `json:"domain"`
	Domainid          string `json:"domainid"`
	Endip             string `json:"endip"`
	Endipv6           string `json:"endipv6"`
	Forsystemvms      bool   `json:"forsystemvms"`
	Forvirtualnetwork bool   `json:"forvirtualnetwork"`
	Gateway           string `json:"gateway"`
	Id                string `json:"id"`
	Ip6cidr           string `json:"ip6cidr"`
	Ip6gateway        string `json:"ip6gateway"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Netmask           string `json:"netmask"`
	Networkid         string `json:"networkid"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Project           string `json:"project"`
	Projectid         string `json:"projectid"`
	Startip           string `json:"startip"`
	Startipv6         string `json:"startipv6"`
	Vlan              string `json:"vlan"`
	Zoneid            string `json:"zoneid"`
}

type DeleteNetworkParams struct {
	P map[string]interface{}
}

func (P *DeleteNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteNetworkParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *DeleteNetworkParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *DeleteNetworkParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteNetworkParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteNetworkParams(id string) *DeleteNetworkParams {
	P := &DeleteNetworkParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a network
func (s *NetworkService) DeleteNetwork(p *DeleteNetworkParams) (*DeleteNetworkResponse, error) {
	resp, err := s.cs.newRequest("deleteNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkResponse
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

type DeleteNetworkResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteNetworkServiceProviderParams struct {
	P map[string]interface{}
}

func (P *DeleteNetworkServiceProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteNetworkServiceProviderParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteNetworkServiceProviderParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkServiceProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteNetworkServiceProviderParams(id string) *DeleteNetworkServiceProviderParams {
	P := &DeleteNetworkServiceProviderParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a Network Service Provider.
func (s *NetworkService) DeleteNetworkServiceProvider(p *DeleteNetworkServiceProviderParams) (*DeleteNetworkServiceProviderResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkServiceProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkServiceProviderResponse
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

type DeleteNetworkServiceProviderResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteOpenDaylightControllerParams struct {
	P map[string]interface{}
}

func (P *DeleteOpenDaylightControllerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteOpenDaylightControllerParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteOpenDaylightControllerParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteOpenDaylightControllerParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteOpenDaylightControllerParams(id string) *DeleteOpenDaylightControllerParams {
	P := &DeleteOpenDaylightControllerParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Removes an OpenDyalight controler
func (s *NetworkService) DeleteOpenDaylightController(p *DeleteOpenDaylightControllerParams) (*DeleteOpenDaylightControllerResponse, error) {
	resp, err := s.cs.newRequest("deleteOpenDaylightController", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteOpenDaylightControllerResponse
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

type DeleteOpenDaylightControllerResponse struct {
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Name              string `json:"name"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Url               string `json:"url"`
	Username          string `json:"username"`
}

type DeletePhysicalNetworkParams struct {
	P map[string]interface{}
}

func (P *DeletePhysicalNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeletePhysicalNetworkParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeletePhysicalNetworkParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeletePhysicalNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeletePhysicalNetworkParams(id string) *DeletePhysicalNetworkParams {
	P := &DeletePhysicalNetworkParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a Physical Network.
func (s *NetworkService) DeletePhysicalNetwork(p *DeletePhysicalNetworkParams) (*DeletePhysicalNetworkResponse, error) {
	resp, err := s.cs.newRequest("deletePhysicalNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePhysicalNetworkResponse
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

type DeletePhysicalNetworkResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteStorageNetworkIpRangeParams struct {
	P map[string]interface{}
}

func (P *DeleteStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteStorageNetworkIpRangeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteStorageNetworkIpRangeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewDeleteStorageNetworkIpRangeParams(id string) *DeleteStorageNetworkIpRangeParams {
	P := &DeleteStorageNetworkIpRangeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a storage network IP Range.
func (s *NetworkService) DeleteStorageNetworkIpRange(p *DeleteStorageNetworkIpRangeParams) (*DeleteStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("deleteStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteStorageNetworkIpRangeResponse
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

type DeleteStorageNetworkIpRangeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListNetscalerLoadBalancerNetworksParams struct {
	P map[string]interface{}
}

func (P *ListNetscalerLoadBalancerNetworksParams) toURLValues() url.Values {
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
	return u
}

func (P *ListNetscalerLoadBalancerNetworksParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNetscalerLoadBalancerNetworksParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNetscalerLoadBalancerNetworksParams) SetLbdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbdeviceid"] = v
}

func (P *ListNetscalerLoadBalancerNetworksParams) GetLbdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbdeviceid"].(string)
	return value, ok
}

func (P *ListNetscalerLoadBalancerNetworksParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNetscalerLoadBalancerNetworksParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNetscalerLoadBalancerNetworksParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNetscalerLoadBalancerNetworksParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListNetscalerLoadBalancerNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetscalerLoadBalancerNetworksParams(lbdeviceid string) *ListNetscalerLoadBalancerNetworksParams {
	P := &ListNetscalerLoadBalancerNetworksParams{}
	P.P = make(map[string]interface{})
	P.P["lbdeviceid"] = lbdeviceid
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetscalerLoadBalancerNetworkID(keyword string, lbdeviceid string, opts ...OptionFunc) (string, int, error) {
	P := &ListNetscalerLoadBalancerNetworksParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword
	P.P["lbdeviceid"] = lbdeviceid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetscalerLoadBalancerNetworks(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.NetscalerLoadBalancerNetworks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NetscalerLoadBalancerNetworks {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// lists network that are using a netscaler load balancer device
func (s *NetworkService) ListNetscalerLoadBalancerNetworks(p *ListNetscalerLoadBalancerNetworksParams) (*ListNetscalerLoadBalancerNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNetscalerLoadBalancerNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetscalerLoadBalancerNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetscalerLoadBalancerNetworksResponse struct {
	Count                         int                             `json:"count"`
	NetscalerLoadBalancerNetworks []*NetscalerLoadBalancerNetwork `json:"netscalerloadbalancernetwork"`
}

type NetscalerLoadBalancerNetwork struct {
	Account                     string                                `json:"account"`
	Aclid                       string                                `json:"aclid"`
	Aclname                     string                                `json:"aclname"`
	Acltype                     string                                `json:"acltype"`
	Broadcastdomaintype         string                                `json:"broadcastdomaintype"`
	Broadcasturi                string                                `json:"broadcasturi"`
	Canusefordeploy             bool                                  `json:"canusefordeploy"`
	Cidr                        string                                `json:"cidr"`
	Created                     string                                `json:"created"`
	Details                     map[string]string                     `json:"details"`
	Displaynetwork              bool                                  `json:"displaynetwork"`
	Displaytext                 string                                `json:"displaytext"`
	Dns1                        string                                `json:"dns1"`
	Dns2                        string                                `json:"dns2"`
	Domain                      string                                `json:"domain"`
	Domainid                    string                                `json:"domainid"`
	Externalid                  string                                `json:"externalid"`
	Gateway                     string                                `json:"gateway"`
	Hasannotations              bool                                  `json:"hasannotations"`
	Icon                        string                                `json:"icon"`
	Id                          string                                `json:"id"`
	Ip6cidr                     string                                `json:"ip6cidr"`
	Ip6gateway                  string                                `json:"ip6gateway"`
	Isdefault                   bool                                  `json:"isdefault"`
	Ispersistent                bool                                  `json:"ispersistent"`
	Issystem                    bool                                  `json:"issystem"`
	JobID                       string                                `json:"jobid"`
	Jobstatus                   int                                   `json:"jobstatus"`
	Name                        string                                `json:"name"`
	Netmask                     string                                `json:"netmask"`
	Networkcidr                 string                                `json:"networkcidr"`
	Networkdomain               string                                `json:"networkdomain"`
	Networkofferingavailability string                                `json:"networkofferingavailability"`
	Networkofferingconservemode bool                                  `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                                `json:"networkofferingdisplaytext"`
	Networkofferingid           string                                `json:"networkofferingid"`
	Networkofferingname         string                                `json:"networkofferingname"`
	Physicalnetworkid           string                                `json:"physicalnetworkid"`
	Project                     string                                `json:"project"`
	Projectid                   string                                `json:"projectid"`
	Receivedbytes               int64                                 `json:"receivedbytes"`
	Redundantrouter             bool                                  `json:"redundantrouter"`
	Related                     string                                `json:"related"`
	Reservediprange             string                                `json:"reservediprange"`
	Restartrequired             bool                                  `json:"restartrequired"`
	Sentbytes                   int64                                 `json:"sentbytes"`
	Service                     []NetscalerLoadBalancerNetworkService `json:"service"`
	Specifyipranges             bool                                  `json:"specifyipranges"`
	State                       string                                `json:"state"`
	Strechedl2subnet            bool                                  `json:"strechedl2subnet"`
	Subdomainaccess             bool                                  `json:"subdomainaccess"`
	Tags                        []Tags                                `json:"tags"`
	Traffictype                 string                                `json:"traffictype"`
	Type                        string                                `json:"type"`
	Vlan                        string                                `json:"vlan"`
	Vpcid                       string                                `json:"vpcid"`
	Vpcname                     string                                `json:"vpcname"`
	Zoneid                      string                                `json:"zoneid"`
	Zonename                    string                                `json:"zonename"`
	Zonesnetworkspans           []interface{}                         `json:"zonesnetworkspans"`
}

type NetscalerLoadBalancerNetworkService struct {
	Capability []NetscalerLoadBalancerNetworkServiceCapability `json:"capability"`
	Name       string                                          `json:"name"`
	Provider   []NetscalerLoadBalancerNetworkServiceProvider   `json:"provider"`
}

type NetscalerLoadBalancerNetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type NetscalerLoadBalancerNetworkServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListNetworkIsolationMethodsParams struct {
	P map[string]interface{}
}

func (P *ListNetworkIsolationMethodsParams) toURLValues() url.Values {
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

func (P *ListNetworkIsolationMethodsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNetworkIsolationMethodsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNetworkIsolationMethodsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNetworkIsolationMethodsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNetworkIsolationMethodsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNetworkIsolationMethodsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListNetworkIsolationMethodsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworkIsolationMethodsParams() *ListNetworkIsolationMethodsParams {
	P := &ListNetworkIsolationMethodsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists supported methods of network isolation
func (s *NetworkService) ListNetworkIsolationMethods(p *ListNetworkIsolationMethodsParams) (*ListNetworkIsolationMethodsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkIsolationMethods", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkIsolationMethodsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkIsolationMethodsResponse struct {
	Count                   int                       `json:"count"`
	NetworkIsolationMethods []*NetworkIsolationMethod `json:"networkisolationmethod"`
}

type NetworkIsolationMethod struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type ListNetworkServiceProvidersParams struct {
	P map[string]interface{}
}

func (P *ListNetworkServiceProvidersParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
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
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (P *ListNetworkServiceProvidersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNetworkServiceProvidersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNetworkServiceProvidersParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListNetworkServiceProvidersParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListNetworkServiceProvidersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNetworkServiceProvidersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNetworkServiceProvidersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNetworkServiceProvidersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListNetworkServiceProvidersParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListNetworkServiceProvidersParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *ListNetworkServiceProvidersParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListNetworkServiceProvidersParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworkServiceProvidersParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworkServiceProvidersParams() *ListNetworkServiceProvidersParams {
	P := &ListNetworkServiceProvidersParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetworkServiceProviderID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListNetworkServiceProvidersParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetworkServiceProviders(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.NetworkServiceProviders[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NetworkServiceProviders {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// Lists network serviceproviders for a given physical network.
func (s *NetworkService) ListNetworkServiceProviders(p *ListNetworkServiceProvidersParams) (*ListNetworkServiceProvidersResponse, error) {
	resp, err := s.cs.newRequest("listNetworkServiceProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkServiceProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkServiceProvidersResponse struct {
	Count                   int                       `json:"count"`
	NetworkServiceProviders []*NetworkServiceProvider `json:"networkserviceprovider"`
}

type NetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	JobID                        string   `json:"jobid"`
	Jobstatus                    int      `json:"jobstatus"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type ListNetworksParams struct {
	P map[string]interface{}
}

func (P *ListNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["acltype"]; found {
		u.Set("acltype", v.(string))
	}
	if v, found := P.P["canusefordeploy"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("canusefordeploy", vv)
	}
	if v, found := P.P["displaynetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displaynetwork", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := P.P["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := P.P["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["restartrequired"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("restartrequired", vv)
	}
	if v, found := P.P["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := P.P["specifyipranges"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyipranges", vv)
	}
	if v, found := P.P["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
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
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListNetworksParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListNetworksParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetAcltype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["acltype"] = v
}

func (P *ListNetworksParams) GetAcltype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["acltype"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetCanusefordeploy(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["canusefordeploy"] = v
}

func (P *ListNetworksParams) GetCanusefordeploy() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["canusefordeploy"].(bool)
	return value, ok
}

func (P *ListNetworksParams) SetDisplaynetwork(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaynetwork"] = v
}

func (P *ListNetworksParams) GetDisplaynetwork() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaynetwork"].(bool)
	return value, ok
}

func (P *ListNetworksParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListNetworksParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetForvpc(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forvpc"] = v
}

func (P *ListNetworksParams) GetForvpc() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forvpc"].(bool)
	return value, ok
}

func (P *ListNetworksParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListNetworksParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListNetworksParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListNetworksParams) SetIssystem(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["issystem"] = v
}

func (P *ListNetworksParams) GetIssystem() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["issystem"].(bool)
	return value, ok
}

func (P *ListNetworksParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNetworksParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListNetworksParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListNetworksParams) SetNetworkofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkofferingid"] = v
}

func (P *ListNetworksParams) GetNetworkofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkofferingid"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNetworksParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNetworksParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNetworksParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListNetworksParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListNetworksParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListNetworksParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetRestartrequired(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["restartrequired"] = v
}

func (P *ListNetworksParams) GetRestartrequired() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["restartrequired"].(bool)
	return value, ok
}

func (P *ListNetworksParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListNetworksParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

func (P *ListNetworksParams) SetSpecifyipranges(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["specifyipranges"] = v
}

func (P *ListNetworksParams) GetSpecifyipranges() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["specifyipranges"].(bool)
	return value, ok
}

func (P *ListNetworksParams) SetSupportedservices(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["supportedservices"] = v
}

func (P *ListNetworksParams) GetSupportedservices() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["supportedservices"].([]string)
	return value, ok
}

func (P *ListNetworksParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListNetworksParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListNetworksParams) SetTraffictype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["traffictype"] = v
}

func (P *ListNetworksParams) GetTraffictype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["traffictype"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ListNetworksParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *ListNetworksParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

func (P *ListNetworksParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListNetworksParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNetworksParams() *ListNetworksParams {
	P := &ListNetworksParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetworkID(keyword string, opts ...OptionFunc) (string, int, error) {
	P := &ListNetworksParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetworks(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.Networks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Networks {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetworkByName(name string, opts ...OptionFunc) (*Network, int, error) {
	id, count, err := s.GetNetworkID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetNetworkByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNetworkByID(id string, opts ...OptionFunc) (*Network, int, error) {
	P := &ListNetworksParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListNetworks(P)
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
		return l.Networks[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Network UUID: %s!", id)
}

// Lists all available networks.
func (s *NetworkService) ListNetworks(p *ListNetworksParams) (*ListNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworksResponse struct {
	Count    int        `json:"count"`
	Networks []*Network `json:"network"`
}

type Network struct {
	Account                     string                   `json:"account"`
	Aclid                       string                   `json:"aclid"`
	Aclname                     string                   `json:"aclname"`
	Acltype                     string                   `json:"acltype"`
	Broadcastdomaintype         string                   `json:"broadcastdomaintype"`
	Broadcasturi                string                   `json:"broadcasturi"`
	Canusefordeploy             bool                     `json:"canusefordeploy"`
	Cidr                        string                   `json:"cidr"`
	Created                     string                   `json:"created"`
	Details                     map[string]string        `json:"details"`
	Displaynetwork              bool                     `json:"displaynetwork"`
	Displaytext                 string                   `json:"displaytext"`
	Dns1                        string                   `json:"dns1"`
	Dns2                        string                   `json:"dns2"`
	Domain                      string                   `json:"domain"`
	Domainid                    string                   `json:"domainid"`
	Externalid                  string                   `json:"externalid"`
	Gateway                     string                   `json:"gateway"`
	Hasannotations              bool                     `json:"hasannotations"`
	Icon                        string                   `json:"icon"`
	Id                          string                   `json:"id"`
	Ip6cidr                     string                   `json:"ip6cidr"`
	Ip6gateway                  string                   `json:"ip6gateway"`
	Isdefault                   bool                     `json:"isdefault"`
	Ispersistent                bool                     `json:"ispersistent"`
	Issystem                    bool                     `json:"issystem"`
	JobID                       string                   `json:"jobid"`
	Jobstatus                   int                      `json:"jobstatus"`
	Name                        string                   `json:"name"`
	Netmask                     string                   `json:"netmask"`
	Networkcidr                 string                   `json:"networkcidr"`
	Networkdomain               string                   `json:"networkdomain"`
	Networkofferingavailability string                   `json:"networkofferingavailability"`
	Networkofferingconservemode bool                     `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                   `json:"networkofferingdisplaytext"`
	Networkofferingid           string                   `json:"networkofferingid"`
	Networkofferingname         string                   `json:"networkofferingname"`
	Physicalnetworkid           string                   `json:"physicalnetworkid"`
	Project                     string                   `json:"project"`
	Projectid                   string                   `json:"projectid"`
	Receivedbytes               int64                    `json:"receivedbytes"`
	Redundantrouter             bool                     `json:"redundantrouter"`
	Related                     string                   `json:"related"`
	Reservediprange             string                   `json:"reservediprange"`
	Restartrequired             bool                     `json:"restartrequired"`
	Sentbytes                   int64                    `json:"sentbytes"`
	Service                     []NetworkServiceInternal `json:"service"`
	Specifyipranges             bool                     `json:"specifyipranges"`
	State                       string                   `json:"state"`
	Strechedl2subnet            bool                     `json:"strechedl2subnet"`
	Subdomainaccess             bool                     `json:"subdomainaccess"`
	Tags                        []Tags                   `json:"tags"`
	Traffictype                 string                   `json:"traffictype"`
	Type                        string                   `json:"type"`
	Vlan                        string                   `json:"vlan"`
	Vpcid                       string                   `json:"vpcid"`
	Vpcname                     string                   `json:"vpcname"`
	Zoneid                      string                   `json:"zoneid"`
	Zonename                    string                   `json:"zonename"`
	Zonesnetworkspans           []interface{}            `json:"zonesnetworkspans"`
}

type NetworkServiceInternal struct {
	Capability []NetworkServiceInternalCapability `json:"capability"`
	Name       string                             `json:"name"`
	Provider   []NetworkServiceInternalProvider   `json:"provider"`
}

type NetworkServiceInternalProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type NetworkServiceInternalCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListNiciraNvpDeviceNetworksParams struct {
	P map[string]interface{}
}

func (P *ListNiciraNvpDeviceNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["nvpdeviceid"]; found {
		u.Set("nvpdeviceid", v.(string))
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

func (P *ListNiciraNvpDeviceNetworksParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNiciraNvpDeviceNetworksParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNiciraNvpDeviceNetworksParams) SetNvpdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nvpdeviceid"] = v
}

func (P *ListNiciraNvpDeviceNetworksParams) GetNvpdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nvpdeviceid"].(string)
	return value, ok
}

func (P *ListNiciraNvpDeviceNetworksParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNiciraNvpDeviceNetworksParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNiciraNvpDeviceNetworksParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNiciraNvpDeviceNetworksParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListNiciraNvpDeviceNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListNiciraNvpDeviceNetworksParams(nvpdeviceid string) *ListNiciraNvpDeviceNetworksParams {
	P := &ListNiciraNvpDeviceNetworksParams{}
	P.P = make(map[string]interface{})
	P.P["nvpdeviceid"] = nvpdeviceid
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetNiciraNvpDeviceNetworkID(keyword string, nvpdeviceid string, opts ...OptionFunc) (string, int, error) {
	P := &ListNiciraNvpDeviceNetworksParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword
	P.P["nvpdeviceid"] = nvpdeviceid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNiciraNvpDeviceNetworks(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.NiciraNvpDeviceNetworks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NiciraNvpDeviceNetworks {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// lists network that are using a nicira nvp device
func (s *NetworkService) ListNiciraNvpDeviceNetworks(p *ListNiciraNvpDeviceNetworksParams) (*ListNiciraNvpDeviceNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNiciraNvpDeviceNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNiciraNvpDeviceNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNiciraNvpDeviceNetworksResponse struct {
	Count                   int                       `json:"count"`
	NiciraNvpDeviceNetworks []*NiciraNvpDeviceNetwork `json:"niciranvpdevicenetwork"`
}

type NiciraNvpDeviceNetwork struct {
	Account                     string                          `json:"account"`
	Aclid                       string                          `json:"aclid"`
	Aclname                     string                          `json:"aclname"`
	Acltype                     string                          `json:"acltype"`
	Broadcastdomaintype         string                          `json:"broadcastdomaintype"`
	Broadcasturi                string                          `json:"broadcasturi"`
	Canusefordeploy             bool                            `json:"canusefordeploy"`
	Cidr                        string                          `json:"cidr"`
	Created                     string                          `json:"created"`
	Details                     map[string]string               `json:"details"`
	Displaynetwork              bool                            `json:"displaynetwork"`
	Displaytext                 string                          `json:"displaytext"`
	Dns1                        string                          `json:"dns1"`
	Dns2                        string                          `json:"dns2"`
	Domain                      string                          `json:"domain"`
	Domainid                    string                          `json:"domainid"`
	Externalid                  string                          `json:"externalid"`
	Gateway                     string                          `json:"gateway"`
	Hasannotations              bool                            `json:"hasannotations"`
	Icon                        string                          `json:"icon"`
	Id                          string                          `json:"id"`
	Ip6cidr                     string                          `json:"ip6cidr"`
	Ip6gateway                  string                          `json:"ip6gateway"`
	Isdefault                   bool                            `json:"isdefault"`
	Ispersistent                bool                            `json:"ispersistent"`
	Issystem                    bool                            `json:"issystem"`
	JobID                       string                          `json:"jobid"`
	Jobstatus                   int                             `json:"jobstatus"`
	Name                        string                          `json:"name"`
	Netmask                     string                          `json:"netmask"`
	Networkcidr                 string                          `json:"networkcidr"`
	Networkdomain               string                          `json:"networkdomain"`
	Networkofferingavailability string                          `json:"networkofferingavailability"`
	Networkofferingconservemode bool                            `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                          `json:"networkofferingdisplaytext"`
	Networkofferingid           string                          `json:"networkofferingid"`
	Networkofferingname         string                          `json:"networkofferingname"`
	Physicalnetworkid           string                          `json:"physicalnetworkid"`
	Project                     string                          `json:"project"`
	Projectid                   string                          `json:"projectid"`
	Receivedbytes               int64                           `json:"receivedbytes"`
	Redundantrouter             bool                            `json:"redundantrouter"`
	Related                     string                          `json:"related"`
	Reservediprange             string                          `json:"reservediprange"`
	Restartrequired             bool                            `json:"restartrequired"`
	Sentbytes                   int64                           `json:"sentbytes"`
	Service                     []NiciraNvpDeviceNetworkService `json:"service"`
	Specifyipranges             bool                            `json:"specifyipranges"`
	State                       string                          `json:"state"`
	Strechedl2subnet            bool                            `json:"strechedl2subnet"`
	Subdomainaccess             bool                            `json:"subdomainaccess"`
	Tags                        []Tags                          `json:"tags"`
	Traffictype                 string                          `json:"traffictype"`
	Type                        string                          `json:"type"`
	Vlan                        string                          `json:"vlan"`
	Vpcid                       string                          `json:"vpcid"`
	Vpcname                     string                          `json:"vpcname"`
	Zoneid                      string                          `json:"zoneid"`
	Zonename                    string                          `json:"zonename"`
	Zonesnetworkspans           []interface{}                   `json:"zonesnetworkspans"`
}

type NiciraNvpDeviceNetworkService struct {
	Capability []NiciraNvpDeviceNetworkServiceCapability `json:"capability"`
	Name       string                                    `json:"name"`
	Provider   []NiciraNvpDeviceNetworkServiceProvider   `json:"provider"`
}

type NiciraNvpDeviceNetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type NiciraNvpDeviceNetworkServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListOpenDaylightControllersParams struct {
	P map[string]interface{}
}

func (P *ListOpenDaylightControllersParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (P *ListOpenDaylightControllersParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListOpenDaylightControllersParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListOpenDaylightControllersParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListOpenDaylightControllersParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListOpenDaylightControllersParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListOpenDaylightControllersParams() *ListOpenDaylightControllersParams {
	P := &ListOpenDaylightControllersParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetOpenDaylightControllerByID(id string, opts ...OptionFunc) (*OpenDaylightController, int, error) {
	P := &ListOpenDaylightControllersParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListOpenDaylightControllers(P)
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
		return l.OpenDaylightControllers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for OpenDaylightController UUID: %s!", id)
}

// Lists OpenDyalight controllers
func (s *NetworkService) ListOpenDaylightControllers(p *ListOpenDaylightControllersParams) (*ListOpenDaylightControllersResponse, error) {
	resp, err := s.cs.newRequest("listOpenDaylightControllers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListOpenDaylightControllersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListOpenDaylightControllersResponse struct {
	Count                   int                       `json:"count"`
	OpenDaylightControllers []*OpenDaylightController `json:"opendaylightcontroller"`
}

type OpenDaylightController struct {
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Name              string `json:"name"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Url               string `json:"url"`
	Username          string `json:"username"`
}

type ListPaloAltoFirewallNetworksParams struct {
	P map[string]interface{}
}

func (P *ListPaloAltoFirewallNetworksParams) toURLValues() url.Values {
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
	return u
}

func (P *ListPaloAltoFirewallNetworksParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListPaloAltoFirewallNetworksParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListPaloAltoFirewallNetworksParams) SetLbdeviceid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbdeviceid"] = v
}

func (P *ListPaloAltoFirewallNetworksParams) GetLbdeviceid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbdeviceid"].(string)
	return value, ok
}

func (P *ListPaloAltoFirewallNetworksParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListPaloAltoFirewallNetworksParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListPaloAltoFirewallNetworksParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListPaloAltoFirewallNetworksParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListPaloAltoFirewallNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListPaloAltoFirewallNetworksParams(lbdeviceid string) *ListPaloAltoFirewallNetworksParams {
	P := &ListPaloAltoFirewallNetworksParams{}
	P.P = make(map[string]interface{})
	P.P["lbdeviceid"] = lbdeviceid
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetPaloAltoFirewallNetworkID(keyword string, lbdeviceid string, opts ...OptionFunc) (string, int, error) {
	P := &ListPaloAltoFirewallNetworksParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword
	P.P["lbdeviceid"] = lbdeviceid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListPaloAltoFirewallNetworks(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.PaloAltoFirewallNetworks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.PaloAltoFirewallNetworks {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// lists network that are using Palo Alto firewall device
func (s *NetworkService) ListPaloAltoFirewallNetworks(p *ListPaloAltoFirewallNetworksParams) (*ListPaloAltoFirewallNetworksResponse, error) {
	resp, err := s.cs.newRequest("listPaloAltoFirewallNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPaloAltoFirewallNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPaloAltoFirewallNetworksResponse struct {
	Count                    int                        `json:"count"`
	PaloAltoFirewallNetworks []*PaloAltoFirewallNetwork `json:"paloaltofirewallnetwork"`
}

type PaloAltoFirewallNetwork struct {
	Account                     string                           `json:"account"`
	Aclid                       string                           `json:"aclid"`
	Aclname                     string                           `json:"aclname"`
	Acltype                     string                           `json:"acltype"`
	Broadcastdomaintype         string                           `json:"broadcastdomaintype"`
	Broadcasturi                string                           `json:"broadcasturi"`
	Canusefordeploy             bool                             `json:"canusefordeploy"`
	Cidr                        string                           `json:"cidr"`
	Created                     string                           `json:"created"`
	Details                     map[string]string                `json:"details"`
	Displaynetwork              bool                             `json:"displaynetwork"`
	Displaytext                 string                           `json:"displaytext"`
	Dns1                        string                           `json:"dns1"`
	Dns2                        string                           `json:"dns2"`
	Domain                      string                           `json:"domain"`
	Domainid                    string                           `json:"domainid"`
	Externalid                  string                           `json:"externalid"`
	Gateway                     string                           `json:"gateway"`
	Hasannotations              bool                             `json:"hasannotations"`
	Icon                        string                           `json:"icon"`
	Id                          string                           `json:"id"`
	Ip6cidr                     string                           `json:"ip6cidr"`
	Ip6gateway                  string                           `json:"ip6gateway"`
	Isdefault                   bool                             `json:"isdefault"`
	Ispersistent                bool                             `json:"ispersistent"`
	Issystem                    bool                             `json:"issystem"`
	JobID                       string                           `json:"jobid"`
	Jobstatus                   int                              `json:"jobstatus"`
	Name                        string                           `json:"name"`
	Netmask                     string                           `json:"netmask"`
	Networkcidr                 string                           `json:"networkcidr"`
	Networkdomain               string                           `json:"networkdomain"`
	Networkofferingavailability string                           `json:"networkofferingavailability"`
	Networkofferingconservemode bool                             `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                           `json:"networkofferingdisplaytext"`
	Networkofferingid           string                           `json:"networkofferingid"`
	Networkofferingname         string                           `json:"networkofferingname"`
	Physicalnetworkid           string                           `json:"physicalnetworkid"`
	Project                     string                           `json:"project"`
	Projectid                   string                           `json:"projectid"`
	Receivedbytes               int64                            `json:"receivedbytes"`
	Redundantrouter             bool                             `json:"redundantrouter"`
	Related                     string                           `json:"related"`
	Reservediprange             string                           `json:"reservediprange"`
	Restartrequired             bool                             `json:"restartrequired"`
	Sentbytes                   int64                            `json:"sentbytes"`
	Service                     []PaloAltoFirewallNetworkService `json:"service"`
	Specifyipranges             bool                             `json:"specifyipranges"`
	State                       string                           `json:"state"`
	Strechedl2subnet            bool                             `json:"strechedl2subnet"`
	Subdomainaccess             bool                             `json:"subdomainaccess"`
	Tags                        []Tags                           `json:"tags"`
	Traffictype                 string                           `json:"traffictype"`
	Type                        string                           `json:"type"`
	Vlan                        string                           `json:"vlan"`
	Vpcid                       string                           `json:"vpcid"`
	Vpcname                     string                           `json:"vpcname"`
	Zoneid                      string                           `json:"zoneid"`
	Zonename                    string                           `json:"zonename"`
	Zonesnetworkspans           []interface{}                    `json:"zonesnetworkspans"`
}

type PaloAltoFirewallNetworkService struct {
	Capability []PaloAltoFirewallNetworkServiceCapability `json:"capability"`
	Name       string                                     `json:"name"`
	Provider   []PaloAltoFirewallNetworkServiceProvider   `json:"provider"`
}

type PaloAltoFirewallNetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type PaloAltoFirewallNetworkServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ListPhysicalNetworksParams struct {
	P map[string]interface{}
}

func (P *ListPhysicalNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListPhysicalNetworksParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListPhysicalNetworksParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListPhysicalNetworksParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListPhysicalNetworksParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListPhysicalNetworksParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListPhysicalNetworksParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListPhysicalNetworksParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListPhysicalNetworksParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListPhysicalNetworksParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListPhysicalNetworksParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListPhysicalNetworksParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListPhysicalNetworksParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListPhysicalNetworksParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListPhysicalNetworksParams() *ListPhysicalNetworksParams {
	P := &ListPhysicalNetworksParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetPhysicalNetworkID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListPhysicalNetworksParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListPhysicalNetworks(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.PhysicalNetworks[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.PhysicalNetworks {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetPhysicalNetworkByName(name string, opts ...OptionFunc) (*PhysicalNetwork, int, error) {
	id, count, err := s.GetPhysicalNetworkID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetPhysicalNetworkByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetPhysicalNetworkByID(id string, opts ...OptionFunc) (*PhysicalNetwork, int, error) {
	P := &ListPhysicalNetworksParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListPhysicalNetworks(P)
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
		return l.PhysicalNetworks[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for PhysicalNetwork UUID: %s!", id)
}

// Lists physical networks
func (s *NetworkService) ListPhysicalNetworks(p *ListPhysicalNetworksParams) (*ListPhysicalNetworksResponse, error) {
	resp, err := s.cs.newRequest("listPhysicalNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListPhysicalNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPhysicalNetworksResponse struct {
	Count            int                `json:"count"`
	PhysicalNetworks []*PhysicalNetwork `json:"physicalnetwork"`
}

type PhysicalNetwork struct {
	Broadcastdomainrange string `json:"broadcastdomainrange"`
	Domainid             string `json:"domainid"`
	Id                   string `json:"id"`
	Isolationmethods     string `json:"isolationmethods"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Name                 string `json:"name"`
	Networkspeed         string `json:"networkspeed"`
	State                string `json:"state"`
	Tags                 string `json:"tags"`
	Vlan                 string `json:"vlan"`
	Zoneid               string `json:"zoneid"`
	Zonename             string `json:"zonename"`
}

type ListStorageNetworkIpRangeParams struct {
	P map[string]interface{}
}

func (P *ListStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListStorageNetworkIpRangeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListStorageNetworkIpRangeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListStorageNetworkIpRangeParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListStorageNetworkIpRangeParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListStorageNetworkIpRangeParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListStorageNetworkIpRangeParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListStorageNetworkIpRangeParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListStorageNetworkIpRangeParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListStorageNetworkIpRangeParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListStorageNetworkIpRangeParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListStorageNetworkIpRangeParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListStorageNetworkIpRangeParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListStorageNetworkIpRangeParams() *ListStorageNetworkIpRangeParams {
	P := &ListStorageNetworkIpRangeParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkService) GetStorageNetworkIpRangeByID(id string, opts ...OptionFunc) (*StorageNetworkIpRange, int, error) {
	P := &ListStorageNetworkIpRangeParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListStorageNetworkIpRange(P)
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
		return l.StorageNetworkIpRange[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for StorageNetworkIpRange UUID: %s!", id)
}

// List a storage network IP range.
func (s *NetworkService) ListStorageNetworkIpRange(p *ListStorageNetworkIpRangeParams) (*ListStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("listStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStorageNetworkIpRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStorageNetworkIpRangeResponse struct {
	Count                 int                      `json:"count"`
	StorageNetworkIpRange []*StorageNetworkIpRange `json:"storagenetworkiprange"`
}

type StorageNetworkIpRange struct {
	Endip     string `json:"endip"`
	Gateway   string `json:"gateway"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Netmask   string `json:"netmask"`
	Networkid string `json:"networkid"`
	Podid     string `json:"podid"`
	Startip   string `json:"startip"`
	Vlan      int    `json:"vlan"`
	Zoneid    string `json:"zoneid"`
}

type ListSupportedNetworkServicesParams struct {
	P map[string]interface{}
}

func (P *ListSupportedNetworkServicesParams) toURLValues() url.Values {
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
	if v, found := P.P["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := P.P["service"]; found {
		u.Set("service", v.(string))
	}
	return u
}

func (P *ListSupportedNetworkServicesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListSupportedNetworkServicesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListSupportedNetworkServicesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListSupportedNetworkServicesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListSupportedNetworkServicesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListSupportedNetworkServicesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListSupportedNetworkServicesParams) SetProvider(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["provider"] = v
}

func (P *ListSupportedNetworkServicesParams) GetProvider() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["provider"].(string)
	return value, ok
}

func (P *ListSupportedNetworkServicesParams) SetService(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["service"] = v
}

func (P *ListSupportedNetworkServicesParams) GetService() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["service"].(string)
	return value, ok
}

// You should always use this function to get a new ListSupportedNetworkServicesParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewListSupportedNetworkServicesParams() *ListSupportedNetworkServicesParams {
	P := &ListSupportedNetworkServicesParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists all network services provided by CloudStack or for the given Provider.
func (s *NetworkService) ListSupportedNetworkServices(p *ListSupportedNetworkServicesParams) (*ListSupportedNetworkServicesResponse, error) {
	resp, err := s.cs.newRequest("listSupportedNetworkServices", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSupportedNetworkServicesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSupportedNetworkServicesResponse struct {
	Count                    int                        `json:"count"`
	SupportedNetworkServices []*SupportedNetworkService `json:"supportednetworkservice"`
}

type SupportedNetworkService struct {
	Capability []SupportedNetworkServiceCapability `json:"capability"`
	JobID      string                              `json:"jobid"`
	Jobstatus  int                                 `json:"jobstatus"`
	Name       string                              `json:"name"`
	Provider   []SupportedNetworkServiceProvider   `json:"provider"`
}

type SupportedNetworkServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type SupportedNetworkServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type ReleasePublicIpRangeParams struct {
	P map[string]interface{}
}

func (P *ReleasePublicIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ReleasePublicIpRangeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ReleasePublicIpRangeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ReleasePublicIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewReleasePublicIpRangeParams(id string) *ReleasePublicIpRangeParams {
	P := &ReleasePublicIpRangeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Releases a Public IP range back to the system pool
func (s *NetworkService) ReleasePublicIpRange(p *ReleasePublicIpRangeParams) (*ReleasePublicIpRangeResponse, error) {
	resp, err := s.cs.newRequest("releasePublicIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleasePublicIpRangeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ReleasePublicIpRangeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ReleasePublicIpRangeResponse) UnmarshalJSON(b []byte) error {
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

	type alias ReleasePublicIpRangeResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RestartNetworkParams struct {
	P map[string]interface{}
}

func (P *RestartNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["cleanup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanup", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["makeredundant"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("makeredundant", vv)
	}
	return u
}

func (P *RestartNetworkParams) SetCleanup(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cleanup"] = v
}

func (P *RestartNetworkParams) GetCleanup() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cleanup"].(bool)
	return value, ok
}

func (P *RestartNetworkParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RestartNetworkParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *RestartNetworkParams) SetMakeredundant(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["makeredundant"] = v
}

func (P *RestartNetworkParams) GetMakeredundant() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["makeredundant"].(bool)
	return value, ok
}

// You should always use this function to get a new RestartNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewRestartNetworkParams(id string) *RestartNetworkParams {
	P := &RestartNetworkParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Restarts the network; includes 1) restarting network elements - virtual routers, DHCP servers 2) reapplying all public IPs 3) reapplying loadBalancing/portForwarding rules
func (s *NetworkService) RestartNetwork(p *RestartNetworkParams) (*RestartNetworkResponse, error) {
	resp, err := s.cs.newRequest("restartNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestartNetworkResponse
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

type RestartNetworkResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateNetworkParams struct {
	P map[string]interface{}
}

func (P *UpdateNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["changecidr"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("changecidr", vv)
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["displaynetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displaynetwork", vv)
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := P.P["guestvmcidr"]; found {
		u.Set("guestvmcidr", v.(string))
	}
	if v, found := P.P["hideipaddressusage"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("hideipaddressusage", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networkdomain"]; found {
		u.Set("networkdomain", v.(string))
	}
	if v, found := P.P["networkofferingid"]; found {
		u.Set("networkofferingid", v.(string))
	}
	if v, found := P.P["updateinsequence"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("updateinsequence", vv)
	}
	return u
}

func (P *UpdateNetworkParams) SetChangecidr(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["changecidr"] = v
}

func (P *UpdateNetworkParams) GetChangecidr() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["changecidr"].(bool)
	return value, ok
}

func (P *UpdateNetworkParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateNetworkParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateNetworkParams) SetDisplaynetwork(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaynetwork"] = v
}

func (P *UpdateNetworkParams) GetDisplaynetwork() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaynetwork"].(bool)
	return value, ok
}

func (P *UpdateNetworkParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *UpdateNetworkParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *UpdateNetworkParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *UpdateNetworkParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *UpdateNetworkParams) SetGuestvmcidr(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["guestvmcidr"] = v
}

func (P *UpdateNetworkParams) GetGuestvmcidr() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["guestvmcidr"].(string)
	return value, ok
}

func (P *UpdateNetworkParams) SetHideipaddressusage(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hideipaddressusage"] = v
}

func (P *UpdateNetworkParams) GetHideipaddressusage() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hideipaddressusage"].(bool)
	return value, ok
}

func (P *UpdateNetworkParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateNetworkParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateNetworkParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateNetworkParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateNetworkParams) SetNetworkdomain(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkdomain"] = v
}

func (P *UpdateNetworkParams) GetNetworkdomain() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkdomain"].(string)
	return value, ok
}

func (P *UpdateNetworkParams) SetNetworkofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkofferingid"] = v
}

func (P *UpdateNetworkParams) GetNetworkofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkofferingid"].(string)
	return value, ok
}

func (P *UpdateNetworkParams) SetUpdateinsequence(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["updateinsequence"] = v
}

func (P *UpdateNetworkParams) GetUpdateinsequence() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["updateinsequence"].(bool)
	return value, ok
}

// You should always use this function to get a new UpdateNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdateNetworkParams(id string) *UpdateNetworkParams {
	P := &UpdateNetworkParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a network
func (s *NetworkService) UpdateNetwork(p *UpdateNetworkParams) (*UpdateNetworkResponse, error) {
	resp, err := s.cs.newRequest("updateNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkResponse
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

type UpdateNetworkResponse struct {
	Account                     string                         `json:"account"`
	Aclid                       string                         `json:"aclid"`
	Aclname                     string                         `json:"aclname"`
	Acltype                     string                         `json:"acltype"`
	Broadcastdomaintype         string                         `json:"broadcastdomaintype"`
	Broadcasturi                string                         `json:"broadcasturi"`
	Canusefordeploy             bool                           `json:"canusefordeploy"`
	Cidr                        string                         `json:"cidr"`
	Created                     string                         `json:"created"`
	Details                     map[string]string              `json:"details"`
	Displaynetwork              bool                           `json:"displaynetwork"`
	Displaytext                 string                         `json:"displaytext"`
	Dns1                        string                         `json:"dns1"`
	Dns2                        string                         `json:"dns2"`
	Domain                      string                         `json:"domain"`
	Domainid                    string                         `json:"domainid"`
	Externalid                  string                         `json:"externalid"`
	Gateway                     string                         `json:"gateway"`
	Hasannotations              bool                           `json:"hasannotations"`
	Icon                        string                         `json:"icon"`
	Id                          string                         `json:"id"`
	Ip6cidr                     string                         `json:"ip6cidr"`
	Ip6gateway                  string                         `json:"ip6gateway"`
	Isdefault                   bool                           `json:"isdefault"`
	Ispersistent                bool                           `json:"ispersistent"`
	Issystem                    bool                           `json:"issystem"`
	JobID                       string                         `json:"jobid"`
	Jobstatus                   int                            `json:"jobstatus"`
	Name                        string                         `json:"name"`
	Netmask                     string                         `json:"netmask"`
	Networkcidr                 string                         `json:"networkcidr"`
	Networkdomain               string                         `json:"networkdomain"`
	Networkofferingavailability string                         `json:"networkofferingavailability"`
	Networkofferingconservemode bool                           `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  string                         `json:"networkofferingdisplaytext"`
	Networkofferingid           string                         `json:"networkofferingid"`
	Networkofferingname         string                         `json:"networkofferingname"`
	Physicalnetworkid           string                         `json:"physicalnetworkid"`
	Project                     string                         `json:"project"`
	Projectid                   string                         `json:"projectid"`
	Receivedbytes               int64                          `json:"receivedbytes"`
	Redundantrouter             bool                           `json:"redundantrouter"`
	Related                     string                         `json:"related"`
	Reservediprange             string                         `json:"reservediprange"`
	Restartrequired             bool                           `json:"restartrequired"`
	Sentbytes                   int64                          `json:"sentbytes"`
	Service                     []UpdateNetworkResponseService `json:"service"`
	Specifyipranges             bool                           `json:"specifyipranges"`
	State                       string                         `json:"state"`
	Strechedl2subnet            bool                           `json:"strechedl2subnet"`
	Subdomainaccess             bool                           `json:"subdomainaccess"`
	Tags                        []Tags                         `json:"tags"`
	Traffictype                 string                         `json:"traffictype"`
	Type                        string                         `json:"type"`
	Vlan                        string                         `json:"vlan"`
	Vpcid                       string                         `json:"vpcid"`
	Vpcname                     string                         `json:"vpcname"`
	Zoneid                      string                         `json:"zoneid"`
	Zonename                    string                         `json:"zonename"`
	Zonesnetworkspans           []interface{}                  `json:"zonesnetworkspans"`
}

type UpdateNetworkResponseService struct {
	Capability []UpdateNetworkResponseServiceCapability `json:"capability"`
	Name       string                                   `json:"name"`
	Provider   []UpdateNetworkResponseServiceProvider   `json:"provider"`
}

type UpdateNetworkResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type UpdateNetworkResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type UpdateNetworkServiceProviderParams struct {
	P map[string]interface{}
}

func (P *UpdateNetworkServiceProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["servicelist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("servicelist", vv)
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (P *UpdateNetworkServiceProviderParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateNetworkServiceProviderParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateNetworkServiceProviderParams) SetServicelist(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["servicelist"] = v
}

func (P *UpdateNetworkServiceProviderParams) GetServicelist() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["servicelist"].([]string)
	return value, ok
}

func (P *UpdateNetworkServiceProviderParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *UpdateNetworkServiceProviderParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateNetworkServiceProviderParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdateNetworkServiceProviderParams(id string) *UpdateNetworkServiceProviderParams {
	P := &UpdateNetworkServiceProviderParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a network serviceProvider of a physical network
func (s *NetworkService) UpdateNetworkServiceProvider(p *UpdateNetworkServiceProviderParams) (*UpdateNetworkServiceProviderResponse, error) {
	resp, err := s.cs.newRequest("updateNetworkServiceProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkServiceProviderResponse
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

type UpdateNetworkServiceProviderResponse struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	JobID                        string   `json:"jobid"`
	Jobstatus                    int      `json:"jobstatus"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type UpdatePhysicalNetworkParams struct {
	P map[string]interface{}
}

func (P *UpdatePhysicalNetworkParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["networkspeed"]; found {
		u.Set("networkspeed", v.(string))
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["tags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("tags", vv)
	}
	if v, found := P.P["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	return u
}

func (P *UpdatePhysicalNetworkParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdatePhysicalNetworkParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdatePhysicalNetworkParams) SetNetworkspeed(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkspeed"] = v
}

func (P *UpdatePhysicalNetworkParams) GetNetworkspeed() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkspeed"].(string)
	return value, ok
}

func (P *UpdatePhysicalNetworkParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *UpdatePhysicalNetworkParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *UpdatePhysicalNetworkParams) SetTags(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *UpdatePhysicalNetworkParams) GetTags() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].([]string)
	return value, ok
}

func (P *UpdatePhysicalNetworkParams) SetVlan(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vlan"] = v
}

func (P *UpdatePhysicalNetworkParams) GetVlan() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vlan"].(string)
	return value, ok
}

// You should always use this function to get a new UpdatePhysicalNetworkParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdatePhysicalNetworkParams(id string) *UpdatePhysicalNetworkParams {
	P := &UpdatePhysicalNetworkParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a physical network
func (s *NetworkService) UpdatePhysicalNetwork(p *UpdatePhysicalNetworkParams) (*UpdatePhysicalNetworkResponse, error) {
	resp, err := s.cs.newRequest("updatePhysicalNetwork", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdatePhysicalNetworkResponse
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

type UpdatePhysicalNetworkResponse struct {
	Broadcastdomainrange string `json:"broadcastdomainrange"`
	Domainid             string `json:"domainid"`
	Id                   string `json:"id"`
	Isolationmethods     string `json:"isolationmethods"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Name                 string `json:"name"`
	Networkspeed         string `json:"networkspeed"`
	State                string `json:"state"`
	Tags                 string `json:"tags"`
	Vlan                 string `json:"vlan"`
	Zoneid               string `json:"zoneid"`
	Zonename             string `json:"zonename"`
}

type UpdateStorageNetworkIpRangeParams struct {
	P map[string]interface{}
}

func (P *UpdateStorageNetworkIpRangeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["endip"]; found {
		u.Set("endip", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["netmask"]; found {
		u.Set("netmask", v.(string))
	}
	if v, found := P.P["startip"]; found {
		u.Set("startip", v.(string))
	}
	if v, found := P.P["vlan"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("vlan", vv)
	}
	return u
}

func (P *UpdateStorageNetworkIpRangeParams) SetEndip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["endip"] = v
}

func (P *UpdateStorageNetworkIpRangeParams) GetEndip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["endip"].(string)
	return value, ok
}

func (P *UpdateStorageNetworkIpRangeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateStorageNetworkIpRangeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateStorageNetworkIpRangeParams) SetNetmask(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["netmask"] = v
}

func (P *UpdateStorageNetworkIpRangeParams) GetNetmask() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["netmask"].(string)
	return value, ok
}

func (P *UpdateStorageNetworkIpRangeParams) SetStartip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startip"] = v
}

func (P *UpdateStorageNetworkIpRangeParams) GetStartip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startip"].(string)
	return value, ok
}

func (P *UpdateStorageNetworkIpRangeParams) SetVlan(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vlan"] = v
}

func (P *UpdateStorageNetworkIpRangeParams) GetVlan() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vlan"].(int)
	return value, ok
}

// You should always use this function to get a new UpdateStorageNetworkIpRangeParams instance,
// as then you are sure you have configured all required params
func (s *NetworkService) NewUpdateStorageNetworkIpRangeParams(id string) *UpdateStorageNetworkIpRangeParams {
	P := &UpdateStorageNetworkIpRangeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Update a Storage network IP range, only allowed when no IPs in this range have been allocated.
func (s *NetworkService) UpdateStorageNetworkIpRange(p *UpdateStorageNetworkIpRangeParams) (*UpdateStorageNetworkIpRangeResponse, error) {
	resp, err := s.cs.newRequest("updateStorageNetworkIpRange", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateStorageNetworkIpRangeResponse
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

type UpdateStorageNetworkIpRangeResponse struct {
	Endip     string `json:"endip"`
	Gateway   string `json:"gateway"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Netmask   string `json:"netmask"`
	Networkid string `json:"networkid"`
	Podid     string `json:"podid"`
	Startip   string `json:"startip"`
	Vlan      int    `json:"vlan"`
	Zoneid    string `json:"zoneid"`
}
