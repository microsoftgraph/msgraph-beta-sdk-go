package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder provides operations to manage the userRunStates property of the microsoft.graph.deviceManagementScript entity.
type DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderGetQueryParameters list of run states for this script across all users.
type DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderGetQueryParameters
}
// DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderInternal instantiates a new DeviceManagementScriptUserStateItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) {
    m := &DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript%2Did}/userRunStates/{deviceManagementScriptUserState%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder instantiates a new DeviceManagementScriptUserStateItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property userRunStates for deviceManagement
func (m *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation list of run states for this script across all users.
func (m *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property userRunStates in deviceManagement
func (m *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, requestConfiguration *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property userRunStates for deviceManagement
func (m *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceManagementScriptUserState entity.
func (m *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) DeviceRunStates()(*DeviceManagementDeviceManagementScriptsItemUserRunStatesItemDeviceRunStatesRequestBuilder) {
    return NewDeviceManagementDeviceManagementScriptsItemUserRunStatesItemDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById provides operations to manage the deviceRunStates property of the microsoft.graph.deviceManagementScriptUserState entity.
func (m *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) DeviceRunStatesById(id string)(*DeviceManagementDeviceManagementScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState%2Did"] = id
    }
    return NewDeviceManagementDeviceManagementScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get list of run states for this script across all users.
func (m *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptUserStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable), nil
}
// Patch update the navigation property userRunStates in deviceManagement
func (m *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, requestConfiguration *DeviceManagementDeviceManagementScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptUserStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable), nil
}
