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

type NicServiceIface interface {
	AddIpToNic(p *AddIpToNicParams) (*AddIpToNicResponse, error)
	NewAddIpToNicParams(nicid string) *AddIpToNicParams
	ListNics(p *ListNicsParams) (*ListNicsResponse, error)
	NewListNicsParams(virtualmachineid string) *ListNicsParams
	RemoveIpFromNic(p *RemoveIpFromNicParams) (*RemoveIpFromNicResponse, error)
	NewRemoveIpFromNicParams(id string) *RemoveIpFromNicParams
	UpdateVmNicIp(p *UpdateVmNicIpParams) (*UpdateVmNicIpResponse, error)
	NewUpdateVmNicIpParams(nicid string) *UpdateVmNicIpParams
}

type AddIpToNicParams struct {
	P map[string]interface{}
}

func (P *AddIpToNicParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := P.P["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	return u
}

func (P *AddIpToNicParams) SetIpaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddress"] = v
}

func (P *AddIpToNicParams) GetIpaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddress"].(string)
	return value, ok
}

func (P *AddIpToNicParams) SetNicid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nicid"] = v
}

func (P *AddIpToNicParams) GetNicid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nicid"].(string)
	return value, ok
}

// You should always use this function to get a new AddIpToNicParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewAddIpToNicParams(nicid string) *AddIpToNicParams {
	P := &AddIpToNicParams{}
	P.P = make(map[string]interface{})
	P.P["nicid"] = nicid
	return P
}

// Assigns secondary IP to NIC
func (s *NicService) AddIpToNic(p *AddIpToNicParams) (*AddIpToNicResponse, error) {
	resp, err := s.cs.newRequest("addIpToNic", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddIpToNicResponse
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

type AddIpToNicResponse struct {
	Id          string `json:"id"`
	Ipaddress   string `json:"ipaddress"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Networkid   string `json:"networkid"`
	Nicid       string `json:"nicid"`
	Secondaryip []struct {
		Id        string `json:"id"`
		Ipaddress string `json:"ipaddress"`
	} `json:"secondaryip"`
	Virtualmachineid string `json:"virtualmachineid"`
}

type ListNicsParams struct {
	P map[string]interface{}
}

func (P *ListNicsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := P.P["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := P.P["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	if v, found := P.P["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := P.P["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *ListNicsParams) SetFordisplay(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["fordisplay"] = v
}

func (P *ListNicsParams) GetFordisplay() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["fordisplay"].(bool)
	return value, ok
}

func (P *ListNicsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListNicsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListNicsParams) SetNetworkid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["networkid"] = v
}

func (P *ListNicsParams) GetNetworkid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["networkid"].(string)
	return value, ok
}

func (P *ListNicsParams) SetNicid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nicid"] = v
}

func (P *ListNicsParams) GetNicid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nicid"].(string)
	return value, ok
}

func (P *ListNicsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListNicsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListNicsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListNicsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListNicsParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *ListNicsParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new ListNicsParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewListNicsParams(virtualmachineid string) *ListNicsParams {
	P := &ListNicsParams{}
	P.P = make(map[string]interface{})
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// list the vm nics  IP to NIC
func (s *NicService) ListNics(p *ListNicsParams) (*ListNicsResponse, error) {
	resp, err := s.cs.newRequest("listNics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNicsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNicsResponse struct {
	Count int    `json:"count"`
	Nics  []*Nic `json:"nic"`
}

type Nic struct {
	Adaptertype          string   `json:"adaptertype"`
	Broadcasturi         string   `json:"broadcasturi"`
	Deviceid             string   `json:"deviceid"`
	Extradhcpoption      []string `json:"extradhcpoption"`
	Gateway              string   `json:"gateway"`
	Id                   string   `json:"id"`
	Ip6address           string   `json:"ip6address"`
	Ip6cidr              string   `json:"ip6cidr"`
	Ip6gateway           string   `json:"ip6gateway"`
	Ipaddress            string   `json:"ipaddress"`
	Ipaddresses          []string `json:"ipaddresses"`
	Isdefault            bool     `json:"isdefault"`
	Isolatedpvlan        int      `json:"isolatedpvlan"`
	Isolatedpvlantype    string   `json:"isolatedpvlantype"`
	Isolationuri         string   `json:"isolationuri"`
	JobID                string   `json:"jobid"`
	Jobstatus            int      `json:"jobstatus"`
	Macaddress           string   `json:"macaddress"`
	Netmask              string   `json:"netmask"`
	Networkid            string   `json:"networkid"`
	Networkname          string   `json:"networkname"`
	Nsxlogicalswitch     string   `json:"nsxlogicalswitch"`
	Nsxlogicalswitchport string   `json:"nsxlogicalswitchport"`
	Secondaryip          []struct {
		Id        string `json:"id"`
		Ipaddress string `json:"ipaddress"`
	} `json:"secondaryip"`
	Traffictype      string `json:"traffictype"`
	Type             string `json:"type"`
	Virtualmachineid string `json:"virtualmachineid"`
	Vlanid           int    `json:"vlanid"`
}

type RemoveIpFromNicParams struct {
	P map[string]interface{}
}

func (P *RemoveIpFromNicParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *RemoveIpFromNicParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RemoveIpFromNicParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveIpFromNicParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewRemoveIpFromNicParams(id string) *RemoveIpFromNicParams {
	P := &RemoveIpFromNicParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Removes secondary IP from the NIC.
func (s *NicService) RemoveIpFromNic(p *RemoveIpFromNicParams) (*RemoveIpFromNicResponse, error) {
	resp, err := s.cs.newRequest("removeIpFromNic", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveIpFromNicResponse
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

type RemoveIpFromNicResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UpdateVmNicIpParams struct {
	P map[string]interface{}
}

func (P *UpdateVmNicIpParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := P.P["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	return u
}

func (P *UpdateVmNicIpParams) SetIpaddress(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ipaddress"] = v
}

func (P *UpdateVmNicIpParams) GetIpaddress() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ipaddress"].(string)
	return value, ok
}

func (P *UpdateVmNicIpParams) SetNicid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["nicid"] = v
}

func (P *UpdateVmNicIpParams) GetNicid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["nicid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVmNicIpParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewUpdateVmNicIpParams(nicid string) *UpdateVmNicIpParams {
	P := &UpdateVmNicIpParams{}
	P.P = make(map[string]interface{})
	P.P["nicid"] = nicid
	return P
}

// Update the default Ip of a VM Nic
func (s *NicService) UpdateVmNicIp(p *UpdateVmNicIpParams) (*UpdateVmNicIpResponse, error) {
	resp, err := s.cs.newRequest("updateVmNicIp", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVmNicIpResponse
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

type UpdateVmNicIpResponse struct {
	Account               string                               `json:"account"`
	Affinitygroup         []UpdateVmNicIpResponseAffinitygroup `json:"affinitygroup"`
	Backupofferingid      string                               `json:"backupofferingid"`
	Backupofferingname    string                               `json:"backupofferingname"`
	Bootmode              string                               `json:"bootmode"`
	Boottype              string                               `json:"boottype"`
	Cpunumber             int                                  `json:"cpunumber"`
	Cpuspeed              int                                  `json:"cpuspeed"`
	Cpuused               string                               `json:"cpuused"`
	Created               string                               `json:"created"`
	Details               map[string]string                    `json:"details"`
	Diskioread            int64                                `json:"diskioread"`
	Diskiowrite           int64                                `json:"diskiowrite"`
	Diskkbsread           int64                                `json:"diskkbsread"`
	Diskkbswrite          int64                                `json:"diskkbswrite"`
	Diskofferingid        string                               `json:"diskofferingid"`
	Diskofferingname      string                               `json:"diskofferingname"`
	Displayname           string                               `json:"displayname"`
	Displayvm             bool                                 `json:"displayvm"`
	Domain                string                               `json:"domain"`
	Domainid              string                               `json:"domainid"`
	Forvirtualnetwork     bool                                 `json:"forvirtualnetwork"`
	Group                 string                               `json:"group"`
	Groupid               string                               `json:"groupid"`
	Guestosid             string                               `json:"guestosid"`
	Haenable              bool                                 `json:"haenable"`
	Hasannotations        bool                                 `json:"hasannotations"`
	Hostid                string                               `json:"hostid"`
	Hostname              string                               `json:"hostname"`
	Hypervisor            string                               `json:"hypervisor"`
	Icon                  string                               `json:"icon"`
	Id                    string                               `json:"id"`
	Instancename          string                               `json:"instancename"`
	Isdynamicallyscalable bool                                 `json:"isdynamicallyscalable"`
	Isodisplaytext        string                               `json:"isodisplaytext"`
	Isoid                 string                               `json:"isoid"`
	Isoname               string                               `json:"isoname"`
	JobID                 string                               `json:"jobid"`
	Jobstatus             int                                  `json:"jobstatus"`
	Keypair               string                               `json:"keypair"`
	Lastupdated           string                               `json:"lastupdated"`
	Memory                int                                  `json:"memory"`
	Memoryintfreekbs      int64                                `json:"memoryintfreekbs"`
	Memorykbs             int64                                `json:"memorykbs"`
	Memorytargetkbs       int64                                `json:"memorytargetkbs"`
	Name                  string                               `json:"name"`
	Networkkbsread        int64                                `json:"networkkbsread"`
	Networkkbswrite       int64                                `json:"networkkbswrite"`
	Nic                   []Nic                                `json:"nic"`
	Osdisplayname         string                               `json:"osdisplayname"`
	Ostypeid              string                               `json:"ostypeid"`
	Password              string                               `json:"password"`
	Passwordenabled       bool                                 `json:"passwordenabled"`
	Pooltype              string                               `json:"pooltype"`
	Project               string                               `json:"project"`
	Projectid             string                               `json:"projectid"`
	Publicip              string                               `json:"publicip"`
	Publicipid            string                               `json:"publicipid"`
	Readonlydetails       string                               `json:"readonlydetails"`
	Receivedbytes         int64                                `json:"receivedbytes"`
	Rootdeviceid          int64                                `json:"rootdeviceid"`
	Rootdevicetype        string                               `json:"rootdevicetype"`
	Securitygroup         []UpdateVmNicIpResponseSecuritygroup `json:"securitygroup"`
	Sentbytes             int64                                `json:"sentbytes"`
	Serviceofferingid     string                               `json:"serviceofferingid"`
	Serviceofferingname   string                               `json:"serviceofferingname"`
	Servicestate          string                               `json:"servicestate"`
	State                 string                               `json:"state"`
	Tags                  []Tags                               `json:"tags"`
	Templatedisplaytext   string                               `json:"templatedisplaytext"`
	Templateid            string                               `json:"templateid"`
	Templatename          string                               `json:"templatename"`
	Userid                string                               `json:"userid"`
	Username              string                               `json:"username"`
	Vgpu                  string                               `json:"vgpu"`
	Zoneid                string                               `json:"zoneid"`
	Zonename              string                               `json:"zonename"`
}

type UpdateVmNicIpResponseSecuritygroup struct {
	Account             string                                   `json:"account"`
	Description         string                                   `json:"description"`
	Domain              string                                   `json:"domain"`
	Domainid            string                                   `json:"domainid"`
	Egressrule          []UpdateVmNicIpResponseSecuritygroupRule `json:"egressrule"`
	Id                  string                                   `json:"id"`
	Ingressrule         []UpdateVmNicIpResponseSecuritygroupRule `json:"ingressrule"`
	Name                string                                   `json:"name"`
	Project             string                                   `json:"project"`
	Projectid           string                                   `json:"projectid"`
	Tags                []Tags                                   `json:"tags"`
	Virtualmachinecount int                                      `json:"virtualmachinecount"`
	Virtualmachineids   []interface{}                            `json:"virtualmachineids"`
}

type UpdateVmNicIpResponseSecuritygroupRule struct {
	Account           string `json:"account"`
	Cidr              string `json:"cidr"`
	Endport           int    `json:"endport"`
	Icmpcode          int    `json:"icmpcode"`
	Icmptype          int    `json:"icmptype"`
	Protocol          string `json:"protocol"`
	Ruleid            string `json:"ruleid"`
	Securitygroupname string `json:"securitygroupname"`
	Startport         int    `json:"startport"`
	Tags              []Tags `json:"tags"`
}

type UpdateVmNicIpResponseAffinitygroup struct {
	Account           string   `json:"account"`
	Description       string   `json:"description"`
	Domain            string   `json:"domain"`
	Domainid          string   `json:"domainid"`
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	Project           string   `json:"project"`
	Projectid         string   `json:"projectid"`
	Type              string   `json:"type"`
	VirtualmachineIds []string `json:"virtualmachineIds"`
}

func (r *UpdateVmNicIpResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateVmNicIpResponse
	return json.Unmarshal(b, (*alias)(r))
}
