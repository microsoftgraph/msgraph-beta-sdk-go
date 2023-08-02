package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPresenceClearPresenceRequestBuilder provides operations to call the clearPresence method.
type ItemPresenceClearPresenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPresenceClearPresenceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPresenceClearPresenceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPresenceClearPresenceRequestBuilderInternal instantiates a new ClearPresenceRequestBuilder and sets the default values.
func NewItemPresenceClearPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPresenceClearPresenceRequestBuilder) {
    m := &ItemPresenceClearPresenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/presence/clearPresence", pathParameters),
    }
    return m
}
// NewItemPresenceClearPresenceRequestBuilder instantiates a new ClearPresenceRequestBuilder and sets the default values.
func NewItemPresenceClearPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPresenceClearPresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPresenceClearPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear a presence session of an application for a user. If it is the user's only presence session, a successful clearPresence changes the user's presence to Offline/Offline. Read more about presence sessions and their time-out and expiration. 
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/presence-clearpresence?view=graph-rest-1.0
func (m *ItemPresenceClearPresenceRequestBuilder) Post(ctx context.Context, body ItemPresenceClearPresencePostRequestBodyable, requestConfiguration *ItemPresenceClearPresenceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation clear a presence session of an application for a user. If it is the user's only presence session, a successful clearPresence changes the user's presence to Offline/Offline. Read more about presence sessions and their time-out and expiration. 
func (m *ItemPresenceClearPresenceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPresenceClearPresencePostRequestBodyable, requestConfiguration *ItemPresenceClearPresenceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
