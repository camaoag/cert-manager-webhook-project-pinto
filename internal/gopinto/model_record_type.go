/*
 * Fava - OpenApi Gateway - DNS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gopinto

import (
	"encoding/json"
	"fmt"
)

// RecordType Resource record types  as defined in  <see href=\"https://tools.ietf.org/html/rfc1035#section-3.2.2\">rfc1035</see>
type RecordType string

// List of RecordType
const (
	RECORDTYPE_A     RecordType = "A"
	RECORDTYPE_NS    RecordType = "NS"
	RECORDTYPE_CNAME RecordType = "CNAME"
	RECORDTYPE_SOA   RecordType = "SOA"
	RECORDTYPE_PTR   RecordType = "PTR"
	RECORDTYPE_MX    RecordType = "MX"
	RECORDTYPE_TXT   RecordType = "TXT"
	RECORDTYPE_SRV   RecordType = "SRV"
	RECORDTYPE_AAAA  RecordType = "AAAA"
	RECORDTYPE_SPF   RecordType = "SPF"
)

var allowedRecordTypeEnumValues = []RecordType{
	"A",
	"NS",
	"CNAME",
	"SOA",
	"PTR",
	"MX",
	"TXT",
	"SRV",
	"AAAA",
	"SPF",
}

func (v *RecordType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecordType(value)
	for _, existing := range allowedRecordTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecordType", value)
}

// NewRecordTypeFromValue returns a pointer to a valid RecordType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecordTypeFromValue(v string) (*RecordType, error) {
	ev := RecordType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecordType: valid values are %v", v, allowedRecordTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecordType) IsValid() bool {
	for _, existing := range allowedRecordTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecordType value
func (v RecordType) Ptr() *RecordType {
	return &v
}

type NullableRecordType struct {
	value *RecordType
	isSet bool
}

func (v NullableRecordType) Get() *RecordType {
	return v.value
}

func (v *NullableRecordType) Set(val *RecordType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordType(val *RecordType) *NullableRecordType {
	return &NullableRecordType{value: val, isSet: true}
}

func (v NullableRecordType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}