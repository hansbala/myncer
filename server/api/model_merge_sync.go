/*
Myncer API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MergeSync type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeSync{}

// MergeSync Representative of merging sources into a master source and writing to all.
type MergeSync struct {
	SyncVariant SyncVariant `json:"syncVariant"`
	// All sources that will be merged into one and written back to sources.
	Sources []MusicSource `json:"sources"`
}

type _MergeSync MergeSync

// NewMergeSync instantiates a new MergeSync object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeSync(syncVariant SyncVariant, sources []MusicSource) *MergeSync {
	this := MergeSync{}
	this.SyncVariant = syncVariant
	this.Sources = sources
	return &this
}

// NewMergeSyncWithDefaults instantiates a new MergeSync object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeSyncWithDefaults() *MergeSync {
	this := MergeSync{}
	return &this
}

// GetSyncVariant returns the SyncVariant field value
func (o *MergeSync) GetSyncVariant() SyncVariant {
	if o == nil {
		var ret SyncVariant
		return ret
	}

	return o.SyncVariant
}

// GetSyncVariantOk returns a tuple with the SyncVariant field value
// and a boolean to check if the value has been set.
func (o *MergeSync) GetSyncVariantOk() (*SyncVariant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncVariant, true
}

// SetSyncVariant sets field value
func (o *MergeSync) SetSyncVariant(v SyncVariant) {
	o.SyncVariant = v
}

// GetSources returns the Sources field value
func (o *MergeSync) GetSources() []MusicSource {
	if o == nil {
		var ret []MusicSource
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *MergeSync) GetSourcesOk() ([]MusicSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *MergeSync) SetSources(v []MusicSource) {
	o.Sources = v
}

func (o MergeSync) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeSync) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["syncVariant"] = o.SyncVariant
	toSerialize["sources"] = o.Sources
	return toSerialize, nil
}

func (o *MergeSync) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"syncVariant",
		"sources",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMergeSync := _MergeSync{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMergeSync)

	if err != nil {
		return err
	}

	*o = MergeSync(varMergeSync)

	return err
}

type NullableMergeSync struct {
	value *MergeSync
	isSet bool
}

func (v NullableMergeSync) Get() *MergeSync {
	return v.value
}

func (v *NullableMergeSync) Set(val *MergeSync) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeSync) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeSync) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeSync(val *MergeSync) *NullableMergeSync {
	return &NullableMergeSync{value: val, isSet: true}
}

func (v NullableMergeSync) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeSync) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


