package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder provides operations to manage the calendar property of the microsoft.graph.event entity.
type ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderGetQueryParameters the calendar that contains the event. Navigation property. Read-only.
type ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderGetQueryParameters
}
// NewItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    m := &ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}/calendar{?%24select}", pathParameters),
    }
    return m
}
// NewItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder instantiates a new CalendarRequestBuilder and sets the default values.
func NewItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the calendar that contains the event. Navigation property. Read-only.
func (m *ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCalendarFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Calendarable), nil
}
// ToGetRequestInformation the calendar that contains the event. Navigation property. Read-only.
func (m *ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) WithUrl(rawUrl string)(*ItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder) {
    return NewItemEventsItemInstancesItemExceptionOccurrencesItemCalendarRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
