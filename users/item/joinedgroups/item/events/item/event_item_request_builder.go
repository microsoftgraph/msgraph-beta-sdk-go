package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2bf57bca818b725517c2b0fdfaba95e9cf249ffca89b44b5231ac50e5fe5668c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/calendar"
    i2fbe31b9e58c8011a506e42b69dbe725cda7df6eddc2fccdf90ab188cef4eaef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances"
    i37379edb51e9bfb9fd96d7da9032ae2324aa15be9d7a9593a73fe1707949ce77 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/dismissreminder"
    i3caae7dfc96b790d05e69ce05f15176feab39242d0d147d046d9648c0711b057 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/extensions"
    i3d338a113c93e7582304bd8b5b44d868bab87df56164db1a008ce9c79f8d7a6b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/cancel"
    i3d9ce5dab909d126d804878ab7addf19f236946a9029b4e6b79907fbc4473fc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/snoozereminder"
    i4e0b1e518a057b91a764866e9a3c6ac53c7c39cd008446a44671dd71f6d755a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/decline"
    i52752e866f43a0d233927d97c89683b8a55b2ecbf7c25680731432bcd5f6f1e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/forward"
    ia6f821902d72def2f58fb1b2edd42472cd97755c5aff5b5a4982268f16e9e022 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/singlevalueextendedproperties"
    ibf1de74c8e4433073be00d6896214dc19351e0df6099361504be7fcff7415d5d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/tentativelyaccept"
    ic013201e8bea67e93a07faa2cc26102980b3b9f59ff7afad3a906f09b11d1341 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/attachments"
    idc8b43e2375981ed5de51066606660782cfb6e62d9a75e763f268863af9cf276 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/accept"
    iddb073a82babc7a8ee88b9b01f25b2ce7d83c84b612dd18854de25c97023e858 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences"
    ifa9b467e05db5decfa9232c7a8a223f16c1e6278eed45c68c314f87e7e1d42cd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/multivalueextendedproperties"
    i0c96259699a345fe039fa8d0001e9e2089a7094e0cb981e1d79777183e05ca12 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/exceptionoccurrences/item"
    i0fcf4bc4ce61037e3f51b19d7dbbde464c111ab6d2ffb0c21d6bdbdea50ac2b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/multivalueextendedproperties/item"
    i203532e59bec5ec86eecac6f4d413fb6881434643f3ea823a3b5800e005b2bfd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/singlevalueextendedproperties/item"
    i4d2aaa0543f77339259e856537edf846b432a1aae21c75f488a02a3cc938a9e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/instances/item"
    i8f3c6f0678de6e805062cc35366671414f52f990bfe080b72c065e289320dc73 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/extensions/item"
    if3ac1315c1e64d8393271fe42605f2daee779315fef1168d270d252549d29384 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/events/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the events property of the microsoft.graph.group entity.
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
// EventItemRequestBuilderGetQueryParameters the group's events.
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
func (m *EventItemRequestBuilder) Accept()(*idc8b43e2375981ed5de51066606660782cfb6e62d9a75e763f268863af9cf276.AcceptRequestBuilder) {
    return idc8b43e2375981ed5de51066606660782cfb6e62d9a75e763f268863af9cf276.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ic013201e8bea67e93a07faa2cc26102980b3b9f59ff7afad3a906f09b11d1341.AttachmentsRequestBuilder) {
    return ic013201e8bea67e93a07faa2cc26102980b3b9f59ff7afad3a906f09b11d1341.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*if3ac1315c1e64d8393271fe42605f2daee779315fef1168d270d252549d29384.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return if3ac1315c1e64d8393271fe42605f2daee779315fef1168d270d252549d29384.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i2bf57bca818b725517c2b0fdfaba95e9cf249ffca89b44b5231ac50e5fe5668c.CalendarRequestBuilder) {
    return i2bf57bca818b725517c2b0fdfaba95e9cf249ffca89b44b5231ac50e5fe5668c.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i3d338a113c93e7582304bd8b5b44d868bab87df56164db1a008ce9c79f8d7a6b.CancelRequestBuilder) {
    return i3d338a113c93e7582304bd8b5b44d868bab87df56164db1a008ce9c79f8d7a6b.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/events/{event%2Did}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property events for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property events for users
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
// CreateGetRequestInformation the group's events.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the group's events.
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
// CreatePatchRequestInformation update the navigation property events in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property events in users
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
func (m *EventItemRequestBuilder) Decline()(*i4e0b1e518a057b91a764866e9a3c6ac53c7c39cd008446a44671dd71f6d755a2.DeclineRequestBuilder) {
    return i4e0b1e518a057b91a764866e9a3c6ac53c7c39cd008446a44671dd71f6d755a2.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property events for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property events for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i37379edb51e9bfb9fd96d7da9032ae2324aa15be9d7a9593a73fe1707949ce77.DismissReminderRequestBuilder) {
    return i37379edb51e9bfb9fd96d7da9032ae2324aa15be9d7a9593a73fe1707949ce77.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*iddb073a82babc7a8ee88b9b01f25b2ce7d83c84b612dd18854de25c97023e858.ExceptionOccurrencesRequestBuilder) {
    return iddb073a82babc7a8ee88b9b01f25b2ce7d83c84b612dd18854de25c97023e858.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i0c96259699a345fe039fa8d0001e9e2089a7094e0cb981e1d79777183e05ca12.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i0c96259699a345fe039fa8d0001e9e2089a7094e0cb981e1d79777183e05ca12.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i3caae7dfc96b790d05e69ce05f15176feab39242d0d147d046d9648c0711b057.ExtensionsRequestBuilder) {
    return i3caae7dfc96b790d05e69ce05f15176feab39242d0d147d046d9648c0711b057.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i8f3c6f0678de6e805062cc35366671414f52f990bfe080b72c065e289320dc73.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8f3c6f0678de6e805062cc35366671414f52f990bfe080b72c065e289320dc73.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i52752e866f43a0d233927d97c89683b8a55b2ecbf7c25680731432bcd5f6f1e7.ForwardRequestBuilder) {
    return i52752e866f43a0d233927d97c89683b8a55b2ecbf7c25680731432bcd5f6f1e7.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the group's events.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the group's events.
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
func (m *EventItemRequestBuilder) Instances()(*i2fbe31b9e58c8011a506e42b69dbe725cda7df6eddc2fccdf90ab188cef4eaef.InstancesRequestBuilder) {
    return i2fbe31b9e58c8011a506e42b69dbe725cda7df6eddc2fccdf90ab188cef4eaef.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i4d2aaa0543f77339259e856537edf846b432a1aae21c75f488a02a3cc938a9e6.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i4d2aaa0543f77339259e856537edf846b432a1aae21c75f488a02a3cc938a9e6.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ifa9b467e05db5decfa9232c7a8a223f16c1e6278eed45c68c314f87e7e1d42cd.MultiValueExtendedPropertiesRequestBuilder) {
    return ifa9b467e05db5decfa9232c7a8a223f16c1e6278eed45c68c314f87e7e1d42cd.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0fcf4bc4ce61037e3f51b19d7dbbde464c111ab6d2ffb0c21d6bdbdea50ac2b7.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i0fcf4bc4ce61037e3f51b19d7dbbde464c111ab6d2ffb0c21d6bdbdea50ac2b7.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property events in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property events in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ia6f821902d72def2f58fb1b2edd42472cd97755c5aff5b5a4982268f16e9e022.SingleValueExtendedPropertiesRequestBuilder) {
    return ia6f821902d72def2f58fb1b2edd42472cd97755c5aff5b5a4982268f16e9e022.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i203532e59bec5ec86eecac6f4d413fb6881434643f3ea823a3b5800e005b2bfd.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i203532e59bec5ec86eecac6f4d413fb6881434643f3ea823a3b5800e005b2bfd.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i3d9ce5dab909d126d804878ab7addf19f236946a9029b4e6b79907fbc4473fc7.SnoozeReminderRequestBuilder) {
    return i3d9ce5dab909d126d804878ab7addf19f236946a9029b4e6b79907fbc4473fc7.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ibf1de74c8e4433073be00d6896214dc19351e0df6099361504be7fcff7415d5d.TentativelyAcceptRequestBuilder) {
    return ibf1de74c8e4433073be00d6896214dc19351e0df6099361504be7fcff7415d5d.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
