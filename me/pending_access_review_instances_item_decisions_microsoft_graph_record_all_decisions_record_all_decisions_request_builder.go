package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilder provides operations to call the recordAllDecisions method.
type PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilderInternal instantiates a new RecordAllDecisionsRequestBuilder and sets the default values.
func NewPendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilder) {
    m := &PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance%2Did}/decisions/microsoft.graph.recordAllDecisions";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilder instantiates a new RecordAllDecisionsRequestBuilder and sets the default values.
func NewPendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post as a reviewer of an access review, record a decision for an accessReviewInstanceDecisionItem that is assigned to you and that matches the principal or resource IDs specified. If no IDs are specified, the decisions will apply to every **accessReviewInstanceDecisionItem** for which you are the reviewer.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accessreviewinstancedecisionitem-recordalldecisions?view=graph-rest-1.0
func (m *PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilder) Post(ctx context.Context, body PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsPostRequestBodyable, requestConfiguration *PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation as a reviewer of an access review, record a decision for an accessReviewInstanceDecisionItem that is assigned to you and that matches the principal or resource IDs specified. If no IDs are specified, the decisions will apply to every **accessReviewInstanceDecisionItem** for which you are the reviewer.
func (m *PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsPostRequestBodyable, requestConfiguration *PendingAccessReviewInstancesItemDecisionsMicrosoftGraphRecordAllDecisionsRecordAllDecisionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
