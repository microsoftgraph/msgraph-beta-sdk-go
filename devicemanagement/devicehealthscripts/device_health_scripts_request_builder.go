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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceHealthScriptsRequestBuilderGetOptions options for Get
type DeviceHealthScriptsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DeviceHealthScriptsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceHealthScriptsRequestBuilderGetQueryParameters the list of device health scripts associated with the tenant.
type DeviceHealthScriptsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// DeviceHealthScriptsRequestBuilderPostOptions options for Post
type DeviceHealthScriptsRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AreGlobalScriptsAvailable provides operations to call the areGlobalScriptsAvailable method.
func (m *DeviceHealthScriptsRequestBuilder) AreGlobalScriptsAvailable()(*i6a873728e2a8769c250276e0c4dc8fa971581156a07a4c225120f5c5242df0a3.AreGlobalScriptsAvailableRequestBuilder) {
    return i6a873728e2a8769c250276e0c4dc8fa971581156a07a4c225120f5c5242df0a3.NewAreGlobalScriptsAvailableRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceHealthScriptsRequestBuilderInternal instantiates a new DeviceHealthScriptsRequestBuilder and sets the default values.
func NewDeviceHealthScriptsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsRequestBuilder) {
    m := &DeviceHealthScriptsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts{?top,skip,search,filter,count,orderby,select,expand}";
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
// CreateGetRequestInformation the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsRequestBuilder) CreateGetRequestInformation(options *DeviceHealthScriptsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptsRequestBuilder) CreatePostRequestInformation(options *DeviceHealthScriptsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// EnableGlobalScripts the enableGlobalScripts property
func (m *DeviceHealthScriptsRequestBuilder) EnableGlobalScripts()(*i87802417fac8dafea49fd0889f538e357ba4f700228dd931789b2c3a9194e714.EnableGlobalScriptsRequestBuilder) {
    return i87802417fac8dafea49fd0889f538e357ba4f700228dd931789b2c3a9194e714.NewEnableGlobalScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsRequestBuilder) Get(options *DeviceHealthScriptsRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptCollectionResponseable), nil
}
// GetRemediationSummary provides operations to call the getRemediationSummary method.
func (m *DeviceHealthScriptsRequestBuilder) GetRemediationSummary()(*i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547.GetRemediationSummaryRequestBuilder) {
    return i2e7b8f78fa151b32792a059a6a4155ce021c6574bebe2d5c93dfeeaa4fb76547.NewGetRemediationSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptsRequestBuilder) Post(options *DeviceHealthScriptsRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable), nil
}
