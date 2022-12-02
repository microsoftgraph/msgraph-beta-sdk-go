package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder provides operations to call the cancel method.
type MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilderInternal instantiates a new CancelRequestBuilder and sets the default values.
func NewMeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder) {
    m := &MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarGroups/{calendarGroup%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/microsoft.graph.cancel";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder instantiates a new CancelRequestBuilder and sets the default values.
func NewMeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation this action allows the organizer of a meeting to send a cancellation message and cancel the event.  The action moves the event to the Deleted Items folder. The organizer can also cancel an occurrence of a recurring meeting by providing the occurrence event ID. An attendee calling this action gets an error (HTTP 400 Bad Request), with the followingerror message: 'Your request can't be completed. You need to be an organizer to cancel a meeting.' This action differs from Delete in that **Cancel** is available to only the organizer, and letsthe organizer send a custom message to the attendees about the cancellation.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelPostRequestBodyable, requestConfiguration *MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post this action allows the organizer of a meeting to send a cancellation message and cancel the event.  The action moves the event to the Deleted Items folder. The organizer can also cancel an occurrence of a recurring meeting by providing the occurrence event ID. An attendee calling this action gets an error (HTTP 400 Bad Request), with the followingerror message: 'Your request can't be completed. You need to be an organizer to cancel a meeting.' This action differs from Delete in that **Cancel** is available to only the organizer, and letsthe organizer send a custom message to the attendees about the cancellation.
func (m *MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelPostRequestBodyable, requestConfiguration *MeCalendarGroupsItemCalendarsItemCalendarViewItemCancelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
