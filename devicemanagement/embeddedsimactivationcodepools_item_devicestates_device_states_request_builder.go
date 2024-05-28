package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder provides operations to manage the deviceStates property of the microsoft.graph.embeddedSIMActivationCodePool entity.
type EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderGetQueryParameters navigational property to a list of device states for this pool.
type EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderGetQueryParameters struct {
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
// EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderGetQueryParameters
}
// EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEmbeddedSIMDeviceStateId provides operations to manage the deviceStates property of the microsoft.graph.embeddedSIMActivationCodePool entity.
// returns a *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) ByEmbeddedSIMDeviceStateId(embeddedSIMDeviceStateId string)(*EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if embeddedSIMDeviceStateId != "" {
        urlTplParams["embeddedSIMDeviceState%2Did"] = embeddedSIMDeviceStateId
    }
    return NewEmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderInternal instantiates a new EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) {
    m := &EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool%2Did}/deviceStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder instantiates a new EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EmbeddedsimactivationcodepoolsItemDevicestatesCountRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) Count()(*EmbeddedsimactivationcodepoolsItemDevicestatesCountRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsItemDevicestatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get navigational property to a list of device states for this pool.
// returns a EmbeddedSIMDeviceStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMDeviceStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateCollectionResponseable), nil
}
// Post create new navigation property to deviceStates for deviceManagement
// returns a EmbeddedSIMDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, requestConfiguration *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable), nil
}
// ToGetRequestInformation navigational property to a list of device states for this pool.
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceStates for deviceManagement
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, requestConfiguration *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) WithUrl(rawUrl string)(*EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
