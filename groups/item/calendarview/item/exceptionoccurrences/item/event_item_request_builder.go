package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04fd641e5a747104f4e0952e98213abc3186667104445a5618d14edf07a6772f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/accept"
    i0fd40eb9949327ae3e38a12a37d8b889a9631e6d019eb079118b1de218b3a98f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i2c59e99a1d1b0518a61a8d556af40d5e5fe456c174ae32eb81a450a2b24ce35c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i5349e999a824e9cbe5608bee94437dd38a319cc623f6ebbbcb20b48284a07301 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/forward"
    i5952b899fd7ad196b67e16874d1e062638b772ec2c510012e1c80ca74060fd3b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/attachments"
    i606eabc5b24235e1dfe5eda3ef9afb4b61fdb7829898669879ea22faca31f014 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/decline"
    i6c23de99d6e19538949dd0a1167c1f38757696973c1e8a2653dccaeb0f21001e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i75a0f8c972b66cf74259bd9d39194d06f5b12edab8842b7b01535f9d7ddfc5ff "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/calendar"
    i75b0bb687cc2cb0e33e9dc3c36651978a7b25ce6a0bce0d1664b36167919e00d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/cancel"
    i77956e6123d355433964150d0c3b33dbb67aa42ab7897ac20e593fe2645c189a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/dismissreminder"
    i880391b83bd4a880827c4fea0d059381b921905d4de77ce3149acbac5ca1b2b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances"
    i9c550bfac2e91800570b150240869141407885a38fdbebfc217a6cd5b70c39e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    ia2f47a8345604b1cb825cc4c0f191b0ee97048e0bbaf45dde6d3c0ec566708ab "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/extensions"
    i08619015a395f2dd82b2fbc2b9782d5ae07f6659e4525de00f3e820177ae8af5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/extensions/item"
    i49367fd7902603d0bb5c5b72858052db34b775a821866c9f83431b213d3b7416 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i4bb153a846f9374b2b40acb7ce407c7465ef82ac64d3f554f200662b0b8dc364 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/attachments/item"
    ib63af9909af4567eefbd6aae62ac98ec8c8d26d16d216a9e306e48f64c58b21e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ic87ae678f816a9e2aaffe8ffeea1c20689c1a543fcfd919fe16903a838f12f65 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Accept()(*i04fd641e5a747104f4e0952e98213abc3186667104445a5618d14edf07a6772f.AcceptRequestBuilder) {
    return i04fd641e5a747104f4e0952e98213abc3186667104445a5618d14edf07a6772f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i5952b899fd7ad196b67e16874d1e062638b772ec2c510012e1c80ca74060fd3b.AttachmentsRequestBuilder) {
    return i5952b899fd7ad196b67e16874d1e062638b772ec2c510012e1c80ca74060fd3b.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i4bb153a846f9374b2b40acb7ce407c7465ef82ac64d3f554f200662b0b8dc364.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i4bb153a846f9374b2b40acb7ce407c7465ef82ac64d3f554f200662b0b8dc364.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i75a0f8c972b66cf74259bd9d39194d06f5b12edab8842b7b01535f9d7ddfc5ff.CalendarRequestBuilder) {
    return i75a0f8c972b66cf74259bd9d39194d06f5b12edab8842b7b01535f9d7ddfc5ff.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*i75b0bb687cc2cb0e33e9dc3c36651978a7b25ce6a0bce0d1664b36167919e00d.CancelRequestBuilder) {
    return i75b0bb687cc2cb0e33e9dc3c36651978a7b25ce6a0bce0d1664b36167919e00d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
// CreateGetRequestInformation get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Decline()(*i606eabc5b24235e1dfe5eda3ef9afb4b61fdb7829898669879ea22faca31f014.DeclineRequestBuilder) {
    return i606eabc5b24235e1dfe5eda3ef9afb4b61fdb7829898669879ea22faca31f014.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i77956e6123d355433964150d0c3b33dbb67aa42ab7897ac20e593fe2645c189a.DismissReminderRequestBuilder) {
    return i77956e6123d355433964150d0c3b33dbb67aa42ab7897ac20e593fe2645c189a.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*ia2f47a8345604b1cb825cc4c0f191b0ee97048e0bbaf45dde6d3c0ec566708ab.ExtensionsRequestBuilder) {
    return ia2f47a8345604b1cb825cc4c0f191b0ee97048e0bbaf45dde6d3c0ec566708ab.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i08619015a395f2dd82b2fbc2b9782d5ae07f6659e4525de00f3e820177ae8af5.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i08619015a395f2dd82b2fbc2b9782d5ae07f6659e4525de00f3e820177ae8af5.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i5349e999a824e9cbe5608bee94437dd38a319cc623f6ebbbcb20b48284a07301.ForwardRequestBuilder) {
    return i5349e999a824e9cbe5608bee94437dd38a319cc623f6ebbbcb20b48284a07301.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Instances()(*i880391b83bd4a880827c4fea0d059381b921905d4de77ce3149acbac5ca1b2b3.InstancesRequestBuilder) {
    return i880391b83bd4a880827c4fea0d059381b921905d4de77ce3149acbac5ca1b2b3.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) InstancesById(id string)(*ic87ae678f816a9e2aaffe8ffeea1c20689c1a543fcfd919fe16903a838f12f65.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return ic87ae678f816a9e2aaffe8ffeea1c20689c1a543fcfd919fe16903a838f12f65.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i9c550bfac2e91800570b150240869141407885a38fdbebfc217a6cd5b70c39e4.MultiValueExtendedPropertiesRequestBuilder) {
    return i9c550bfac2e91800570b150240869141407885a38fdbebfc217a6cd5b70c39e4.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i49367fd7902603d0bb5c5b72858052db34b775a821866c9f83431b213d3b7416.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i49367fd7902603d0bb5c5b72858052db34b775a821866c9f83431b213d3b7416.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i6c23de99d6e19538949dd0a1167c1f38757696973c1e8a2653dccaeb0f21001e.SingleValueExtendedPropertiesRequestBuilder) {
    return i6c23de99d6e19538949dd0a1167c1f38757696973c1e8a2653dccaeb0f21001e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib63af9909af4567eefbd6aae62ac98ec8c8d26d16d216a9e306e48f64c58b21e.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ib63af9909af4567eefbd6aae62ac98ec8c8d26d16d216a9e306e48f64c58b21e.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i0fd40eb9949327ae3e38a12a37d8b889a9631e6d019eb079118b1de218b3a98f.SnoozeReminderRequestBuilder) {
    return i0fd40eb9949327ae3e38a12a37d8b889a9631e6d019eb079118b1de218b3a98f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i2c59e99a1d1b0518a61a8d556af40d5e5fe456c174ae32eb81a450a2b24ce35c.TentativelyAcceptRequestBuilder) {
    return i2c59e99a1d1b0518a61a8d556af40d5e5fe456c174ae32eb81a450a2b24ce35c.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
