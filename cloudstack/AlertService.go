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

type AlertServiceIface interface {
	ArchiveAlerts(p *ArchiveAlertsParams) (*ArchiveAlertsResponse, error)
	NewArchiveAlertsParams() *ArchiveAlertsParams
	DeleteAlerts(p *DeleteAlertsParams) (*DeleteAlertsResponse, error)
	NewDeleteAlertsParams() *DeleteAlertsParams
	GenerateAlert(p *GenerateAlertParams) (*GenerateAlertResponse, error)
	NewGenerateAlertParams(description string, name string, alertType int) *GenerateAlertParams
	ListAlerts(p *ListAlertsParams) (*ListAlertsResponse, error)
	NewListAlertsParams() *ListAlertsParams
	GetAlertID(name string, opts ...OptionFunc) (string, int, error)
	GetAlertByName(name string, opts ...OptionFunc) (*Alert, int, error)
	GetAlertByID(id string, opts ...OptionFunc) (*Alert, int, error)
}

type ArchiveAlertsParams struct {
	P map[string]interface{}
}

func (P *ArchiveAlertsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := P.P["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := P.P["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (P *ArchiveAlertsParams) SetEnddate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enddate"] = v
}

func (P *ArchiveAlertsParams) GetEnddate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enddate"].(string)
	return value, ok
}

func (P *ArchiveAlertsParams) SetIds(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ids"] = v
}

func (P *ArchiveAlertsParams) GetIds() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ids"].([]string)
	return value, ok
}

func (P *ArchiveAlertsParams) SetStartdate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startdate"] = v
}

func (P *ArchiveAlertsParams) GetStartdate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startdate"].(string)
	return value, ok
}

func (P *ArchiveAlertsParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ArchiveAlertsParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new ArchiveAlertsParams instance,
// as then you are sure you have configured all required params
func (s *AlertService) NewArchiveAlertsParams() *ArchiveAlertsParams {
	P := &ArchiveAlertsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Archive one or more alerts.
func (s *AlertService) ArchiveAlerts(p *ArchiveAlertsParams) (*ArchiveAlertsResponse, error) {
	resp, err := s.cs.newRequest("archiveAlerts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ArchiveAlertsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ArchiveAlertsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ArchiveAlertsResponse) UnmarshalJSON(b []byte) error {
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

	type alias ArchiveAlertsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteAlertsParams struct {
	P map[string]interface{}
}

func (P *DeleteAlertsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := P.P["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
	}
	if v, found := P.P["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (P *DeleteAlertsParams) SetEnddate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enddate"] = v
}

func (P *DeleteAlertsParams) GetEnddate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enddate"].(string)
	return value, ok
}

func (P *DeleteAlertsParams) SetIds(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ids"] = v
}

func (P *DeleteAlertsParams) GetIds() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ids"].([]string)
	return value, ok
}

func (P *DeleteAlertsParams) SetStartdate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startdate"] = v
}

func (P *DeleteAlertsParams) GetStartdate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startdate"].(string)
	return value, ok
}

func (P *DeleteAlertsParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *DeleteAlertsParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAlertsParams instance,
// as then you are sure you have configured all required params
func (s *AlertService) NewDeleteAlertsParams() *DeleteAlertsParams {
	P := &DeleteAlertsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Delete one or more alerts.
func (s *AlertService) DeleteAlerts(p *DeleteAlertsParams) (*DeleteAlertsResponse, error) {
	resp, err := s.cs.newRequest("deleteAlerts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAlertsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteAlertsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteAlertsResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteAlertsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type GenerateAlertParams struct {
	P map[string]interface{}
}

func (P *GenerateAlertParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["type"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("type", vv)
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *GenerateAlertParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *GenerateAlertParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *GenerateAlertParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *GenerateAlertParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *GenerateAlertParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *GenerateAlertParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *GenerateAlertParams) SetType(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *GenerateAlertParams) GetType() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(int)
	return value, ok
}

func (P *GenerateAlertParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *GenerateAlertParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new GenerateAlertParams instance,
// as then you are sure you have configured all required params
func (s *AlertService) NewGenerateAlertParams(description string, name string, alertType int) *GenerateAlertParams {
	P := &GenerateAlertParams{}
	P.P = make(map[string]interface{})
	P.P["description"] = description
	P.P["name"] = name
	P.P["type"] = alertType
	return P
}

// Generates an alert
func (s *AlertService) GenerateAlert(p *GenerateAlertParams) (*GenerateAlertResponse, error) {
	resp, err := s.cs.newRequest("generateAlert", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GenerateAlertResponse
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

type GenerateAlertResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListAlertsParams struct {
	P map[string]interface{}
}

func (P *ListAlertsParams) toURLValues() url.Values {
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
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (P *ListAlertsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListAlertsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListAlertsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListAlertsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListAlertsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListAlertsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListAlertsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListAlertsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListAlertsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListAlertsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListAlertsParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ListAlertsParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new ListAlertsParams instance,
// as then you are sure you have configured all required params
func (s *AlertService) NewListAlertsParams() *ListAlertsParams {
	P := &ListAlertsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AlertService) GetAlertID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListAlertsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListAlerts(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Alerts[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Alerts {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AlertService) GetAlertByName(name string, opts ...OptionFunc) (*Alert, int, error) {
	id, count, err := s.GetAlertID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetAlertByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AlertService) GetAlertByID(id string, opts ...OptionFunc) (*Alert, int, error) {
	P := &ListAlertsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAlerts(P)
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
		return l.Alerts[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Alert UUID: %s!", id)
}

// Lists all alerts.
func (s *AlertService) ListAlerts(p *ListAlertsParams) (*ListAlertsResponse, error) {
	resp, err := s.cs.newRequest("listAlerts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAlertsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAlertsResponse struct {
	Count  int      `json:"count"`
	Alerts []*Alert `json:"alert"`
}

type Alert struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Sent        string `json:"sent"`
	Type        int    `json:"type"`
}
