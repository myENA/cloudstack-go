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

type EventServiceIface interface {
	ArchiveEvents(P *ArchiveEventsParams) (*ArchiveEventsResponse, error)
	NewArchiveEventsParams() *ArchiveEventsParams
	DeleteEvents(P *DeleteEventsParams) (*DeleteEventsResponse, error)
	NewDeleteEventsParams() *DeleteEventsParams
	ListEventTypes(P *ListEventTypesParams) (*ListEventTypesResponse, error)
	NewListEventTypesParams() *ListEventTypesParams
	ListEvents(P *ListEventsParams) (*ListEventsResponse, error)
	NewListEventsParams() *ListEventsParams
	GetEventByID(id string, opts ...OptionFunc) (*Event, int, error)
}

type ArchiveEventsParams struct {
	P map[string]interface{}
}

func (P *ArchiveEventsParams) toURLValues() url.Values {
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

func (P *ArchiveEventsParams) SetEnddate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enddate"] = v
}

func (P *ArchiveEventsParams) GetEnddate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enddate"].(string)
	return value, ok
}

func (P *ArchiveEventsParams) SetIds(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ids"] = v
}

func (P *ArchiveEventsParams) GetIds() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ids"].([]string)
	return value, ok
}

func (P *ArchiveEventsParams) SetStartdate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startdate"] = v
}

func (P *ArchiveEventsParams) GetStartdate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startdate"].(string)
	return value, ok
}

func (P *ArchiveEventsParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ArchiveEventsParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new ArchiveEventsParams instance,
// as then you are sure you have configured all required params
func (s *EventService) NewArchiveEventsParams() *ArchiveEventsParams {
	P := &ArchiveEventsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Archive one or more events.
func (s *EventService) ArchiveEvents(p *ArchiveEventsParams) (*ArchiveEventsResponse, error) {
	resp, err := s.cs.newRequest("archiveEvents", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ArchiveEventsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ArchiveEventsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ArchiveEventsResponse) UnmarshalJSON(b []byte) error {
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

	type alias ArchiveEventsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteEventsParams struct {
	P map[string]interface{}
}

func (P *DeleteEventsParams) toURLValues() url.Values {
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

func (P *DeleteEventsParams) SetEnddate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enddate"] = v
}

func (P *DeleteEventsParams) GetEnddate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enddate"].(string)
	return value, ok
}

func (P *DeleteEventsParams) SetIds(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ids"] = v
}

func (P *DeleteEventsParams) GetIds() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ids"].([]string)
	return value, ok
}

func (P *DeleteEventsParams) SetStartdate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startdate"] = v
}

func (P *DeleteEventsParams) GetStartdate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startdate"].(string)
	return value, ok
}

func (P *DeleteEventsParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *DeleteEventsParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteEventsParams instance,
// as then you are sure you have configured all required params
func (s *EventService) NewDeleteEventsParams() *DeleteEventsParams {
	P := &DeleteEventsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Delete one or more events.
func (s *EventService) DeleteEvents(p *DeleteEventsParams) (*DeleteEventsResponse, error) {
	resp, err := s.cs.newRequest("deleteEvents", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteEventsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteEventsResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteEventsResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteEventsResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListEventTypesParams struct {
	P map[string]interface{}
}

func (P *ListEventTypesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	return u
}

// You should always use this function to get a new ListEventTypesParams instance,
// as then you are sure you have configured all required params
func (s *EventService) NewListEventTypesParams() *ListEventTypesParams {
	P := &ListEventTypesParams{}
	P.P = make(map[string]interface{})
	return P
}

// List Event Types
func (s *EventService) ListEventTypes(p *ListEventTypesParams) (*ListEventTypesResponse, error) {
	resp, err := s.cs.newRequest("listEventTypes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListEventTypesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListEventTypesResponse struct {
	Count      int          `json:"count"`
	EventTypes []*EventType `json:"eventtype"`
}

type EventType struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
}

type ListEventsParams struct {
	P map[string]interface{}
}

func (P *ListEventsParams) toURLValues() url.Values {
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
	if v, found := P.P["duration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("duration", vv)
	}
	if v, found := P.P["enddate"]; found {
		u.Set("enddate", v.(string))
	}
	if v, found := P.P["entrytime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("entrytime", vv)
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
	if v, found := P.P["level"]; found {
		u.Set("level", v.(string))
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	if v, found := P.P["startid"]; found {
		u.Set("startid", v.(string))
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (P *ListEventsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListEventsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListEventsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListEventsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListEventsParams) SetDuration(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["duration"] = v
}

func (P *ListEventsParams) GetDuration() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["duration"].(int)
	return value, ok
}

func (P *ListEventsParams) SetEnddate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enddate"] = v
}

func (P *ListEventsParams) GetEnddate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enddate"].(string)
	return value, ok
}

func (P *ListEventsParams) SetEntrytime(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["entrytime"] = v
}

func (P *ListEventsParams) GetEntrytime() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["entrytime"].(int)
	return value, ok
}

func (P *ListEventsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListEventsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListEventsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListEventsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListEventsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListEventsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListEventsParams) SetLevel(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["level"] = v
}

func (P *ListEventsParams) GetLevel() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["level"].(string)
	return value, ok
}

func (P *ListEventsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListEventsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListEventsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListEventsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListEventsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListEventsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListEventsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListEventsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListEventsParams) SetStartdate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startdate"] = v
}

func (P *ListEventsParams) GetStartdate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startdate"].(string)
	return value, ok
}

func (P *ListEventsParams) SetStartid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startid"] = v
}

func (P *ListEventsParams) GetStartid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startid"].(string)
	return value, ok
}

func (P *ListEventsParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ListEventsParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new ListEventsParams instance,
// as then you are sure you have configured all required params
func (s *EventService) NewListEventsParams() *ListEventsParams {
	P := &ListEventsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *EventService) GetEventByID(id string, opts ...OptionFunc) (*Event, int, error) {
	P := &ListEventsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListEvents(P)
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
		return l.Events[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Event UUID: %s!", id)
}

// A command to list events.
func (s *EventService) ListEvents(p *ListEventsParams) (*ListEventsResponse, error) {
	resp, err := s.cs.newRequest("listEvents", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListEventsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListEventsResponse struct {
	Count  int      `json:"count"`
	Events []*Event `json:"event"`
}

type Event struct {
	Account     string `json:"account"`
	Created     string `json:"created"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	Domainid    string `json:"domainid"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Level       string `json:"level"`
	Parentid    string `json:"parentid"`
	Project     string `json:"project"`
	Projectid   string `json:"projectid"`
	State       string `json:"state"`
	Type        string `json:"type"`
	Username    string `json:"username"`
}
