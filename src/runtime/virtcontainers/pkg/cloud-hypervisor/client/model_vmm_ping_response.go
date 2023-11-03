/*
Cloud Hypervisor API

Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.

API version: 0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VmmPingResponse Virtual Machine Monitor information
type VmmPingResponse struct {
	BuildVersion *string   `json:"build_version,omitempty"`
	Version      string    `json:"version"`
	Pid          *int64    `json:"pid,omitempty"`
	Features     *[]string `json:"features,omitempty"`
}

// NewVmmPingResponse instantiates a new VmmPingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmmPingResponse(version string) *VmmPingResponse {
	this := VmmPingResponse{}
	this.Version = version
	return &this
}

// NewVmmPingResponseWithDefaults instantiates a new VmmPingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmmPingResponseWithDefaults() *VmmPingResponse {
	this := VmmPingResponse{}
	return &this
}

// GetBuildVersion returns the BuildVersion field value if set, zero value otherwise.
func (o *VmmPingResponse) GetBuildVersion() string {
	if o == nil || o.BuildVersion == nil {
		var ret string
		return ret
	}
	return *o.BuildVersion
}

// GetBuildVersionOk returns a tuple with the BuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmmPingResponse) GetBuildVersionOk() (*string, bool) {
	if o == nil || o.BuildVersion == nil {
		return nil, false
	}
	return o.BuildVersion, true
}

// HasBuildVersion returns a boolean if a field has been set.
func (o *VmmPingResponse) HasBuildVersion() bool {
	if o != nil && o.BuildVersion != nil {
		return true
	}

	return false
}

// SetBuildVersion gets a reference to the given string and assigns it to the BuildVersion field.
func (o *VmmPingResponse) SetBuildVersion(v string) {
	o.BuildVersion = &v
}

// GetVersion returns the Version field value
func (o *VmmPingResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VmmPingResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VmmPingResponse) SetVersion(v string) {
	o.Version = v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *VmmPingResponse) GetPid() int64 {
	if o == nil || o.Pid == nil {
		var ret int64
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmmPingResponse) GetPidOk() (*int64, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *VmmPingResponse) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given int64 and assigns it to the Pid field.
func (o *VmmPingResponse) SetPid(v int64) {
	o.Pid = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *VmmPingResponse) GetFeatures() []string {
	if o == nil || o.Features == nil {
		var ret []string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmmPingResponse) GetFeaturesOk() (*[]string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *VmmPingResponse) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *VmmPingResponse) SetFeatures(v []string) {
	o.Features = &v
}

func (o VmmPingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BuildVersion != nil {
		toSerialize["build_version"] = o.BuildVersion
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.Pid != nil {
		toSerialize["pid"] = o.Pid
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableVmmPingResponse struct {
	value *VmmPingResponse
	isSet bool
}

func (v NullableVmmPingResponse) Get() *VmmPingResponse {
	return v.value
}

func (v *NullableVmmPingResponse) Set(val *VmmPingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVmmPingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVmmPingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmmPingResponse(val *VmmPingResponse) *NullableVmmPingResponse {
	return &NullableVmmPingResponse{value: val, isSet: true}
}

func (v NullableVmmPingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmmPingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
