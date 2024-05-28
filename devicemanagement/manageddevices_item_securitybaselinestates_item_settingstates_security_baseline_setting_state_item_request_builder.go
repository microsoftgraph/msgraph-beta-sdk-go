package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder provides operations to manage the settingStates property of the microsoft.graph.securityBaselineState entity.
type ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetQueryParameters the security baseline state for different settings for a device
type ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetQueryParameters
}
// ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderInternal instantiates a new ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder and sets the default values.
func NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) {
    m := &ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/securityBaselineStates/{securityBaselineState%2Did}/settingStates/{securityBaselineSettingState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder instantiates a new ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder and sets the default values.
func NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property settingStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the security baseline state for different settings for a device
// returns a SecurityBaselineSettingStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityBaselineSettingStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable), nil
}
// Patch update the navigation property settingStates in deviceManagement
// returns a SecurityBaselineSettingStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, requestConfiguration *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityBaselineSettingStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable), nil
}
// ToDeleteRequestInformation delete navigation property settingStates for deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the security baseline state for different settings for a device
// returns a *RequestInformation when successful
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property settingStates in deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, requestConfiguration *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder when successful
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) {
    return NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
