package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder provides operations to manage the deviceHealthScriptStates property of the microsoft.graph.managedDevice entity.
type ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters results of device health scripts that ran for this device. Default is empty list. This property is read-only.
type ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetQueryParameters
}
// ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal instantiates a new ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder and sets the default values.
func NewComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceId *string, id *string, policyId *string)(*ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    m := &ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder{
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
// NewComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder instantiates a new ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder and sets the default values.
func NewComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Delete delete navigation property deviceHealthScriptStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, error) {
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
func (m *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, requestConfiguration *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, error) {
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
func (m *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptPolicyStateable, requestConfiguration *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder when successful
func (m *ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder) {
    return NewComanagedDevicesItemDeviceHealthScriptStatesWithIdWithPolicyIdWithDeviceIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
