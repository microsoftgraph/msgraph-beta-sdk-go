package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicehealthscriptsDeviceHealthScriptsRequestBuilder provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
type DevicehealthscriptsDeviceHealthScriptsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicehealthscriptsDeviceHealthScriptsRequestBuilderGetQueryParameters the list of device health scripts associated with the tenant.
type DevicehealthscriptsDeviceHealthScriptsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DevicehealthscriptsDeviceHealthScriptsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicehealthscriptsDeviceHealthScriptsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicehealthscriptsDeviceHealthScriptsRequestBuilderGetQueryParameters
}
// DevicehealthscriptsDeviceHealthScriptsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicehealthscriptsDeviceHealthScriptsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AreGlobalScriptsAvailable provides operations to call the areGlobalScriptsAvailable method.
// returns a *DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptsRequestBuilder) AreGlobalScriptsAvailable()(*DevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilder) {
    return NewDevicehealthscriptsAreglobalscriptsavailableAreGlobalScriptsAvailableRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByDeviceHealthScriptId provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptsRequestBuilder) ByDeviceHealthScriptId(deviceHealthScriptId string)(*DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceHealthScriptId != "" {
        urlTplParams["deviceHealthScript%2Did"] = deviceHealthScriptId
    }
    return NewDevicehealthscriptsDeviceHealthScriptItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicehealthscriptsDeviceHealthScriptsRequestBuilderInternal instantiates a new DevicehealthscriptsDeviceHealthScriptsRequestBuilder and sets the default values.
func NewDevicehealthscriptsDeviceHealthScriptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsDeviceHealthScriptsRequestBuilder) {
    m := &DevicehealthscriptsDeviceHealthScriptsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicehealthscriptsDeviceHealthScriptsRequestBuilder instantiates a new DevicehealthscriptsDeviceHealthScriptsRequestBuilder and sets the default values.
func NewDevicehealthscriptsDeviceHealthScriptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsDeviceHealthScriptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicehealthscriptsDeviceHealthScriptsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicehealthscriptsCountRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptsRequestBuilder) Count()(*DevicehealthscriptsCountRequestBuilder) {
    return NewDevicehealthscriptsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableGlobalScripts provides operations to call the enableGlobalScripts method.
// returns a *DevicehealthscriptsEnableglobalscriptsEnableGlobalScriptsRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptsRequestBuilder) EnableGlobalScripts()(*DevicehealthscriptsEnableglobalscriptsEnableGlobalScriptsRequestBuilder) {
    return NewDevicehealthscriptsEnableglobalscriptsEnableGlobalScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of device health scripts associated with the tenant.
// returns a DeviceHealthScriptCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsDeviceHealthScriptsRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicehealthscriptsDeviceHealthScriptsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptCollectionResponseable), nil
}
// GetRemediationSummary provides operations to call the getRemediationSummary method.
// returns a *DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptsRequestBuilder) GetRemediationSummary()(*DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder) {
    return NewDevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to deviceHealthScripts for deviceManagement
// returns a DeviceHealthScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsDeviceHealthScriptsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DevicehealthscriptsDeviceHealthScriptsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable), nil
}
// ToGetRequestInformation the list of device health scripts associated with the tenant.
// returns a *RequestInformation when successful
func (m *DevicehealthscriptsDeviceHealthScriptsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicehealthscriptsDeviceHealthScriptsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceHealthScripts for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicehealthscriptsDeviceHealthScriptsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DevicehealthscriptsDeviceHealthScriptsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicehealthscriptsDeviceHealthScriptsRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptsRequestBuilder) WithUrl(rawUrl string)(*DevicehealthscriptsDeviceHealthScriptsRequestBuilder) {
    return NewDevicehealthscriptsDeviceHealthScriptsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
