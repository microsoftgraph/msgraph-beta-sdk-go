package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarViewItemExceptionOccurrencesRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type ItemCalendarViewItemExceptionOccurrencesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarViewItemExceptionOccurrencesRequestBuilderGetQueryParameters get exceptionOccurrences from users
type ItemCalendarViewItemExceptionOccurrencesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemCalendarViewItemExceptionOccurrencesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarViewItemExceptionOccurrencesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarViewItemExceptionOccurrencesRequestBuilderGetQueryParameters
}
// ByEventId1 provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
// returns a *ItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder when successful
func (m *ItemCalendarViewItemExceptionOccurrencesRequestBuilder) ByEventId1(eventId1 string)(*ItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if eventId1 != "" {
        urlTplParams["event%2Did1"] = eventId1
    }
    return NewItemCalendarViewItemExceptionOccurrencesEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarViewItemExceptionOccurrencesRequestBuilderInternal instantiates a new ItemCalendarViewItemExceptionOccurrencesRequestBuilder and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemExceptionOccurrencesRequestBuilder) {
    m := &ItemCalendarViewItemExceptionOccurrencesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}/exceptionOccurrences{?%24count,%24expand,%24filter,%24orderby,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCalendarViewItemExceptionOccurrencesRequestBuilder instantiates a new ItemCalendarViewItemExceptionOccurrencesRequestBuilder and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemExceptionOccurrencesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarViewItemExceptionOccurrencesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemCalendarViewItemExceptionOccurrencesCountRequestBuilder when successful
func (m *ItemCalendarViewItemExceptionOccurrencesRequestBuilder) Count()(*ItemCalendarViewItemExceptionOccurrencesCountRequestBuilder) {
    return NewItemCalendarViewItemExceptionOccurrencesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
// returns a *ItemCalendarViewItemExceptionOccurrencesDeltaRequestBuilder when successful
func (m *ItemCalendarViewItemExceptionOccurrencesRequestBuilder) Delta()(*ItemCalendarViewItemExceptionOccurrencesDeltaRequestBuilder) {
    return NewItemCalendarViewItemExceptionOccurrencesDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get exceptionOccurrences from users
// returns a EventCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarViewItemExceptionOccurrencesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable), nil
}
// ToGetRequestInformation get exceptionOccurrences from users
// returns a *RequestInformation when successful
func (m *ItemCalendarViewItemExceptionOccurrencesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarViewItemExceptionOccurrencesRequestBuilder when successful
func (m *ItemCalendarViewItemExceptionOccurrencesRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarViewItemExceptionOccurrencesRequestBuilder) {
    return NewItemCalendarViewItemExceptionOccurrencesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
