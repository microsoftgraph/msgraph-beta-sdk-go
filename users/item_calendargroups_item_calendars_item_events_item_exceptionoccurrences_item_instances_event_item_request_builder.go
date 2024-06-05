package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
type ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAcceptRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCalendarRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCancelRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderInternal instantiates a new ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder instantiates a new ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDeclineRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemForwardRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder when successful
func (m *ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) {
    return NewItemCalendargroupsItemCalendarsItemEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
