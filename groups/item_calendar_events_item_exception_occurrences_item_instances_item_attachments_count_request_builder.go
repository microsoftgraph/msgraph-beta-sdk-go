package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder provides operations to count the resources in the collection.
type ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderGetQueryParameters get the number of the resource
type ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
}
// ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderGetQueryParameters
}
// NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderInternal instantiates a new ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder) {
    m := &ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/attachments/$count{?%24filter}", pathParameters),
    }
    return m
}
// NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder instantiates a new ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder when successful
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder) {
    return NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemAttachmentsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
