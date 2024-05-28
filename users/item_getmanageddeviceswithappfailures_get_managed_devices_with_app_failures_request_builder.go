package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder provides operations to call the getManagedDevicesWithAppFailures method.
type ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderGetQueryParameters retrieves the list of devices with failed apps
type ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderGetQueryParameters struct {
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
// ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderGetQueryParameters
}
// NewItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderInternal instantiates a new ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder and sets the default values.
func NewItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder) {
    m := &ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/getManagedDevicesWithAppFailures(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder instantiates a new ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder and sets the default values.
func NewItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderInternal(urlParams, requestAdapter)
}
// Get retrieves the list of devices with failed apps
// Deprecated: This method is obsolete. Use GetAsGetManagedDevicesWithAppFailuresGetResponse instead.
// returns a ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderGetRequestConfiguration)(ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresResponseable), nil
}
// GetAsGetManagedDevicesWithAppFailuresGetResponse retrieves the list of devices with failed apps
// returns a ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder) GetAsGetManagedDevicesWithAppFailuresGetResponse(ctx context.Context, requestConfiguration *ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderGetRequestConfiguration)(ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresGetResponseable), nil
}
// ToGetRequestInformation retrieves the list of devices with failed apps
// returns a *RequestInformation when successful
func (m *ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder when successful
func (m *ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder) WithUrl(rawUrl string)(*ItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder) {
    return NewItemGetmanageddeviceswithappfailuresGetManagedDevicesWithAppFailuresRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
