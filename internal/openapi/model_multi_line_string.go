/*
ETSI ISG CIM / NGSI-LD API

This OAS file describes the NGSI-LD API defined by the ETSI ISG CIM group. This Cross-domain Context Information Management API allows to provide, consume and subscribe to context information in multiple scenarios and involving multiple stakeholders

API version: latest
Contact: NGSI-LD@etsi.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MultiLineString struct for MultiLineString
type MultiLineString struct {
	Type *string `json:"type,omitempty"`
	Coordinates [][][]float32 `json:"coordinates,omitempty"`
}

// NewMultiLineString instantiates a new MultiLineString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiLineString() *MultiLineString {
	this := MultiLineString{}
	return &this
}

// NewMultiLineStringWithDefaults instantiates a new MultiLineString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiLineStringWithDefaults() *MultiLineString {
	this := MultiLineString{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MultiLineString) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiLineString) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MultiLineString) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MultiLineString) SetType(v string) {
	o.Type = &v
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise.
func (o *MultiLineString) GetCoordinates() [][][]float32 {
	if o == nil || o.Coordinates == nil {
		var ret [][][]float32
		return ret
	}
	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiLineString) GetCoordinatesOk() ([][][]float32, bool) {
	if o == nil || o.Coordinates == nil {
		return nil, false
	}
	return o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *MultiLineString) HasCoordinates() bool {
	if o != nil && o.Coordinates != nil {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given [][][]float32 and assigns it to the Coordinates field.
func (o *MultiLineString) SetCoordinates(v [][][]float32) {
	o.Coordinates = v
}

func (o MultiLineString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Coordinates != nil {
		toSerialize["coordinates"] = o.Coordinates
	}
	return json.Marshal(toSerialize)
}

type NullableMultiLineString struct {
	value *MultiLineString
	isSet bool
}

func (v NullableMultiLineString) Get() *MultiLineString {
	return v.value
}

func (v *NullableMultiLineString) Set(val *MultiLineString) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiLineString) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiLineString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiLineString(val *MultiLineString) *NullableMultiLineString {
	return &NullableMultiLineString{value: val, isSet: true}
}

func (v NullableMultiLineString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiLineString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

