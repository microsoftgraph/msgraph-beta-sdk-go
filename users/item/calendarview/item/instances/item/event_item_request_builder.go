package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04dd445e58ac2dc8444371ea8bf66d600ada99f00ecdfdc1bd49cb55d8721d2a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/cancel"
    i1b9d0ba5972e24686ad4e34c2e350df1009e5b96944e4f9e46001c5bb7543f54 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/singlevalueextendedproperties"
    i3945772f0f66626567d385c013343a1c18775fd6d2bfed256f0b93cbb4727b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/extensions"
    i4b62240b9f6c61d7c563fd1e0be066e64f088c5f3d0f877c1fa2301fe3a5456b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/tentativelyaccept"
    i5f563fb78caead28abfd54c986e4059f7909c1d0e7515a7b52c6d746cce80998 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/decline"
    i69754caef6507f32520c138342e448dba6615f13385b5031dc8e90a18f04152f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/snoozereminder"
    i6cf3a7f81197f53a9b0a92a91567bf9e7daec09ab930db2ee7aa56be6d70c5ae "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/forward"
    i93e284c96de720a935e8a900e39f6dde33ffd9f3a356da8783c7e8fefdb9faa2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences"
    i9903f93be66c9b70cedcb3981a1413c4079ffe1df3d8be841aa33fa916737b47 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/calendar"
    ia2dec4870d660e87e2cbf09cdbd87ec8daae5ec8591bf55e3c34b0fed77cad1a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/dismissreminder"
    ib7a78248b9a6c8acf89bb8a04a7069e720a0141e34937c28602a77a69d33d39a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/attachments"
    ibe9c853c6f498babeef0e6f88b41351e8e3198c2ad04c13846100e89b0c76867 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/accept"
    if05012549238b8d05860cfffcd1d1e08c7d74e2698d24a2c9b6a7dc5a479a970 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/multivalueextendedproperties"
    i1d1ec8c66027c70de040976e7c6774fe6af352d4ac490e229116083db15fe62f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/exceptionoccurrences/item"
    i5414e31f123f2f2e1c1f6071ab0a90290de7509755ac67e2e992949c87a79467 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/extensions/item"
    ibd4b38c3cf0d2fc5fe064eba976de1d165acebf85532f4f749af3d39273d156d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/attachments/item"
    ic2c6cd996034b82771104046b2cf7e493458a10e3a4315e9e7753d81cccbdf86 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
    iebe71885860c0b08e085b212d5c92dcf7c05c9420de956a6e50a4808b8a46323 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item/multivalueextendedproperties/item"
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
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*ibe9c853c6f498babeef0e6f88b41351e8e3198c2ad04c13846100e89b0c76867.AcceptRequestBuilder) {
    return ibe9c853c6f498babeef0e6f88b41351e8e3198c2ad04c13846100e89b0c76867.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ib7a78248b9a6c8acf89bb8a04a7069e720a0141e34937c28602a77a69d33d39a.AttachmentsRequestBuilder) {
    return ib7a78248b9a6c8acf89bb8a04a7069e720a0141e34937c28602a77a69d33d39a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ibd4b38c3cf0d2fc5fe064eba976de1d165acebf85532f4f749af3d39273d156d.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ibd4b38c3cf0d2fc5fe064eba976de1d165acebf85532f4f749af3d39273d156d.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i9903f93be66c9b70cedcb3981a1413c4079ffe1df3d8be841aa33fa916737b47.CalendarRequestBuilder) {
    return i9903f93be66c9b70cedcb3981a1413c4079ffe1df3d8be841aa33fa916737b47.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i04dd445e58ac2dc8444371ea8bf66d600ada99f00ecdfdc1bd49cb55d8721d2a.CancelRequestBuilder) {
    return i04dd445e58ac2dc8444371ea8bf66d600ada99f00ecdfdc1bd49cb55d8721d2a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property instances for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property instances in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i5f563fb78caead28abfd54c986e4059f7909c1d0e7515a7b52c6d746cce80998.DeclineRequestBuilder) {
    return i5f563fb78caead28abfd54c986e4059f7909c1d0e7515a7b52c6d746cce80998.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*ia2dec4870d660e87e2cbf09cdbd87ec8daae5ec8591bf55e3c34b0fed77cad1a.DismissReminderRequestBuilder) {
    return ia2dec4870d660e87e2cbf09cdbd87ec8daae5ec8591bf55e3c34b0fed77cad1a.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i93e284c96de720a935e8a900e39f6dde33ffd9f3a356da8783c7e8fefdb9faa2.ExceptionOccurrencesRequestBuilder) {
    return i93e284c96de720a935e8a900e39f6dde33ffd9f3a356da8783c7e8fefdb9faa2.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i1d1ec8c66027c70de040976e7c6774fe6af352d4ac490e229116083db15fe62f.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i1d1ec8c66027c70de040976e7c6774fe6af352d4ac490e229116083db15fe62f.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i3945772f0f66626567d385c013343a1c18775fd6d2bfed256f0b93cbb4727b3c.ExtensionsRequestBuilder) {
    return i3945772f0f66626567d385c013343a1c18775fd6d2bfed256f0b93cbb4727b3c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i5414e31f123f2f2e1c1f6071ab0a90290de7509755ac67e2e992949c87a79467.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i5414e31f123f2f2e1c1f6071ab0a90290de7509755ac67e2e992949c87a79467.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i6cf3a7f81197f53a9b0a92a91567bf9e7daec09ab930db2ee7aa56be6d70c5ae.ForwardRequestBuilder) {
    return i6cf3a7f81197f53a9b0a92a91567bf9e7daec09ab930db2ee7aa56be6d70c5ae.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*if05012549238b8d05860cfffcd1d1e08c7d74e2698d24a2c9b6a7dc5a479a970.MultiValueExtendedPropertiesRequestBuilder) {
    return if05012549238b8d05860cfffcd1d1e08c7d74e2698d24a2c9b6a7dc5a479a970.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iebe71885860c0b08e085b212d5c92dcf7c05c9420de956a6e50a4808b8a46323.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iebe71885860c0b08e085b212d5c92dcf7c05c9420de956a6e50a4808b8a46323.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i1b9d0ba5972e24686ad4e34c2e350df1009e5b96944e4f9e46001c5bb7543f54.SingleValueExtendedPropertiesRequestBuilder) {
    return i1b9d0ba5972e24686ad4e34c2e350df1009e5b96944e4f9e46001c5bb7543f54.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ic2c6cd996034b82771104046b2cf7e493458a10e3a4315e9e7753d81cccbdf86.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ic2c6cd996034b82771104046b2cf7e493458a10e3a4315e9e7753d81cccbdf86.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i69754caef6507f32520c138342e448dba6615f13385b5031dc8e90a18f04152f.SnoozeReminderRequestBuilder) {
    return i69754caef6507f32520c138342e448dba6615f13385b5031dc8e90a18f04152f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i4b62240b9f6c61d7c563fd1e0be066e64f088c5f3d0f877c1fa2301fe3a5456b.TentativelyAcceptRequestBuilder) {
    return i4b62240b9f6c61d7c563fd1e0be066e64f088c5f3d0f877c1fa2301fe3a5456b.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
