package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i25c258057200bb0f49fe695e4a04b563a49a78ea1b8d683224229ecc2216f2a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances"
    i48852bc54e98648603615f21e8fe236e812e000865f7f4c50233739217e4d1e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/tentativelyaccept"
    i55f0917eabd722c0b1a208a3899d358fa90ca8d177365df27f6e8b3741b91b11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/cancel"
    i6ac110054bf4866ebd4c71125f70f8ed4944d4e5a1e4a4f60ec33866136a21dc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/dismissreminder"
    i77dbc8b0b312e9997d696a64628fb76fbd74dca2b57dad9b55a8aaeac2226883 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/snoozereminder"
    i7cb1b521cdcddebe65f30672b8c67221c12155ef96f0ce5cb7f15ba40d2b53b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/attachments"
    i8fc8df027938e6c57d6a20fed58f975aff99e1b08ccd8cd6d0155b73942df792 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties"
    i91ef35b5ec9cae5a9c036e761d8e2be8be1d20a8e9150854611f0bf715cf73d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/calendar"
    i93edb1ee8f13a1618935932133f24016779d6669f623b4490fbbdd678515d191 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/forward"
    iaebfa9415ee51c6dc7a11bd735bcb48cc6d59a447064f4ad81c5d165950b3e8e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/extensions"
    ibdc76841b079a0f188daa2353850e249b532969e1cb0485c6094353bdc5f3930 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/accept"
    icc729204df03f40d1aa2f95863ea112bb68d4be2c94ca964621c4a7014a3ac33 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/decline"
    id6ca25b6555ad7c56f58e1bf21643fa0cfd6ef1dc4e6414360aa3b916d23d32d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/multivalueextendedproperties"
    i291a6f0c68faa3f60f8ebad637c88cff44b789be9064d3c42c7af21f0c1041a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i2d93095d60bcd79483609d3222e456dda61fc9ca8bb9109b27f24e904d333421 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/attachments/item"
    i39a537deb3ece6da6f5c941a4a6c877fafe08898d46300e4a18dd7b209d2752e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i9e21713a2ffb327ba755e9e84494cca171566fabae8c73f12284ce0b5ea1710e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/instances/item"
    ia35ae6531a2d15573b6ddbc296f2787be67a3cd79e5e7eec7365225122114987 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendargroups/item/calendars/item/events/item/exceptionoccurrences/item/extensions/item"
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
func (m *EventItemRequestBuilder) Accept()(*ibdc76841b079a0f188daa2353850e249b532969e1cb0485c6094353bdc5f3930.AcceptRequestBuilder) {
    return ibdc76841b079a0f188daa2353850e249b532969e1cb0485c6094353bdc5f3930.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i7cb1b521cdcddebe65f30672b8c67221c12155ef96f0ce5cb7f15ba40d2b53b4.AttachmentsRequestBuilder) {
    return i7cb1b521cdcddebe65f30672b8c67221c12155ef96f0ce5cb7f15ba40d2b53b4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i2d93095d60bcd79483609d3222e456dda61fc9ca8bb9109b27f24e904d333421.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i2d93095d60bcd79483609d3222e456dda61fc9ca8bb9109b27f24e904d333421.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i91ef35b5ec9cae5a9c036e761d8e2be8be1d20a8e9150854611f0bf715cf73d3.CalendarRequestBuilder) {
    return i91ef35b5ec9cae5a9c036e761d8e2be8be1d20a8e9150854611f0bf715cf73d3.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i55f0917eabd722c0b1a208a3899d358fa90ca8d177365df27f6e8b3741b91b11.CancelRequestBuilder) {
    return i55f0917eabd722c0b1a208a3899d358fa90ca8d177365df27f6e8b3741b91b11.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*icc729204df03f40d1aa2f95863ea112bb68d4be2c94ca964621c4a7014a3ac33.DeclineRequestBuilder) {
    return icc729204df03f40d1aa2f95863ea112bb68d4be2c94ca964621c4a7014a3ac33.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i6ac110054bf4866ebd4c71125f70f8ed4944d4e5a1e4a4f60ec33866136a21dc.DismissReminderRequestBuilder) {
    return i6ac110054bf4866ebd4c71125f70f8ed4944d4e5a1e4a4f60ec33866136a21dc.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*iaebfa9415ee51c6dc7a11bd735bcb48cc6d59a447064f4ad81c5d165950b3e8e.ExtensionsRequestBuilder) {
    return iaebfa9415ee51c6dc7a11bd735bcb48cc6d59a447064f4ad81c5d165950b3e8e.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ia35ae6531a2d15573b6ddbc296f2787be67a3cd79e5e7eec7365225122114987.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ia35ae6531a2d15573b6ddbc296f2787be67a3cd79e5e7eec7365225122114987.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i93edb1ee8f13a1618935932133f24016779d6669f623b4490fbbdd678515d191.ForwardRequestBuilder) {
    return i93edb1ee8f13a1618935932133f24016779d6669f623b4490fbbdd678515d191.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// Instances the instances property
func (m *EventItemRequestBuilder) Instances()(*i25c258057200bb0f49fe695e4a04b563a49a78ea1b8d683224229ecc2216f2a3.InstancesRequestBuilder) {
    return i25c258057200bb0f49fe695e4a04b563a49a78ea1b8d683224229ecc2216f2a3.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i9e21713a2ffb327ba755e9e84494cca171566fabae8c73f12284ce0b5ea1710e.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i9e21713a2ffb327ba755e9e84494cca171566fabae8c73f12284ce0b5ea1710e.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*id6ca25b6555ad7c56f58e1bf21643fa0cfd6ef1dc4e6414360aa3b916d23d32d.MultiValueExtendedPropertiesRequestBuilder) {
    return id6ca25b6555ad7c56f58e1bf21643fa0cfd6ef1dc4e6414360aa3b916d23d32d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i39a537deb3ece6da6f5c941a4a6c877fafe08898d46300e4a18dd7b209d2752e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i39a537deb3ece6da6f5c941a4a6c877fafe08898d46300e4a18dd7b209d2752e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i8fc8df027938e6c57d6a20fed58f975aff99e1b08ccd8cd6d0155b73942df792.SingleValueExtendedPropertiesRequestBuilder) {
    return i8fc8df027938e6c57d6a20fed58f975aff99e1b08ccd8cd6d0155b73942df792.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarGroups.item.calendars.item.events.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i291a6f0c68faa3f60f8ebad637c88cff44b789be9064d3c42c7af21f0c1041a7.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i291a6f0c68faa3f60f8ebad637c88cff44b789be9064d3c42c7af21f0c1041a7.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i77dbc8b0b312e9997d696a64628fb76fbd74dca2b57dad9b55a8aaeac2226883.SnoozeReminderRequestBuilder) {
    return i77dbc8b0b312e9997d696a64628fb76fbd74dca2b57dad9b55a8aaeac2226883.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i48852bc54e98648603615f21e8fe236e812e000865f7f4c50233739217e4d1e0.TentativelyAcceptRequestBuilder) {
    return i48852bc54e98648603615f21e8fe236e812e000865f7f4c50233739217e4d1e0.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
