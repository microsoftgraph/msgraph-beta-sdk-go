package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0b9f80f8fac79c3a9cf2e537f859b1c2e21df1fa7ae49ddb0e041c048e2a9b92 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    i1ed714b778d8b7d2dea09689275b50005e4d6639baaec71a0dbe9aa6adc2f31f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/accept"
    i22dbf9f458bfeb5e2a193184182a4c810cd7d72886f8040446e2625460755d45 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    i4895c52c8b3aef844d9b277586d6bd877f0354270c0fb4615e2752e505d9435d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    i53e0452291f095707b5942cca00bac9131e66dd37adfcea3833107ea99f4b1ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i77f20c4b4594205b842d02086d88fd51932204c7169f06e344b6189cb87d468a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    ia22b68672e63e164db65f2d6433cd00e6dfef6588859a76f352301ecbade573c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    ic2577072b12a75e0f45116330ffea1ed21cd5ce0880c405b313b87ef925038e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    icde9fc17d59fb2667ba21b41bb23a93f3e7ef492ef80db24a749927f783187a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    id87abc09906ded7d7954d4fd07f8614f5ef2e06fe29a2001e0aad0a02e6cd304 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/forward"
    ie01b911721cc06fe9220cc6e311e48277d2c5d1e83c2f85c954f45d058023887 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    ieb053bc43aed21e38a0dc172dfa07fcdf56eb34b1a45324f6429b41778dcc174 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/decline"
    i1b58f0aecd662360ac86db0cb8d5dc6847453e51a07519e2eaeecf55c28f53c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    i5787a60d94ced5c7eacb9c5b114c730d82e74fe1740d214aaeeb47a39a2fe71d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    i8e3cf8a70d3ff34d13e9fb27e223a5ac2f39b2dc223bf83bc8e66b97e1132129 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    i9349f129ad1cf4d0c89e359bfda1097c31ea5c39cbc03c7c1aa4b773be138b31 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i1ed714b778d8b7d2dea09689275b50005e4d6639baaec71a0dbe9aa6adc2f31f.AcceptRequestBuilder) {
    return i1ed714b778d8b7d2dea09689275b50005e4d6639baaec71a0dbe9aa6adc2f31f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i4895c52c8b3aef844d9b277586d6bd877f0354270c0fb4615e2752e505d9435d.AttachmentsRequestBuilder) {
    return i4895c52c8b3aef844d9b277586d6bd877f0354270c0fb4615e2752e505d9435d.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i5787a60d94ced5c7eacb9c5b114c730d82e74fe1740d214aaeeb47a39a2fe71d.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i5787a60d94ced5c7eacb9c5b114c730d82e74fe1740d214aaeeb47a39a2fe71d.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i77f20c4b4594205b842d02086d88fd51932204c7169f06e344b6189cb87d468a.CalendarRequestBuilder) {
    return i77f20c4b4594205b842d02086d88fd51932204c7169f06e344b6189cb87d468a.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i22dbf9f458bfeb5e2a193184182a4c810cd7d72886f8040446e2625460755d45.CancelRequestBuilder) {
    return i22dbf9f458bfeb5e2a193184182a4c810cd7d72886f8040446e2625460755d45.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*ieb053bc43aed21e38a0dc172dfa07fcdf56eb34b1a45324f6429b41778dcc174.DeclineRequestBuilder) {
    return ieb053bc43aed21e38a0dc172dfa07fcdf56eb34b1a45324f6429b41778dcc174.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*ia22b68672e63e164db65f2d6433cd00e6dfef6588859a76f352301ecbade573c.DismissReminderRequestBuilder) {
    return ia22b68672e63e164db65f2d6433cd00e6dfef6588859a76f352301ecbade573c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i0b9f80f8fac79c3a9cf2e537f859b1c2e21df1fa7ae49ddb0e041c048e2a9b92.ExtensionsRequestBuilder) {
    return i0b9f80f8fac79c3a9cf2e537f859b1c2e21df1fa7ae49ddb0e041c048e2a9b92.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i8e3cf8a70d3ff34d13e9fb27e223a5ac2f39b2dc223bf83bc8e66b97e1132129.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8e3cf8a70d3ff34d13e9fb27e223a5ac2f39b2dc223bf83bc8e66b97e1132129.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*id87abc09906ded7d7954d4fd07f8614f5ef2e06fe29a2001e0aad0a02e6cd304.ForwardRequestBuilder) {
    return id87abc09906ded7d7954d4fd07f8614f5ef2e06fe29a2001e0aad0a02e6cd304.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic2577072b12a75e0f45116330ffea1ed21cd5ce0880c405b313b87ef925038e6.MultiValueExtendedPropertiesRequestBuilder) {
    return ic2577072b12a75e0f45116330ffea1ed21cd5ce0880c405b313b87ef925038e6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i9349f129ad1cf4d0c89e359bfda1097c31ea5c39cbc03c7c1aa4b773be138b31.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i9349f129ad1cf4d0c89e359bfda1097c31ea5c39cbc03c7c1aa4b773be138b31.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i53e0452291f095707b5942cca00bac9131e66dd37adfcea3833107ea99f4b1ef.SingleValueExtendedPropertiesRequestBuilder) {
    return i53e0452291f095707b5942cca00bac9131e66dd37adfcea3833107ea99f4b1ef.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i1b58f0aecd662360ac86db0cb8d5dc6847453e51a07519e2eaeecf55c28f53c6.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i1b58f0aecd662360ac86db0cb8d5dc6847453e51a07519e2eaeecf55c28f53c6.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ie01b911721cc06fe9220cc6e311e48277d2c5d1e83c2f85c954f45d058023887.SnoozeReminderRequestBuilder) {
    return ie01b911721cc06fe9220cc6e311e48277d2c5d1e83c2f85c954f45d058023887.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*icde9fc17d59fb2667ba21b41bb23a93f3e7ef492ef80db24a749927f783187a3.TentativelyAcceptRequestBuilder) {
    return icde9fc17d59fb2667ba21b41bb23a93f3e7ef492ef80db24a749927f783187a3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
