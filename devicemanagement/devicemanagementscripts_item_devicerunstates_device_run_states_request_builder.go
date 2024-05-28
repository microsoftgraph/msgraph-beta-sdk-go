package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder provides operations to manage the deviceRunStates property of the microsoft.graph.deviceManagementScript entity.
type DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderGetQueryParameters list of run states for this script across all devices.
type DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderGetQueryParameters struct {
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
// DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderGetQueryParameters
}
// DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementScriptDeviceStateId provides operations to manage the deviceRunStates property of the microsoft.graph.deviceManagementScript entity.
// returns a *DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder when successful
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) ByDeviceManagementScriptDeviceStateId(deviceManagementScriptDeviceStateId string)(*DevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementScriptDeviceStateId != "" {
        urlTplParams["deviceManagementScriptDeviceState%2Did"] = deviceManagementScriptDeviceStateId
    }
    return NewDevicemanagementscriptsItemDevicerunstatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderInternal instantiates a new DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder and sets the default values.
func NewDevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    m := &DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript%2Did}/deviceRunStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder instantiates a new DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder and sets the default values.
func NewDevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicemanagementscriptsItemDevicerunstatesCountRequestBuilder when successful
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) Count()(*DevicemanagementscriptsItemDevicerunstatesCountRequestBuilder) {
    return NewDevicemanagementscriptsItemDevicerunstatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of run states for this script across all devices.
// returns a DeviceManagementScriptDeviceStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateCollectionResponseable, error) {
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
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, requestConfiguration *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, error) {
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
// ToGetRequestInformation list of run states for this script across all devices.
// returns a *RequestInformation when successful
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptDeviceStateable, requestConfiguration *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder when successful
func (m *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) WithUrl(rawUrl string)(*DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    return NewDevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
