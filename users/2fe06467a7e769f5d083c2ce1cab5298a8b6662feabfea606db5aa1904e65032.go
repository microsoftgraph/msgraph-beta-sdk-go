package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder provides operations to call the permanentDelete method.
type ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilderInternal instantiates a new ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder) {
    m := &ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/permanentDelete", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder instantiates a new ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post permanently delete an event and place it in the Purges folder in the dumpster in the user's mailbox. Email clients such as Outlook or the Outlook on the web can't access permanently deleted items. Unless there's a hold set on the mailbox, the items are permanently deleted after a set period of time. For more information about item retention, see Configure Deleted Item retention and Recoverable Items quotas.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/event-permanentdelete?view=graph-rest-beta
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation permanently delete an event and place it in the Purges folder in the dumpster in the user's mailbox. Email clients such as Outlook or the Outlook on the web can't access permanently deleted items. Unless there's a hold set on the mailbox, the items are permanently deleted after a set period of time. For more information about item retention, see Configure Deleted Item retention and Recoverable Items quotas.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder when successful
func (m *ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder) WithUrl(rawUrl string)(*ItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesItemInstancesItemPermanentDeleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
