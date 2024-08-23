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

// RoutingRule Specifies the redirect behavior and when a redirect is applied.
type RoutingRule struct {
	XMLName   xml.Name              `xml:"RoutingRule"`
	Condition *RoutingRuleCondition `json:"Condition,omitempty" xml:"Condition"`
	Redirect  *Redirect             `json:"Redirect" xml:"Redirect"`
}

// NewRoutingRule instantiates a new RoutingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingRule(redirect Redirect) *RoutingRule {
	this := RoutingRule{}

	this.Redirect = &redirect

	return &this
}

// NewRoutingRuleWithDefaults instantiates a new RoutingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingRuleWithDefaults() *RoutingRule {
	this := RoutingRule{}
	return &this
}

// GetCondition returns the Condition field value
// If the value is explicit nil, the zero value for RoutingRuleCondition will be returned
func (o *RoutingRule) GetCondition() *RoutingRuleCondition {
	if o == nil {
		return nil
	}

	return o.Condition

}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingRule) GetConditionOk() (*RoutingRuleCondition, bool) {
	if o == nil {
		return nil, false
	}

	return o.Condition, true
}

// SetCondition sets field value
func (o *RoutingRule) SetCondition(v RoutingRuleCondition) {

	o.Condition = &v

}

// HasCondition returns a boolean if a field has been set.
func (o *RoutingRule) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// GetRedirect returns the Redirect field value
// If the value is explicit nil, the zero value for Redirect will be returned
func (o *RoutingRule) GetRedirect() *Redirect {
	if o == nil {
		return nil
	}

	return o.Redirect

}

// GetRedirectOk returns a tuple with the Redirect field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingRule) GetRedirectOk() (*Redirect, bool) {
	if o == nil {
		return nil, false
	}

	return o.Redirect, true
}

// SetRedirect sets field value
func (o *RoutingRule) SetRedirect(v Redirect) {

	o.Redirect = &v

}

// HasRedirect returns a boolean if a field has been set.
func (o *RoutingRule) HasRedirect() bool {
	if o != nil && o.Redirect != nil {
		return true
	}

	return false
}

func (o RoutingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Condition != nil {
		toSerialize["Condition"] = o.Condition
	}

	if o.Redirect != nil {
		toSerialize["Redirect"] = o.Redirect
	}

	return json.Marshal(toSerialize)
}

type NullableRoutingRule struct {
	value *RoutingRule
	isSet bool
}

func (v NullableRoutingRule) Get() *RoutingRule {
	return v.value
}

func (v *NullableRoutingRule) Set(val *RoutingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingRule(val *RoutingRule) *NullableRoutingRule {
	return &NullableRoutingRule{value: val, isSet: true}
}

func (v NullableRoutingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}