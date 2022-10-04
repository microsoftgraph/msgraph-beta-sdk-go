package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1d63ccc415a2bbf541b9614f6eb0e657f4405ea836748d6e5a444aeca1cda369 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i2df6c0ffe311e222cc43f3c5c9b9b0eab46263354f6503d324e662b1f482c21b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/dismissreminder"
    i89045cd833e17bc2c440dea0765851e45126a94f6e79005670d0c283adf0381a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances"
    i895b91699f960b44fb654d7e9c171e84b54bcb0ab1b52b18eddf8f02c467d295 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    i8acf00073e73a3171f4c587d2cdfdc7e3536158d1667729740172f60689146da "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/attachments"
    i95e2251f06fb8e130e9d603a600e097ef56c24fb533e7ea10b8bf9753a57769b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/decline"
    i9e20f6c5ac30fdefa7c15a6f1cea3cf7455d9ba2b7a818ab0fdc702e79418f7e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/extensions"
    ibc85e661acac9ac2cd9dd76b7b360b074582e5bc6cb69b091ef925c82e8a42fa "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/cancel"
    ibcfb1244367c25c3f1cc50287473910283bb6789e16767a3331343440c3e4683 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/accept"
    id669929c5b24aae0fd80a71b4b066df78a9a0b89af2f4b1f1da82114be83c370 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/snoozereminder"
    id7512aac6b609eb5d7ee0134960f895648e28a4453aacd221f4091f803d2d0ec "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ie0e2915c58785b51b3401a6e97165778562e7a3707951ae5c366251325343011 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/forward"
    if9f622d3619b48d56f92e05cf17e2e91e2b883c8b5d8a724251083f87f2241dc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/calendar"
    i1fa57022e27d9222d1da8189f48a7bec257ba3bdb46f2a496c7e132eae9f5c86 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item"
    i66285810dbb0f78f31e742889b203ba950c309debbb8e49da1550887a9211ff0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/attachments/item"
    i6d06fc9e4048ff898cd856b76ef3fb5847f18a99e7ca22848f5770515a8063b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/extensions/item"
    ic231cf3454eff45ab85f77c6165fbe381b953c0f7f21c32c784a9c7f95f9c335 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    icb64d2ed76ad9a277dfc29fdbcb22396affeaeee3d3fd99a2fa880fd74c38e21 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
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
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*ibcfb1244367c25c3f1cc50287473910283bb6789e16767a3331343440c3e4683.AcceptRequestBuilder) {
    return ibcfb1244367c25c3f1cc50287473910283bb6789e16767a3331343440c3e4683.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i8acf00073e73a3171f4c587d2cdfdc7e3536158d1667729740172f60689146da.AttachmentsRequestBuilder) {
    return i8acf00073e73a3171f4c587d2cdfdc7e3536158d1667729740172f60689146da.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i66285810dbb0f78f31e742889b203ba950c309debbb8e49da1550887a9211ff0.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i66285810dbb0f78f31e742889b203ba950c309debbb8e49da1550887a9211ff0.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*if9f622d3619b48d56f92e05cf17e2e91e2b883c8b5d8a724251083f87f2241dc.CalendarRequestBuilder) {
    return if9f622d3619b48d56f92e05cf17e2e91e2b883c8b5d8a724251083f87f2241dc.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ibc85e661acac9ac2cd9dd76b7b360b074582e5bc6cb69b091ef925c82e8a42fa.CancelRequestBuilder) {
    return ibc85e661acac9ac2cd9dd76b7b360b074582e5bc6cb69b091ef925c82e8a42fa.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i95e2251f06fb8e130e9d603a600e097ef56c24fb533e7ea10b8bf9753a57769b.DeclineRequestBuilder) {
    return i95e2251f06fb8e130e9d603a600e097ef56c24fb533e7ea10b8bf9753a57769b.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i2df6c0ffe311e222cc43f3c5c9b9b0eab46263354f6503d324e662b1f482c21b.DismissReminderRequestBuilder) {
    return i2df6c0ffe311e222cc43f3c5c9b9b0eab46263354f6503d324e662b1f482c21b.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i9e20f6c5ac30fdefa7c15a6f1cea3cf7455d9ba2b7a818ab0fdc702e79418f7e.ExtensionsRequestBuilder) {
    return i9e20f6c5ac30fdefa7c15a6f1cea3cf7455d9ba2b7a818ab0fdc702e79418f7e.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i6d06fc9e4048ff898cd856b76ef3fb5847f18a99e7ca22848f5770515a8063b1.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i6d06fc9e4048ff898cd856b76ef3fb5847f18a99e7ca22848f5770515a8063b1.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ie0e2915c58785b51b3401a6e97165778562e7a3707951ae5c366251325343011.ForwardRequestBuilder) {
    return ie0e2915c58785b51b3401a6e97165778562e7a3707951ae5c366251325343011.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*i89045cd833e17bc2c440dea0765851e45126a94f6e79005670d0c283adf0381a.InstancesRequestBuilder) {
    return i89045cd833e17bc2c440dea0765851e45126a94f6e79005670d0c283adf0381a.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i1fa57022e27d9222d1da8189f48a7bec257ba3bdb46f2a496c7e132eae9f5c86.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i1fa57022e27d9222d1da8189f48a7bec257ba3bdb46f2a496c7e132eae9f5c86.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i895b91699f960b44fb654d7e9c171e84b54bcb0ab1b52b18eddf8f02c467d295.MultiValueExtendedPropertiesRequestBuilder) {
    return i895b91699f960b44fb654d7e9c171e84b54bcb0ab1b52b18eddf8f02c467d295.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic231cf3454eff45ab85f77c6165fbe381b953c0f7f21c32c784a9c7f95f9c335.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ic231cf3454eff45ab85f77c6165fbe381b953c0f7f21c32c784a9c7f95f9c335.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*id7512aac6b609eb5d7ee0134960f895648e28a4453aacd221f4091f803d2d0ec.SingleValueExtendedPropertiesRequestBuilder) {
    return id7512aac6b609eb5d7ee0134960f895648e28a4453aacd221f4091f803d2d0ec.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*icb64d2ed76ad9a277dfc29fdbcb22396affeaeee3d3fd99a2fa880fd74c38e21.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return icb64d2ed76ad9a277dfc29fdbcb22396affeaeee3d3fd99a2fa880fd74c38e21.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*id669929c5b24aae0fd80a71b4b066df78a9a0b89af2f4b1f1da82114be83c370.SnoozeReminderRequestBuilder) {
    return id669929c5b24aae0fd80a71b4b066df78a9a0b89af2f4b1f1da82114be83c370.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i1d63ccc415a2bbf541b9614f6eb0e657f4405ea836748d6e5a444aeca1cda369.TentativelyAcceptRequestBuilder) {
    return i1d63ccc415a2bbf541b9614f6eb0e657f4405ea836748d6e5a444aeca1cda369.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
