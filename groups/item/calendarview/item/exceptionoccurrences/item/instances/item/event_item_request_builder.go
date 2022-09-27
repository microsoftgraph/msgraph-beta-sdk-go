package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03c355ed068c6a7385daac9b4ba7c3851e323e658f20d431199852de45ba5781 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    i08ca72de4e866621a9949969339e87ad0db291ba6cbdbc0b31f586ad0dd104cb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i0d53845ad4c9b02308b537fa5937785c1fb91e1428dc458647c74f370fbe2d20 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i24564f34e5353ab2405e9efe2d9c146f52164f1068ddb72a6e02cdb2e17c1cfe "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i3bca677d38b3cc9593ff85eef6eb2d6e33cb0078ca494e7894dd521d844bd5c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    i3e6e3f7ba4a6f979cfccef30e453f7f6041fe3df0a0ea4755bb95f20f0e33b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    i461ef239e3efa7fc9d788c15aa550acea79bc948a1d8a0a3780e2a243def986a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    i8995f8a7bf29894e3868e11efb83cf43c08586f5789d309a1a1ef2ab7a10498a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i9723713f3dbd5ad48f1d184a1e7d1d6168c376b6a26723867d99ad7310169dbf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    ic07285bc27b4cdd05809b8320be3252542ec79f9a72a507850a252f3aba78a36 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    id3e743259b7650f67d2c5ddff1de9ce83493ecabe7fc082d3a4ef76c36dab0fa "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    ie493b644164faa2185855a98cba00f87c626da12fd9f0b2a842d5486378c8c32 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    i4e45ca22802f537c290d0cf43a7c92104d693740acea383f603c3294ef368ee9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    iaa5bd559d854cdb092a6e723f97a6ad726ee0a26fa54f91b67992e1430d83694 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    ibe15a1df40d7d55b7e223449604eba15892bea99e19367d221786e2e52d2a87b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
    iee606450eb9c7d7f0ce88e813fdb68878885faa2e99762f240a53bed95ed4c18 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
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
func (m *EventItemRequestBuilder) Accept()(*id3e743259b7650f67d2c5ddff1de9ce83493ecabe7fc082d3a4ef76c36dab0fa.AcceptRequestBuilder) {
    return id3e743259b7650f67d2c5ddff1de9ce83493ecabe7fc082d3a4ef76c36dab0fa.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ie493b644164faa2185855a98cba00f87c626da12fd9f0b2a842d5486378c8c32.AttachmentsRequestBuilder) {
    return ie493b644164faa2185855a98cba00f87c626da12fd9f0b2a842d5486378c8c32.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ibe15a1df40d7d55b7e223449604eba15892bea99e19367d221786e2e52d2a87b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ibe15a1df40d7d55b7e223449604eba15892bea99e19367d221786e2e52d2a87b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i461ef239e3efa7fc9d788c15aa550acea79bc948a1d8a0a3780e2a243def986a.CalendarRequestBuilder) {
    return i461ef239e3efa7fc9d788c15aa550acea79bc948a1d8a0a3780e2a243def986a.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i3bca677d38b3cc9593ff85eef6eb2d6e33cb0078ca494e7894dd521d844bd5c5.CancelRequestBuilder) {
    return i3bca677d38b3cc9593ff85eef6eb2d6e33cb0078ca494e7894dd521d844bd5c5.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Decline()(*i9723713f3dbd5ad48f1d184a1e7d1d6168c376b6a26723867d99ad7310169dbf.DeclineRequestBuilder) {
    return i9723713f3dbd5ad48f1d184a1e7d1d6168c376b6a26723867d99ad7310169dbf.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i0d53845ad4c9b02308b537fa5937785c1fb91e1428dc458647c74f370fbe2d20.DismissReminderRequestBuilder) {
    return i0d53845ad4c9b02308b537fa5937785c1fb91e1428dc458647c74f370fbe2d20.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i03c355ed068c6a7385daac9b4ba7c3851e323e658f20d431199852de45ba5781.ExtensionsRequestBuilder) {
    return i03c355ed068c6a7385daac9b4ba7c3851e323e658f20d431199852de45ba5781.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*iaa5bd559d854cdb092a6e723f97a6ad726ee0a26fa54f91b67992e1430d83694.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return iaa5bd559d854cdb092a6e723f97a6ad726ee0a26fa54f91b67992e1430d83694.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i3e6e3f7ba4a6f979cfccef30e453f7f6041fe3df0a0ea4755bb95f20f0e33b3c.ForwardRequestBuilder) {
    return i3e6e3f7ba4a6f979cfccef30e453f7f6041fe3df0a0ea4755bb95f20f0e33b3c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i08ca72de4e866621a9949969339e87ad0db291ba6cbdbc0b31f586ad0dd104cb.MultiValueExtendedPropertiesRequestBuilder) {
    return i08ca72de4e866621a9949969339e87ad0db291ba6cbdbc0b31f586ad0dd104cb.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4e45ca22802f537c290d0cf43a7c92104d693740acea383f603c3294ef368ee9.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i4e45ca22802f537c290d0cf43a7c92104d693740acea383f603c3294ef368ee9.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i24564f34e5353ab2405e9efe2d9c146f52164f1068ddb72a6e02cdb2e17c1cfe.SingleValueExtendedPropertiesRequestBuilder) {
    return i24564f34e5353ab2405e9efe2d9c146f52164f1068ddb72a6e02cdb2e17c1cfe.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iee606450eb9c7d7f0ce88e813fdb68878885faa2e99762f240a53bed95ed4c18.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return iee606450eb9c7d7f0ce88e813fdb68878885faa2e99762f240a53bed95ed4c18.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ic07285bc27b4cdd05809b8320be3252542ec79f9a72a507850a252f3aba78a36.SnoozeReminderRequestBuilder) {
    return ic07285bc27b4cdd05809b8320be3252542ec79f9a72a507850a252f3aba78a36.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i8995f8a7bf29894e3868e11efb83cf43c08586f5789d309a1a1ef2ab7a10498a.TentativelyAcceptRequestBuilder) {
    return i8995f8a7bf29894e3868e11efb83cf43c08586f5789d309a1a1ef2ab7a10498a.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
