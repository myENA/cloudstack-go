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

type NetworkOfferingServiceIface interface {
	CreateNetworkOffering(P *CreateNetworkOfferingParams) (*CreateNetworkOfferingResponse, error)
	NewCreateNetworkOfferingParams(displaytext string, guestiptype string, name string, traffictype string) *CreateNetworkOfferingParams
	DeleteNetworkOffering(P *DeleteNetworkOfferingParams) (*DeleteNetworkOfferingResponse, error)
	NewDeleteNetworkOfferingParams(id string) *DeleteNetworkOfferingParams
	ListNetworkOfferings(P *ListNetworkOfferingsParams) (*ListNetworkOfferingsResponse, error)
	NewListNetworkOfferingsParams() *ListNetworkOfferingsParams
	GetNetworkOfferingID(name string, opts ...OptionFunc) (string, int, error)
	GetNetworkOfferingByName(name string, opts ...OptionFunc) (*NetworkOffering, int, error)
	GetNetworkOfferingByID(id string, opts ...OptionFunc) (*NetworkOffering, int, error)
	UpdateNetworkOffering(P *UpdateNetworkOfferingParams) (*UpdateNetworkOfferingResponse, error)
	NewUpdateNetworkOfferingParams() *UpdateNetworkOfferingParams
}

type CreateNetworkOfferingParams struct {
	P map[string]interface{}
}

func (P *CreateNetworkOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["availability"]; found {
		u.Set("availability", v.(string))
	}
	if v, found := P.P["conservemode"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("conservemode", vv)
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("domainid", vv)
	}
	if v, found := P.P["egressdefaultpolicy"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("egressdefaultpolicy", vv)
	}
	if v, found := P.P["enable"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enable", vv)
	}
	if v, found := P.P["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := P.P["guestiptype"]; found {
		u.Set("guestiptype", v.(string))
	}
	if v, found := P.P["ispersistent"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispersistent", vv)
	}
	if v, found := P.P["keepaliveenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("keepaliveenabled", vv)
	}
	if v, found := P.P["maxconnections"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxconnections", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networkrate"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("networkrate", vv)
	}
	if v, found := P.P["servicecapabilitylist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("servicecapabilitylist[%d].key", i), k)
			u.Set(fmt.Sprintf("servicecapabilitylist[%d].value", i), m[k])
		}
	}
	if v, found := P.P["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := P.P["serviceproviderlist"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("serviceproviderlist[%d].service", i), k)
			u.Set(fmt.Sprintf("serviceproviderlist[%d].provider", i), m[k])
		}
	}
	if v, found := P.P["specifyipranges"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyipranges", vv)
	}
	if v, found := P.P["specifyvlan"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyvlan", vv)
	}
	if v, found := P.P["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
	}
	if v, found := P.P["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := P.P["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneid", vv)
	}
	return u
}

func (P *CreateNetworkOfferingParams) SetAvailability(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["availability"] = v
}

func (P *CreateNetworkOfferingParams) GetAvailability() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["availability"].(string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetConservemode(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["conservemode"] = v
}

func (P *CreateNetworkOfferingParams) GetConservemode() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["conservemode"].(bool)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *CreateNetworkOfferingParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *CreateNetworkOfferingParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetDomainid(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateNetworkOfferingParams) GetDomainid() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].([]string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetEgressdefaultpolicy(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["egressdefaultpolicy"] = v
}

func (P *CreateNetworkOfferingParams) GetEgressdefaultpolicy() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["egressdefaultpolicy"].(bool)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetEnable(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["enable"] = v
}

func (P *CreateNetworkOfferingParams) GetEnable() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["enable"].(bool)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetForvpc(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forvpc"] = v
}

func (P *CreateNetworkOfferingParams) GetForvpc() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forvpc"].(bool)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetGuestiptype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["guestiptype"] = v
}

func (P *CreateNetworkOfferingParams) GetGuestiptype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["guestiptype"].(string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetIspersistent(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ispersistent"] = v
}

func (P *CreateNetworkOfferingParams) GetIspersistent() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ispersistent"].(bool)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetKeepaliveenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keepaliveenabled"] = v
}

func (P *CreateNetworkOfferingParams) GetKeepaliveenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keepaliveenabled"].(bool)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetMaxconnections(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxconnections"] = v
}

func (P *CreateNetworkOfferingParams) GetMaxconnections() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxconnections"].(int)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateNetworkOfferingParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetNetworkrate(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkrate"] = v
}

func (P *CreateNetworkOfferingParams) GetNetworkrate() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkrate"].(int)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetServicecapabilitylist(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["servicecapabilitylist"] = v
}

func (P *CreateNetworkOfferingParams) GetServicecapabilitylist() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["servicecapabilitylist"].(map[string]string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetServiceofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingid"] = v
}

func (P *CreateNetworkOfferingParams) GetServiceofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingid"].(string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetServiceproviderlist(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceproviderlist"] = v
}

func (P *CreateNetworkOfferingParams) GetServiceproviderlist() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceproviderlist"].(map[string]string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetSpecifyipranges(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["specifyipranges"] = v
}

func (P *CreateNetworkOfferingParams) GetSpecifyipranges() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["specifyipranges"].(bool)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetSpecifyvlan(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["specifyvlan"] = v
}

func (P *CreateNetworkOfferingParams) GetSpecifyvlan() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["specifyvlan"].(bool)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetSupportedservices(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["supportedservices"] = v
}

func (P *CreateNetworkOfferingParams) GetSupportedservices() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["supportedservices"].([]string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetTags(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *CreateNetworkOfferingParams) GetTags() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetTraffictype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["traffictype"] = v
}

func (P *CreateNetworkOfferingParams) GetTraffictype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["traffictype"].(string)
	return value, ok
}

func (P *CreateNetworkOfferingParams) SetZoneid(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateNetworkOfferingParams) GetZoneid() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].([]string)
	return value, ok
}

// You should always use this function to get a new CreateNetworkOfferingParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewCreateNetworkOfferingParams(displaytext string, guestiptype string, name string, traffictype string) *CreateNetworkOfferingParams {
	P := &CreateNetworkOfferingParams{}
	P.P = make(map[string]interface{})
	P.P["displaytext"] = displaytext
	P.P["guestiptype"] = guestiptype
	P.P["name"] = name
	P.P["traffictype"] = traffictype
	return P
}

// Creates a network offering.
func (s *NetworkOfferingService) CreateNetworkOffering(p *CreateNetworkOfferingParams) (*CreateNetworkOfferingResponse, error) {
	resp, err := s.cs.newRequest("createNetworkOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateNetworkOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateNetworkOfferingResponse struct {
	Availability             string                                 `json:"availability"`
	Conservemode             bool                                   `json:"conservemode"`
	Created                  string                                 `json:"created"`
	Details                  map[string]string                      `json:"details"`
	Displaytext              string                                 `json:"displaytext"`
	Domain                   string                                 `json:"domain"`
	Domainid                 string                                 `json:"domainid"`
	Egressdefaultpolicy      bool                                   `json:"egressdefaultpolicy"`
	Forvpc                   bool                                   `json:"forvpc"`
	Guestiptype              string                                 `json:"guestiptype"`
	Hasannotations           bool                                   `json:"hasannotations"`
	Id                       string                                 `json:"id"`
	Isdefault                bool                                   `json:"isdefault"`
	Ispersistent             bool                                   `json:"ispersistent"`
	JobID                    string                                 `json:"jobid"`
	Jobstatus                int                                    `json:"jobstatus"`
	Maxconnections           int                                    `json:"maxconnections"`
	Name                     string                                 `json:"name"`
	Networkrate              int                                    `json:"networkrate"`
	Service                  []CreateNetworkOfferingResponseService `json:"service"`
	Serviceofferingid        string                                 `json:"serviceofferingid"`
	Specifyipranges          bool                                   `json:"specifyipranges"`
	Specifyvlan              bool                                   `json:"specifyvlan"`
	State                    string                                 `json:"state"`
	Supportspublicaccess     bool                                   `json:"supportspublicaccess"`
	Supportsstrechedl2subnet bool                                   `json:"supportsstrechedl2subnet"`
	Tags                     string                                 `json:"tags"`
	Traffictype              string                                 `json:"traffictype"`
	Zone                     string                                 `json:"zone"`
	Zoneid                   string                                 `json:"zoneid"`
}

type CreateNetworkOfferingResponseService struct {
	Capability []CreateNetworkOfferingResponseServiceCapability `json:"capability"`
	Name       string                                           `json:"name"`
	Provider   []CreateNetworkOfferingResponseServiceProvider   `json:"provider"`
}

type CreateNetworkOfferingResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type CreateNetworkOfferingResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type DeleteNetworkOfferingParams struct {
	P map[string]interface{}
}

func (P *DeleteNetworkOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteNetworkOfferingParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteNetworkOfferingParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteNetworkOfferingParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewDeleteNetworkOfferingParams(id string) *DeleteNetworkOfferingParams {
	P := &DeleteNetworkOfferingParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a network offering.
func (s *NetworkOfferingService) DeleteNetworkOffering(p *DeleteNetworkOfferingParams) (*DeleteNetworkOfferingResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteNetworkOfferingResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteNetworkOfferingResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteNetworkOfferingResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListNetworkOfferingsParams struct {
	P map[string]interface{}
}

func (P *ListNetworkOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["availability"]; found {
		u.Set("availability", v.(string))
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := P.P["guestiptype"]; found {
		u.Set("guestiptype", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["isdefault"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdefault", vv)
	}
	if v, found := P.P["istagged"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("istagged", vv)
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
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
	if v, found := P.P["sourcenatsupported"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sourcenatsupported", vv)
	}
	if v, found := P.P["specifyipranges"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyipranges", vv)
	}
	if v, found := P.P["specifyvlan"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyvlan", vv)
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["supportedservices"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("supportedservices", vv)
	}
	if v, found := P.P["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := P.P["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListNetworkOfferingsParams) SetAvailability(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["availability"] = v
}

func (P *ListNetworkOfferingsParams) GetAvailability() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["availability"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *ListNetworkOfferingsParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListNetworkOfferingsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetForvpc(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["forvpc"] = v
}

func (P *ListNetworkOfferingsParams) GetForvpc() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["forvpc"].(bool)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetGuestiptype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["guestiptype"] = v
}

func (P *ListNetworkOfferingsParams) GetGuestiptype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["guestiptype"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListNetworkOfferingsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetIsdefault(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isdefault"] = v
}

func (P *ListNetworkOfferingsParams) GetIsdefault() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isdefault"].(bool)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetIstagged(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["istagged"] = v
}

func (P *ListNetworkOfferingsParams) GetIstagged() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["istagged"].(bool)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNetworkOfferingsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListNetworkOfferingsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListNetworkOfferingsParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNetworkOfferingsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNetworkOfferingsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetSourcenatsupported(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sourcenatsupported"] = v
}

func (P *ListNetworkOfferingsParams) GetSourcenatsupported() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sourcenatsupported"].(bool)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetSpecifyipranges(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["specifyipranges"] = v
}

func (P *ListNetworkOfferingsParams) GetSpecifyipranges() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["specifyipranges"].(bool)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetSpecifyvlan(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["specifyvlan"] = v
}

func (P *ListNetworkOfferingsParams) GetSpecifyvlan() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["specifyvlan"].(bool)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListNetworkOfferingsParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetSupportedservices(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["supportedservices"] = v
}

func (P *ListNetworkOfferingsParams) GetSupportedservices() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["supportedservices"].([]string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetTags(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListNetworkOfferingsParams) GetTags() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetTraffictype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["traffictype"] = v
}

func (P *ListNetworkOfferingsParams) GetTraffictype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["traffictype"].(string)
	return value, ok
}

func (P *ListNetworkOfferingsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListNetworkOfferingsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNetworkOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewListNetworkOfferingsParams() *ListNetworkOfferingsParams {
	P := &ListNetworkOfferingsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkOfferingService) GetNetworkOfferingID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListNetworkOfferingsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListNetworkOfferings(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.NetworkOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.NetworkOfferings {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkOfferingService) GetNetworkOfferingByName(name string, opts ...OptionFunc) (*NetworkOffering, int, error) {
	id, count, err := s.GetNetworkOfferingID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetNetworkOfferingByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkOfferingService) GetNetworkOfferingByID(id string, opts ...OptionFunc) (*NetworkOffering, int, error) {
	P := &ListNetworkOfferingsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListNetworkOfferings(P)
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
		return l.NetworkOfferings[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for NetworkOffering UUID: %s!", id)
}

// Lists all available network offerings.
func (s *NetworkOfferingService) ListNetworkOfferings(p *ListNetworkOfferingsParams) (*ListNetworkOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworkOfferingsResponse struct {
	Count            int                `json:"count"`
	NetworkOfferings []*NetworkOffering `json:"networkoffering"`
}

type NetworkOffering struct {
	Availability             string                           `json:"availability"`
	Conservemode             bool                             `json:"conservemode"`
	Created                  string                           `json:"created"`
	Details                  map[string]string                `json:"details"`
	Displaytext              string                           `json:"displaytext"`
	Domain                   string                           `json:"domain"`
	Domainid                 string                           `json:"domainid"`
	Egressdefaultpolicy      bool                             `json:"egressdefaultpolicy"`
	Forvpc                   bool                             `json:"forvpc"`
	Guestiptype              string                           `json:"guestiptype"`
	Hasannotations           bool                             `json:"hasannotations"`
	Id                       string                           `json:"id"`
	Isdefault                bool                             `json:"isdefault"`
	Ispersistent             bool                             `json:"ispersistent"`
	JobID                    string                           `json:"jobid"`
	Jobstatus                int                              `json:"jobstatus"`
	Maxconnections           int                              `json:"maxconnections"`
	Name                     string                           `json:"name"`
	Networkrate              int                              `json:"networkrate"`
	Service                  []NetworkOfferingServiceInternal `json:"service"`
	Serviceofferingid        string                           `json:"serviceofferingid"`
	Specifyipranges          bool                             `json:"specifyipranges"`
	Specifyvlan              bool                             `json:"specifyvlan"`
	State                    string                           `json:"state"`
	Supportspublicaccess     bool                             `json:"supportspublicaccess"`
	Supportsstrechedl2subnet bool                             `json:"supportsstrechedl2subnet"`
	Tags                     string                           `json:"tags"`
	Traffictype              string                           `json:"traffictype"`
	Zone                     string                           `json:"zone"`
	Zoneid                   string                           `json:"zoneid"`
}

type NetworkOfferingServiceInternal struct {
	Capability []NetworkOfferingServiceInternalCapability `json:"capability"`
	Name       string                                     `json:"name"`
	Provider   []NetworkOfferingServiceInternalProvider   `json:"provider"`
}

type NetworkOfferingServiceInternalProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type NetworkOfferingServiceInternalCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}

type UpdateNetworkOfferingParams struct {
	P map[string]interface{}
}

func (P *UpdateNetworkOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["availability"]; found {
		u.Set("availability", v.(string))
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keepaliveenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("keepaliveenabled", vv)
	}
	if v, found := P.P["maxconnections"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxconnections", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *UpdateNetworkOfferingParams) SetAvailability(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["availability"] = v
}

func (P *UpdateNetworkOfferingParams) GetAvailability() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["availability"].(string)
	return value, ok
}

func (P *UpdateNetworkOfferingParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *UpdateNetworkOfferingParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *UpdateNetworkOfferingParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UpdateNetworkOfferingParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UpdateNetworkOfferingParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateNetworkOfferingParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateNetworkOfferingParams) SetKeepaliveenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keepaliveenabled"] = v
}

func (P *UpdateNetworkOfferingParams) GetKeepaliveenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keepaliveenabled"].(bool)
	return value, ok
}

func (P *UpdateNetworkOfferingParams) SetMaxconnections(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxconnections"] = v
}

func (P *UpdateNetworkOfferingParams) GetMaxconnections() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxconnections"].(int)
	return value, ok
}

func (P *UpdateNetworkOfferingParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateNetworkOfferingParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateNetworkOfferingParams) SetSortkey(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sortkey"] = v
}

func (P *UpdateNetworkOfferingParams) GetSortkey() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sortkey"].(int)
	return value, ok
}

func (P *UpdateNetworkOfferingParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *UpdateNetworkOfferingParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *UpdateNetworkOfferingParams) SetTags(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *UpdateNetworkOfferingParams) GetTags() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(string)
	return value, ok
}

func (P *UpdateNetworkOfferingParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *UpdateNetworkOfferingParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateNetworkOfferingParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewUpdateNetworkOfferingParams() *UpdateNetworkOfferingParams {
	P := &UpdateNetworkOfferingParams{}
	P.P = make(map[string]interface{})
	return P
}

// Updates a network offering.
func (s *NetworkOfferingService) UpdateNetworkOffering(p *UpdateNetworkOfferingParams) (*UpdateNetworkOfferingResponse, error) {
	resp, err := s.cs.newRequest("updateNetworkOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r UpdateNetworkOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateNetworkOfferingResponse struct {
	Availability             string                                 `json:"availability"`
	Conservemode             bool                                   `json:"conservemode"`
	Created                  string                                 `json:"created"`
	Details                  map[string]string                      `json:"details"`
	Displaytext              string                                 `json:"displaytext"`
	Domain                   string                                 `json:"domain"`
	Domainid                 string                                 `json:"domainid"`
	Egressdefaultpolicy      bool                                   `json:"egressdefaultpolicy"`
	Forvpc                   bool                                   `json:"forvpc"`
	Guestiptype              string                                 `json:"guestiptype"`
	Hasannotations           bool                                   `json:"hasannotations"`
	Id                       string                                 `json:"id"`
	Isdefault                bool                                   `json:"isdefault"`
	Ispersistent             bool                                   `json:"ispersistent"`
	JobID                    string                                 `json:"jobid"`
	Jobstatus                int                                    `json:"jobstatus"`
	Maxconnections           int                                    `json:"maxconnections"`
	Name                     string                                 `json:"name"`
	Networkrate              int                                    `json:"networkrate"`
	Service                  []UpdateNetworkOfferingResponseService `json:"service"`
	Serviceofferingid        string                                 `json:"serviceofferingid"`
	Specifyipranges          bool                                   `json:"specifyipranges"`
	Specifyvlan              bool                                   `json:"specifyvlan"`
	State                    string                                 `json:"state"`
	Supportspublicaccess     bool                                   `json:"supportspublicaccess"`
	Supportsstrechedl2subnet bool                                   `json:"supportsstrechedl2subnet"`
	Tags                     string                                 `json:"tags"`
	Traffictype              string                                 `json:"traffictype"`
	Zone                     string                                 `json:"zone"`
	Zoneid                   string                                 `json:"zoneid"`
}

type UpdateNetworkOfferingResponseService struct {
	Capability []UpdateNetworkOfferingResponseServiceCapability `json:"capability"`
	Name       string                                           `json:"name"`
	Provider   []UpdateNetworkOfferingResponseServiceProvider   `json:"provider"`
}

type UpdateNetworkOfferingResponseServiceProvider struct {
	Canenableindividualservice   bool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid"`
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Physicalnetworkid            string   `json:"physicalnetworkid"`
	Servicelist                  []string `json:"servicelist"`
	State                        string   `json:"state"`
}

type UpdateNetworkOfferingResponseServiceCapability struct {
	Canchooseservicecapability bool   `json:"canchooseservicecapability"`
	Name                       string `json:"name"`
	Value                      string `json:"value"`
}
