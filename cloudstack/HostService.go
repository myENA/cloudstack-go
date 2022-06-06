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

type HostServiceIface interface {
	AddBaremetalHost(P *AddBaremetalHostParams) (*AddBaremetalHostResponse, error)
	NewAddBaremetalHostParams(hypervisor string, podid string, url string, zoneid string) *AddBaremetalHostParams
	AddGloboDnsHost(P *AddGloboDnsHostParams) (*AddGloboDnsHostResponse, error)
	NewAddGloboDnsHostParams(password string, physicalnetworkid string, url string, username string) *AddGloboDnsHostParams
	AddHost(P *AddHostParams) (*AddHostResponse, error)
	NewAddHostParams(hypervisor string, podid string, url string, zoneid string) *AddHostParams
	AddSecondaryStorage(P *AddSecondaryStorageParams) (*AddSecondaryStorageResponse, error)
	NewAddSecondaryStorageParams(url string) *AddSecondaryStorageParams
	CancelHostMaintenance(P *CancelHostMaintenanceParams) (*CancelHostMaintenanceResponse, error)
	NewCancelHostMaintenanceParams(id string) *CancelHostMaintenanceParams
	ConfigureHAForHost(P *ConfigureHAForHostParams) (*HAForHostResponse, error)
	NewConfigureHAForHostParams(hostid string, provider string) *ConfigureHAForHostParams
	EnableHAForHost(P *EnableHAForHostParams) (*EnableHAForHostResponse, error)
	NewEnableHAForHostParams(hostid string) *EnableHAForHostParams
	DedicateHost(P *DedicateHostParams) (*DedicateHostResponse, error)
	NewDedicateHostParams(domainid string, hostid string) *DedicateHostParams
	DeleteHost(P *DeleteHostParams) (*DeleteHostResponse, error)
	NewDeleteHostParams(id string) *DeleteHostParams
	DisableOutOfBandManagementForHost(P *DisableOutOfBandManagementForHostParams) (*DisableOutOfBandManagementForHostResponse, error)
	NewDisableOutOfBandManagementForHostParams(hostid string) *DisableOutOfBandManagementForHostParams
	EnableOutOfBandManagementForHost(P *EnableOutOfBandManagementForHostParams) (*EnableOutOfBandManagementForHostResponse, error)
	NewEnableOutOfBandManagementForHostParams(hostid string) *EnableOutOfBandManagementForHostParams
	FindHostsForMigration(P *FindHostsForMigrationParams) (*FindHostsForMigrationResponse, error)
	NewFindHostsForMigrationParams(virtualmachineid string) *FindHostsForMigrationParams
	ListDedicatedHosts(P *ListDedicatedHostsParams) (*ListDedicatedHostsResponse, error)
	NewListDedicatedHostsParams() *ListDedicatedHostsParams
	ListHostTags(P *ListHostTagsParams) (*ListHostTagsResponse, error)
	NewListHostTagsParams() *ListHostTagsParams
	GetHostTagID(keyword string, opts ...OptionFunc) (string, int, error)
	ListHosts(P *ListHostsParams) (*ListHostsResponse, error)
	NewListHostsParams() *ListHostsParams
	GetHostID(name string, opts ...OptionFunc) (string, int, error)
	GetHostByName(name string, opts ...OptionFunc) (*Host, int, error)
	GetHostByID(id string, opts ...OptionFunc) (*Host, int, error)
	ListHostsMetrics(P *ListHostsMetricsParams) (*ListHostsMetricsResponse, error)
	NewListHostsMetricsParams() *ListHostsMetricsParams
	GetHostsMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetHostsMetricByName(name string, opts ...OptionFunc) (*HostsMetric, int, error)
	GetHostsMetricByID(id string, opts ...OptionFunc) (*HostsMetric, int, error)
	PrepareHostForMaintenance(P *PrepareHostForMaintenanceParams) (*PrepareHostForMaintenanceResponse, error)
	NewPrepareHostForMaintenanceParams(id string) *PrepareHostForMaintenanceParams
	ReconnectHost(P *ReconnectHostParams) (*ReconnectHostResponse, error)
	NewReconnectHostParams(id string) *ReconnectHostParams
	ReleaseDedicatedHost(P *ReleaseDedicatedHostParams) (*ReleaseDedicatedHostResponse, error)
	NewReleaseDedicatedHostParams(hostid string) *ReleaseDedicatedHostParams
	ReleaseHostReservation(P *ReleaseHostReservationParams) (*ReleaseHostReservationResponse, error)
	NewReleaseHostReservationParams(id string) *ReleaseHostReservationParams
	UpdateHost(P *UpdateHostParams) (*UpdateHostResponse, error)
	NewUpdateHostParams(id string) *UpdateHostParams
	UpdateHostPassword(P *UpdateHostPasswordParams) (*UpdateHostPasswordResponse, error)
	NewUpdateHostPasswordParams(password string, username string) *UpdateHostPasswordParams
}

type AddBaremetalHostParams struct {
	P map[string]interface{}
}

func (P *AddBaremetalHostParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := P.P["hosttags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("hosttags", vv)
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *AddBaremetalHostParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *AddBaremetalHostParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *AddBaremetalHostParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *AddBaremetalHostParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *AddBaremetalHostParams) SetClustername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clustername"] = v
}

func (P *AddBaremetalHostParams) GetClustername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clustername"].(string)
	return value, ok
}

func (P *AddBaremetalHostParams) SetHosttags(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hosttags"] = v
}

func (P *AddBaremetalHostParams) GetHosttags() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hosttags"].([]string)
	return value, ok
}

func (P *AddBaremetalHostParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *AddBaremetalHostParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *AddBaremetalHostParams) SetIpaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddress"] = v
}

func (P *AddBaremetalHostParams) GetIpaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddress"].(string)
	return value, ok
}

func (P *AddBaremetalHostParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddBaremetalHostParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddBaremetalHostParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *AddBaremetalHostParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *AddBaremetalHostParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddBaremetalHostParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddBaremetalHostParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddBaremetalHostParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

func (P *AddBaremetalHostParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *AddBaremetalHostParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddBaremetalHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddBaremetalHostParams(hypervisor string, podid string, url string, zoneid string) *AddBaremetalHostParams {
	P := &AddBaremetalHostParams{}
	P.P = make(map[string]interface{})
	P.P["hypervisor"] = hypervisor
	P.P["podid"] = podid
	P.P["url"] = url
	P.P["zoneid"] = zoneid
	return P
}

// add a baremetal host
func (s *HostService) AddBaremetalHost(p *AddBaremetalHostParams) (*AddBaremetalHostResponse, error) {
	resp, err := s.cs.newRequest("addBaremetalHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddBaremetalHostResponse struct {
	Annotation                       string                             `json:"annotation"`
	Capabilities                     string                             `json:"capabilities"`
	Clusterid                        string                             `json:"clusterid"`
	Clustername                      string                             `json:"clustername"`
	Clustertype                      string                             `json:"clustertype"`
	Cpuallocated                     string                             `json:"cpuallocated"`
	Cpuallocatedpercentage           string                             `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                              `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                             `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                            `json:"cpuloadaverage"`
	Cpunumber                        int                                `json:"cpunumber"`
	Cpusockets                       int                                `json:"cpusockets"`
	Cpuspeed                         int64                              `json:"cpuspeed"`
	Cpuused                          string                             `json:"cpuused"`
	Cpuwithoverprovisioning          string                             `json:"cpuwithoverprovisioning"`
	Created                          string                             `json:"created"`
	Details                          map[string]string                  `json:"details"`
	Disconnected                     string                             `json:"disconnected"`
	Disksizeallocated                int64                              `json:"disksizeallocated"`
	Disksizetotal                    int64                              `json:"disksizetotal"`
	Events                           string                             `json:"events"`
	Gpugroup                         []AddBaremetalHostResponseGpugroup `json:"gpugroup"`
	Hahost                           bool                               `json:"hahost"`
	Hasannotations                   bool                               `json:"hasannotations"`
	Hasenoughcapacity                bool                               `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse                  `json:"hostha"`
	Hosttags                         string                             `json:"hosttags"`
	Hypervisor                       string                             `json:"hypervisor"`
	Hypervisorversion                string                             `json:"hypervisorversion"`
	Id                               string                             `json:"id"`
	Ipaddress                        string                             `json:"ipaddress"`
	Islocalstorageactive             bool                               `json:"islocalstorageactive"`
	JobID                            string                             `json:"jobid"`
	Jobstatus                        int                                `json:"jobstatus"`
	Lastannotated                    string                             `json:"lastannotated"`
	Lastpinged                       string                             `json:"lastpinged"`
	Managementserverid               UUID                               `json:"managementserverid"`
	Memoryallocated                  int64                              `json:"memoryallocated"`
	Memoryallocatedbytes             int64                              `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                             `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                              `json:"memorytotal"`
	Memoryused                       int64                              `json:"memoryused"`
	Memorywithoverprovisioning       string                             `json:"memorywithoverprovisioning"`
	Name                             string                             `json:"name"`
	Networkkbsread                   int64                              `json:"networkkbsread"`
	Networkkbswrite                  int64                              `json:"networkkbswrite"`
	Oscategoryid                     string                             `json:"oscategoryid"`
	Oscategoryname                   string                             `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse        `json:"outofbandmanagement"`
	Podid                            string                             `json:"podid"`
	Podname                          string                             `json:"podname"`
	Removed                          string                             `json:"removed"`
	Resourcestate                    string                             `json:"resourcestate"`
	State                            string                             `json:"state"`
	Suitableformigration             bool                               `json:"suitableformigration"`
	Type                             string                             `json:"type"`
	Ueficapability                   bool                               `json:"ueficapability"`
	Username                         string                             `json:"username"`
	Version                          string                             `json:"version"`
	Zoneid                           string                             `json:"zoneid"`
	Zonename                         string                             `json:"zonename"`
}

type AddBaremetalHostResponseGpugroup struct {
	Gpugroupname string                                 `json:"gpugroupname"`
	Vgpu         []AddBaremetalHostResponseGpugroupVgpu `json:"vgpu"`
}

type AddBaremetalHostResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type AddGloboDnsHostParams struct {
	P map[string]interface{}
}

func (P *AddGloboDnsHostParams) toURLValues() url.Values {
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

func (P *AddGloboDnsHostParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddGloboDnsHostParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddGloboDnsHostParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddGloboDnsHostParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddGloboDnsHostParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddGloboDnsHostParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddGloboDnsHostParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddGloboDnsHostParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddGloboDnsHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddGloboDnsHostParams(password string, physicalnetworkid string, url string, username string) *AddGloboDnsHostParams {
	P := &AddGloboDnsHostParams{}
	P.P = make(map[string]interface{})
	P.P["password"] = password
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["url"] = url
	P.P["username"] = username
	return P
}

// Adds the GloboDNS external host
func (s *HostService) AddGloboDnsHost(p *AddGloboDnsHostParams) (*AddGloboDnsHostResponse, error) {
	resp, err := s.cs.newRequest("addGloboDnsHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddGloboDnsHostResponse
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

type AddGloboDnsHostResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type AddHostParams struct {
	P map[string]interface{}
}

func (P *AddHostParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := P.P["hosttags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("hosttags", vv)
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *AddHostParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *AddHostParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *AddHostParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *AddHostParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *AddHostParams) SetClustername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clustername"] = v
}

func (P *AddHostParams) GetClustername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clustername"].(string)
	return value, ok
}

func (P *AddHostParams) SetHosttags(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hosttags"] = v
}

func (P *AddHostParams) GetHosttags() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hosttags"].([]string)
	return value, ok
}

func (P *AddHostParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *AddHostParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *AddHostParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddHostParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddHostParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *AddHostParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *AddHostParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddHostParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddHostParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddHostParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

func (P *AddHostParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *AddHostParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddHostParams(hypervisor string, podid string, url string, zoneid string) *AddHostParams {
	P := &AddHostParams{}
	P.P = make(map[string]interface{})
	P.P["hypervisor"] = hypervisor
	P.P["podid"] = podid
	P.P["url"] = url
	P.P["zoneid"] = zoneid
	return P
}

// Adds a new host.
func (s *HostService) AddHost(p *AddHostParams) (*AddHostResponse, error) {
	resp, err := s.cs.newRequest("addHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddHostResponse struct {
	Annotation                       string                      `json:"annotation"`
	Capabilities                     string                      `json:"capabilities"`
	Clusterid                        string                      `json:"clusterid"`
	Clustername                      string                      `json:"clustername"`
	Clustertype                      string                      `json:"clustertype"`
	Cpuallocated                     string                      `json:"cpuallocated"`
	Cpuallocatedpercentage           string                      `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                       `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                      `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                     `json:"cpuloadaverage"`
	Cpunumber                        int                         `json:"cpunumber"`
	Cpusockets                       int                         `json:"cpusockets"`
	Cpuspeed                         int64                       `json:"cpuspeed"`
	Cpuused                          string                      `json:"cpuused"`
	Cpuwithoverprovisioning          string                      `json:"cpuwithoverprovisioning"`
	Created                          string                      `json:"created"`
	Details                          map[string]string           `json:"details"`
	Disconnected                     string                      `json:"disconnected"`
	Disksizeallocated                int64                       `json:"disksizeallocated"`
	Disksizetotal                    int64                       `json:"disksizetotal"`
	Events                           string                      `json:"events"`
	Gpugroup                         []AddHostResponseGpugroup   `json:"gpugroup"`
	Hahost                           bool                        `json:"hahost"`
	Hasannotations                   bool                        `json:"hasannotations"`
	Hasenoughcapacity                bool                        `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse           `json:"hostha"`
	Hosttags                         string                      `json:"hosttags"`
	Hypervisor                       string                      `json:"hypervisor"`
	Hypervisorversion                string                      `json:"hypervisorversion"`
	Id                               string                      `json:"id"`
	Ipaddress                        string                      `json:"ipaddress"`
	Islocalstorageactive             bool                        `json:"islocalstorageactive"`
	JobID                            string                      `json:"jobid"`
	Jobstatus                        int                         `json:"jobstatus"`
	Lastannotated                    string                      `json:"lastannotated"`
	Lastpinged                       string                      `json:"lastpinged"`
	Managementserverid               UUID                        `json:"managementserverid"`
	Memoryallocated                  int64                       `json:"memoryallocated"`
	Memoryallocatedbytes             int64                       `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                      `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                       `json:"memorytotal"`
	Memoryused                       int64                       `json:"memoryused"`
	Memorywithoverprovisioning       string                      `json:"memorywithoverprovisioning"`
	Name                             string                      `json:"name"`
	Networkkbsread                   int64                       `json:"networkkbsread"`
	Networkkbswrite                  int64                       `json:"networkkbswrite"`
	Oscategoryid                     string                      `json:"oscategoryid"`
	Oscategoryname                   string                      `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse `json:"outofbandmanagement"`
	Podid                            string                      `json:"podid"`
	Podname                          string                      `json:"podname"`
	Removed                          string                      `json:"removed"`
	Resourcestate                    string                      `json:"resourcestate"`
	State                            string                      `json:"state"`
	Suitableformigration             bool                        `json:"suitableformigration"`
	Type                             string                      `json:"type"`
	Ueficapability                   bool                        `json:"ueficapability"`
	Username                         string                      `json:"username"`
	Version                          string                      `json:"version"`
	Zoneid                           string                      `json:"zoneid"`
	Zonename                         string                      `json:"zonename"`
}

type AddHostResponseGpugroup struct {
	Gpugroupname string                        `json:"gpugroupname"`
	Vgpu         []AddHostResponseGpugroupVgpu `json:"vgpu"`
}

type AddHostResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type AddSecondaryStorageParams struct {
	P map[string]interface{}
}

func (P *AddSecondaryStorageParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *AddSecondaryStorageParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddSecondaryStorageParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddSecondaryStorageParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *AddSecondaryStorageParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddSecondaryStorageParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddSecondaryStorageParams(url string) *AddSecondaryStorageParams {
	P := &AddSecondaryStorageParams{}
	P.P = make(map[string]interface{})
	P.P["url"] = url
	return P
}

// Adds secondary storage.
func (s *HostService) AddSecondaryStorage(p *AddSecondaryStorageParams) (*AddSecondaryStorageResponse, error) {
	resp, err := s.cs.newRequest("addSecondaryStorage", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddSecondaryStorageResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddSecondaryStorageResponse struct {
	Disksizetotal  int64  `json:"disksizetotal"`
	Disksizeused   int64  `json:"disksizeused"`
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	Providername   string `json:"providername"`
	Readonly       bool   `json:"readonly"`
	Scope          string `json:"scope"`
	Url            string `json:"url"`
	Zoneid         string `json:"zoneid"`
	Zonename       string `json:"zonename"`
}

type CancelHostMaintenanceParams struct {
	P map[string]interface{}
}

func (P *CancelHostMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *CancelHostMaintenanceParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *CancelHostMaintenanceParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new CancelHostMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewCancelHostMaintenanceParams(id string) *CancelHostMaintenanceParams {
	P := &CancelHostMaintenanceParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Cancels host maintenance.
func (s *HostService) CancelHostMaintenance(p *CancelHostMaintenanceParams) (*CancelHostMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("cancelHostMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CancelHostMaintenanceResponse
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

type CancelHostMaintenanceResponse struct {
	Annotation                       string                                  `json:"annotation"`
	Capabilities                     string                                  `json:"capabilities"`
	Clusterid                        string                                  `json:"clusterid"`
	Clustername                      string                                  `json:"clustername"`
	Clustertype                      string                                  `json:"clustertype"`
	Cpuallocated                     string                                  `json:"cpuallocated"`
	Cpuallocatedpercentage           string                                  `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                                   `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                                  `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                                 `json:"cpuloadaverage"`
	Cpunumber                        int                                     `json:"cpunumber"`
	Cpusockets                       int                                     `json:"cpusockets"`
	Cpuspeed                         int64                                   `json:"cpuspeed"`
	Cpuused                          string                                  `json:"cpuused"`
	Cpuwithoverprovisioning          string                                  `json:"cpuwithoverprovisioning"`
	Created                          string                                  `json:"created"`
	Details                          map[string]string                       `json:"details"`
	Disconnected                     string                                  `json:"disconnected"`
	Disksizeallocated                int64                                   `json:"disksizeallocated"`
	Disksizetotal                    int64                                   `json:"disksizetotal"`
	Events                           string                                  `json:"events"`
	Gpugroup                         []CancelHostMaintenanceResponseGpugroup `json:"gpugroup"`
	Hahost                           bool                                    `json:"hahost"`
	Hasannotations                   bool                                    `json:"hasannotations"`
	Hasenoughcapacity                bool                                    `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse                       `json:"hostha"`
	Hosttags                         string                                  `json:"hosttags"`
	Hypervisor                       string                                  `json:"hypervisor"`
	Hypervisorversion                string                                  `json:"hypervisorversion"`
	Id                               string                                  `json:"id"`
	Ipaddress                        string                                  `json:"ipaddress"`
	Islocalstorageactive             bool                                    `json:"islocalstorageactive"`
	JobID                            string                                  `json:"jobid"`
	Jobstatus                        int                                     `json:"jobstatus"`
	Lastannotated                    string                                  `json:"lastannotated"`
	Lastpinged                       string                                  `json:"lastpinged"`
	Managementserverid               UUID                                    `json:"managementserverid"`
	Memoryallocated                  int64                                   `json:"memoryallocated"`
	Memoryallocatedbytes             int64                                   `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                                  `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                                   `json:"memorytotal"`
	Memoryused                       int64                                   `json:"memoryused"`
	Memorywithoverprovisioning       string                                  `json:"memorywithoverprovisioning"`
	Name                             string                                  `json:"name"`
	Networkkbsread                   int64                                   `json:"networkkbsread"`
	Networkkbswrite                  int64                                   `json:"networkkbswrite"`
	Oscategoryid                     string                                  `json:"oscategoryid"`
	Oscategoryname                   string                                  `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse             `json:"outofbandmanagement"`
	Podid                            string                                  `json:"podid"`
	Podname                          string                                  `json:"podname"`
	Removed                          string                                  `json:"removed"`
	Resourcestate                    string                                  `json:"resourcestate"`
	State                            string                                  `json:"state"`
	Suitableformigration             bool                                    `json:"suitableformigration"`
	Type                             string                                  `json:"type"`
	Ueficapability                   bool                                    `json:"ueficapability"`
	Username                         string                                  `json:"username"`
	Version                          string                                  `json:"version"`
	Zoneid                           string                                  `json:"zoneid"`
	Zonename                         string                                  `json:"zonename"`
}

type CancelHostMaintenanceResponseGpugroup struct {
	Gpugroupname string                                      `json:"gpugroupname"`
	Vgpu         []CancelHostMaintenanceResponseGpugroupVgpu `json:"vgpu"`
}

type CancelHostMaintenanceResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type ConfigureHAForHostParams struct {
	P map[string]interface{}
}

func (P *ConfigureHAForHostParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["provider"]; found {
		u.Set("provider", v.(string))
	}
	return u
}

func (P *ConfigureHAForHostParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ConfigureHAForHostParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *ConfigureHAForHostParams) SetProvider(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["provider"] = v
}

func (P *ConfigureHAForHostParams) GetProvider() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["provider"].(string)
	return value, ok
}

// You should always use this function to get a new ConfigureHAForHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewConfigureHAForHostParams(hostid string, provider string) *ConfigureHAForHostParams {
	P := &ConfigureHAForHostParams{}
	P.P = make(map[string]interface{})
	P.P["hostid"] = hostid
	P.P["provider"] = provider
	return P
}

// Configures HA for a host
func (s *HostService) ConfigureHAForHost(p *ConfigureHAForHostParams) (*HAForHostResponse, error) {
	resp, err := s.cs.newRequest("configureHAForHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r HAForHostResponse
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

type HAForHostResponse struct {
	Haenable   bool   `json:"haenable"`
	Haprovider string `json:"haprovider"`
	Hastate    string `json:"hastate"`
	Hostid     string `json:"hostid"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Status     bool   `json:"status"`
}

type EnableHAForHostParams struct {
	P map[string]interface{}
}

func (P *EnableHAForHostParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (P *EnableHAForHostParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *EnableHAForHostParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableHAForHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewEnableHAForHostParams(hostid string) *EnableHAForHostParams {
	P := &EnableHAForHostParams{}
	P.P = make(map[string]interface{})
	P.P["hostid"] = hostid
	return P
}

// Enables HA for a host
func (s *HostService) EnableHAForHost(p *EnableHAForHostParams) (*EnableHAForHostResponse, error) {
	resp, err := s.cs.newRequest("enableHAForHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableHAForHostResponse
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

type EnableHAForHostResponse struct {
	Haenable   bool   `json:"haenable"`
	Haprovider string `json:"haprovider"`
	Hastate    string `json:"hastate"`
	Hostid     string `json:"hostid"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Status     bool   `json:"status"`
}

type DedicateHostParams struct {
	P map[string]interface{}
}

func (P *DedicateHostParams) toURLValues() url.Values {
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
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (P *DedicateHostParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DedicateHostParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DedicateHostParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DedicateHostParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *DedicateHostParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *DedicateHostParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicateHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewDedicateHostParams(domainid string, hostid string) *DedicateHostParams {
	P := &DedicateHostParams{}
	P.P = make(map[string]interface{})
	P.P["domainid"] = domainid
	P.P["hostid"] = hostid
	return P
}

// Dedicates a host.
func (s *HostService) DedicateHost(p *DedicateHostParams) (*DedicateHostResponse, error) {
	resp, err := s.cs.newRequest("dedicateHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateHostResponse
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

type DedicateHostResponse struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Hostid          string `json:"hostid"`
	Hostname        string `json:"hostname"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
}

type DeleteHostParams struct {
	P map[string]interface{}
}

func (P *DeleteHostParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := P.P["forcedestroylocalstorage"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forcedestroylocalstorage", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteHostParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *DeleteHostParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *DeleteHostParams) SetForcedestroylocalstorage(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forcedestroylocalstorage"] = v
}

func (P *DeleteHostParams) GetForcedestroylocalstorage() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forcedestroylocalstorage"].(bool)
	return value, ok
}

func (P *DeleteHostParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteHostParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewDeleteHostParams(id string) *DeleteHostParams {
	P := &DeleteHostParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a host.
func (s *HostService) DeleteHost(p *DeleteHostParams) (*DeleteHostResponse, error) {
	resp, err := s.cs.newRequest("deleteHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteHostResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteHostResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteHostResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DisableOutOfBandManagementForHostParams struct {
	P map[string]interface{}
}

func (P *DisableOutOfBandManagementForHostParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (P *DisableOutOfBandManagementForHostParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *DisableOutOfBandManagementForHostParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableOutOfBandManagementForHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewDisableOutOfBandManagementForHostParams(hostid string) *DisableOutOfBandManagementForHostParams {
	P := &DisableOutOfBandManagementForHostParams{}
	P.P = make(map[string]interface{})
	P.P["hostid"] = hostid
	return P
}

// Disables out-of-band management for a host
func (s *HostService) DisableOutOfBandManagementForHost(p *DisableOutOfBandManagementForHostParams) (*DisableOutOfBandManagementForHostResponse, error) {
	resp, err := s.cs.newRequest("disableOutOfBandManagementForHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableOutOfBandManagementForHostResponse
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

type DisableOutOfBandManagementForHostResponse struct {
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

type EnableOutOfBandManagementForHostParams struct {
	P map[string]interface{}
}

func (P *EnableOutOfBandManagementForHostParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (P *EnableOutOfBandManagementForHostParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *EnableOutOfBandManagementForHostParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableOutOfBandManagementForHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewEnableOutOfBandManagementForHostParams(hostid string) *EnableOutOfBandManagementForHostParams {
	P := &EnableOutOfBandManagementForHostParams{}
	P.P = make(map[string]interface{})
	P.P["hostid"] = hostid
	return P
}

// Enables out-of-band management for a host
func (s *HostService) EnableOutOfBandManagementForHost(p *EnableOutOfBandManagementForHostParams) (*EnableOutOfBandManagementForHostResponse, error) {
	resp, err := s.cs.newRequest("enableOutOfBandManagementForHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableOutOfBandManagementForHostResponse
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

type EnableOutOfBandManagementForHostResponse struct {
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

type FindHostsForMigrationParams struct {
	P map[string]interface{}
}

func (P *FindHostsForMigrationParams) toURLValues() url.Values {
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
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *FindHostsForMigrationParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *FindHostsForMigrationParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *FindHostsForMigrationParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *FindHostsForMigrationParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *FindHostsForMigrationParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *FindHostsForMigrationParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *FindHostsForMigrationParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *FindHostsForMigrationParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new FindHostsForMigrationParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewFindHostsForMigrationParams(virtualmachineid string) *FindHostsForMigrationParams {
	P := &FindHostsForMigrationParams{}
	P.P = make(map[string]interface{})
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Find hosts suitable for migrating a virtual machine.
func (s *HostService) FindHostsForMigration(p *FindHostsForMigrationParams) (*FindHostsForMigrationResponse, error) {
	resp, err := s.cs.newRequest("findHostsForMigration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r FindHostsForMigrationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type FindHostsForMigrationResponse struct {
	Averageload                      int64  `json:"averageload"`
	Capabilities                     string `json:"capabilities"`
	Clusterid                        string `json:"clusterid"`
	Clustername                      string `json:"clustername"`
	Clustertype                      string `json:"clustertype"`
	Cpuallocated                     string `json:"cpuallocated"`
	Cpuallocatedpercentage           string `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64  `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string `json:"cpuallocatedwithoverprovisioning"`
	Cpunumber                        int    `json:"cpunumber"`
	Cpuspeed                         int64  `json:"cpuspeed"`
	Cpuused                          string `json:"cpuused"`
	Cpuwithoverprovisioning          string `json:"cpuwithoverprovisioning"`
	Created                          string `json:"created"`
	Disconnected                     string `json:"disconnected"`
	Disksizeallocated                int64  `json:"disksizeallocated"`
	Disksizetotal                    int64  `json:"disksizetotal"`
	Events                           string `json:"events"`
	Hahost                           bool   `json:"hahost"`
	Hasenoughcapacity                bool   `json:"hasenoughcapacity"`
	Hosttags                         string `json:"hosttags"`
	Hypervisor                       string `json:"hypervisor"`
	Hypervisorversion                string `json:"hypervisorversion"`
	Id                               string `json:"id"`
	Ipaddress                        string `json:"ipaddress"`
	Islocalstorageactive             bool   `json:"islocalstorageactive"`
	JobID                            string `json:"jobid"`
	Jobstatus                        int    `json:"jobstatus"`
	Lastpinged                       string `json:"lastpinged"`
	Managementserverid               UUID   `json:"managementserverid"`
	Memoryallocated                  string `json:"memoryallocated"`
	Memoryallocatedbytes             int64  `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string `json:"memoryallocatedpercentage"`
	Memorytotal                      int64  `json:"memorytotal"`
	Memoryused                       int64  `json:"memoryused"`
	Memorywithoverprovisioning       string `json:"memorywithoverprovisioning"`
	Name                             string `json:"name"`
	Networkkbsread                   int64  `json:"networkkbsread"`
	Networkkbswrite                  int64  `json:"networkkbswrite"`
	Oscategoryid                     string `json:"oscategoryid"`
	Oscategoryname                   string `json:"oscategoryname"`
	Podid                            string `json:"podid"`
	Podname                          string `json:"podname"`
	Removed                          string `json:"removed"`
	RequiresStorageMotion            bool   `json:"requiresStorageMotion"`
	Resourcestate                    string `json:"resourcestate"`
	State                            string `json:"state"`
	Suitableformigration             bool   `json:"suitableformigration"`
	Type                             string `json:"type"`
	Version                          string `json:"version"`
	Zoneid                           string `json:"zoneid"`
	Zonename                         string `json:"zonename"`
}

type ListDedicatedHostsParams struct {
	P map[string]interface{}
}

func (P *ListDedicatedHostsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["affinitygroupid"]; found {
		u.Set("affinitygroupid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
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

func (P *ListDedicatedHostsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListDedicatedHostsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListDedicatedHostsParams) SetAffinitygroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["affinitygroupid"] = v
}

func (P *ListDedicatedHostsParams) GetAffinitygroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["affinitygroupid"].(string)
	return value, ok
}

func (P *ListDedicatedHostsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListDedicatedHostsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListDedicatedHostsParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ListDedicatedHostsParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *ListDedicatedHostsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListDedicatedHostsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListDedicatedHostsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListDedicatedHostsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListDedicatedHostsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListDedicatedHostsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListDedicatedHostsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListDedicatedHostsParams() *ListDedicatedHostsParams {
	P := &ListDedicatedHostsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists dedicated hosts.
func (s *HostService) ListDedicatedHosts(p *ListDedicatedHostsParams) (*ListDedicatedHostsResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedHosts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedHostsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedHostsResponse struct {
	Count          int              `json:"count"`
	DedicatedHosts []*DedicatedHost `json:"dedicatedhost"`
}

type DedicatedHost struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Domainid        string `json:"domainid"`
	Hostid          string `json:"hostid"`
	Hostname        string `json:"hostname"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
}

type ListHostTagsParams struct {
	P map[string]interface{}
}

func (P *ListHostTagsParams) toURLValues() url.Values {
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

func (P *ListHostTagsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListHostTagsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListHostTagsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListHostTagsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListHostTagsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListHostTagsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListHostTagsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListHostTagsParams() *ListHostTagsParams {
	P := &ListHostTagsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostTagID(keyword string, opts ...OptionFunc) (string, int, error) {
	P := &ListHostTagsParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListHostTags(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.HostTags[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.HostTags {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Lists host tags
func (s *HostService) ListHostTags(p *ListHostTagsParams) (*ListHostTagsResponse, error) {
	resp, err := s.cs.newRequest("listHostTags", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHostTagsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListHostTagsResponse struct {
	Count    int        `json:"count"`
	HostTags []*HostTag `json:"hosttag"`
}

type HostTag struct {
	Hostid    int64  `json:"hostid"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type ListHostsParams struct {
	P map[string]interface{}
}

func (P *ListHostsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := P.P["hahost"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("hahost", vv)
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
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
	if v, found := P.P["outofbandmanagementenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("outofbandmanagementenabled", vv)
	}
	if v, found := P.P["outofbandmanagementpowerstate"]; found {
		u.Set("outofbandmanagementpowerstate", v.(string))
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
	if v, found := P.P["resourcestate"]; found {
		u.Set("resourcestate", v.(string))
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListHostsParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *ListHostsParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *ListHostsParams) SetDetails(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ListHostsParams) GetDetails() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].([]string)
	return value, ok
}

func (P *ListHostsParams) SetHahost(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hahost"] = v
}

func (P *ListHostsParams) GetHahost() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hahost"].(bool)
	return value, ok
}

func (P *ListHostsParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *ListHostsParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *ListHostsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListHostsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListHostsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListHostsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListHostsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListHostsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListHostsParams) SetOutofbandmanagementenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["outofbandmanagementenabled"] = v
}

func (P *ListHostsParams) GetOutofbandmanagementenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["outofbandmanagementenabled"].(bool)
	return value, ok
}

func (P *ListHostsParams) SetOutofbandmanagementpowerstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["outofbandmanagementpowerstate"] = v
}

func (P *ListHostsParams) GetOutofbandmanagementpowerstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["outofbandmanagementpowerstate"].(string)
	return value, ok
}

func (P *ListHostsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListHostsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListHostsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListHostsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListHostsParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListHostsParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListHostsParams) SetResourcestate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcestate"] = v
}

func (P *ListHostsParams) GetResourcestate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcestate"].(string)
	return value, ok
}

func (P *ListHostsParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListHostsParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListHostsParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ListHostsParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

func (P *ListHostsParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *ListHostsParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *ListHostsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListHostsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListHostsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListHostsParams() *ListHostsParams {
	P := &ListHostsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListHostsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListHosts(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Hosts[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Hosts {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostByName(name string, opts ...OptionFunc) (*Host, int, error) {
	id, count, err := s.GetHostID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetHostByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostByID(id string, opts ...OptionFunc) (*Host, int, error) {
	P := &ListHostsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListHosts(P)
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
		return l.Hosts[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Host UUID: %s!", id)
}

// Lists hosts.
func (s *HostService) ListHosts(p *ListHostsParams) (*ListHostsResponse, error) {
	resp, err := s.cs.newRequest("listHosts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHostsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListHostsResponse struct {
	Count int     `json:"count"`
	Hosts []*Host `json:"host"`
}

type Host struct {
	Annotation                       string                      `json:"annotation"`
	Capabilities                     string                      `json:"capabilities"`
	Clusterid                        string                      `json:"clusterid"`
	Clustername                      string                      `json:"clustername"`
	Clustertype                      string                      `json:"clustertype"`
	Cpuallocated                     string                      `json:"cpuallocated"`
	Cpuallocatedpercentage           string                      `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                       `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                      `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                     `json:"cpuloadaverage"`
	Cpunumber                        int                         `json:"cpunumber"`
	Cpusockets                       int                         `json:"cpusockets"`
	Cpuspeed                         int64                       `json:"cpuspeed"`
	Cpuused                          string                      `json:"cpuused"`
	Cpuwithoverprovisioning          string                      `json:"cpuwithoverprovisioning"`
	Created                          string                      `json:"created"`
	Details                          map[string]string           `json:"details"`
	Disconnected                     string                      `json:"disconnected"`
	Disksizeallocated                int64                       `json:"disksizeallocated"`
	Disksizetotal                    int64                       `json:"disksizetotal"`
	Events                           string                      `json:"events"`
	Gpugroup                         []HostGpugroup              `json:"gpugroup"`
	Hahost                           bool                        `json:"hahost"`
	Hasannotations                   bool                        `json:"hasannotations"`
	Hasenoughcapacity                bool                        `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse           `json:"hostha"`
	Hosttags                         string                      `json:"hosttags"`
	Hypervisor                       string                      `json:"hypervisor"`
	Hypervisorversion                string                      `json:"hypervisorversion"`
	Id                               string                      `json:"id"`
	Ipaddress                        string                      `json:"ipaddress"`
	Islocalstorageactive             bool                        `json:"islocalstorageactive"`
	JobID                            string                      `json:"jobid"`
	Jobstatus                        int                         `json:"jobstatus"`
	Lastannotated                    string                      `json:"lastannotated"`
	Lastpinged                       string                      `json:"lastpinged"`
	Managementserverid               UUID                        `json:"managementserverid"`
	Memoryallocated                  int64                       `json:"memoryallocated"`
	Memoryallocatedbytes             int64                       `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                      `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                       `json:"memorytotal"`
	Memoryused                       int64                       `json:"memoryused"`
	Memorywithoverprovisioning       string                      `json:"memorywithoverprovisioning"`
	Name                             string                      `json:"name"`
	Networkkbsread                   int64                       `json:"networkkbsread"`
	Networkkbswrite                  int64                       `json:"networkkbswrite"`
	Oscategoryid                     string                      `json:"oscategoryid"`
	Oscategoryname                   string                      `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse `json:"outofbandmanagement"`
	Podid                            string                      `json:"podid"`
	Podname                          string                      `json:"podname"`
	Removed                          string                      `json:"removed"`
	Resourcestate                    string                      `json:"resourcestate"`
	State                            string                      `json:"state"`
	Suitableformigration             bool                        `json:"suitableformigration"`
	Type                             string                      `json:"type"`
	Ueficapability                   bool                        `json:"ueficapability"`
	Username                         string                      `json:"username"`
	Version                          string                      `json:"version"`
	Zoneid                           string                      `json:"zoneid"`
	Zonename                         string                      `json:"zonename"`
}

type HostGpugroup struct {
	Gpugroupname string             `json:"gpugroupname"`
	Vgpu         []HostGpugroupVgpu `json:"vgpu"`
}

type HostGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type ListHostsMetricsParams struct {
	P map[string]interface{}
}

func (P *ListHostsMetricsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := P.P["hahost"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("hahost", vv)
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
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
	if v, found := P.P["outofbandmanagementenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("outofbandmanagementenabled", vv)
	}
	if v, found := P.P["outofbandmanagementpowerstate"]; found {
		u.Set("outofbandmanagementpowerstate", v.(string))
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
	if v, found := P.P["resourcestate"]; found {
		u.Set("resourcestate", v.(string))
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListHostsMetricsParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *ListHostsMetricsParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetDetails(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ListHostsMetricsParams) GetDetails() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].([]string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetHahost(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hahost"] = v
}

func (P *ListHostsMetricsParams) GetHahost() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hahost"].(bool)
	return value, ok
}

func (P *ListHostsMetricsParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *ListHostsMetricsParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListHostsMetricsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListHostsMetricsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListHostsMetricsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetOutofbandmanagementenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["outofbandmanagementenabled"] = v
}

func (P *ListHostsMetricsParams) GetOutofbandmanagementenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["outofbandmanagementenabled"].(bool)
	return value, ok
}

func (P *ListHostsMetricsParams) SetOutofbandmanagementpowerstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["outofbandmanagementpowerstate"] = v
}

func (P *ListHostsMetricsParams) GetOutofbandmanagementpowerstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["outofbandmanagementpowerstate"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListHostsMetricsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListHostsMetricsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListHostsMetricsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListHostsMetricsParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListHostsMetricsParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetResourcestate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["resourcestate"] = v
}

func (P *ListHostsMetricsParams) GetResourcestate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["resourcestate"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListHostsMetricsParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ListHostsMetricsParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *ListHostsMetricsParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *ListHostsMetricsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListHostsMetricsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListHostsMetricsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListHostsMetricsParams() *ListHostsMetricsParams {
	P := &ListHostsMetricsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostsMetricID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListHostsMetricsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListHostsMetrics(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.HostsMetrics[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.HostsMetrics {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostsMetricByName(name string, opts ...OptionFunc) (*HostsMetric, int, error) {
	id, count, err := s.GetHostsMetricID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetHostsMetricByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *HostService) GetHostsMetricByID(id string, opts ...OptionFunc) (*HostsMetric, int, error) {
	P := &ListHostsMetricsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListHostsMetrics(P)
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
		return l.HostsMetrics[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for HostsMetric UUID: %s!", id)
}

// Lists hosts metrics
func (s *HostService) ListHostsMetrics(p *ListHostsMetricsParams) (*ListHostsMetricsResponse, error) {
	resp, err := s.cs.newRequest("listHostsMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListHostsMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListHostsMetricsResponse struct {
	Count        int            `json:"count"`
	HostsMetrics []*HostsMetric `json:"hostsmetric"`
}

type HostsMetric struct {
	Annotation                       string                      `json:"annotation"`
	Capabilities                     string                      `json:"capabilities"`
	Clusterid                        string                      `json:"clusterid"`
	Clustername                      string                      `json:"clustername"`
	Clustertype                      string                      `json:"clustertype"`
	Cpuallocated                     string                      `json:"cpuallocated"`
	Cpuallocateddisablethreshold     bool                        `json:"cpuallocateddisablethreshold"`
	Cpuallocatedghz                  string                      `json:"cpuallocatedghz"`
	Cpuallocatedpercentage           string                      `json:"cpuallocatedpercentage"`
	Cpuallocatedthreshold            bool                        `json:"cpuallocatedthreshold"`
	Cpuallocatedvalue                int64                       `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                      `json:"cpuallocatedwithoverprovisioning"`
	Cpudisablethreshold              bool                        `json:"cpudisablethreshold"`
	Cpuloadaverage                   float64                     `json:"cpuloadaverage"`
	Cpunumber                        int                         `json:"cpunumber"`
	Cpusockets                       int                         `json:"cpusockets"`
	Cpuspeed                         int64                       `json:"cpuspeed"`
	Cputhreshold                     bool                        `json:"cputhreshold"`
	Cputotalghz                      string                      `json:"cputotalghz"`
	Cpuused                          string                      `json:"cpuused"`
	Cpuusedghz                       string                      `json:"cpuusedghz"`
	Cpuwithoverprovisioning          string                      `json:"cpuwithoverprovisioning"`
	Created                          string                      `json:"created"`
	Details                          map[string]string           `json:"details"`
	Disconnected                     string                      `json:"disconnected"`
	Disksizeallocated                int64                       `json:"disksizeallocated"`
	Disksizetotal                    int64                       `json:"disksizetotal"`
	Events                           string                      `json:"events"`
	Gpugroup                         []HostsMetricGpugroup       `json:"gpugroup"`
	Hahost                           bool                        `json:"hahost"`
	Hasannotations                   bool                        `json:"hasannotations"`
	Hasenoughcapacity                bool                        `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse           `json:"hostha"`
	Hosttags                         string                      `json:"hosttags"`
	Hypervisor                       string                      `json:"hypervisor"`
	Hypervisorversion                string                      `json:"hypervisorversion"`
	Id                               string                      `json:"id"`
	Instances                        string                      `json:"instances"`
	Ipaddress                        string                      `json:"ipaddress"`
	Islocalstorageactive             bool                        `json:"islocalstorageactive"`
	JobID                            string                      `json:"jobid"`
	Jobstatus                        int                         `json:"jobstatus"`
	Lastannotated                    string                      `json:"lastannotated"`
	Lastpinged                       string                      `json:"lastpinged"`
	Managementserverid               UUID                        `json:"managementserverid"`
	Memoryallocated                  int64                       `json:"memoryallocated"`
	Memoryallocatedbytes             int64                       `json:"memoryallocatedbytes"`
	Memoryallocateddisablethreshold  bool                        `json:"memoryallocateddisablethreshold"`
	Memoryallocatedgb                string                      `json:"memoryallocatedgb"`
	Memoryallocatedpercentage        string                      `json:"memoryallocatedpercentage"`
	Memoryallocatedthreshold         bool                        `json:"memoryallocatedthreshold"`
	Memorydisablethreshold           bool                        `json:"memorydisablethreshold"`
	Memorythreshold                  bool                        `json:"memorythreshold"`
	Memorytotal                      int64                       `json:"memorytotal"`
	Memorytotalgb                    string                      `json:"memorytotalgb"`
	Memoryused                       int64                       `json:"memoryused"`
	Memoryusedgb                     string                      `json:"memoryusedgb"`
	Memorywithoverprovisioning       string                      `json:"memorywithoverprovisioning"`
	Name                             string                      `json:"name"`
	Networkkbsread                   int64                       `json:"networkkbsread"`
	Networkkbswrite                  int64                       `json:"networkkbswrite"`
	Networkread                      string                      `json:"networkread"`
	Networkwrite                     string                      `json:"networkwrite"`
	Oscategoryid                     string                      `json:"oscategoryid"`
	Oscategoryname                   string                      `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse `json:"outofbandmanagement"`
	Podid                            string                      `json:"podid"`
	Podname                          string                      `json:"podname"`
	Powerstate                       string                      `json:"powerstate"`
	Removed                          string                      `json:"removed"`
	Resourcestate                    string                      `json:"resourcestate"`
	State                            string                      `json:"state"`
	Suitableformigration             bool                        `json:"suitableformigration"`
	Type                             string                      `json:"type"`
	Ueficapability                   bool                        `json:"ueficapability"`
	Username                         string                      `json:"username"`
	Version                          string                      `json:"version"`
	Zoneid                           string                      `json:"zoneid"`
	Zonename                         string                      `json:"zonename"`
}

type HostsMetricGpugroup struct {
	Gpugroupname string                    `json:"gpugroupname"`
	Vgpu         []HostsMetricGpugroupVgpu `json:"vgpu"`
}

type HostsMetricGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type PrepareHostForMaintenanceParams struct {
	P map[string]interface{}
}

func (P *PrepareHostForMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *PrepareHostForMaintenanceParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *PrepareHostForMaintenanceParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new PrepareHostForMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewPrepareHostForMaintenanceParams(id string) *PrepareHostForMaintenanceParams {
	P := &PrepareHostForMaintenanceParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Prepares a host for maintenance.
func (s *HostService) PrepareHostForMaintenance(p *PrepareHostForMaintenanceParams) (*PrepareHostForMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("prepareHostForMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PrepareHostForMaintenanceResponse
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

type PrepareHostForMaintenanceResponse struct {
	Annotation                       string                                      `json:"annotation"`
	Capabilities                     string                                      `json:"capabilities"`
	Clusterid                        string                                      `json:"clusterid"`
	Clustername                      string                                      `json:"clustername"`
	Clustertype                      string                                      `json:"clustertype"`
	Cpuallocated                     string                                      `json:"cpuallocated"`
	Cpuallocatedpercentage           string                                      `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                                       `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                                      `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                                     `json:"cpuloadaverage"`
	Cpunumber                        int                                         `json:"cpunumber"`
	Cpusockets                       int                                         `json:"cpusockets"`
	Cpuspeed                         int64                                       `json:"cpuspeed"`
	Cpuused                          string                                      `json:"cpuused"`
	Cpuwithoverprovisioning          string                                      `json:"cpuwithoverprovisioning"`
	Created                          string                                      `json:"created"`
	Details                          map[string]string                           `json:"details"`
	Disconnected                     string                                      `json:"disconnected"`
	Disksizeallocated                int64                                       `json:"disksizeallocated"`
	Disksizetotal                    int64                                       `json:"disksizetotal"`
	Events                           string                                      `json:"events"`
	Gpugroup                         []PrepareHostForMaintenanceResponseGpugroup `json:"gpugroup"`
	Hahost                           bool                                        `json:"hahost"`
	Hasannotations                   bool                                        `json:"hasannotations"`
	Hasenoughcapacity                bool                                        `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse                           `json:"hostha"`
	Hosttags                         string                                      `json:"hosttags"`
	Hypervisor                       string                                      `json:"hypervisor"`
	Hypervisorversion                string                                      `json:"hypervisorversion"`
	Id                               string                                      `json:"id"`
	Ipaddress                        string                                      `json:"ipaddress"`
	Islocalstorageactive             bool                                        `json:"islocalstorageactive"`
	JobID                            string                                      `json:"jobid"`
	Jobstatus                        int                                         `json:"jobstatus"`
	Lastannotated                    string                                      `json:"lastannotated"`
	Lastpinged                       string                                      `json:"lastpinged"`
	Managementserverid               UUID                                        `json:"managementserverid"`
	Memoryallocated                  int64                                       `json:"memoryallocated"`
	Memoryallocatedbytes             int64                                       `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                                      `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                                       `json:"memorytotal"`
	Memoryused                       int64                                       `json:"memoryused"`
	Memorywithoverprovisioning       string                                      `json:"memorywithoverprovisioning"`
	Name                             string                                      `json:"name"`
	Networkkbsread                   int64                                       `json:"networkkbsread"`
	Networkkbswrite                  int64                                       `json:"networkkbswrite"`
	Oscategoryid                     string                                      `json:"oscategoryid"`
	Oscategoryname                   string                                      `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse                 `json:"outofbandmanagement"`
	Podid                            string                                      `json:"podid"`
	Podname                          string                                      `json:"podname"`
	Removed                          string                                      `json:"removed"`
	Resourcestate                    string                                      `json:"resourcestate"`
	State                            string                                      `json:"state"`
	Suitableformigration             bool                                        `json:"suitableformigration"`
	Type                             string                                      `json:"type"`
	Ueficapability                   bool                                        `json:"ueficapability"`
	Username                         string                                      `json:"username"`
	Version                          string                                      `json:"version"`
	Zoneid                           string                                      `json:"zoneid"`
	Zonename                         string                                      `json:"zonename"`
}

type PrepareHostForMaintenanceResponseGpugroup struct {
	Gpugroupname string                                          `json:"gpugroupname"`
	Vgpu         []PrepareHostForMaintenanceResponseGpugroupVgpu `json:"vgpu"`
}

type PrepareHostForMaintenanceResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type ReconnectHostParams struct {
	P map[string]interface{}
}

func (P *ReconnectHostParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ReconnectHostParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ReconnectHostParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ReconnectHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewReconnectHostParams(id string) *ReconnectHostParams {
	P := &ReconnectHostParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Reconnects a host.
func (s *HostService) ReconnectHost(p *ReconnectHostParams) (*ReconnectHostResponse, error) {
	resp, err := s.cs.newRequest("reconnectHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReconnectHostResponse
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

type ReconnectHostResponse struct {
	Annotation                       string                          `json:"annotation"`
	Capabilities                     string                          `json:"capabilities"`
	Clusterid                        string                          `json:"clusterid"`
	Clustername                      string                          `json:"clustername"`
	Clustertype                      string                          `json:"clustertype"`
	Cpuallocated                     string                          `json:"cpuallocated"`
	Cpuallocatedpercentage           string                          `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                           `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                          `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                         `json:"cpuloadaverage"`
	Cpunumber                        int                             `json:"cpunumber"`
	Cpusockets                       int                             `json:"cpusockets"`
	Cpuspeed                         int64                           `json:"cpuspeed"`
	Cpuused                          string                          `json:"cpuused"`
	Cpuwithoverprovisioning          string                          `json:"cpuwithoverprovisioning"`
	Created                          string                          `json:"created"`
	Details                          map[string]string               `json:"details"`
	Disconnected                     string                          `json:"disconnected"`
	Disksizeallocated                int64                           `json:"disksizeallocated"`
	Disksizetotal                    int64                           `json:"disksizetotal"`
	Events                           string                          `json:"events"`
	Gpugroup                         []ReconnectHostResponseGpugroup `json:"gpugroup"`
	Hahost                           bool                            `json:"hahost"`
	Hasannotations                   bool                            `json:"hasannotations"`
	Hasenoughcapacity                bool                            `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse               `json:"hostha"`
	Hosttags                         string                          `json:"hosttags"`
	Hypervisor                       string                          `json:"hypervisor"`
	Hypervisorversion                string                          `json:"hypervisorversion"`
	Id                               string                          `json:"id"`
	Ipaddress                        string                          `json:"ipaddress"`
	Islocalstorageactive             bool                            `json:"islocalstorageactive"`
	JobID                            string                          `json:"jobid"`
	Jobstatus                        int                             `json:"jobstatus"`
	Lastannotated                    string                          `json:"lastannotated"`
	Lastpinged                       string                          `json:"lastpinged"`
	Managementserverid               UUID                            `json:"managementserverid"`
	Memoryallocated                  int64                           `json:"memoryallocated"`
	Memoryallocatedbytes             int64                           `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                          `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                           `json:"memorytotal"`
	Memoryused                       int64                           `json:"memoryused"`
	Memorywithoverprovisioning       string                          `json:"memorywithoverprovisioning"`
	Name                             string                          `json:"name"`
	Networkkbsread                   int64                           `json:"networkkbsread"`
	Networkkbswrite                  int64                           `json:"networkkbswrite"`
	Oscategoryid                     string                          `json:"oscategoryid"`
	Oscategoryname                   string                          `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse     `json:"outofbandmanagement"`
	Podid                            string                          `json:"podid"`
	Podname                          string                          `json:"podname"`
	Removed                          string                          `json:"removed"`
	Resourcestate                    string                          `json:"resourcestate"`
	State                            string                          `json:"state"`
	Suitableformigration             bool                            `json:"suitableformigration"`
	Type                             string                          `json:"type"`
	Ueficapability                   bool                            `json:"ueficapability"`
	Username                         string                          `json:"username"`
	Version                          string                          `json:"version"`
	Zoneid                           string                          `json:"zoneid"`
	Zonename                         string                          `json:"zonename"`
}

type ReconnectHostResponseGpugroup struct {
	Gpugroupname string                              `json:"gpugroupname"`
	Vgpu         []ReconnectHostResponseGpugroupVgpu `json:"vgpu"`
}

type ReconnectHostResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type ReleaseDedicatedHostParams struct {
	P map[string]interface{}
}

func (P *ReleaseDedicatedHostParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	return u
}

func (P *ReleaseDedicatedHostParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ReleaseDedicatedHostParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseDedicatedHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewReleaseDedicatedHostParams(hostid string) *ReleaseDedicatedHostParams {
	P := &ReleaseDedicatedHostParams{}
	P.P = make(map[string]interface{})
	P.P["hostid"] = hostid
	return P
}

// Release the dedication for host
func (s *HostService) ReleaseDedicatedHost(p *ReleaseDedicatedHostParams) (*ReleaseDedicatedHostResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedHostResponse
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

type ReleaseDedicatedHostResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ReleaseHostReservationParams struct {
	P map[string]interface{}
}

func (P *ReleaseHostReservationParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ReleaseHostReservationParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ReleaseHostReservationParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseHostReservationParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewReleaseHostReservationParams(id string) *ReleaseHostReservationParams {
	P := &ReleaseHostReservationParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Releases host reservation.
func (s *HostService) ReleaseHostReservation(p *ReleaseHostReservationParams) (*ReleaseHostReservationResponse, error) {
	resp, err := s.cs.newRequest("releaseHostReservation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseHostReservationResponse
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

type ReleaseHostReservationResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateHostParams struct {
	P map[string]interface{}
}

func (P *UpdateHostParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["annotation"]; found {
		u.Set("annotation", v.(string))
	}
	if v, found := P.P["hosttags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("hosttags", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["oscategoryid"]; found {
		u.Set("oscategoryid", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	return u
}

func (P *UpdateHostParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *UpdateHostParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *UpdateHostParams) SetAnnotation(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["annotation"] = v
}

func (P *UpdateHostParams) GetAnnotation() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["annotation"].(string)
	return value, ok
}

func (P *UpdateHostParams) SetHosttags(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hosttags"] = v
}

func (P *UpdateHostParams) GetHosttags() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hosttags"].([]string)
	return value, ok
}

func (P *UpdateHostParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateHostParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateHostParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateHostParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateHostParams) SetOscategoryid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["oscategoryid"] = v
}

func (P *UpdateHostParams) GetOscategoryid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["oscategoryid"].(string)
	return value, ok
}

func (P *UpdateHostParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *UpdateHostParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewUpdateHostParams(id string) *UpdateHostParams {
	P := &UpdateHostParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a host.
func (s *HostService) UpdateHost(p *UpdateHostParams) (*UpdateHostResponse, error) {
	resp, err := s.cs.newRequest("updateHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateHostResponse struct {
	Annotation                       string                       `json:"annotation"`
	Capabilities                     string                       `json:"capabilities"`
	Clusterid                        string                       `json:"clusterid"`
	Clustername                      string                       `json:"clustername"`
	Clustertype                      string                       `json:"clustertype"`
	Cpuallocated                     string                       `json:"cpuallocated"`
	Cpuallocatedpercentage           string                       `json:"cpuallocatedpercentage"`
	Cpuallocatedvalue                int64                        `json:"cpuallocatedvalue"`
	Cpuallocatedwithoverprovisioning string                       `json:"cpuallocatedwithoverprovisioning"`
	Cpuloadaverage                   float64                      `json:"cpuloadaverage"`
	Cpunumber                        int                          `json:"cpunumber"`
	Cpusockets                       int                          `json:"cpusockets"`
	Cpuspeed                         int64                        `json:"cpuspeed"`
	Cpuused                          string                       `json:"cpuused"`
	Cpuwithoverprovisioning          string                       `json:"cpuwithoverprovisioning"`
	Created                          string                       `json:"created"`
	Details                          map[string]string            `json:"details"`
	Disconnected                     string                       `json:"disconnected"`
	Disksizeallocated                int64                        `json:"disksizeallocated"`
	Disksizetotal                    int64                        `json:"disksizetotal"`
	Events                           string                       `json:"events"`
	Gpugroup                         []UpdateHostResponseGpugroup `json:"gpugroup"`
	Hahost                           bool                         `json:"hahost"`
	Hasannotations                   bool                         `json:"hasannotations"`
	Hasenoughcapacity                bool                         `json:"hasenoughcapacity"`
	Hostha                           HAForHostResponse            `json:"hostha"`
	Hosttags                         string                       `json:"hosttags"`
	Hypervisor                       string                       `json:"hypervisor"`
	Hypervisorversion                string                       `json:"hypervisorversion"`
	Id                               string                       `json:"id"`
	Ipaddress                        string                       `json:"ipaddress"`
	Islocalstorageactive             bool                         `json:"islocalstorageactive"`
	JobID                            string                       `json:"jobid"`
	Jobstatus                        int                          `json:"jobstatus"`
	Lastannotated                    string                       `json:"lastannotated"`
	Lastpinged                       string                       `json:"lastpinged"`
	Managementserverid               UUID                         `json:"managementserverid"`
	Memoryallocated                  int64                        `json:"memoryallocated"`
	Memoryallocatedbytes             int64                        `json:"memoryallocatedbytes"`
	Memoryallocatedpercentage        string                       `json:"memoryallocatedpercentage"`
	Memorytotal                      int64                        `json:"memorytotal"`
	Memoryused                       int64                        `json:"memoryused"`
	Memorywithoverprovisioning       string                       `json:"memorywithoverprovisioning"`
	Name                             string                       `json:"name"`
	Networkkbsread                   int64                        `json:"networkkbsread"`
	Networkkbswrite                  int64                        `json:"networkkbswrite"`
	Oscategoryid                     string                       `json:"oscategoryid"`
	Oscategoryname                   string                       `json:"oscategoryname"`
	Outofbandmanagement              OutOfBandManagementResponse  `json:"outofbandmanagement"`
	Podid                            string                       `json:"podid"`
	Podname                          string                       `json:"podname"`
	Removed                          string                       `json:"removed"`
	Resourcestate                    string                       `json:"resourcestate"`
	State                            string                       `json:"state"`
	Suitableformigration             bool                         `json:"suitableformigration"`
	Type                             string                       `json:"type"`
	Ueficapability                   bool                         `json:"ueficapability"`
	Username                         string                       `json:"username"`
	Version                          string                       `json:"version"`
	Zoneid                           string                       `json:"zoneid"`
	Zonename                         string                       `json:"zonename"`
}

type UpdateHostResponseGpugroup struct {
	Gpugroupname string                           `json:"gpugroupname"`
	Vgpu         []UpdateHostResponseGpugroupVgpu `json:"vgpu"`
}

type UpdateHostResponseGpugroupVgpu struct {
	Maxcapacity       int64  `json:"maxcapacity"`
	Maxheads          int64  `json:"maxheads"`
	Maxresolutionx    int64  `json:"maxresolutionx"`
	Maxresolutiony    int64  `json:"maxresolutiony"`
	Maxvgpuperpgpu    int64  `json:"maxvgpuperpgpu"`
	Remainingcapacity int64  `json:"remainingcapacity"`
	Vgputype          string `json:"vgputype"`
	Videoram          int64  `json:"videoram"`
}

type UpdateHostPasswordParams struct {
	P map[string]interface{}
}

func (P *UpdateHostPasswordParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["update_passwd_on_host"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("update_passwd_on_host", vv)
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *UpdateHostPasswordParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *UpdateHostPasswordParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *UpdateHostPasswordParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *UpdateHostPasswordParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *UpdateHostPasswordParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *UpdateHostPasswordParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *UpdateHostPasswordParams) SetUpdate_passwd_on_host(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["update_passwd_on_host"] = v
}

func (P *UpdateHostPasswordParams) GetUpdate_passwd_on_host() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["update_passwd_on_host"].(bool)
	return value, ok
}

func (P *UpdateHostPasswordParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *UpdateHostPasswordParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateHostPasswordParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewUpdateHostPasswordParams(password string, username string) *UpdateHostPasswordParams {
	P := &UpdateHostPasswordParams{}
	P.P = make(map[string]interface{})
	P.P["password"] = password
	P.P["username"] = username
	return P
}

// Update password of a host/pool on management server.
func (s *HostService) UpdateHostPassword(p *UpdateHostPasswordParams) (*UpdateHostPasswordResponse, error) {
	resp, err := s.cs.newRequest("updateHostPassword", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateHostPasswordResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateHostPasswordResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateHostPasswordResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateHostPasswordResponse
	return json.Unmarshal(b, (*alias)(r))
}
