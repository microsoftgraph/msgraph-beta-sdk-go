package eventmessagerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i920522cb06e0a16d53c9143591ce48d22f59cacbef7c8a4ee3bacebd9398be2c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/eventmessagerequest/accept"
    idcc893039853b4c11de778ee04afbd751030092eb0d700d765da9c395b7222b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/eventmessagerequest/tentativelyaccept"
    ifd85e2e63f84f55a2706d537fcb113ef2ea17d054c46e82c720853d6df713e74 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/eventmessagerequest/decline"
)

// EventMessageRequestRequestBuilder builds and executes requests for operations under \me\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages\{message-id}\microsoft.graph.eventMessageRequest
type EventMessageRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *EventMessageRequestRequestBuilder) Accept()(*i920522cb06e0a16d53c9143591ce48d22f59cacbef7c8a4ee3bacebd9398be2c.AcceptRequestBuilder) {
    return i920522cb06e0a16d53c9143591ce48d22f59cacbef7c8a4ee3bacebd9398be2c.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventMessageRequestRequestBuilderInternal instantiates a new EventMessageRequestRequestBuilder and sets the default values.
func NewEventMessageRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    m := &EventMessageRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder_id}/childFolders/{mailFolder_id1}/messages/{message_id}/microsoft.graph.eventMessageRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventMessageRequestRequestBuilder instantiates a new EventMessageRequestRequestBuilder and sets the default values.
func NewEventMessageRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventMessageRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EventMessageRequestRequestBuilder) Decline()(*ifd85e2e63f84f55a2706d537fcb113ef2ea17d054c46e82c720853d6df713e74.DeclineRequestBuilder) {
    return ifd85e2e63f84f55a2706d537fcb113ef2ea17d054c46e82c720853d6df713e74.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventMessageRequestRequestBuilder) TentativelyAccept()(*idcc893039853b4c11de778ee04afbd751030092eb0d700d765da9c395b7222b9.TentativelyAcceptRequestBuilder) {
    return idcc893039853b4c11de778ee04afbd751030092eb0d700d765da9c395b7222b9.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
