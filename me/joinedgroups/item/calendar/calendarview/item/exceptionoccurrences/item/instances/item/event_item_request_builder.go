package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3c89bfd65920400c43fc742d303e8c57fc47ea1c1a05a446bfbac376b668f845 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    i5a6e7d28a08001fc4a358505f17f645b466f8fee88131e6d0eaa4e087172e4c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i5aa3ecf098a4711b47308b51d022d12beacb52d164a2692ce51647b28378aecc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    i65cf3e5b2dc0d43f3353ea93662e21a729a41af0d0267b7b966b99d913978b18 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i7576ba38ad041650f20a21edfa06d72e6982e9b16b1837a42b64c72f2a7db2aa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i7f46b018ba2873ff617fef5684194592416d3de4706663395fdc96f19bea84cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    i94973edce7445a8864c18f665412fc79a47968004ab268fa0b19c1e001974ee3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i9b69120e569c5ad2fd941d5083b9fd15211fb764cff617370b7d3e80882592cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    ia52f958c7a96e4e22713d067ae188859cefa4ca6005aac38177063c344ce3372 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    ib8e8cb2ab4026e7192d0358e58d577e0c9c266f6d4e08eab7aad7b2ec435d940 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    ie12474e462c96c998789c9569a055d9fbea482560ea4ad83f327d27eee0162ce "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    ife2fc0637378d7ca14ba4b282dad1c304a87a13cf717242151ec5e13488fc203 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    i1c18134f6f07c2ab5e8da11010ba59f9b04bb2e337d6972f9a31737f2824486f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    i8396ed055df92d1ae75181a3a1a83121f4c5882b54e6a137b6d1441005fb3855 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
    i9708cf02dee7a32cd99aad28f579aa32d6034cbf765b2bbe134a42e5e6e0641c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    iaa3b0089e93d44c50c83ad3a25361adc1c8aedcb3f4d2b7bd69863e1bbefd761 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*ie12474e462c96c998789c9569a055d9fbea482560ea4ad83f327d27eee0162ce.AcceptRequestBuilder) {
    return ie12474e462c96c998789c9569a055d9fbea482560ea4ad83f327d27eee0162ce.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i7f46b018ba2873ff617fef5684194592416d3de4706663395fdc96f19bea84cd.AttachmentsRequestBuilder) {
    return i7f46b018ba2873ff617fef5684194592416d3de4706663395fdc96f19bea84cd.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i8396ed055df92d1ae75181a3a1a83121f4c5882b54e6a137b6d1441005fb3855.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i8396ed055df92d1ae75181a3a1a83121f4c5882b54e6a137b6d1441005fb3855.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ife2fc0637378d7ca14ba4b282dad1c304a87a13cf717242151ec5e13488fc203.CalendarRequestBuilder) {
    return ife2fc0637378d7ca14ba4b282dad1c304a87a13cf717242151ec5e13488fc203.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i3c89bfd65920400c43fc742d303e8c57fc47ea1c1a05a446bfbac376b668f845.CancelRequestBuilder) {
    return i3c89bfd65920400c43fc742d303e8c57fc47ea1c1a05a446bfbac376b668f845.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property instances for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instances for me
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
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property instances in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instances in me
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
func (m *EventItemRequestBuilder) Decline()(*i5aa3ecf098a4711b47308b51d022d12beacb52d164a2692ce51647b28378aecc.DeclineRequestBuilder) {
    return i5aa3ecf098a4711b47308b51d022d12beacb52d164a2692ce51647b28378aecc.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property instances for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i65cf3e5b2dc0d43f3353ea93662e21a729a41af0d0267b7b966b99d913978b18.DismissReminderRequestBuilder) {
    return i65cf3e5b2dc0d43f3353ea93662e21a729a41af0d0267b7b966b99d913978b18.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i9b69120e569c5ad2fd941d5083b9fd15211fb764cff617370b7d3e80882592cd.ExtensionsRequestBuilder) {
    return i9b69120e569c5ad2fd941d5083b9fd15211fb764cff617370b7d3e80882592cd.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i9708cf02dee7a32cd99aad28f579aa32d6034cbf765b2bbe134a42e5e6e0641c.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i9708cf02dee7a32cd99aad28f579aa32d6034cbf765b2bbe134a42e5e6e0641c.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ia52f958c7a96e4e22713d067ae188859cefa4ca6005aac38177063c344ce3372.ForwardRequestBuilder) {
    return ia52f958c7a96e4e22713d067ae188859cefa4ca6005aac38177063c344ce3372.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i5a6e7d28a08001fc4a358505f17f645b466f8fee88131e6d0eaa4e087172e4c4.MultiValueExtendedPropertiesRequestBuilder) {
    return i5a6e7d28a08001fc4a358505f17f645b466f8fee88131e6d0eaa4e087172e4c4.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i1c18134f6f07c2ab5e8da11010ba59f9b04bb2e337d6972f9a31737f2824486f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i1c18134f6f07c2ab5e8da11010ba59f9b04bb2e337d6972f9a31737f2824486f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property instances in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ib8e8cb2ab4026e7192d0358e58d577e0c9c266f6d4e08eab7aad7b2ec435d940.SingleValueExtendedPropertiesRequestBuilder) {
    return ib8e8cb2ab4026e7192d0358e58d577e0c9c266f6d4e08eab7aad7b2ec435d940.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iaa3b0089e93d44c50c83ad3a25361adc1c8aedcb3f4d2b7bd69863e1bbefd761.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return iaa3b0089e93d44c50c83ad3a25361adc1c8aedcb3f4d2b7bd69863e1bbefd761.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i94973edce7445a8864c18f665412fc79a47968004ab268fa0b19c1e001974ee3.SnoozeReminderRequestBuilder) {
    return i94973edce7445a8864c18f665412fc79a47968004ab268fa0b19c1e001974ee3.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i7576ba38ad041650f20a21edfa06d72e6982e9b16b1837a42b64c72f2a7db2aa.TentativelyAcceptRequestBuilder) {
    return i7576ba38ad041650f20a21edfa06d72e6982e9b16b1837a42b64c72f2a7db2aa.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
