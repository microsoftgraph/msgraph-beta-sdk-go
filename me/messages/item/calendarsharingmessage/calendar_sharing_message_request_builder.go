package calendarsharingmessage

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i181e9cf4130b18c1f40044e856a6b727ecd1daa908cb274d54825301360d2efe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/calendarsharingmessage/accept"
)

// CalendarSharingMessageRequestBuilder builds and executes requests for operations under \me\messages\{message-id}\microsoft.graph.calendarSharingMessage
type CalendarSharingMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
func (m *CalendarSharingMessageRequestBuilder) Accept()(*i181e9cf4130b18c1f40044e856a6b727ecd1daa908cb274d54825301360d2efe.AcceptRequestBuilder) {
    return i181e9cf4130b18c1f40044e856a6b727ecd1daa908cb274d54825301360d2efe.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCalendarSharingMessageRequestBuilderInternal instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
func NewCalendarSharingMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    m := &CalendarSharingMessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/messages/{message_id}/microsoft.graph.calendarSharingMessage";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarSharingMessageRequestBuilder instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
func NewCalendarSharingMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarSharingMessageRequestBuilderInternal(urlParams, requestAdapter)
}
