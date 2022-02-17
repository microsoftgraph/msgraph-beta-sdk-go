package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i276be7fad31bd680ac1411035eef4f3a00776a60d1125d214d24b184f7469220 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances"
    i3145dece064d967ee1c16e3e936c47ea8011ed2847cccb8dedaaa5c1eedaa013 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/cancel"
    i5986037676fe795074e810e4246ff753605940169fbbeed27b2aefa176b11751 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/calendar"
    i6caf8586c74aeb53bd8b7a86089cf9aa13e361aa8d0dd14f84486a74e16e0222 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/dismissreminder"
    i88425e8a17893a7e305a95729332a64c30460d352b0bf5845f5d7090f92c8f2b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/snoozereminder"
    i88ecb5106cc0bb521ce42c551cf99da444f6f2720897e5c4464c18a087b9fa03 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/decline"
    ia22ed1a7b043dc5283f91ddeb06193ce26ef17e32a8a4abd851e6155127ed218 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/tentativelyaccept"
    iaa58f3091eded3bb8f8022ff383adf6893ae04638de3510fb6f50e685897d70d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/accept"
    iab57ad64a283353ed19ad1d9c7814f25835c7ab4a0d769767a6db218d4703ee7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties"
    ib0aa93c8bad1e825bd5cdfdc4fa22e31b84830d4f2814d7e3ccfe19465129484 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/attachments"
    ib53b9735ba043d622c998452495258595cc62dbf7581e5a3605dfc29007b5997 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/forward"
    ic53f6b01f6b3da1d49fabe4f45413f9d94236e324fbc5f83450593600644550a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties"
    ic99712d0b0d1392d14e91b9fb4b95688da8938561196adae18f98062d303b315 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences"
    ie128b633259361fd2a945ba339cd6061e92d56f2cdc9422dd860b687dc19ebfb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/extensions"
    i3e68bb94a10239ef0489773e8e20a5d33a576d00b16ded587e7df49c942e261a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/extensions/item"
    ib8afe549ec5fafbb94854314181e9eac838b20474be31174184b3169d64d40a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/attachments/item"
    ic9ed65a5fbed6446506612533b9a4560f521e68e07c8a6e70fc6eb495cf13716 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/multivalueextendedproperties/item"
    id21bb72bf65597e6d902b0acc913993108ac302ca65b2d1f3c3fcaaada16028d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/instances/item"
    id97c351c3207d208122c8654eec8fa4fae5d90b3136fc2c79c3605d5a1fee25d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item"
    if2b4feb3f0e58d1cd206b56ee504e531df602a972bcdd86373ae72f7f5884e6f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/calendarview/item/singlevalueextendedproperties/item"
)

// EventRequestBuilder builds and executes requests for operations under \me\calendarGroups\{calendarGroup-id}\calendars\{calendar-id}\calendarView\{event-id}
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
func (m *EventRequestBuilder) Accept()(*iaa58f3091eded3bb8f8022ff383adf6893ae04638de3510fb6f50e685897d70d.AcceptRequestBuilder) {
    return iaa58f3091eded3bb8f8022ff383adf6893ae04638de3510fb6f50e685897d70d.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*ib0aa93c8bad1e825bd5cdfdc4fa22e31b84830d4f2814d7e3ccfe19465129484.AttachmentsRequestBuilder) {
    return ib0aa93c8bad1e825bd5cdfdc4fa22e31b84830d4f2814d7e3ccfe19465129484.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.attachments.item collection
func (m *EventRequestBuilder) AttachmentsById(id string)(*ib8afe549ec5fafbb94854314181e9eac838b20474be31174184b3169d64d40a6.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ib8afe549ec5fafbb94854314181e9eac838b20474be31174184b3169d64d40a6.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i5986037676fe795074e810e4246ff753605940169fbbeed27b2aefa176b11751.CalendarRequestBuilder) {
    return i5986037676fe795074e810e4246ff753605940169fbbeed27b2aefa176b11751.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i3145dece064d967ee1c16e3e936c47ea8011ed2847cccb8dedaaa5c1eedaa013.CancelRequestBuilder) {
    return i3145dece064d967ee1c16e3e936c47ea8011ed2847cccb8dedaaa5c1eedaa013.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventRequestBuilderInternal instantiates a new EventRequestBuilder and sets the default values.
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup_id}/calendars/{calendar_id}/calendarView/{event_id}{?select}";
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
func (m *EventRequestBuilder) Decline()(*i88ecb5106cc0bb521ce42c551cf99da444f6f2720897e5c4464c18a087b9fa03.DeclineRequestBuilder) {
    return i88ecb5106cc0bb521ce42c551cf99da444f6f2720897e5c4464c18a087b9fa03.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i6caf8586c74aeb53bd8b7a86089cf9aa13e361aa8d0dd14f84486a74e16e0222.DismissReminderRequestBuilder) {
    return i6caf8586c74aeb53bd8b7a86089cf9aa13e361aa8d0dd14f84486a74e16e0222.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*ic99712d0b0d1392d14e91b9fb4b95688da8938561196adae18f98062d303b315.ExceptionOccurrencesRequestBuilder) {
    return ic99712d0b0d1392d14e91b9fb4b95688da8938561196adae18f98062d303b315.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*id97c351c3207d208122c8654eec8fa4fae5d90b3136fc2c79c3605d5a1fee25d.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return id97c351c3207d208122c8654eec8fa4fae5d90b3136fc2c79c3605d5a1fee25d.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*ie128b633259361fd2a945ba339cd6061e92d56f2cdc9422dd860b687dc19ebfb.ExtensionsRequestBuilder) {
    return ie128b633259361fd2a945ba339cd6061e92d56f2cdc9422dd860b687dc19ebfb.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.extensions.item collection
func (m *EventRequestBuilder) ExtensionsById(id string)(*i3e68bb94a10239ef0489773e8e20a5d33a576d00b16ded587e7df49c942e261a.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i3e68bb94a10239ef0489773e8e20a5d33a576d00b16ded587e7df49c942e261a.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ib53b9735ba043d622c998452495258595cc62dbf7581e5a3605dfc29007b5997.ForwardRequestBuilder) {
    return ib53b9735ba043d622c998452495258595cc62dbf7581e5a3605dfc29007b5997.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*i276be7fad31bd680ac1411035eef4f3a00776a60d1125d214d24b184f7469220.InstancesRequestBuilder) {
    return i276be7fad31bd680ac1411035eef4f3a00776a60d1125d214d24b184f7469220.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.instances.item collection
func (m *EventRequestBuilder) InstancesById(id string)(*id21bb72bf65597e6d902b0acc913993108ac302ca65b2d1f3c3fcaaada16028d.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return id21bb72bf65597e6d902b0acc913993108ac302ca65b2d1f3c3fcaaada16028d.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*iab57ad64a283353ed19ad1d9c7814f25835c7ab4a0d769767a6db218d4703ee7.MultiValueExtendedPropertiesRequestBuilder) {
    return iab57ad64a283353ed19ad1d9c7814f25835c7ab4a0d769767a6db218d4703ee7.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic9ed65a5fbed6446506612533b9a4560f521e68e07c8a6e70fc6eb495cf13716.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ic9ed65a5fbed6446506612533b9a4560f521e68e07c8a6e70fc6eb495cf13716.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*ic53f6b01f6b3da1d49fabe4f45413f9d94236e324fbc5f83450593600644550a.SingleValueExtendedPropertiesRequestBuilder) {
    return ic53f6b01f6b3da1d49fabe4f45413f9d94236e324fbc5f83450593600644550a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if2b4feb3f0e58d1cd206b56ee504e531df602a972bcdd86373ae72f7f5884e6f.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if2b4feb3f0e58d1cd206b56ee504e531df602a972bcdd86373ae72f7f5884e6f.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i88425e8a17893a7e305a95729332a64c30460d352b0bf5845f5d7090f92c8f2b.SnoozeReminderRequestBuilder) {
    return i88425e8a17893a7e305a95729332a64c30460d352b0bf5845f5d7090f92c8f2b.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ia22ed1a7b043dc5283f91ddeb06193ce26ef17e32a8a4abd851e6155127ed218.TentativelyAcceptRequestBuilder) {
    return ia22ed1a7b043dc5283f91ddeb06193ce26ef17e32a8a4abd851e6155127ed218.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
