package calendarsharingmessage

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0894db477956d58d8df90f6c1a340da96b19893b03af59d696aae5f292a6451b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/messages/item/calendarsharingmessage/accept"
)

type CalendarSharingMessageRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *CalendarSharingMessageRequestBuilder) Accept()(*i0894db477956d58d8df90f6c1a340da96b19893b03af59d696aae5f292a6451b.AcceptRequestBuilder) {
    return i0894db477956d58d8df90f6c1a340da96b19893b03af59d696aae5f292a6451b.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewCalendarSharingMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    m := &CalendarSharingMessageRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/messages/{message_id}/microsoft.graph.calendarSharingMessage";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewCalendarSharingMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarSharingMessageRequestBuilderInternal(urlParams, requestAdapter)
}
