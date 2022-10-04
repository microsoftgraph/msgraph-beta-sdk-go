package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09f5dfd1a5df4250c4f4913a6d0e0c4df6021f4405e7daec3f75047bdaafc0b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences"
    i1f783971eac6488869345e0e6a5eb53ea7eab44c20658ac485cfa0b3a2b3c0c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/dismissreminder"
    i2915adc28448f5ea89a2c857fadf46dccfe589ce01919fb14d71512a9c1a5eac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/accept"
    i360706ada10a10e0ac8a34a18afa8a17dcb9ee384d751c270271f83088b05561 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/forward"
    i61a6287eb18c020d67a5548a580ee2d40f023308d9c9f67db1f1069ab1ecd8dd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/extensions"
    i8718987829e1f6e3a6434f466a8e49143e850fa9085ed5c172f0e1e8ff3e111b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/tentativelyaccept"
    ib0f3a5424fc6febcc130148818f2725ec5d8c4ae4fd7aa9c8e27189798348fc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/snoozereminder"
    ib0f51dbb6b0208bb71d0cc1c76bed069accaad0cc61aa5478448eeed29025e68 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/attachments"
    ibe58f8d597afefd0f759b866eb299cd108c2bd87094a52ea5a8cb2a495fd6c69 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/singlevalueextendedproperties"
    ie104a25b434ba5b5e09cefb9b47bf7e245461481e8923c6fd32c4dccafc6dd14 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/calendar"
    ie73b70cbdef3146e2c7c8bfd4a6631be2dabe6f13ef2dba00ab326cf4c3ab35b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/decline"
    iee1de0441f81825e304e53df19414a9f4f9cff20abaf2850a31ef3587dfc238e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/multivalueextendedproperties"
    if44b8b29f72e9409df216d6626d1a53d36959d97b7eb6bb85095f4ba5699bc49 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/cancel"
    i4022809c8375537f7a834e6923cdf25c51595bfdb121d476c9660943ea6e4a88 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/extensions/item"
    i6c72d4ee462949237fab5c8077331dc5e0a723652747ffaf7704ccdc7e285152 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/multivalueextendedproperties/item"
    ia82de5bf8ab9ec6a33f9da498937d5bbff4016e9d593ba2b192ed7198b814784 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/singlevalueextendedproperties/item"
    ia8dbf3b16458048333cca86618e5ed994bcaca9ce9b3eb23800113e6d52f03a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item"
    ie6931a9993174abe9d7da89d161347de53e1a2d1399aa606b845504bf863ebce "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*i2915adc28448f5ea89a2c857fadf46dccfe589ce01919fb14d71512a9c1a5eac.AcceptRequestBuilder) {
    return i2915adc28448f5ea89a2c857fadf46dccfe589ce01919fb14d71512a9c1a5eac.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ib0f51dbb6b0208bb71d0cc1c76bed069accaad0cc61aa5478448eeed29025e68.AttachmentsRequestBuilder) {
    return ib0f51dbb6b0208bb71d0cc1c76bed069accaad0cc61aa5478448eeed29025e68.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ie6931a9993174abe9d7da89d161347de53e1a2d1399aa606b845504bf863ebce.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ie6931a9993174abe9d7da89d161347de53e1a2d1399aa606b845504bf863ebce.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ie104a25b434ba5b5e09cefb9b47bf7e245461481e8923c6fd32c4dccafc6dd14.CalendarRequestBuilder) {
    return ie104a25b434ba5b5e09cefb9b47bf7e245461481e8923c6fd32c4dccafc6dd14.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*if44b8b29f72e9409df216d6626d1a53d36959d97b7eb6bb85095f4ba5699bc49.CancelRequestBuilder) {
    return if44b8b29f72e9409df216d6626d1a53d36959d97b7eb6bb85095f4ba5699bc49.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EventItemRequestBuilder) Decline()(*ie73b70cbdef3146e2c7c8bfd4a6631be2dabe6f13ef2dba00ab326cf4c3ab35b.DeclineRequestBuilder) {
    return ie73b70cbdef3146e2c7c8bfd4a6631be2dabe6f13ef2dba00ab326cf4c3ab35b.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i1f783971eac6488869345e0e6a5eb53ea7eab44c20658ac485cfa0b3a2b3c0c2.DismissReminderRequestBuilder) {
    return i1f783971eac6488869345e0e6a5eb53ea7eab44c20658ac485cfa0b3a2b3c0c2.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i09f5dfd1a5df4250c4f4913a6d0e0c4df6021f4405e7daec3f75047bdaafc0b6.ExceptionOccurrencesRequestBuilder) {
    return i09f5dfd1a5df4250c4f4913a6d0e0c4df6021f4405e7daec3f75047bdaafc0b6.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*ia8dbf3b16458048333cca86618e5ed994bcaca9ce9b3eb23800113e6d52f03a8.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return ia8dbf3b16458048333cca86618e5ed994bcaca9ce9b3eb23800113e6d52f03a8.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i61a6287eb18c020d67a5548a580ee2d40f023308d9c9f67db1f1069ab1ecd8dd.ExtensionsRequestBuilder) {
    return i61a6287eb18c020d67a5548a580ee2d40f023308d9c9f67db1f1069ab1ecd8dd.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i4022809c8375537f7a834e6923cdf25c51595bfdb121d476c9660943ea6e4a88.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i4022809c8375537f7a834e6923cdf25c51595bfdb121d476c9660943ea6e4a88.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i360706ada10a10e0ac8a34a18afa8a17dcb9ee384d751c270271f83088b05561.ForwardRequestBuilder) {
    return i360706ada10a10e0ac8a34a18afa8a17dcb9ee384d751c270271f83088b05561.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iee1de0441f81825e304e53df19414a9f4f9cff20abaf2850a31ef3587dfc238e.MultiValueExtendedPropertiesRequestBuilder) {
    return iee1de0441f81825e304e53df19414a9f4f9cff20abaf2850a31ef3587dfc238e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i6c72d4ee462949237fab5c8077331dc5e0a723652747ffaf7704ccdc7e285152.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i6c72d4ee462949237fab5c8077331dc5e0a723652747ffaf7704ccdc7e285152.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ibe58f8d597afefd0f759b866eb299cd108c2bd87094a52ea5a8cb2a495fd6c69.SingleValueExtendedPropertiesRequestBuilder) {
    return ibe58f8d597afefd0f759b866eb299cd108c2bd87094a52ea5a8cb2a495fd6c69.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia82de5bf8ab9ec6a33f9da498937d5bbff4016e9d593ba2b192ed7198b814784.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ia82de5bf8ab9ec6a33f9da498937d5bbff4016e9d593ba2b192ed7198b814784.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ib0f3a5424fc6febcc130148818f2725ec5d8c4ae4fd7aa9c8e27189798348fc0.SnoozeReminderRequestBuilder) {
    return ib0f3a5424fc6febcc130148818f2725ec5d8c4ae4fd7aa9c8e27189798348fc0.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i8718987829e1f6e3a6434f466a8e49143e850fa9085ed5c172f0e1e8ff3e111b.TentativelyAcceptRequestBuilder) {
    return i8718987829e1f6e3a6434f466a8e49143e850fa9085ed5c172f0e1e8ff3e111b.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
