package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
type ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters results of device health scripts that ran for this device. Default is empty list. This property is read-only.
type ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters
}
// ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal instantiates a new ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder and sets the default values.
func NewComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceId *string, id *string, policyId *string)(*ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    m := &ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/deviceHealthScriptStates/id='{id}',policyId='{policyId}',deviceId='{deviceId}'{?%24expand,%24select}", pathParameters),
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
// NewComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder instantiates a new ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder and sets the default values.
func NewComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Delete delete navigation property deviceHealthScriptStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, error) {
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
func (m *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, requestConfiguration *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, error) {
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
func (m *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation results of device health scripts that ran for this device. Default is empty list. This property is read-only.
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, requestConfiguration *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder when successful
func (m *ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    return NewComanageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
