package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
type ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetQueryParameters security baseline states for this device.
type ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetQueryParameters
}
// ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal instantiates a new ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder and sets the default values.
func NewManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    m := &ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/securityBaselineStates/{securityBaselineState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder instantiates a new ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder and sets the default values.
func NewManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property securityBaselineStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get security baseline states for this device.
// returns a SecurityBaselineStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityBaselineStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable), nil
}
// Patch update the navigation property securityBaselineStates in deviceManagement
// returns a SecurityBaselineStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable, requestConfiguration *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityBaselineStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable), nil
}
// SettingStates provides operations to manage the settingStates property of the microsoft.graph.securityBaselineState entity.
// returns a *ManagedDevicesItemSecurityBaselineStatesItemSettingStatesRequestBuilder when successful
func (m *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) SettingStates()(*ManagedDevicesItemSecurityBaselineStatesItemSettingStatesRequestBuilder) {
    return NewManagedDevicesItemSecurityBaselineStatesItemSettingStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property securityBaselineStates for deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/securityBaselineStates/{securityBaselineState%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation security baseline states for this device.
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property securityBaselineStates in deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable, requestConfiguration *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/securityBaselineStates/{securityBaselineState%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder when successful
func (m *ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    return NewManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
