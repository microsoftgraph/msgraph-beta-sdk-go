package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1a62b90911d2b15be7dc165dba44dfd2113030f9f9059c096ef1b302ef4c4e1e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    i1c6afa562f7836937a411805e22acd011a77c5b9aa3250c74214aad769a4c30e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    i4c8cd9a96e98f2ac27d13fcb62743c20a2066ab1d331431c6e35a0d543776dc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    i51ffe4971dc9d6f04febfb038acafa1a56a90b4e0e46f4d71733df1c4ae6b9ac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i5517f181917105f7ef98e8c382ce445267f01e06c079059d078ead792c67e6d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    i5b1c6047e6a917355bc9be0feacc39bf4b562bff0b369ba68d4b4522cbdb4e95 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    i7045d2b025173472e6196bf63f7e04ba3472037cf57cd99ae2a23e3e20667c33 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i79c099c1c97fcfbaaaf8a59ccf2342bfb1e1ee2ab25277a43e523968e9444779 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i84aec0f0f45d84db356768abf6f085e0fc79b6928ed32048f6e28faa6d90fdc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    i89d5461ff77549a50cdd2557666fc164985139cd24d5b98bdb3cc08d23022ade "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    if5686551850d2aa97403c52b36d610b1ad772c143542323bca08fcf3b8be5c1b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    if995c06bf4c58b3ec328e38a2c4f10ecb3dfebcdc9b8735b44715543a485bce3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i0fadeb49753ed8a0663dc6d56629d28724748cda4c440bbddb1b5e9cb84c827a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    i2cedeb995f4b45129646faa674853ecbf6a57947cad1831338b37e93e02c935f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
    ib51bca3e04f3cd5f5952ac6f94fcddf71acb735d085a9109a734489adb39c3e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    ic11e0ad5bf8552a9a59cb665fe685e4e2b29a668163b2da40029313760c3c2b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendargroups/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
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
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i84aec0f0f45d84db356768abf6f085e0fc79b6928ed32048f6e28faa6d90fdc2.AcceptRequestBuilder) {
    return i84aec0f0f45d84db356768abf6f085e0fc79b6928ed32048f6e28faa6d90fdc2.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i1a62b90911d2b15be7dc165dba44dfd2113030f9f9059c096ef1b302ef4c4e1e.AttachmentsRequestBuilder) {
    return i1a62b90911d2b15be7dc165dba44dfd2113030f9f9059c096ef1b302ef4c4e1e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i2cedeb995f4b45129646faa674853ecbf6a57947cad1831338b37e93e02c935f.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i2cedeb995f4b45129646faa674853ecbf6a57947cad1831338b37e93e02c935f.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i89d5461ff77549a50cdd2557666fc164985139cd24d5b98bdb3cc08d23022ade.CalendarRequestBuilder) {
    return i89d5461ff77549a50cdd2557666fc164985139cd24d5b98bdb3cc08d23022ade.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i5b1c6047e6a917355bc9be0feacc39bf4b562bff0b369ba68d4b4522cbdb4e95.CancelRequestBuilder) {
    return i5b1c6047e6a917355bc9be0feacc39bf4b562bff0b369ba68d4b4522cbdb4e95.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i4c8cd9a96e98f2ac27d13fcb62743c20a2066ab1d331431c6e35a0d543776dc7.DeclineRequestBuilder) {
    return i4c8cd9a96e98f2ac27d13fcb62743c20a2066ab1d331431c6e35a0d543776dc7.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*if995c06bf4c58b3ec328e38a2c4f10ecb3dfebcdc9b8735b44715543a485bce3.DismissReminderRequestBuilder) {
    return if995c06bf4c58b3ec328e38a2c4f10ecb3dfebcdc9b8735b44715543a485bce3.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i5517f181917105f7ef98e8c382ce445267f01e06c079059d078ead792c67e6d6.ExtensionsRequestBuilder) {
    return i5517f181917105f7ef98e8c382ce445267f01e06c079059d078ead792c67e6d6.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic11e0ad5bf8552a9a59cb665fe685e4e2b29a668163b2da40029313760c3c2b0.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ic11e0ad5bf8552a9a59cb665fe685e4e2b29a668163b2da40029313760c3c2b0.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i1c6afa562f7836937a411805e22acd011a77c5b9aa3250c74214aad769a4c30e.ForwardRequestBuilder) {
    return i1c6afa562f7836937a411805e22acd011a77c5b9aa3250c74214aad769a4c30e.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i7045d2b025173472e6196bf63f7e04ba3472037cf57cd99ae2a23e3e20667c33.MultiValueExtendedPropertiesRequestBuilder) {
    return i7045d2b025173472e6196bf63f7e04ba3472037cf57cd99ae2a23e3e20667c33.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0fadeb49753ed8a0663dc6d56629d28724748cda4c440bbddb1b5e9cb84c827a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i0fadeb49753ed8a0663dc6d56629d28724748cda4c440bbddb1b5e9cb84c827a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i79c099c1c97fcfbaaaf8a59ccf2342bfb1e1ee2ab25277a43e523968e9444779.SingleValueExtendedPropertiesRequestBuilder) {
    return i79c099c1c97fcfbaaaf8a59ccf2342bfb1e1ee2ab25277a43e523968e9444779.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarGroups.item.calendars.item.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib51bca3e04f3cd5f5952ac6f94fcddf71acb735d085a9109a734489adb39c3e8.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ib51bca3e04f3cd5f5952ac6f94fcddf71acb735d085a9109a734489adb39c3e8.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*if5686551850d2aa97403c52b36d610b1ad772c143542323bca08fcf3b8be5c1b.SnoozeReminderRequestBuilder) {
    return if5686551850d2aa97403c52b36d610b1ad772c143542323bca08fcf3b8be5c1b.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i51ffe4971dc9d6f04febfb038acafa1a56a90b4e0e46f4d71733df1c4ae6b9ac.TentativelyAcceptRequestBuilder) {
    return i51ffe4971dc9d6f04febfb038acafa1a56a90b4e0e46f4d71733df1c4ae6b9ac.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
