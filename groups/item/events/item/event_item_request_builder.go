package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i05a42ea7d5e1884a8e1dae2793ab110b1e1f8d988db71891436832059ff2ba3c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/cancel"
    i5fb746eca0c242c8984c8a879184ad9d08b364aa7b9c285a3bc182d2ddfc503e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/attachments"
    i79a7405307cbbdb21bedf1701ef943fd18c4517f47986e866d5dd325fbd2c118 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/decline"
    i8aae6ec506cc62628c5af6e10369ee326d5030e6a9f6dd3af2f33eed318509cb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/accept"
    iae63acc96b63dde476df5d572aa9d992f16b648b4365e066f9f52d4d87bb923c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/tentativelyaccept"
    ib4e29bfd4281e36f35dc582f9c214a6f560296bfeb558b3a54ea7b8bf8ae5b8e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/extensions"
    ib9a7b515d2c33bd771c3637dd035f38dd054dfc1c3bedb54b7c1c3fc35bd0c4f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/snoozereminder"
    ibc5ca5cafb1fc0ba72521ab6774586b362f987eb4515dfa602dbad582436708c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/multivalueextendedproperties"
    ic0681472e18a3022a4b2b729ab4483580fb7b4e6b7ae1d4f181ec9dbd1da88d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences"
    id83297e5327484cd1ced89db57bf0a85fa33d5cae6cdbb44636d29ebe5ae6f36 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances"
    iea920758cfc75c1ccf77797bcdbc550af58e71c9f109ec8a94722a33d74c7cc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/singlevalueextendedproperties"
    ieea5b0e4f6283da8e5578128baf2e8078539a018e85731134404d8b5aa938c76 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/calendar"
    if96aa91366f276971bcdb05b3ac53fb867427708b271afba983647fc1f3609cd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/forward"
    if9e65d705bb4056db822cab18c1d189374708a5ba4fbdacd32700342862d9952 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/dismissreminder"
    i1e6b45e04eb5c15511813d0040d37a7573d659ce2474195041adf65e7a3da179 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/exceptionoccurrences/item"
    i51135f6535d37302d8527c79893cfa3567ae1a2cc4fec29b051f79c3f3b706b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/extensions/item"
    i792a485ff185e2fdffa2e1feaabbe0871f5a892a62e0d8d85b4f2a593554431a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/instances/item"
    i81aa3005b6422532d1fe02077a5f622ae66cc374dced98aa4b6578b09fbfba01 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/singlevalueextendedproperties/item"
    iaf6ee02f9752677d650b6f2d4444fe3664c88db77ed2d14f02597638d09e85d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/multivalueextendedproperties/item"
    ife7ef1d1a994cdb0804ec1d15de22eaf8f8e9bc910b45c55f19b4d13e63ecff9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/events/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.group entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EventItemRequestBuilderGetQueryParameters the group's events.
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
// EventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i8aae6ec506cc62628c5af6e10369ee326d5030e6a9f6dd3af2f33eed318509cb.AcceptRequestBuilder) {
    return i8aae6ec506cc62628c5af6e10369ee326d5030e6a9f6dd3af2f33eed318509cb.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i5fb746eca0c242c8984c8a879184ad9d08b364aa7b9c285a3bc182d2ddfc503e.AttachmentsRequestBuilder) {
    return i5fb746eca0c242c8984c8a879184ad9d08b364aa7b9c285a3bc182d2ddfc503e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ife7ef1d1a994cdb0804ec1d15de22eaf8f8e9bc910b45c55f19b4d13e63ecff9.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ife7ef1d1a994cdb0804ec1d15de22eaf8f8e9bc910b45c55f19b4d13e63ecff9.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ieea5b0e4f6283da8e5578128baf2e8078539a018e85731134404d8b5aa938c76.CalendarRequestBuilder) {
    return ieea5b0e4f6283da8e5578128baf2e8078539a018e85731134404d8b5aa938c76.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i05a42ea7d5e1884a8e1dae2793ab110b1e1f8d988db71891436832059ff2ba3c.CancelRequestBuilder) {
    return i05a42ea7d5e1884a8e1dae2793ab110b1e1f8d988db71891436832059ff2ba3c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/events/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property events for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property events for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the group's events.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the group's events.
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
// CreatePatchRequestInformation update the navigation property events in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property events in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i79a7405307cbbdb21bedf1701ef943fd18c4517f47986e866d5dd325fbd2c118.DeclineRequestBuilder) {
    return i79a7405307cbbdb21bedf1701ef943fd18c4517f47986e866d5dd325fbd2c118.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for groups
func (m *EventItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*if9e65d705bb4056db822cab18c1d189374708a5ba4fbdacd32700342862d9952.DismissReminderRequestBuilder) {
    return if9e65d705bb4056db822cab18c1d189374708a5ba4fbdacd32700342862d9952.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ic0681472e18a3022a4b2b729ab4483580fb7b4e6b7ae1d4f181ec9dbd1da88d9.ExceptionOccurrencesRequestBuilder) {
    return ic0681472e18a3022a4b2b729ab4483580fb7b4e6b7ae1d4f181ec9dbd1da88d9.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i1e6b45e04eb5c15511813d0040d37a7573d659ce2474195041adf65e7a3da179.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i1e6b45e04eb5c15511813d0040d37a7573d659ce2474195041adf65e7a3da179.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ib4e29bfd4281e36f35dc582f9c214a6f560296bfeb558b3a54ea7b8bf8ae5b8e.ExtensionsRequestBuilder) {
    return ib4e29bfd4281e36f35dc582f9c214a6f560296bfeb558b3a54ea7b8bf8ae5b8e.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i51135f6535d37302d8527c79893cfa3567ae1a2cc4fec29b051f79c3f3b706b1.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i51135f6535d37302d8527c79893cfa3567ae1a2cc4fec29b051f79c3f3b706b1.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*if96aa91366f276971bcdb05b3ac53fb867427708b271afba983647fc1f3609cd.ForwardRequestBuilder) {
    return if96aa91366f276971bcdb05b3ac53fb867427708b271afba983647fc1f3609cd.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the group's events.
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
func (m *EventItemRequestBuilder) Instances()(*id83297e5327484cd1ced89db57bf0a85fa33d5cae6cdbb44636d29ebe5ae6f36.InstancesRequestBuilder) {
    return id83297e5327484cd1ced89db57bf0a85fa33d5cae6cdbb44636d29ebe5ae6f36.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i792a485ff185e2fdffa2e1feaabbe0871f5a892a62e0d8d85b4f2a593554431a.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i792a485ff185e2fdffa2e1feaabbe0871f5a892a62e0d8d85b4f2a593554431a.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ibc5ca5cafb1fc0ba72521ab6774586b362f987eb4515dfa602dbad582436708c.MultiValueExtendedPropertiesRequestBuilder) {
    return ibc5ca5cafb1fc0ba72521ab6774586b362f987eb4515dfa602dbad582436708c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iaf6ee02f9752677d650b6f2d4444fe3664c88db77ed2d14f02597638d09e85d1.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iaf6ee02f9752677d650b6f2d4444fe3664c88db77ed2d14f02597638d09e85d1.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in groups
func (m *EventItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*iea920758cfc75c1ccf77797bcdbc550af58e71c9f109ec8a94722a33d74c7cc4.SingleValueExtendedPropertiesRequestBuilder) {
    return iea920758cfc75c1ccf77797bcdbc550af58e71c9f109ec8a94722a33d74c7cc4.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i81aa3005b6422532d1fe02077a5f622ae66cc374dced98aa4b6578b09fbfba01.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i81aa3005b6422532d1fe02077a5f622ae66cc374dced98aa4b6578b09fbfba01.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ib9a7b515d2c33bd771c3637dd035f38dd054dfc1c3bedb54b7c1c3fc35bd0c4f.SnoozeReminderRequestBuilder) {
    return ib9a7b515d2c33bd771c3637dd035f38dd054dfc1c3bedb54b7c1c3fc35bd0c4f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*iae63acc96b63dde476df5d572aa9d992f16b648b4365e066f9f52d4d87bb923c.TentativelyAcceptRequestBuilder) {
    return iae63acc96b63dde476df5d572aa9d992f16b648b4365e066f9f52d4d87bb923c.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
