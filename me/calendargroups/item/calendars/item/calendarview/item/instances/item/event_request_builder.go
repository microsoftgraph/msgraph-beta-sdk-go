package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i47276488d225dcc7e3f72798310bc6ca89481b6777b407d0bb37c20be25b000f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/tentativelyaccept"
    i688c7d91e4531849f710689d0c250faa616a1a4cdab0254425a4f7a6211a2b38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/snoozereminder"
    i86ed9ef973e47ab7658aac2c1c0a852789334dc7816d55cb3d8e8942d2b7a0fb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/decline"
    i96a181755a49c2932e63bfe6a5aab36214e41f47d5bc94c50c84370fbf3a9a0d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/cancel"
    ia4bfb9757654f9c5b3007c22b7a788e359bff63d25987b7ece8a6012cd7057b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/accept"
    ibc1c823b12bc0be47f58db00c06ca1883de39606def2526e225da84305fffe47 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/forward"
    ieaa4264c9c7f4a01a1387aba7de5a990f329172121dc139778f4a1acf4809626 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item/dismissreminder"
)

// EventRequestBuilder builds and executes requests for operations under \me\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\calendarView\{event-id}\instances\{event-id1}
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
// EventRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventRequestBuilderGetQueryParameters struct {
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
func (m *EventRequestBuilder) Accept()(*ia4bfb9757654f9c5b3007c22b7a788e359bff63d25987b7ece8a6012cd7057b7.AcceptRequestBuilder) {
    return ia4bfb9757654f9c5b3007c22b7a788e359bff63d25987b7ece8a6012cd7057b7.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i96a181755a49c2932e63bfe6a5aab36214e41f47d5bc94c50c84370fbf3a9a0d.CancelRequestBuilder) {
    return i96a181755a49c2932e63bfe6a5aab36214e41f47d5bc94c50c84370fbf3a9a0d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/calendarView/{event_id}/instances/{event_id1}{?select}";
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
// CreateDeleteRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
// CreatePatchRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventRequestBuilder) Decline()(*i86ed9ef973e47ab7658aac2c1c0a852789334dc7816d55cb3d8e8942d2b7a0fb.DeclineRequestBuilder) {
    return i86ed9ef973e47ab7658aac2c1c0a852789334dc7816d55cb3d8e8942d2b7a0fb.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventRequestBuilder) DismissReminder()(*ieaa4264c9c7f4a01a1387aba7de5a990f329172121dc139778f4a1acf4809626.DismissReminderRequestBuilder) {
    return ieaa4264c9c7f4a01a1387aba7de5a990f329172121dc139778f4a1acf4809626.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ibc1c823b12bc0be47f58db00c06ca1883de39606def2526e225da84305fffe47.ForwardRequestBuilder) {
    return ibc1c823b12bc0be47f58db00c06ca1883de39606def2526e225da84305fffe47.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
// Patch the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventRequestBuilder) SnoozeReminder()(*i688c7d91e4531849f710689d0c250faa616a1a4cdab0254425a4f7a6211a2b38.SnoozeReminderRequestBuilder) {
    return i688c7d91e4531849f710689d0c250faa616a1a4cdab0254425a4f7a6211a2b38.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i47276488d225dcc7e3f72798310bc6ca89481b6777b407d0bb37c20be25b000f.TentativelyAcceptRequestBuilder) {
    return i47276488d225dcc7e3f72798310bc6ca89481b6777b407d0bb37c20be25b000f.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
