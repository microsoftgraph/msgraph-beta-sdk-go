package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0c3995287d1b4bbcac2996824ccfa6bd0042948aa8beedf7f2506da0880c8f57 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i11a58a628af76076d57a3c984c0ceb0bef49bc6a9389499b2f576e01d294a9cd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    i14d561ec4d8986189ec738a1f74ce30575c2d20fd9e51a9b05968a43eef64991 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    i1fe8a36de23d439eea59c93683fbda09f544890f89c324f9dfcde72e8c57cf10 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    i2f558291f74e8633215378da1df32d6b02781fb5368d287952baca29e76a5865 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i8cd08175b7c660c07c68b4a50634430ea9d32a8ed07548d6a634548178cbf934 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    ia216054458b0ab1ddfa5b06b713031c4a6dba3352b21b962e5067c2de1d88161 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    iabf07b3bea9017f8424f235c28fa81e3bf56af38da93f2e4b8d6429c40b4e9db "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    ibd171122d61c7f727e2a73dd62cd0aa7ea5eca1ccf80b468e7adce2b394a7512 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    id7772fea5a533e0020fefda3c2c96f51aaa7a850380a8ace7e8e8a27e78e5354 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    ie2f85f708b1460d715838fadd2f859b6107b03e817c3758b0370c6de0991f7d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    ie9bfebdc3d75960a689eaaf3ed9f6e2c4c0130557a6882be2764cd22951b322e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i0740ff41752510227152c765a034e3ce3e428339603678132d218e6383fd5921 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    i7528bf2313ae037c80fb355568fe7d5599c09fdc6deb10271cea89d15a6084ad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    ibf2868d4347d0c2036da774ec868bb17d74df06c76c720de5a4a05b9a58202de "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    ied617b938e8d20da088146865e730cf8a5f9452ba92cca4dd4c70bd7e21862c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
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
// Accept provides operations to call the accept method.
func (m *EventItemRequestBuilder) Accept()(*id7772fea5a533e0020fefda3c2c96f51aaa7a850380a8ace7e8e8a27e78e5354.AcceptRequestBuilder) {
    return id7772fea5a533e0020fefda3c2c96f51aaa7a850380a8ace7e8e8a27e78e5354.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*i14d561ec4d8986189ec738a1f74ce30575c2d20fd9e51a9b05968a43eef64991.AttachmentsRequestBuilder) {
    return i14d561ec4d8986189ec738a1f74ce30575c2d20fd9e51a9b05968a43eef64991.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ied617b938e8d20da088146865e730cf8a5f9452ba92cca4dd4c70bd7e21862c2.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return ied617b938e8d20da088146865e730cf8a5f9452ba92cca4dd4c70bd7e21862c2.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*ia216054458b0ab1ddfa5b06b713031c4a6dba3352b21b962e5067c2de1d88161.CalendarRequestBuilder) {
    return ia216054458b0ab1ddfa5b06b713031c4a6dba3352b21b962e5067c2de1d88161.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*ie2f85f708b1460d715838fadd2f859b6107b03e817c3758b0370c6de0991f7d5.CancelRequestBuilder) {
    return ie2f85f708b1460d715838fadd2f859b6107b03e817c3758b0370c6de0991f7d5.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Decline provides operations to call the decline method.
func (m *EventItemRequestBuilder) Decline()(*i1fe8a36de23d439eea59c93683fbda09f544890f89c324f9dfcde72e8c57cf10.DeclineRequestBuilder) {
    return i1fe8a36de23d439eea59c93683fbda09f544890f89c324f9dfcde72e8c57cf10.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i0c3995287d1b4bbcac2996824ccfa6bd0042948aa8beedf7f2506da0880c8f57.DismissReminderRequestBuilder) {
    return i0c3995287d1b4bbcac2996824ccfa6bd0042948aa8beedf7f2506da0880c8f57.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*ibd171122d61c7f727e2a73dd62cd0aa7ea5eca1ccf80b468e7adce2b394a7512.ExtensionsRequestBuilder) {
    return ibd171122d61c7f727e2a73dd62cd0aa7ea5eca1ccf80b468e7adce2b394a7512.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i7528bf2313ae037c80fb355568fe7d5599c09fdc6deb10271cea89d15a6084ad.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i7528bf2313ae037c80fb355568fe7d5599c09fdc6deb10271cea89d15a6084ad.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i8cd08175b7c660c07c68b4a50634430ea9d32a8ed07548d6a634548178cbf934.ForwardRequestBuilder) {
    return i8cd08175b7c660c07c68b4a50634430ea9d32a8ed07548d6a634548178cbf934.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ie9bfebdc3d75960a689eaaf3ed9f6e2c4c0130557a6882be2764cd22951b322e.MultiValueExtendedPropertiesRequestBuilder) {
    return ie9bfebdc3d75960a689eaaf3ed9f6e2c4c0130557a6882be2764cd22951b322e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ibf2868d4347d0c2036da774ec868bb17d74df06c76c720de5a4a05b9a58202de.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ibf2868d4347d0c2036da774ec868bb17d74df06c76c720de5a4a05b9a58202de.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i2f558291f74e8633215378da1df32d6b02781fb5368d287952baca29e76a5865.SingleValueExtendedPropertiesRequestBuilder) {
    return i2f558291f74e8633215378da1df32d6b02781fb5368d287952baca29e76a5865.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0740ff41752510227152c765a034e3ce3e428339603678132d218e6383fd5921.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i0740ff41752510227152c765a034e3ce3e428339603678132d218e6383fd5921.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*iabf07b3bea9017f8424f235c28fa81e3bf56af38da93f2e4b8d6429c40b4e9db.SnoozeReminderRequestBuilder) {
    return iabf07b3bea9017f8424f235c28fa81e3bf56af38da93f2e4b8d6429c40b4e9db.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*i11a58a628af76076d57a3c984c0ceb0bef49bc6a9389499b2f576e01d294a9cd.TentativelyAcceptRequestBuilder) {
    return i11a58a628af76076d57a3c984c0ceb0bef49bc6a9389499b2f576e01d294a9cd.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
