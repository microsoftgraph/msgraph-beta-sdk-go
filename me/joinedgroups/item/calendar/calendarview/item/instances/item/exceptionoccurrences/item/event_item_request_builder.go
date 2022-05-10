package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i16e7399e0f6158531b334a20ee7c439efed975043f485874a88cc3dca40f757d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i1d1af58ce3ff10571abc16853e55a16414b9e0b6faa3fdae3d311ef4aab70a68 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i2778af276ad6b8a52405e509c83cfd1ef2bbb994fee04ed4d4f0187d4154eea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i3752d9245a32a2c6d07b9dfda6e6a79f64664072ec46653d9e80bc47abe5b8ec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i7a622d6d090b7444b69442afeef11ef916b689b7bb9b7424fc51c154e7afb665 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i7f73651273c7aa747cfc8beb5f597a871d2a981ae6a247a9242e36de40e1b07a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    i80075baff1e429f5d466ddc6942885527f97b017bc691bf0c24fbdc5dde45461 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i8204d9e63aba41ec1b830110b2496bfdf9d0ed098c8818334be51430bbac4357 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    iac4c0624d0975b34543936e3725b55b43944c4b9a6d792453b7c4d1e644a29be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ib93216e64a0787adbcc76bdc9819234c01e4d148073066e0ce2eb1d922a667ef "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    ic346c2265de9797c5687bc567f5313b4595d8206645f452e8d495ebe36a38c5c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    if272446c783a8bbabdb6c998fcbcb4d39f31e07e27faa7851718c8a4743531a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i24e6842412d080c7c7d6f7241058ea31ead5c0a12eb6f2805461528b8f3a65ac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i7b63fd48689c5734b185c11a31e6a953ffc69ad49860e02bb53e99a002d60644 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    i8401c94866cab982b86aa46a0bb24daffb52b72ccc49fddee013ee69de533e23 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ie25a4ec5a2b1dba0410e27db0f189af46c1a2127a8aca3e58acc603d3fc220ed "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
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
// EventItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EventItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*ic346c2265de9797c5687bc567f5313b4595d8206645f452e8d495ebe36a38c5c.AcceptRequestBuilder) {
    return ic346c2265de9797c5687bc567f5313b4595d8206645f452e8d495ebe36a38c5c.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i16e7399e0f6158531b334a20ee7c439efed975043f485874a88cc3dca40f757d.AttachmentsRequestBuilder) {
    return i16e7399e0f6158531b334a20ee7c439efed975043f485874a88cc3dca40f757d.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i7b63fd48689c5734b185c11a31e6a953ffc69ad49860e02bb53e99a002d60644.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i7b63fd48689c5734b185c11a31e6a953ffc69ad49860e02bb53e99a002d60644.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i7f73651273c7aa747cfc8beb5f597a871d2a981ae6a247a9242e36de40e1b07a.CalendarRequestBuilder) {
    return i7f73651273c7aa747cfc8beb5f597a871d2a981ae6a247a9242e36de40e1b07a.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ib93216e64a0787adbcc76bdc9819234c01e4d148073066e0ce2eb1d922a667ef.CancelRequestBuilder) {
    return ib93216e64a0787adbcc76bdc9819234c01e4d148073066e0ce2eb1d922a667ef.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property exceptionOccurrences for me
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
// CreateGetRequestInformation get exceptionOccurrences from me
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get exceptionOccurrences from me
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) Decline()(*i8204d9e63aba41ec1b830110b2496bfdf9d0ed098c8818334be51430bbac4357.DeclineRequestBuilder) {
    return i8204d9e63aba41ec1b830110b2496bfdf9d0ed098c8818334be51430bbac4357.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property exceptionOccurrences for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i1d1af58ce3ff10571abc16853e55a16414b9e0b6faa3fdae3d311ef4aab70a68.DismissReminderRequestBuilder) {
    return i1d1af58ce3ff10571abc16853e55a16414b9e0b6faa3fdae3d311ef4aab70a68.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i7a622d6d090b7444b69442afeef11ef916b689b7bb9b7424fc51c154e7afb665.ExtensionsRequestBuilder) {
    return i7a622d6d090b7444b69442afeef11ef916b689b7bb9b7424fc51c154e7afb665.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ie25a4ec5a2b1dba0410e27db0f189af46c1a2127a8aca3e58acc603d3fc220ed.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ie25a4ec5a2b1dba0410e27db0f189af46c1a2127a8aca3e58acc603d3fc220ed.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i80075baff1e429f5d466ddc6942885527f97b017bc691bf0c24fbdc5dde45461.ForwardRequestBuilder) {
    return i80075baff1e429f5d466ddc6942885527f97b017bc691bf0c24fbdc5dde45461.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*if272446c783a8bbabdb6c998fcbcb4d39f31e07e27faa7851718c8a4743531a9.MultiValueExtendedPropertiesRequestBuilder) {
    return if272446c783a8bbabdb6c998fcbcb4d39f31e07e27faa7851718c8a4743531a9.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i8401c94866cab982b86aa46a0bb24daffb52b72ccc49fddee013ee69de533e23.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i8401c94866cab982b86aa46a0bb24daffb52b72ccc49fddee013ee69de533e23.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i2778af276ad6b8a52405e509c83cfd1ef2bbb994fee04ed4d4f0187d4154eea1.SingleValueExtendedPropertiesRequestBuilder) {
    return i2778af276ad6b8a52405e509c83cfd1ef2bbb994fee04ed4d4f0187d4154eea1.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i24e6842412d080c7c7d6f7241058ea31ead5c0a12eb6f2805461528b8f3a65ac.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i24e6842412d080c7c7d6f7241058ea31ead5c0a12eb6f2805461528b8f3a65ac.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i3752d9245a32a2c6d07b9dfda6e6a79f64664072ec46653d9e80bc47abe5b8ec.SnoozeReminderRequestBuilder) {
    return i3752d9245a32a2c6d07b9dfda6e6a79f64664072ec46653d9e80bc47abe5b8ec.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*iac4c0624d0975b34543936e3725b55b43944c4b9a6d792453b7c4d1e644a29be.TentativelyAcceptRequestBuilder) {
    return iac4c0624d0975b34543936e3725b55b43944c4b9a6d792453b7c4d1e644a29be.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
