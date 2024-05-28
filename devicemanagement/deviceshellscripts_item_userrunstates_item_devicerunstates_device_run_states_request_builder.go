package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder provides operations to manage the deviceRunStates property of the microsoft.graph.deviceManagementScriptUserState entity.
type DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderGetQueryParameters list of run states for this script across all devices of specific user.
type DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderGetQueryParameters struct {
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
// DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderGetQueryParameters
}
// DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementScriptDeviceStateId provides operations to manage the deviceRunStates property of the microsoft.graph.deviceManagementScriptUserState entity.
// returns a *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder when successful
func (m *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) ByDeviceManagementScriptDeviceStateId(deviceManagementScriptDeviceStateId string)(*DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementScriptDeviceStateId != "" {
        urlTplParams["deviceManagementScriptDeviceState%2Did"] = deviceManagementScriptDeviceStateId
    }
    return NewDeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderInternal instantiates a new DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder and sets the default values.
func NewDeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    m := &DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceShellScripts/{deviceShellScript%2Did}/userRunStates/{deviceManagementScriptUserState%2Did}/deviceRunStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder instantiates a new DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder and sets the default values.
func NewDeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder when successful
func (m *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) Count()(*DeviceshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder) {
    return NewDeviceshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of run states for this script across all devices of specific user.
// returns a DeviceManagementScriptDeviceStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptDeviceStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateCollectionResponseable), nil
}
// Post create new navigation property to deviceRunStates for deviceManagement
// returns a DeviceManagementScriptDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, requestConfiguration *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation list of run states for this script across all devices of specific user.
// returns a *RequestInformation when successful
func (m *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceRunStates for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, requestConfiguration *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder when successful
func (m *DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) WithUrl(rawUrl string)(*DeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    return NewDeviceshellscriptsItemUserrunstatesItemDevicerunstatesDeviceRunStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
