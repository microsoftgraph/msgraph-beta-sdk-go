package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i189cf07452004a59fdc9a312b85613cf42e674b1268f40892744850e80393bfa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i18f5bf7222ff32fb24ccc8f6d32737d602e9f7bc09f143b35f9d132c2a2021aa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i245482dd46f7d17bb831b43d6b1ad96ff3e81020caf86ffb9b6d599037540c0f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i30d9d0230e5d2db706ed53b91c53ae308e2b8358c9cd7c22d46c1c29877c1af0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    i33447d1fb62e68bdd02fc6c549dfff416b77c7baeee5cb43d7ecc14cb29557b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i3c17378b532dfd56069cdb9f5ea5ea3b8978bbd5e503ae433d1354f572cecad9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i603d786505660fcad5c74082c44a538a12268bbe066125af061fc34e8fbf6311 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i83c05d9ac07f820cfa6d1ad0162434a1c6dd103ab5066bcab1cd38cb219140d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    iaac5fefd553b7432398b65c83dc236c81f390dff713b631a9a2ccd7de84936d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    id09cdc1412d932b6e8acaff79d58551587c44cdbab5dc27b7a61f10007f4ff85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    idc76f6a8ee03d34512ffad53b611cabbd8b9015cc991aae550d8697bde22588e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    if75001e11938f08c700df05d71ae477ea3cde9a9c6cf9d17a91b35a55ab4261d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i0bf2d76aebec29acdb979979d4b8a74c9a8352845eb6dd72584313051773f121 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i3288dde21fb1b1441b04dff2b72ba0b3a4daf0269a8bde7c09b26d780dc938e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    ia3428c7cb98f457b22c0c9745b715d96ec19ce770aed036ac37ac456f2e90bdf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ida2e6f17dbd1a69a5e2c6bc9d98e7ffc20ce68a999e0c960b044302c74ccd843 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
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
func (m *EventItemRequestBuilder) Accept()(*iaac5fefd553b7432398b65c83dc236c81f390dff713b631a9a2ccd7de84936d5.AcceptRequestBuilder) {
    return iaac5fefd553b7432398b65c83dc236c81f390dff713b631a9a2ccd7de84936d5.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i603d786505660fcad5c74082c44a538a12268bbe066125af061fc34e8fbf6311.AttachmentsRequestBuilder) {
    return i603d786505660fcad5c74082c44a538a12268bbe066125af061fc34e8fbf6311.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ida2e6f17dbd1a69a5e2c6bc9d98e7ffc20ce68a999e0c960b044302c74ccd843.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ida2e6f17dbd1a69a5e2c6bc9d98e7ffc20ce68a999e0c960b044302c74ccd843.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*idc76f6a8ee03d34512ffad53b611cabbd8b9015cc991aae550d8697bde22588e.CalendarRequestBuilder) {
    return idc76f6a8ee03d34512ffad53b611cabbd8b9015cc991aae550d8697bde22588e.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i33447d1fb62e68bdd02fc6c549dfff416b77c7baeee5cb43d7ecc14cb29557b0.CancelRequestBuilder) {
    return i33447d1fb62e68bdd02fc6c549dfff416b77c7baeee5cb43d7ecc14cb29557b0.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendar/calendarView/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for users
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
// CreateGetRequestInformation get exceptionOccurrences from users
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) Decline()(*i30d9d0230e5d2db706ed53b91c53ae308e2b8358c9cd7c22d46c1c29877c1af0.DeclineRequestBuilder) {
    return i30d9d0230e5d2db706ed53b91c53ae308e2b8358c9cd7c22d46c1c29877c1af0.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i245482dd46f7d17bb831b43d6b1ad96ff3e81020caf86ffb9b6d599037540c0f.DismissReminderRequestBuilder) {
    return i245482dd46f7d17bb831b43d6b1ad96ff3e81020caf86ffb9b6d599037540c0f.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i18f5bf7222ff32fb24ccc8f6d32737d602e9f7bc09f143b35f9d132c2a2021aa.ExtensionsRequestBuilder) {
    return i18f5bf7222ff32fb24ccc8f6d32737d602e9f7bc09f143b35f9d132c2a2021aa.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i3288dde21fb1b1441b04dff2b72ba0b3a4daf0269a8bde7c09b26d780dc938e9.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i3288dde21fb1b1441b04dff2b72ba0b3a4daf0269a8bde7c09b26d780dc938e9.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i189cf07452004a59fdc9a312b85613cf42e674b1268f40892744850e80393bfa.ForwardRequestBuilder) {
    return i189cf07452004a59fdc9a312b85613cf42e674b1268f40892744850e80393bfa.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*if75001e11938f08c700df05d71ae477ea3cde9a9c6cf9d17a91b35a55ab4261d.MultiValueExtendedPropertiesRequestBuilder) {
    return if75001e11938f08c700df05d71ae477ea3cde9a9c6cf9d17a91b35a55ab4261d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0bf2d76aebec29acdb979979d4b8a74c9a8352845eb6dd72584313051773f121.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0bf2d76aebec29acdb979979d4b8a74c9a8352845eb6dd72584313051773f121.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i83c05d9ac07f820cfa6d1ad0162434a1c6dd103ab5066bcab1cd38cb219140d9.SingleValueExtendedPropertiesRequestBuilder) {
    return i83c05d9ac07f820cfa6d1ad0162434a1c6dd103ab5066bcab1cd38cb219140d9.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia3428c7cb98f457b22c0c9745b715d96ec19ce770aed036ac37ac456f2e90bdf.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ia3428c7cb98f457b22c0c9745b715d96ec19ce770aed036ac37ac456f2e90bdf.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i3c17378b532dfd56069cdb9f5ea5ea3b8978bbd5e503ae433d1354f572cecad9.SnoozeReminderRequestBuilder) {
    return i3c17378b532dfd56069cdb9f5ea5ea3b8978bbd5e503ae433d1354f572cecad9.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*id09cdc1412d932b6e8acaff79d58551587c44cdbab5dc27b7a61f10007f4ff85.TentativelyAcceptRequestBuilder) {
    return id09cdc1412d932b6e8acaff79d58551587c44cdbab5dc27b7a61f10007f4ff85.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
