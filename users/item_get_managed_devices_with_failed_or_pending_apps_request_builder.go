package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
type ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetQueryParameters retrieves the list of devices with failed or pending apps
type ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetQueryParameters struct {
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
// ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetQueryParameters
}
// NewItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal instantiates a new GetManagedDevicesWithFailedOrPendingAppsRequestBuilder and sets the default values.
func NewItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    m := &ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/getManagedDevicesWithFailedOrPendingApps(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder instantiates a new GetManagedDevicesWithFailedOrPendingAppsRequestBuilder and sets the default values.
func NewItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieves the list of devices with failed or pending apps
// Deprecated: This method is obsolete. Use GetAsGetManagedDevicesWithFailedOrPendingAppsGetResponse instead.
func (m *ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration)(ItemGetManagedDevicesWithFailedOrPendingAppsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetManagedDevicesWithFailedOrPendingAppsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetManagedDevicesWithFailedOrPendingAppsResponseable), nil
}
// GetAsGetManagedDevicesWithFailedOrPendingAppsGetResponse retrieves the list of devices with failed or pending apps
func (m *ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) GetAsGetManagedDevicesWithFailedOrPendingAppsGetResponse(ctx context.Context, requestConfiguration *ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration)(ItemGetManagedDevicesWithFailedOrPendingAppsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetManagedDevicesWithFailedOrPendingAppsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetManagedDevicesWithFailedOrPendingAppsGetResponseable), nil
}
// ToGetRequestInformation retrieves the list of devices with failed or pending apps
func (m *ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) WithUrl(rawUrl string)(*ItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return NewItemGetManagedDevicesWithFailedOrPendingAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
