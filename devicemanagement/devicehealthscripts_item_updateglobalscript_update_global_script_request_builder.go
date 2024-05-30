package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder provides operations to call the updateGlobalScript method.
type DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilderInternal instantiates a new DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder and sets the default values.
func NewDevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder) {
    m := &DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}/updateGlobalScript", pathParameters),
    }
    return m
}
// NewDevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder instantiates a new DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder and sets the default values.
func NewDevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the Proprietary Device Health Script
// Deprecated: This method is obsolete. Use PostAsUpdateGlobalScriptPostResponse instead.
// returns a DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder) Post(ctx context.Context, body DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptPostRequestBodyable, requestConfiguration *DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilderPostRequestConfiguration)(DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptResponseable), nil
}
// PostAsUpdateGlobalScriptPostResponse update the Proprietary Device Health Script
// returns a DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder) PostAsUpdateGlobalScriptPostResponse(ctx context.Context, body DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptPostRequestBodyable, requestConfiguration *DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilderPostRequestConfiguration)(DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptPostResponseable), nil
}
// ToPostRequestInformation update the Proprietary Device Health Script
// returns a *RequestInformation when successful
func (m *DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder) ToPostRequestInformation(ctx context.Context, body DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptPostRequestBodyable, requestConfiguration *DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder when successful
func (m *DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder) WithUrl(rawUrl string)(*DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder) {
    return NewDevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
