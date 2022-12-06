package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i33c237b80872e74abd5582ee9e98c570c7e6f02883241a95a3c88ae46b984552 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i36f8c02fbc92e3e0caa5522773d18daf2beeefb8505d123eee62e3dc03b8da7d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/accept"
    i47f85055d10bbc2ce5596ccc6b1cf375e39ff7636bbd2a8a85a26b04c8ec760b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/forward"
    i4b92cda71aa39cf43f1f4117d98af10c1b939188ac1f64900b459cbc86ce63d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/decline"
    i5da7820459f807892d24ff65fd96a91f4aac050c463293da376172f19b75a513 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i65f1f92552ffb48c728459d8e357479ff8ec797e294da0afafc24659ef8c3c24 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    i6c8729e6fa598d8c1a1def789a86190e6c425f33d3adeb1666d815fbf8521c04 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i92c1eb63d4144be862f92fc808e1688959b28b6423b622f3cfccb057eddfba26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    i940bde31fa32e362016c8393cc476a6f0a102a2ba874a3f21c82490cfb412bd4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    ic74c15eb17958e3d5aedbd7aa4e0d9baf8d5434a1f33184dc193d8e698f2d520 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    icd2d22f9a6b46c7b3ed6518f51fc18003eea511e431f001ec72a8d90370868b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    ie5cdec5f899951505bb7ff57572af18941b64136ff8cc51747b7c00af4ba5448 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    i0a6f89794fdbde0788f0c2ff964c175f2d1e14b59909b07026acceeaf74bbb15 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    i5ccfa8d8cf0255bec3dfff58300167af470e01e32723b254a929e2c54b59eada "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    i97d84954950d52b00814176753ad6fce764f2a130bdb4c3f31de8904563dcb49 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    ibf13cd72807bb3c8eb9ac05dcacf04c61cdd52496d8431cfdb6fbcd120da847f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
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
// Accept provides operations to call the accept method.
func (m *EventItemRequestBuilder) Accept()(*i36f8c02fbc92e3e0caa5522773d18daf2beeefb8505d123eee62e3dc03b8da7d.AcceptRequestBuilder) {
    return i36f8c02fbc92e3e0caa5522773d18daf2beeefb8505d123eee62e3dc03b8da7d.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*ie5cdec5f899951505bb7ff57572af18941b64136ff8cc51747b7c00af4ba5448.AttachmentsRequestBuilder) {
    return ie5cdec5f899951505bb7ff57572af18941b64136ff8cc51747b7c00af4ba5448.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i5ccfa8d8cf0255bec3dfff58300167af470e01e32723b254a929e2c54b59eada.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i5ccfa8d8cf0255bec3dfff58300167af470e01e32723b254a929e2c54b59eada.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i65f1f92552ffb48c728459d8e357479ff8ec797e294da0afafc24659ef8c3c24.CalendarRequestBuilder) {
    return i65f1f92552ffb48c728459d8e357479ff8ec797e294da0afafc24659ef8c3c24.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*ic74c15eb17958e3d5aedbd7aa4e0d9baf8d5434a1f33184dc193d8e698f2d520.CancelRequestBuilder) {
    return ic74c15eb17958e3d5aedbd7aa4e0d9baf8d5434a1f33184dc193d8e698f2d520.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
// Decline provides operations to call the decline method.
func (m *EventItemRequestBuilder) Decline()(*i4b92cda71aa39cf43f1f4117d98af10c1b939188ac1f64900b459cbc86ce63d3.DeclineRequestBuilder) {
    return i4b92cda71aa39cf43f1f4117d98af10c1b939188ac1f64900b459cbc86ce63d3.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i33c237b80872e74abd5582ee9e98c570c7e6f02883241a95a3c88ae46b984552.DismissReminderRequestBuilder) {
    return i33c237b80872e74abd5582ee9e98c570c7e6f02883241a95a3c88ae46b984552.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i92c1eb63d4144be862f92fc808e1688959b28b6423b622f3cfccb057eddfba26.ExtensionsRequestBuilder) {
    return i92c1eb63d4144be862f92fc808e1688959b28b6423b622f3cfccb057eddfba26.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i97d84954950d52b00814176753ad6fce764f2a130bdb4c3f31de8904563dcb49.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i97d84954950d52b00814176753ad6fce764f2a130bdb4c3f31de8904563dcb49.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i47f85055d10bbc2ce5596ccc6b1cf375e39ff7636bbd2a8a85a26b04c8ec760b.ForwardRequestBuilder) {
    return i47f85055d10bbc2ce5596ccc6b1cf375e39ff7636bbd2a8a85a26b04c8ec760b.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i940bde31fa32e362016c8393cc476a6f0a102a2ba874a3f21c82490cfb412bd4.MultiValueExtendedPropertiesRequestBuilder) {
    return i940bde31fa32e362016c8393cc476a6f0a102a2ba874a3f21c82490cfb412bd4.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ibf13cd72807bb3c8eb9ac05dcacf04c61cdd52496d8431cfdb6fbcd120da847f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ibf13cd72807bb3c8eb9ac05dcacf04c61cdd52496d8431cfdb6fbcd120da847f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i5da7820459f807892d24ff65fd96a91f4aac050c463293da376172f19b75a513.SingleValueExtendedPropertiesRequestBuilder) {
    return i5da7820459f807892d24ff65fd96a91f4aac050c463293da376172f19b75a513.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0a6f89794fdbde0788f0c2ff964c175f2d1e14b59909b07026acceeaf74bbb15.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i0a6f89794fdbde0788f0c2ff964c175f2d1e14b59909b07026acceeaf74bbb15.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i6c8729e6fa598d8c1a1def789a86190e6c425f33d3adeb1666d815fbf8521c04.SnoozeReminderRequestBuilder) {
    return i6c8729e6fa598d8c1a1def789a86190e6c425f33d3adeb1666d815fbf8521c04.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*icd2d22f9a6b46c7b3ed6518f51fc18003eea511e431f001ec72a8d90370868b5.TentativelyAcceptRequestBuilder) {
    return icd2d22f9a6b46c7b3ed6518f51fc18003eea511e431f001ec72a8d90370868b5.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
