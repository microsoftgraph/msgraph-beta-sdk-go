package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i019db5d653a643e1a549e917348bbdacf1efb3a8caf89ca71a40988c2d3ca692 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    i36bb761cf2b3fd08e22dd0629b6d3414537d80a143da72167abdc43db9fe5176 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i40b295700d1a3dc9d4856ca82158a88dc5d2d35c5958521997664ba4fb1148e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i51fe48770386ecc7a9855817ee2ac09be025a6a9b143895262f400e1b3714012 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i55d48ec4278f3950b9d37bd21c26e0c1f021183da3015589d261582b9ce9978f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i7a291b26188996d4cdd5a95ee9429a4386af4d615fa05f98cba8442103f1969f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/accept"
    i824af3ce9c56f248451ff666dab98ed0a74003056e040dbb1bf5e36d52a9a78b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i8992ab597ecd5e75d6fac72ebb8df63711765f795e1150140c8ca3b713450fc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    i8e6836cb48c3fcffb55577d809f7a3db58227602f22c9c5d5145099bc8b03b20 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    ib97d7ab381f4d3986575b3f368addd550b65ae2dd39dff078276ced80a3047ac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/forward"
    ibc8dd027671410060ae3e28393b66723e6b3fa7d64029404faa96df20501488f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/decline"
    ie5295b5cffb686191c2e1152e4dfe27d2a49572591853453888b9e997af70e80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    i56325cc99232a0ede8e79a8e95284320babc1e40ece138d8773f76f51025cf3b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    i66a83b5642595f7c324d4fae0b85ed53b98f278d9fbd18fc37cbcc3484190981 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    id8c880825b71c3b54ed5b24f987ef412bf93ec0d06b3b9ac1cf44402719b34aa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    if75358db76287af9368cb02f46e1d783229c00f6337c1d11b9013fb349cf584c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*i7a291b26188996d4cdd5a95ee9429a4386af4d615fa05f98cba8442103f1969f.AcceptRequestBuilder) {
    return i7a291b26188996d4cdd5a95ee9429a4386af4d615fa05f98cba8442103f1969f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*ie5295b5cffb686191c2e1152e4dfe27d2a49572591853453888b9e997af70e80.AttachmentsRequestBuilder) {
    return ie5295b5cffb686191c2e1152e4dfe27d2a49572591853453888b9e997af70e80.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*id8c880825b71c3b54ed5b24f987ef412bf93ec0d06b3b9ac1cf44402719b34aa.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return id8c880825b71c3b54ed5b24f987ef412bf93ec0d06b3b9ac1cf44402719b34aa.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i8992ab597ecd5e75d6fac72ebb8df63711765f795e1150140c8ca3b713450fc2.CalendarRequestBuilder) {
    return i8992ab597ecd5e75d6fac72ebb8df63711765f795e1150140c8ca3b713450fc2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*i019db5d653a643e1a549e917348bbdacf1efb3a8caf89ca71a40988c2d3ca692.CancelRequestBuilder) {
    return i019db5d653a643e1a549e917348bbdacf1efb3a8caf89ca71a40988c2d3ca692.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*ibc8dd027671410060ae3e28393b66723e6b3fa7d64029404faa96df20501488f.DeclineRequestBuilder) {
    return ibc8dd027671410060ae3e28393b66723e6b3fa7d64029404faa96df20501488f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i51fe48770386ecc7a9855817ee2ac09be025a6a9b143895262f400e1b3714012.DismissReminderRequestBuilder) {
    return i51fe48770386ecc7a9855817ee2ac09be025a6a9b143895262f400e1b3714012.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i8e6836cb48c3fcffb55577d809f7a3db58227602f22c9c5d5145099bc8b03b20.ExtensionsRequestBuilder) {
    return i8e6836cb48c3fcffb55577d809f7a3db58227602f22c9c5d5145099bc8b03b20.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*if75358db76287af9368cb02f46e1d783229c00f6337c1d11b9013fb349cf584c.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return if75358db76287af9368cb02f46e1d783229c00f6337c1d11b9013fb349cf584c.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*ib97d7ab381f4d3986575b3f368addd550b65ae2dd39dff078276ced80a3047ac.ForwardRequestBuilder) {
    return ib97d7ab381f4d3986575b3f368addd550b65ae2dd39dff078276ced80a3047ac.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i40b295700d1a3dc9d4856ca82158a88dc5d2d35c5958521997664ba4fb1148e2.MultiValueExtendedPropertiesRequestBuilder) {
    return i40b295700d1a3dc9d4856ca82158a88dc5d2d35c5958521997664ba4fb1148e2.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i56325cc99232a0ede8e79a8e95284320babc1e40ece138d8773f76f51025cf3b.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i56325cc99232a0ede8e79a8e95284320babc1e40ece138d8773f76f51025cf3b.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i824af3ce9c56f248451ff666dab98ed0a74003056e040dbb1bf5e36d52a9a78b.SingleValueExtendedPropertiesRequestBuilder) {
    return i824af3ce9c56f248451ff666dab98ed0a74003056e040dbb1bf5e36d52a9a78b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i66a83b5642595f7c324d4fae0b85ed53b98f278d9fbd18fc37cbcc3484190981.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i66a83b5642595f7c324d4fae0b85ed53b98f278d9fbd18fc37cbcc3484190981.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i55d48ec4278f3950b9d37bd21c26e0c1f021183da3015589d261582b9ce9978f.SnoozeReminderRequestBuilder) {
    return i55d48ec4278f3950b9d37bd21c26e0c1f021183da3015589d261582b9ce9978f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i36bb761cf2b3fd08e22dd0629b6d3414537d80a143da72167abdc43db9fe5176.TentativelyAcceptRequestBuilder) {
    return i36bb761cf2b3fd08e22dd0629b6d3414537d80a143da72167abdc43db9fe5176.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
