package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder provides operations to manage the settingStates property of the microsoft.graph.securityBaselineState entity.
type ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderGetQueryParameters the security baseline state for different settings for a device
type ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderGetQueryParameters struct {
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
// ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderGetQueryParameters
}
// ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySecurityBaselineSettingStateId provides operations to manage the settingStates property of the microsoft.graph.securityBaselineState entity.
// returns a *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder when successful
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder) BySecurityBaselineSettingStateId(securityBaselineSettingStateId string)(*ManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if securityBaselineSettingStateId != "" {
        urlTplParams["securityBaselineSettingState%2Did"] = securityBaselineSettingStateId
    }
    return NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSecurityBaselineSettingStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderInternal instantiates a new ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder and sets the default values.
func NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder) {
    m := &ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/securityBaselineStates/{securityBaselineState%2Did}/settingStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder instantiates a new ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder and sets the default values.
func NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManageddevicesItemSecuritybaselinestatesItemSettingstatesCountRequestBuilder when successful
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder) Count()(*ManageddevicesItemSecuritybaselinestatesItemSettingstatesCountRequestBuilder) {
    return NewManageddevicesItemSecuritybaselinestatesItemSettingstatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the security baseline state for different settings for a device
// returns a SecurityBaselineSettingStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityBaselineSettingStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateCollectionResponseable), nil
}
// Post create new navigation property to settingStates for deviceManagement
// returns a SecurityBaselineSettingStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, requestConfiguration *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the security baseline state for different settings for a device
// returns a *RequestInformation when successful
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to settingStates for deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SecurityBaselineSettingStateable, requestConfiguration *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder when successful
func (m *ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder) {
    return NewManageddevicesItemSecuritybaselinestatesItemSettingstatesSettingStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
