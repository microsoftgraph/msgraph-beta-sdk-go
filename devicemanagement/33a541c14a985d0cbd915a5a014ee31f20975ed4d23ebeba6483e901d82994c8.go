package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder provides operations to manage the managedDevice property of the microsoft.graph.deviceManagementScriptDeviceState entity.
type DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderGetQueryParameters the managed devices that executes the device management script.
type DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderGetQueryParameters
}
// NewDeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderInternal instantiates a new DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder and sets the default values.
func NewDeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder) {
    m := &DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}/userRunStates/{deviceManagementScriptUserState%2Did}/deviceRunStates/{deviceManagementScriptDeviceState%2Did}/managedDevice{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder instantiates a new DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder and sets the default values.
func NewDeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the managed devices that executes the device management script.
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
// ToGetRequestInformation the managed devices that executes the device management script.
// returns a *RequestInformation when successful
func (m *DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder when successful
func (m *DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder) WithUrl(rawUrl string)(*DeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemUserRunStatesItemDeviceRunStatesItemManagedDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
