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

type VPNServiceIface interface {
	AddVpnUser(P *AddVpnUserParams) (*AddVpnUserResponse, error)
	NewAddVpnUserParams(password string, username string) *AddVpnUserParams
	CreateRemoteAccessVpn(P *CreateRemoteAccessVpnParams) (*CreateRemoteAccessVpnResponse, error)
	NewCreateRemoteAccessVpnParams(publicipid string) *CreateRemoteAccessVpnParams
	CreateVpnConnection(P *CreateVpnConnectionParams) (*CreateVpnConnectionResponse, error)
	NewCreateVpnConnectionParams(s2scustomergatewayid string, s2svpngatewayid string) *CreateVpnConnectionParams
	CreateVpnCustomerGateway(P *CreateVpnCustomerGatewayParams) (*CreateVpnCustomerGatewayResponse, error)
	NewCreateVpnCustomerGatewayParams(cidrlist string, esppolicy string, gateway string, ikepolicy string, ipsecpsk string) *CreateVpnCustomerGatewayParams
	CreateVpnGateway(P *CreateVpnGatewayParams) (*CreateVpnGatewayResponse, error)
	NewCreateVpnGatewayParams(vpcid string) *CreateVpnGatewayParams
	DeleteRemoteAccessVpn(P *DeleteRemoteAccessVpnParams) (*DeleteRemoteAccessVpnResponse, error)
	NewDeleteRemoteAccessVpnParams(publicipid string) *DeleteRemoteAccessVpnParams
	DeleteVpnConnection(P *DeleteVpnConnectionParams) (*DeleteVpnConnectionResponse, error)
	NewDeleteVpnConnectionParams(id string) *DeleteVpnConnectionParams
	DeleteVpnCustomerGateway(P *DeleteVpnCustomerGatewayParams) (*DeleteVpnCustomerGatewayResponse, error)
	NewDeleteVpnCustomerGatewayParams(id string) *DeleteVpnCustomerGatewayParams
	DeleteVpnGateway(P *DeleteVpnGatewayParams) (*DeleteVpnGatewayResponse, error)
	NewDeleteVpnGatewayParams(id string) *DeleteVpnGatewayParams
	ListRemoteAccessVpns(P *ListRemoteAccessVpnsParams) (*ListRemoteAccessVpnsResponse, error)
	NewListRemoteAccessVpnsParams() *ListRemoteAccessVpnsParams
	GetRemoteAccessVpnByID(id string, opts ...OptionFunc) (*RemoteAccessVpn, int, error)
	ListVpnConnections(P *ListVpnConnectionsParams) (*ListVpnConnectionsResponse, error)
	NewListVpnConnectionsParams() *ListVpnConnectionsParams
	GetVpnConnectionByID(id string, opts ...OptionFunc) (*VpnConnection, int, error)
	ListVpnCustomerGateways(P *ListVpnCustomerGatewaysParams) (*ListVpnCustomerGatewaysResponse, error)
	NewListVpnCustomerGatewaysParams() *ListVpnCustomerGatewaysParams
	GetVpnCustomerGatewayID(keyword string, opts ...OptionFunc) (string, int, error)
	GetVpnCustomerGatewayByName(name string, opts ...OptionFunc) (*VpnCustomerGateway, int, error)
	GetVpnCustomerGatewayByID(id string, opts ...OptionFunc) (*VpnCustomerGateway, int, error)
	ListVpnGateways(P *ListVpnGatewaysParams) (*ListVpnGatewaysResponse, error)
	NewListVpnGatewaysParams() *ListVpnGatewaysParams
	GetVpnGatewayByID(id string, opts ...OptionFunc) (*VpnGateway, int, error)
	ListVpnUsers(P *ListVpnUsersParams) (*ListVpnUsersResponse, error)
	NewListVpnUsersParams() *ListVpnUsersParams
	GetVpnUserByID(id string, opts ...OptionFunc) (*VpnUser, int, error)
	RemoveVpnUser(P *RemoveVpnUserParams) (*RemoveVpnUserResponse, error)
	NewRemoveVpnUserParams(username string) *RemoveVpnUserParams
	ResetVpnConnection(P *ResetVpnConnectionParams) (*ResetVpnConnectionResponse, error)
	NewResetVpnConnectionParams(id string) *ResetVpnConnectionParams
	UpdateRemoteAccessVpn(P *UpdateRemoteAccessVpnParams) (*UpdateRemoteAccessVpnResponse, error)
	NewUpdateRemoteAccessVpnParams(id string) *UpdateRemoteAccessVpnParams
	UpdateVpnConnection(P *UpdateVpnConnectionParams) (*UpdateVpnConnectionResponse, error)
	NewUpdateVpnConnectionParams(id string) *UpdateVpnConnectionParams
	UpdateVpnCustomerGateway(P *UpdateVpnCustomerGatewayParams) (*UpdateVpnCustomerGatewayResponse, error)
	NewUpdateVpnCustomerGatewayParams(cidrlist string, esppolicy string, gateway string, id string, ikepolicy string, ipsecpsk string) *UpdateVpnCustomerGatewayParams
	UpdateVpnGateway(P *UpdateVpnGatewayParams) (*UpdateVpnGatewayResponse, error)
	NewUpdateVpnGatewayParams(id string) *UpdateVpnGatewayParams
}

type AddVpnUserParams struct {
	P map[string]interface{}
}

func (P *AddVpnUserParams) toURLValues() url.Values {
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
	if v, found := P.P["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *AddVpnUserParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *AddVpnUserParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *AddVpnUserParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *AddVpnUserParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *AddVpnUserParams) SetPassword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["password"] = v
}

func (P *AddVpnUserParams) GetPassword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["password"].(string)
	return value, ok
}

func (P *AddVpnUserParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *AddVpnUserParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *AddVpnUserParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddVpnUserParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddVpnUserParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewAddVpnUserParams(password string, username string) *AddVpnUserParams {
	P := &AddVpnUserParams{}
	P.P = make(map[string]interface{})
	P.P["password"] = password
	P.P["username"] = username
	return P
}

// Adds vpn users
func (s *VPNService) AddVpnUser(p *AddVpnUserParams) (*AddVpnUserResponse, error) {
	resp, err := s.cs.newRequest("addVpnUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddVpnUserResponse
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

type AddVpnUserResponse struct {
	Account   string `json:"account"`
	Domain    string `json:"domain"`
	Domainid  string `json:"domainid"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
	State     string `json:"state"`
	Username  string `json:"username"`
}

type CreateRemoteAccessVpnParams struct {
	P map[string]interface{}
}

func (P *CreateRemoteAccessVpnParams) toURLValues() url.Values {
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
	if v, found := P.P["iprange"]; found {
		u.Set("iprange", v.(string))
	}
	if v, found := P.P["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := P.P["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	return u
}

func (P *CreateRemoteAccessVpnParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateRemoteAccessVpnParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateRemoteAccessVpnParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateRemoteAccessVpnParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateRemoteAccessVpnParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateRemoteAccessVpnParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateRemoteAccessVpnParams) SetIprange(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iprange"] = v
}

func (P *CreateRemoteAccessVpnParams) GetIprange() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iprange"].(string)
	return value, ok
}

func (P *CreateRemoteAccessVpnParams) SetOpenfirewall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["openfirewall"] = v
}

func (P *CreateRemoteAccessVpnParams) GetOpenfirewall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["openfirewall"].(bool)
	return value, ok
}

func (P *CreateRemoteAccessVpnParams) SetPublicipid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publicipid"] = v
}

func (P *CreateRemoteAccessVpnParams) GetPublicipid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publicipid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateRemoteAccessVpnParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateRemoteAccessVpnParams(publicipid string) *CreateRemoteAccessVpnParams {
	P := &CreateRemoteAccessVpnParams{}
	P.P = make(map[string]interface{})
	P.P["publicipid"] = publicipid
	return P
}

// Creates a l2tp/ipsec remote access vpn
func (s *VPNService) CreateRemoteAccessVpn(p *CreateRemoteAccessVpnParams) (*CreateRemoteAccessVpnResponse, error) {
	resp, err := s.cs.newRequest("createRemoteAccessVpn", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateRemoteAccessVpnResponse
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

type CreateRemoteAccessVpnResponse struct {
	Account      string `json:"account"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Fordisplay   bool   `json:"fordisplay"`
	Id           string `json:"id"`
	Iprange      string `json:"iprange"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Presharedkey string `json:"presharedkey"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Publicip     string `json:"publicip"`
	Publicipid   string `json:"publicipid"`
	State        string `json:"state"`
}

type CreateVpnConnectionParams struct {
	P map[string]interface{}
}

func (P *CreateVpnConnectionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["passive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passive", vv)
	}
	if v, found := P.P["s2scustomergatewayid"]; found {
		u.Set("s2scustomergatewayid", v.(string))
	}
	if v, found := P.P["s2svpngatewayid"]; found {
		u.Set("s2svpngatewayid", v.(string))
	}
	return u
}

func (P *CreateVpnConnectionParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateVpnConnectionParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateVpnConnectionParams) SetPassive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["passive"] = v
}

func (P *CreateVpnConnectionParams) GetPassive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["passive"].(bool)
	return value, ok
}

func (P *CreateVpnConnectionParams) SetS2scustomergatewayid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["s2scustomergatewayid"] = v
}

func (P *CreateVpnConnectionParams) GetS2scustomergatewayid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["s2scustomergatewayid"].(string)
	return value, ok
}

func (P *CreateVpnConnectionParams) SetS2svpngatewayid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["s2svpngatewayid"] = v
}

func (P *CreateVpnConnectionParams) GetS2svpngatewayid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["s2svpngatewayid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateVpnConnectionParams(s2scustomergatewayid string, s2svpngatewayid string) *CreateVpnConnectionParams {
	P := &CreateVpnConnectionParams{}
	P.P = make(map[string]interface{})
	P.P["s2scustomergatewayid"] = s2scustomergatewayid
	P.P["s2svpngatewayid"] = s2svpngatewayid
	return P
}

// Create site to site vpn connection
func (s *VPNService) CreateVpnConnection(p *CreateVpnConnectionParams) (*CreateVpnConnectionResponse, error) {
	resp, err := s.cs.newRequest("createVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVpnConnectionResponse
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

type CreateVpnConnectionResponse struct {
	Account              string `json:"account"`
	Cidrlist             string `json:"cidrlist"`
	Created              string `json:"created"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Dpd                  bool   `json:"dpd"`
	Esplifetime          int64  `json:"esplifetime"`
	Esppolicy            string `json:"esppolicy"`
	Forceencap           bool   `json:"forceencap"`
	Fordisplay           bool   `json:"fordisplay"`
	Gateway              string `json:"gateway"`
	Id                   string `json:"id"`
	Ikelifetime          int64  `json:"ikelifetime"`
	Ikepolicy            string `json:"ikepolicy"`
	Ikeversion           string `json:"ikeversion"`
	Ipsecpsk             string `json:"ipsecpsk"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Passive              bool   `json:"passive"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Publicip             string `json:"publicip"`
	Removed              string `json:"removed"`
	S2scustomergatewayid string `json:"s2scustomergatewayid"`
	S2svpngatewayid      string `json:"s2svpngatewayid"`
	Splitconnections     bool   `json:"splitconnections"`
	State                string `json:"state"`
}

type CreateVpnCustomerGatewayParams struct {
	P map[string]interface{}
}

func (P *CreateVpnCustomerGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["cidrlist"]; found {
		u.Set("cidrlist", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["dpd"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("dpd", vv)
	}
	if v, found := P.P["esplifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("esplifetime", vv)
	}
	if v, found := P.P["esppolicy"]; found {
		u.Set("esppolicy", v.(string))
	}
	if v, found := P.P["forceencap"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forceencap", vv)
	}
	if v, found := P.P["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := P.P["ikelifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("ikelifetime", vv)
	}
	if v, found := P.P["ikepolicy"]; found {
		u.Set("ikepolicy", v.(string))
	}
	if v, found := P.P["ikeversion"]; found {
		u.Set("ikeversion", v.(string))
	}
	if v, found := P.P["ipsecpsk"]; found {
		u.Set("ipsecpsk", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["splitconnections"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("splitconnections", vv)
	}
	return u
}

func (P *CreateVpnCustomerGatewayParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetCidrlist(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetCidrlist() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].(string)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetDpd(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dpd"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetDpd() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dpd"].(bool)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetEsplifetime(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["esplifetime"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetEsplifetime() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["esplifetime"].(int64)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetEsppolicy(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["esppolicy"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetEsppolicy() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["esppolicy"].(string)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetForceencap(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forceencap"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetForceencap() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forceencap"].(bool)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetGateway(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gateway"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetGateway() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gateway"].(string)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetIkelifetime(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ikelifetime"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetIkelifetime() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ikelifetime"].(int64)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetIkepolicy(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ikepolicy"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetIkepolicy() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ikepolicy"].(string)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetIkeversion(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ikeversion"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetIkeversion() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ikeversion"].(string)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetIpsecpsk(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipsecpsk"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetIpsecpsk() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipsecpsk"].(string)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *CreateVpnCustomerGatewayParams) SetSplitconnections(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["splitconnections"] = v
}

func (P *CreateVpnCustomerGatewayParams) GetSplitconnections() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["splitconnections"].(bool)
	return value, ok
}

// You should always use this function to get a new CreateVpnCustomerGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateVpnCustomerGatewayParams(cidrlist string, esppolicy string, gateway string, ikepolicy string, ipsecpsk string) *CreateVpnCustomerGatewayParams {
	P := &CreateVpnCustomerGatewayParams{}
	P.P = make(map[string]interface{})
	P.P["cidrlist"] = cidrlist
	P.P["esppolicy"] = esppolicy
	P.P["gateway"] = gateway
	P.P["ikepolicy"] = ikepolicy
	P.P["ipsecpsk"] = ipsecpsk
	return P
}

// Creates site to site vpn customer gateway
func (s *VPNService) CreateVpnCustomerGateway(p *CreateVpnCustomerGatewayParams) (*CreateVpnCustomerGatewayResponse, error) {
	resp, err := s.cs.newRequest("createVpnCustomerGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVpnCustomerGatewayResponse
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

type CreateVpnCustomerGatewayResponse struct {
	Account          string `json:"account"`
	Cidrlist         string `json:"cidrlist"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Dpd              bool   `json:"dpd"`
	Esplifetime      int64  `json:"esplifetime"`
	Esppolicy        string `json:"esppolicy"`
	Forceencap       bool   `json:"forceencap"`
	Gateway          string `json:"gateway"`
	Hasannotations   bool   `json:"hasannotations"`
	Id               string `json:"id"`
	Ikelifetime      int64  `json:"ikelifetime"`
	Ikepolicy        string `json:"ikepolicy"`
	Ikeversion       string `json:"ikeversion"`
	Ipaddress        string `json:"ipaddress"`
	Ipsecpsk         string `json:"ipsecpsk"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Name             string `json:"name"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Removed          string `json:"removed"`
	Splitconnections bool   `json:"splitconnections"`
}

type CreateVpnGatewayParams struct {
	P map[string]interface{}
}

func (P *CreateVpnGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (P *CreateVpnGatewayParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *CreateVpnGatewayParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *CreateVpnGatewayParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *CreateVpnGatewayParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVpnGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateVpnGatewayParams(vpcid string) *CreateVpnGatewayParams {
	P := &CreateVpnGatewayParams{}
	P.P = make(map[string]interface{})
	P.P["vpcid"] = vpcid
	return P
}

// Creates site to site vpn local gateway
func (s *VPNService) CreateVpnGateway(p *CreateVpnGatewayParams) (*CreateVpnGatewayResponse, error) {
	resp, err := s.cs.newRequest("createVpnGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVpnGatewayResponse
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

type CreateVpnGatewayResponse struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Fordisplay bool   `json:"fordisplay"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	Publicip   string `json:"publicip"`
	Removed    string `json:"removed"`
	Vpcid      string `json:"vpcid"`
	Vpcname    string `json:"vpcname"`
}

type DeleteRemoteAccessVpnParams struct {
	P map[string]interface{}
}

func (P *DeleteRemoteAccessVpnParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	return u
}

func (P *DeleteRemoteAccessVpnParams) SetPublicipid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publicipid"] = v
}

func (P *DeleteRemoteAccessVpnParams) GetPublicipid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publicipid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteRemoteAccessVpnParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteRemoteAccessVpnParams(publicipid string) *DeleteRemoteAccessVpnParams {
	P := &DeleteRemoteAccessVpnParams{}
	P.P = make(map[string]interface{})
	P.P["publicipid"] = publicipid
	return P
}

// Destroys a l2tp/ipsec remote access vpn
func (s *VPNService) DeleteRemoteAccessVpn(p *DeleteRemoteAccessVpnParams) (*DeleteRemoteAccessVpnResponse, error) {
	resp, err := s.cs.newRequest("deleteRemoteAccessVpn", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteRemoteAccessVpnResponse
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

type DeleteRemoteAccessVpnResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteVpnConnectionParams struct {
	P map[string]interface{}
}

func (P *DeleteVpnConnectionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteVpnConnectionParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteVpnConnectionParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteVpnConnectionParams(id string) *DeleteVpnConnectionParams {
	P := &DeleteVpnConnectionParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Delete site to site vpn connection
func (s *VPNService) DeleteVpnConnection(p *DeleteVpnConnectionParams) (*DeleteVpnConnectionResponse, error) {
	resp, err := s.cs.newRequest("deleteVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVpnConnectionResponse
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

type DeleteVpnConnectionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteVpnCustomerGatewayParams struct {
	P map[string]interface{}
}

func (P *DeleteVpnCustomerGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteVpnCustomerGatewayParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteVpnCustomerGatewayParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVpnCustomerGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteVpnCustomerGatewayParams(id string) *DeleteVpnCustomerGatewayParams {
	P := &DeleteVpnCustomerGatewayParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Delete site to site vpn customer gateway
func (s *VPNService) DeleteVpnCustomerGateway(p *DeleteVpnCustomerGatewayParams) (*DeleteVpnCustomerGatewayResponse, error) {
	resp, err := s.cs.newRequest("deleteVpnCustomerGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVpnCustomerGatewayResponse
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

type DeleteVpnCustomerGatewayResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteVpnGatewayParams struct {
	P map[string]interface{}
}

func (P *DeleteVpnGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteVpnGatewayParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteVpnGatewayParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVpnGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteVpnGatewayParams(id string) *DeleteVpnGatewayParams {
	P := &DeleteVpnGatewayParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Delete site to site vpn gateway
func (s *VPNService) DeleteVpnGateway(p *DeleteVpnGatewayParams) (*DeleteVpnGatewayResponse, error) {
	resp, err := s.cs.newRequest("deleteVpnGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVpnGatewayResponse
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

type DeleteVpnGatewayResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListRemoteAccessVpnsParams struct {
	P map[string]interface{}
}

func (P *ListRemoteAccessVpnsParams) toURLValues() url.Values {
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	return u
}

func (P *ListRemoteAccessVpnsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListRemoteAccessVpnsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListRemoteAccessVpnsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListRemoteAccessVpnsParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListRemoteAccessVpnsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListRemoteAccessVpnsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListRemoteAccessVpnsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListRemoteAccessVpnsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListRemoteAccessVpnsParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListRemoteAccessVpnsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListRemoteAccessVpnsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListRemoteAccessVpnsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListRemoteAccessVpnsParams) SetPublicipid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["publicipid"] = v
}

func (P *ListRemoteAccessVpnsParams) GetPublicipid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["publicipid"].(string)
	return value, ok
}

// You should always use this function to get a new ListRemoteAccessVpnsParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListRemoteAccessVpnsParams() *ListRemoteAccessVpnsParams {
	P := &ListRemoteAccessVpnsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetRemoteAccessVpnByID(id string, opts ...OptionFunc) (*RemoteAccessVpn, int, error) {
	P := &ListRemoteAccessVpnsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListRemoteAccessVpns(P)
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
		return l.RemoteAccessVpns[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for RemoteAccessVpn UUID: %s!", id)
}

// Lists remote access vpns
func (s *VPNService) ListRemoteAccessVpns(p *ListRemoteAccessVpnsParams) (*ListRemoteAccessVpnsResponse, error) {
	resp, err := s.cs.newRequest("listRemoteAccessVpns", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRemoteAccessVpnsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListRemoteAccessVpnsResponse struct {
	Count            int                `json:"count"`
	RemoteAccessVpns []*RemoteAccessVpn `json:"remoteaccessvpn"`
}

type RemoteAccessVpn struct {
	Account      string `json:"account"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Fordisplay   bool   `json:"fordisplay"`
	Id           string `json:"id"`
	Iprange      string `json:"iprange"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Presharedkey string `json:"presharedkey"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Publicip     string `json:"publicip"`
	Publicipid   string `json:"publicipid"`
	State        string `json:"state"`
}

type ListVpnConnectionsParams struct {
	P map[string]interface{}
}

func (P *ListVpnConnectionsParams) toURLValues() url.Values {
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
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (P *ListVpnConnectionsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListVpnConnectionsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListVpnConnectionsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListVpnConnectionsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListVpnConnectionsParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListVpnConnectionsParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListVpnConnectionsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListVpnConnectionsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListVpnConnectionsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListVpnConnectionsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListVpnConnectionsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListVpnConnectionsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListVpnConnectionsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListVpnConnectionsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListVpnConnectionsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListVpnConnectionsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListVpnConnectionsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListVpnConnectionsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListVpnConnectionsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListVpnConnectionsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListVpnConnectionsParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *ListVpnConnectionsParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVpnConnectionsParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnConnectionsParams() *ListVpnConnectionsParams {
	P := &ListVpnConnectionsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnConnectionByID(id string, opts ...OptionFunc) (*VpnConnection, int, error) {
	P := &ListVpnConnectionsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVpnConnections(P)
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
		return l.VpnConnections[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VpnConnection UUID: %s!", id)
}

// Lists site to site vpn connection gateways
func (s *VPNService) ListVpnConnections(p *ListVpnConnectionsParams) (*ListVpnConnectionsResponse, error) {
	resp, err := s.cs.newRequest("listVpnConnections", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnConnectionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVpnConnectionsResponse struct {
	Count          int              `json:"count"`
	VpnConnections []*VpnConnection `json:"vpnconnection"`
}

type VpnConnection struct {
	Account              string `json:"account"`
	Cidrlist             string `json:"cidrlist"`
	Created              string `json:"created"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Dpd                  bool   `json:"dpd"`
	Esplifetime          int64  `json:"esplifetime"`
	Esppolicy            string `json:"esppolicy"`
	Forceencap           bool   `json:"forceencap"`
	Fordisplay           bool   `json:"fordisplay"`
	Gateway              string `json:"gateway"`
	Id                   string `json:"id"`
	Ikelifetime          int64  `json:"ikelifetime"`
	Ikepolicy            string `json:"ikepolicy"`
	Ikeversion           string `json:"ikeversion"`
	Ipsecpsk             string `json:"ipsecpsk"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Passive              bool   `json:"passive"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Publicip             string `json:"publicip"`
	Removed              string `json:"removed"`
	S2scustomergatewayid string `json:"s2scustomergatewayid"`
	S2svpngatewayid      string `json:"s2svpngatewayid"`
	Splitconnections     bool   `json:"splitconnections"`
	State                string `json:"state"`
}

type ListVpnCustomerGatewaysParams struct {
	P map[string]interface{}
}

func (P *ListVpnCustomerGatewaysParams) toURLValues() url.Values {
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *ListVpnCustomerGatewaysParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListVpnCustomerGatewaysParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListVpnCustomerGatewaysParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListVpnCustomerGatewaysParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListVpnCustomerGatewaysParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListVpnCustomerGatewaysParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListVpnCustomerGatewaysParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListVpnCustomerGatewaysParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListVpnCustomerGatewaysParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListVpnCustomerGatewaysParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListVpnCustomerGatewaysParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListVpnCustomerGatewaysParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListVpnCustomerGatewaysParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListVpnCustomerGatewaysParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListVpnCustomerGatewaysParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListVpnCustomerGatewaysParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListVpnCustomerGatewaysParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListVpnCustomerGatewaysParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVpnCustomerGatewaysParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnCustomerGatewaysParams() *ListVpnCustomerGatewaysParams {
	P := &ListVpnCustomerGatewaysParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnCustomerGatewayID(keyword string, opts ...OptionFunc) (string, int, error) {
	P := &ListVpnCustomerGatewaysParams{}
	P.P = make(map[string]interface{})

	P.P["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVpnCustomerGateways(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.VpnCustomerGateways[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VpnCustomerGateways {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnCustomerGatewayByName(name string, opts ...OptionFunc) (*VpnCustomerGateway, int, error) {
	id, count, err := s.GetVpnCustomerGatewayID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVpnCustomerGatewayByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnCustomerGatewayByID(id string, opts ...OptionFunc) (*VpnCustomerGateway, int, error) {
	P := &ListVpnCustomerGatewaysParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVpnCustomerGateways(P)
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
		return l.VpnCustomerGateways[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VpnCustomerGateway UUID: %s!", id)
}

// Lists site to site vpn customer gateways
func (s *VPNService) ListVpnCustomerGateways(p *ListVpnCustomerGatewaysParams) (*ListVpnCustomerGatewaysResponse, error) {
	resp, err := s.cs.newRequest("listVpnCustomerGateways", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnCustomerGatewaysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVpnCustomerGatewaysResponse struct {
	Count               int                   `json:"count"`
	VpnCustomerGateways []*VpnCustomerGateway `json:"vpncustomergateway"`
}

type VpnCustomerGateway struct {
	Account          string `json:"account"`
	Cidrlist         string `json:"cidrlist"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Dpd              bool   `json:"dpd"`
	Esplifetime      int64  `json:"esplifetime"`
	Esppolicy        string `json:"esppolicy"`
	Forceencap       bool   `json:"forceencap"`
	Gateway          string `json:"gateway"`
	Hasannotations   bool   `json:"hasannotations"`
	Id               string `json:"id"`
	Ikelifetime      int64  `json:"ikelifetime"`
	Ikepolicy        string `json:"ikepolicy"`
	Ikeversion       string `json:"ikeversion"`
	Ipaddress        string `json:"ipaddress"`
	Ipsecpsk         string `json:"ipsecpsk"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Name             string `json:"name"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Removed          string `json:"removed"`
	Splitconnections bool   `json:"splitconnections"`
}

type ListVpnGatewaysParams struct {
	P map[string]interface{}
}

func (P *ListVpnGatewaysParams) toURLValues() url.Values {
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
	if v, found := P.P["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (P *ListVpnGatewaysParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListVpnGatewaysParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListVpnGatewaysParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListVpnGatewaysParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListVpnGatewaysParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListVpnGatewaysParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListVpnGatewaysParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListVpnGatewaysParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListVpnGatewaysParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListVpnGatewaysParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListVpnGatewaysParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListVpnGatewaysParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListVpnGatewaysParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListVpnGatewaysParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListVpnGatewaysParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListVpnGatewaysParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListVpnGatewaysParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListVpnGatewaysParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListVpnGatewaysParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListVpnGatewaysParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListVpnGatewaysParams) SetVpcid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["vpcid"] = v
}

func (P *ListVpnGatewaysParams) GetVpcid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["vpcid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVpnGatewaysParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnGatewaysParams() *ListVpnGatewaysParams {
	P := &ListVpnGatewaysParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnGatewayByID(id string, opts ...OptionFunc) (*VpnGateway, int, error) {
	P := &ListVpnGatewaysParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVpnGateways(P)
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
		return l.VpnGateways[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VpnGateway UUID: %s!", id)
}

// Lists site 2 site vpn gateways
func (s *VPNService) ListVpnGateways(p *ListVpnGatewaysParams) (*ListVpnGatewaysResponse, error) {
	resp, err := s.cs.newRequest("listVpnGateways", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnGatewaysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVpnGatewaysResponse struct {
	Count       int           `json:"count"`
	VpnGateways []*VpnGateway `json:"vpngateway"`
}

type VpnGateway struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Fordisplay bool   `json:"fordisplay"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	Publicip   string `json:"publicip"`
	Removed    string `json:"removed"`
	Vpcid      string `json:"vpcid"`
	Vpcname    string `json:"vpcname"`
}

type ListVpnUsersParams struct {
	P map[string]interface{}
}

func (P *ListVpnUsersParams) toURLValues() url.Values {
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *ListVpnUsersParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListVpnUsersParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListVpnUsersParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListVpnUsersParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListVpnUsersParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListVpnUsersParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListVpnUsersParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListVpnUsersParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListVpnUsersParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListVpnUsersParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListVpnUsersParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListVpnUsersParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListVpnUsersParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListVpnUsersParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListVpnUsersParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListVpnUsersParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListVpnUsersParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListVpnUsersParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListVpnUsersParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *ListVpnUsersParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new ListVpnUsersParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnUsersParams() *ListVpnUsersParams {
	P := &ListVpnUsersParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnUserByID(id string, opts ...OptionFunc) (*VpnUser, int, error) {
	P := &ListVpnUsersParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVpnUsers(P)
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
		return l.VpnUsers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for VpnUser UUID: %s!", id)
}

// Lists vpn users
func (s *VPNService) ListVpnUsers(p *ListVpnUsersParams) (*ListVpnUsersResponse, error) {
	resp, err := s.cs.newRequest("listVpnUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVpnUsersResponse struct {
	Count    int        `json:"count"`
	VpnUsers []*VpnUser `json:"vpnuser"`
}

type VpnUser struct {
	Account   string `json:"account"`
	Domain    string `json:"domain"`
	Domainid  string `json:"domainid"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
	State     string `json:"state"`
	Username  string `json:"username"`
}

type RemoveVpnUserParams struct {
	P map[string]interface{}
}

func (P *RemoveVpnUserParams) toURLValues() url.Values {
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *RemoveVpnUserParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *RemoveVpnUserParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *RemoveVpnUserParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *RemoveVpnUserParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *RemoveVpnUserParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *RemoveVpnUserParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *RemoveVpnUserParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *RemoveVpnUserParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveVpnUserParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewRemoveVpnUserParams(username string) *RemoveVpnUserParams {
	P := &RemoveVpnUserParams{}
	P.P = make(map[string]interface{})
	P.P["username"] = username
	return P
}

// Removes vpn user
func (s *VPNService) RemoveVpnUser(p *RemoveVpnUserParams) (*RemoveVpnUserResponse, error) {
	resp, err := s.cs.newRequest("removeVpnUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveVpnUserResponse
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

type RemoveVpnUserResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ResetVpnConnectionParams struct {
	P map[string]interface{}
}

func (P *ResetVpnConnectionParams) toURLValues() url.Values {
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
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ResetVpnConnectionParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ResetVpnConnectionParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ResetVpnConnectionParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ResetVpnConnectionParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ResetVpnConnectionParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ResetVpnConnectionParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ResetVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewResetVpnConnectionParams(id string) *ResetVpnConnectionParams {
	P := &ResetVpnConnectionParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Reset site to site vpn connection
func (s *VPNService) ResetVpnConnection(p *ResetVpnConnectionParams) (*ResetVpnConnectionResponse, error) {
	resp, err := s.cs.newRequest("resetVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetVpnConnectionResponse
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

type ResetVpnConnectionResponse struct {
	Account              string `json:"account"`
	Cidrlist             string `json:"cidrlist"`
	Created              string `json:"created"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Dpd                  bool   `json:"dpd"`
	Esplifetime          int64  `json:"esplifetime"`
	Esppolicy            string `json:"esppolicy"`
	Forceencap           bool   `json:"forceencap"`
	Fordisplay           bool   `json:"fordisplay"`
	Gateway              string `json:"gateway"`
	Id                   string `json:"id"`
	Ikelifetime          int64  `json:"ikelifetime"`
	Ikepolicy            string `json:"ikepolicy"`
	Ikeversion           string `json:"ikeversion"`
	Ipsecpsk             string `json:"ipsecpsk"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Passive              bool   `json:"passive"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Publicip             string `json:"publicip"`
	Removed              string `json:"removed"`
	S2scustomergatewayid string `json:"s2scustomergatewayid"`
	S2svpngatewayid      string `json:"s2svpngatewayid"`
	Splitconnections     bool   `json:"splitconnections"`
	State                string `json:"state"`
}

type UpdateRemoteAccessVpnParams struct {
	P map[string]interface{}
}

func (P *UpdateRemoteAccessVpnParams) toURLValues() url.Values {
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
	return u
}

func (P *UpdateRemoteAccessVpnParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateRemoteAccessVpnParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateRemoteAccessVpnParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateRemoteAccessVpnParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateRemoteAccessVpnParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateRemoteAccessVpnParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateRemoteAccessVpnParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewUpdateRemoteAccessVpnParams(id string) *UpdateRemoteAccessVpnParams {
	P := &UpdateRemoteAccessVpnParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates remote access vpn
func (s *VPNService) UpdateRemoteAccessVpn(p *UpdateRemoteAccessVpnParams) (*UpdateRemoteAccessVpnResponse, error) {
	resp, err := s.cs.newRequest("updateRemoteAccessVpn", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateRemoteAccessVpnResponse
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

type UpdateRemoteAccessVpnResponse struct {
	Account      string `json:"account"`
	Domain       string `json:"domain"`
	Domainid     string `json:"domainid"`
	Fordisplay   bool   `json:"fordisplay"`
	Id           string `json:"id"`
	Iprange      string `json:"iprange"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Presharedkey string `json:"presharedkey"`
	Project      string `json:"project"`
	Projectid    string `json:"projectid"`
	Publicip     string `json:"publicip"`
	Publicipid   string `json:"publicipid"`
	State        string `json:"state"`
}

type UpdateVpnConnectionParams struct {
	P map[string]interface{}
}

func (P *UpdateVpnConnectionParams) toURLValues() url.Values {
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
	return u
}

func (P *UpdateVpnConnectionParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateVpnConnectionParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateVpnConnectionParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateVpnConnectionParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateVpnConnectionParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateVpnConnectionParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewUpdateVpnConnectionParams(id string) *UpdateVpnConnectionParams {
	P := &UpdateVpnConnectionParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates site to site vpn connection
func (s *VPNService) UpdateVpnConnection(p *UpdateVpnConnectionParams) (*UpdateVpnConnectionResponse, error) {
	resp, err := s.cs.newRequest("updateVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVpnConnectionResponse
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

type UpdateVpnConnectionResponse struct {
	Account              string `json:"account"`
	Cidrlist             string `json:"cidrlist"`
	Created              string `json:"created"`
	Domain               string `json:"domain"`
	Domainid             string `json:"domainid"`
	Dpd                  bool   `json:"dpd"`
	Esplifetime          int64  `json:"esplifetime"`
	Esppolicy            string `json:"esppolicy"`
	Forceencap           bool   `json:"forceencap"`
	Fordisplay           bool   `json:"fordisplay"`
	Gateway              string `json:"gateway"`
	Id                   string `json:"id"`
	Ikelifetime          int64  `json:"ikelifetime"`
	Ikepolicy            string `json:"ikepolicy"`
	Ikeversion           string `json:"ikeversion"`
	Ipsecpsk             string `json:"ipsecpsk"`
	JobID                string `json:"jobid"`
	Jobstatus            int    `json:"jobstatus"`
	Passive              bool   `json:"passive"`
	Project              string `json:"project"`
	Projectid            string `json:"projectid"`
	Publicip             string `json:"publicip"`
	Removed              string `json:"removed"`
	S2scustomergatewayid string `json:"s2scustomergatewayid"`
	S2svpngatewayid      string `json:"s2svpngatewayid"`
	Splitconnections     bool   `json:"splitconnections"`
	State                string `json:"state"`
}

type UpdateVpnCustomerGatewayParams struct {
	P map[string]interface{}
}

func (P *UpdateVpnCustomerGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["cidrlist"]; found {
		u.Set("cidrlist", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["dpd"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("dpd", vv)
	}
	if v, found := P.P["esplifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("esplifetime", vv)
	}
	if v, found := P.P["esppolicy"]; found {
		u.Set("esppolicy", v.(string))
	}
	if v, found := P.P["forceencap"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forceencap", vv)
	}
	if v, found := P.P["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["ikelifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("ikelifetime", vv)
	}
	if v, found := P.P["ikepolicy"]; found {
		u.Set("ikepolicy", v.(string))
	}
	if v, found := P.P["ikeversion"]; found {
		u.Set("ikeversion", v.(string))
	}
	if v, found := P.P["ipsecpsk"]; found {
		u.Set("ipsecpsk", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["splitconnections"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("splitconnections", vv)
	}
	return u
}

func (P *UpdateVpnCustomerGatewayParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetCidrlist(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cidrlist"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetCidrlist() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cidrlist"].(string)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetDpd(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dpd"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetDpd() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dpd"].(bool)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetEsplifetime(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["esplifetime"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetEsplifetime() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["esplifetime"].(int64)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetEsppolicy(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["esppolicy"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetEsppolicy() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["esppolicy"].(string)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetForceencap(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forceencap"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetForceencap() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forceencap"].(bool)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetGateway(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["gateway"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetGateway() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["gateway"].(string)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetIkelifetime(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ikelifetime"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetIkelifetime() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ikelifetime"].(int64)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetIkepolicy(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ikepolicy"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetIkepolicy() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ikepolicy"].(string)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetIkeversion(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ikeversion"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetIkeversion() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ikeversion"].(string)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetIpsecpsk(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipsecpsk"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetIpsecpsk() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipsecpsk"].(string)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateVpnCustomerGatewayParams) SetSplitconnections(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["splitconnections"] = v
}

func (P *UpdateVpnCustomerGatewayParams) GetSplitconnections() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["splitconnections"].(bool)
	return value, ok
}

// You should always use this function to get a new UpdateVpnCustomerGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewUpdateVpnCustomerGatewayParams(cidrlist string, esppolicy string, gateway string, id string, ikepolicy string, ipsecpsk string) *UpdateVpnCustomerGatewayParams {
	P := &UpdateVpnCustomerGatewayParams{}
	P.P = make(map[string]interface{})
	P.P["cidrlist"] = cidrlist
	P.P["esppolicy"] = esppolicy
	P.P["gateway"] = gateway
	P.P["id"] = id
	P.P["ikepolicy"] = ikepolicy
	P.P["ipsecpsk"] = ipsecpsk
	return P
}

// Update site to site vpn customer gateway
func (s *VPNService) UpdateVpnCustomerGateway(p *UpdateVpnCustomerGatewayParams) (*UpdateVpnCustomerGatewayResponse, error) {
	resp, err := s.cs.newRequest("updateVpnCustomerGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVpnCustomerGatewayResponse
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

type UpdateVpnCustomerGatewayResponse struct {
	Account          string `json:"account"`
	Cidrlist         string `json:"cidrlist"`
	Domain           string `json:"domain"`
	Domainid         string `json:"domainid"`
	Dpd              bool   `json:"dpd"`
	Esplifetime      int64  `json:"esplifetime"`
	Esppolicy        string `json:"esppolicy"`
	Forceencap       bool   `json:"forceencap"`
	Gateway          string `json:"gateway"`
	Hasannotations   bool   `json:"hasannotations"`
	Id               string `json:"id"`
	Ikelifetime      int64  `json:"ikelifetime"`
	Ikepolicy        string `json:"ikepolicy"`
	Ikeversion       string `json:"ikeversion"`
	Ipaddress        string `json:"ipaddress"`
	Ipsecpsk         string `json:"ipsecpsk"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Name             string `json:"name"`
	Project          string `json:"project"`
	Projectid        string `json:"projectid"`
	Removed          string `json:"removed"`
	Splitconnections bool   `json:"splitconnections"`
}

type UpdateVpnGatewayParams struct {
	P map[string]interface{}
}

func (P *UpdateVpnGatewayParams) toURLValues() url.Values {
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
	return u
}

func (P *UpdateVpnGatewayParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateVpnGatewayParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateVpnGatewayParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *UpdateVpnGatewayParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *UpdateVpnGatewayParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateVpnGatewayParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVpnGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewUpdateVpnGatewayParams(id string) *UpdateVpnGatewayParams {
	P := &UpdateVpnGatewayParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates site to site vpn local gateway
func (s *VPNService) UpdateVpnGateway(p *UpdateVpnGatewayParams) (*UpdateVpnGatewayResponse, error) {
	resp, err := s.cs.newRequest("updateVpnGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVpnGatewayResponse
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

type UpdateVpnGatewayResponse struct {
	Account    string `json:"account"`
	Domain     string `json:"domain"`
	Domainid   string `json:"domainid"`
	Fordisplay bool   `json:"fordisplay"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Project    string `json:"project"`
	Projectid  string `json:"projectid"`
	Publicip   string `json:"publicip"`
	Removed    string `json:"removed"`
	Vpcid      string `json:"vpcid"`
	Vpcname    string `json:"vpcname"`
}
