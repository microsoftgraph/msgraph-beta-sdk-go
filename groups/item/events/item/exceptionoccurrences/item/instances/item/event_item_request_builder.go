package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0bd35b1fb9b5cd3a509560e24ec57494bb35171b5ffa0a6e007c91ce2ba419c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/accept"
    i187c9e6816ffbc23be0deb69b1a127130ae37c10edf1254e9b5826ff6031ba01 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    i332f925e36b2fe945c50b1b6d90c949acde78fca9bdddf4f0487fd904ab368b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i4476c63e8cf22dbd915a41d4ab33b7425e126d395bceb40bad6ab35569ae35d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i4956ec783cf4129a29c7531c7d13a3f7bb3f5c830d1c90ee045cecaefe3288be "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    i4cab08f986acca114d1529f5b639030b91f1da7c714ffc08f80e31221192dcd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i7a60b93f7d0221f3b1402a67920fc82eb7d36db595295f22b24ff301b1ec7098 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    i9248a10c6abd9ed71d01c49cf1e461e9e9b06c0ef44578f7d33b2c9e99cc2aba "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    ia4ef4f75b96a0768357c29a43a6454f3b1a8ae188bee0c585960a7babc17654f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/forward"
    ia812ae1dc4b1f30fa3f2a099e78c863c203a7f19deb5c36ac0ed96f2e1514a3c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/decline"
    ic66635521102afb8cd7214f483378adefcbe7d0f153f0a7ef85b1da938364825 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    icca8571980095ab1b8115b8579ce64b33631d293ab438a53752a6360e4fd2154 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    i4f5f241ea2012d39eb24707ce6d64e3cb86fc78310cdccb1e0caeff5547fa816 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    i558f96fb156da918ec2245ed183a904c7c5e6ce361c4ca7e14bb110bbf083ee2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    ic3ef47a5c760d3b34d87fdce711d26557f319a1afff6d092446c7dc99656869b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    ie3ae8b75285bede0c6f450453901c8aea50863d5c526729b8f466fc8a489ce60 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*i0bd35b1fb9b5cd3a509560e24ec57494bb35171b5ffa0a6e007c91ce2ba419c9.AcceptRequestBuilder) {
    return i0bd35b1fb9b5cd3a509560e24ec57494bb35171b5ffa0a6e007c91ce2ba419c9.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i7a60b93f7d0221f3b1402a67920fc82eb7d36db595295f22b24ff301b1ec7098.AttachmentsRequestBuilder) {
    return i7a60b93f7d0221f3b1402a67920fc82eb7d36db595295f22b24ff301b1ec7098.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i558f96fb156da918ec2245ed183a904c7c5e6ce361c4ca7e14bb110bbf083ee2.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i558f96fb156da918ec2245ed183a904c7c5e6ce361c4ca7e14bb110bbf083ee2.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i187c9e6816ffbc23be0deb69b1a127130ae37c10edf1254e9b5826ff6031ba01.CalendarRequestBuilder) {
    return i187c9e6816ffbc23be0deb69b1a127130ae37c10edf1254e9b5826ff6031ba01.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*icca8571980095ab1b8115b8579ce64b33631d293ab438a53752a6360e4fd2154.CancelRequestBuilder) {
    return icca8571980095ab1b8115b8579ce64b33631d293ab438a53752a6360e4fd2154.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*ia812ae1dc4b1f30fa3f2a099e78c863c203a7f19deb5c36ac0ed96f2e1514a3c.DeclineRequestBuilder) {
    return ia812ae1dc4b1f30fa3f2a099e78c863c203a7f19deb5c36ac0ed96f2e1514a3c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i4cab08f986acca114d1529f5b639030b91f1da7c714ffc08f80e31221192dcd3.DismissReminderRequestBuilder) {
    return i4cab08f986acca114d1529f5b639030b91f1da7c714ffc08f80e31221192dcd3.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i4956ec783cf4129a29c7531c7d13a3f7bb3f5c830d1c90ee045cecaefe3288be.ExtensionsRequestBuilder) {
    return i4956ec783cf4129a29c7531c7d13a3f7bb3f5c830d1c90ee045cecaefe3288be.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ie3ae8b75285bede0c6f450453901c8aea50863d5c526729b8f466fc8a489ce60.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ie3ae8b75285bede0c6f450453901c8aea50863d5c526729b8f466fc8a489ce60.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*ia4ef4f75b96a0768357c29a43a6454f3b1a8ae188bee0c585960a7babc17654f.ForwardRequestBuilder) {
    return ia4ef4f75b96a0768357c29a43a6454f3b1a8ae188bee0c585960a7babc17654f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i4476c63e8cf22dbd915a41d4ab33b7425e126d395bceb40bad6ab35569ae35d9.MultiValueExtendedPropertiesRequestBuilder) {
    return i4476c63e8cf22dbd915a41d4ab33b7425e126d395bceb40bad6ab35569ae35d9.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4f5f241ea2012d39eb24707ce6d64e3cb86fc78310cdccb1e0caeff5547fa816.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i4f5f241ea2012d39eb24707ce6d64e3cb86fc78310cdccb1e0caeff5547fa816.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i332f925e36b2fe945c50b1b6d90c949acde78fca9bdddf4f0487fd904ab368b1.SingleValueExtendedPropertiesRequestBuilder) {
    return i332f925e36b2fe945c50b1b6d90c949acde78fca9bdddf4f0487fd904ab368b1.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic3ef47a5c760d3b34d87fdce711d26557f319a1afff6d092446c7dc99656869b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ic3ef47a5c760d3b34d87fdce711d26557f319a1afff6d092446c7dc99656869b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i9248a10c6abd9ed71d01c49cf1e461e9e9b06c0ef44578f7d33b2c9e99cc2aba.SnoozeReminderRequestBuilder) {
    return i9248a10c6abd9ed71d01c49cf1e461e9e9b06c0ef44578f7d33b2c9e99cc2aba.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*ic66635521102afb8cd7214f483378adefcbe7d0f153f0a7ef85b1da938364825.TentativelyAcceptRequestBuilder) {
    return ic66635521102afb8cd7214f483378adefcbe7d0f153f0a7ef85b1da938364825.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
