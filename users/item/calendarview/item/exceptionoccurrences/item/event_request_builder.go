package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3c9cfac66a8ffdac7c174188decaed53d9a1421ddf2e8337a5b9469759acf107 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/forward"
    i47d7ddc9026bc47718982620dae541994b3aff78a6d3c6209e12464488531513 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i4d9faee480292c3fb3968a9d7b1d6294c6bff3c494cc164bb2f6d98cbe3acec1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i63e99f8731989a64d54014f801fde09e691078a487be0f2bef59417c1e149ea5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/decline"
    i7753f0930d8de9052fe364f67381cab9b05b6f24a32e80c9df5484bbc775eda4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/cancel"
    ia45ffc28a0b64efc01cd2426317cc2c5ae8508c40a50b461779b760a71b881fa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/dismissreminder"
    icb1e1d55ab4aa707679af47394f91954a69329f97d05bf04e1264225ccad1744 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/accept"
)

// Builds and executes requests for operations under \users\{user-id}\calendarView\{event-id}\exceptionOccurrences\{event-id1}
type EventRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type EventRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// Get exceptionOccurrences from users
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
func (m *EventRequestBuilder) Accept()(*icb1e1d55ab4aa707679af47394f91954a69329f97d05bf04e1264225ccad1744.AcceptRequestBuilder) {
    return icb1e1d55ab4aa707679af47394f91954a69329f97d05bf04e1264225ccad1744.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i7753f0930d8de9052fe364f67381cab9b05b6f24a32e80c9df5484bbc775eda4.CancelRequestBuilder) {
    return i7753f0930d8de9052fe364f67381cab9b05b6f24a32e80c9df5484bbc775eda4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/calendarView/{event_id}/exceptionOccurrences/{event_id1}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new EventRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property exceptionOccurrences for users
// Parameters:
//  - options : Options for the request
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
// Get exceptionOccurrences from users
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) CreateGetRequestInformation(options *EventRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update the navigation property exceptionOccurrences in users
// Parameters:
//  - options : Options for the request
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
func (m *EventRequestBuilder) Decline()(*i63e99f8731989a64d54014f801fde09e691078a487be0f2bef59417c1e149ea5.DeclineRequestBuilder) {
    return i63e99f8731989a64d54014f801fde09e691078a487be0f2bef59417c1e149ea5.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete navigation property exceptionOccurrences for users
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) Delete(options *EventRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) DismissReminder()(*ia45ffc28a0b64efc01cd2426317cc2c5ae8508c40a50b461779b760a71b881fa.DismissReminderRequestBuilder) {
    return ia45ffc28a0b64efc01cd2426317cc2c5ae8508c40a50b461779b760a71b881fa.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i3c9cfac66a8ffdac7c174188decaed53d9a1421ddf2e8337a5b9469759acf107.ForwardRequestBuilder) {
    return i3c9cfac66a8ffdac7c174188decaed53d9a1421ddf2e8337a5b9469759acf107.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get exceptionOccurrences from users
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) Get(options *EventRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvent() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event), nil
}
// Update the navigation property exceptionOccurrences in users
// Parameters:
//  - options : Options for the request
func (m *EventRequestBuilder) Patch(options *EventRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventRequestBuilder) SnoozeReminder()(*i4d9faee480292c3fb3968a9d7b1d6294c6bff3c494cc164bb2f6d98cbe3acec1.SnoozeReminderRequestBuilder) {
    return i4d9faee480292c3fb3968a9d7b1d6294c6bff3c494cc164bb2f6d98cbe3acec1.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i47d7ddc9026bc47718982620dae541994b3aff78a6d3c6209e12464488531513.TentativelyAcceptRequestBuilder) {
    return i47d7ddc9026bc47718982620dae541994b3aff78a6d3c6209e12464488531513.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
