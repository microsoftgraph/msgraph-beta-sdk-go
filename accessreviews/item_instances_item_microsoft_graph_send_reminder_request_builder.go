package accessreviews

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInstancesItemMicrosoftGraphSendReminderRequestBuilder provides operations to call the sendReminder method.
type ItemInstancesItemMicrosoftGraphSendReminderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemInstancesItemMicrosoftGraphSendReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstancesItemMicrosoftGraphSendReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInstancesItemMicrosoftGraphSendReminderRequestBuilderInternal instantiates a new MicrosoftGraphSendReminderRequestBuilder and sets the default values.
func NewItemInstancesItemMicrosoftGraphSendReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesItemMicrosoftGraphSendReminderRequestBuilder) {
    m := &ItemInstancesItemMicrosoftGraphSendReminderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/accessReviews/{accessReview%2Did}/instances/{accessReview%2Did1}/microsoft.graph.sendReminder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemInstancesItemMicrosoftGraphSendReminderRequestBuilder instantiates a new MicrosoftGraphSendReminderRequestBuilder and sets the default values.
func NewItemInstancesItemMicrosoftGraphSendReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstancesItemMicrosoftGraphSendReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstancesItemMicrosoftGraphSendReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in the Azure AD access reviews feature, send a reminder to the reviewers of a currently active accessReview.  The target object can be either a one-time access review, or an instance of a recurring access review. 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accessreview-sendreminder?view=graph-rest-1.0
func (m *ItemInstancesItemMicrosoftGraphSendReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemInstancesItemMicrosoftGraphSendReminderRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation in the Azure AD access reviews feature, send a reminder to the reviewers of a currently active accessReview.  The target object can be either a one-time access review, or an instance of a recurring access review. 
func (m *ItemInstancesItemMicrosoftGraphSendReminderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemInstancesItemMicrosoftGraphSendReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
