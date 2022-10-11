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
)

type UsageServiceIface interface {
	AddTrafficMonitor(p *AddTrafficMonitorParams) (*AddTrafficMonitorResponse, error)
	NewAddTrafficMonitorParams(url string, zoneid string) *AddTrafficMonitorParams
	AddTrafficType(p *AddTrafficTypeParams) (*AddTrafficTypeResponse, error)
	NewAddTrafficTypeParams(physicalnetworkid string, traffictype string) *AddTrafficTypeParams
	DeleteTrafficMonitor(p *DeleteTrafficMonitorParams) (*DeleteTrafficMonitorResponse, error)
	NewDeleteTrafficMonitorParams(id string) *DeleteTrafficMonitorParams
	DeleteTrafficType(p *DeleteTrafficTypeParams) (*DeleteTrafficTypeResponse, error)
	NewDeleteTrafficTypeParams(id string) *DeleteTrafficTypeParams
	GenerateUsageRecords(p *GenerateUsageRecordsParams) (*GenerateUsageRecordsResponse, error)
	NewGenerateUsageRecordsParams(enddate string, startdate string) *GenerateUsageRecordsParams
	ListTrafficMonitors(p *ListTrafficMonitorsParams) (*ListTrafficMonitorsResponse, error)
	NewListTrafficMonitorsParams(zoneid string) *ListTrafficMonitorsParams
	ListTrafficTypeImplementors(p *ListTrafficTypeImplementorsParams) (*ListTrafficTypeImplementorsResponse, error)
	NewListTrafficTypeImplementorsParams() *ListTrafficTypeImplementorsParams
	ListTrafficTypes(p *ListTrafficTypesParams) (*ListTrafficTypesResponse, error)
	NewListTrafficTypesParams(physicalnetworkid string) *ListTrafficTypesParams
	GetTrafficTypeID(keyword string, physicalnetworkid string, opts ...OptionFunc) (string, int, error)
	ListUsageRecords(p *ListUsageRecordsParams) (*ListUsageRecordsResponse, error)
	NewListUsageRecordsParams(enddate string, startdate string) *ListUsageRecordsParams
	ListUsageTypes(p *ListUsageTypesParams) (*ListUsageTypesResponse, error)
	NewListUsageTypesParams() *ListUsageTypesParams
	RemoveRawUsageRecords(p *RemoveRawUsageRecordsParams) (*RemoveRawUsageRecordsResponse, error)
	NewRemoveRawUsageRecordsParams(interval int) *RemoveRawUsageRecordsParams
	UpdateTrafficType(p *UpdateTrafficTypeParams) (*UpdateTrafficTypeResponse, error)
	NewUpdateTrafficTypeParams(id string) *UpdateTrafficTypeParams
}

type AddTrafficMonitorParams struct {
	P map[string]interface{}
}

func (P *AddTrafficMonitorParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["excludezones"]; found {
		u.Set("excludezones", v.(string))
	}
	if v, found := P.P["includezones"]; found {
		u.Set("includezones", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *AddTrafficMonitorParams) SetExcludezones(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["excludezones"] = v
}

func (P *AddTrafficMonitorParams) GetExcludezones() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["excludezones"].(string)
	return value, ok
}

func (P *AddTrafficMonitorParams) SetIncludezones(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["includezones"] = v
}

func (P *AddTrafficMonitorParams) GetIncludezones() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["includezones"].(string)
	return value, ok
}

func (P *AddTrafficMonitorParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddTrafficMonitorParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddTrafficMonitorParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *AddTrafficMonitorParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddTrafficMonitorParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewAddTrafficMonitorParams(url string, zoneid string) *AddTrafficMonitorParams {
	P := &AddTrafficMonitorParams{}
	P.P = make(map[string]interface{})
	P.P["url"] = url
	P.P["zoneid"] = zoneid
	return P
}

// Adds Traffic Monitor Host for Direct Network Usage
func (s *UsageService) AddTrafficMonitor(p *AddTrafficMonitorParams) (*AddTrafficMonitorResponse, error) {
	resp, err := s.cs.newRequest("addTrafficMonitor", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddTrafficMonitorResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddTrafficMonitorResponse struct {
	Id         string `json:"id"`
	Ipaddress  string `json:"ipaddress"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Numretries string `json:"numretries"`
	Timeout    string `json:"timeout"`
	Zoneid     string `json:"zoneid"`
}

type AddTrafficTypeParams struct {
	P map[string]interface{}
}

func (P *AddTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hypervnetworklabel"]; found {
		u.Set("hypervnetworklabel", v.(string))
	}
	if v, found := P.P["isolationmethod"]; found {
		u.Set("isolationmethod", v.(string))
	}
	if v, found := P.P["kvmnetworklabel"]; found {
		u.Set("kvmnetworklabel", v.(string))
	}
	if v, found := P.P["ovm3networklabel"]; found {
		u.Set("ovm3networklabel", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := P.P["vlan"]; found {
		u.Set("vlan", v.(string))
	}
	if v, found := P.P["vmwarenetworklabel"]; found {
		u.Set("vmwarenetworklabel", v.(string))
	}
	if v, found := P.P["xennetworklabel"]; found {
		u.Set("xennetworklabel", v.(string))
	}
	return u
}

func (P *AddTrafficTypeParams) SetHypervnetworklabel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervnetworklabel"] = v
}

func (P *AddTrafficTypeParams) GetHypervnetworklabel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervnetworklabel"].(string)
	return value, ok
}

func (P *AddTrafficTypeParams) SetIsolationmethod(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isolationmethod"] = v
}

func (P *AddTrafficTypeParams) GetIsolationmethod() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isolationmethod"].(string)
	return value, ok
}

func (P *AddTrafficTypeParams) SetKvmnetworklabel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["kvmnetworklabel"] = v
}

func (P *AddTrafficTypeParams) GetKvmnetworklabel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["kvmnetworklabel"].(string)
	return value, ok
}

func (P *AddTrafficTypeParams) SetOvm3networklabel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ovm3networklabel"] = v
}

func (P *AddTrafficTypeParams) GetOvm3networklabel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ovm3networklabel"].(string)
	return value, ok
}

func (P *AddTrafficTypeParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddTrafficTypeParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddTrafficTypeParams) SetTraffictype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["traffictype"] = v
}

func (P *AddTrafficTypeParams) GetTraffictype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["traffictype"].(string)
	return value, ok
}

func (P *AddTrafficTypeParams) SetVlan(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vlan"] = v
}

func (P *AddTrafficTypeParams) GetVlan() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vlan"].(string)
	return value, ok
}

func (P *AddTrafficTypeParams) SetVmwarenetworklabel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmwarenetworklabel"] = v
}

func (P *AddTrafficTypeParams) GetVmwarenetworklabel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmwarenetworklabel"].(string)
	return value, ok
}

func (P *AddTrafficTypeParams) SetXennetworklabel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["xennetworklabel"] = v
}

func (P *AddTrafficTypeParams) GetXennetworklabel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["xennetworklabel"].(string)
	return value, ok
}

// You should always use this function to get a new AddTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewAddTrafficTypeParams(physicalnetworkid string, traffictype string) *AddTrafficTypeParams {
	P := &AddTrafficTypeParams{}
	P.P = make(map[string]interface{})
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["traffictype"] = traffictype
	return P
}

// Adds traffic type to a physical network
func (s *UsageService) AddTrafficType(p *AddTrafficTypeParams) (*AddTrafficTypeResponse, error) {
	resp, err := s.cs.newRequest("addTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddTrafficTypeResponse
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

type AddTrafficTypeResponse struct {
	Hypervnetworklabel string `json:"hypervnetworklabel"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Kvmnetworklabel    string `json:"kvmnetworklabel"`
	Ovm3networklabel   string `json:"ovm3networklabel"`
	Physicalnetworkid  string `json:"physicalnetworkid"`
	Traffictype        string `json:"traffictype"`
	Vmwarenetworklabel string `json:"vmwarenetworklabel"`
	Xennetworklabel    string `json:"xennetworklabel"`
}

type DeleteTrafficMonitorParams struct {
	P map[string]interface{}
}

func (P *DeleteTrafficMonitorParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteTrafficMonitorParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteTrafficMonitorParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTrafficMonitorParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewDeleteTrafficMonitorParams(id string) *DeleteTrafficMonitorParams {
	P := &DeleteTrafficMonitorParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes an traffic monitor host.
func (s *UsageService) DeleteTrafficMonitor(p *DeleteTrafficMonitorParams) (*DeleteTrafficMonitorResponse, error) {
	resp, err := s.cs.newRequest("deleteTrafficMonitor", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTrafficMonitorResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteTrafficMonitorResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteTrafficMonitorResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteTrafficMonitorResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteTrafficTypeParams struct {
	P map[string]interface{}
}

func (P *DeleteTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteTrafficTypeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteTrafficTypeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewDeleteTrafficTypeParams(id string) *DeleteTrafficTypeParams {
	P := &DeleteTrafficTypeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes traffic type of a physical network
func (s *UsageService) DeleteTrafficType(p *DeleteTrafficTypeParams) (*DeleteTrafficTypeResponse, error) {
	resp, err := s.cs.newRequest("deleteTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTrafficTypeResponse
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

type DeleteTrafficTypeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type GenerateUsageRecordsParams struct {
	P map[string]interface{}
}

func (P *GenerateUsageRecordsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := P.P["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (P *GenerateUsageRecordsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *GenerateUsageRecordsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *GenerateUsageRecordsParams) SetEnddate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enddate"] = v
}

func (P *GenerateUsageRecordsParams) GetEnddate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enddate"].(string)
	return value, ok
}

func (P *GenerateUsageRecordsParams) SetStartdate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startdate"] = v
}

func (P *GenerateUsageRecordsParams) GetStartdate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startdate"].(string)
	return value, ok
}

// You should always use this function to get a new GenerateUsageRecordsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewGenerateUsageRecordsParams(enddate string, startdate string) *GenerateUsageRecordsParams {
	P := &GenerateUsageRecordsParams{}
	P.P = make(map[string]interface{})
	P.P["enddate"] = enddate
	P.P["startdate"] = startdate
	return P
}

// Generates usage records. This will generate records only if there any records to be generated, i.e if the scheduled usage job was not run or failed
func (s *UsageService) GenerateUsageRecords(p *GenerateUsageRecordsParams) (*GenerateUsageRecordsResponse, error) {
	resp, err := s.cs.newRequest("generateUsageRecords", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GenerateUsageRecordsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GenerateUsageRecordsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *GenerateUsageRecordsResponse) UnmarshalJSON(b []byte) error {
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

	type alias GenerateUsageRecordsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListTrafficMonitorsParams struct {
	P map[string]interface{}
}

func (P *ListTrafficMonitorsParams) toURLValues() url.Values {
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
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListTrafficMonitorsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListTrafficMonitorsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListTrafficMonitorsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListTrafficMonitorsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListTrafficMonitorsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListTrafficMonitorsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListTrafficMonitorsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListTrafficMonitorsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTrafficMonitorsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListTrafficMonitorsParams(zoneid string) *ListTrafficMonitorsParams {
	P := &ListTrafficMonitorsParams{}
	P.P = make(map[string]interface{})
	P.P["zoneid"] = zoneid
	return P
}

// List traffic monitor Hosts.
func (s *UsageService) ListTrafficMonitors(p *ListTrafficMonitorsParams) (*ListTrafficMonitorsResponse, error) {
	resp, err := s.cs.newRequest("listTrafficMonitors", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTrafficMonitorsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTrafficMonitorsResponse struct {
	Count           int               `json:"count"`
	TrafficMonitors []*TrafficMonitor `json:"trafficmonitor"`
}

type TrafficMonitor struct {
	Id         string `json:"id"`
	Ipaddress  string `json:"ipaddress"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Numretries string `json:"numretries"`
	Timeout    string `json:"timeout"`
	Zoneid     string `json:"zoneid"`
}

type ListTrafficTypeImplementorsParams struct {
	P map[string]interface{}
}

func (P *ListTrafficTypeImplementorsParams) toURLValues() url.Values {
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
	if v, found := P.P["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (P *ListTrafficTypeImplementorsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListTrafficTypeImplementorsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListTrafficTypeImplementorsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListTrafficTypeImplementorsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListTrafficTypeImplementorsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListTrafficTypeImplementorsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListTrafficTypeImplementorsParams) SetTraffictype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["traffictype"] = v
}

func (P *ListTrafficTypeImplementorsParams) GetTraffictype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["traffictype"].(string)
	return value, ok
}

// You should always use this function to get a new ListTrafficTypeImplementorsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListTrafficTypeImplementorsParams() *ListTrafficTypeImplementorsParams {
	P := &ListTrafficTypeImplementorsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists implementors of implementor of a network traffic type or implementors of all network traffic types
func (s *UsageService) ListTrafficTypeImplementors(p *ListTrafficTypeImplementorsParams) (*ListTrafficTypeImplementorsResponse, error) {
	resp, err := s.cs.newRequest("listTrafficTypeImplementors", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTrafficTypeImplementorsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTrafficTypeImplementorsResponse struct {
	Count                   int                       `json:"count"`
	TrafficTypeImplementors []*TrafficTypeImplementor `json:"traffictypeimplementor"`
}

type TrafficTypeImplementor struct {
	JobID                  string `json:"jobid"`
	Jobstatus              int    `json:"jobstatus"`
	Traffictype            string `json:"traffictype"`
	Traffictypeimplementor string `json:"traffictypeimplementor"`
}

type ListTrafficTypesParams struct {
	P map[string]interface{}
}

func (P *ListTrafficTypesParams) toURLValues() url.Values {
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
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (P *ListTrafficTypesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListTrafficTypesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListTrafficTypesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListTrafficTypesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListTrafficTypesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListTrafficTypesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListTrafficTypesParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListTrafficTypesParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTrafficTypesParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListTrafficTypesParams(physicalnetworkid string) *ListTrafficTypesParams {
	P := &ListTrafficTypesParams{}
	P.P = make(map[string]interface{})
	P.P["physicalnetworkid"] = physicalnetworkid
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *UsageService) GetTrafficTypeID(keyword string, physicalnetworkid string, opts ...OptionFunc) (string, int, error) {
	P := &ListTrafficTypesParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword
	P.P["physicalnetworkid"] = physicalnetworkid

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListTrafficTypes(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.TrafficTypes[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.TrafficTypes {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// Lists traffic types of a given physical network.
func (s *UsageService) ListTrafficTypes(p *ListTrafficTypesParams) (*ListTrafficTypesResponse, error) {
	resp, err := s.cs.newRequest("listTrafficTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTrafficTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTrafficTypesResponse struct {
	Count        int            `json:"count"`
	TrafficTypes []*TrafficType `json:"traffictype"`
}

type TrafficType struct {
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

type ListUsageRecordsParams struct {
	P map[string]interface{}
}

func (P *ListUsageRecordsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := P.P["includetags"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("includetags", vv)
	}
	if v, found := P.P["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["oldformat"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("oldformat", vv)
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
	if v, found := P.P["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := P.P["type"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("type", vv)
	}
	if v, found := P.P["usageid"]; found {
		u.Set("usageid", v.(string))
	}
	return u
}

func (P *ListUsageRecordsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListUsageRecordsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListUsageRecordsParams) SetAccountid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountid"] = v
}

func (P *ListUsageRecordsParams) GetAccountid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountid"].(string)
	return value, ok
}

func (P *ListUsageRecordsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListUsageRecordsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListUsageRecordsParams) SetEnddate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enddate"] = v
}

func (P *ListUsageRecordsParams) GetEnddate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enddate"].(string)
	return value, ok
}

func (P *ListUsageRecordsParams) SetIncludetags(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["includetags"] = v
}

func (P *ListUsageRecordsParams) GetIncludetags() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["includetags"].(bool)
	return value, ok
}

func (P *ListUsageRecordsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListUsageRecordsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListUsageRecordsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListUsageRecordsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListUsageRecordsParams) SetOldformat(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["oldformat"] = v
}

func (P *ListUsageRecordsParams) GetOldformat() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["oldformat"].(bool)
	return value, ok
}

func (P *ListUsageRecordsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListUsageRecordsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListUsageRecordsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListUsageRecordsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListUsageRecordsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListUsageRecordsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListUsageRecordsParams) SetStartdate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startdate"] = v
}

func (P *ListUsageRecordsParams) GetStartdate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startdate"].(string)
	return value, ok
}

func (P *ListUsageRecordsParams) SetType(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ListUsageRecordsParams) GetType() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(int64)
	return value, ok
}

func (P *ListUsageRecordsParams) SetUsageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["usageid"] = v
}

func (P *ListUsageRecordsParams) GetUsageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["usageid"].(string)
	return value, ok
}

// You should always use this function to get a new ListUsageRecordsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListUsageRecordsParams(enddate string, startdate string) *ListUsageRecordsParams {
	P := &ListUsageRecordsParams{}
	P.P = make(map[string]interface{})
	P.P["enddate"] = enddate
	P.P["startdate"] = startdate
	return P
}

// Lists usage records for accounts
func (s *UsageService) ListUsageRecords(p *ListUsageRecordsParams) (*ListUsageRecordsResponse, error) {
	resp, err := s.cs.newRequest("listUsageRecords", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsageRecordsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUsageRecordsResponse struct {
	Count        int            `json:"count"`
	UsageRecords []*UsageRecord `json:"usagerecord"`
}

type UsageRecord struct {
	Account          string `json:"account"`
	Accountid        string `json:"accountid"`
	Cpunumber        int64  `json:"cpunumber"`
	Cpuspeed         int64  `json:"cpuspeed"`
	Description      string `json:"description"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Enddate          string `json:"enddate"`
	Hasannotations   bool   `json:"hasannotations"`
	Isdefault        bool   `json:"isdefault"`
	Issourcenat      bool   `json:"issourcenat"`
	Issystem         bool   `json:"issystem"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Memory           int64  `json:"memory"`
	Name             string `json:"name"`
	Networkid        string `json:"networkid"`
	Offeringid       string `json:"offeringid"`
	Oscategoryid     string `json:"oscategoryid"`
	Oscategoryname   string `json:"oscategoryname"`
	Osdisplayname    string `json:"osdisplayname"`
	Ostypeid         string `json:"ostypeid"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Rawusage         string `json:"rawusage"`
	Size             int64  `json:"size"`
	Startdate        string `json:"startdate"`
	Tags             []Tags `json:"tags"`
	Templateid       string `json:"templateid"`
	Type             string `json:"type"`
	Usage            string `json:"usage"`
	Usageid          string `json:"usageid"`
	Usagetype        int    `json:"usagetype"`
	Virtualmachineid string `json:"virtualmachineid"`
	Virtualsize      int64  `json:"virtualsize"`
	Vpcid            string `json:"vpcid"`
	Zoneid           string `json:"zoneid"`
}

func (r *UsageRecord) UnmarshalJSON(b []byte) error {
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

	type alias UsageRecord
	return json.Unmarshal(b, (*alias)(r))
}

type ListUsageTypesParams struct {
	P map[string]interface{}
}

func (P *ListUsageTypesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	return u
}

// You should always use this function to get a new ListUsageTypesParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewListUsageTypesParams() *ListUsageTypesParams {
	P := &ListUsageTypesParams{}
	P.P = make(map[string]interface{})
	return P
}

// List Usage Types
func (s *UsageService) ListUsageTypes(p *ListUsageTypesParams) (*ListUsageTypesResponse, error) {
	resp, err := s.cs.newRequest("listUsageTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsageTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUsageTypesResponse struct {
	Count      int          `json:"count"`
	UsageTypes []*UsageType `json:"usagetype"`
}

type UsageType struct {
	Description string `json:"description"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Usagetypeid int    `json:"usagetypeid"`
}

type RemoveRawUsageRecordsParams struct {
	P map[string]interface{}
}

func (P *RemoveRawUsageRecordsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["interval"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("interval", vv)
	}
	return u
}

func (P *RemoveRawUsageRecordsParams) SetInterval(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["interval"] = v
}

func (P *RemoveRawUsageRecordsParams) GetInterval() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["interval"].(int)
	return value, ok
}

// You should always use this function to get a new RemoveRawUsageRecordsParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewRemoveRawUsageRecordsParams(interval int) *RemoveRawUsageRecordsParams {
	P := &RemoveRawUsageRecordsParams{}
	P.P = make(map[string]interface{})
	P.P["interval"] = interval
	return P
}

// Safely removes raw records from cloud_usage table
func (s *UsageService) RemoveRawUsageRecords(p *RemoveRawUsageRecordsParams) (*RemoveRawUsageRecordsResponse, error) {
	resp, err := s.cs.newRequest("removeRawUsageRecords", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveRawUsageRecordsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveRawUsageRecordsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *RemoveRawUsageRecordsResponse) UnmarshalJSON(b []byte) error {
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

	type alias RemoveRawUsageRecordsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateTrafficTypeParams struct {
	P map[string]interface{}
}

func (P *UpdateTrafficTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["hypervnetworklabel"]; found {
		u.Set("hypervnetworklabel", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["kvmnetworklabel"]; found {
		u.Set("kvmnetworklabel", v.(string))
	}
	if v, found := P.P["ovm3networklabel"]; found {
		u.Set("ovm3networklabel", v.(string))
	}
	if v, found := P.P["vmwarenetworklabel"]; found {
		u.Set("vmwarenetworklabel", v.(string))
	}
	if v, found := P.P["xennetworklabel"]; found {
		u.Set("xennetworklabel", v.(string))
	}
	return u
}

func (P *UpdateTrafficTypeParams) SetHypervnetworklabel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervnetworklabel"] = v
}

func (P *UpdateTrafficTypeParams) GetHypervnetworklabel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervnetworklabel"].(string)
	return value, ok
}

func (P *UpdateTrafficTypeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateTrafficTypeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateTrafficTypeParams) SetKvmnetworklabel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["kvmnetworklabel"] = v
}

func (P *UpdateTrafficTypeParams) GetKvmnetworklabel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["kvmnetworklabel"].(string)
	return value, ok
}

func (P *UpdateTrafficTypeParams) SetOvm3networklabel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ovm3networklabel"] = v
}

func (P *UpdateTrafficTypeParams) GetOvm3networklabel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ovm3networklabel"].(string)
	return value, ok
}

func (P *UpdateTrafficTypeParams) SetVmwarenetworklabel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vmwarenetworklabel"] = v
}

func (P *UpdateTrafficTypeParams) GetVmwarenetworklabel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vmwarenetworklabel"].(string)
	return value, ok
}

func (P *UpdateTrafficTypeParams) SetXennetworklabel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["xennetworklabel"] = v
}

func (P *UpdateTrafficTypeParams) GetXennetworklabel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["xennetworklabel"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateTrafficTypeParams instance,
// as then you are sure you have configured all required params
func (s *UsageService) NewUpdateTrafficTypeParams(id string) *UpdateTrafficTypeParams {
	P := &UpdateTrafficTypeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates traffic type of a physical network
func (s *UsageService) UpdateTrafficType(p *UpdateTrafficTypeParams) (*UpdateTrafficTypeResponse, error) {
	resp, err := s.cs.newRequest("updateTrafficType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateTrafficTypeResponse
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

type UpdateTrafficTypeResponse struct {
	Hypervnetworklabel string `json:"hypervnetworklabel"`
	Id                 string `json:"id"`
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Kvmnetworklabel    string `json:"kvmnetworklabel"`
	Ovm3networklabel   string `json:"ovm3networklabel"`
	Physicalnetworkid  string `json:"physicalnetworkid"`
	Traffictype        string `json:"traffictype"`
	Vmwarenetworklabel string `json:"vmwarenetworklabel"`
	Xennetworklabel    string `json:"xennetworklabel"`
}
