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
)

type CloudIdentifierServiceIface interface {
	GetCloudIdentifier(p *GetCloudIdentifierParams) (*GetCloudIdentifierResponse, error)
	NewGetCloudIdentifierParams(userid string) *GetCloudIdentifierParams
}

type GetCloudIdentifierParams struct {
	P map[string]interface{}
}

func (P *GetCloudIdentifierParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (P *GetCloudIdentifierParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *GetCloudIdentifierParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

// You should always use this function to get a new GetCloudIdentifierParams instance,
// as then you are sure you have configured all required params
func (s *CloudIdentifierService) NewGetCloudIdentifierParams(userid string) *GetCloudIdentifierParams {
	P := &GetCloudIdentifierParams{}
	P.P = make(map[string]interface{})
	P.P["userid"] = userid
	return P
}

// Retrieves a cloud identifier.
func (s *CloudIdentifierService) GetCloudIdentifier(p *GetCloudIdentifierParams) (*GetCloudIdentifierResponse, error) {
	resp, err := s.cs.newRequest("getCloudIdentifier", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetCloudIdentifierResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetCloudIdentifierResponse struct {
	Cloudidentifier string `json:"cloudidentifier"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Signature       string `json:"signature"`
	Userid          string `json:"userid"`
}
