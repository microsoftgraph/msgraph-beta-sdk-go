package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1a83e59eb57380c3a658a4b1024accce312df73ce038980af6ba9c617ade4f25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/forward"
    i1bef5e006a0d505d71f4671715fec55a93e9ecf403a2799872aa57b59e90617d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/cancel"
    i2aa4af366132e6693047bfec94521c36f1c5b4dbdbbf1655a8146892e156d12f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/decline"
    i316b301fa50729bedc547e0e238e1765d3c385551bb5bebf5c07c13e86520d50 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i38b845df59dccc1202cd959ac3f11b936fe6e7d81ca54ffac709c77e7cd67f84 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/accept"
    i3e475b83eed76216d2c113b9ccf504e34c383a6ba7175a407749f7b78489d561 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i3f9b8b4f6e160cc82ed455d5aa759803e457b693dafd6abab10371e374f16799 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/attachments"
    i9074cecfe0c46c337eb919bfa0e5cc966fd8c92dba95157275bb6f0fea16d96d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/extensions"
    ib6fce88e994b6edfde03ab7ddc638ac51696f68eac20d527c7f328b03eaa87d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/calendar"
    id08320318ad3cbfa2808800522162425dafa821fc32281ade6e78223d733ee60 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ieb1965442f708c862020f2259403b002a767802b37c3ea8d74325b6c755e949d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    ifabc5440e6442ca757d7798d1b5e6989cc29fb0e03d9ba78ac7c9c9309d5d4d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i2aafd6a994207ec13e8b1c7da283da944964014026e5f201f95013e1af35aff2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ic65f67470295cc89cb57e3ba6a6e5cdd92256aac14f9f148dfc3054cbe6d0d1b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    if46bfdeeeb4dc793c50e79f445ce75d7fab86e16092ccd7da2740ce94f5874c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    iffc4e0baae71a43c6686d88e5c91fa706968a7572971007ed81ab4a52b560e12 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) Accept()(*i38b845df59dccc1202cd959ac3f11b936fe6e7d81ca54ffac709c77e7cd67f84.AcceptRequestBuilder) {
    return i38b845df59dccc1202cd959ac3f11b936fe6e7d81ca54ffac709c77e7cd67f84.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i3f9b8b4f6e160cc82ed455d5aa759803e457b693dafd6abab10371e374f16799.AttachmentsRequestBuilder) {
    return i3f9b8b4f6e160cc82ed455d5aa759803e457b693dafd6abab10371e374f16799.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*if46bfdeeeb4dc793c50e79f445ce75d7fab86e16092ccd7da2740ce94f5874c0.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return if46bfdeeeb4dc793c50e79f445ce75d7fab86e16092ccd7da2740ce94f5874c0.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ib6fce88e994b6edfde03ab7ddc638ac51696f68eac20d527c7f328b03eaa87d5.CalendarRequestBuilder) {
    return ib6fce88e994b6edfde03ab7ddc638ac51696f68eac20d527c7f328b03eaa87d5.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i1bef5e006a0d505d71f4671715fec55a93e9ecf403a2799872aa57b59e90617d.CancelRequestBuilder) {
    return i1bef5e006a0d505d71f4671715fec55a93e9ecf403a2799872aa57b59e90617d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
// CreateGetRequestInformation get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) Decline()(*i2aa4af366132e6693047bfec94521c36f1c5b4dbdbbf1655a8146892e156d12f.DeclineRequestBuilder) {
    return i2aa4af366132e6693047bfec94521c36f1c5b4dbdbbf1655a8146892e156d12f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*ieb1965442f708c862020f2259403b002a767802b37c3ea8d74325b6c755e949d.DismissReminderRequestBuilder) {
    return ieb1965442f708c862020f2259403b002a767802b37c3ea8d74325b6c755e949d.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i9074cecfe0c46c337eb919bfa0e5cc966fd8c92dba95157275bb6f0fea16d96d.ExtensionsRequestBuilder) {
    return i9074cecfe0c46c337eb919bfa0e5cc966fd8c92dba95157275bb6f0fea16d96d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic65f67470295cc89cb57e3ba6a6e5cdd92256aac14f9f148dfc3054cbe6d0d1b.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ic65f67470295cc89cb57e3ba6a6e5cdd92256aac14f9f148dfc3054cbe6d0d1b.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i1a83e59eb57380c3a658a4b1024accce312df73ce038980af6ba9c617ade4f25.ForwardRequestBuilder) {
    return i1a83e59eb57380c3a658a4b1024accce312df73ce038980af6ba9c617ade4f25.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i316b301fa50729bedc547e0e238e1765d3c385551bb5bebf5c07c13e86520d50.MultiValueExtendedPropertiesRequestBuilder) {
    return i316b301fa50729bedc547e0e238e1765d3c385551bb5bebf5c07c13e86520d50.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iffc4e0baae71a43c6686d88e5c91fa706968a7572971007ed81ab4a52b560e12.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iffc4e0baae71a43c6686d88e5c91fa706968a7572971007ed81ab4a52b560e12.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i3e475b83eed76216d2c113b9ccf504e34c383a6ba7175a407749f7b78489d561.SingleValueExtendedPropertiesRequestBuilder) {
    return i3e475b83eed76216d2c113b9ccf504e34c383a6ba7175a407749f7b78489d561.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i2aafd6a994207ec13e8b1c7da283da944964014026e5f201f95013e1af35aff2.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i2aafd6a994207ec13e8b1c7da283da944964014026e5f201f95013e1af35aff2.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ifabc5440e6442ca757d7798d1b5e6989cc29fb0e03d9ba78ac7c9c9309d5d4d6.SnoozeReminderRequestBuilder) {
    return ifabc5440e6442ca757d7798d1b5e6989cc29fb0e03d9ba78ac7c9c9309d5d4d6.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*id08320318ad3cbfa2808800522162425dafa821fc32281ade6e78223d733ee60.TentativelyAcceptRequestBuilder) {
    return id08320318ad3cbfa2808800522162425dafa821fc32281ade6e78223d733ee60.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
