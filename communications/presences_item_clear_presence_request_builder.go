package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PresencesItemClearPresenceRequestBuilder provides operations to call the clearPresence method.
type PresencesItemClearPresenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PresencesItemClearPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresencesItemClearPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPresencesItemClearPresenceRequestBuilderInternal instantiates a new ClearPresenceRequestBuilder and sets the default values.
func NewPresencesItemClearPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemClearPresenceRequestBuilder) {
    m := &PresencesItemClearPresenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/presences/{presence%2Did}/clearPresence", pathParameters),
    }
    return m
}
// NewPresencesItemClearPresenceRequestBuilder instantiates a new ClearPresenceRequestBuilder and sets the default values.
func NewPresencesItemClearPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemClearPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresencesItemClearPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear a presence session of an application for a user. If it is the user's only presence session, a successful clearPresence changes the user's presence to Offline/Offline. Read more about presence sessions and their time-out and expiration.  This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/presence-clearpresence?view=graph-rest-1.0
func (m *PresencesItemClearPresenceRequestBuilder) Post(ctx context.Context, body PresencesItemClearPresencePostRequestBodyable, requestConfiguration *PresencesItemClearPresenceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation clear a presence session of an application for a user. If it is the user's only presence session, a successful clearPresence changes the user's presence to Offline/Offline. Read more about presence sessions and their time-out and expiration.  This API is available in the following national cloud deployments.
func (m *PresencesItemClearPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, body PresencesItemClearPresencePostRequestBodyable, requestConfiguration *PresencesItemClearPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *PresencesItemClearPresenceRequestBuilder) WithUrl(rawUrl string)(*PresencesItemClearPresenceRequestBuilder) {
    return NewPresencesItemClearPresenceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
