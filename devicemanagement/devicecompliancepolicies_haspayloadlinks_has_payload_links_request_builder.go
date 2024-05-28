package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder provides operations to call the hasPayloadLinks method.
type DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilderInternal instantiates a new DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) {
    m := &DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/hasPayloadLinks", pathParameters),
    }
    return m
}
// NewDevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder instantiates a new DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasPayloadLinks
// Deprecated: This method is obsolete. Use PostAsHasPayloadLinksPostResponse instead.
// returns a DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) Post(ctx context.Context, body DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksPostRequestBodyable, requestConfiguration *DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDevicecompliancepoliciesHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksResponseable), nil
}
// PostAsHasPayloadLinksPostResponse invoke action hasPayloadLinks
// returns a DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) PostAsHasPayloadLinksPostResponse(ctx context.Context, body DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksPostRequestBodyable, requestConfiguration *DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDevicecompliancepoliciesHaspayloadlinksHasPayloadLinksPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksPostResponseable), nil
}
// ToPostRequestInformation invoke action hasPayloadLinks
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksPostRequestBodyable, requestConfiguration *DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder when successful
func (m *DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) {
    return NewDevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
