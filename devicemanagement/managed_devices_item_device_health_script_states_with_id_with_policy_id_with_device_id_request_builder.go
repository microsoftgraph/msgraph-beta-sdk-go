package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
type ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters results of device health scripts that ran for this device. Default is empty list. This property is read-only.
type ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters
}
// ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal instantiates a new ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder and sets the default values.
func NewManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceId *string, id *string, policyId *string)(*ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    m := &ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceHealthScriptStates/id='{id}',policyId='{policyId}',deviceId='{deviceId}'{?%24expand,%24select}", pathParameters),
    }
    if deviceId != nil {
        m.BaseRequestBuilder.PathParameters["deviceId"] = *deviceId
    }
    if id != nil {
        m.BaseRequestBuilder.PathParameters["id"] = *id
    }
    if policyId != nil {
        m.BaseRequestBuilder.PathParameters["policyId"] = *policyId
    }
    return m
}
// NewManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder instantiates a new ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder and sets the default values.
func NewManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Delete delete navigation property deviceHealthScriptStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get results of device health scripts that ran for this device. Default is empty list. This property is read-only.
// returns a DeviceHealthScriptPolicyStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property deviceHealthScriptStates in deviceManagement
// returns a DeviceHealthScriptPolicyStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property deviceHealthScriptStates for deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceHealthScriptStates/id='{id}',policyId='{policyId}',deviceId='{deviceId}'", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation results of device health scripts that ran for this device. Default is empty list. This property is read-only.
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceHealthScriptStates in deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, requestConfiguration *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/deviceHealthScriptStates/id='{id}',policyId='{policyId}',deviceId='{deviceId}'", m.BaseRequestBuilder.PathParameters)
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
// returns a *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder when successful
func (m *ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    return NewManagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
