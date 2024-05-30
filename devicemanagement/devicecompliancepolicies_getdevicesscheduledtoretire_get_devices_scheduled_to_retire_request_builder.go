package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder provides operations to call the getDevicesScheduledToRetire method.
type DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderGetQueryParameters invoke function getDevicesScheduledToRetire
type DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderGetQueryParameters struct {
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
// DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderGetQueryParameters
}
// NewDevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderInternal instantiates a new DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder) {
    m := &DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/getDevicesScheduledToRetire(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder instantiates a new DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getDevicesScheduledToRetire
// Deprecated: This method is obsolete. Use GetAsGetDevicesScheduledToRetireGetResponse instead.
// returns a DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderGetRequestConfiguration)(DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireResponseable), nil
}
// GetAsGetDevicesScheduledToRetireGetResponse invoke function getDevicesScheduledToRetire
// returns a DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder) GetAsGetDevicesScheduledToRetireGetResponse(ctx context.Context, requestConfiguration *DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderGetRequestConfiguration)(DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireGetResponseable), nil
}
// ToGetRequestInformation invoke function getDevicesScheduledToRetire
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder when successful
func (m *DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder) {
    return NewDevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
