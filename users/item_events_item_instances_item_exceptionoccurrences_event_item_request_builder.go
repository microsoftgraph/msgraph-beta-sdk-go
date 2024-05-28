package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Accept()(*ItemEventsItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Attachments()(*ItemEventsItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Calendar()(*ItemEventsItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Cancel()(*ItemEventsItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal instantiates a new ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder and sets the default values.
func NewItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    m := &ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder instantiates a new ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder and sets the default values.
func NewItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Decline()(*ItemEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) DismissReminder()(*ItemEventsItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Extensions()(*ItemEventsItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesItemForwardRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Forward()(*ItemEventsItemInstancesItemExceptionoccurrencesItemForwardRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemEventsItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemEventsItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder when successful
func (m *ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionoccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
