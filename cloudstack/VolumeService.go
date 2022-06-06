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

type VolumeServiceIface interface {
	AttachVolume(P *AttachVolumeParams) (*AttachVolumeResponse, error)
	NewAttachVolumeParams(id string, virtualmachineid string) *AttachVolumeParams
	CreateVolume(P *CreateVolumeParams) (*CreateVolumeResponse, error)
	NewCreateVolumeParams() *CreateVolumeParams
	DeleteVolume(P *DeleteVolumeParams) (*DeleteVolumeResponse, error)
	NewDeleteVolumeParams(id string) *DeleteVolumeParams
	DetachVolume(P *DetachVolumeParams) (*DetachVolumeResponse, error)
	NewDetachVolumeParams() *DetachVolumeParams
	ExtractVolume(P *ExtractVolumeParams) (*ExtractVolumeResponse, error)
	NewExtractVolumeParams(id string, mode string, zoneid string) *ExtractVolumeParams
	GetPathForVolume(P *GetPathForVolumeParams) (*GetPathForVolumeResponse, error)
	NewGetPathForVolumeParams(volumeid string) *GetPathForVolumeParams
	GetSolidFireVolumeSize(P *GetSolidFireVolumeSizeParams) (*GetSolidFireVolumeSizeResponse, error)
	NewGetSolidFireVolumeSizeParams(volumeid string) *GetSolidFireVolumeSizeParams
	GetUploadParamsForVolume(P *GetUploadParamsForVolumeParams) (*GetUploadParamsForVolumeResponse, error)
	NewGetUploadParamsForVolumeParams(format string, name string, zoneid string) *GetUploadParamsForVolumeParams
	GetVolumeiScsiName(P *GetVolumeiScsiNameParams) (*GetVolumeiScsiNameResponse, error)
	NewGetVolumeiScsiNameParams(volumeid string) *GetVolumeiScsiNameParams
	ListVolumes(P *ListVolumesParams) (*ListVolumesResponse, error)
	NewListVolumesParams() *ListVolumesParams
	GetVolumeID(name string, opts ...OptionFunc) (string, int, error)
	GetVolumeByName(name string, opts ...OptionFunc) (*Volume, int, error)
	GetVolumeByID(id string, opts ...OptionFunc) (*Volume, int, error)
	MigrateVolume(P *MigrateVolumeParams) (*MigrateVolumeResponse, error)
	NewMigrateVolumeParams(storageid string, volumeid string) *MigrateVolumeParams
	ResizeVolume(P *ResizeVolumeParams) (*ResizeVolumeResponse, error)
	NewResizeVolumeParams(id string) *ResizeVolumeParams
	UpdateVolume(P *UpdateVolumeParams) (*UpdateVolumeResponse, error)
	NewUpdateVolumeParams() *UpdateVolumeParams
	UploadVolume(P *UploadVolumeParams) (*UploadVolumeResponse, error)
	NewUploadVolumeParams(format string, name string, url string, zoneid string) *UploadVolumeParams
}

type AttachVolumeParams struct {
	P map[string]interface{}
}

func (P *AttachVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["deviceid"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("deviceid", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *AttachVolumeParams) SetDeviceid(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["deviceid"] = v
}

func (P *AttachVolumeParams) GetDeviceid() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["deviceid"].(int64)
	return value, ok
}

func (P *AttachVolumeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *AttachVolumeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *AttachVolumeParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *AttachVolumeParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new AttachVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewAttachVolumeParams(id string, virtualmachineid string) *AttachVolumeParams {
	P := &AttachVolumeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["virtualmachineid"] = virtualmachineid
	return P
}

// Attaches a disk volume to a virtual machine.
func (s *VolumeService) AttachVolume(p *AttachVolumeParams) (*AttachVolumeResponse, error) {
	resp, err := s.cs.newRequest("attachVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AttachVolumeResponse
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

type AttachVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type CreateVolumeParams struct {
	P map[string]interface{}
}

func (P *CreateVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := P.P["displayvolume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvolume", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
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
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	if v, found := P.P["snapshotid"]; found {
		u.Set("snapshotid", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *CreateVolumeParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateVolumeParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateVolumeParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *CreateVolumeParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *CreateVolumeParams) SetDiskofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["diskofferingid"] = v
}

func (P *CreateVolumeParams) GetDiskofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["diskofferingid"].(string)
	return value, ok
}

func (P *CreateVolumeParams) SetDisplayvolume(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayvolume"] = v
}

func (P *CreateVolumeParams) GetDisplayvolume() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayvolume"].(bool)
	return value, ok
}

func (P *CreateVolumeParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateVolumeParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateVolumeParams) SetMaxiops(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxiops"] = v
}

func (P *CreateVolumeParams) GetMaxiops() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxiops"].(int64)
	return value, ok
}

func (P *CreateVolumeParams) SetMiniops(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["miniops"] = v
}

func (P *CreateVolumeParams) GetMiniops() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["miniops"].(int64)
	return value, ok
}

func (P *CreateVolumeParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateVolumeParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateVolumeParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateVolumeParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *CreateVolumeParams) SetSize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["size"] = v
}

func (P *CreateVolumeParams) GetSize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["size"].(int64)
	return value, ok
}

func (P *CreateVolumeParams) SetSnapshotid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["snapshotid"] = v
}

func (P *CreateVolumeParams) GetSnapshotid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["snapshotid"].(string)
	return value, ok
}

func (P *CreateVolumeParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *CreateVolumeParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *CreateVolumeParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *CreateVolumeParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewCreateVolumeParams() *CreateVolumeParams {
	P := &CreateVolumeParams{}
	P.P = make(map[string]interface{})
	return P
}

// Creates a disk volume from a disk offering. This disk volume must still be attached to a virtual machine to make use of it.
func (s *VolumeService) CreateVolume(p *CreateVolumeParams) (*CreateVolumeResponse, error) {
	resp, err := s.cs.newRequest("createVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVolumeResponse
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

type CreateVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type DeleteVolumeParams struct {
	P map[string]interface{}
}

func (P *DeleteVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteVolumeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteVolumeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewDeleteVolumeParams(id string) *DeleteVolumeParams {
	P := &DeleteVolumeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a detached disk volume.
func (s *VolumeService) DeleteVolume(p *DeleteVolumeParams) (*DeleteVolumeResponse, error) {
	resp, err := s.cs.newRequest("deleteVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVolumeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteVolumeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteVolumeResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteVolumeResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DetachVolumeParams struct {
	P map[string]interface{}
}

func (P *DetachVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["deviceid"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("deviceid", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (P *DetachVolumeParams) SetDeviceid(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["deviceid"] = v
}

func (P *DetachVolumeParams) GetDeviceid() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["deviceid"].(int64)
	return value, ok
}

func (P *DetachVolumeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DetachVolumeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *DetachVolumeParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *DetachVolumeParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

// You should always use this function to get a new DetachVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewDetachVolumeParams() *DetachVolumeParams {
	P := &DetachVolumeParams{}
	P.P = make(map[string]interface{})
	return P
}

// Detaches a disk volume from a virtual machine.
func (s *VolumeService) DetachVolume(p *DetachVolumeParams) (*DetachVolumeResponse, error) {
	resp, err := s.cs.newRequest("detachVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DetachVolumeResponse
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

type DetachVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type ExtractVolumeParams struct {
	P map[string]interface{}
}

func (P *ExtractVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["mode"]; found {
		u.Set("mode", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ExtractVolumeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ExtractVolumeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ExtractVolumeParams) SetMode(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["mode"] = v
}

func (P *ExtractVolumeParams) GetMode() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["mode"].(string)
	return value, ok
}

func (P *ExtractVolumeParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *ExtractVolumeParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *ExtractVolumeParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ExtractVolumeParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ExtractVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewExtractVolumeParams(id string, mode string, zoneid string) *ExtractVolumeParams {
	P := &ExtractVolumeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["mode"] = mode
	P.P["zoneid"] = zoneid
	return P
}

// Extracts volume
func (s *VolumeService) ExtractVolume(p *ExtractVolumeParams) (*ExtractVolumeResponse, error) {
	resp, err := s.cs.newRequest("extractVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ExtractVolumeResponse
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

type ExtractVolumeResponse struct {
	Accountid        string `json:"accountid"`
	Created          string `json:"created"`
	ExtractId        string `json:"extractId"`
	ExtractMode      string `json:"extractMode"`
	Id               string `json:"id"`
	JobID            string `json:"jobid"`
	Jobstatus        int    `json:"jobstatus"`
	Name             string `json:"name"`
	Resultstring     string `json:"resultstring"`
	State            string `json:"state"`
	Status           string `json:"status"`
	Storagetype      string `json:"storagetype"`
	Uploadpercentage int    `json:"uploadpercentage"`
	Url              string `json:"url"`
	Zoneid           string `json:"zoneid"`
	Zonename         string `json:"zonename"`
}

type GetPathForVolumeParams struct {
	P map[string]interface{}
}

func (P *GetPathForVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (P *GetPathForVolumeParams) SetVolumeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["volumeid"] = v
}

func (P *GetPathForVolumeParams) GetVolumeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new GetPathForVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewGetPathForVolumeParams(volumeid string) *GetPathForVolumeParams {
	P := &GetPathForVolumeParams{}
	P.P = make(map[string]interface{})
	P.P["volumeid"] = volumeid
	return P
}

// Get the path associated with the provided volume UUID
func (s *VolumeService) GetPathForVolume(p *GetPathForVolumeParams) (*GetPathForVolumeResponse, error) {
	resp, err := s.cs.newRequest("getPathForVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetPathForVolumeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetPathForVolumeResponse struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Path      string `json:"path"`
}

type GetSolidFireVolumeSizeParams struct {
	P map[string]interface{}
}

func (P *GetSolidFireVolumeSizeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (P *GetSolidFireVolumeSizeParams) SetVolumeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["volumeid"] = v
}

func (P *GetSolidFireVolumeSizeParams) GetVolumeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new GetSolidFireVolumeSizeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewGetSolidFireVolumeSizeParams(volumeid string) *GetSolidFireVolumeSizeParams {
	P := &GetSolidFireVolumeSizeParams{}
	P.P = make(map[string]interface{})
	P.P["volumeid"] = volumeid
	return P
}

// Get the SF volume size including Hypervisor Snapshot Reserve
func (s *VolumeService) GetSolidFireVolumeSize(p *GetSolidFireVolumeSizeParams) (*GetSolidFireVolumeSizeResponse, error) {
	resp, err := s.cs.newRequest("getSolidFireVolumeSize", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetSolidFireVolumeSizeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetSolidFireVolumeSizeResponse struct {
	JobID               string `json:"jobid"`
	Jobstatus           int    `json:"jobstatus"`
	SolidFireVolumeSize int64  `json:"solidFireVolumeSize"`
}

type GetUploadParamsForVolumeParams struct {
	P map[string]interface{}
}

func (P *GetUploadParamsForVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := P.P["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := P.P["imagestoreuuid"]; found {
		u.Set("imagestoreuuid", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *GetUploadParamsForVolumeParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *GetUploadParamsForVolumeParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *GetUploadParamsForVolumeParams) SetChecksum(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["checksum"] = v
}

func (P *GetUploadParamsForVolumeParams) GetChecksum() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["checksum"].(string)
	return value, ok
}

func (P *GetUploadParamsForVolumeParams) SetDiskofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["diskofferingid"] = v
}

func (P *GetUploadParamsForVolumeParams) GetDiskofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["diskofferingid"].(string)
	return value, ok
}

func (P *GetUploadParamsForVolumeParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *GetUploadParamsForVolumeParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *GetUploadParamsForVolumeParams) SetFormat(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["format"] = v
}

func (P *GetUploadParamsForVolumeParams) GetFormat() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["format"].(string)
	return value, ok
}

func (P *GetUploadParamsForVolumeParams) SetImagestoreuuid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["imagestoreuuid"] = v
}

func (P *GetUploadParamsForVolumeParams) GetImagestoreuuid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["imagestoreuuid"].(string)
	return value, ok
}

func (P *GetUploadParamsForVolumeParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *GetUploadParamsForVolumeParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *GetUploadParamsForVolumeParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *GetUploadParamsForVolumeParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *GetUploadParamsForVolumeParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *GetUploadParamsForVolumeParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new GetUploadParamsForVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewGetUploadParamsForVolumeParams(format string, name string, zoneid string) *GetUploadParamsForVolumeParams {
	P := &GetUploadParamsForVolumeParams{}
	P.P = make(map[string]interface{})
	P.P["format"] = format
	P.P["name"] = name
	P.P["zoneid"] = zoneid
	return P
}

// Upload a data disk to the cloudstack cloud.
func (s *VolumeService) GetUploadParamsForVolume(p *GetUploadParamsForVolumeParams) (*GetUploadParamsForVolumeResponse, error) {
	resp, err := s.cs.newRequest("getUploadParamsForVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var nested struct {
		Response GetUploadParamsForVolumeResponse `json:"getuploadparams"`
	}
	if err := json.Unmarshal(resp, &nested); err != nil {
		return nil, err
	}
	r := nested.Response

	return &r, nil
}

type GetUploadParamsForVolumeResponse struct {
	Expires   string `json:"expires"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Metadata  string `json:"metadata"`
	PostURL   string `json:"postURL"`
	Signature string `json:"signature"`
}

type GetVolumeiScsiNameParams struct {
	P map[string]interface{}
}

func (P *GetVolumeiScsiNameParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (P *GetVolumeiScsiNameParams) SetVolumeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["volumeid"] = v
}

func (P *GetVolumeiScsiNameParams) GetVolumeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new GetVolumeiScsiNameParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewGetVolumeiScsiNameParams(volumeid string) *GetVolumeiScsiNameParams {
	P := &GetVolumeiScsiNameParams{}
	P.P = make(map[string]interface{})
	P.P["volumeid"] = volumeid
	return P
}

// Get Volume's iSCSI Name
func (s *VolumeService) GetVolumeiScsiName(p *GetVolumeiScsiNameParams) (*GetVolumeiScsiNameResponse, error) {
	resp, err := s.cs.newRequest("getVolumeiScsiName", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r GetVolumeiScsiNameResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type GetVolumeiScsiNameResponse struct {
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	VolumeiScsiName string `json:"volumeiScsiName"`
}

type ListVolumesParams struct {
	P map[string]interface{}
}

func (P *ListVolumesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := P.P["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := P.P["displayvolume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvolume", vv)
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["ids"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ids", vv)
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
	if v, found := P.P["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := P.P["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *ListVolumesParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListVolumesParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetClusterid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["clusterid"] = v
}

func (P *ListVolumesParams) GetClusterid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["clusterid"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetDiskofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["diskofferingid"] = v
}

func (P *ListVolumesParams) GetDiskofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["diskofferingid"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetDisplayvolume(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayvolume"] = v
}

func (P *ListVolumesParams) GetDisplayvolume() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayvolume"].(bool)
	return value, ok
}

func (P *ListVolumesParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListVolumesParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetHostid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["hostid"] = v
}

func (P *ListVolumesParams) GetHostid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["hostid"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListVolumesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetIds(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ids"] = v
}

func (P *ListVolumesParams) GetIds() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ids"].([]string)
	return value, ok
}

func (P *ListVolumesParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListVolumesParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListVolumesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListVolumesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListVolumesParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListVolumesParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListVolumesParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListVolumesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListVolumesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListVolumesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListVolumesParams) SetPodid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["podid"] = v
}

func (P *ListVolumesParams) GetPodid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["podid"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListVolumesParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListVolumesParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *ListVolumesParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListVolumesParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListVolumesParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ListVolumesParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetVirtualmachineid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["virtualmachineid"] = v
}

func (P *ListVolumesParams) GetVirtualmachineid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["virtualmachineid"].(string)
	return value, ok
}

func (P *ListVolumesParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *ListVolumesParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVolumesParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewListVolumesParams() *ListVolumesParams {
	P := &ListVolumesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListVolumesParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVolumes(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Volumes[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Volumes {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeByName(name string, opts ...OptionFunc) (*Volume, int, error) {
	id, count, err := s.GetVolumeID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVolumeByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeByID(id string, opts ...OptionFunc) (*Volume, int, error) {
	P := &ListVolumesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVolumes(P)
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
		return l.Volumes[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Volume UUID: %s!", id)
}

// Lists all volumes.
func (s *VolumeService) ListVolumes(p *ListVolumesParams) (*ListVolumesResponse, error) {
	resp, err := s.cs.newRequest("listVolumes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVolumesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVolumesResponse struct {
	Count   int       `json:"count"`
	Volumes []*Volume `json:"volume"`
}

type Volume struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type MigrateVolumeParams struct {
	P map[string]interface{}
}

func (P *MigrateVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["livemigrate"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("livemigrate", vv)
	}
	if v, found := P.P["newdiskofferingid"]; found {
		u.Set("newdiskofferingid", v.(string))
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := P.P["volumeid"]; found {
		u.Set("volumeid", v.(string))
	}
	return u
}

func (P *MigrateVolumeParams) SetLivemigrate(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["livemigrate"] = v
}

func (P *MigrateVolumeParams) GetLivemigrate() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["livemigrate"].(bool)
	return value, ok
}

func (P *MigrateVolumeParams) SetNewdiskofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["newdiskofferingid"] = v
}

func (P *MigrateVolumeParams) GetNewdiskofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["newdiskofferingid"].(string)
	return value, ok
}

func (P *MigrateVolumeParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *MigrateVolumeParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

func (P *MigrateVolumeParams) SetVolumeid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["volumeid"] = v
}

func (P *MigrateVolumeParams) GetVolumeid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["volumeid"].(string)
	return value, ok
}

// You should always use this function to get a new MigrateVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewMigrateVolumeParams(storageid string, volumeid string) *MigrateVolumeParams {
	P := &MigrateVolumeParams{}
	P.P = make(map[string]interface{})
	P.P["storageid"] = storageid
	P.P["volumeid"] = volumeid
	return P
}

// Migrate volume
func (s *VolumeService) MigrateVolume(p *MigrateVolumeParams) (*MigrateVolumeResponse, error) {
	resp, err := s.cs.newRequest("migrateVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateVolumeResponse
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

type MigrateVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type ResizeVolumeParams struct {
	P map[string]interface{}
}

func (P *ResizeVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["maxiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxiops", vv)
	}
	if v, found := P.P["miniops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("miniops", vv)
	}
	if v, found := P.P["shrinkok"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("shrinkok", vv)
	}
	if v, found := P.P["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	return u
}

func (P *ResizeVolumeParams) SetDiskofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["diskofferingid"] = v
}

func (P *ResizeVolumeParams) GetDiskofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["diskofferingid"].(string)
	return value, ok
}

func (P *ResizeVolumeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ResizeVolumeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ResizeVolumeParams) SetMaxiops(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["maxiops"] = v
}

func (P *ResizeVolumeParams) GetMaxiops() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["maxiops"].(int64)
	return value, ok
}

func (P *ResizeVolumeParams) SetMiniops(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["miniops"] = v
}

func (P *ResizeVolumeParams) GetMiniops() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["miniops"].(int64)
	return value, ok
}

func (P *ResizeVolumeParams) SetShrinkok(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["shrinkok"] = v
}

func (P *ResizeVolumeParams) GetShrinkok() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["shrinkok"].(bool)
	return value, ok
}

func (P *ResizeVolumeParams) SetSize(v int64) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["size"] = v
}

func (P *ResizeVolumeParams) GetSize() (int64, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["size"].(int64)
	return value, ok
}

// You should always use this function to get a new ResizeVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewResizeVolumeParams(id string) *ResizeVolumeParams {
	P := &ResizeVolumeParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Resizes a volume
func (s *VolumeService) ResizeVolume(p *ResizeVolumeParams) (*ResizeVolumeResponse, error) {
	resp, err := s.cs.newRequest("resizeVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResizeVolumeResponse
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

type ResizeVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type UpdateVolumeParams struct {
	P map[string]interface{}
}

func (P *UpdateVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["chaininfo"]; found {
		u.Set("chaininfo", v.(string))
	}
	if v, found := P.P["customid"]; found {
		u.Set("customid", v.(string))
	}
	if v, found := P.P["displayvolume"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("displayvolume", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["path"]; found {
		u.Set("path", v.(string))
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	return u
}

func (P *UpdateVolumeParams) SetChaininfo(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["chaininfo"] = v
}

func (P *UpdateVolumeParams) GetChaininfo() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["chaininfo"].(string)
	return value, ok
}

func (P *UpdateVolumeParams) SetCustomid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["customid"] = v
}

func (P *UpdateVolumeParams) GetCustomid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["customid"].(string)
	return value, ok
}

func (P *UpdateVolumeParams) SetDisplayvolume(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displayvolume"] = v
}

func (P *UpdateVolumeParams) GetDisplayvolume() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displayvolume"].(bool)
	return value, ok
}

func (P *UpdateVolumeParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateVolumeParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateVolumeParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateVolumeParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateVolumeParams) SetPath(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["path"] = v
}

func (P *UpdateVolumeParams) GetPath() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["path"].(string)
	return value, ok
}

func (P *UpdateVolumeParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *UpdateVolumeParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *UpdateVolumeParams) SetStorageid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["storageid"] = v
}

func (P *UpdateVolumeParams) GetStorageid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["storageid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewUpdateVolumeParams() *UpdateVolumeParams {
	P := &UpdateVolumeParams{}
	P.P = make(map[string]interface{})
	return P
}

// Updates the volume.
func (s *VolumeService) UpdateVolume(p *UpdateVolumeParams) (*UpdateVolumeResponse, error) {
	resp, err := s.cs.newRequest("updateVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVolumeResponse
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

type UpdateVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}

type UploadVolumeParams struct {
	P map[string]interface{}
}

func (P *UploadVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["checksum"]; found {
		u.Set("checksum", v.(string))
	}
	if v, found := P.P["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["format"]; found {
		u.Set("format", v.(string))
	}
	if v, found := P.P["imagestoreuuid"]; found {
		u.Set("imagestoreuuid", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := P.P["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (P *UploadVolumeParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *UploadVolumeParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *UploadVolumeParams) SetChecksum(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["checksum"] = v
}

func (P *UploadVolumeParams) GetChecksum() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["checksum"].(string)
	return value, ok
}

func (P *UploadVolumeParams) SetDiskofferingid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["diskofferingid"] = v
}

func (P *UploadVolumeParams) GetDiskofferingid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["diskofferingid"].(string)
	return value, ok
}

func (P *UploadVolumeParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *UploadVolumeParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *UploadVolumeParams) SetFormat(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["format"] = v
}

func (P *UploadVolumeParams) GetFormat() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["format"].(string)
	return value, ok
}

func (P *UploadVolumeParams) SetImagestoreuuid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["imagestoreuuid"] = v
}

func (P *UploadVolumeParams) GetImagestoreuuid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["imagestoreuuid"].(string)
	return value, ok
}

func (P *UploadVolumeParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UploadVolumeParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UploadVolumeParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *UploadVolumeParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *UploadVolumeParams) SetUrl(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["url"] = v
}

func (P *UploadVolumeParams) GetUrl() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["url"].(string)
	return value, ok
}

func (P *UploadVolumeParams) SetZoneid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["zoneid"] = v
}

func (P *UploadVolumeParams) GetZoneid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UploadVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewUploadVolumeParams(format string, name string, url string, zoneid string) *UploadVolumeParams {
	P := &UploadVolumeParams{}
	P.P = make(map[string]interface{})
	P.P["format"] = format
	P.P["name"] = name
	P.P["url"] = url
	P.P["zoneid"] = zoneid
	return P
}

// Uploads a data disk.
func (s *VolumeService) UploadVolume(p *UploadVolumeParams) (*UploadVolumeResponse, error) {
	resp, err := s.cs.newRequest("uploadVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadVolumeResponse
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

type UploadVolumeResponse struct {
	Account                    string `json:"account"`
	Attached                   string `json:"attached"`
	Chaininfo                  string `json:"chaininfo"`
	Clusterid                  string `json:"clusterid"`
	Clustername                string `json:"clustername"`
	Created                    string `json:"created"`
	Destroyed                  bool   `json:"destroyed"`
	Deviceid                   int64  `json:"deviceid"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate"`
	Diskioread                 int64  `json:"diskioread"`
	Diskiowrite                int64  `json:"diskiowrite"`
	Diskkbsread                int64  `json:"diskkbsread"`
	Diskkbswrite               int64  `json:"diskkbswrite"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext"`
	Diskofferingid             string `json:"diskofferingid"`
	Diskofferingname           string `json:"diskofferingname"`
	Displayvolume              bool   `json:"displayvolume"`
	Domain                     string `json:"domain"`
	Domainid                   string `json:"domainid"`
	Hasannotations             bool   `json:"hasannotations"`
	Hypervisor                 string `json:"hypervisor"`
	Id                         string `json:"id"`
	Isextractable              bool   `json:"isextractable"`
	Isodisplaytext             string `json:"isodisplaytext"`
	Isoid                      string `json:"isoid"`
	Isoname                    string `json:"isoname"`
	JobID                      string `json:"jobid"`
	Jobstatus                  int    `json:"jobstatus"`
	Maxiops                    int64  `json:"maxiops"`
	Miniops                    int64  `json:"miniops"`
	Name                       string `json:"name"`
	Path                       string `json:"path"`
	Physicalsize               int64  `json:"physicalsize"`
	Podid                      string `json:"podid"`
	Podname                    string `json:"podname"`
	Project                    string `json:"project"`
	Projectid                  string `json:"projectid"`
	Provisioningtype           string `json:"provisioningtype"`
	Quiescevm                  bool   `json:"quiescevm"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext"`
	Serviceofferingid          string `json:"serviceofferingid"`
	Serviceofferingname        string `json:"serviceofferingname"`
	Size                       int64  `json:"size"`
	Snapshotid                 string `json:"snapshotid"`
	State                      string `json:"state"`
	Status                     string `json:"status"`
	Storage                    string `json:"storage"`
	Storageid                  string `json:"storageid"`
	Storagetype                string `json:"storagetype"`
	Supportsstoragesnapshot    bool   `json:"supportsstoragesnapshot"`
	Tags                       []Tags `json:"tags"`
	Templatedisplaytext        string `json:"templatedisplaytext"`
	Templateid                 string `json:"templateid"`
	Templatename               string `json:"templatename"`
	Type                       string `json:"type"`
	Utilization                string `json:"utilization"`
	Virtualmachineid           string `json:"virtualmachineid"`
	Virtualsize                int64  `json:"virtualsize"`
	Vmdisplayname              string `json:"vmdisplayname"`
	Vmname                     string `json:"vmname"`
	Vmstate                    string `json:"vmstate"`
	Zoneid                     string `json:"zoneid"`
	Zonename                   string `json:"zonename"`
}
