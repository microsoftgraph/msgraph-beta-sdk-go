package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder provides operations to manage the settingStates property of the microsoft.graph.securityBaselineState entity.
type ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetQueryParameters the security baseline state for different settings for a device
type ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetQueryParameters
}
// ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderInternal instantiates a new ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder and sets the default values.
func NewItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) {
    m := &ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/securityBaselineStates/{securityBaselineState%2Did}/settingStates/{securityBaselineSettingState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder instantiates a new ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder and sets the default values.
func NewItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property settingStates for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, error) {
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
// Patch update the navigation property settingStates in users
// returns a SecurityBaselineSettingStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, requestConfiguration *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, error) {
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
// ToDeleteRequestInformation delete navigation property settingStates for users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property settingStates in users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, requestConfiguration *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder when successful
func (m *ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) {
    return NewItemManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
