// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceCompliancePoliciesHasPayloadLinksRequestBuilder provides operations to call the hasPayloadLinks method.
type DeviceCompliancePoliciesHasPayloadLinksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceCompliancePoliciesHasPayloadLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesHasPayloadLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceCompliancePoliciesHasPayloadLinksRequestBuilderInternal instantiates a new DeviceCompliancePoliciesHasPayloadLinksRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesHasPayloadLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesHasPayloadLinksRequestBuilder) {
    m := &DeviceCompliancePoliciesHasPayloadLinksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/hasPayloadLinks", pathParameters),
    }
    return m
}
// NewDeviceCompliancePoliciesHasPayloadLinksRequestBuilder instantiates a new DeviceCompliancePoliciesHasPayloadLinksRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesHasPayloadLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesHasPayloadLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesHasPayloadLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasPayloadLinks
// Deprecated: This method is obsolete. Use PostAsHasPayloadLinksPostResponse instead.
// returns a DeviceCompliancePoliciesHasPayloadLinksResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceCompliancePoliciesHasPayloadLinksRequestBuilder) Post(ctx context.Context, body DeviceCompliancePoliciesHasPayloadLinksPostRequestBodyable, requestConfiguration *DeviceCompliancePoliciesHasPayloadLinksRequestBuilderPostRequestConfiguration)(DeviceCompliancePoliciesHasPayloadLinksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceCompliancePoliciesHasPayloadLinksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceCompliancePoliciesHasPayloadLinksResponseable), nil
}
// PostAsHasPayloadLinksPostResponse invoke action hasPayloadLinks
// returns a DeviceCompliancePoliciesHasPayloadLinksPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceCompliancePoliciesHasPayloadLinksRequestBuilder) PostAsHasPayloadLinksPostResponse(ctx context.Context, body DeviceCompliancePoliciesHasPayloadLinksPostRequestBodyable, requestConfiguration *DeviceCompliancePoliciesHasPayloadLinksRequestBuilderPostRequestConfiguration)(DeviceCompliancePoliciesHasPayloadLinksPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceCompliancePoliciesHasPayloadLinksPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceCompliancePoliciesHasPayloadLinksPostResponseable), nil
}
// ToPostRequestInformation invoke action hasPayloadLinks
// returns a *RequestInformation when successful
func (m *DeviceCompliancePoliciesHasPayloadLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceCompliancePoliciesHasPayloadLinksPostRequestBodyable, requestConfiguration *DeviceCompliancePoliciesHasPayloadLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceCompliancePoliciesHasPayloadLinksRequestBuilder when successful
func (m *DeviceCompliancePoliciesHasPayloadLinksRequestBuilder) WithUrl(rawUrl string)(*DeviceCompliancePoliciesHasPayloadLinksRequestBuilder) {
    return NewDeviceCompliancePoliciesHasPayloadLinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
