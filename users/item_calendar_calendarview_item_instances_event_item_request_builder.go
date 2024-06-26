package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarviewItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarCalendarviewItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
}
// ItemCalendarCalendarviewItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarviewItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarCalendarviewItemInstancesItemAcceptRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendarCalendarviewItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemInstancesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarCalendarviewItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemInstancesItemCalendarRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarCalendarviewItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarCalendarviewItemInstancesItemCancelRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendarCalendarviewItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarCalendarviewItemInstancesEventItemRequestBuilderInternal instantiates a new ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarviewItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarviewItemInstancesEventItemRequestBuilder instantiates a new ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarviewItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarviewItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarCalendarviewItemInstancesItemDeclineRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendarCalendarviewItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendarCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesExceptionOccurrencesRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) ExceptionOccurrences()(*ItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesExceptionOccurrencesRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExceptionoccurrencesExceptionOccurrencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarviewItemInstancesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarCalendarviewItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarCalendarviewItemInstancesItemForwardRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendarCalendarviewItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarCalendarviewItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarCalendarviewItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder when successful
func (m *ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarCalendarviewItemInstancesEventItemRequestBuilder) {
    return NewItemCalendarCalendarviewItemInstancesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
