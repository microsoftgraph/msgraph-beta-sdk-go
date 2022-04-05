package eventmessagerequest

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0899d26b0b0522b710df806f0454178eaf5c74fca4ec1761f9293365fb10a400 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/eventmessagerequest/accept"
    i1b1fb416c2a097fa7fb6d7f58517668193a1f68c1c86940328d84a96e121c128 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/eventmessagerequest/decline"
    i97f7c5bb3fb3d0e102fdbcc51d3c0673e7144cf63586517eaf7160cf5c7452e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/eventmessagerequest/tentativelyaccept"
)

// EventMessageRequestRequestBuilder builds and executes requests for operations under \me\mailFolders\{mailFolder-id}\messages\{message-id}\microsoft.graph.eventMessageRequest
type EventMessageRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Accept the accept property
func (m *EventMessageRequestRequestBuilder) Accept()(*i0899d26b0b0522b710df806f0454178eaf5c74fca4ec1761f9293365fb10a400.AcceptRequestBuilder) {
    return i0899d26b0b0522b710df806f0454178eaf5c74fca4ec1761f9293365fb10a400.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventMessageRequestRequestBuilderInternal instantiates a new EventMessageRequestRequestBuilder and sets the default values.
func NewEventMessageRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    m := &EventMessageRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder_id}/messages/{message_id}/microsoft.graph.eventMessageRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventMessageRequestRequestBuilder instantiates a new EventMessageRequestRequestBuilder and sets the default values.
func NewEventMessageRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventMessageRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline the decline property
func (m *EventMessageRequestRequestBuilder) Decline()(*i1b1fb416c2a097fa7fb6d7f58517668193a1f68c1c86940328d84a96e121c128.DeclineRequestBuilder) {
    return i1b1fb416c2a097fa7fb6d7f58517668193a1f68c1c86940328d84a96e121c128.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventMessageRequestRequestBuilder) TentativelyAccept()(*i97f7c5bb3fb3d0e102fdbcc51d3c0673e7144cf63586517eaf7160cf5c7452e8.TentativelyAcceptRequestBuilder) {
    return i97f7c5bb3fb3d0e102fdbcc51d3c0673e7144cf63586517eaf7160cf5c7452e8.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
