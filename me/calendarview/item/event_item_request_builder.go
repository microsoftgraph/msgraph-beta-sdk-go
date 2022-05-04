package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3bd895fd978dad1672adc3f0311f66fdb03005344906384f3ab769063e50829f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/extensions"
    i3c3a9a8e0b281a1eed928493de2e7d54cad87b5cc7ce112c860a09f13fffdb4f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/multivalueextendedproperties"
    i5072dcc71f81fbbc8e1cfabf9d699e903478b616baee2be98a65da249f1fa954 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances"
    i5f5f630d5ced06d8366f7b861f39f535ba283864a9c0b987a235807d16bcd2d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/attachments"
    i62212df5f59e6343c09b1b76ea98ab81a010543533fdc29e488c313f56875602 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/accept"
    i7d64b836ebe68769a340fbfe80e29dc1e3d2d62c42cf1e7f6994fcbbc89fcb34 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/dismissreminder"
    i87f4103909d29d7e9abba659ab2e44d5ca7a48aaa0ab7a174cdffe70e67072f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences"
    i8fd84426a95c0c0cb3b5115bd277b9e826eab1ca15cbb4d40df967d7fd862d63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/tentativelyaccept"
    ia46e16175ce133a4275988f6c7f673681cbd68546351722f3943a53f9f9cdcf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/snoozereminder"
    ic00b975abeb3690b8a537e2dff81241198d1561dbf386671dc806f3228917476 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/singlevalueextendedproperties"
    iddaddb07fe634e5727af12a88fa5d19867a17bf51a0ca2e7db6fc889d2dd561b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/cancel"
    ie9f89f59a45fb8f2559548a58c139cde22ddbbdebe5d326c2d897ca32fb418e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/decline"
    iecc8853e211513ba4b57daba10b08033aa72795dc491f48ddc834c2b85b35b95 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/forward"
    if0ee07e429daf61efbe896fc2c452f25eb4fb0f148f699739371a571c28387af "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar"
    i4892f6b0282f166f307584c158a0b18f350bd5e674fd73a2f4a68c01793b5ddd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/singlevalueextendedproperties/item"
    i8b61b6c9f100a3ca9dcf838c2b4ff31850b4d63e0a3898010d0740c45ec3fe53 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item"
    ic542911019ed2e6fb318ff87489a78092fc04495dfbe0c06e9f2f797b3f1457a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/extensions/item"
    ief7c421ee9694c46920200caadcf67aba8d4d0106872a7a7850705606effdc0f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/multivalueextendedproperties/item"
    if2a79ac6e19e9a10ee61e59ae32c00ed73834be0f7c7760ea2432801f5af3a7a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/attachments/item"
    if55b6d3ec50fcb8d34741fdcda5010c79e181eeee3537556b0cfaaa29dcef6cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*i62212df5f59e6343c09b1b76ea98ab81a010543533fdc29e488c313f56875602.AcceptRequestBuilder) {
    return i62212df5f59e6343c09b1b76ea98ab81a010543533fdc29e488c313f56875602.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i5f5f630d5ced06d8366f7b861f39f535ba283864a9c0b987a235807d16bcd2d9.AttachmentsRequestBuilder) {
    return i5f5f630d5ced06d8366f7b861f39f535ba283864a9c0b987a235807d16bcd2d9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*if2a79ac6e19e9a10ee61e59ae32c00ed73834be0f7c7760ea2432801f5af3a7a.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return if2a79ac6e19e9a10ee61e59ae32c00ed73834be0f7c7760ea2432801f5af3a7a.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*if0ee07e429daf61efbe896fc2c452f25eb4fb0f148f699739371a571c28387af.CalendarRequestBuilder) {
    return if0ee07e429daf61efbe896fc2c452f25eb4fb0f148f699739371a571c28387af.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*iddaddb07fe634e5727af12a88fa5d19867a17bf51a0ca2e7db6fc889d2dd561b.CancelRequestBuilder) {
    return iddaddb07fe634e5727af12a88fa5d19867a17bf51a0ca2e7db6fc889d2dd561b.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event%2Did}{?startDateTime,endDateTime,%24select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calendarView for me
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
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property calendarView in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calendarView in me
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
func (m *EventItemRequestBuilder) Decline()(*ie9f89f59a45fb8f2559548a58c139cde22ddbbdebe5d326c2d897ca32fb418e6.DeclineRequestBuilder) {
    return ie9f89f59a45fb8f2559548a58c139cde22ddbbdebe5d326c2d897ca32fb418e6.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property calendarView for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i7d64b836ebe68769a340fbfe80e29dc1e3d2d62c42cf1e7f6994fcbbc89fcb34.DismissReminderRequestBuilder) {
    return i7d64b836ebe68769a340fbfe80e29dc1e3d2d62c42cf1e7f6994fcbbc89fcb34.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i87f4103909d29d7e9abba659ab2e44d5ca7a48aaa0ab7a174cdffe70e67072f5.ExceptionOccurrencesRequestBuilder) {
    return i87f4103909d29d7e9abba659ab2e44d5ca7a48aaa0ab7a174cdffe70e67072f5.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*if55b6d3ec50fcb8d34741fdcda5010c79e181eeee3537556b0cfaaa29dcef6cb.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return if55b6d3ec50fcb8d34741fdcda5010c79e181eeee3537556b0cfaaa29dcef6cb.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i3bd895fd978dad1672adc3f0311f66fdb03005344906384f3ab769063e50829f.ExtensionsRequestBuilder) {
    return i3bd895fd978dad1672adc3f0311f66fdb03005344906384f3ab769063e50829f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic542911019ed2e6fb318ff87489a78092fc04495dfbe0c06e9f2f797b3f1457a.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ic542911019ed2e6fb318ff87489a78092fc04495dfbe0c06e9f2f797b3f1457a.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*iecc8853e211513ba4b57daba10b08033aa72795dc491f48ddc834c2b85b35b95.ForwardRequestBuilder) {
    return iecc8853e211513ba4b57daba10b08033aa72795dc491f48ddc834c2b85b35b95.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*i5072dcc71f81fbbc8e1cfabf9d699e903478b616baee2be98a65da249f1fa954.InstancesRequestBuilder) {
    return i5072dcc71f81fbbc8e1cfabf9d699e903478b616baee2be98a65da249f1fa954.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i8b61b6c9f100a3ca9dcf838c2b4ff31850b4d63e0a3898010d0740c45ec3fe53.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i8b61b6c9f100a3ca9dcf838c2b4ff31850b4d63e0a3898010d0740c45ec3fe53.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i3c3a9a8e0b281a1eed928493de2e7d54cad87b5cc7ce112c860a09f13fffdb4f.MultiValueExtendedPropertiesRequestBuilder) {
    return i3c3a9a8e0b281a1eed928493de2e7d54cad87b5cc7ce112c860a09f13fffdb4f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ief7c421ee9694c46920200caadcf67aba8d4d0106872a7a7850705606effdc0f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ief7c421ee9694c46920200caadcf67aba8d4d0106872a7a7850705606effdc0f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property calendarView in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ic00b975abeb3690b8a537e2dff81241198d1561dbf386671dc806f3228917476.SingleValueExtendedPropertiesRequestBuilder) {
    return ic00b975abeb3690b8a537e2dff81241198d1561dbf386671dc806f3228917476.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4892f6b0282f166f307584c158a0b18f350bd5e674fd73a2f4a68c01793b5ddd.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i4892f6b0282f166f307584c158a0b18f350bd5e674fd73a2f4a68c01793b5ddd.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ia46e16175ce133a4275988f6c7f673681cbd68546351722f3943a53f9f9cdcf4.SnoozeReminderRequestBuilder) {
    return ia46e16175ce133a4275988f6c7f673681cbd68546351722f3943a53f9f9cdcf4.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i8fd84426a95c0c0cb3b5115bd277b9e826eab1ca15cbb4d40df967d7fd862d63.TentativelyAcceptRequestBuilder) {
    return i8fd84426a95c0c0cb3b5115bd277b9e826eab1ca15cbb4d40df967d7fd862d63.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
