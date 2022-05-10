package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0422d04c95be8bbfd84c82c414f364c62ce96964a303ac9fff43fcabcd6fa062 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/forward"
    i0e452fd12b5b875ada5bdf50102bf2a25fae778540616c2cd837b5a7c891d7c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/cancel"
    i34fee5672a5340dcc0fe1e36a379e50cb521fb2dd64f508e3727769c16a38a3c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/decline"
    i4463f64974aee91a14bbe7d308bd483f3fb27df5330489aff5599a3954034ac0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i56b81648f4d65c7f1cc88aa7c13b8b20e6fb3fa03f892126f85b370924dfae10 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/extensions"
    i77c6db270495e0b82aa21013b66a1b2fb572ea6e04ad5eb82ccde0ec5430d610 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i9a3cdff718dea975af25a185cde356034a493a90504214bf5572c77a71468ad4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i9f633f3821995649564b6ac5d3173918ac64a4f317c4d488f52f781388cf028e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/accept"
    ibe57a92c7ff05abcd389cf584941529a70c91182417431f3e25ca5f17b43ca72 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    ic8f47c2462f757b0ee7070212699e8157c2afc29a61186b52085b4b8c27524cd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/attachments"
    idc273b2d569cf4cd17b017dfc3069e03e01e84944c6332837929ff8f65e50aa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    ifa139de6160a9637d146c65690d88a27f9c82e7c6aae067773fe2f22fee1dcb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/calendar"
    i5a5ad989809719e279b7cadedc87dfbf506a89a9fe3f5210d96ebb15fcc44785 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    i5fa23c1207887d36cda91a3cb68b18a8f782a3943d179165aba1049d65e64d74 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    i72a66cac4f09d7fe487bd2ec8ffdbbd4db09345183a7509761c1f6b19877dd92 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    ifed8fd02bfbe9377aa430bf8724b9d47c471434ba519e92ec3ee2275a3be35a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i9f633f3821995649564b6ac5d3173918ac64a4f317c4d488f52f781388cf028e.AcceptRequestBuilder) {
    return i9f633f3821995649564b6ac5d3173918ac64a4f317c4d488f52f781388cf028e.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ic8f47c2462f757b0ee7070212699e8157c2afc29a61186b52085b4b8c27524cd.AttachmentsRequestBuilder) {
    return ic8f47c2462f757b0ee7070212699e8157c2afc29a61186b52085b4b8c27524cd.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i72a66cac4f09d7fe487bd2ec8ffdbbd4db09345183a7509761c1f6b19877dd92.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i72a66cac4f09d7fe487bd2ec8ffdbbd4db09345183a7509761c1f6b19877dd92.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ifa139de6160a9637d146c65690d88a27f9c82e7c6aae067773fe2f22fee1dcb7.CalendarRequestBuilder) {
    return ifa139de6160a9637d146c65690d88a27f9c82e7c6aae067773fe2f22fee1dcb7.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i0e452fd12b5b875ada5bdf50102bf2a25fae778540616c2cd837b5a7c891d7c0.CancelRequestBuilder) {
    return i0e452fd12b5b875ada5bdf50102bf2a25fae778540616c2cd837b5a7c891d7c0.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i34fee5672a5340dcc0fe1e36a379e50cb521fb2dd64f508e3727769c16a38a3c.DeclineRequestBuilder) {
    return i34fee5672a5340dcc0fe1e36a379e50cb521fb2dd64f508e3727769c16a38a3c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*ibe57a92c7ff05abcd389cf584941529a70c91182417431f3e25ca5f17b43ca72.DismissReminderRequestBuilder) {
    return ibe57a92c7ff05abcd389cf584941529a70c91182417431f3e25ca5f17b43ca72.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i56b81648f4d65c7f1cc88aa7c13b8b20e6fb3fa03f892126f85b370924dfae10.ExtensionsRequestBuilder) {
    return i56b81648f4d65c7f1cc88aa7c13b8b20e6fb3fa03f892126f85b370924dfae10.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i5a5ad989809719e279b7cadedc87dfbf506a89a9fe3f5210d96ebb15fcc44785.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i5a5ad989809719e279b7cadedc87dfbf506a89a9fe3f5210d96ebb15fcc44785.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i0422d04c95be8bbfd84c82c414f364c62ce96964a303ac9fff43fcabcd6fa062.ForwardRequestBuilder) {
    return i0422d04c95be8bbfd84c82c414f364c62ce96964a303ac9fff43fcabcd6fa062.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i4463f64974aee91a14bbe7d308bd483f3fb27df5330489aff5599a3954034ac0.MultiValueExtendedPropertiesRequestBuilder) {
    return i4463f64974aee91a14bbe7d308bd483f3fb27df5330489aff5599a3954034ac0.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5fa23c1207887d36cda91a3cb68b18a8f782a3943d179165aba1049d65e64d74.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i5fa23c1207887d36cda91a3cb68b18a8f782a3943d179165aba1049d65e64d74.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*idc273b2d569cf4cd17b017dfc3069e03e01e84944c6332837929ff8f65e50aa3.SingleValueExtendedPropertiesRequestBuilder) {
    return idc273b2d569cf4cd17b017dfc3069e03e01e84944c6332837929ff8f65e50aa3.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ifed8fd02bfbe9377aa430bf8724b9d47c471434ba519e92ec3ee2275a3be35a4.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ifed8fd02bfbe9377aa430bf8724b9d47c471434ba519e92ec3ee2275a3be35a4.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i9a3cdff718dea975af25a185cde356034a493a90504214bf5572c77a71468ad4.SnoozeReminderRequestBuilder) {
    return i9a3cdff718dea975af25a185cde356034a493a90504214bf5572c77a71468ad4.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i77c6db270495e0b82aa21013b66a1b2fb572ea6e04ad5eb82ccde0ec5430d610.TentativelyAcceptRequestBuilder) {
    return i77c6db270495e0b82aa21013b66a1b2fb572ea6e04ad5eb82ccde0ec5430d610.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
