package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01e0892348e731467322610217cce58fefedab3334f67badc6002285bca059f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i21c239d72cfde932c0bdf50f372cf21fe79cd075723d1f0c53b9226bc3c6708a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/accept"
    i247070e7a0c24fb82cdc7f7465f8ab756c357a895cc2705d14c58ba13776105c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i25bfd266d768e15f969ac6f58214d77d01f8ac7c8e597dde68ebfa718a60050d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i38a242a961bb900d3bfdc162d29be37d1965287e3b49641aa9177865b978f4ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/decline"
    i41709df556b2760f1eae0f95308285452c1cba2e044dbc3795553b1d4b6c4dbd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/forward"
    i5424996a49e39090ca8eeaa6d88c6e066aafc31305e75783e59de63f2eb596f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i75fe2537ed6ea5d884ca1b6b0443dcb7c6af24f46fa49f6344cebc0b9b2b6fb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/extensions"
    i880c311d3aaaa8ebd42c68e40a617dc91817e05984516b6eea2dacff1569691d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i8e71c4477779b43130c561e5f3d4851d5de99133e4b177b792b29c2d833536bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/cancel"
    ic84f4e61ce2ce572648eab429339198e63051a66555fcbc240004fad0a9c867b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/calendar"
    if639f55d515e96ba9f6349824c8f40f745f7f4bb0fda0bfc0e96d8b5eb80ab0a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/attachments"
    i38047f0d60afe4ad849f4e3462c17f34bb200e49ccca08c16bbcfb3372d9e23b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i7251a127dbef106aac7ff3d5912444fe4fbfdd65863097348b011c2d0643d36d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    iaa41235be8ac05b3bf267e0ef73118e67449a14f52211279782f1009c598df2c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    ic6f41d4cd05390235ab7de5b50698ea31168b580e178315b4cfb8eb0a56ffe62 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item/exceptionoccurrences/item/extensions/item"
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
// EventItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
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
// EventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i21c239d72cfde932c0bdf50f372cf21fe79cd075723d1f0c53b9226bc3c6708a.AcceptRequestBuilder) {
    return i21c239d72cfde932c0bdf50f372cf21fe79cd075723d1f0c53b9226bc3c6708a.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*if639f55d515e96ba9f6349824c8f40f745f7f4bb0fda0bfc0e96d8b5eb80ab0a.AttachmentsRequestBuilder) {
    return if639f55d515e96ba9f6349824c8f40f745f7f4bb0fda0bfc0e96d8b5eb80ab0a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*iaa41235be8ac05b3bf267e0ef73118e67449a14f52211279782f1009c598df2c.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return iaa41235be8ac05b3bf267e0ef73118e67449a14f52211279782f1009c598df2c.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ic84f4e61ce2ce572648eab429339198e63051a66555fcbc240004fad0a9c867b.CalendarRequestBuilder) {
    return ic84f4e61ce2ce572648eab429339198e63051a66555fcbc240004fad0a9c867b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i8e71c4477779b43130c561e5f3d4851d5de99133e4b177b792b29c2d833536bd.CancelRequestBuilder) {
    return i8e71c4477779b43130c561e5f3d4851d5de99133e4b177b792b29c2d833536bd.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property exceptionOccurrences for users
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) Decline()(*i38a242a961bb900d3bfdc162d29be37d1965287e3b49641aa9177865b978f4ef.DeclineRequestBuilder) {
    return i38a242a961bb900d3bfdc162d29be37d1965287e3b49641aa9177865b978f4ef.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property exceptionOccurrences for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i5424996a49e39090ca8eeaa6d88c6e066aafc31305e75783e59de63f2eb596f7.DismissReminderRequestBuilder) {
    return i5424996a49e39090ca8eeaa6d88c6e066aafc31305e75783e59de63f2eb596f7.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i75fe2537ed6ea5d884ca1b6b0443dcb7c6af24f46fa49f6344cebc0b9b2b6fb4.ExtensionsRequestBuilder) {
    return i75fe2537ed6ea5d884ca1b6b0443dcb7c6af24f46fa49f6344cebc0b9b2b6fb4.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic6f41d4cd05390235ab7de5b50698ea31168b580e178315b4cfb8eb0a56ffe62.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ic6f41d4cd05390235ab7de5b50698ea31168b580e178315b4cfb8eb0a56ffe62.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i41709df556b2760f1eae0f95308285452c1cba2e044dbc3795553b1d4b6c4dbd.ForwardRequestBuilder) {
    return i41709df556b2760f1eae0f95308285452c1cba2e044dbc3795553b1d4b6c4dbd.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from users
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get exceptionOccurrences from users
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i247070e7a0c24fb82cdc7f7465f8ab756c357a895cc2705d14c58ba13776105c.MultiValueExtendedPropertiesRequestBuilder) {
    return i247070e7a0c24fb82cdc7f7465f8ab756c357a895cc2705d14c58ba13776105c.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7251a127dbef106aac7ff3d5912444fe4fbfdd65863097348b011c2d0643d36d.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i7251a127dbef106aac7ff3d5912444fe4fbfdd65863097348b011c2d0643d36d.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property exceptionOccurrences in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i880c311d3aaaa8ebd42c68e40a617dc91817e05984516b6eea2dacff1569691d.SingleValueExtendedPropertiesRequestBuilder) {
    return i880c311d3aaaa8ebd42c68e40a617dc91817e05984516b6eea2dacff1569691d.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i38047f0d60afe4ad849f4e3462c17f34bb200e49ccca08c16bbcfb3372d9e23b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i38047f0d60afe4ad849f4e3462c17f34bb200e49ccca08c16bbcfb3372d9e23b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i25bfd266d768e15f969ac6f58214d77d01f8ac7c8e597dde68ebfa718a60050d.SnoozeReminderRequestBuilder) {
    return i25bfd266d768e15f969ac6f58214d77d01f8ac7c8e597dde68ebfa718a60050d.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i01e0892348e731467322610217cce58fefedab3334f67badc6002285bca059f1.TentativelyAcceptRequestBuilder) {
    return i01e0892348e731467322610217cce58fefedab3334f67badc6002285bca059f1.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
