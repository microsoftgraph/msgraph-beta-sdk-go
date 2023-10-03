package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsValidateBulkResizeRequestBuilder provides operations to call the validateBulkResize method.
type VirtualEndpointCloudPCsValidateBulkResizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsValidateBulkResizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsValidateBulkResizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsValidateBulkResizeRequestBuilderInternal instantiates a new ValidateBulkResizeRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsValidateBulkResizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsValidateBulkResizeRequestBuilder) {
    m := &VirtualEndpointCloudPCsValidateBulkResizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/validateBulkResize", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsValidateBulkResizeRequestBuilder instantiates a new ValidateBulkResizeRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsValidateBulkResizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsValidateBulkResizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsValidateBulkResizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post validate that a set of cloudPC devices meet the requirements to be bulk resized. This API is supported in the following national cloud deployments.
// Deprecated: This method is obsolete. Use PostAsValidateBulkResizePostResponse instead.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-validatebulkresize?view=graph-rest-1.0
func (m *VirtualEndpointCloudPCsValidateBulkResizeRequestBuilder) Post(ctx context.Context, body VirtualEndpointCloudPCsValidateBulkResizePostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsValidateBulkResizeRequestBuilderPostRequestConfiguration)(VirtualEndpointCloudPCsValidateBulkResizeResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsValidateBulkResizeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsValidateBulkResizeResponseable), nil
}
// PostAsValidateBulkResizePostResponse validate that a set of cloudPC devices meet the requirements to be bulk resized. This API is supported in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-validatebulkresize?view=graph-rest-1.0
func (m *VirtualEndpointCloudPCsValidateBulkResizeRequestBuilder) PostAsValidateBulkResizePostResponse(ctx context.Context, body VirtualEndpointCloudPCsValidateBulkResizePostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsValidateBulkResizeRequestBuilderPostRequestConfiguration)(VirtualEndpointCloudPCsValidateBulkResizePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsValidateBulkResizePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsValidateBulkResizePostResponseable), nil
}
// ToPostRequestInformation validate that a set of cloudPC devices meet the requirements to be bulk resized. This API is supported in the following national cloud deployments.
func (m *VirtualEndpointCloudPCsValidateBulkResizeRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointCloudPCsValidateBulkResizePostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsValidateBulkResizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointCloudPCsValidateBulkResizeRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsValidateBulkResizeRequestBuilder) {
    return NewVirtualEndpointCloudPCsValidateBulkResizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
