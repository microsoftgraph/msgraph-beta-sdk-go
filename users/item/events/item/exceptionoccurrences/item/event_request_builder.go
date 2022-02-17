package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i24af1de4524059322182c9dd8cae9ebf8a39ca79890918a10f59b308d9a8ff7a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/decline"
    i33a5a3f17c3c57c4d9ecdbf5dc6b8dca428e7cea934705d5ba0434606ef702b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/snoozereminder"
    i49a1cc0de2bf20f44e8d26b7db5c08b1126860270fc8cdbeb73ec0e23dee368b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/forward"
    i56809098fcd1533ed7a3215b7a76c743f5332ffe162cbff495da959cfaf9be26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/accept"
    i8a3f3551ceec8769968f60c75e1d1e3f77020bea63406e6fbcdc0f9cc8ace867 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/dismissreminder"
    ic7bf6b5e3783a5d1aa22a4bbf19207ae67dfbbb1ca79567297d3316e2bd8b57c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/cancel"
    id88b88ba958aa11563129cac9b477d009998b987629f23e8db87302be9c85604 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/tentativelyaccept"
)

// EventRequestBuilder builds and executes requests for operations under \users\{user-id}\events\{event-id}\exceptionOccurrences\{event-id1}
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
func (m *EventRequestBuilder) Accept()(*i56809098fcd1533ed7a3215b7a76c743f5332ffe162cbff495da959cfaf9be26.AcceptRequestBuilder) {
    return i56809098fcd1533ed7a3215b7a76c743f5332ffe162cbff495da959cfaf9be26.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*ic7bf6b5e3783a5d1aa22a4bbf19207ae67dfbbb1ca79567297d3316e2bd8b57c.CancelRequestBuilder) {
    return ic7bf6b5e3783a5d1aa22a4bbf19207ae67dfbbb1ca79567297d3316e2bd8b57c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/events/{event_id}/exceptionOccurrences/{event_id1}{?select,expand}";
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
func (m *EventRequestBuilder) Decline()(*i24af1de4524059322182c9dd8cae9ebf8a39ca79890918a10f59b308d9a8ff7a.DeclineRequestBuilder) {
    return i24af1de4524059322182c9dd8cae9ebf8a39ca79890918a10f59b308d9a8ff7a.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i8a3f3551ceec8769968f60c75e1d1e3f77020bea63406e6fbcdc0f9cc8ace867.DismissReminderRequestBuilder) {
    return i8a3f3551ceec8769968f60c75e1d1e3f77020bea63406e6fbcdc0f9cc8ace867.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i49a1cc0de2bf20f44e8d26b7db5c08b1126860270fc8cdbeb73ec0e23dee368b.ForwardRequestBuilder) {
    return i49a1cc0de2bf20f44e8d26b7db5c08b1126860270fc8cdbeb73ec0e23dee368b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) SnoozeReminder()(*i33a5a3f17c3c57c4d9ecdbf5dc6b8dca428e7cea934705d5ba0434606ef702b2.SnoozeReminderRequestBuilder) {
    return i33a5a3f17c3c57c4d9ecdbf5dc6b8dca428e7cea934705d5ba0434606ef702b2.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*id88b88ba958aa11563129cac9b477d009998b987629f23e8db87302be9c85604.TentativelyAcceptRequestBuilder) {
    return id88b88ba958aa11563129cac9b477d009998b987629f23e8db87302be9c85604.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
