package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder provides operations to call the getStorageAccounts method.
type VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderGetQueryParameters invoke function getStorageAccounts
type VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderGetQueryParameters struct {
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
// VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderGetQueryParameters
}
// NewVirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderInternal instantiates a new GetStorageAccountsWithSubscriptionIdRequestBuilder and sets the default values.
func NewVirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, subscriptionId *string)(*VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder) {
    m := &VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/snapshots/getStorageAccounts(subscriptionId='{subscriptionId}'){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    if subscriptionId != nil {
        m.BaseRequestBuilder.PathParameters["subscriptionId"] = *subscriptionId
    }
    return m
}
// NewVirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder instantiates a new GetStorageAccountsWithSubscriptionIdRequestBuilder and sets the default values.
func NewVirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getStorageAccounts
// Deprecated: This method is obsolete. Use GetAsGetStorageAccountsWithSubscriptionIdGetResponse instead.
func (m *VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration)(VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdResponseable), nil
}
// GetAsGetStorageAccountsWithSubscriptionIdGetResponse invoke function getStorageAccounts
func (m *VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder) GetAsGetStorageAccountsWithSubscriptionIdGetResponse(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration)(VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdGetResponseable), nil
}
// ToGetRequestInformation invoke function getStorageAccounts
func (m *VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder) {
    return NewVirtualEndpointSnapshotsGetStorageAccountsWithSubscriptionIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
