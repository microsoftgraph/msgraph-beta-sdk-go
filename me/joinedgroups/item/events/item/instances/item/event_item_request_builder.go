package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i017a715501a5ab4f66e77ae5b4bd5374f4e8b5b4ff4303c4f1760e18ce649a67 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/multivalueextendedproperties"
    i03c8d4da7643b9e7219f7467bfa0d76efd6f53f5dcbcd3bc55bde0e2b781b6bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/singlevalueextendedproperties"
    i1a557d4fc16aa13d0108453d1c17e66dbd07e63c05715e5f9d35051e88a24a24 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/calendar"
    i3fd0ef49c857e171f8937f421b54c5927f498ab413bf6825cebf23a7e9d0ea01 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/accept"
    i4d57b8aaed5da088236afa277a95d913be0366ead59755ed6a66f88667b89356 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/tentativelyaccept"
    i667ed9dad4d0c7e77f1a855f41d77968445096cfee0fe4e2d47e140b8d90b8c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences"
    i8e9ecd6e780ee4bb2c786ff4a010f0c3f58b8ad11c339198ee764edbe774e97f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/snoozereminder"
    i9da0074ea8a34e5df0ba5be77bbfaa546a3789dc60a880171c2e4829efb3ecbe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/extensions"
    i9e9edad3e83bc93a3623a987af5f56542c1c84e452cc45b229c9a594be8d62b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/attachments"
    iae3ac1c45f912bc7620b50de6f38592a1ecf12075e336c96f755cdb5b6b71b06 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/forward"
    ibe0f523247e72d807de19887754eb0e311dfa79b099a4b762d9b4dfe259edd95 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/cancel"
    iccc659f8b7944aa988b5bfd93abbb93d536a4c19d69893dabf8d7d4860528c87 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/decline"
    id40d08ce5fefbde3ff290f0251f54e2d66b05399d8cdf74aa00ca24dc8dbb8e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/dismissreminder"
    i57fba512fdd9c018cea1ea18a1c19859d0dc72eb1271b7ef534e62a89f822751 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/extensions/item"
    i5c96f88b3ec2612713c9d12f6793cc22baf5bf2de02b2a1fbfa9092f9640d7d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/exceptionoccurrences/item"
    i6ab0f05b73e2398f67634938632afbb24fbf7929e1a0075ebc5f4309841a5683 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/attachments/item"
    ic860d18bcf9afeee3897ed4e52f55e614afd045fcf573e65d9fd194225b99773 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/multivalueextendedproperties/item"
    ie98d034aecbf9bae9b33a1db86754c1684a20b4959b168ca7abaf8b9100b7cea "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/events/item/instances/item/singlevalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
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
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) Accept()(*i3fd0ef49c857e171f8937f421b54c5927f498ab413bf6825cebf23a7e9d0ea01.AcceptRequestBuilder) {
    return i3fd0ef49c857e171f8937f421b54c5927f498ab413bf6825cebf23a7e9d0ea01.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i9e9edad3e83bc93a3623a987af5f56542c1c84e452cc45b229c9a594be8d62b5.AttachmentsRequestBuilder) {
    return i9e9edad3e83bc93a3623a987af5f56542c1c84e452cc45b229c9a594be8d62b5.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i6ab0f05b73e2398f67634938632afbb24fbf7929e1a0075ebc5f4309841a5683.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i6ab0f05b73e2398f67634938632afbb24fbf7929e1a0075ebc5f4309841a5683.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i1a557d4fc16aa13d0108453d1c17e66dbd07e63c05715e5f9d35051e88a24a24.CalendarRequestBuilder) {
    return i1a557d4fc16aa13d0108453d1c17e66dbd07e63c05715e5f9d35051e88a24a24.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ibe0f523247e72d807de19887754eb0e311dfa79b099a4b762d9b4dfe259edd95.CancelRequestBuilder) {
    return ibe0f523247e72d807de19887754eb0e311dfa79b099a4b762d9b4dfe259edd95.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property instances for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instances for me
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
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property instances in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instances in me
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
func (m *EventItemRequestBuilder) Decline()(*iccc659f8b7944aa988b5bfd93abbb93d536a4c19d69893dabf8d7d4860528c87.DeclineRequestBuilder) {
    return iccc659f8b7944aa988b5bfd93abbb93d536a4c19d69893dabf8d7d4860528c87.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for me
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property instances for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*id40d08ce5fefbde3ff290f0251f54e2d66b05399d8cdf74aa00ca24dc8dbb8e0.DismissReminderRequestBuilder) {
    return id40d08ce5fefbde3ff290f0251f54e2d66b05399d8cdf74aa00ca24dc8dbb8e0.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i667ed9dad4d0c7e77f1a855f41d77968445096cfee0fe4e2d47e140b8d90b8c9.ExceptionOccurrencesRequestBuilder) {
    return i667ed9dad4d0c7e77f1a855f41d77968445096cfee0fe4e2d47e140b8d90b8c9.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i5c96f88b3ec2612713c9d12f6793cc22baf5bf2de02b2a1fbfa9092f9640d7d2.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i5c96f88b3ec2612713c9d12f6793cc22baf5bf2de02b2a1fbfa9092f9640d7d2.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i9da0074ea8a34e5df0ba5be77bbfaa546a3789dc60a880171c2e4829efb3ecbe.ExtensionsRequestBuilder) {
    return i9da0074ea8a34e5df0ba5be77bbfaa546a3789dc60a880171c2e4829efb3ecbe.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i57fba512fdd9c018cea1ea18a1c19859d0dc72eb1271b7ef534e62a89f822751.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i57fba512fdd9c018cea1ea18a1c19859d0dc72eb1271b7ef534e62a89f822751.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*iae3ac1c45f912bc7620b50de6f38592a1ecf12075e336c96f755cdb5b6b71b06.ForwardRequestBuilder) {
    return iae3ac1c45f912bc7620b50de6f38592a1ecf12075e336c96f755cdb5b6b71b06.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i017a715501a5ab4f66e77ae5b4bd5374f4e8b5b4ff4303c4f1760e18ce649a67.MultiValueExtendedPropertiesRequestBuilder) {
    return i017a715501a5ab4f66e77ae5b4bd5374f4e8b5b4ff4303c4f1760e18ce649a67.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ic860d18bcf9afeee3897ed4e52f55e614afd045fcf573e65d9fd194225b99773.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ic860d18bcf9afeee3897ed4e52f55e614afd045fcf573e65d9fd194225b99773.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in me
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property instances in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i03c8d4da7643b9e7219f7467bfa0d76efd6f53f5dcbcd3bc55bde0e2b781b6bd.SingleValueExtendedPropertiesRequestBuilder) {
    return i03c8d4da7643b9e7219f7467bfa0d76efd6f53f5dcbcd3bc55bde0e2b781b6bd.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie98d034aecbf9bae9b33a1db86754c1684a20b4959b168ca7abaf8b9100b7cea.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ie98d034aecbf9bae9b33a1db86754c1684a20b4959b168ca7abaf8b9100b7cea.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i8e9ecd6e780ee4bb2c786ff4a010f0c3f58b8ad11c339198ee764edbe774e97f.SnoozeReminderRequestBuilder) {
    return i8e9ecd6e780ee4bb2c786ff4a010f0c3f58b8ad11c339198ee764edbe774e97f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i4d57b8aaed5da088236afa277a95d913be0366ead59755ed6a66f88667b89356.TentativelyAcceptRequestBuilder) {
    return i4d57b8aaed5da088236afa277a95d913be0366ead59755ed6a66f88667b89356.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
