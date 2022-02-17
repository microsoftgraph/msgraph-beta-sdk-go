package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i04c719c895acec6c592a8d7a8c99daedee4d6e03576774242ac518f4ca97f22d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/cancel"
    i0894b6f6226727a9418b410b051c30ca7a83dbf5724623855be38e9b8fc84136 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/forward"
    i08d9c816ec9504a1ccd4dd0d6cd8314afa08cb82d1b91e11064c54e552604d27 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences"
    i0d950e836e4744d2f557f6e0ef4ec3b62f3c1080cd446b0354fdbfec72951f7d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/dismissreminder"
    i0fcaa0c479a3c654579f8181a39c814649f1d5464b954e178ccbc7be2305818a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/attachments"
    i1cd61b81d19facac87465cbccad8f16fa4806738ff2736b95214e7814016ce2c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances"
    i6bcd189817e40fc5716d15a67c6a972c603e467a6275999b2cda08503ca1b274 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/accept"
    iaaec2ffe60b7e88a7eec860d9dda02fca6f1fd301915aa6e7ed5433699a22ece "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/extensions"
    ib4dcf73ff3c2130c3449325e737109b0274434aa8c7a4e3586acaa00f0698c32 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/snoozereminder"
    ibd0ab74cf035804e4a297f74c17c9116e6d51fbbf5ab1e68f8a3a787f4072ea2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/calendar"
    id2289c2a6f838883659e046083563cf60dcd5e984c949f14b689fe8418054ed1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/multivalueextendedproperties"
    idcfc0f6f13cfe31fd44775b67fb3c4db7708a9473bba0eb0ec30fefbeca28ba6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties"
    ie03c784e30aa926a3fc3ca0c4ee7437288f53eb208b57774aaf64385b3bc24dd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/tentativelyaccept"
    if72ad66e54b22cbc558eeb86a695bb7ac2d89c6beaf41cb86224d55c3fe01396 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/decline"
    i094f05a8ac64a75e4ef8691bd5e1036db5cd0689451270938b4712f947bd811c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties/item"
    i687127d3891d370eb8d20a210a67df658676908c6ec8ff74ffc0251bc61b5039 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/attachments/item"
    i69d244b2e4785f363c5b08786031cfdc6f1ea6786860dbba8bd696eb8bdf1eb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/multivalueextendedproperties/item"
    ia4845a22710fbc6645ada0a9de578d7004512c8ce99f327697a9e3148117b4e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item"
    ib911879fbf568651ae56e76797cf7895cbd515a33dfeb908d9088187feb3d43b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/extensions/item"
    if4a386b83ac6acb1bc708653d19283aa35f0d4ef6552877c7f9095c9bc2a807a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item"
)

// EventRequestBuilder builds and executes requests for operations under \me\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\events\{event-id}
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
// EventRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Accept()(*i6bcd189817e40fc5716d15a67c6a972c603e467a6275999b2cda08503ca1b274.AcceptRequestBuilder) {
    return i6bcd189817e40fc5716d15a67c6a972c603e467a6275999b2cda08503ca1b274.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i0fcaa0c479a3c654579f8181a39c814649f1d5464b954e178ccbc7be2305818a.AttachmentsRequestBuilder) {
    return i0fcaa0c479a3c654579f8181a39c814649f1d5464b954e178ccbc7be2305818a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*i687127d3891d370eb8d20a210a67df658676908c6ec8ff74ffc0251bc61b5039.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i687127d3891d370eb8d20a210a67df658676908c6ec8ff74ffc0251bc61b5039.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*ibd0ab74cf035804e4a297f74c17c9116e6d51fbbf5ab1e68f8a3a787f4072ea2.CalendarRequestBuilder) {
    return ibd0ab74cf035804e4a297f74c17c9116e6d51fbbf5ab1e68f8a3a787f4072ea2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i04c719c895acec6c592a8d7a8c99daedee4d6e03576774242ac518f4ca97f22d.CancelRequestBuilder) {
    return i04c719c895acec6c592a8d7a8c99daedee4d6e03576774242ac518f4ca97f22d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/events/{event_id}{?select}";
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
// CreateDeleteRequestInformation the events in the calendar. Navigation property. Read-only.
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
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Decline()(*if72ad66e54b22cbc558eeb86a695bb7ac2d89c6beaf41cb86224d55c3fe01396.DeclineRequestBuilder) {
    return if72ad66e54b22cbc558eeb86a695bb7ac2d89c6beaf41cb86224d55c3fe01396.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) DismissReminder()(*i0d950e836e4744d2f557f6e0ef4ec3b62f3c1080cd446b0354fdbfec72951f7d.DismissReminderRequestBuilder) {
    return i0d950e836e4744d2f557f6e0ef4ec3b62f3c1080cd446b0354fdbfec72951f7d.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*i08d9c816ec9504a1ccd4dd0d6cd8314afa08cb82d1b91e11064c54e552604d27.ExceptionOccurrencesRequestBuilder) {
    return i08d9c816ec9504a1ccd4dd0d6cd8314afa08cb82d1b91e11064c54e552604d27.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*if4a386b83ac6acb1bc708653d19283aa35f0d4ef6552877c7f9095c9bc2a807a.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return if4a386b83ac6acb1bc708653d19283aa35f0d4ef6552877c7f9095c9bc2a807a.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*iaaec2ffe60b7e88a7eec860d9dda02fca6f1fd301915aa6e7ed5433699a22ece.ExtensionsRequestBuilder) {
    return iaaec2ffe60b7e88a7eec860d9dda02fca6f1fd301915aa6e7ed5433699a22ece.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*ib911879fbf568651ae56e76797cf7895cbd515a33dfeb908d9088187feb3d43b.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ib911879fbf568651ae56e76797cf7895cbd515a33dfeb908d9088187feb3d43b.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i0894b6f6226727a9418b410b051c30ca7a83dbf5724623855be38e9b8fc84136.ForwardRequestBuilder) {
    return i0894b6f6226727a9418b410b051c30ca7a83dbf5724623855be38e9b8fc84136.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) Instances()(*i1cd61b81d19facac87465cbccad8f16fa4806738ff2736b95214e7814016ce2c.InstancesRequestBuilder) {
    return i1cd61b81d19facac87465cbccad8f16fa4806738ff2736b95214e7814016ce2c.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*ia4845a22710fbc6645ada0a9de578d7004512c8ce99f327697a9e3148117b4e2.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ia4845a22710fbc6645ada0a9de578d7004512c8ce99f327697a9e3148117b4e2.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*id2289c2a6f838883659e046083563cf60dcd5e984c949f14b689fe8418054ed1.MultiValueExtendedPropertiesRequestBuilder) {
    return id2289c2a6f838883659e046083563cf60dcd5e984c949f14b689fe8418054ed1.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i69d244b2e4785f363c5b08786031cfdc6f1ea6786860dbba8bd696eb8bdf1eb9.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i69d244b2e4785f363c5b08786031cfdc6f1ea6786860dbba8bd696eb8bdf1eb9.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the events in the calendar. Navigation property. Read-only.
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*idcfc0f6f13cfe31fd44775b67fb3c4db7708a9473bba0eb0ec30fefbeca28ba6.SingleValueExtendedPropertiesRequestBuilder) {
    return idcfc0f6f13cfe31fd44775b67fb3c4db7708a9473bba0eb0ec30fefbeca28ba6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i094f05a8ac64a75e4ef8691bd5e1036db5cd0689451270938b4712f947bd811c.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i094f05a8ac64a75e4ef8691bd5e1036db5cd0689451270938b4712f947bd811c.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*ib4dcf73ff3c2130c3449325e737109b0274434aa8c7a4e3586acaa00f0698c32.SnoozeReminderRequestBuilder) {
    return ib4dcf73ff3c2130c3449325e737109b0274434aa8c7a4e3586acaa00f0698c32.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ie03c784e30aa926a3fc3ca0c4ee7437288f53eb208b57774aaf64385b3bc24dd.TentativelyAcceptRequestBuilder) {
    return ie03c784e30aa926a3fc3ca0c4ee7437288f53eb208b57774aaf64385b3bc24dd.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
