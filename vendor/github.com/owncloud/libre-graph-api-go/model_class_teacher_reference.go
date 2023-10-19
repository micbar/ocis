/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// checks if the ClassTeacherReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClassTeacherReference{}

// ClassTeacherReference struct for ClassTeacherReference
type ClassTeacherReference struct {
	OdataId *string `json:"@odata.id,omitempty"`
}

// NewClassTeacherReference instantiates a new ClassTeacherReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassTeacherReference() *ClassTeacherReference {
	this := ClassTeacherReference{}
	return &this
}

// NewClassTeacherReferenceWithDefaults instantiates a new ClassTeacherReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassTeacherReferenceWithDefaults() *ClassTeacherReference {
	this := ClassTeacherReference{}
	return &this
}

// GetOdataId returns the OdataId field value if set, zero value otherwise.
func (o *ClassTeacherReference) GetOdataId() string {
	if o == nil || IsNil(o.OdataId) {
		var ret string
		return ret
	}
	return *o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassTeacherReference) GetOdataIdOk() (*string, bool) {
	if o == nil || IsNil(o.OdataId) {
		return nil, false
	}
	return o.OdataId, true
}

// HasOdataId returns a boolean if a field has been set.
func (o *ClassTeacherReference) HasOdataId() bool {
	if o != nil && !IsNil(o.OdataId) {
		return true
	}

	return false
}

// SetOdataId gets a reference to the given string and assigns it to the OdataId field.
func (o *ClassTeacherReference) SetOdataId(v string) {
	o.OdataId = &v
}

func (o ClassTeacherReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClassTeacherReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OdataId) {
		toSerialize["@odata.id"] = o.OdataId
	}
	return toSerialize, nil
}

type NullableClassTeacherReference struct {
	value *ClassTeacherReference
	isSet bool
}

func (v NullableClassTeacherReference) Get() *ClassTeacherReference {
	return v.value
}

func (v *NullableClassTeacherReference) Set(val *ClassTeacherReference) {
	v.value = val
	v.isSet = true
}

func (v NullableClassTeacherReference) IsSet() bool {
	return v.isSet
}

func (v *NullableClassTeacherReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassTeacherReference(val *ClassTeacherReference) *NullableClassTeacherReference {
	return &NullableClassTeacherReference{value: val, isSet: true}
}

func (v NullableClassTeacherReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassTeacherReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
