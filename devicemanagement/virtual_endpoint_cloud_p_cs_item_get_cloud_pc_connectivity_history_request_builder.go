package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder provides operations to call the getCloudPcConnectivityHistory method.
type VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetQueryParameters invoke function getCloudPcConnectivityHistory
type VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetQueryParameters struct {
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
// VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetQueryParameters
}
// NewVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal instantiates a new VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    m := &VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/getCloudPcConnectivityHistory(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder instantiates a new VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getCloudPcConnectivityHistory
// Deprecated: This method is obsolete. Use GetAsGetCloudPcConnectivityHistoryGetResponse instead.
// returns a VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryResponseable), nil
}
// GetAsGetCloudPcConnectivityHistoryGetResponse invoke function getCloudPcConnectivityHistory
// returns a VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) GetAsGetCloudPcConnectivityHistoryGetResponse(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration)(VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryGetResponseable), nil
}
// ToGetRequestInformation invoke function getCloudPcConnectivityHistory
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder when successful
func (m *VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
