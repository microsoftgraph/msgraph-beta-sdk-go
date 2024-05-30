package chats

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder provides operations to call the removeAllAccessForUser method.
type ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilderInternal instantiates a new ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder and sets the default values.
func NewItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder) {
    m := &ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/chats/{chat%2Did}/removeAllAccessForUser", pathParameters),
    }
    return m
}
// NewItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder instantiates a new ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder and sets the default values.
func NewItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove access to a chat for a user.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chat-removeallaccessforuser?view=graph-rest-beta
func (m *ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder) Post(ctx context.Context, body ItemRemoveallaccessforuserRemoveAllAccessForUserPostRequestBodyable, requestConfiguration *ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation remove access to a chat for a user.
// returns a *RequestInformation when successful
func (m *ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemRemoveallaccessforuserRemoveAllAccessForUserPostRequestBodyable, requestConfiguration *ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder when successful
func (m *ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder) WithUrl(rawUrl string)(*ItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder) {
    return NewItemRemoveallaccessforuserRemoveAllAccessForUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
