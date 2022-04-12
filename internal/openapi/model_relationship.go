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

// Relationship struct for Relationship
type Relationship struct {
	Type string `json:"type"`
	Object string `json:"object"`
	ObservedAt *time.Time `json:"observedAt,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	DatasetId *string `json:"datasetId,omitempty"`
	InstanceId *string `json:"instanceId,omitempty"`
}

// NewRelationship instantiates a new Relationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationship(type_ string, object string) *Relationship {
	this := Relationship{}
	this.Type = type_
	this.Object = object
	return &this
}

// NewRelationshipWithDefaults instantiates a new Relationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWithDefaults() *Relationship {
	this := Relationship{}
	return &this
}

// GetType returns the Type field value
func (o *Relationship) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Relationship) SetType(v string) {
	o.Type = v
}

// GetObject returns the Object field value
func (o *Relationship) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Relationship) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Relationship) SetObject(v string) {
	o.Object = v
}

// GetObservedAt returns the ObservedAt field value if set, zero value otherwise.
func (o *Relationship) GetObservedAt() time.Time {
	if o == nil || o.ObservedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Relationship) GetObservedAtOk() (*time.Time, bool) {
	if o == nil || o.ObservedAt == nil {
		return nil, false
	}
	return o.ObservedAt, true
}

// HasObservedAt returns a boolean if a field has been set.
func (o *Relationship) HasObservedAt() bool {
	if o != nil && o.ObservedAt != nil {
		return true
	}

	return false
}

// SetObservedAt gets a reference to the given time.Time and assigns it to the ObservedAt field.
func (o *Relationship) SetObservedAt(v time.Time) {
	o.ObservedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Relationship) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Relationship) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Relationship) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Relationship) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Relationship) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Relationship) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Relationship) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *Relationship) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *Relationship) GetDatasetId() string {
	if o == nil || o.DatasetId == nil {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Relationship) GetDatasetIdOk() (*string, bool) {
	if o == nil || o.DatasetId == nil {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *Relationship) HasDatasetId() bool {
	if o != nil && o.DatasetId != nil {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *Relationship) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *Relationship) GetInstanceId() string {
	if o == nil || o.InstanceId == nil {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Relationship) GetInstanceIdOk() (*string, bool) {
	if o == nil || o.InstanceId == nil {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *Relationship) HasInstanceId() bool {
	if o != nil && o.InstanceId != nil {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *Relationship) SetInstanceId(v string) {
	o.InstanceId = &v
}

func (o Relationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["object"] = o.Object
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
	return json.Marshal(toSerialize)
}

type NullableRelationship struct {
	value *Relationship
	isSet bool
}

func (v NullableRelationship) Get() *Relationship {
	return v.value
}

func (v *NullableRelationship) Set(val *Relationship) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationship(val *Relationship) *NullableRelationship {
	return &NullableRelationship{value: val, isSet: true}
}

func (v NullableRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


