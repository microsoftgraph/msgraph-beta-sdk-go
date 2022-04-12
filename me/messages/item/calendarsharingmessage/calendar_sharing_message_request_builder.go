package calendarsharingmessage

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i181e9cf4130b18c1f40044e856a6b727ecd1daa908cb274d54825301360d2efe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/messages/item/calendarsharingmessage/accept"
)

// CalendarSharingMessageRequestBuilder builds and executes requests for operations under \me\messages\{message-id}\microsoft.graph.calendarSharingMessage
type CalendarSharingMessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Accept the accept property
func (m *CalendarSharingMessageRequestBuilder) Accept()(*i181e9cf4130b18c1f40044e856a6b727ecd1daa908cb274d54825301360d2efe.AcceptRequestBuilder) {
    return i181e9cf4130b18c1f40044e856a6b727ecd1daa908cb274d54825301360d2efe.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCalendarSharingMessageRequestBuilderInternal instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
func NewCalendarSharingMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    m := &CalendarSharingMessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/messages/{message%2Did}/microsoft.graph.calendarSharingMessage";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarSharingMessageRequestBuilder instantiates a new CalendarSharingMessageRequestBuilder and sets the default values.
func NewCalendarSharingMessageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarSharingMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarSharingMessageRequestBuilderInternal(urlParams, requestAdapter)
}
