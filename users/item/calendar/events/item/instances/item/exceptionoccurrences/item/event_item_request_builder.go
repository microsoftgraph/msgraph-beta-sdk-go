package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i02138ebcd8c54ad66fcc1b83c5aba691862df535a3f8cd0e5677dcddeb5322a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i0d5d0b4ad13238336e9693fd8966907ac00388d551ab79ac0d593835617b34ab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/cancel"
    i0fc36d4a45386d570212af1896e142e6704d07e98752280e60a7661f7f5ecc49 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i16df64c8d20d87a6609bcf56f4455b0f9c8fd3cdb27d9a6169b3a2ef60ec8360 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i5cf6d8c5f4f353d9d389158ac22f9dcaae558a31c2f6c99650d644edb641ddb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i684131f8a9a189e3f27ad66b5d5356f66131fe0d85d8b4a4597886a0d28f0fa0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/forward"
    i6be21aff574b2becf273f1fabc83ec951e8ef79c132598c72836600ba7d2c962 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i8b24e9964b7c61504f543f7577e033a90854a5aaee3f167b1dfc0eb75a657124 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/attachments"
    i938486f4acbe372b7aadcf4078813716807e26d4204315e7c962fd0761bd7013 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/accept"
    i9c3a2ce49501f8a32b329eacfb9c472ce25f422938a806c3a50cca37a16ed465 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/calendar"
    id6fc161013bc8ec2332f0381c2efc6de32567458d77dc23487e740c0776c344c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/extensions"
    id79b50e5381bfe03a214a9398ec00b575ecceb92f3af4ffa0278cef9b3f2c4d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/decline"
    i1f332daed746484fef47995403c11dfd599d87f9e0f48d7eb0195e1f0d78461d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i4d1d20ce6486a01a14fbf938b2b458c7c73ba14533b80dcbf2eff8451f3d26ec "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    i9cef18a06a3c57f209883112eedc73f1b6ad564b50bb98be203409d2bd7b36d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    iab4639258e6785c5299a99811e24d3487e45d5c5b1f7831007604d0fad936c90 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i938486f4acbe372b7aadcf4078813716807e26d4204315e7c962fd0761bd7013.AcceptRequestBuilder) {
    return i938486f4acbe372b7aadcf4078813716807e26d4204315e7c962fd0761bd7013.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i8b24e9964b7c61504f543f7577e033a90854a5aaee3f167b1dfc0eb75a657124.AttachmentsRequestBuilder) {
    return i8b24e9964b7c61504f543f7577e033a90854a5aaee3f167b1dfc0eb75a657124.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i9cef18a06a3c57f209883112eedc73f1b6ad564b50bb98be203409d2bd7b36d1.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i9cef18a06a3c57f209883112eedc73f1b6ad564b50bb98be203409d2bd7b36d1.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i9c3a2ce49501f8a32b329eacfb9c472ce25f422938a806c3a50cca37a16ed465.CalendarRequestBuilder) {
    return i9c3a2ce49501f8a32b329eacfb9c472ce25f422938a806c3a50cca37a16ed465.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i0d5d0b4ad13238336e9693fd8966907ac00388d551ab79ac0d593835617b34ab.CancelRequestBuilder) {
    return i0d5d0b4ad13238336e9693fd8966907ac00388d551ab79ac0d593835617b34ab.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*id79b50e5381bfe03a214a9398ec00b575ecceb92f3af4ffa0278cef9b3f2c4d3.DeclineRequestBuilder) {
    return id79b50e5381bfe03a214a9398ec00b575ecceb92f3af4ffa0278cef9b3f2c4d3.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i0fc36d4a45386d570212af1896e142e6704d07e98752280e60a7661f7f5ecc49.DismissReminderRequestBuilder) {
    return i0fc36d4a45386d570212af1896e142e6704d07e98752280e60a7661f7f5ecc49.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*id6fc161013bc8ec2332f0381c2efc6de32567458d77dc23487e740c0776c344c.ExtensionsRequestBuilder) {
    return id6fc161013bc8ec2332f0381c2efc6de32567458d77dc23487e740c0776c344c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i4d1d20ce6486a01a14fbf938b2b458c7c73ba14533b80dcbf2eff8451f3d26ec.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i4d1d20ce6486a01a14fbf938b2b458c7c73ba14533b80dcbf2eff8451f3d26ec.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i684131f8a9a189e3f27ad66b5d5356f66131fe0d85d8b4a4597886a0d28f0fa0.ForwardRequestBuilder) {
    return i684131f8a9a189e3f27ad66b5d5356f66131fe0d85d8b4a4597886a0d28f0fa0.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i02138ebcd8c54ad66fcc1b83c5aba691862df535a3f8cd0e5677dcddeb5322a8.MultiValueExtendedPropertiesRequestBuilder) {
    return i02138ebcd8c54ad66fcc1b83c5aba691862df535a3f8cd0e5677dcddeb5322a8.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i1f332daed746484fef47995403c11dfd599d87f9e0f48d7eb0195e1f0d78461d.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i1f332daed746484fef47995403c11dfd599d87f9e0f48d7eb0195e1f0d78461d.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i6be21aff574b2becf273f1fabc83ec951e8ef79c132598c72836600ba7d2c962.SingleValueExtendedPropertiesRequestBuilder) {
    return i6be21aff574b2becf273f1fabc83ec951e8ef79c132598c72836600ba7d2c962.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iab4639258e6785c5299a99811e24d3487e45d5c5b1f7831007604d0fad936c90.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return iab4639258e6785c5299a99811e24d3487e45d5c5b1f7831007604d0fad936c90.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i5cf6d8c5f4f353d9d389158ac22f9dcaae558a31c2f6c99650d644edb641ddb8.SnoozeReminderRequestBuilder) {
    return i5cf6d8c5f4f353d9d389158ac22f9dcaae558a31c2f6c99650d644edb641ddb8.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i16df64c8d20d87a6609bcf56f4455b0f9c8fd3cdb27d9a6169b3a2ef60ec8360.TentativelyAcceptRequestBuilder) {
    return i16df64c8d20d87a6609bcf56f4455b0f9c8fd3cdb27d9a6169b3a2ef60ec8360.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
