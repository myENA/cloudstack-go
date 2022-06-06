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

type VirtualMachineServiceIface interface {
	AddNicToVirtualMachine(P *AddNicToVirtualMachineParams) (*AddNicToVirtualMachineResponse, error)
	NewAddNicToVirtualMachineParams(networkid string, virtualmachineid string) *AddNicToVirtualMachineParams
	AssignVirtualMachine(P *AssignVirtualMachineParams) (*AssignVirtualMachineResponse, error)
	NewAssignVirtualMachineParams(virtualmachineid string) *AssignVirtualMachineParams
	ChangeServiceForVirtualMachine(P *ChangeServiceForVirtualMachineParams) (*ChangeServiceForVirtualMachineResponse, error)
	NewChangeServiceForVirtualMachineParams(id string, serviceofferingid string) *ChangeServiceForVirtualMachineParams
	CleanVMReservations(P *CleanVMReservationsParams) (*CleanVMReservationsResponse, error)
	NewCleanVMReservationsParams() *CleanVMReservationsParams
	DeployVirtualMachine(P *DeployVirtualMachineParams) (*DeployVirtualMachineResponse, error)
	NewDeployVirtualMachineParams(serviceofferingid string, templateid string, zoneid string) *DeployVirtualMachineParams
	DestroyVirtualMachine(P *DestroyVirtualMachineParams) (*DestroyVirtualMachineResponse, error)
	NewDestroyVirtualMachineParams(id string) *DestroyVirtualMachineParams
	ExpungeVirtualMachine(P *ExpungeVirtualMachineParams) (*ExpungeVirtualMachineResponse, error)
	NewExpungeVirtualMachineParams(id string) *ExpungeVirtualMachineParams
	GetVMPassword(P *GetVMPasswordParams) (*GetVMPasswordResponse, error)
	NewGetVMPasswordParams(id string) *GetVMPasswordParams
	ListVirtualMachines(P *ListVirtualMachinesParams) (*ListVirtualMachinesResponse, error)
	NewListVirtualMachinesParams() *ListVirtualMachinesParams
	GetVirtualMachineID(name string, opts ...OptionFunc) (string, int, error)
	GetVirtualMachineByName(name string, opts ...OptionFunc) (*VirtualMachine, int, error)
	GetVirtualMachineByID(id string, opts ...OptionFunc) (*VirtualMachine, int, error)
	ListVirtualMachinesMetrics(P *ListVirtualMachinesMetricsParams) (*ListVirtualMachinesMetricsResponse, error)
	NewListVirtualMachinesMetricsParams() *ListVirtualMachinesMetricsParams
	GetVirtualMachinesMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetVirtualMachinesMetricByName(name string, opts ...OptionFunc) (*VirtualMachinesMetric, int, error)
	GetVirtualMachinesMetricByID(id string, opts ...OptionFunc) (*VirtualMachinesMetric, int, error)
	MigrateVirtualMachine(P *MigrateVirtualMachineParams) (*MigrateVirtualMachineResponse, error)
	NewMigrateVirtualMachineParams(virtualmachineid string) *MigrateVirtualMachineParams
	MigrateVirtualMachineWithVolume(P *MigrateVirtualMachineWithVolumeParams) (*MigrateVirtualMachineWithVolumeResponse, error)
	NewMigrateVirtualMachineWithVolumeParams(virtualmachineid string) *MigrateVirtualMachineWithVolumeParams
	RebootVirtualMachine(P *RebootVirtualMachineParams) (*RebootVirtualMachineResponse, error)
	NewRebootVirtualMachineParams(id string) *RebootVirtualMachineParams
	RecoverVirtualMachine(P *RecoverVirtualMachineParams) (*RecoverVirtualMachineResponse, error)
	NewRecoverVirtualMachineParams(id string) *RecoverVirtualMachineParams
	RemoveNicFromVirtualMachine(P *RemoveNicFromVirtualMachineParams) (*RemoveNicFromVirtualMachineResponse, error)
	NewRemoveNicFromVirtualMachineParams(nicid string, virtualmachineid string) *RemoveNicFromVirtualMachineParams
	ResetPasswordForVirtualMachine(P *ResetPasswordForVirtualMachineParams) (*ResetPasswordForVirtualMachineResponse, error)
	NewResetPasswordForVirtualMachineParams(id string) *ResetPasswordForVirtualMachineParams
	RestoreVirtualMachine(P *RestoreVirtualMachineParams) (*RestoreVirtualMachineResponse, error)
	NewRestoreVirtualMachineParams(virtualmachineid string) *RestoreVirtualMachineParams
	ScaleVirtualMachine(P *ScaleVirtualMachineParams) (*ScaleVirtualMachineResponse, error)
	NewScaleVirtualMachineParams(id string, serviceofferingid string) *ScaleVirtualMachineParams
	StartVirtualMachine(P *StartVirtualMachineParams) (*StartVirtualMachineResponse, error)
	NewStartVirtualMachineParams(id string) *StartVirtualMachineParams
	StopVirtualMachine(P *StopVirtualMachineParams) (*StopVirtualMachineResponse, error)
	NewStopVirtualMachineParams(id string) *StopVirtualMachineParams
	UpdateDefaultNicForVirtualMachine(P *UpdateDefaultNicForVirtualMachineParams) (*UpdateDefaultNicForVirtualMachineResponse, error)
	NewUpdateDefaultNicForVirtualMachineParams(nicid string, virtualmachineid string) *UpdateDefaultNicForVirtualMachineParams
	UpdateVirtualMachine(P *UpdateVirtualMachineParams) (*UpdateVirtualMachineResponse, error)
	NewUpdateVirtualMachineParams(id string) *UpdateVirtualMachineParams
}

type AddNicToVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *AddNicToVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["dhcpoptions"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("dhcpoptions[%d].key", i), k)
			u.Set(fmt.Sprintf("dhcpoptions[%d].value", i), m[k])
		}
	}
	if v, found := P.P["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := P.P["macaddress"]; found {
		u.Set("macaddress", v.(string))
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *AddNicToVirtualMachineParams) SetDhcpoptions(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dhcpoptions"] = v
}

func (P *AddNicToVirtualMachineParams) GetDhcpoptions() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dhcpoptions"].(map[string]string)
	return value, ok
}

func (P *AddNicToVirtualMachineParams) SetIpaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddress"] = v
}

func (P *AddNicToVirtualMachineParams) GetIpaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddress"].(string)
	return value, ok
}

func (P *AddNicToVirtualMachineParams) SetMacaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["macaddress"] = v
}

func (P *AddNicToVirtualMachineParams) GetMacaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["macaddress"].(string)
	return value, ok
}

func (P *AddNicToVirtualMachineParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *AddNicToVirtualMachineParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *AddNicToVirtualMachineParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *AddNicToVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new AddNicToVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewAddNicToVirtualMachineParams(networkid string, virtualmachineid string) *AddNicToVirtualMachineParams {
	P := &AddNicToVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["networkid"] = networkid
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Adds VM to specified network by creating a NIC
func (s *VirtualMachineService) AddNicToVirtualMachine(p *AddNicToVirtualMachineParams) (*AddNicToVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("addNicToVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNicToVirtualMachineResponse
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

type AddNicToVirtualMachineResponse struct {
	Account               string                                        `json:"account"`
	Affinitygroup         []AddNicToVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                        `json:"backupofferingid"`
	Backupofferingname    string                                        `json:"backupofferingname"`
	Bootmode              string                                        `json:"bootmode"`
	Boottype              string                                        `json:"boottype"`
	Cpunumber             int                                           `json:"cpunumber"`
	Cpuspeed              int                                           `json:"cpuspeed"`
	Cpuused               string                                        `json:"cpuused"`
	Created               string                                        `json:"created"`
	Details               map[string]string                             `json:"details"`
	Diskioread            int64                                         `json:"diskioread"`
	Diskiowrite           int64                                         `json:"diskiowrite"`
	Diskkbsread           int64                                         `json:"diskkbsread"`
	Diskkbswrite          int64                                         `json:"diskkbswrite"`
	Diskofferingid        string                                        `json:"diskofferingid"`
	Diskofferingname      string                                        `json:"diskofferingname"`
	Displayname           string                                        `json:"displayname"`
	Displayvm             bool                                          `json:"displayvm"`
	Domain                string                                        `json:"domain"`
	Domainid              string                                        `json:"domainid"`
	Forvirtualnetwork     bool                                          `json:"forvirtualnetwork"`
	Group                 string                                        `json:"group"`
	Groupid               string                                        `json:"groupid"`
	Guestosid             string                                        `json:"guestosid"`
	Haenable              bool                                          `json:"haenable"`
	Hasannotations        bool                                          `json:"hasannotations"`
	Hostid                string                                        `json:"hostid"`
	Hostname              string                                        `json:"hostname"`
	Hypervisor            string                                        `json:"hypervisor"`
	Icon                  string                                        `json:"icon"`
	Id                    string                                        `json:"id"`
	Instancename          string                                        `json:"instancename"`
	Isdynamicallyscalable bool                                          `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                        `json:"isodisplaytext"`
	Isoid                 string                                        `json:"isoid"`
	Isoname               string                                        `json:"isoname"`
	JobID                 string                                        `json:"jobid"`
	Jobstatus             int                                           `json:"jobstatus"`
	Keypair               string                                        `json:"keypair"`
	Lastupdated           string                                        `json:"lastupdated"`
	Memory                int                                           `json:"memory"`
	Memoryintfreekbs      int64                                         `json:"memoryintfreekbs"`
	Memorykbs             int64                                         `json:"memorykbs"`
	Memorytargetkbs       int64                                         `json:"memorytargetkbs"`
	Name                  string                                        `json:"name"`
	Networkkbsread        int64                                         `json:"networkkbsread"`
	Networkkbswrite       int64                                         `json:"networkkbswrite"`
	Nic                   []Nic                                         `json:"nic"`
	Osdisplayname         string                                        `json:"osdisplayname"`
	Ostypeid              string                                        `json:"ostypeid"`
	Password              string                                        `json:"password"`
	Passwordenabled       bool                                          `json:"passwordenabled"`
	Pooltype              string                                        `json:"pooltype"`
	Project               string                                        `json:"project"`
	Projectid             string                                        `json:"projectid"`
	Publicip              string                                        `json:"publicip"`
	Publicipid            string                                        `json:"publicipid"`
	Readonlydetails       string                                        `json:"readonlydetails"`
	Receivedbytes         int64                                         `json:"receivedbytes"`
	Rootdeviceid          int64                                         `json:"rootdeviceid"`
	Rootdevicetype        string                                        `json:"rootdevicetype"`
	Securitygroup         []AddNicToVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                         `json:"sentbytes"`
	Serviceofferingid     string                                        `json:"serviceofferingid"`
	Serviceofferingname   string                                        `json:"serviceofferingname"`
	Servicestate          string                                        `json:"servicestate"`
	State                 string                                        `json:"state"`
	Tags                  []Tags                                        `json:"tags"`
	Templatedisplaytext   string                                        `json:"templatedisplaytext"`
	Templateid            string                                        `json:"templateid"`
	Templatename          string                                        `json:"templatename"`
	Userid                string                                        `json:"userid"`
	Username              string                                        `json:"username"`
	Vgpu                  string                                        `json:"vgpu"`
	Zoneid                string                                        `json:"zoneid"`
	Zonename              string                                        `json:"zonename"`
}

type AddNicToVirtualMachineResponseSecuritygroup struct {
	Account             string                                            `json:"account"`
	Description         string                                            `json:"description"`
	Domain              string                                            `json:"domain"`
	Domainid            string                                            `json:"domainid"`
	Egressrule          []AddNicToVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                            `json:"id"`
	Ingressrule         []AddNicToVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                            `json:"name"`
	Project             string                                            `json:"project"`
	Projectid           string                                            `json:"projectid"`
	Tags                []Tags                                            `json:"tags"`
	Virtualmachinecount int                                               `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                     `json:"virtualmachineids"`
}

type AddNicToVirtualMachineResponseSecuritygroupRule struct {
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

type AddNicToVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *AddNicToVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias AddNicToVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type AssignVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *AssignVirtualMachineParams) toURLValues() url.Values {
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
	if v, found := P.P["networkids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("networkids", vv)
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["securitygroupids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("securitygroupids", vv)
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *AssignVirtualMachineParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *AssignVirtualMachineParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *AssignVirtualMachineParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *AssignVirtualMachineParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *AssignVirtualMachineParams) SetNetworkids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkids"] = v
}

func (P *AssignVirtualMachineParams) GetNetworkids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkids"].([]string)
	return value, ok
}

func (P *AssignVirtualMachineParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *AssignVirtualMachineParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *AssignVirtualMachineParams) SetSecuritygroupids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupids"] = v
}

func (P *AssignVirtualMachineParams) GetSecuritygroupids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupids"].([]string)
	return value, ok
}

func (P *AssignVirtualMachineParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *AssignVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new AssignVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewAssignVirtualMachineParams(virtualmachineid string) *AssignVirtualMachineParams {
	P := &AssignVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Change ownership of a VM from one account to another. This API is available for Basic zones with security groups and Advanced zones with guest networks. A root administrator can reassign a VM from any account to any other account in any domain. A domain administrator can reassign a VM to any account in the same domain.
func (s *VirtualMachineService) AssignVirtualMachine(p *AssignVirtualMachineParams) (*AssignVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("assignVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AssignVirtualMachineResponse struct {
	Account               string                                      `json:"account"`
	Affinitygroup         []AssignVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                      `json:"backupofferingid"`
	Backupofferingname    string                                      `json:"backupofferingname"`
	Bootmode              string                                      `json:"bootmode"`
	Boottype              string                                      `json:"boottype"`
	Cpunumber             int                                         `json:"cpunumber"`
	Cpuspeed              int                                         `json:"cpuspeed"`
	Cpuused               string                                      `json:"cpuused"`
	Created               string                                      `json:"created"`
	Details               map[string]string                           `json:"details"`
	Diskioread            int64                                       `json:"diskioread"`
	Diskiowrite           int64                                       `json:"diskiowrite"`
	Diskkbsread           int64                                       `json:"diskkbsread"`
	Diskkbswrite          int64                                       `json:"diskkbswrite"`
	Diskofferingid        string                                      `json:"diskofferingid"`
	Diskofferingname      string                                      `json:"diskofferingname"`
	Displayname           string                                      `json:"displayname"`
	Displayvm             bool                                        `json:"displayvm"`
	Domain                string                                      `json:"domain"`
	Domainid              string                                      `json:"domainid"`
	Forvirtualnetwork     bool                                        `json:"forvirtualnetwork"`
	Group                 string                                      `json:"group"`
	Groupid               string                                      `json:"groupid"`
	Guestosid             string                                      `json:"guestosid"`
	Haenable              bool                                        `json:"haenable"`
	Hasannotations        bool                                        `json:"hasannotations"`
	Hostid                string                                      `json:"hostid"`
	Hostname              string                                      `json:"hostname"`
	Hypervisor            string                                      `json:"hypervisor"`
	Icon                  string                                      `json:"icon"`
	Id                    string                                      `json:"id"`
	Instancename          string                                      `json:"instancename"`
	Isdynamicallyscalable bool                                        `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                      `json:"isodisplaytext"`
	Isoid                 string                                      `json:"isoid"`
	Isoname               string                                      `json:"isoname"`
	JobID                 string                                      `json:"jobid"`
	Jobstatus             int                                         `json:"jobstatus"`
	Keypair               string                                      `json:"keypair"`
	Lastupdated           string                                      `json:"lastupdated"`
	Memory                int                                         `json:"memory"`
	Memoryintfreekbs      int64                                       `json:"memoryintfreekbs"`
	Memorykbs             int64                                       `json:"memorykbs"`
	Memorytargetkbs       int64                                       `json:"memorytargetkbs"`
	Name                  string                                      `json:"name"`
	Networkkbsread        int64                                       `json:"networkkbsread"`
	Networkkbswrite       int64                                       `json:"networkkbswrite"`
	Nic                   []Nic                                       `json:"nic"`
	Osdisplayname         string                                      `json:"osdisplayname"`
	Ostypeid              string                                      `json:"ostypeid"`
	Password              string                                      `json:"password"`
	Passwordenabled       bool                                        `json:"passwordenabled"`
	Pooltype              string                                      `json:"pooltype"`
	Project               string                                      `json:"project"`
	Projectid             string                                      `json:"projectid"`
	Publicip              string                                      `json:"publicip"`
	Publicipid            string                                      `json:"publicipid"`
	Readonlydetails       string                                      `json:"readonlydetails"`
	Receivedbytes         int64                                       `json:"receivedbytes"`
	Rootdeviceid          int64                                       `json:"rootdeviceid"`
	Rootdevicetype        string                                      `json:"rootdevicetype"`
	Securitygroup         []AssignVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                       `json:"sentbytes"`
	Serviceofferingid     string                                      `json:"serviceofferingid"`
	Serviceofferingname   string                                      `json:"serviceofferingname"`
	Servicestate          string                                      `json:"servicestate"`
	State                 string                                      `json:"state"`
	Tags                  []Tags                                      `json:"tags"`
	Templatedisplaytext   string                                      `json:"templatedisplaytext"`
	Templateid            string                                      `json:"templateid"`
	Templatename          string                                      `json:"templatename"`
	Userid                string                                      `json:"userid"`
	Username              string                                      `json:"username"`
	Vgpu                  string                                      `json:"vgpu"`
	Zoneid                string                                      `json:"zoneid"`
	Zonename              string                                      `json:"zonename"`
}

type AssignVirtualMachineResponseSecuritygroup struct {
	Account             string                                          `json:"account"`
	Description         string                                          `json:"description"`
	Domain              string                                          `json:"domain"`
	Domainid            string                                          `json:"domainid"`
	Egressrule          []AssignVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                          `json:"id"`
	Ingressrule         []AssignVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                          `json:"name"`
	Project             string                                          `json:"project"`
	Projectid           string                                          `json:"projectid"`
	Tags                []Tags                                          `json:"tags"`
	Virtualmachinecount int                                             `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                   `json:"virtualmachineids"`
}

type AssignVirtualMachineResponseSecuritygroupRule struct {
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

type AssignVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *AssignVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias AssignVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ChangeServiceForVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *ChangeServiceForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (P *ChangeServiceForVirtualMachineParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ChangeServiceForVirtualMachineParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *ChangeServiceForVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ChangeServiceForVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ChangeServiceForVirtualMachineParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *ChangeServiceForVirtualMachineParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

// You should always use this function to get a new ChangeServiceForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewChangeServiceForVirtualMachineParams(id string, serviceofferingid string) *ChangeServiceForVirtualMachineParams {
	P := &ChangeServiceForVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["serviceofferingid"] = serviceofferingid
	return P
}

// Changes the service offering for a virtual machine. The virtual machine must be in a "Stopped" state for this command to take effect. Note that it only changes the VM's compute offering and it does not update the root volume offering. If the Service Offering has a root disk size the volume will be resized only if using API command 'scaleVirtualMachine'.
func (s *VirtualMachineService) ChangeServiceForVirtualMachine(p *ChangeServiceForVirtualMachineParams) (*ChangeServiceForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("changeServiceForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ChangeServiceForVirtualMachineResponse struct {
	Account               string                                                `json:"account"`
	Affinitygroup         []ChangeServiceForVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                                `json:"backupofferingid"`
	Backupofferingname    string                                                `json:"backupofferingname"`
	Bootmode              string                                                `json:"bootmode"`
	Boottype              string                                                `json:"boottype"`
	Cpunumber             int                                                   `json:"cpunumber"`
	Cpuspeed              int                                                   `json:"cpuspeed"`
	Cpuused               string                                                `json:"cpuused"`
	Created               string                                                `json:"created"`
	Details               map[string]string                                     `json:"details"`
	Diskioread            int64                                                 `json:"diskioread"`
	Diskiowrite           int64                                                 `json:"diskiowrite"`
	Diskkbsread           int64                                                 `json:"diskkbsread"`
	Diskkbswrite          int64                                                 `json:"diskkbswrite"`
	Diskofferingid        string                                                `json:"diskofferingid"`
	Diskofferingname      string                                                `json:"diskofferingname"`
	Displayname           string                                                `json:"displayname"`
	Displayvm             bool                                                  `json:"displayvm"`
	Domain                string                                                `json:"domain"`
	Domainid              string                                                `json:"domainid"`
	Forvirtualnetwork     bool                                                  `json:"forvirtualnetwork"`
	Group                 string                                                `json:"group"`
	Groupid               string                                                `json:"groupid"`
	Guestosid             string                                                `json:"guestosid"`
	Haenable              bool                                                  `json:"haenable"`
	Hasannotations        bool                                                  `json:"hasannotations"`
	Hostid                string                                                `json:"hostid"`
	Hostname              string                                                `json:"hostname"`
	Hypervisor            string                                                `json:"hypervisor"`
	Icon                  string                                                `json:"icon"`
	Id                    string                                                `json:"id"`
	Instancename          string                                                `json:"instancename"`
	Isdynamicallyscalable bool                                                  `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                                `json:"isodisplaytext"`
	Isoid                 string                                                `json:"isoid"`
	Isoname               string                                                `json:"isoname"`
	JobID                 string                                                `json:"jobid"`
	Jobstatus             int                                                   `json:"jobstatus"`
	Keypair               string                                                `json:"keypair"`
	Lastupdated           string                                                `json:"lastupdated"`
	Memory                int                                                   `json:"memory"`
	Memoryintfreekbs      int64                                                 `json:"memoryintfreekbs"`
	Memorykbs             int64                                                 `json:"memorykbs"`
	Memorytargetkbs       int64                                                 `json:"memorytargetkbs"`
	Name                  string                                                `json:"name"`
	Networkkbsread        int64                                                 `json:"networkkbsread"`
	Networkkbswrite       int64                                                 `json:"networkkbswrite"`
	Nic                   []Nic                                                 `json:"nic"`
	Osdisplayname         string                                                `json:"osdisplayname"`
	Ostypeid              string                                                `json:"ostypeid"`
	Password              string                                                `json:"password"`
	Passwordenabled       bool                                                  `json:"passwordenabled"`
	Pooltype              string                                                `json:"pooltype"`
	Project               string                                                `json:"project"`
	Projectid             string                                                `json:"projectid"`
	Publicip              string                                                `json:"publicip"`
	Publicipid            string                                                `json:"publicipid"`
	Readonlydetails       string                                                `json:"readonlydetails"`
	Receivedbytes         int64                                                 `json:"receivedbytes"`
	Rootdeviceid          int64                                                 `json:"rootdeviceid"`
	Rootdevicetype        string                                                `json:"rootdevicetype"`
	Securitygroup         []ChangeServiceForVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                                 `json:"sentbytes"`
	Serviceofferingid     string                                                `json:"serviceofferingid"`
	Serviceofferingname   string                                                `json:"serviceofferingname"`
	Servicestate          string                                                `json:"servicestate"`
	State                 string                                                `json:"state"`
	Tags                  []Tags                                                `json:"tags"`
	Templatedisplaytext   string                                                `json:"templatedisplaytext"`
	Templateid            string                                                `json:"templateid"`
	Templatename          string                                                `json:"templatename"`
	Userid                string                                                `json:"userid"`
	Username              string                                                `json:"username"`
	Vgpu                  string                                                `json:"vgpu"`
	Zoneid                string                                                `json:"zoneid"`
	Zonename              string                                                `json:"zonename"`
}

type ChangeServiceForVirtualMachineResponseSecuritygroup struct {
	Account             string                                                    `json:"account"`
	Description         string                                                    `json:"description"`
	Domain              string                                                    `json:"domain"`
	Domainid            string                                                    `json:"domainid"`
	Egressrule          []ChangeServiceForVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                    `json:"id"`
	Ingressrule         []ChangeServiceForVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                    `json:"name"`
	Project             string                                                    `json:"project"`
	Projectid           string                                                    `json:"projectid"`
	Tags                []Tags                                                    `json:"tags"`
	Virtualmachinecount int                                                       `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                             `json:"virtualmachineids"`
}

type ChangeServiceForVirtualMachineResponseSecuritygroupRule struct {
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

type ChangeServiceForVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *ChangeServiceForVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias ChangeServiceForVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type CleanVMReservationsParams struct {
	P map[string]interface{}
}

func (P *CleanVMReservationsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	return u
}

// You should always use this function to get a new CleanVMReservationsParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewCleanVMReservationsParams() *CleanVMReservationsParams {
	P := &CleanVMReservationsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Cleanups VM reservations in the database.
func (s *VirtualMachineService) CleanVMReservations(p *CleanVMReservationsParams) (*CleanVMReservationsResponse, error) {
	resp, err := s.cs.newRequest("cleanVMReservations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CleanVMReservationsResponse
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

type CleanVMReservationsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeployVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *DeployVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["affinitygroupids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("affinitygroupids", vv)
	}
	if v, found := P.P["affinitygroupnames"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("affinitygroupnames", vv)
	}
	if v, found := P.P["bootintosetup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootintosetup", vv)
	}
	if v, found := P.P["bootmode"]; found {
		u.Set("bootmode", v.(string))
	}
	if v, found := P.P["boottype"]; found {
		u.Set("boottype", v.(string))
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["copyimagetags"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("copyimagetags", vv)
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["datadiskofferinglist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("datadiskofferinglist[%d].key", i), k)
			u.Set(fmt.Sprintf("datadiskofferinglist[%d].value", i), m[k])
		}
	}
	if v, found := P.P["deploymentplanner"]; found {
		u.Set("deploymentplanner", v.(string))
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := P.P["dhcpoptionsnetworklist"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("dhcpoptionsnetworklist[%d].%s", i, key), val)
			}
		}
	}
	if v, found := P.P["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := P.P["displayname"]; found {
		u.Set("displayname", v.(string))
	}
	if v, found := P.P["displayvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvm", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["dynamicscalingenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("dynamicscalingenabled", vv)
	}
	if v, found := P.P["extraconfig"]; found {
		u.Set("extraconfig", v.(string))
	}
	if v, found := P.P["group"]; found {
		u.Set("group", v.(string))
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["ip6address"]; found {
		u.Set("ip6address", v.(string))
	}
	if v, found := P.P["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := P.P["iptonetworklist"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("iptonetworklist[%d].%s", i, key), val)
			}
		}
	}
	if v, found := P.P["keyboard"]; found {
		u.Set("keyboard", v.(string))
	}
	if v, found := P.P["keypair"]; found {
		u.Set("keypair", v.(string))
	}
	if v, found := P.P["macaddress"]; found {
		u.Set("macaddress", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networkids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("networkids", vv)
	}
	if v, found := P.P["nicnetworklist"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("nicnetworklist[%d].%s", i, key), val)
			}
		}
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["properties"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("properties[%d].key", i), k)
			u.Set(fmt.Sprintf("properties[%d].value", i), m[k])
		}
	}
	if v, found := P.P["rootdisksize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("rootdisksize", vv)
	}
	if v, found := P.P["securitygroupids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("securitygroupids", vv)
	}
	if v, found := P.P["securitygroupnames"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("securitygroupnames", vv)
	}
	if v, found := P.P["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := P.P["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	if v, found := P.P["startvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("startvm", vv)
	}
	if v, found := P.P["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := P.P["userdata"]; found {
		u.Set("userdata", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *DeployVirtualMachineParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DeployVirtualMachineParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetAffinitygroupids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["affinitygroupids"] = v
}

func (P *DeployVirtualMachineParams) GetAffinitygroupids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["affinitygroupids"].([]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetAffinitygroupnames(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["affinitygroupnames"] = v
}

func (P *DeployVirtualMachineParams) GetAffinitygroupnames() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["affinitygroupnames"].([]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetBootintosetup(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bootintosetup"] = v
}

func (P *DeployVirtualMachineParams) GetBootintosetup() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bootintosetup"].(bool)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetBootmode(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bootmode"] = v
}

func (P *DeployVirtualMachineParams) GetBootmode() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bootmode"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetBoottype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["boottype"] = v
}

func (P *DeployVirtualMachineParams) GetBoottype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["boottype"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *DeployVirtualMachineParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetCopyimagetags(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["copyimagetags"] = v
}

func (P *DeployVirtualMachineParams) GetCopyimagetags() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["copyimagetags"].(bool)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *DeployVirtualMachineParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetDatadiskofferinglist(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["datadiskofferinglist"] = v
}

func (P *DeployVirtualMachineParams) GetDatadiskofferinglist() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["datadiskofferinglist"].(map[string]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetDeploymentplanner(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["deploymentplanner"] = v
}

func (P *DeployVirtualMachineParams) GetDeploymentplanner() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["deploymentplanner"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *DeployVirtualMachineParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetDhcpoptionsnetworklist(v []map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dhcpoptionsnetworklist"] = v
}

func (P *DeployVirtualMachineParams) GetDhcpoptionsnetworklist() ([]map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dhcpoptionsnetworklist"].([]map[string]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) AddDhcpoptionsnetworklist(item map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	val, found := P.P["dhcpoptionsnetworklist"]
	if !found {
		P.P["dhcpoptionsnetworklist"] = []map[string]string{}
		val = P.P["dhcpoptionsnetworklist"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	P.P["dhcpoptionsnetworklist"] = l
}

func (P *DeployVirtualMachineParams) SetDiskofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["diskofferingid"] = v
}

func (P *DeployVirtualMachineParams) GetDiskofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["diskofferingid"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetDisplayname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayname"] = v
}

func (P *DeployVirtualMachineParams) GetDisplayname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayname"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetDisplayvm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayvm"] = v
}

func (P *DeployVirtualMachineParams) GetDisplayvm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayvm"].(bool)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *DeployVirtualMachineParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetDynamicscalingenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dynamicscalingenabled"] = v
}

func (P *DeployVirtualMachineParams) GetDynamicscalingenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dynamicscalingenabled"].(bool)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetExtraconfig(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["extraconfig"] = v
}

func (P *DeployVirtualMachineParams) GetExtraconfig() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["extraconfig"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetGroup(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["group"] = v
}

func (P *DeployVirtualMachineParams) GetGroup() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["group"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *DeployVirtualMachineParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *DeployVirtualMachineParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetIp6address(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ip6address"] = v
}

func (P *DeployVirtualMachineParams) GetIp6address() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ip6address"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetIpaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddress"] = v
}

func (P *DeployVirtualMachineParams) GetIpaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddress"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetIptonetworklist(v []map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iptonetworklist"] = v
}

func (P *DeployVirtualMachineParams) GetIptonetworklist() ([]map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iptonetworklist"].([]map[string]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) AddIptonetworklist(item map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	val, found := P.P["iptonetworklist"]
	if !found {
		P.P["iptonetworklist"] = []map[string]string{}
		val = P.P["iptonetworklist"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	P.P["iptonetworklist"] = l
}

func (P *DeployVirtualMachineParams) SetKeyboard(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyboard"] = v
}

func (P *DeployVirtualMachineParams) GetKeyboard() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyboard"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetKeypair(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keypair"] = v
}

func (P *DeployVirtualMachineParams) GetKeypair() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keypair"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetMacaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["macaddress"] = v
}

func (P *DeployVirtualMachineParams) GetMacaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["macaddress"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *DeployVirtualMachineParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetNetworkids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkids"] = v
}

func (P *DeployVirtualMachineParams) GetNetworkids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkids"].([]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetNicnetworklist(v []map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nicnetworklist"] = v
}

func (P *DeployVirtualMachineParams) GetNicnetworklist() ([]map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nicnetworklist"].([]map[string]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) AddNicnetworklist(item map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	val, found := P.P["nicnetworklist"]
	if !found {
		P.P["nicnetworklist"] = []map[string]string{}
		val = P.P["nicnetworklist"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	P.P["nicnetworklist"] = l
}

func (P *DeployVirtualMachineParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *DeployVirtualMachineParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *DeployVirtualMachineParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetProperties(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["properties"] = v
}

func (P *DeployVirtualMachineParams) GetProperties() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["properties"].(map[string]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetRootdisksize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["rootdisksize"] = v
}

func (P *DeployVirtualMachineParams) GetRootdisksize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["rootdisksize"].(int64)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetSecuritygroupids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupids"] = v
}

func (P *DeployVirtualMachineParams) GetSecuritygroupids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupids"].([]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetSecuritygroupnames(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupnames"] = v
}

func (P *DeployVirtualMachineParams) GetSecuritygroupnames() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupnames"].([]string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *DeployVirtualMachineParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetSize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["size"] = v
}

func (P *DeployVirtualMachineParams) GetSize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["size"].(int64)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetStartvm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startvm"] = v
}

func (P *DeployVirtualMachineParams) GetStartvm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startvm"].(bool)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetTemplateid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templateid"] = v
}

func (P *DeployVirtualMachineParams) GetTemplateid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templateid"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetUserdata(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userdata"] = v
}

func (P *DeployVirtualMachineParams) GetUserdata() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userdata"].(string)
	return value, ok
}

func (P *DeployVirtualMachineParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *DeployVirtualMachineParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeployVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewDeployVirtualMachineParams(serviceofferingid string, templateid string, zoneid string) *DeployVirtualMachineParams {
	P := &DeployVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["serviceofferingid"] = serviceofferingid
	P.P["templateid"] = templateid
	P.P["zoneid"] = zoneid
	return P
}

// Creates and automatically starts a virtual machine based on a service offering, disk offering, and template.
func (s *VirtualMachineService) DeployVirtualMachine(p *DeployVirtualMachineParams) (*DeployVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("deployVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeployVirtualMachineResponse
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

type DeployVirtualMachineResponse struct {
	Account               string                                      `json:"account"`
	Affinitygroup         []DeployVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                      `json:"backupofferingid"`
	Backupofferingname    string                                      `json:"backupofferingname"`
	Bootmode              string                                      `json:"bootmode"`
	Boottype              string                                      `json:"boottype"`
	Cpunumber             int                                         `json:"cpunumber"`
	Cpuspeed              int                                         `json:"cpuspeed"`
	Cpuused               string                                      `json:"cpuused"`
	Created               string                                      `json:"created"`
	Details               map[string]string                           `json:"details"`
	Diskioread            int64                                       `json:"diskioread"`
	Diskiowrite           int64                                       `json:"diskiowrite"`
	Diskkbsread           int64                                       `json:"diskkbsread"`
	Diskkbswrite          int64                                       `json:"diskkbswrite"`
	Diskofferingid        string                                      `json:"diskofferingid"`
	Diskofferingname      string                                      `json:"diskofferingname"`
	Displayname           string                                      `json:"displayname"`
	Displayvm             bool                                        `json:"displayvm"`
	Domain                string                                      `json:"domain"`
	Domainid              string                                      `json:"domainid"`
	Forvirtualnetwork     bool                                        `json:"forvirtualnetwork"`
	Group                 string                                      `json:"group"`
	Groupid               string                                      `json:"groupid"`
	Guestosid             string                                      `json:"guestosid"`
	Haenable              bool                                        `json:"haenable"`
	Hasannotations        bool                                        `json:"hasannotations"`
	Hostid                string                                      `json:"hostid"`
	Hostname              string                                      `json:"hostname"`
	Hypervisor            string                                      `json:"hypervisor"`
	Icon                  string                                      `json:"icon"`
	Id                    string                                      `json:"id"`
	Instancename          string                                      `json:"instancename"`
	Isdynamicallyscalable bool                                        `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                      `json:"isodisplaytext"`
	Isoid                 string                                      `json:"isoid"`
	Isoname               string                                      `json:"isoname"`
	JobID                 string                                      `json:"jobid"`
	Jobstatus             int                                         `json:"jobstatus"`
	Keypair               string                                      `json:"keypair"`
	Lastupdated           string                                      `json:"lastupdated"`
	Memory                int                                         `json:"memory"`
	Memoryintfreekbs      int64                                       `json:"memoryintfreekbs"`
	Memorykbs             int64                                       `json:"memorykbs"`
	Memorytargetkbs       int64                                       `json:"memorytargetkbs"`
	Name                  string                                      `json:"name"`
	Networkkbsread        int64                                       `json:"networkkbsread"`
	Networkkbswrite       int64                                       `json:"networkkbswrite"`
	Nic                   []Nic                                       `json:"nic"`
	Osdisplayname         string                                      `json:"osdisplayname"`
	Ostypeid              string                                      `json:"ostypeid"`
	Password              string                                      `json:"password"`
	Passwordenabled       bool                                        `json:"passwordenabled"`
	Pooltype              string                                      `json:"pooltype"`
	Project               string                                      `json:"project"`
	Projectid             string                                      `json:"projectid"`
	Publicip              string                                      `json:"publicip"`
	Publicipid            string                                      `json:"publicipid"`
	Readonlydetails       string                                      `json:"readonlydetails"`
	Receivedbytes         int64                                       `json:"receivedbytes"`
	Rootdeviceid          int64                                       `json:"rootdeviceid"`
	Rootdevicetype        string                                      `json:"rootdevicetype"`
	Securitygroup         []DeployVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                       `json:"sentbytes"`
	Serviceofferingid     string                                      `json:"serviceofferingid"`
	Serviceofferingname   string                                      `json:"serviceofferingname"`
	Servicestate          string                                      `json:"servicestate"`
	State                 string                                      `json:"state"`
	Tags                  []Tags                                      `json:"tags"`
	Templatedisplaytext   string                                      `json:"templatedisplaytext"`
	Templateid            string                                      `json:"templateid"`
	Templatename          string                                      `json:"templatename"`
	Userid                string                                      `json:"userid"`
	Username              string                                      `json:"username"`
	Vgpu                  string                                      `json:"vgpu"`
	Zoneid                string                                      `json:"zoneid"`
	Zonename              string                                      `json:"zonename"`
}

type DeployVirtualMachineResponseSecuritygroup struct {
	Account             string                                          `json:"account"`
	Description         string                                          `json:"description"`
	Domain              string                                          `json:"domain"`
	Domainid            string                                          `json:"domainid"`
	Egressrule          []DeployVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                          `json:"id"`
	Ingressrule         []DeployVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                          `json:"name"`
	Project             string                                          `json:"project"`
	Projectid           string                                          `json:"projectid"`
	Tags                []Tags                                          `json:"tags"`
	Virtualmachinecount int                                             `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                   `json:"virtualmachineids"`
}

type DeployVirtualMachineResponseSecuritygroupRule struct {
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

type DeployVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *DeployVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeployVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DestroyVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *DestroyVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["expunge"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("expunge", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["volumeids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("volumeids", vv)
	}
	return u
}

func (P *DestroyVirtualMachineParams) SetExpunge(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["expunge"] = v
}

func (P *DestroyVirtualMachineParams) GetExpunge() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["expunge"].(bool)
	return value, ok
}

func (P *DestroyVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DestroyVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *DestroyVirtualMachineParams) SetVolumeids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["volumeids"] = v
}

func (P *DestroyVirtualMachineParams) GetVolumeids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["volumeids"].([]string)
	return value, ok
}

// You should always use this function to get a new DestroyVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewDestroyVirtualMachineParams(id string) *DestroyVirtualMachineParams {
	P := &DestroyVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Destroys a virtual machine. Once destroyed, only the administrator can recover it.
func (s *VirtualMachineService) DestroyVirtualMachine(p *DestroyVirtualMachineParams) (*DestroyVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("destroyVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroyVirtualMachineResponse
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

type DestroyVirtualMachineResponse struct {
	Account               string                                       `json:"account"`
	Affinitygroup         []DestroyVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                       `json:"backupofferingid"`
	Backupofferingname    string                                       `json:"backupofferingname"`
	Bootmode              string                                       `json:"bootmode"`
	Boottype              string                                       `json:"boottype"`
	Cpunumber             int                                          `json:"cpunumber"`
	Cpuspeed              int                                          `json:"cpuspeed"`
	Cpuused               string                                       `json:"cpuused"`
	Created               string                                       `json:"created"`
	Details               map[string]string                            `json:"details"`
	Diskioread            int64                                        `json:"diskioread"`
	Diskiowrite           int64                                        `json:"diskiowrite"`
	Diskkbsread           int64                                        `json:"diskkbsread"`
	Diskkbswrite          int64                                        `json:"diskkbswrite"`
	Diskofferingid        string                                       `json:"diskofferingid"`
	Diskofferingname      string                                       `json:"diskofferingname"`
	Displayname           string                                       `json:"displayname"`
	Displayvm             bool                                         `json:"displayvm"`
	Domain                string                                       `json:"domain"`
	Domainid              string                                       `json:"domainid"`
	Forvirtualnetwork     bool                                         `json:"forvirtualnetwork"`
	Group                 string                                       `json:"group"`
	Groupid               string                                       `json:"groupid"`
	Guestosid             string                                       `json:"guestosid"`
	Haenable              bool                                         `json:"haenable"`
	Hasannotations        bool                                         `json:"hasannotations"`
	Hostid                string                                       `json:"hostid"`
	Hostname              string                                       `json:"hostname"`
	Hypervisor            string                                       `json:"hypervisor"`
	Icon                  string                                       `json:"icon"`
	Id                    string                                       `json:"id"`
	Instancename          string                                       `json:"instancename"`
	Isdynamicallyscalable bool                                         `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                       `json:"isodisplaytext"`
	Isoid                 string                                       `json:"isoid"`
	Isoname               string                                       `json:"isoname"`
	JobID                 string                                       `json:"jobid"`
	Jobstatus             int                                          `json:"jobstatus"`
	Keypair               string                                       `json:"keypair"`
	Lastupdated           string                                       `json:"lastupdated"`
	Memory                int                                          `json:"memory"`
	Memoryintfreekbs      int64                                        `json:"memoryintfreekbs"`
	Memorykbs             int64                                        `json:"memorykbs"`
	Memorytargetkbs       int64                                        `json:"memorytargetkbs"`
	Name                  string                                       `json:"name"`
	Networkkbsread        int64                                        `json:"networkkbsread"`
	Networkkbswrite       int64                                        `json:"networkkbswrite"`
	Nic                   []Nic                                        `json:"nic"`
	Osdisplayname         string                                       `json:"osdisplayname"`
	Ostypeid              string                                       `json:"ostypeid"`
	Password              string                                       `json:"password"`
	Passwordenabled       bool                                         `json:"passwordenabled"`
	Pooltype              string                                       `json:"pooltype"`
	Project               string                                       `json:"project"`
	Projectid             string                                       `json:"projectid"`
	Publicip              string                                       `json:"publicip"`
	Publicipid            string                                       `json:"publicipid"`
	Readonlydetails       string                                       `json:"readonlydetails"`
	Receivedbytes         int64                                        `json:"receivedbytes"`
	Rootdeviceid          int64                                        `json:"rootdeviceid"`
	Rootdevicetype        string                                       `json:"rootdevicetype"`
	Securitygroup         []DestroyVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                        `json:"sentbytes"`
	Serviceofferingid     string                                       `json:"serviceofferingid"`
	Serviceofferingname   string                                       `json:"serviceofferingname"`
	Servicestate          string                                       `json:"servicestate"`
	State                 string                                       `json:"state"`
	Tags                  []Tags                                       `json:"tags"`
	Templatedisplaytext   string                                       `json:"templatedisplaytext"`
	Templateid            string                                       `json:"templateid"`
	Templatename          string                                       `json:"templatename"`
	Userid                string                                       `json:"userid"`
	Username              string                                       `json:"username"`
	Vgpu                  string                                       `json:"vgpu"`
	Zoneid                string                                       `json:"zoneid"`
	Zonename              string                                       `json:"zonename"`
}

type DestroyVirtualMachineResponseSecuritygroup struct {
	Account             string                                           `json:"account"`
	Description         string                                           `json:"description"`
	Domain              string                                           `json:"domain"`
	Domainid            string                                           `json:"domainid"`
	Egressrule          []DestroyVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                           `json:"id"`
	Ingressrule         []DestroyVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                           `json:"name"`
	Project             string                                           `json:"project"`
	Projectid           string                                           `json:"projectid"`
	Tags                []Tags                                           `json:"tags"`
	Virtualmachinecount int                                              `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                    `json:"virtualmachineids"`
}

type DestroyVirtualMachineResponseSecuritygroupRule struct {
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

type DestroyVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *DestroyVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias DestroyVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ExpungeVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *ExpungeVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ExpungeVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ExpungeVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ExpungeVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewExpungeVirtualMachineParams(id string) *ExpungeVirtualMachineParams {
	P := &ExpungeVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Expunge a virtual machine. Once expunged, it cannot be recoverd.
func (s *VirtualMachineService) ExpungeVirtualMachine(p *ExpungeVirtualMachineParams) (*ExpungeVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("expungeVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExpungeVirtualMachineResponse
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

type ExpungeVirtualMachineResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type GetVMPasswordParams struct {
	P map[string]interface{}
}

func (P *GetVMPasswordParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *GetVMPasswordParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *GetVMPasswordParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new GetVMPasswordParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewGetVMPasswordParams(id string) *GetVMPasswordParams {
	P := &GetVMPasswordParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Returns an encrypted password for the VM
func (s *VirtualMachineService) GetVMPassword(p *GetVMPasswordParams) (*GetVMPasswordResponse, error) {
	resp, err := s.cs.newRequest("getVMPassword", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetVMPasswordResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetVMPasswordResponse struct {
	Encryptedpassword string `json:"encryptedpassword"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
}

type ListVirtualMachinesParams struct {
	P map[string]interface{}
}

func (P *ListVirtualMachinesParams) toURLValues() url.Values {
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
	if v, found := P.P["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := P.P["displayvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvm", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["forvirtualnetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvirtualnetwork", vv)
	}
	if v, found := P.P["groupid"]; found {
		u.Set("groupid", v.(string))
	}
	if v, found := P.P["haenable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("haenable", vv)
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := P.P["isoid"]; found {
		u.Set("isoid", v.(string))
	}
	if v, found := P.P["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := P.P["keypair"]; found {
		u.Set("keypair", v.(string))
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
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["securitygroupid"]; found {
		u.Set("securitygroupid", v.(string))
	}
	if v, found := P.P["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := P.P["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListVirtualMachinesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListVirtualMachinesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetAffinitygroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["affinitygroupid"] = v
}

func (P *ListVirtualMachinesParams) GetAffinitygroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["affinitygroupid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *ListVirtualMachinesParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetDetails(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ListVirtualMachinesParams) GetDetails() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].([]string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetDisplayvm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayvm"] = v
}

func (P *ListVirtualMachinesParams) GetDisplayvm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayvm"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListVirtualMachinesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetForvirtualnetwork(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forvirtualnetwork"] = v
}

func (P *ListVirtualMachinesParams) GetForvirtualnetwork() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forvirtualnetwork"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetGroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["groupid"] = v
}

func (P *ListVirtualMachinesParams) GetGroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["groupid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetHaenable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["haenable"] = v
}

func (P *ListVirtualMachinesParams) GetHaenable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["haenable"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ListVirtualMachinesParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *ListVirtualMachinesParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListVirtualMachinesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetIds(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ids"] = v
}

func (P *ListVirtualMachinesParams) GetIds() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ids"].([]string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetIsoid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isoid"] = v
}

func (P *ListVirtualMachinesParams) GetIsoid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isoid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListVirtualMachinesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetKeypair(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keypair"] = v
}

func (P *ListVirtualMachinesParams) GetKeypair() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keypair"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListVirtualMachinesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListVirtualMachinesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListVirtualMachinesParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListVirtualMachinesParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListVirtualMachinesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListVirtualMachinesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListVirtualMachinesParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListVirtualMachinesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetSecuritygroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupid"] = v
}

func (P *ListVirtualMachinesParams) GetSecuritygroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *ListVirtualMachinesParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListVirtualMachinesParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListVirtualMachinesParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *ListVirtualMachinesParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListVirtualMachinesParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetTemplateid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templateid"] = v
}

func (P *ListVirtualMachinesParams) GetTemplateid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templateid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *ListVirtualMachinesParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *ListVirtualMachinesParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListVirtualMachinesParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVirtualMachinesParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewListVirtualMachinesParams() *ListVirtualMachinesParams {
	P := &ListVirtualMachinesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachineID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListVirtualMachinesParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVirtualMachines(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VirtualMachines[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VirtualMachines {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachineByName(name string, opts ...OptionFunc) (*VirtualMachine, int, error) {
	id, count, err := s.GetVirtualMachineID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVirtualMachineByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachineByID(id string, opts ...OptionFunc) (*VirtualMachine, int, error) {
	P := &ListVirtualMachinesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVirtualMachines(P)
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
		return l.VirtualMachines[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VirtualMachine UUID: %s!", id)
}

// List the virtual machines owned by the account.
func (s *VirtualMachineService) ListVirtualMachines(p *ListVirtualMachinesParams) (*ListVirtualMachinesResponse, error) {
	resp, err := s.cs.newRequest("listVirtualMachines", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVirtualMachinesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVirtualMachinesResponse struct {
	Count           int               `json:"count"`
	VirtualMachines []*VirtualMachine `json:"virtualmachine"`
}

type VirtualMachine struct {
	Account               string                        `json:"account"`
	Affinitygroup         []VirtualMachineAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                        `json:"backupofferingid"`
	Backupofferingname    string                        `json:"backupofferingname"`
	Bootmode              string                        `json:"bootmode"`
	Boottype              string                        `json:"boottype"`
	Cpunumber             int                           `json:"cpunumber"`
	Cpuspeed              int                           `json:"cpuspeed"`
	Cpuused               string                        `json:"cpuused"`
	Created               string                        `json:"created"`
	Details               map[string]string             `json:"details"`
	Diskioread            int64                         `json:"diskioread"`
	Diskiowrite           int64                         `json:"diskiowrite"`
	Diskkbsread           int64                         `json:"diskkbsread"`
	Diskkbswrite          int64                         `json:"diskkbswrite"`
	Diskofferingid        string                        `json:"diskofferingid"`
	Diskofferingname      string                        `json:"diskofferingname"`
	Displayname           string                        `json:"displayname"`
	Displayvm             bool                          `json:"displayvm"`
	Domain                string                        `json:"domain"`
	Domainid              string                        `json:"domainid"`
	Forvirtualnetwork     bool                          `json:"forvirtualnetwork"`
	Group                 string                        `json:"group"`
	Groupid               string                        `json:"groupid"`
	Guestosid             string                        `json:"guestosid"`
	Haenable              bool                          `json:"haenable"`
	Hasannotations        bool                          `json:"hasannotations"`
	Hostid                string                        `json:"hostid"`
	Hostname              string                        `json:"hostname"`
	Hypervisor            string                        `json:"hypervisor"`
	Icon                  string                        `json:"icon"`
	Id                    string                        `json:"id"`
	Instancename          string                        `json:"instancename"`
	Isdynamicallyscalable bool                          `json:"isdynamicallyscalable"`
	Isodisplaytext        string                        `json:"isodisplaytext"`
	Isoid                 string                        `json:"isoid"`
	Isoname               string                        `json:"isoname"`
	JobID                 string                        `json:"jobid"`
	Jobstatus             int                           `json:"jobstatus"`
	Keypair               string                        `json:"keypair"`
	Lastupdated           string                        `json:"lastupdated"`
	Memory                int                           `json:"memory"`
	Memoryintfreekbs      int64                         `json:"memoryintfreekbs"`
	Memorykbs             int64                         `json:"memorykbs"`
	Memorytargetkbs       int64                         `json:"memorytargetkbs"`
	Name                  string                        `json:"name"`
	Networkkbsread        int64                         `json:"networkkbsread"`
	Networkkbswrite       int64                         `json:"networkkbswrite"`
	Nic                   []Nic                         `json:"nic"`
	Osdisplayname         string                        `json:"osdisplayname"`
	Ostypeid              string                        `json:"ostypeid"`
	Password              string                        `json:"password"`
	Passwordenabled       bool                          `json:"passwordenabled"`
	Pooltype              string                        `json:"pooltype"`
	Project               string                        `json:"project"`
	Projectid             string                        `json:"projectid"`
	Publicip              string                        `json:"publicip"`
	Publicipid            string                        `json:"publicipid"`
	Readonlydetails       string                        `json:"readonlydetails"`
	Receivedbytes         int64                         `json:"receivedbytes"`
	Rootdeviceid          int64                         `json:"rootdeviceid"`
	Rootdevicetype        string                        `json:"rootdevicetype"`
	Securitygroup         []VirtualMachineSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                         `json:"sentbytes"`
	Serviceofferingid     string                        `json:"serviceofferingid"`
	Serviceofferingname   string                        `json:"serviceofferingname"`
	Servicestate          string                        `json:"servicestate"`
	State                 string                        `json:"state"`
	Tags                  []Tags                        `json:"tags"`
	Templatedisplaytext   string                        `json:"templatedisplaytext"`
	Templateid            string                        `json:"templateid"`
	Templatename          string                        `json:"templatename"`
	Userid                string                        `json:"userid"`
	Username              string                        `json:"username"`
	Vgpu                  string                        `json:"vgpu"`
	Zoneid                string                        `json:"zoneid"`
	Zonename              string                        `json:"zonename"`
}

type VirtualMachineSecuritygroup struct {
	Account             string                            `json:"account"`
	Description         string                            `json:"description"`
	Domain              string                            `json:"domain"`
	Domainid            string                            `json:"domainid"`
	Egressrule          []VirtualMachineSecuritygroupRule `json:"egressrule"`
	Id                  string                            `json:"id"`
	Ingressrule         []VirtualMachineSecuritygroupRule `json:"ingressrule"`
	Name                string                            `json:"name"`
	Project             string                            `json:"project"`
	Projectid           string                            `json:"projectid"`
	Tags                []Tags                            `json:"tags"`
	Virtualmachinecount int                               `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                     `json:"virtualmachineids"`
}

type VirtualMachineSecuritygroupRule struct {
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

type VirtualMachineAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *VirtualMachine) UnmarshalJSON(b []byte) error {
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

	type alias VirtualMachine
	return json.Unmarshal(b, (*alias)(r))
}

type ListVirtualMachinesMetricsParams struct {
	P map[string]interface{}
}

func (P *ListVirtualMachinesMetricsParams) toURLValues() url.Values {
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
	if v, found := P.P["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
	}
	if v, found := P.P["displayvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvm", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["forvirtualnetwork"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvirtualnetwork", vv)
	}
	if v, found := P.P["groupid"]; found {
		u.Set("groupid", v.(string))
	}
	if v, found := P.P["haenable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("haenable", vv)
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := P.P["isoid"]; found {
		u.Set("isoid", v.(string))
	}
	if v, found := P.P["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := P.P["keypair"]; found {
		u.Set("keypair", v.(string))
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
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["securitygroupid"]; found {
		u.Set("securitygroupid", v.(string))
	}
	if v, found := P.P["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := P.P["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListVirtualMachinesMetricsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetAffinitygroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["affinitygroupid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetAffinitygroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["affinitygroupid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetDetails(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetDetails() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].([]string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetDisplayvm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayvm"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetDisplayvm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayvm"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetForvirtualnetwork(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forvirtualnetwork"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetForvirtualnetwork() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forvirtualnetwork"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetGroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["groupid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetGroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["groupid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetHaenable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["haenable"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetHaenable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["haenable"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetHypervisor(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisor"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetHypervisor() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisor"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetIds(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ids"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetIds() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ids"].([]string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetIsoid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isoid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetIsoid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isoid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetKeypair(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keypair"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetKeypair() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keypair"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetSecuritygroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetSecuritygroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetTemplateid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templateid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetTemplateid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templateid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

func (P *ListVirtualMachinesMetricsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListVirtualMachinesMetricsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVirtualMachinesMetricsParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewListVirtualMachinesMetricsParams() *ListVirtualMachinesMetricsParams {
	P := &ListVirtualMachinesMetricsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachinesMetricID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListVirtualMachinesMetricsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVirtualMachinesMetrics(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.VirtualMachinesMetrics[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VirtualMachinesMetrics {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachinesMetricByName(name string, opts ...OptionFunc) (*VirtualMachinesMetric, int, error) {
	id, count, err := s.GetVirtualMachinesMetricID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVirtualMachinesMetricByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VirtualMachineService) GetVirtualMachinesMetricByID(id string, opts ...OptionFunc) (*VirtualMachinesMetric, int, error) {
	P := &ListVirtualMachinesMetricsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVirtualMachinesMetrics(P)
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
		return l.VirtualMachinesMetrics[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VirtualMachinesMetric UUID: %s!", id)
}

// Lists VM metrics
func (s *VirtualMachineService) ListVirtualMachinesMetrics(p *ListVirtualMachinesMetricsParams) (*ListVirtualMachinesMetricsResponse, error) {
	resp, err := s.cs.newRequest("listVirtualMachinesMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVirtualMachinesMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVirtualMachinesMetricsResponse struct {
	Count                  int                      `json:"count"`
	VirtualMachinesMetrics []*VirtualMachinesMetric `json:"virtualmachine"`
}

type VirtualMachinesMetric struct {
	Account               string                               `json:"account"`
	Affinitygroup         []VirtualMachinesMetricAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                               `json:"backupofferingid"`
	Backupofferingname    string                               `json:"backupofferingname"`
	Bootmode              string                               `json:"bootmode"`
	Boottype              string                               `json:"boottype"`
	Cpunumber             int                                  `json:"cpunumber"`
	Cpuspeed              int                                  `json:"cpuspeed"`
	Cputotal              string                               `json:"cputotal"`
	Cpuused               string                               `json:"cpuused"`
	Created               string                               `json:"created"`
	Details               map[string]string                    `json:"details"`
	Diskiopstotal         int64                                `json:"diskiopstotal"`
	Diskioread            int64                                `json:"diskioread"`
	Diskiowrite           int64                                `json:"diskiowrite"`
	Diskkbsread           int64                                `json:"diskkbsread"`
	Diskkbswrite          int64                                `json:"diskkbswrite"`
	Diskofferingid        string                               `json:"diskofferingid"`
	Diskofferingname      string                               `json:"diskofferingname"`
	Diskread              string                               `json:"diskread"`
	Diskwrite             string                               `json:"diskwrite"`
	Displayname           string                               `json:"displayname"`
	Displayvm             bool                                 `json:"displayvm"`
	Domain                string                               `json:"domain"`
	Domainid              string                               `json:"domainid"`
	Forvirtualnetwork     bool                                 `json:"forvirtualnetwork"`
	Group                 string                               `json:"group"`
	Groupid               string                               `json:"groupid"`
	Guestosid             string                               `json:"guestosid"`
	Haenable              bool                                 `json:"haenable"`
	Hasannotations        bool                                 `json:"hasannotations"`
	Hostid                string                               `json:"hostid"`
	Hostname              string                               `json:"hostname"`
	Hypervisor            string                               `json:"hypervisor"`
	Icon                  string                               `json:"icon"`
	Id                    string                               `json:"id"`
	Instancename          string                               `json:"instancename"`
	Ipaddress             string                               `json:"ipaddress"`
	Isdynamicallyscalable bool                                 `json:"isdynamicallyscalable"`
	Isodisplaytext        string                               `json:"isodisplaytext"`
	Isoid                 string                               `json:"isoid"`
	Isoname               string                               `json:"isoname"`
	JobID                 string                               `json:"jobid"`
	Jobstatus             int                                  `json:"jobstatus"`
	Keypair               string                               `json:"keypair"`
	Lastupdated           string                               `json:"lastupdated"`
	Memory                int                                  `json:"memory"`
	Memoryintfreekbs      int64                                `json:"memoryintfreekbs"`
	Memorykbs             int64                                `json:"memorykbs"`
	Memorytargetkbs       int64                                `json:"memorytargetkbs"`
	Memorytotal           string                               `json:"memorytotal"`
	Name                  string                               `json:"name"`
	Networkkbsread        int64                                `json:"networkkbsread"`
	Networkkbswrite       int64                                `json:"networkkbswrite"`
	Networkread           string                               `json:"networkread"`
	Networkwrite          string                               `json:"networkwrite"`
	Nic                   []Nic                                `json:"nic"`
	Osdisplayname         string                               `json:"osdisplayname"`
	Ostypeid              string                               `json:"ostypeid"`
	Password              string                               `json:"password"`
	Passwordenabled       bool                                 `json:"passwordenabled"`
	Pooltype              string                               `json:"pooltype"`
	Project               string                               `json:"project"`
	Projectid             string                               `json:"projectid"`
	Publicip              string                               `json:"publicip"`
	Publicipid            string                               `json:"publicipid"`
	Readonlydetails       string                               `json:"readonlydetails"`
	Receivedbytes         int64                                `json:"receivedbytes"`
	Rootdeviceid          int64                                `json:"rootdeviceid"`
	Rootdevicetype        string                               `json:"rootdevicetype"`
	Securitygroup         []VirtualMachinesMetricSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                `json:"sentbytes"`
	Serviceofferingid     string                               `json:"serviceofferingid"`
	Serviceofferingname   string                               `json:"serviceofferingname"`
	Servicestate          string                               `json:"servicestate"`
	State                 string                               `json:"state"`
	Tags                  []Tags                               `json:"tags"`
	Templatedisplaytext   string                               `json:"templatedisplaytext"`
	Templateid            string                               `json:"templateid"`
	Templatename          string                               `json:"templatename"`
	Userid                string                               `json:"userid"`
	Username              string                               `json:"username"`
	Vgpu                  string                               `json:"vgpu"`
	Zoneid                string                               `json:"zoneid"`
	Zonename              string                               `json:"zonename"`
}

type VirtualMachinesMetricSecuritygroup struct {
	Account             string                                   `json:"account"`
	Description         string                                   `json:"description"`
	Domain              string                                   `json:"domain"`
	Domainid            string                                   `json:"domainid"`
	Egressrule          []VirtualMachinesMetricSecuritygroupRule `json:"egressrule"`
	Id                  string                                   `json:"id"`
	Ingressrule         []VirtualMachinesMetricSecuritygroupRule `json:"ingressrule"`
	Name                string                                   `json:"name"`
	Project             string                                   `json:"project"`
	Projectid           string                                   `json:"projectid"`
	Tags                []Tags                                   `json:"tags"`
	Virtualmachinecount int                                      `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                            `json:"virtualmachineids"`
}

type VirtualMachinesMetricSecuritygroupRule struct {
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

type VirtualMachinesMetricAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *VirtualMachinesMetric) UnmarshalJSON(b []byte) error {
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

	type alias VirtualMachinesMetric
	return json.Unmarshal(b, (*alias)(r))
}

type MigrateVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *MigrateVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["autoselect"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("autoselect", vv)
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *MigrateVirtualMachineParams) SetAutoselect(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["autoselect"] = v
}

func (P *MigrateVirtualMachineParams) GetAutoselect() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["autoselect"].(bool)
	return value, ok
}

func (P *MigrateVirtualMachineParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *MigrateVirtualMachineParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *MigrateVirtualMachineParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *MigrateVirtualMachineParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

func (P *MigrateVirtualMachineParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *MigrateVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new MigrateVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewMigrateVirtualMachineParams(virtualmachineid string) *MigrateVirtualMachineParams {
	P := &MigrateVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Attempts Migration of a VM to a different host or Root volume of the vm to a different storage pool
func (s *VirtualMachineService) MigrateVirtualMachine(p *MigrateVirtualMachineParams) (*MigrateVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("migrateVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateVirtualMachineResponse
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

type MigrateVirtualMachineResponse struct {
	Account               string                                       `json:"account"`
	Affinitygroup         []MigrateVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                       `json:"backupofferingid"`
	Backupofferingname    string                                       `json:"backupofferingname"`
	Bootmode              string                                       `json:"bootmode"`
	Boottype              string                                       `json:"boottype"`
	Cpunumber             int                                          `json:"cpunumber"`
	Cpuspeed              int                                          `json:"cpuspeed"`
	Cpuused               string                                       `json:"cpuused"`
	Created               string                                       `json:"created"`
	Details               map[string]string                            `json:"details"`
	Diskioread            int64                                        `json:"diskioread"`
	Diskiowrite           int64                                        `json:"diskiowrite"`
	Diskkbsread           int64                                        `json:"diskkbsread"`
	Diskkbswrite          int64                                        `json:"diskkbswrite"`
	Diskofferingid        string                                       `json:"diskofferingid"`
	Diskofferingname      string                                       `json:"diskofferingname"`
	Displayname           string                                       `json:"displayname"`
	Displayvm             bool                                         `json:"displayvm"`
	Domain                string                                       `json:"domain"`
	Domainid              string                                       `json:"domainid"`
	Forvirtualnetwork     bool                                         `json:"forvirtualnetwork"`
	Group                 string                                       `json:"group"`
	Groupid               string                                       `json:"groupid"`
	Guestosid             string                                       `json:"guestosid"`
	Haenable              bool                                         `json:"haenable"`
	Hasannotations        bool                                         `json:"hasannotations"`
	Hostid                string                                       `json:"hostid"`
	Hostname              string                                       `json:"hostname"`
	Hypervisor            string                                       `json:"hypervisor"`
	Icon                  string                                       `json:"icon"`
	Id                    string                                       `json:"id"`
	Instancename          string                                       `json:"instancename"`
	Isdynamicallyscalable bool                                         `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                       `json:"isodisplaytext"`
	Isoid                 string                                       `json:"isoid"`
	Isoname               string                                       `json:"isoname"`
	JobID                 string                                       `json:"jobid"`
	Jobstatus             int                                          `json:"jobstatus"`
	Keypair               string                                       `json:"keypair"`
	Lastupdated           string                                       `json:"lastupdated"`
	Memory                int                                          `json:"memory"`
	Memoryintfreekbs      int64                                        `json:"memoryintfreekbs"`
	Memorykbs             int64                                        `json:"memorykbs"`
	Memorytargetkbs       int64                                        `json:"memorytargetkbs"`
	Name                  string                                       `json:"name"`
	Networkkbsread        int64                                        `json:"networkkbsread"`
	Networkkbswrite       int64                                        `json:"networkkbswrite"`
	Nic                   []Nic                                        `json:"nic"`
	Osdisplayname         string                                       `json:"osdisplayname"`
	Ostypeid              string                                       `json:"ostypeid"`
	Password              string                                       `json:"password"`
	Passwordenabled       bool                                         `json:"passwordenabled"`
	Pooltype              string                                       `json:"pooltype"`
	Project               string                                       `json:"project"`
	Projectid             string                                       `json:"projectid"`
	Publicip              string                                       `json:"publicip"`
	Publicipid            string                                       `json:"publicipid"`
	Readonlydetails       string                                       `json:"readonlydetails"`
	Receivedbytes         int64                                        `json:"receivedbytes"`
	Rootdeviceid          int64                                        `json:"rootdeviceid"`
	Rootdevicetype        string                                       `json:"rootdevicetype"`
	Securitygroup         []MigrateVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                        `json:"sentbytes"`
	Serviceofferingid     string                                       `json:"serviceofferingid"`
	Serviceofferingname   string                                       `json:"serviceofferingname"`
	Servicestate          string                                       `json:"servicestate"`
	State                 string                                       `json:"state"`
	Tags                  []Tags                                       `json:"tags"`
	Templatedisplaytext   string                                       `json:"templatedisplaytext"`
	Templateid            string                                       `json:"templateid"`
	Templatename          string                                       `json:"templatename"`
	Userid                string                                       `json:"userid"`
	Username              string                                       `json:"username"`
	Vgpu                  string                                       `json:"vgpu"`
	Zoneid                string                                       `json:"zoneid"`
	Zonename              string                                       `json:"zonename"`
}

type MigrateVirtualMachineResponseSecuritygroup struct {
	Account             string                                           `json:"account"`
	Description         string                                           `json:"description"`
	Domain              string                                           `json:"domain"`
	Domainid            string                                           `json:"domainid"`
	Egressrule          []MigrateVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                           `json:"id"`
	Ingressrule         []MigrateVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                           `json:"name"`
	Project             string                                           `json:"project"`
	Projectid           string                                           `json:"projectid"`
	Tags                []Tags                                           `json:"tags"`
	Virtualmachinecount int                                              `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                    `json:"virtualmachineids"`
}

type MigrateVirtualMachineResponseSecuritygroupRule struct {
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

type MigrateVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *MigrateVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias MigrateVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type MigrateVirtualMachineWithVolumeParams struct {
	P map[string]interface{}
}

func (P *MigrateVirtualMachineWithVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["migrateto"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("migrateto[%d].%s", i, key), val)
			}
		}
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *MigrateVirtualMachineWithVolumeParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *MigrateVirtualMachineWithVolumeParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *MigrateVirtualMachineWithVolumeParams) SetMigrateto(v []map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["migrateto"] = v
}

func (P *MigrateVirtualMachineWithVolumeParams) GetMigrateto() ([]map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["migrateto"].([]map[string]string)
	return value, ok
}

func (P *MigrateVirtualMachineWithVolumeParams) AddMigrateto(item map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	val, found := P.P["migrateto"]
	if !found {
		P.P["migrateto"] = []map[string]string{}
		val = P.P["migrateto"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	P.P["migrateto"] = l
}

func (P *MigrateVirtualMachineWithVolumeParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *MigrateVirtualMachineWithVolumeParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new MigrateVirtualMachineWithVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewMigrateVirtualMachineWithVolumeParams(virtualmachineid string) *MigrateVirtualMachineWithVolumeParams {
	P := &MigrateVirtualMachineWithVolumeParams{}
	P.P = make(map[string]interface{})
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Attempts Migration of a VM with its volumes to a different host
func (s *VirtualMachineService) MigrateVirtualMachineWithVolume(p *MigrateVirtualMachineWithVolumeParams) (*MigrateVirtualMachineWithVolumeResponse, error) {
	resp, err := s.cs.newRequest("migrateVirtualMachineWithVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateVirtualMachineWithVolumeResponse
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

type MigrateVirtualMachineWithVolumeResponse struct {
	Account               string                                                 `json:"account"`
	Affinitygroup         []MigrateVirtualMachineWithVolumeResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                                 `json:"backupofferingid"`
	Backupofferingname    string                                                 `json:"backupofferingname"`
	Bootmode              string                                                 `json:"bootmode"`
	Boottype              string                                                 `json:"boottype"`
	Cpunumber             int                                                    `json:"cpunumber"`
	Cpuspeed              int                                                    `json:"cpuspeed"`
	Cpuused               string                                                 `json:"cpuused"`
	Created               string                                                 `json:"created"`
	Details               map[string]string                                      `json:"details"`
	Diskioread            int64                                                  `json:"diskioread"`
	Diskiowrite           int64                                                  `json:"diskiowrite"`
	Diskkbsread           int64                                                  `json:"diskkbsread"`
	Diskkbswrite          int64                                                  `json:"diskkbswrite"`
	Diskofferingid        string                                                 `json:"diskofferingid"`
	Diskofferingname      string                                                 `json:"diskofferingname"`
	Displayname           string                                                 `json:"displayname"`
	Displayvm             bool                                                   `json:"displayvm"`
	Domain                string                                                 `json:"domain"`
	Domainid              string                                                 `json:"domainid"`
	Forvirtualnetwork     bool                                                   `json:"forvirtualnetwork"`
	Group                 string                                                 `json:"group"`
	Groupid               string                                                 `json:"groupid"`
	Guestosid             string                                                 `json:"guestosid"`
	Haenable              bool                                                   `json:"haenable"`
	Hasannotations        bool                                                   `json:"hasannotations"`
	Hostid                string                                                 `json:"hostid"`
	Hostname              string                                                 `json:"hostname"`
	Hypervisor            string                                                 `json:"hypervisor"`
	Icon                  string                                                 `json:"icon"`
	Id                    string                                                 `json:"id"`
	Instancename          string                                                 `json:"instancename"`
	Isdynamicallyscalable bool                                                   `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                                 `json:"isodisplaytext"`
	Isoid                 string                                                 `json:"isoid"`
	Isoname               string                                                 `json:"isoname"`
	JobID                 string                                                 `json:"jobid"`
	Jobstatus             int                                                    `json:"jobstatus"`
	Keypair               string                                                 `json:"keypair"`
	Lastupdated           string                                                 `json:"lastupdated"`
	Memory                int                                                    `json:"memory"`
	Memoryintfreekbs      int64                                                  `json:"memoryintfreekbs"`
	Memorykbs             int64                                                  `json:"memorykbs"`
	Memorytargetkbs       int64                                                  `json:"memorytargetkbs"`
	Name                  string                                                 `json:"name"`
	Networkkbsread        int64                                                  `json:"networkkbsread"`
	Networkkbswrite       int64                                                  `json:"networkkbswrite"`
	Nic                   []Nic                                                  `json:"nic"`
	Osdisplayname         string                                                 `json:"osdisplayname"`
	Ostypeid              string                                                 `json:"ostypeid"`
	Password              string                                                 `json:"password"`
	Passwordenabled       bool                                                   `json:"passwordenabled"`
	Pooltype              string                                                 `json:"pooltype"`
	Project               string                                                 `json:"project"`
	Projectid             string                                                 `json:"projectid"`
	Publicip              string                                                 `json:"publicip"`
	Publicipid            string                                                 `json:"publicipid"`
	Readonlydetails       string                                                 `json:"readonlydetails"`
	Receivedbytes         int64                                                  `json:"receivedbytes"`
	Rootdeviceid          int64                                                  `json:"rootdeviceid"`
	Rootdevicetype        string                                                 `json:"rootdevicetype"`
	Securitygroup         []MigrateVirtualMachineWithVolumeResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                                  `json:"sentbytes"`
	Serviceofferingid     string                                                 `json:"serviceofferingid"`
	Serviceofferingname   string                                                 `json:"serviceofferingname"`
	Servicestate          string                                                 `json:"servicestate"`
	State                 string                                                 `json:"state"`
	Tags                  []Tags                                                 `json:"tags"`
	Templatedisplaytext   string                                                 `json:"templatedisplaytext"`
	Templateid            string                                                 `json:"templateid"`
	Templatename          string                                                 `json:"templatename"`
	Userid                string                                                 `json:"userid"`
	Username              string                                                 `json:"username"`
	Vgpu                  string                                                 `json:"vgpu"`
	Zoneid                string                                                 `json:"zoneid"`
	Zonename              string                                                 `json:"zonename"`
}

type MigrateVirtualMachineWithVolumeResponseSecuritygroup struct {
	Account             string                                                     `json:"account"`
	Description         string                                                     `json:"description"`
	Domain              string                                                     `json:"domain"`
	Domainid            string                                                     `json:"domainid"`
	Egressrule          []MigrateVirtualMachineWithVolumeResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                     `json:"id"`
	Ingressrule         []MigrateVirtualMachineWithVolumeResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                     `json:"name"`
	Project             string                                                     `json:"project"`
	Projectid           string                                                     `json:"projectid"`
	Tags                []Tags                                                     `json:"tags"`
	Virtualmachinecount int                                                        `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                              `json:"virtualmachineids"`
}

type MigrateVirtualMachineWithVolumeResponseSecuritygroupRule struct {
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

type MigrateVirtualMachineWithVolumeResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *MigrateVirtualMachineWithVolumeResponse) UnmarshalJSON(b []byte) error {
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

	type alias MigrateVirtualMachineWithVolumeResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RebootVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *RebootVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["bootintosetup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootintosetup", vv)
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

func (P *RebootVirtualMachineParams) SetBootintosetup(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bootintosetup"] = v
}

func (P *RebootVirtualMachineParams) GetBootintosetup() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bootintosetup"].(bool)
	return value, ok
}

func (P *RebootVirtualMachineParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *RebootVirtualMachineParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *RebootVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RebootVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RebootVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRebootVirtualMachineParams(id string) *RebootVirtualMachineParams {
	P := &RebootVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Reboots a virtual machine.
func (s *VirtualMachineService) RebootVirtualMachine(p *RebootVirtualMachineParams) (*RebootVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("rebootVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootVirtualMachineResponse
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

type RebootVirtualMachineResponse struct {
	Account               string                                      `json:"account"`
	Affinitygroup         []RebootVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                      `json:"backupofferingid"`
	Backupofferingname    string                                      `json:"backupofferingname"`
	Bootmode              string                                      `json:"bootmode"`
	Boottype              string                                      `json:"boottype"`
	Cpunumber             int                                         `json:"cpunumber"`
	Cpuspeed              int                                         `json:"cpuspeed"`
	Cpuused               string                                      `json:"cpuused"`
	Created               string                                      `json:"created"`
	Details               map[string]string                           `json:"details"`
	Diskioread            int64                                       `json:"diskioread"`
	Diskiowrite           int64                                       `json:"diskiowrite"`
	Diskkbsread           int64                                       `json:"diskkbsread"`
	Diskkbswrite          int64                                       `json:"diskkbswrite"`
	Diskofferingid        string                                      `json:"diskofferingid"`
	Diskofferingname      string                                      `json:"diskofferingname"`
	Displayname           string                                      `json:"displayname"`
	Displayvm             bool                                        `json:"displayvm"`
	Domain                string                                      `json:"domain"`
	Domainid              string                                      `json:"domainid"`
	Forvirtualnetwork     bool                                        `json:"forvirtualnetwork"`
	Group                 string                                      `json:"group"`
	Groupid               string                                      `json:"groupid"`
	Guestosid             string                                      `json:"guestosid"`
	Haenable              bool                                        `json:"haenable"`
	Hasannotations        bool                                        `json:"hasannotations"`
	Hostid                string                                      `json:"hostid"`
	Hostname              string                                      `json:"hostname"`
	Hypervisor            string                                      `json:"hypervisor"`
	Icon                  string                                      `json:"icon"`
	Id                    string                                      `json:"id"`
	Instancename          string                                      `json:"instancename"`
	Isdynamicallyscalable bool                                        `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                      `json:"isodisplaytext"`
	Isoid                 string                                      `json:"isoid"`
	Isoname               string                                      `json:"isoname"`
	JobID                 string                                      `json:"jobid"`
	Jobstatus             int                                         `json:"jobstatus"`
	Keypair               string                                      `json:"keypair"`
	Lastupdated           string                                      `json:"lastupdated"`
	Memory                int                                         `json:"memory"`
	Memoryintfreekbs      int64                                       `json:"memoryintfreekbs"`
	Memorykbs             int64                                       `json:"memorykbs"`
	Memorytargetkbs       int64                                       `json:"memorytargetkbs"`
	Name                  string                                      `json:"name"`
	Networkkbsread        int64                                       `json:"networkkbsread"`
	Networkkbswrite       int64                                       `json:"networkkbswrite"`
	Nic                   []Nic                                       `json:"nic"`
	Osdisplayname         string                                      `json:"osdisplayname"`
	Ostypeid              string                                      `json:"ostypeid"`
	Password              string                                      `json:"password"`
	Passwordenabled       bool                                        `json:"passwordenabled"`
	Pooltype              string                                      `json:"pooltype"`
	Project               string                                      `json:"project"`
	Projectid             string                                      `json:"projectid"`
	Publicip              string                                      `json:"publicip"`
	Publicipid            string                                      `json:"publicipid"`
	Readonlydetails       string                                      `json:"readonlydetails"`
	Receivedbytes         int64                                       `json:"receivedbytes"`
	Rootdeviceid          int64                                       `json:"rootdeviceid"`
	Rootdevicetype        string                                      `json:"rootdevicetype"`
	Securitygroup         []RebootVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                       `json:"sentbytes"`
	Serviceofferingid     string                                      `json:"serviceofferingid"`
	Serviceofferingname   string                                      `json:"serviceofferingname"`
	Servicestate          string                                      `json:"servicestate"`
	State                 string                                      `json:"state"`
	Tags                  []Tags                                      `json:"tags"`
	Templatedisplaytext   string                                      `json:"templatedisplaytext"`
	Templateid            string                                      `json:"templateid"`
	Templatename          string                                      `json:"templatename"`
	Userid                string                                      `json:"userid"`
	Username              string                                      `json:"username"`
	Vgpu                  string                                      `json:"vgpu"`
	Zoneid                string                                      `json:"zoneid"`
	Zonename              string                                      `json:"zonename"`
}

type RebootVirtualMachineResponseSecuritygroup struct {
	Account             string                                          `json:"account"`
	Description         string                                          `json:"description"`
	Domain              string                                          `json:"domain"`
	Domainid            string                                          `json:"domainid"`
	Egressrule          []RebootVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                          `json:"id"`
	Ingressrule         []RebootVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                          `json:"name"`
	Project             string                                          `json:"project"`
	Projectid           string                                          `json:"projectid"`
	Tags                []Tags                                          `json:"tags"`
	Virtualmachinecount int                                             `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                   `json:"virtualmachineids"`
}

type RebootVirtualMachineResponseSecuritygroupRule struct {
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

type RebootVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *RebootVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias RebootVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RecoverVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *RecoverVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *RecoverVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RecoverVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RecoverVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRecoverVirtualMachineParams(id string) *RecoverVirtualMachineParams {
	P := &RecoverVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Recovers a virtual machine.
func (s *VirtualMachineService) RecoverVirtualMachine(p *RecoverVirtualMachineParams) (*RecoverVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("recoverVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RecoverVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RecoverVirtualMachineResponse struct {
	Account               string                                       `json:"account"`
	Affinitygroup         []RecoverVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                       `json:"backupofferingid"`
	Backupofferingname    string                                       `json:"backupofferingname"`
	Bootmode              string                                       `json:"bootmode"`
	Boottype              string                                       `json:"boottype"`
	Cpunumber             int                                          `json:"cpunumber"`
	Cpuspeed              int                                          `json:"cpuspeed"`
	Cpuused               string                                       `json:"cpuused"`
	Created               string                                       `json:"created"`
	Details               map[string]string                            `json:"details"`
	Diskioread            int64                                        `json:"diskioread"`
	Diskiowrite           int64                                        `json:"diskiowrite"`
	Diskkbsread           int64                                        `json:"diskkbsread"`
	Diskkbswrite          int64                                        `json:"diskkbswrite"`
	Diskofferingid        string                                       `json:"diskofferingid"`
	Diskofferingname      string                                       `json:"diskofferingname"`
	Displayname           string                                       `json:"displayname"`
	Displayvm             bool                                         `json:"displayvm"`
	Domain                string                                       `json:"domain"`
	Domainid              string                                       `json:"domainid"`
	Forvirtualnetwork     bool                                         `json:"forvirtualnetwork"`
	Group                 string                                       `json:"group"`
	Groupid               string                                       `json:"groupid"`
	Guestosid             string                                       `json:"guestosid"`
	Haenable              bool                                         `json:"haenable"`
	Hasannotations        bool                                         `json:"hasannotations"`
	Hostid                string                                       `json:"hostid"`
	Hostname              string                                       `json:"hostname"`
	Hypervisor            string                                       `json:"hypervisor"`
	Icon                  string                                       `json:"icon"`
	Id                    string                                       `json:"id"`
	Instancename          string                                       `json:"instancename"`
	Isdynamicallyscalable bool                                         `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                       `json:"isodisplaytext"`
	Isoid                 string                                       `json:"isoid"`
	Isoname               string                                       `json:"isoname"`
	JobID                 string                                       `json:"jobid"`
	Jobstatus             int                                          `json:"jobstatus"`
	Keypair               string                                       `json:"keypair"`
	Lastupdated           string                                       `json:"lastupdated"`
	Memory                int                                          `json:"memory"`
	Memoryintfreekbs      int64                                        `json:"memoryintfreekbs"`
	Memorykbs             int64                                        `json:"memorykbs"`
	Memorytargetkbs       int64                                        `json:"memorytargetkbs"`
	Name                  string                                       `json:"name"`
	Networkkbsread        int64                                        `json:"networkkbsread"`
	Networkkbswrite       int64                                        `json:"networkkbswrite"`
	Nic                   []Nic                                        `json:"nic"`
	Osdisplayname         string                                       `json:"osdisplayname"`
	Ostypeid              string                                       `json:"ostypeid"`
	Password              string                                       `json:"password"`
	Passwordenabled       bool                                         `json:"passwordenabled"`
	Pooltype              string                                       `json:"pooltype"`
	Project               string                                       `json:"project"`
	Projectid             string                                       `json:"projectid"`
	Publicip              string                                       `json:"publicip"`
	Publicipid            string                                       `json:"publicipid"`
	Readonlydetails       string                                       `json:"readonlydetails"`
	Receivedbytes         int64                                        `json:"receivedbytes"`
	Rootdeviceid          int64                                        `json:"rootdeviceid"`
	Rootdevicetype        string                                       `json:"rootdevicetype"`
	Securitygroup         []RecoverVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                        `json:"sentbytes"`
	Serviceofferingid     string                                       `json:"serviceofferingid"`
	Serviceofferingname   string                                       `json:"serviceofferingname"`
	Servicestate          string                                       `json:"servicestate"`
	State                 string                                       `json:"state"`
	Tags                  []Tags                                       `json:"tags"`
	Templatedisplaytext   string                                       `json:"templatedisplaytext"`
	Templateid            string                                       `json:"templateid"`
	Templatename          string                                       `json:"templatename"`
	Userid                string                                       `json:"userid"`
	Username              string                                       `json:"username"`
	Vgpu                  string                                       `json:"vgpu"`
	Zoneid                string                                       `json:"zoneid"`
	Zonename              string                                       `json:"zonename"`
}

type RecoverVirtualMachineResponseSecuritygroup struct {
	Account             string                                           `json:"account"`
	Description         string                                           `json:"description"`
	Domain              string                                           `json:"domain"`
	Domainid            string                                           `json:"domainid"`
	Egressrule          []RecoverVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                           `json:"id"`
	Ingressrule         []RecoverVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                           `json:"name"`
	Project             string                                           `json:"project"`
	Projectid           string                                           `json:"projectid"`
	Tags                []Tags                                           `json:"tags"`
	Virtualmachinecount int                                              `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                    `json:"virtualmachineids"`
}

type RecoverVirtualMachineResponseSecuritygroupRule struct {
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

type RecoverVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *RecoverVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias RecoverVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RemoveNicFromVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *RemoveNicFromVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *RemoveNicFromVirtualMachineParams) SetNicid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nicid"] = v
}

func (P *RemoveNicFromVirtualMachineParams) GetNicid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nicid"].(string)
	return value, ok
}

func (P *RemoveNicFromVirtualMachineParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *RemoveNicFromVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveNicFromVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRemoveNicFromVirtualMachineParams(nicid string, virtualmachineid string) *RemoveNicFromVirtualMachineParams {
	P := &RemoveNicFromVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["nicid"] = nicid
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Removes VM from specified network by deleting a NIC
func (s *VirtualMachineService) RemoveNicFromVirtualMachine(p *RemoveNicFromVirtualMachineParams) (*RemoveNicFromVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("removeNicFromVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveNicFromVirtualMachineResponse
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

type RemoveNicFromVirtualMachineResponse struct {
	Account               string                                             `json:"account"`
	Affinitygroup         []RemoveNicFromVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                             `json:"backupofferingid"`
	Backupofferingname    string                                             `json:"backupofferingname"`
	Bootmode              string                                             `json:"bootmode"`
	Boottype              string                                             `json:"boottype"`
	Cpunumber             int                                                `json:"cpunumber"`
	Cpuspeed              int                                                `json:"cpuspeed"`
	Cpuused               string                                             `json:"cpuused"`
	Created               string                                             `json:"created"`
	Details               map[string]string                                  `json:"details"`
	Diskioread            int64                                              `json:"diskioread"`
	Diskiowrite           int64                                              `json:"diskiowrite"`
	Diskkbsread           int64                                              `json:"diskkbsread"`
	Diskkbswrite          int64                                              `json:"diskkbswrite"`
	Diskofferingid        string                                             `json:"diskofferingid"`
	Diskofferingname      string                                             `json:"diskofferingname"`
	Displayname           string                                             `json:"displayname"`
	Displayvm             bool                                               `json:"displayvm"`
	Domain                string                                             `json:"domain"`
	Domainid              string                                             `json:"domainid"`
	Forvirtualnetwork     bool                                               `json:"forvirtualnetwork"`
	Group                 string                                             `json:"group"`
	Groupid               string                                             `json:"groupid"`
	Guestosid             string                                             `json:"guestosid"`
	Haenable              bool                                               `json:"haenable"`
	Hasannotations        bool                                               `json:"hasannotations"`
	Hostid                string                                             `json:"hostid"`
	Hostname              string                                             `json:"hostname"`
	Hypervisor            string                                             `json:"hypervisor"`
	Icon                  string                                             `json:"icon"`
	Id                    string                                             `json:"id"`
	Instancename          string                                             `json:"instancename"`
	Isdynamicallyscalable bool                                               `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                             `json:"isodisplaytext"`
	Isoid                 string                                             `json:"isoid"`
	Isoname               string                                             `json:"isoname"`
	JobID                 string                                             `json:"jobid"`
	Jobstatus             int                                                `json:"jobstatus"`
	Keypair               string                                             `json:"keypair"`
	Lastupdated           string                                             `json:"lastupdated"`
	Memory                int                                                `json:"memory"`
	Memoryintfreekbs      int64                                              `json:"memoryintfreekbs"`
	Memorykbs             int64                                              `json:"memorykbs"`
	Memorytargetkbs       int64                                              `json:"memorytargetkbs"`
	Name                  string                                             `json:"name"`
	Networkkbsread        int64                                              `json:"networkkbsread"`
	Networkkbswrite       int64                                              `json:"networkkbswrite"`
	Nic                   []Nic                                              `json:"nic"`
	Osdisplayname         string                                             `json:"osdisplayname"`
	Ostypeid              string                                             `json:"ostypeid"`
	Password              string                                             `json:"password"`
	Passwordenabled       bool                                               `json:"passwordenabled"`
	Pooltype              string                                             `json:"pooltype"`
	Project               string                                             `json:"project"`
	Projectid             string                                             `json:"projectid"`
	Publicip              string                                             `json:"publicip"`
	Publicipid            string                                             `json:"publicipid"`
	Readonlydetails       string                                             `json:"readonlydetails"`
	Receivedbytes         int64                                              `json:"receivedbytes"`
	Rootdeviceid          int64                                              `json:"rootdeviceid"`
	Rootdevicetype        string                                             `json:"rootdevicetype"`
	Securitygroup         []RemoveNicFromVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                              `json:"sentbytes"`
	Serviceofferingid     string                                             `json:"serviceofferingid"`
	Serviceofferingname   string                                             `json:"serviceofferingname"`
	Servicestate          string                                             `json:"servicestate"`
	State                 string                                             `json:"state"`
	Tags                  []Tags                                             `json:"tags"`
	Templatedisplaytext   string                                             `json:"templatedisplaytext"`
	Templateid            string                                             `json:"templateid"`
	Templatename          string                                             `json:"templatename"`
	Userid                string                                             `json:"userid"`
	Username              string                                             `json:"username"`
	Vgpu                  string                                             `json:"vgpu"`
	Zoneid                string                                             `json:"zoneid"`
	Zonename              string                                             `json:"zonename"`
}

type RemoveNicFromVirtualMachineResponseSecuritygroup struct {
	Account             string                                                 `json:"account"`
	Description         string                                                 `json:"description"`
	Domain              string                                                 `json:"domain"`
	Domainid            string                                                 `json:"domainid"`
	Egressrule          []RemoveNicFromVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                 `json:"id"`
	Ingressrule         []RemoveNicFromVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                 `json:"name"`
	Project             string                                                 `json:"project"`
	Projectid           string                                                 `json:"projectid"`
	Tags                []Tags                                                 `json:"tags"`
	Virtualmachinecount int                                                    `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                          `json:"virtualmachineids"`
}

type RemoveNicFromVirtualMachineResponseSecuritygroupRule struct {
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

type RemoveNicFromVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *RemoveNicFromVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias RemoveNicFromVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ResetPasswordForVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *ResetPasswordForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ResetPasswordForVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ResetPasswordForVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ResetPasswordForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewResetPasswordForVirtualMachineParams(id string) *ResetPasswordForVirtualMachineParams {
	P := &ResetPasswordForVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Resets the password for virtual machine. The virtual machine must be in a "Stopped" state and the template must already support this feature for this command to take effect. [async]
func (s *VirtualMachineService) ResetPasswordForVirtualMachine(p *ResetPasswordForVirtualMachineParams) (*ResetPasswordForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("resetPasswordForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetPasswordForVirtualMachineResponse
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

type ResetPasswordForVirtualMachineResponse struct {
	Account               string                                                `json:"account"`
	Affinitygroup         []ResetPasswordForVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                                `json:"backupofferingid"`
	Backupofferingname    string                                                `json:"backupofferingname"`
	Bootmode              string                                                `json:"bootmode"`
	Boottype              string                                                `json:"boottype"`
	Cpunumber             int                                                   `json:"cpunumber"`
	Cpuspeed              int                                                   `json:"cpuspeed"`
	Cpuused               string                                                `json:"cpuused"`
	Created               string                                                `json:"created"`
	Details               map[string]string                                     `json:"details"`
	Diskioread            int64                                                 `json:"diskioread"`
	Diskiowrite           int64                                                 `json:"diskiowrite"`
	Diskkbsread           int64                                                 `json:"diskkbsread"`
	Diskkbswrite          int64                                                 `json:"diskkbswrite"`
	Diskofferingid        string                                                `json:"diskofferingid"`
	Diskofferingname      string                                                `json:"diskofferingname"`
	Displayname           string                                                `json:"displayname"`
	Displayvm             bool                                                  `json:"displayvm"`
	Domain                string                                                `json:"domain"`
	Domainid              string                                                `json:"domainid"`
	Forvirtualnetwork     bool                                                  `json:"forvirtualnetwork"`
	Group                 string                                                `json:"group"`
	Groupid               string                                                `json:"groupid"`
	Guestosid             string                                                `json:"guestosid"`
	Haenable              bool                                                  `json:"haenable"`
	Hasannotations        bool                                                  `json:"hasannotations"`
	Hostid                string                                                `json:"hostid"`
	Hostname              string                                                `json:"hostname"`
	Hypervisor            string                                                `json:"hypervisor"`
	Icon                  string                                                `json:"icon"`
	Id                    string                                                `json:"id"`
	Instancename          string                                                `json:"instancename"`
	Isdynamicallyscalable bool                                                  `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                                `json:"isodisplaytext"`
	Isoid                 string                                                `json:"isoid"`
	Isoname               string                                                `json:"isoname"`
	JobID                 string                                                `json:"jobid"`
	Jobstatus             int                                                   `json:"jobstatus"`
	Keypair               string                                                `json:"keypair"`
	Lastupdated           string                                                `json:"lastupdated"`
	Memory                int                                                   `json:"memory"`
	Memoryintfreekbs      int64                                                 `json:"memoryintfreekbs"`
	Memorykbs             int64                                                 `json:"memorykbs"`
	Memorytargetkbs       int64                                                 `json:"memorytargetkbs"`
	Name                  string                                                `json:"name"`
	Networkkbsread        int64                                                 `json:"networkkbsread"`
	Networkkbswrite       int64                                                 `json:"networkkbswrite"`
	Nic                   []Nic                                                 `json:"nic"`
	Osdisplayname         string                                                `json:"osdisplayname"`
	Ostypeid              string                                                `json:"ostypeid"`
	Password              string                                                `json:"password"`
	Passwordenabled       bool                                                  `json:"passwordenabled"`
	Pooltype              string                                                `json:"pooltype"`
	Project               string                                                `json:"project"`
	Projectid             string                                                `json:"projectid"`
	Publicip              string                                                `json:"publicip"`
	Publicipid            string                                                `json:"publicipid"`
	Readonlydetails       string                                                `json:"readonlydetails"`
	Receivedbytes         int64                                                 `json:"receivedbytes"`
	Rootdeviceid          int64                                                 `json:"rootdeviceid"`
	Rootdevicetype        string                                                `json:"rootdevicetype"`
	Securitygroup         []ResetPasswordForVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                                 `json:"sentbytes"`
	Serviceofferingid     string                                                `json:"serviceofferingid"`
	Serviceofferingname   string                                                `json:"serviceofferingname"`
	Servicestate          string                                                `json:"servicestate"`
	State                 string                                                `json:"state"`
	Tags                  []Tags                                                `json:"tags"`
	Templatedisplaytext   string                                                `json:"templatedisplaytext"`
	Templateid            string                                                `json:"templateid"`
	Templatename          string                                                `json:"templatename"`
	Userid                string                                                `json:"userid"`
	Username              string                                                `json:"username"`
	Vgpu                  string                                                `json:"vgpu"`
	Zoneid                string                                                `json:"zoneid"`
	Zonename              string                                                `json:"zonename"`
}

type ResetPasswordForVirtualMachineResponseSecuritygroup struct {
	Account             string                                                    `json:"account"`
	Description         string                                                    `json:"description"`
	Domain              string                                                    `json:"domain"`
	Domainid            string                                                    `json:"domainid"`
	Egressrule          []ResetPasswordForVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                    `json:"id"`
	Ingressrule         []ResetPasswordForVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                    `json:"name"`
	Project             string                                                    `json:"project"`
	Projectid           string                                                    `json:"projectid"`
	Tags                []Tags                                                    `json:"tags"`
	Virtualmachinecount int                                                       `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                             `json:"virtualmachineids"`
}

type ResetPasswordForVirtualMachineResponseSecuritygroupRule struct {
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

type ResetPasswordForVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *ResetPasswordForVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias ResetPasswordForVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type RestoreVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *RestoreVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *RestoreVirtualMachineParams) SetTemplateid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templateid"] = v
}

func (P *RestoreVirtualMachineParams) GetTemplateid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templateid"].(string)
	return value, ok
}

func (P *RestoreVirtualMachineParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *RestoreVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new RestoreVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRestoreVirtualMachineParams(virtualmachineid string) *RestoreVirtualMachineParams {
	P := &RestoreVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Restore a VM to original template/ISO or new template/ISO
func (s *VirtualMachineService) RestoreVirtualMachine(p *RestoreVirtualMachineParams) (*RestoreVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("restoreVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RestoreVirtualMachineResponse
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

type RestoreVirtualMachineResponse struct {
	Account               string                                       `json:"account"`
	Affinitygroup         []RestoreVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                       `json:"backupofferingid"`
	Backupofferingname    string                                       `json:"backupofferingname"`
	Bootmode              string                                       `json:"bootmode"`
	Boottype              string                                       `json:"boottype"`
	Cpunumber             int                                          `json:"cpunumber"`
	Cpuspeed              int                                          `json:"cpuspeed"`
	Cpuused               string                                       `json:"cpuused"`
	Created               string                                       `json:"created"`
	Details               map[string]string                            `json:"details"`
	Diskioread            int64                                        `json:"diskioread"`
	Diskiowrite           int64                                        `json:"diskiowrite"`
	Diskkbsread           int64                                        `json:"diskkbsread"`
	Diskkbswrite          int64                                        `json:"diskkbswrite"`
	Diskofferingid        string                                       `json:"diskofferingid"`
	Diskofferingname      string                                       `json:"diskofferingname"`
	Displayname           string                                       `json:"displayname"`
	Displayvm             bool                                         `json:"displayvm"`
	Domain                string                                       `json:"domain"`
	Domainid              string                                       `json:"domainid"`
	Forvirtualnetwork     bool                                         `json:"forvirtualnetwork"`
	Group                 string                                       `json:"group"`
	Groupid               string                                       `json:"groupid"`
	Guestosid             string                                       `json:"guestosid"`
	Haenable              bool                                         `json:"haenable"`
	Hasannotations        bool                                         `json:"hasannotations"`
	Hostid                string                                       `json:"hostid"`
	Hostname              string                                       `json:"hostname"`
	Hypervisor            string                                       `json:"hypervisor"`
	Icon                  string                                       `json:"icon"`
	Id                    string                                       `json:"id"`
	Instancename          string                                       `json:"instancename"`
	Isdynamicallyscalable bool                                         `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                       `json:"isodisplaytext"`
	Isoid                 string                                       `json:"isoid"`
	Isoname               string                                       `json:"isoname"`
	JobID                 string                                       `json:"jobid"`
	Jobstatus             int                                          `json:"jobstatus"`
	Keypair               string                                       `json:"keypair"`
	Lastupdated           string                                       `json:"lastupdated"`
	Memory                int                                          `json:"memory"`
	Memoryintfreekbs      int64                                        `json:"memoryintfreekbs"`
	Memorykbs             int64                                        `json:"memorykbs"`
	Memorytargetkbs       int64                                        `json:"memorytargetkbs"`
	Name                  string                                       `json:"name"`
	Networkkbsread        int64                                        `json:"networkkbsread"`
	Networkkbswrite       int64                                        `json:"networkkbswrite"`
	Nic                   []Nic                                        `json:"nic"`
	Osdisplayname         string                                       `json:"osdisplayname"`
	Ostypeid              string                                       `json:"ostypeid"`
	Password              string                                       `json:"password"`
	Passwordenabled       bool                                         `json:"passwordenabled"`
	Pooltype              string                                       `json:"pooltype"`
	Project               string                                       `json:"project"`
	Projectid             string                                       `json:"projectid"`
	Publicip              string                                       `json:"publicip"`
	Publicipid            string                                       `json:"publicipid"`
	Readonlydetails       string                                       `json:"readonlydetails"`
	Receivedbytes         int64                                        `json:"receivedbytes"`
	Rootdeviceid          int64                                        `json:"rootdeviceid"`
	Rootdevicetype        string                                       `json:"rootdevicetype"`
	Securitygroup         []RestoreVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                        `json:"sentbytes"`
	Serviceofferingid     string                                       `json:"serviceofferingid"`
	Serviceofferingname   string                                       `json:"serviceofferingname"`
	Servicestate          string                                       `json:"servicestate"`
	State                 string                                       `json:"state"`
	Tags                  []Tags                                       `json:"tags"`
	Templatedisplaytext   string                                       `json:"templatedisplaytext"`
	Templateid            string                                       `json:"templateid"`
	Templatename          string                                       `json:"templatename"`
	Userid                string                                       `json:"userid"`
	Username              string                                       `json:"username"`
	Vgpu                  string                                       `json:"vgpu"`
	Zoneid                string                                       `json:"zoneid"`
	Zonename              string                                       `json:"zonename"`
}

type RestoreVirtualMachineResponseSecuritygroup struct {
	Account             string                                           `json:"account"`
	Description         string                                           `json:"description"`
	Domain              string                                           `json:"domain"`
	Domainid            string                                           `json:"domainid"`
	Egressrule          []RestoreVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                           `json:"id"`
	Ingressrule         []RestoreVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                           `json:"name"`
	Project             string                                           `json:"project"`
	Projectid           string                                           `json:"projectid"`
	Tags                []Tags                                           `json:"tags"`
	Virtualmachinecount int                                              `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                    `json:"virtualmachineids"`
}

type RestoreVirtualMachineResponseSecuritygroupRule struct {
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

type RestoreVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *RestoreVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias RestoreVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ScaleVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *ScaleVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (P *ScaleVirtualMachineParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ScaleVirtualMachineParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *ScaleVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ScaleVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ScaleVirtualMachineParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *ScaleVirtualMachineParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

// You should always use this function to get a new ScaleVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewScaleVirtualMachineParams(id string, serviceofferingid string) *ScaleVirtualMachineParams {
	P := &ScaleVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["serviceofferingid"] = serviceofferingid
	return P
}

// Scales the virtual machine to a new service offering. This command also takes into account the Volume and it may resize the root disk size according to the service offering.
func (s *VirtualMachineService) ScaleVirtualMachine(p *ScaleVirtualMachineParams) (*ScaleVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("scaleVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ScaleVirtualMachineResponse
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

type ScaleVirtualMachineResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type StartVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *StartVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["bootintosetup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("bootintosetup", vv)
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["deploymentplanner"]; found {
		u.Set("deploymentplanner", v.(string))
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	return u
}

func (P *StartVirtualMachineParams) SetBootintosetup(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bootintosetup"] = v
}

func (P *StartVirtualMachineParams) GetBootintosetup() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bootintosetup"].(bool)
	return value, ok
}

func (P *StartVirtualMachineParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *StartVirtualMachineParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *StartVirtualMachineParams) SetDeploymentplanner(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["deploymentplanner"] = v
}

func (P *StartVirtualMachineParams) GetDeploymentplanner() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["deploymentplanner"].(string)
	return value, ok
}

func (P *StartVirtualMachineParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *StartVirtualMachineParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *StartVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *StartVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *StartVirtualMachineParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *StartVirtualMachineParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

// You should always use this function to get a new StartVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewStartVirtualMachineParams(id string) *StartVirtualMachineParams {
	P := &StartVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Starts a virtual machine.
func (s *VirtualMachineService) StartVirtualMachine(p *StartVirtualMachineParams) (*StartVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("startVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartVirtualMachineResponse
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

type StartVirtualMachineResponse struct {
	Account               string                                     `json:"account"`
	Affinitygroup         []StartVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                     `json:"backupofferingid"`
	Backupofferingname    string                                     `json:"backupofferingname"`
	Bootmode              string                                     `json:"bootmode"`
	Boottype              string                                     `json:"boottype"`
	Cpunumber             int                                        `json:"cpunumber"`
	Cpuspeed              int                                        `json:"cpuspeed"`
	Cpuused               string                                     `json:"cpuused"`
	Created               string                                     `json:"created"`
	Details               map[string]string                          `json:"details"`
	Diskioread            int64                                      `json:"diskioread"`
	Diskiowrite           int64                                      `json:"diskiowrite"`
	Diskkbsread           int64                                      `json:"diskkbsread"`
	Diskkbswrite          int64                                      `json:"diskkbswrite"`
	Diskofferingid        string                                     `json:"diskofferingid"`
	Diskofferingname      string                                     `json:"diskofferingname"`
	Displayname           string                                     `json:"displayname"`
	Displayvm             bool                                       `json:"displayvm"`
	Domain                string                                     `json:"domain"`
	Domainid              string                                     `json:"domainid"`
	Forvirtualnetwork     bool                                       `json:"forvirtualnetwork"`
	Group                 string                                     `json:"group"`
	Groupid               string                                     `json:"groupid"`
	Guestosid             string                                     `json:"guestosid"`
	Haenable              bool                                       `json:"haenable"`
	Hasannotations        bool                                       `json:"hasannotations"`
	Hostid                string                                     `json:"hostid"`
	Hostname              string                                     `json:"hostname"`
	Hypervisor            string                                     `json:"hypervisor"`
	Icon                  string                                     `json:"icon"`
	Id                    string                                     `json:"id"`
	Instancename          string                                     `json:"instancename"`
	Isdynamicallyscalable bool                                       `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                     `json:"isodisplaytext"`
	Isoid                 string                                     `json:"isoid"`
	Isoname               string                                     `json:"isoname"`
	JobID                 string                                     `json:"jobid"`
	Jobstatus             int                                        `json:"jobstatus"`
	Keypair               string                                     `json:"keypair"`
	Lastupdated           string                                     `json:"lastupdated"`
	Memory                int                                        `json:"memory"`
	Memoryintfreekbs      int64                                      `json:"memoryintfreekbs"`
	Memorykbs             int64                                      `json:"memorykbs"`
	Memorytargetkbs       int64                                      `json:"memorytargetkbs"`
	Name                  string                                     `json:"name"`
	Networkkbsread        int64                                      `json:"networkkbsread"`
	Networkkbswrite       int64                                      `json:"networkkbswrite"`
	Nic                   []Nic                                      `json:"nic"`
	Osdisplayname         string                                     `json:"osdisplayname"`
	Ostypeid              string                                     `json:"ostypeid"`
	Password              string                                     `json:"password"`
	Passwordenabled       bool                                       `json:"passwordenabled"`
	Pooltype              string                                     `json:"pooltype"`
	Project               string                                     `json:"project"`
	Projectid             string                                     `json:"projectid"`
	Publicip              string                                     `json:"publicip"`
	Publicipid            string                                     `json:"publicipid"`
	Readonlydetails       string                                     `json:"readonlydetails"`
	Receivedbytes         int64                                      `json:"receivedbytes"`
	Rootdeviceid          int64                                      `json:"rootdeviceid"`
	Rootdevicetype        string                                     `json:"rootdevicetype"`
	Securitygroup         []StartVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                      `json:"sentbytes"`
	Serviceofferingid     string                                     `json:"serviceofferingid"`
	Serviceofferingname   string                                     `json:"serviceofferingname"`
	Servicestate          string                                     `json:"servicestate"`
	State                 string                                     `json:"state"`
	Tags                  []Tags                                     `json:"tags"`
	Templatedisplaytext   string                                     `json:"templatedisplaytext"`
	Templateid            string                                     `json:"templateid"`
	Templatename          string                                     `json:"templatename"`
	Userid                string                                     `json:"userid"`
	Username              string                                     `json:"username"`
	Vgpu                  string                                     `json:"vgpu"`
	Zoneid                string                                     `json:"zoneid"`
	Zonename              string                                     `json:"zonename"`
}

type StartVirtualMachineResponseSecuritygroup struct {
	Account             string                                         `json:"account"`
	Description         string                                         `json:"description"`
	Domain              string                                         `json:"domain"`
	Domainid            string                                         `json:"domainid"`
	Egressrule          []StartVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                         `json:"id"`
	Ingressrule         []StartVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                         `json:"name"`
	Project             string                                         `json:"project"`
	Projectid           string                                         `json:"projectid"`
	Tags                []Tags                                         `json:"tags"`
	Virtualmachinecount int                                            `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                  `json:"virtualmachineids"`
}

type StartVirtualMachineResponseSecuritygroupRule struct {
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

type StartVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *StartVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias StartVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type StopVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *StopVirtualMachineParams) toURLValues() url.Values {
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

func (P *StopVirtualMachineParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *StopVirtualMachineParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *StopVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *StopVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new StopVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewStopVirtualMachineParams(id string) *StopVirtualMachineParams {
	P := &StopVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Stops a virtual machine.
func (s *VirtualMachineService) StopVirtualMachine(p *StopVirtualMachineParams) (*StopVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("stopVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopVirtualMachineResponse
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

type StopVirtualMachineResponse struct {
	Account               string                                    `json:"account"`
	Affinitygroup         []StopVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                    `json:"backupofferingid"`
	Backupofferingname    string                                    `json:"backupofferingname"`
	Bootmode              string                                    `json:"bootmode"`
	Boottype              string                                    `json:"boottype"`
	Cpunumber             int                                       `json:"cpunumber"`
	Cpuspeed              int                                       `json:"cpuspeed"`
	Cpuused               string                                    `json:"cpuused"`
	Created               string                                    `json:"created"`
	Details               map[string]string                         `json:"details"`
	Diskioread            int64                                     `json:"diskioread"`
	Diskiowrite           int64                                     `json:"diskiowrite"`
	Diskkbsread           int64                                     `json:"diskkbsread"`
	Diskkbswrite          int64                                     `json:"diskkbswrite"`
	Diskofferingid        string                                    `json:"diskofferingid"`
	Diskofferingname      string                                    `json:"diskofferingname"`
	Displayname           string                                    `json:"displayname"`
	Displayvm             bool                                      `json:"displayvm"`
	Domain                string                                    `json:"domain"`
	Domainid              string                                    `json:"domainid"`
	Forvirtualnetwork     bool                                      `json:"forvirtualnetwork"`
	Group                 string                                    `json:"group"`
	Groupid               string                                    `json:"groupid"`
	Guestosid             string                                    `json:"guestosid"`
	Haenable              bool                                      `json:"haenable"`
	Hasannotations        bool                                      `json:"hasannotations"`
	Hostid                string                                    `json:"hostid"`
	Hostname              string                                    `json:"hostname"`
	Hypervisor            string                                    `json:"hypervisor"`
	Icon                  string                                    `json:"icon"`
	Id                    string                                    `json:"id"`
	Instancename          string                                    `json:"instancename"`
	Isdynamicallyscalable bool                                      `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                    `json:"isodisplaytext"`
	Isoid                 string                                    `json:"isoid"`
	Isoname               string                                    `json:"isoname"`
	JobID                 string                                    `json:"jobid"`
	Jobstatus             int                                       `json:"jobstatus"`
	Keypair               string                                    `json:"keypair"`
	Lastupdated           string                                    `json:"lastupdated"`
	Memory                int                                       `json:"memory"`
	Memoryintfreekbs      int64                                     `json:"memoryintfreekbs"`
	Memorykbs             int64                                     `json:"memorykbs"`
	Memorytargetkbs       int64                                     `json:"memorytargetkbs"`
	Name                  string                                    `json:"name"`
	Networkkbsread        int64                                     `json:"networkkbsread"`
	Networkkbswrite       int64                                     `json:"networkkbswrite"`
	Nic                   []Nic                                     `json:"nic"`
	Osdisplayname         string                                    `json:"osdisplayname"`
	Ostypeid              string                                    `json:"ostypeid"`
	Password              string                                    `json:"password"`
	Passwordenabled       bool                                      `json:"passwordenabled"`
	Pooltype              string                                    `json:"pooltype"`
	Project               string                                    `json:"project"`
	Projectid             string                                    `json:"projectid"`
	Publicip              string                                    `json:"publicip"`
	Publicipid            string                                    `json:"publicipid"`
	Readonlydetails       string                                    `json:"readonlydetails"`
	Receivedbytes         int64                                     `json:"receivedbytes"`
	Rootdeviceid          int64                                     `json:"rootdeviceid"`
	Rootdevicetype        string                                    `json:"rootdevicetype"`
	Securitygroup         []StopVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                     `json:"sentbytes"`
	Serviceofferingid     string                                    `json:"serviceofferingid"`
	Serviceofferingname   string                                    `json:"serviceofferingname"`
	Servicestate          string                                    `json:"servicestate"`
	State                 string                                    `json:"state"`
	Tags                  []Tags                                    `json:"tags"`
	Templatedisplaytext   string                                    `json:"templatedisplaytext"`
	Templateid            string                                    `json:"templateid"`
	Templatename          string                                    `json:"templatename"`
	Userid                string                                    `json:"userid"`
	Username              string                                    `json:"username"`
	Vgpu                  string                                    `json:"vgpu"`
	Zoneid                string                                    `json:"zoneid"`
	Zonename              string                                    `json:"zonename"`
}

type StopVirtualMachineResponseSecuritygroup struct {
	Account             string                                        `json:"account"`
	Description         string                                        `json:"description"`
	Domain              string                                        `json:"domain"`
	Domainid            string                                        `json:"domainid"`
	Egressrule          []StopVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                        `json:"id"`
	Ingressrule         []StopVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                        `json:"name"`
	Project             string                                        `json:"project"`
	Projectid           string                                        `json:"projectid"`
	Tags                []Tags                                        `json:"tags"`
	Virtualmachinecount int                                           `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                 `json:"virtualmachineids"`
}

type StopVirtualMachineResponseSecuritygroupRule struct {
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

type StopVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *StopVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias StopVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateDefaultNicForVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *UpdateDefaultNicForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *UpdateDefaultNicForVirtualMachineParams) SetNicid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nicid"] = v
}

func (P *UpdateDefaultNicForVirtualMachineParams) GetNicid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nicid"].(string)
	return value, ok
}

func (P *UpdateDefaultNicForVirtualMachineParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *UpdateDefaultNicForVirtualMachineParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateDefaultNicForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewUpdateDefaultNicForVirtualMachineParams(nicid string, virtualmachineid string) *UpdateDefaultNicForVirtualMachineParams {
	P := &UpdateDefaultNicForVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["nicid"] = nicid
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Changes the default NIC on a VM
func (s *VirtualMachineService) UpdateDefaultNicForVirtualMachine(p *UpdateDefaultNicForVirtualMachineParams) (*UpdateDefaultNicForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("updateDefaultNicForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateDefaultNicForVirtualMachineResponse
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

type UpdateDefaultNicForVirtualMachineResponse struct {
	Account               string                                                   `json:"account"`
	Affinitygroup         []UpdateDefaultNicForVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                                   `json:"backupofferingid"`
	Backupofferingname    string                                                   `json:"backupofferingname"`
	Bootmode              string                                                   `json:"bootmode"`
	Boottype              string                                                   `json:"boottype"`
	Cpunumber             int                                                      `json:"cpunumber"`
	Cpuspeed              int                                                      `json:"cpuspeed"`
	Cpuused               string                                                   `json:"cpuused"`
	Created               string                                                   `json:"created"`
	Details               map[string]string                                        `json:"details"`
	Diskioread            int64                                                    `json:"diskioread"`
	Diskiowrite           int64                                                    `json:"diskiowrite"`
	Diskkbsread           int64                                                    `json:"diskkbsread"`
	Diskkbswrite          int64                                                    `json:"diskkbswrite"`
	Diskofferingid        string                                                   `json:"diskofferingid"`
	Diskofferingname      string                                                   `json:"diskofferingname"`
	Displayname           string                                                   `json:"displayname"`
	Displayvm             bool                                                     `json:"displayvm"`
	Domain                string                                                   `json:"domain"`
	Domainid              string                                                   `json:"domainid"`
	Forvirtualnetwork     bool                                                     `json:"forvirtualnetwork"`
	Group                 string                                                   `json:"group"`
	Groupid               string                                                   `json:"groupid"`
	Guestosid             string                                                   `json:"guestosid"`
	Haenable              bool                                                     `json:"haenable"`
	Hasannotations        bool                                                     `json:"hasannotations"`
	Hostid                string                                                   `json:"hostid"`
	Hostname              string                                                   `json:"hostname"`
	Hypervisor            string                                                   `json:"hypervisor"`
	Icon                  string                                                   `json:"icon"`
	Id                    string                                                   `json:"id"`
	Instancename          string                                                   `json:"instancename"`
	Isdynamicallyscalable bool                                                     `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                                   `json:"isodisplaytext"`
	Isoid                 string                                                   `json:"isoid"`
	Isoname               string                                                   `json:"isoname"`
	JobID                 string                                                   `json:"jobid"`
	Jobstatus             int                                                      `json:"jobstatus"`
	Keypair               string                                                   `json:"keypair"`
	Lastupdated           string                                                   `json:"lastupdated"`
	Memory                int                                                      `json:"memory"`
	Memoryintfreekbs      int64                                                    `json:"memoryintfreekbs"`
	Memorykbs             int64                                                    `json:"memorykbs"`
	Memorytargetkbs       int64                                                    `json:"memorytargetkbs"`
	Name                  string                                                   `json:"name"`
	Networkkbsread        int64                                                    `json:"networkkbsread"`
	Networkkbswrite       int64                                                    `json:"networkkbswrite"`
	Nic                   []Nic                                                    `json:"nic"`
	Osdisplayname         string                                                   `json:"osdisplayname"`
	Ostypeid              string                                                   `json:"ostypeid"`
	Password              string                                                   `json:"password"`
	Passwordenabled       bool                                                     `json:"passwordenabled"`
	Pooltype              string                                                   `json:"pooltype"`
	Project               string                                                   `json:"project"`
	Projectid             string                                                   `json:"projectid"`
	Publicip              string                                                   `json:"publicip"`
	Publicipid            string                                                   `json:"publicipid"`
	Readonlydetails       string                                                   `json:"readonlydetails"`
	Receivedbytes         int64                                                    `json:"receivedbytes"`
	Rootdeviceid          int64                                                    `json:"rootdeviceid"`
	Rootdevicetype        string                                                   `json:"rootdevicetype"`
	Securitygroup         []UpdateDefaultNicForVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                                    `json:"sentbytes"`
	Serviceofferingid     string                                                   `json:"serviceofferingid"`
	Serviceofferingname   string                                                   `json:"serviceofferingname"`
	Servicestate          string                                                   `json:"servicestate"`
	State                 string                                                   `json:"state"`
	Tags                  []Tags                                                   `json:"tags"`
	Templatedisplaytext   string                                                   `json:"templatedisplaytext"`
	Templateid            string                                                   `json:"templateid"`
	Templatename          string                                                   `json:"templatename"`
	Userid                string                                                   `json:"userid"`
	Username              string                                                   `json:"username"`
	Vgpu                  string                                                   `json:"vgpu"`
	Zoneid                string                                                   `json:"zoneid"`
	Zonename              string                                                   `json:"zonename"`
}

type UpdateDefaultNicForVirtualMachineResponseSecuritygroup struct {
	Account             string                                                       `json:"account"`
	Description         string                                                       `json:"description"`
	Domain              string                                                       `json:"domain"`
	Domainid            string                                                       `json:"domainid"`
	Egressrule          []UpdateDefaultNicForVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                                       `json:"id"`
	Ingressrule         []UpdateDefaultNicForVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                                       `json:"name"`
	Project             string                                                       `json:"project"`
	Projectid           string                                                       `json:"projectid"`
	Tags                []Tags                                                       `json:"tags"`
	Virtualmachinecount int                                                          `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                                `json:"virtualmachineids"`
}

type UpdateDefaultNicForVirtualMachineResponseSecuritygroupRule struct {
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

type UpdateDefaultNicForVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *UpdateDefaultNicForVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateDefaultNicForVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateVirtualMachineParams struct {
	P map[string]interface{}
}

func (P *UpdateVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["cleanupdetails"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanupdetails", vv)
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := P.P["dhcpoptionsnetworklist"]; found {
		l := v.([]map[string]string)
		for i, m := range l {
			for key, val := range m {
				u.Set(fmt.Sprintf("dhcpoptionsnetworklist[%d].%s", i, key), val)
			}
		}
	}
	if v, found := P.P["displayname"]; found {
		u.Set("displayname", v.(string))
	}
	if v, found := P.P["displayvm"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvm", vv)
	}
	if v, found := P.P["extraconfig"]; found {
		u.Set("extraconfig", v.(string))
	}
	if v, found := P.P["group"]; found {
		u.Set("group", v.(string))
	}
	if v, found := P.P["haenable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("haenable", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["instancename"]; found {
		u.Set("instancename", v.(string))
	}
	if v, found := P.P["isdynamicallyscalable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdynamicallyscalable", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["ostypeid"]; found {
		u.Set("ostypeid", v.(string))
	}
	if v, found := P.P["securitygroupids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("securitygroupids", vv)
	}
	if v, found := P.P["securitygroupnames"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("securitygroupnames", vv)
	}
	if v, found := P.P["userdata"]; found {
		u.Set("userdata", v.(string))
	}
	return u
}

func (P *UpdateVirtualMachineParams) SetCleanupdetails(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cleanupdetails"] = v
}

func (P *UpdateVirtualMachineParams) GetCleanupdetails() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cleanupdetails"].(bool)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateVirtualMachineParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *UpdateVirtualMachineParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetDhcpoptionsnetworklist(v []map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dhcpoptionsnetworklist"] = v
}

func (P *UpdateVirtualMachineParams) GetDhcpoptionsnetworklist() ([]map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dhcpoptionsnetworklist"].([]map[string]string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) AddDhcpoptionsnetworklist(item map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	val, found := P.P["dhcpoptionsnetworklist"]
	if !found {
		P.P["dhcpoptionsnetworklist"] = []map[string]string{}
		val = P.P["dhcpoptionsnetworklist"]
	}
	l := val.([]map[string]string)
	l = append(l, item)
	P.P["dhcpoptionsnetworklist"] = l
}

func (P *UpdateVirtualMachineParams) SetDisplayname(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayname"] = v
}

func (P *UpdateVirtualMachineParams) GetDisplayname() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayname"].(string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetDisplayvm(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayvm"] = v
}

func (P *UpdateVirtualMachineParams) GetDisplayvm() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayvm"].(bool)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetExtraconfig(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["extraconfig"] = v
}

func (P *UpdateVirtualMachineParams) GetExtraconfig() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["extraconfig"].(string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetGroup(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["group"] = v
}

func (P *UpdateVirtualMachineParams) GetGroup() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["group"].(string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetHaenable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["haenable"] = v
}

func (P *UpdateVirtualMachineParams) GetHaenable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["haenable"].(bool)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateVirtualMachineParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetInstancename(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["instancename"] = v
}

func (P *UpdateVirtualMachineParams) GetInstancename() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["instancename"].(string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetIsdynamicallyscalable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isdynamicallyscalable"] = v
}

func (P *UpdateVirtualMachineParams) GetIsdynamicallyscalable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isdynamicallyscalable"].(bool)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateVirtualMachineParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetOstypeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ostypeid"] = v
}

func (P *UpdateVirtualMachineParams) GetOstypeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ostypeid"].(string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetSecuritygroupids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupids"] = v
}

func (P *UpdateVirtualMachineParams) GetSecuritygroupids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupids"].([]string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetSecuritygroupnames(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["securitygroupnames"] = v
}

func (P *UpdateVirtualMachineParams) GetSecuritygroupnames() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["securitygroupnames"].([]string)
	return value, ok
}

func (P *UpdateVirtualMachineParams) SetUserdata(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userdata"] = v
}

func (P *UpdateVirtualMachineParams) GetUserdata() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userdata"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewUpdateVirtualMachineParams(id string) *UpdateVirtualMachineParams {
	P := &UpdateVirtualMachineParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates properties of a virtual machine. The VM has to be stopped and restarted for the new properties to take effect. UpdateVirtualMachine does not first check whether the VM is stopped. Therefore, stop the VM manually before issuing this call.
func (s *VirtualMachineService) UpdateVirtualMachine(p *UpdateVirtualMachineParams) (*UpdateVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("updateVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateVirtualMachineResponse struct {
	Account               string                                      `json:"account"`
	Affinitygroup         []UpdateVirtualMachineResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                                      `json:"backupofferingid"`
	Backupofferingname    string                                      `json:"backupofferingname"`
	Bootmode              string                                      `json:"bootmode"`
	Boottype              string                                      `json:"boottype"`
	Cpunumber             int                                         `json:"cpunumber"`
	Cpuspeed              int                                         `json:"cpuspeed"`
	Cpuused               string                                      `json:"cpuused"`
	Created               string                                      `json:"created"`
	Details               map[string]string                           `json:"details"`
	Diskioread            int64                                       `json:"diskioread"`
	Diskiowrite           int64                                       `json:"diskiowrite"`
	Diskkbsread           int64                                       `json:"diskkbsread"`
	Diskkbswrite          int64                                       `json:"diskkbswrite"`
	Diskofferingid        string                                      `json:"diskofferingid"`
	Diskofferingname      string                                      `json:"diskofferingname"`
	Displayname           string                                      `json:"displayname"`
	Displayvm             bool                                        `json:"displayvm"`
	Domain                string                                      `json:"domain"`
	Domainid              string                                      `json:"domainid"`
	Forvirtualnetwork     bool                                        `json:"forvirtualnetwork"`
	Group                 string                                      `json:"group"`
	Groupid               string                                      `json:"groupid"`
	Guestosid             string                                      `json:"guestosid"`
	Haenable              bool                                        `json:"haenable"`
	Hasannotations        bool                                        `json:"hasannotations"`
	Hostid                string                                      `json:"hostid"`
	Hostname              string                                      `json:"hostname"`
	Hypervisor            string                                      `json:"hypervisor"`
	Icon                  string                                      `json:"icon"`
	Id                    string                                      `json:"id"`
	Instancename          string                                      `json:"instancename"`
	Isdynamicallyscalable bool                                        `json:"isdynamicallyscalable"`
	Isodisplaytext        string                                      `json:"isodisplaytext"`
	Isoid                 string                                      `json:"isoid"`
	Isoname               string                                      `json:"isoname"`
	JobID                 string                                      `json:"jobid"`
	Jobstatus             int                                         `json:"jobstatus"`
	Keypair               string                                      `json:"keypair"`
	Lastupdated           string                                      `json:"lastupdated"`
	Memory                int                                         `json:"memory"`
	Memoryintfreekbs      int64                                       `json:"memoryintfreekbs"`
	Memorykbs             int64                                       `json:"memorykbs"`
	Memorytargetkbs       int64                                       `json:"memorytargetkbs"`
	Name                  string                                      `json:"name"`
	Networkkbsread        int64                                       `json:"networkkbsread"`
	Networkkbswrite       int64                                       `json:"networkkbswrite"`
	Nic                   []Nic                                       `json:"nic"`
	Osdisplayname         string                                      `json:"osdisplayname"`
	Ostypeid              string                                      `json:"ostypeid"`
	Password              string                                      `json:"password"`
	Passwordenabled       bool                                        `json:"passwordenabled"`
	Pooltype              string                                      `json:"pooltype"`
	Project               string                                      `json:"project"`
	Projectid             string                                      `json:"projectid"`
	Publicip              string                                      `json:"publicip"`
	Publicipid            string                                      `json:"publicipid"`
	Readonlydetails       string                                      `json:"readonlydetails"`
	Receivedbytes         int64                                       `json:"receivedbytes"`
	Rootdeviceid          int64                                       `json:"rootdeviceid"`
	Rootdevicetype        string                                      `json:"rootdevicetype"`
	Securitygroup         []UpdateVirtualMachineResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                       `json:"sentbytes"`
	Serviceofferingid     string                                      `json:"serviceofferingid"`
	Serviceofferingname   string                                      `json:"serviceofferingname"`
	Servicestate          string                                      `json:"servicestate"`
	State                 string                                      `json:"state"`
	Tags                  []Tags                                      `json:"tags"`
	Templatedisplaytext   string                                      `json:"templatedisplaytext"`
	Templateid            string                                      `json:"templateid"`
	Templatename          string                                      `json:"templatename"`
	Userid                string                                      `json:"userid"`
	Username              string                                      `json:"username"`
	Vgpu                  string                                      `json:"vgpu"`
	Zoneid                string                                      `json:"zoneid"`
	Zonename              string                                      `json:"zonename"`
}

type UpdateVirtualMachineResponseSecuritygroup struct {
	Account             string                                          `json:"account"`
	Description         string                                          `json:"description"`
	Domain              string                                          `json:"domain"`
	Domainid            string                                          `json:"domainid"`
	Egressrule          []UpdateVirtualMachineResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                          `json:"id"`
	Ingressrule         []UpdateVirtualMachineResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                          `json:"name"`
	Project             string                                          `json:"project"`
	Projectid           string                                          `json:"projectid"`
	Tags                []Tags                                          `json:"tags"`
	Virtualmachinecount int                                             `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                                   `json:"virtualmachineids"`
}

type UpdateVirtualMachineResponseSecuritygroupRule struct {
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

type UpdateVirtualMachineResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *UpdateVirtualMachineResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateVirtualMachineResponse
	return json.Unmarshal(b, (*alias)(r))
}
