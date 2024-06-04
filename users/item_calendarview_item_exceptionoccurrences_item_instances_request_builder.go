package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string `uriparametername:"endDateTime"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string `uriparametername:"startDateTime"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilderGetQueryParameters
}
// ByEventId2 provides operations to manage the instances property of the microsoft.graph.event entity.
// returns a *ItemCalendarviewItemExceptionoccurrencesItemInstancesEventItemRequestBuilder when successful
func (m *ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder) ByEventId2(eventId2 string)(*ItemCalendarviewItemExceptionoccurrencesItemInstancesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if eventId2 != "" {
        urlTplParams["event%2Did2"] = eventId2
    }
    return NewItemCalendarviewItemExceptionoccurrencesItemInstancesEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilderInternal instantiates a new ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder and sets the default values.
func NewItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder) {
    m := &ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances?endDateTime={endDateTime}&startDateTime={startDateTime}{&%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder instantiates a new ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder and sets the default values.
func NewItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemCalendarviewItemExceptionoccurrencesItemInstancesCountRequestBuilder when successful
func (m *ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder) Count()(*ItemCalendarviewItemExceptionoccurrencesItemInstancesCountRequestBuilder) {
    return NewItemCalendarviewItemExceptionoccurrencesItemInstancesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delta provides operations to call the delta method.
// returns a *ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaRequestBuilder when successful
func (m *ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder) Delta()(*ItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaRequestBuilder) {
    return NewItemCalendarviewItemExceptionoccurrencesItemInstancesDeltaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a EventCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EventCollectionResponseable, error) {
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
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but doesn't include occurrences that have been canceled from the series. Navigation property. Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder when successful
func (m *ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder) {
    return NewItemCalendarviewItemExceptionoccurrencesItemInstancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
