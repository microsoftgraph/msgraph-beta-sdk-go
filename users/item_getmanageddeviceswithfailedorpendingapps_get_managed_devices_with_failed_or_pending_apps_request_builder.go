package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
type ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetQueryParameters retrieves the list of devices with failed or pending apps
type ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetQueryParameters struct {
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
// ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetQueryParameters
}
// NewItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal instantiates a new ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder and sets the default values.
func NewItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    m := &ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/getManagedDevicesWithFailedOrPendingApps(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder instantiates a new ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder and sets the default values.
func NewItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieves the list of devices with failed or pending apps
// Deprecated: This method is obsolete. Use GetAsGetManagedDevicesWithFailedOrPendingAppsGetResponse instead.
// returns a ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration)(ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsResponseable), nil
}
// GetAsGetManagedDevicesWithFailedOrPendingAppsGetResponse retrieves the list of devices with failed or pending apps
// returns a ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) GetAsGetManagedDevicesWithFailedOrPendingAppsGetResponse(ctx context.Context, requestConfiguration *ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration)(ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsGetResponseable), nil
}
// ToGetRequestInformation retrieves the list of devices with failed or pending apps
// returns a *RequestInformation when successful
func (m *ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder when successful
func (m *ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) WithUrl(rawUrl string)(*ItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    return NewItemGetmanageddeviceswithfailedorpendingappsGetManagedDevicesWithFailedOrPendingAppsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
