package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1392060414ef191e570bf2441a06bf0aa02916b747db286a28fb7dc3c2c92e09 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i2138bb645bcb8a4c21f46dd5c1c182aae362eb9ad12f1032dc2a2a80b4234d40 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i4b74cf1475be79d8c5bd5a96a857dfee3b4a001104265d86db42bc4ccdb139f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i5a9277e9b29a102a4ac15815f43c255e33b63212fd83179cfb06ba9c0095a417 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i5c74e0eceb86a90abe58ab377605dfd9a92559258fc4e477917da5c5ec800911 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i638365c56d4a3a9d8db8ef7a8f8cb86cb08f3af8ba56b38f8f1b65bdf27df1bf "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    i6bc39607fa012755f044c52399c42bb807b680526e3542b3df42d1f675ce2edc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    i759d7b0b434060f3aaed7df848350d3517136226a983ee0c485acde76a53fed9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ic2e3264dc769002999eaf12fe5eec70ab157b8c926d2fd9e8e8b474cc4aed762 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    ice2f5260d34d3c032174d7c833245dd7952e01872c8c372f831f6ccda3b49aa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    ie00181941fec3c9663ffb77f6f64ef400b43cab518310760d1e5894cee015858 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    if08749e141f1fdc527f44f4b0e1c235ad9ab3135dd7323ba322054e32b36a03d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i0067dc107b6785a4cbaa9516df188178c3ed39eca93dd9f677c93380cc0a0d1a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i4235545d855e808026b5af63ffee955a888c59e75d548a24093ddf21321260f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i700bc3bbf501af9afe6e7f3fbfa81a604caccf7870a64a2b888153813b9c749a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    ia0ef39ad909bd7275646fee8900363fa776a983b1e646976e57cfe18a7c5bcbd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
func (m *EventItemRequestBuilder) Accept()(*ice2f5260d34d3c032174d7c833245dd7952e01872c8c372f831f6ccda3b49aa9.AcceptRequestBuilder) {
    return ice2f5260d34d3c032174d7c833245dd7952e01872c8c372f831f6ccda3b49aa9.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i5c74e0eceb86a90abe58ab377605dfd9a92559258fc4e477917da5c5ec800911.AttachmentsRequestBuilder) {
    return i5c74e0eceb86a90abe58ab377605dfd9a92559258fc4e477917da5c5ec800911.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i700bc3bbf501af9afe6e7f3fbfa81a604caccf7870a64a2b888153813b9c749a.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i700bc3bbf501af9afe6e7f3fbfa81a604caccf7870a64a2b888153813b9c749a.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i638365c56d4a3a9d8db8ef7a8f8cb86cb08f3af8ba56b38f8f1b65bdf27df1bf.CalendarRequestBuilder) {
    return i638365c56d4a3a9d8db8ef7a8f8cb86cb08f3af8ba56b38f8f1b65bdf27df1bf.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*i1392060414ef191e570bf2441a06bf0aa02916b747db286a28fb7dc3c2c92e09.CancelRequestBuilder) {
    return i1392060414ef191e570bf2441a06bf0aa02916b747db286a28fb7dc3c2c92e09.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
// CreateGetRequestInformation get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) Decline()(*i6bc39607fa012755f044c52399c42bb807b680526e3542b3df42d1f675ce2edc.DeclineRequestBuilder) {
    return i6bc39607fa012755f044c52399c42bb807b680526e3542b3df42d1f675ce2edc.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*ic2e3264dc769002999eaf12fe5eec70ab157b8c926d2fd9e8e8b474cc4aed762.DismissReminderRequestBuilder) {
    return ic2e3264dc769002999eaf12fe5eec70ab157b8c926d2fd9e8e8b474cc4aed762.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i5a9277e9b29a102a4ac15815f43c255e33b63212fd83179cfb06ba9c0095a417.ExtensionsRequestBuilder) {
    return i5a9277e9b29a102a4ac15815f43c255e33b63212fd83179cfb06ba9c0095a417.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ia0ef39ad909bd7275646fee8900363fa776a983b1e646976e57cfe18a7c5bcbd.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ia0ef39ad909bd7275646fee8900363fa776a983b1e646976e57cfe18a7c5bcbd.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i4b74cf1475be79d8c5bd5a96a857dfee3b4a001104265d86db42bc4ccdb139f2.ForwardRequestBuilder) {
    return i4b74cf1475be79d8c5bd5a96a857dfee3b4a001104265d86db42bc4ccdb139f2.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie00181941fec3c9663ffb77f6f64ef400b43cab518310760d1e5894cee015858.MultiValueExtendedPropertiesRequestBuilder) {
    return ie00181941fec3c9663ffb77f6f64ef400b43cab518310760d1e5894cee015858.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0067dc107b6785a4cbaa9516df188178c3ed39eca93dd9f677c93380cc0a0d1a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i0067dc107b6785a4cbaa9516df188178c3ed39eca93dd9f677c93380cc0a0d1a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i2138bb645bcb8a4c21f46dd5c1c182aae362eb9ad12f1032dc2a2a80b4234d40.SingleValueExtendedPropertiesRequestBuilder) {
    return i2138bb645bcb8a4c21f46dd5c1c182aae362eb9ad12f1032dc2a2a80b4234d40.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4235545d855e808026b5af63ffee955a888c59e75d548a24093ddf21321260f9.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i4235545d855e808026b5af63ffee955a888c59e75d548a24093ddf21321260f9.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*if08749e141f1fdc527f44f4b0e1c235ad9ab3135dd7323ba322054e32b36a03d.SnoozeReminderRequestBuilder) {
    return if08749e141f1fdc527f44f4b0e1c235ad9ab3135dd7323ba322054e32b36a03d.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i759d7b0b434060f3aaed7df848350d3517136226a983ee0c485acde76a53fed9.TentativelyAcceptRequestBuilder) {
    return i759d7b0b434060f3aaed7df848350d3517136226a983ee0c485acde76a53fed9.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
