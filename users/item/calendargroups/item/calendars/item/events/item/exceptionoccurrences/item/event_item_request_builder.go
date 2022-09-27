package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i11a2298f817d20027061f562981fdfd8f1fb9f300ee82790dbc6ed294d808318 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    i3c0c53b8c8d6a38cfb4486f1b5534946c9d68076720994aa4220c3fbfac42c3b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/calendar"
    i3d0ff5611157d7552ee59914cd4e7114ca57186b4d11f03168ca84ddc1849481 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/attachments"
    i473cf07cd47eca0525deb0d000e5667e706a12ed9107cf7a10a98b1fef601e6c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances"
    i5593685ccf7132f2b322d72c7f6190a49e272ba79796551cb489be2b74af524f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/cancel"
    i57332911775b2c8d06d7a5ec6a0181b964d10ead1a8ad0594e53d7bf944092a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/tentativelyaccept"
    i7f72755c435a384033bee52a976326640bd6d1b3fe52ec1d522d1718f3d6acd5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/snoozereminder"
    i8115d40777816dc52915a104ff18869f165cd32f3ae8ecb2740a6a9cbdf2af2c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/forward"
    i945aeffc35337602807e6b9c73c243730e3651ec6bc0045fd68b850242667045 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/extensions"
    i9c7f3e4cde7019b659d5a1eb0156635e2a87eacffdd75a9352dfb7debde4d289 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ia16eaac4103fc4b7aff74f7c928df3fa66c1acc3bd5976fc3c93c29176030028 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/dismissreminder"
    ia9bf548a1db34e1a13b86e5b5f8f7297f8c5c9cdf08b00efaf8f47edc57c530d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/decline"
    id92bca78a768b46b0a9ae7f9bb4158b0042091ffb95f227d3e6174994267a64b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/accept"
    i0c187b5ad81f2594b6848896d53784fcc719c6e87a858ba91876cc777c5ccf1f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/extensions/item"
    i209698a62df63a99c74c7376511b6d963f2072196d726fbcc8d45354cf9cef35 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item"
    i3e3a17e24e7d47b4f7b401638ca3732888fc9b1c6fa3cace56f507e4b8a2425d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i63ea09eeedba8bbcb11f1d743dc373f10e2da0df976be1e1aca5a3b76ee1558e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    iec2c6029b139d1e3fcd38fba8597cacf78f292ab71c39989975656787a54d7dc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*id92bca78a768b46b0a9ae7f9bb4158b0042091ffb95f227d3e6174994267a64b.AcceptRequestBuilder) {
    return id92bca78a768b46b0a9ae7f9bb4158b0042091ffb95f227d3e6174994267a64b.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i3d0ff5611157d7552ee59914cd4e7114ca57186b4d11f03168ca84ddc1849481.AttachmentsRequestBuilder) {
    return i3d0ff5611157d7552ee59914cd4e7114ca57186b4d11f03168ca84ddc1849481.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*iec2c6029b139d1e3fcd38fba8597cacf78f292ab71c39989975656787a54d7dc.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return iec2c6029b139d1e3fcd38fba8597cacf78f292ab71c39989975656787a54d7dc.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i3c0c53b8c8d6a38cfb4486f1b5534946c9d68076720994aa4220c3fbfac42c3b.CalendarRequestBuilder) {
    return i3c0c53b8c8d6a38cfb4486f1b5534946c9d68076720994aa4220c3fbfac42c3b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i5593685ccf7132f2b322d72c7f6190a49e272ba79796551cb489be2b74af524f.CancelRequestBuilder) {
    return i5593685ccf7132f2b322d72c7f6190a49e272ba79796551cb489be2b74af524f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*ia9bf548a1db34e1a13b86e5b5f8f7297f8c5c9cdf08b00efaf8f47edc57c530d.DeclineRequestBuilder) {
    return ia9bf548a1db34e1a13b86e5b5f8f7297f8c5c9cdf08b00efaf8f47edc57c530d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*ia16eaac4103fc4b7aff74f7c928df3fa66c1acc3bd5976fc3c93c29176030028.DismissReminderRequestBuilder) {
    return ia16eaac4103fc4b7aff74f7c928df3fa66c1acc3bd5976fc3c93c29176030028.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i945aeffc35337602807e6b9c73c243730e3651ec6bc0045fd68b850242667045.ExtensionsRequestBuilder) {
    return i945aeffc35337602807e6b9c73c243730e3651ec6bc0045fd68b850242667045.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i0c187b5ad81f2594b6848896d53784fcc719c6e87a858ba91876cc777c5ccf1f.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i0c187b5ad81f2594b6848896d53784fcc719c6e87a858ba91876cc777c5ccf1f.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i8115d40777816dc52915a104ff18869f165cd32f3ae8ecb2740a6a9cbdf2af2c.ForwardRequestBuilder) {
    return i8115d40777816dc52915a104ff18869f165cd32f3ae8ecb2740a6a9cbdf2af2c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*i473cf07cd47eca0525deb0d000e5667e706a12ed9107cf7a10a98b1fef601e6c.InstancesRequestBuilder) {
    return i473cf07cd47eca0525deb0d000e5667e706a12ed9107cf7a10a98b1fef601e6c.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i209698a62df63a99c74c7376511b6d963f2072196d726fbcc8d45354cf9cef35.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i209698a62df63a99c74c7376511b6d963f2072196d726fbcc8d45354cf9cef35.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i11a2298f817d20027061f562981fdfd8f1fb9f300ee82790dbc6ed294d808318.MultiValueExtendedPropertiesRequestBuilder) {
    return i11a2298f817d20027061f562981fdfd8f1fb9f300ee82790dbc6ed294d808318.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i63ea09eeedba8bbcb11f1d743dc373f10e2da0df976be1e1aca5a3b76ee1558e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i63ea09eeedba8bbcb11f1d743dc373f10e2da0df976be1e1aca5a3b76ee1558e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i9c7f3e4cde7019b659d5a1eb0156635e2a87eacffdd75a9352dfb7debde4d289.SingleValueExtendedPropertiesRequestBuilder) {
    return i9c7f3e4cde7019b659d5a1eb0156635e2a87eacffdd75a9352dfb7debde4d289.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i3e3a17e24e7d47b4f7b401638ca3732888fc9b1c6fa3cace56f507e4b8a2425d.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i3e3a17e24e7d47b4f7b401638ca3732888fc9b1c6fa3cace56f507e4b8a2425d.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i7f72755c435a384033bee52a976326640bd6d1b3fe52ec1d522d1718f3d6acd5.SnoozeReminderRequestBuilder) {
    return i7f72755c435a384033bee52a976326640bd6d1b3fe52ec1d522d1718f3d6acd5.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i57332911775b2c8d06d7a5ec6a0181b964d10ead1a8ad0594e53d7bf944092a7.TentativelyAcceptRequestBuilder) {
    return i57332911775b2c8d06d7a5ec6a0181b964d10ead1a8ad0594e53d7bf944092a7.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
