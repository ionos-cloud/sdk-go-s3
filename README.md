# Go API client for ionoscloud

## Overview
The IONOS S3 Object Storage API for contract-owned buckets is a REST-based API that allows developers and applications to interact directly with
IONOS' scalable storage solution, leveraging the S3 protocol for object storage operations. Its design ensures seamless
compatibility with existing tools and libraries tailored for S3 systems.

### API References
- [S3 API Reference for contract-owned buckets](https://api.ionos.com/docs/s3-contract-owned-buckets/v2/)
### User documentation
[IONOS S3 Object Storage User Guide](https://docs.ionos.com/cloud/managed-services/s3-object-storage)
* [Documentation on S3 API Compatibility](https://docs.ionos.com/cloud/managed-services/s3-object-storage/concepts/s3-api-compatibility)

## Endpoints for contract-owned buckets
| Location | Region Name | Bucket Type | Endpoint |
| --- | --- | --- | --- |
| **Berlin, Germany** | **eu-central-3** | Contract-owned | `https://s3.eu-central-3.ionoscloud.com` |

## Changelog
- 30.05.2024 Initial version


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.2
- Package version: 1.1.0
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
*CORSApi* | [**DeleteBucketCors**](docs/api/CORSApi.md#deletebucketcors) | **Delete** /{Bucket}?cors | DeleteBucketCors
*CORSApi* | [**GetBucketCors**](docs/api/CORSApi.md#getbucketcors) | **Get** /{Bucket}?cors | GetBucketCors
*CORSApi* | [**PutBucketCors**](docs/api/CORSApi.md#putbucketcors) | **Put** /{Bucket}?cors | PutBucketCors
*EncryptionApi* | [**DeleteBucketEncryption**](docs/api/EncryptionApi.md#deletebucketencryption) | **Delete** /{Bucket}?encryption | DeleteBucketEncryption
*EncryptionApi* | [**GetBucketEncryption**](docs/api/EncryptionApi.md#getbucketencryption) | **Get** /{Bucket}?encryption | GetBucketEncryption
*EncryptionApi* | [**PutBucketEncryption**](docs/api/EncryptionApi.md#putbucketencryption) | **Put** /{Bucket}?encryption | PutBucketEncryption
*LifecycleApi* | [**DeleteBucketLifecycle**](docs/api/LifecycleApi.md#deletebucketlifecycle) | **Delete** /{Bucket}?lifecycle | DeleteBucketLifecycle
*LifecycleApi* | [**GetBucketLifecycle**](docs/api/LifecycleApi.md#getbucketlifecycle) | **Get** /{Bucket}?lifecycle | GetBucketLifecycle
*LifecycleApi* | [**PutBucketLifecycle**](docs/api/LifecycleApi.md#putbucketlifecycle) | **Put** /{Bucket}?lifecycle | PutBucketLifecycle
*ObjectLockApi* | [**GetObjectLegalHold**](docs/api/ObjectLockApi.md#getobjectlegalhold) | **Get** /{Bucket}/{Key}?legal-hold | GetObjectLegalHold
*ObjectLockApi* | [**GetObjectLockConfiguration**](docs/api/ObjectLockApi.md#getobjectlockconfiguration) | **Get** /{Bucket}?object-lock | GetObjectLockConfiguration
*ObjectLockApi* | [**GetObjectRetention**](docs/api/ObjectLockApi.md#getobjectretention) | **Get** /{Bucket}/{Key}?retention | GetObjectRetention
*ObjectLockApi* | [**PutObjectLegalHold**](docs/api/ObjectLockApi.md#putobjectlegalhold) | **Put** /{Bucket}/{Key}?legal-hold | PutObjectLegalHold
*ObjectLockApi* | [**PutObjectLockConfiguration**](docs/api/ObjectLockApi.md#putobjectlockconfiguration) | **Put** /{Bucket}?object-lock | PutObjectLockConfiguration
*ObjectLockApi* | [**PutObjectRetention**](docs/api/ObjectLockApi.md#putobjectretention) | **Put** /{Bucket}/{Key}?retention | PutObjectRetention
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
*ReplicationApi* | [**GetBucketReplication**](docs/api/ReplicationApi.md#getbucketreplication) | **Get** /{Bucket}?replication | GetBucketReplication
*TaggingApi* | [**DeleteBucketTagging**](docs/api/TaggingApi.md#deletebuckettagging) | **Delete** /{Bucket}?tagging | DeleteBucketTagging
*TaggingApi* | [**DeleteObjectTagging**](docs/api/TaggingApi.md#deleteobjecttagging) | **Delete** /{Bucket}/{Key}?tagging | DeleteObjectTagging
*TaggingApi* | [**GetBucketTagging**](docs/api/TaggingApi.md#getbuckettagging) | **Get** /{Bucket}?tagging | GetBucketTagging
*TaggingApi* | [**GetObjectTagging**](docs/api/TaggingApi.md#getobjecttagging) | **Get** /{Bucket}/{Key}?tagging | GetObjectTagging
*TaggingApi* | [**PutBucketTagging**](docs/api/TaggingApi.md#putbuckettagging) | **Put** /{Bucket}?tagging | PutBucketTagging
*TaggingApi* | [**PutObjectTagging**](docs/api/TaggingApi.md#putobjecttagging) | **Put** /{Bucket}/{Key}?tagging | PutObjectTagging
*UploadsApi* | [**AbortMultipartUpload**](docs/api/UploadsApi.md#abortmultipartupload) | **Delete** /{Bucket}/{Key}?uploadId | AbortMultipartUpload
*UploadsApi* | [**CompleteMultipartUpload**](docs/api/UploadsApi.md#completemultipartupload) | **Post** /{Bucket}/{Key}?uploadId | CompleteMultipartUpload
*UploadsApi* | [**CreateMultipartUpload**](docs/api/UploadsApi.md#createmultipartupload) | **Post** /{Bucket}/{Key}?uploads | CreateMultipartUpload
*UploadsApi* | [**ListMultipartUploads**](docs/api/UploadsApi.md#listmultipartuploads) | **Get** /{Bucket}?uploads | ListMultipartUploads
*UploadsApi* | [**ListParts**](docs/api/UploadsApi.md#listparts) | **Get** /{Bucket}/{Key}?uploadId | ListParts
*UploadsApi* | [**UploadPart**](docs/api/UploadsApi.md#uploadpart) | **Put** /{Bucket}/{Key}?uploadId | UploadPart
*UploadsApi* | [**UploadPartCopy**](docs/api/UploadsApi.md#uploadpartcopy) | **Put** /{Bucket}/{Key}?x-amz-copy-source&amp;partNumber&amp;uploadId | UploadPartCopy
*VersioningApi* | [**GetBucketVersioning**](docs/api/VersioningApi.md#getbucketversioning) | **Get** /{Bucket}?versioning | GetBucketVersioning
*VersioningApi* | [**PutBucketVersioning**](docs/api/VersioningApi.md#putbucketversioning) | **Put** /{Bucket}?versioning | PutBucketVersioning
*VersionsApi* | [**ListObjectVersions**](docs/api/VersionsApi.md#listobjectversions) | **Get** /{Bucket}?versions | ListObjectVersions
*WebsiteApi* | [**DeleteBucketWebsite**](docs/api/WebsiteApi.md#deletebucketwebsite) | **Delete** /{Bucket}?website | DeleteBucketWebsite
*WebsiteApi* | [**GetBucketWebsite**](docs/api/WebsiteApi.md#getbucketwebsite) | **Get** /{Bucket}?website | GetBucketWebsite
*WebsiteApi* | [**PutBucketWebsite**](docs/api/WebsiteApi.md#putbucketwebsite) | **Put** /{Bucket}?website | PutBucketWebsite


## Documentation For Models

 - [AbortIncompleteMultipartUpload](docs/models/AbortIncompleteMultipartUpload.md)
 - [BlockPublicAccessOutput](docs/models/BlockPublicAccessOutput.md)
 - [BlockPublicAccessPayload](docs/models/BlockPublicAccessPayload.md)
 - [Bucket](docs/models/Bucket.md)
 - [BucketLocation](docs/models/BucketLocation.md)
 - [BucketPolicy](docs/models/BucketPolicy.md)
 - [BucketPolicyCondition](docs/models/BucketPolicyCondition.md)
 - [BucketPolicyConditionDate](docs/models/BucketPolicyConditionDate.md)
 - [BucketPolicyConditionIpAddress](docs/models/BucketPolicyConditionIpAddress.md)
 - [BucketPolicyStatement](docs/models/BucketPolicyStatement.md)
 - [BucketVersioningStatus](docs/models/BucketVersioningStatus.md)
 - [CORSRule](docs/models/CORSRule.md)
 - [CSVInput](docs/models/CSVInput.md)
 - [CSVOutput](docs/models/CSVOutput.md)
 - [CommonPrefix](docs/models/CommonPrefix.md)
 - [CompleteMultipartUploadOutput](docs/models/CompleteMultipartUploadOutput.md)
 - [CompletedPart](docs/models/CompletedPart.md)
 - [CopyObjectRequest](docs/models/CopyObjectRequest.md)
 - [CopyObjectResult](docs/models/CopyObjectResult.md)
 - [CopyPartResult](docs/models/CopyPartResult.md)
 - [CreateBucketConfiguration](docs/models/CreateBucketConfiguration.md)
 - [CreateMultipartUploadOutput](docs/models/CreateMultipartUploadOutput.md)
 - [DefaultRetention](docs/models/DefaultRetention.md)
 - [DeleteMarkerEntry](docs/models/DeleteMarkerEntry.md)
 - [DeleteObjectsOutput](docs/models/DeleteObjectsOutput.md)
 - [DeleteObjectsRequest](docs/models/DeleteObjectsRequest.md)
 - [DeletedObject](docs/models/DeletedObject.md)
 - [DeletionError](docs/models/DeletionError.md)
 - [Destination](docs/models/Destination.md)
 - [EncodingType](docs/models/EncodingType.md)
 - [Encryption](docs/models/Encryption.md)
 - [Error](docs/models/Error.md)
 - [ErrorDocument](docs/models/ErrorDocument.md)
 - [Example](docs/models/Example.md)
 - [ExampleCompleteMultipartUpload](docs/models/ExampleCompleteMultipartUpload.md)
 - [ExpirationStatus](docs/models/ExpirationStatus.md)
 - [ExpressionType](docs/models/ExpressionType.md)
 - [GetBucketCorsOutput](docs/models/GetBucketCorsOutput.md)
 - [GetBucketLifecycleOutput](docs/models/GetBucketLifecycleOutput.md)
 - [GetBucketReplicationOutput](docs/models/GetBucketReplicationOutput.md)
 - [GetBucketTaggingOutput](docs/models/GetBucketTaggingOutput.md)
 - [GetBucketVersioningOutput](docs/models/GetBucketVersioningOutput.md)
 - [GetBucketWebsiteOutput](docs/models/GetBucketWebsiteOutput.md)
 - [GetObjectLockConfigurationOutput](docs/models/GetObjectLockConfigurationOutput.md)
 - [GetObjectTaggingOutput](docs/models/GetObjectTaggingOutput.md)
 - [HeadObjectOutput](docs/models/HeadObjectOutput.md)
 - [IndexDocument](docs/models/IndexDocument.md)
 - [Initiator](docs/models/Initiator.md)
 - [InputSerialization](docs/models/InputSerialization.md)
 - [InputSerializationJSON](docs/models/InputSerializationJSON.md)
 - [JSONOutput](docs/models/JSONOutput.md)
 - [LifecycleExpiration](docs/models/LifecycleExpiration.md)
 - [ListAllMyBucketsResult](docs/models/ListAllMyBucketsResult.md)
 - [ListBucketResultV2](docs/models/ListBucketResultV2.md)
 - [ListMultipartUploadsOutput](docs/models/ListMultipartUploadsOutput.md)
 - [ListObjectVersionsOutput](docs/models/ListObjectVersionsOutput.md)
 - [ListObjectsOutput](docs/models/ListObjectsOutput.md)
 - [ListPartsOutput](docs/models/ListPartsOutput.md)
 - [MetadataEntry](docs/models/MetadataEntry.md)
 - [MfaDeleteStatus](docs/models/MfaDeleteStatus.md)
 - [MultipartUpload](docs/models/MultipartUpload.md)
 - [NoncurrentVersionExpiration](docs/models/NoncurrentVersionExpiration.md)
 - [Object](docs/models/Object.md)
 - [ObjectIdentifier](docs/models/ObjectIdentifier.md)
 - [ObjectLegalHoldConfiguration](docs/models/ObjectLegalHoldConfiguration.md)
 - [ObjectLockRetention](docs/models/ObjectLockRetention.md)
 - [ObjectLockRule](docs/models/ObjectLockRule.md)
 - [ObjectStorageClass](docs/models/ObjectStorageClass.md)
 - [ObjectVersion](docs/models/ObjectVersion.md)
 - [ObjectVersionStorageClass](docs/models/ObjectVersionStorageClass.md)
 - [OutputSerialization](docs/models/OutputSerialization.md)
 - [Owner](docs/models/Owner.md)
 - [POSTObjectRequest](docs/models/POSTObjectRequest.md)
 - [Part](docs/models/Part.md)
 - [PolicyStatus](docs/models/PolicyStatus.md)
 - [Principal](docs/models/Principal.md)
 - [PrincipalAllOf](docs/models/PrincipalAllOf.md)
 - [PutBucketCorsRequest](docs/models/PutBucketCorsRequest.md)
 - [PutBucketEncryptionRequest](docs/models/PutBucketEncryptionRequest.md)
 - [PutBucketLifecycleRequest](docs/models/PutBucketLifecycleRequest.md)
 - [PutBucketTaggingRequest](docs/models/PutBucketTaggingRequest.md)
 - [PutBucketVersioningRequest](docs/models/PutBucketVersioningRequest.md)
 - [PutBucketWebsiteRequest](docs/models/PutBucketWebsiteRequest.md)
 - [PutObjectLockConfigurationRequest](docs/models/PutObjectLockConfigurationRequest.md)
 - [PutObjectLockConfigurationRequestRule](docs/models/PutObjectLockConfigurationRequestRule.md)
 - [PutObjectRetentionRequest](docs/models/PutObjectRetentionRequest.md)
 - [PutObjectTaggingRequest](docs/models/PutObjectTaggingRequest.md)
 - [Redirect](docs/models/Redirect.md)
 - [RedirectAllRequestsTo](docs/models/RedirectAllRequestsTo.md)
 - [ReplicaModificationsStatus](docs/models/ReplicaModificationsStatus.md)
 - [ReplicationConfiguration](docs/models/ReplicationConfiguration.md)
 - [ReplicationRule](docs/models/ReplicationRule.md)
 - [RoutingRule](docs/models/RoutingRule.md)
 - [RoutingRuleCondition](docs/models/RoutingRuleCondition.md)
 - [Rule](docs/models/Rule.md)
 - [ServerSideEncryption](docs/models/ServerSideEncryption.md)
 - [ServerSideEncryptionByDefault](docs/models/ServerSideEncryptionByDefault.md)
 - [ServerSideEncryptionConfiguration](docs/models/ServerSideEncryptionConfiguration.md)
 - [ServerSideEncryptionRule](docs/models/ServerSideEncryptionRule.md)
 - [StorageClass](docs/models/StorageClass.md)
 - [Tag](docs/models/Tag.md)
 - [UploadPartCopyOutput](docs/models/UploadPartCopyOutput.md)
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

