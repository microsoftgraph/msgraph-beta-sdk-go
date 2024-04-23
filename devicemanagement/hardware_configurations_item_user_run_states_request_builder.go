package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// HardwareConfigurationsItemUserRunStatesRequestBuilder provides operations to manage the userRunStates property of the microsoft.graph.hardwareConfiguration entity.
type HardwareConfigurationsItemUserRunStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HardwareConfigurationsItemUserRunStatesRequestBuilderGetQueryParameters list of run states for the hardware configuration across all users. Read-Only.
type HardwareConfigurationsItemUserRunStatesRequestBuilderGetQueryParameters struct {
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
// HardwareConfigurationsItemUserRunStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareConfigurationsItemUserRunStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HardwareConfigurationsItemUserRunStatesRequestBuilderGetQueryParameters
}
// HardwareConfigurationsItemUserRunStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareConfigurationsItemUserRunStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHardwareConfigurationUserStateId provides operations to manage the userRunStates property of the microsoft.graph.hardwareConfiguration entity.
// returns a *HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder when successful
func (m *HardwareConfigurationsItemUserRunStatesRequestBuilder) ByHardwareConfigurationUserStateId(hardwareConfigurationUserStateId string)(*HardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hardwareConfigurationUserStateId != "" {
        urlTplParams["hardwareConfigurationUserState%2Did"] = hardwareConfigurationUserStateId
    }
    return NewHardwareConfigurationsItemUserRunStatesHardwareConfigurationUserStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewHardwareConfigurationsItemUserRunStatesRequestBuilderInternal instantiates a new HardwareConfigurationsItemUserRunStatesRequestBuilder and sets the default values.
func NewHardwareConfigurationsItemUserRunStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareConfigurationsItemUserRunStatesRequestBuilder) {
    m := &HardwareConfigurationsItemUserRunStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/hardwareConfigurations/{hardwareConfiguration%2Did}/userRunStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewHardwareConfigurationsItemUserRunStatesRequestBuilder instantiates a new HardwareConfigurationsItemUserRunStatesRequestBuilder and sets the default values.
func NewHardwareConfigurationsItemUserRunStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareConfigurationsItemUserRunStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHardwareConfigurationsItemUserRunStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *HardwareConfigurationsItemUserRunStatesCountRequestBuilder when successful
func (m *HardwareConfigurationsItemUserRunStatesRequestBuilder) Count()(*HardwareConfigurationsItemUserRunStatesCountRequestBuilder) {
    return NewHardwareConfigurationsItemUserRunStatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of run states for the hardware configuration across all users. Read-Only.
// returns a HardwareConfigurationUserStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareConfigurationsItemUserRunStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *HardwareConfigurationsItemUserRunStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareConfigurationUserStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateCollectionResponseable), nil
}
// Post create new navigation property to userRunStates for deviceManagement
// returns a HardwareConfigurationUserStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareConfigurationsItemUserRunStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateable, requestConfiguration *HardwareConfigurationsItemUserRunStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation list of run states for the hardware configuration across all users. Read-Only.
// returns a *RequestInformation when successful
func (m *HardwareConfigurationsItemUserRunStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HardwareConfigurationsItemUserRunStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userRunStates for deviceManagement
// returns a *RequestInformation when successful
func (m *HardwareConfigurationsItemUserRunStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationUserStateable, requestConfiguration *HardwareConfigurationsItemUserRunStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HardwareConfigurationsItemUserRunStatesRequestBuilder when successful
func (m *HardwareConfigurationsItemUserRunStatesRequestBuilder) WithUrl(rawUrl string)(*HardwareConfigurationsItemUserRunStatesRequestBuilder) {
    return NewHardwareConfigurationsItemUserRunStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
