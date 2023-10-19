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

// checks if the SharePointIdentitySet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharePointIdentitySet{}

// SharePointIdentitySet This resource is used to represent a set of identities associated with various events for an item, such as created by or last modified by.
type SharePointIdentitySet struct {
	User  *Identity `json:"user,omitempty"`
	Group *Identity `json:"group,omitempty"`
}

// NewSharePointIdentitySet instantiates a new SharePointIdentitySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharePointIdentitySet() *SharePointIdentitySet {
	this := SharePointIdentitySet{}
	return &this
}

// NewSharePointIdentitySetWithDefaults instantiates a new SharePointIdentitySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharePointIdentitySetWithDefaults() *SharePointIdentitySet {
	this := SharePointIdentitySet{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SharePointIdentitySet) GetUser() Identity {
	if o == nil || IsNil(o.User) {
		var ret Identity
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharePointIdentitySet) GetUserOk() (*Identity, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SharePointIdentitySet) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given Identity and assigns it to the User field.
func (o *SharePointIdentitySet) SetUser(v Identity) {
	o.User = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *SharePointIdentitySet) GetGroup() Identity {
	if o == nil || IsNil(o.Group) {
		var ret Identity
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharePointIdentitySet) GetGroupOk() (*Identity, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *SharePointIdentitySet) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given Identity and assigns it to the Group field.
func (o *SharePointIdentitySet) SetGroup(v Identity) {
	o.Group = &v
}

func (o SharePointIdentitySet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharePointIdentitySet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

type NullableSharePointIdentitySet struct {
	value *SharePointIdentitySet
	isSet bool
}

func (v NullableSharePointIdentitySet) Get() *SharePointIdentitySet {
	return v.value
}

func (v *NullableSharePointIdentitySet) Set(val *SharePointIdentitySet) {
	v.value = val
	v.isSet = true
}

func (v NullableSharePointIdentitySet) IsSet() bool {
	return v.isSet
}

func (v *NullableSharePointIdentitySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharePointIdentitySet(val *SharePointIdentitySet) *NullableSharePointIdentitySet {
	return &NullableSharePointIdentitySet{value: val, isSet: true}
}

func (v NullableSharePointIdentitySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharePointIdentitySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
