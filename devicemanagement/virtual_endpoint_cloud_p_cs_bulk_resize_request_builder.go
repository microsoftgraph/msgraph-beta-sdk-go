package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsBulkResizeRequestBuilder provides operations to call the bulkResize method.
type VirtualEndpointCloudPCsBulkResizeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsBulkResizeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsBulkResizeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointCloudPCsBulkResizeRequestBuilderInternal instantiates a new BulkResizeRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsBulkResizeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsBulkResizeRequestBuilder) {
    m := &VirtualEndpointCloudPCsBulkResizeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/bulkResize", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsBulkResizeRequestBuilder instantiates a new BulkResizeRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsBulkResizeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsBulkResizeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsBulkResizeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post perform a bulk resize action to resize a group of cloudPCs that have successfully passed validation. If any devices can't be resized, those devices indicate 'resize failed'. The remaining devices are provisioned for the resize process. This API is available in the following national cloud deployments.
// Deprecated: This method is obsolete. Use PostAsBulkResizePostResponse instead.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-bulkresize?view=graph-rest-1.0
func (m *VirtualEndpointCloudPCsBulkResizeRequestBuilder) Post(ctx context.Context, body VirtualEndpointCloudPCsBulkResizePostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsBulkResizeRequestBuilderPostRequestConfiguration)(VirtualEndpointCloudPCsBulkResizeResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsBulkResizeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsBulkResizeResponseable), nil
}
// PostAsBulkResizePostResponse perform a bulk resize action to resize a group of cloudPCs that have successfully passed validation. If any devices can't be resized, those devices indicate 'resize failed'. The remaining devices are provisioned for the resize process. This API is available in the following national cloud deployments.
// Deprecated: The bulkResize action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkResize on 2023-05-24 and will be removed 2023-09-24
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-bulkresize?view=graph-rest-1.0
func (m *VirtualEndpointCloudPCsBulkResizeRequestBuilder) PostAsBulkResizePostResponse(ctx context.Context, body VirtualEndpointCloudPCsBulkResizePostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsBulkResizeRequestBuilderPostRequestConfiguration)(VirtualEndpointCloudPCsBulkResizePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsBulkResizePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsBulkResizePostResponseable), nil
}
// ToPostRequestInformation perform a bulk resize action to resize a group of cloudPCs that have successfully passed validation. If any devices can't be resized, those devices indicate 'resize failed'. The remaining devices are provisioned for the resize process. This API is available in the following national cloud deployments.
// Deprecated: The bulkResize action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkResize on 2023-05-24 and will be removed 2023-09-24
func (m *VirtualEndpointCloudPCsBulkResizeRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointCloudPCsBulkResizePostRequestBodyable, requestConfiguration *VirtualEndpointCloudPCsBulkResizeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The bulkResize action is deprecated and will stop supporting on September 24, 2023. Please use bulk action entity api. as of 2023-05/bulkResize on 2023-05-24 and will be removed 2023-09-24
func (m *VirtualEndpointCloudPCsBulkResizeRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsBulkResizeRequestBuilder) {
    return NewVirtualEndpointCloudPCsBulkResizeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
