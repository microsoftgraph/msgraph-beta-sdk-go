package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemPendingAccessReviewInstancesItemStopRequestBuilder provides operations to call the stop method.
type UsersItemPendingAccessReviewInstancesItemStopRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemPendingAccessReviewInstancesItemStopRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPendingAccessReviewInstancesItemStopRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemPendingAccessReviewInstancesItemStopRequestBuilderInternal instantiates a new StopRequestBuilder and sets the default values.
func NewUsersItemPendingAccessReviewInstancesItemStopRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPendingAccessReviewInstancesItemStopRequestBuilder) {
    m := &UsersItemPendingAccessReviewInstancesItemStopRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/pendingAccessReviewInstances/{accessReviewInstance%2Did}/microsoft.graph.stop";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemPendingAccessReviewInstancesItemStopRequestBuilder instantiates a new StopRequestBuilder and sets the default values.
func NewUsersItemPendingAccessReviewInstancesItemStopRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPendingAccessReviewInstancesItemStopRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemPendingAccessReviewInstancesItemStopRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation stop a currently active accessReviewInstance. After the access review instance stops, the instance status will be `Completed`, the reviewers can no longer give input, and the access review decisions can be applied. Stopping an instance will not effect future instances. To prevent a recurring access review from starting future instances, update the schedule definition to change its scheduled end date.
func (m *UsersItemPendingAccessReviewInstancesItemStopRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemStopRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post stop a currently active accessReviewInstance. After the access review instance stops, the instance status will be `Completed`, the reviewers can no longer give input, and the access review decisions can be applied. Stopping an instance will not effect future instances. To prevent a recurring access review from starting future instances, update the schedule definition to change its scheduled end date.
func (m *UsersItemPendingAccessReviewInstancesItemStopRequestBuilder) Post(ctx context.Context, requestConfiguration *UsersItemPendingAccessReviewInstancesItemStopRequestBuilderPostRequestConfiguration)(error) {
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
