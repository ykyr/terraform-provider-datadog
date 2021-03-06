/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// RoleCreatePayload Create a role.
type RoleCreatePayload struct {
	Data *RoleCreateData `json:"data,omitempty"`
}

// NewRoleCreatePayload instantiates a new RoleCreatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleCreatePayload() *RoleCreatePayload {
	this := RoleCreatePayload{}
	return &this
}

// NewRoleCreatePayloadWithDefaults instantiates a new RoleCreatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleCreatePayloadWithDefaults() *RoleCreatePayload {
	this := RoleCreatePayload{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RoleCreatePayload) GetData() RoleCreateData {
	if o == nil || o.Data == nil {
		var ret RoleCreateData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCreatePayload) GetDataOk() (*RoleCreateData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RoleCreatePayload) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RoleCreateData and assigns it to the Data field.
func (o *RoleCreatePayload) SetData(v RoleCreateData) {
	o.Data = &v
}

func (o RoleCreatePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableRoleCreatePayload struct {
	value *RoleCreatePayload
	isSet bool
}

func (v NullableRoleCreatePayload) Get() *RoleCreatePayload {
	return v.value
}

func (v *NullableRoleCreatePayload) Set(val *RoleCreatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleCreatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleCreatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleCreatePayload(val *RoleCreatePayload) *NullableRoleCreatePayload {
	return &NullableRoleCreatePayload{value: val, isSet: true}
}

func (v NullableRoleCreatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleCreatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
