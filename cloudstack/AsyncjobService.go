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
	"net/url"
	"strconv"
	"time"
)

type AsyncjobServiceIface interface {
	ListAsyncJobs(P *ListAsyncJobsParams) (*ListAsyncJobsResponse, error)
	NewListAsyncJobsParams() *ListAsyncJobsParams
	QueryAsyncJobResult(P *QueryAsyncJobResultParams) (*QueryAsyncJobResultResponse, error)
	NewQueryAsyncJobResultParams(jobid string) *QueryAsyncJobResultParams
}

type ListAsyncJobsParams struct {
	P map[string]interface{}
}

func (P *ListAsyncJobsParams) toURLValues() url.Values {
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
	if v, found := P.P["startdate"]; found {
		u.Set("startdate", v.(string))
	}
	return u
}

func (P *ListAsyncJobsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListAsyncJobsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListAsyncJobsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListAsyncJobsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListAsyncJobsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListAsyncJobsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListAsyncJobsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListAsyncJobsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListAsyncJobsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListAsyncJobsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListAsyncJobsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListAsyncJobsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListAsyncJobsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListAsyncJobsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListAsyncJobsParams) SetStartdate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["startdate"] = v
}

func (P *ListAsyncJobsParams) GetStartdate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["startdate"].(string)
	return value, ok
}

// You should always use this function to get a new ListAsyncJobsParams instance,
// as then you are sure you have configured all required params
func (s *AsyncjobService) NewListAsyncJobsParams() *ListAsyncJobsParams {
	P := &ListAsyncJobsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists all pending asynchronous jobs for the account.
func (s *AsyncjobService) ListAsyncJobs(p *ListAsyncJobsParams) (*ListAsyncJobsResponse, error) {
	resp, err := s.cs.newRequest("listAsyncJobs", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAsyncJobsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAsyncJobsResponse struct {
	Count     int         `json:"count"`
	AsyncJobs []*AsyncJob `json:"asyncjobs"`
}

type AsyncJob struct {
	Accountid       string          `json:"accountid"`
	Cmd             string          `json:"cmd"`
	Completed       string          `json:"completed"`
	Created         string          `json:"created"`
	JobID           string          `json:"jobid"`
	Jobinstanceid   string          `json:"jobinstanceid"`
	Jobinstancetype string          `json:"jobinstancetype"`
	Jobprocstatus   int             `json:"jobprocstatus"`
	Jobresult       json.RawMessage `json:"jobresult"`
	Jobresultcode   int             `json:"jobresultcode"`
	Jobresulttype   string          `json:"jobresulttype"`
	Jobstatus       int             `json:"jobstatus"`
	Userid          string          `json:"userid"`
}

type QueryAsyncJobResultParams struct {
	P map[string]interface{}
}

func (P *QueryAsyncJobResultParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["jobid"]; found {
		u.Set("jobid", v.(string))
	}
	return u
}

func (P *QueryAsyncJobResultParams) SetJobID(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["jobid"] = v
}

func (P *QueryAsyncJobResultParams) GetJobID() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["jobid"].(string)
	return value, ok
}

// You should always use this function to get a new QueryAsyncJobResultParams instance,
// as then you are sure you have configured all required params
func (s *AsyncjobService) NewQueryAsyncJobResultParams(jobid string) *QueryAsyncJobResultParams {
	P := &QueryAsyncJobResultParams{}
	P.P = make(map[string]interface{})
	P.P["jobid"] = jobid
	return P
}

// Retrieves the current status of asynchronous job.
func (s *AsyncjobService) QueryAsyncJobResult(p *QueryAsyncJobResultParams) (*QueryAsyncJobResultResponse, error) {
	var resp json.RawMessage
	var err error

	// We should be able to retry on failure as this call is idempotent
	for i := 0; i < 3; i++ {
		resp, err = s.cs.newRequest("queryAsyncJobResult", p.toURLValues())
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		return nil, err
	}

	var r QueryAsyncJobResultResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type QueryAsyncJobResultResponse struct {
	Accountid       string          `json:"accountid"`
	Cmd             string          `json:"cmd"`
	Completed       string          `json:"completed"`
	Created         string          `json:"created"`
	JobID           string          `json:"jobid"`
	Jobinstanceid   string          `json:"jobinstanceid"`
	Jobinstancetype string          `json:"jobinstancetype"`
	Jobprocstatus   int             `json:"jobprocstatus"`
	Jobresult       json.RawMessage `json:"jobresult"`
	Jobresultcode   int             `json:"jobresultcode"`
	Jobresulttype   string          `json:"jobresulttype"`
	Jobstatus       int             `json:"jobstatus"`
	Userid          string          `json:"userid"`
}
