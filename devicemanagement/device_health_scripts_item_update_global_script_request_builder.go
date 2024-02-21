package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder provides operations to call the updateGlobalScript method.
type DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderInternal instantiates a new DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder and sets the default values.
func NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) {
    m := &DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}/updateGlobalScript", pathParameters),
    }
    return m
}
// NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder instantiates a new DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder and sets the default values.
func NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the Proprietary Device Health Script
// Deprecated: This method is obsolete. Use PostAsUpdateGlobalScriptPostResponse instead.
// returns a DeviceHealthScriptsItemUpdateGlobalScriptResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) Post(ctx context.Context, body DeviceHealthScriptsItemUpdateGlobalScriptPostRequestBodyable, requestConfiguration *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderPostRequestConfiguration)(DeviceHealthScriptsItemUpdateGlobalScriptResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceHealthScriptsItemUpdateGlobalScriptResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceHealthScriptsItemUpdateGlobalScriptResponseable), nil
}
// PostAsUpdateGlobalScriptPostResponse update the Proprietary Device Health Script
// returns a DeviceHealthScriptsItemUpdateGlobalScriptPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) PostAsUpdateGlobalScriptPostResponse(ctx context.Context, body DeviceHealthScriptsItemUpdateGlobalScriptPostRequestBodyable, requestConfiguration *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderPostRequestConfiguration)(DeviceHealthScriptsItemUpdateGlobalScriptPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceHealthScriptsItemUpdateGlobalScriptPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceHealthScriptsItemUpdateGlobalScriptPostResponseable), nil
}
// ToPostRequestInformation update the Proprietary Device Health Script
// returns a *RequestInformation when successful
func (m *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceHealthScriptsItemUpdateGlobalScriptPostRequestBodyable, requestConfiguration *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder when successful
func (m *DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) WithUrl(rawUrl string)(*DeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) {
    return NewDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
