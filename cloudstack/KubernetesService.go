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

type KubernetesServiceIface interface {
	AddKubernetesSupportedVersion(p *AddKubernetesSupportedVersionParams) (*AddKubernetesSupportedVersionResponse, error)
	NewAddKubernetesSupportedVersionParams(mincpunumber int, minmemory int, semanticversion string) *AddKubernetesSupportedVersionParams
	CreateKubernetesCluster(p *CreateKubernetesClusterParams) (*CreateKubernetesClusterResponse, error)
	NewCreateKubernetesClusterParams(description string, kubernetesversionid string, name string, serviceofferingid string, size int64, zoneid string) *CreateKubernetesClusterParams
	DeleteKubernetesCluster(p *DeleteKubernetesClusterParams) (*DeleteKubernetesClusterResponse, error)
	NewDeleteKubernetesClusterParams(id string) *DeleteKubernetesClusterParams
	DeleteKubernetesSupportedVersion(p *DeleteKubernetesSupportedVersionParams) (*DeleteKubernetesSupportedVersionResponse, error)
	NewDeleteKubernetesSupportedVersionParams(id string) *DeleteKubernetesSupportedVersionParams
	GetKubernetesClusterConfig(p *GetKubernetesClusterConfigParams) (*GetKubernetesClusterConfigResponse, error)
	NewGetKubernetesClusterConfigParams() *GetKubernetesClusterConfigParams
	ListKubernetesClusters(p *ListKubernetesClustersParams) (*ListKubernetesClustersResponse, error)
	NewListKubernetesClustersParams() *ListKubernetesClustersParams
	GetKubernetesClusterID(name string, opts ...OptionFunc) (string, int, error)
	GetKubernetesClusterByName(name string, opts ...OptionFunc) (*KubernetesCluster, int, error)
	GetKubernetesClusterByID(id string, opts ...OptionFunc) (*KubernetesCluster, int, error)
	ListKubernetesSupportedVersions(p *ListKubernetesSupportedVersionsParams) (*ListKubernetesSupportedVersionsResponse, error)
	NewListKubernetesSupportedVersionsParams() *ListKubernetesSupportedVersionsParams
	GetKubernetesSupportedVersionID(keyword string, opts ...OptionFunc) (string, int, error)
	GetKubernetesSupportedVersionByName(name string, opts ...OptionFunc) (*KubernetesSupportedVersion, int, error)
	GetKubernetesSupportedVersionByID(id string, opts ...OptionFunc) (*KubernetesSupportedVersion, int, error)
	ScaleKubernetesCluster(p *ScaleKubernetesClusterParams) (*ScaleKubernetesClusterResponse, error)
	NewScaleKubernetesClusterParams(id string) *ScaleKubernetesClusterParams
	StartKubernetesCluster(p *StartKubernetesClusterParams) (*StartKubernetesClusterResponse, error)
	NewStartKubernetesClusterParams(id string) *StartKubernetesClusterParams
	StopKubernetesCluster(p *StopKubernetesClusterParams) (*StopKubernetesClusterResponse, error)
	NewStopKubernetesClusterParams(id string) *StopKubernetesClusterParams
	UpdateKubernetesSupportedVersion(p *UpdateKubernetesSupportedVersionParams) (*UpdateKubernetesSupportedVersionResponse, error)
	NewUpdateKubernetesSupportedVersionParams(id string, state string) *UpdateKubernetesSupportedVersionParams
	UpgradeKubernetesCluster(p *UpgradeKubernetesClusterParams) (*UpgradeKubernetesClusterResponse, error)
	NewUpgradeKubernetesClusterParams(id string, kubernetesversionid string) *UpgradeKubernetesClusterParams
}

type AddKubernetesSupportedVersionParams struct {
	P map[string]interface{}
}

func (P *AddKubernetesSupportedVersionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := P.P["mincpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("mincpunumber", vv)
	}
	if v, found := P.P["minmemory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minmemory", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["semanticversion"]; found {
		u.Set("semanticversion", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *AddKubernetesSupportedVersionParams) SetChecksum(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["checksum"] = v
}

func (P *AddKubernetesSupportedVersionParams) GetChecksum() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["checksum"].(string)
	return value, ok
}

func (P *AddKubernetesSupportedVersionParams) SetMincpunumber(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["mincpunumber"] = v
}

func (P *AddKubernetesSupportedVersionParams) GetMincpunumber() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["mincpunumber"].(int)
	return value, ok
}

func (P *AddKubernetesSupportedVersionParams) SetMinmemory(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["minmemory"] = v
}

func (P *AddKubernetesSupportedVersionParams) GetMinmemory() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["minmemory"].(int)
	return value, ok
}

func (P *AddKubernetesSupportedVersionParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *AddKubernetesSupportedVersionParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *AddKubernetesSupportedVersionParams) SetSemanticversion(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["semanticversion"] = v
}

func (P *AddKubernetesSupportedVersionParams) GetSemanticversion() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["semanticversion"].(string)
	return value, ok
}

func (P *AddKubernetesSupportedVersionParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddKubernetesSupportedVersionParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddKubernetesSupportedVersionParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *AddKubernetesSupportedVersionParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddKubernetesSupportedVersionParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewAddKubernetesSupportedVersionParams(mincpunumber int, minmemory int, semanticversion string) *AddKubernetesSupportedVersionParams {
	P := &AddKubernetesSupportedVersionParams{}
	P.P = make(map[string]interface{})
	P.P["mincpunumber"] = mincpunumber
	P.P["minmemory"] = minmemory
	P.P["semanticversion"] = semanticversion
	return P
}

// Add a supported Kubernetes version
func (s *KubernetesService) AddKubernetesSupportedVersion(p *AddKubernetesSupportedVersionParams) (*AddKubernetesSupportedVersionResponse, error) {
	resp, err := s.cs.newRequest("addKubernetesSupportedVersion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r AddKubernetesSupportedVersionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddKubernetesSupportedVersionResponse struct {
	Id                  string `json:"id"`
	Isoid               string `json:"isoid"`
	Isoname             string `json:"isoname"`
	Isostate            string `json:"isostate"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Mincpunumber        int    `json:"mincpunumber"`
	Minmemory           int    `json:"minmemory"`
	Name                string `json:"name"`
	Semanticversion     string `json:"semanticversion"`
	State               string `json:"state"`
	Supportsautoscaling bool   `json:"supportsautoscaling"`
	Supportsha          bool   `json:"supportsha"`
	Zoneid              string `json:"zoneid"`
	Zonename            string `json:"zonename"`
}

type CreateKubernetesClusterParams struct {
	P map[string]interface{}
}

func (P *CreateKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["controlnodes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("controlnodes", vv)
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["dockerregistrypassword"]; found {
		u.Set("dockerregistrypassword", v.(string))
	}
	if v, found := P.P["dockerregistryurl"]; found {
		u.Set("dockerregistryurl", v.(string))
	}
	if v, found := P.P["dockerregistryusername"]; found {
		u.Set("dockerregistryusername", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["externalloadbalanceripaddress"]; found {
		u.Set("externalloadbalanceripaddress", v.(string))
	}
	if v, found := P.P["keypair"]; found {
		u.Set("keypair", v.(string))
	}
	if v, found := P.P["kubernetesversionid"]; found {
		u.Set("kubernetesversionid", v.(string))
	}
	if v, found := P.P["masternodes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("masternodes", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := P.P["noderootdisksize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("noderootdisksize", vv)
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := P.P["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *CreateKubernetesClusterParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateKubernetesClusterParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetControlnodes(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["controlnodes"] = v
}

func (P *CreateKubernetesClusterParams) GetControlnodes() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["controlnodes"].(int64)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateKubernetesClusterParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetDockerregistrypassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dockerregistrypassword"] = v
}

func (P *CreateKubernetesClusterParams) GetDockerregistrypassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dockerregistrypassword"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetDockerregistryurl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dockerregistryurl"] = v
}

func (P *CreateKubernetesClusterParams) GetDockerregistryurl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dockerregistryurl"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetDockerregistryusername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dockerregistryusername"] = v
}

func (P *CreateKubernetesClusterParams) GetDockerregistryusername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dockerregistryusername"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateKubernetesClusterParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetExternalloadbalanceripaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["externalloadbalanceripaddress"] = v
}

func (P *CreateKubernetesClusterParams) GetExternalloadbalanceripaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["externalloadbalanceripaddress"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetKeypair(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keypair"] = v
}

func (P *CreateKubernetesClusterParams) GetKeypair() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keypair"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetKubernetesversionid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["kubernetesversionid"] = v
}

func (P *CreateKubernetesClusterParams) GetKubernetesversionid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["kubernetesversionid"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetMasternodes(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["masternodes"] = v
}

func (P *CreateKubernetesClusterParams) GetMasternodes() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["masternodes"].(int64)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateKubernetesClusterParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *CreateKubernetesClusterParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetNoderootdisksize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["noderootdisksize"] = v
}

func (P *CreateKubernetesClusterParams) GetNoderootdisksize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["noderootdisksize"].(int64)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateKubernetesClusterParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *CreateKubernetesClusterParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetSize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["size"] = v
}

func (P *CreateKubernetesClusterParams) GetSize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["size"].(int64)
	return value, ok
}

func (P *CreateKubernetesClusterParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateKubernetesClusterParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewCreateKubernetesClusterParams(description string, kubernetesversionid string, name string, serviceofferingid string, size int64, zoneid string) *CreateKubernetesClusterParams {
	P := &CreateKubernetesClusterParams{}
	P.P = make(map[string]interface{})
	P.P["description"] = description
	P.P["kubernetesversionid"] = kubernetesversionid
	P.P["name"] = name
	P.P["serviceofferingid"] = serviceofferingid
	P.P["size"] = size
	P.P["zoneid"] = zoneid
	return P
}

// Creates a Kubernetes cluster
func (s *KubernetesService) CreateKubernetesCluster(p *CreateKubernetesClusterParams) (*CreateKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("createKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateKubernetesClusterResponse
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

type CreateKubernetesClusterResponse struct {
	Account               string            `json:"account"`
	Associatednetworkname string            `json:"associatednetworkname"`
	Autoscalingenabled    bool              `json:"autoscalingenabled"`
	Consoleendpoint       string            `json:"consoleendpoint"`
	Controlnodes          int64             `json:"controlnodes"`
	Cpunumber             string            `json:"cpunumber"`
	Description           string            `json:"description"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Endpoint              string            `json:"endpoint"`
	Hasannotations        bool              `json:"hasannotations"`
	Id                    string            `json:"id"`
	Ipaddress             string            `json:"ipaddress"`
	Ipaddressid           string            `json:"ipaddressid"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Keypair               string            `json:"keypair"`
	Kubernetesversionid   string            `json:"kubernetesversionid"`
	Kubernetesversionname string            `json:"kubernetesversionname"`
	Masternodes           int64             `json:"masternodes"`
	Maxsize               int64             `json:"maxsize"`
	Memory                string            `json:"memory"`
	Minsize               int64             `json:"minsize"`
	Name                  string            `json:"name"`
	Networkid             string            `json:"networkid"`
	Project               string            `json:"project"`
	Projectid             string            `json:"projectid"`
	Serviceofferingid     string            `json:"serviceofferingid"`
	Serviceofferingname   string            `json:"serviceofferingname"`
	Size                  int64             `json:"size"`
	State                 string            `json:"state"`
	Templateid            string            `json:"templateid"`
	Virtualmachines       []*VirtualMachine `json:"virtualmachines"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type DeleteKubernetesClusterParams struct {
	P map[string]interface{}
}

func (P *DeleteKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteKubernetesClusterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteKubernetesClusterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewDeleteKubernetesClusterParams(id string) *DeleteKubernetesClusterParams {
	P := &DeleteKubernetesClusterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a Kubernetes cluster
func (s *KubernetesService) DeleteKubernetesCluster(p *DeleteKubernetesClusterParams) (*DeleteKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("deleteKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteKubernetesClusterResponse
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

type DeleteKubernetesClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteKubernetesSupportedVersionParams struct {
	P map[string]interface{}
}

func (P *DeleteKubernetesSupportedVersionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteKubernetesSupportedVersionParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteKubernetesSupportedVersionParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteKubernetesSupportedVersionParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewDeleteKubernetesSupportedVersionParams(id string) *DeleteKubernetesSupportedVersionParams {
	P := &DeleteKubernetesSupportedVersionParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a Kubernetes cluster
func (s *KubernetesService) DeleteKubernetesSupportedVersion(p *DeleteKubernetesSupportedVersionParams) (*DeleteKubernetesSupportedVersionResponse, error) {
	resp, err := s.cs.newRequest("deleteKubernetesSupportedVersion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteKubernetesSupportedVersionResponse
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

type DeleteKubernetesSupportedVersionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type GetKubernetesClusterConfigParams struct {
	P map[string]interface{}
}

func (P *GetKubernetesClusterConfigParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *GetKubernetesClusterConfigParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *GetKubernetesClusterConfigParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new GetKubernetesClusterConfigParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewGetKubernetesClusterConfigParams() *GetKubernetesClusterConfigParams {
	P := &GetKubernetesClusterConfigParams{}
	P.P = make(map[string]interface{})
	return P
}

// Get Kubernetes cluster config
func (s *KubernetesService) GetKubernetesClusterConfig(p *GetKubernetesClusterConfigParams) (*GetKubernetesClusterConfigResponse, error) {
	resp, err := s.cs.newRequest("getKubernetesClusterConfig", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetKubernetesClusterConfigResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetKubernetesClusterConfigResponse struct {
	Configdata string `json:"configdata"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Name       string `json:"name"`
}

type ListKubernetesClustersParams struct {
	P map[string]interface{}
}

func (P *ListKubernetesClustersParams) toURLValues() url.Values {
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (P *ListKubernetesClustersParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListKubernetesClustersParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListKubernetesClustersParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListKubernetesClustersParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListKubernetesClustersParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListKubernetesClustersParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListKubernetesClustersParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListKubernetesClustersParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListKubernetesClustersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListKubernetesClustersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListKubernetesClustersParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListKubernetesClustersParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListKubernetesClustersParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListKubernetesClustersParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListKubernetesClustersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListKubernetesClustersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListKubernetesClustersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListKubernetesClustersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListKubernetesClustersParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListKubernetesClustersParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListKubernetesClustersParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListKubernetesClustersParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

// You should always use this function to get a new ListKubernetesClustersParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewListKubernetesClustersParams() *ListKubernetesClustersParams {
	P := &ListKubernetesClustersParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesClusterID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListKubernetesClustersParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListKubernetesClusters(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.KubernetesClusters[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.KubernetesClusters {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesClusterByName(name string, opts ...OptionFunc) (*KubernetesCluster, int, error) {
	id, count, err := s.GetKubernetesClusterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetKubernetesClusterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesClusterByID(id string, opts ...OptionFunc) (*KubernetesCluster, int, error) {
	P := &ListKubernetesClustersParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListKubernetesClusters(P)
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
		return l.KubernetesClusters[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for KubernetesCluster UUID: %s!", id)
}

// Lists Kubernetes clusters
func (s *KubernetesService) ListKubernetesClusters(p *ListKubernetesClustersParams) (*ListKubernetesClustersResponse, error) {
	resp, err := s.cs.newRequest("listKubernetesClusters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListKubernetesClustersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListKubernetesClustersResponse struct {
	Count              int                  `json:"count"`
	KubernetesClusters []*KubernetesCluster `json:"kubernetescluster"`
}

type KubernetesCluster struct {
	Account               string            `json:"account"`
	Associatednetworkname string            `json:"associatednetworkname"`
	Autoscalingenabled    bool              `json:"autoscalingenabled"`
	Consoleendpoint       string            `json:"consoleendpoint"`
	Controlnodes          int64             `json:"controlnodes"`
	Cpunumber             string            `json:"cpunumber"`
	Description           string            `json:"description"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Endpoint              string            `json:"endpoint"`
	Hasannotations        bool              `json:"hasannotations"`
	Id                    string            `json:"id"`
	Ipaddress             string            `json:"ipaddress"`
	Ipaddressid           string            `json:"ipaddressid"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Keypair               string            `json:"keypair"`
	Kubernetesversionid   string            `json:"kubernetesversionid"`
	Kubernetesversionname string            `json:"kubernetesversionname"`
	Masternodes           int64             `json:"masternodes"`
	Maxsize               int64             `json:"maxsize"`
	Memory                string            `json:"memory"`
	Minsize               int64             `json:"minsize"`
	Name                  string            `json:"name"`
	Networkid             string            `json:"networkid"`
	Project               string            `json:"project"`
	Projectid             string            `json:"projectid"`
	Serviceofferingid     string            `json:"serviceofferingid"`
	Serviceofferingname   string            `json:"serviceofferingname"`
	Size                  int64             `json:"size"`
	State                 string            `json:"state"`
	Templateid            string            `json:"templateid"`
	Virtualmachines       []*VirtualMachine `json:"virtualmachines"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type ListKubernetesSupportedVersionsParams struct {
	P map[string]interface{}
}

func (P *ListKubernetesSupportedVersionsParams) toURLValues() url.Values {
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
	if v, found := P.P["minimumkubernetesversionid"]; found {
		u.Set("minimumkubernetesversionid", v.(string))
	}
	if v, found := P.P["minimumsemanticversion"]; found {
		u.Set("minimumsemanticversion", v.(string))
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

func (P *ListKubernetesSupportedVersionsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListKubernetesSupportedVersionsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListKubernetesSupportedVersionsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListKubernetesSupportedVersionsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListKubernetesSupportedVersionsParams) SetMinimumkubernetesversionid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["minimumkubernetesversionid"] = v
}

func (P *ListKubernetesSupportedVersionsParams) GetMinimumkubernetesversionid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["minimumkubernetesversionid"].(string)
	return value, ok
}

func (P *ListKubernetesSupportedVersionsParams) SetMinimumsemanticversion(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["minimumsemanticversion"] = v
}

func (P *ListKubernetesSupportedVersionsParams) GetMinimumsemanticversion() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["minimumsemanticversion"].(string)
	return value, ok
}

func (P *ListKubernetesSupportedVersionsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListKubernetesSupportedVersionsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListKubernetesSupportedVersionsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListKubernetesSupportedVersionsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListKubernetesSupportedVersionsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListKubernetesSupportedVersionsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListKubernetesSupportedVersionsParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewListKubernetesSupportedVersionsParams() *ListKubernetesSupportedVersionsParams {
	P := &ListKubernetesSupportedVersionsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesSupportedVersionID(keyword string, opts ...OptionFunc) (string, int, error) {
	P := &ListKubernetesSupportedVersionsParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListKubernetesSupportedVersions(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.KubernetesSupportedVersions[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.KubernetesSupportedVersions {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesSupportedVersionByName(name string, opts ...OptionFunc) (*KubernetesSupportedVersion, int, error) {
	id, count, err := s.GetKubernetesSupportedVersionID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetKubernetesSupportedVersionByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *KubernetesService) GetKubernetesSupportedVersionByID(id string, opts ...OptionFunc) (*KubernetesSupportedVersion, int, error) {
	P := &ListKubernetesSupportedVersionsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListKubernetesSupportedVersions(P)
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
		return l.KubernetesSupportedVersions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for KubernetesSupportedVersion UUID: %s!", id)
}

// Lists supported Kubernetes version
func (s *KubernetesService) ListKubernetesSupportedVersions(p *ListKubernetesSupportedVersionsParams) (*ListKubernetesSupportedVersionsResponse, error) {
	resp, err := s.cs.newRequest("listKubernetesSupportedVersions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListKubernetesSupportedVersionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListKubernetesSupportedVersionsResponse struct {
	Count                       int                           `json:"count"`
	KubernetesSupportedVersions []*KubernetesSupportedVersion `json:"kubernetessupportedversion"`
}

type KubernetesSupportedVersion struct {
	Id                  string `json:"id"`
	Isoid               string `json:"isoid"`
	Isoname             string `json:"isoname"`
	Isostate            string `json:"isostate"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Mincpunumber        int    `json:"mincpunumber"`
	Minmemory           int    `json:"minmemory"`
	Name                string `json:"name"`
	Semanticversion     string `json:"semanticversion"`
	State               string `json:"state"`
	Supportsautoscaling bool   `json:"supportsautoscaling"`
	Supportsha          bool   `json:"supportsha"`
	Zoneid              string `json:"zoneid"`
	Zonename            string `json:"zonename"`
}

type ScaleKubernetesClusterParams struct {
	P map[string]interface{}
}

func (P *ScaleKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["autoscalingenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("autoscalingenabled", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["maxsize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxsize", vv)
	}
	if v, found := P.P["minsize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("minsize", vv)
	}
	if v, found := P.P["nodeids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("nodeids", vv)
	}
	if v, found := P.P["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := P.P["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	return u
}

func (P *ScaleKubernetesClusterParams) SetAutoscalingenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["autoscalingenabled"] = v
}

func (P *ScaleKubernetesClusterParams) GetAutoscalingenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["autoscalingenabled"].(bool)
	return value, ok
}

func (P *ScaleKubernetesClusterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ScaleKubernetesClusterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ScaleKubernetesClusterParams) SetMaxsize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxsize"] = v
}

func (P *ScaleKubernetesClusterParams) GetMaxsize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxsize"].(int64)
	return value, ok
}

func (P *ScaleKubernetesClusterParams) SetMinsize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["minsize"] = v
}

func (P *ScaleKubernetesClusterParams) GetMinsize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["minsize"].(int64)
	return value, ok
}

func (P *ScaleKubernetesClusterParams) SetNodeids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nodeids"] = v
}

func (P *ScaleKubernetesClusterParams) GetNodeids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nodeids"].([]string)
	return value, ok
}

func (P *ScaleKubernetesClusterParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *ScaleKubernetesClusterParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

func (P *ScaleKubernetesClusterParams) SetSize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["size"] = v
}

func (P *ScaleKubernetesClusterParams) GetSize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["size"].(int64)
	return value, ok
}

// You should always use this function to get a new ScaleKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewScaleKubernetesClusterParams(id string) *ScaleKubernetesClusterParams {
	P := &ScaleKubernetesClusterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Scales a created, running or stopped Kubernetes cluster
func (s *KubernetesService) ScaleKubernetesCluster(p *ScaleKubernetesClusterParams) (*ScaleKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("scaleKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ScaleKubernetesClusterResponse
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

type ScaleKubernetesClusterResponse struct {
	Account               string            `json:"account"`
	Associatednetworkname string            `json:"associatednetworkname"`
	Autoscalingenabled    bool              `json:"autoscalingenabled"`
	Consoleendpoint       string            `json:"consoleendpoint"`
	Controlnodes          int64             `json:"controlnodes"`
	Cpunumber             string            `json:"cpunumber"`
	Description           string            `json:"description"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Endpoint              string            `json:"endpoint"`
	Hasannotations        bool              `json:"hasannotations"`
	Id                    string            `json:"id"`
	Ipaddress             string            `json:"ipaddress"`
	Ipaddressid           string            `json:"ipaddressid"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Keypair               string            `json:"keypair"`
	Kubernetesversionid   string            `json:"kubernetesversionid"`
	Kubernetesversionname string            `json:"kubernetesversionname"`
	Masternodes           int64             `json:"masternodes"`
	Maxsize               int64             `json:"maxsize"`
	Memory                string            `json:"memory"`
	Minsize               int64             `json:"minsize"`
	Name                  string            `json:"name"`
	Networkid             string            `json:"networkid"`
	Project               string            `json:"project"`
	Projectid             string            `json:"projectid"`
	Serviceofferingid     string            `json:"serviceofferingid"`
	Serviceofferingname   string            `json:"serviceofferingname"`
	Size                  int64             `json:"size"`
	State                 string            `json:"state"`
	Templateid            string            `json:"templateid"`
	Virtualmachines       []*VirtualMachine `json:"virtualmachines"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type StartKubernetesClusterParams struct {
	P map[string]interface{}
}

func (P *StartKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *StartKubernetesClusterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *StartKubernetesClusterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new StartKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewStartKubernetesClusterParams(id string) *StartKubernetesClusterParams {
	P := &StartKubernetesClusterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Starts a stopped Kubernetes cluster
func (s *KubernetesService) StartKubernetesCluster(p *StartKubernetesClusterParams) (*StartKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("startKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartKubernetesClusterResponse
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

type StartKubernetesClusterResponse struct {
	Account               string            `json:"account"`
	Associatednetworkname string            `json:"associatednetworkname"`
	Autoscalingenabled    bool              `json:"autoscalingenabled"`
	Consoleendpoint       string            `json:"consoleendpoint"`
	Controlnodes          int64             `json:"controlnodes"`
	Cpunumber             string            `json:"cpunumber"`
	Description           string            `json:"description"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Endpoint              string            `json:"endpoint"`
	Hasannotations        bool              `json:"hasannotations"`
	Id                    string            `json:"id"`
	Ipaddress             string            `json:"ipaddress"`
	Ipaddressid           string            `json:"ipaddressid"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Keypair               string            `json:"keypair"`
	Kubernetesversionid   string            `json:"kubernetesversionid"`
	Kubernetesversionname string            `json:"kubernetesversionname"`
	Masternodes           int64             `json:"masternodes"`
	Maxsize               int64             `json:"maxsize"`
	Memory                string            `json:"memory"`
	Minsize               int64             `json:"minsize"`
	Name                  string            `json:"name"`
	Networkid             string            `json:"networkid"`
	Project               string            `json:"project"`
	Projectid             string            `json:"projectid"`
	Serviceofferingid     string            `json:"serviceofferingid"`
	Serviceofferingname   string            `json:"serviceofferingname"`
	Size                  int64             `json:"size"`
	State                 string            `json:"state"`
	Templateid            string            `json:"templateid"`
	Virtualmachines       []*VirtualMachine `json:"virtualmachines"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}

type StopKubernetesClusterParams struct {
	P map[string]interface{}
}

func (P *StopKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *StopKubernetesClusterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *StopKubernetesClusterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new StopKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewStopKubernetesClusterParams(id string) *StopKubernetesClusterParams {
	P := &StopKubernetesClusterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Stops a running Kubernetes cluster
func (s *KubernetesService) StopKubernetesCluster(p *StopKubernetesClusterParams) (*StopKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("stopKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopKubernetesClusterResponse
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

type StopKubernetesClusterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateKubernetesSupportedVersionParams struct {
	P map[string]interface{}
}

func (P *UpdateKubernetesSupportedVersionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (P *UpdateKubernetesSupportedVersionParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateKubernetesSupportedVersionParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateKubernetesSupportedVersionParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *UpdateKubernetesSupportedVersionParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateKubernetesSupportedVersionParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewUpdateKubernetesSupportedVersionParams(id string, state string) *UpdateKubernetesSupportedVersionParams {
	P := &UpdateKubernetesSupportedVersionParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["state"] = state
	return P
}

// Update a supported Kubernetes version
func (s *KubernetesService) UpdateKubernetesSupportedVersion(p *UpdateKubernetesSupportedVersionParams) (*UpdateKubernetesSupportedVersionResponse, error) {
	resp, err := s.cs.newRequest("updateKubernetesSupportedVersion", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateKubernetesSupportedVersionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateKubernetesSupportedVersionResponse struct {
	Id                  string `json:"id"`
	Isoid               string `json:"isoid"`
	Isoname             string `json:"isoname"`
	Isostate            string `json:"isostate"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Mincpunumber        int    `json:"mincpunumber"`
	Minmemory           int    `json:"minmemory"`
	Name                string `json:"name"`
	Semanticversion     string `json:"semanticversion"`
	State               string `json:"state"`
	Supportsautoscaling bool   `json:"supportsautoscaling"`
	Supportsha          bool   `json:"supportsha"`
	Zoneid              string `json:"zoneid"`
	Zonename            string `json:"zonename"`
}

type UpgradeKubernetesClusterParams struct {
	P map[string]interface{}
}

func (P *UpgradeKubernetesClusterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["kubernetesversionid"]; found {
		u.Set("kubernetesversionid", v.(string))
	}
	return u
}

func (P *UpgradeKubernetesClusterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpgradeKubernetesClusterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpgradeKubernetesClusterParams) SetKubernetesversionid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["kubernetesversionid"] = v
}

func (P *UpgradeKubernetesClusterParams) GetKubernetesversionid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["kubernetesversionid"].(string)
	return value, ok
}

// You should always use this function to get a new UpgradeKubernetesClusterParams instance,
// as then you are sure you have configured all required params
func (s *KubernetesService) NewUpgradeKubernetesClusterParams(id string, kubernetesversionid string) *UpgradeKubernetesClusterParams {
	P := &UpgradeKubernetesClusterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["kubernetesversionid"] = kubernetesversionid
	return P
}

// Upgrades a running Kubernetes cluster
func (s *KubernetesService) UpgradeKubernetesCluster(p *UpgradeKubernetesClusterParams) (*UpgradeKubernetesClusterResponse, error) {
	resp, err := s.cs.newRequest("upgradeKubernetesCluster", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpgradeKubernetesClusterResponse
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

type UpgradeKubernetesClusterResponse struct {
	Account               string            `json:"account"`
	Associatednetworkname string            `json:"associatednetworkname"`
	Autoscalingenabled    bool              `json:"autoscalingenabled"`
	Consoleendpoint       string            `json:"consoleendpoint"`
	Controlnodes          int64             `json:"controlnodes"`
	Cpunumber             string            `json:"cpunumber"`
	Description           string            `json:"description"`
	Domain                string            `json:"domain"`
	Domainid              string            `json:"domainid"`
	Endpoint              string            `json:"endpoint"`
	Hasannotations        bool              `json:"hasannotations"`
	Id                    string            `json:"id"`
	Ipaddress             string            `json:"ipaddress"`
	Ipaddressid           string            `json:"ipaddressid"`
	JobID                 string            `json:"jobid"`
	Jobstatus             int               `json:"jobstatus"`
	Keypair               string            `json:"keypair"`
	Kubernetesversionid   string            `json:"kubernetesversionid"`
	Kubernetesversionname string            `json:"kubernetesversionname"`
	Masternodes           int64             `json:"masternodes"`
	Maxsize               int64             `json:"maxsize"`
	Memory                string            `json:"memory"`
	Minsize               int64             `json:"minsize"`
	Name                  string            `json:"name"`
	Networkid             string            `json:"networkid"`
	Project               string            `json:"project"`
	Projectid             string            `json:"projectid"`
	Serviceofferingid     string            `json:"serviceofferingid"`
	Serviceofferingname   string            `json:"serviceofferingname"`
	Size                  int64             `json:"size"`
	State                 string            `json:"state"`
	Templateid            string            `json:"templateid"`
	Virtualmachines       []*VirtualMachine `json:"virtualmachines"`
	Zoneid                string            `json:"zoneid"`
	Zonename              string            `json:"zonename"`
}
