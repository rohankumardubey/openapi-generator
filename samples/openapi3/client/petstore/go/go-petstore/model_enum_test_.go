/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"fmt"
)

// checks if the EnumTest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnumTest{}

// EnumTest struct for EnumTest
type EnumTest struct {
	EnumString *string `json:"enum_string,omitempty"`
	EnumStringRequired string `json:"enum_string_required"`
	EnumInteger *int32 `json:"enum_integer,omitempty"`
	EnumNumber *float64 `json:"enum_number,omitempty"`
	OuterEnum NullableOuterEnum `json:"outerEnum,omitempty"`
	OuterEnumInteger *OuterEnumInteger `json:"outerEnumInteger,omitempty"`
	OuterEnumDefaultValue *OuterEnumDefaultValue `json:"outerEnumDefaultValue,omitempty"`
	OuterEnumIntegerDefaultValue *OuterEnumIntegerDefaultValue `json:"outerEnumIntegerDefaultValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnumTest EnumTest

// NewEnumTest instantiates a new EnumTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnumTest(enumStringRequired string) *EnumTest {
	this := EnumTest{}
	this.EnumStringRequired = enumStringRequired
	var outerEnumDefaultValue OuterEnumDefaultValue = OUTERENUMDEFAULTVALUE_PLACED
	this.OuterEnumDefaultValue = &outerEnumDefaultValue
	var outerEnumIntegerDefaultValue OuterEnumIntegerDefaultValue = OUTERENUMINTEGERDEFAULTVALUE__0
	this.OuterEnumIntegerDefaultValue = &outerEnumIntegerDefaultValue
	return &this
}

// NewEnumTestWithDefaults instantiates a new EnumTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnumTestWithDefaults() *EnumTest {
	this := EnumTest{}
	var outerEnumDefaultValue OuterEnumDefaultValue = OUTERENUMDEFAULTVALUE_PLACED
	this.OuterEnumDefaultValue = &outerEnumDefaultValue
	var outerEnumIntegerDefaultValue OuterEnumIntegerDefaultValue = OUTERENUMINTEGERDEFAULTVALUE__0
	this.OuterEnumIntegerDefaultValue = &outerEnumIntegerDefaultValue
	return &this
}

// GetEnumString returns the EnumString field value if set, zero value otherwise.
func (o *EnumTest) GetEnumString() string {
	if o == nil || IsNil(o.EnumString) {
		var ret string
		return ret
	}
	return *o.EnumString
}

// GetEnumStringOk returns a tuple with the EnumString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumStringOk() (*string, bool) {
	if o == nil || IsNil(o.EnumString) {
		return nil, false
	}
	return o.EnumString, true
}

// HasEnumString returns a boolean if a field has been set.
func (o *EnumTest) HasEnumString() bool {
	if o != nil && !IsNil(o.EnumString) {
		return true
	}

	return false
}

// SetEnumString gets a reference to the given string and assigns it to the EnumString field.
func (o *EnumTest) SetEnumString(v string) {
	o.EnumString = &v
}

// GetEnumStringRequired returns the EnumStringRequired field value
func (o *EnumTest) GetEnumStringRequired() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnumStringRequired
}

// GetEnumStringRequiredOk returns a tuple with the EnumStringRequired field value
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumStringRequiredOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnumStringRequired, true
}

// SetEnumStringRequired sets field value
func (o *EnumTest) SetEnumStringRequired(v string) {
	o.EnumStringRequired = v
}

// GetEnumInteger returns the EnumInteger field value if set, zero value otherwise.
func (o *EnumTest) GetEnumInteger() int32 {
	if o == nil || IsNil(o.EnumInteger) {
		var ret int32
		return ret
	}
	return *o.EnumInteger
}

// GetEnumIntegerOk returns a tuple with the EnumInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumIntegerOk() (*int32, bool) {
	if o == nil || IsNil(o.EnumInteger) {
		return nil, false
	}
	return o.EnumInteger, true
}

// HasEnumInteger returns a boolean if a field has been set.
func (o *EnumTest) HasEnumInteger() bool {
	if o != nil && !IsNil(o.EnumInteger) {
		return true
	}

	return false
}

// SetEnumInteger gets a reference to the given int32 and assigns it to the EnumInteger field.
func (o *EnumTest) SetEnumInteger(v int32) {
	o.EnumInteger = &v
}

// GetEnumNumber returns the EnumNumber field value if set, zero value otherwise.
func (o *EnumTest) GetEnumNumber() float64 {
	if o == nil || IsNil(o.EnumNumber) {
		var ret float64
		return ret
	}
	return *o.EnumNumber
}

// GetEnumNumberOk returns a tuple with the EnumNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetEnumNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.EnumNumber) {
		return nil, false
	}
	return o.EnumNumber, true
}

// HasEnumNumber returns a boolean if a field has been set.
func (o *EnumTest) HasEnumNumber() bool {
	if o != nil && !IsNil(o.EnumNumber) {
		return true
	}

	return false
}

// SetEnumNumber gets a reference to the given float64 and assigns it to the EnumNumber field.
func (o *EnumTest) SetEnumNumber(v float64) {
	o.EnumNumber = &v
}

// GetOuterEnum returns the OuterEnum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnumTest) GetOuterEnum() OuterEnum {
	if o == nil || IsNil(o.OuterEnum.Get()) {
		var ret OuterEnum
		return ret
	}
	return *o.OuterEnum.Get()
}

// GetOuterEnumOk returns a tuple with the OuterEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnumTest) GetOuterEnumOk() (*OuterEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.OuterEnum.Get(), o.OuterEnum.IsSet()
}

// HasOuterEnum returns a boolean if a field has been set.
func (o *EnumTest) HasOuterEnum() bool {
	if o != nil && o.OuterEnum.IsSet() {
		return true
	}

	return false
}

// SetOuterEnum gets a reference to the given NullableOuterEnum and assigns it to the OuterEnum field.
func (o *EnumTest) SetOuterEnum(v OuterEnum) {
	o.OuterEnum.Set(&v)
}
// SetOuterEnumNil sets the value for OuterEnum to be an explicit nil
func (o *EnumTest) SetOuterEnumNil() {
	o.OuterEnum.Set(nil)
}

// UnsetOuterEnum ensures that no value is present for OuterEnum, not even an explicit nil
func (o *EnumTest) UnsetOuterEnum() {
	o.OuterEnum.Unset()
}

// GetOuterEnumInteger returns the OuterEnumInteger field value if set, zero value otherwise.
func (o *EnumTest) GetOuterEnumInteger() OuterEnumInteger {
	if o == nil || IsNil(o.OuterEnumInteger) {
		var ret OuterEnumInteger
		return ret
	}
	return *o.OuterEnumInteger
}

// GetOuterEnumIntegerOk returns a tuple with the OuterEnumInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetOuterEnumIntegerOk() (*OuterEnumInteger, bool) {
	if o == nil || IsNil(o.OuterEnumInteger) {
		return nil, false
	}
	return o.OuterEnumInteger, true
}

// HasOuterEnumInteger returns a boolean if a field has been set.
func (o *EnumTest) HasOuterEnumInteger() bool {
	if o != nil && !IsNil(o.OuterEnumInteger) {
		return true
	}

	return false
}

// SetOuterEnumInteger gets a reference to the given OuterEnumInteger and assigns it to the OuterEnumInteger field.
func (o *EnumTest) SetOuterEnumInteger(v OuterEnumInteger) {
	o.OuterEnumInteger = &v
}

// GetOuterEnumDefaultValue returns the OuterEnumDefaultValue field value if set, zero value otherwise.
func (o *EnumTest) GetOuterEnumDefaultValue() OuterEnumDefaultValue {
	if o == nil || IsNil(o.OuterEnumDefaultValue) {
		var ret OuterEnumDefaultValue
		return ret
	}
	return *o.OuterEnumDefaultValue
}

// GetOuterEnumDefaultValueOk returns a tuple with the OuterEnumDefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetOuterEnumDefaultValueOk() (*OuterEnumDefaultValue, bool) {
	if o == nil || IsNil(o.OuterEnumDefaultValue) {
		return nil, false
	}
	return o.OuterEnumDefaultValue, true
}

// HasOuterEnumDefaultValue returns a boolean if a field has been set.
func (o *EnumTest) HasOuterEnumDefaultValue() bool {
	if o != nil && !IsNil(o.OuterEnumDefaultValue) {
		return true
	}

	return false
}

// SetOuterEnumDefaultValue gets a reference to the given OuterEnumDefaultValue and assigns it to the OuterEnumDefaultValue field.
func (o *EnumTest) SetOuterEnumDefaultValue(v OuterEnumDefaultValue) {
	o.OuterEnumDefaultValue = &v
}

// GetOuterEnumIntegerDefaultValue returns the OuterEnumIntegerDefaultValue field value if set, zero value otherwise.
func (o *EnumTest) GetOuterEnumIntegerDefaultValue() OuterEnumIntegerDefaultValue {
	if o == nil || IsNil(o.OuterEnumIntegerDefaultValue) {
		var ret OuterEnumIntegerDefaultValue
		return ret
	}
	return *o.OuterEnumIntegerDefaultValue
}

// GetOuterEnumIntegerDefaultValueOk returns a tuple with the OuterEnumIntegerDefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumTest) GetOuterEnumIntegerDefaultValueOk() (*OuterEnumIntegerDefaultValue, bool) {
	if o == nil || IsNil(o.OuterEnumIntegerDefaultValue) {
		return nil, false
	}
	return o.OuterEnumIntegerDefaultValue, true
}

// HasOuterEnumIntegerDefaultValue returns a boolean if a field has been set.
func (o *EnumTest) HasOuterEnumIntegerDefaultValue() bool {
	if o != nil && !IsNil(o.OuterEnumIntegerDefaultValue) {
		return true
	}

	return false
}

// SetOuterEnumIntegerDefaultValue gets a reference to the given OuterEnumIntegerDefaultValue and assigns it to the OuterEnumIntegerDefaultValue field.
func (o *EnumTest) SetOuterEnumIntegerDefaultValue(v OuterEnumIntegerDefaultValue) {
	o.OuterEnumIntegerDefaultValue = &v
}

func (o EnumTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnumTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnumString) {
		toSerialize["enum_string"] = o.EnumString
	}
	toSerialize["enum_string_required"] = o.EnumStringRequired
	if !IsNil(o.EnumInteger) {
		toSerialize["enum_integer"] = o.EnumInteger
	}
	if !IsNil(o.EnumNumber) {
		toSerialize["enum_number"] = o.EnumNumber
	}
	if o.OuterEnum.IsSet() {
		toSerialize["outerEnum"] = o.OuterEnum.Get()
	}
	if !IsNil(o.OuterEnumInteger) {
		toSerialize["outerEnumInteger"] = o.OuterEnumInteger
	}
	if !IsNil(o.OuterEnumDefaultValue) {
		toSerialize["outerEnumDefaultValue"] = o.OuterEnumDefaultValue
	}
	if !IsNil(o.OuterEnumIntegerDefaultValue) {
		toSerialize["outerEnumIntegerDefaultValue"] = o.OuterEnumIntegerDefaultValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnumTest) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enum_string_required",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEnumTest := _EnumTest{}

	err = json.Unmarshal(bytes, &varEnumTest)

	if err != nil {
		return err
	}

	*o = EnumTest(varEnumTest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enum_string")
		delete(additionalProperties, "enum_string_required")
		delete(additionalProperties, "enum_integer")
		delete(additionalProperties, "enum_number")
		delete(additionalProperties, "outerEnum")
		delete(additionalProperties, "outerEnumInteger")
		delete(additionalProperties, "outerEnumDefaultValue")
		delete(additionalProperties, "outerEnumIntegerDefaultValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnumTest struct {
	value *EnumTest
	isSet bool
}

func (v NullableEnumTest) Get() *EnumTest {
	return v.value
}

func (v *NullableEnumTest) Set(val *EnumTest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumTest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumTest(val *EnumTest) *NullableEnumTest {
	return &NullableEnumTest{value: val, isSet: true}
}

func (v NullableEnumTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


