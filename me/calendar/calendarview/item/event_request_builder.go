package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i07b0f5c126d3fd8e32f7c54a9d6f30a2c41a47efc37fbc9e8e49bdae5c2fd9e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/dismissreminder"
    i1a941c95f5627fa057d3dd8f69e7a76f5d91a7980d086a7888814dfd3e470968 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences"
    i296b15e1aad1ec85f3f8b45773230524b885d2afe41813071aa2b54c663963a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/forward"
    i4198a050c6fc737c2f34080e3829e36d908a821a9d7591fb186578ed061ee36d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/accept"
    i45997cdc7a813a982ecba7b8d3a586df5f2d88422358d670e2214804e3b53d00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/attachments"
    i4d8785343bb818131232baf9fc68c1d519ffd7bdf81da717ad76b235e251b9c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/multivalueextendedproperties"
    i535cf1cace249b664640330f949a728d46b263902138f84849c221b9abe894c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/calendar"
    i9a40e362aa3e36ede13cbac3bcd8bfe04901a57f8fb7b3971b67b06bf4459783 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances"
    ibe4fece240dca74238f03c2fa5a7bbcdd654c947e8b78e285b49ab9c7dfea505 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/tentativelyaccept"
    ibf789e58fb65fce33e2cae00e3c293e6ed19d7976b391f80d636389ef81f352c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/decline"
    icf3fd4bf51eebd15a020fc5021462c2e0a5518f7116ed3243ab0a1db3517572c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/cancel"
    id520c21d1389ffd38b7be74ec9a55a88e91677de395badf00b7ec8d8fa5400be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/extensions"
    if239cd3f23a0eda2f101672148a42c42a844663475919c4ecf0efcbe29330f19 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/snoozereminder"
    ifbc956c1dc35585d5237ea921bfac64de28ea44bcca4be498442c82a7a1f0a1e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/singlevalueextendedproperties"
    ia87adf0c3cc4103016c1134c4bb596fb8685d587b22a450dd7199fe89763d0ad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item"
    iadba4b231f387f8389de2e976dffa453258e027a2e2a2e9b08e4283bd5bc7c45 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/extensions/item"
    ib8a97a009d82433874d4147f1a44d955f32a3435ad75792434ce1d59aaac5ae0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/attachments/item"
    ic0c0093229c77d1529a83c452145260e3b4924f194817de9810cfd2218f001e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/multivalueextendedproperties/item"
    if12270c62b6873dd5445fa0f6853dc371ddec134c8027ed30812ed5c01cbd2a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/instances/item"
    ifae272c2a7c5a0611ecb939bd341b17d95d1884ae7c2800fcad35ac9b30a8fcb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/singlevalueextendedproperties/item"
)

// EventRequestBuilder builds and executes requests for operations under \me\calendar\calendarView\{event-id}
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
// EventRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type EventRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
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
func (m *EventRequestBuilder) Accept()(*i4198a050c6fc737c2f34080e3829e36d908a821a9d7591fb186578ed061ee36d.AcceptRequestBuilder) {
    return i4198a050c6fc737c2f34080e3829e36d908a821a9d7591fb186578ed061ee36d.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i45997cdc7a813a982ecba7b8d3a586df5f2d88422358d670e2214804e3b53d00.AttachmentsRequestBuilder) {
    return i45997cdc7a813a982ecba7b8d3a586df5f2d88422358d670e2214804e3b53d00.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ib8a97a009d82433874d4147f1a44d955f32a3435ad75792434ce1d59aaac5ae0.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ib8a97a009d82433874d4147f1a44d955f32a3435ad75792434ce1d59aaac5ae0.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i535cf1cace249b664640330f949a728d46b263902138f84849c221b9abe894c5.CalendarRequestBuilder) {
    return i535cf1cace249b664640330f949a728d46b263902138f84849c221b9abe894c5.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*icf3fd4bf51eebd15a020fc5021462c2e0a5518f7116ed3243ab0a1db3517572c.CancelRequestBuilder) {
    return icf3fd4bf51eebd15a020fc5021462c2e0a5518f7116ed3243ab0a1db3517572c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event_id}{?startDateTime,endDateTime,select}";
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
// CreateDeleteRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*ibf789e58fb65fce33e2cae00e3c293e6ed19d7976b391f80d636389ef81f352c.DeclineRequestBuilder) {
    return ibf789e58fb65fce33e2cae00e3c293e6ed19d7976b391f80d636389ef81f352c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*i07b0f5c126d3fd8e32f7c54a9d6f30a2c41a47efc37fbc9e8e49bdae5c2fd9e0.DismissReminderRequestBuilder) {
    return i07b0f5c126d3fd8e32f7c54a9d6f30a2c41a47efc37fbc9e8e49bdae5c2fd9e0.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*i1a941c95f5627fa057d3dd8f69e7a76f5d91a7980d086a7888814dfd3e470968.ExceptionOccurrencesRequestBuilder) {
    return i1a941c95f5627fa057d3dd8f69e7a76f5d91a7980d086a7888814dfd3e470968.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*ia87adf0c3cc4103016c1134c4bb596fb8685d587b22a450dd7199fe89763d0ad.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ia87adf0c3cc4103016c1134c4bb596fb8685d587b22a450dd7199fe89763d0ad.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*id520c21d1389ffd38b7be74ec9a55a88e91677de395badf00b7ec8d8fa5400be.ExtensionsRequestBuilder) {
    return id520c21d1389ffd38b7be74ec9a55a88e91677de395badf00b7ec8d8fa5400be.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*iadba4b231f387f8389de2e976dffa453258e027a2e2a2e9b08e4283bd5bc7c45.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return iadba4b231f387f8389de2e976dffa453258e027a2e2a2e9b08e4283bd5bc7c45.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i296b15e1aad1ec85f3f8b45773230524b885d2afe41813071aa2b54c663963a6.ForwardRequestBuilder) {
    return i296b15e1aad1ec85f3f8b45773230524b885d2afe41813071aa2b54c663963a6.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Instances()(*i9a40e362aa3e36ede13cbac3bcd8bfe04901a57f8fb7b3971b67b06bf4459783.InstancesRequestBuilder) {
    return i9a40e362aa3e36ede13cbac3bcd8bfe04901a57f8fb7b3971b67b06bf4459783.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*if12270c62b6873dd5445fa0f6853dc371ddec134c8027ed30812ed5c01cbd2a1.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return if12270c62b6873dd5445fa0f6853dc371ddec134c8027ed30812ed5c01cbd2a1.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i4d8785343bb818131232baf9fc68c1d519ffd7bdf81da717ad76b235e251b9c0.MultiValueExtendedPropertiesRequestBuilder) {
    return i4d8785343bb818131232baf9fc68c1d519ffd7bdf81da717ad76b235e251b9c0.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic0c0093229c77d1529a83c452145260e3b4924f194817de9810cfd2218f001e2.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ic0c0093229c77d1529a83c452145260e3b4924f194817de9810cfd2218f001e2.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ifbc956c1dc35585d5237ea921bfac64de28ea44bcca4be498442c82a7a1f0a1e.SingleValueExtendedPropertiesRequestBuilder) {
    return ifbc956c1dc35585d5237ea921bfac64de28ea44bcca4be498442c82a7a1f0a1e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifae272c2a7c5a0611ecb939bd341b17d95d1884ae7c2800fcad35ac9b30a8fcb.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ifae272c2a7c5a0611ecb939bd341b17d95d1884ae7c2800fcad35ac9b30a8fcb.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*if239cd3f23a0eda2f101672148a42c42a844663475919c4ecf0efcbe29330f19.SnoozeReminderRequestBuilder) {
    return if239cd3f23a0eda2f101672148a42c42a844663475919c4ecf0efcbe29330f19.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ibe4fece240dca74238f03c2fa5a7bbcdd654c947e8b78e285b49ab9c7dfea505.TentativelyAcceptRequestBuilder) {
    return ibe4fece240dca74238f03c2fa5a7bbcdd654c947e8b78e285b49ab9c7dfea505.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
