package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder provides operations to call the filterByCurrentUser method.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderGetQueryParameters invoke function filterByCurrentUser
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderGetQueryParameters struct {
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
// AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderGetQueryParameters
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderInternal instantiates a new AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, on *string)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/stages/{accessReviewStage%2Did}/decisions/filterByCurrentUser(on='{on}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if on != nil {
        m.BaseRequestBuilder.PathParameters["on"] = *on
    }
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder instantiates a new AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function filterByCurrentUser
// Deprecated: This method is obsolete. Use GetAsFilterByCurrentUserWithOnGetResponse instead.
// returns a AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnResponseable), nil
}
// GetAsFilterByCurrentUserWithOnGetResponse invoke function filterByCurrentUser
// returns a AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder) GetAsFilterByCurrentUserWithOnGetResponse(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnGetResponseable), nil
}
// ToGetRequestInformation invoke function filterByCurrentUser
// returns a *RequestInformation when successful
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder when successful
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder) WithUrl(rawUrl string)(*AccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsItemInstanceStagesItemDecisionsFilterByCurrentUserWithOnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
