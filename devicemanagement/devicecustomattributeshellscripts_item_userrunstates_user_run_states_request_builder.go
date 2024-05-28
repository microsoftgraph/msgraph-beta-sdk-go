package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder provides operations to manage the userRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
type DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderGetQueryParameters list of run states for this script across all users.
type DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderGetQueryParameters struct {
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
// DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderGetQueryParameters
}
// DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementScriptUserStateId provides operations to manage the userRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
// returns a *DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) ByDeviceManagementScriptUserStateId(deviceManagementScriptUserStateId string)(*DevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementScriptUserStateId != "" {
        urlTplParams["deviceManagementScriptUserState%2Did"] = deviceManagementScriptUserStateId
    }
    return NewDevicecustomattributeshellscriptsItemUserrunstatesDeviceManagementScriptUserStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderInternal instantiates a new DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) {
    m := &DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}/userRunStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder instantiates a new DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicecustomattributeshellscriptsItemUserrunstatesCountRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) Count()(*DevicecustomattributeshellscriptsItemUserrunstatesCountRequestBuilder) {
    return NewDevicecustomattributeshellscriptsItemUserrunstatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of run states for this script across all users.
// returns a DeviceManagementScriptUserStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptUserStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateCollectionResponseable), nil
}
// Post create new navigation property to userRunStates for deviceManagement
// returns a DeviceManagementScriptUserStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptUserStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable), nil
}
// ToGetRequestInformation list of run states for this script across all users.
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userRunStates for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptUserStateable, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) WithUrl(rawUrl string)(*DevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) {
    return NewDevicecustomattributeshellscriptsItemUserrunstatesUserRunStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
