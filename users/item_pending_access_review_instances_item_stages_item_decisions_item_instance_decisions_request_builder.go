package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
type ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetQueryParameters retrieve the accessReviewInstanceDecisionItem objects for a specific accessReviewInstance. A list of zero or more accessReviewInstanceDecisionItem objects are returned, including all of their nested properties. This API is available in the following national cloud deployments.
type ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetQueryParameters struct {
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
// ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetQueryParameters
}
// ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessReviewInstanceDecisionItemId1 provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) ByAccessReviewInstanceDecisionItemId1(accessReviewInstanceDecisionItemId1 string)(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessReviewInstanceDecisionItemId1 != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did1"] = accessReviewInstanceDecisionItemId1
    }
    return NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderInternal instantiates a new DecisionsRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) {
    m := &ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/decisions{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder instantiates a new DecisionsRequestBuilder and sets the default values.
func NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) Count()(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get retrieve the accessReviewInstanceDecisionItem objects for a specific accessReviewInstance. A list of zero or more accessReviewInstanceDecisionItem objects are returned, including all of their nested properties. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-list-decisions?view=graph-rest-1.0
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceDecisionItemCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemCollectionResponseable), nil
}
// Post create new navigation property to decisions for users
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable), nil
}
// RecordAllDecisions provides operations to call the recordAllDecisions method.
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) RecordAllDecisions()(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRecordAllDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRecordAllDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation retrieve the accessReviewInstanceDecisionItem objects for a specific accessReviewInstance. A list of zero or more accessReviewInstanceDecisionItem objects are returned, including all of their nested properties. This API is available in the following national cloud deployments.
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to decisions for users
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) WithUrl(rawUrl string)(*ItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) {
    return NewItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
