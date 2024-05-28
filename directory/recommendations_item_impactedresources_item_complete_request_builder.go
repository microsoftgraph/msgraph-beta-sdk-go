package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RecommendationsItemImpactedresourcesItemCompleteRequestBuilder provides operations to call the complete method.
type RecommendationsItemImpactedresourcesItemCompleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RecommendationsItemImpactedresourcesItemCompleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RecommendationsItemImpactedresourcesItemCompleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRecommendationsItemImpactedresourcesItemCompleteRequestBuilderInternal instantiates a new RecommendationsItemImpactedresourcesItemCompleteRequestBuilder and sets the default values.
func NewRecommendationsItemImpactedresourcesItemCompleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RecommendationsItemImpactedresourcesItemCompleteRequestBuilder) {
    m := &RecommendationsItemImpactedresourcesItemCompleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/recommendations/{recommendation%2Did}/impactedResources/{impactedResource%2Did}/complete", pathParameters),
    }
    return m
}
// NewRecommendationsItemImpactedresourcesItemCompleteRequestBuilder instantiates a new RecommendationsItemImpactedresourcesItemCompleteRequestBuilder and sets the default values.
func NewRecommendationsItemImpactedresourcesItemCompleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RecommendationsItemImpactedresourcesItemCompleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRecommendationsItemImpactedresourcesItemCompleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post complete an impactedResource object and update its status to completedByUser.
// returns a ImpactedResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/impactedresource-complete?view=graph-rest-beta
func (m *RecommendationsItemImpactedresourcesItemCompleteRequestBuilder) Post(ctx context.Context, requestConfiguration *RecommendationsItemImpactedresourcesItemCompleteRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImpactedResourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation complete an impactedResource object and update its status to completedByUser.
// returns a *RequestInformation when successful
func (m *RecommendationsItemImpactedresourcesItemCompleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *RecommendationsItemImpactedresourcesItemCompleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RecommendationsItemImpactedresourcesItemCompleteRequestBuilder when successful
func (m *RecommendationsItemImpactedresourcesItemCompleteRequestBuilder) WithUrl(rawUrl string)(*RecommendationsItemImpactedresourcesItemCompleteRequestBuilder) {
    return NewRecommendationsItemImpactedresourcesItemCompleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
