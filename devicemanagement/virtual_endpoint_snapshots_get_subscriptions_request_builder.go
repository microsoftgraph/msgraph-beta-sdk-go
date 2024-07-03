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
// VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetQueryParameters list all subscriptions cloudPcSubscription that can be used to store a snapshot or snapshots of a Cloud PC for forensic analysis.
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
// NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilderInternal instantiates a new VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder and sets the default values.
func NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) {
    m := &VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/snapshots/getSubscriptions(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilder instantiates a new VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder and sets the default values.
func NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all subscriptions cloudPcSubscription that can be used to store a snapshot or snapshots of a Cloud PC for forensic analysis.
// Deprecated: This method is obsolete. Use GetAsGetSubscriptionsGetResponse instead.
// returns a VirtualEndpointSnapshotsGetSubscriptionsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcsnapshot-getsubscriptions?view=graph-rest-beta
func (m *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetRequestConfiguration)(VirtualEndpointSnapshotsGetSubscriptionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// GetAsGetSubscriptionsGetResponse list all subscriptions cloudPcSubscription that can be used to store a snapshot or snapshots of a Cloud PC for forensic analysis.
// returns a VirtualEndpointSnapshotsGetSubscriptionsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcsnapshot-getsubscriptions?view=graph-rest-beta
func (m *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) GetAsGetSubscriptionsGetResponse(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetRequestConfiguration)(VirtualEndpointSnapshotsGetSubscriptionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation list all subscriptions cloudPcSubscription that can be used to store a snapshot or snapshots of a Cloud PC for forensic analysis.
// returns a *RequestInformation when successful
func (m *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder when successful
func (m *VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointSnapshotsGetSubscriptionsRequestBuilder) {
    return NewVirtualEndpointSnapshotsGetSubscriptionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
