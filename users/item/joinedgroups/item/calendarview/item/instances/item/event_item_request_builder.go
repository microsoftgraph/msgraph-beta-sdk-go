package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i041fb10032ac10f3938964c081f2b9183edeb605ce65756ce10f17d113e97163 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/accept"
    i153fac7b9fcd09f461791aa9f034f17a4e063239824db55ac0a6b55188ed5243 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/multivalueextendedproperties"
    i28d5b841955e95b17324bd59ebfea256f2036b390320ea73bc36d201750b42fd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/cancel"
    i551b1637f3eaa87470573f4c043ed6e2342d167040d96a4b364b234b18f395ee "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/tentativelyaccept"
    i80ea1a246a1e6d1f3f6ea9250e4f530acb1ef27f4d4dd3572f5b96714da3457c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/extensions"
    i82f94a297b400a58cfc36eecaadc0e239a591b14e4b19a6b2f2f80e1ecec4665 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences"
    i8e90421e5340491d785b6d0d4465673a2af529a06acb6b68c679ee8f17d0815c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/calendar"
    i967a0507269674b056e5315471167c7231ba00e90d10a8110f2afbe464412996 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/forward"
    i98396e10385048afadd21710b4e51aad4c1cc7ad9bb8506debd38412417f8d53 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/decline"
    i99877bfcd6ff1d4aecaf990572a774848a5e45757649dbc22d0008afdbd068b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/dismissreminder"
    i99a528bb387ad329b7be7215ea3095a09f61ac034509e46ec26ceabe126b5b7b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/singlevalueextendedproperties"
    ib7ee3b2f85f4ac57854ac2bd9bdf83cb441ff9a9fda3483a383af50d726096b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/snoozereminder"
    ic816482ba62d7eeca8293279613773d1773c83765097f0b79bf189a9609613b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/attachments"
    i7851002e7e69831973ce1d990f022548b6b0334042c63645f351b61130efeb1b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/attachments/item"
    i8b11cd0918fa752a39f085c2931633c7f90c2441b469851e99a7383c7095df3e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
    i947ea5a4030aaa38103570ea7df4801d4df61ec52e4592cfe86ec0055a8924fd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/extensions/item"
    ib59430d6eb73fa755d5794ee3a67d009578b31abf4a2d7e2cb4257d685a957e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/multivalueextendedproperties/item"
    if437dc9fac1fe2d015cffed0645a83936a3013eed9528a881c8fda33e7f6634b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/instances/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*i041fb10032ac10f3938964c081f2b9183edeb605ce65756ce10f17d113e97163.AcceptRequestBuilder) {
    return i041fb10032ac10f3938964c081f2b9183edeb605ce65756ce10f17d113e97163.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ic816482ba62d7eeca8293279613773d1773c83765097f0b79bf189a9609613b6.AttachmentsRequestBuilder) {
    return ic816482ba62d7eeca8293279613773d1773c83765097f0b79bf189a9609613b6.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i7851002e7e69831973ce1d990f022548b6b0334042c63645f351b61130efeb1b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i7851002e7e69831973ce1d990f022548b6b0334042c63645f351b61130efeb1b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i8e90421e5340491d785b6d0d4465673a2af529a06acb6b68c679ee8f17d0815c.CalendarRequestBuilder) {
    return i8e90421e5340491d785b6d0d4465673a2af529a06acb6b68c679ee8f17d0815c.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i28d5b841955e95b17324bd59ebfea256f2036b390320ea73bc36d201750b42fd.CancelRequestBuilder) {
    return i28d5b841955e95b17324bd59ebfea256f2036b390320ea73bc36d201750b42fd.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instances for users
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
// CreatePatchRequestInformation update the navigation property instances in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instances in users
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
func (m *EventItemRequestBuilder) Decline()(*i98396e10385048afadd21710b4e51aad4c1cc7ad9bb8506debd38412417f8d53.DeclineRequestBuilder) {
    return i98396e10385048afadd21710b4e51aad4c1cc7ad9bb8506debd38412417f8d53.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property instances for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i99877bfcd6ff1d4aecaf990572a774848a5e45757649dbc22d0008afdbd068b3.DismissReminderRequestBuilder) {
    return i99877bfcd6ff1d4aecaf990572a774848a5e45757649dbc22d0008afdbd068b3.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i82f94a297b400a58cfc36eecaadc0e239a591b14e4b19a6b2f2f80e1ecec4665.ExceptionOccurrencesRequestBuilder) {
    return i82f94a297b400a58cfc36eecaadc0e239a591b14e4b19a6b2f2f80e1ecec4665.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*if437dc9fac1fe2d015cffed0645a83936a3013eed9528a881c8fda33e7f6634b.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return if437dc9fac1fe2d015cffed0645a83936a3013eed9528a881c8fda33e7f6634b.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i80ea1a246a1e6d1f3f6ea9250e4f530acb1ef27f4d4dd3572f5b96714da3457c.ExtensionsRequestBuilder) {
    return i80ea1a246a1e6d1f3f6ea9250e4f530acb1ef27f4d4dd3572f5b96714da3457c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i947ea5a4030aaa38103570ea7df4801d4df61ec52e4592cfe86ec0055a8924fd.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i947ea5a4030aaa38103570ea7df4801d4df61ec52e4592cfe86ec0055a8924fd.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i967a0507269674b056e5315471167c7231ba00e90d10a8110f2afbe464412996.ForwardRequestBuilder) {
    return i967a0507269674b056e5315471167c7231ba00e90d10a8110f2afbe464412996.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i153fac7b9fcd09f461791aa9f034f17a4e063239824db55ac0a6b55188ed5243.MultiValueExtendedPropertiesRequestBuilder) {
    return i153fac7b9fcd09f461791aa9f034f17a4e063239824db55ac0a6b55188ed5243.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib59430d6eb73fa755d5794ee3a67d009578b31abf4a2d7e2cb4257d685a957e0.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ib59430d6eb73fa755d5794ee3a67d009578b31abf4a2d7e2cb4257d685a957e0.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property instances in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i99a528bb387ad329b7be7215ea3095a09f61ac034509e46ec26ceabe126b5b7b.SingleValueExtendedPropertiesRequestBuilder) {
    return i99a528bb387ad329b7be7215ea3095a09f61ac034509e46ec26ceabe126b5b7b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8b11cd0918fa752a39f085c2931633c7f90c2441b469851e99a7383c7095df3e.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i8b11cd0918fa752a39f085c2931633c7f90c2441b469851e99a7383c7095df3e.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ib7ee3b2f85f4ac57854ac2bd9bdf83cb441ff9a9fda3483a383af50d726096b6.SnoozeReminderRequestBuilder) {
    return ib7ee3b2f85f4ac57854ac2bd9bdf83cb441ff9a9fda3483a383af50d726096b6.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i551b1637f3eaa87470573f4c043ed6e2342d167040d96a4b364b234b18f395ee.TentativelyAcceptRequestBuilder) {
    return i551b1637f3eaa87470573f4c043ed6e2342d167040d96a4b364b234b18f395ee.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
