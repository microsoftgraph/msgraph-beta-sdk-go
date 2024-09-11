package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
type ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemAcceptRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Accept()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemAcceptRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Attachments()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Calendar()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Cancel()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal instantiates a new ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    m := &ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder instantiates a new ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Decline()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemDismissReminderRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) DismissReminder()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemDismissReminderRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Extensions()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Forward()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from groups
// returns a Eventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
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
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) Instances()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PermanentDelete provides operations to call the permanentDelete method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemPermanentDeleteRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) PermanentDelete()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemPermanentDeleteRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemPermanentDeleteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarCalendarViewItemExceptionOccurrencesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get exceptionOccurrences from groups
// returns a *RequestInformation when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesEventItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
