package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder provides operations to call the getNonCompliantSettings method.
type ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetQueryParameters invoke function getNonCompliantSettings
type ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetQueryParameters struct {
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
// ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetQueryParameters
}
// NewItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal instantiates a new ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder and sets the default values.
func NewItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    m := &ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/getNonCompliantSettings(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder instantiates a new ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder and sets the default values.
func NewItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getNonCompliantSettings
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a ItemManagedDevicesItemGetNonCompliantSettingsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetRequestConfiguration)(ItemManagedDevicesItemGetNonCompliantSettingsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesItemGetNonCompliantSettingsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesItemGetNonCompliantSettingsResponseable), nil
}
// GetAsGetNonCompliantSettingsGetResponse invoke function getNonCompliantSettings
// returns a ItemManagedDevicesItemGetNonCompliantSettingsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder) GetAsGetNonCompliantSettingsGetResponse(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetRequestConfiguration)(ItemManagedDevicesItemGetNonCompliantSettingsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesItemGetNonCompliantSettingsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesItemGetNonCompliantSettingsGetResponseable), nil
}
// ToGetRequestInformation invoke function getNonCompliantSettings
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder when successful
func (m *ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewItemManagedDevicesItemGetNonCompliantSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
