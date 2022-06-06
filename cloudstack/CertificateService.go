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
)

type CertificateServiceIface interface {
	UploadCustomCertificate(P *UploadCustomCertificateParams) (*UploadCustomCertificateResponse, error)
	NewUploadCustomCertificateParams(certificate string, domainsuffix string) *UploadCustomCertificateParams
}

type UploadCustomCertificateParams struct {
	P map[string]interface{}
}

func (P *UploadCustomCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["certificate"]; found {
		u.Set("certificate", v.(string))
	}
	if v, found := P.P["domainsuffix"]; found {
		u.Set("domainsuffix", v.(string))
	}
	if v, found := P.P["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["privatekey"]; found {
		u.Set("privatekey", v.(string))
	}
	return u
}

func (P *UploadCustomCertificateParams) SetCertificate(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["certificate"] = v
}

func (P *UploadCustomCertificateParams) GetCertificate() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["certificate"].(string)
	return value, ok
}

func (P *UploadCustomCertificateParams) SetDomainsuffix(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainsuffix"] = v
}

func (P *UploadCustomCertificateParams) GetDomainsuffix() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainsuffix"].(string)
	return value, ok
}

func (P *UploadCustomCertificateParams) SetId(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UploadCustomCertificateParams) GetId() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(int)
	return value, ok
}

func (P *UploadCustomCertificateParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UploadCustomCertificateParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UploadCustomCertificateParams) SetPrivatekey(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["privatekey"] = v
}

func (P *UploadCustomCertificateParams) GetPrivatekey() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["privatekey"].(string)
	return value, ok
}

// You should always use this function to get a new UploadCustomCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewUploadCustomCertificateParams(certificate string, domainsuffix string) *UploadCustomCertificateParams {
	P := &UploadCustomCertificateParams{}
	P.P = make(map[string]interface{})
	P.P["certificate"] = certificate
	P.P["domainsuffix"] = domainsuffix
	return P
}

// Uploads a custom certificate for the console proxy VMs to use for SSL. Can be used to upload a single certificate signed by a known CA. Can also be used, through multiple calls, to upload a chain of certificates from CA to the custom certificate itself.
func (s *CertificateService) UploadCustomCertificate(p *UploadCustomCertificateParams) (*UploadCustomCertificateResponse, error) {
	resp, err := s.cs.newRequest("uploadCustomCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadCustomCertificateResponse
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

type UploadCustomCertificateResponse struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Message   string `json:"message"`
}
