package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder provides operations to call the sendVirtualAppointmentReminderSms method.
type ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderInternal instantiates a new ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) {
    m := &ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/sendVirtualAppointmentReminderSms", pathParameters),
    }
    return m
}
// NewItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder instantiates a new ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder and sets the default values.
func NewItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send an SMS reminder to external attendees for a Teams Virtual Appointment. This feature requires Teams Premium and attendees must have a valid United States phone number to receive SMS notifications.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-sendvirtualappointmentremindersms?view=graph-rest-beta
func (m *ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) Post(ctx context.Context, body ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBodyable, requestConfiguration *ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send an SMS reminder to external attendees for a Teams Virtual Appointment. This feature requires Teams Premium and attendees must have a valid United States phone number to receive SMS notifications.
// returns a *RequestInformation when successful
func (m *ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsPostRequestBodyable, requestConfiguration *ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder when successful
func (m *ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) WithUrl(rawUrl string)(*ItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder) {
    return NewItemOnlinemeetingsItemSendvirtualappointmentremindersmsSendVirtualAppointmentReminderSmsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
