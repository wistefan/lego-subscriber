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

// Subscription struct for Subscription
type Subscription struct {
	Context map[string]interface{} `json:"@context,omitempty"`
	Entities []EntityInfo `json:"entities,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	WatchedAttributes []string `json:"watchedAttributes,omitempty"`
	TimeInterval *float32 `json:"timeInterval,omitempty"`
	Expires *time.Time `json:"expires,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	Throttling *float32 `json:"throttling,omitempty"`
	Q *string `json:"q,omitempty"`
	GeoQ *GeoQuery `json:"geoQ,omitempty"`
	Csf *string `json:"csf,omitempty"`
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Notification *NotificationParams `json:"notification,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription() *Subscription {
	this := Subscription{}
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *Subscription) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetContextOk() (map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *Subscription) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *Subscription) SetContext(v map[string]interface{}) {
	o.Context = v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *Subscription) GetEntities() []EntityInfo {
	if o == nil || o.Entities == nil {
		var ret []EntityInfo
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetEntitiesOk() ([]EntityInfo, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *Subscription) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []EntityInfo and assigns it to the Entities field.
func (o *Subscription) SetEntities(v []EntityInfo) {
	o.Entities = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Subscription) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Subscription) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Subscription) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Subscription) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Subscription) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Subscription) SetDescription(v string) {
	o.Description = &v
}

// GetWatchedAttributes returns the WatchedAttributes field value if set, zero value otherwise.
func (o *Subscription) GetWatchedAttributes() []string {
	if o == nil || o.WatchedAttributes == nil {
		var ret []string
		return ret
	}
	return o.WatchedAttributes
}

// GetWatchedAttributesOk returns a tuple with the WatchedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetWatchedAttributesOk() ([]string, bool) {
	if o == nil || o.WatchedAttributes == nil {
		return nil, false
	}
	return o.WatchedAttributes, true
}

// HasWatchedAttributes returns a boolean if a field has been set.
func (o *Subscription) HasWatchedAttributes() bool {
	if o != nil && o.WatchedAttributes != nil {
		return true
	}

	return false
}

// SetWatchedAttributes gets a reference to the given []string and assigns it to the WatchedAttributes field.
func (o *Subscription) SetWatchedAttributes(v []string) {
	o.WatchedAttributes = v
}

// GetTimeInterval returns the TimeInterval field value if set, zero value otherwise.
func (o *Subscription) GetTimeInterval() float32 {
	if o == nil || o.TimeInterval == nil {
		var ret float32
		return ret
	}
	return *o.TimeInterval
}

// GetTimeIntervalOk returns a tuple with the TimeInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetTimeIntervalOk() (*float32, bool) {
	if o == nil || o.TimeInterval == nil {
		return nil, false
	}
	return o.TimeInterval, true
}

// HasTimeInterval returns a boolean if a field has been set.
func (o *Subscription) HasTimeInterval() bool {
	if o != nil && o.TimeInterval != nil {
		return true
	}

	return false
}

// SetTimeInterval gets a reference to the given float32 and assigns it to the TimeInterval field.
func (o *Subscription) SetTimeInterval(v float32) {
	o.TimeInterval = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *Subscription) GetExpires() time.Time {
	if o == nil || o.Expires == nil {
		var ret time.Time
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetExpiresOk() (*time.Time, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *Subscription) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given time.Time and assigns it to the Expires field.
func (o *Subscription) SetExpires(v time.Time) {
	o.Expires = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *Subscription) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *Subscription) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *Subscription) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetThrottling returns the Throttling field value if set, zero value otherwise.
func (o *Subscription) GetThrottling() float32 {
	if o == nil || o.Throttling == nil {
		var ret float32
		return ret
	}
	return *o.Throttling
}

// GetThrottlingOk returns a tuple with the Throttling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetThrottlingOk() (*float32, bool) {
	if o == nil || o.Throttling == nil {
		return nil, false
	}
	return o.Throttling, true
}

// HasThrottling returns a boolean if a field has been set.
func (o *Subscription) HasThrottling() bool {
	if o != nil && o.Throttling != nil {
		return true
	}

	return false
}

// SetThrottling gets a reference to the given float32 and assigns it to the Throttling field.
func (o *Subscription) SetThrottling(v float32) {
	o.Throttling = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *Subscription) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *Subscription) HasQ() bool {
	if o != nil && o.Q != nil {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *Subscription) SetQ(v string) {
	o.Q = &v
}

// GetGeoQ returns the GeoQ field value if set, zero value otherwise.
func (o *Subscription) GetGeoQ() GeoQuery {
	if o == nil || o.GeoQ == nil {
		var ret GeoQuery
		return ret
	}
	return *o.GeoQ
}

// GetGeoQOk returns a tuple with the GeoQ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetGeoQOk() (*GeoQuery, bool) {
	if o == nil || o.GeoQ == nil {
		return nil, false
	}
	return o.GeoQ, true
}

// HasGeoQ returns a boolean if a field has been set.
func (o *Subscription) HasGeoQ() bool {
	if o != nil && o.GeoQ != nil {
		return true
	}

	return false
}

// SetGeoQ gets a reference to the given GeoQuery and assigns it to the GeoQ field.
func (o *Subscription) SetGeoQ(v GeoQuery) {
	o.GeoQ = &v
}

// GetCsf returns the Csf field value if set, zero value otherwise.
func (o *Subscription) GetCsf() string {
	if o == nil || o.Csf == nil {
		var ret string
		return ret
	}
	return *o.Csf
}

// GetCsfOk returns a tuple with the Csf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCsfOk() (*string, bool) {
	if o == nil || o.Csf == nil {
		return nil, false
	}
	return o.Csf, true
}

// HasCsf returns a boolean if a field has been set.
func (o *Subscription) HasCsf() bool {
	if o != nil && o.Csf != nil {
		return true
	}

	return false
}

// SetCsf gets a reference to the given string and assigns it to the Csf field.
func (o *Subscription) SetCsf(v string) {
	o.Csf = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Subscription) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Subscription) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Subscription) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Subscription) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Subscription) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Subscription) SetType(v string) {
	o.Type = &v
}

// GetNotification returns the Notification field value if set, zero value otherwise.
func (o *Subscription) GetNotification() NotificationParams {
	if o == nil || o.Notification == nil {
		var ret NotificationParams
		return ret
	}
	return *o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetNotificationOk() (*NotificationParams, bool) {
	if o == nil || o.Notification == nil {
		return nil, false
	}
	return o.Notification, true
}

// HasNotification returns a boolean if a field has been set.
func (o *Subscription) HasNotification() bool {
	if o != nil && o.Notification != nil {
		return true
	}

	return false
}

// SetNotification gets a reference to the given NotificationParams and assigns it to the Notification field.
func (o *Subscription) SetNotification(v NotificationParams) {
	o.Notification = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Subscription) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Subscription) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Subscription) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Subscription) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Subscription) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Subscription) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Subscription) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Subscription) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *Subscription) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["@context"] = o.Context
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.WatchedAttributes != nil {
		toSerialize["watchedAttributes"] = o.WatchedAttributes
	}
	if o.TimeInterval != nil {
		toSerialize["timeInterval"] = o.TimeInterval
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	if o.Throttling != nil {
		toSerialize["throttling"] = o.Throttling
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	if o.GeoQ != nil {
		toSerialize["geoQ"] = o.GeoQ
	}
	if o.Csf != nil {
		toSerialize["csf"] = o.Csf
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Notification != nil {
		toSerialize["notification"] = o.Notification
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

