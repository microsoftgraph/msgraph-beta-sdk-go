package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01fb2df2892ac0830c87fcf62df419a5b957e82a94a89a349b90c39f6326679f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i13c2c43066b385d189c402f1f3f02386c685d77e6f9642b8d003c4bd88889914 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    i15bbfef361fe7f43c4a0cd40a59b5ec6588061e6a4197451c9519398b88a3cfa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i699849b0e1615887fb1c6d4922bcb1386bae6fd4b85c8e86f0c9da8ab8202e40 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    i6feed7c975422a5f9086f766d03ab0a24ed1b17fba89878fafb57aabe69c93dd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i72df339a7df0d82419938509403b21bf282ed3744852880a72066041563e6758 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i799529df9f53556c8af69e2520763915522836d6aafef34252ada69b68d59e9c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    ia1da8b1ea95847818016a9356e6270d909e6d418a0dfd0e51e8de07b93f58774 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ic1ad4ca23474b81e1226dfd7992a27af68a71109111a888e15e745ba61403c56 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    icc15b479516c13d95b6f684dfe0766b6706c9d7d143ab550d41c2ec0a10dfd63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    ide723326b393119009a6f0c17c949555227f138f4d770bd2a2b909f1d33c93df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    iebfd6610409ba87469bf28d1f2ca772c68e234f4525294e6e0ef11f925f69988 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i2c01bdb15aebb4c846849d6e815ae936248297bb97e58ed479654e0c45204217 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
    i72e02a1bf8eaa5d8a8bbb72de7ef1af36bf0bcdb78c896a4043e45a41db53468 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    ic7bcb62e34588a967c7d97ac3c1f396f8903ff050ce91b506399f6b2a748bec6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    ie6d99adf383d06bf9aead85124fcf4eccb5db52694831a69544b844c86e96fea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*ide723326b393119009a6f0c17c949555227f138f4d770bd2a2b909f1d33c93df.AcceptRequestBuilder) {
    return ide723326b393119009a6f0c17c949555227f138f4d770bd2a2b909f1d33c93df.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i13c2c43066b385d189c402f1f3f02386c685d77e6f9642b8d003c4bd88889914.AttachmentsRequestBuilder) {
    return i13c2c43066b385d189c402f1f3f02386c685d77e6f9642b8d003c4bd88889914.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i2c01bdb15aebb4c846849d6e815ae936248297bb97e58ed479654e0c45204217.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i2c01bdb15aebb4c846849d6e815ae936248297bb97e58ed479654e0c45204217.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i699849b0e1615887fb1c6d4922bcb1386bae6fd4b85c8e86f0c9da8ab8202e40.CalendarRequestBuilder) {
    return i699849b0e1615887fb1c6d4922bcb1386bae6fd4b85c8e86f0c9da8ab8202e40.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ic1ad4ca23474b81e1226dfd7992a27af68a71109111a888e15e745ba61403c56.CancelRequestBuilder) {
    return ic1ad4ca23474b81e1226dfd7992a27af68a71109111a888e15e745ba61403c56.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24select,%24expand}";
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
func (m *EventItemRequestBuilder) Decline()(*i799529df9f53556c8af69e2520763915522836d6aafef34252ada69b68d59e9c.DeclineRequestBuilder) {
    return i799529df9f53556c8af69e2520763915522836d6aafef34252ada69b68d59e9c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i15bbfef361fe7f43c4a0cd40a59b5ec6588061e6a4197451c9519398b88a3cfa.DismissReminderRequestBuilder) {
    return i15bbfef361fe7f43c4a0cd40a59b5ec6588061e6a4197451c9519398b88a3cfa.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*icc15b479516c13d95b6f684dfe0766b6706c9d7d143ab550d41c2ec0a10dfd63.ExtensionsRequestBuilder) {
    return icc15b479516c13d95b6f684dfe0766b6706c9d7d143ab550d41c2ec0a10dfd63.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic7bcb62e34588a967c7d97ac3c1f396f8903ff050ce91b506399f6b2a748bec6.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ic7bcb62e34588a967c7d97ac3c1f396f8903ff050ce91b506399f6b2a748bec6.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i01fb2df2892ac0830c87fcf62df419a5b957e82a94a89a349b90c39f6326679f.ForwardRequestBuilder) {
    return i01fb2df2892ac0830c87fcf62df419a5b957e82a94a89a349b90c39f6326679f.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*iebfd6610409ba87469bf28d1f2ca772c68e234f4525294e6e0ef11f925f69988.MultiValueExtendedPropertiesRequestBuilder) {
    return iebfd6610409ba87469bf28d1f2ca772c68e234f4525294e6e0ef11f925f69988.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i72e02a1bf8eaa5d8a8bbb72de7ef1af36bf0bcdb78c896a4043e45a41db53468.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i72e02a1bf8eaa5d8a8bbb72de7ef1af36bf0bcdb78c896a4043e45a41db53468.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ia1da8b1ea95847818016a9356e6270d909e6d418a0dfd0e51e8de07b93f58774.SingleValueExtendedPropertiesRequestBuilder) {
    return ia1da8b1ea95847818016a9356e6270d909e6d418a0dfd0e51e8de07b93f58774.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie6d99adf383d06bf9aead85124fcf4eccb5db52694831a69544b844c86e96fea.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ie6d99adf383d06bf9aead85124fcf4eccb5db52694831a69544b844c86e96fea.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i72df339a7df0d82419938509403b21bf282ed3744852880a72066041563e6758.SnoozeReminderRequestBuilder) {
    return i72df339a7df0d82419938509403b21bf282ed3744852880a72066041563e6758.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i6feed7c975422a5f9086f766d03ab0a24ed1b17fba89878fafb57aabe69c93dd.TentativelyAcceptRequestBuilder) {
    return i6feed7c975422a5f9086f766d03ab0a24ed1b17fba89878fafb57aabe69c93dd.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
