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

type ClusterServiceIface interface {
	AddCluster(P *AddClusterParams) (*AddClusterResponse, error)
	NewAddClusterParams(clustername string, clustertype string, hypervisor string, podid string, zoneid string) *AddClusterParams
	DedicateCluster(P *DedicateClusterParams) (*DedicateClusterResponse, error)
	NewDedicateClusterParams(clusterid string, domainid string) *DedicateClusterParams
	DeleteCluster(P *DeleteClusterParams) (*DeleteClusterResponse, error)
	NewDeleteClusterParams(id string) *DeleteClusterParams
	DisableOutOfBandManagementForCluster(P *DisableOutOfBandManagementForClusterParams) (*DisableOutOfBandManagementForClusterResponse, error)
	NewDisableOutOfBandManagementForClusterParams(clusterid string) *DisableOutOfBandManagementForClusterParams
	EnableOutOfBandManagementForCluster(P *EnableOutOfBandManagementForClusterParams) (*EnableOutOfBandManagementForClusterResponse, error)
	NewEnableOutOfBandManagementForClusterParams(clusterid string) *EnableOutOfBandManagementForClusterParams
	EnableHAForCluster(P *EnableHAForClusterParams) (*EnableHAForClusterResponse, error)
	NewEnableHAForClusterParams(clusterid string) *EnableHAForClusterParams
	DisableHAForCluster(P *DisableHAForClusterParams) (*DisableHAForClusterResponse, error)
	NewDisableHAForClusterParams(clusterid string) *DisableHAForClusterParams
	ListClusters(P *ListClustersParams) (*ListClustersResponse, error)
	NewListClustersParams() *ListClustersParams
	GetClusterID(name string, opts ...OptionFunc) (string, int, error)
	GetClusterByName(name string, opts ...OptionFunc) (*Cluster, int, error)
	GetClusterByID(id string, opts ...OptionFunc) (*Cluster, int, error)
	ListClustersMetrics(P *ListClustersMetricsParams) (*ListClustersMetricsResponse, error)
	NewListClustersMetricsParams() *ListClustersMetricsParams
	GetClustersMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetClustersMetricByName(name string, opts ...OptionFunc) (*ClustersMetric, int, error)
	GetClustersMetricByID(id string, opts ...OptionFunc) (*ClustersMetric, int, error)
	ListDedicatedClusters(P *ListDedicatedClustersParams) (*ListDedicatedClustersResponse, error)
	NewListDedicatedClustersParams() *ListDedicatedClustersParams
	ReleaseDedicatedCluster(P *ReleaseDedicatedClusterParams) (*ReleaseDedicatedClusterResponse, error)
	NewReleaseDedicatedClusterParams(clusterid string) *ReleaseDedicatedClusterParams
	UpdateCluster(P *UpdateClusterParams) (*UpdateClusterResponse, error)
	NewUpdateClusterParams(id string) *UpdateClusterParams
}

type AddClusterParams struct {
	P map[string]interface{}
}

func (P *AddClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := P.P["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := P.P["guestvswitchname"]; found {
		u.Set("guestvswitchname", v.(string))
	}
	if v, found := P.P["guestvswitchtype"]; found {
		u.Set("guestvswitchtype", v.(string))
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["ovm3cluster"]; found {
		u.Set("ovm3cluster", v.(string))
	}
	if v, found := P.P["ovm3pool"]; found {
		u.Set("ovm3pool", v.(string))
	}
	if v, found := P.P["ovm3vip"]; found {
		u.Set("ovm3vip", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["publicvswitchname"]; found {
		u.Set("publicvswitchname", v.(string))
	}
	if v, found := P.P["publicvswitchtype"]; found {
		u.Set("publicvswitchtype", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	if v, found := P.P["vsmipaddress"]; found {
		u.Set("vsmipaddress", v.(string))
	}
	if v, found := P.P["vsmpassword"]; found {
		u.Set("vsmpassword", v.(string))
	}
	if v, found := P.P["vsmusername"]; found {
		u.Set("vsmusername", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *AddClusterParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *AddClusterParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *AddClusterParams) SetClustername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clustername"] = v
}

func (P *AddClusterParams) GetClustername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clustername"].(string)
	return value, ok
}

func (P *AddClusterParams) SetClustertype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clustertype"] = v
}

func (P *AddClusterParams) GetClustertype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clustertype"].(string)
	return value, ok
}

func (P *AddClusterParams) SetGuestvswitchname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["guestvswitchname"] = v
}

func (P *AddClusterParams) GetGuestvswitchname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["guestvswitchname"].(string)
	return value, ok
}

func (P *AddClusterParams) SetGuestvswitchtype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["guestvswitchtype"] = v
}

func (P *AddClusterParams) GetGuestvswitchtype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["guestvswitchtype"].(string)
	return value, ok
}

func (P *AddClusterParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *AddClusterParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *AddClusterParams) SetOvm3cluster(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ovm3cluster"] = v
}

func (P *AddClusterParams) GetOvm3cluster() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ovm3cluster"].(string)
	return value, ok
}

func (P *AddClusterParams) SetOvm3pool(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ovm3pool"] = v
}

func (P *AddClusterParams) GetOvm3pool() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ovm3pool"].(string)
	return value, ok
}

func (P *AddClusterParams) SetOvm3vip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ovm3vip"] = v
}

func (P *AddClusterParams) GetOvm3vip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ovm3vip"].(string)
	return value, ok
}

func (P *AddClusterParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddClusterParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddClusterParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *AddClusterParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *AddClusterParams) SetPublicvswitchname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publicvswitchname"] = v
}

func (P *AddClusterParams) GetPublicvswitchname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publicvswitchname"].(string)
	return value, ok
}

func (P *AddClusterParams) SetPublicvswitchtype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publicvswitchtype"] = v
}

func (P *AddClusterParams) GetPublicvswitchtype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publicvswitchtype"].(string)
	return value, ok
}

func (P *AddClusterParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddClusterParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddClusterParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddClusterParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

func (P *AddClusterParams) SetVsmipaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vsmipaddress"] = v
}

func (P *AddClusterParams) GetVsmipaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vsmipaddress"].(string)
	return value, ok
}

func (P *AddClusterParams) SetVsmpassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vsmpassword"] = v
}

func (P *AddClusterParams) GetVsmpassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vsmpassword"].(string)
	return value, ok
}

func (P *AddClusterParams) SetVsmusername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vsmusername"] = v
}

func (P *AddClusterParams) GetVsmusername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vsmusername"].(string)
	return value, ok
}

func (P *AddClusterParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *AddClusterParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewAddClusterParams(clustername string, clustertype string, hypervisor string, podid string, zoneid string) *AddClusterParams {
	P := &AddClusterParams{}
	P.P = make(map[string]interface{})
	P.P["clustername"] = clustername
	P.P["clustertype"] = clustertype
	P.P["hypervisor"] = hypervisor
	P.P["podid"] = podid
	P.P["zoneid"] = zoneid
	return P
}

// Adds a new cluster
func (s *ClusterService) AddCluster(p *AddClusterParams) (*AddClusterResponse, error) {
	resp, err := s.cs.newRequest("addCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddClusterResponse struct {
	Allocationstate       string                       `json:"allocationstate"`
	Capacity              []AddClusterResponseCapacity `json:"capacity"`
	Clustertype           string                       `json:"clustertype"`
	Cpuovercommitratio    string                       `json:"cpuovercommitratio"`
	Hasannotations        bool                         `json:"hasannotations"`
	Hypervisortype        string                       `json:"hypervisortype"`
	Id                    string                       `json:"id"`
	JobID                 string                       `json:"jobid"`
	Jobstatus             int                          `json:"jobstatus"`
	Managedstate          string                       `json:"managedstate"`
	Memoryovercommitratio string                       `json:"memoryovercommitratio"`
	Name                  string                       `json:"name"`
	Ovm3vip               string                       `json:"ovm3vip"`
	Podid                 string                       `json:"podid"`
	Podname               string                       `json:"podname"`
	Resourcedetails       map[string]string            `json:"resourcedetails"`
	Zoneid                string                       `json:"zoneid"`
	Zonename              string                       `json:"zonename"`
}

type AddClusterResponseCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}

type DedicateClusterParams struct {
	P map[string]interface{}
}

func (P *DedicateClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	return u
}

func (P *DedicateClusterParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DedicateClusterParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DedicateClusterParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *DedicateClusterParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *DedicateClusterParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DedicateClusterParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

// You should always use this function to get a new DedicateClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDedicateClusterParams(clusterid string, domainid string) *DedicateClusterParams {
	P := &DedicateClusterParams{}
	P.P = make(map[string]interface{})
	P.P["clusterid"] = clusterid
	P.P["domainid"] = domainid
	return P
}

// Dedicate an existing cluster
func (s *ClusterService) DedicateCluster(p *DedicateClusterParams) (*DedicateClusterResponse, error) {
	resp, err := s.cs.newRequest("dedicateCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DedicateClusterResponse
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

type DedicateClusterResponse struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Clusterid       string `json:"clusterid"`
	Clustername     string `json:"clustername"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
}

type DeleteClusterParams struct {
	P map[string]interface{}
}

func (P *DeleteClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteClusterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteClusterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDeleteClusterParams(id string) *DeleteClusterParams {
	P := &DeleteClusterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a cluster.
func (s *ClusterService) DeleteCluster(p *DeleteClusterParams) (*DeleteClusterResponse, error) {
	resp, err := s.cs.newRequest("deleteCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteClusterResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteClusterResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DisableOutOfBandManagementForClusterParams struct {
	P map[string]interface{}
}

func (P *DisableOutOfBandManagementForClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (P *DisableOutOfBandManagementForClusterParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *DisableOutOfBandManagementForClusterParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableOutOfBandManagementForClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDisableOutOfBandManagementForClusterParams(clusterid string) *DisableOutOfBandManagementForClusterParams {
	P := &DisableOutOfBandManagementForClusterParams{}
	P.P = make(map[string]interface{})
	P.P["clusterid"] = clusterid
	return P
}

// Disables out-of-band management for a cluster
func (s *ClusterService) DisableOutOfBandManagementForCluster(p *DisableOutOfBandManagementForClusterParams) (*DisableOutOfBandManagementForClusterResponse, error) {
	resp, err := s.cs.newRequest("disableOutOfBandManagementForCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableOutOfBandManagementForClusterResponse
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

type DisableOutOfBandManagementForClusterResponse struct {
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

type EnableOutOfBandManagementForClusterParams struct {
	P map[string]interface{}
}

func (P *EnableOutOfBandManagementForClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (P *EnableOutOfBandManagementForClusterParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *EnableOutOfBandManagementForClusterParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableOutOfBandManagementForClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewEnableOutOfBandManagementForClusterParams(clusterid string) *EnableOutOfBandManagementForClusterParams {
	P := &EnableOutOfBandManagementForClusterParams{}
	P.P = make(map[string]interface{})
	P.P["clusterid"] = clusterid
	return P
}

// Enables out-of-band management for a cluster
func (s *ClusterService) EnableOutOfBandManagementForCluster(p *EnableOutOfBandManagementForClusterParams) (*EnableOutOfBandManagementForClusterResponse, error) {
	resp, err := s.cs.newRequest("enableOutOfBandManagementForCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableOutOfBandManagementForClusterResponse
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

type EnableOutOfBandManagementForClusterResponse struct {
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

type EnableHAForClusterParams struct {
	P map[string]interface{}
}

func (P *EnableHAForClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (P *EnableHAForClusterParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *EnableHAForClusterParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

// You should always use this function to get a new EnableHAForClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewEnableHAForClusterParams(clusterid string) *EnableHAForClusterParams {
	P := &EnableHAForClusterParams{}
	P.P = make(map[string]interface{})
	P.P["clusterid"] = clusterid
	return P
}

// Enables HA cluster-wide
func (s *ClusterService) EnableHAForCluster(p *EnableHAForClusterParams) (*EnableHAForClusterResponse, error) {
	resp, err := s.cs.newRequest("enableHAForCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableHAForClusterResponse
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

type EnableHAForClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DisableHAForClusterParams struct {
	P map[string]interface{}
}

func (P *DisableHAForClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (P *DisableHAForClusterParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *DisableHAForClusterParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

// You should always use this function to get a new DisableHAForClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewDisableHAForClusterParams(clusterid string) *DisableHAForClusterParams {
	P := &DisableHAForClusterParams{}
	P.P = make(map[string]interface{})
	P.P["clusterid"] = clusterid
	return P
}

// Disables HA cluster-wide
func (s *ClusterService) DisableHAForCluster(p *DisableHAForClusterParams) (*DisableHAForClusterResponse, error) {
	resp, err := s.cs.newRequest("disableHAForCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableHAForClusterResponse
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

type DisableHAForClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListClustersParams struct {
	P map[string]interface{}
}

func (P *ListClustersParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["clustertype"]; found {
		u.Set("clustertype", v.(string))
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
	if v, found := P.P["managedstate"]; found {
		u.Set("managedstate", v.(string))
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
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListClustersParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *ListClustersParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *ListClustersParams) SetClustertype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clustertype"] = v
}

func (P *ListClustersParams) GetClustertype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clustertype"].(string)
	return value, ok
}

func (P *ListClustersParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *ListClustersParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *ListClustersParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListClustersParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListClustersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListClustersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListClustersParams) SetManagedstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["managedstate"] = v
}

func (P *ListClustersParams) GetManagedstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["managedstate"].(string)
	return value, ok
}

func (P *ListClustersParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListClustersParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListClustersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListClustersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListClustersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListClustersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListClustersParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListClustersParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListClustersParams) SetShowcapacities(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showcapacities"] = v
}

func (P *ListClustersParams) GetShowcapacities() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showcapacities"].(bool)
	return value, ok
}

func (P *ListClustersParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListClustersParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListClustersParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListClustersParams() *ListClustersParams {
	P := &ListClustersParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListClustersParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListClusters(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Clusters[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Clusters {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterByName(name string, opts ...OptionFunc) (*Cluster, int, error) {
	id, count, err := s.GetClusterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetClusterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClusterByID(id string, opts ...OptionFunc) (*Cluster, int, error) {
	P := &ListClustersParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListClusters(P)
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
		return l.Clusters[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Cluster UUID: %s!", id)
}

// Lists clusters.
func (s *ClusterService) ListClusters(p *ListClustersParams) (*ListClustersResponse, error) {
	resp, err := s.cs.newRequest("listClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListClustersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListClustersResponse struct {
	Count    int        `json:"count"`
	Clusters []*Cluster `json:"cluster"`
}

type Cluster struct {
	Allocationstate       string            `json:"allocationstate"`
	Capacity              []ClusterCapacity `json:"capacity"`
	Clustertype           string            `json:"clustertype"`
	Cpuovercommitratio    string            `json:"cpuovercommitratio"`
	Hasannotations        bool              `json:"hasannotations"`
	Hypervisortype        string            `json:"hypervisortype"`
	Id                    string            `json:"id"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Managedstate          string            `json:"managedstate"`
	Memoryovercommitratio string            `json:"memoryovercommitratio"`
	Name                  string            `json:"name"`
	Ovm3vip               string            `json:"ovm3vip"`
	Podid                 string            `json:"podid"`
	Podname               string            `json:"podname"`
	Resourcedetails       map[string]string `json:"resourcedetails"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type ClusterCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}

type ListClustersMetricsParams struct {
	P map[string]interface{}
}

func (P *ListClustersMetricsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["clustertype"]; found {
		u.Set("clustertype", v.(string))
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
	if v, found := P.P["managedstate"]; found {
		u.Set("managedstate", v.(string))
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
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListClustersMetricsParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *ListClustersMetricsParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *ListClustersMetricsParams) SetClustertype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clustertype"] = v
}

func (P *ListClustersMetricsParams) GetClustertype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clustertype"].(string)
	return value, ok
}

func (P *ListClustersMetricsParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *ListClustersMetricsParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *ListClustersMetricsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListClustersMetricsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListClustersMetricsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListClustersMetricsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListClustersMetricsParams) SetManagedstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["managedstate"] = v
}

func (P *ListClustersMetricsParams) GetManagedstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["managedstate"].(string)
	return value, ok
}

func (P *ListClustersMetricsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListClustersMetricsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListClustersMetricsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListClustersMetricsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListClustersMetricsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListClustersMetricsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListClustersMetricsParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListClustersMetricsParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListClustersMetricsParams) SetShowcapacities(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showcapacities"] = v
}

func (P *ListClustersMetricsParams) GetShowcapacities() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showcapacities"].(bool)
	return value, ok
}

func (P *ListClustersMetricsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListClustersMetricsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListClustersMetricsParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListClustersMetricsParams() *ListClustersMetricsParams {
	P := &ListClustersMetricsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClustersMetricID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListClustersMetricsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListClustersMetrics(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ClustersMetrics[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ClustersMetrics {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClustersMetricByName(name string, opts ...OptionFunc) (*ClustersMetric, int, error) {
	id, count, err := s.GetClustersMetricID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetClustersMetricByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ClusterService) GetClustersMetricByID(id string, opts ...OptionFunc) (*ClustersMetric, int, error) {
	P := &ListClustersMetricsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListClustersMetrics(P)
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
		return l.ClustersMetrics[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ClustersMetric UUID: %s!", id)
}

// Lists clusters metrics
func (s *ClusterService) ListClustersMetrics(p *ListClustersMetricsParams) (*ListClustersMetricsResponse, error) {
	resp, err := s.cs.newRequest("listClustersMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListClustersMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListClustersMetricsResponse struct {
	Count           int               `json:"count"`
	ClustersMetrics []*ClustersMetric `json:"clustersmetric"`
}

type ClustersMetric struct {
	Allocationstate                 string                   `json:"allocationstate"`
	Capacity                        []ClustersMetricCapacity `json:"capacity"`
	Clustertype                     string                   `json:"clustertype"`
	Cpuallocated                    string                   `json:"cpuallocated"`
	Cpuallocateddisablethreshold    bool                     `json:"cpuallocateddisablethreshold"`
	Cpuallocatedthreshold           bool                     `json:"cpuallocatedthreshold"`
	Cpudisablethreshold             bool                     `json:"cpudisablethreshold"`
	Cpumaxdeviation                 string                   `json:"cpumaxdeviation"`
	Cpuovercommitratio              string                   `json:"cpuovercommitratio"`
	Cputhreshold                    bool                     `json:"cputhreshold"`
	Cputotal                        string                   `json:"cputotal"`
	Cpuused                         string                   `json:"cpuused"`
	Hasannotations                  bool                     `json:"hasannotations"`
	Hosts                           string                   `json:"hosts"`
	Hypervisortype                  string                   `json:"hypervisortype"`
	Id                              string                   `json:"id"`
	JobID                           string                   `json:"jobid"`
	Jobstatus                       int                      `json:"jobstatus"`
	Managedstate                    string                   `json:"managedstate"`
	Memoryallocated                 string                   `json:"memoryallocated"`
	Memoryallocateddisablethreshold bool                     `json:"memoryallocateddisablethreshold"`
	Memoryallocatedthreshold        bool                     `json:"memoryallocatedthreshold"`
	Memorydisablethreshold          bool                     `json:"memorydisablethreshold"`
	Memorymaxdeviation              string                   `json:"memorymaxdeviation"`
	Memoryovercommitratio           string                   `json:"memoryovercommitratio"`
	Memorythreshold                 bool                     `json:"memorythreshold"`
	Memorytotal                     string                   `json:"memorytotal"`
	Memoryused                      string                   `json:"memoryused"`
	Name                            string                   `json:"name"`
	Ovm3vip                         string                   `json:"ovm3vip"`
	Podid                           string                   `json:"podid"`
	Podname                         string                   `json:"podname"`
	Resourcedetails                 map[string]string        `json:"resourcedetails"`
	State                           string                   `json:"state"`
	Zoneid                          string                   `json:"zoneid"`
	Zonename                        string                   `json:"zonename"`
}

type ClustersMetricCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}

type ListDedicatedClustersParams struct {
	P map[string]interface{}
}

func (P *ListDedicatedClustersParams) toURLValues() url.Values {
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
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
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

func (P *ListDedicatedClustersParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListDedicatedClustersParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListDedicatedClustersParams) SetAffinitygroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["affinitygroupid"] = v
}

func (P *ListDedicatedClustersParams) GetAffinitygroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["affinitygroupid"].(string)
	return value, ok
}

func (P *ListDedicatedClustersParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *ListDedicatedClustersParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *ListDedicatedClustersParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListDedicatedClustersParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListDedicatedClustersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListDedicatedClustersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListDedicatedClustersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListDedicatedClustersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListDedicatedClustersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListDedicatedClustersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListDedicatedClustersParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewListDedicatedClustersParams() *ListDedicatedClustersParams {
	P := &ListDedicatedClustersParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists dedicated clusters.
func (s *ClusterService) ListDedicatedClusters(p *ListDedicatedClustersParams) (*ListDedicatedClustersResponse, error) {
	resp, err := s.cs.newRequest("listDedicatedClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDedicatedClustersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDedicatedClustersResponse struct {
	Count             int                 `json:"count"`
	DedicatedClusters []*DedicatedCluster `json:"dedicatedcluster"`
}

type DedicatedCluster struct {
	Accountid       string `json:"accountid"`
	Affinitygroupid string `json:"affinitygroupid"`
	Clusterid       string `json:"clusterid"`
	Clustername     string `json:"clustername"`
	Domainid        string `json:"domainid"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
}

type ReleaseDedicatedClusterParams struct {
	P map[string]interface{}
}

func (P *ReleaseDedicatedClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	return u
}

func (P *ReleaseDedicatedClusterParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *ReleaseDedicatedClusterParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

// You should always use this function to get a new ReleaseDedicatedClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewReleaseDedicatedClusterParams(clusterid string) *ReleaseDedicatedClusterParams {
	P := &ReleaseDedicatedClusterParams{}
	P.P = make(map[string]interface{})
	P.P["clusterid"] = clusterid
	return P
}

// Release the dedication for cluster
func (s *ClusterService) ReleaseDedicatedCluster(p *ReleaseDedicatedClusterParams) (*ReleaseDedicatedClusterResponse, error) {
	resp, err := s.cs.newRequest("releaseDedicatedCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReleaseDedicatedClusterResponse
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

type ReleaseDedicatedClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateClusterParams struct {
	P map[string]interface{}
}

func (P *UpdateClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["allocationstate"]; found {
		u.Set("allocationstate", v.(string))
	}
	if v, found := P.P["clustername"]; found {
		u.Set("clustername", v.(string))
	}
	if v, found := P.P["clustertype"]; found {
		u.Set("clustertype", v.(string))
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["managedstate"]; found {
		u.Set("managedstate", v.(string))
	}
	return u
}

func (P *UpdateClusterParams) SetAllocationstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["allocationstate"] = v
}

func (P *UpdateClusterParams) GetAllocationstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["allocationstate"].(string)
	return value, ok
}

func (P *UpdateClusterParams) SetClustername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clustername"] = v
}

func (P *UpdateClusterParams) GetClustername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clustername"].(string)
	return value, ok
}

func (P *UpdateClusterParams) SetClustertype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clustertype"] = v
}

func (P *UpdateClusterParams) GetClustertype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clustertype"].(string)
	return value, ok
}

func (P *UpdateClusterParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *UpdateClusterParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *UpdateClusterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateClusterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateClusterParams) SetManagedstate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["managedstate"] = v
}

func (P *UpdateClusterParams) GetManagedstate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["managedstate"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateClusterParams instance,
// as then you are sure you have configured all required params
func (s *ClusterService) NewUpdateClusterParams(id string) *UpdateClusterParams {
	P := &UpdateClusterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates an existing cluster
func (s *ClusterService) UpdateCluster(p *UpdateClusterParams) (*UpdateClusterResponse, error) {
	resp, err := s.cs.newRequest("updateCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r UpdateClusterResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateClusterResponse struct {
	Allocationstate       string                          `json:"allocationstate"`
	Capacity              []UpdateClusterResponseCapacity `json:"capacity"`
	Clustertype           string                          `json:"clustertype"`
	Cpuovercommitratio    string                          `json:"cpuovercommitratio"`
	Hasannotations        bool                            `json:"hasannotations"`
	Hypervisortype        string                          `json:"hypervisortype"`
	Id                    string                          `json:"id"`
	JobID                 string                          `json:"jobid"`
	Jobstatus             int                             `json:"jobstatus"`
	Managedstate          string                          `json:"managedstate"`
	Memoryovercommitratio string                          `json:"memoryovercommitratio"`
	Name                  string                          `json:"name"`
	Ovm3vip               string                          `json:"ovm3vip"`
	Podid                 string                          `json:"podid"`
	Podname               string                          `json:"podname"`
	Resourcedetails       map[string]string               `json:"resourcedetails"`
	Zoneid                string                          `json:"zoneid"`
	Zonename              string                          `json:"zonename"`
}

type UpdateClusterResponseCapacity struct {
	Capacityallocated int64  `json:"capacityallocated"`
	Capacitytotal     int64  `json:"capacitytotal"`
	Capacityused      int64  `json:"capacityused"`
	Clusterid         string `json:"clusterid"`
	Clustername       string `json:"clustername"`
	Name              string `json:"name"`
	Percentused       string `json:"percentused"`
	Podid             string `json:"podid"`
	Podname           string `json:"podname"`
	Type              int    `json:"type"`
	Zoneid            string `json:"zoneid"`
	Zonename          string `json:"zonename"`
}
