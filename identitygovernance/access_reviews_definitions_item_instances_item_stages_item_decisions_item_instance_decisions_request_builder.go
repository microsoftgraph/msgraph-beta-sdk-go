package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetQueryParameters each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetQueryParameters struct {
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
// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetQueryParameters
}
// AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessReviewInstanceDecisionItemId1 provides operations to manage the decisions property of the microsoft.graph.accessReviewInstance entity.
// returns a *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) ByAccessReviewInstanceDecisionItemId1(accessReviewInstanceDecisionItemId1 string)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessReviewInstanceDecisionItemId1 != "" {
        urlTplParams["accessReviewInstanceDecisionItem%2Did1"] = accessReviewInstanceDecisionItemId1
    }
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsAccessReviewInstanceDecisionItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderInternal instantiates a new AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/decisions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder instantiates a new AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) Count()(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnRequestBuilder when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
// returns a AccessReviewInstanceDecisionItemCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Post create new navigation property to decisions for identityGovernance
// returns a AccessReviewInstanceDecisionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRecordAllDecisionsRequestBuilder when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) RecordAllDecisions()(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRecordAllDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRecordAllDecisionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
// returns a *RequestInformation when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to decisions for identityGovernance
// returns a *RequestInformation when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewInstanceDecisionItemable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder when successful
func (m *AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) WithUrl(rawUrl string)(*AccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemStagesItemDecisionsItemInstanceDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
