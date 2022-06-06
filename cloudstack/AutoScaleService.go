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

type AutoScaleServiceIface interface {
	CreateAutoScalePolicy(P *CreateAutoScalePolicyParams) (*CreateAutoScalePolicyResponse, error)
	NewCreateAutoScalePolicyParams(action string, conditionids []string, duration int) *CreateAutoScalePolicyParams
	CreateAutoScaleVmGroup(P *CreateAutoScaleVmGroupParams) (*CreateAutoScaleVmGroupResponse, error)
	NewCreateAutoScaleVmGroupParams(lbruleid string, maxmembers int, minmembers int, scaledownpolicyids []string, scaleuppolicyids []string, vmprofileid string) *CreateAutoScaleVmGroupParams
	CreateAutoScaleVmProfile(P *CreateAutoScaleVmProfileParams) (*CreateAutoScaleVmProfileResponse, error)
	NewCreateAutoScaleVmProfileParams(serviceofferingid string, templateid string, zoneid string) *CreateAutoScaleVmProfileParams
	CreateCondition(P *CreateConditionParams) (*CreateConditionResponse, error)
	NewCreateConditionParams(counterid string, relationaloperator string, threshold int64) *CreateConditionParams
	CreateCounter(P *CreateCounterParams) (*CreateCounterResponse, error)
	NewCreateCounterParams(name string, source string, value string) *CreateCounterParams
	DeleteAutoScalePolicy(P *DeleteAutoScalePolicyParams) (*DeleteAutoScalePolicyResponse, error)
	NewDeleteAutoScalePolicyParams(id string) *DeleteAutoScalePolicyParams
	DeleteAutoScaleVmGroup(P *DeleteAutoScaleVmGroupParams) (*DeleteAutoScaleVmGroupResponse, error)
	NewDeleteAutoScaleVmGroupParams(id string) *DeleteAutoScaleVmGroupParams
	DeleteAutoScaleVmProfile(P *DeleteAutoScaleVmProfileParams) (*DeleteAutoScaleVmProfileResponse, error)
	NewDeleteAutoScaleVmProfileParams(id string) *DeleteAutoScaleVmProfileParams
	DeleteCondition(P *DeleteConditionParams) (*DeleteConditionResponse, error)
	NewDeleteConditionParams(id string) *DeleteConditionParams
	DeleteCounter(P *DeleteCounterParams) (*DeleteCounterResponse, error)
	NewDeleteCounterParams(id string) *DeleteCounterParams
	DisableAutoScaleVmGroup(P *DisableAutoScaleVmGroupParams) (*DisableAutoScaleVmGroupResponse, error)
	NewDisableAutoScaleVmGroupParams(id string) *DisableAutoScaleVmGroupParams
	EnableAutoScaleVmGroup(P *EnableAutoScaleVmGroupParams) (*EnableAutoScaleVmGroupResponse, error)
	NewEnableAutoScaleVmGroupParams(id string) *EnableAutoScaleVmGroupParams
	ListAutoScalePolicies(P *ListAutoScalePoliciesParams) (*ListAutoScalePoliciesResponse, error)
	NewListAutoScalePoliciesParams() *ListAutoScalePoliciesParams
	GetAutoScalePolicyByID(id string, opts ...OptionFunc) (*AutoScalePolicy, int, error)
	ListAutoScaleVmGroups(P *ListAutoScaleVmGroupsParams) (*ListAutoScaleVmGroupsResponse, error)
	NewListAutoScaleVmGroupsParams() *ListAutoScaleVmGroupsParams
	GetAutoScaleVmGroupByID(id string, opts ...OptionFunc) (*AutoScaleVmGroup, int, error)
	ListAutoScaleVmProfiles(P *ListAutoScaleVmProfilesParams) (*ListAutoScaleVmProfilesResponse, error)
	NewListAutoScaleVmProfilesParams() *ListAutoScaleVmProfilesParams
	GetAutoScaleVmProfileByID(id string, opts ...OptionFunc) (*AutoScaleVmProfile, int, error)
	ListConditions(P *ListConditionsParams) (*ListConditionsResponse, error)
	NewListConditionsParams() *ListConditionsParams
	GetConditionByID(id string, opts ...OptionFunc) (*Condition, int, error)
	ListCounters(P *ListCountersParams) (*ListCountersResponse, error)
	NewListCountersParams() *ListCountersParams
	GetCounterID(name string, opts ...OptionFunc) (string, int, error)
	GetCounterByName(name string, opts ...OptionFunc) (*Counter, int, error)
	GetCounterByID(id string, opts ...OptionFunc) (*Counter, int, error)
	UpdateAutoScalePolicy(P *UpdateAutoScalePolicyParams) (*UpdateAutoScalePolicyResponse, error)
	NewUpdateAutoScalePolicyParams(id string) *UpdateAutoScalePolicyParams
	UpdateAutoScaleVmGroup(P *UpdateAutoScaleVmGroupParams) (*UpdateAutoScaleVmGroupResponse, error)
	NewUpdateAutoScaleVmGroupParams(id string) *UpdateAutoScaleVmGroupParams
	UpdateAutoScaleVmProfile(P *UpdateAutoScaleVmProfileParams) (*UpdateAutoScaleVmProfileResponse, error)
	NewUpdateAutoScaleVmProfileParams(id string) *UpdateAutoScaleVmProfileParams
}

type CreateAutoScalePolicyParams struct {
	P map[string]interface{}
}

func (P *CreateAutoScalePolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := P.P["conditionids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("conditionids", vv)
	}
	if v, found := P.P["duration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("duration", vv)
	}
	if v, found := P.P["quiettime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("quiettime", vv)
	}
	return u
}

func (P *CreateAutoScalePolicyParams) SetAction(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["action"] = v
}

func (P *CreateAutoScalePolicyParams) GetAction() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["action"].(string)
	return value, ok
}

func (P *CreateAutoScalePolicyParams) SetConditionids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["conditionids"] = v
}

func (P *CreateAutoScalePolicyParams) GetConditionids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["conditionids"].([]string)
	return value, ok
}

func (P *CreateAutoScalePolicyParams) SetDuration(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["duration"] = v
}

func (P *CreateAutoScalePolicyParams) GetDuration() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["duration"].(int)
	return value, ok
}

func (P *CreateAutoScalePolicyParams) SetQuiettime(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["quiettime"] = v
}

func (P *CreateAutoScalePolicyParams) GetQuiettime() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["quiettime"].(int)
	return value, ok
}

// You should always use this function to get a new CreateAutoScalePolicyParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateAutoScalePolicyParams(action string, conditionids []string, duration int) *CreateAutoScalePolicyParams {
	P := &CreateAutoScalePolicyParams{}
	P.P = make(map[string]interface{})
	P.P["action"] = action
	P.P["conditionids"] = conditionids
	P.P["duration"] = duration
	return P
}

// Creates an autoscale policy for a provision or deprovision action, the action is taken when the all the conditions evaluates to true for the specified duration. The policy is in effect once it is attached to a autscale vm group.
func (s *AutoScaleService) CreateAutoScalePolicy(p *CreateAutoScalePolicyParams) (*CreateAutoScalePolicyResponse, error) {
	resp, err := s.cs.newRequest("createAutoScalePolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAutoScalePolicyResponse
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

type CreateAutoScalePolicyResponse struct {
	Account    string   `json:"account"`
	Action     string   `json:"action"`
	Conditions []string `json:"conditions"`
	Domain     string   `json:"domain"`
	Domainid   string   `json:"domainid"`
	Duration   int      `json:"duration"`
	Id         string   `json:"id"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Project    string   `json:"project"`
	Projectid  string   `json:"projectid"`
	Quiettime  int      `json:"quiettime"`
}

type CreateAutoScaleVmGroupParams struct {
	P map[string]interface{}
}

func (P *CreateAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["interval"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("interval", vv)
	}
	if v, found := P.P["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := P.P["maxmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxmembers", vv)
	}
	if v, found := P.P["minmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minmembers", vv)
	}
	if v, found := P.P["scaledownpolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaledownpolicyids", vv)
	}
	if v, found := P.P["scaleuppolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaleuppolicyids", vv)
	}
	if v, found := P.P["vmprofileid"]; found {
		u.Set("vmprofileid", v.(string))
	}
	return u
}

func (P *CreateAutoScaleVmGroupParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateAutoScaleVmGroupParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateAutoScaleVmGroupParams) SetInterval(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["interval"] = v
}

func (P *CreateAutoScaleVmGroupParams) GetInterval() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["interval"].(int)
	return value, ok
}

func (P *CreateAutoScaleVmGroupParams) SetLbruleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbruleid"] = v
}

func (P *CreateAutoScaleVmGroupParams) GetLbruleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbruleid"].(string)
	return value, ok
}

func (P *CreateAutoScaleVmGroupParams) SetMaxmembers(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxmembers"] = v
}

func (P *CreateAutoScaleVmGroupParams) GetMaxmembers() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxmembers"].(int)
	return value, ok
}

func (P *CreateAutoScaleVmGroupParams) SetMinmembers(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["minmembers"] = v
}

func (P *CreateAutoScaleVmGroupParams) GetMinmembers() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["minmembers"].(int)
	return value, ok
}

func (P *CreateAutoScaleVmGroupParams) SetScaledownpolicyids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["scaledownpolicyids"] = v
}

func (P *CreateAutoScaleVmGroupParams) GetScaledownpolicyids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["scaledownpolicyids"].([]string)
	return value, ok
}

func (P *CreateAutoScaleVmGroupParams) SetScaleuppolicyids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["scaleuppolicyids"] = v
}

func (P *CreateAutoScaleVmGroupParams) GetScaleuppolicyids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["scaleuppolicyids"].([]string)
	return value, ok
}

func (P *CreateAutoScaleVmGroupParams) SetVmprofileid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmprofileid"] = v
}

func (P *CreateAutoScaleVmGroupParams) GetVmprofileid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmprofileid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateAutoScaleVmGroupParams(lbruleid string, maxmembers int, minmembers int, scaledownpolicyids []string, scaleuppolicyids []string, vmprofileid string) *CreateAutoScaleVmGroupParams {
	P := &CreateAutoScaleVmGroupParams{}
	P.P = make(map[string]interface{})
	P.P["lbruleid"] = lbruleid
	P.P["maxmembers"] = maxmembers
	P.P["minmembers"] = minmembers
	P.P["scaledownpolicyids"] = scaledownpolicyids
	P.P["scaleuppolicyids"] = scaleuppolicyids
	P.P["vmprofileid"] = vmprofileid
	return P
}

// Creates and automatically starts a virtual machine based on a service offering, disk offering, and template.
func (s *AutoScaleService) CreateAutoScaleVmGroup(p *CreateAutoScaleVmGroupParams) (*CreateAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newRequest("createAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAutoScaleVmGroupResponse
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

type CreateAutoScaleVmGroupResponse struct {
	Account           string   `json:"account"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Fordisplay        bool     `json:"fordisplay"`
	Id                string   `json:"id"`
	Interval          int      `json:"interval"`
	JobID             string   `json:"jobid"`
	Jobstatus         int      `json:"jobstatus"`
	Lbruleid          string   `json:"lbruleid"`
	Maxmembers        int      `json:"maxmembers"`
	Minmembers        int      `json:"minmembers"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Scaledownpolicies []string `json:"scaledownpolicies"`
	Scaleuppolicies   []string `json:"scaleuppolicies"`
	State             string   `json:"state"`
	Vmprofileid       string   `json:"vmprofileid"`
}

type CreateAutoScaleVmProfileParams struct {
	P map[string]interface{}
}

func (P *CreateAutoScaleVmProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["autoscaleuserid"]; found {
		u.Set("autoscaleuserid", v.(string))
	}
	if v, found := P.P["counterparam"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("counterparam[%d].key", i), k)
			u.Set(fmt.Sprintf("counterparam[%d].value", i), m[k])
		}
	}
	if v, found := P.P["destroyvmgraceperiod"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("destroyvmgraceperiod", vv)
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["otherdeployparams"]; found {
		u.Set("otherdeployparams", v.(string))
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

func (P *CreateAutoScaleVmProfileParams) SetAutoscaleuserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["autoscaleuserid"] = v
}

func (P *CreateAutoScaleVmProfileParams) GetAutoscaleuserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["autoscaleuserid"].(string)
	return value, ok
}

func (P *CreateAutoScaleVmProfileParams) SetCounterparam(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["counterparam"] = v
}

func (P *CreateAutoScaleVmProfileParams) GetCounterparam() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["counterparam"].(map[string]string)
	return value, ok
}

func (P *CreateAutoScaleVmProfileParams) SetDestroyvmgraceperiod(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["destroyvmgraceperiod"] = v
}

func (P *CreateAutoScaleVmProfileParams) GetDestroyvmgraceperiod() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["destroyvmgraceperiod"].(int)
	return value, ok
}

func (P *CreateAutoScaleVmProfileParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateAutoScaleVmProfileParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateAutoScaleVmProfileParams) SetOtherdeployparams(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["otherdeployparams"] = v
}

func (P *CreateAutoScaleVmProfileParams) GetOtherdeployparams() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["otherdeployparams"].(string)
	return value, ok
}

func (P *CreateAutoScaleVmProfileParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *CreateAutoScaleVmProfileParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

func (P *CreateAutoScaleVmProfileParams) SetTemplateid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templateid"] = v
}

func (P *CreateAutoScaleVmProfileParams) GetTemplateid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templateid"].(string)
	return value, ok
}

func (P *CreateAutoScaleVmProfileParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateAutoScaleVmProfileParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateAutoScaleVmProfileParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateAutoScaleVmProfileParams(serviceofferingid string, templateid string, zoneid string) *CreateAutoScaleVmProfileParams {
	P := &CreateAutoScaleVmProfileParams{}
	P.P = make(map[string]interface{})
	P.P["serviceofferingid"] = serviceofferingid
	P.P["templateid"] = templateid
	P.P["zoneid"] = zoneid
	return P
}

// Creates a profile that contains information about the virtual machine which will be provisioned automatically by autoscale feature.
func (s *AutoScaleService) CreateAutoScaleVmProfile(p *CreateAutoScaleVmProfileParams) (*CreateAutoScaleVmProfileResponse, error) {
	resp, err := s.cs.newRequest("createAutoScaleVmProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateAutoScaleVmProfileResponse
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

type CreateAutoScaleVmProfileResponse struct {
	Account              string `json:"account"`
	Autoscaleuserid      string `json:"autoscaleuserid"`
	Destroyvmgraceperiod int    `json:"destroyvmgraceperiod"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Fordisplay           bool   `json:"fordisplay"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Otherdeployparams    string `json:"otherdeployparams"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Serviceofferingid    string `json:"serviceofferingid"`
	Templateid           string `json:"templateid"`
	Zoneid               string `json:"zoneid"`
}

type CreateConditionParams struct {
	P map[string]interface{}
}

func (P *CreateConditionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["counterid"]; found {
		u.Set("counterid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["relationaloperator"]; found {
		u.Set("relationaloperator", v.(string))
	}
	if v, found := P.P["threshold"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("threshold", vv)
	}
	return u
}

func (P *CreateConditionParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateConditionParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateConditionParams) SetCounterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["counterid"] = v
}

func (P *CreateConditionParams) GetCounterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["counterid"].(string)
	return value, ok
}

func (P *CreateConditionParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateConditionParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateConditionParams) SetRelationaloperator(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["relationaloperator"] = v
}

func (P *CreateConditionParams) GetRelationaloperator() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["relationaloperator"].(string)
	return value, ok
}

func (P *CreateConditionParams) SetThreshold(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["threshold"] = v
}

func (P *CreateConditionParams) GetThreshold() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["threshold"].(int64)
	return value, ok
}

// You should always use this function to get a new CreateConditionParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateConditionParams(counterid string, relationaloperator string, threshold int64) *CreateConditionParams {
	P := &CreateConditionParams{}
	P.P = make(map[string]interface{})
	P.P["counterid"] = counterid
	P.P["relationaloperator"] = relationaloperator
	P.P["threshold"] = threshold
	return P
}

// Creates a condition
func (s *AutoScaleService) CreateCondition(p *CreateConditionParams) (*CreateConditionResponse, error) {
	resp, err := s.cs.newRequest("createCondition", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateConditionResponse
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

type CreateConditionResponse struct {
	Account            string   `json:"account"`
	Counter            []string `json:"counter"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Id                 string   `json:"id"`
	JobID              string   `json:"jobid"`
	Jobstatus          int      `json:"jobstatus"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Relationaloperator string   `json:"relationaloperator"`
	Threshold          int64    `json:"threshold"`
	Zoneid             string   `json:"zoneid"`
}

type CreateCounterParams struct {
	P map[string]interface{}
}

func (P *CreateCounterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["source"]; found {
		u.Set("source", v.(string))
	}
	if v, found := P.P["value"]; found {
		u.Set("value", v.(string))
	}
	return u
}

func (P *CreateCounterParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateCounterParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateCounterParams) SetSource(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["source"] = v
}

func (P *CreateCounterParams) GetSource() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["source"].(string)
	return value, ok
}

func (P *CreateCounterParams) SetValue(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["value"] = v
}

func (P *CreateCounterParams) GetValue() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["value"].(string)
	return value, ok
}

// You should always use this function to get a new CreateCounterParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewCreateCounterParams(name string, source string, value string) *CreateCounterParams {
	P := &CreateCounterParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	P.P["source"] = source
	P.P["value"] = value
	return P
}

// Adds metric counter
func (s *AutoScaleService) CreateCounter(p *CreateCounterParams) (*CreateCounterResponse, error) {
	resp, err := s.cs.newRequest("createCounter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateCounterResponse
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

type CreateCounterResponse struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Source    string `json:"source"`
	Value     string `json:"value"`
	Zoneid    string `json:"zoneid"`
}

type DeleteAutoScalePolicyParams struct {
	P map[string]interface{}
}

func (P *DeleteAutoScalePolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteAutoScalePolicyParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteAutoScalePolicyParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAutoScalePolicyParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteAutoScalePolicyParams(id string) *DeleteAutoScalePolicyParams {
	P := &DeleteAutoScalePolicyParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a autoscale policy.
func (s *AutoScaleService) DeleteAutoScalePolicy(p *DeleteAutoScalePolicyParams) (*DeleteAutoScalePolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteAutoScalePolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAutoScalePolicyResponse
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

type DeleteAutoScalePolicyResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteAutoScaleVmGroupParams struct {
	P map[string]interface{}
}

func (P *DeleteAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteAutoScaleVmGroupParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteAutoScaleVmGroupParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteAutoScaleVmGroupParams(id string) *DeleteAutoScaleVmGroupParams {
	P := &DeleteAutoScaleVmGroupParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a autoscale vm group.
func (s *AutoScaleService) DeleteAutoScaleVmGroup(p *DeleteAutoScaleVmGroupParams) (*DeleteAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newRequest("deleteAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAutoScaleVmGroupResponse
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

type DeleteAutoScaleVmGroupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteAutoScaleVmProfileParams struct {
	P map[string]interface{}
}

func (P *DeleteAutoScaleVmProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteAutoScaleVmProfileParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteAutoScaleVmProfileParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAutoScaleVmProfileParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteAutoScaleVmProfileParams(id string) *DeleteAutoScaleVmProfileParams {
	P := &DeleteAutoScaleVmProfileParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a autoscale vm profile.
func (s *AutoScaleService) DeleteAutoScaleVmProfile(p *DeleteAutoScaleVmProfileParams) (*DeleteAutoScaleVmProfileResponse, error) {
	resp, err := s.cs.newRequest("deleteAutoScaleVmProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAutoScaleVmProfileResponse
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

type DeleteAutoScaleVmProfileResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteConditionParams struct {
	P map[string]interface{}
}

func (P *DeleteConditionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteConditionParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteConditionParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteConditionParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteConditionParams(id string) *DeleteConditionParams {
	P := &DeleteConditionParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Removes a condition
func (s *AutoScaleService) DeleteCondition(p *DeleteConditionParams) (*DeleteConditionResponse, error) {
	resp, err := s.cs.newRequest("deleteCondition", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteConditionResponse
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

type DeleteConditionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteCounterParams struct {
	P map[string]interface{}
}

func (P *DeleteCounterParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteCounterParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteCounterParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteCounterParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDeleteCounterParams(id string) *DeleteCounterParams {
	P := &DeleteCounterParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a counter
func (s *AutoScaleService) DeleteCounter(p *DeleteCounterParams) (*DeleteCounterResponse, error) {
	resp, err := s.cs.newRequest("deleteCounter", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteCounterResponse
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

type DeleteCounterResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DisableAutoScaleVmGroupParams struct {
	P map[string]interface{}
}

func (P *DisableAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DisableAutoScaleVmGroupParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DisableAutoScaleVmGroupParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DisableAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewDisableAutoScaleVmGroupParams(id string) *DisableAutoScaleVmGroupParams {
	P := &DisableAutoScaleVmGroupParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Disables an AutoScale Vm Group
func (s *AutoScaleService) DisableAutoScaleVmGroup(p *DisableAutoScaleVmGroupParams) (*DisableAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newRequest("disableAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableAutoScaleVmGroupResponse
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

type DisableAutoScaleVmGroupResponse struct {
	Account           string   `json:"account"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Fordisplay        bool     `json:"fordisplay"`
	Id                string   `json:"id"`
	Interval          int      `json:"interval"`
	JobID             string   `json:"jobid"`
	Jobstatus         int      `json:"jobstatus"`
	Lbruleid          string   `json:"lbruleid"`
	Maxmembers        int      `json:"maxmembers"`
	Minmembers        int      `json:"minmembers"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Scaledownpolicies []string `json:"scaledownpolicies"`
	Scaleuppolicies   []string `json:"scaleuppolicies"`
	State             string   `json:"state"`
	Vmprofileid       string   `json:"vmprofileid"`
}

type EnableAutoScaleVmGroupParams struct {
	P map[string]interface{}
}

func (P *EnableAutoScaleVmGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *EnableAutoScaleVmGroupParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *EnableAutoScaleVmGroupParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new EnableAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewEnableAutoScaleVmGroupParams(id string) *EnableAutoScaleVmGroupParams {
	P := &EnableAutoScaleVmGroupParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Enables an AutoScale Vm Group
func (s *AutoScaleService) EnableAutoScaleVmGroup(p *EnableAutoScaleVmGroupParams) (*EnableAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newRequest("enableAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableAutoScaleVmGroupResponse
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

type EnableAutoScaleVmGroupResponse struct {
	Account           string   `json:"account"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Fordisplay        bool     `json:"fordisplay"`
	Id                string   `json:"id"`
	Interval          int      `json:"interval"`
	JobID             string   `json:"jobid"`
	Jobstatus         int      `json:"jobstatus"`
	Lbruleid          string   `json:"lbruleid"`
	Maxmembers        int      `json:"maxmembers"`
	Minmembers        int      `json:"minmembers"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Scaledownpolicies []string `json:"scaledownpolicies"`
	Scaleuppolicies   []string `json:"scaleuppolicies"`
	State             string   `json:"state"`
	Vmprofileid       string   `json:"vmprofileid"`
}

type ListAutoScalePoliciesParams struct {
	P map[string]interface{}
}

func (P *ListAutoScalePoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := P.P["conditionid"]; found {
		u.Set("conditionid", v.(string))
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
	if v, found := P.P["vmgroupid"]; found {
		u.Set("vmgroupid", v.(string))
	}
	return u
}

func (P *ListAutoScalePoliciesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListAutoScalePoliciesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListAutoScalePoliciesParams) SetAction(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["action"] = v
}

func (P *ListAutoScalePoliciesParams) GetAction() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["action"].(string)
	return value, ok
}

func (P *ListAutoScalePoliciesParams) SetConditionid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["conditionid"] = v
}

func (P *ListAutoScalePoliciesParams) GetConditionid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["conditionid"].(string)
	return value, ok
}

func (P *ListAutoScalePoliciesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListAutoScalePoliciesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListAutoScalePoliciesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListAutoScalePoliciesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListAutoScalePoliciesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListAutoScalePoliciesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListAutoScalePoliciesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListAutoScalePoliciesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListAutoScalePoliciesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListAutoScalePoliciesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListAutoScalePoliciesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListAutoScalePoliciesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListAutoScalePoliciesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListAutoScalePoliciesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListAutoScalePoliciesParams) SetVmgroupid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmgroupid"] = v
}

func (P *ListAutoScalePoliciesParams) GetVmgroupid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmgroupid"].(string)
	return value, ok
}

// You should always use this function to get a new ListAutoScalePoliciesParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListAutoScalePoliciesParams() *ListAutoScalePoliciesParams {
	P := &ListAutoScalePoliciesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScalePolicyByID(id string, opts ...OptionFunc) (*AutoScalePolicy, int, error) {
	P := &ListAutoScalePoliciesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAutoScalePolicies(P)
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
		return l.AutoScalePolicies[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for AutoScalePolicy UUID: %s!", id)
}

// Lists autoscale policies.
func (s *AutoScaleService) ListAutoScalePolicies(p *ListAutoScalePoliciesParams) (*ListAutoScalePoliciesResponse, error) {
	resp, err := s.cs.newRequest("listAutoScalePolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAutoScalePoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAutoScalePoliciesResponse struct {
	Count             int                `json:"count"`
	AutoScalePolicies []*AutoScalePolicy `json:"autoscalepolicy"`
}

type AutoScalePolicy struct {
	Account    string   `json:"account"`
	Action     string   `json:"action"`
	Conditions []string `json:"conditions"`
	Domain     string   `json:"domain"`
	Domainid   string   `json:"domainid"`
	Duration   int      `json:"duration"`
	Id         string   `json:"id"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Project    string   `json:"project"`
	Projectid  string   `json:"projectid"`
	Quiettime  int      `json:"quiettime"`
}

type ListAutoScaleVmGroupsParams struct {
	P map[string]interface{}
}

func (P *ListAutoScaleVmGroupsParams) toURLValues() url.Values {
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
	if v, found := P.P["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
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
	if v, found := P.P["policyid"]; found {
		u.Set("policyid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["vmprofileid"]; found {
		u.Set("vmprofileid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListAutoScaleVmGroupsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetLbruleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["lbruleid"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetLbruleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["lbruleid"].(string)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetPolicyid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["policyid"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetPolicyid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["policyid"].(string)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetVmprofileid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmprofileid"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetVmprofileid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmprofileid"].(string)
	return value, ok
}

func (P *ListAutoScaleVmGroupsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListAutoScaleVmGroupsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListAutoScaleVmGroupsParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListAutoScaleVmGroupsParams() *ListAutoScaleVmGroupsParams {
	P := &ListAutoScaleVmGroupsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScaleVmGroupByID(id string, opts ...OptionFunc) (*AutoScaleVmGroup, int, error) {
	P := &ListAutoScaleVmGroupsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAutoScaleVmGroups(P)
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
		return l.AutoScaleVmGroups[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for AutoScaleVmGroup UUID: %s!", id)
}

// Lists autoscale vm groups.
func (s *AutoScaleService) ListAutoScaleVmGroups(p *ListAutoScaleVmGroupsParams) (*ListAutoScaleVmGroupsResponse, error) {
	resp, err := s.cs.newRequest("listAutoScaleVmGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAutoScaleVmGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAutoScaleVmGroupsResponse struct {
	Count             int                 `json:"count"`
	AutoScaleVmGroups []*AutoScaleVmGroup `json:"autoscalevmgroup"`
}

type AutoScaleVmGroup struct {
	Account           string   `json:"account"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Fordisplay        bool     `json:"fordisplay"`
	Id                string   `json:"id"`
	Interval          int      `json:"interval"`
	JobID             string   `json:"jobid"`
	Jobstatus         int      `json:"jobstatus"`
	Lbruleid          string   `json:"lbruleid"`
	Maxmembers        int      `json:"maxmembers"`
	Minmembers        int      `json:"minmembers"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Scaledownpolicies []string `json:"scaledownpolicies"`
	Scaleuppolicies   []string `json:"scaleuppolicies"`
	State             string   `json:"state"`
	Vmprofileid       string   `json:"vmprofileid"`
}

type ListAutoScaleVmProfilesParams struct {
	P map[string]interface{}
}

func (P *ListAutoScaleVmProfilesParams) toURLValues() url.Values {
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
	if v, found := P.P["otherdeployparams"]; found {
		u.Set("otherdeployparams", v.(string))
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

func (P *ListAutoScaleVmProfilesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetOtherdeployparams(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["otherdeployparams"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetOtherdeployparams() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["otherdeployparams"].(string)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetTemplateid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templateid"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetTemplateid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templateid"].(string)
	return value, ok
}

func (P *ListAutoScaleVmProfilesParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListAutoScaleVmProfilesParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListAutoScaleVmProfilesParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListAutoScaleVmProfilesParams() *ListAutoScaleVmProfilesParams {
	P := &ListAutoScaleVmProfilesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetAutoScaleVmProfileByID(id string, opts ...OptionFunc) (*AutoScaleVmProfile, int, error) {
	P := &ListAutoScaleVmProfilesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAutoScaleVmProfiles(P)
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
		return l.AutoScaleVmProfiles[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for AutoScaleVmProfile UUID: %s!", id)
}

// Lists autoscale vm profiles.
func (s *AutoScaleService) ListAutoScaleVmProfiles(p *ListAutoScaleVmProfilesParams) (*ListAutoScaleVmProfilesResponse, error) {
	resp, err := s.cs.newRequest("listAutoScaleVmProfiles", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAutoScaleVmProfilesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAutoScaleVmProfilesResponse struct {
	Count               int                   `json:"count"`
	AutoScaleVmProfiles []*AutoScaleVmProfile `json:"autoscalevmprofile"`
}

type AutoScaleVmProfile struct {
	Account              string `json:"account"`
	Autoscaleuserid      string `json:"autoscaleuserid"`
	Destroyvmgraceperiod int    `json:"destroyvmgraceperiod"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Fordisplay           bool   `json:"fordisplay"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Otherdeployparams    string `json:"otherdeployparams"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Serviceofferingid    string `json:"serviceofferingid"`
	Templateid           string `json:"templateid"`
	Zoneid               string `json:"zoneid"`
}

type ListConditionsParams struct {
	P map[string]interface{}
}

func (P *ListConditionsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["counterid"]; found {
		u.Set("counterid", v.(string))
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
	if v, found := P.P["policyid"]; found {
		u.Set("policyid", v.(string))
	}
	return u
}

func (P *ListConditionsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListConditionsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListConditionsParams) SetCounterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["counterid"] = v
}

func (P *ListConditionsParams) GetCounterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["counterid"].(string)
	return value, ok
}

func (P *ListConditionsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListConditionsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListConditionsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListConditionsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListConditionsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListConditionsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListConditionsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListConditionsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListConditionsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListConditionsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListConditionsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListConditionsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListConditionsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListConditionsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListConditionsParams) SetPolicyid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["policyid"] = v
}

func (P *ListConditionsParams) GetPolicyid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["policyid"].(string)
	return value, ok
}

// You should always use this function to get a new ListConditionsParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListConditionsParams() *ListConditionsParams {
	P := &ListConditionsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetConditionByID(id string, opts ...OptionFunc) (*Condition, int, error) {
	P := &ListConditionsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListConditions(P)
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
		return l.Conditions[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Condition UUID: %s!", id)
}

// List Conditions for the specific user
func (s *AutoScaleService) ListConditions(p *ListConditionsParams) (*ListConditionsResponse, error) {
	resp, err := s.cs.newRequest("listConditions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListConditionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListConditionsResponse struct {
	Count      int          `json:"count"`
	Conditions []*Condition `json:"condition"`
}

type Condition struct {
	Account            string   `json:"account"`
	Counter            []string `json:"counter"`
	Domain             string   `json:"domain"`
	Domainid           string   `json:"domainid"`
	Id                 string   `json:"id"`
	JobID              string   `json:"jobid"`
	Jobstatus          int      `json:"jobstatus"`
	Project            string   `json:"project"`
	Projectid          string   `json:"projectid"`
	Relationaloperator string   `json:"relationaloperator"`
	Threshold          int64    `json:"threshold"`
	Zoneid             string   `json:"zoneid"`
}

type ListCountersParams struct {
	P map[string]interface{}
}

func (P *ListCountersParams) toURLValues() url.Values {
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
	if v, found := P.P["source"]; found {
		u.Set("source", v.(string))
	}
	return u
}

func (P *ListCountersParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListCountersParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListCountersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListCountersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListCountersParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListCountersParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListCountersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListCountersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListCountersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListCountersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListCountersParams) SetSource(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["source"] = v
}

func (P *ListCountersParams) GetSource() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["source"].(string)
	return value, ok
}

// You should always use this function to get a new ListCountersParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewListCountersParams() *ListCountersParams {
	P := &ListCountersParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetCounterID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListCountersParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListCounters(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Counters[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Counters {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetCounterByName(name string, opts ...OptionFunc) (*Counter, int, error) {
	id, count, err := s.GetCounterID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetCounterByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AutoScaleService) GetCounterByID(id string, opts ...OptionFunc) (*Counter, int, error) {
	P := &ListCountersParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListCounters(P)
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
		return l.Counters[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Counter UUID: %s!", id)
}

// List the counters
func (s *AutoScaleService) ListCounters(p *ListCountersParams) (*ListCountersResponse, error) {
	resp, err := s.cs.newRequest("listCounters", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCountersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCountersResponse struct {
	Count    int        `json:"count"`
	Counters []*Counter `json:"counter"`
}

type Counter struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Source    string `json:"source"`
	Value     string `json:"value"`
	Zoneid    string `json:"zoneid"`
}

type UpdateAutoScalePolicyParams struct {
	P map[string]interface{}
}

func (P *UpdateAutoScalePolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["conditionids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("conditionids", vv)
	}
	if v, found := P.P["duration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("duration", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["quiettime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("quiettime", vv)
	}
	return u
}

func (P *UpdateAutoScalePolicyParams) SetConditionids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["conditionids"] = v
}

func (P *UpdateAutoScalePolicyParams) GetConditionids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["conditionids"].([]string)
	return value, ok
}

func (P *UpdateAutoScalePolicyParams) SetDuration(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["duration"] = v
}

func (P *UpdateAutoScalePolicyParams) GetDuration() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["duration"].(int)
	return value, ok
}

func (P *UpdateAutoScalePolicyParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateAutoScalePolicyParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateAutoScalePolicyParams) SetQuiettime(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["quiettime"] = v
}

func (P *UpdateAutoScalePolicyParams) GetQuiettime() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["quiettime"].(int)
	return value, ok
}

// You should always use this function to get a new UpdateAutoScalePolicyParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewUpdateAutoScalePolicyParams(id string) *UpdateAutoScalePolicyParams {
	P := &UpdateAutoScalePolicyParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates an existing autoscale policy.
func (s *AutoScaleService) UpdateAutoScalePolicy(p *UpdateAutoScalePolicyParams) (*UpdateAutoScalePolicyResponse, error) {
	resp, err := s.cs.newRequest("updateAutoScalePolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAutoScalePolicyResponse
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

type UpdateAutoScalePolicyResponse struct {
	Account    string   `json:"account"`
	Action     string   `json:"action"`
	Conditions []string `json:"conditions"`
	Domain     string   `json:"domain"`
	Domainid   string   `json:"domainid"`
	Duration   int      `json:"duration"`
	Id         string   `json:"id"`
	JobID      string   `json:"jobid"`
	Jobstatus  int      `json:"jobstatus"`
	Project    string   `json:"project"`
	Projectid  string   `json:"projectid"`
	Quiettime  int      `json:"quiettime"`
}

type UpdateAutoScaleVmGroupParams struct {
	P map[string]interface{}
}

func (P *UpdateAutoScaleVmGroupParams) toURLValues() url.Values {
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
	if v, found := P.P["interval"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("interval", vv)
	}
	if v, found := P.P["maxmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxmembers", vv)
	}
	if v, found := P.P["minmembers"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minmembers", vv)
	}
	if v, found := P.P["scaledownpolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaledownpolicyids", vv)
	}
	if v, found := P.P["scaleuppolicyids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("scaleuppolicyids", vv)
	}
	return u
}

func (P *UpdateAutoScaleVmGroupParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateAutoScaleVmGroupParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateAutoScaleVmGroupParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateAutoScaleVmGroupParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateAutoScaleVmGroupParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateAutoScaleVmGroupParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateAutoScaleVmGroupParams) SetInterval(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["interval"] = v
}

func (P *UpdateAutoScaleVmGroupParams) GetInterval() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["interval"].(int)
	return value, ok
}

func (P *UpdateAutoScaleVmGroupParams) SetMaxmembers(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxmembers"] = v
}

func (P *UpdateAutoScaleVmGroupParams) GetMaxmembers() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxmembers"].(int)
	return value, ok
}

func (P *UpdateAutoScaleVmGroupParams) SetMinmembers(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["minmembers"] = v
}

func (P *UpdateAutoScaleVmGroupParams) GetMinmembers() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["minmembers"].(int)
	return value, ok
}

func (P *UpdateAutoScaleVmGroupParams) SetScaledownpolicyids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["scaledownpolicyids"] = v
}

func (P *UpdateAutoScaleVmGroupParams) GetScaledownpolicyids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["scaledownpolicyids"].([]string)
	return value, ok
}

func (P *UpdateAutoScaleVmGroupParams) SetScaleuppolicyids(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["scaleuppolicyids"] = v
}

func (P *UpdateAutoScaleVmGroupParams) GetScaleuppolicyids() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["scaleuppolicyids"].([]string)
	return value, ok
}

// You should always use this function to get a new UpdateAutoScaleVmGroupParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewUpdateAutoScaleVmGroupParams(id string) *UpdateAutoScaleVmGroupParams {
	P := &UpdateAutoScaleVmGroupParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates an existing autoscale vm group.
func (s *AutoScaleService) UpdateAutoScaleVmGroup(p *UpdateAutoScaleVmGroupParams) (*UpdateAutoScaleVmGroupResponse, error) {
	resp, err := s.cs.newRequest("updateAutoScaleVmGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAutoScaleVmGroupResponse
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

type UpdateAutoScaleVmGroupResponse struct {
	Account           string   `json:"account"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Fordisplay        bool     `json:"fordisplay"`
	Id                string   `json:"id"`
	Interval          int      `json:"interval"`
	JobID             string   `json:"jobid"`
	Jobstatus         int      `json:"jobstatus"`
	Lbruleid          string   `json:"lbruleid"`
	Maxmembers        int      `json:"maxmembers"`
	Minmembers        int      `json:"minmembers"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Scaledownpolicies []string `json:"scaledownpolicies"`
	Scaleuppolicies   []string `json:"scaleuppolicies"`
	State             string   `json:"state"`
	Vmprofileid       string   `json:"vmprofileid"`
}

type UpdateAutoScaleVmProfileParams struct {
	P map[string]interface{}
}

func (P *UpdateAutoScaleVmProfileParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["autoscaleuserid"]; found {
		u.Set("autoscaleuserid", v.(string))
	}
	if v, found := P.P["counterparam"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("counterparam[%d].key", i), k)
			u.Set(fmt.Sprintf("counterparam[%d].value", i), m[k])
		}
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["destroyvmgraceperiod"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("destroyvmgraceperiod", vv)
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	return u
}

func (P *UpdateAutoScaleVmProfileParams) SetAutoscaleuserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["autoscaleuserid"] = v
}

func (P *UpdateAutoScaleVmProfileParams) GetAutoscaleuserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["autoscaleuserid"].(string)
	return value, ok
}

func (P *UpdateAutoScaleVmProfileParams) SetCounterparam(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["counterparam"] = v
}

func (P *UpdateAutoScaleVmProfileParams) GetCounterparam() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["counterparam"].(map[string]string)
	return value, ok
}

func (P *UpdateAutoScaleVmProfileParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateAutoScaleVmProfileParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateAutoScaleVmProfileParams) SetDestroyvmgraceperiod(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["destroyvmgraceperiod"] = v
}

func (P *UpdateAutoScaleVmProfileParams) GetDestroyvmgraceperiod() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["destroyvmgraceperiod"].(int)
	return value, ok
}

func (P *UpdateAutoScaleVmProfileParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateAutoScaleVmProfileParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateAutoScaleVmProfileParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateAutoScaleVmProfileParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateAutoScaleVmProfileParams) SetTemplateid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["templateid"] = v
}

func (P *UpdateAutoScaleVmProfileParams) GetTemplateid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["templateid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateAutoScaleVmProfileParams instance,
// as then you are sure you have configured all required params
func (s *AutoScaleService) NewUpdateAutoScaleVmProfileParams(id string) *UpdateAutoScaleVmProfileParams {
	P := &UpdateAutoScaleVmProfileParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates an existing autoscale vm profile.
func (s *AutoScaleService) UpdateAutoScaleVmProfile(p *UpdateAutoScaleVmProfileParams) (*UpdateAutoScaleVmProfileResponse, error) {
	resp, err := s.cs.newRequest("updateAutoScaleVmProfile", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAutoScaleVmProfileResponse
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

type UpdateAutoScaleVmProfileResponse struct {
	Account              string `json:"account"`
	Autoscaleuserid      string `json:"autoscaleuserid"`
	Destroyvmgraceperiod int    `json:"destroyvmgraceperiod"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Fordisplay           bool   `json:"fordisplay"`
	Id                   string `json:"id"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Otherdeployparams    string `json:"otherdeployparams"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Serviceofferingid    string `json:"serviceofferingid"`
	Templateid           string `json:"templateid"`
	Zoneid               string `json:"zoneid"`
}
