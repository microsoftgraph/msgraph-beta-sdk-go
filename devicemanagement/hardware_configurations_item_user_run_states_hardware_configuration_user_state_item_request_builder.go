package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder provides operations to manage the userRunStates property of the microsoft.graph.hardwareConfiguration entity.
type HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderGetQueryParameters list of run states for the hardware configuration across all users. Read-Only.
type HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderGetQueryParameters
}
// HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewHardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderInternal instantiates a new HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder and sets the default values.
func NewHardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) {
    m := &HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/hardwareConfigurations/{hardwareConfiguration%2Did}/userRunStates/{hardwareConfigurationUserState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewHardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder instantiates a new HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder and sets the default values.
func NewHardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userRunStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of run states for the hardware configuration across all users. Read-Only.
// returns a HardwareConfigurationUserStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareConfigurationUserStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateable), nil
}
// Patch update the navigation property userRunStates in deviceManagement
// returns a HardwareConfigurationUserStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateable, requestConfiguration *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareConfigurationUserStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateable), nil
}
// ToDeleteRequestInformation delete navigation property userRunStates for deviceManagement
// returns a *RequestInformation when successful
func (m *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of run states for the hardware configuration across all users. Read-Only.
// returns a *RequestInformation when successful
func (m *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userRunStates in deviceManagement
// returns a *RequestInformation when successful
func (m *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateable, requestConfiguration *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder when successful
func (m *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) WithUrl(rawUrl string)(*HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) {
    return NewHardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
