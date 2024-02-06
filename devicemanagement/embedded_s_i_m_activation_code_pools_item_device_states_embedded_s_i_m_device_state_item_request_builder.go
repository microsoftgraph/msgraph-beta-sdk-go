package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder provides operations to manage the deviceStates property of the microsoft.graph.embeddedSIMActivationCodePool entity.
type EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetQueryParameters navigational property to a list of device states for this pool.
type EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetQueryParameters
}
// EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderInternal instantiates a new EmbeddedSIMDeviceStateItemRequestBuilder and sets the default values.
func NewEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) {
    m := &EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool%2Did}/deviceStates/{embeddedSIMDeviceState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder instantiates a new EmbeddedSIMDeviceStateItemRequestBuilder and sets the default values.
func NewEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceStates for deviceManagement
func (m *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get navigational property to a list of device states for this pool.
func (m *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, requestConfiguration *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation navigational property to a list of device states for this pool.
func (m *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, requestConfiguration *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) WithUrl(rawUrl string)(*EmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) {
    return NewEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
