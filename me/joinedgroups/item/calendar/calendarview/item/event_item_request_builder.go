package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09fe811e7817b216670fcd0bbb6f01b6a656b609ed37cadd42b76bc845dfde42 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/attachments"
    i20eab9b2716b01ce47ed2ac38559245133ca432dc4f419627e2693d382901f32 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/forward"
    i299a789abb0595f84b846e00cdada3a494f47fd5dd41db4b6a7717765689ebad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences"
    i3c92f193c61e7aa7bc9bf562f9ca49cdca1d5a7b5edfc885b1bcb0d54cab4655 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/cancel"
    i3e8560cda037bfc90d7a3b440ed8e00a5a93950d84828b20101f13c319670c5a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/multivalueextendedproperties"
    i6942dacd5610322900cf5e1f1ac0bbc9cd3de4bda5277544418f21c017b1a3f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/calendar"
    i72eca700e1f8e503d934e6cad00415e27e6e471fe4c77f93b03e3206ddbef1a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/singlevalueextendedproperties"
    i75473ed1b1a9b5d0143270466a05f97cc4602ac753d2c765ff4f1fcbc6c805e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/decline"
    ia8a5ad4649c2c0e08140337c3b426cbeebb9550c9ca91036d6432f99cce79e0f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/extensions"
    iac0babf3df60b355cb656f8f6766d7770372fe64c4c3c21bf5f3a5b31566bcd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/snoozereminder"
    id6b3c46f929db123a9757475dd6bea248f327e7f78b7562beb1c79ea95b127b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/tentativelyaccept"
    id7744fa2c85618d48bca505a2b416d98217f5a7e179eb4f2691c2501bcfc76a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances"
    ie19c25abf17f829d7b8880d87c25139b059928b8e54ac38afbb648d0eac31d04 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/accept"
    ie729ce2b51cd17cb6f66ab87c86389efa6bd8b197ab5be11f7964f9823790d7e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/dismissreminder"
    i01af0fe462e29f75d26daa638af47010e9ae6752b9a962d9d4ceeae1e91a650b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/exceptionoccurrences/item"
    i08c3cda5f51e8a0c4673391fdd805bc31da9763ec26fb1b24c1d9da325a73fd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/extensions/item"
    i45deda81dc9cbd4703121dfdd236b6e5740bc72617367cc7688310b9bd59d675 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/attachments/item"
    i79b267922b9e1333a4f86613021dcde6c5b787a45ba22b9b19fd094e2bf82f2d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/multivalueextendedproperties/item"
    ia2e520ec941e0dc2338ccf7ec9acec324f7843ea4cca3af15412cf61ca5853e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/instances/item"
    ib6e93320e6518f50a7fa8eeb528bbd8ac97bb033ed77c35d3eeb58793e1aa21c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/calendarview/item/singlevalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.calendar entity.
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
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Accept()(*ie19c25abf17f829d7b8880d87c25139b059928b8e54ac38afbb648d0eac31d04.AcceptRequestBuilder) {
    return ie19c25abf17f829d7b8880d87c25139b059928b8e54ac38afbb648d0eac31d04.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i09fe811e7817b216670fcd0bbb6f01b6a656b609ed37cadd42b76bc845dfde42.AttachmentsRequestBuilder) {
    return i09fe811e7817b216670fcd0bbb6f01b6a656b609ed37cadd42b76bc845dfde42.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i45deda81dc9cbd4703121dfdd236b6e5740bc72617367cc7688310b9bd59d675.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i45deda81dc9cbd4703121dfdd236b6e5740bc72617367cc7688310b9bd59d675.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i6942dacd5610322900cf5e1f1ac0bbc9cd3de4bda5277544418f21c017b1a3f0.CalendarRequestBuilder) {
    return i6942dacd5610322900cf5e1f1ac0bbc9cd3de4bda5277544418f21c017b1a3f0.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i3c92f193c61e7aa7bc9bf562f9ca49cdca1d5a7b5edfc885b1bcb0d54cab4655.CancelRequestBuilder) {
    return i3c92f193c61e7aa7bc9bf562f9ca49cdca1d5a7b5edfc885b1bcb0d54cab4655.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendar/calendarView/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calendarView for me
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
// CreateGetRequestInformation the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the calendar view for the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation update the navigation property calendarView in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calendarView in me
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
func (m *EventItemRequestBuilder) Decline()(*i75473ed1b1a9b5d0143270466a05f97cc4602ac753d2c765ff4f1fcbc6c805e3.DeclineRequestBuilder) {
    return i75473ed1b1a9b5d0143270466a05f97cc4602ac753d2c765ff4f1fcbc6c805e3.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property calendarView for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*ie729ce2b51cd17cb6f66ab87c86389efa6bd8b197ab5be11f7964f9823790d7e.DismissReminderRequestBuilder) {
    return ie729ce2b51cd17cb6f66ab87c86389efa6bd8b197ab5be11f7964f9823790d7e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i299a789abb0595f84b846e00cdada3a494f47fd5dd41db4b6a7717765689ebad.ExceptionOccurrencesRequestBuilder) {
    return i299a789abb0595f84b846e00cdada3a494f47fd5dd41db4b6a7717765689ebad.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i01af0fe462e29f75d26daa638af47010e9ae6752b9a962d9d4ceeae1e91a650b.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i01af0fe462e29f75d26daa638af47010e9ae6752b9a962d9d4ceeae1e91a650b.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ia8a5ad4649c2c0e08140337c3b426cbeebb9550c9ca91036d6432f99cce79e0f.ExtensionsRequestBuilder) {
    return ia8a5ad4649c2c0e08140337c3b426cbeebb9550c9ca91036d6432f99cce79e0f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i08c3cda5f51e8a0c4673391fdd805bc31da9763ec26fb1b24c1d9da325a73fd6.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i08c3cda5f51e8a0c4673391fdd805bc31da9763ec26fb1b24c1d9da325a73fd6.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i20eab9b2716b01ce47ed2ac38559245133ca432dc4f419627e2693d382901f32.ForwardRequestBuilder) {
    return i20eab9b2716b01ce47ed2ac38559245133ca432dc4f419627e2693d382901f32.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the calendar view for the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*id7744fa2c85618d48bca505a2b416d98217f5a7e179eb4f2691c2501bcfc76a6.InstancesRequestBuilder) {
    return id7744fa2c85618d48bca505a2b416d98217f5a7e179eb4f2691c2501bcfc76a6.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ia2e520ec941e0dc2338ccf7ec9acec324f7843ea4cca3af15412cf61ca5853e7.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return ia2e520ec941e0dc2338ccf7ec9acec324f7843ea4cca3af15412cf61ca5853e7.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i3e8560cda037bfc90d7a3b440ed8e00a5a93950d84828b20101f13c319670c5a.MultiValueExtendedPropertiesRequestBuilder) {
    return i3e8560cda037bfc90d7a3b440ed8e00a5a93950d84828b20101f13c319670c5a.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i79b267922b9e1333a4f86613021dcde6c5b787a45ba22b9b19fd094e2bf82f2d.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i79b267922b9e1333a4f86613021dcde6c5b787a45ba22b9b19fd094e2bf82f2d.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property calendarView in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i72eca700e1f8e503d934e6cad00415e27e6e471fe4c77f93b03e3206ddbef1a5.SingleValueExtendedPropertiesRequestBuilder) {
    return i72eca700e1f8e503d934e6cad00415e27e6e471fe4c77f93b03e3206ddbef1a5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib6e93320e6518f50a7fa8eeb528bbd8ac97bb033ed77c35d3eeb58793e1aa21c.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ib6e93320e6518f50a7fa8eeb528bbd8ac97bb033ed77c35d3eeb58793e1aa21c.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*iac0babf3df60b355cb656f8f6766d7770372fe64c4c3c21bf5f3a5b31566bcd0.SnoozeReminderRequestBuilder) {
    return iac0babf3df60b355cb656f8f6766d7770372fe64c4c3c21bf5f3a5b31566bcd0.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*id6b3c46f929db123a9757475dd6bea248f327e7f78b7562beb1c79ea95b127b9.TentativelyAcceptRequestBuilder) {
    return id6b3c46f929db123a9757475dd6bea248f327e7f78b7562beb1c79ea95b127b9.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
