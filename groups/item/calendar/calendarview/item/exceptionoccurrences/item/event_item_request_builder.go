package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06614cf97b25e4f1cc95b099db1bff5aa2d15eb9fe48ceeec06d30f041cec7b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i20818696d027c8339014df26f5cff39fca939fea5e71a491efd9e8a8feb34d96 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i245e54c9299bc73f21d507de0abc77a91a3ee273ede39d36f8cc4b30d3d929d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/cancel"
    i2c6f242e034d183656c5cbb51a72e875198d972c7654e24ca256a037cd60fa2a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/attachments"
    i342c127c12be2e6048b7dff13a836a0b2d04c12eba2f598c2e211d0dd0666e5a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/accept"
    i380e854f772bf5bc9a128e8415ee4ad551ae07887f6a329120e013966758c29c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/forward"
    i7968aa204c89c1bd548200a458a991bebe04ca6a16068768c1269ebf170e26f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/extensions"
    ibd54d8108ed0df1725e8c7a096107a962956e60e3b9b4d0851e4b09aa0d4bd1e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances"
    ic5673a6bd5616076eb617c73ae999646f4b1e1254759dd671a8f249fb54a8a37 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/dismissreminder"
    id0d2742f6d706f93431cd309866bd747cae6f6fb721136a8478052f87a40452d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/decline"
    ide9cc24b38c043ff97d2882e83ae818dc56b4d7ed8bfa6dab5ac8b8e766435da "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/calendar"
    ie2bc1bf78859cacf27c559fad849551bf1cd6552993f78d6f1e7e4c484dfc5d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    ie48d50a4a462cf7ba0136119b50d82579a435f443e9085a25c31f262598ced12 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i15f5819471788997963c36e490f6a51e9f5550aa4436173f156b761ef03a91ff "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/extensions/item"
    i30151f5bb8a6840dfd789ab1d79f1791912103c60c88280ee2b9070461ca730f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i50f22ae6af4b1db80a0da3d8151a9ddaba5fcae0f9669dd5503fe1945048a7b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ic98255458c7c86b6d4c11d49c572809d991cd1cfd656b375f4e29a936576cdd4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item"
    ifa2e8f16b45156022dbea12e9f778c0713f79a848df0dd18efd079d168eba76e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/exceptionoccurrences/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*i342c127c12be2e6048b7dff13a836a0b2d04c12eba2f598c2e211d0dd0666e5a.AcceptRequestBuilder) {
    return i342c127c12be2e6048b7dff13a836a0b2d04c12eba2f598c2e211d0dd0666e5a.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i2c6f242e034d183656c5cbb51a72e875198d972c7654e24ca256a037cd60fa2a.AttachmentsRequestBuilder) {
    return i2c6f242e034d183656c5cbb51a72e875198d972c7654e24ca256a037cd60fa2a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ifa2e8f16b45156022dbea12e9f778c0713f79a848df0dd18efd079d168eba76e.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ifa2e8f16b45156022dbea12e9f778c0713f79a848df0dd18efd079d168eba76e.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*ide9cc24b38c043ff97d2882e83ae818dc56b4d7ed8bfa6dab5ac8b8e766435da.CalendarRequestBuilder) {
    return ide9cc24b38c043ff97d2882e83ae818dc56b4d7ed8bfa6dab5ac8b8e766435da.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*i245e54c9299bc73f21d507de0abc77a91a3ee273ede39d36f8cc4b30d3d929d0.CancelRequestBuilder) {
    return i245e54c9299bc73f21d507de0abc77a91a3ee273ede39d36f8cc4b30d3d929d0.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*id0d2742f6d706f93431cd309866bd747cae6f6fb721136a8478052f87a40452d.DeclineRequestBuilder) {
    return id0d2742f6d706f93431cd309866bd747cae6f6fb721136a8478052f87a40452d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*ic5673a6bd5616076eb617c73ae999646f4b1e1254759dd671a8f249fb54a8a37.DismissReminderRequestBuilder) {
    return ic5673a6bd5616076eb617c73ae999646f4b1e1254759dd671a8f249fb54a8a37.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*i7968aa204c89c1bd548200a458a991bebe04ca6a16068768c1269ebf170e26f3.ExtensionsRequestBuilder) {
    return i7968aa204c89c1bd548200a458a991bebe04ca6a16068768c1269ebf170e26f3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i15f5819471788997963c36e490f6a51e9f5550aa4436173f156b761ef03a91ff.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i15f5819471788997963c36e490f6a51e9f5550aa4436173f156b761ef03a91ff.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i380e854f772bf5bc9a128e8415ee4ad551ae07887f6a329120e013966758c29c.ForwardRequestBuilder) {
    return i380e854f772bf5bc9a128e8415ee4ad551ae07887f6a329120e013966758c29c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*ibd54d8108ed0df1725e8c7a096107a962956e60e3b9b4d0851e4b09aa0d4bd1e.InstancesRequestBuilder) {
    return ibd54d8108ed0df1725e8c7a096107a962956e60e3b9b4d0851e4b09aa0d4bd1e.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById provides operations to manage the instances property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) InstancesById(id string)(*ic98255458c7c86b6d4c11d49c572809d991cd1cfd656b375f4e29a936576cdd4.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return ic98255458c7c86b6d4c11d49c572809d991cd1cfd656b375f4e29a936576cdd4.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie2bc1bf78859cacf27c559fad849551bf1cd6552993f78d6f1e7e4c484dfc5d0.MultiValueExtendedPropertiesRequestBuilder) {
    return ie2bc1bf78859cacf27c559fad849551bf1cd6552993f78d6f1e7e4c484dfc5d0.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i50f22ae6af4b1db80a0da3d8151a9ddaba5fcae0f9669dd5503fe1945048a7b5.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i50f22ae6af4b1db80a0da3d8151a9ddaba5fcae0f9669dd5503fe1945048a7b5.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i20818696d027c8339014df26f5cff39fca939fea5e71a491efd9e8a8feb34d96.SingleValueExtendedPropertiesRequestBuilder) {
    return i20818696d027c8339014df26f5cff39fca939fea5e71a491efd9e8a8feb34d96.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i30151f5bb8a6840dfd789ab1d79f1791912103c60c88280ee2b9070461ca730f.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i30151f5bb8a6840dfd789ab1d79f1791912103c60c88280ee2b9070461ca730f.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*i06614cf97b25e4f1cc95b099db1bff5aa2d15eb9fe48ceeec06d30f041cec7b3.SnoozeReminderRequestBuilder) {
    return i06614cf97b25e4f1cc95b099db1bff5aa2d15eb9fe48ceeec06d30f041cec7b3.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*ie48d50a4a462cf7ba0136119b50d82579a435f443e9085a25c31f262598ced12.TentativelyAcceptRequestBuilder) {
    return ie48d50a4a462cf7ba0136119b50d82579a435f443e9085a25c31f262598ced12.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
