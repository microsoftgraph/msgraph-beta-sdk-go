package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2764131d6064d55a6fb28f2c0abcae47736dde3f40e965b0eda7a971dde5de10 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i27a3a352b6ff7d9499e5d09898d04ce38661929729d0d9a6d496952e9f70c30e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i34b1fded7e2d95e2cdfa2275ee4419128351ddd305cd95a46e6b51cc48b67976 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i3574e76278c0a6f91b08c48f88683e4dba058986d9592c2aba37b074d381c179 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i5ee649c1233ae64934fe34f5783abab787512a61101c1f720f81b969f0b1decb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    i665ca788afaaae309174146e34bb77894cf47bb2bba420c30a1eb0912158594d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i991412ae494f86f4d381089968289cae90002d2cb421d0f94c320e30b5561d25 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ibed6e83805e6013a18a6e19fbddd1cbd35347d5c1914d17080c3413d73bcbba9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    id19ec0743b30d2cbd1758bd779f1e99ddc91c5b14dc665e6c945eaf4623aa27d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    ida7bdbd0c64c580394b466feda1a6af9066a6919f0535f64131a9a1866d55c3b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    idbf0a00c08b54cfd7fac35e38445dc93328b6d746213307f32af7de5b59220ac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    ie102b9d30203f964bbf8844901a37d39d3c1917d954fe13a686d3f0efd36a9e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i3f862a33674900febc35502d782a7c3b473a6235fe1834b43fa3e93a0d3756c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i46222f284134be595ab501e2db9e55e6883b67f43c32ddf4bd9a55f9d96d2744 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    i8883f5c2ffbaa391bb23a2ccdb377fcc9971efd46b510c8a5d361908abfa4ab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    ib8c9f1b1d8fc0e2d1b4c7f619e86de8a7306494c87eccea322c28ca34745d939 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i5ee649c1233ae64934fe34f5783abab787512a61101c1f720f81b969f0b1decb.AcceptRequestBuilder) {
    return i5ee649c1233ae64934fe34f5783abab787512a61101c1f720f81b969f0b1decb.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i27a3a352b6ff7d9499e5d09898d04ce38661929729d0d9a6d496952e9f70c30e.AttachmentsRequestBuilder) {
    return i27a3a352b6ff7d9499e5d09898d04ce38661929729d0d9a6d496952e9f70c30e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i8883f5c2ffbaa391bb23a2ccdb377fcc9971efd46b510c8a5d361908abfa4ab5.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i8883f5c2ffbaa391bb23a2ccdb377fcc9971efd46b510c8a5d361908abfa4ab5.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ida7bdbd0c64c580394b466feda1a6af9066a6919f0535f64131a9a1866d55c3b.CalendarRequestBuilder) {
    return ida7bdbd0c64c580394b466feda1a6af9066a6919f0535f64131a9a1866d55c3b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i2764131d6064d55a6fb28f2c0abcae47736dde3f40e965b0eda7a971dde5de10.CancelRequestBuilder) {
    return i2764131d6064d55a6fb28f2c0abcae47736dde3f40e965b0eda7a971dde5de10.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*id19ec0743b30d2cbd1758bd779f1e99ddc91c5b14dc665e6c945eaf4623aa27d.DeclineRequestBuilder) {
    return id19ec0743b30d2cbd1758bd779f1e99ddc91c5b14dc665e6c945eaf4623aa27d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i34b1fded7e2d95e2cdfa2275ee4419128351ddd305cd95a46e6b51cc48b67976.DismissReminderRequestBuilder) {
    return i34b1fded7e2d95e2cdfa2275ee4419128351ddd305cd95a46e6b51cc48b67976.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i665ca788afaaae309174146e34bb77894cf47bb2bba420c30a1eb0912158594d.ExtensionsRequestBuilder) {
    return i665ca788afaaae309174146e34bb77894cf47bb2bba420c30a1eb0912158594d.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i46222f284134be595ab501e2db9e55e6883b67f43c32ddf4bd9a55f9d96d2744.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i46222f284134be595ab501e2db9e55e6883b67f43c32ddf4bd9a55f9d96d2744.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ibed6e83805e6013a18a6e19fbddd1cbd35347d5c1914d17080c3413d73bcbba9.ForwardRequestBuilder) {
    return ibed6e83805e6013a18a6e19fbddd1cbd35347d5c1914d17080c3413d73bcbba9.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie102b9d30203f964bbf8844901a37d39d3c1917d954fe13a686d3f0efd36a9e5.MultiValueExtendedPropertiesRequestBuilder) {
    return ie102b9d30203f964bbf8844901a37d39d3c1917d954fe13a686d3f0efd36a9e5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib8c9f1b1d8fc0e2d1b4c7f619e86de8a7306494c87eccea322c28ca34745d939.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ib8c9f1b1d8fc0e2d1b4c7f619e86de8a7306494c87eccea322c28ca34745d939.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i991412ae494f86f4d381089968289cae90002d2cb421d0f94c320e30b5561d25.SingleValueExtendedPropertiesRequestBuilder) {
    return i991412ae494f86f4d381089968289cae90002d2cb421d0f94c320e30b5561d25.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i3f862a33674900febc35502d782a7c3b473a6235fe1834b43fa3e93a0d3756c3.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i3f862a33674900febc35502d782a7c3b473a6235fe1834b43fa3e93a0d3756c3.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*idbf0a00c08b54cfd7fac35e38445dc93328b6d746213307f32af7de5b59220ac.SnoozeReminderRequestBuilder) {
    return idbf0a00c08b54cfd7fac35e38445dc93328b6d746213307f32af7de5b59220ac.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i3574e76278c0a6f91b08c48f88683e4dba058986d9592c2aba37b074d381c179.TentativelyAcceptRequestBuilder) {
    return i3574e76278c0a6f91b08c48f88683e4dba058986d9592c2aba37b074d381c179.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
