package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder provides operations to call the getIosAvailableUpdateVersions method.
type DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderGetQueryParameters invoke function getIosAvailableUpdateVersions
type DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderGetQueryParameters struct {
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
// DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderGetQueryParameters
}
// NewDeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderInternal instantiates a new DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder and sets the default values.
func NewDeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder) {
    m := &DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/getIosAvailableUpdateVersions(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder instantiates a new DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder and sets the default values.
func NewDeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getIosAvailableUpdateVersions
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a DeviceConfigurationsGetIosAvailableUpdateVersionsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderGetRequestConfiguration)(DeviceConfigurationsGetIosAvailableUpdateVersionsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceConfigurationsGetIosAvailableUpdateVersionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceConfigurationsGetIosAvailableUpdateVersionsResponseable), nil
}
// GetAsGetIosAvailableUpdateVersionsGetResponse invoke function getIosAvailableUpdateVersions
// returns a DeviceConfigurationsGetIosAvailableUpdateVersionsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder) GetAsGetIosAvailableUpdateVersionsGetResponse(ctx context.Context, requestConfiguration *DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderGetRequestConfiguration)(DeviceConfigurationsGetIosAvailableUpdateVersionsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceConfigurationsGetIosAvailableUpdateVersionsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceConfigurationsGetIosAvailableUpdateVersionsGetResponseable), nil
}
// ToGetRequestInformation invoke function getIosAvailableUpdateVersions
// returns a *RequestInformation when successful
func (m *DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder when successful
func (m *DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder) WithUrl(rawUrl string)(*DeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder) {
    return NewDeviceConfigurationsGetIosAvailableUpdateVersionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
