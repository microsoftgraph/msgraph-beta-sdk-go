package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0685118e139fcb848d59800d418daa4532ae388a4c9dad71e85e7dc66595d84c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/decline"
    i140815eababf3d3a8f8adcd68a7e5c3fcc05ff78506230177ca9402afaa389ac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/calendar"
    i20182d9c8015abb77377593ee2ec22de4a0de5b82e2931493a15ae22f1a89442 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/dismissreminder"
    i5470d0741c76cc51b5bcfc084d78c0c7dd707966a2e24964ed5338a53f0fdf5a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/singlevalueextendedproperties"
    i742b5c8460601a65705785b19afc7e8916b9c065c3040d7433c8b8559a1dda8a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences"
    i78b383ee012702182e57a257c65454b8e89c67a59e3373b3e3c9b4f0655658ed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/attachments"
    i85c6711913d8125cc394b5abca8e6b64b01e731e9f7a7808d776b24f1df58ede "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/accept"
    i93f1324d095c6e3176913302bdcb0280c717c0f8143463c11e80b637568aad11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/forward"
    i9bdbae22ae811450f012cbada6c498744fbc08a1807344b1d4da67d9e4f74cac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/cancel"
    iafb4c38e8bb94fddfea79c758ce2374b46c76fc1c083bb61d1b0209562b66c62 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/snoozereminder"
    id5cda663c28ff918cacf9a980047094bc8d07bf741d9a408379d59e91e60f157 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/extensions"
    idc485cf3c6f6b3fbf7a4b9f74595a5274efefbeef6a1c267aa6678d76295fc7d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances"
    ieddcec3480c9c607f44a9741427a19c80f4f463d2a7e20f693d62b7ac2a312fc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/multivalueextendedproperties"
    if3184f542cbb8b7c726a991685ee6a99362ce2ca926bec7586a79ac011f47ee5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/tentativelyaccept"
    i21eea82952b9a6a307ae751fe4a971dfad0180c0d38c3bf42efa3e6e524978bf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item"
    i6e8562073f5a8a463b68bd0da295e28855aa3bc85b3a841ff98cce37c5b90329 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/extensions/item"
    i9ae076e1010d857aca9d451f87b5372abc57a6b750145a6b041dd6d621ae90e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/exceptionoccurrences/item"
    ibf360dd4027c0826877c6a9d0097e99664182d0a5d6e0932edf3083429282690 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/attachments/item"
    ie87906fcdaca43b80683c3fec1e42631e58c1774dbffd64091c6dd1e7746e910 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/multivalueextendedproperties/item"
    if1309d2e33962b90d2f0d28b11874b6123980ea48d7b11ffa71cf5646f889d43 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/singlevalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
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
// EventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Accept()(*i85c6711913d8125cc394b5abca8e6b64b01e731e9f7a7808d776b24f1df58ede.AcceptRequestBuilder) {
    return i85c6711913d8125cc394b5abca8e6b64b01e731e9f7a7808d776b24f1df58ede.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i78b383ee012702182e57a257c65454b8e89c67a59e3373b3e3c9b4f0655658ed.AttachmentsRequestBuilder) {
    return i78b383ee012702182e57a257c65454b8e89c67a59e3373b3e3c9b4f0655658ed.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ibf360dd4027c0826877c6a9d0097e99664182d0a5d6e0932edf3083429282690.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ibf360dd4027c0826877c6a9d0097e99664182d0a5d6e0932edf3083429282690.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i140815eababf3d3a8f8adcd68a7e5c3fcc05ff78506230177ca9402afaa389ac.CalendarRequestBuilder) {
    return i140815eababf3d3a8f8adcd68a7e5c3fcc05ff78506230177ca9402afaa389ac.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i9bdbae22ae811450f012cbada6c498744fbc08a1807344b1d4da67d9e4f74cac.CancelRequestBuilder) {
    return i9bdbae22ae811450f012cbada6c498744fbc08a1807344b1d4da67d9e4f74cac.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar_id}/events/{event_id}{?select}";
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
// CreateDeleteRequestInformation delete navigation property events for me
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
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation update the navigation property events in me
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
func (m *EventItemRequestBuilder) Decline()(*i0685118e139fcb848d59800d418daa4532ae388a4c9dad71e85e7dc66595d84c.DeclineRequestBuilder) {
    return i0685118e139fcb848d59800d418daa4532ae388a4c9dad71e85e7dc66595d84c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i20182d9c8015abb77377593ee2ec22de4a0de5b82e2931493a15ae22f1a89442.DismissReminderRequestBuilder) {
    return i20182d9c8015abb77377593ee2ec22de4a0de5b82e2931493a15ae22f1a89442.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i742b5c8460601a65705785b19afc7e8916b9c065c3040d7433c8b8559a1dda8a.ExceptionOccurrencesRequestBuilder) {
    return i742b5c8460601a65705785b19afc7e8916b9c065c3040d7433c8b8559a1dda8a.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i9ae076e1010d857aca9d451f87b5372abc57a6b750145a6b041dd6d621ae90e8.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i9ae076e1010d857aca9d451f87b5372abc57a6b750145a6b041dd6d621ae90e8.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*id5cda663c28ff918cacf9a980047094bc8d07bf741d9a408379d59e91e60f157.ExtensionsRequestBuilder) {
    return id5cda663c28ff918cacf9a980047094bc8d07bf741d9a408379d59e91e60f157.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i6e8562073f5a8a463b68bd0da295e28855aa3bc85b3a841ff98cce37c5b90329.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i6e8562073f5a8a463b68bd0da295e28855aa3bc85b3a841ff98cce37c5b90329.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i93f1324d095c6e3176913302bdcb0280c717c0f8143463c11e80b637568aad11.ForwardRequestBuilder) {
    return i93f1324d095c6e3176913302bdcb0280c717c0f8143463c11e80b637568aad11.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*idc485cf3c6f6b3fbf7a4b9f74595a5274efefbeef6a1c267aa6678d76295fc7d.InstancesRequestBuilder) {
    return idc485cf3c6f6b3fbf7a4b9f74595a5274efefbeef6a1c267aa6678d76295fc7d.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i21eea82952b9a6a307ae751fe4a971dfad0180c0d38c3bf42efa3e6e524978bf.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i21eea82952b9a6a307ae751fe4a971dfad0180c0d38c3bf42efa3e6e524978bf.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ieddcec3480c9c607f44a9741427a19c80f4f463d2a7e20f693d62b7ac2a312fc.MultiValueExtendedPropertiesRequestBuilder) {
    return ieddcec3480c9c607f44a9741427a19c80f4f463d2a7e20f693d62b7ac2a312fc.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie87906fcdaca43b80683c3fec1e42631e58c1774dbffd64091c6dd1e7746e910.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ie87906fcdaca43b80683c3fec1e42631e58c1774dbffd64091c6dd1e7746e910.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i5470d0741c76cc51b5bcfc084d78c0c7dd707966a2e24964ed5338a53f0fdf5a.SingleValueExtendedPropertiesRequestBuilder) {
    return i5470d0741c76cc51b5bcfc084d78c0c7dd707966a2e24964ed5338a53f0fdf5a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if1309d2e33962b90d2f0d28b11874b6123980ea48d7b11ffa71cf5646f889d43.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if1309d2e33962b90d2f0d28b11874b6123980ea48d7b11ffa71cf5646f889d43.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*iafb4c38e8bb94fddfea79c758ce2374b46c76fc1c083bb61d1b0209562b66c62.SnoozeReminderRequestBuilder) {
    return iafb4c38e8bb94fddfea79c758ce2374b46c76fc1c083bb61d1b0209562b66c62.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*if3184f542cbb8b7c726a991685ee6a99362ce2ca926bec7586a79ac011f47ee5.TentativelyAcceptRequestBuilder) {
    return if3184f542cbb8b7c726a991685ee6a99362ce2ca926bec7586a79ac011f47ee5.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
