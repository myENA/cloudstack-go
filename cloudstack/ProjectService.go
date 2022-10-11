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

type ProjectServiceIface interface {
	ActivateProject(p *ActivateProjectParams) (*ActivateProjectResponse, error)
	NewActivateProjectParams(id string) *ActivateProjectParams
	AddAccountToProject(p *AddAccountToProjectParams) (*AddAccountToProjectResponse, error)
	NewAddAccountToProjectParams(projectid string) *AddAccountToProjectParams
	AddUserToProject(p *AddUserToProjectParams) (*AddUserToProjectResponse, error)
	NewAddUserToProjectParams(projectid string, username string) *AddUserToProjectParams
	CreateProject(p *CreateProjectParams) (*CreateProjectResponse, error)
	NewCreateProjectParams(displaytext string, name string) *CreateProjectParams
	DeleteAccountFromProject(p *DeleteAccountFromProjectParams) (*DeleteAccountFromProjectResponse, error)
	NewDeleteAccountFromProjectParams(account string, projectid string) *DeleteAccountFromProjectParams
	DeleteUserFromProject(p *DeleteUserFromProjectParams) (*DeleteUserFromProjectResponse, error)
	NewDeleteUserFromProjectParams(projectid string, userid string) *DeleteUserFromProjectParams
	DeleteProject(p *DeleteProjectParams) (*DeleteProjectResponse, error)
	NewDeleteProjectParams(id string) *DeleteProjectParams
	DeleteProjectInvitation(p *DeleteProjectInvitationParams) (*DeleteProjectInvitationResponse, error)
	NewDeleteProjectInvitationParams(id string) *DeleteProjectInvitationParams
	ListProjectInvitations(p *ListProjectInvitationsParams) (*ListProjectInvitationsResponse, error)
	NewListProjectInvitationsParams() *ListProjectInvitationsParams
	GetProjectInvitationByID(id string, opts ...OptionFunc) (*ProjectInvitation, int, error)
	ListProjects(p *ListProjectsParams) (*ListProjectsResponse, error)
	NewListProjectsParams() *ListProjectsParams
	GetProjectID(name string, opts ...OptionFunc) (string, int, error)
	GetProjectByName(name string, opts ...OptionFunc) (*Project, int, error)
	GetProjectByID(id string, opts ...OptionFunc) (*Project, int, error)
	SuspendProject(p *SuspendProjectParams) (*SuspendProjectResponse, error)
	NewSuspendProjectParams(id string) *SuspendProjectParams
	UpdateProject(p *UpdateProjectParams) (*UpdateProjectResponse, error)
	NewUpdateProjectParams(id string) *UpdateProjectParams
	UpdateProjectInvitation(p *UpdateProjectInvitationParams) (*UpdateProjectInvitationResponse, error)
	NewUpdateProjectInvitationParams(projectid string) *UpdateProjectInvitationParams
	ListProjectRolePermissions(p *ListProjectRolePermissionsParams) (*ListProjectRolePermissionsResponse, error)
	NewListProjectRolePermissionsParams(projectid string) *ListProjectRolePermissionsParams
	CreateProjectRolePermission(p *CreateProjectRolePermissionParams) (*CreateProjectRolePermissionResponse, error)
	NewCreateProjectRolePermissionParams(permission string, projectid string, projectroleid string, rule string) *CreateProjectRolePermissionParams
	UpdateProjectRolePermission(p *UpdateProjectRolePermissionParams) (*UpdateProjectRolePermissionResponse, error)
	NewUpdateProjectRolePermissionParams(projectid string, projectroleid string) *UpdateProjectRolePermissionParams
	DeleteProjectRolePermission(p *DeleteProjectRolePermissionParams) (*DeleteProjectRolePermissionResponse, error)
	NewDeleteProjectRolePermissionParams(id string, projectid string) *DeleteProjectRolePermissionParams
}

type ActivateProjectParams struct {
	P map[string]interface{}
}

func (P *ActivateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *ActivateProjectParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ActivateProjectParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new ActivateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewActivateProjectParams(id string) *ActivateProjectParams {
	P := &ActivateProjectParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Activates a project
func (s *ProjectService) ActivateProject(p *ActivateProjectParams) (*ActivateProjectResponse, error) {
	resp, err := s.cs.newRequest("activateProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ActivateProjectResponse
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

type ActivateProjectResponse struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      string              `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type AddAccountToProjectParams struct {
	P map[string]interface{}
}

func (P *AddAccountToProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	if v, found := P.P["roletype"]; found {
		u.Set("roletype", v.(string))
	}
	return u
}

func (P *AddAccountToProjectParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *AddAccountToProjectParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *AddAccountToProjectParams) SetEmail(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["email"] = v
}

func (P *AddAccountToProjectParams) GetEmail() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["email"].(string)
	return value, ok
}

func (P *AddAccountToProjectParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *AddAccountToProjectParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *AddAccountToProjectParams) SetProjectroleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectroleid"] = v
}

func (P *AddAccountToProjectParams) GetProjectroleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectroleid"].(string)
	return value, ok
}

func (P *AddAccountToProjectParams) SetRoletype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roletype"] = v
}

func (P *AddAccountToProjectParams) GetRoletype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roletype"].(string)
	return value, ok
}

// You should always use this function to get a new AddAccountToProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewAddAccountToProjectParams(projectid string) *AddAccountToProjectParams {
	P := &AddAccountToProjectParams{}
	P.P = make(map[string]interface{})
	P.P["projectid"] = projectid
	return P
}

// Adds account to a project
func (s *ProjectService) AddAccountToProject(p *AddAccountToProjectParams) (*AddAccountToProjectResponse, error) {
	resp, err := s.cs.newRequest("addAccountToProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddAccountToProjectResponse
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

type AddAccountToProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type AddUserToProjectParams struct {
	P map[string]interface{}
}

func (P *AddUserToProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["email"]; found {
		u.Set("email", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	if v, found := P.P["roletype"]; found {
		u.Set("roletype", v.(string))
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *AddUserToProjectParams) SetEmail(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["email"] = v
}

func (P *AddUserToProjectParams) GetEmail() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["email"].(string)
	return value, ok
}

func (P *AddUserToProjectParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *AddUserToProjectParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *AddUserToProjectParams) SetProjectroleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectroleid"] = v
}

func (P *AddUserToProjectParams) GetProjectroleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectroleid"].(string)
	return value, ok
}

func (P *AddUserToProjectParams) SetRoletype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roletype"] = v
}

func (P *AddUserToProjectParams) GetRoletype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roletype"].(string)
	return value, ok
}

func (P *AddUserToProjectParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *AddUserToProjectParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new AddUserToProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewAddUserToProjectParams(projectid string, username string) *AddUserToProjectParams {
	P := &AddUserToProjectParams{}
	P.P = make(map[string]interface{})
	P.P["projectid"] = projectid
	P.P["username"] = username
	return P
}

// Adds user to a project
func (s *ProjectService) AddUserToProject(p *AddUserToProjectParams) (*AddUserToProjectResponse, error) {
	resp, err := s.cs.newRequest("addUserToProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddUserToProjectResponse
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

type AddUserToProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type CreateProjectParams struct {
	P map[string]interface{}
}

func (P *CreateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (P *CreateProjectParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *CreateProjectParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *CreateProjectParams) SetAccountid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accountid"] = v
}

func (P *CreateProjectParams) GetAccountid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accountid"].(string)
	return value, ok
}

func (P *CreateProjectParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *CreateProjectParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *CreateProjectParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *CreateProjectParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *CreateProjectParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateProjectParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateProjectParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *CreateProjectParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewCreateProjectParams(displaytext string, name string) *CreateProjectParams {
	P := &CreateProjectParams{}
	P.P = make(map[string]interface{})
	P.P["displaytext"] = displaytext
	P.P["name"] = name
	return P
}

// Creates a project
func (s *ProjectService) CreateProject(p *CreateProjectParams) (*CreateProjectResponse, error) {
	resp, err := s.cs.newRequest("createProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateProjectResponse
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

type CreateProjectResponse struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      string              `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type DeleteAccountFromProjectParams struct {
	P map[string]interface{}
}

func (P *DeleteAccountFromProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *DeleteAccountFromProjectParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *DeleteAccountFromProjectParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *DeleteAccountFromProjectParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *DeleteAccountFromProjectParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteAccountFromProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteAccountFromProjectParams(account string, projectid string) *DeleteAccountFromProjectParams {
	P := &DeleteAccountFromProjectParams{}
	P.P = make(map[string]interface{})
	P.P["account"] = account
	P.P["projectid"] = projectid
	return P
}

// Deletes account from the project
func (s *ProjectService) DeleteAccountFromProject(p *DeleteAccountFromProjectParams) (*DeleteAccountFromProjectResponse, error) {
	resp, err := s.cs.newRequest("deleteAccountFromProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteAccountFromProjectResponse
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

type DeleteAccountFromProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteUserFromProjectParams struct {
	P map[string]interface{}
}

func (P *DeleteUserFromProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (P *DeleteUserFromProjectParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *DeleteUserFromProjectParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *DeleteUserFromProjectParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *DeleteUserFromProjectParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteUserFromProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteUserFromProjectParams(projectid string, userid string) *DeleteUserFromProjectParams {
	P := &DeleteUserFromProjectParams{}
	P.P = make(map[string]interface{})
	P.P["projectid"] = projectid
	P.P["userid"] = userid
	return P
}

// Deletes user from the project
func (s *ProjectService) DeleteUserFromProject(p *DeleteUserFromProjectParams) (*DeleteUserFromProjectResponse, error) {
	resp, err := s.cs.newRequest("deleteUserFromProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteUserFromProjectResponse
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

type DeleteUserFromProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteProjectParams struct {
	P map[string]interface{}
}

func (P *DeleteProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["cleanup"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("cleanup", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteProjectParams) SetCleanup(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["cleanup"] = v
}

func (P *DeleteProjectParams) GetCleanup() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["cleanup"].(bool)
	return value, ok
}

func (P *DeleteProjectParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteProjectParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteProjectParams(id string) *DeleteProjectParams {
	P := &DeleteProjectParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a project
func (s *ProjectService) DeleteProject(p *DeleteProjectParams) (*DeleteProjectResponse, error) {
	resp, err := s.cs.newRequest("deleteProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteProjectResponse
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

type DeleteProjectResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteProjectInvitationParams struct {
	P map[string]interface{}
}

func (P *DeleteProjectInvitationParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteProjectInvitationParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteProjectInvitationParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteProjectInvitationParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteProjectInvitationParams(id string) *DeleteProjectInvitationParams {
	P := &DeleteProjectInvitationParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes project invitation
func (s *ProjectService) DeleteProjectInvitation(p *DeleteProjectInvitationParams) (*DeleteProjectInvitationResponse, error) {
	resp, err := s.cs.newRequest("deleteProjectInvitation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteProjectInvitationResponse
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

type DeleteProjectInvitationResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListProjectInvitationsParams struct {
	P map[string]interface{}
}

func (P *ListProjectInvitationsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["activeonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("activeonly", vv)
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
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (P *ListProjectInvitationsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListProjectInvitationsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetActiveonly(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["activeonly"] = v
}

func (P *ListProjectInvitationsParams) GetActiveonly() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["activeonly"].(bool)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListProjectInvitationsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListProjectInvitationsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListProjectInvitationsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListProjectInvitationsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListProjectInvitationsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListProjectInvitationsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListProjectInvitationsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListProjectInvitationsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListProjectInvitationsParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListProjectInvitationsParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *ListProjectInvitationsParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

// You should always use this function to get a new ListProjectInvitationsParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewListProjectInvitationsParams() *ListProjectInvitationsParams {
	P := &ListProjectInvitationsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectInvitationByID(id string, opts ...OptionFunc) (*ProjectInvitation, int, error) {
	P := &ListProjectInvitationsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListProjectInvitations(P)
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
		return l.ProjectInvitations[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ProjectInvitation UUID: %s!", id)
}

// Lists project invitations and provides detailed information for listed invitations
func (s *ProjectService) ListProjectInvitations(p *ListProjectInvitationsParams) (*ListProjectInvitationsResponse, error) {
	resp, err := s.cs.newRequest("listProjectInvitations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectInvitationsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListProjectInvitationsResponse struct {
	Count              int                  `json:"count"`
	ProjectInvitations []*ProjectInvitation `json:"projectinvitation"`
}

type ProjectInvitation struct {
	Account   string `json:"account"`
	Domain    string `json:"domain"`
	Domainid  string `json:"domainid"`
	Email     string `json:"email"`
	Id        string `json:"id"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Project   string `json:"project"`
	Projectid string `json:"projectid"`
	State     string `json:"state"`
	Userid    string `json:"userid"`
}

type ListProjectsParams struct {
	P map[string]interface{}
}

func (P *ListProjectsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["details"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("details", vv)
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
	if v, found := P.P["showicon"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showicon", vv)
	}
	if v, found := P.P["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := P.P["tags"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), m[k])
		}
	}
	if v, found := P.P["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (P *ListProjectsParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *ListProjectsParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *ListProjectsParams) SetDetails(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["details"] = v
}

func (P *ListProjectsParams) GetDetails() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["details"].([]string)
	return value, ok
}

func (P *ListProjectsParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *ListProjectsParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *ListProjectsParams) SetDomainid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["domainid"] = v
}

func (P *ListProjectsParams) GetDomainid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["domainid"].(string)
	return value, ok
}

func (P *ListProjectsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListProjectsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListProjectsParams) SetIsrecursive(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["isrecursive"] = v
}

func (P *ListProjectsParams) GetIsrecursive() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["isrecursive"].(bool)
	return value, ok
}

func (P *ListProjectsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListProjectsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListProjectsParams) SetListall(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["listall"] = v
}

func (P *ListProjectsParams) GetListall() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["listall"].(bool)
	return value, ok
}

func (P *ListProjectsParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListProjectsParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListProjectsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListProjectsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListProjectsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListProjectsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListProjectsParams) SetShowicon(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["showicon"] = v
}

func (P *ListProjectsParams) GetShowicon() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["showicon"].(bool)
	return value, ok
}

func (P *ListProjectsParams) SetState(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["state"] = v
}

func (P *ListProjectsParams) GetState() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["state"].(string)
	return value, ok
}

func (P *ListProjectsParams) SetTags(v map[string]string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["tags"] = v
}

func (P *ListProjectsParams) GetTags() (map[string]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["tags"].(map[string]string)
	return value, ok
}

func (P *ListProjectsParams) SetUsername(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["username"] = v
}

func (P *ListProjectsParams) GetUsername() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["username"].(string)
	return value, ok
}

// You should always use this function to get a new ListProjectsParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewListProjectsParams() *ListProjectsParams {
	P := &ListProjectsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListProjectsParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListProjects(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Projects[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Projects {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectByName(name string, opts ...OptionFunc) (*Project, int, error) {
	id, count, err := s.GetProjectID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetProjectByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *ProjectService) GetProjectByID(id string, opts ...OptionFunc) (*Project, int, error) {
	P := &ListProjectsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListProjects(P)
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
		return l.Projects[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Project UUID: %s!", id)
}

// Lists projects and provides detailed information for listed projects
func (s *ProjectService) ListProjects(p *ListProjectsParams) (*ListProjectsResponse, error) {
	resp, err := s.cs.newRequest("listProjects", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListProjectsResponse struct {
	Count    int        `json:"count"`
	Projects []*Project `json:"project"`
}

type Project struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      string              `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type SuspendProjectParams struct {
	P map[string]interface{}
}

func (P *SuspendProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *SuspendProjectParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *SuspendProjectParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new SuspendProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewSuspendProjectParams(id string) *SuspendProjectParams {
	P := &SuspendProjectParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Suspends a project
func (s *ProjectService) SuspendProject(p *SuspendProjectParams) (*SuspendProjectResponse, error) {
	resp, err := s.cs.newRequest("suspendProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r SuspendProjectResponse
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

type SuspendProjectResponse struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      string              `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type UpdateProjectParams struct {
	P map[string]interface{}
}

func (P *UpdateProjectParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["roletype"]; found {
		u.Set("roletype", v.(string))
	}
	if v, found := P.P["swapowner"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("swapowner", vv)
	}
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (P *UpdateProjectParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *UpdateProjectParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *UpdateProjectParams) SetDisplaytext(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["displaytext"] = v
}

func (P *UpdateProjectParams) GetDisplaytext() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["displaytext"].(string)
	return value, ok
}

func (P *UpdateProjectParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateProjectParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateProjectParams) SetRoletype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roletype"] = v
}

func (P *UpdateProjectParams) GetRoletype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roletype"].(string)
	return value, ok
}

func (P *UpdateProjectParams) SetSwapowner(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["swapowner"] = v
}

func (P *UpdateProjectParams) GetSwapowner() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["swapowner"].(bool)
	return value, ok
}

func (P *UpdateProjectParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *UpdateProjectParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateProjectParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewUpdateProjectParams(id string) *UpdateProjectParams {
	P := &UpdateProjectParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a project
func (s *ProjectService) UpdateProject(p *UpdateProjectParams) (*UpdateProjectResponse, error) {
	resp, err := s.cs.newRequest("updateProject", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateProjectResponse
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

type UpdateProjectResponse struct {
	Cpuavailable              string              `json:"cpuavailable"`
	Cpulimit                  string              `json:"cpulimit"`
	Cputotal                  int64               `json:"cputotal"`
	Created                   string              `json:"created"`
	Displaytext               string              `json:"displaytext"`
	Domain                    string              `json:"domain"`
	Domainid                  string              `json:"domainid"`
	Icon                      string              `json:"icon"`
	Id                        string              `json:"id"`
	Ipavailable               string              `json:"ipavailable"`
	Iplimit                   string              `json:"iplimit"`
	Iptotal                   int64               `json:"iptotal"`
	JobID                     string              `json:"jobid"`
	Jobstatus                 int                 `json:"jobstatus"`
	Memoryavailable           string              `json:"memoryavailable"`
	Memorylimit               string              `json:"memorylimit"`
	Memorytotal               int64               `json:"memorytotal"`
	Name                      string              `json:"name"`
	Networkavailable          string              `json:"networkavailable"`
	Networklimit              string              `json:"networklimit"`
	Networktotal              int64               `json:"networktotal"`
	Owner                     []map[string]string `json:"owner"`
	Primarystorageavailable   string              `json:"primarystorageavailable"`
	Primarystoragelimit       string              `json:"primarystoragelimit"`
	Primarystoragetotal       int64               `json:"primarystoragetotal"`
	Projectaccountname        string              `json:"projectaccountname"`
	Secondarystorageavailable string              `json:"secondarystorageavailable"`
	Secondarystoragelimit     string              `json:"secondarystoragelimit"`
	Secondarystoragetotal     float64             `json:"secondarystoragetotal"`
	Snapshotavailable         string              `json:"snapshotavailable"`
	Snapshotlimit             string              `json:"snapshotlimit"`
	Snapshottotal             int64               `json:"snapshottotal"`
	State                     string              `json:"state"`
	Tags                      []Tags              `json:"tags"`
	Templateavailable         string              `json:"templateavailable"`
	Templatelimit             string              `json:"templatelimit"`
	Templatetotal             int64               `json:"templatetotal"`
	Vmavailable               string              `json:"vmavailable"`
	Vmlimit                   string              `json:"vmlimit"`
	Vmrunning                 int                 `json:"vmrunning"`
	Vmstopped                 int                 `json:"vmstopped"`
	Vmtotal                   int64               `json:"vmtotal"`
	Volumeavailable           string              `json:"volumeavailable"`
	Volumelimit               string              `json:"volumelimit"`
	Volumetotal               int64               `json:"volumetotal"`
	Vpcavailable              string              `json:"vpcavailable"`
	Vpclimit                  string              `json:"vpclimit"`
	Vpctotal                  int64               `json:"vpctotal"`
}

type UpdateProjectInvitationParams struct {
	P map[string]interface{}
}

func (P *UpdateProjectInvitationParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["accept"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("accept", vv)
	}
	if v, found := P.P["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["token"]; found {
		u.Set("token", v.(string))
	}
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (P *UpdateProjectInvitationParams) SetAccept(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["accept"] = v
}

func (P *UpdateProjectInvitationParams) GetAccept() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["accept"].(bool)
	return value, ok
}

func (P *UpdateProjectInvitationParams) SetAccount(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["account"] = v
}

func (P *UpdateProjectInvitationParams) GetAccount() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["account"].(string)
	return value, ok
}

func (P *UpdateProjectInvitationParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *UpdateProjectInvitationParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *UpdateProjectInvitationParams) SetToken(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["token"] = v
}

func (P *UpdateProjectInvitationParams) GetToken() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["token"].(string)
	return value, ok
}

func (P *UpdateProjectInvitationParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *UpdateProjectInvitationParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateProjectInvitationParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewUpdateProjectInvitationParams(projectid string) *UpdateProjectInvitationParams {
	P := &UpdateProjectInvitationParams{}
	P.P = make(map[string]interface{})
	P.P["projectid"] = projectid
	return P
}

// Accepts or declines project invitation
func (s *ProjectService) UpdateProjectInvitation(p *UpdateProjectInvitationParams) (*UpdateProjectInvitationResponse, error) {
	resp, err := s.cs.newRequest("updateProjectInvitation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateProjectInvitationResponse
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

type UpdateProjectInvitationResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListProjectRolePermissionsParams struct {
	P map[string]interface{}
}

func (P *ListProjectRolePermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	return u
}

func (P *ListProjectRolePermissionsParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *ListProjectRolePermissionsParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *ListProjectRolePermissionsParams) SetProjectroleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectroleid"] = v
}

func (P *ListProjectRolePermissionsParams) GetProjectroleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectroleid"].(string)
	return value, ok
}

// You should always use this function to get a new ListProjectRolePermissionsParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewListProjectRolePermissionsParams(projectid string) *ListProjectRolePermissionsParams {
	P := &ListProjectRolePermissionsParams{}
	P.P = make(map[string]interface{})
	P.P["projectid"] = projectid
	return P
}

// Lists a project's project role permissions
func (s *ProjectService) ListProjectRolePermissions(p *ListProjectRolePermissionsParams) (*ListProjectRolePermissionsResponse, error) {
	resp, err := s.cs.newRequest("listProjectRolePermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListProjectRolePermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListProjectRolePermissionsResponse struct {
	Count                  int                      `json:"count"`
	ProjectRolePermissions []*ProjectRolePermission `json:"projectrolepermission"`
}

type ProjectRolePermission struct {
	Description     string `json:"description"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Permission      string `json:"permission"`
	Projectid       string `json:"projectid"`
	Projectroleid   string `json:"projectroleid"`
	Projectrolename string `json:"projectrolename"`
	Rule            string `json:"rule"`
}

type CreateProjectRolePermissionParams struct {
	P map[string]interface{}
}

func (P *CreateProjectRolePermissionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["permission"]; found {
		u.Set("permission", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	if v, found := P.P["rule"]; found {
		u.Set("rule", v.(string))
	}
	return u
}

func (P *CreateProjectRolePermissionParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateProjectRolePermissionParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateProjectRolePermissionParams) SetPermission(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["permission"] = v
}

func (P *CreateProjectRolePermissionParams) GetPermission() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["permission"].(string)
	return value, ok
}

func (P *CreateProjectRolePermissionParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *CreateProjectRolePermissionParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *CreateProjectRolePermissionParams) SetProjectroleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectroleid"] = v
}

func (P *CreateProjectRolePermissionParams) GetProjectroleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectroleid"].(string)
	return value, ok
}

func (P *CreateProjectRolePermissionParams) SetRule(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["rule"] = v
}

func (P *CreateProjectRolePermissionParams) GetRule() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["rule"].(string)
	return value, ok
}

// You should always use this function to get a new CreateProjectRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewCreateProjectRolePermissionParams(permission string, projectid string, projectroleid string, rule string) *CreateProjectRolePermissionParams {
	P := &CreateProjectRolePermissionParams{}
	P.P = make(map[string]interface{})
	P.P["permission"] = permission
	P.P["projectid"] = projectid
	P.P["projectroleid"] = projectroleid
	P.P["rule"] = rule
	return P
}

// Adds API permissions to a project role
func (s *ProjectService) CreateProjectRolePermission(p *CreateProjectRolePermissionParams) (*CreateProjectRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("createProjectRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateProjectRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateProjectRolePermissionResponse struct {
	Description     string `json:"description"`
	Id              string `json:"id"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Permission      string `json:"permission"`
	Projectid       string `json:"projectid"`
	Projectroleid   string `json:"projectroleid"`
	Projectrolename string `json:"projectrolename"`
	Rule            string `json:"rule"`
}

type UpdateProjectRolePermissionParams struct {
	P map[string]interface{}
}

func (P *UpdateProjectRolePermissionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["permission"]; found {
		u.Set("permission", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := P.P["projectroleid"]; found {
		u.Set("projectroleid", v.(string))
	}
	if v, found := P.P["projectrolepermissionid"]; found {
		u.Set("projectrolepermissionid", v.(string))
	}
	if v, found := P.P["ruleorder"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ruleorder", vv)
	}
	return u
}

func (P *UpdateProjectRolePermissionParams) SetPermission(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["permission"] = v
}

func (P *UpdateProjectRolePermissionParams) GetPermission() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["permission"].(string)
	return value, ok
}

func (P *UpdateProjectRolePermissionParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *UpdateProjectRolePermissionParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

func (P *UpdateProjectRolePermissionParams) SetProjectroleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectroleid"] = v
}

func (P *UpdateProjectRolePermissionParams) GetProjectroleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectroleid"].(string)
	return value, ok
}

func (P *UpdateProjectRolePermissionParams) SetProjectrolepermissionid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectrolepermissionid"] = v
}

func (P *UpdateProjectRolePermissionParams) GetProjectrolepermissionid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectrolepermissionid"].(string)
	return value, ok
}

func (P *UpdateProjectRolePermissionParams) SetRuleorder(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ruleorder"] = v
}

func (P *UpdateProjectRolePermissionParams) GetRuleorder() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ruleorder"].([]string)
	return value, ok
}

// You should always use this function to get a new UpdateProjectRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewUpdateProjectRolePermissionParams(projectid string, projectroleid string) *UpdateProjectRolePermissionParams {
	P := &UpdateProjectRolePermissionParams{}
	P.P = make(map[string]interface{})
	P.P["projectid"] = projectid
	P.P["projectroleid"] = projectroleid
	return P
}

// Updates a project role permission and/or order
func (s *ProjectService) UpdateProjectRolePermission(p *UpdateProjectRolePermissionParams) (*UpdateProjectRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("updateProjectRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateProjectRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateProjectRolePermissionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateProjectRolePermissionResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateProjectRolePermissionResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteProjectRolePermissionParams struct {
	P map[string]interface{}
}

func (P *DeleteProjectRolePermissionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (P *DeleteProjectRolePermissionParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteProjectRolePermissionParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *DeleteProjectRolePermissionParams) SetProjectid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["projectid"] = v
}

func (P *DeleteProjectRolePermissionParams) GetProjectid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["projectid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteProjectRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *ProjectService) NewDeleteProjectRolePermissionParams(id string, projectid string) *DeleteProjectRolePermissionParams {
	P := &DeleteProjectRolePermissionParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	P.P["projectid"] = projectid
	return P
}

// Deletes a project role permission in the project
func (s *ProjectService) DeleteProjectRolePermission(p *DeleteProjectRolePermissionParams) (*DeleteProjectRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("deleteProjectRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteProjectRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteProjectRolePermissionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteProjectRolePermissionResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteProjectRolePermissionResponse
	return json.Unmarshal(b, (*alias)(r))
}
