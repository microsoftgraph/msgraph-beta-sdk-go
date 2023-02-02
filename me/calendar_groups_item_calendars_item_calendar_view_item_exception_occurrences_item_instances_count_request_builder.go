package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder provides operations to count the resources in the collection.
type CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetQueryParameters get the number of the resource
type CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
}
// CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetQueryParameters
}
// NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder) {
    m := &CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/$count{?%24filter}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
