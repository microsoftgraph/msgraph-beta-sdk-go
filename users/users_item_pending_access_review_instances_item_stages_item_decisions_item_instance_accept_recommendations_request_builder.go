package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder provides operations to call the acceptRecommendations method.
type UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilderInternal instantiates a new AcceptRecommendationsRequestBuilder and sets the default values.
func NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder) {
    m := &UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/microsoft.graph.acceptRecommendations";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder instantiates a new AcceptRecommendationsRequestBuilder and sets the default values.
func NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation allows the acceptance of recommendations on all accessReviewInstanceDecisionItem objects that have not been reviewed for an accessReviewInstance object for which the calling user is a reviewer. Recommendations are generated if **recommendationsEnabled** is `true` on the accessReviewScheduleDefinition object. If there is not a recommendation on an accessReviewInstanceDecisionItem object no decision will be recorded.
func (m *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post allows the acceptance of recommendations on all accessReviewInstanceDecisionItem objects that have not been reviewed for an accessReviewInstance object for which the calling user is a reviewer. Recommendations are generated if **recommendationsEnabled** is `true` on the accessReviewScheduleDefinition object. If there is not a recommendation on an accessReviewInstanceDecisionItem object no decision will be recorded.
func (m *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilder) Post(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceAcceptRecommendationsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
