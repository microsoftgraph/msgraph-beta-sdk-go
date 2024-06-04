package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
type AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderGetQueryParameters get the reviewers for an access review instance, irrespective of whether or not they have received a notification. The reviewers are represented by an accessReviewReviewer object. A list of zero or more objects are returned, including all of their nested properties.
type AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderGetQueryParameters struct {
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
// AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderGetQueryParameters
}
// AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessReviewReviewerId provides operations to manage the contactedReviewers property of the microsoft.graph.accessReviewInstance entity.
// returns a *AccessreviewsDefinitionsItemInstancesItemContactedreviewersAccessReviewReviewerItemRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) ByAccessReviewReviewerId(accessReviewReviewerId string)(*AccessreviewsDefinitionsItemInstancesItemContactedreviewersAccessReviewReviewerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessReviewReviewerId != "" {
        urlTplParams["accessReviewReviewer%2Did"] = accessReviewReviewerId
    }
    return NewAccessreviewsDefinitionsItemInstancesItemContactedreviewersAccessReviewReviewerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderInternal instantiates a new AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) {
    m := &AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/contactedReviewers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder instantiates a new AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder and sets the default values.
func NewAccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AccessreviewsDefinitionsItemInstancesItemContactedreviewersCountRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) Count()(*AccessreviewsDefinitionsItemInstancesItemContactedreviewersCountRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemContactedreviewersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the reviewers for an access review instance, irrespective of whether or not they have received a notification. The reviewers are represented by an accessReviewReviewer object. A list of zero or more objects are returned, including all of their nested properties.
// returns a AccessReviewReviewerCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstance-list-contactedreviewers?view=graph-rest-beta
func (m *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewReviewerCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerCollectionResponseable), nil
}
// Post create new navigation property to contactedReviewers for identityGovernance
// returns a AccessReviewReviewerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerable, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessReviewReviewerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerable), nil
}
// ToGetRequestInformation get the reviewers for an access review instance, irrespective of whether or not they have received a notification. The reviewers are represented by an accessReviewReviewer object. A list of zero or more objects are returned, including all of their nested properties.
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to contactedReviewers for identityGovernance
// returns a *RequestInformation when successful
func (m *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessReviewReviewerable, requestConfiguration *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder when successful
func (m *AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder) {
    return NewAccessreviewsDefinitionsItemInstancesItemContactedreviewersContactedReviewersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
