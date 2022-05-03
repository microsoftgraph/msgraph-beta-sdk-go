package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i26748f3b82efe6dd34888eee41065858c3a3e7e659c22f2c94447a37d3662a8b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/attachments"
    i41cafad21014522f56f3d33f6b2da69078e47c2055eb1dc7d5aabff3dd6622cd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/forward"
    i5eb61ecb78730bc2c3ad575ffe87c5367b968f709cad0817e2fd5eeed10953e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i5ec90d122c8eea0bc4755f573ef8e7b85396a0799234015650c38ea8baa53a7c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/decline"
    i62024e0e912fe04a643b48f00c42e38a84248141ba69d1680e72b03c9bb80495 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/extensions"
    i8d02fff45bccaec2ae6e8864741639324a55edac01027508e32a664617febd27 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    i911030a672e539aae3de6ddfe1bf95540afc5ae7ed7d4240f2a4de0d2c87c1d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i997f2833a1fe607f8aaa98d463e65ff4d16d5e51a71e17621d781a8fae2f3d7c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i9e8f8173bb92f3d53cb90550b6401379e73d006c9d7150e8671ac72828b22cfd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/accept"
    ic0435f44eb0f660f9d9cb6a2159d343b4096f9c8f3ee1ba459dd96e16cfdeb59 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/calendar"
    ic966907eb7d0213581b24d3da583a939d9b333100d6d5b43dc85b8315530edb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/dismissreminder"
    icdcf759ed5568b0365d41c65ec2cb504fe017b3316b0a185294d0cdcfb5a1cb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances"
    ifcaa57dd328055e0183c1bab880aa0e1d7e613804099801e5c201031ae9c1e29 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/cancel"
    i2698471ac0526318c600feec3cca3fa6b0464c485340df96ce4946c9ac8b05ad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item"
    i6056de85764c487309ec6fcf68bec715b17a8200ecdc7222e03c80d3b42151e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i8a670c1f6b71fca49b869fcff1553f423f74f57708266eef9569e254765113da "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/extensions/item"
    ia8afd0840ce8785a77fde9341b7ad7d166e146d457f70fd09d5a2059e99cf241 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    if563f512c2d4b5231e1ac2fd148948f2da7181a6406904b46d3cb8db58eb7fae "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendar/calendarview/item/exceptionoccurrences/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*i9e8f8173bb92f3d53cb90550b6401379e73d006c9d7150e8671ac72828b22cfd.AcceptRequestBuilder) {
    return i9e8f8173bb92f3d53cb90550b6401379e73d006c9d7150e8671ac72828b22cfd.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i26748f3b82efe6dd34888eee41065858c3a3e7e659c22f2c94447a37d3662a8b.AttachmentsRequestBuilder) {
    return i26748f3b82efe6dd34888eee41065858c3a3e7e659c22f2c94447a37d3662a8b.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*if563f512c2d4b5231e1ac2fd148948f2da7181a6406904b46d3cb8db58eb7fae.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return if563f512c2d4b5231e1ac2fd148948f2da7181a6406904b46d3cb8db58eb7fae.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ic0435f44eb0f660f9d9cb6a2159d343b4096f9c8f3ee1ba459dd96e16cfdeb59.CalendarRequestBuilder) {
    return ic0435f44eb0f660f9d9cb6a2159d343b4096f9c8f3ee1ba459dd96e16cfdeb59.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ifcaa57dd328055e0183c1bab880aa0e1d7e613804099801e5c201031ae9c1e29.CancelRequestBuilder) {
    return ifcaa57dd328055e0183c1bab880aa0e1d7e613804099801e5c201031ae9c1e29.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i5ec90d122c8eea0bc4755f573ef8e7b85396a0799234015650c38ea8baa53a7c.DeclineRequestBuilder) {
    return i5ec90d122c8eea0bc4755f573ef8e7b85396a0799234015650c38ea8baa53a7c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*ic966907eb7d0213581b24d3da583a939d9b333100d6d5b43dc85b8315530edb9.DismissReminderRequestBuilder) {
    return ic966907eb7d0213581b24d3da583a939d9b333100d6d5b43dc85b8315530edb9.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i62024e0e912fe04a643b48f00c42e38a84248141ba69d1680e72b03c9bb80495.ExtensionsRequestBuilder) {
    return i62024e0e912fe04a643b48f00c42e38a84248141ba69d1680e72b03c9bb80495.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i8a670c1f6b71fca49b869fcff1553f423f74f57708266eef9569e254765113da.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8a670c1f6b71fca49b869fcff1553f423f74f57708266eef9569e254765113da.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i41cafad21014522f56f3d33f6b2da69078e47c2055eb1dc7d5aabff3dd6622cd.ForwardRequestBuilder) {
    return i41cafad21014522f56f3d33f6b2da69078e47c2055eb1dc7d5aabff3dd6622cd.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*icdcf759ed5568b0365d41c65ec2cb504fe017b3316b0a185294d0cdcfb5a1cb6.InstancesRequestBuilder) {
    return icdcf759ed5568b0365d41c65ec2cb504fe017b3316b0a185294d0cdcfb5a1cb6.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i2698471ac0526318c600feec3cca3fa6b0464c485340df96ce4946c9ac8b05ad.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i2698471ac0526318c600feec3cca3fa6b0464c485340df96ce4946c9ac8b05ad.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i8d02fff45bccaec2ae6e8864741639324a55edac01027508e32a664617febd27.MultiValueExtendedPropertiesRequestBuilder) {
    return i8d02fff45bccaec2ae6e8864741639324a55edac01027508e32a664617febd27.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i6056de85764c487309ec6fcf68bec715b17a8200ecdc7222e03c80d3b42151e8.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i6056de85764c487309ec6fcf68bec715b17a8200ecdc7222e03c80d3b42151e8.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i911030a672e539aae3de6ddfe1bf95540afc5ae7ed7d4240f2a4de0d2c87c1d5.SingleValueExtendedPropertiesRequestBuilder) {
    return i911030a672e539aae3de6ddfe1bf95540afc5ae7ed7d4240f2a4de0d2c87c1d5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendar.calendarView.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia8afd0840ce8785a77fde9341b7ad7d166e146d457f70fd09d5a2059e99cf241.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ia8afd0840ce8785a77fde9341b7ad7d166e146d457f70fd09d5a2059e99cf241.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i997f2833a1fe607f8aaa98d463e65ff4d16d5e51a71e17621d781a8fae2f3d7c.SnoozeReminderRequestBuilder) {
    return i997f2833a1fe607f8aaa98d463e65ff4d16d5e51a71e17621d781a8fae2f3d7c.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i5eb61ecb78730bc2c3ad575ffe87c5367b968f709cad0817e2fd5eeed10953e0.TentativelyAcceptRequestBuilder) {
    return i5eb61ecb78730bc2c3ad575ffe87c5367b968f709cad0817e2fd5eeed10953e0.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
