package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03b3809121ebbb4ba1b1d15fd81be35c33cc1a35cc0146cb18c18ecccb41d527 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i2e254c66b6056240e450a8cf9a3af654d462493cd47da925da8908d589f99f84 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    i37b21e0db6b9accd70ed5809e6903df7f199b8ad6405f3dc79105de9ba372d35 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    i3e02f7bceaf559a5d5cbf716df2c3b1cb419b5184a93c46b6bd96af89663e1a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    i421f006db32e222dd3eb536917647689a121aa8e05b8ab8a2ea7dede2aac7e03 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i53e0e9bd522bc9d743ffd152bcbe922d8e50fb529152d7e82feb312d3a15dafb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    i5a580ff2748a6abe72c5ac9a12cbb3aabf5485331237fa1297afcb9abf95337e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    i6ba051e31f44b920f1c90221dd72a7ea5e85b4bbb621f4fc24552ba7511e66ff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i7e24a17eec3b5d851dc6385138718e6dd3a5e6399e978df43eaa0feb30d87f2c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    ica3d48dc4c37ddbc964ce06bdd03c200f840a5029d773491d0c3e03df2b04fb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    ied33059619ae6a4d893b4b53c1ef14afb557e6453336d62b060064e285ba4f64 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    if4112c5a2c132239ddbeb5fa41da8ed548aaee3df95d3e544a68e82ba254b078 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    i66a70b7136972ef22913fbff1d666a3e71272f1fb839071555b56ac80e189ba1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    i93d92afabc944dbb98a4ec3bc3ec83f7c15e39fa86c23481cd0bbbe0f10b9e5d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
    ieed57b6f47a051553245939912c9aa0205f1496fe233ed1ae6f7209c1e76bc7b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    if7f9fc0becac793e4caad12234ed60b21a9365e9531176783bdc06d8bcebab0a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*if4112c5a2c132239ddbeb5fa41da8ed548aaee3df95d3e544a68e82ba254b078.AcceptRequestBuilder) {
    return if4112c5a2c132239ddbeb5fa41da8ed548aaee3df95d3e544a68e82ba254b078.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i3e02f7bceaf559a5d5cbf716df2c3b1cb419b5184a93c46b6bd96af89663e1a8.AttachmentsRequestBuilder) {
    return i3e02f7bceaf559a5d5cbf716df2c3b1cb419b5184a93c46b6bd96af89663e1a8.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i93d92afabc944dbb98a4ec3bc3ec83f7c15e39fa86c23481cd0bbbe0f10b9e5d.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i93d92afabc944dbb98a4ec3bc3ec83f7c15e39fa86c23481cd0bbbe0f10b9e5d.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ica3d48dc4c37ddbc964ce06bdd03c200f840a5029d773491d0c3e03df2b04fb2.CalendarRequestBuilder) {
    return ica3d48dc4c37ddbc964ce06bdd03c200f840a5029d773491d0c3e03df2b04fb2.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i37b21e0db6b9accd70ed5809e6903df7f199b8ad6405f3dc79105de9ba372d35.CancelRequestBuilder) {
    return i37b21e0db6b9accd70ed5809e6903df7f199b8ad6405f3dc79105de9ba372d35.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i2e254c66b6056240e450a8cf9a3af654d462493cd47da925da8908d589f99f84.DeclineRequestBuilder) {
    return i2e254c66b6056240e450a8cf9a3af654d462493cd47da925da8908d589f99f84.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) DismissReminder()(*i421f006db32e222dd3eb536917647689a121aa8e05b8ab8a2ea7dede2aac7e03.DismissReminderRequestBuilder) {
    return i421f006db32e222dd3eb536917647689a121aa8e05b8ab8a2ea7dede2aac7e03.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i53e0e9bd522bc9d743ffd152bcbe922d8e50fb529152d7e82feb312d3a15dafb.ExtensionsRequestBuilder) {
    return i53e0e9bd522bc9d743ffd152bcbe922d8e50fb529152d7e82feb312d3a15dafb.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i66a70b7136972ef22913fbff1d666a3e71272f1fb839071555b56ac80e189ba1.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i66a70b7136972ef22913fbff1d666a3e71272f1fb839071555b56ac80e189ba1.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i5a580ff2748a6abe72c5ac9a12cbb3aabf5485331237fa1297afcb9abf95337e.ForwardRequestBuilder) {
    return i5a580ff2748a6abe72c5ac9a12cbb3aabf5485331237fa1297afcb9abf95337e.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i03b3809121ebbb4ba1b1d15fd81be35c33cc1a35cc0146cb18c18ecccb41d527.MultiValueExtendedPropertiesRequestBuilder) {
    return i03b3809121ebbb4ba1b1d15fd81be35c33cc1a35cc0146cb18c18ecccb41d527.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if7f9fc0becac793e4caad12234ed60b21a9365e9531176783bdc06d8bcebab0a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return if7f9fc0becac793e4caad12234ed60b21a9365e9531176783bdc06d8bcebab0a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i7e24a17eec3b5d851dc6385138718e6dd3a5e6399e978df43eaa0feb30d87f2c.SingleValueExtendedPropertiesRequestBuilder) {
    return i7e24a17eec3b5d851dc6385138718e6dd3a5e6399e978df43eaa0feb30d87f2c.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ieed57b6f47a051553245939912c9aa0205f1496fe233ed1ae6f7209c1e76bc7b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ieed57b6f47a051553245939912c9aa0205f1496fe233ed1ae6f7209c1e76bc7b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ied33059619ae6a4d893b4b53c1ef14afb557e6453336d62b060064e285ba4f64.SnoozeReminderRequestBuilder) {
    return ied33059619ae6a4d893b4b53c1ef14afb557e6453336d62b060064e285ba4f64.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i6ba051e31f44b920f1c90221dd72a7ea5e85b4bbb621f4fc24552ba7511e66ff.TentativelyAcceptRequestBuilder) {
    return i6ba051e31f44b920f1c90221dd72a7ea5e85b4bbb621f4fc24552ba7511e66ff.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
