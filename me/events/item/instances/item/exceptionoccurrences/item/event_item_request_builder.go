package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i235c5ce0d1f74d524ef46dbde3ad46ff72b77628085dbc6bfdac40c1fb48c23c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i29a8545864fc805f0e2701a8b3eec5af1142195c0187b2c6968094d3b343fc07 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/accept"
    i3bd23c78fcfbc6b0f4baf49aa2633f97e3d708ae8375a9cce895ad4a6a73a76a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/attachments"
    i5e5077c321fc6e74b2a81f7770dae70a1666eda8cfcd968f487525cb202d1690 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/forward"
    i62af91d8429dd5f4cb597ef41ed6e71511fb493b4b5fcbff8a4e59b7b859e6be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i70176bee18712bf6effaf06a937591a269c52462ddd3254f0c639fa8193e8257 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/decline"
    i76a3e45766626f05b2e1bfda4f0a8e614736e1ed200ed9d1ecd953a7c21f29b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i782c18674076087ebdcded13180ab6c2573441036e3b850dbacd6508343032fb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/extensions"
    i782f63d0c32eb12865e32070ff0d30a940280f9a88402ee6b94497170ac21eda "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i81438a5134e8af20e7bb64724e4a869db0f1040f8c7d4b6651b4a8c4e66758bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/cancel"
    iaef4803cc86320b3f857eae04b3c7a0ea4111e2a487f8e3cc015656cfbd2f476 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/calendar"
    ibe378cc2630ababf55a71d5b746f507d3433811e5b1747b548c290c5f774ebd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i5a7159f62dfc28b5ce002654b03796b6b64a8bf6334f48f84d36040205821d66 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    i6030ecc1ae42e4cc03f29341a683734a1c1a632a2e1bc1df69f744da21821d07 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    iacd3eefe1e8b57f639bf54342be2037e4646af4aa9296f8557c19be991af64bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    ife47cecb8e20b7ce7a99756a6a698281988672e18019307b3efe508888854f1e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i29a8545864fc805f0e2701a8b3eec5af1142195c0187b2c6968094d3b343fc07.AcceptRequestBuilder) {
    return i29a8545864fc805f0e2701a8b3eec5af1142195c0187b2c6968094d3b343fc07.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i3bd23c78fcfbc6b0f4baf49aa2633f97e3d708ae8375a9cce895ad4a6a73a76a.AttachmentsRequestBuilder) {
    return i3bd23c78fcfbc6b0f4baf49aa2633f97e3d708ae8375a9cce895ad4a6a73a76a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i5a7159f62dfc28b5ce002654b03796b6b64a8bf6334f48f84d36040205821d66.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i5a7159f62dfc28b5ce002654b03796b6b64a8bf6334f48f84d36040205821d66.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*iaef4803cc86320b3f857eae04b3c7a0ea4111e2a487f8e3cc015656cfbd2f476.CalendarRequestBuilder) {
    return iaef4803cc86320b3f857eae04b3c7a0ea4111e2a487f8e3cc015656cfbd2f476.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i81438a5134e8af20e7bb64724e4a869db0f1040f8c7d4b6651b4a8c4e66758bd.CancelRequestBuilder) {
    return i81438a5134e8af20e7bb64724e4a869db0f1040f8c7d4b6651b4a8c4e66758bd.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i70176bee18712bf6effaf06a937591a269c52462ddd3254f0c639fa8193e8257.DeclineRequestBuilder) {
    return i70176bee18712bf6effaf06a937591a269c52462ddd3254f0c639fa8193e8257.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i235c5ce0d1f74d524ef46dbde3ad46ff72b77628085dbc6bfdac40c1fb48c23c.DismissReminderRequestBuilder) {
    return i235c5ce0d1f74d524ef46dbde3ad46ff72b77628085dbc6bfdac40c1fb48c23c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i782c18674076087ebdcded13180ab6c2573441036e3b850dbacd6508343032fb.ExtensionsRequestBuilder) {
    return i782c18674076087ebdcded13180ab6c2573441036e3b850dbacd6508343032fb.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*iacd3eefe1e8b57f639bf54342be2037e4646af4aa9296f8557c19be991af64bd.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return iacd3eefe1e8b57f639bf54342be2037e4646af4aa9296f8557c19be991af64bd.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i5e5077c321fc6e74b2a81f7770dae70a1666eda8cfcd968f487525cb202d1690.ForwardRequestBuilder) {
    return i5e5077c321fc6e74b2a81f7770dae70a1666eda8cfcd968f487525cb202d1690.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i62af91d8429dd5f4cb597ef41ed6e71511fb493b4b5fcbff8a4e59b7b859e6be.MultiValueExtendedPropertiesRequestBuilder) {
    return i62af91d8429dd5f4cb597ef41ed6e71511fb493b4b5fcbff8a4e59b7b859e6be.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ife47cecb8e20b7ce7a99756a6a698281988672e18019307b3efe508888854f1e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ife47cecb8e20b7ce7a99756a6a698281988672e18019307b3efe508888854f1e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i76a3e45766626f05b2e1bfda4f0a8e614736e1ed200ed9d1ecd953a7c21f29b5.SingleValueExtendedPropertiesRequestBuilder) {
    return i76a3e45766626f05b2e1bfda4f0a8e614736e1ed200ed9d1ecd953a7c21f29b5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i6030ecc1ae42e4cc03f29341a683734a1c1a632a2e1bc1df69f744da21821d07.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i6030ecc1ae42e4cc03f29341a683734a1c1a632a2e1bc1df69f744da21821d07.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i782f63d0c32eb12865e32070ff0d30a940280f9a88402ee6b94497170ac21eda.SnoozeReminderRequestBuilder) {
    return i782f63d0c32eb12865e32070ff0d30a940280f9a88402ee6b94497170ac21eda.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ibe378cc2630ababf55a71d5b746f507d3433811e5b1747b548c290c5f774ebd7.TentativelyAcceptRequestBuilder) {
    return ibe378cc2630ababf55a71d5b746f507d3433811e5b1747b548c290c5f774ebd7.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
