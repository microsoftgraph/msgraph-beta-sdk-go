package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0afcb2c03d1ee02a759843837718d498b3b7a31bb6a8f8567f8f2ee82e08f6b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/cancel"
    i0cbde926d382023c0322c5a02756b36f5cc3555c3a3e42ccf81f0fe2328d60da "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i9157189a72bc4b79acc0436a007027d935b773485d5138c2005b2ebad608dc94 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/dismissreminder"
    ia546eba81a33eaf6dea2e0a5a4af88df1192817310dd311ca8049116110ef0ba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/snoozereminder"
    iac85e1e6b35961199fed24ddde940883c14819fdd25988d6de0a6de5a9af0adf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/decline"
    ic72278706e4ce7b845b83533a17c670656d7d0710d904d3ef6218f5fa7b3d5cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/accept"
    ie61c01ef4a777c72a9b781ade8c1b5f6894824bed5986603029a602562eca688 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/forward"
)

// EventRequestBuilder builds and executes requests for operations under \me\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\calendarView\{event-id}\exceptionOccurrences\{event-id1}
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
// EventRequestBuilderGetQueryParameters get exceptionOccurrences from me
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
func (m *EventRequestBuilder) Accept()(*ic72278706e4ce7b845b83533a17c670656d7d0710d904d3ef6218f5fa7b3d5cd.AcceptRequestBuilder) {
    return ic72278706e4ce7b845b83533a17c670656d7d0710d904d3ef6218f5fa7b3d5cd.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i0afcb2c03d1ee02a759843837718d498b3b7a31bb6a8f8567f8f2ee82e08f6b0.CancelRequestBuilder) {
    return i0afcb2c03d1ee02a759843837718d498b3b7a31bb6a8f8567f8f2ee82e08f6b0.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/calendarView/{event_id}/exceptionOccurrences/{event_id1}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for me
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
// CreateGetRequestInformation get exceptionOccurrences from me
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in me
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
func (m *EventRequestBuilder) Decline()(*iac85e1e6b35961199fed24ddde940883c14819fdd25988d6de0a6de5a9af0adf.DeclineRequestBuilder) {
    return iac85e1e6b35961199fed24ddde940883c14819fdd25988d6de0a6de5a9af0adf.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for me
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
func (m *EventRequestBuilder) DismissReminder()(*i9157189a72bc4b79acc0436a007027d935b773485d5138c2005b2ebad608dc94.DismissReminderRequestBuilder) {
    return i9157189a72bc4b79acc0436a007027d935b773485d5138c2005b2ebad608dc94.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ie61c01ef4a777c72a9b781ade8c1b5f6894824bed5986603029a602562eca688.ForwardRequestBuilder) {
    return ie61c01ef4a777c72a9b781ade8c1b5f6894824bed5986603029a602562eca688.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
// Patch update the navigation property exceptionOccurrences in me
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
func (m *EventRequestBuilder) SnoozeReminder()(*ia546eba81a33eaf6dea2e0a5a4af88df1192817310dd311ca8049116110ef0ba.SnoozeReminderRequestBuilder) {
    return ia546eba81a33eaf6dea2e0a5a4af88df1192817310dd311ca8049116110ef0ba.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i0cbde926d382023c0322c5a02756b36f5cc3555c3a3e42ccf81f0fe2328d60da.TentativelyAcceptRequestBuilder) {
    return i0cbde926d382023c0322c5a02756b36f5cc3555c3a3e42ccf81f0fe2328d60da.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
