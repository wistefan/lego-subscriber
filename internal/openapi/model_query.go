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

// Query struct for Query
type Query struct {
	Type *string `json:"type,omitempty"`
	Entities *EntityInfo `json:"entities,omitempty"`
	Attrs []string `json:"attrs,omitempty"`
	Q *string `json:"q,omitempty"`
	GeoQ *GeoQuery `json:"geoQ,omitempty"`
	TemporalQ *TemporalQuery `json:"temporalQ,omitempty"`
	Csf *string `json:"csf,omitempty"`
}

// NewQuery instantiates a new Query object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuery() *Query {
	this := Query{}
	return &this
}

// NewQueryWithDefaults instantiates a new Query object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWithDefaults() *Query {
	this := Query{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Query) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Query) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Query) SetType(v string) {
	o.Type = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *Query) GetEntities() EntityInfo {
	if o == nil || o.Entities == nil {
		var ret EntityInfo
		return ret
	}
	return *o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetEntitiesOk() (*EntityInfo, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *Query) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given EntityInfo and assigns it to the Entities field.
func (o *Query) SetEntities(v EntityInfo) {
	o.Entities = &v
}

// GetAttrs returns the Attrs field value if set, zero value otherwise.
func (o *Query) GetAttrs() []string {
	if o == nil || o.Attrs == nil {
		var ret []string
		return ret
	}
	return o.Attrs
}

// GetAttrsOk returns a tuple with the Attrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetAttrsOk() ([]string, bool) {
	if o == nil || o.Attrs == nil {
		return nil, false
	}
	return o.Attrs, true
}

// HasAttrs returns a boolean if a field has been set.
func (o *Query) HasAttrs() bool {
	if o != nil && o.Attrs != nil {
		return true
	}

	return false
}

// SetAttrs gets a reference to the given []string and assigns it to the Attrs field.
func (o *Query) SetAttrs(v []string) {
	o.Attrs = v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *Query) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *Query) HasQ() bool {
	if o != nil && o.Q != nil {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *Query) SetQ(v string) {
	o.Q = &v
}

// GetGeoQ returns the GeoQ field value if set, zero value otherwise.
func (o *Query) GetGeoQ() GeoQuery {
	if o == nil || o.GeoQ == nil {
		var ret GeoQuery
		return ret
	}
	return *o.GeoQ
}

// GetGeoQOk returns a tuple with the GeoQ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetGeoQOk() (*GeoQuery, bool) {
	if o == nil || o.GeoQ == nil {
		return nil, false
	}
	return o.GeoQ, true
}

// HasGeoQ returns a boolean if a field has been set.
func (o *Query) HasGeoQ() bool {
	if o != nil && o.GeoQ != nil {
		return true
	}

	return false
}

// SetGeoQ gets a reference to the given GeoQuery and assigns it to the GeoQ field.
func (o *Query) SetGeoQ(v GeoQuery) {
	o.GeoQ = &v
}

// GetTemporalQ returns the TemporalQ field value if set, zero value otherwise.
func (o *Query) GetTemporalQ() TemporalQuery {
	if o == nil || o.TemporalQ == nil {
		var ret TemporalQuery
		return ret
	}
	return *o.TemporalQ
}

// GetTemporalQOk returns a tuple with the TemporalQ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetTemporalQOk() (*TemporalQuery, bool) {
	if o == nil || o.TemporalQ == nil {
		return nil, false
	}
	return o.TemporalQ, true
}

// HasTemporalQ returns a boolean if a field has been set.
func (o *Query) HasTemporalQ() bool {
	if o != nil && o.TemporalQ != nil {
		return true
	}

	return false
}

// SetTemporalQ gets a reference to the given TemporalQuery and assigns it to the TemporalQ field.
func (o *Query) SetTemporalQ(v TemporalQuery) {
	o.TemporalQ = &v
}

// GetCsf returns the Csf field value if set, zero value otherwise.
func (o *Query) GetCsf() string {
	if o == nil || o.Csf == nil {
		var ret string
		return ret
	}
	return *o.Csf
}

// GetCsfOk returns a tuple with the Csf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetCsfOk() (*string, bool) {
	if o == nil || o.Csf == nil {
		return nil, false
	}
	return o.Csf, true
}

// HasCsf returns a boolean if a field has been set.
func (o *Query) HasCsf() bool {
	if o != nil && o.Csf != nil {
		return true
	}

	return false
}

// SetCsf gets a reference to the given string and assigns it to the Csf field.
func (o *Query) SetCsf(v string) {
	o.Csf = &v
}

func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.Attrs != nil {
		toSerialize["attrs"] = o.Attrs
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	if o.GeoQ != nil {
		toSerialize["geoQ"] = o.GeoQ
	}
	if o.TemporalQ != nil {
		toSerialize["temporalQ"] = o.TemporalQ
	}
	if o.Csf != nil {
		toSerialize["csf"] = o.Csf
	}
	return json.Marshal(toSerialize)
}

type NullableQuery struct {
	value *Query
	isSet bool
}

func (v NullableQuery) Get() *Query {
	return v.value
}

func (v *NullableQuery) Set(val *Query) {
	v.value = val
	v.isSet = true
}

func (v NullableQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuery(val *Query) *NullableQuery {
	return &NullableQuery{value: val, isSet: true}
}

func (v NullableQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

