# Go API client for ionoscloud

## Overview
The IONOS S3 Object Storage API for contract-owned buckets is a REST-based API that allows developers and applications to interact directly with
IONOS' scalable storage solution, leveraging the S3 protocol for object storage operations. Its design ensures seamless
compatibility with existing tools and libraries tailored for S3 systems.

### API References
- [S3 Management API Reference](https://api.ionos.com/docs/s3-management/v1/) for managing Access Keys
- S3 API Reference for contract-owned buckets - current document
- [S3 API Reference for user-owned buckets](https://api.ionos.com/docs/s3-user-owned-buckets/v2/)

### User documentation
[IONOS S3 Object Storage User Guide](https://docs.ionos.com/cloud/managed-services/s3-object-storage)
* [Documentation on user-owned and contract-owned buckets](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/buckets)
* [Documentation on S3 API Compatibility](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/s3-api-compatibility)
* [S3 Tools](https://docs.ionos.com/cloud/managed-services/s3-object-storage/s3-tools)

## Endpoints for contract-owned buckets
| Location | Region Name | Bucket Type | Endpoint |
| --- | --- | --- | --- |
| **Berlin, Germany** | **eu-central-3** | Contract-owned | `https://s3.eu-central-3.ionoscloud.com` |

## Changelog
- 30.05.2024 Initial version


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.2
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://docs.ionos.com/support/general-information/contact-information](https://docs.ionos.com/support/general-information/contact-information)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import ionoscloud "github.com/ionos-cloud/sdk-go-s3"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), ionoscloud.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), ionoscloud.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

## Documentation for API Endpoints

All URIs are relative to *https://s3.eu-central-3.ionoscloud.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BucketsApi* | [**CreateBucket**](docs/api/BucketsApi.md#createbucket) | **Put** /{Bucket} | CreateBucket
*BucketsApi* | [**DeleteBucket**](docs/api/BucketsApi.md#deletebucket) | **Delete** /{Bucket} | DeleteBucket
*BucketsApi* | [**GetBucketLocation**](docs/api/BucketsApi.md#getbucketlocation) | **Get** /{Bucket}?location | GetBucketLocation
*BucketsApi* | [**HeadBucket**](docs/api/BucketsApi.md#headbucket) | **Head** /{Bucket} | HeadBucket
*BucketsApi* | [**ListBuckets**](docs/api/BucketsApi.md#listbuckets) | **Get** / | ListBuckets
*ObjectsApi* | [**CopyObject**](docs/api/ObjectsApi.md#copyobject) | **Put** /{Bucket}/{Key}?x-amz-copy-source | CopyObject
*ObjectsApi* | [**DeleteObject**](docs/api/ObjectsApi.md#deleteobject) | **Delete** /{Bucket}/{Key} | DeleteObject
*ObjectsApi* | [**DeleteObjects**](docs/api/ObjectsApi.md#deleteobjects) | **Post** /{Bucket}?delete | DeleteObjects
*ObjectsApi* | [**GetObject**](docs/api/ObjectsApi.md#getobject) | **Get** /{Bucket}/{Key} | GetObject
*ObjectsApi* | [**HeadObject**](docs/api/ObjectsApi.md#headobject) | **Head** /{Bucket}/{Key} | HeadObject
*ObjectsApi* | [**ListObjects**](docs/api/ObjectsApi.md#listobjects) | **Get** /{Bucket} | ListObjects
*ObjectsApi* | [**ListObjectsV2**](docs/api/ObjectsApi.md#listobjectsv2) | **Get** /{Bucket}?list-type&#x3D;2 | ListObjectsV2
*ObjectsApi* | [**OPTIONSObject**](docs/api/ObjectsApi.md#optionsobject) | **Options** /{Bucket} | OPTIONSObject
*ObjectsApi* | [**POSTObject**](docs/api/ObjectsApi.md#postobject) | **Post** /{Bucket}/{Key} | POSTObject
*ObjectsApi* | [**PutObject**](docs/api/ObjectsApi.md#putobject) | **Put** /{Bucket}/{Key} | PutObject
*PolicyApi* | [**DeleteBucketPolicy**](docs/api/PolicyApi.md#deletebucketpolicy) | **Delete** /{Bucket}?policy | DeleteBucketPolicy
*PolicyApi* | [**GetBucketPolicy**](docs/api/PolicyApi.md#getbucketpolicy) | **Get** /{Bucket}?policy | GetBucketPolicy
*PolicyApi* | [**GetBucketPolicyStatus**](docs/api/PolicyApi.md#getbucketpolicystatus) | **Get** /{Bucket}?policyStatus | GetBucketPolicyStatus
*PolicyApi* | [**PutBucketPolicy**](docs/api/PolicyApi.md#putbucketpolicy) | **Put** /{Bucket}?policy | PutBucketPolicy
*PublicAccessBlockApi* | [**DeletePublicAccessBlock**](docs/api/PublicAccessBlockApi.md#deletepublicaccessblock) | **Delete** /{Bucket}?publicAccessBlock | DeletePublicAccessBlock
*PublicAccessBlockApi* | [**GetPublicAccessBlock**](docs/api/PublicAccessBlockApi.md#getpublicaccessblock) | **Get** /{Bucket}?publicAccessBlock | GetPublicAccessBlock
*PublicAccessBlockApi* | [**PutPublicAccessBlock**](docs/api/PublicAccessBlockApi.md#putpublicaccessblock) | **Put** /{Bucket}?publicAccessBlock | PutPublicAccessBlock
*TaggingApi* | [**DeleteBucketTagging**](docs/api/TaggingApi.md#deletebuckettagging) | **Delete** /{Bucket}?tagging | DeleteBucketTagging
*TaggingApi* | [**DeleteObjectTagging**](docs/api/TaggingApi.md#deleteobjecttagging) | **Delete** /{Bucket}/{Key}?tagging | DeleteObjectTagging
*TaggingApi* | [**GetBucketTagging**](docs/api/TaggingApi.md#getbuckettagging) | **Get** /{Bucket}?tagging | GetBucketTagging
*TaggingApi* | [**GetObjectTagging**](docs/api/TaggingApi.md#getobjecttagging) | **Get** /{Bucket}/{Key}?tagging | GetObjectTagging
*TaggingApi* | [**PutBucketTagging**](docs/api/TaggingApi.md#putbuckettagging) | **Put** /{Bucket}?tagging | PutBucketTagging
*TaggingApi* | [**PutObjectTagging**](docs/api/TaggingApi.md#putobjecttagging) | **Put** /{Bucket}/{Key}?tagging | PutObjectTagging


## Documentation For Models

 - [BlockPublicAccessOutput](docs/models/BlockPublicAccessOutput.md)
 - [BlockPublicAccessPayload](docs/models/BlockPublicAccessPayload.md)
 - [Bucket](docs/models/Bucket.md)
 - [BucketLocation](docs/models/BucketLocation.md)
 - [BucketPolicy](docs/models/BucketPolicy.md)
 - [BucketPolicyCondition](docs/models/BucketPolicyCondition.md)
 - [BucketPolicyConditionDate](docs/models/BucketPolicyConditionDate.md)
 - [BucketPolicyConditionIpAddress](docs/models/BucketPolicyConditionIpAddress.md)
 - [BucketPolicyStatement](docs/models/BucketPolicyStatement.md)
 - [CommonPrefix](docs/models/CommonPrefix.md)
 - [CopyObjectRequest](docs/models/CopyObjectRequest.md)
 - [CopyObjectResult](docs/models/CopyObjectResult.md)
 - [CreateBucketConfiguration](docs/models/CreateBucketConfiguration.md)
 - [DeleteObjectsOutput](docs/models/DeleteObjectsOutput.md)
 - [DeleteObjectsRequest](docs/models/DeleteObjectsRequest.md)
 - [DeletedObject](docs/models/DeletedObject.md)
 - [DeletionError](docs/models/DeletionError.md)
 - [EncodingType](docs/models/EncodingType.md)
 - [Error](docs/models/Error.md)
 - [GetBucketTaggingOutput](docs/models/GetBucketTaggingOutput.md)
 - [GetObjectTaggingOutput](docs/models/GetObjectTaggingOutput.md)
 - [HeadObjectOutput](docs/models/HeadObjectOutput.md)
 - [ListAllMyBucketsResult](docs/models/ListAllMyBucketsResult.md)
 - [ListBucketResultV2](docs/models/ListBucketResultV2.md)
 - [ListObjectsOutput](docs/models/ListObjectsOutput.md)
 - [Object](docs/models/Object.md)
 - [ObjectIdentifier](docs/models/ObjectIdentifier.md)
 - [ObjectStorageClass](docs/models/ObjectStorageClass.md)
 - [Owner](docs/models/Owner.md)
 - [POSTObjectRequest](docs/models/POSTObjectRequest.md)
 - [PolicyStatus](docs/models/PolicyStatus.md)
 - [Principal](docs/models/Principal.md)
 - [PutBucketTaggingRequest](docs/models/PutBucketTaggingRequest.md)
 - [PutBucketTaggingRequestTagging](docs/models/PutBucketTaggingRequestTagging.md)
 - [PutObjectTaggingRequest](docs/models/PutObjectTaggingRequest.md)
 - [PutObjectTaggingRequestTagging](docs/models/PutObjectTaggingRequestTagging.md)
 - [Tag](docs/models/Tag.md)
 - [UploadPartRequest](docs/models/UploadPartRequest.md)


## Documentation For Authorization

For authentication you need to set the following environment variables with your credentials
```bash
export IONOS_S3_ACCESS_KEY="accesskey"
export IONOS_S3_SECRET_KEY="secretkey"
```

```golang
cfg := s3.NewConfigurationFromEnv()
apiClient:=s3.NewApiClient(cfg)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@cloud.ionos.com

