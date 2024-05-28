package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder provides operations to manage the impactedResources property of the microsoft.graph.recommendationBase entity.
type RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderGetQueryParameters read the properties and relationships of an impactedResource object.
type RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderGetQueryParameters
}
// RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Complete provides operations to call the complete method.
// returns a *RecommendationsItemImpactedresourcesItemCompleteRequestBuilder when successful
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) Complete()(*RecommendationsItemImpactedresourcesItemCompleteRequestBuilder) {
    return NewRecommendationsItemImpactedresourcesItemCompleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewRecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderInternal instantiates a new RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder and sets the default values.
func NewRecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) {
    m := &RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/recommendations/{recommendation%2Did}/impactedResources/{impactedResource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder instantiates a new RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder and sets the default values.
func NewRecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property impactedResources for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *RecommendationsItemImpactedresourcesItemDismissRequestBuilder when successful
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) Dismiss()(*RecommendationsItemImpactedresourcesItemDismissRequestBuilder) {
    return NewRecommendationsItemImpactedresourcesItemDismissRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of an impactedResource object.
// returns a ImpactedResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/impactedresource-get?view=graph-rest-beta
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable, error) {
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
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable, requestConfiguration *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable, error) {
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
// returns a *RecommendationsItemImpactedresourcesItemPostponeRequestBuilder when successful
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) Postpone()(*RecommendationsItemImpactedresourcesItemPostponeRequestBuilder) {
    return NewRecommendationsItemImpactedresourcesItemPostponeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reactivate provides operations to call the reactivate method.
// returns a *RecommendationsItemImpactedresourcesItemReactivateRequestBuilder when successful
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) Reactivate()(*RecommendationsItemImpactedresourcesItemReactivateRequestBuilder) {
    return NewRecommendationsItemImpactedresourcesItemReactivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property impactedResources for directory
// returns a *RequestInformation when successful
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an impactedResource object.
// returns a *RequestInformation when successful
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable, requestConfiguration *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder when successful
func (m *RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) WithUrl(rawUrl string)(*RecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder) {
    return NewRecommendationsItemImpactedresourcesImpactedResourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
