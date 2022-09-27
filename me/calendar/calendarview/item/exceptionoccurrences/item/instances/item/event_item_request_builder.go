package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i181c2c9eac293439f8cf3a1f5170a0aed081a554ce8a84da9916cd5d5ec27e0e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i1d1d1f3e99bc69739a9ac47d2d637140717d86e1edba01560c181ca0eaadea7e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/cancel"
    i3363382a5f05954f78675568b4812fe0d91175c6be6dca062cd983f084d2bbc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i381efd2a9df0e50fde81059df3648bf45c2910fa95dd4cd99d23092675e0938e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    i585ee49a0b10bc7c3fe8cdbae364c4ff3b6dd315f6a3d61d8d146dab2b81caf9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/attachments"
    i5e6ca5e447c6b2fb5c8b8e9b7e305108ceaaaf9c770fe3ae4091912f77aace8d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/decline"
    i6c888ed2d05f7ce3273a76c590f8bd61cd15582ea0907921d983e3be49587634 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/forward"
    iac50d23b30ad6247a74379f7e24926b880b5a82f30e99f0fec9ea5cecf4f8c08 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/accept"
    iba37c222f898cf7c4cb964a04beaaaa200f647067de269d1da0bc71780b6c2c5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/extensions"
    id78983fc9b9eadc34bb4b1bc264f0570592fd423e7b1bfb9a19b9666e3e097ac "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    if58d4a4f0cf6748201e69067f741372ae66d7fa091d7a308d8713421a246723e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/dismissreminder"
    ifeee5428ad3ab30094b24f38a1e43e88856b35b5e06acacc8c0f31523448c73b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/calendar"
    i5dfe18ab162f1e564bc1c5996483362efaa220ce142beca4d11da0db951fa80e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
    i7b2ccd2d9199f6d362cef63c7825e38f27e1744e27dd91af2e83b4fa1f88d7f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    ia76ccb8069b513a589407f7654de23dfe49509cc6d0693dc21640a87a00efc00 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/extensions/item"
    icde4da04efc347562262fe157c6720d92f9a57754d263ac1be4ad87c784bde6a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendar/calendarview/item/exceptionoccurrences/item/instances/item/attachments/item"
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
func (m *EventItemRequestBuilder) Accept()(*iac50d23b30ad6247a74379f7e24926b880b5a82f30e99f0fec9ea5cecf4f8c08.AcceptRequestBuilder) {
    return iac50d23b30ad6247a74379f7e24926b880b5a82f30e99f0fec9ea5cecf4f8c08.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*i585ee49a0b10bc7c3fe8cdbae364c4ff3b6dd315f6a3d61d8d146dab2b81caf9.AttachmentsRequestBuilder) {
    return i585ee49a0b10bc7c3fe8cdbae364c4ff3b6dd315f6a3d61d8d146dab2b81caf9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*icde4da04efc347562262fe157c6720d92f9a57754d263ac1be4ad87c784bde6a.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return icde4da04efc347562262fe157c6720d92f9a57754d263ac1be4ad87c784bde6a.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*ifeee5428ad3ab30094b24f38a1e43e88856b35b5e06acacc8c0f31523448c73b.CalendarRequestBuilder) {
    return ifeee5428ad3ab30094b24f38a1e43e88856b35b5e06acacc8c0f31523448c73b.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i1d1d1f3e99bc69739a9ac47d2d637140717d86e1edba01560c181ca0eaadea7e.CancelRequestBuilder) {
    return i1d1d1f3e99bc69739a9ac47d2d637140717d86e1edba01560c181ca0eaadea7e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
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
func (m *EventItemRequestBuilder) Decline()(*i5e6ca5e447c6b2fb5c8b8e9b7e305108ceaaaf9c770fe3ae4091912f77aace8d.DeclineRequestBuilder) {
    return i5e6ca5e447c6b2fb5c8b8e9b7e305108ceaaaf9c770fe3ae4091912f77aace8d.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*if58d4a4f0cf6748201e69067f741372ae66d7fa091d7a308d8713421a246723e.DismissReminderRequestBuilder) {
    return if58d4a4f0cf6748201e69067f741372ae66d7fa091d7a308d8713421a246723e.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*iba37c222f898cf7c4cb964a04beaaaa200f647067de269d1da0bc71780b6c2c5.ExtensionsRequestBuilder) {
    return iba37c222f898cf7c4cb964a04beaaaa200f647067de269d1da0bc71780b6c2c5.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ia76ccb8069b513a589407f7654de23dfe49509cc6d0693dc21640a87a00efc00.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ia76ccb8069b513a589407f7654de23dfe49509cc6d0693dc21640a87a00efc00.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i6c888ed2d05f7ce3273a76c590f8bd61cd15582ea0907921d983e3be49587634.ForwardRequestBuilder) {
    return i6c888ed2d05f7ce3273a76c590f8bd61cd15582ea0907921d983e3be49587634.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i181c2c9eac293439f8cf3a1f5170a0aed081a554ce8a84da9916cd5d5ec27e0e.MultiValueExtendedPropertiesRequestBuilder) {
    return i181c2c9eac293439f8cf3a1f5170a0aed081a554ce8a84da9916cd5d5ec27e0e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5dfe18ab162f1e564bc1c5996483362efaa220ce142beca4d11da0db951fa80e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i5dfe18ab162f1e564bc1c5996483362efaa220ce142beca4d11da0db951fa80e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i381efd2a9df0e50fde81059df3648bf45c2910fa95dd4cd99d23092675e0938e.SingleValueExtendedPropertiesRequestBuilder) {
    return i381efd2a9df0e50fde81059df3648bf45c2910fa95dd4cd99d23092675e0938e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendar.calendarView.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7b2ccd2d9199f6d362cef63c7825e38f27e1744e27dd91af2e83b4fa1f88d7f3.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i7b2ccd2d9199f6d362cef63c7825e38f27e1744e27dd91af2e83b4fa1f88d7f3.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i3363382a5f05954f78675568b4812fe0d91175c6be6dca062cd983f084d2bbc5.SnoozeReminderRequestBuilder) {
    return i3363382a5f05954f78675568b4812fe0d91175c6be6dca062cd983f084d2bbc5.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*id78983fc9b9eadc34bb4b1bc264f0570592fd423e7b1bfb9a19b9666e3e097ac.TentativelyAcceptRequestBuilder) {
    return id78983fc9b9eadc34bb4b1bc264f0570592fd423e7b1bfb9a19b9666e3e097ac.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
