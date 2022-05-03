package devicehealthscripts

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/getremediationsummary"
    i6a873728e2a8769c250276e0c4dc8fa971581156a07a4c225120f5c5242df0a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/areglobalscriptsavailable"
    i87802417fac8dafea49fd0889f538e357ba4f700228dd931789b2c3a9194e714 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/enableglobalscripts"
    ic4de1ae6615c5df9eec3108346be9478c79d3a7aa8645c40d6f5091d453580f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/count"
)

// DeviceHealthScriptsRequestBuilder provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
type DeviceHealthScriptsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceHealthScriptsRequestBuilderGetQueryParameters the list of device health scripts associated with the tenant.
type DeviceHealthScriptsRequestBuilderGetQueryParameters struct {
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
// DeviceHealthScriptsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceHealthScriptsRequestBuilderGetQueryParameters
}
// DeviceHealthScriptsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AreGlobalScriptsAvailable provides operations to call the areGlobalScriptsAvailable method.
func (m *DeviceHealthScriptsRequestBuilder) AreGlobalScriptsAvailable()(*i6a873728e2a8769c250276e0c4dc8fa971581156a07a4c225120f5c5242df0a3.AreGlobalScriptsAvailableRequestBuilder) {
    return i6a873728e2a8769c250276e0c4dc8fa971581156a07a4c225120f5c5242df0a3.NewAreGlobalScriptsAvailableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceHealthScriptsRequestBuilderInternal instantiates a new DeviceHealthScriptsRequestBuilder and sets the default values.
func NewDeviceHealthScriptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsRequestBuilder) {
    m := &DeviceHealthScriptsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceHealthScriptsRequestBuilder instantiates a new DeviceHealthScriptsRequestBuilder and sets the default values.
func NewDeviceHealthScriptsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *DeviceHealthScriptsRequestBuilder) Count()(*ic4de1ae6615c5df9eec3108346be9478c79d3a7aa8645c40d6f5091d453580f1.CountRequestBuilder) {
    return ic4de1ae6615c5df9eec3108346be9478c79d3a7aa8645c40d6f5091d453580f1.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformationWithRequestConfiguration the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceHealthScriptsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DeviceHealthScriptsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// EnableGlobalScripts the enableGlobalScripts property
func (m *DeviceHealthScriptsRequestBuilder) EnableGlobalScripts()(*i87802417fac8dafea49fd0889f538e357ba4f700228dd931789b2c3a9194e714.EnableGlobalScriptsRequestBuilder) {
    return i87802417fac8dafea49fd0889f538e357ba4f700228dd931789b2c3a9194e714.NewEnableGlobalScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRemediationSummary provides operations to call the getRemediationSummary method.
func (m *DeviceHealthScriptsRequestBuilder) GetRemediationSummary()(*i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547.GetRemediationSummaryRequestBuilder) {
    return i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547.NewGetRemediationSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithResponseHandler the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsRequestBuilder) GetWithResponseHandler(requestConfiguration *DeviceHealthScriptsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptCollectionResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsRequestBuilder) GetWithResponseHandler(requestConfiguration *DeviceHealthScriptsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptCollectionResponseable), nil
}
// PostWithResponseHandler create new navigation property to deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptsRequestBuilder) PostWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DeviceHealthScriptsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler create new navigation property to deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptsRequestBuilder) PostWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DeviceHealthScriptsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable), nil
}
