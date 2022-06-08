package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04c719c895acec6c592a8d7a8c99daedee4d6e03576774242ac518f4ca97f22d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/cancel"
    i0894b6f6226727a9418b410b051c30ca7a83dbf5724623855be38e9b8fc84136 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/forward"
    i08d9c816ec9504a1ccd4dd0d6cd8314afa08cb82d1b91e11064c54e552604d27 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences"
    i0d950e836e4744d2f557f6e0ef4ec3b62f3c1080cd446b0354fdbfec72951f7d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/dismissreminder"
    i0fcaa0c479a3c654579f8181a39c814649f1d5464b954e178ccbc7be2305818a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/attachments"
    i1cd61b81d19facac87465cbccad8f16fa4806738ff2736b95214e7814016ce2c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances"
    i6bcd189817e40fc5716d15a67c6a972c603e467a6275999b2cda08503ca1b274 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/accept"
    iaaec2ffe60b7e88a7eec860d9dda02fca6f1fd301915aa6e7ed5433699a22ece "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/extensions"
    ib4dcf73ff3c2130c3449325e737109b0274434aa8c7a4e3586acaa00f0698c32 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/snoozereminder"
    ibd0ab74cf035804e4a297f74c17c9116e6d51fbbf5ab1e68f8a3a787f4072ea2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/calendar"
    id2289c2a6f838883659e046083563cf60dcd5e984c949f14b689fe8418054ed1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/multivalueextendedproperties"
    idcfc0f6f13cfe31fd44775b67fb3c4db7708a9473bba0eb0ec30fefbeca28ba6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties"
    ie03c784e30aa926a3fc3ca0c4ee7437288f53eb208b57774aaf64385b3bc24dd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/tentativelyaccept"
    if72ad66e54b22cbc558eeb86a695bb7ac2d89c6beaf41cb86224d55c3fe01396 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/decline"
    i094f05a8ac64a75e4ef8691bd5e1036db5cd0689451270938b4712f947bd811c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/singlevalueextendedproperties/item"
    i687127d3891d370eb8d20a210a67df658676908c6ec8ff74ffc0251bc61b5039 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/attachments/item"
    i69d244b2e4785f363c5b08786031cfdc6f1ea6786860dbba8bd696eb8bdf1eb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/multivalueextendedproperties/item"
    ia4845a22710fbc6645ada0a9de578d7004512c8ce99f327697a9e3148117b4e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/instances/item"
    ib911879fbf568651ae56e76797cf7895cbd515a33dfeb908d9088187feb3d43b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/extensions/item"
    if4a386b83ac6acb1bc708653d19283aa35f0d4ef6552877c7f9095c9bc2a807a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.calendar entity.
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
// EventItemRequestBuilderGetQueryParameters the events in the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Accept()(*i6bcd189817e40fc5716d15a67c6a972c603e467a6275999b2cda08503ca1b274.AcceptRequestBuilder) {
    return i6bcd189817e40fc5716d15a67c6a972c603e467a6275999b2cda08503ca1b274.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i0fcaa0c479a3c654579f8181a39c814649f1d5464b954e178ccbc7be2305818a.AttachmentsRequestBuilder) {
    return i0fcaa0c479a3c654579f8181a39c814649f1d5464b954e178ccbc7be2305818a.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i687127d3891d370eb8d20a210a67df658676908c6ec8ff74ffc0251bc61b5039.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i687127d3891d370eb8d20a210a67df658676908c6ec8ff74ffc0251bc61b5039.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ibd0ab74cf035804e4a297f74c17c9116e6d51fbbf5ab1e68f8a3a787f4072ea2.CalendarRequestBuilder) {
    return ibd0ab74cf035804e4a297f74c17c9116e6d51fbbf5ab1e68f8a3a787f4072ea2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i04c719c895acec6c592a8d7a8c99daedee4d6e03576774242ac518f4ca97f22d.CancelRequestBuilder) {
    return i04c719c895acec6c592a8d7a8c99daedee4d6e03576774242ac518f4ca97f22d.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property events for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property events for me
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
// CreateGetRequestInformation the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the events in the calendar. Navigation property. Read-only.
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
// CreatePatchRequestInformation update the navigation property events in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property events in me
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
func (m *EventItemRequestBuilder) Decline()(*if72ad66e54b22cbc558eeb86a695bb7ac2d89c6beaf41cb86224d55c3fe01396.DeclineRequestBuilder) {
    return if72ad66e54b22cbc558eeb86a695bb7ac2d89c6beaf41cb86224d55c3fe01396.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property events for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i0d950e836e4744d2f557f6e0ef4ec3b62f3c1080cd446b0354fdbfec72951f7d.DismissReminderRequestBuilder) {
    return i0d950e836e4744d2f557f6e0ef4ec3b62f3c1080cd446b0354fdbfec72951f7d.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i08d9c816ec9504a1ccd4dd0d6cd8314afa08cb82d1b91e11064c54e552604d27.ExceptionOccurrencesRequestBuilder) {
    return i08d9c816ec9504a1ccd4dd0d6cd8314afa08cb82d1b91e11064c54e552604d27.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*if4a386b83ac6acb1bc708653d19283aa35f0d4ef6552877c7f9095c9bc2a807a.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return if4a386b83ac6acb1bc708653d19283aa35f0d4ef6552877c7f9095c9bc2a807a.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*iaaec2ffe60b7e88a7eec860d9dda02fca6f1fd301915aa6e7ed5433699a22ece.ExtensionsRequestBuilder) {
    return iaaec2ffe60b7e88a7eec860d9dda02fca6f1fd301915aa6e7ed5433699a22ece.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ib911879fbf568651ae56e76797cf7895cbd515a33dfeb908d9088187feb3d43b.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ib911879fbf568651ae56e76797cf7895cbd515a33dfeb908d9088187feb3d43b.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i0894b6f6226727a9418b410b051c30ca7a83dbf5724623855be38e9b8fc84136.ForwardRequestBuilder) {
    return i0894b6f6226727a9418b410b051c30ca7a83dbf5724623855be38e9b8fc84136.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the events in the calendar. Navigation property. Read-only.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the events in the calendar. Navigation property. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*i1cd61b81d19facac87465cbccad8f16fa4806738ff2736b95214e7814016ce2c.InstancesRequestBuilder) {
    return i1cd61b81d19facac87465cbccad8f16fa4806738ff2736b95214e7814016ce2c.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*ia4845a22710fbc6645ada0a9de578d7004512c8ce99f327697a9e3148117b4e2.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return ia4845a22710fbc6645ada0a9de578d7004512c8ce99f327697a9e3148117b4e2.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*id2289c2a6f838883659e046083563cf60dcd5e984c949f14b689fe8418054ed1.MultiValueExtendedPropertiesRequestBuilder) {
    return id2289c2a6f838883659e046083563cf60dcd5e984c949f14b689fe8418054ed1.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i69d244b2e4785f363c5b08786031cfdc6f1ea6786860dbba8bd696eb8bdf1eb9.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i69d244b2e4785f363c5b08786031cfdc6f1ea6786860dbba8bd696eb8bdf1eb9.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property events in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*idcfc0f6f13cfe31fd44775b67fb3c4db7708a9473bba0eb0ec30fefbeca28ba6.SingleValueExtendedPropertiesRequestBuilder) {
    return idcfc0f6f13cfe31fd44775b67fb3c4db7708a9473bba0eb0ec30fefbeca28ba6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i094f05a8ac64a75e4ef8691bd5e1036db5cd0689451270938b4712f947bd811c.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i094f05a8ac64a75e4ef8691bd5e1036db5cd0689451270938b4712f947bd811c.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ib4dcf73ff3c2130c3449325e737109b0274434aa8c7a4e3586acaa00f0698c32.SnoozeReminderRequestBuilder) {
    return ib4dcf73ff3c2130c3449325e737109b0274434aa8c7a4e3586acaa00f0698c32.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ie03c784e30aa926a3fc3ca0c4ee7437288f53eb208b57774aaf64385b3bc24dd.TentativelyAcceptRequestBuilder) {
    return ie03c784e30aa926a3fc3ca0c4ee7437288f53eb208b57774aaf64385b3bc24dd.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
