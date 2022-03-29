package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i43ca2d478fd8d0c13a1e40866838b59480e161915793db1bc9b8f186fafe966b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/messages/item/hostedcontents"
    i4cc99e3bb3eca7b191b592fb411b58b665ec62f59f863310967fddcda939f753 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/messages/item/softdelete"
    i5d9f63d54fd491adf4d5dcc502464707c1447d44452ece8a476719fed0e7d4e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/messages/item/undosoftdelete"
    i83b530ab58fd2e86db39073f225e00e20c06a8b1febf4dc19211b1d5e8f13f01 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/messages/item/replies"
    i4f8a42a2a0b2ae782a5a5dccdca1b7a4b5cda360e48e4d2a7fc38be40a22e1df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/messages/item/replies/item"
    i64eda43b61ebd8abf184713d19fb8592f07da0b8d76264c2a08b2b2a773a9bb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/chats/item/messages/item/hostedcontents/item"
)

// ChatMessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.chat entity.
type ChatMessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChatMessageItemRequestBuilderDeleteOptions options for Delete
type ChatMessageItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatMessageItemRequestBuilderGetOptions options for Get
type ChatMessageItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChatMessageItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatMessageItemRequestBuilderGetQueryParameters a collection of all the messages in the chat. Nullable.
type ChatMessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ChatMessageItemRequestBuilderPatchOptions options for Patch
type ChatMessageItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ChatMessageable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewChatMessageItemRequestBuilderInternal instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewChatMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatMessageItemRequestBuilder) {
    m := &ChatMessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats/{chat_id}/messages/{chatMessage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatMessageItemRequestBuilder instantiates a new ChatMessageItemRequestBuilder and sets the default values.
func NewChatMessageItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatMessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property messages for me
func (m *ChatMessageItemRequestBuilder) CreateDeleteRequestInformation(options *ChatMessageItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a collection of all the messages in the chat. Nullable.
func (m *ChatMessageItemRequestBuilder) CreateGetRequestInformation(options *ChatMessageItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property messages in me
func (m *ChatMessageItemRequestBuilder) CreatePatchRequestInformation(options *ChatMessageItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property messages for me
func (m *ChatMessageItemRequestBuilder) Delete(options *ChatMessageItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of all the messages in the chat. Nullable.
func (m *ChatMessageItemRequestBuilder) Get(options *ChatMessageItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ChatMessageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateChatMessageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ChatMessageable), nil
}
func (m *ChatMessageItemRequestBuilder) HostedContents()(*i43ca2d478fd8d0c13a1e40866838b59480e161915793db1bc9b8f186fafe966b.HostedContentsRequestBuilder) {
    return i43ca2d478fd8d0c13a1e40866838b59480e161915793db1bc9b8f186fafe966b.NewHostedContentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostedContentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.messages.item.hostedContents.item collection
func (m *ChatMessageItemRequestBuilder) HostedContentsById(id string)(*i64eda43b61ebd8abf184713d19fb8592f07da0b8d76264c2a08b2b2a773a9bb7.ChatMessageHostedContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessageHostedContent_id"] = id
    }
    return i64eda43b61ebd8abf184713d19fb8592f07da0b8d76264c2a08b2b2a773a9bb7.NewChatMessageHostedContentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in me
func (m *ChatMessageItemRequestBuilder) Patch(options *ChatMessageItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ChatMessageItemRequestBuilder) Replies()(*i83b530ab58fd2e86db39073f225e00e20c06a8b1febf4dc19211b1d5e8f13f01.RepliesRequestBuilder) {
    return i83b530ab58fd2e86db39073f225e00e20c06a8b1febf4dc19211b1d5e8f13f01.NewRepliesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RepliesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.chats.item.messages.item.replies.item collection
func (m *ChatMessageItemRequestBuilder) RepliesById(id string)(*i4f8a42a2a0b2ae782a5a5dccdca1b7a4b5cda360e48e4d2a7fc38be40a22e1df.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id1"] = id
    }
    return i4f8a42a2a0b2ae782a5a5dccdca1b7a4b5cda360e48e4d2a7fc38be40a22e1df.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatMessageItemRequestBuilder) SoftDelete()(*i4cc99e3bb3eca7b191b592fb411b58b665ec62f59f863310967fddcda939f753.SoftDeleteRequestBuilder) {
    return i4cc99e3bb3eca7b191b592fb411b58b665ec62f59f863310967fddcda939f753.NewSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatMessageItemRequestBuilder) UndoSoftDelete()(*i5d9f63d54fd491adf4d5dcc502464707c1447d44452ece8a476719fed0e7d4e5.UndoSoftDeleteRequestBuilder) {
    return i5d9f63d54fd491adf4d5dcc502464707c1447d44452ece8a476719fed0e7d4e5.NewUndoSoftDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
