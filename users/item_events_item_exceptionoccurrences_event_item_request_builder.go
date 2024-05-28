package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemEventsItemExceptionoccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemEventsItemExceptionoccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEventsItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemEventsItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemEventsItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEventsItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemEventsItemExceptionoccurrencesItemAcceptRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) Accept()(*ItemEventsItemExceptionoccurrencesItemAcceptRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemEventsItemExceptionoccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) Attachments()(*ItemEventsItemExceptionoccurrencesItemAttachmentsRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) Calendar()(*ItemEventsItemExceptionoccurrencesItemCalendarRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemEventsItemExceptionoccurrencesItemCancelRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) Cancel()(*ItemEventsItemExceptionoccurrencesItemCancelRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemEventsItemExceptionoccurrencesEventItemRequestBuilderInternal instantiates a new ItemEventsItemExceptionoccurrencesEventItemRequestBuilder and sets the default values.
func NewItemEventsItemExceptionoccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) {
    m := &ItemEventsItemExceptionoccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/events/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemEventsItemExceptionoccurrencesEventItemRequestBuilder instantiates a new ItemEventsItemExceptionoccurrencesEventItemRequestBuilder and sets the default values.
func NewItemEventsItemExceptionoccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsItemExceptionoccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) Decline()(*ItemEventsItemExceptionoccurrencesItemDeclineRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemEventsItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) DismissReminder()(*ItemEventsItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemEventsItemExceptionoccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) Extensions()(*ItemEventsItemExceptionoccurrencesItemExtensionsRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemEventsItemExceptionoccurrencesItemForwardRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) Forward()(*ItemEventsItemExceptionoccurrencesItemForwardRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEventsItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemEventsItemExceptionoccurrencesItemInstancesRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) Instances()(*ItemEventsItemExceptionoccurrencesItemInstancesRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemEventsItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemEventsItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemEventsItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemEventsItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEventsItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder when successful
func (m *ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemEventsItemExceptionoccurrencesEventItemRequestBuilder) {
    return NewItemEventsItemExceptionoccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
