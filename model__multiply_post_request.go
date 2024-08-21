/*
Arithmetic Operations API

API for performing basic arithmetic operations

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calcssdkgo

import (
	"encoding/json"
)

// checks if the MultiplyPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiplyPostRequest{}

// MultiplyPostRequest struct for MultiplyPostRequest
type MultiplyPostRequest struct {
	// List of numbers to multiply
	Numbers []float32 `json:"numbers,omitempty"`
}

// NewMultiplyPostRequest instantiates a new MultiplyPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiplyPostRequest() *MultiplyPostRequest {
	this := MultiplyPostRequest{}
	return &this
}

// NewMultiplyPostRequestWithDefaults instantiates a new MultiplyPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiplyPostRequestWithDefaults() *MultiplyPostRequest {
	this := MultiplyPostRequest{}
	return &this
}

// GetNumbers returns the Numbers field value if set, zero value otherwise.
func (o *MultiplyPostRequest) GetNumbers() []float32 {
	if o == nil || IsNil(o.Numbers) {
		var ret []float32
		return ret
	}
	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiplyPostRequest) GetNumbersOk() ([]float32, bool) {
	if o == nil || IsNil(o.Numbers) {
		return nil, false
	}
	return o.Numbers, true
}

// HasNumbers returns a boolean if a field has been set.
func (o *MultiplyPostRequest) HasNumbers() bool {
	if o != nil && !IsNil(o.Numbers) {
		return true
	}

	return false
}

// SetNumbers gets a reference to the given []float32 and assigns it to the Numbers field.
func (o *MultiplyPostRequest) SetNumbers(v []float32) {
	o.Numbers = v
}

func (o MultiplyPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiplyPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Numbers) {
		toSerialize["numbers"] = o.Numbers
	}
	return toSerialize, nil
}

type NullableMultiplyPostRequest struct {
	value *MultiplyPostRequest
	isSet bool
}

func (v NullableMultiplyPostRequest) Get() *MultiplyPostRequest {
	return v.value
}

func (v *NullableMultiplyPostRequest) Set(val *MultiplyPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiplyPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiplyPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiplyPostRequest(val *MultiplyPostRequest) *NullableMultiplyPostRequest {
	return &NullableMultiplyPostRequest{value: val, isSet: true}
}

func (v NullableMultiplyPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiplyPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


