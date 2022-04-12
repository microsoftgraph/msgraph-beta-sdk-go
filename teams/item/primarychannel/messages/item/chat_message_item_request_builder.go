package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0576c0f39898532fc04e6824bf28aa7cff8ea81dbbab73b73eefcea817d16a7e "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/messages/item/softdelete"
    i2c260db421138b4b9d16bc781872e9b2bcc7212fcb3df78fccc27b29fa3afdd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/messages/item/replies"
    i58f2982db5184e602b87372e05e45b924a186cb16c2823dc58e2e0d4aefc9433 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/messages/item/hostedcontents"
    id0944412a69a3f355253d839fae1747dfa15fc71c4f32c0625c32ab3087cb973 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/messages/item/undosoftdelete"
    i3bf93b00c9d8a64f21168607b55a5b671b8629d16d677ef881f8c67fc3d08c86 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/messages/item/hostedcontents/item"
    i5b233e37d8baee4d7a51632cc955aa8905e1db9922a952b63423775c3a43570a "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/primarychannel/messages/item/replies/item"
)

// ChatMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.channel entity.
type ChatMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatMessageItemRequestBuilderDeleteOptions options for Delete
type ChatMessageItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChatMessageItemRequestBuilderGetOptions options for Get
type ChatMessageItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChatMessageItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChatMessageItemRequestBuilderGetQueryParameters a collection of all the messages in the channel. A navigation property. Nullable.
type ChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChatMessageItemRequestBuilderPatchOptions options for Patch
type ChatMessageItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewChatMessageItemRequestBuilderInternal instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatMessageItemRequestBuilder) {
    m := &ChatMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}/primaryChannel/messages/{chatMessage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatMessageItemRequestBuilder instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewChatMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property messages for teams
func (m *ChatMessageItemRequestBuilder) CreateDeleteRequestInformation(options *ChatMessageItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a collection of all the messages in the channel. A navigation property. Nullable.
func (m *ChatMessageItemRequestBuilder) CreateGetRequestInformation(options *ChatMessageItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property messages in teams
func (m *ChatMessageItemRequestBuilder) CreatePatchRequestInformation(options *ChatMessageItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property messages for teams
func (m *ChatMessageItemRequestBuilder) Delete(options *ChatMessageItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of all the messages in the channel. A navigation property. Nullable.
func (m *ChatMessageItemRequestBuilder) Get(options *ChatMessageItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChatMessageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChatMessageable), nil
}
// HostedContents the hostedContents property
func (m *ChatMessageItemRequestBuilder) HostedContents()(*i58f2982db5184e602b87372e05e45b924a186cb16c2823dc58e2e0d4aefc9433.HostedContentsRequestBuilder) {
    return i58f2982db5184e602b87372e05e45b924a186cb16c2823dc58e2e0d4aefc9433.NewHostedContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostedContentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.primaryChannel.messages.item.hostedContents.item collection
func (m *ChatMessageItemRequestBuilder) HostedContentsById(id string)(*i3bf93b00c9d8a64f21168607b55a5b671b8629d16d677ef881f8c67fc3d08c86.ChatMessageHostedContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessageHostedContent%2Did"] = id
    }
    return i3bf93b00c9d8a64f21168607b55a5b671b8629d16d677ef881f8c67fc3d08c86.NewChatMessageHostedContentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in teams
func (m *ChatMessageItemRequestBuilder) Patch(options *ChatMessageItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Replies the replies property
func (m *ChatMessageItemRequestBuilder) Replies()(*i2c260db421138b4b9d16bc781872e9b2bcc7212fcb3df78fccc27b29fa3afdd0.RepliesRequestBuilder) {
    return i2c260db421138b4b9d16bc781872e9b2bcc7212fcb3df78fccc27b29fa3afdd0.NewRepliesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RepliesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.teams.item.primaryChannel.messages.item.replies.item collection
func (m *ChatMessageItemRequestBuilder) RepliesById(id string)(*i5b233e37d8baee4d7a51632cc955aa8905e1db9922a952b63423775c3a43570a.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did1"] = id
    }
    return i5b233e37d8baee4d7a51632cc955aa8905e1db9922a952b63423775c3a43570a.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SoftDelete the softDelete property
func (m *ChatMessageItemRequestBuilder) SoftDelete()(*i0576c0f39898532fc04e6824bf28aa7cff8ea81dbbab73b73eefcea817d16a7e.SoftDeleteRequestBuilder) {
    return i0576c0f39898532fc04e6824bf28aa7cff8ea81dbbab73b73eefcea817d16a7e.NewSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UndoSoftDelete the undoSoftDelete property
func (m *ChatMessageItemRequestBuilder) UndoSoftDelete()(*id0944412a69a3f355253d839fae1747dfa15fc71c4f32c0625c32ab3087cb973.UndoSoftDeleteRequestBuilder) {
    return id0944412a69a3f355253d839fae1747dfa15fc71c4f32c0625c32ab3087cb973.NewUndoSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
