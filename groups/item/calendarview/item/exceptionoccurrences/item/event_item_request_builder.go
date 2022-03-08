package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i04fd641e5a747104f4e0952e98213abc3186667104445a5618d14edf07a6772f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/accept"
    i0fd40eb9949327ae3e38a12a37d8b889a9631e6d019eb079118b1de218b3a98f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i2c59e99a1d1b0518a61a8d556af40d5e5fe456c174ae32eb81a450a2b24ce35c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i5349e999a824e9cbe5608bee94437dd38a319cc623f6ebbbcb20b48284a07301 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/forward"
    i5952b899fd7ad196b67e16874d1e062638b772ec2c510012e1c80ca74060fd3b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/attachments"
    i606eabc5b24235e1dfe5eda3ef9afb4b61fdb7829898669879ea22faca31f014 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/decline"
    i6c23de99d6e19538949dd0a1167c1f38757696973c1e8a2653dccaeb0f21001e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i75a0f8c972b66cf74259bd9d39194d06f5b12edab8842b7b01535f9d7ddfc5ff "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/calendar"
    i75b0bb687cc2cb0e33e9dc3c36651978a7b25ce6a0bce0d1664b36167919e00d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/cancel"
    i77956e6123d355433964150d0c3b33dbb67aa42ab7897ac20e593fe2645c189a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/dismissreminder"
    i880391b83bd4a880827c4fea0d059381b921905d4de77ce3149acbac5ca1b2b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances"
    i9c550bfac2e91800570b150240869141407885a38fdbebfc217a6cd5b70c39e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    ia2f47a8345604b1cb825cc4c0f191b0ee97048e0bbaf45dde6d3c0ec566708ab "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/extensions"
    i08619015a395f2dd82b2fbc2b9782d5ae07f6659e4525de00f3e820177ae8af5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/extensions/item"
    i49367fd7902603d0bb5c5b72858052db34b775a821866c9f83431b213d3b7416 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i4bb153a846f9374b2b40acb7ce407c7465ef82ac64d3f554f200662b0b8dc364 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/attachments/item"
    ib63af9909af4567eefbd6aae62ac98ec8c8d26d16d216a9e306e48f64c58b21e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ic87ae678f816a9e2aaffe8ffeea1c20689c1a543fcfd919fe16903a838f12f65 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Accept()(*i04fd641e5a747104f4e0952e98213abc3186667104445a5618d14edf07a6772f.AcceptRequestBuilder) {
    return i04fd641e5a747104f4e0952e98213abc3186667104445a5618d14edf07a6772f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i5952b899fd7ad196b67e16874d1e062638b772ec2c510012e1c80ca74060fd3b.AttachmentsRequestBuilder) {
    return i5952b899fd7ad196b67e16874d1e062638b772ec2c510012e1c80ca74060fd3b.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i4bb153a846f9374b2b40acb7ce407c7465ef82ac64d3f554f200662b0b8dc364.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i4bb153a846f9374b2b40acb7ce407c7465ef82ac64d3f554f200662b0b8dc364.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i75a0f8c972b66cf74259bd9d39194d06f5b12edab8842b7b01535f9d7ddfc5ff.CalendarRequestBuilder) {
    return i75a0f8c972b66cf74259bd9d39194d06f5b12edab8842b7b01535f9d7ddfc5ff.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i75b0bb687cc2cb0e33e9dc3c36651978a7b25ce6a0bce0d1664b36167919e00d.CancelRequestBuilder) {
    return i75b0bb687cc2cb0e33e9dc3c36651978a7b25ce6a0bce0d1664b36167919e00d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendarView/{event_id}/exceptionOccurrences/{event_id1}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for groups
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
// CreateGetRequestInformation get exceptionOccurrences from groups
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in groups
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
func (m *EventItemRequestBuilder) Decline()(*i606eabc5b24235e1dfe5eda3ef9afb4b61fdb7829898669879ea22faca31f014.DeclineRequestBuilder) {
    return i606eabc5b24235e1dfe5eda3ef9afb4b61fdb7829898669879ea22faca31f014.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*i77956e6123d355433964150d0c3b33dbb67aa42ab7897ac20e593fe2645c189a.DismissReminderRequestBuilder) {
    return i77956e6123d355433964150d0c3b33dbb67aa42ab7897ac20e593fe2645c189a.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*ia2f47a8345604b1cb825cc4c0f191b0ee97048e0bbaf45dde6d3c0ec566708ab.ExtensionsRequestBuilder) {
    return ia2f47a8345604b1cb825cc4c0f191b0ee97048e0bbaf45dde6d3c0ec566708ab.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i08619015a395f2dd82b2fbc2b9782d5ae07f6659e4525de00f3e820177ae8af5.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i08619015a395f2dd82b2fbc2b9782d5ae07f6659e4525de00f3e820177ae8af5.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i5349e999a824e9cbe5608bee94437dd38a319cc623f6ebbbcb20b48284a07301.ForwardRequestBuilder) {
    return i5349e999a824e9cbe5608bee94437dd38a319cc623f6ebbbcb20b48284a07301.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Instances()(*i880391b83bd4a880827c4fea0d059381b921905d4de77ce3149acbac5ca1b2b3.InstancesRequestBuilder) {
    return i880391b83bd4a880827c4fea0d059381b921905d4de77ce3149acbac5ca1b2b3.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ic87ae678f816a9e2aaffe8ffeea1c20689c1a543fcfd919fe16903a838f12f65.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id2"] = id
    }
    return ic87ae678f816a9e2aaffe8ffeea1c20689c1a543fcfd919fe16903a838f12f65.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i9c550bfac2e91800570b150240869141407885a38fdbebfc217a6cd5b70c39e4.MultiValueExtendedPropertiesRequestBuilder) {
    return i9c550bfac2e91800570b150240869141407885a38fdbebfc217a6cd5b70c39e4.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i49367fd7902603d0bb5c5b72858052db34b775a821866c9f83431b213d3b7416.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i49367fd7902603d0bb5c5b72858052db34b775a821866c9f83431b213d3b7416.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i6c23de99d6e19538949dd0a1167c1f38757696973c1e8a2653dccaeb0f21001e.SingleValueExtendedPropertiesRequestBuilder) {
    return i6c23de99d6e19538949dd0a1167c1f38757696973c1e8a2653dccaeb0f21001e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib63af9909af4567eefbd6aae62ac98ec8c8d26d16d216a9e306e48f64c58b21e.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ib63af9909af4567eefbd6aae62ac98ec8c8d26d16d216a9e306e48f64c58b21e.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i0fd40eb9949327ae3e38a12a37d8b889a9631e6d019eb079118b1de218b3a98f.SnoozeReminderRequestBuilder) {
    return i0fd40eb9949327ae3e38a12a37d8b889a9631e6d019eb079118b1de218b3a98f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i2c59e99a1d1b0518a61a8d556af40d5e5fe456c174ae32eb81a450a2b24ce35c.TentativelyAcceptRequestBuilder) {
    return i2c59e99a1d1b0518a61a8d556af40d5e5fe456c174ae32eb81a450a2b24ce35c.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
