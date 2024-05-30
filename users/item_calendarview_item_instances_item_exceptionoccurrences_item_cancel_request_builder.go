package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder provides operations to call the cancel method.
type ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderInternal instantiates a new ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder and sets the default values.
func NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder) {
    m := &ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}/instances/{event%2Did1}/exceptionOccurrences/{event%2Did2}/cancel", pathParameters),
    }
    return m
}
// NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder instantiates a new ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder and sets the default values.
func NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this action allows the organizer of a meeting to send a cancellation message and cancel the event.  The action moves the event to the Deleted Items folder. The organizer can also cancel an occurrence of a recurring meeting by providing the occurrence event ID. An attendee calling this action gets an error (HTTP 400 Bad Request), with the followingerror message: 'Your request can't be completed. You need to be an organizer to cancel a meeting.' This action differs from Delete in that Cancel is available to only the organizer, and letsthe organizer send a custom message to the attendees about the cancellation.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-cancel?view=graph-rest-beta
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder) Post(ctx context.Context, body ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelPostRequestBodyable, requestConfiguration *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation this action allows the organizer of a meeting to send a cancellation message and cancel the event.  The action moves the event to the Deleted Items folder. The organizer can also cancel an occurrence of a recurring meeting by providing the occurrence event ID. An attendee calling this action gets an error (HTTP 400 Bad Request), with the followingerror message: 'Your request can't be completed. You need to be an organizer to cancel a meeting.' This action differs from Delete in that Cancel is available to only the organizer, and letsthe organizer send a custom message to the attendees about the cancellation.
// returns a *RequestInformation when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelPostRequestBodyable, requestConfiguration *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder when successful
func (m *ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder) {
    return NewItemCalendarviewItemInstancesItemExceptionoccurrencesItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
