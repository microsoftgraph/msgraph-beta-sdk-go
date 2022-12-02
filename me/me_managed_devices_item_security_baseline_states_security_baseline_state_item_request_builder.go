package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder provides operations to manage the securityBaselineStates property of the microsoft.graph.managedDevice entity.
type MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetQueryParameters security baseline states for this device.
type MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetQueryParameters
}
// MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal instantiates a new SecurityBaselineStateItemRequestBuilder and sets the default values.
func NewMeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    m := &MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/managedDevices/{managedDevice%2Did}/securityBaselineStates/{securityBaselineState%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder instantiates a new SecurityBaselineStateItemRequestBuilder and sets the default values.
func NewMeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property securityBaselineStates for me
func (m *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation security baseline states for this device.
func (m *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property securityBaselineStates in me
func (m *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable, requestConfiguration *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property securityBaselineStates for me
func (m *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get security baseline states for this device.
func (m *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityBaselineStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable), nil
}
// Patch update the navigation property securityBaselineStates in me
func (m *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable, requestConfiguration *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityBaselineStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineStateable), nil
}
// SettingStates provides operations to manage the settingStates property of the microsoft.graph.securityBaselineState entity.
func (m *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) SettingStates()(*MeManagedDevicesItemSecurityBaselineStatesItemSettingStatesRequestBuilder) {
    return NewMeManagedDevicesItemSecurityBaselineStatesItemSettingStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingStatesById provides operations to manage the settingStates property of the microsoft.graph.securityBaselineState entity.
func (m *MeManagedDevicesItemSecurityBaselineStatesSecurityBaselineStateItemRequestBuilder) SettingStatesById(id string)(*MeManagedDevicesItemSecurityBaselineStatesItemSettingStatesSecurityBaselineSettingStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityBaselineSettingState%2Did"] = id
    }
    return NewMeManagedDevicesItemSecurityBaselineStatesItemSettingStatesSecurityBaselineSettingStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
