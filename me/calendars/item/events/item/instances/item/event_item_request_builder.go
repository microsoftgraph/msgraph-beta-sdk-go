package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i048cae98afed2bcfc52124c5a73b9724614574aa8636518cd4d1321fda7bfc07 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/tentativelyaccept"
    i09bbc4cb000c31d7d0dcc2d8d2a19d825cc947d91157459d7037271c58d21cf2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/multivalueextendedproperties"
    i3d9b649c8ca7312d3ab05eed2b11021d247349a0d6684c22e352a75aa4b85780 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/extensions"
    i628be3bc640d9f7ad97770af1fb243685ad0495937c654c8bf3b409bbe2206fa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/decline"
    i6bfd098884ab420a6c8080f2ea836f7a50f00fb47a739fcf71f3da14fad8daa6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/calendar"
    i6d79d0caec33d046d9854c5df0a59a36dd43c13a8931767a705d1709b177c1fd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/dismissreminder"
    i8e37990dab8ad7d081bb17c5a28ae24aab95bd004c238f26afbbc7ba02f88102 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/singlevalueextendedproperties"
    iccc31c4251a5fd741c25404a7d87955e7597506a851e61a35c3e472804baf2db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/cancel"
    id83e39dd29a8e3c2e9469f256dc18b44ff99d8b8ff288cabe59845d7f6f5f434 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/snoozereminder"
    id9f64ebcc1f739b750d6541599359d7f296e285ad3ee55ce90c2e2cf891d96d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/forward"
    ie914662dfee924e7297194e72446e1746c92203c8f0dea99255cdbdc185d3879 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/attachments"
    if11bd22064e35e4b39c13be59c77670bcc446ce0c6fbf2201c1d2c1246f4b44f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/accept"
    if7c0420b86216b7cf8e153fed9f0be790a80daf529bca27757ba6a5276a1762c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences"
    i1b4e4991f3d8bd757f3f6adaba892d558311427a930d1d93a3e84446a9c57fc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/singlevalueextendedproperties/item"
    i207bb0827311b7ceeaa1ab3b1ce079beb84fe98b8884787b8d0d2d54a5e7ee25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/multivalueextendedproperties/item"
    i94f4d1e66dfeaf32c3c9b4151d666ac69c7f00e8fa0a68c672c447c6009326c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/attachments/item"
    i9816b27bba149eee28e30c1ef89d862feb5f58d24fc827c4bcbde62ee5999e7c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/extensions/item"
    id22bc80b512b5084443604a788667020ca343bef03de2bf08262a8a317d05750 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/events/item/instances/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*if11bd22064e35e4b39c13be59c77670bcc446ce0c6fbf2201c1d2c1246f4b44f.AcceptRequestBuilder) {
    return if11bd22064e35e4b39c13be59c77670bcc446ce0c6fbf2201c1d2c1246f4b44f.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ie914662dfee924e7297194e72446e1746c92203c8f0dea99255cdbdc185d3879.AttachmentsRequestBuilder) {
    return ie914662dfee924e7297194e72446e1746c92203c8f0dea99255cdbdc185d3879.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i94f4d1e66dfeaf32c3c9b4151d666ac69c7f00e8fa0a68c672c447c6009326c0.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i94f4d1e66dfeaf32c3c9b4151d666ac69c7f00e8fa0a68c672c447c6009326c0.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*i6bfd098884ab420a6c8080f2ea836f7a50f00fb47a739fcf71f3da14fad8daa6.CalendarRequestBuilder) {
    return i6bfd098884ab420a6c8080f2ea836f7a50f00fb47a739fcf71f3da14fad8daa6.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*iccc31c4251a5fd741c25404a7d87955e7597506a851e61a35c3e472804baf2db.CancelRequestBuilder) {
    return iccc31c4251a5fd741c25404a7d87955e7597506a851e61a35c3e472804baf2db.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendars/{calendar%2Did}/events/{event%2Did}/instances/{event%2Did1}{?%24select}";
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
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i628be3bc640d9f7ad97770af1fb243685ad0495937c654c8bf3b409bbe2206fa.DeclineRequestBuilder) {
    return i628be3bc640d9f7ad97770af1fb243685ad0495937c654c8bf3b409bbe2206fa.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i6d79d0caec33d046d9854c5df0a59a36dd43c13a8931767a705d1709b177c1fd.DismissReminderRequestBuilder) {
    return i6d79d0caec33d046d9854c5df0a59a36dd43c13a8931767a705d1709b177c1fd.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrences the exceptionOccurrences property
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*if7c0420b86216b7cf8e153fed9f0be790a80daf529bca27757ba6a5276a1762c.ExceptionOccurrencesRequestBuilder) {
    return if7c0420b86216b7cf8e153fed9f0be790a80daf529bca27757ba6a5276a1762c.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*id22bc80b512b5084443604a788667020ca343bef03de2bf08262a8a317d05750.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return id22bc80b512b5084443604a788667020ca343bef03de2bf08262a8a317d05750.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i3d9b649c8ca7312d3ab05eed2b11021d247349a0d6684c22e352a75aa4b85780.ExtensionsRequestBuilder) {
    return i3d9b649c8ca7312d3ab05eed2b11021d247349a0d6684c22e352a75aa4b85780.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i9816b27bba149eee28e30c1ef89d862feb5f58d24fc827c4bcbde62ee5999e7c.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i9816b27bba149eee28e30c1ef89d862feb5f58d24fc827c4bcbde62ee5999e7c.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*id9f64ebcc1f739b750d6541599359d7f296e285ad3ee55ce90c2e2cf891d96d6.ForwardRequestBuilder) {
    return id9f64ebcc1f739b750d6541599359d7f296e285ad3ee55ce90c2e2cf891d96d6.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i09bbc4cb000c31d7d0dcc2d8d2a19d825cc947d91157459d7037271c58d21cf2.MultiValueExtendedPropertiesRequestBuilder) {
    return i09bbc4cb000c31d7d0dcc2d8d2a19d825cc947d91157459d7037271c58d21cf2.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i207bb0827311b7ceeaa1ab3b1ce079beb84fe98b8884787b8d0d2d54a5e7ee25.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i207bb0827311b7ceeaa1ab3b1ce079beb84fe98b8884787b8d0d2d54a5e7ee25.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i8e37990dab8ad7d081bb17c5a28ae24aab95bd004c238f26afbbc7ba02f88102.SingleValueExtendedPropertiesRequestBuilder) {
    return i8e37990dab8ad7d081bb17c5a28ae24aab95bd004c238f26afbbc7ba02f88102.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendars.item.events.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i1b4e4991f3d8bd757f3f6adaba892d558311427a930d1d93a3e84446a9c57fc5.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i1b4e4991f3d8bd757f3f6adaba892d558311427a930d1d93a3e84446a9c57fc5.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*id83e39dd29a8e3c2e9469f256dc18b44ff99d8b8ff288cabe59845d7f6f5f434.SnoozeReminderRequestBuilder) {
    return id83e39dd29a8e3c2e9469f256dc18b44ff99d8b8ff288cabe59845d7f6f5f434.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i048cae98afed2bcfc52124c5a73b9724614574aa8636518cd4d1321fda7bfc07.TentativelyAcceptRequestBuilder) {
    return i048cae98afed2bcfc52124c5a73b9724614574aa8636518cd4d1321fda7bfc07.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
