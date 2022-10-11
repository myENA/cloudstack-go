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

type SystemVMServiceIface interface {
	ChangeServiceForSystemVm(p *ChangeServiceForSystemVmParams) (*ChangeServiceForSystemVmResponse, error)
	NewChangeServiceForSystemVmParams(id string, serviceofferingid string) *ChangeServiceForSystemVmParams
	DestroySystemVm(p *DestroySystemVmParams) (*DestroySystemVmResponse, error)
	NewDestroySystemVmParams(id string) *DestroySystemVmParams
	ListSystemVms(p *ListSystemVmsParams) (*ListSystemVmsResponse, error)
	NewListSystemVmsParams() *ListSystemVmsParams
	GetSystemVmID(name string, opts ...OptionFunc) (string, int, error)
	GetSystemVmByName(name string, opts ...OptionFunc) (*SystemVm, int, error)
	GetSystemVmByID(id string, opts ...OptionFunc) (*SystemVm, int, error)
	MigrateSystemVm(p *MigrateSystemVmParams) (*MigrateSystemVmResponse, error)
	NewMigrateSystemVmParams(virtualmachineid string) *MigrateSystemVmParams
	RebootSystemVm(p *RebootSystemVmParams) (*RebootSystemVmResponse, error)
	NewRebootSystemVmParams(id string) *RebootSystemVmParams
	ScaleSystemVm(p *ScaleSystemVmParams) (*ScaleSystemVmResponse, error)
	NewScaleSystemVmParams(id string, serviceofferingid string) *ScaleSystemVmParams
	StartSystemVm(p *StartSystemVmParams) (*StartSystemVmResponse, error)
	NewStartSystemVmParams(id string) *StartSystemVmParams
	StopSystemVm(p *StopSystemVmParams) (*StopSystemVmResponse, error)
	NewStopSystemVmParams(id string) *StopSystemVmParams
}

type ChangeServiceForSystemVmParams struct {
	P map[string]interface{}
}

func (P *ChangeServiceForSystemVmParams) toURLValues() url.Values {
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

func (P *ChangeServiceForSystemVmParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ChangeServiceForSystemVmParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *ChangeServiceForSystemVmParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ChangeServiceForSystemVmParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ChangeServiceForSystemVmParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *ChangeServiceForSystemVmParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

// You should always use this function to get a new ChangeServiceForSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewChangeServiceForSystemVmParams(id string, serviceofferingid string) *ChangeServiceForSystemVmParams {
	P := &ChangeServiceForSystemVmParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["serviceofferingid"] = serviceofferingid
	return P
}

// Changes the service offering for a system vm (console proxy or secondary storage). The system vm must be in a "Stopped" state for this command to take effect.
func (s *SystemVMService) ChangeServiceForSystemVm(p *ChangeServiceForSystemVmParams) (*ChangeServiceForSystemVmResponse, error) {
	resp, err := s.cs.newRequest("changeServiceForSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ChangeServiceForSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type DestroySystemVmParams struct {
	P map[string]interface{}
}

func (P *DestroySystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DestroySystemVmParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DestroySystemVmParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DestroySystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewDestroySystemVmParams(id string) *DestroySystemVmParams {
	P := &DestroySystemVmParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Destroyes a system virtual machine.
func (s *SystemVMService) DestroySystemVm(p *DestroySystemVmParams) (*DestroySystemVmResponse, error) {
	resp, err := s.cs.newRequest("destroySystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroySystemVmResponse
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

type DestroySystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type ListSystemVmsParams struct {
	P map[string]interface{}
}

func (P *ListSystemVmsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
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
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := P.P["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListSystemVmsParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ListSystemVmsParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *ListSystemVmsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListSystemVmsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListSystemVmsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListSystemVmsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListSystemVmsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListSystemVmsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListSystemVmsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListSystemVmsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListSystemVmsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListSystemVmsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListSystemVmsParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListSystemVmsParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListSystemVmsParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListSystemVmsParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListSystemVmsParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *ListSystemVmsParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

func (P *ListSystemVmsParams) SetSystemvmtype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["systemvmtype"] = v
}

func (P *ListSystemVmsParams) GetSystemvmtype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["systemvmtype"].(string)
	return value, ok
}

func (P *ListSystemVmsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListSystemVmsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListSystemVmsParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewListSystemVmsParams() *ListSystemVmsParams {
	P := &ListSystemVmsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListSystemVmsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListSystemVms(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.SystemVms[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.SystemVms {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmByName(name string, opts ...OptionFunc) (*SystemVm, int, error) {
	id, count, err := s.GetSystemVmID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetSystemVmByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmByID(id string, opts ...OptionFunc) (*SystemVm, int, error) {
	P := &ListSystemVmsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListSystemVms(P)
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
		return l.SystemVms[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SystemVm UUID: %s!", id)
}

// List system virtual machines.
func (s *SystemVMService) ListSystemVms(p *ListSystemVmsParams) (*ListSystemVmsResponse, error) {
	resp, err := s.cs.newRequest("listSystemVms", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSystemVmsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListSystemVmsResponse struct {
	Count     int         `json:"count"`
	SystemVms []*SystemVm `json:"systemvm"`
}

type SystemVm struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type MigrateSystemVmParams struct {
	P map[string]interface{}
}

func (P *MigrateSystemVmParams) toURLValues() url.Values {
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

func (P *MigrateSystemVmParams) SetAutoselect(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["autoselect"] = v
}

func (P *MigrateSystemVmParams) GetAutoselect() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["autoselect"].(bool)
	return value, ok
}

func (P *MigrateSystemVmParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *MigrateSystemVmParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *MigrateSystemVmParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *MigrateSystemVmParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

func (P *MigrateSystemVmParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *MigrateSystemVmParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new MigrateSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewMigrateSystemVmParams(virtualmachineid string) *MigrateSystemVmParams {
	P := &MigrateSystemVmParams{}
	P.P = make(map[string]interface{})
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Attempts Migration of a system virtual machine to the host specified.
func (s *SystemVMService) MigrateSystemVm(p *MigrateSystemVmParams) (*MigrateSystemVmResponse, error) {
	resp, err := s.cs.newRequest("migrateSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateSystemVmResponse
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

type MigrateSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type RebootSystemVmParams struct {
	P map[string]interface{}
}

func (P *RebootSystemVmParams) toURLValues() url.Values {
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

func (P *RebootSystemVmParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *RebootSystemVmParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *RebootSystemVmParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RebootSystemVmParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RebootSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewRebootSystemVmParams(id string) *RebootSystemVmParams {
	P := &RebootSystemVmParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Reboots a system VM.
func (s *SystemVMService) RebootSystemVm(p *RebootSystemVmParams) (*RebootSystemVmResponse, error) {
	resp, err := s.cs.newRequest("rebootSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootSystemVmResponse
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

type RebootSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type ScaleSystemVmParams struct {
	P map[string]interface{}
}

func (P *ScaleSystemVmParams) toURLValues() url.Values {
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

func (P *ScaleSystemVmParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ScaleSystemVmParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *ScaleSystemVmParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ScaleSystemVmParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ScaleSystemVmParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *ScaleSystemVmParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

// You should always use this function to get a new ScaleSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewScaleSystemVmParams(id string, serviceofferingid string) *ScaleSystemVmParams {
	P := &ScaleSystemVmParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["serviceofferingid"] = serviceofferingid
	return P
}

// Scale the service offering for a system vm (console proxy or secondary storage). The system vm must be in a "Stopped" state for this command to take effect.
func (s *SystemVMService) ScaleSystemVm(p *ScaleSystemVmParams) (*ScaleSystemVmResponse, error) {
	resp, err := s.cs.newRequest("scaleSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ScaleSystemVmResponse
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

type ScaleSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type StartSystemVmParams struct {
	P map[string]interface{}
}

func (P *StartSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *StartSystemVmParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *StartSystemVmParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new StartSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewStartSystemVmParams(id string) *StartSystemVmParams {
	P := &StartSystemVmParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Starts a system virtual machine.
func (s *SystemVMService) StartSystemVm(p *StartSystemVmParams) (*StartSystemVmResponse, error) {
	resp, err := s.cs.newRequest("startSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartSystemVmResponse
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

type StartSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}

type StopSystemVmParams struct {
	P map[string]interface{}
}

func (P *StopSystemVmParams) toURLValues() url.Values {
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

func (P *StopSystemVmParams) SetForced(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forced"] = v
}

func (P *StopSystemVmParams) GetForced() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forced"].(bool)
	return value, ok
}

func (P *StopSystemVmParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *StopSystemVmParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new StopSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewStopSystemVmParams(id string) *StopSystemVmParams {
	P := &StopSystemVmParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Stops a system VM.
func (s *SystemVMService) StopSystemVm(p *StopSystemVmParams) (*StopSystemVmResponse, error) {
	resp, err := s.cs.newRequest("stopSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopSystemVmResponse
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

type StopSystemVmResponse struct {
	Activeviewersessions  int      `json:"activeviewersessions"`
	Agentstate            string   `json:"agentstate"`
	Created               string   `json:"created"`
	Disconnected          string   `json:"disconnected"`
	Dns1                  string   `json:"dns1"`
	Dns2                  string   `json:"dns2"`
	Gateway               string   `json:"gateway"`
	Guestvlan             string   `json:"guestvlan"`
	Hasannotations        bool     `json:"hasannotations"`
	Hostid                string   `json:"hostid"`
	Hostname              string   `json:"hostname"`
	Hypervisor            string   `json:"hypervisor"`
	Id                    string   `json:"id"`
	Isdynamicallyscalable bool     `json:"isdynamicallyscalable"`
	JobID                 string   `json:"jobid"`
	Jobstatus             int      `json:"jobstatus"`
	Linklocalip           string   `json:"linklocalip"`
	Linklocalmacaddress   string   `json:"linklocalmacaddress"`
	Linklocalnetmask      string   `json:"linklocalnetmask"`
	Name                  string   `json:"name"`
	Networkdomain         string   `json:"networkdomain"`
	Podid                 string   `json:"podid"`
	Podname               string   `json:"podname"`
	Privateip             string   `json:"privateip"`
	Privatemacaddress     string   `json:"privatemacaddress"`
	Privatenetmask        string   `json:"privatenetmask"`
	Publicip              string   `json:"publicip"`
	Publicmacaddress      string   `json:"publicmacaddress"`
	Publicnetmask         string   `json:"publicnetmask"`
	Publicvlan            []string `json:"publicvlan"`
	State                 string   `json:"state"`
	Systemvmtype          string   `json:"systemvmtype"`
	Templateid            string   `json:"templateid"`
	Templatename          string   `json:"templatename"`
	Version               string   `json:"version"`
	Zoneid                string   `json:"zoneid"`
	Zonename              string   `json:"zonename"`
}
