package accessreviews

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AccessReviewsItemInstancesItemSendReminderRequestBuilder provides operations to call the sendReminder method.
type AccessReviewsItemInstancesItemSendReminderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AccessReviewsItemInstancesItemSendReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessReviewsItemInstancesItemSendReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessReviewsItemInstancesItemSendReminderRequestBuilderInternal instantiates a new SendReminderRequestBuilder and sets the default values.
func NewAccessReviewsItemInstancesItemSendReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsItemInstancesItemSendReminderRequestBuilder) {
    m := &AccessReviewsItemInstancesItemSendReminderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/accessReviews/{accessReview%2Did}/instances/{accessReview%2Did1}/microsoft.graph.sendReminder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAccessReviewsItemInstancesItemSendReminderRequestBuilder instantiates a new SendReminderRequestBuilder and sets the default values.
func NewAccessReviewsItemInstancesItemSendReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessReviewsItemInstancesItemSendReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewsItemInstancesItemSendReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation in the Azure AD access reviews feature, send a reminder to the reviewers of a currently active accessReview.  The target object can be either a one-time access review, or an instance of a recurring access review. 
func (m *AccessReviewsItemInstancesItemSendReminderRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *AccessReviewsItemInstancesItemSendReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post in the Azure AD access reviews feature, send a reminder to the reviewers of a currently active accessReview.  The target object can be either a one-time access review, or an instance of a recurring access review. 
func (m *AccessReviewsItemInstancesItemSendReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *AccessReviewsItemInstancesItemSendReminderRequestBuilderPostRequestConfiguration)(error) {
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
