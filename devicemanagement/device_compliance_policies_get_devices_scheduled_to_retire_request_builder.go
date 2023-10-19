package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder provides operations to call the getDevicesScheduledToRetire method.
type DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderGetQueryParameters invoke function getDevicesScheduledToRetire
type DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderGetQueryParameters struct {
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
// DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderGetQueryParameters
}
// NewDeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderInternal instantiates a new GetDevicesScheduledToRetireRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder) {
    m := &DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/getDevicesScheduledToRetire(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewDeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder instantiates a new GetDevicesScheduledToRetireRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getDevicesScheduledToRetire
// Deprecated: This method is obsolete. Use GetAsGetDevicesScheduledToRetireGetResponse instead.
func (m *DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderGetRequestConfiguration)(DeviceCompliancePoliciesGetDevicesScheduledToRetireResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceCompliancePoliciesGetDevicesScheduledToRetireResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceCompliancePoliciesGetDevicesScheduledToRetireResponseable), nil
}
// GetAsGetDevicesScheduledToRetireGetResponse invoke function getDevicesScheduledToRetire
func (m *DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder) GetAsGetDevicesScheduledToRetireGetResponse(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderGetRequestConfiguration)(DeviceCompliancePoliciesGetDevicesScheduledToRetireGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceCompliancePoliciesGetDevicesScheduledToRetireGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceCompliancePoliciesGetDevicesScheduledToRetireGetResponseable), nil
}
// ToGetRequestInformation invoke function getDevicesScheduledToRetire
func (m *DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder) WithUrl(rawUrl string)(*DeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder) {
    return NewDeviceCompliancePoliciesGetDevicesScheduledToRetireRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
