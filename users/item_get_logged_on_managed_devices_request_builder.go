package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetLoggedOnManagedDevicesRequestBuilder provides operations to call the getLoggedOnManagedDevices method.
type ItemGetLoggedOnManagedDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetLoggedOnManagedDevicesRequestBuilderGetQueryParameters invoke function getLoggedOnManagedDevices
type ItemGetLoggedOnManagedDevicesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemGetLoggedOnManagedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetLoggedOnManagedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetLoggedOnManagedDevicesRequestBuilderGetQueryParameters
}
// NewItemGetLoggedOnManagedDevicesRequestBuilderInternal instantiates a new ItemGetLoggedOnManagedDevicesRequestBuilder and sets the default values.
func NewItemGetLoggedOnManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetLoggedOnManagedDevicesRequestBuilder) {
    m := &ItemGetLoggedOnManagedDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/getLoggedOnManagedDevices(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemGetLoggedOnManagedDevicesRequestBuilder instantiates a new ItemGetLoggedOnManagedDevicesRequestBuilder and sets the default values.
func NewItemGetLoggedOnManagedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetLoggedOnManagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetLoggedOnManagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getLoggedOnManagedDevices
// Deprecated: This method is obsolete. Use GetAsGetLoggedOnManagedDevicesGetResponse instead.
// returns a ItemGetLoggedOnManagedDevicesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetLoggedOnManagedDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetLoggedOnManagedDevicesRequestBuilderGetRequestConfiguration)(ItemGetLoggedOnManagedDevicesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetLoggedOnManagedDevicesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetLoggedOnManagedDevicesResponseable), nil
}
// GetAsGetLoggedOnManagedDevicesGetResponse invoke function getLoggedOnManagedDevices
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ItemGetLoggedOnManagedDevicesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetLoggedOnManagedDevicesRequestBuilder) GetAsGetLoggedOnManagedDevicesGetResponse(ctx context.Context, requestConfiguration *ItemGetLoggedOnManagedDevicesRequestBuilderGetRequestConfiguration)(ItemGetLoggedOnManagedDevicesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetLoggedOnManagedDevicesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetLoggedOnManagedDevicesGetResponseable), nil
}
// ToGetRequestInformation invoke function getLoggedOnManagedDevices
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemGetLoggedOnManagedDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetLoggedOnManagedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemGetLoggedOnManagedDevicesRequestBuilder when successful
func (m *ItemGetLoggedOnManagedDevicesRequestBuilder) WithUrl(rawUrl string)(*ItemGetLoggedOnManagedDevicesRequestBuilder) {
    return NewItemGetLoggedOnManagedDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
