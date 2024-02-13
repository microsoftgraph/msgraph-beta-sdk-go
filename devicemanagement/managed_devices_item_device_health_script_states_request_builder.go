package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
type ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetQueryParameters results of device health scripts that ran for this device. Default is empty list. This property is read-only.
type ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetQueryParameters struct {
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
// ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetQueryParameters
}
// ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal instantiates a new ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder and sets the default values.
func NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    m := &ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceHealthScriptStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilder instantiates a new ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder and sets the default values.
func NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedDevicesItemDeviceHealthScriptStatesCountRequestBuilder when successful
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) Count()(*ManagedDevicesItemDeviceHealthScriptStatesCountRequestBuilder) {
    return NewManagedDevicesItemDeviceHealthScriptStatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get results of device health scripts that ran for this device. Default is empty list. This property is read-only.
// returns a DeviceHealthScriptPolicyStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptPolicyStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateCollectionResponseable), nil
}
// Post create new navigation property to deviceHealthScriptStates for deviceManagement
// returns a DeviceHealthScriptPolicyStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptPolicyStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable), nil
}
// ToGetRequestInformation results of device health scripts that ran for this device. Default is empty list. This property is read-only.
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceHealthScriptStates for deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceHealthScriptStates", m.BaseRequestBuilder.PathParameters)
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
// WithIdWithPolicyIdWithDeviceId provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
// returns a *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder when successful
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) WithIdWithPolicyIdWithDeviceId(deviceId *string, id *string, policyId *string)(*ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    return NewManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, deviceId, id, policyId)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder when successful
func (m *ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemDeviceHealthScriptStatesRequestBuilder) {
    return NewManagedDevicesItemDeviceHealthScriptStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
