package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAcceptRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCalendarRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCancelRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderInternal instantiates a new ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder instantiates a new ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDeclineRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemForwardRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder when successful
func (m *ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) {
    return NewItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
