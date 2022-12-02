package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilder provides operations to call the replyAll method.
type UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemMailFoldersItemMessagesItemReplyAllRequestBuilderInternal instantiates a new ReplyAllRequestBuilder and sets the default values.
func NewUsersItemMailFoldersItemMessagesItemReplyAllRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilder) {
    m := &UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/microsoft.graph.replyAll";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemMailFoldersItemMessagesItemReplyAllRequestBuilder instantiates a new ReplyAllRequestBuilder and sets the default values.
func NewUsersItemMailFoldersItemMessagesItemReplyAllRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemMailFoldersItemMessagesItemReplyAllRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation reply to all recipients of a message using either JSON or MIME format. When using JSON format:- Specify either a comment or the **body** property of the `message` parameter. Specifying both will return an HTTP 400 Bad Request error.- If the original message specifies a recipient in the **replyTo** property, per Internet Message Format (RFC 2822), send the reply to the recipients in **replyTo** and not the recipient in the **from** property. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- Add any attachments and S/MIME properties to the MIME content. This method saves the message in the **Sent Items** folder. Alternatively, create a draft to reply-all to a message, and send it later.
func (m *UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemMailFoldersItemMessagesItemReplyAllPostRequestBodyable, requestConfiguration *UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post reply to all recipients of a message using either JSON or MIME format. When using JSON format:- Specify either a comment or the **body** property of the `message` parameter. Specifying both will return an HTTP 400 Bad Request error.- If the original message specifies a recipient in the **replyTo** property, per Internet Message Format (RFC 2822), send the reply to the recipients in **replyTo** and not the recipient in the **from** property. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in **base64** format in the request body.- Add any attachments and S/MIME properties to the MIME content. This method saves the message in the **Sent Items** folder. Alternatively, create a draft to reply-all to a message, and send it later.
func (m *UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemMailFoldersItemMessagesItemReplyAllPostRequestBodyable, requestConfiguration *UsersItemMailFoldersItemMessagesItemReplyAllRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
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
