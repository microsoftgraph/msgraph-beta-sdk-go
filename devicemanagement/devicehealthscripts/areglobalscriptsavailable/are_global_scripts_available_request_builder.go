package areglobalscriptsavailable

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AreGlobalScriptsAvailableRequestBuilder provides operations to call the areGlobalScriptsAvailable method.
type AreGlobalScriptsAvailableRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAreGlobalScriptsAvailableRequestBuilderInternal instantiates a new AreGlobalScriptsAvailableRequestBuilder and sets the default values.
func NewAreGlobalScriptsAvailableRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AreGlobalScriptsAvailableRequestBuilder) {
    m := &AreGlobalScriptsAvailableRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/microsoft.graph.areGlobalScriptsAvailable()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAreGlobalScriptsAvailableRequestBuilder instantiates a new AreGlobalScriptsAvailableRequestBuilder and sets the default values.
func NewAreGlobalScriptsAvailableRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AreGlobalScriptsAvailableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAreGlobalScriptsAvailableRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function areGlobalScriptsAvailable
func (m *AreGlobalScriptsAvailableRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function areGlobalScriptsAvailable
func (m *AreGlobalScriptsAvailableRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function areGlobalScriptsAvailable
func (m *AreGlobalScriptsAvailableRequestBuilder) Get()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function areGlobalScriptsAvailable
func (m *AreGlobalScriptsAvailableRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *AreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendEnumAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseGlobalDeviceHealthScriptState, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState), nil
}
