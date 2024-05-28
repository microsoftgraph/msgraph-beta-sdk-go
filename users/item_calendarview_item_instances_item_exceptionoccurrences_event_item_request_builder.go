package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Accept()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Attachments()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Calendar()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Cancel()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal instantiates a new ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    m := &ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder instantiates a new ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Decline()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) DismissReminder()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemDismissreminderDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Extensions()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Forward()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemSnoozereminderSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemTentativelyacceptTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
