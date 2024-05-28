package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
type EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters the embedded SIM activation code pools created by this account.
type EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters
}
// EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *EmbeddedsimactivationcodepoolsItemAssignRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Assign()(*EmbeddedsimactivationcodepoolsItemAssignRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.embeddedSIMActivationCodePool entity.
// returns a *EmbeddedsimactivationcodepoolsItemAssignmentsRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Assignments()(*EmbeddedsimactivationcodepoolsItemAssignmentsRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal instantiates a new EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) {
    m := &EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder instantiates a new EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property embeddedSIMActivationCodePools for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStates provides operations to manage the deviceStates property of the microsoft.graph.embeddedSIMActivationCodePool entity.
// returns a *EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) DeviceStates()(*EmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsItemDevicestatesDeviceStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the embedded SIM activation code pools created by this account.
// returns a EmbeddedSIMActivationCodePoolable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable), nil
}
// Patch update the navigation property embeddedSIMActivationCodePools in deviceManagement
// returns a EmbeddedSIMActivationCodePoolable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable), nil
}
// ToDeleteRequestInformation delete navigation property embeddedSIMActivationCodePools for deviceManagement
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the embedded SIM activation code pools created by this account.
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property embeddedSIMActivationCodePools in deviceManagement
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) WithUrl(rawUrl string)(*EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
