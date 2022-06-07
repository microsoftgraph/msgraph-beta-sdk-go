package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i19df0ab3d061d701ae7c98c2427ce2fa10b87e3325e37698f8f582e592e57a82 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/multivalueextendedproperties"
    i230d2b67bc5ce07238d33f9511acd4155630b958ab43aa19f44e82db7251388d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/dismissreminder"
    i329f53cda9d26a100576fd55e49bfb1631a4a4cc14f81413d47eb4de9ef7bc83 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/snoozereminder"
    i33d49beeb8d637d8f70d956515fb4eb71e74e2fcffc2e6fac4fe1c209e6d519c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/cancel"
    i76dbdebfd0ba40974265fbbd458a9873c2c1cc676ea506c3eabc7cb56611abc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/attachments"
    i7edc42ee14e3746bb2f507d3e07365aa6041e048cfe7476b79438d78ccd092b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/decline"
    i9eb9259230c7bb1bff775ff5e1c6418a0a210562ba1823e35b52840c9078f423 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/calendar"
    ic353fa9d30f10509120433d9258cdf547a623db16784d3250f18186a66a7a1ab "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/forward"
    id7bddaa58f8d7c0f2fbefa3379c51cd50652f9e40203a33e998b1ba9be46e375 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/accept"
    id802b2d6d3c0f3f509fe2e2ac81003f6cc0fba5f0403bd96d4538c70be09ce71 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences"
    ida3da94d938e8837c31b6d6f832613b3be2b7a0a243cad86d626d14c8f15b452 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/extensions"
    if16e10298df837937a2788fadd430c8e3c1f4e013ee6ff8b4e79323587f77e10 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/tentativelyaccept"
    if64fe8feea9af0d49e1bfad9e4391adae4728e31837d09e03f5f1c427a5c2a85 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties"
    i10b0a17e1d4c46e96fbacd22c9a233cd99490e87c2404e0ec97d86a4d04f46df "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/multivalueextendedproperties/item"
    i307a9153df02d5652c3462486f9884235fbea2c01f98dc032dfb9565e66abff9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
    i5ac9f3f6d1248471b0eb1ef74aed334784b1c2086c2c9d1eb8d9ce31171b6e45 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/extensions/item"
    i664e25718be471c50930925f8171d81ed927ae5d1428fee023d7dd9a59707c4c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/attachments/item"
    i97f92c90f82f9463545346d23d267dc2848663994ae298fdf2b488cc8f6b81a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*id7bddaa58f8d7c0f2fbefa3379c51cd50652f9e40203a33e998b1ba9be46e375.AcceptRequestBuilder) {
    return id7bddaa58f8d7c0f2fbefa3379c51cd50652f9e40203a33e998b1ba9be46e375.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i76dbdebfd0ba40974265fbbd458a9873c2c1cc676ea506c3eabc7cb56611abc0.AttachmentsRequestBuilder) {
    return i76dbdebfd0ba40974265fbbd458a9873c2c1cc676ea506c3eabc7cb56611abc0.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i664e25718be471c50930925f8171d81ed927ae5d1428fee023d7dd9a59707c4c.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i664e25718be471c50930925f8171d81ed927ae5d1428fee023d7dd9a59707c4c.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i9eb9259230c7bb1bff775ff5e1c6418a0a210562ba1823e35b52840c9078f423.CalendarRequestBuilder) {
    return i9eb9259230c7bb1bff775ff5e1c6418a0a210562ba1823e35b52840c9078f423.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i33d49beeb8d637d8f70d956515fb4eb71e74e2fcffc2e6fac4fe1c209e6d519c.CancelRequestBuilder) {
    return i33d49beeb8d637d8f70d956515fb4eb71e74e2fcffc2e6fac4fe1c209e6d519c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i7edc42ee14e3746bb2f507d3e07365aa6041e048cfe7476b79438d78ccd092b0.DeclineRequestBuilder) {
    return i7edc42ee14e3746bb2f507d3e07365aa6041e048cfe7476b79438d78ccd092b0.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i230d2b67bc5ce07238d33f9511acd4155630b958ab43aa19f44e82db7251388d.DismissReminderRequestBuilder) {
    return i230d2b67bc5ce07238d33f9511acd4155630b958ab43aa19f44e82db7251388d.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*id802b2d6d3c0f3f509fe2e2ac81003f6cc0fba5f0403bd96d4538c70be09ce71.ExceptionOccurrencesRequestBuilder) {
    return id802b2d6d3c0f3f509fe2e2ac81003f6cc0fba5f0403bd96d4538c70be09ce71.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i97f92c90f82f9463545346d23d267dc2848663994ae298fdf2b488cc8f6b81a7.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return i97f92c90f82f9463545346d23d267dc2848663994ae298fdf2b488cc8f6b81a7.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*ida3da94d938e8837c31b6d6f832613b3be2b7a0a243cad86d626d14c8f15b452.ExtensionsRequestBuilder) {
    return ida3da94d938e8837c31b6d6f832613b3be2b7a0a243cad86d626d14c8f15b452.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i5ac9f3f6d1248471b0eb1ef74aed334784b1c2086c2c9d1eb8d9ce31171b6e45.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i5ac9f3f6d1248471b0eb1ef74aed334784b1c2086c2c9d1eb8d9ce31171b6e45.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*ic353fa9d30f10509120433d9258cdf547a623db16784d3250f18186a66a7a1ab.ForwardRequestBuilder) {
    return ic353fa9d30f10509120433d9258cdf547a623db16784d3250f18186a66a7a1ab.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i19df0ab3d061d701ae7c98c2427ce2fa10b87e3325e37698f8f582e592e57a82.MultiValueExtendedPropertiesRequestBuilder) {
    return i19df0ab3d061d701ae7c98c2427ce2fa10b87e3325e37698f8f582e592e57a82.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i10b0a17e1d4c46e96fbacd22c9a233cd99490e87c2404e0ec97d86a4d04f46df.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i10b0a17e1d4c46e96fbacd22c9a233cd99490e87c2404e0ec97d86a4d04f46df.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*if64fe8feea9af0d49e1bfad9e4391adae4728e31837d09e03f5f1c427a5c2a85.SingleValueExtendedPropertiesRequestBuilder) {
    return if64fe8feea9af0d49e1bfad9e4391adae4728e31837d09e03f5f1c427a5c2a85.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i307a9153df02d5652c3462486f9884235fbea2c01f98dc032dfb9565e66abff9.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i307a9153df02d5652c3462486f9884235fbea2c01f98dc032dfb9565e66abff9.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i329f53cda9d26a100576fd55e49bfb1631a4a4cc14f81413d47eb4de9ef7bc83.SnoozeReminderRequestBuilder) {
    return i329f53cda9d26a100576fd55e49bfb1631a4a4cc14f81413d47eb4de9ef7bc83.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*if16e10298df837937a2788fadd430c8e3c1f4e013ee6ff8b4e79323587f77e10.TentativelyAcceptRequestBuilder) {
    return if16e10298df837937a2788fadd430c8e3c1f4e013ee6ff8b4e79323587f77e10.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
