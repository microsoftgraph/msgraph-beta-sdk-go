package eventmessagerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4d3541020132803aaabd3cb95aebaf982a6f832643bbeab622ac5caea33037c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/eventmessagerequest/accept"
    i91bfb4cde0e46cc0325ed184fbacc00907512f081a333138933fb1b956c24ace "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/eventmessagerequest/tentativelyaccept"
    ic549cd957a762c9b708cf3f98e3df1f9ff42c54339a2f5975e4ece29f7338f37 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/eventmessagerequest/decline"
)

// Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}\microsoft.graph.eventMessageRequest
type EventMessageRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *EventMessageRequestRequestBuilder) Accept()(*i4d3541020132803aaabd3cb95aebaf982a6f832643bbeab622ac5caea33037c0.AcceptRequestBuilder) {
    return i4d3541020132803aaabd3cb95aebaf982a6f832643bbeab622ac5caea33037c0.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventMessageRequestRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventMessageRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    m := &EventMessageRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/mailFolders/{mailFolder_id}/messages/{message_id}/microsoft.graph.eventMessageRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new EventMessageRequestRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventMessageRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventMessageRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EventMessageRequestRequestBuilder) Decline()(*ic549cd957a762c9b708cf3f98e3df1f9ff42c54339a2f5975e4ece29f7338f37.DeclineRequestBuilder) {
    return ic549cd957a762c9b708cf3f98e3df1f9ff42c54339a2f5975e4ece29f7338f37.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventMessageRequestRequestBuilder) TentativelyAccept()(*i91bfb4cde0e46cc0325ed184fbacc00907512f081a333138933fb1b956c24ace.TentativelyAcceptRequestBuilder) {
    return i91bfb4cde0e46cc0325ed184fbacc00907512f081a333138933fb1b956c24ace.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
