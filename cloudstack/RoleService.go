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

type RoleServiceIface interface {
	CreateRole(P *CreateRoleParams) (*CreateRoleResponse, error)
	NewCreateRoleParams(name string) *CreateRoleParams
	CreateRolePermission(P *CreateRolePermissionParams) (*CreateRolePermissionResponse, error)
	NewCreateRolePermissionParams(permission string, roleid string, rule string) *CreateRolePermissionParams
	DeleteRole(P *DeleteRoleParams) (*DeleteRoleResponse, error)
	NewDeleteRoleParams(id string) *DeleteRoleParams
	DeleteRolePermission(P *DeleteRolePermissionParams) (*DeleteRolePermissionResponse, error)
	NewDeleteRolePermissionParams(id string) *DeleteRolePermissionParams
	ListRolePermissions(P *ListRolePermissionsParams) (*ListRolePermissionsResponse, error)
	NewListRolePermissionsParams() *ListRolePermissionsParams
	ListRoles(P *ListRolesParams) (*ListRolesResponse, error)
	NewListRolesParams() *ListRolesParams
	GetRoleID(name string, opts ...OptionFunc) (string, int, error)
	GetRoleByName(name string, opts ...OptionFunc) (*Role, int, error)
	GetRoleByID(id string, opts ...OptionFunc) (*Role, int, error)
	UpdateRole(P *UpdateRoleParams) (*UpdateRoleResponse, error)
	NewUpdateRoleParams(id string) *UpdateRoleParams
	UpdateRolePermission(P *UpdateRolePermissionParams) (*UpdateRolePermissionResponse, error)
	NewUpdateRolePermissionParams(roleid string) *UpdateRolePermissionParams
}

type CreateRoleParams struct {
	P map[string]interface{}
}

func (P *CreateRoleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (P *CreateRoleParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateRoleParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateRoleParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *CreateRoleParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *CreateRoleParams) SetRoleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roleid"] = v
}

func (P *CreateRoleParams) GetRoleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roleid"].(string)
	return value, ok
}

func (P *CreateRoleParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *CreateRoleParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new CreateRoleParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewCreateRoleParams(name string) *CreateRoleParams {
	P := &CreateRoleParams{}
	P.P = make(map[string]interface{})
	P.P["name"] = name
	return P
}

// Creates a role
func (s *RoleService) CreateRole(p *CreateRoleParams) (*CreateRoleResponse, error) {
	resp, err := s.cs.newRequest("createRole", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateRoleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateRoleResponse struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Isdefault   bool   `json:"isdefault"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type CreateRolePermissionParams struct {
	P map[string]interface{}
}

func (P *CreateRolePermissionParams) toURLValues() url.Values {
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
	if v, found := P.P["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := P.P["rule"]; found {
		u.Set("rule", v.(string))
	}
	return u
}

func (P *CreateRolePermissionParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *CreateRolePermissionParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *CreateRolePermissionParams) SetPermission(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["permission"] = v
}

func (P *CreateRolePermissionParams) GetPermission() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["permission"].(string)
	return value, ok
}

func (P *CreateRolePermissionParams) SetRoleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roleid"] = v
}

func (P *CreateRolePermissionParams) GetRoleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roleid"].(string)
	return value, ok
}

func (P *CreateRolePermissionParams) SetRule(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["rule"] = v
}

func (P *CreateRolePermissionParams) GetRule() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["rule"].(string)
	return value, ok
}

// You should always use this function to get a new CreateRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewCreateRolePermissionParams(permission string, roleid string, rule string) *CreateRolePermissionParams {
	P := &CreateRolePermissionParams{}
	P.P = make(map[string]interface{})
	P.P["permission"] = permission
	P.P["roleid"] = roleid
	P.P["rule"] = rule
	return P
}

// Adds an API permission to a role
func (s *RoleService) CreateRolePermission(p *CreateRolePermissionParams) (*CreateRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("createRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateRolePermissionResponse struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Permission  string `json:"permission"`
	Roleid      string `json:"roleid"`
	Rolename    string `json:"rolename"`
	Rule        string `json:"rule"`
}

type DeleteRoleParams struct {
	P map[string]interface{}
}

func (P *DeleteRoleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteRoleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteRoleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteRoleParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewDeleteRoleParams(id string) *DeleteRoleParams {
	P := &DeleteRoleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a role
func (s *RoleService) DeleteRole(p *DeleteRoleParams) (*DeleteRoleResponse, error) {
	resp, err := s.cs.newRequest("deleteRole", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteRoleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteRoleResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteRoleResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteRoleResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteRolePermissionParams struct {
	P map[string]interface{}
}

func (P *DeleteRolePermissionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *DeleteRolePermissionParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *DeleteRolePermissionParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewDeleteRolePermissionParams(id string) *DeleteRolePermissionParams {
	P := &DeleteRolePermissionParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Deletes a role permission
func (s *RoleService) DeleteRolePermission(p *DeleteRolePermissionParams) (*DeleteRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("deleteRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteRolePermissionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteRolePermissionResponse) UnmarshalJSON(b []byte) error {
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

	type alias DeleteRolePermissionResponse
	return json.Unmarshal(b, (*alias)(r))
}

type ListRolePermissionsParams struct {
	P map[string]interface{}
}

func (P *ListRolePermissionsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	return u
}

func (P *ListRolePermissionsParams) SetRoleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roleid"] = v
}

func (P *ListRolePermissionsParams) GetRoleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roleid"].(string)
	return value, ok
}

// You should always use this function to get a new ListRolePermissionsParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewListRolePermissionsParams() *ListRolePermissionsParams {
	P := &ListRolePermissionsParams{}
	P.P = make(map[string]interface{})
	return P
}

// Lists role permissions
func (s *RoleService) ListRolePermissions(p *ListRolePermissionsParams) (*ListRolePermissionsResponse, error) {
	resp, err := s.cs.newRequest("listRolePermissions", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRolePermissionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListRolePermissionsResponse struct {
	Count           int               `json:"count"`
	RolePermissions []*RolePermission `json:"rolepermission"`
}

type RolePermission struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Permission  string `json:"permission"`
	Roleid      string `json:"roleid"`
	Rolename    string `json:"rolename"`
	Rule        string `json:"rule"`
}

type ListRolesParams struct {
	P map[string]interface{}
}

func (P *ListRolesParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["keyword"]; found {
		u.Set("keyword", v.(string))
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
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (P *ListRolesParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListRolesParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListRolesParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListRolesParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListRolesParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *ListRolesParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *ListRolesParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListRolesParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListRolesParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListRolesParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListRolesParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *ListRolesParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new ListRolesParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewListRolesParams() *ListRolesParams {
	P := &ListRolesParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RoleService) GetRoleID(name string, opts ...OptionFunc) (string, int, error) {
	P := &ListRolesParams{}
	P.P = make(map[string]interface{})

	P.P["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListRoles(P)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Roles[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Roles {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RoleService) GetRoleByName(name string, opts ...OptionFunc) (*Role, int, error) {
	id, count, err := s.GetRoleID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetRoleByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *RoleService) GetRoleByID(id string, opts ...OptionFunc) (*Role, int, error) {
	P := &ListRolesParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListRoles(P)
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
		return l.Roles[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Role UUID: %s!", id)
}

// Lists dynamic roles in CloudStack
func (s *RoleService) ListRoles(p *ListRolesParams) (*ListRolesResponse, error) {
	resp, err := s.cs.newRequest("listRoles", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRolesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListRolesResponse struct {
	Count int     `json:"count"`
	Roles []*Role `json:"role"`
}

type Role struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Isdefault   bool   `json:"isdefault"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type UpdateRoleParams struct {
	P map[string]interface{}
}

func (P *UpdateRoleParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := P.P["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := P.P["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (P *UpdateRoleParams) SetDescription(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["description"] = v
}

func (P *UpdateRoleParams) GetDescription() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["description"].(string)
	return value, ok
}

func (P *UpdateRoleParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateRoleParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *UpdateRoleParams) SetName(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["name"] = v
}

func (P *UpdateRoleParams) GetName() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["name"].(string)
	return value, ok
}

func (P *UpdateRoleParams) SetType(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["type"] = v
}

func (P *UpdateRoleParams) GetType() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["type"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateRoleParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewUpdateRoleParams(id string) *UpdateRoleParams {
	P := &UpdateRoleParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// Updates a role
func (s *RoleService) UpdateRole(p *UpdateRoleParams) (*UpdateRoleResponse, error) {
	resp, err := s.cs.newRequest("updateRole", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateRoleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateRoleResponse struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Isdefault   bool   `json:"isdefault"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type UpdateRolePermissionParams struct {
	P map[string]interface{}
}

func (P *UpdateRolePermissionParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["permission"]; found {
		u.Set("permission", v.(string))
	}
	if v, found := P.P["roleid"]; found {
		u.Set("roleid", v.(string))
	}
	if v, found := P.P["ruleid"]; found {
		u.Set("ruleid", v.(string))
	}
	if v, found := P.P["ruleorder"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("ruleorder", vv)
	}
	return u
}

func (P *UpdateRolePermissionParams) SetPermission(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["permission"] = v
}

func (P *UpdateRolePermissionParams) GetPermission() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["permission"].(string)
	return value, ok
}

func (P *UpdateRolePermissionParams) SetRoleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["roleid"] = v
}

func (P *UpdateRolePermissionParams) GetRoleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["roleid"].(string)
	return value, ok
}

func (P *UpdateRolePermissionParams) SetRuleid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ruleid"] = v
}

func (P *UpdateRolePermissionParams) GetRuleid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ruleid"].(string)
	return value, ok
}

func (P *UpdateRolePermissionParams) SetRuleorder(v []string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["ruleorder"] = v
}

func (P *UpdateRolePermissionParams) GetRuleorder() ([]string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["ruleorder"].([]string)
	return value, ok
}

// You should always use this function to get a new UpdateRolePermissionParams instance,
// as then you are sure you have configured all required params
func (s *RoleService) NewUpdateRolePermissionParams(roleid string) *UpdateRolePermissionParams {
	P := &UpdateRolePermissionParams{}
	P.P = make(map[string]interface{})
	P.P["roleid"] = roleid
	return P
}

// Updates a role permission order
func (s *RoleService) UpdateRolePermission(p *UpdateRolePermissionParams) (*UpdateRolePermissionResponse, error) {
	resp, err := s.cs.newRequest("updateRolePermission", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateRolePermissionResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateRolePermissionResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *UpdateRolePermissionResponse) UnmarshalJSON(b []byte) error {
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

	type alias UpdateRolePermissionResponse
	return json.Unmarshal(b, (*alias)(r))
}
