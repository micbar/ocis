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

// checks if the SharingLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharingLink{}

// SharingLink The `SharingLink` resource groups link-related data items into a single structure.  If a `permission` resource has a non-null `sharingLink` facet, the permission represents a sharing link (as opposed to permissions granted to a person or group).
type SharingLink struct {
	Type *SharingLinkType `json:"type,omitempty"`
	// If `true` then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item.
	PreventsDownload *bool `json:"preventsDownload,omitempty"`
	// A URL that opens the item in the browser on the website.
	WebUrl *string `json:"webUrl,omitempty"`
	// Provides a user-visible display name of the link. Optional. Libregraph only.
	LibreGraphDisplayName *string `json:"@libre.graph.displayName,omitempty"`
}

// NewSharingLink instantiates a new SharingLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharingLink() *SharingLink {
	this := SharingLink{}
	return &this
}

// NewSharingLinkWithDefaults instantiates a new SharingLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharingLinkWithDefaults() *SharingLink {
	this := SharingLink{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SharingLink) GetType() SharingLinkType {
	if o == nil || IsNil(o.Type) {
		var ret SharingLinkType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingLink) GetTypeOk() (*SharingLinkType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SharingLink) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SharingLinkType and assigns it to the Type field.
func (o *SharingLink) SetType(v SharingLinkType) {
	o.Type = &v
}

// GetPreventsDownload returns the PreventsDownload field value if set, zero value otherwise.
func (o *SharingLink) GetPreventsDownload() bool {
	if o == nil || IsNil(o.PreventsDownload) {
		var ret bool
		return ret
	}
	return *o.PreventsDownload
}

// GetPreventsDownloadOk returns a tuple with the PreventsDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingLink) GetPreventsDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.PreventsDownload) {
		return nil, false
	}
	return o.PreventsDownload, true
}

// HasPreventsDownload returns a boolean if a field has been set.
func (o *SharingLink) HasPreventsDownload() bool {
	if o != nil && !IsNil(o.PreventsDownload) {
		return true
	}

	return false
}

// SetPreventsDownload gets a reference to the given bool and assigns it to the PreventsDownload field.
func (o *SharingLink) SetPreventsDownload(v bool) {
	o.PreventsDownload = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *SharingLink) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl) {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingLink) GetWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebUrl) {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *SharingLink) HasWebUrl() bool {
	if o != nil && !IsNil(o.WebUrl) {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *SharingLink) SetWebUrl(v string) {
	o.WebUrl = &v
}

// GetLibreGraphDisplayName returns the LibreGraphDisplayName field value if set, zero value otherwise.
func (o *SharingLink) GetLibreGraphDisplayName() string {
	if o == nil || IsNil(o.LibreGraphDisplayName) {
		var ret string
		return ret
	}
	return *o.LibreGraphDisplayName
}

// GetLibreGraphDisplayNameOk returns a tuple with the LibreGraphDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharingLink) GetLibreGraphDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.LibreGraphDisplayName) {
		return nil, false
	}
	return o.LibreGraphDisplayName, true
}

// HasLibreGraphDisplayName returns a boolean if a field has been set.
func (o *SharingLink) HasLibreGraphDisplayName() bool {
	if o != nil && !IsNil(o.LibreGraphDisplayName) {
		return true
	}

	return false
}

// SetLibreGraphDisplayName gets a reference to the given string and assigns it to the LibreGraphDisplayName field.
func (o *SharingLink) SetLibreGraphDisplayName(v string) {
	o.LibreGraphDisplayName = &v
}

func (o SharingLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharingLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.PreventsDownload) {
		toSerialize["preventsDownload"] = o.PreventsDownload
	}
	if !IsNil(o.WebUrl) {
		toSerialize["webUrl"] = o.WebUrl
	}
	if !IsNil(o.LibreGraphDisplayName) {
		toSerialize["@libre.graph.displayName"] = o.LibreGraphDisplayName
	}
	return toSerialize, nil
}

type NullableSharingLink struct {
	value *SharingLink
	isSet bool
}

func (v NullableSharingLink) Get() *SharingLink {
	return v.value
}

func (v *NullableSharingLink) Set(val *SharingLink) {
	v.value = val
	v.isSet = true
}

func (v NullableSharingLink) IsSet() bool {
	return v.isSet
}

func (v *NullableSharingLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharingLink(val *SharingLink) *NullableSharingLink {
	return &NullableSharingLink{value: val, isSet: true}
}

func (v NullableSharingLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharingLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
