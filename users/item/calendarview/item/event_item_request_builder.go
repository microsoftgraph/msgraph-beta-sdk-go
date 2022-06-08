package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1bb040a6172da0f3751a7ceda22fda1a4e82f431532706fa23be932b15ce4da7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/accept"
    i49b3399d3f94916e2dcd2961f6cc29b41beca49a880c4d215fb7071373610633 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/multivalueextendedproperties"
    i5273a0eb53fd57878ca0ece6fb7c5aa1c6b53a2713f8fb3cfda89307d1ebbb18 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/extensions"
    i59caea15222b4e5de6959c735972cffb8c4402d6efdf6aaaf8c7c239a8d57d39 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/singlevalueextendedproperties"
    i6229e0dcc67d4b2439cec0174ff358ae49a1d7c4ba1b18919ea2f4ae0ade1775 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/tentativelyaccept"
    i84de11bd1dac1de0c3b8e7b1613efa73af6db81a30531e727741b72681135b66 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/attachments"
    i892c6aafc6d41ce4059451ac393b4a64e888cc586703223009a6984d4044678c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/dismissreminder"
    ia8dfdfa9be3f0f71b1319d005e93ad8157e2145d947f2e0d992b9c8e2aa1edb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/forward"
    ib714f5d74222e2dc5b36353082172e5399fd0349109b9d44f53ac74ae21d9436 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/snoozereminder"
    ic6ce6a431d330b0c7010d3239fc771d8757d7df7f7020314f17078102f22f534 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances"
    id0d2b6dbed65b3df6307ebd38618fdef6a17c12e18ec709bdb429080d1148b9e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar"
    id1287a1e327e26acfdcd299e4732689834b3895fb4db5fae88ccf989bd24a46d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences"
    id822bec9817ce7b201d36abc0694c419616196694c625ded53a514e4b502ad26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/cancel"
    ie4c764f177dd767bef6e8417b11477361fb6c82afd5afecf584ba40ba776a5bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/decline"
    i236482608f31b0fda79b0cbbe06fe1399a6be22f9b861ae3e2b2dc11325de539 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/attachments/item"
    i848daab0afa54e7503df8363a3b46fa1b112c1ddc7abb259a29792f4d55276c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item"
    i85ba546a2606ea05abdcd903570e68f1f0153bb14bedc01ef322f040d645e05e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/multivalueextendedproperties/item"
    ib521e23bb7525673c729e3af728ba92bac1d9ff85f2443be14ac282fba44de57 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/singlevalueextendedproperties/item"
    icdda30c8f71f344f627842a9ad82860cf8d02d3a7f04be4929db3d975aa70ee8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/extensions/item"
    iee92f89d05dcbe8d2cd9b4a2f793efd2c5a69122576cc5ea9173539a69c33ce8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item"
)

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.user entity.
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
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string
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
func (m *EventItemRequestBuilder) Accept()(*i1bb040a6172da0f3751a7ceda22fda1a4e82f431532706fa23be932b15ce4da7.AcceptRequestBuilder) {
    return i1bb040a6172da0f3751a7ceda22fda1a4e82f431532706fa23be932b15ce4da7.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i84de11bd1dac1de0c3b8e7b1613efa73af6db81a30531e727741b72681135b66.AttachmentsRequestBuilder) {
    return i84de11bd1dac1de0c3b8e7b1613efa73af6db81a30531e727741b72681135b66.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i236482608f31b0fda79b0cbbe06fe1399a6be22f9b861ae3e2b2dc11325de539.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i236482608f31b0fda79b0cbbe06fe1399a6be22f9b861ae3e2b2dc11325de539.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*id0d2b6dbed65b3df6307ebd38618fdef6a17c12e18ec709bdb429080d1148b9e.CalendarRequestBuilder) {
    return id0d2b6dbed65b3df6307ebd38618fdef6a17c12e18ec709bdb429080d1148b9e.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*id822bec9817ce7b201d36abc0694c419616196694c625ded53a514e4b502ad26.CancelRequestBuilder) {
    return id822bec9817ce7b201d36abc0694c419616196694c625ded53a514e4b502ad26.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}{?startDateTime,endDateTime,%24select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calendarView for users
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
// CreateGetRequestInformation the calendar view for the calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the calendar view for the calendar. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property calendarView in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calendarView in users
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
func (m *EventItemRequestBuilder) Decline()(*ie4c764f177dd767bef6e8417b11477361fb6c82afd5afecf584ba40ba776a5bd.DeclineRequestBuilder) {
    return ie4c764f177dd767bef6e8417b11477361fb6c82afd5afecf584ba40ba776a5bd.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property calendarView for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i892c6aafc6d41ce4059451ac393b4a64e888cc586703223009a6984d4044678c.DismissReminderRequestBuilder) {
    return i892c6aafc6d41ce4059451ac393b4a64e888cc586703223009a6984d4044678c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*id1287a1e327e26acfdcd299e4732689834b3895fb4db5fae88ccf989bd24a46d.ExceptionOccurrencesRequestBuilder) {
    return id1287a1e327e26acfdcd299e4732689834b3895fb4db5fae88ccf989bd24a46d.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*iee92f89d05dcbe8d2cd9b4a2f793efd2c5a69122576cc5ea9173539a69c33ce8.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return iee92f89d05dcbe8d2cd9b4a2f793efd2c5a69122576cc5ea9173539a69c33ce8.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i5273a0eb53fd57878ca0ece6fb7c5aa1c6b53a2713f8fb3cfda89307d1ebbb18.ExtensionsRequestBuilder) {
    return i5273a0eb53fd57878ca0ece6fb7c5aa1c6b53a2713f8fb3cfda89307d1ebbb18.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*icdda30c8f71f344f627842a9ad82860cf8d02d3a7f04be4929db3d975aa70ee8.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return icdda30c8f71f344f627842a9ad82860cf8d02d3a7f04be4929db3d975aa70ee8.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ia8dfdfa9be3f0f71b1319d005e93ad8157e2145d947f2e0d992b9c8e2aa1edb4.ForwardRequestBuilder) {
    return ia8dfdfa9be3f0f71b1319d005e93ad8157e2145d947f2e0d992b9c8e2aa1edb4.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the calendar view for the calendar. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Instances()(*ic6ce6a431d330b0c7010d3239fc771d8757d7df7f7020314f17078102f22f534.InstancesRequestBuilder) {
    return ic6ce6a431d330b0c7010d3239fc771d8757d7df7f7020314f17078102f22f534.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i848daab0afa54e7503df8363a3b46fa1b112c1ddc7abb259a29792f4d55276c2.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i848daab0afa54e7503df8363a3b46fa1b112c1ddc7abb259a29792f4d55276c2.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i49b3399d3f94916e2dcd2961f6cc29b41beca49a880c4d215fb7071373610633.MultiValueExtendedPropertiesRequestBuilder) {
    return i49b3399d3f94916e2dcd2961f6cc29b41beca49a880c4d215fb7071373610633.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i85ba546a2606ea05abdcd903570e68f1f0153bb14bedc01ef322f040d645e05e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i85ba546a2606ea05abdcd903570e68f1f0153bb14bedc01ef322f040d645e05e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property calendarView in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i59caea15222b4e5de6959c735972cffb8c4402d6efdf6aaaf8c7c239a8d57d39.SingleValueExtendedPropertiesRequestBuilder) {
    return i59caea15222b4e5de6959c735972cffb8c4402d6efdf6aaaf8c7c239a8d57d39.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib521e23bb7525673c729e3af728ba92bac1d9ff85f2443be14ac282fba44de57.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ib521e23bb7525673c729e3af728ba92bac1d9ff85f2443be14ac282fba44de57.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ib714f5d74222e2dc5b36353082172e5399fd0349109b9d44f53ac74ae21d9436.SnoozeReminderRequestBuilder) {
    return ib714f5d74222e2dc5b36353082172e5399fd0349109b9d44f53ac74ae21d9436.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i6229e0dcc67d4b2439cec0174ff358ae49a1d7c4ba1b18919ea2f4ae0ade1775.TentativelyAcceptRequestBuilder) {
    return i6229e0dcc67d4b2439cec0174ff358ae49a1d7c4ba1b18919ea2f4ae0ade1775.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
