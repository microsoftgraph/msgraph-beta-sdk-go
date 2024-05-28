package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Accept()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Attachments()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Calendar()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Cancel()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal instantiates a new ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    m := &ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder instantiates a new ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Decline()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) DismissReminder()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Extensions()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Forward()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    return NewItemCalendarsItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
