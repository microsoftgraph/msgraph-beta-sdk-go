package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder provides operations to count the resources in the collection.
type CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetQueryParameters get the number of the resource
type CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
}
// CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetQueryParameters
}
// NewCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder) {
    m := &CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/$count{?%24filter}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
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
func (m *CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CalendarViewItemExceptionOccurrencesItemInstancesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
