package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder provides operations to call the areGlobalScriptsAvailable method.
type DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilderInternal instantiates a new DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder and sets the default values.
func NewDevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder) {
    m := &DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts/areGlobalScriptsAvailable()", pathParameters),
    }
    return m
}
// NewDevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder instantiates a new DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder and sets the default values.
func NewDevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function areGlobalScriptsAvailable
// returns a *GlobalDeviceHealthScriptState when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration)(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendEnum(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseGlobalDeviceHealthScriptState, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState), nil
}
// ToGetRequestInformation invoke function areGlobalScriptsAvailable
// returns a *RequestInformation when successful
func (m *DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder when successful
func (m *DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder) WithUrl(rawUrl string)(*DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder) {
    return NewDevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
