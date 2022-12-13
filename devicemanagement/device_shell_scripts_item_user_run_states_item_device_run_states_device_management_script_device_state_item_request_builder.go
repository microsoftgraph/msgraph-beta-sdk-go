package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder provides operations to manage the deviceRunStates property of the microsoft.graph.deviceManagementScriptUserState entity.
type DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderGetQueryParameters list of run states for this script across all devices of specific user.
type DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderGetQueryParameters
}
// DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal instantiates a new DeviceManagementScriptDeviceStateItemRequestBuilder and sets the default values.
func NewDeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    m := &DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceShellScripts/{deviceShellScript%2Did}/userRunStates/{deviceManagementScriptUserState%2Did}/deviceRunStates/{deviceManagementScriptDeviceState%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder instantiates a new DeviceManagementScriptDeviceStateItemRequestBuilder and sets the default values.
func NewDeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceRunStates for deviceManagement
func (m *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation list of run states for this script across all devices of specific user.
func (m *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property deviceRunStates in deviceManagement
func (m *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, requestConfiguration *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property deviceRunStates for deviceManagement
func (m *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of run states for this script across all devices of specific user.
func (m *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable), nil
}
// ManagedDevice provides operations to manage the managedDevice property of the microsoft.graph.deviceManagementScriptDeviceState entity.
func (m *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) ManagedDevice()(*DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder) {
    return NewDeviceShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property deviceRunStates in deviceManagement
func (m *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, requestConfiguration *DeviceShellScriptsItemUserRunStatesItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable), nil
}
