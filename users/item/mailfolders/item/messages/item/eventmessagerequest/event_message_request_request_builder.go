package eventmessagerequest

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4d3541020132803aaabd3cb95aebaf982a6f832643bbeab622ac5caea33037c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/eventmessagerequest/accept"
    i91bfb4cde0e46cc0325ed184fbacc00907512f081a333138933fb1b956c24ace "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/eventmessagerequest/tentativelyaccept"
    ic549cd957a762c9b708cf3f98e3df1f9ff42c54339a2f5975e4ece29f7338f37 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/eventmessagerequest/decline"
)

// EventMessageRequestRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}\microsoft.graph.eventMessageRequest
type EventMessageRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Accept the accept property
func (m *EventMessageRequestRequestBuilder) Accept()(*i4d3541020132803aaabd3cb95aebaf982a6f832643bbeab622ac5caea33037c0.AcceptRequestBuilder) {
    return i4d3541020132803aaabd3cb95aebaf982a6f832643bbeab622ac5caea33037c0.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventMessageRequestRequestBuilderInternal instantiates a new EventMessageRequestRequestBuilder and sets the default values.
func NewEventMessageRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    m := &EventMessageRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/microsoft.graph.eventMessageRequest";
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
func (m *EventMessageRequestRequestBuilder) Decline()(*ic549cd957a762c9b708cf3f98e3df1f9ff42c54339a2f5975e4ece29f7338f37.DeclineRequestBuilder) {
    return ic549cd957a762c9b708cf3f98e3df1f9ff42c54339a2f5975e4ece29f7338f37.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventMessageRequestRequestBuilder) TentativelyAccept()(*i91bfb4cde0e46cc0325ed184fbacc00907512f081a333138933fb1b956c24ace.TentativelyAcceptRequestBuilder) {
    return i91bfb4cde0e46cc0325ed184fbacc00907512f081a333138933fb1b956c24ace.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
