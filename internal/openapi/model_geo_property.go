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
	"time"
)

// GeoProperty struct for GeoProperty
type GeoProperty struct {
	Type string `json:"type"`
	Value map[string]interface{} `json:"value"`
	ObservedAt *time.Time `json:"observedAt,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	DatasetId *string `json:"datasetId,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
	UnitCode *string `json:"unitCode,omitempty"`
}

// NewGeoProperty instantiates a new GeoProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoProperty(type_ string, value map[string]interface{}) *GeoProperty {
	this := GeoProperty{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewGeoPropertyWithDefaults instantiates a new GeoProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoPropertyWithDefaults() *GeoProperty {
	this := GeoProperty{}
	return &this
}

// GetType returns the Type field value
func (o *GeoProperty) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GeoProperty) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GeoProperty) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *GeoProperty) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GeoProperty) GetValueOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *GeoProperty) SetValue(v map[string]interface{}) {
	o.Value = v
}

// GetObservedAt returns the ObservedAt field value if set, zero value otherwise.
func (o *GeoProperty) GetObservedAt() time.Time {
	if o == nil || o.ObservedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoProperty) GetObservedAtOk() (*time.Time, bool) {
	if o == nil || o.ObservedAt == nil {
		return nil, false
	}
	return o.ObservedAt, true
}

// HasObservedAt returns a boolean if a field has been set.
func (o *GeoProperty) HasObservedAt() bool {
	if o != nil && o.ObservedAt != nil {
		return true
	}

	return false
}

// SetObservedAt gets a reference to the given time.Time and assigns it to the ObservedAt field.
func (o *GeoProperty) SetObservedAt(v time.Time) {
	o.ObservedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GeoProperty) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoProperty) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GeoProperty) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GeoProperty) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *GeoProperty) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoProperty) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *GeoProperty) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *GeoProperty) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *GeoProperty) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoProperty) GetDatasetIdOk() (*string, bool) {
	if o == nil || o.DatasetId == nil {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *GeoProperty) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *GeoProperty) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *GeoProperty) GetInstanceId() string {
	if o == nil || o.InstanceId == nil {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoProperty) GetInstanceIdOk() (*string, bool) {
	if o == nil || o.InstanceId == nil {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *GeoProperty) HasInstanceId() bool {
	if o != nil && o.InstanceId != nil {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *GeoProperty) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetUnitCode returns the UnitCode field value if set, zero value otherwise.
func (o *GeoProperty) GetUnitCode() string {
	if o == nil || o.UnitCode == nil {
		var ret string
		return ret
	}
	return *o.UnitCode
}

// GetUnitCodeOk returns a tuple with the UnitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeoProperty) GetUnitCodeOk() (*string, bool) {
	if o == nil || o.UnitCode == nil {
		return nil, false
	}
	return o.UnitCode, true
}

// HasUnitCode returns a boolean if a field has been set.
func (o *GeoProperty) HasUnitCode() bool {
	if o != nil && o.UnitCode != nil {
		return true
	}

	return false
}

// SetUnitCode gets a reference to the given string and assigns it to the UnitCode field.
func (o *GeoProperty) SetUnitCode(v string) {
	o.UnitCode = &v
}

func (o GeoProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.ObservedAt != nil {
		toSerialize["observedAt"] = o.ObservedAt
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.DatasetId != nil {
		toSerialize["datasetId"] = o.DatasetId
	}
	if o.InstanceId != nil {
		toSerialize["instanceId"] = o.InstanceId
	}
	if o.UnitCode != nil {
		toSerialize["unitCode"] = o.UnitCode
	}
	return json.Marshal(toSerialize)
}

type NullableGeoProperty struct {
	value *GeoProperty
	isSet bool
}

func (v NullableGeoProperty) Get() *GeoProperty {
	return v.value
}

func (v *NullableGeoProperty) Set(val *GeoProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoProperty(val *GeoProperty) *NullableGeoProperty {
	return &NullableGeoProperty{value: val, isSet: true}
}

func (v NullableGeoProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

