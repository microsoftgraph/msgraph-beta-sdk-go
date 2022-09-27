package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09464dfde46ea9b5614331a441462441658822ee493fbccebc73230d1194d842 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences"
    i0ae79ce850436f0587525bcaec70b835a2cef9711988a12de9cb93e24414b430 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/snoozereminder"
    i0b099560e4b2efb93ab572b92e7f1a76ec97b902d61ab9becc2defcef681ac25 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/dismissreminder"
    i2bcbda220eee3efb3381726cc4041421ea7e20a03d0304ac407dea06fa354f27 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/singlevalueextendedproperties"
    i4b428c5cf9a69e44aa972695ae72775d3e28ccc810fe60f6b821b483200c436d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances"
    i5debd82ad007fe3e7b83a2eeb29c28f5da3b9563dd9ca463f34104a008b9eee9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/calendar"
    i88e3730e4bce338f224d440e746456a0359fb801e5e58784dbfe1799783810bf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/forward"
    i8c425fcf68846ce1fd43b66e601e98461e65f7444df3be897e0da2f43a8df63c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/cancel"
    i9dd05f797cf579bc4dc0e79bc00cc79d1351d71707d91b30098d9d3ae9ea28b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/attachments"
    iac7fd37b42a89ce1966b5f0ea72249cc45d99faf64966f63c5ebe88241edbbe9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/extensions"
    iaf2491bc5b9ff1ab39daeb0d781b18170fe54aa51fb51480ae6422705aa157c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/accept"
    ic31b3a2cd538f90211d1e9cf6de4bb7092f7aacee83586f1ab97621d45d8b815 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/decline"
    id6ac304b8cd98e6aec12a2df7396b7ef53304b70df9d56b58094a06bb593da74 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/multivalueextendedproperties"
    ifd62164a910c345faf910bed37d632b88189cfa74bdd65db738cf2f49dabf4d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/tentativelyaccept"
    i0f9b7621256a1b0f4c88adedcc6a9f1bbfd62cef422da8242f344ebe0128feda "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/extensions/item"
    i47e3e498cae25dc161c5f4d25d399c625874e2e39d1692d4216a0b825941c4ed "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/multivalueextendedproperties/item"
    i91c7117765263d0e129264c87641bfc35cd44799a05cde4e3181a86d7ab46ccd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item"
    ib0b41a465e6f0cb684c983aa2ebdcc8e6295fb076929a3ad3e741eb6fc148feb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/singlevalueextendedproperties/item"
    ib21fde6c453d56158f7bf6413ec85dcf8c8b15c7d91c031a7cccedeacb0d80c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item"
    ief3b9e1dfba882027be7bcb7ecb89490b77e45266ac9933e953debb528098d78 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
type EventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string
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
func (m *EventItemRequestBuilder) Accept()(*iaf2491bc5b9ff1ab39daeb0d781b18170fe54aa51fb51480ae6422705aa157c8.AcceptRequestBuilder) {
    return iaf2491bc5b9ff1ab39daeb0d781b18170fe54aa51fb51480ae6422705aa157c8.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i9dd05f797cf579bc4dc0e79bc00cc79d1351d71707d91b30098d9d3ae9ea28b5.AttachmentsRequestBuilder) {
    return i9dd05f797cf579bc4dc0e79bc00cc79d1351d71707d91b30098d9d3ae9ea28b5.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ief3b9e1dfba882027be7bcb7ecb89490b77e45266ac9933e953debb528098d78.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ief3b9e1dfba882027be7bcb7ecb89490b77e45266ac9933e953debb528098d78.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i5debd82ad007fe3e7b83a2eeb29c28f5da3b9563dd9ca463f34104a008b9eee9.CalendarRequestBuilder) {
    return i5debd82ad007fe3e7b83a2eeb29c28f5da3b9563dd9ca463f34104a008b9eee9.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i8c425fcf68846ce1fd43b66e601e98461e65f7444df3be897e0da2f43a8df63c.CancelRequestBuilder) {
    return i8c425fcf68846ce1fd43b66e601e98461e65f7444df3be897e0da2f43a8df63c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}{?startDateTime*,endDateTime*,%24select}";
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Decline()(*ic31b3a2cd538f90211d1e9cf6de4bb7092f7aacee83586f1ab97621d45d8b815.DeclineRequestBuilder) {
    return ic31b3a2cd538f90211d1e9cf6de4bb7092f7aacee83586f1ab97621d45d8b815.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i0b099560e4b2efb93ab572b92e7f1a76ec97b902d61ab9becc2defcef681ac25.DismissReminderRequestBuilder) {
    return i0b099560e4b2efb93ab572b92e7f1a76ec97b902d61ab9becc2defcef681ac25.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i09464dfde46ea9b5614331a441462441658822ee493fbccebc73230d1194d842.ExceptionOccurrencesRequestBuilder) {
    return i09464dfde46ea9b5614331a441462441658822ee493fbccebc73230d1194d842.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i91c7117765263d0e129264c87641bfc35cd44799a05cde4e3181a86d7ab46ccd.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i91c7117765263d0e129264c87641bfc35cd44799a05cde4e3181a86d7ab46ccd.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*iac7fd37b42a89ce1966b5f0ea72249cc45d99faf64966f63c5ebe88241edbbe9.ExtensionsRequestBuilder) {
    return iac7fd37b42a89ce1966b5f0ea72249cc45d99faf64966f63c5ebe88241edbbe9.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i0f9b7621256a1b0f4c88adedcc6a9f1bbfd62cef422da8242f344ebe0128feda.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i0f9b7621256a1b0f4c88adedcc6a9f1bbfd62cef422da8242f344ebe0128feda.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i88e3730e4bce338f224d440e746456a0359fb801e5e58784dbfe1799783810bf.ForwardRequestBuilder) {
    return i88e3730e4bce338f224d440e746456a0359fb801e5e58784dbfe1799783810bf.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
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
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*i4b428c5cf9a69e44aa972695ae72775d3e28ccc810fe60f6b821b483200c436d.InstancesRequestBuilder) {
    return i4b428c5cf9a69e44aa972695ae72775d3e28ccc810fe60f6b821b483200c436d.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ib21fde6c453d56158f7bf6413ec85dcf8c8b15c7d91c031a7cccedeacb0d80c9.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return ib21fde6c453d56158f7bf6413ec85dcf8c8b15c7d91c031a7cccedeacb0d80c9.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*id6ac304b8cd98e6aec12a2df7396b7ef53304b70df9d56b58094a06bb593da74.MultiValueExtendedPropertiesRequestBuilder) {
    return id6ac304b8cd98e6aec12a2df7396b7ef53304b70df9d56b58094a06bb593da74.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i47e3e498cae25dc161c5f4d25d399c625874e2e39d1692d4216a0b825941c4ed.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i47e3e498cae25dc161c5f4d25d399c625874e2e39d1692d4216a0b825941c4ed.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i2bcbda220eee3efb3381726cc4041421ea7e20a03d0304ac407dea06fa354f27.SingleValueExtendedPropertiesRequestBuilder) {
    return i2bcbda220eee3efb3381726cc4041421ea7e20a03d0304ac407dea06fa354f27.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib0b41a465e6f0cb684c983aa2ebdcc8e6295fb076929a3ad3e741eb6fc148feb.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ib0b41a465e6f0cb684c983aa2ebdcc8e6295fb076929a3ad3e741eb6fc148feb.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i0ae79ce850436f0587525bcaec70b835a2cef9711988a12de9cb93e24414b430.SnoozeReminderRequestBuilder) {
    return i0ae79ce850436f0587525bcaec70b835a2cef9711988a12de9cb93e24414b430.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ifd62164a910c345faf910bed37d632b88189cfa74bdd65db738cf2f49dabf4d1.TentativelyAcceptRequestBuilder) {
    return ifd62164a910c345faf910bed37d632b88189cfa74bdd65db738cf2f49dabf4d1.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
