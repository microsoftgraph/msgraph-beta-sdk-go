package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilder provides operations to call the acceptRecommendations method.
type AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilderInternal instantiates a new MicrosoftGraphAcceptRecommendationsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/microsoft.graph.acceptRecommendations";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilder instantiates a new MicrosoftGraphAcceptRecommendationsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post allows the acceptance of recommendations on all accessReviewInstanceDecisionItem objects that have not been reviewed for an accessReviewInstance object for which the calling user is a reviewer. Recommendations are generated if **recommendationsEnabled** is `true` on the accessReviewScheduleDefinition object. If there is not a recommendation on an accessReviewInstanceDecisionItem object no decision will be recorded.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accessreviewinstance-acceptrecommendations?view=graph-rest-1.0
func (m *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation allows the acceptance of recommendations on all accessReviewInstanceDecisionItem objects that have not been reviewed for an accessReviewInstance object for which the calling user is a reviewer. Recommendations are generated if **recommendationsEnabled** is `true` on the accessReviewScheduleDefinition object. If there is not a recommendation on an accessReviewInstanceDecisionItem object no decision will be recorded.
func (m *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemMicrosoftGraphAcceptRecommendationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
