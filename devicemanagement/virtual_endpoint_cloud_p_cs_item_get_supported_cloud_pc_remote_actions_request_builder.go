package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder provides operations to call the getSupportedCloudPcRemoteActions method.
type VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters invoke function getSupportedCloudPcRemoteActions
type VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters
}
// NewVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderInternal instantiates a new GetSupportedCloudPcRemoteActionsRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/getSupportedCloudPcRemoteActions(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder instantiates a new GetSupportedCloudPcRemoteActionsRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getSupportedCloudPcRemoteActions
// Deprecated: This method is obsolete. Use GetAsGetSupportedCloudPcRemoteActionsGetResponse instead.
func (m *VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsResponseable), nil
}
// GetAsGetSupportedCloudPcRemoteActionsGetResponse invoke function getSupportedCloudPcRemoteActions
func (m *VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) GetAsGetSupportedCloudPcRemoteActionsGetResponse(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsGetResponseable), nil
}
// ToGetRequestInformation invoke function getSupportedCloudPcRemoteActions
func (m *VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
