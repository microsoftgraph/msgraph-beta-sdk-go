package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i15c209653a211d9777ec2d1ed8b878b064e002f7ea48551614588ed56730cde9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    i1c5922e79c6b3359cea7fd4593e78e5a6e27617ec8e29d178b2d579f194ee26c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i2b44d42cea3020c91d54fc71e978be6107fee41ea0e2504f321fb69b89fc65b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i56f760d98dfc74d1875dedb098ecfcc9d7ea1d7fce874649b9e5697a384f6b1f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i62385dc821ac307d2ef5dca0c390c2821828649863fc960ae04a09b7cb027266 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    i7964f486890ccf355de39769c25eb35ad46f79afe7172e18782b7be7ffc15ac5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i7a49f43f057f14be019e7094d2cdf879307e2d1ddf9bbcf6c8510ad4b02fdf60 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    i7c5ec3d3bd0ee25ec803487e7f108d0c536bd96aa5eed2ce4be483fb0b8a922a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    ia349e2bc849f8e17f5b91ed44e682b64884592ad964ba7eb2f3c0a308f81c1c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    ie265499f2fa9a9a5fe299ea9a50b3be1c1aaf310bbdf3bc4c753e1ebddf0cc10 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    if47277ffd8c37dd6ebbc7c51fa22abedbb69451d9158718fda225aafeac4c179 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    if54352e1a5a47c0ac23ffcd83973d4626320ecfc628e8c5999363e6a67d702c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i66d929b2b6c4d9d882ed4daa7f664ea4c0d12f84f181934f516b883de9e29f7d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    i9fe50f684237f9011738361bdb77d2f8b3a0aa38567f5e63ea60c806bbe30a4b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    iaaa5806f96fe89f94a648a8d174c122a6b9e05309a7ccf9bcc195cc6c8ad5868 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ib210197b72aa0ad39cf94d53249e84a44c98896b6f08d020e9b4b7a9874b9041 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
func (m *EventItemRequestBuilder) Accept()(*i7a49f43f057f14be019e7094d2cdf879307e2d1ddf9bbcf6c8510ad4b02fdf60.AcceptRequestBuilder) {
    return i7a49f43f057f14be019e7094d2cdf879307e2d1ddf9bbcf6c8510ad4b02fdf60.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ia349e2bc849f8e17f5b91ed44e682b64884592ad964ba7eb2f3c0a308f81c1c3.AttachmentsRequestBuilder) {
    return ia349e2bc849f8e17f5b91ed44e682b64884592ad964ba7eb2f3c0a308f81c1c3.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i66d929b2b6c4d9d882ed4daa7f664ea4c0d12f84f181934f516b883de9e29f7d.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i66d929b2b6c4d9d882ed4daa7f664ea4c0d12f84f181934f516b883de9e29f7d.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i62385dc821ac307d2ef5dca0c390c2821828649863fc960ae04a09b7cb027266.CalendarRequestBuilder) {
    return i62385dc821ac307d2ef5dca0c390c2821828649863fc960ae04a09b7cb027266.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*if54352e1a5a47c0ac23ffcd83973d4626320ecfc628e8c5999363e6a67d702c1.CancelRequestBuilder) {
    return if54352e1a5a47c0ac23ffcd83973d4626320ecfc628e8c5999363e6a67d702c1.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for groups
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
// CreateGetRequestInformation get exceptionOccurrences from groups
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in groups
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
func (m *EventItemRequestBuilder) Decline()(*i15c209653a211d9777ec2d1ed8b878b064e002f7ea48551614588ed56730cde9.DeclineRequestBuilder) {
    return i15c209653a211d9777ec2d1ed8b878b064e002f7ea48551614588ed56730cde9.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*i2b44d42cea3020c91d54fc71e978be6107fee41ea0e2504f321fb69b89fc65b5.DismissReminderRequestBuilder) {
    return i2b44d42cea3020c91d54fc71e978be6107fee41ea0e2504f321fb69b89fc65b5.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*if47277ffd8c37dd6ebbc7c51fa22abedbb69451d9158718fda225aafeac4c179.ExtensionsRequestBuilder) {
    return if47277ffd8c37dd6ebbc7c51fa22abedbb69451d9158718fda225aafeac4c179.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i9fe50f684237f9011738361bdb77d2f8b3a0aa38567f5e63ea60c806bbe30a4b.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i9fe50f684237f9011738361bdb77d2f8b3a0aa38567f5e63ea60c806bbe30a4b.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i7c5ec3d3bd0ee25ec803487e7f108d0c536bd96aa5eed2ce4be483fb0b8a922a.ForwardRequestBuilder) {
    return i7c5ec3d3bd0ee25ec803487e7f108d0c536bd96aa5eed2ce4be483fb0b8a922a.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i7964f486890ccf355de39769c25eb35ad46f79afe7172e18782b7be7ffc15ac5.MultiValueExtendedPropertiesRequestBuilder) {
    return i7964f486890ccf355de39769c25eb35ad46f79afe7172e18782b7be7ffc15ac5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iaaa5806f96fe89f94a648a8d174c122a6b9e05309a7ccf9bcc195cc6c8ad5868.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iaaa5806f96fe89f94a648a8d174c122a6b9e05309a7ccf9bcc195cc6c8ad5868.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i1c5922e79c6b3359cea7fd4593e78e5a6e27617ec8e29d178b2d579f194ee26c.SingleValueExtendedPropertiesRequestBuilder) {
    return i1c5922e79c6b3359cea7fd4593e78e5a6e27617ec8e29d178b2d579f194ee26c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib210197b72aa0ad39cf94d53249e84a44c98896b6f08d020e9b4b7a9874b9041.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ib210197b72aa0ad39cf94d53249e84a44c98896b6f08d020e9b4b7a9874b9041.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ie265499f2fa9a9a5fe299ea9a50b3be1c1aaf310bbdf3bc4c753e1ebddf0cc10.SnoozeReminderRequestBuilder) {
    return ie265499f2fa9a9a5fe299ea9a50b3be1c1aaf310bbdf3bc4c753e1ebddf0cc10.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i56f760d98dfc74d1875dedb098ecfcc9d7ea1d7fce874649b9e5697a384f6b1f.TentativelyAcceptRequestBuilder) {
    return i56f760d98dfc74d1875dedb098ecfcc9d7ea1d7fce874649b9e5697a384f6b1f.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
