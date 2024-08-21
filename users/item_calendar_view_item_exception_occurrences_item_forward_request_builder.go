package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder provides operations to call the forward method.
type ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilderInternal instantiates a new ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder) {
    m := &ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/forward", pathParameters),
    }
    return m
}
// NewItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder instantiates a new ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilderInternal(urlParams, requestAdapter)
}
// Post this action allows the organizer or attendee of a meeting event to forward themeeting request to a new recipient. If the meeting event is forwarded from an attendee's Microsoft 365 mailbox to another recipient, this actionalso sends a message to notify the organizer of the forwarding, and adds the recipient to the organizer'scopy of the meeting event. This convenience is not available when forwarding from an Outlook.com account.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-forward?view=graph-rest-beta
func (m *ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder) Post(ctx context.Context, body ItemCalendarViewItemExceptionOccurrencesItemForwardPostRequestBodyable, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation this action allows the organizer or attendee of a meeting event to forward themeeting request to a new recipient. If the meeting event is forwarded from an attendee's Microsoft 365 mailbox to another recipient, this actionalso sends a message to notify the organizer of the forwarding, and adds the recipient to the organizer'scopy of the meeting event. This convenience is not available when forwarding from an Outlook.com account.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarViewItemExceptionOccurrencesItemForwardPostRequestBodyable, requestConfiguration *ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder when successful
func (m *ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder) {
    return NewItemCalendarViewItemExceptionOccurrencesItemForwardRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
