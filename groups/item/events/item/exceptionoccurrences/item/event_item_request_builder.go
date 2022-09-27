package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0e63c98082b16d34427276844ba48abea6f4a4df04dc240935637148c0128228 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/tentativelyaccept"
    i14761d40f4ea6381da8a8733ac61011e3db7d7b15cfc5aaf7f5d3404bf7c3d67 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/calendar"
    i54ffd054a501e8df2e6ee11f0b07bb762c2f783afbafbb248a1b1d1135f62e1e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i62577b4dadaa090eca88460c637ed911f14d6024cdebebcbecc6591db22e7096 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/cancel"
    i6e80cff0f4ab82a01472195edd354e751e0817994c6eea624f4d8eecebdfd93e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/dismissreminder"
    iacd403a6a900f1e55bd7fe438f49e2a4bf5b5051112e3fbcc63ef98a167fc028 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances"
    ic333bab5725eafd70731cd8ad5c7648759f53cff523dece27cda0b809f809576 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/forward"
    ic74dd51ffaf7075bdc9cda367e10a2040083334f5003182af42fb71f261949c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/snoozereminder"
    icf60a8e6df2767e998560595845b1f272d9f9e194f6473e114b8d7c6fc6d0bea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/accept"
    id0236475c0793f7e46ee7454719fd55c76add053c6c74899c821514313501a0f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/decline"
    ie4665a795c926980f4c9279fda9a00cca2d00e9899401a6b383d954fca19a50e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/extensions"
    ie9f7c64c2bcbcdcae69a43a35b8d19510edefe345493912588894e7dfeb7cd4b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    if0c7358bbc864648a0924622cf54bd4b4bf83ade2d9a7d3edafca829c721230c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/attachments"
    i1b5b8745445c4c02a9c94937f36e007bd63773c8121e24bc84b5b9dbb274a5d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/instances/item"
    i2e5b6a6924eb5cbc9a6b15ad5c05de91721ae4de67dd7447a37dbc98b18a894c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/extensions/item"
    i6e97d308af981e53de9342506001a0584abbc966e22837194cb53d1ca18f0257 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ie56ffcf5015fb2fc001270158e778df63af920cee01bb97f2eee4bc6da6c7f97 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/attachments/item"
    if5b70906256cbf7a0066c664c28f947f1263867315c750675f07f9d3f67052d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*icf60a8e6df2767e998560595845b1f272d9f9e194f6473e114b8d7c6fc6d0bea.AcceptRequestBuilder) {
    return icf60a8e6df2767e998560595845b1f272d9f9e194f6473e114b8d7c6fc6d0bea.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*if0c7358bbc864648a0924622cf54bd4b4bf83ade2d9a7d3edafca829c721230c.AttachmentsRequestBuilder) {
    return if0c7358bbc864648a0924622cf54bd4b4bf83ade2d9a7d3edafca829c721230c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ie56ffcf5015fb2fc001270158e778df63af920cee01bb97f2eee4bc6da6c7f97.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ie56ffcf5015fb2fc001270158e778df63af920cee01bb97f2eee4bc6da6c7f97.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i14761d40f4ea6381da8a8733ac61011e3db7d7b15cfc5aaf7f5d3404bf7c3d67.CalendarRequestBuilder) {
    return i14761d40f4ea6381da8a8733ac61011e3db7d7b15cfc5aaf7f5d3404bf7c3d67.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i62577b4dadaa090eca88460c637ed911f14d6024cdebebcbecc6591db22e7096.CancelRequestBuilder) {
    return i62577b4dadaa090eca88460c637ed911f14d6024cdebebcbecc6591db22e7096.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Decline()(*id0236475c0793f7e46ee7454719fd55c76add053c6c74899c821514313501a0f.DeclineRequestBuilder) {
    return id0236475c0793f7e46ee7454719fd55c76add053c6c74899c821514313501a0f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i6e80cff0f4ab82a01472195edd354e751e0817994c6eea624f4d8eecebdfd93e.DismissReminderRequestBuilder) {
    return i6e80cff0f4ab82a01472195edd354e751e0817994c6eea624f4d8eecebdfd93e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ie4665a795c926980f4c9279fda9a00cca2d00e9899401a6b383d954fca19a50e.ExtensionsRequestBuilder) {
    return ie4665a795c926980f4c9279fda9a00cca2d00e9899401a6b383d954fca19a50e.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i2e5b6a6924eb5cbc9a6b15ad5c05de91721ae4de67dd7447a37dbc98b18a894c.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i2e5b6a6924eb5cbc9a6b15ad5c05de91721ae4de67dd7447a37dbc98b18a894c.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ic333bab5725eafd70731cd8ad5c7648759f53cff523dece27cda0b809f809576.ForwardRequestBuilder) {
    return ic333bab5725eafd70731cd8ad5c7648759f53cff523dece27cda0b809f809576.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) Instances()(*iacd403a6a900f1e55bd7fe438f49e2a4bf5b5051112e3fbcc63ef98a167fc028.InstancesRequestBuilder) {
    return iacd403a6a900f1e55bd7fe438f49e2a4bf5b5051112e3fbcc63ef98a167fc028.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i1b5b8745445c4c02a9c94937f36e007bd63773c8121e24bc84b5b9dbb274a5d4.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i1b5b8745445c4c02a9c94937f36e007bd63773c8121e24bc84b5b9dbb274a5d4.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie9f7c64c2bcbcdcae69a43a35b8d19510edefe345493912588894e7dfeb7cd4b.MultiValueExtendedPropertiesRequestBuilder) {
    return ie9f7c64c2bcbcdcae69a43a35b8d19510edefe345493912588894e7dfeb7cd4b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if5b70906256cbf7a0066c664c28f947f1263867315c750675f07f9d3f67052d4.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return if5b70906256cbf7a0066c664c28f947f1263867315c750675f07f9d3f67052d4.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i54ffd054a501e8df2e6ee11f0b07bb762c2f783afbafbb248a1b1d1135f62e1e.SingleValueExtendedPropertiesRequestBuilder) {
    return i54ffd054a501e8df2e6ee11f0b07bb762c2f783afbafbb248a1b1d1135f62e1e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i6e97d308af981e53de9342506001a0584abbc966e22837194cb53d1ca18f0257.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i6e97d308af981e53de9342506001a0584abbc966e22837194cb53d1ca18f0257.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ic74dd51ffaf7075bdc9cda367e10a2040083334f5003182af42fb71f261949c5.SnoozeReminderRequestBuilder) {
    return ic74dd51ffaf7075bdc9cda367e10a2040083334f5003182af42fb71f261949c5.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i0e63c98082b16d34427276844ba48abea6f4a4df04dc240935637148c0128228.TentativelyAcceptRequestBuilder) {
    return i0e63c98082b16d34427276844ba48abea6f4a4df04dc240935637148c0128228.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
