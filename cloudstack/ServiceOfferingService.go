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

type ServiceOfferingServiceIface interface {
	CreateServiceOffering(P *CreateServiceOfferingParams) (*CreateServiceOfferingResponse, error)
	NewCreateServiceOfferingParams(displaytext string, name string) *CreateServiceOfferingParams
	DeleteServiceOffering(P *DeleteServiceOfferingParams) (*DeleteServiceOfferingResponse, error)
	NewDeleteServiceOfferingParams(id string) *DeleteServiceOfferingParams
	ListServiceOfferings(P *ListServiceOfferingsParams) (*ListServiceOfferingsResponse, error)
	NewListServiceOfferingsParams() *ListServiceOfferingsParams
	GetServiceOfferingID(name string, opts ...OptionFunc) (string, int, error)
	GetServiceOfferingByName(name string, opts ...OptionFunc) (*ServiceOffering, int, error)
	GetServiceOfferingByID(id string, opts ...OptionFunc) (*ServiceOffering, int, error)
	UpdateServiceOffering(P *UpdateServiceOfferingParams) (*UpdateServiceOfferingResponse, error)
	NewUpdateServiceOfferingParams(id string) *UpdateServiceOfferingParams
}

type CreateServiceOfferingParams struct {
	P map[string]interface{}
}

func (P *CreateServiceOfferingParams) toURLValues() url.Values {
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
	if v, found := P.P["cpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cpunumber", vv)
	}
	if v, found := P.P["cpuspeed"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cpuspeed", vv)
	}
	if v, found := P.P["customized"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customized", vv)
	}
	if v, found := P.P["customizediops"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("customizediops", vv)
	}
	if v, found := P.P["deploymentplanner"]; found {
		u.Set("deploymentplanner", v.(string))
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("domainid", vv)
	}
	if v, found := P.P["dynamicscalingenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("dynamicscalingenabled", vv)
	}
	if v, found := P.P["hosttags"]; found {
		u.Set("hosttags", v.(string))
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
	if v, found := P.P["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
	}
	if v, found := P.P["isvolatile"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isvolatile", vv)
	}
	if v, found := P.P["limitcpuuse"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("limitcpuuse", vv)
	}
	if v, found := P.P["maxcpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxcpunumber", vv)
	}
	if v, found := P.P["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := P.P["maxmemory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("maxmemory", vv)
	}
	if v, found := P.P["memory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("memory", vv)
	}
	if v, found := P.P["mincpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("mincpunumber", vv)
	}
	if v, found := P.P["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := P.P["minmemory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minmemory", vv)
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["networkrate"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("networkrate", vv)
	}
	if v, found := P.P["offerha"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("offerha", vv)
	}
	if v, found := P.P["provisioningtype"]; found {
		u.Set("provisioningtype", v.(string))
	}
	if v, found := P.P["rootdisksize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("rootdisksize", vv)
	}
	if v, found := P.P["serviceofferingdetails"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("serviceofferingdetails[%d].key", i), k)
			u.Set(fmt.Sprintf("serviceofferingdetails[%d].value", i), m[k])
		}
	}
	if v, found := P.P["storagepolicy"]; found {
		u.Set("storagepolicy", v.(string))
	}
	if v, found := P.P["storagetype"]; found {
		u.Set("storagetype", v.(string))
	}
	if v, found := P.P["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
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

func (P *CreateServiceOfferingParams) SetBytesreadrate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadrate"] = v
}

func (P *CreateServiceOfferingParams) GetBytesreadrate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadrate"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetBytesreadratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadratemax"] = v
}

func (P *CreateServiceOfferingParams) GetBytesreadratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadratemax"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetBytesreadratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadratemaxlength"] = v
}

func (P *CreateServiceOfferingParams) GetBytesreadratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadratemaxlength"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetByteswriterate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriterate"] = v
}

func (P *CreateServiceOfferingParams) GetByteswriterate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriterate"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetByteswriteratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriteratemax"] = v
}

func (P *CreateServiceOfferingParams) GetByteswriteratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriteratemax"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetByteswriteratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriteratemaxlength"] = v
}

func (P *CreateServiceOfferingParams) GetByteswriteratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriteratemaxlength"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetCachemode(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cachemode"] = v
}

func (P *CreateServiceOfferingParams) GetCachemode() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cachemode"].(string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetCpunumber(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cpunumber"] = v
}

func (P *CreateServiceOfferingParams) GetCpunumber() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cpunumber"].(int)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetCpuspeed(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cpuspeed"] = v
}

func (P *CreateServiceOfferingParams) GetCpuspeed() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cpuspeed"].(int)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetCustomized(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customized"] = v
}

func (P *CreateServiceOfferingParams) GetCustomized() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customized"].(bool)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetCustomizediops(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customizediops"] = v
}

func (P *CreateServiceOfferingParams) GetCustomizediops() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customizediops"].(bool)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetDeploymentplanner(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["deploymentplanner"] = v
}

func (P *CreateServiceOfferingParams) GetDeploymentplanner() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["deploymentplanner"].(string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *CreateServiceOfferingParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetDomainid(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateServiceOfferingParams) GetDomainid() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].([]string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetDynamicscalingenabled(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["dynamicscalingenabled"] = v
}

func (P *CreateServiceOfferingParams) GetDynamicscalingenabled() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["dynamicscalingenabled"].(bool)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetHosttags(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hosttags"] = v
}

func (P *CreateServiceOfferingParams) GetHosttags() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hosttags"].(string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetHypervisorsnapshotreserve(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hypervisorsnapshotreserve"] = v
}

func (P *CreateServiceOfferingParams) GetHypervisorsnapshotreserve() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hypervisorsnapshotreserve"].(int)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetIopsreadrate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadrate"] = v
}

func (P *CreateServiceOfferingParams) GetIopsreadrate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadrate"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetIopsreadratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadratemax"] = v
}

func (P *CreateServiceOfferingParams) GetIopsreadratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadratemax"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetIopsreadratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadratemaxlength"] = v
}

func (P *CreateServiceOfferingParams) GetIopsreadratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadratemaxlength"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetIopswriterate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriterate"] = v
}

func (P *CreateServiceOfferingParams) GetIopswriterate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriterate"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetIopswriteratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriteratemax"] = v
}

func (P *CreateServiceOfferingParams) GetIopswriteratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriteratemax"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetIopswriteratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriteratemaxlength"] = v
}

func (P *CreateServiceOfferingParams) GetIopswriteratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriteratemaxlength"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetIssystem(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["issystem"] = v
}

func (P *CreateServiceOfferingParams) GetIssystem() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["issystem"].(bool)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetIsvolatile(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isvolatile"] = v
}

func (P *CreateServiceOfferingParams) GetIsvolatile() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isvolatile"].(bool)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetLimitcpuuse(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["limitcpuuse"] = v
}

func (P *CreateServiceOfferingParams) GetLimitcpuuse() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["limitcpuuse"].(bool)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetMaxcpunumber(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxcpunumber"] = v
}

func (P *CreateServiceOfferingParams) GetMaxcpunumber() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxcpunumber"].(int)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetMaxiops(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxiops"] = v
}

func (P *CreateServiceOfferingParams) GetMaxiops() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxiops"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetMaxmemory(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxmemory"] = v
}

func (P *CreateServiceOfferingParams) GetMaxmemory() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxmemory"].(int)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetMemory(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["memory"] = v
}

func (P *CreateServiceOfferingParams) GetMemory() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["memory"].(int)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetMincpunumber(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["mincpunumber"] = v
}

func (P *CreateServiceOfferingParams) GetMincpunumber() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["mincpunumber"].(int)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetMiniops(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["miniops"] = v
}

func (P *CreateServiceOfferingParams) GetMiniops() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["miniops"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetMinmemory(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["minmemory"] = v
}

func (P *CreateServiceOfferingParams) GetMinmemory() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["minmemory"].(int)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateServiceOfferingParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetNetworkrate(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkrate"] = v
}

func (P *CreateServiceOfferingParams) GetNetworkrate() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkrate"].(int)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetOfferha(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["offerha"] = v
}

func (P *CreateServiceOfferingParams) GetOfferha() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["offerha"].(bool)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetProvisioningtype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["provisioningtype"] = v
}

func (P *CreateServiceOfferingParams) GetProvisioningtype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["provisioningtype"].(string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetRootdisksize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["rootdisksize"] = v
}

func (P *CreateServiceOfferingParams) GetRootdisksize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["rootdisksize"].(int64)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetServiceofferingdetails(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["serviceofferingdetails"] = v
}

func (P *CreateServiceOfferingParams) GetServiceofferingdetails() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["serviceofferingdetails"].(map[string]string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetStoragepolicy(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storagepolicy"] = v
}

func (P *CreateServiceOfferingParams) GetStoragepolicy() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storagepolicy"].(string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetStoragetype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storagetype"] = v
}

func (P *CreateServiceOfferingParams) GetStoragetype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storagetype"].(string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetSystemvmtype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["systemvmtype"] = v
}

func (P *CreateServiceOfferingParams) GetSystemvmtype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["systemvmtype"].(string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetTags(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *CreateServiceOfferingParams) GetTags() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(string)
	return value, ok
}

func (P *CreateServiceOfferingParams) SetZoneid(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateServiceOfferingParams) GetZoneid() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].([]string)
	return value, ok
}

// You should always use this function to get a new CreateServiceOfferingParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewCreateServiceOfferingParams(displaytext string, name string) *CreateServiceOfferingParams {
	P := &CreateServiceOfferingParams{}
	P.P = make(map[string]interface{})
	P.P["displaytext"] = displaytext
	P.P["name"] = name
	return P
}

// Creates a service offering.
func (s *ServiceOfferingService) CreateServiceOffering(p *CreateServiceOfferingParams) (*CreateServiceOfferingResponse, error) {
	resp, err := s.cs.newRequest("createServiceOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateServiceOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateServiceOfferingResponse struct {
	CacheMode                   string            `json:"cacheMode"`
	Cpunumber                   int               `json:"cpunumber"`
	Cpuspeed                    int               `json:"cpuspeed"`
	Created                     string            `json:"created"`
	Defaultuse                  bool              `json:"defaultuse"`
	Deploymentplanner           string            `json:"deploymentplanner"`
	DiskBytesReadRate           int64             `json:"diskBytesReadRate"`
	DiskBytesReadRateMax        int64             `json:"diskBytesReadRateMax"`
	DiskBytesReadRateMaxLength  int64             `json:"diskBytesReadRateMaxLength"`
	DiskBytesWriteRate          int64             `json:"diskBytesWriteRate"`
	DiskBytesWriteRateMax       int64             `json:"diskBytesWriteRateMax"`
	DiskBytesWriteRateMaxLength int64             `json:"diskBytesWriteRateMaxLength"`
	DiskIopsReadRate            int64             `json:"diskIopsReadRate"`
	DiskIopsReadRateMax         int64             `json:"diskIopsReadRateMax"`
	DiskIopsReadRateMaxLength   int64             `json:"diskIopsReadRateMaxLength"`
	DiskIopsWriteRate           int64             `json:"diskIopsWriteRate"`
	DiskIopsWriteRateMax        int64             `json:"diskIopsWriteRateMax"`
	DiskIopsWriteRateMaxLength  int64             `json:"diskIopsWriteRateMaxLength"`
	Displaytext                 string            `json:"displaytext"`
	Domain                      string            `json:"domain"`
	Domainid                    string            `json:"domainid"`
	Dynamicscalingenabled       bool              `json:"dynamicscalingenabled"`
	Hasannotations              bool              `json:"hasannotations"`
	Hosttags                    string            `json:"hosttags"`
	Hypervisorsnapshotreserve   int               `json:"hypervisorsnapshotreserve"`
	Id                          string            `json:"id"`
	Iscustomized                bool              `json:"iscustomized"`
	Iscustomizediops            bool              `json:"iscustomizediops"`
	Issystem                    bool              `json:"issystem"`
	Isvolatile                  bool              `json:"isvolatile"`
	JobID                       string            `json:"jobid"`
	Jobstatus                   int               `json:"jobstatus"`
	Limitcpuuse                 bool              `json:"limitcpuuse"`
	Maxiops                     int64             `json:"maxiops"`
	Memory                      int               `json:"memory"`
	Miniops                     int64             `json:"miniops"`
	Name                        string            `json:"name"`
	Networkrate                 int               `json:"networkrate"`
	Offerha                     bool              `json:"offerha"`
	Provisioningtype            string            `json:"provisioningtype"`
	Rootdisksize                int64             `json:"rootdisksize"`
	Serviceofferingdetails      map[string]string `json:"serviceofferingdetails"`
	Storagetags                 string            `json:"storagetags"`
	Storagetype                 string            `json:"storagetype"`
	Systemvmtype                string            `json:"systemvmtype"`
	Vspherestoragepolicy        string            `json:"vspherestoragepolicy"`
	Zone                        string            `json:"zone"`
	Zoneid                      string            `json:"zoneid"`
}

type DeleteServiceOfferingParams struct {
	P map[string]interface{}
}

func (P *DeleteServiceOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteServiceOfferingParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteServiceOfferingParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteServiceOfferingParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewDeleteServiceOfferingParams(id string) *DeleteServiceOfferingParams {
	P := &DeleteServiceOfferingParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a service offering.
func (s *ServiceOfferingService) DeleteServiceOffering(p *DeleteServiceOfferingParams) (*DeleteServiceOfferingResponse, error) {
	resp, err := s.cs.newRequest("deleteServiceOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteServiceOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteServiceOfferingResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteServiceOfferingResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteServiceOfferingResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListServiceOfferingsParams struct {
	P map[string]interface{}
}

func (P *ListServiceOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["cpunumber"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cpunumber", vv)
	}
	if v, found := P.P["cpuspeed"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("cpuspeed", vv)
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
	if v, found := P.P["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := P.P["memory"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("memory", vv)
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
	if v, found := P.P["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListServiceOfferingsParams) SetCpunumber(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cpunumber"] = v
}

func (P *ListServiceOfferingsParams) GetCpunumber() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cpunumber"].(int)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetCpuspeed(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cpuspeed"] = v
}

func (P *ListServiceOfferingsParams) GetCpuspeed() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cpuspeed"].(int)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListServiceOfferingsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListServiceOfferingsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListServiceOfferingsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetIssystem(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["issystem"] = v
}

func (P *ListServiceOfferingsParams) GetIssystem() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["issystem"].(bool)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListServiceOfferingsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListServiceOfferingsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetMemory(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["memory"] = v
}

func (P *ListServiceOfferingsParams) GetMemory() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["memory"].(int)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListServiceOfferingsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListServiceOfferingsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListServiceOfferingsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetSystemvmtype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["systemvmtype"] = v
}

func (P *ListServiceOfferingsParams) GetSystemvmtype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["systemvmtype"].(string)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *ListServiceOfferingsParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *ListServiceOfferingsParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListServiceOfferingsParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListServiceOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewListServiceOfferingsParams() *ListServiceOfferingsParams {
	P := &ListServiceOfferingsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ServiceOfferingService) GetServiceOfferingID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListServiceOfferingsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListServiceOfferings(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ServiceOfferings[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ServiceOfferings {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ServiceOfferingService) GetServiceOfferingByName(name string, opts ...OptionFunc) (*ServiceOffering, int, error) {
	id, count, err := s.GetServiceOfferingID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetServiceOfferingByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ServiceOfferingService) GetServiceOfferingByID(id string, opts ...OptionFunc) (*ServiceOffering, int, error) {
	P := &ListServiceOfferingsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListServiceOfferings(P)
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
		return l.ServiceOfferings[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ServiceOffering UUID: %s!", id)
}

// Lists all available service offerings.
func (s *ServiceOfferingService) ListServiceOfferings(p *ListServiceOfferingsParams) (*ListServiceOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listServiceOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListServiceOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListServiceOfferingsResponse struct {
	Count            int                `json:"count"`
	ServiceOfferings []*ServiceOffering `json:"serviceoffering"`
}

type ServiceOffering struct {
	CacheMode                   string            `json:"cacheMode"`
	Cpunumber                   int               `json:"cpunumber"`
	Cpuspeed                    int               `json:"cpuspeed"`
	Created                     string            `json:"created"`
	Defaultuse                  bool              `json:"defaultuse"`
	Deploymentplanner           string            `json:"deploymentplanner"`
	DiskBytesReadRate           int64             `json:"diskBytesReadRate"`
	DiskBytesReadRateMax        int64             `json:"diskBytesReadRateMax"`
	DiskBytesReadRateMaxLength  int64             `json:"diskBytesReadRateMaxLength"`
	DiskBytesWriteRate          int64             `json:"diskBytesWriteRate"`
	DiskBytesWriteRateMax       int64             `json:"diskBytesWriteRateMax"`
	DiskBytesWriteRateMaxLength int64             `json:"diskBytesWriteRateMaxLength"`
	DiskIopsReadRate            int64             `json:"diskIopsReadRate"`
	DiskIopsReadRateMax         int64             `json:"diskIopsReadRateMax"`
	DiskIopsReadRateMaxLength   int64             `json:"diskIopsReadRateMaxLength"`
	DiskIopsWriteRate           int64             `json:"diskIopsWriteRate"`
	DiskIopsWriteRateMax        int64             `json:"diskIopsWriteRateMax"`
	DiskIopsWriteRateMaxLength  int64             `json:"diskIopsWriteRateMaxLength"`
	Displaytext                 string            `json:"displaytext"`
	Domain                      string            `json:"domain"`
	Domainid                    string            `json:"domainid"`
	Dynamicscalingenabled       bool              `json:"dynamicscalingenabled"`
	Hasannotations              bool              `json:"hasannotations"`
	Hosttags                    string            `json:"hosttags"`
	Hypervisorsnapshotreserve   int               `json:"hypervisorsnapshotreserve"`
	Id                          string            `json:"id"`
	Iscustomized                bool              `json:"iscustomized"`
	Iscustomizediops            bool              `json:"iscustomizediops"`
	Issystem                    bool              `json:"issystem"`
	Isvolatile                  bool              `json:"isvolatile"`
	JobID                       string            `json:"jobid"`
	Jobstatus                   int               `json:"jobstatus"`
	Limitcpuuse                 bool              `json:"limitcpuuse"`
	Maxiops                     int64             `json:"maxiops"`
	Memory                      int               `json:"memory"`
	Miniops                     int64             `json:"miniops"`
	Name                        string            `json:"name"`
	Networkrate                 int               `json:"networkrate"`
	Offerha                     bool              `json:"offerha"`
	Provisioningtype            string            `json:"provisioningtype"`
	Rootdisksize                int64             `json:"rootdisksize"`
	Serviceofferingdetails      map[string]string `json:"serviceofferingdetails"`
	Storagetags                 string            `json:"storagetags"`
	Storagetype                 string            `json:"storagetype"`
	Systemvmtype                string            `json:"systemvmtype"`
	Vspherestoragepolicy        string            `json:"vspherestoragepolicy"`
	Zone                        string            `json:"zone"`
	Zoneid                      string            `json:"zoneid"`
}

type UpdateServiceOfferingParams struct {
	P map[string]interface{}
}

func (P *UpdateServiceOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["hosttags"]; found {
		u.Set("hosttags", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["sortkey"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sortkey", vv)
	}
	if v, found := P.P["storagetags"]; found {
		u.Set("storagetags", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
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
	return u
}

func (P *UpdateServiceOfferingParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *UpdateServiceOfferingParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UpdateServiceOfferingParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetHosttags(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hosttags"] = v
}

func (P *UpdateServiceOfferingParams) GetHosttags() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hosttags"].(string)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateServiceOfferingParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateServiceOfferingParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetSortkey(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["sortkey"] = v
}

func (P *UpdateServiceOfferingParams) GetSortkey() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["sortkey"].(int)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetStoragetags(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storagetags"] = v
}

func (P *UpdateServiceOfferingParams) GetStoragetags() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storagetags"].(string)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *UpdateServiceOfferingParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetBytesreadrate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadrate"] = v
}

func (P *UpdateServiceOfferingParams) GetBytesreadrate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadrate"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetBytesreadratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadratemax"] = v
}

func (P *UpdateServiceOfferingParams) GetBytesreadratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadratemax"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetBytesreadratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["bytesreadratemaxlength"] = v
}

func (P *UpdateServiceOfferingParams) GetBytesreadratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["bytesreadratemaxlength"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetByteswriterate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriterate"] = v
}

func (P *UpdateServiceOfferingParams) GetByteswriterate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriterate"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetByteswriteratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriteratemax"] = v
}

func (P *UpdateServiceOfferingParams) GetByteswriteratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriteratemax"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetByteswriteratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["byteswriteratemaxlength"] = v
}

func (P *UpdateServiceOfferingParams) GetByteswriteratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["byteswriteratemaxlength"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetIopsreadrate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadrate"] = v
}

func (P *UpdateServiceOfferingParams) GetIopsreadrate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadrate"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetIopsreadratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadratemax"] = v
}

func (P *UpdateServiceOfferingParams) GetIopsreadratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadratemax"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetIopsreadratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopsreadratemaxlength"] = v
}

func (P *UpdateServiceOfferingParams) GetIopsreadratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopsreadratemaxlength"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetIopswriterate(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriterate"] = v
}

func (P *UpdateServiceOfferingParams) GetIopswriterate() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriterate"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetIopswriteratemax(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriteratemax"] = v
}

func (P *UpdateServiceOfferingParams) GetIopswriteratemax() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriteratemax"].(int64)
	return value, ok
}

func (P *UpdateServiceOfferingParams) SetIopswriteratemaxlength(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["iopswriteratemaxlength"] = v
}

func (P *UpdateServiceOfferingParams) GetIopswriteratemaxlength() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["iopswriteratemaxlength"].(int64)
	return value, ok
}

// You should always use this function to get a new UpdateServiceOfferingParams instance,
// as then you are sure you have configured all required params
func (s *ServiceOfferingService) NewUpdateServiceOfferingParams(id string) *UpdateServiceOfferingParams {
	P := &UpdateServiceOfferingParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a service offering.
func (s *ServiceOfferingService) UpdateServiceOffering(p *UpdateServiceOfferingParams) (*UpdateServiceOfferingResponse, error) {
	resp, err := s.cs.newRequest("updateServiceOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r UpdateServiceOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateServiceOfferingResponse struct {
	CacheMode                   string            `json:"cacheMode"`
	Cpunumber                   int               `json:"cpunumber"`
	Cpuspeed                    int               `json:"cpuspeed"`
	Created                     string            `json:"created"`
	Defaultuse                  bool              `json:"defaultuse"`
	Deploymentplanner           string            `json:"deploymentplanner"`
	DiskBytesReadRate           int64             `json:"diskBytesReadRate"`
	DiskBytesReadRateMax        int64             `json:"diskBytesReadRateMax"`
	DiskBytesReadRateMaxLength  int64             `json:"diskBytesReadRateMaxLength"`
	DiskBytesWriteRate          int64             `json:"diskBytesWriteRate"`
	DiskBytesWriteRateMax       int64             `json:"diskBytesWriteRateMax"`
	DiskBytesWriteRateMaxLength int64             `json:"diskBytesWriteRateMaxLength"`
	DiskIopsReadRate            int64             `json:"diskIopsReadRate"`
	DiskIopsReadRateMax         int64             `json:"diskIopsReadRateMax"`
	DiskIopsReadRateMaxLength   int64             `json:"diskIopsReadRateMaxLength"`
	DiskIopsWriteRate           int64             `json:"diskIopsWriteRate"`
	DiskIopsWriteRateMax        int64             `json:"diskIopsWriteRateMax"`
	DiskIopsWriteRateMaxLength  int64             `json:"diskIopsWriteRateMaxLength"`
	Displaytext                 string            `json:"displaytext"`
	Domain                      string            `json:"domain"`
	Domainid                    string            `json:"domainid"`
	Dynamicscalingenabled       bool              `json:"dynamicscalingenabled"`
	Hasannotations              bool              `json:"hasannotations"`
	Hosttags                    string            `json:"hosttags"`
	Hypervisorsnapshotreserve   int               `json:"hypervisorsnapshotreserve"`
	Id                          string            `json:"id"`
	Iscustomized                bool              `json:"iscustomized"`
	Iscustomizediops            bool              `json:"iscustomizediops"`
	Issystem                    bool              `json:"issystem"`
	Isvolatile                  bool              `json:"isvolatile"`
	JobID                       string            `json:"jobid"`
	Jobstatus                   int               `json:"jobstatus"`
	Limitcpuuse                 bool              `json:"limitcpuuse"`
	Maxiops                     int64             `json:"maxiops"`
	Memory                      int               `json:"memory"`
	Miniops                     int64             `json:"miniops"`
	Name                        string            `json:"name"`
	Networkrate                 int               `json:"networkrate"`
	Offerha                     bool              `json:"offerha"`
	Provisioningtype            string            `json:"provisioningtype"`
	Rootdisksize                int64             `json:"rootdisksize"`
	Serviceofferingdetails      map[string]string `json:"serviceofferingdetails"`
	Storagetags                 string            `json:"storagetags"`
	Storagetype                 string            `json:"storagetype"`
	Systemvmtype                string            `json:"systemvmtype"`
	Vspherestoragepolicy        string            `json:"vspherestoragepolicy"`
	Zone                        string            `json:"zone"`
	Zoneid                      string            `json:"zoneid"`
}
