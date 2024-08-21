# \ArithmeticOperationsAPI

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DividePost**](ArithmeticOperationsAPI.md#DividePost) | **Post** /divide | Divide a list of numbers
[**MultiplyPost**](ArithmeticOperationsAPI.md#MultiplyPost) | **Post** /multiply | Multiply a list of numbers
[**SubtractPost**](ArithmeticOperationsAPI.md#SubtractPost) | **Post** /subtract | Subtract a list of numbers
[**SumPost**](ArithmeticOperationsAPI.md#SumPost) | **Post** /sum | Sum a list of numbers



## DividePost

> DividePost200Response DividePost(ctx).Body(body).Execute()

Divide a list of numbers

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sec4you/SEC4YOU-LABS"
)

func main() {
	body := *openapiclient.NewDividePostRequest() // DividePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArithmeticOperationsAPI.DividePost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArithmeticOperationsAPI.DividePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DividePost`: DividePost200Response
	fmt.Fprintf(os.Stdout, "Response from `ArithmeticOperationsAPI.DividePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDividePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DividePostRequest**](DividePostRequest.md) |  | 

### Return type

[**DividePost200Response**](DividePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MultiplyPost

> MultiplyPost200Response MultiplyPost(ctx).Body(body).Execute()

Multiply a list of numbers

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sec4you/SEC4YOU-LABS"
)

func main() {
	body := *openapiclient.NewMultiplyPostRequest() // MultiplyPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArithmeticOperationsAPI.MultiplyPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArithmeticOperationsAPI.MultiplyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultiplyPost`: MultiplyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ArithmeticOperationsAPI.MultiplyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultiplyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MultiplyPostRequest**](MultiplyPostRequest.md) |  | 

### Return type

[**MultiplyPost200Response**](MultiplyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubtractPost

> SubtractPost200Response SubtractPost(ctx).Body(body).Execute()

Subtract a list of numbers

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sec4you/SEC4YOU-LABS"
)

func main() {
	body := *openapiclient.NewSubtractPostRequest() // SubtractPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArithmeticOperationsAPI.SubtractPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArithmeticOperationsAPI.SubtractPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubtractPost`: SubtractPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ArithmeticOperationsAPI.SubtractPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubtractPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SubtractPostRequest**](SubtractPostRequest.md) |  | 

### Return type

[**SubtractPost200Response**](SubtractPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SumPost

> SumPost200Response SumPost(ctx).Body(body).Execute()

Sum a list of numbers

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sec4you/SEC4YOU-LABS"
)

func main() {
	body := *openapiclient.NewSumPostRequest() // SumPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArithmeticOperationsAPI.SumPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArithmeticOperationsAPI.SumPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SumPost`: SumPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ArithmeticOperationsAPI.SumPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSumPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SumPostRequest**](SumPostRequest.md) |  | 

### Return type

[**SumPost200Response**](SumPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

