package eventmessagerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i82d5228b240b684460e802c686fbcd142766754c5e906e6b306f048773f26024 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/eventmessagerequest/tentativelyaccept"
    i966adf25d5a9aca2e07c5266c5bd2af4d940fed67c462c966f492d0707603884 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/eventmessagerequest/decline"
    if72a00e0c7ee025d9ed69cf1124fa2a3d07a64f8117214c488334090dd8bf6e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/eventmessagerequest/accept"
)

type EventMessageRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *EventMessageRequestRequestBuilder) Accept()(*if72a00e0c7ee025d9ed69cf1124fa2a3d07a64f8117214c488334090dd8bf6e2.AcceptRequestBuilder) {
    return if72a00e0c7ee025d9ed69cf1124fa2a3d07a64f8117214c488334090dd8bf6e2.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventMessageRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    m := &EventMessageRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/messages/{message_id}/microsoft.graph.eventMessageRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewEventMessageRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventMessageRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EventMessageRequestRequestBuilder) Decline()(*i966adf25d5a9aca2e07c5266c5bd2af4d940fed67c462c966f492d0707603884.DeclineRequestBuilder) {
    return i966adf25d5a9aca2e07c5266c5bd2af4d940fed67c462c966f492d0707603884.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventMessageRequestRequestBuilder) TentativelyAccept()(*i82d5228b240b684460e802c686fbcd142766754c5e906e6b306f048773f26024.TentativelyAcceptRequestBuilder) {
    return i82d5228b240b684460e802c686fbcd142766754c5e906e6b306f048773f26024.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
