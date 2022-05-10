package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0c7906a46cd3570c58378ded74411de15227e2d779e5655aebe8225e8af872e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances"
    i1073697e475927811aad904547f6513fa5f74982d0d3ff0e75282b340778206f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/forward"
    i28d4b967a0a453124dd6bd0e0197bbaa48accac70c41fa2dbb27689e69adfb9e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties"
    i29e043eaee31491b11c98cf95ab9a206543eb504e91e849ce5dc75c11db368d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/tentativelyaccept"
    i6715da1957ca86f82d8b761208b8373f87296676acffecc74f236992ddd7ff4a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/snoozereminder"
    i6a98992877b2e129de526cfa1aca5e316bc346081012dab77bb14ef48e3bce07 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/cancel"
    i9a8c6b4340d6510871a4b91322948f851de0b3619b8ad8736e7986c8384387ab "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/attachments"
    ibda2764d6365b6907754db6ca7f56ef5fc34fe39ecfcbe8b94801a1987537b3b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ibf32607336aa93147540fe4f97af179c1e03e2bac45ef40bbde5c3a10d73a633 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/accept"
    ic710dd6c5471d6d93304043e7d091713f6b427d95f336ebeb7aafda37bd0573e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/dismissreminder"
    ie8fac8bea605bb51b108d5398ca815fe8b8db776d57ea99f33548dde8fb09abc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/extensions"
    ieb3cb10eba4179bf61ec7539bd73d8c7afebf75eeafec9071b08a74bc74eb32d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/decline"
    if2bfcc8d0068bceb5a99eae0bdbcebe16c10bbf9fb489ee164b456738a35e486 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/calendar"
    i296762b0003986a4bc9f5bb4934735dbd6c510000fe4111f0c99c7ca58893ee3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/extensions/item"
    i40ad88377bdcea87a9bf81f4f20ac56e52c135f41bcf17f343b2f17f8c17ed2d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item"
    i421482755bd37443917e084a854e179974c6a4fdc954958616023e9509133dc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i81024da70f9ee34f7d5b4be69d2ccdd5628df29fdbc0e2ceef48cb4962b2ba50 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i88cba6b682a6c3a30375e2b158fe63b0052632fa669889a93064b2a4d74ead56 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendarview/item/exceptionoccurrences/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*ibf32607336aa93147540fe4f97af179c1e03e2bac45ef40bbde5c3a10d73a633.AcceptRequestBuilder) {
    return ibf32607336aa93147540fe4f97af179c1e03e2bac45ef40bbde5c3a10d73a633.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i9a8c6b4340d6510871a4b91322948f851de0b3619b8ad8736e7986c8384387ab.AttachmentsRequestBuilder) {
    return i9a8c6b4340d6510871a4b91322948f851de0b3619b8ad8736e7986c8384387ab.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i88cba6b682a6c3a30375e2b158fe63b0052632fa669889a93064b2a4d74ead56.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i88cba6b682a6c3a30375e2b158fe63b0052632fa669889a93064b2a4d74ead56.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*if2bfcc8d0068bceb5a99eae0bdbcebe16c10bbf9fb489ee164b456738a35e486.CalendarRequestBuilder) {
    return if2bfcc8d0068bceb5a99eae0bdbcebe16c10bbf9fb489ee164b456738a35e486.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i6a98992877b2e129de526cfa1aca5e316bc346081012dab77bb14ef48e3bce07.CancelRequestBuilder) {
    return i6a98992877b2e129de526cfa1aca5e316bc346081012dab77bb14ef48e3bce07.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*ieb3cb10eba4179bf61ec7539bd73d8c7afebf75eeafec9071b08a74bc74eb32d.DeclineRequestBuilder) {
    return ieb3cb10eba4179bf61ec7539bd73d8c7afebf75eeafec9071b08a74bc74eb32d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*ic710dd6c5471d6d93304043e7d091713f6b427d95f336ebeb7aafda37bd0573e.DismissReminderRequestBuilder) {
    return ic710dd6c5471d6d93304043e7d091713f6b427d95f336ebeb7aafda37bd0573e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ie8fac8bea605bb51b108d5398ca815fe8b8db776d57ea99f33548dde8fb09abc.ExtensionsRequestBuilder) {
    return ie8fac8bea605bb51b108d5398ca815fe8b8db776d57ea99f33548dde8fb09abc.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i296762b0003986a4bc9f5bb4934735dbd6c510000fe4111f0c99c7ca58893ee3.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i296762b0003986a4bc9f5bb4934735dbd6c510000fe4111f0c99c7ca58893ee3.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i1073697e475927811aad904547f6513fa5f74982d0d3ff0e75282b340778206f.ForwardRequestBuilder) {
    return i1073697e475927811aad904547f6513fa5f74982d0d3ff0e75282b340778206f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*i0c7906a46cd3570c58378ded74411de15227e2d779e5655aebe8225e8af872e6.InstancesRequestBuilder) {
    return i0c7906a46cd3570c58378ded74411de15227e2d779e5655aebe8225e8af872e6.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.exceptionOccurrences.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i40ad88377bdcea87a9bf81f4f20ac56e52c135f41bcf17f343b2f17f8c17ed2d.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i40ad88377bdcea87a9bf81f4f20ac56e52c135f41bcf17f343b2f17f8c17ed2d.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i28d4b967a0a453124dd6bd0e0197bbaa48accac70c41fa2dbb27689e69adfb9e.MultiValueExtendedPropertiesRequestBuilder) {
    return i28d4b967a0a453124dd6bd0e0197bbaa48accac70c41fa2dbb27689e69adfb9e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i421482755bd37443917e084a854e179974c6a4fdc954958616023e9509133dc5.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i421482755bd37443917e084a854e179974c6a4fdc954958616023e9509133dc5.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ibda2764d6365b6907754db6ca7f56ef5fc34fe39ecfcbe8b94801a1987537b3b.SingleValueExtendedPropertiesRequestBuilder) {
    return ibda2764d6365b6907754db6ca7f56ef5fc34fe39ecfcbe8b94801a1987537b3b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendarView.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i81024da70f9ee34f7d5b4be69d2ccdd5628df29fdbc0e2ceef48cb4962b2ba50.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i81024da70f9ee34f7d5b4be69d2ccdd5628df29fdbc0e2ceef48cb4962b2ba50.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i6715da1957ca86f82d8b761208b8373f87296676acffecc74f236992ddd7ff4a.SnoozeReminderRequestBuilder) {
    return i6715da1957ca86f82d8b761208b8373f87296676acffecc74f236992ddd7ff4a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i29e043eaee31491b11c98cf95ab9a206543eb504e91e849ce5dc75c11db368d1.TentativelyAcceptRequestBuilder) {
    return i29e043eaee31491b11c98cf95ab9a206543eb504e91e849ce5dc75c11db368d1.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
