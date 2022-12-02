package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilder provides operations to call the dismissReminder method.
type MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilderInternal instantiates a new DismissReminderRequestBuilder and sets the default values.
func NewMeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilder) {
    m := &MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendar/calendarView/{event%2Did}/instances/{event%2Did1}/microsoft.graph.dismissReminder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilder instantiates a new DismissReminderRequestBuilder and sets the default values.
func NewMeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation dismiss a reminder that has been triggered for an event in a user calendar.
func (m *MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post dismiss a reminder that has been triggered for an event in a user calendar.
func (m *MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilder) Post(ctx context.Context, requestConfiguration *MeCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
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
