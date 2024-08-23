/*
 * IONOS S3 Object Storage API for contract-owned buckets
 *
 * ## Overview The IONOS S3 Object Storage API for contract-owned buckets is a REST-based API that allows developers and applications to interact directly with IONOS' scalable storage solution, leveraging the S3 protocol for object storage operations. Its design ensures seamless compatibility with existing tools and libraries tailored for S3 systems.  ### API References - [S3 API Reference for contract-owned buckets](https://api.ionos.com/docs/s3-contract-owned-buckets/v2/) ### User documentation [IONOS S3 Object Storage User Guide](https://docs.ionos.com/cloud/managed-services/s3-object-storage) * [Documentation on user-owned and contract-owned buckets](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/buckets) * [Documentation on S3 API Compatibility](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/s3-api-compatibility) * [S3 Tools](https://docs.ionos.com/cloud/managed-services/s3-object-storage/s3-tools)  ## Endpoints for contract-owned buckets | Location | Region Name | Bucket Type | Endpoint | | --- | --- | --- | --- | | **Berlin, Germany** | **eu-central-3** | Contract-owned | `https://s3.eu-central-3.ionoscloud.com` |  ## Changelog - 30.05.2024 Initial version
 *
 * API version: 2.0.2
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

import "encoding/xml"

// BucketPolicy struct for BucketPolicy
type BucketPolicy struct {
	XMLName xml.Name `xml:"BucketPolicy"`
	// Specifies an optional identifier for the policy.
	Id *string `json:"Id,omitempty" xml:"Id"`
	// Policy version
	Version   *string                  `json:"Version,omitempty" xml:"Version"`
	Statement *[]BucketPolicyStatement `json:"Statement" xml:"Statement"`
}

// NewBucketPolicy instantiates a new BucketPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketPolicy(statement []BucketPolicyStatement) *BucketPolicy {
	this := BucketPolicy{}

	this.Statement = &statement

	return &this
}

// NewBucketPolicyWithDefaults instantiates a new BucketPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketPolicyWithDefaults() *BucketPolicy {
	this := BucketPolicy{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BucketPolicy) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BucketPolicy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *BucketPolicy) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *BucketPolicy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetVersion returns the Version field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BucketPolicy) GetVersion() *string {
	if o == nil {
		return nil
	}

	return o.Version

}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BucketPolicy) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Version, true
}

// SetVersion sets field value
func (o *BucketPolicy) SetVersion(v string) {

	o.Version = &v

}

// HasVersion returns a boolean if a field has been set.
func (o *BucketPolicy) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// GetStatement returns the Statement field value
// If the value is explicit nil, the zero value for []BucketPolicyStatement will be returned
func (o *BucketPolicy) GetStatement() *[]BucketPolicyStatement {
	if o == nil {
		return nil
	}

	return o.Statement

}

// GetStatementOk returns a tuple with the Statement field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BucketPolicy) GetStatementOk() (*[]BucketPolicyStatement, bool) {
	if o == nil {
		return nil, false
	}

	return o.Statement, true
}

// SetStatement sets field value
func (o *BucketPolicy) SetStatement(v []BucketPolicyStatement) {

	o.Statement = &v

}

// HasStatement returns a boolean if a field has been set.
func (o *BucketPolicy) HasStatement() bool {
	if o != nil && o.Statement != nil {
		return true
	}

	return false
}

func (o BucketPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}

	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	if o.Statement != nil {
		toSerialize["Statement"] = o.Statement
	}

	return json.Marshal(toSerialize)
}

type NullableBucketPolicy struct {
	value *BucketPolicy
	isSet bool
}

func (v NullableBucketPolicy) Get() *BucketPolicy {
	return v.value
}

func (v *NullableBucketPolicy) Set(val *BucketPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketPolicy(val *BucketPolicy) *NullableBucketPolicy {
	return &NullableBucketPolicy{value: val, isSet: true}
}

func (v NullableBucketPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
