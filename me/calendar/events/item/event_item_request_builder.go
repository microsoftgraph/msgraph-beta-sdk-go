package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i07fe24d1dec166fbef03b00ab5b46e976baffb6c06b5ef5e6347d446ff25bca3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/extensions"
    i5dbf769c1f25e1b0b17d30509a4a645cfa242a1c4d7c348c060dccc53cf78dd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/multivalueextendedproperties"
    i7cda4870e61701a73f65b06cb1c454e1fa2c00f69cd309e8d501fd4a70e14a22 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/dismissreminder"
    i86ef866cc2d984e3066c38be80d8b147c286544fa4fecd6ab59fe1d8ff248c1d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/decline"
    i8f07c66b2cf1a3f6d8da9f8cf5d8442d87cb10892aad82c5dd44556cc8ec85bb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/forward"
    i97b528d359cdfaf66013514b9af5ce682328c510bc03248fb0810ebf226f19e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/attachments"
    ib2a51806722e21c6ec788fdf6b0eee10e907a4dc8e354b7d914a32f4368d0864 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/accept"
    ibd61d302d825b6e73918b5027ff3b14b0b17c3c96a153d5cffabf8cf98cbd736 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/tentativelyaccept"
    icae7b5ff991f4631b46cc3802546c55d85ad10eaeb0238bf2f6e3ff7bf554f0e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/cancel"
    idf6baae4d48e4fcb9e8ccf745a9b91eb8ec7f7f41d6f2c4074fca23712f7aa70 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/calendar"
    ie4cb09c95e0c35a55010b7327a11926471dad8ff28619f453f661c22aee3f94a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/singlevalueextendedproperties"
    ie506b5a0873b4dc903852664779e9434b8ef3c1ada875aa7f0551abf90de2d3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/snoozereminder"
    if0032e7643530222733c0d1a82c01d7975494927aa723dfe9e7505cb06dc839e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences"
    if607cd43935ec64edb92794109173f3c0d5910d8dfebec67b6e22ea9bffd7433 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances"
    i12298c30cb81e9f690ef9152f8caab34fa4d1864ebd36ed5577020082370bd73 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/extensions/item"
    i5956806f1c4b998abe4fcd54db1586d7465f29b5e18161bba26e275d033eab11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/instances/item"
    i73e700365b2c2d1e86684192b02ce3e5f7775ba91bb65b0fe4eb04ec97e7b592 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/exceptionoccurrences/item"
    i916f3cc62a9632691b92c06ae4a196bfbbab686ba1d7e0ec5f5b48c10806fd6b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/singlevalueextendedproperties/item"
    ia14123345e5038209c9a2a037b69700ff79c35780a03e5454afed921703f8320 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/attachments/item"
    if2b1e2a9807085574644316d2e02a3a9aa2335ecc40c7e9a8828e945fbc62be9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/events/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*ib2a51806722e21c6ec788fdf6b0eee10e907a4dc8e354b7d914a32f4368d0864.AcceptRequestBuilder) {
    return ib2a51806722e21c6ec788fdf6b0eee10e907a4dc8e354b7d914a32f4368d0864.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i97b528d359cdfaf66013514b9af5ce682328c510bc03248fb0810ebf226f19e4.AttachmentsRequestBuilder) {
    return i97b528d359cdfaf66013514b9af5ce682328c510bc03248fb0810ebf226f19e4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ia14123345e5038209c9a2a037b69700ff79c35780a03e5454afed921703f8320.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ia14123345e5038209c9a2a037b69700ff79c35780a03e5454afed921703f8320.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*idf6baae4d48e4fcb9e8ccf745a9b91eb8ec7f7f41d6f2c4074fca23712f7aa70.CalendarRequestBuilder) {
    return idf6baae4d48e4fcb9e8ccf745a9b91eb8ec7f7f41d6f2c4074fca23712f7aa70.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*icae7b5ff991f4631b46cc3802546c55d85ad10eaeb0238bf2f6e3ff7bf554f0e.CancelRequestBuilder) {
    return icae7b5ff991f4631b46cc3802546c55d85ad10eaeb0238bf2f6e3ff7bf554f0e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/events/{event%2Did}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i86ef866cc2d984e3066c38be80d8b147c286544fa4fecd6ab59fe1d8ff248c1d.DeclineRequestBuilder) {
    return i86ef866cc2d984e3066c38be80d8b147c286544fa4fecd6ab59fe1d8ff248c1d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i7cda4870e61701a73f65b06cb1c454e1fa2c00f69cd309e8d501fd4a70e14a22.DismissReminderRequestBuilder) {
    return i7cda4870e61701a73f65b06cb1c454e1fa2c00f69cd309e8d501fd4a70e14a22.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*if0032e7643530222733c0d1a82c01d7975494927aa723dfe9e7505cb06dc839e.ExceptionOccurrencesRequestBuilder) {
    return if0032e7643530222733c0d1a82c01d7975494927aa723dfe9e7505cb06dc839e.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i73e700365b2c2d1e86684192b02ce3e5f7775ba91bb65b0fe4eb04ec97e7b592.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i73e700365b2c2d1e86684192b02ce3e5f7775ba91bb65b0fe4eb04ec97e7b592.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i07fe24d1dec166fbef03b00ab5b46e976baffb6c06b5ef5e6347d446ff25bca3.ExtensionsRequestBuilder) {
    return i07fe24d1dec166fbef03b00ab5b46e976baffb6c06b5ef5e6347d446ff25bca3.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i12298c30cb81e9f690ef9152f8caab34fa4d1864ebd36ed5577020082370bd73.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i12298c30cb81e9f690ef9152f8caab34fa4d1864ebd36ed5577020082370bd73.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i8f07c66b2cf1a3f6d8da9f8cf5d8442d87cb10892aad82c5dd44556cc8ec85bb.ForwardRequestBuilder) {
    return i8f07c66b2cf1a3f6d8da9f8cf5d8442d87cb10892aad82c5dd44556cc8ec85bb.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*if607cd43935ec64edb92794109173f3c0d5910d8dfebec67b6e22ea9bffd7433.InstancesRequestBuilder) {
    return if607cd43935ec64edb92794109173f3c0d5910d8dfebec67b6e22ea9bffd7433.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i5956806f1c4b998abe4fcd54db1586d7465f29b5e18161bba26e275d033eab11.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did1"] = id
    }
    return i5956806f1c4b998abe4fcd54db1586d7465f29b5e18161bba26e275d033eab11.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i5dbf769c1f25e1b0b17d30509a4a645cfa242a1c4d7c348c060dccc53cf78dd6.MultiValueExtendedPropertiesRequestBuilder) {
    return i5dbf769c1f25e1b0b17d30509a4a645cfa242a1c4d7c348c060dccc53cf78dd6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if2b1e2a9807085574644316d2e02a3a9aa2335ecc40c7e9a8828e945fbc62be9.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return if2b1e2a9807085574644316d2e02a3a9aa2335ecc40c7e9a8828e945fbc62be9.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ie4cb09c95e0c35a55010b7327a11926471dad8ff28619f453f661c22aee3f94a.SingleValueExtendedPropertiesRequestBuilder) {
    return ie4cb09c95e0c35a55010b7327a11926471dad8ff28619f453f661c22aee3f94a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.events.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i916f3cc62a9632691b92c06ae4a196bfbbab686ba1d7e0ec5f5b48c10806fd6b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i916f3cc62a9632691b92c06ae4a196bfbbab686ba1d7e0ec5f5b48c10806fd6b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ie506b5a0873b4dc903852664779e9434b8ef3c1ada875aa7f0551abf90de2d3a.SnoozeReminderRequestBuilder) {
    return ie506b5a0873b4dc903852664779e9434b8ef3c1ada875aa7f0551abf90de2d3a.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ibd61d302d825b6e73918b5027ff3b14b0b17c3c96a153d5cffabf8cf98cbd736.TentativelyAcceptRequestBuilder) {
    return ibd61d302d825b6e73918b5027ff3b14b0b17c3c96a153d5cffabf8cf98cbd736.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
