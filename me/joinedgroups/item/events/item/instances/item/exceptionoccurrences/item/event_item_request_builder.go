package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0637b1b044557e7b5b8c3e6c70dc5733df7fa76bbe90427bee500ee6652a261c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/extensions"
    i2ef2359d0033ea337e3295b53c19a6b626e831b1d55a60c06d6c1576066752b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/accept"
    i459ddc92b87b18fdf71dea8264df298f148def9536419f7defa96ce3f4d510c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i4ef2b31c0d6b0ed8bb9f80f3c5a2d2bdb1d71482334c06c9f97b3dc5b71edd38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/decline"
    i5943fab14ec1838ed09357994199fd9bbce376c0383798fc91e0e2b1e04a69d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i75669a6e22f23d1197683cfb0ae52280bc1324f8bffe50cf6d1464a129f4a8f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i882f2d94c002fd30f9e57f18d6a42068a659d7f09a923fed7d2fc8095460df1e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/cancel"
    i8f26076830125efabc65aefc231eaf67039146d3935c685cfcd50644da7c7171 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/attachments"
    i90f3a2fc1b0953e92717035be173ee1a14288390e2c2a5d5ab36cd6e48dfbb33 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    ib0b98b629b75b0010847eb99d1cfc91590e570896327aea6ac849c35ad9b165d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/calendar"
    ib58ddc4a294b2c6b12dcc6819a0bd17f5f2ec9bff60aec34f1ad6d7e1a856fac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ifea97e5ce93ccfdcebd2292996e99243535133c1c3d40cd5656b31315d1d7299 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/forward"
    i342dcfa63a14c6c5f779c4c2365feea1b5fea0f6fba4eb0b58411ce20f22703d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i5fc6bffc427c34882bb24d373a67d7f26b9f4b040deb072b95f41eb731ac1c86 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    i6c201a558cd1290a8aecd3397896e869806213ae412e4e35865176d8067b34f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    i7888e02bb29e7b80d9e6d14b80a01edb9216d845ce305c09d03fc75643d60ab3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i2ef2359d0033ea337e3295b53c19a6b626e831b1d55a60c06d6c1576066752b9.AcceptRequestBuilder) {
    return i2ef2359d0033ea337e3295b53c19a6b626e831b1d55a60c06d6c1576066752b9.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i8f26076830125efabc65aefc231eaf67039146d3935c685cfcd50644da7c7171.AttachmentsRequestBuilder) {
    return i8f26076830125efabc65aefc231eaf67039146d3935c685cfcd50644da7c7171.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i6c201a558cd1290a8aecd3397896e869806213ae412e4e35865176d8067b34f9.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i6c201a558cd1290a8aecd3397896e869806213ae412e4e35865176d8067b34f9.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ib0b98b629b75b0010847eb99d1cfc91590e570896327aea6ac849c35ad9b165d.CalendarRequestBuilder) {
    return ib0b98b629b75b0010847eb99d1cfc91590e570896327aea6ac849c35ad9b165d.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i882f2d94c002fd30f9e57f18d6a42068a659d7f09a923fed7d2fc8095460df1e.CancelRequestBuilder) {
    return i882f2d94c002fd30f9e57f18d6a42068a659d7f09a923fed7d2fc8095460df1e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i4ef2b31c0d6b0ed8bb9f80f3c5a2d2bdb1d71482334c06c9f97b3dc5b71edd38.DeclineRequestBuilder) {
    return i4ef2b31c0d6b0ed8bb9f80f3c5a2d2bdb1d71482334c06c9f97b3dc5b71edd38.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i75669a6e22f23d1197683cfb0ae52280bc1324f8bffe50cf6d1464a129f4a8f1.DismissReminderRequestBuilder) {
    return i75669a6e22f23d1197683cfb0ae52280bc1324f8bffe50cf6d1464a129f4a8f1.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i0637b1b044557e7b5b8c3e6c70dc5733df7fa76bbe90427bee500ee6652a261c.ExtensionsRequestBuilder) {
    return i0637b1b044557e7b5b8c3e6c70dc5733df7fa76bbe90427bee500ee6652a261c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i5fc6bffc427c34882bb24d373a67d7f26b9f4b040deb072b95f41eb731ac1c86.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i5fc6bffc427c34882bb24d373a67d7f26b9f4b040deb072b95f41eb731ac1c86.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ifea97e5ce93ccfdcebd2292996e99243535133c1c3d40cd5656b31315d1d7299.ForwardRequestBuilder) {
    return ifea97e5ce93ccfdcebd2292996e99243535133c1c3d40cd5656b31315d1d7299.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i90f3a2fc1b0953e92717035be173ee1a14288390e2c2a5d5ab36cd6e48dfbb33.MultiValueExtendedPropertiesRequestBuilder) {
    return i90f3a2fc1b0953e92717035be173ee1a14288390e2c2a5d5ab36cd6e48dfbb33.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7888e02bb29e7b80d9e6d14b80a01edb9216d845ce305c09d03fc75643d60ab3.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i7888e02bb29e7b80d9e6d14b80a01edb9216d845ce305c09d03fc75643d60ab3.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ib58ddc4a294b2c6b12dcc6819a0bd17f5f2ec9bff60aec34f1ad6d7e1a856fac.SingleValueExtendedPropertiesRequestBuilder) {
    return ib58ddc4a294b2c6b12dcc6819a0bd17f5f2ec9bff60aec34f1ad6d7e1a856fac.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i342dcfa63a14c6c5f779c4c2365feea1b5fea0f6fba4eb0b58411ce20f22703d.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i342dcfa63a14c6c5f779c4c2365feea1b5fea0f6fba4eb0b58411ce20f22703d.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i5943fab14ec1838ed09357994199fd9bbce376c0383798fc91e0e2b1e04a69d5.SnoozeReminderRequestBuilder) {
    return i5943fab14ec1838ed09357994199fd9bbce376c0383798fc91e0e2b1e04a69d5.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i459ddc92b87b18fdf71dea8264df298f148def9536419f7defa96ce3f4d510c4.TentativelyAcceptRequestBuilder) {
    return i459ddc92b87b18fdf71dea8264df298f148def9536419f7defa96ce3f4d510c4.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
