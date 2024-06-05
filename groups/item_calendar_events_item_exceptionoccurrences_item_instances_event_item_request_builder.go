package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemAcceptRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemCalendarRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemCancelRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderInternal instantiates a new ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder instantiates a new ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemDeclineRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemForwardRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) {
    return NewItemCalendarEventsItemExceptionoccurrencesItemInstancesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
