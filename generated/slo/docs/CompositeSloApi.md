# \CompositeSloAPI

All URIs are relative to *http://localhost:5601*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompositeSloOp**](CompositeSloAPI.md#CreateCompositeSloOp) | **Post** /s/{spaceId}/api/observability/composite_slos | Creates a Composite SLO
[**DeleteCompositeSloOp**](CompositeSloAPI.md#DeleteCompositeSloOp) | **Delete** /s/{spaceId}/api/observability/composite_slos/{compositeSloId} | Deletes a composite SLO
[**FindCompositeSloOp**](CompositeSloAPI.md#FindCompositeSloOp) | **Get** /s/{spaceId}/api/observability/composite_slos | Retrieves a paginated list of composite SLOs with summary
[**GetCompositeSloOp**](CompositeSloAPI.md#GetCompositeSloOp) | **Get** /s/{spaceId}/api/observability/composite_slos/{compositeSloId} | Retrieves a composite SLO
[**UpdateCompositeSloOp**](CompositeSloAPI.md#UpdateCompositeSloOp) | **Put** /s/{spaceId}/api/observability/composite_slos/{compositeSloId} | Updates a composite SLO



## CreateCompositeSloOp

> CreateCompositeSloResponse CreateCompositeSloOp(ctx, spaceId).KbnXsrf(kbnXsrf).CreateCompositeSloRequest(createCompositeSloRequest).Execute()

Creates a Composite SLO



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/slo"
)

func main() {
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    spaceId := "default" // string | An identifier for the space. If `/s/` and the identifier are omitted from the path, the default space is used.
    createCompositeSloRequest := *openapiclient.NewCreateCompositeSloRequest("Name_example", *openapiclient.NewTimeWindow("28d", "rolling"), openapiclient.budgeting_method("occurrences"), openapiclient.composite_method("weightedAverage"), *openapiclient.NewObjective(float32(0.99)), []openapiclient.WeightedCompositeSourcesInner{*openapiclient.NewWeightedCompositeSourcesInner("8853df00-ae2e-11ed-90af-09bb6422b258", float32(2), float32(3))}) // CreateCompositeSloRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompositeSloAPI.CreateCompositeSloOp(context.Background(), spaceId).KbnXsrf(kbnXsrf).CreateCompositeSloRequest(createCompositeSloRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositeSloAPI.CreateCompositeSloOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompositeSloOp`: CreateCompositeSloResponse
    fmt.Fprintf(os.Stdout, "Response from `CompositeSloAPI.CreateCompositeSloOp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | An identifier for the space. If &#x60;/s/&#x60; and the identifier are omitted from the path, the default space is used. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompositeSloOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbnXsrf** | **string** | Cross-site request forgery protection | 

 **createCompositeSloRequest** | [**CreateCompositeSloRequest**](CreateCompositeSloRequest.md) |  | 

### Return type

[**CreateCompositeSloResponse**](CreateCompositeSloResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompositeSloOp

> DeleteCompositeSloOp(ctx, spaceId, compositeSloId).KbnXsrf(kbnXsrf).Execute()

Deletes a composite SLO



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/slo"
)

func main() {
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    spaceId := "default" // string | An identifier for the space. If `/s/` and the identifier are omitted from the path, the default space is used.
    compositeSloId := "9c235211-6834-11ea-a78c-6feb38a34414" // string | An identifier for the composite slo.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CompositeSloAPI.DeleteCompositeSloOp(context.Background(), spaceId, compositeSloId).KbnXsrf(kbnXsrf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositeSloAPI.DeleteCompositeSloOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | An identifier for the space. If &#x60;/s/&#x60; and the identifier are omitted from the path, the default space is used. | 
**compositeSloId** | **string** | An identifier for the composite slo. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompositeSloOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbnXsrf** | **string** | Cross-site request forgery protection | 



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCompositeSloOp

> FindCompositeSloResponse FindCompositeSloOp(ctx, spaceId).KbnXsrf(kbnXsrf).Page(page).PerPage(perPage).SortBy(sortBy).SortDirection(sortDirection).Execute()

Retrieves a paginated list of composite SLOs with summary



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/slo"
)

func main() {
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    spaceId := "default" // string | An identifier for the space. If `/s/` and the identifier are omitted from the path, the default space is used.
    page := int32(1) // int32 | The page number to return (optional) (default to 1)
    perPage := int32(20) // int32 | The number of SLOs to return per page (optional) (default to 25)
    sortBy := "creationTime" // string | Sort by field (optional) (default to "creationTime")
    sortDirection := "asc" // string | Sort order (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompositeSloAPI.FindCompositeSloOp(context.Background(), spaceId).KbnXsrf(kbnXsrf).Page(page).PerPage(perPage).SortBy(sortBy).SortDirection(sortDirection).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositeSloAPI.FindCompositeSloOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCompositeSloOp`: FindCompositeSloResponse
    fmt.Fprintf(os.Stdout, "Response from `CompositeSloAPI.FindCompositeSloOp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | An identifier for the space. If &#x60;/s/&#x60; and the identifier are omitted from the path, the default space is used. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindCompositeSloOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbnXsrf** | **string** | Cross-site request forgery protection | 

 **page** | **int32** | The page number to return | [default to 1]
 **perPage** | **int32** | The number of SLOs to return per page | [default to 25]
 **sortBy** | **string** | Sort by field | [default to &quot;creationTime&quot;]
 **sortDirection** | **string** | Sort order | [default to &quot;asc&quot;]

### Return type

[**FindCompositeSloResponse**](FindCompositeSloResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompositeSloOp

> CompositeSloResponse GetCompositeSloOp(ctx, spaceId, compositeSloId).KbnXsrf(kbnXsrf).Execute()

Retrieves a composite SLO



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/slo"
)

func main() {
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    spaceId := "default" // string | An identifier for the space. If `/s/` and the identifier are omitted from the path, the default space is used.
    compositeSloId := "9c235211-6834-11ea-a78c-6feb38a34414" // string | An identifier for the composite slo.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompositeSloAPI.GetCompositeSloOp(context.Background(), spaceId, compositeSloId).KbnXsrf(kbnXsrf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositeSloAPI.GetCompositeSloOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompositeSloOp`: CompositeSloResponse
    fmt.Fprintf(os.Stdout, "Response from `CompositeSloAPI.GetCompositeSloOp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | An identifier for the space. If &#x60;/s/&#x60; and the identifier are omitted from the path, the default space is used. | 
**compositeSloId** | **string** | An identifier for the composite slo. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompositeSloOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbnXsrf** | **string** | Cross-site request forgery protection | 



### Return type

[**CompositeSloResponse**](CompositeSloResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompositeSloOp

> BaseCompositeSloResponse UpdateCompositeSloOp(ctx, spaceId, compositeSloId).KbnXsrf(kbnXsrf).UpdateCompositeSloRequest(updateCompositeSloRequest).Execute()

Updates a composite SLO



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/slo"
)

func main() {
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    spaceId := "default" // string | An identifier for the space. If `/s/` and the identifier are omitted from the path, the default space is used.
    compositeSloId := "9c235211-6834-11ea-a78c-6feb38a34414" // string | An identifier for the composite slo.
    updateCompositeSloRequest := *openapiclient.NewUpdateCompositeSloRequest() // UpdateCompositeSloRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompositeSloAPI.UpdateCompositeSloOp(context.Background(), spaceId, compositeSloId).KbnXsrf(kbnXsrf).UpdateCompositeSloRequest(updateCompositeSloRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositeSloAPI.UpdateCompositeSloOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCompositeSloOp`: BaseCompositeSloResponse
    fmt.Fprintf(os.Stdout, "Response from `CompositeSloAPI.UpdateCompositeSloOp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** | An identifier for the space. If &#x60;/s/&#x60; and the identifier are omitted from the path, the default space is used. | 
**compositeSloId** | **string** | An identifier for the composite slo. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompositeSloOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbnXsrf** | **string** | Cross-site request forgery protection | 


 **updateCompositeSloRequest** | [**UpdateCompositeSloRequest**](UpdateCompositeSloRequest.md) |  | 

### Return type

[**BaseCompositeSloResponse**](BaseCompositeSloResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

