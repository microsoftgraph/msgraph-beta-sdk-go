package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.embeddedSIMActivationCodePool entity.
type EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderGetQueryParameters navigational property to a list of targets to which this pool is assigned.
type EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderGetQueryParameters
}
// EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderInternal instantiates a new EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) {
    m := &EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool%2Did}/assignments/{embeddedSIMActivationCodePoolAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder instantiates a new EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get navigational property to a list of targets to which this pool is assigned.
// returns a EmbeddedSIMActivationCodePoolAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolAssignmentable), nil
}
// Patch update the navigation property assignments in deviceManagement
// returns a EmbeddedSIMActivationCodePoolAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolAssignmentable, requestConfiguration *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolAssignmentable), nil
}
// ToDeleteRequestInformation delete navigation property assignments for deviceManagement
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation navigational property to a list of targets to which this pool is assigned.
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in deviceManagement
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolAssignmentable, requestConfiguration *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*EmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsItemAssignmentsEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
