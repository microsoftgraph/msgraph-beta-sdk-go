package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder provides operations to call the decline method.
type ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderInternal instantiates a new ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder and sets the default values.
func NewItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder) {
    m := &ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}/decline", pathParameters),
    }
    return m
}
// NewItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder instantiates a new ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder and sets the default values.
func NewItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderInternal(urlParams, requestAdapter)
}
// Post decline invitation to the specified event in a user calendar. If the event allows proposals for new times, on declining the event, an invitee can choose to suggest an alternative time by including the proposedNewTime parameter. For more information on how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-decline?view=graph-rest-beta
func (m *ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder) Post(ctx context.Context, body ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclinePostRequestBodyable, requestConfiguration *ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation decline invitation to the specified event in a user calendar. If the event allows proposals for new times, on declining the event, an invitee can choose to suggest an alternative time by including the proposedNewTime parameter. For more information on how to propose a time, and how to receive and accept a new time proposal, see Propose new meeting times.
// returns a *RequestInformation when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclinePostRequestBodyable, requestConfiguration *ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder when successful
func (m *ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder) {
    return NewItemCalendarEventsItemInstancesItemExceptionoccurrencesItemDeclineRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
