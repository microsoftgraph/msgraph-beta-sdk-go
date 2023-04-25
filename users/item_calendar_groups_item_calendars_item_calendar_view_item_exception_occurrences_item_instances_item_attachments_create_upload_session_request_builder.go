package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilder provides operations to call the createUploadSession method.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilderInternal instantiates a new CreateUploadSessionRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilder) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}/attachments/createUploadSession", pathParameters),
    }
    return m
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilder instantiates a new CreateUploadSessionRequestBuilder and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createUploadSession
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilder) Post(ctx context.Context, body ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionPostRequestBodyable, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UploadSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUploadSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UploadSessionable), nil
}
// ToPostRequestInformation invoke action createUploadSession
func (m *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionPostRequestBodyable, requestConfiguration *ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsCreateUploadSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
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
