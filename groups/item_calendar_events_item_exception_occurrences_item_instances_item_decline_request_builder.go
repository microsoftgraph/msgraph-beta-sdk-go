package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder provides operations to call the decline method.
type ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderInternal instantiates a new DeclineRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder) {
    m := &ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/calendar/events/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/decline", pathParameters),
    }
    return m
}
// NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder instantiates a new DeclineRequestBuilder and sets the default values.
func NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action decline
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder) Post(ctx context.Context, body ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclinePostRequestBodyable, requestConfiguration *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action decline
func (m *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclinePostRequestBodyable, requestConfiguration *ItemCalendarEventsItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
