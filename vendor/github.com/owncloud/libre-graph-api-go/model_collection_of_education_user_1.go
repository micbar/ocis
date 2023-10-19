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

// checks if the CollectionOfEducationUser1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfEducationUser1{}

// CollectionOfEducationUser1 struct for CollectionOfEducationUser1
type CollectionOfEducationUser1 struct {
	Value []EducationClass `json:"value,omitempty"`
}

// NewCollectionOfEducationUser1 instantiates a new CollectionOfEducationUser1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfEducationUser1() *CollectionOfEducationUser1 {
	this := CollectionOfEducationUser1{}
	return &this
}

// NewCollectionOfEducationUser1WithDefaults instantiates a new CollectionOfEducationUser1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfEducationUser1WithDefaults() *CollectionOfEducationUser1 {
	this := CollectionOfEducationUser1{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfEducationUser1) GetValue() []EducationClass {
	if o == nil || IsNil(o.Value) {
		var ret []EducationClass
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfEducationUser1) GetValueOk() ([]EducationClass, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfEducationUser1) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []EducationClass and assigns it to the Value field.
func (o *CollectionOfEducationUser1) SetValue(v []EducationClass) {
	o.Value = v
}

func (o CollectionOfEducationUser1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfEducationUser1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCollectionOfEducationUser1 struct {
	value *CollectionOfEducationUser1
	isSet bool
}

func (v NullableCollectionOfEducationUser1) Get() *CollectionOfEducationUser1 {
	return v.value
}

func (v *NullableCollectionOfEducationUser1) Set(val *CollectionOfEducationUser1) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfEducationUser1) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfEducationUser1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfEducationUser1(val *CollectionOfEducationUser1) *NullableCollectionOfEducationUser1 {
	return &NullableCollectionOfEducationUser1{value: val, isSet: true}
}

func (v NullableCollectionOfEducationUser1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfEducationUser1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
