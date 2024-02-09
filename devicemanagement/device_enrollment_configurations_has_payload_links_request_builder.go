package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder provides operations to call the hasPayloadLinks method.
type DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilderInternal instantiates a new DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder and sets the default values.
func NewDeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder) {
    m := &DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceEnrollmentConfigurations/hasPayloadLinks", pathParameters),
    }
    return m
}
// NewDeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder instantiates a new DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder and sets the default values.
func NewDeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasPayloadLinks
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a DeviceEnrollmentConfigurationsHasPayloadLinksResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder) Post(ctx context.Context, body DeviceEnrollmentConfigurationsHasPayloadLinksPostRequestBodyable, requestConfiguration *DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilderPostRequestConfiguration)(DeviceEnrollmentConfigurationsHasPayloadLinksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceEnrollmentConfigurationsHasPayloadLinksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceEnrollmentConfigurationsHasPayloadLinksResponseable), nil
}
// PostAsHasPayloadLinksPostResponse invoke action hasPayloadLinks
// returns a DeviceEnrollmentConfigurationsHasPayloadLinksPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder) PostAsHasPayloadLinksPostResponse(ctx context.Context, body DeviceEnrollmentConfigurationsHasPayloadLinksPostRequestBodyable, requestConfiguration *DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilderPostRequestConfiguration)(DeviceEnrollmentConfigurationsHasPayloadLinksPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceEnrollmentConfigurationsHasPayloadLinksPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceEnrollmentConfigurationsHasPayloadLinksPostResponseable), nil
}
// ToPostRequestInformation invoke action hasPayloadLinks
// returns a *RequestInformation when successful
func (m *DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceEnrollmentConfigurationsHasPayloadLinksPostRequestBodyable, requestConfiguration *DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder when successful
func (m *DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder) WithUrl(rawUrl string)(*DeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder) {
    return NewDeviceEnrollmentConfigurationsHasPayloadLinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
