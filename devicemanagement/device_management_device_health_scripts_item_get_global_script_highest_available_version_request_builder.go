package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder provides operations to call the getGlobalScriptHighestAvailableVersion method.
type DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderInternal instantiates a new GetGlobalScriptHighestAvailableVersionRequestBuilder and sets the default values.
func NewDeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    m := &DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}/microsoft.graph.getGlobalScriptHighestAvailableVersion";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder instantiates a new GetGlobalScriptHighestAvailableVersionRequestBuilder and sets the default values.
func NewDeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation update the Proprietary Device Health Script
func (m *DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post update the Proprietary Device Health Script
func (m *DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) Post(ctx context.Context, requestConfiguration *DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionResponseable), nil
}
