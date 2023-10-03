package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder provides operations to call the getSubscriptions method.
type VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetQueryParameters invoke function getSubscriptions
type VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetQueryParameters struct {
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
// VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetQueryParameters
}
// NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilderInternal instantiates a new GetSubscriptionsRequestBuilder and sets the default values.
func NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) {
    m := &VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/snapshots/getSubscriptions(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilder instantiates a new GetSubscriptionsRequestBuilder and sets the default values.
func NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getSubscriptions
// Deprecated: This method is obsolete. Use GetAsGetSubscriptionsGetResponse instead.
func (m *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetRequestConfiguration)(VirtualEndpointSnapshotsGetSubscriptionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointSnapshotsGetSubscriptionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointSnapshotsGetSubscriptionsResponseable), nil
}
// GetAsGetSubscriptionsGetResponse invoke function getSubscriptions
func (m *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) GetAsGetSubscriptionsGetResponse(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetRequestConfiguration)(VirtualEndpointSnapshotsGetSubscriptionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointSnapshotsGetSubscriptionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointSnapshotsGetSubscriptionsGetResponseable), nil
}
// ToGetRequestInformation invoke function getSubscriptions
func (m *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) {
    return NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
