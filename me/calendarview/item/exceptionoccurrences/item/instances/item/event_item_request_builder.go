package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0486566865131f6d8b456f530e737b5e3fbe0217cb7948c373742b0033e8a44f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    i30e83456fb8329f3c046a914b62a9f85b284d4d04c01feab7ff48ddbe62aea73 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i6c41df34d0e9dcd10bd6bb9b55035c117190b5387cae4492a79016e8676f9151 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    i75f58a5c84622a9825f6c2b5aa1eaa354240f9c87c606d9aec2425c8e2363bdd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    i7a49d356c2f1e4891ed56582750b0ac8f9736d0e4452115febb4078b3f1ce278 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i8f7cdd594578e792fbacc4e8a352a95bf9edd7d2a5db3d89cf209e5b20a02370 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i95932a60acd652212cad648f475122406ee5fe09cbcff26ad2034963b5f8420d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    ib140cc00876380ead7c1ff7081b142ee49cea86bf343b4cc1b90bd5ca632d12b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    ib8f085c38c95eae4eaf7188bbc43dd54ac9ee40bddb34f21cec53e3fab2f65c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    id949bac63d6a4ebfd35170de09add8e6f96d142f6a1a64b9ba300b8af31ecc35 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    ide9ad148ca8a2edd82554f35a0b42dd5d246351e3b9fe7026b87aacb4ed7b98a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    ie9e28506d74ed610ad7baf1e6773582a5660d05a5fac19e222df6b382b1be031 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    ia9da6a2a739c1f9be2b81d9f0aa732b17555ae6c19884c33fa53e0af0fde959c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
    ib5cd9b637892b36b96975b66063dd5fc834dec6f22b2e0e4350843184de21066 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    id3b184451d6724dd598ca03eab1b7292640b90abcdbc0e3236d39601f80dec3f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    if2302784e90922c8ce421533f42bb5f00b91095587f48cb22208020f307bf040 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventItemRequestBuilder) Accept()(*i6c41df34d0e9dcd10bd6bb9b55035c117190b5387cae4492a79016e8676f9151.AcceptRequestBuilder) {
    return i6c41df34d0e9dcd10bd6bb9b55035c117190b5387cae4492a79016e8676f9151.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i75f58a5c84622a9825f6c2b5aa1eaa354240f9c87c606d9aec2425c8e2363bdd.AttachmentsRequestBuilder) {
    return i75f58a5c84622a9825f6c2b5aa1eaa354240f9c87c606d9aec2425c8e2363bdd.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia9da6a2a739c1f9be2b81d9f0aa732b17555ae6c19884c33fa53e0af0fde959c.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia9da6a2a739c1f9be2b81d9f0aa732b17555ae6c19884c33fa53e0af0fde959c.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*id949bac63d6a4ebfd35170de09add8e6f96d142f6a1a64b9ba300b8af31ecc35.CalendarRequestBuilder) {
    return id949bac63d6a4ebfd35170de09add8e6f96d142f6a1a64b9ba300b8af31ecc35.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i0486566865131f6d8b456f530e737b5e3fbe0217cb7948c373742b0033e8a44f.CancelRequestBuilder) {
    return i0486566865131f6d8b456f530e737b5e3fbe0217cb7948c373742b0033e8a44f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event_id}/exceptionOccurrences/{event_id1}/instances/{event_id2}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property instances for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property instances in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventItemRequestBuilder) Decline()(*ide9ad148ca8a2edd82554f35a0b42dd5d246351e3b9fe7026b87aacb4ed7b98a.DeclineRequestBuilder) {
    return ide9ad148ca8a2edd82554f35a0b42dd5d246351e3b9fe7026b87aacb4ed7b98a.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*i30e83456fb8329f3c046a914b62a9f85b284d4d04c01feab7ff48ddbe62aea73.DismissReminderRequestBuilder) {
    return i30e83456fb8329f3c046a914b62a9f85b284d4d04c01feab7ff48ddbe62aea73.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*ib8f085c38c95eae4eaf7188bbc43dd54ac9ee40bddb34f21cec53e3fab2f65c0.ExtensionsRequestBuilder) {
    return ib8f085c38c95eae4eaf7188bbc43dd54ac9ee40bddb34f21cec53e3fab2f65c0.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ib5cd9b637892b36b96975b66063dd5fc834dec6f22b2e0e4350843184de21066.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ib5cd9b637892b36b96975b66063dd5fc834dec6f22b2e0e4350843184de21066.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i95932a60acd652212cad648f475122406ee5fe09cbcff26ad2034963b5f8420d.ForwardRequestBuilder) {
    return i95932a60acd652212cad648f475122406ee5fe09cbcff26ad2034963b5f8420d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable), nil
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie9e28506d74ed610ad7baf1e6773582a5660d05a5fac19e222df6b382b1be031.MultiValueExtendedPropertiesRequestBuilder) {
    return ie9e28506d74ed610ad7baf1e6773582a5660d05a5fac19e222df6b382b1be031.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*id3b184451d6724dd598ca03eab1b7292640b90abcdbc0e3236d39601f80dec3f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return id3b184451d6724dd598ca03eab1b7292640b90abcdbc0e3236d39601f80dec3f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ib140cc00876380ead7c1ff7081b142ee49cea86bf343b4cc1b90bd5ca632d12b.SingleValueExtendedPropertiesRequestBuilder) {
    return ib140cc00876380ead7c1ff7081b142ee49cea86bf343b4cc1b90bd5ca632d12b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if2302784e90922c8ce421533f42bb5f00b91095587f48cb22208020f307bf040.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if2302784e90922c8ce421533f42bb5f00b91095587f48cb22208020f307bf040.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i8f7cdd594578e792fbacc4e8a352a95bf9edd7d2a5db3d89cf209e5b20a02370.SnoozeReminderRequestBuilder) {
    return i8f7cdd594578e792fbacc4e8a352a95bf9edd7d2a5db3d89cf209e5b20a02370.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i7a49d356c2f1e4891ed56582750b0ac8f9736d0e4452115febb4078b3f1ce278.TentativelyAcceptRequestBuilder) {
    return i7a49d356c2f1e4891ed56582750b0ac8f9736d0e4452115febb4078b3f1ce278.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
