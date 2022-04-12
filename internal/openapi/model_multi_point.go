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

// MultiPoint struct for MultiPoint
type MultiPoint struct {
	Type *string `json:"type,omitempty"`
	// An array of positions
	Coordinates [][]float32 `json:"coordinates,omitempty"`
}

// NewMultiPoint instantiates a new MultiPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiPoint() *MultiPoint {
	this := MultiPoint{}
	return &this
}

// NewMultiPointWithDefaults instantiates a new MultiPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiPointWithDefaults() *MultiPoint {
	this := MultiPoint{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MultiPoint) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPoint) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MultiPoint) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MultiPoint) SetType(v string) {
	o.Type = &v
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise.
func (o *MultiPoint) GetCoordinates() [][]float32 {
	if o == nil || o.Coordinates == nil {
		var ret [][]float32
		return ret
	}
	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiPoint) GetCoordinatesOk() ([][]float32, bool) {
	if o == nil || o.Coordinates == nil {
		return nil, false
	}
	return o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *MultiPoint) HasCoordinates() bool {
	if o != nil && o.Coordinates != nil {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given [][]float32 and assigns it to the Coordinates field.
func (o *MultiPoint) SetCoordinates(v [][]float32) {
	o.Coordinates = v
}

func (o MultiPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Coordinates != nil {
		toSerialize["coordinates"] = o.Coordinates
	}
	return json.Marshal(toSerialize)
}

type NullableMultiPoint struct {
	value *MultiPoint
	isSet bool
}

func (v NullableMultiPoint) Get() *MultiPoint {
	return v.value
}

func (v *NullableMultiPoint) Set(val *MultiPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiPoint(val *MultiPoint) *NullableMultiPoint {
	return &NullableMultiPoint{value: val, isSet: true}
}

func (v NullableMultiPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

