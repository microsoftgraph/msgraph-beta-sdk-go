package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MessagesItemUnsubscribeRequestBuilder provides operations to call the unsubscribe method.
type MessagesItemUnsubscribeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MessagesItemUnsubscribeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessagesItemUnsubscribeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMessagesItemUnsubscribeRequestBuilderInternal instantiates a new UnsubscribeRequestBuilder and sets the default values.
func NewMessagesItemUnsubscribeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessagesItemUnsubscribeRequestBuilder) {
    m := &MessagesItemUnsubscribeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/messages/{message%2Did}/microsoft.graph.unsubscribe";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessagesItemUnsubscribeRequestBuilder instantiates a new UnsubscribeRequestBuilder and sets the default values.
func NewMessagesItemUnsubscribeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessagesItemUnsubscribeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessagesItemUnsubscribeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post submits a email request on behalf of the signed-in user to unsubscribe from an email distribution list. Uses the information in the `List-Unsubscribe` header. Message senders can use mailing lists in a user-friendly way by including an option for recipients to opt out. They can do so by specifying the `List-Unsubscribe` header in each message following RFC-2369. **Note** In particular, for the **unsubscribe** action to work, the sender must specify `mailto:` and not URL-based unsubscribe information. Setting that header would also set the **unsubscribeEnabled** property of the message instance to `true`, and the **unsubscribeData** property to the header data. If the **unsubscribeEnabled** property of a message is `true`, you can use the **unsubscribe** action to unsubscribe the user from similar future messages as managed by the message sender. A successful **unsubscribe** action moves the message to the **Deleted Items** folder. The actual exclusion of the user from future mail distribution is managed by the sender.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/message-unsubscribe?view=graph-rest-1.0
func (m *MessagesItemUnsubscribeRequestBuilder) Post(ctx context.Context, requestConfiguration *MessagesItemUnsubscribeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation submits a email request on behalf of the signed-in user to unsubscribe from an email distribution list. Uses the information in the `List-Unsubscribe` header. Message senders can use mailing lists in a user-friendly way by including an option for recipients to opt out. They can do so by specifying the `List-Unsubscribe` header in each message following RFC-2369. **Note** In particular, for the **unsubscribe** action to work, the sender must specify `mailto:` and not URL-based unsubscribe information. Setting that header would also set the **unsubscribeEnabled** property of the message instance to `true`, and the **unsubscribeData** property to the header data. If the **unsubscribeEnabled** property of a message is `true`, you can use the **unsubscribe** action to unsubscribe the user from similar future messages as managed by the message sender. A successful **unsubscribe** action moves the message to the **Deleted Items** folder. The actual exclusion of the user from future mail distribution is managed by the sender.
func (m *MessagesItemUnsubscribeRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MessagesItemUnsubscribeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
