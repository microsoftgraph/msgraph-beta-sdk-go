package accessreviews

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInstancesItemSendReminderRequestBuilder provides operations to call the sendReminder method.
type ItemInstancesItemSendReminderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInstancesItemSendReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstancesItemSendReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInstancesItemSendReminderRequestBuilderInternal instantiates a new ItemInstancesItemSendReminderRequestBuilder and sets the default values.
func NewItemInstancesItemSendReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesItemSendReminderRequestBuilder) {
    m := &ItemInstancesItemSendReminderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/accessReviews/{accessReview%2Did}/instances/{accessReview%2Did1}/sendReminder", pathParameters),
    }
    return m
}
// NewItemInstancesItemSendReminderRequestBuilder instantiates a new ItemInstancesItemSendReminderRequestBuilder and sets the default values.
func NewItemInstancesItemSendReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesItemSendReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstancesItemSendReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in the Microsoft Entra access reviews feature, send a reminder to the reviewers of a currently active accessReview.  The target object can be either a one-time access review, or an instance of a recurring access review. 
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accessreview-sendreminder?view=graph-rest-beta
func (m *ItemInstancesItemSendReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemInstancesItemSendReminderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation in the Microsoft Entra access reviews feature, send a reminder to the reviewers of a currently active accessReview.  The target object can be either a one-time access review, or an instance of a recurring access review. 
// returns a *RequestInformation when successful
func (m *ItemInstancesItemSendReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemInstancesItemSendReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInstancesItemSendReminderRequestBuilder when successful
func (m *ItemInstancesItemSendReminderRequestBuilder) WithUrl(rawUrl string)(*ItemInstancesItemSendReminderRequestBuilder) {
    return NewItemInstancesItemSendReminderRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
