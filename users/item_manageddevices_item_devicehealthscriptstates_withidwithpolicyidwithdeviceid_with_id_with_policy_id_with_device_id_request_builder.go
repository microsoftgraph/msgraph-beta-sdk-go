package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
type ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters results of device health scripts that ran for this device. Default is empty list. This property is read-only.
type ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters
}
// ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal instantiates a new ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder and sets the default values.
func NewItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceId *string, id *string, policyId *string)(*ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    m := &ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/deviceHealthScriptStates/id='{id}',policyId='{policyId}',deviceId='{deviceId}'{?%24expand,%24select}", pathParameters),
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
// NewItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder instantiates a new ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder and sets the default values.
func NewItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Delete delete navigation property deviceHealthScriptStates for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, error) {
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
// Patch update the navigation property deviceHealthScriptStates in users
// returns a DeviceHealthScriptPolicyStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, requestConfiguration *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, error) {
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
// ToDeleteRequestInformation delete navigation property deviceHealthScriptStates for users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceHealthScriptStates in users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, requestConfiguration *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder when successful
func (m *ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    return NewItemManageddevicesItemDevicehealthscriptstatesWithidwithpolicyidwithdeviceidWithIdWithPolicyIdWithDeviceIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
