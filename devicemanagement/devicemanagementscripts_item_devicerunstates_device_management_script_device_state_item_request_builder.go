package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder provides operations to manage the deviceRunStates property of the microsoft.graph.deviceManagementScript entity.
type DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderGetQueryParameters list of run states for this script across all devices.
type DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderGetQueryParameters
}
// DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal instantiates a new DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder and sets the default values.
func NewDevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    m := &DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript%2Did}/deviceRunStates/{deviceManagementScriptDeviceState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder instantiates a new DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder and sets the default values.
func NewDevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceRunStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of run states for this script across all devices.
// returns a DeviceManagementScriptDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable), nil
}
// ManagedDevice provides operations to manage the managedDevice property of the microsoft.graph.deviceManagementScriptDeviceState entity.
// returns a *DevicemanagementscriptsItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder when successful
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) ManagedDevice()(*DevicemanagementscriptsItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder) {
    return NewDevicemanagementscriptsItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deviceRunStates in deviceManagement
// returns a DeviceManagementScriptDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, requestConfiguration *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable), nil
}
// ToDeleteRequestInformation delete navigation property deviceRunStates for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of run states for this script across all devices.
// returns a *RequestInformation when successful
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceRunStates in deviceManagement
// returns a *RequestInformation when successful
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, requestConfiguration *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder when successful
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) WithUrl(rawUrl string)(*DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    return NewDevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
