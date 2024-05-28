package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder provides operations to call the getGlobalScriptHighestAvailableVersion method.
type DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilderInternal instantiates a new DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder and sets the default values.
func NewDevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    m := &DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}/getGlobalScriptHighestAvailableVersion", pathParameters),
    }
    return m
}
// NewDevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder instantiates a new DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder and sets the default values.
func NewDevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the Proprietary Device Health Script
// Deprecated: This method is obsolete. Use PostAsGetGlobalScriptHighestAvailableVersionPostResponse instead.
// returns a DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder) Post(ctx context.Context, requestConfiguration *DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration)(DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionResponseable), nil
}
// PostAsGetGlobalScriptHighestAvailableVersionPostResponse update the Proprietary Device Health Script
// returns a DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder) PostAsGetGlobalScriptHighestAvailableVersionPostResponse(ctx context.Context, requestConfiguration *DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration)(DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionPostResponseable), nil
}
// ToPostRequestInformation update the Proprietary Device Health Script
// returns a *RequestInformation when successful
func (m *DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder when successful
func (m *DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder) WithUrl(rawUrl string)(*DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    return NewDevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
