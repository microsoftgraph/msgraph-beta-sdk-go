package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0ae57be6544be5491d8f07f6c22d1d5c1194c936a08b56655250cefa18825664 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/dismissreminder"
    i175888625ff4bf869f5e0269cb0a6a0d91fb7fbe373b8c172b444d5db7f47082 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/attachments"
    i431d82cb500dd73f3ac073c0b45bb7d45c76c3b3e52f053fbcc6418322c4663a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties"
    i487feed4b0115f5ced94e601107a88f9b150fd40eaff3706d8209e8f99c96c77 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/extensions"
    i5040ecdd01e7f93adab613a821d0a8efd669cc1f48405f015fd466b880c73ac7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/forward"
    i5a094799dfca88d645558c9eefa74361db275a78e371f8153f179f81c5b21dd8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/tentativelyaccept"
    i7a6e2ccf12d5a83f94f2d21882e0d96bf56ed960e4689b2d4f796860412c74ea "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/calendar"
    i835c1906813c89b638c317d9dc615c2cecc25378b54c7d801a0b4d114f981539 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/multivalueextendedproperties"
    i88663444369e0125d1e82990dcc795b203c3bddcb45380d0623ad46de2711b1c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/decline"
    ib2cf79f9b281e92476da2100e6d3bc7c4432fa7eebc4abef5550dc753338665d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/snoozereminder"
    ib2e9f3cddad30b495532ff026269dae4c863bb2cda7f30e4a902683de14c5310 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences"
    ica85e02b4122e3327dd706cafb60db8367b4aa904920ea126622cdc9c1f9586c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/cancel"
    icfd497893f32ed18238b56e923f2694e380178d8d31c557b6d402ac7914cdcc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/accept"
    i1885f611e03764b8e68cc9f8559186312be120f673f0301a0bf9b18540733a34 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/multivalueextendedproperties/item"
    i445026c3831dd7e4c48babc6cd0beed2d470b2cdd428b8421fc6dcf7df1caa19 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/singlevalueextendedproperties/item"
    iac328d6d4b61570ae7f06cd56092beaeb0eb5940ff8417b79af8ad5d701d0a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/attachments/item"
    iafe4d248276388987d7cdf9ff83b9a107a480d16448c88d4517740c6cf0304c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/extensions/item"
    ic434547d141e4d72908ae3b00133c5bf2a8199ca025c608868aed8b016aa2890 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/instances/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*icfd497893f32ed18238b56e923f2694e380178d8d31c557b6d402ac7914cdcc4.AcceptRequestBuilder) {
    return icfd497893f32ed18238b56e923f2694e380178d8d31c557b6d402ac7914cdcc4.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i175888625ff4bf869f5e0269cb0a6a0d91fb7fbe373b8c172b444d5db7f47082.AttachmentsRequestBuilder) {
    return i175888625ff4bf869f5e0269cb0a6a0d91fb7fbe373b8c172b444d5db7f47082.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*iac328d6d4b61570ae7f06cd56092beaeb0eb5940ff8417b79af8ad5d701d0a91.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return iac328d6d4b61570ae7f06cd56092beaeb0eb5940ff8417b79af8ad5d701d0a91.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i7a6e2ccf12d5a83f94f2d21882e0d96bf56ed960e4689b2d4f796860412c74ea.CalendarRequestBuilder) {
    return i7a6e2ccf12d5a83f94f2d21882e0d96bf56ed960e4689b2d4f796860412c74ea.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*ica85e02b4122e3327dd706cafb60db8367b4aa904920ea126622cdc9c1f9586c.CancelRequestBuilder) {
    return ica85e02b4122e3327dd706cafb60db8367b4aa904920ea126622cdc9c1f9586c.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
// CreateDeleteRequestInformation delete navigation property instances for users
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property instances for users
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
// CreatePatchRequestInformation update the navigation property instances in users
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property instances in users
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
func (m *EventItemRequestBuilder) Decline()(*i88663444369e0125d1e82990dcc795b203c3bddcb45380d0623ad46de2711b1c.DeclineRequestBuilder) {
    return i88663444369e0125d1e82990dcc795b203c3bddcb45380d0623ad46de2711b1c.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for users
func (m *EventItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property instances for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i0ae57be6544be5491d8f07f6c22d1d5c1194c936a08b56655250cefa18825664.DismissReminderRequestBuilder) {
    return i0ae57be6544be5491d8f07f6c22d1d5c1194c936a08b56655250cefa18825664.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ib2e9f3cddad30b495532ff026269dae4c863bb2cda7f30e4a902683de14c5310.ExceptionOccurrencesRequestBuilder) {
    return ib2e9f3cddad30b495532ff026269dae4c863bb2cda7f30e4a902683de14c5310.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*ic434547d141e4d72908ae3b00133c5bf2a8199ca025c608868aed8b016aa2890.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return ic434547d141e4d72908ae3b00133c5bf2a8199ca025c608868aed8b016aa2890.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i487feed4b0115f5ced94e601107a88f9b150fd40eaff3706d8209e8f99c96c77.ExtensionsRequestBuilder) {
    return i487feed4b0115f5ced94e601107a88f9b150fd40eaff3706d8209e8f99c96c77.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*iafe4d248276388987d7cdf9ff83b9a107a480d16448c88d4517740c6cf0304c8.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return iafe4d248276388987d7cdf9ff83b9a107a480d16448c88d4517740c6cf0304c8.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i5040ecdd01e7f93adab613a821d0a8efd669cc1f48405f015fd466b880c73ac7.ForwardRequestBuilder) {
    return i5040ecdd01e7f93adab613a821d0a8efd669cc1f48405f015fd466b880c73ac7.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i835c1906813c89b638c317d9dc615c2cecc25378b54c7d801a0b4d114f981539.MultiValueExtendedPropertiesRequestBuilder) {
    return i835c1906813c89b638c317d9dc615c2cecc25378b54c7d801a0b4d114f981539.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i1885f611e03764b8e68cc9f8559186312be120f673f0301a0bf9b18540733a34.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i1885f611e03764b8e68cc9f8559186312be120f673f0301a0bf9b18540733a34.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in users
func (m *EventItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property instances in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i431d82cb500dd73f3ac073c0b45bb7d45c76c3b3e52f053fbcc6418322c4663a.SingleValueExtendedPropertiesRequestBuilder) {
    return i431d82cb500dd73f3ac073c0b45bb7d45c76c3b3e52f053fbcc6418322c4663a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendars.item.calendarView.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i445026c3831dd7e4c48babc6cd0beed2d470b2cdd428b8421fc6dcf7df1caa19.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i445026c3831dd7e4c48babc6cd0beed2d470b2cdd428b8421fc6dcf7df1caa19.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ib2cf79f9b281e92476da2100e6d3bc7c4432fa7eebc4abef5550dc753338665d.SnoozeReminderRequestBuilder) {
    return ib2cf79f9b281e92476da2100e6d3bc7c4432fa7eebc4abef5550dc753338665d.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i5a094799dfca88d645558c9eefa74361db275a78e371f8153f179f81c5b21dd8.TentativelyAcceptRequestBuilder) {
    return i5a094799dfca88d645558c9eefa74361db275a78e371f8153f179f81c5b21dd8.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
