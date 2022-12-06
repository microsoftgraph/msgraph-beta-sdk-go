package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i14ee84d132a9c3d8926db67682cef112db0cb08d61907eacf46e71c90773a168 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i191a2656a54e6d98f5a1b2f5ceaf27dda6668d6b40e62af29f9cd88c04f7ed79 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    i1d04871dc730aa67e925b820c690a45170dfbae60446824842e9889a561f68b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/decline"
    i28b430b4ea7bf7a83aa2e1eb3adb0ed17e0645edb3fe7005566c78e427579a31 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/attachments"
    i41366d4861c412e3ea08b69ee6437b2acdd8456b62b46fdd36e2657132daecc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i4d36c25838a234196c1798c8f82271254fb5e24540b2842d9990f24ce985d9e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/dismissreminder"
    i50fa47fa627101d9d4d37ae0724349fec8ff003ae30c00cd25c74fc926b5d870 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i5b141e0023ea6fc1b69122d94b6567b01ae861c1fc9fc244d545a1c244188fcc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/calendar"
    i8a0351ba7ade483be42dce2132220dc831f9baa2a8516cfddb07533af5898942 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/forward"
    i98fbb10cb38a7bfeb9bd2ae9159c623bbff0711095953f41a676efb8994f1ab2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances"
    ib1c3e913aa3c8ee51789fc49d9b0a537c2575698dd64657f4b68648d894f0c72 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/extensions"
    ic87ecf61420a4556d7d841ea43309af1046805ac4a60e4b8a1962e334f65418e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/accept"
    ide812d7b455c638ddacec6fdf3dcf7407902c381a4af126dd8c239e5606d34d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/cancel"
    i0466fed42dfe3428ea80bfe7746d30b2398dcd9a31788d1fe8af4d16dc513214 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i159ca1d5c2afad34162ceb364b677128fb5db6c9a39142ca0d92d553b6e4cb62 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/attachments/item"
    i67bacd62c51e8801bb4850ef42fb6042cac7672c15c68b351a62bfb3e98ea00a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/extensions/item"
    ibff5857072d6f9e31bebc0cdb4d79c01dd5c1d03b369c84d0da2a06c033f8755 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item"
    if87420bae8cf0ec1873b6781a34d18c378909c5afdcac9e288b9f836fdeb31f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*ic87ecf61420a4556d7d841ea43309af1046805ac4a60e4b8a1962e334f65418e.AcceptRequestBuilder) {
    return ic87ecf61420a4556d7d841ea43309af1046805ac4a60e4b8a1962e334f65418e.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i28b430b4ea7bf7a83aa2e1eb3adb0ed17e0645edb3fe7005566c78e427579a31.AttachmentsRequestBuilder) {
    return i28b430b4ea7bf7a83aa2e1eb3adb0ed17e0645edb3fe7005566c78e427579a31.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i159ca1d5c2afad34162ceb364b677128fb5db6c9a39142ca0d92d553b6e4cb62.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i159ca1d5c2afad34162ceb364b677128fb5db6c9a39142ca0d92d553b6e4cb62.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i5b141e0023ea6fc1b69122d94b6567b01ae861c1fc9fc244d545a1c244188fcc.CalendarRequestBuilder) {
    return i5b141e0023ea6fc1b69122d94b6567b01ae861c1fc9fc244d545a1c244188fcc.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*ide812d7b455c638ddacec6fdf3dcf7407902c381a4af126dd8c239e5606d34d1.CancelRequestBuilder) {
    return ide812d7b455c638ddacec6fdf3dcf7407902c381a4af126dd8c239e5606d34d1.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i1d04871dc730aa67e925b820c690a45170dfbae60446824842e9889a561f68b6.DeclineRequestBuilder) {
    return i1d04871dc730aa67e925b820c690a45170dfbae60446824842e9889a561f68b6.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i4d36c25838a234196c1798c8f82271254fb5e24540b2842d9990f24ce985d9e6.DismissReminderRequestBuilder) {
    return i4d36c25838a234196c1798c8f82271254fb5e24540b2842d9990f24ce985d9e6.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*ib1c3e913aa3c8ee51789fc49d9b0a537c2575698dd64657f4b68648d894f0c72.ExtensionsRequestBuilder) {
    return ib1c3e913aa3c8ee51789fc49d9b0a537c2575698dd64657f4b68648d894f0c72.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i67bacd62c51e8801bb4850ef42fb6042cac7672c15c68b351a62bfb3e98ea00a.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i67bacd62c51e8801bb4850ef42fb6042cac7672c15c68b351a62bfb3e98ea00a.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i8a0351ba7ade483be42dce2132220dc831f9baa2a8516cfddb07533af5898942.ForwardRequestBuilder) {
    return i8a0351ba7ade483be42dce2132220dc831f9baa2a8516cfddb07533af5898942.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Instances provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Instances()(*i98fbb10cb38a7bfeb9bd2ae9159c623bbff0711095953f41a676efb8994f1ab2.InstancesRequestBuilder) {
    return i98fbb10cb38a7bfeb9bd2ae9159c623bbff0711095953f41a676efb8994f1ab2.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) InstancesById(id string)(*ibff5857072d6f9e31bebc0cdb4d79c01dd5c1d03b369c84d0da2a06c033f8755.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return ibff5857072d6f9e31bebc0cdb4d79c01dd5c1d03b369c84d0da2a06c033f8755.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i191a2656a54e6d98f5a1b2f5ceaf27dda6668d6b40e62af29f9cd88c04f7ed79.MultiValueExtendedPropertiesRequestBuilder) {
    return i191a2656a54e6d98f5a1b2f5ceaf27dda6668d6b40e62af29f9cd88c04f7ed79.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if87420bae8cf0ec1873b6781a34d18c378909c5afdcac9e288b9f836fdeb31f7.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return if87420bae8cf0ec1873b6781a34d18c378909c5afdcac9e288b9f836fdeb31f7.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i50fa47fa627101d9d4d37ae0724349fec8ff003ae30c00cd25c74fc926b5d870.SingleValueExtendedPropertiesRequestBuilder) {
    return i50fa47fa627101d9d4d37ae0724349fec8ff003ae30c00cd25c74fc926b5d870.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0466fed42dfe3428ea80bfe7746d30b2398dcd9a31788d1fe8af4d16dc513214.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i0466fed42dfe3428ea80bfe7746d30b2398dcd9a31788d1fe8af4d16dc513214.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i14ee84d132a9c3d8926db67682cef112db0cb08d61907eacf46e71c90773a168.SnoozeReminderRequestBuilder) {
    return i14ee84d132a9c3d8926db67682cef112db0cb08d61907eacf46e71c90773a168.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i41366d4861c412e3ea08b69ee6437b2acdd8456b62b46fdd36e2657132daecc3.TentativelyAcceptRequestBuilder) {
    return i41366d4861c412e3ea08b69ee6437b2acdd8456b62b46fdd36e2657132daecc3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
