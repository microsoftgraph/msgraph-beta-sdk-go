package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder provides operations to call the recordAllDecisions method.
type AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilderInternal instantiates a new RecordAllDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder) {
    m := &AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/definitions/{accessReviewScheduleDefinition%2Did}/instances/{accessReviewInstance%2Did}/decisions/recordAllDecisions", pathParameters),
    }
    return m
}
// NewAccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder instantiates a new RecordAllDecisionsRequestBuilder and sets the default values.
func NewAccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post as a reviewer of an access review, record a decision for an accessReviewInstanceDecisionItem that is assigned to you and that matches the principal or resource IDs specified. If no IDs are specified, the decisions will apply to every accessReviewInstanceDecisionItem for which you are the reviewer.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreviewinstancedecisionitem-recordalldecisions?view=graph-rest-1.0
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder) Post(ctx context.Context, body AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsPostRequestBodyable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation as a reviewer of an access review, record a decision for an accessReviewInstanceDecisionItem that is assigned to you and that matches the principal or resource IDs specified. If no IDs are specified, the decisions will apply to every accessReviewInstanceDecisionItem for which you are the reviewer.
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsPostRequestBodyable, requestConfiguration *AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder) WithUrl(rawUrl string)(*AccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder) {
    return NewAccessReviewsDefinitionsItemInstancesItemDecisionsRecordAllDecisionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
