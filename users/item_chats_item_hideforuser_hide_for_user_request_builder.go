package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemChatsItemHideforuserHideForUserRequestBuilder provides operations to call the hideForUser method.
type ItemChatsItemHideforuserHideForUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemHideforuserHideForUserRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemHideforuserHideForUserRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemHideforuserHideForUserRequestBuilderInternal instantiates a new ItemChatsItemHideforuserHideForUserRequestBuilder and sets the default values.
func NewItemChatsItemHideforuserHideForUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemHideforuserHideForUserRequestBuilder) {
    m := &ItemChatsItemHideforuserHideForUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/hideForUser", pathParameters),
    }
    return m
}
// NewItemChatsItemHideforuserHideForUserRequestBuilder instantiates a new ItemChatsItemHideforuserHideForUserRequestBuilder and sets the default values.
func NewItemChatsItemHideforuserHideForUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemHideforuserHideForUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemHideforuserHideForUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Post hide a chat for a user.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chat-hideforuser?view=graph-rest-beta
func (m *ItemChatsItemHideforuserHideForUserRequestBuilder) Post(ctx context.Context, body ItemChatsItemHideforuserHideForUserPostRequestBodyable, requestConfiguration *ItemChatsItemHideforuserHideForUserRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation hide a chat for a user.
// returns a *RequestInformation when successful
func (m *ItemChatsItemHideforuserHideForUserRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemChatsItemHideforuserHideForUserPostRequestBodyable, requestConfiguration *ItemChatsItemHideforuserHideForUserRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemChatsItemHideforuserHideForUserRequestBuilder when successful
func (m *ItemChatsItemHideforuserHideForUserRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemHideforuserHideForUserRequestBuilder) {
    return NewItemChatsItemHideforuserHideForUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
