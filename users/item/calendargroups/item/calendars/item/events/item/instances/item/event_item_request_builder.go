package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3c30ca65e90b69ba0281d155d28471459339d9357416e7fa20920dd7fb53a46a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/tentativelyaccept"
    i5bdead2970e6a3e7ea3d7996bbf16c058395ab6603f2bf38685a695ee9467e42 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/forward"
    i640521ec093bc0556b892715f6a7fc2b6bae71b6d513a04cee8ffdf63dbf4790 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/cancel"
    i6cdb6139b308ed2eb2f2859f2aafbb7d7b9f5ff985fb89608fae6a0a6f5cab4d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/decline"
    i821e450cede61f6adabc700b0976fb2dc28b82ae1e5897c462bb508f1f8607ca "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/extensions"
    i8a3ac25440a21e51394e0ef569129148ece7381700c78f125ab59921249c98f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/snoozereminder"
    i9649ae8562ffc8d27234cae2bf408ba6fa96b6176d520e1b11ec6320ed693c34 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/attachments"
    i9e4faf44eb040d129716df670d72044b9aeff37a4af4bcae7e726350f24d4ca9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/multivalueextendedproperties"
    ia77ff343f97c45f5fcaf79c67210b4f781431345cec76aff6739d5427cb472a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/accept"
    ib6854270cc64cdc03fe7f48f7ca1c6b2b509b32e4e45cd071e58dbe31d35e3d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/singlevalueextendedproperties"
    ib83d6d6607466b314487e87fb023466e1b1eeb293740d36b806095d251938966 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences"
    ie6a268da82c710cce50997de3860f32faaf552a060eb4f2c4a341bc6add64f97 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/calendar"
    if5345c5ac4ac3fc447471af3ae043eec0cfbbf5c679743efde76993702848ef0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/dismissreminder"
    i46d71118fd91d57f6c18c845a5d9c1ea6e1333b1e733cf06869c815e49726092 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/attachments/item"
    i6e7a935205120176c42da76e2537e59883049561e9c277d8f7066b98d91dd3dc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item"
    i77f001ef271ad08c69a5081d7babf6eb1e60cd5c5dfc749d82359fdfc0d55c1f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/singlevalueextendedproperties/item"
    ida3021ac9c0588e0cfb7c1fad1b655eae5a139e4b7e99b3a892bc6a98126fa73 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/extensions/item"
    ie117fe513f75824907279a273dab72b4749d2c9cf060bf6c318e06880e623116 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*ia77ff343f97c45f5fcaf79c67210b4f781431345cec76aff6739d5427cb472a5.AcceptRequestBuilder) {
    return ia77ff343f97c45f5fcaf79c67210b4f781431345cec76aff6739d5427cb472a5.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i9649ae8562ffc8d27234cae2bf408ba6fa96b6176d520e1b11ec6320ed693c34.AttachmentsRequestBuilder) {
    return i9649ae8562ffc8d27234cae2bf408ba6fa96b6176d520e1b11ec6320ed693c34.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i46d71118fd91d57f6c18c845a5d9c1ea6e1333b1e733cf06869c815e49726092.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i46d71118fd91d57f6c18c845a5d9c1ea6e1333b1e733cf06869c815e49726092.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ie6a268da82c710cce50997de3860f32faaf552a060eb4f2c4a341bc6add64f97.CalendarRequestBuilder) {
    return ie6a268da82c710cce50997de3860f32faaf552a060eb4f2c4a341bc6add64f97.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i640521ec093bc0556b892715f6a7fc2b6bae71b6d513a04cee8ffdf63dbf4790.CancelRequestBuilder) {
    return i640521ec093bc0556b892715f6a7fc2b6bae71b6d513a04cee8ffdf63dbf4790.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i6cdb6139b308ed2eb2f2859f2aafbb7d7b9f5ff985fb89608fae6a0a6f5cab4d.DeclineRequestBuilder) {
    return i6cdb6139b308ed2eb2f2859f2aafbb7d7b9f5ff985fb89608fae6a0a6f5cab4d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*if5345c5ac4ac3fc447471af3ae043eec0cfbbf5c679743efde76993702848ef0.DismissReminderRequestBuilder) {
    return if5345c5ac4ac3fc447471af3ae043eec0cfbbf5c679743efde76993702848ef0.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ib83d6d6607466b314487e87fb023466e1b1eeb293740d36b806095d251938966.ExceptionOccurrencesRequestBuilder) {
    return ib83d6d6607466b314487e87fb023466e1b1eeb293740d36b806095d251938966.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i6e7a935205120176c42da76e2537e59883049561e9c277d8f7066b98d91dd3dc.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i6e7a935205120176c42da76e2537e59883049561e9c277d8f7066b98d91dd3dc.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i821e450cede61f6adabc700b0976fb2dc28b82ae1e5897c462bb508f1f8607ca.ExtensionsRequestBuilder) {
    return i821e450cede61f6adabc700b0976fb2dc28b82ae1e5897c462bb508f1f8607ca.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ida3021ac9c0588e0cfb7c1fad1b655eae5a139e4b7e99b3a892bc6a98126fa73.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ida3021ac9c0588e0cfb7c1fad1b655eae5a139e4b7e99b3a892bc6a98126fa73.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i5bdead2970e6a3e7ea3d7996bbf16c058395ab6603f2bf38685a695ee9467e42.ForwardRequestBuilder) {
    return i5bdead2970e6a3e7ea3d7996bbf16c058395ab6603f2bf38685a695ee9467e42.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i9e4faf44eb040d129716df670d72044b9aeff37a4af4bcae7e726350f24d4ca9.MultiValueExtendedPropertiesRequestBuilder) {
    return i9e4faf44eb040d129716df670d72044b9aeff37a4af4bcae7e726350f24d4ca9.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ie117fe513f75824907279a273dab72b4749d2c9cf060bf6c318e06880e623116.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ie117fe513f75824907279a273dab72b4749d2c9cf060bf6c318e06880e623116.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ib6854270cc64cdc03fe7f48f7ca1c6b2b509b32e4e45cd071e58dbe31d35e3d8.SingleValueExtendedPropertiesRequestBuilder) {
    return ib6854270cc64cdc03fe7f48f7ca1c6b2b509b32e4e45cd071e58dbe31d35e3d8.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i77f001ef271ad08c69a5081d7babf6eb1e60cd5c5dfc749d82359fdfc0d55c1f.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i77f001ef271ad08c69a5081d7babf6eb1e60cd5c5dfc749d82359fdfc0d55c1f.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i8a3ac25440a21e51394e0ef569129148ece7381700c78f125ab59921249c98f7.SnoozeReminderRequestBuilder) {
    return i8a3ac25440a21e51394e0ef569129148ece7381700c78f125ab59921249c98f7.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3c30ca65e90b69ba0281d155d28471459339d9357416e7fa20920dd7fb53a46a.TentativelyAcceptRequestBuilder) {
    return i3c30ca65e90b69ba0281d155d28471459339d9357416e7fa20920dd7fb53a46a.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
