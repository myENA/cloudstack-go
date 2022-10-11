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

type AnnotationServiceIface interface {
	AddAnnotation(p *AddAnnotationParams) (*AddAnnotationResponse, error)
	NewAddAnnotationParams() *AddAnnotationParams
	ListAnnotations(p *ListAnnotationsParams) (*ListAnnotationsResponse, error)
	NewListAnnotationsParams() *ListAnnotationsParams
	GetAnnotationByID(id string, opts ...OptionFunc) (*Annotation, int, error)
	RemoveAnnotation(p *RemoveAnnotationParams) (*RemoveAnnotationResponse, error)
	NewRemoveAnnotationParams(id string) *RemoveAnnotationParams
	UpdateAnnotationVisibility(p *UpdateAnnotationVisibilityParams) (*UpdateAnnotationVisibilityResponse, error)
	NewUpdateAnnotationVisibilityParams(adminsonly bool, id string) *UpdateAnnotationVisibilityParams
}

type AddAnnotationParams struct {
	P map[string]interface{}
}

func (P *AddAnnotationParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["adminsonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("adminsonly", vv)
	}
	if v, found := P.P["annotation"]; found {
		u.Set("annotation", v.(string))
	}
	if v, found := P.P["entityid"]; found {
		u.Set("entityid", v.(string))
	}
	if v, found := P.P["entitytype"]; found {
		u.Set("entitytype", v.(string))
	}
	return u
}

func (P *AddAnnotationParams) SetAdminsonly(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["adminsonly"] = v
}

func (P *AddAnnotationParams) GetAdminsonly() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["adminsonly"].(bool)
	return value, ok
}

func (P *AddAnnotationParams) SetAnnotation(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["annotation"] = v
}

func (P *AddAnnotationParams) GetAnnotation() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["annotation"].(string)
	return value, ok
}

func (P *AddAnnotationParams) SetEntityid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["entityid"] = v
}

func (P *AddAnnotationParams) GetEntityid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["entityid"].(string)
	return value, ok
}

func (P *AddAnnotationParams) SetEntitytype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["entitytype"] = v
}

func (P *AddAnnotationParams) GetEntitytype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["entitytype"].(string)
	return value, ok
}

// You should always use this function to get a new AddAnnotationParams instance,
// as then you are sure you have configured all required params
func (s *AnnotationService) NewAddAnnotationParams() *AddAnnotationParams {
	P := &AddAnnotationParams{}
	P.P = make(map[string]interface{})
	return P
}

// add an annotation.
func (s *AnnotationService) AddAnnotation(p *AddAnnotationParams) (*AddAnnotationResponse, error) {
	resp, err := s.cs.newRequest("addAnnotation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r AddAnnotationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddAnnotationResponse struct {
	Adminsonly bool   `json:"adminsonly"`
	Annotation string `json:"annotation"`
	Created    string `json:"created"`
	Entityid   string `json:"entityid"`
	Entityname string `json:"entityname"`
	Entitytype string `json:"entitytype"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Removed    string `json:"removed"`
	Userid     string `json:"userid"`
	Username   string `json:"username"`
}

type ListAnnotationsParams struct {
	P map[string]interface{}
}

func (P *ListAnnotationsParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["annotationfilter"]; found {
		u.Set("annotationfilter", v.(string))
	}
	if v, found := P.P["entityid"]; found {
		u.Set("entityid", v.(string))
	}
	if v, found := P.P["entitytype"]; found {
		u.Set("entitytype", v.(string))
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
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
	if v, found := P.P["userid"]; found {
		u.Set("userid", v.(string))
	}
	return u
}

func (P *ListAnnotationsParams) SetAnnotationfilter(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["annotationfilter"] = v
}

func (P *ListAnnotationsParams) GetAnnotationfilter() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["annotationfilter"].(string)
	return value, ok
}

func (P *ListAnnotationsParams) SetEntityid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["entityid"] = v
}

func (P *ListAnnotationsParams) GetEntityid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["entityid"].(string)
	return value, ok
}

func (P *ListAnnotationsParams) SetEntitytype(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["entitytype"] = v
}

func (P *ListAnnotationsParams) GetEntitytype() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["entitytype"].(string)
	return value, ok
}

func (P *ListAnnotationsParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *ListAnnotationsParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

func (P *ListAnnotationsParams) SetKeyword(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["keyword"] = v
}

func (P *ListAnnotationsParams) GetKeyword() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["keyword"].(string)
	return value, ok
}

func (P *ListAnnotationsParams) SetPage(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["page"] = v
}

func (P *ListAnnotationsParams) GetPage() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["page"].(int)
	return value, ok
}

func (P *ListAnnotationsParams) SetPagesize(v int) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["pagesize"] = v
}

func (P *ListAnnotationsParams) GetPagesize() (int, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["pagesize"].(int)
	return value, ok
}

func (P *ListAnnotationsParams) SetUserid(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["userid"] = v
}

func (P *ListAnnotationsParams) GetUserid() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["userid"].(string)
	return value, ok
}

// You should always use this function to get a new ListAnnotationsParams instance,
// as then you are sure you have configured all required params
func (s *AnnotationService) NewListAnnotationsParams() *ListAnnotationsParams {
	P := &ListAnnotationsParams{}
	P.P = make(map[string]interface{})
	return P
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *AnnotationService) GetAnnotationByID(id string, opts ...OptionFunc) (*Annotation, int, error) {
	P := &ListAnnotationsParams{}
	P.P = make(map[string]interface{})

	P.P["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, P); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListAnnotations(P)
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
		return l.Annotations[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Annotation UUID: %s!", id)
}

// Lists annotations.
func (s *AnnotationService) ListAnnotations(p *ListAnnotationsParams) (*ListAnnotationsResponse, error) {
	resp, err := s.cs.newRequest("listAnnotations", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListAnnotationsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListAnnotationsResponse struct {
	Count       int           `json:"count"`
	Annotations []*Annotation `json:"annotation"`
}

type Annotation struct {
	Adminsonly bool   `json:"adminsonly"`
	Annotation string `json:"annotation"`
	Created    string `json:"created"`
	Entityid   string `json:"entityid"`
	Entityname string `json:"entityname"`
	Entitytype string `json:"entitytype"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Removed    string `json:"removed"`
	Userid     string `json:"userid"`
	Username   string `json:"username"`
}

type RemoveAnnotationParams struct {
	P map[string]interface{}
}

func (P *RemoveAnnotationParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *RemoveAnnotationParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *RemoveAnnotationParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveAnnotationParams instance,
// as then you are sure you have configured all required params
func (s *AnnotationService) NewRemoveAnnotationParams(id string) *RemoveAnnotationParams {
	P := &RemoveAnnotationParams{}
	P.P = make(map[string]interface{})
	P.P["id"] = id
	return P
}

// remove an annotation.
func (s *AnnotationService) RemoveAnnotation(p *RemoveAnnotationParams) (*RemoveAnnotationResponse, error) {
	resp, err := s.cs.newRequest("removeAnnotation", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r RemoveAnnotationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemoveAnnotationResponse struct {
	Adminsonly bool   `json:"adminsonly"`
	Annotation string `json:"annotation"`
	Created    string `json:"created"`
	Entityid   string `json:"entityid"`
	Entityname string `json:"entityname"`
	Entitytype string `json:"entitytype"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Removed    string `json:"removed"`
	Userid     string `json:"userid"`
	Username   string `json:"username"`
}

type UpdateAnnotationVisibilityParams struct {
	P map[string]interface{}
}

func (P *UpdateAnnotationVisibilityParams) toURLValues() url.Values {
	u := url.Values{}
	if P.P == nil {
		return u
	}
	if v, found := P.P["adminsonly"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("adminsonly", vv)
	}
	if v, found := P.P["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (P *UpdateAnnotationVisibilityParams) SetAdminsonly(v bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["adminsonly"] = v
}

func (P *UpdateAnnotationVisibilityParams) GetAdminsonly() (bool, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["adminsonly"].(bool)
	return value, ok
}

func (P *UpdateAnnotationVisibilityParams) SetId(v string) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	P.P["id"] = v
}

func (P *UpdateAnnotationVisibilityParams) GetId() (string, bool) {
	if P.P == nil {
		P.P = make(map[string]interface{})
	}
	value, ok := P.P["id"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateAnnotationVisibilityParams instance,
// as then you are sure you have configured all required params
func (s *AnnotationService) NewUpdateAnnotationVisibilityParams(adminsonly bool, id string) *UpdateAnnotationVisibilityParams {
	P := &UpdateAnnotationVisibilityParams{}
	P.P = make(map[string]interface{})
	P.P["adminsonly"] = adminsonly
	P.P["id"] = id
	return P
}

// update an annotation visibility.
func (s *AnnotationService) UpdateAnnotationVisibility(p *UpdateAnnotationVisibilityParams) (*UpdateAnnotationVisibilityResponse, error) {
	resp, err := s.cs.newRequest("updateAnnotationVisibility", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateAnnotationVisibilityResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateAnnotationVisibilityResponse struct {
	Adminsonly bool   `json:"adminsonly"`
	Annotation string `json:"annotation"`
	Created    string `json:"created"`
	Entityid   string `json:"entityid"`
	Entityname string `json:"entityname"`
	Entitytype string `json:"entitytype"`
	Id         string `json:"id"`
	JobID      string `json:"jobid"`
	Jobstatus  int    `json:"jobstatus"`
	Removed    string `json:"removed"`
	Userid     string `json:"userid"`
	Username   string `json:"username"`
}
