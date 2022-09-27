package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i12f6208fa635d862ebcbb2efc3fc2ad520e8b11af30ce47c1b7aae81087c3fc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/cancel"
    i24648f3b72cb6913a9588b8f30f65a5861a94ceec8a08eafb72d128c878c149c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances"
    i51f82170e9d7e02f39adef69c5ad83f6c68a8391a942267ee25404779c09df02 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/multivalueextendedproperties"
    i689722744aa2c7a040ce1eea267991c6a8df229fedaf72403c3cd9fd504dc586 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/forward"
    i6a81bfdd259fd587f1bd56d9b047bb56371d6496c53c194845624ba2a958a171 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/snoozereminder"
    i6ccf1ed39358cfe85d5d5e1ed7cc04329f170b3143984349ea66a828df15091d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/attachments"
    i7ea3353b3d2a5a1c8706154adf849e0629e8535acc87cb16c1b741d012b14551 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/calendar"
    i990c9d56acce6df7a509f9f537bf1c3651710e2d54e9785ef179a00a350d7624 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/singlevalueextendedproperties"
    i9f03e8930612bb38b0ee316af407e0390228dbd9c1ce9d5b920322626d4a7ccb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/dismissreminder"
    ia519bb57124ce78ce804d4c8fde3bad9e5de6768daf6ccd8b683b306570d66b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences"
    ib35ba17b740b364bef12429dc54a7fc3e5647944b9c94a261b2afb41557bdedb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/tentativelyaccept"
    id0b046c29ca78c81c5fc1cdab2b1e6aad652b41020d22dd887163e374abd4714 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/accept"
    idf4c71a10741a39450889d693f4932da9fbbd7fd71fa3872d803f21687c7eeb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/decline"
    ieb846c3f403e62d9249e7d035caf7a3aae7a2c964e81219a2fdbf7417e1f6fbb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/extensions"
    i2ed60d96750ee240e18770a7ac42e716a5f27688d6f8e7757c055b8d4564d81c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item"
    i3ab7e30c83492e2917ff34bdeecc0fca059eeb741d94af5ebf45410534e47d9f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/attachments/item"
    i7a7c70eae23ab3508021e2f69db1d970a3bd21e3036bbb44f0a1291701a6d4da "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/singlevalueextendedproperties/item"
    ia1b7d92a7d10b8e6e611cc5be140d3a64be2800b960d7d8781caf1650e2f0c08 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item"
    ib2a745f0d686f44f5e53733974718efc67fcdb195c1f77b50aa122d3f058cb4e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/multivalueextendedproperties/item"
    id96b14704d4aef544449833c6800c5d3620733e0995ae628d5280e53eed51574 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*id0b046c29ca78c81c5fc1cdab2b1e6aad652b41020d22dd887163e374abd4714.AcceptRequestBuilder) {
    return id0b046c29ca78c81c5fc1cdab2b1e6aad652b41020d22dd887163e374abd4714.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i6ccf1ed39358cfe85d5d5e1ed7cc04329f170b3143984349ea66a828df15091d.AttachmentsRequestBuilder) {
    return i6ccf1ed39358cfe85d5d5e1ed7cc04329f170b3143984349ea66a828df15091d.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i3ab7e30c83492e2917ff34bdeecc0fca059eeb741d94af5ebf45410534e47d9f.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i3ab7e30c83492e2917ff34bdeecc0fca059eeb741d94af5ebf45410534e47d9f.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i7ea3353b3d2a5a1c8706154adf849e0629e8535acc87cb16c1b741d012b14551.CalendarRequestBuilder) {
    return i7ea3353b3d2a5a1c8706154adf849e0629e8535acc87cb16c1b741d012b14551.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i12f6208fa635d862ebcbb2efc3fc2ad520e8b11af30ce47c1b7aae81087c3fc4.CancelRequestBuilder) {
    return i12f6208fa635d862ebcbb2efc3fc2ad520e8b11af30ce47c1b7aae81087c3fc4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/calendarView/{event%2Did}{?startDateTime*,endDateTime*,%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*idf4c71a10741a39450889d693f4932da9fbbd7fd71fa3872d803f21687c7eeb0.DeclineRequestBuilder) {
    return idf4c71a10741a39450889d693f4932da9fbbd7fd71fa3872d803f21687c7eeb0.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i9f03e8930612bb38b0ee316af407e0390228dbd9c1ce9d5b920322626d4a7ccb.DismissReminderRequestBuilder) {
    return i9f03e8930612bb38b0ee316af407e0390228dbd9c1ce9d5b920322626d4a7ccb.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ia519bb57124ce78ce804d4c8fde3bad9e5de6768daf6ccd8b683b306570d66b9.ExceptionOccurrencesRequestBuilder) {
    return ia519bb57124ce78ce804d4c8fde3bad9e5de6768daf6ccd8b683b306570d66b9.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i2ed60d96750ee240e18770a7ac42e716a5f27688d6f8e7757c055b8d4564d81c.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i2ed60d96750ee240e18770a7ac42e716a5f27688d6f8e7757c055b8d4564d81c.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ieb846c3f403e62d9249e7d035caf7a3aae7a2c964e81219a2fdbf7417e1f6fbb.ExtensionsRequestBuilder) {
    return ieb846c3f403e62d9249e7d035caf7a3aae7a2c964e81219a2fdbf7417e1f6fbb.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*id96b14704d4aef544449833c6800c5d3620733e0995ae628d5280e53eed51574.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return id96b14704d4aef544449833c6800c5d3620733e0995ae628d5280e53eed51574.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i689722744aa2c7a040ce1eea267991c6a8df229fedaf72403c3cd9fd504dc586.ForwardRequestBuilder) {
    return i689722744aa2c7a040ce1eea267991c6a8df229fedaf72403c3cd9fd504dc586.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*i24648f3b72cb6913a9588b8f30f65a5861a94ceec8a08eafb72d128c878c149c.InstancesRequestBuilder) {
    return i24648f3b72cb6913a9588b8f30f65a5861a94ceec8a08eafb72d128c878c149c.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ia1b7d92a7d10b8e6e611cc5be140d3a64be2800b960d7d8781caf1650e2f0c08.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return ia1b7d92a7d10b8e6e611cc5be140d3a64be2800b960d7d8781caf1650e2f0c08.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i51f82170e9d7e02f39adef69c5ad83f6c68a8391a942267ee25404779c09df02.MultiValueExtendedPropertiesRequestBuilder) {
    return i51f82170e9d7e02f39adef69c5ad83f6c68a8391a942267ee25404779c09df02.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib2a745f0d686f44f5e53733974718efc67fcdb195c1f77b50aa122d3f058cb4e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ib2a745f0d686f44f5e53733974718efc67fcdb195c1f77b50aa122d3f058cb4e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i990c9d56acce6df7a509f9f537bf1c3651710e2d54e9785ef179a00a350d7624.SingleValueExtendedPropertiesRequestBuilder) {
    return i990c9d56acce6df7a509f9f537bf1c3651710e2d54e9785ef179a00a350d7624.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7a7c70eae23ab3508021e2f69db1d970a3bd21e3036bbb44f0a1291701a6d4da.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i7a7c70eae23ab3508021e2f69db1d970a3bd21e3036bbb44f0a1291701a6d4da.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i6a81bfdd259fd587f1bd56d9b047bb56371d6496c53c194845624ba2a958a171.SnoozeReminderRequestBuilder) {
    return i6a81bfdd259fd587f1bd56d9b047bb56371d6496c53c194845624ba2a958a171.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ib35ba17b740b364bef12429dc54a7fc3e5647944b9c94a261b2afb41557bdedb.TentativelyAcceptRequestBuilder) {
    return ib35ba17b740b364bef12429dc54a7fc3e5647944b9c94a261b2afb41557bdedb.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
