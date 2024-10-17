package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*ItemCalendarEventsItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*ItemCalendarEventsItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemCalendarRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*ItemCalendarEventsItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*ItemCalendarEventsItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendar/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder instantiates a new ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*ItemCalendarEventsItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemDismissReminderRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*ItemCalendarEventsItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*ItemCalendarEventsItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*ItemCalendarEventsItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// Instances provides operations to manage the instances property of the microsoft.graph.event entity.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemInstancesRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) Instances()(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemPermanentDeleteRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) PermanentDelete()(*ItemCalendarEventsItemExceptionOccurrencesItemPermanentDeleteRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
