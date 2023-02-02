package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilder provides operations to call the areGlobalScriptsAvailable method.
type DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilderInternal instantiates a new AreGlobalScriptsAvailableRequestBuilder and sets the default values.
func NewDeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilder) {
    m := &DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/microsoft.graph.areGlobalScriptsAvailable()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilder instantiates a new AreGlobalScriptsAvailableRequestBuilder and sets the default values.
func NewDeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function areGlobalScriptsAvailable
func (m *DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration)(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendEnum(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseGlobalDeviceHealthScriptState, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GlobalDeviceHealthScriptState), nil
}
// ToGetRequestInformation invoke function areGlobalScriptsAvailable
func (m *DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceHealthScriptsMicrosoftGraphAreGlobalScriptsAvailableAreGlobalScriptsAvailableRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
