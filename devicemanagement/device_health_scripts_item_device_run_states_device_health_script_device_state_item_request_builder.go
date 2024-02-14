package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder provides operations to manage the deviceRunStates property of the microsoft.graph.deviceHealthScript entity.
type DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderGetQueryParameters list of run states for the device health script across all devices
type DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderGetQueryParameters
}
// DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderInternal instantiates a new DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder and sets the default values.
func NewDeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) {
    m := &DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}/deviceRunStates/{deviceHealthScriptDeviceState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder instantiates a new DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder and sets the default values.
func NewDeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceRunStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of run states for the device health script across all devices
// returns a DeviceHealthScriptDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptDeviceStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptDeviceStateable), nil
}
// ManagedDevice provides operations to manage the managedDevice property of the microsoft.graph.deviceHealthScriptDeviceState entity.
// returns a *DeviceHealthScriptsItemDeviceRunStatesItemManagedDeviceRequestBuilder when successful
func (m *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) ManagedDevice()(*DeviceHealthScriptsItemDeviceRunStatesItemManagedDeviceRequestBuilder) {
    return NewDeviceHealthScriptsItemDeviceRunStatesItemManagedDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deviceRunStates in deviceManagement
// returns a DeviceHealthScriptDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptDeviceStateable, requestConfiguration *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptDeviceStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptDeviceStateable), nil
}
// ToDeleteRequestInformation delete navigation property deviceRunStates for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}/deviceRunStates/{deviceHealthScriptDeviceState%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of run states for the device health script across all devices
// returns a *RequestInformation when successful
func (m *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptDeviceStateable, requestConfiguration *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}/deviceRunStates/{deviceHealthScriptDeviceState%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder when successful
func (m *DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) WithUrl(rawUrl string)(*DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) {
    return NewDeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
