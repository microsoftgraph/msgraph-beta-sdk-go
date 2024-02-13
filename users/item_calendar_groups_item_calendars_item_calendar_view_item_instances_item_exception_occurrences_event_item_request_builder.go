package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder instantiates a new ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
