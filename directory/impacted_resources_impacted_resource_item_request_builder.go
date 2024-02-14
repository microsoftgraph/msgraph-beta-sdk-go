package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImpactedResourcesImpactedResourceItemRequestBuilder provides operations to manage the impactedResources property of the microsoft.graph.directory entity.
type ImpactedResourcesImpactedResourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ImpactedResourcesImpactedResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImpactedResourcesImpactedResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ImpactedResourcesImpactedResourceItemRequestBuilderGetQueryParameters get impactedResources from directory
type ImpactedResourcesImpactedResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ImpactedResourcesImpactedResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImpactedResourcesImpactedResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ImpactedResourcesImpactedResourceItemRequestBuilderGetQueryParameters
}
// ImpactedResourcesImpactedResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImpactedResourcesImpactedResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Complete provides operations to call the complete method.
// returns a *ImpactedResourcesItemCompleteRequestBuilder when successful
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) Complete()(*ImpactedResourcesItemCompleteRequestBuilder) {
    return NewImpactedResourcesItemCompleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewImpactedResourcesImpactedResourceItemRequestBuilderInternal instantiates a new ImpactedResourcesImpactedResourceItemRequestBuilder and sets the default values.
func NewImpactedResourcesImpactedResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImpactedResourcesImpactedResourceItemRequestBuilder) {
    m := &ImpactedResourcesImpactedResourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/impactedResources/{impactedResource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewImpactedResourcesImpactedResourceItemRequestBuilder instantiates a new ImpactedResourcesImpactedResourceItemRequestBuilder and sets the default values.
func NewImpactedResourcesImpactedResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImpactedResourcesImpactedResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImpactedResourcesImpactedResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property impactedResources for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ImpactedResourcesImpactedResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Dismiss provides operations to call the dismiss method.
// returns a *ImpactedResourcesItemDismissRequestBuilder when successful
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) Dismiss()(*ImpactedResourcesItemDismissRequestBuilder) {
    return NewImpactedResourcesItemDismissRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get impactedResources from directory
// returns a ImpactedResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ImpactedResourcesImpactedResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImpactedResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable), nil
}
// Patch update the navigation property impactedResources in directory
// returns a ImpactedResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable, requestConfiguration *ImpactedResourcesImpactedResourceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImpactedResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable), nil
}
// Postpone provides operations to call the postpone method.
// returns a *ImpactedResourcesItemPostponeRequestBuilder when successful
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) Postpone()(*ImpactedResourcesItemPostponeRequestBuilder) {
    return NewImpactedResourcesItemPostponeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reactivate provides operations to call the reactivate method.
// returns a *ImpactedResourcesItemReactivateRequestBuilder when successful
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) Reactivate()(*ImpactedResourcesItemReactivateRequestBuilder) {
    return NewImpactedResourcesItemReactivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property impactedResources for directory
// returns a *RequestInformation when successful
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ImpactedResourcesImpactedResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/directory/impactedResources/{impactedResource%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get impactedResources from directory
// returns a *RequestInformation when successful
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ImpactedResourcesImpactedResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property impactedResources in directory
// returns a *RequestInformation when successful
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable, requestConfiguration *ImpactedResourcesImpactedResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/directory/impactedResources/{impactedResource%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ImpactedResourcesImpactedResourceItemRequestBuilder when successful
func (m *ImpactedResourcesImpactedResourceItemRequestBuilder) WithUrl(rawUrl string)(*ImpactedResourcesImpactedResourceItemRequestBuilder) {
    return NewImpactedResourcesImpactedResourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
