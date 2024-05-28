package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder provides operations to call the getRemediationSummary method.
type DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilderInternal instantiates a new DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder and sets the default values.
func NewDevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder) {
    m := &DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts/getRemediationSummary()", pathParameters),
    }
    return m
}
// NewDevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder instantiates a new DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder and sets the default values.
func NewDevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getRemediationSummary
// returns a DeviceHealthScriptRemediationSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptRemediationSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptRemediationSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptRemediationSummaryable), nil
}
// ToGetRequestInformation invoke function getRemediationSummary
// returns a *RequestInformation when successful
func (m *DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder when successful
func (m *DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder) WithUrl(rawUrl string)(*DevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder) {
    return NewDevicehealthscriptsGetremediationsummaryGetRemediationSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
