package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04547fc25f9a19b56375a31791a646bb9a7291300dab821083dcd9dff492eec5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i0cc9b39109a903e3c7493bc95907d6b99e23de8914847bebc5b2b7718aa48e85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    i1b93fb71e20e87d0c94ea6fbd6fb298fc08fb5d0e4a2e8f9c34137b47ee5e9d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    i2061516be4b73f307ef27998a5436c0d6263ff0f2dfe896d758d303753de3249 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    i3c0518909ebf9cd483b89bfb6b30484ba04b2be163059af0f1fb66f6749d7f13 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    iaaf46f7a2ead40b95e11dc338f68f4cc605581625ff09115cf3dfa212c70c55d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    iba1ffd0464b70bb45422cece7b13a6e10464f2457707c0bf58770857e0fccd33 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    ibbcd29199f49a1e2c6067cb3d21b38496815d0537ce16e70bbffe8a002c63ac3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    ibdcbb5181e62a6c318e7e080ec9929125febe7b6feb6ff10777c7a7626f2ccf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    ic8e9ae378ae88a52d0b547ce11c9cff604379b3f87d4a0da0c87d617167b5c54 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    ica549a59bb07a4b59d1250ec60c1c7f20cc7b2c95841d2af82d8b77c130ebf03 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    if9d4d521baacac2476f774dd755d2919846176c52baa2d6a33807091c4f57372 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    i64550b35c048770208de1ceba65993bad9d9150195e269fb6b146ca69ce8afc7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
    i9b67f74d1e1e9cf957ac05505c98e12f20babaf600449186355ccc1ef198c643 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    ib7178ba408e138ce3cdd93a42a6b5a070b0b033794f27a9ebb32ed3651994b60 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    id6bbec9cfa87470cc5313b0414204e4191e208f07a93b995e7462d3584214ebb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
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
func (m *EventItemRequestBuilder) Accept()(*i1b93fb71e20e87d0c94ea6fbd6fb298fc08fb5d0e4a2e8f9c34137b47ee5e9d3.AcceptRequestBuilder) {
    return i1b93fb71e20e87d0c94ea6fbd6fb298fc08fb5d0e4a2e8f9c34137b47ee5e9d3.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Attachments()(*iba1ffd0464b70bb45422cece7b13a6e10464f2457707c0bf58770857e0fccd33.AttachmentsRequestBuilder) {
    return iba1ffd0464b70bb45422cece7b13a6e10464f2457707c0bf58770857e0fccd33.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i64550b35c048770208de1ceba65993bad9d9150195e269fb6b146ca69ce8afc7.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i64550b35c048770208de1ceba65993bad9d9150195e269fb6b146ca69ce8afc7.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Calendar()(*i2061516be4b73f307ef27998a5436c0d6263ff0f2dfe896d758d303753de3249.CalendarRequestBuilder) {
    return i2061516be4b73f307ef27998a5436c0d6263ff0f2dfe896d758d303753de3249.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *EventItemRequestBuilder) Cancel()(*ibdcbb5181e62a6c318e7e080ec9929125febe7b6feb6ff10777c7a7626f2ccf2.CancelRequestBuilder) {
    return ibdcbb5181e62a6c318e7e080ec9929125febe7b6feb6ff10777c7a7626f2ccf2.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*ica549a59bb07a4b59d1250ec60c1c7f20cc7b2c95841d2af82d8b77c130ebf03.DeclineRequestBuilder) {
    return ica549a59bb07a4b59d1250ec60c1c7f20cc7b2c95841d2af82d8b77c130ebf03.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *EventItemRequestBuilder) DismissReminder()(*i04547fc25f9a19b56375a31791a646bb9a7291300dab821083dcd9dff492eec5.DismissReminderRequestBuilder) {
    return i04547fc25f9a19b56375a31791a646bb9a7291300dab821083dcd9dff492eec5.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) Extensions()(*if9d4d521baacac2476f774dd755d2919846176c52baa2d6a33807091c4f57372.ExtensionsRequestBuilder) {
    return if9d4d521baacac2476f774dd755d2919846176c52baa2d6a33807091c4f57372.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i9b67f74d1e1e9cf957ac05505c98e12f20babaf600449186355ccc1ef198c643.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i9b67f74d1e1e9cf957ac05505c98e12f20babaf600449186355ccc1ef198c643.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *EventItemRequestBuilder) Forward()(*i0cc9b39109a903e3c7493bc95907d6b99e23de8914847bebc5b2b7718aa48e85.ForwardRequestBuilder) {
    return i0cc9b39109a903e3c7493bc95907d6b99e23de8914847bebc5b2b7718aa48e85.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*ic8e9ae378ae88a52d0b547ce11c9cff604379b3f87d4a0da0c87d617167b5c54.MultiValueExtendedPropertiesRequestBuilder) {
    return ic8e9ae378ae88a52d0b547ce11c9cff604379b3f87d4a0da0c87d617167b5c54.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib7178ba408e138ce3cdd93a42a6b5a070b0b033794f27a9ebb32ed3651994b60.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return ib7178ba408e138ce3cdd93a42a6b5a070b0b033794f27a9ebb32ed3651994b60.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i3c0518909ebf9cd483b89bfb6b30484ba04b2be163059af0f1fb66f6749d7f13.SingleValueExtendedPropertiesRequestBuilder) {
    return i3c0518909ebf9cd483b89bfb6b30484ba04b2be163059af0f1fb66f6749d7f13.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id6bbec9cfa87470cc5313b0414204e4191e208f07a93b995e7462d3584214ebb.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return id6bbec9cfa87470cc5313b0414204e4191e208f07a93b995e7462d3584214ebb.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *EventItemRequestBuilder) SnoozeReminder()(*iaaf46f7a2ead40b95e11dc338f68f4cc605581625ff09115cf3dfa212c70c55d.SnoozeReminderRequestBuilder) {
    return iaaf46f7a2ead40b95e11dc338f68f4cc605581625ff09115cf3dfa212c70c55d.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *EventItemRequestBuilder) TentativelyAccept()(*ibbcd29199f49a1e2c6067cb3d21b38496815d0537ce16e70bbffe8a002c63ac3.TentativelyAcceptRequestBuilder) {
    return ibbcd29199f49a1e2c6067cb3d21b38496815d0537ce16e70bbffe8a002c63ac3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
