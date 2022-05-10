package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i405b2d56f34519559149250c12e7d43b6e47d5ab9c198f009bc20501b63f8f2c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i48411056f705ce9de2703670ab279b7ca00fb7b024fc4dbcb411987b24fd4a85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    i6ad1b10f1c3f36819e60623b8ca6a6526eb742956accaefa7397e52c6dfaf0d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i7b6cf93600bbf2c59272fb1a96419d03634399125150869ca86fdacbfbc60629 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i81796e8e2b93871407c3b0ec56e6c7afeb3925154921ff54dcc33136258bc49c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    i86ec6b457759062550f70ea9b194016df3dd7b1f574887f334f83572d925c27b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    ia04fc7bf3030b08f73c59021380bdff93641a6ee04898a1033cfa16a081f83cc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    ic181865131096b2da119bfefef520775d4ba7385c68836ceb96daaaff452b6e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ic41966c08226172b27a25ac14ff05336a3a143e68dcd379f329fb1e2aa07097a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    id634916872597020284e347bdf458ae91d203163707958ebf111f8a6cfd2c4bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    iea033355efd039d9457a7f08713e979f966e0f0eab57d4f18a1b1ed65ce678c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    ifb65fe5ada7ff03f0a44c5a954a76e168a6d8417552736fd9c3bf5cdf376434b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    i4b0bb9c262ec948e789735daef8096fe3274392ed92c67dd2b27b105284e7e34 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ia0f1dc29ce12f23b8c291c1a9ad08236603c8c4af6d55071c6c450dd53057c19 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    ib2762f55b200767004aea91cd9bb2a2c585a57653bf92672ff2926f675f81402 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    if16656d90dd70cdb72c7f8f0ff4e8da4e0424468ba731070bd7145bf968758f0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendar/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*i48411056f705ce9de2703670ab279b7ca00fb7b024fc4dbcb411987b24fd4a85.AcceptRequestBuilder) {
    return i48411056f705ce9de2703670ab279b7ca00fb7b024fc4dbcb411987b24fd4a85.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*id634916872597020284e347bdf458ae91d203163707958ebf111f8a6cfd2c4bd.AttachmentsRequestBuilder) {
    return id634916872597020284e347bdf458ae91d203163707958ebf111f8a6cfd2c4bd.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ib2762f55b200767004aea91cd9bb2a2c585a57653bf92672ff2926f675f81402.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ib2762f55b200767004aea91cd9bb2a2c585a57653bf92672ff2926f675f81402.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ifb65fe5ada7ff03f0a44c5a954a76e168a6d8417552736fd9c3bf5cdf376434b.CalendarRequestBuilder) {
    return ifb65fe5ada7ff03f0a44c5a954a76e168a6d8417552736fd9c3bf5cdf376434b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i86ec6b457759062550f70ea9b194016df3dd7b1f574887f334f83572d925c27b.CancelRequestBuilder) {
    return i86ec6b457759062550f70ea9b194016df3dd7b1f574887f334f83572d925c27b.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i81796e8e2b93871407c3b0ec56e6c7afeb3925154921ff54dcc33136258bc49c.DeclineRequestBuilder) {
    return i81796e8e2b93871407c3b0ec56e6c7afeb3925154921ff54dcc33136258bc49c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i6ad1b10f1c3f36819e60623b8ca6a6526eb742956accaefa7397e52c6dfaf0d1.DismissReminderRequestBuilder) {
    return i6ad1b10f1c3f36819e60623b8ca6a6526eb742956accaefa7397e52c6dfaf0d1.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i405b2d56f34519559149250c12e7d43b6e47d5ab9c198f009bc20501b63f8f2c.ExtensionsRequestBuilder) {
    return i405b2d56f34519559149250c12e7d43b6e47d5ab9c198f009bc20501b63f8f2c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*if16656d90dd70cdb72c7f8f0ff4e8da4e0424468ba731070bd7145bf968758f0.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return if16656d90dd70cdb72c7f8f0ff4e8da4e0424468ba731070bd7145bf968758f0.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ia04fc7bf3030b08f73c59021380bdff93641a6ee04898a1033cfa16a081f83cc.ForwardRequestBuilder) {
    return ia04fc7bf3030b08f73c59021380bdff93641a6ee04898a1033cfa16a081f83cc.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i7b6cf93600bbf2c59272fb1a96419d03634399125150869ca86fdacbfbc60629.MultiValueExtendedPropertiesRequestBuilder) {
    return i7b6cf93600bbf2c59272fb1a96419d03634399125150869ca86fdacbfbc60629.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4b0bb9c262ec948e789735daef8096fe3274392ed92c67dd2b27b105284e7e34.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i4b0bb9c262ec948e789735daef8096fe3274392ed92c67dd2b27b105284e7e34.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ic41966c08226172b27a25ac14ff05336a3a143e68dcd379f329fb1e2aa07097a.SingleValueExtendedPropertiesRequestBuilder) {
    return ic41966c08226172b27a25ac14ff05336a3a143e68dcd379f329fb1e2aa07097a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendar.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia0f1dc29ce12f23b8c291c1a9ad08236603c8c4af6d55071c6c450dd53057c19.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ia0f1dc29ce12f23b8c291c1a9ad08236603c8c4af6d55071c6c450dd53057c19.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*iea033355efd039d9457a7f08713e979f966e0f0eab57d4f18a1b1ed65ce678c7.SnoozeReminderRequestBuilder) {
    return iea033355efd039d9457a7f08713e979f966e0f0eab57d4f18a1b1ed65ce678c7.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ic181865131096b2da119bfefef520775d4ba7385c68836ceb96daaaff452b6e8.TentativelyAcceptRequestBuilder) {
    return ic181865131096b2da119bfefef520775d4ba7385c68836ceb96daaaff452b6e8.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
