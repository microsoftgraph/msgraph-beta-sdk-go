package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5593685ccf7132f2b322d72c7f6190a49e272ba79796551cb489be2b74af524f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/cancel"
    i57332911775b2c8d06d7a5ec6a0181b964d10ead1a8ad0594e53d7bf944092a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/tentativelyaccept"
    i7f72755c435a384033bee52a976326640bd6d1b3fe52ec1d522d1718f3d6acd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/snoozereminder"
    i8115d40777816dc52915a104ff18869f165cd32f3ae8ecb2740a6a9cbdf2af2c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/forward"
    ia16eaac4103fc4b7aff74f7c928df3fa66c1acc3bd5976fc3c93c29176030028 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/dismissreminder"
    ia9bf548a1db34e1a13b86e5b5f8f7297f8c5c9cdf08b00efaf8f47edc57c530d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/decline"
    id92bca78a768b46b0a9ae7f9bb4158b0042091ffb95f227d3e6174994267a64b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/accept"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\events\{event-id}\exceptionOccurrences\{event-id1}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventRequestBuilderDeleteOptions options for Delete
type EventRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventRequestBuilderGetOptions options for Get
type EventRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventRequestBuilderGetQueryParameters get exceptionOccurrences from users
type EventRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EventRequestBuilderPatchOptions options for Patch
type EventRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventRequestBuilder) Accept()(*id92bca78a768b46b0a9ae7f9bb4158b0042091ffb95f227d3e6174994267a64b.AcceptRequestBuilder) {
    return id92bca78a768b46b0a9ae7f9bb4158b0042091ffb95f227d3e6174994267a64b.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i5593685ccf7132f2b322d72c7f6190a49e272ba79796551cb489be2b74af524f.CancelRequestBuilder) {
    return i5593685ccf7132f2b322d72c7f6190a49e272ba79796551cb489be2b74af524f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/events/{event_id}/exceptionOccurrences/{event_id1}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventRequestBuilder instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for users
func (m *EventRequestBuilder) CreateDeleteRequestInformation(options *EventRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get exceptionOccurrences from users
func (m *EventRequestBuilder) CreateGetRequestInformation(options *EventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in users
func (m *EventRequestBuilder) CreatePatchRequestInformation(options *EventRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) Decline()(*ia9bf548a1db34e1a13b86e5b5f8f7297f8c5c9cdf08b00efaf8f47edc57c530d.DeclineRequestBuilder) {
    return ia9bf548a1db34e1a13b86e5b5f8f7297f8c5c9cdf08b00efaf8f47edc57c530d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for users
func (m *EventRequestBuilder) Delete(options *EventRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) DismissReminder()(*ia16eaac4103fc4b7aff74f7c928df3fa66c1acc3bd5976fc3c93c29176030028.DismissReminderRequestBuilder) {
    return ia16eaac4103fc4b7aff74f7c928df3fa66c1acc3bd5976fc3c93c29176030028.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i8115d40777816dc52915a104ff18869f165cd32f3ae8ecb2740a6a9cbdf2af2c.ForwardRequestBuilder) {
    return i8115d40777816dc52915a104ff18869f165cd32f3ae8ecb2740a6a9cbdf2af2c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
func (m *EventRequestBuilder) Get(options *EventRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvent() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event), nil
}
// Patch update the navigation property exceptionOccurrences in users
func (m *EventRequestBuilder) Patch(options *EventRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) SnoozeReminder()(*i7f72755c435a384033bee52a976326640bd6d1b3fe52ec1d522d1718f3d6acd5.SnoozeReminderRequestBuilder) {
    return i7f72755c435a384033bee52a976326640bd6d1b3fe52ec1d522d1718f3d6acd5.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i57332911775b2c8d06d7a5ec6a0181b964d10ead1a8ad0594e53d7bf944092a7.TentativelyAcceptRequestBuilder) {
    return i57332911775b2c8d06d7a5ec6a0181b964d10ead1a8ad0594e53d7bf944092a7.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
