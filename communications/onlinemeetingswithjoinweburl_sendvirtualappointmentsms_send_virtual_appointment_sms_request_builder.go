package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder provides operations to call the sendVirtualAppointmentSms method.
type OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilderInternal instantiates a new OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder and sets the default values.
func NewOnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder) {
    m := &OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/onlineMeetings(joinWebUrl='{joinWebUrl}')/sendVirtualAppointmentSms", pathParameters),
    }
    return m
}
// NewOnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder instantiates a new OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder and sets the default values.
func NewOnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send an SMS notification to external attendees when a Teams Virtual Appointment is confirmed, rescheduled, or canceled. This feature requires Teams Premium. Attendees must have a valid United States phone number to receive these SMS notifications.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-sendvirtualappointmentsms?view=graph-rest-beta
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder) Post(ctx context.Context, body OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsPostRequestBodyable, requestConfiguration *OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation send an SMS notification to external attendees when a Teams Virtual Appointment is confirmed, rescheduled, or canceled. This feature requires Teams Premium. Attendees must have a valid United States phone number to receive these SMS notifications.
// returns a *RequestInformation when successful
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder) ToPostRequestInformation(ctx context.Context, body OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsPostRequestBodyable, requestConfiguration *OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder when successful
func (m *OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder) WithUrl(rawUrl string)(*OnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder) {
    return NewOnlinemeetingswithjoinweburlSendvirtualappointmentsmsSendVirtualAppointmentSmsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
