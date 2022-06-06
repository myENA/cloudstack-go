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

type BaremetalServiceIface interface {
	AddBaremetalDhcp(P *AddBaremetalDhcpParams) (*AddBaremetalDhcpResponse, error)
	NewAddBaremetalDhcpParams(dhcpservertype string, password string, physicalnetworkid string, url string, username string) *AddBaremetalDhcpParams
	AddBaremetalPxeKickStartServer(P *AddBaremetalPxeKickStartServerParams) (*AddBaremetalPxeKickStartServerResponse, error)
	NewAddBaremetalPxeKickStartServerParams(password string, physicalnetworkid string, pxeservertype string, tftpdir string, url string, username string) *AddBaremetalPxeKickStartServerParams
	AddBaremetalPxePingServer(P *AddBaremetalPxePingServerParams) (*AddBaremetalPxePingServerResponse, error)
	NewAddBaremetalPxePingServerParams(password string, physicalnetworkid string, pingdir string, pingstorageserverip string, pxeservertype string, tftpdir string, url string, username string) *AddBaremetalPxePingServerParams
	AddBaremetalRct(P *AddBaremetalRctParams) (*AddBaremetalRctResponse, error)
	NewAddBaremetalRctParams(baremetalrcturl string) *AddBaremetalRctParams
	DeleteBaremetalRct(P *DeleteBaremetalRctParams) (*DeleteBaremetalRctResponse, error)
	NewDeleteBaremetalRctParams(id string) *DeleteBaremetalRctParams
	ListBaremetalDhcp(P *ListBaremetalDhcpParams) (*ListBaremetalDhcpResponse, error)
	NewListBaremetalDhcpParams(physicalnetworkid string) *ListBaremetalDhcpParams
	ListBaremetalPxeServers(P *ListBaremetalPxeServersParams) (*ListBaremetalPxeServersResponse, error)
	NewListBaremetalPxeServersParams(physicalnetworkid string) *ListBaremetalPxeServersParams
	ListBaremetalRct(P *ListBaremetalRctParams) (*ListBaremetalRctResponse, error)
	NewListBaremetalRctParams() *ListBaremetalRctParams
	NotifyBaremetalProvisionDone(P *NotifyBaremetalProvisionDoneParams) (*NotifyBaremetalProvisionDoneResponse, error)
	NewNotifyBaremetalProvisionDoneParams(mac string) *NotifyBaremetalProvisionDoneParams
}

type AddBaremetalDhcpParams struct {
	P map[string]interface{}
}

func (P *AddBaremetalDhcpParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["dhcpservertype"]; found {
		u.Set("dhcpservertype", v.(string))
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *AddBaremetalDhcpParams) SetDhcpservertype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dhcpservertype"] = v
}

func (P *AddBaremetalDhcpParams) GetDhcpservertype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dhcpservertype"].(string)
	return value, ok
}

func (P *AddBaremetalDhcpParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddBaremetalDhcpParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddBaremetalDhcpParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddBaremetalDhcpParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddBaremetalDhcpParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddBaremetalDhcpParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddBaremetalDhcpParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddBaremetalDhcpParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddBaremetalDhcpParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewAddBaremetalDhcpParams(dhcpservertype string, password string, physicalnetworkid string, url string, username string) *AddBaremetalDhcpParams {
	P := &AddBaremetalDhcpParams{}
	P.P = make(map[string]interface{})
	P.P["dhcpservertype"] = dhcpservertype
	P.P["password"] = password
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["url"] = url
	P.P["username"] = username
	return P
}

// adds a baremetal dhcp server
func (s *BaremetalService) AddBaremetalDhcp(p *AddBaremetalDhcpParams) (*AddBaremetalDhcpResponse, error) {
	resp, err := s.cs.newRequest("addBaremetalDhcp", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalDhcpResponse
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

type AddBaremetalDhcpResponse struct {
	Dhcpservertype    string `json:"dhcpservertype"`
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Url               string `json:"url"`
}

type AddBaremetalPxeKickStartServerParams struct {
	P map[string]interface{}
}

func (P *AddBaremetalPxeKickStartServerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["pxeservertype"]; found {
		u.Set("pxeservertype", v.(string))
	}
	if v, found := P.P["tftpdir"]; found {
		u.Set("tftpdir", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *AddBaremetalPxeKickStartServerParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddBaremetalPxeKickStartServerParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddBaremetalPxeKickStartServerParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddBaremetalPxeKickStartServerParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddBaremetalPxeKickStartServerParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *AddBaremetalPxeKickStartServerParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *AddBaremetalPxeKickStartServerParams) SetPxeservertype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pxeservertype"] = v
}

func (P *AddBaremetalPxeKickStartServerParams) GetPxeservertype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pxeservertype"].(string)
	return value, ok
}

func (P *AddBaremetalPxeKickStartServerParams) SetTftpdir(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tftpdir"] = v
}

func (P *AddBaremetalPxeKickStartServerParams) GetTftpdir() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tftpdir"].(string)
	return value, ok
}

func (P *AddBaremetalPxeKickStartServerParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddBaremetalPxeKickStartServerParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddBaremetalPxeKickStartServerParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddBaremetalPxeKickStartServerParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddBaremetalPxeKickStartServerParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewAddBaremetalPxeKickStartServerParams(password string, physicalnetworkid string, pxeservertype string, tftpdir string, url string, username string) *AddBaremetalPxeKickStartServerParams {
	P := &AddBaremetalPxeKickStartServerParams{}
	P.P = make(map[string]interface{})
	P.P["password"] = password
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["pxeservertype"] = pxeservertype
	P.P["tftpdir"] = tftpdir
	P.P["url"] = url
	P.P["username"] = username
	return P
}

// add a baremetal pxe server
func (s *BaremetalService) AddBaremetalPxeKickStartServer(p *AddBaremetalPxeKickStartServerParams) (*AddBaremetalPxeKickStartServerResponse, error) {
	resp, err := s.cs.newRequest("addBaremetalPxeKickStartServer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalPxeKickStartServerResponse
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

type AddBaremetalPxeKickStartServerResponse struct {
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Tftpdir           string `json:"tftpdir"`
	Url               string `json:"url"`
}

type AddBaremetalPxePingServerParams struct {
	P map[string]interface{}
}

func (P *AddBaremetalPxePingServerParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := P.P["pingcifspassword"]; found {
		u.Set("pingcifspassword", v.(string))
	}
	if v, found := P.P["pingcifsusername"]; found {
		u.Set("pingcifsusername", v.(string))
	}
	if v, found := P.P["pingdir"]; found {
		u.Set("pingdir", v.(string))
	}
	if v, found := P.P["pingstorageserverip"]; found {
		u.Set("pingstorageserverip", v.(string))
	}
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["pxeservertype"]; found {
		u.Set("pxeservertype", v.(string))
	}
	if v, found := P.P["tftpdir"]; found {
		u.Set("tftpdir", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *AddBaremetalPxePingServerParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddBaremetalPxePingServerParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddBaremetalPxePingServerParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *AddBaremetalPxePingServerParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

func (P *AddBaremetalPxePingServerParams) SetPingcifspassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pingcifspassword"] = v
}

func (P *AddBaremetalPxePingServerParams) GetPingcifspassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pingcifspassword"].(string)
	return value, ok
}

func (P *AddBaremetalPxePingServerParams) SetPingcifsusername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pingcifsusername"] = v
}

func (P *AddBaremetalPxePingServerParams) GetPingcifsusername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pingcifsusername"].(string)
	return value, ok
}

func (P *AddBaremetalPxePingServerParams) SetPingdir(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pingdir"] = v
}

func (P *AddBaremetalPxePingServerParams) GetPingdir() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pingdir"].(string)
	return value, ok
}

func (P *AddBaremetalPxePingServerParams) SetPingstorageserverip(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pingstorageserverip"] = v
}

func (P *AddBaremetalPxePingServerParams) GetPingstorageserverip() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pingstorageserverip"].(string)
	return value, ok
}

func (P *AddBaremetalPxePingServerParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *AddBaremetalPxePingServerParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *AddBaremetalPxePingServerParams) SetPxeservertype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pxeservertype"] = v
}

func (P *AddBaremetalPxePingServerParams) GetPxeservertype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pxeservertype"].(string)
	return value, ok
}

func (P *AddBaremetalPxePingServerParams) SetTftpdir(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tftpdir"] = v
}

func (P *AddBaremetalPxePingServerParams) GetTftpdir() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tftpdir"].(string)
	return value, ok
}

func (P *AddBaremetalPxePingServerParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *AddBaremetalPxePingServerParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *AddBaremetalPxePingServerParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddBaremetalPxePingServerParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddBaremetalPxePingServerParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewAddBaremetalPxePingServerParams(password string, physicalnetworkid string, pingdir string, pingstorageserverip string, pxeservertype string, tftpdir string, url string, username string) *AddBaremetalPxePingServerParams {
	P := &AddBaremetalPxePingServerParams{}
	P.P = make(map[string]interface{})
	P.P["password"] = password
	P.P["physicalnetworkid"] = physicalnetworkid
	P.P["pingdir"] = pingdir
	P.P["pingstorageserverip"] = pingstorageserverip
	P.P["pxeservertype"] = pxeservertype
	P.P["tftpdir"] = tftpdir
	P.P["url"] = url
	P.P["username"] = username
	return P
}

// add a baremetal ping pxe server
func (s *BaremetalService) AddBaremetalPxePingServer(p *AddBaremetalPxePingServerParams) (*AddBaremetalPxePingServerResponse, error) {
	resp, err := s.cs.newRequest("addBaremetalPxePingServer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalPxePingServerResponse
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

type AddBaremetalPxePingServerResponse struct {
	Id                  string `json:"id"`
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	Physicalnetworkid   string `json:"physicalnetworkid"`
	Pingdir             string `json:"pingdir"`
	Pingstorageserverip string `json:"pingstorageserverip"`
	Provider            string `json:"provider"`
	Tftpdir             string `json:"tftpdir"`
	Url                 string `json:"url"`
}

type AddBaremetalRctParams struct {
	P map[string]interface{}
}

func (P *AddBaremetalRctParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["baremetalrcturl"]; found {
		u.Set("baremetalrcturl", v.(string))
	}
	return u
}

func (P *AddBaremetalRctParams) SetBaremetalrcturl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["baremetalrcturl"] = v
}

func (P *AddBaremetalRctParams) GetBaremetalrcturl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["baremetalrcturl"].(string)
	return value, ok
}

// You should always use this function to get a new AddBaremetalRctParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewAddBaremetalRctParams(baremetalrcturl string) *AddBaremetalRctParams {
	P := &AddBaremetalRctParams{}
	P.P = make(map[string]interface{})
	P.P["baremetalrcturl"] = baremetalrcturl
	return P
}

// adds baremetal rack configuration text
func (s *BaremetalService) AddBaremetalRct(p *AddBaremetalRctParams) (*AddBaremetalRctResponse, error) {
	resp, err := s.cs.newRequest("addBaremetalRct", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddBaremetalRctResponse
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

type AddBaremetalRctResponse struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Url       string `json:"url"`
}

type DeleteBaremetalRctParams struct {
	P map[string]interface{}
}

func (P *DeleteBaremetalRctParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteBaremetalRctParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteBaremetalRctParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteBaremetalRctParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewDeleteBaremetalRctParams(id string) *DeleteBaremetalRctParams {
	P := &DeleteBaremetalRctParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// deletes baremetal rack configuration text
func (s *BaremetalService) DeleteBaremetalRct(p *DeleteBaremetalRctParams) (*DeleteBaremetalRctResponse, error) {
	resp, err := s.cs.newRequest("deleteBaremetalRct", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteBaremetalRctResponse
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

type DeleteBaremetalRctResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListBaremetalDhcpParams struct {
	P map[string]interface{}
}

func (P *ListBaremetalDhcpParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["dhcpservertype"]; found {
		u.Set("dhcpservertype", v.(string))
	}
	if v, found := P.P["id"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("id", vv)
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

func (P *ListBaremetalDhcpParams) SetDhcpservertype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dhcpservertype"] = v
}

func (P *ListBaremetalDhcpParams) GetDhcpservertype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dhcpservertype"].(string)
	return value, ok
}

func (P *ListBaremetalDhcpParams) SetId(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListBaremetalDhcpParams) GetId() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(int64)
	return value, ok
}

func (P *ListBaremetalDhcpParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListBaremetalDhcpParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListBaremetalDhcpParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListBaremetalDhcpParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListBaremetalDhcpParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListBaremetalDhcpParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListBaremetalDhcpParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListBaremetalDhcpParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBaremetalDhcpParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewListBaremetalDhcpParams(physicalnetworkid string) *ListBaremetalDhcpParams {
	P := &ListBaremetalDhcpParams{}
	P.P = make(map[string]interface{})
	P.P["physicalnetworkid"] = physicalnetworkid
	return P
}

// list baremetal dhcp servers
func (s *BaremetalService) ListBaremetalDhcp(p *ListBaremetalDhcpParams) (*ListBaremetalDhcpResponse, error) {
	resp, err := s.cs.newRequest("listBaremetalDhcp", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBaremetalDhcpResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBaremetalDhcpResponse struct {
	Count         int              `json:"count"`
	BaremetalDhcp []*BaremetalDhcp `json:"baremetaldhcp"`
}

type BaremetalDhcp struct {
	Dhcpservertype    string `json:"dhcpservertype"`
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Url               string `json:"url"`
}

type ListBaremetalPxeServersParams struct {
	P map[string]interface{}
}

func (P *ListBaremetalPxeServersParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("id", vv)
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

func (P *ListBaremetalPxeServersParams) SetId(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListBaremetalPxeServersParams) GetId() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(int64)
	return value, ok
}

func (P *ListBaremetalPxeServersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListBaremetalPxeServersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListBaremetalPxeServersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListBaremetalPxeServersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListBaremetalPxeServersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListBaremetalPxeServersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListBaremetalPxeServersParams) SetPhysicalnetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["physicalnetworkid"] = v
}

func (P *ListBaremetalPxeServersParams) GetPhysicalnetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["physicalnetworkid"].(string)
	return value, ok
}

// You should always use this function to get a new ListBaremetalPxeServersParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewListBaremetalPxeServersParams(physicalnetworkid string) *ListBaremetalPxeServersParams {
	P := &ListBaremetalPxeServersParams{}
	P.P = make(map[string]interface{})
	P.P["physicalnetworkid"] = physicalnetworkid
	return P
}

// list baremetal pxe server
func (s *BaremetalService) ListBaremetalPxeServers(p *ListBaremetalPxeServersParams) (*ListBaremetalPxeServersResponse, error) {
	resp, err := s.cs.newRequest("listBaremetalPxeServers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBaremetalPxeServersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBaremetalPxeServersResponse struct {
	Count               int                   `json:"count"`
	BaremetalPxeServers []*BaremetalPxeServer `json:"baremetalpxeserver"`
}

type BaremetalPxeServer struct {
	Id                string `json:"id"`
	JobID             string `json:"jobid"`
	Jobstatus         int    `json:"jobstatus"`
	Physicalnetworkid string `json:"physicalnetworkid"`
	Provider          string `json:"provider"`
	Url               string `json:"url"`
}

type ListBaremetalRctParams struct {
	P map[string]interface{}
}

func (P *ListBaremetalRctParams) toURLValues() url.Values {
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
	return u
}

func (P *ListBaremetalRctParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListBaremetalRctParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListBaremetalRctParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListBaremetalRctParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListBaremetalRctParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListBaremetalRctParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListBaremetalRctParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewListBaremetalRctParams() *ListBaremetalRctParams {
	P := &ListBaremetalRctParams{}
	P.P = make(map[string]interface{})
	return P
}

// list baremetal rack configuration
func (s *BaremetalService) ListBaremetalRct(p *ListBaremetalRctParams) (*ListBaremetalRctResponse, error) {
	resp, err := s.cs.newRequest("listBaremetalRct", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListBaremetalRctResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListBaremetalRctResponse struct {
	Count        int             `json:"count"`
	BaremetalRct []*BaremetalRct `json:"baremetalrct"`
}

type BaremetalRct struct {
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Url       string `json:"url"`
}

type NotifyBaremetalProvisionDoneParams struct {
	P map[string]interface{}
}

func (P *NotifyBaremetalProvisionDoneParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["mac"]; found {
		u.Set("mac", v.(string))
	}
	return u
}

func (P *NotifyBaremetalProvisionDoneParams) SetMac(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["mac"] = v
}

func (P *NotifyBaremetalProvisionDoneParams) GetMac() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["mac"].(string)
	return value, ok
}

// You should always use this function to get a new NotifyBaremetalProvisionDoneParams instance,
// as then you are sure you have configured all required params
func (s *BaremetalService) NewNotifyBaremetalProvisionDoneParams(mac string) *NotifyBaremetalProvisionDoneParams {
	P := &NotifyBaremetalProvisionDoneParams{}
	P.P = make(map[string]interface{})
	P.P["mac"] = mac
	return P
}

// Notify provision has been done on a host. This api is for baremetal virtual router service, not for end user
func (s *BaremetalService) NotifyBaremetalProvisionDone(p *NotifyBaremetalProvisionDoneParams) (*NotifyBaremetalProvisionDoneResponse, error) {
	resp, err := s.cs.newRequest("notifyBaremetalProvisionDone", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r NotifyBaremetalProvisionDoneResponse
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

type NotifyBaremetalProvisionDoneResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}
