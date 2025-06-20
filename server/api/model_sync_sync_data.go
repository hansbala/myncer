/*
Myncer API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// SyncSyncData - struct for SyncSyncData
type SyncSyncData struct {
	MergeSync *MergeSync
	OneWaySync *OneWaySync
}

// MergeSyncAsSyncSyncData is a convenience function that returns MergeSync wrapped in SyncSyncData
func MergeSyncAsSyncSyncData(v *MergeSync) SyncSyncData {
	return SyncSyncData{
		MergeSync: v,
	}
}

// OneWaySyncAsSyncSyncData is a convenience function that returns OneWaySync wrapped in SyncSyncData
func OneWaySyncAsSyncSyncData(v *OneWaySync) SyncSyncData {
	return SyncSyncData{
		OneWaySync: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SyncSyncData) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MergeSync
	err = newStrictDecoder(data).Decode(&dst.MergeSync)
	if err == nil {
		jsonMergeSync, _ := json.Marshal(dst.MergeSync)
		if string(jsonMergeSync) == "{}" { // empty struct
			dst.MergeSync = nil
		} else {
			if err = validator.Validate(dst.MergeSync); err != nil {
				dst.MergeSync = nil
			} else {
				match++
			}
		}
	} else {
		dst.MergeSync = nil
	}

	// try to unmarshal data into OneWaySync
	err = newStrictDecoder(data).Decode(&dst.OneWaySync)
	if err == nil {
		jsonOneWaySync, _ := json.Marshal(dst.OneWaySync)
		if string(jsonOneWaySync) == "{}" { // empty struct
			dst.OneWaySync = nil
		} else {
			if err = validator.Validate(dst.OneWaySync); err != nil {
				dst.OneWaySync = nil
			} else {
				match++
			}
		}
	} else {
		dst.OneWaySync = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MergeSync = nil
		dst.OneWaySync = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SyncSyncData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SyncSyncData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SyncSyncData) MarshalJSON() ([]byte, error) {
	if src.MergeSync != nil {
		return json.Marshal(&src.MergeSync)
	}

	if src.OneWaySync != nil {
		return json.Marshal(&src.OneWaySync)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SyncSyncData) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MergeSync != nil {
		return obj.MergeSync
	}

	if obj.OneWaySync != nil {
		return obj.OneWaySync
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SyncSyncData) GetActualInstanceValue() (interface{}) {
	if obj.MergeSync != nil {
		return *obj.MergeSync
	}

	if obj.OneWaySync != nil {
		return *obj.OneWaySync
	}

	// all schemas are nil
	return nil
}

type NullableSyncSyncData struct {
	value *SyncSyncData
	isSet bool
}

func (v NullableSyncSyncData) Get() *SyncSyncData {
	return v.value
}

func (v *NullableSyncSyncData) Set(val *SyncSyncData) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncSyncData) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncSyncData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncSyncData(val *SyncSyncData) *NullableSyncSyncData {
	return &NullableSyncSyncData{value: val, isSet: true}
}

func (v NullableSyncSyncData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncSyncData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


