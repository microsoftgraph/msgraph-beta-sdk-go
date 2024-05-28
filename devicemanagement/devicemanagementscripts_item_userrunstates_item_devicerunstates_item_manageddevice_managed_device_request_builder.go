package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder provides operations to manage the managedDevice property of the microsoft.graph.deviceManagementScriptDeviceState entity.
type DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderGetQueryParameters the managed devices that executes the device management script.
type DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderGetQueryParameters
}
// NewDevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderInternal instantiates a new DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder and sets the default values.
func NewDevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder) {
    m := &DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript%2Did}/userRunStates/{deviceManagementScriptUserState%2Did}/deviceRunStates/{deviceManagementScriptDeviceState%2Did}/managedDevice{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder instantiates a new DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder and sets the default values.
func NewDevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the managed devices that executes the device management script.
// returns a ManagedDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
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
func (m *DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder when successful
func (m *DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder) WithUrl(rawUrl string)(*DevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder) {
    return NewDevicemanagementscriptsItemUserrunstatesItemDevicerunstatesItemManageddeviceManagedDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
