package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2764131d6064d55a6fb28f2c0abcae47736dde3f40e965b0eda7a971dde5de10 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i27a3a352b6ff7d9499e5d09898d04ce38661929729d0d9a6d496952e9f70c30e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i34b1fded7e2d95e2cdfa2275ee4419128351ddd305cd95a46e6b51cc48b67976 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i3574e76278c0a6f91b08c48f88683e4dba058986d9592c2aba37b074d381c179 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i5ee649c1233ae64934fe34f5783abab787512a61101c1f720f81b969f0b1decb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    i665ca788afaaae309174146e34bb77894cf47bb2bba420c30a1eb0912158594d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i991412ae494f86f4d381089968289cae90002d2cb421d0f94c320e30b5561d25 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ibed6e83805e6013a18a6e19fbddd1cbd35347d5c1914d17080c3413d73bcbba9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    id19ec0743b30d2cbd1758bd779f1e99ddc91c5b14dc665e6c945eaf4623aa27d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    ida7bdbd0c64c580394b466feda1a6af9066a6919f0535f64131a9a1866d55c3b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    idbf0a00c08b54cfd7fac35e38445dc93328b6d746213307f32af7de5b59220ac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    ie102b9d30203f964bbf8844901a37d39d3c1917d954fe13a686d3f0efd36a9e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i3f862a33674900febc35502d782a7c3b473a6235fe1834b43fa3e93a0d3756c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i46222f284134be595ab501e2db9e55e6883b67f43c32ddf4bd9a55f9d96d2744 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    i8883f5c2ffbaa391bb23a2ccdb377fcc9971efd46b510c8a5d361908abfa4ab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    ib8c9f1b1d8fc0e2d1b4c7f619e86de8a7306494c87eccea322c28ca34745d939 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i5ee649c1233ae64934fe34f5783abab787512a61101c1f720f81b969f0b1decb.AcceptRequestBuilder) {
    return i5ee649c1233ae64934fe34f5783abab787512a61101c1f720f81b969f0b1decb.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i27a3a352b6ff7d9499e5d09898d04ce38661929729d0d9a6d496952e9f70c30e.AttachmentsRequestBuilder) {
    return i27a3a352b6ff7d9499e5d09898d04ce38661929729d0d9a6d496952e9f70c30e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i8883f5c2ffbaa391bb23a2ccdb377fcc9971efd46b510c8a5d361908abfa4ab5.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i8883f5c2ffbaa391bb23a2ccdb377fcc9971efd46b510c8a5d361908abfa4ab5.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ida7bdbd0c64c580394b466feda1a6af9066a6919f0535f64131a9a1866d55c3b.CalendarRequestBuilder) {
    return ida7bdbd0c64c580394b466feda1a6af9066a6919f0535f64131a9a1866d55c3b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i2764131d6064d55a6fb28f2c0abcae47736dde3f40e965b0eda7a971dde5de10.CancelRequestBuilder) {
    return i2764131d6064d55a6fb28f2c0abcae47736dde3f40e965b0eda7a971dde5de10.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarView/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
func (m *EventItemRequestBuilder) Decline()(*id19ec0743b30d2cbd1758bd779f1e99ddc91c5b14dc665e6c945eaf4623aa27d.DeclineRequestBuilder) {
    return id19ec0743b30d2cbd1758bd779f1e99ddc91c5b14dc665e6c945eaf4623aa27d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i34b1fded7e2d95e2cdfa2275ee4419128351ddd305cd95a46e6b51cc48b67976.DismissReminderRequestBuilder) {
    return i34b1fded7e2d95e2cdfa2275ee4419128351ddd305cd95a46e6b51cc48b67976.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i665ca788afaaae309174146e34bb77894cf47bb2bba420c30a1eb0912158594d.ExtensionsRequestBuilder) {
    return i665ca788afaaae309174146e34bb77894cf47bb2bba420c30a1eb0912158594d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i46222f284134be595ab501e2db9e55e6883b67f43c32ddf4bd9a55f9d96d2744.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i46222f284134be595ab501e2db9e55e6883b67f43c32ddf4bd9a55f9d96d2744.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*ibed6e83805e6013a18a6e19fbddd1cbd35347d5c1914d17080c3413d73bcbba9.ForwardRequestBuilder) {
    return ibed6e83805e6013a18a6e19fbddd1cbd35347d5c1914d17080c3413d73bcbba9.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie102b9d30203f964bbf8844901a37d39d3c1917d954fe13a686d3f0efd36a9e5.MultiValueExtendedPropertiesRequestBuilder) {
    return ie102b9d30203f964bbf8844901a37d39d3c1917d954fe13a686d3f0efd36a9e5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib8c9f1b1d8fc0e2d1b4c7f619e86de8a7306494c87eccea322c28ca34745d939.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ib8c9f1b1d8fc0e2d1b4c7f619e86de8a7306494c87eccea322c28ca34745d939.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i991412ae494f86f4d381089968289cae90002d2cb421d0f94c320e30b5561d25.SingleValueExtendedPropertiesRequestBuilder) {
    return i991412ae494f86f4d381089968289cae90002d2cb421d0f94c320e30b5561d25.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i3f862a33674900febc35502d782a7c3b473a6235fe1834b43fa3e93a0d3756c3.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i3f862a33674900febc35502d782a7c3b473a6235fe1834b43fa3e93a0d3756c3.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*idbf0a00c08b54cfd7fac35e38445dc93328b6d746213307f32af7de5b59220ac.SnoozeReminderRequestBuilder) {
    return idbf0a00c08b54cfd7fac35e38445dc93328b6d746213307f32af7de5b59220ac.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3574e76278c0a6f91b08c48f88683e4dba058986d9592c2aba37b074d381c179.TentativelyAcceptRequestBuilder) {
    return i3574e76278c0a6f91b08c48f88683e4dba058986d9592c2aba37b074d381c179.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
