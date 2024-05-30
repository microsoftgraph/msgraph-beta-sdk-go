package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder provides operations to call the getLoggedOnManagedDevices method.
type ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderGetQueryParameters invoke function getLoggedOnManagedDevices
type ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderGetQueryParameters struct {
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
// ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderGetQueryParameters
}
// NewItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderInternal instantiates a new ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder and sets the default values.
func NewItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder) {
    m := &ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/getLoggedOnManagedDevices(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder instantiates a new ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder and sets the default values.
func NewItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getLoggedOnManagedDevices
// Deprecated: This method is obsolete. Use GetAsGetLoggedOnManagedDevicesGetResponse instead.
// returns a ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderGetRequestConfiguration)(ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesResponseable), nil
}
// GetAsGetLoggedOnManagedDevicesGetResponse invoke function getLoggedOnManagedDevices
// returns a ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder) GetAsGetLoggedOnManagedDevicesGetResponse(ctx context.Context, requestConfiguration *ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderGetRequestConfiguration)(ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesGetResponseable), nil
}
// ToGetRequestInformation invoke function getLoggedOnManagedDevices
// returns a *RequestInformation when successful
func (m *ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder when successful
func (m *ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder) WithUrl(rawUrl string)(*ItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder) {
    return NewItemGetloggedonmanageddevicesGetLoggedOnManagedDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
