package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder provides operations to call the markAsJunk method.
type ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilderInternal instantiates a new ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder and sets the default values.
func NewItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder) {
    m := &ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/markAsJunk", pathParameters),
    }
    return m
}
// NewItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder instantiates a new ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder and sets the default values.
func NewItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mark a message as junk. This API adds the sender to the list of blocked senders and moves the message to the Junk Email folder, when moveToJunk is true.
// returns a Messageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/message-markasjunk?view=graph-rest-beta
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder) Post(ctx context.Context, body ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkPostRequestBodyable, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable), nil
}
// ToPostRequestInformation mark a message as junk. This API adds the sender to the list of blocked senders and moves the message to the Junk Email folder, when moveToJunk is true.
// returns a *RequestInformation when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkPostRequestBodyable, requestConfiguration *ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder when successful
func (m *ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder) WithUrl(rawUrl string)(*ItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder) {
    return NewItemMailfoldersItemChildfoldersItemMessagesItemMarkasjunkMarkAsJunkRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
