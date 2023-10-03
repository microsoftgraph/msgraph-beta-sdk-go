package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementScriptsHasPayloadLinksRequestBuilder provides operations to call the hasPayloadLinks method.
type DeviceManagementScriptsHasPayloadLinksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceManagementScriptsHasPayloadLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementScriptsHasPayloadLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementScriptsHasPayloadLinksRequestBuilderInternal instantiates a new HasPayloadLinksRequestBuilder and sets the default values.
func NewDeviceManagementScriptsHasPayloadLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementScriptsHasPayloadLinksRequestBuilder) {
    m := &DeviceManagementScriptsHasPayloadLinksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceManagementScripts/hasPayloadLinks", pathParameters),
    }
    return m
}
// NewDeviceManagementScriptsHasPayloadLinksRequestBuilder instantiates a new HasPayloadLinksRequestBuilder and sets the default values.
func NewDeviceManagementScriptsHasPayloadLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementScriptsHasPayloadLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementScriptsHasPayloadLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasPayloadLinks
// Deprecated: This method is obsolete. Use PostAsHasPayloadLinksPostResponse instead.
func (m *DeviceManagementScriptsHasPayloadLinksRequestBuilder) Post(ctx context.Context, body DeviceManagementScriptsHasPayloadLinksPostRequestBodyable, requestConfiguration *DeviceManagementScriptsHasPayloadLinksRequestBuilderPostRequestConfiguration)(DeviceManagementScriptsHasPayloadLinksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceManagementScriptsHasPayloadLinksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceManagementScriptsHasPayloadLinksResponseable), nil
}
// PostAsHasPayloadLinksPostResponse invoke action hasPayloadLinks
func (m *DeviceManagementScriptsHasPayloadLinksRequestBuilder) PostAsHasPayloadLinksPostResponse(ctx context.Context, body DeviceManagementScriptsHasPayloadLinksPostRequestBodyable, requestConfiguration *DeviceManagementScriptsHasPayloadLinksRequestBuilderPostRequestConfiguration)(DeviceManagementScriptsHasPayloadLinksPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceManagementScriptsHasPayloadLinksPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceManagementScriptsHasPayloadLinksPostResponseable), nil
}
// ToPostRequestInformation invoke action hasPayloadLinks
func (m *DeviceManagementScriptsHasPayloadLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceManagementScriptsHasPayloadLinksPostRequestBodyable, requestConfiguration *DeviceManagementScriptsHasPayloadLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeviceManagementScriptsHasPayloadLinksRequestBuilder) WithUrl(rawUrl string)(*DeviceManagementScriptsHasPayloadLinksRequestBuilder) {
    return NewDeviceManagementScriptsHasPayloadLinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
