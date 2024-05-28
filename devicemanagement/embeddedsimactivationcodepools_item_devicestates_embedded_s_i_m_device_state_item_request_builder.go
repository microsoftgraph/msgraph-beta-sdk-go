package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder provides operations to manage the deviceStates property of the microsoft.graph.embeddedSIMActivationCodePool entity.
type EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderGetQueryParameters navigational property to a list of device states for this pool.
type EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderGetQueryParameters
}
// EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderInternal instantiates a new EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) {
    m := &EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool%2Did}/deviceStates/{embeddedSIMDeviceState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder instantiates a new EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get navigational property to a list of device states for this pool.
// returns a EmbeddedSIMDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property deviceStates in deviceManagement
// returns a EmbeddedSIMDeviceStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, requestConfiguration *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property deviceStates for deviceManagement
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation navigational property to a list of device states for this pool.
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceStates in deviceManagement
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, requestConfiguration *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) WithUrl(rawUrl string)(*EmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsItemDevicestatesEmbeddedSIMDeviceStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
