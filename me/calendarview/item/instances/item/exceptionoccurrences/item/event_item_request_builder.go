package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i152f667946903f3142a179ce6b8207b72be5d83dd570c317088c5f68e6426de1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i19008bdd393ac510836f89abdfc5961bbe94940d8ca6650e1dd92f85abdf1a89 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i1db01cd34af40fffde05b40f2ebd8f98ca26f2ed39afa357f54e38851ef29ddf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i1e5f1ccd25bca6841a67513eaf26575240a16533ebf30d66da8c1f9ab6317246 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    i20e9cdb6b9ecc1ab3d960602c9613673912c729196ffb07bd83dd5ddf4722522 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i21cb246107a672cced00b657124cc184c3bc3468f1b53f7ee8a6afb99469be41 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    i5b08890a5c58b7714fc9fef5bb602698db543a14ae9d904f36efa40e9ec4cca5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i762fedaf6d516d048dc4cbb2c5bc4b5fb0b695ee49d85a441fa6ebbdc22996ff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ia35c643ab7ea66cb5904e1e8413b9d9750d134553af0285e58ada35bf037c966 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    ie536ed9529f4fe16b02d0801081385eb18aea8bbefc370965aa9b15dc9900045 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ied2d4e78ae16ad635fa91183d46cfab11b73a2d44d9cdd29a311e5528614909d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    ife7a07962e2bb49d0cdb81333af66776561905ba176340ecc5a9b2b49a09e8dc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i012601c386deb427e2ad383144b42a711462b3c5bb87560521efacd757400c56 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i266648a9362e43d946c4da032ba11dd1968a12c53707315c92a23dca9b936a7f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i7e26973f0110103fc737f73342f2351a2be0efeff7b2f04e2e2c876d76e47478 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    icaa62508cd21e490f961e360752b4fa560dea6978c255091494ded05f6ab9d6b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*i21cb246107a672cced00b657124cc184c3bc3468f1b53f7ee8a6afb99469be41.AcceptRequestBuilder) {
    return i21cb246107a672cced00b657124cc184c3bc3468f1b53f7ee8a6afb99469be41.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ia35c643ab7ea66cb5904e1e8413b9d9750d134553af0285e58ada35bf037c966.AttachmentsRequestBuilder) {
    return ia35c643ab7ea66cb5904e1e8413b9d9750d134553af0285e58ada35bf037c966.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*icaa62508cd21e490f961e360752b4fa560dea6978c255091494ded05f6ab9d6b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return icaa62508cd21e490f961e360752b4fa560dea6978c255091494ded05f6ab9d6b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ied2d4e78ae16ad635fa91183d46cfab11b73a2d44d9cdd29a311e5528614909d.CalendarRequestBuilder) {
    return ied2d4e78ae16ad635fa91183d46cfab11b73a2d44d9cdd29a311e5528614909d.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i19008bdd393ac510836f89abdfc5961bbe94940d8ca6650e1dd92f85abdf1a89.CancelRequestBuilder) {
    return i19008bdd393ac510836f89abdfc5961bbe94940d8ca6650e1dd92f85abdf1a89.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i1e5f1ccd25bca6841a67513eaf26575240a16533ebf30d66da8c1f9ab6317246.DeclineRequestBuilder) {
    return i1e5f1ccd25bca6841a67513eaf26575240a16533ebf30d66da8c1f9ab6317246.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i152f667946903f3142a179ce6b8207b72be5d83dd570c317088c5f68e6426de1.DismissReminderRequestBuilder) {
    return i152f667946903f3142a179ce6b8207b72be5d83dd570c317088c5f68e6426de1.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i20e9cdb6b9ecc1ab3d960602c9613673912c729196ffb07bd83dd5ddf4722522.ExtensionsRequestBuilder) {
    return i20e9cdb6b9ecc1ab3d960602c9613673912c729196ffb07bd83dd5ddf4722522.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i7e26973f0110103fc737f73342f2351a2be0efeff7b2f04e2e2c876d76e47478.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i7e26973f0110103fc737f73342f2351a2be0efeff7b2f04e2e2c876d76e47478.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ife7a07962e2bb49d0cdb81333af66776561905ba176340ecc5a9b2b49a09e8dc.ForwardRequestBuilder) {
    return ife7a07962e2bb49d0cdb81333af66776561905ba176340ecc5a9b2b49a09e8dc.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i5b08890a5c58b7714fc9fef5bb602698db543a14ae9d904f36efa40e9ec4cca5.MultiValueExtendedPropertiesRequestBuilder) {
    return i5b08890a5c58b7714fc9fef5bb602698db543a14ae9d904f36efa40e9ec4cca5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i266648a9362e43d946c4da032ba11dd1968a12c53707315c92a23dca9b936a7f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i266648a9362e43d946c4da032ba11dd1968a12c53707315c92a23dca9b936a7f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i762fedaf6d516d048dc4cbb2c5bc4b5fb0b695ee49d85a441fa6ebbdc22996ff.SingleValueExtendedPropertiesRequestBuilder) {
    return i762fedaf6d516d048dc4cbb2c5bc4b5fb0b695ee49d85a441fa6ebbdc22996ff.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i012601c386deb427e2ad383144b42a711462b3c5bb87560521efacd757400c56.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i012601c386deb427e2ad383144b42a711462b3c5bb87560521efacd757400c56.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i1db01cd34af40fffde05b40f2ebd8f98ca26f2ed39afa357f54e38851ef29ddf.SnoozeReminderRequestBuilder) {
    return i1db01cd34af40fffde05b40f2ebd8f98ca26f2ed39afa357f54e38851ef29ddf.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ie536ed9529f4fe16b02d0801081385eb18aea8bbefc370965aa9b15dc9900045.TentativelyAcceptRequestBuilder) {
    return ie536ed9529f4fe16b02d0801081385eb18aea8bbefc370965aa9b15dc9900045.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
