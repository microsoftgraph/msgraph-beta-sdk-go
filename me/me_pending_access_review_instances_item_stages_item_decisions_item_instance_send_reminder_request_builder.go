package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder provides operations to call the sendReminder method.
type MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderInternal instantiates a new SendReminderRequestBuilder and sets the default values.
func NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    m := &MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/pendingAccessReviewInstances/{accessReviewInstance%2Did}/stages/{accessReviewStage%2Did}/decisions/{accessReviewInstanceDecisionItem%2Did}/instance/microsoft.graph.sendReminder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder instantiates a new SendReminderRequestBuilder and sets the default values.
func NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation send a reminder to the reviewers of a currently active accessReviewInstance.
func (m *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post send a reminder to the reviewers of a currently active accessReviewInstance.
func (m *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *MePendingAccessReviewInstancesItemStagesItemDecisionsItemInstanceSendReminderRequestBuilderPostRequestConfiguration)(error) {
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
