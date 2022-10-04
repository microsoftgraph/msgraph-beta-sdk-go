package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i00e8cee33d801300cbb4ad14f617a9ae72146de6e7faf2e3173a8e87c5fd4177 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i06a51171e29aa4d9148d44bc6fc0133efcf2139fc7d2eebac311e0c1659c40f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i0d6bcc2245d4301efd92326df96c4309b1ec268749435c05c22b43916445e60e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i10588310ef114a7303177041c9c7b6b2f05c88762f0e11f13b439a739f31b463 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    i1fc4130db0da2d667f018c8ade40a0f1158d2ffd7a10ef4f9bcdc3d3a0480cbb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i37ba7a547c8c3eea8c32baf13b26626b4a51a51c9a248edecfe0ca438e1f7b8f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    i87c7eaf268f775c95bc5c8ee3af4a2b44a32bd00cd4a81802b0380502c277ea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    ica13a7ef2877dfc068e740a476213a8a2e737cf87469a7df060da2177eea40bf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    idd791a8914249115db63545178c54a9d142f11d5f2db25ebebaa585b33141342 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    iede12bcee681e4c6d9fd1c0b33b30675372f8ce158f56cbcc6be12fdf08fce5e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    ifa190757e759a10d6f019a4cc23b89bdbab7075416a2b8140da20fb74d997428 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    ifa80757feb55d099fda174fcd1fb4a480fe4651215ec7931aa86494b52d883be "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    i414b8a28da16e4dec436ebff6d2662881f1ff22e8de7559f1d2cf429edf0e953 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i443776f5b8c58e75c22569c39fe5bdefd83bab185878ad8105e7bd45850bd519 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    i7af0f6c45539d2dad182638a03cd0fa8eae8bdeb9fee4192d4d90198894ef83f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    iacaf2fb0f3d2d0e6b04e2fdf45c115097b8dde9ccae1c353dc8a2a0efaf00af5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i10588310ef114a7303177041c9c7b6b2f05c88762f0e11f13b439a739f31b463.AcceptRequestBuilder) {
    return i10588310ef114a7303177041c9c7b6b2f05c88762f0e11f13b439a739f31b463.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i0d6bcc2245d4301efd92326df96c4309b1ec268749435c05c22b43916445e60e.AttachmentsRequestBuilder) {
    return i0d6bcc2245d4301efd92326df96c4309b1ec268749435c05c22b43916445e60e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i443776f5b8c58e75c22569c39fe5bdefd83bab185878ad8105e7bd45850bd519.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i443776f5b8c58e75c22569c39fe5bdefd83bab185878ad8105e7bd45850bd519.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i37ba7a547c8c3eea8c32baf13b26626b4a51a51c9a248edecfe0ca438e1f7b8f.CalendarRequestBuilder) {
    return i37ba7a547c8c3eea8c32baf13b26626b4a51a51c9a248edecfe0ca438e1f7b8f.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ifa190757e759a10d6f019a4cc23b89bdbab7075416a2b8140da20fb74d997428.CancelRequestBuilder) {
    return ifa190757e759a10d6f019a4cc23b89bdbab7075416a2b8140da20fb74d997428.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*ifa80757feb55d099fda174fcd1fb4a480fe4651215ec7931aa86494b52d883be.DeclineRequestBuilder) {
    return ifa80757feb55d099fda174fcd1fb4a480fe4651215ec7931aa86494b52d883be.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i00e8cee33d801300cbb4ad14f617a9ae72146de6e7faf2e3173a8e87c5fd4177.DismissReminderRequestBuilder) {
    return i00e8cee33d801300cbb4ad14f617a9ae72146de6e7faf2e3173a8e87c5fd4177.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i87c7eaf268f775c95bc5c8ee3af4a2b44a32bd00cd4a81802b0380502c277ea1.ExtensionsRequestBuilder) {
    return i87c7eaf268f775c95bc5c8ee3af4a2b44a32bd00cd4a81802b0380502c277ea1.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i7af0f6c45539d2dad182638a03cd0fa8eae8bdeb9fee4192d4d90198894ef83f.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i7af0f6c45539d2dad182638a03cd0fa8eae8bdeb9fee4192d4d90198894ef83f.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ica13a7ef2877dfc068e740a476213a8a2e737cf87469a7df060da2177eea40bf.ForwardRequestBuilder) {
    return ica13a7ef2877dfc068e740a476213a8a2e737cf87469a7df060da2177eea40bf.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iede12bcee681e4c6d9fd1c0b33b30675372f8ce158f56cbcc6be12fdf08fce5e.MultiValueExtendedPropertiesRequestBuilder) {
    return iede12bcee681e4c6d9fd1c0b33b30675372f8ce158f56cbcc6be12fdf08fce5e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iacaf2fb0f3d2d0e6b04e2fdf45c115097b8dde9ccae1c353dc8a2a0efaf00af5.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iacaf2fb0f3d2d0e6b04e2fdf45c115097b8dde9ccae1c353dc8a2a0efaf00af5.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i1fc4130db0da2d667f018c8ade40a0f1158d2ffd7a10ef4f9bcdc3d3a0480cbb.SingleValueExtendedPropertiesRequestBuilder) {
    return i1fc4130db0da2d667f018c8ade40a0f1158d2ffd7a10ef4f9bcdc3d3a0480cbb.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i414b8a28da16e4dec436ebff6d2662881f1ff22e8de7559f1d2cf429edf0e953.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i414b8a28da16e4dec436ebff6d2662881f1ff22e8de7559f1d2cf429edf0e953.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*idd791a8914249115db63545178c54a9d142f11d5f2db25ebebaa585b33141342.SnoozeReminderRequestBuilder) {
    return idd791a8914249115db63545178c54a9d142f11d5f2db25ebebaa585b33141342.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i06a51171e29aa4d9148d44bc6fc0133efcf2139fc7d2eebac311e0c1659c40f1.TentativelyAcceptRequestBuilder) {
    return i06a51171e29aa4d9148d44bc6fc0133efcf2139fc7d2eebac311e0c1659c40f1.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
