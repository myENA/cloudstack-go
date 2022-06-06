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

type DiskOfferingServiceIface interface {
	CreateDiskOffering(P *CreateDiskOfferingParams) (*CreateDiskOfferingResponse, error)
	NewCreateDiskOfferingParams(displaytext string, name string) *CreateDiskOfferingParams
	DeleteDiskOffering(P *DeleteDiskOfferingParams) (*DeleteDiskOfferingResponse, error)
	NewDeleteDiskOfferingParams(id string) *DeleteDiskOfferingParams
	ListDiskOfferings(P *ListDiskOfferingsParams) (*ListDiskOfferingsResponse, error)
	NewListDiskOfferingsParams() *ListDiskOfferingsParams
	GetDiskOfferingID(name string, opts ...OptionFunc) (string, int, error)
	GetDiskOfferingByName(name string, opts ...OptionFunc) (*DiskOffering, int, error)
	GetDiskOfferingByID(id string, opts ...OptionFunc) (*DiskOffering, int, error)
	UpdateDiskOffering(P *UpdateDiskOfferingParams) (*UpdateDiskOfferingResponse, error)
	NewUpdateDiskOfferingParams(id string) *UpdateDiskOfferingParams
}

type CreateDiskOfferingParams struct {
	P map[string]interface{}
}

func (P *CreateDiskOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["bytesreadrate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("bytesreadrate", vv)
	}
	if v, found := P.P["bytesreadratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("bytesreadratemax", vv)
	}
	if v, found := P.P["bytesreadratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("bytesreadratemaxlength", vv)
	}
	if v, found := P.P["byteswriterate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("byteswriterate", vv)
	}
	if v, found := P.P["byteswriteratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("byteswriteratemax", vv)
	}
	if v, found := P.P["byteswriteratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("byteswriteratemaxlength", vv)
	}
	if v, found := P.P["cachemode"]; found {
		u.Set("cachemode", v.(string))
	}
	if v, found := P.P["customized"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customized", vv)
	}
	if v, found := P.P["customizediops"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customizediops", vv)
	}
	if v, found := P.P["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := P.P["disksize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("disksize", vv)
	}
	if v, found := P.P["displayoffering"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayoffering", vv)
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("domainid", vv)
	}
	if v, found := P.P["hypervisorsnapshotreserve"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("hypervisorsnapshotreserve", vv)
	}
	if v, found := P.P["iopsreadrate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopsreadrate", vv)
	}
	if v, found := P.P["iopsreadratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopsreadratemax", vv)
	}
	if v, found := P.P["iopsreadratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopsreadratemaxlength", vv)
	}
	if v, found := P.P["iopswriterate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopswriterate", vv)
	}
	if v, found := P.P["iopswriteratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopswriteratemax", vv)
	}
	if v, found := P.P["iopswriteratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopswriteratemaxlength", vv)
	}
	if v, found := P.P["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := P.P["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["provisioningtype"]; found {
		u.Set("provisioningtype", v.(string))
	}
	if v, found := P.P["storagepolicy"]; found {
		u.Set("storagepolicy", v.(string))
	}
	if v, found := P.P["storagetype"]; found {
		u.Set("storagetype", v.(string))
	}
	if v, found := P.P["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("zoneid", vv)
	}
	return u
}

func (P *CreateDiskOfferingParams) SetBytesreadrate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadrate"] = v
}

func (P *CreateDiskOfferingParams) GetBytesreadrate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadrate"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetBytesreadratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadratemax"] = v
}

func (P *CreateDiskOfferingParams) GetBytesreadratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadratemax"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetBytesreadratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadratemaxlength"] = v
}

func (P *CreateDiskOfferingParams) GetBytesreadratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadratemaxlength"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetByteswriterate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriterate"] = v
}

func (P *CreateDiskOfferingParams) GetByteswriterate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriterate"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetByteswriteratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriteratemax"] = v
}

func (P *CreateDiskOfferingParams) GetByteswriteratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriteratemax"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetByteswriteratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriteratemaxlength"] = v
}

func (P *CreateDiskOfferingParams) GetByteswriteratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriteratemaxlength"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetCachemode(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cachemode"] = v
}

func (P *CreateDiskOfferingParams) GetCachemode() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cachemode"].(string)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetCustomized(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customized"] = v
}

func (P *CreateDiskOfferingParams) GetCustomized() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customized"].(bool)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetCustomizediops(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customizediops"] = v
}

func (P *CreateDiskOfferingParams) GetCustomizediops() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customizediops"].(bool)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetDetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *CreateDiskOfferingParams) GetDetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].(map[string]string)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetDisksize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["disksize"] = v
}

func (P *CreateDiskOfferingParams) GetDisksize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["disksize"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetDisplayoffering(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayoffering"] = v
}

func (P *CreateDiskOfferingParams) GetDisplayoffering() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayoffering"].(bool)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *CreateDiskOfferingParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetDomainid(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateDiskOfferingParams) GetDomainid() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].([]string)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetHypervisorsnapshotreserve(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisorsnapshotreserve"] = v
}

func (P *CreateDiskOfferingParams) GetHypervisorsnapshotreserve() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisorsnapshotreserve"].(int)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetIopsreadrate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadrate"] = v
}

func (P *CreateDiskOfferingParams) GetIopsreadrate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadrate"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetIopsreadratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadratemax"] = v
}

func (P *CreateDiskOfferingParams) GetIopsreadratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadratemax"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetIopsreadratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadratemaxlength"] = v
}

func (P *CreateDiskOfferingParams) GetIopsreadratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadratemaxlength"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetIopswriterate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriterate"] = v
}

func (P *CreateDiskOfferingParams) GetIopswriterate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriterate"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetIopswriteratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriteratemax"] = v
}

func (P *CreateDiskOfferingParams) GetIopswriteratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriteratemax"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetIopswriteratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriteratemaxlength"] = v
}

func (P *CreateDiskOfferingParams) GetIopswriteratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriteratemaxlength"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetMaxiops(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxiops"] = v
}

func (P *CreateDiskOfferingParams) GetMaxiops() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxiops"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetMiniops(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["miniops"] = v
}

func (P *CreateDiskOfferingParams) GetMiniops() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["miniops"].(int64)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateDiskOfferingParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetProvisioningtype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["provisioningtype"] = v
}

func (P *CreateDiskOfferingParams) GetProvisioningtype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["provisioningtype"].(string)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetStoragepolicy(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storagepolicy"] = v
}

func (P *CreateDiskOfferingParams) GetStoragepolicy() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storagepolicy"].(string)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetStoragetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storagetype"] = v
}

func (P *CreateDiskOfferingParams) GetStoragetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storagetype"].(string)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetTags(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *CreateDiskOfferingParams) GetTags() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(string)
	return value, ok
}

func (P *CreateDiskOfferingParams) SetZoneid(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateDiskOfferingParams) GetZoneid() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].([]string)
	return value, ok
}

// You should always use this function to get a new CreateDiskOfferingParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewCreateDiskOfferingParams(displaytext string, name string) *CreateDiskOfferingParams {
	P := &CreateDiskOfferingParams{}
	P.P = make(map[string]interface{})
	P.P["displaytext"] = displaytext
	P.P["name"] = name
	return P
}

// Creates a disk offering.
func (s *DiskOfferingService) CreateDiskOffering(p *CreateDiskOfferingParams) (*CreateDiskOfferingResponse, error) {
	resp, err := s.cs.newRequest("createDiskOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateDiskOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateDiskOfferingResponse struct {
	CacheMode                   string `json:"cacheMode"`
	Created                     string `json:"created"`
	DiskBytesReadRate           int64  `json:"diskBytesReadRate"`
	DiskBytesReadRateMax        int64  `json:"diskBytesReadRateMax"`
	DiskBytesReadRateMaxLength  int64  `json:"diskBytesReadRateMaxLength"`
	DiskBytesWriteRate          int64  `json:"diskBytesWriteRate"`
	DiskBytesWriteRateMax       int64  `json:"diskBytesWriteRateMax"`
	DiskBytesWriteRateMaxLength int64  `json:"diskBytesWriteRateMaxLength"`
	DiskIopsReadRate            int64  `json:"diskIopsReadRate"`
	DiskIopsReadRateMax         int64  `json:"diskIopsReadRateMax"`
	DiskIopsReadRateMaxLength   int64  `json:"diskIopsReadRateMaxLength"`
	DiskIopsWriteRate           int64  `json:"diskIopsWriteRate"`
	DiskIopsWriteRateMax        int64  `json:"diskIopsWriteRateMax"`
	DiskIopsWriteRateMaxLength  int64  `json:"diskIopsWriteRateMaxLength"`
	Disksize                    int64  `json:"disksize"`
	Displayoffering             bool   `json:"displayoffering"`
	Displaytext                 string `json:"displaytext"`
	Domain                      string `json:"domain"`
	Domainid                    string `json:"domainid"`
	Hasannotations              bool   `json:"hasannotations"`
	Hypervisorsnapshotreserve   int    `json:"hypervisorsnapshotreserve"`
	Id                          string `json:"id"`
	Iscustomized                bool   `json:"iscustomized"`
	Iscustomizediops            bool   `json:"iscustomizediops"`
	JobID                       string `json:"jobid"`
	Jobstatus                   int    `json:"jobstatus"`
	Maxiops                     int64  `json:"maxiops"`
	Miniops                     int64  `json:"miniops"`
	Name                        string `json:"name"`
	Provisioningtype            string `json:"provisioningtype"`
	Storagetype                 string `json:"storagetype"`
	Tags                        string `json:"tags"`
	Vspherestoragepolicy        string `json:"vspherestoragepolicy"`
	Zone                        string `json:"zone"`
	Zoneid                      string `json:"zoneid"`
}

type DeleteDiskOfferingParams struct {
	P map[string]interface{}
}

func (P *DeleteDiskOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteDiskOfferingParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteDiskOfferingParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteDiskOfferingParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewDeleteDiskOfferingParams(id string) *DeleteDiskOfferingParams {
	P := &DeleteDiskOfferingParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a disk offering.
func (s *DiskOfferingService) DeleteDiskOffering(p *DeleteDiskOfferingParams) (*DeleteDiskOfferingResponse, error) {
	resp, err := s.cs.newRequest("deleteDiskOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteDiskOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteDiskOfferingResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteDiskOfferingResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteDiskOfferingResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListDiskOfferingsParams struct {
	P map[string]interface{}
}

func (P *ListDiskOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
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
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListDiskOfferingsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListDiskOfferingsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListDiskOfferingsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListDiskOfferingsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListDiskOfferingsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListDiskOfferingsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListDiskOfferingsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListDiskOfferingsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListDiskOfferingsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListDiskOfferingsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListDiskOfferingsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListDiskOfferingsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListDiskOfferingsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListDiskOfferingsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListDiskOfferingsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListDiskOfferingsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListDiskOfferingsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListDiskOfferingsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListDiskOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewListDiskOfferingsParams() *ListDiskOfferingsParams {
	P := &ListDiskOfferingsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DiskOfferingService) GetDiskOfferingID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListDiskOfferingsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListDiskOfferings(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.DiskOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.DiskOfferings {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DiskOfferingService) GetDiskOfferingByName(name string, opts ...OptionFunc) (*DiskOffering, int, error) {
	id, count, err := s.GetDiskOfferingID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetDiskOfferingByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *DiskOfferingService) GetDiskOfferingByID(id string, opts ...OptionFunc) (*DiskOffering, int, error) {
	P := &ListDiskOfferingsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListDiskOfferings(P)
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
		return l.DiskOfferings[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for DiskOffering UUID: %s!", id)
}

// Lists all available disk offerings.
func (s *DiskOfferingService) ListDiskOfferings(p *ListDiskOfferingsParams) (*ListDiskOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listDiskOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDiskOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDiskOfferingsResponse struct {
	Count         int             `json:"count"`
	DiskOfferings []*DiskOffering `json:"diskoffering"`
}

type DiskOffering struct {
	CacheMode                   string `json:"cacheMode"`
	Created                     string `json:"created"`
	DiskBytesReadRate           int64  `json:"diskBytesReadRate"`
	DiskBytesReadRateMax        int64  `json:"diskBytesReadRateMax"`
	DiskBytesReadRateMaxLength  int64  `json:"diskBytesReadRateMaxLength"`
	DiskBytesWriteRate          int64  `json:"diskBytesWriteRate"`
	DiskBytesWriteRateMax       int64  `json:"diskBytesWriteRateMax"`
	DiskBytesWriteRateMaxLength int64  `json:"diskBytesWriteRateMaxLength"`
	DiskIopsReadRate            int64  `json:"diskIopsReadRate"`
	DiskIopsReadRateMax         int64  `json:"diskIopsReadRateMax"`
	DiskIopsReadRateMaxLength   int64  `json:"diskIopsReadRateMaxLength"`
	DiskIopsWriteRate           int64  `json:"diskIopsWriteRate"`
	DiskIopsWriteRateMax        int64  `json:"diskIopsWriteRateMax"`
	DiskIopsWriteRateMaxLength  int64  `json:"diskIopsWriteRateMaxLength"`
	Disksize                    int64  `json:"disksize"`
	Displayoffering             bool   `json:"displayoffering"`
	Displaytext                 string `json:"displaytext"`
	Domain                      string `json:"domain"`
	Domainid                    string `json:"domainid"`
	Hasannotations              bool   `json:"hasannotations"`
	Hypervisorsnapshotreserve   int    `json:"hypervisorsnapshotreserve"`
	Id                          string `json:"id"`
	Iscustomized                bool   `json:"iscustomized"`
	Iscustomizediops            bool   `json:"iscustomizediops"`
	JobID                       string `json:"jobid"`
	Jobstatus                   int    `json:"jobstatus"`
	Maxiops                     int64  `json:"maxiops"`
	Miniops                     int64  `json:"miniops"`
	Name                        string `json:"name"`
	Provisioningtype            string `json:"provisioningtype"`
	Storagetype                 string `json:"storagetype"`
	Tags                        string `json:"tags"`
	Vspherestoragepolicy        string `json:"vspherestoragepolicy"`
	Zone                        string `json:"zone"`
	Zoneid                      string `json:"zoneid"`
}

type UpdateDiskOfferingParams struct {
	P map[string]interface{}
}

func (P *UpdateDiskOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["bytesreadrate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("bytesreadrate", vv)
	}
	if v, found := P.P["bytesreadratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("bytesreadratemax", vv)
	}
	if v, found := P.P["bytesreadratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("bytesreadratemaxlength", vv)
	}
	if v, found := P.P["byteswriterate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("byteswriterate", vv)
	}
	if v, found := P.P["byteswriteratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("byteswriteratemax", vv)
	}
	if v, found := P.P["byteswriteratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("byteswriteratemaxlength", vv)
	}
	if v, found := P.P["cachemode"]; found {
		u.Set("cachemode", v.(string))
	}
	if v, found := P.P["displayoffering"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayoffering", vv)
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
	if v, found := P.P["iopsreadrate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopsreadrate", vv)
	}
	if v, found := P.P["iopsreadratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopsreadratemax", vv)
	}
	if v, found := P.P["iopsreadratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopsreadratemaxlength", vv)
	}
	if v, found := P.P["iopswriterate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopswriterate", vv)
	}
	if v, found := P.P["iopswriteratemax"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopswriteratemax", vv)
	}
	if v, found := P.P["iopswriteratemaxlength"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("iopswriteratemaxlength", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	if v, found := P.P["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *UpdateDiskOfferingParams) SetBytesreadrate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadrate"] = v
}

func (P *UpdateDiskOfferingParams) GetBytesreadrate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadrate"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetBytesreadratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadratemax"] = v
}

func (P *UpdateDiskOfferingParams) GetBytesreadratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadratemax"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetBytesreadratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadratemaxlength"] = v
}

func (P *UpdateDiskOfferingParams) GetBytesreadratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadratemaxlength"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetByteswriterate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriterate"] = v
}

func (P *UpdateDiskOfferingParams) GetByteswriterate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriterate"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetByteswriteratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriteratemax"] = v
}

func (P *UpdateDiskOfferingParams) GetByteswriteratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriteratemax"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetByteswriteratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriteratemaxlength"] = v
}

func (P *UpdateDiskOfferingParams) GetByteswriteratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriteratemaxlength"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetCachemode(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cachemode"] = v
}

func (P *UpdateDiskOfferingParams) GetCachemode() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cachemode"].(string)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetDisplayoffering(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayoffering"] = v
}

func (P *UpdateDiskOfferingParams) GetDisplayoffering() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayoffering"].(bool)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *UpdateDiskOfferingParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UpdateDiskOfferingParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateDiskOfferingParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetIopsreadrate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadrate"] = v
}

func (P *UpdateDiskOfferingParams) GetIopsreadrate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadrate"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetIopsreadratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadratemax"] = v
}

func (P *UpdateDiskOfferingParams) GetIopsreadratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadratemax"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetIopsreadratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadratemaxlength"] = v
}

func (P *UpdateDiskOfferingParams) GetIopsreadratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadratemaxlength"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetIopswriterate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriterate"] = v
}

func (P *UpdateDiskOfferingParams) GetIopswriterate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriterate"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetIopswriteratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriteratemax"] = v
}

func (P *UpdateDiskOfferingParams) GetIopswriteratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriteratemax"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetIopswriteratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriteratemaxlength"] = v
}

func (P *UpdateDiskOfferingParams) GetIopswriteratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriteratemaxlength"].(int64)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateDiskOfferingParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetSortkey(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sortkey"] = v
}

func (P *UpdateDiskOfferingParams) GetSortkey() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sortkey"].(int)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetTags(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *UpdateDiskOfferingParams) GetTags() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(string)
	return value, ok
}

func (P *UpdateDiskOfferingParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *UpdateDiskOfferingParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateDiskOfferingParams instance,
// as then you are sure you have configured all required params
func (s *DiskOfferingService) NewUpdateDiskOfferingParams(id string) *UpdateDiskOfferingParams {
	P := &UpdateDiskOfferingParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a disk offering.
func (s *DiskOfferingService) UpdateDiskOffering(p *UpdateDiskOfferingParams) (*UpdateDiskOfferingResponse, error) {
	resp, err := s.cs.newRequest("updateDiskOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateDiskOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateDiskOfferingResponse struct {
	CacheMode                   string `json:"cacheMode"`
	Created                     string `json:"created"`
	DiskBytesReadRate           int64  `json:"diskBytesReadRate"`
	DiskBytesReadRateMax        int64  `json:"diskBytesReadRateMax"`
	DiskBytesReadRateMaxLength  int64  `json:"diskBytesReadRateMaxLength"`
	DiskBytesWriteRate          int64  `json:"diskBytesWriteRate"`
	DiskBytesWriteRateMax       int64  `json:"diskBytesWriteRateMax"`
	DiskBytesWriteRateMaxLength int64  `json:"diskBytesWriteRateMaxLength"`
	DiskIopsReadRate            int64  `json:"diskIopsReadRate"`
	DiskIopsReadRateMax         int64  `json:"diskIopsReadRateMax"`
	DiskIopsReadRateMaxLength   int64  `json:"diskIopsReadRateMaxLength"`
	DiskIopsWriteRate           int64  `json:"diskIopsWriteRate"`
	DiskIopsWriteRateMax        int64  `json:"diskIopsWriteRateMax"`
	DiskIopsWriteRateMaxLength  int64  `json:"diskIopsWriteRateMaxLength"`
	Disksize                    int64  `json:"disksize"`
	Displayoffering             bool   `json:"displayoffering"`
	Displaytext                 string `json:"displaytext"`
	Domain                      string `json:"domain"`
	Domainid                    string `json:"domainid"`
	Hasannotations              bool   `json:"hasannotations"`
	Hypervisorsnapshotreserve   int    `json:"hypervisorsnapshotreserve"`
	Id                          string `json:"id"`
	Iscustomized                bool   `json:"iscustomized"`
	Iscustomizediops            bool   `json:"iscustomizediops"`
	JobID                       string `json:"jobid"`
	Jobstatus                   int    `json:"jobstatus"`
	Maxiops                     int64  `json:"maxiops"`
	Miniops                     int64  `json:"miniops"`
	Name                        string `json:"name"`
	Provisioningtype            string `json:"provisioningtype"`
	Storagetype                 string `json:"storagetype"`
	Tags                        string `json:"tags"`
	Vspherestoragepolicy        string `json:"vspherestoragepolicy"`
	Zone                        string `json:"zone"`
	Zoneid                      string `json:"zoneid"`
}
