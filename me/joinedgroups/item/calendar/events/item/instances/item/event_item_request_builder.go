package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03e890f2e3e48d5e42767b7b08111e199f7f7cf6536ea1f2342c3873cf8c9102 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/tentativelyaccept"
    i229e2c55b88fa6315ed95d09f0a91093cbc85bf005247544adc647cd4af937e7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/extensions"
    i6dbeec2eee6a96b33f586c7bde1e995ed6ebbdfc677681d69a1e6b44401f83d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/accept"
    i7bd4a7a0edb412724a05e7ad4be6226a3cfe5822865ec6e91f6c01975a56076e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/forward"
    i8318cacc6a55d01241f68f301e47cdb9b2d27173327185516993853890d906a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/calendar"
    i85cf645f89a0dfd16bb403c031f2456bf17f437b2041313dc8a29dd592e6086f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/dismissreminder"
    i876b09c0e799e616b30d02170a3e0360eb0419e735069defc206a887899919d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/cancel"
    ib0662937155354cade70207fea7846387b987f12345817b366f3ea9f52195330 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/singlevalueextendedproperties"
    ib6305fe4da10dcf1d1ea0d976a48d76e3cac454340ac5773d8fca9da3a8523c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/snoozereminder"
    ic99a14c71118136304d4f8b55ccbf2e0ca2e2959eb919dfb3d219e795c341ff4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/attachments"
    if0a45dbb22bace5b353a66e17dca935ba871330fdc19d0228cb421f2c2d53095 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/decline"
    if17488e00f34b349a47d3bfbf2585194e041a748b35b86b670be04b8c432ea2f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/multivalueextendedproperties"
    if425449d72a894953db480b7b1dc63a42e7eeacba10b9ddcac29ab1da0253af1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences"
    i47f0327a7f89228ce0ed47fab0caa2175d693b06a80bc618d4a262b15a19a345 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/attachments/item"
    i59418681a154ac27f8065e25ca11ff075c1aae2164baa7b6d6abd4506fede5ae "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/singlevalueextendedproperties/item"
    i80199bd852a27b26586137b4575e4992ff786ab6c4a31382cc8546bec03627cf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/multivalueextendedproperties/item"
    ia434f24bf86aa24406c19164bd3c4bb9a38f64d1f6e17adb8dae5a2d44e4ff05 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/extensions/item"
    iab6b69feea02c9906197e9ef5bc6460d1a1fd8a1f3c32fb32dc92c8fa5bed584 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/calendar/events/item/instances/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*i6dbeec2eee6a96b33f586c7bde1e995ed6ebbdfc677681d69a1e6b44401f83d1.AcceptRequestBuilder) {
    return i6dbeec2eee6a96b33f586c7bde1e995ed6ebbdfc677681d69a1e6b44401f83d1.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ic99a14c71118136304d4f8b55ccbf2e0ca2e2959eb919dfb3d219e795c341ff4.AttachmentsRequestBuilder) {
    return ic99a14c71118136304d4f8b55ccbf2e0ca2e2959eb919dfb3d219e795c341ff4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i47f0327a7f89228ce0ed47fab0caa2175d693b06a80bc618d4a262b15a19a345.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i47f0327a7f89228ce0ed47fab0caa2175d693b06a80bc618d4a262b15a19a345.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i8318cacc6a55d01241f68f301e47cdb9b2d27173327185516993853890d906a4.CalendarRequestBuilder) {
    return i8318cacc6a55d01241f68f301e47cdb9b2d27173327185516993853890d906a4.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i876b09c0e799e616b30d02170a3e0360eb0419e735069defc206a887899919d4.CancelRequestBuilder) {
    return i876b09c0e799e616b30d02170a3e0360eb0419e735069defc206a887899919d4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*if0a45dbb22bace5b353a66e17dca935ba871330fdc19d0228cb421f2c2d53095.DeclineRequestBuilder) {
    return if0a45dbb22bace5b353a66e17dca935ba871330fdc19d0228cb421f2c2d53095.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i85cf645f89a0dfd16bb403c031f2456bf17f437b2041313dc8a29dd592e6086f.DismissReminderRequestBuilder) {
    return i85cf645f89a0dfd16bb403c031f2456bf17f437b2041313dc8a29dd592e6086f.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*if425449d72a894953db480b7b1dc63a42e7eeacba10b9ddcac29ab1da0253af1.ExceptionOccurrencesRequestBuilder) {
    return if425449d72a894953db480b7b1dc63a42e7eeacba10b9ddcac29ab1da0253af1.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*iab6b69feea02c9906197e9ef5bc6460d1a1fd8a1f3c32fb32dc92c8fa5bed584.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return iab6b69feea02c9906197e9ef5bc6460d1a1fd8a1f3c32fb32dc92c8fa5bed584.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i229e2c55b88fa6315ed95d09f0a91093cbc85bf005247544adc647cd4af937e7.ExtensionsRequestBuilder) {
    return i229e2c55b88fa6315ed95d09f0a91093cbc85bf005247544adc647cd4af937e7.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ia434f24bf86aa24406c19164bd3c4bb9a38f64d1f6e17adb8dae5a2d44e4ff05.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ia434f24bf86aa24406c19164bd3c4bb9a38f64d1f6e17adb8dae5a2d44e4ff05.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i7bd4a7a0edb412724a05e7ad4be6226a3cfe5822865ec6e91f6c01975a56076e.ForwardRequestBuilder) {
    return i7bd4a7a0edb412724a05e7ad4be6226a3cfe5822865ec6e91f6c01975a56076e.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*if17488e00f34b349a47d3bfbf2585194e041a748b35b86b670be04b8c432ea2f.MultiValueExtendedPropertiesRequestBuilder) {
    return if17488e00f34b349a47d3bfbf2585194e041a748b35b86b670be04b8c432ea2f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i80199bd852a27b26586137b4575e4992ff786ab6c4a31382cc8546bec03627cf.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i80199bd852a27b26586137b4575e4992ff786ab6c4a31382cc8546bec03627cf.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ib0662937155354cade70207fea7846387b987f12345817b366f3ea9f52195330.SingleValueExtendedPropertiesRequestBuilder) {
    return ib0662937155354cade70207fea7846387b987f12345817b366f3ea9f52195330.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.calendar.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i59418681a154ac27f8065e25ca11ff075c1aae2164baa7b6d6abd4506fede5ae.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i59418681a154ac27f8065e25ca11ff075c1aae2164baa7b6d6abd4506fede5ae.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ib6305fe4da10dcf1d1ea0d976a48d76e3cac454340ac5773d8fca9da3a8523c4.SnoozeReminderRequestBuilder) {
    return ib6305fe4da10dcf1d1ea0d976a48d76e3cac454340ac5773d8fca9da3a8523c4.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i03e890f2e3e48d5e42767b7b08111e199f7f7cf6536ea1f2342c3873cf8c9102.TentativelyAcceptRequestBuilder) {
    return i03e890f2e3e48d5e42767b7b08111e199f7f7cf6536ea1f2342c3873cf8c9102.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
