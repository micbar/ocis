// Code generated by protoc-gen-microweb. DO NOT EDIT.
// source: v0.proto

package v0

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
)

// ValueWithIdentifierJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ValueWithIdentifier. This struct is safe to replace or modify but
// should not be done so concurrently.
var ValueWithIdentifierJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ValueWithIdentifier) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ValueWithIdentifierJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ValueWithIdentifier)(nil)

// ValueWithIdentifierJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ValueWithIdentifier. This struct is safe to replace or modify but
// should not be done so concurrently.
var ValueWithIdentifierJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ValueWithIdentifier) UnmarshalJSON(b []byte) error {
	return ValueWithIdentifierJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ValueWithIdentifier)(nil)

// IdentifierJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Identifier. This struct is safe to replace or modify but
// should not be done so concurrently.
var IdentifierJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Identifier) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := IdentifierJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Identifier)(nil)

// IdentifierJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Identifier. This struct is safe to replace or modify but
// should not be done so concurrently.
var IdentifierJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Identifier) UnmarshalJSON(b []byte) error {
	return IdentifierJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Identifier)(nil)

// UserRoleAssignmentJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of UserRoleAssignment. This struct is safe to replace or modify but
// should not be done so concurrently.
var UserRoleAssignmentJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *UserRoleAssignment) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := UserRoleAssignmentJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*UserRoleAssignment)(nil)

// UserRoleAssignmentJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of UserRoleAssignment. This struct is safe to replace or modify but
// should not be done so concurrently.
var UserRoleAssignmentJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *UserRoleAssignment) UnmarshalJSON(b []byte) error {
	return UserRoleAssignmentJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*UserRoleAssignment)(nil)

// UserRoleAssignmentFilterJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of UserRoleAssignmentFilter. This struct is safe to replace or modify but
// should not be done so concurrently.
var UserRoleAssignmentFilterJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *UserRoleAssignmentFilter) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := UserRoleAssignmentFilterJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*UserRoleAssignmentFilter)(nil)

// UserRoleAssignmentFilterJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of UserRoleAssignmentFilter. This struct is safe to replace or modify but
// should not be done so concurrently.
var UserRoleAssignmentFilterJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *UserRoleAssignmentFilter) UnmarshalJSON(b []byte) error {
	return UserRoleAssignmentFilterJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*UserRoleAssignmentFilter)(nil)

// ResourceJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Resource. This struct is safe to replace or modify but
// should not be done so concurrently.
var ResourceJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Resource) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ResourceJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Resource)(nil)

// ResourceJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Resource. This struct is safe to replace or modify but
// should not be done so concurrently.
var ResourceJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Resource) UnmarshalJSON(b []byte) error {
	return ResourceJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Resource)(nil)

// BundleJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Bundle. This struct is safe to replace or modify but
// should not be done so concurrently.
var BundleJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Bundle) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := BundleJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Bundle)(nil)

// BundleJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Bundle. This struct is safe to replace or modify but
// should not be done so concurrently.
var BundleJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Bundle) UnmarshalJSON(b []byte) error {
	return BundleJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Bundle)(nil)

// SettingJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Setting. This struct is safe to replace or modify but
// should not be done so concurrently.
var SettingJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Setting) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SettingJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Setting)(nil)

// SettingJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Setting. This struct is safe to replace or modify but
// should not be done so concurrently.
var SettingJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Setting) UnmarshalJSON(b []byte) error {
	return SettingJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Setting)(nil)

// IntJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Int. This struct is safe to replace or modify but
// should not be done so concurrently.
var IntJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Int) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := IntJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Int)(nil)

// IntJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Int. This struct is safe to replace or modify but
// should not be done so concurrently.
var IntJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Int) UnmarshalJSON(b []byte) error {
	return IntJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Int)(nil)

// StringJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of String. This struct is safe to replace or modify but
// should not be done so concurrently.
var StringJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *String) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := StringJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*String)(nil)

// StringJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of String. This struct is safe to replace or modify but
// should not be done so concurrently.
var StringJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *String) UnmarshalJSON(b []byte) error {
	return StringJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*String)(nil)

// BoolJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Bool. This struct is safe to replace or modify but
// should not be done so concurrently.
var BoolJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Bool) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := BoolJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Bool)(nil)

// BoolJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Bool. This struct is safe to replace or modify but
// should not be done so concurrently.
var BoolJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Bool) UnmarshalJSON(b []byte) error {
	return BoolJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Bool)(nil)

// SingleChoiceListJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SingleChoiceList. This struct is safe to replace or modify but
// should not be done so concurrently.
var SingleChoiceListJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SingleChoiceList) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SingleChoiceListJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SingleChoiceList)(nil)

// SingleChoiceListJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SingleChoiceList. This struct is safe to replace or modify but
// should not be done so concurrently.
var SingleChoiceListJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SingleChoiceList) UnmarshalJSON(b []byte) error {
	return SingleChoiceListJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SingleChoiceList)(nil)

// MultiChoiceListJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of MultiChoiceList. This struct is safe to replace or modify but
// should not be done so concurrently.
var MultiChoiceListJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *MultiChoiceList) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := MultiChoiceListJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*MultiChoiceList)(nil)

// MultiChoiceListJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of MultiChoiceList. This struct is safe to replace or modify but
// should not be done so concurrently.
var MultiChoiceListJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *MultiChoiceList) UnmarshalJSON(b []byte) error {
	return MultiChoiceListJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*MultiChoiceList)(nil)

// ListOptionJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListOption. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListOptionJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListOption) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListOptionJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListOption)(nil)

// ListOptionJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListOption. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListOptionJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListOption) UnmarshalJSON(b []byte) error {
	return ListOptionJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListOption)(nil)

// PermissionJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Permission. This struct is safe to replace or modify but
// should not be done so concurrently.
var PermissionJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Permission) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := PermissionJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Permission)(nil)

// PermissionJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Permission. This struct is safe to replace or modify but
// should not be done so concurrently.
var PermissionJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Permission) UnmarshalJSON(b []byte) error {
	return PermissionJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Permission)(nil)

// ValueJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Value. This struct is safe to replace or modify but
// should not be done so concurrently.
var ValueJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Value) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ValueJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Value)(nil)

// ValueJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Value. This struct is safe to replace or modify but
// should not be done so concurrently.
var ValueJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Value) UnmarshalJSON(b []byte) error {
	return ValueJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Value)(nil)

// ListValueJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListValue. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListValueJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListValue) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListValueJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListValue)(nil)

// ListValueJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListValue. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListValueJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListValue) UnmarshalJSON(b []byte) error {
	return ListValueJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListValue)(nil)

// ListOptionValueJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListOptionValue. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListOptionValueJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListOptionValue) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListOptionValueJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListOptionValue)(nil)

// ListOptionValueJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListOptionValue. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListOptionValueJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListOptionValue) UnmarshalJSON(b []byte) error {
	return ListOptionValueJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListOptionValue)(nil)
