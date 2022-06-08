package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0ea8ab0bba5beeede093d83c541be8c66b021320d76401721bd00078a24d45a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/accept"
    i12e0ca19419c92e9242e720ab0014668680aeba2f474a1bf5b84bb72f292e234 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences"
    i399c64eaad805c394195598a977cd68055dd419f057f4679f9282cf17f6ffe47 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/extensions"
    i4f2c2f8126d1858148714b5952c8002615279cfcd4e9d43d02b36d2917038b28 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/cancel"
    i6a5348acf0f084c480b48d8c688aa7eeee8c8af838b0ca5e2f6fb0175faaeb40 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/attachments"
    i754015ce89c9de225717502e2c01d255f74913ca6352d439bfde68fdb1ed8cdf "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/singlevalueextendedproperties"
    i879420c24330f8775fb067bf872459af28093232087684f5a14b5eca048b5ab8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/snoozereminder"
    i8a0b0a3edc91f45bb57496f3170ef173e9f3b61ea75c917150638301f07580db "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/calendar"
    ia85f3d05b719427481232cf3b099b3da173e390bd6c6bf64e345c56caa22348d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/tentativelyaccept"
    ib9745936652db0ebb6668c855f67dd2778ddebdbfde6c0cd78e5fc2d20d608ea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/dismissreminder"
    ic728fa7a737d3ae0b3093286987fa5ed7c8e9d62acad0bcfdc7495a836b6c654 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/multivalueextendedproperties"
    id37f3d625d878bd89d20612fe6a6dfe981b9192b2f6f81d3f64e81c6aef8c74c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/decline"
    ieef52383219d6e7f90de8470496c4cf81609329acf4bbe285e8890e1061ddb51 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/forward"
    i2d7b19aedb40db595dc15e5ba1ba4d41cdd61ed2641580bcb114920e3fc15792 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/extensions/item"
    i4441e05a1a0bd5212eac64f79fd60b92c16b7e0462eeec6045e7eaae1608341b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item"
    i573a154724f0c54849ccec0bf6fcd278c7a61fe486cd05323f1cd0e3147bab1f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/attachments/item"
    i6e143d963336d0eb1fca0303fc558f38a07026efa48e32549400b93db27e94dc "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/singlevalueextendedproperties/item"
    i7d051bb9d8ad6c7bf75445d177b1d6a880eed14999301591a29f9e8cb9b81276 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/multivalueextendedproperties/item"
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
// EventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
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
// EventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i0ea8ab0bba5beeede093d83c541be8c66b021320d76401721bd00078a24d45a6.AcceptRequestBuilder) {
    return i0ea8ab0bba5beeede093d83c541be8c66b021320d76401721bd00078a24d45a6.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i6a5348acf0f084c480b48d8c688aa7eeee8c8af838b0ca5e2f6fb0175faaeb40.AttachmentsRequestBuilder) {
    return i6a5348acf0f084c480b48d8c688aa7eeee8c8af838b0ca5e2f6fb0175faaeb40.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i573a154724f0c54849ccec0bf6fcd278c7a61fe486cd05323f1cd0e3147bab1f.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i573a154724f0c54849ccec0bf6fcd278c7a61fe486cd05323f1cd0e3147bab1f.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i8a0b0a3edc91f45bb57496f3170ef173e9f3b61ea75c917150638301f07580db.CalendarRequestBuilder) {
    return i8a0b0a3edc91f45bb57496f3170ef173e9f3b61ea75c917150638301f07580db.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i4f2c2f8126d1858148714b5952c8002615279cfcd4e9d43d02b36d2917038b28.CancelRequestBuilder) {
    return i4f2c2f8126d1858148714b5952c8002615279cfcd4e9d43d02b36d2917038b28.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property instances for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instances for groups
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
// CreatePatchRequestInformation update the navigation property instances in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instances in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*id37f3d625d878bd89d20612fe6a6dfe981b9192b2f6f81d3f64e81c6aef8c74c.DeclineRequestBuilder) {
    return id37f3d625d878bd89d20612fe6a6dfe981b9192b2f6f81d3f64e81c6aef8c74c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for groups
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property instances for groups
func (m *EventItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *EventItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*ib9745936652db0ebb6668c855f67dd2778ddebdbfde6c0cd78e5fc2d20d608ea.DismissReminderRequestBuilder) {
    return ib9745936652db0ebb6668c855f67dd2778ddebdbfde6c0cd78e5fc2d20d608ea.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i12e0ca19419c92e9242e720ab0014668680aeba2f474a1bf5b84bb72f292e234.ExceptionOccurrencesRequestBuilder) {
    return i12e0ca19419c92e9242e720ab0014668680aeba2f474a1bf5b84bb72f292e234.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i4441e05a1a0bd5212eac64f79fd60b92c16b7e0462eeec6045e7eaae1608341b.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i4441e05a1a0bd5212eac64f79fd60b92c16b7e0462eeec6045e7eaae1608341b.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i399c64eaad805c394195598a977cd68055dd419f057f4679f9282cf17f6ffe47.ExtensionsRequestBuilder) {
    return i399c64eaad805c394195598a977cd68055dd419f057f4679f9282cf17f6ffe47.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i2d7b19aedb40db595dc15e5ba1ba4d41cdd61ed2641580bcb114920e3fc15792.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i2d7b19aedb40db595dc15e5ba1ba4d41cdd61ed2641580bcb114920e3fc15792.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ieef52383219d6e7f90de8470496c4cf81609329acf4bbe285e8890e1061ddb51.ForwardRequestBuilder) {
    return ieef52383219d6e7f90de8470496c4cf81609329acf4bbe285e8890e1061ddb51.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *EventItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic728fa7a737d3ae0b3093286987fa5ed7c8e9d62acad0bcfdc7495a836b6c654.MultiValueExtendedPropertiesRequestBuilder) {
    return ic728fa7a737d3ae0b3093286987fa5ed7c8e9d62acad0bcfdc7495a836b6c654.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7d051bb9d8ad6c7bf75445d177b1d6a880eed14999301591a29f9e8cb9b81276.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i7d051bb9d8ad6c7bf75445d177b1d6a880eed14999301591a29f9e8cb9b81276.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in groups
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property instances in groups
func (m *EventItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, requestConfiguration *EventItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i754015ce89c9de225717502e2c01d255f74913ca6352d439bfde68fdb1ed8cdf.SingleValueExtendedPropertiesRequestBuilder) {
    return i754015ce89c9de225717502e2c01d255f74913ca6352d439bfde68fdb1ed8cdf.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i6e143d963336d0eb1fca0303fc558f38a07026efa48e32549400b93db27e94dc.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i6e143d963336d0eb1fca0303fc558f38a07026efa48e32549400b93db27e94dc.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i879420c24330f8775fb067bf872459af28093232087684f5a14b5eca048b5ab8.SnoozeReminderRequestBuilder) {
    return i879420c24330f8775fb067bf872459af28093232087684f5a14b5eca048b5ab8.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ia85f3d05b719427481232cf3b099b3da173e390bd6c6bf64e345c56caa22348d.TentativelyAcceptRequestBuilder) {
    return ia85f3d05b719427481232cf3b099b3da173e390bd6c6bf64e345c56caa22348d.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
