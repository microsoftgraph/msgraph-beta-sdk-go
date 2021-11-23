package eventmessagerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i19b46588b0adc3999ff320dfdb7d228cf5b26518ff6428a8e63a968e97cf01b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/eventmessagerequest/tentativelyaccept"
    i9c76f5ea50d606fd3274a3b0a443c980ce7301ab87ae564a7ca0d036870493e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/eventmessagerequest/decline"
    ia00434b9a8fa3b48b55c0935acc63a5f19c1dc22afab2c27bf7aaf0042639e0e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/eventmessagerequest/accept"
)

// EventMessageRequestRequestBuilder builds and executes requests for operations under \me\messages\{message-id}\microsoft.graph.eventMessageRequest
type EventMessageRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *EventMessageRequestRequestBuilder) Accept()(*ia00434b9a8fa3b48b55c0935acc63a5f19c1dc22afab2c27bf7aaf0042639e0e.AcceptRequestBuilder) {
    return ia00434b9a8fa3b48b55c0935acc63a5f19c1dc22afab2c27bf7aaf0042639e0e.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventMessageRequestRequestBuilderInternal instantiates a new EventMessageRequestRequestBuilder and sets the default values.
func NewEventMessageRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    m := &EventMessageRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/messages/{message_id}/microsoft.graph.eventMessageRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventMessageRequestRequestBuilder instantiates a new EventMessageRequestRequestBuilder and sets the default values.
func NewEventMessageRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventMessageRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EventMessageRequestRequestBuilder) Decline()(*i9c76f5ea50d606fd3274a3b0a443c980ce7301ab87ae564a7ca0d036870493e4.DeclineRequestBuilder) {
    return i9c76f5ea50d606fd3274a3b0a443c980ce7301ab87ae564a7ca0d036870493e4.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventMessageRequestRequestBuilder) TentativelyAccept()(*i19b46588b0adc3999ff320dfdb7d228cf5b26518ff6428a8e63a968e97cf01b9.TentativelyAcceptRequestBuilder) {
    return i19b46588b0adc3999ff320dfdb7d228cf5b26518ff6428a8e63a968e97cf01b9.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
