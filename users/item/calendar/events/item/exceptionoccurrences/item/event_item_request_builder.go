package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i050fa3cfba1766755e62be09db7509d697402595f146dfa6917836103ea6c418 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/dismissreminder"
    i117d9ebe24172939d0a5f63702c7509338a57c472472a49ee2565208b4d780bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/calendar"
    i17325f41fc1dfe0de0e904a3ae38c725efa0d2265ffce8ce6f3018f12c9a4b4e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/forward"
    i2ad554bf085255bb25cc0a0d6980dd873fa36373fa4837961d9aa588b1712d8f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/attachments"
    i45b9326583488618ead31962b306c42d114f0247787ec0455579853c35a23226 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/tentativelyaccept"
    i6382d1d739bdd919528e2d99dc700679ac6ffd1b44d5683e92a1ebce647f9cf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    i7ec706215b1a4de83953f27df5698ba1502115937a890507c5a1b1952db5bf2d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/extensions"
    i890234a43eddfbb7906d7fc141859ab3228b2bf8edbebe01a56d2e9011f52236 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/decline"
    i9ece5c52e1d32b599990f2019c6d3a778f13f9f332d1ddbab1abb62e8a20c07d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances"
    ia816fb6558bf6de9fca63f7aa79137e2a5c51a9979ce0e65066b9062eb2becdd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/snoozereminder"
    id0ec38d5b88daca23dd7e9f451ae04167b318090b27afd19bf9474f97e48af9e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/accept"
    id814468e09981972f31dbeba72bd3b9bc3d5998b33227510f19290d97aaf643f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ie82bc68c2150b10b6eac9873b0428f1ecff6323accd09b4b729023c9439cf923 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/cancel"
    i06a8ed3557ec38fcf30fa35b273ef30e00d94a0d50b15de290675d1cea7a30e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/attachments/item"
    i4f3760ebaf8befb9f04aa8e0f72101e132c43a2e798e6ff11f716e9e0a766314 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    iaca324da655a515ff1fd42e185390f292ff9e89839ba866e74bc5f3d3f127c36 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/instances/item"
    ieca6997f19b6f68725b8ff66b95f5efaa5c4aba0bd54ed35eab5e4d7ce828403 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/extensions/item"
    ief603bdd91788274032ba500639b4fbbe934af2730bed0cd2ba2199d2b388843 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*id0ec38d5b88daca23dd7e9f451ae04167b318090b27afd19bf9474f97e48af9e.AcceptRequestBuilder) {
    return id0ec38d5b88daca23dd7e9f451ae04167b318090b27afd19bf9474f97e48af9e.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i2ad554bf085255bb25cc0a0d6980dd873fa36373fa4837961d9aa588b1712d8f.AttachmentsRequestBuilder) {
    return i2ad554bf085255bb25cc0a0d6980dd873fa36373fa4837961d9aa588b1712d8f.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i06a8ed3557ec38fcf30fa35b273ef30e00d94a0d50b15de290675d1cea7a30e5.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i06a8ed3557ec38fcf30fa35b273ef30e00d94a0d50b15de290675d1cea7a30e5.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i117d9ebe24172939d0a5f63702c7509338a57c472472a49ee2565208b4d780bd.CalendarRequestBuilder) {
    return i117d9ebe24172939d0a5f63702c7509338a57c472472a49ee2565208b4d780bd.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ie82bc68c2150b10b6eac9873b0428f1ecff6323accd09b4b729023c9439cf923.CancelRequestBuilder) {
    return ie82bc68c2150b10b6eac9873b0428f1ecff6323accd09b4b729023c9439cf923.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) Decline()(*i890234a43eddfbb7906d7fc141859ab3228b2bf8edbebe01a56d2e9011f52236.DeclineRequestBuilder) {
    return i890234a43eddfbb7906d7fc141859ab3228b2bf8edbebe01a56d2e9011f52236.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i050fa3cfba1766755e62be09db7509d697402595f146dfa6917836103ea6c418.DismissReminderRequestBuilder) {
    return i050fa3cfba1766755e62be09db7509d697402595f146dfa6917836103ea6c418.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i7ec706215b1a4de83953f27df5698ba1502115937a890507c5a1b1952db5bf2d.ExtensionsRequestBuilder) {
    return i7ec706215b1a4de83953f27df5698ba1502115937a890507c5a1b1952db5bf2d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ieca6997f19b6f68725b8ff66b95f5efaa5c4aba0bd54ed35eab5e4d7ce828403.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ieca6997f19b6f68725b8ff66b95f5efaa5c4aba0bd54ed35eab5e4d7ce828403.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i17325f41fc1dfe0de0e904a3ae38c725efa0d2265ffce8ce6f3018f12c9a4b4e.ForwardRequestBuilder) {
    return i17325f41fc1dfe0de0e904a3ae38c725efa0d2265ffce8ce6f3018f12c9a4b4e.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) Instances()(*i9ece5c52e1d32b599990f2019c6d3a778f13f9f332d1ddbab1abb62e8a20c07d.InstancesRequestBuilder) {
    return i9ece5c52e1d32b599990f2019c6d3a778f13f9f332d1ddbab1abb62e8a20c07d.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*iaca324da655a515ff1fd42e185390f292ff9e89839ba866e74bc5f3d3f127c36.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return iaca324da655a515ff1fd42e185390f292ff9e89839ba866e74bc5f3d3f127c36.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i6382d1d739bdd919528e2d99dc700679ac6ffd1b44d5683e92a1ebce647f9cf9.MultiValueExtendedPropertiesRequestBuilder) {
    return i6382d1d739bdd919528e2d99dc700679ac6ffd1b44d5683e92a1ebce647f9cf9.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ief603bdd91788274032ba500639b4fbbe934af2730bed0cd2ba2199d2b388843.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ief603bdd91788274032ba500639b4fbbe934af2730bed0cd2ba2199d2b388843.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*id814468e09981972f31dbeba72bd3b9bc3d5998b33227510f19290d97aaf643f.SingleValueExtendedPropertiesRequestBuilder) {
    return id814468e09981972f31dbeba72bd3b9bc3d5998b33227510f19290d97aaf643f.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4f3760ebaf8befb9f04aa8e0f72101e132c43a2e798e6ff11f716e9e0a766314.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i4f3760ebaf8befb9f04aa8e0f72101e132c43a2e798e6ff11f716e9e0a766314.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ia816fb6558bf6de9fca63f7aa79137e2a5c51a9979ce0e65066b9062eb2becdd.SnoozeReminderRequestBuilder) {
    return ia816fb6558bf6de9fca63f7aa79137e2a5c51a9979ce0e65066b9062eb2becdd.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i45b9326583488618ead31962b306c42d114f0247787ec0455579853c35a23226.TentativelyAcceptRequestBuilder) {
    return i45b9326583488618ead31962b306c42d114f0247787ec0455579853c35a23226.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
